/*
 * Copyright: ⚛ <0xe2.0x9a.0x9b@gmail.com> 2010
 *
 * The contents of this file can be used freely,
 * except for usages in immoral contexts.
 */

#include "callback.h"
#include <stdio.h>
#include <pthread.h>
#include <assert.h>
#include <string.h>
#include <time.h>

#define TRUE  1
#define FALSE 0

static pthread_mutex_t m         = PTHREAD_MUTEX_INITIALIZER;
static pthread_cond_t  need      = PTHREAD_COND_INITIALIZER;
static pthread_cond_t  avail     = PTHREAD_COND_INITIALIZER;
static size_t          needed    = 0;	// Number of bytes needed by the consumer
static size_t          available = 0;	// Number of bytes available (from the producer)

static Uint8 *stream;	// Communication buffer between the consumer and the producer

static int64_t get_time() {
	struct timespec ts;
	if(clock_gettime(CLOCK_MONOTONIC, &ts) == 0)
		return 1000000000*(int64_t)ts.tv_sec + (int64_t)ts.tv_nsec;
	else
		return -1;
}

static uint64_t cummulativeLatency = 0;
static unsigned numCallbacks = 0;

static void SDLCALL callback(void *userdata, Uint8 *_stream, int _len) {
	assert(_len > 0);

	size_t len = (size_t)_len;

	pthread_mutex_lock(&m);
	{
		assert(available == 0);
		stream = _stream;

		{
			int64_t t1 = get_time();
			//printf("consumer: t1=%lld µs\n", (long long)t1/1000);

			assert(needed == 0);
			//printf("consumer: needed <- %zu\n", len);
			needed = len;
			pthread_cond_signal(&need);

			//printf("consumer: waiting for data\n");
			pthread_cond_wait(&avail, &m);
			assert(needed == 0);
			assert(available == len);

			int64_t t2 = get_time();
			//printf("consumer: t2=%lld µs\n", (long long)t2/1000);
			if(t1>0 && t2>0) {
				uint64_t latency = t2-t1;
				cummulativeLatency += latency;
				numCallbacks++;
				/*printf("consumer: latency=%lld µs, avg=%u µs\n",
				       (long long)(latency/1000),
				       (unsigned)(cummulativeLatency/numCallbacks/1000));*/
			}
		}

		//printf("consumer: received %zu bytes of data\n", available);

		//printf("consumer: available <- 0\n");
		available = 0;
		stream = NULL;
	}
	pthread_mutex_unlock(&m);
}

callback_t callback_getCallback() {
	return &callback;
}

void callback_fillBuffer(Uint8 *data, size_t numBytes) {
	size_t sent = 0;

	pthread_mutex_lock(&m);

	while(sent < numBytes) {
		//int64_t t = get_time();
		//printf("producer: t=%lld µs\n", (long long)t1/1000);

		if(needed == 0) {
			//printf("producer: waiting until data is needed (1)\n");
			pthread_cond_wait(&need, &m);
		}

		assert(stream != NULL);
		assert(needed > 0);

		// Append a chunk of data to the 'stream'
		size_t n = (needed<(numBytes-sent)) ? needed : (numBytes-sent);
		memcpy(stream+available, data+sent, n);
		available += n;
		sent += n;
		needed -= n;

		//printf("producer: added %zu bytes, available=%zu\n", n, available);

		if(needed == 0) {
			pthread_cond_signal(&avail);
			if(sent < numBytes) {
				//printf("producer: waiting until data is needed (2)\n");
				pthread_cond_wait(&need, &m);
			}
			else {
				break;
			}
		}
	}

	pthread_mutex_unlock(&m);
}

void callback_unblock() {
	pthread_mutex_lock(&m);
	if(needed > 0) {
		// Note: SDL already prefilled the entire 'stream' with silence
		assert(stream != NULL);
		available += needed;
		needed = 0;
		pthread_cond_signal(&avail);
	}
	pthread_mutex_unlock(&m);
}

