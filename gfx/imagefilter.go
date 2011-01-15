package gfx

// #include <SDL/SDL_imageFilter.h>
import "C"

type  u8p *C.uchar
type  s16p *C.short
	   
func ImageFilterMMXdetect() bool {
	if C.SDL_imageFilterMMXdetect() == C.int(0) {
		return false
	}
	return true
}

func ImageFilterMMXoff() { C.SDL_imageFilterMMXoff() }
func ImageFilterMMXon() { C.SDL_imageFilterMMXon() }

func ImageFilterAdd(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterAdd( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterMean(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterMean( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterSub(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterSub( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterAbsDiff(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterAbsDiff( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterMult(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterMult( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterMultNor(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterMultNor( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterMultDivby2(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterMultDivby2( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterMultDivby4(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterMultDivby4( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterBitAnd(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterBitAnd( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterBitOr(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterBitOr( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterDiv(src1, src2, dest *byte, length uint) int {
	return int( C.SDL_imageFilterDiv( u8p(ptr(src1)), u8p(ptr(src2)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterBitNegation(src, dest *byte, length uint) int {
	return int( C.SDL_imageFilterBitNegation( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length)) )
}

func ImageFilterAddByte(src, dest *byte, length uint, c byte) int {
	return int( C.SDL_imageFilterAddByte( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(c)) )
}

func ImageFilterAddUint(src, dest *byte, length uint, c uint) int {
	return int( C.SDL_imageFilterAddUint( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uint(c)) )
}

func ImageFilterAddByteToHalf(src, dest *byte, length uint, c byte) int {
	return int( C.SDL_imageFilterAddByteToHalf( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(c)) )
}

func ImageFilterSubByte(src, dest *byte, length uint, c byte) int {
	return int( C.SDL_imageFilterSubByte( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(c)) )
}

func ImageFilterSubUint(src, dest *byte, length uint, c uint) int {
	return int( C.SDL_imageFilterSubUint( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uint(c)) )
}

func ImageFilterShiftRight(src, dest *byte, length uint, n byte) int {
	return int( C.SDL_imageFilterShiftRight( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(n)) )
}

func ImageFilterShiftRightUint(src, dest *byte, length uint, n byte) int {
	return int( C.SDL_imageFilterShiftRightUint( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(n)) )
}

func ImageFilterMultByByte(src, dest *byte, length uint, n byte) int {
	return int( C.SDL_imageFilterMultByByte( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(n)) )
}

func ImageFilterShiftRightAndMultByByte(src, dest *byte, length uint, n, c byte) int {
	return int( C.SDL_imageFilterShiftRightAndMultByByte( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(n), C.uchar(c)) )
}

func ImageFilterShiftLeftByte(src, dest *byte, length uint, n byte) int {
	return int( C.SDL_imageFilterShiftLeftByte( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(n)) )
}

func ImageFilterShiftLeftUint(src, dest *byte, length uint, n byte) int {
	return int( C.SDL_imageFilterShiftLeftUint( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(n)) )
}

func ImageFilterShiftLeft(src, dest *byte, length uint, n byte) int {
	return int( C.SDL_imageFilterShiftLeft( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(n)) )
}


func ImageFilterBinarizeUsingThreshold(src, dest *byte, length uint, T byte) int {
	return int( C.SDL_imageFilterBinarizeUsingThreshold( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(T)) )
}

func ImageFilterClipToRange(src, dest *byte, length uint, Tmin, Tmax byte) int {
	return int( C.SDL_imageFilterClipToRange( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.uchar(Tmin), C.uchar(Tmax)) )
}

func ImageFilterNormalizeLinear(src, dest *byte, length uint, Cmin, Cmax, Nmin, Nmax int) int {
	return int( C.SDL_imageFilterNormalizeLinear( u8p(ptr(src)), u8p(ptr(dest)), C.uint(length), C.int(Cmin), C.int(Cmax), C.int(Nmin), C.int(Nmax)) )
}

func ImageFilterConvolveKernel3x3Divide(src, dest *byte, rows, columns int, Kernel *int16, Divisor byte) int {
	return int( C.SDL_imageFilterConvolveKernel3x3Divide( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns), s16p(ptr(Kernel)), C.uchar(Divisor)) )
}

func ImageFilterConvolveKernel5x5Divide(src, dest *byte, rows, columns int, Kernel *int16, Divisor byte) int {
	return int( C.SDL_imageFilterConvolveKernel5x5Divide( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns), s16p(ptr(Kernel)), C.uchar(Divisor)) )
}

func ImageFilterConvolveKernel7x7Divide(src, dest *byte, rows, columns int, Kernel *int16, Divisor byte) int {
	return int( C.SDL_imageFilterConvolveKernel7x7Divide( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns), s16p(ptr(Kernel)), C.uchar(Divisor)) )
}

func ImageFilterConvolveKernel9x9Divide(src, dest *byte, rows, columns int, Kernel *int16, Divisor byte) int {
	return int( C.SDL_imageFilterConvolveKernel9x9Divide( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns), s16p(ptr(Kernel)), C.uchar(Divisor)) )
}

func ImageFilterConvolveKernel3x3ShiftRight(src, dest *byte, rows, columns int, Kernel *int16, NRightShift byte) int {
	return int( C.SDL_imageFilterConvolveKernel3x3ShiftRight( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns), s16p(ptr(Kernel)), C.uchar(NRightShift)) )
}

func ImageFilterConvolveKernel5x5ShiftRight(src, dest *byte, rows, columns int, Kernel *int16, NRightShift byte) int {
	return int( C.SDL_imageFilterConvolveKernel5x5ShiftRight( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns), s16p(ptr(Kernel)), C.uchar(NRightShift)) )
}

func ImageFilterConvolveKernel7x7ShiftRight(src, dest *byte, rows, columns int, Kernel *int16, NRightShift byte) int {
	return int( C.SDL_imageFilterConvolveKernel7x7ShiftRight( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns), s16p(ptr(Kernel)), C.uchar(NRightShift)) )
}

func ImageFilterConvolveKernel9x9ShiftRight(src, dest *byte, rows, columns int, Kernel *int16, NRightShift byte) int {
	return int( C.SDL_imageFilterConvolveKernel9x9ShiftRight( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns), s16p(ptr(Kernel)), C.uchar(NRightShift)) )
}

func ImageFilterSobelX(src, dest *byte, rows, columns int) int {
	return int( C.SDL_imageFilterSobelX( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns)) )
}

func SDL_imageFilterSobelXShiftRight(src, dest *byte, rows, columns int, NRightShift byte) int {
	return int( C.SDL_imageFilterSobelXShiftRight( u8p(ptr(src)), u8p(ptr(dest)), C.int(rows), C.int(columns), C.uchar(NRightShift)) )
}

func ImageFilterAlignStack() { C.SDL_imageFilterAlignStack() }
func ImageFilterRestoreStack() { C.SDL_imageFilterRestoreStack() }
