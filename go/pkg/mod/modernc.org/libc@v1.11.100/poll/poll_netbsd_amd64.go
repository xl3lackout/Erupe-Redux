// Code generated by 'ccgo poll/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o poll/poll_netbsd_amd64.go -pkgname poll', DO NOT EDIT.

package poll

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	INFTIM                    = -1
	POLLERR                   = 0x0008
	POLLHUP                   = 0x0010
	POLLIN                    = 0x0001
	POLLNVAL                  = 0x0020
	POLLOUT                   = 0x0004
	POLLPRI                   = 0x0002
	POLLRDBAND                = 0x0080
	POLLRDNORM                = 0x0040
	POLLWRBAND                = 0x0100
	POLLWRNORM                = 4
	X_AMD64_INT_TYPES_H_      = 0
	X_FILE_OFFSET_BITS        = 64
	X_LP64                    = 1
	X_NETBSD_SOURCE           = 1
	X_SYS_CDEFS_ELF_H_        = 0
	X_SYS_CDEFS_H_            = 0
	X_SYS_COMMON_ANSI_H_      = 0
	X_SYS_COMMON_INT_TYPES_H_ = 0
	X_SYS_POLL_H_             = 0
	X_SYS_SIGTYPES_H_         = 0
	X_X86_64_CDEFS_H_         = 0
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

//	$NetBSD: poll.h,v 1.15 2009/11/11 09:48:51 rmind Exp $

// -
// Copyright (c) 1998 The NetBSD Foundation, Inc.
// All rights reserved.
//
// This code is derived from software contributed to The NetBSD Foundation
// by Charles M. Hannum.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE NETBSD FOUNDATION, INC. AND CONTRIBUTORS
// ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED
// TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
// PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE FOUNDATION OR CONTRIBUTORS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

//	$NetBSD: featuretest.h,v 1.10 2013/04/26 18:29:06 christos Exp $

// Written by Klaus Klein <kleink@NetBSD.org>, February 2, 1998.
// Public domain.
//
// NOTE: Do not protect this header against multiple inclusion.  Doing
// so can have subtle side-effects due to header file inclusion order
// and testing of e.g. _POSIX_SOURCE vs. _POSIX_C_SOURCE.  Instead,
// protect each CPP macro that we want to supply.

// Feature-test macros are defined by several standards, and allow an
// application to specify what symbols they want the system headers to
// expose, and hence what standard they want them to conform to.
// There are two classes of feature-test macros.  The first class
// specify complete standards, and if one of these is defined, header
// files will try to conform to the relevant standard.  They are:
//
// ANSI macros:
// _ANSI_SOURCE			ANSI C89
//
// POSIX macros:
// _POSIX_SOURCE == 1		IEEE Std 1003.1 (version?)
// _POSIX_C_SOURCE == 1		IEEE Std 1003.1-1990
// _POSIX_C_SOURCE == 2		IEEE Std 1003.2-1992
// _POSIX_C_SOURCE == 199309L	IEEE Std 1003.1b-1993
// _POSIX_C_SOURCE == 199506L	ISO/IEC 9945-1:1996
// _POSIX_C_SOURCE == 200112L	IEEE Std 1003.1-2001
// _POSIX_C_SOURCE == 200809L   IEEE Std 1003.1-2008
//
// X/Open macros:
// _XOPEN_SOURCE		System Interfaces and Headers, Issue 4, Ver 2
// _XOPEN_SOURCE_EXTENDED == 1	XSH4.2 UNIX extensions
// _XOPEN_SOURCE == 500		System Interfaces and Headers, Issue 5
// _XOPEN_SOURCE == 520		Networking Services (XNS), Issue 5.2
// _XOPEN_SOURCE == 600		IEEE Std 1003.1-2001, XSI option
// _XOPEN_SOURCE == 700		IEEE Std 1003.1-2008, XSI option
//
// NetBSD macros:
// _NETBSD_SOURCE == 1		Make all NetBSD features available.
//
// If more than one of these "major" feature-test macros is defined,
// then the set of facilities provided (and namespace used) is the
// union of that specified by the relevant standards, and in case of
// conflict, the earlier standard in the above list has precedence (so
// if both _POSIX_C_SOURCE and _NETBSD_SOURCE are defined, the version
// of rename() that's used is the POSIX one).  If none of the "major"
// feature-test macros is defined, _NETBSD_SOURCE is assumed.
//
// There are also "minor" feature-test macros, which enable extra
// functionality in addition to some base standard.  They should be
// defined along with one of the "major" macros.  The "minor" macros
// are:
//
// _REENTRANT
// _ISOC99_SOURCE
// _ISOC11_SOURCE
// _LARGEFILE_SOURCE		Large File Support
//		<http://ftp.sas.com/standards/large.file/x_open.20Mar96.html>

type Nfds_t = uint32 /* poll.h:37:22 */

type Pollfd = struct {
	Ffd      int32
	Fevents  int16
	Frevents int16
} /* poll.h:39:1 */

//	$NetBSD: sigtypes.h,v 1.11 2017/01/12 18:29:14 christos Exp $

// Copyright (c) 1982, 1986, 1989, 1991, 1993
//	The Regents of the University of California.  All rights reserved.
// (c) UNIX System Laboratories, Inc.
// All or some portions of this file are derived from material licensed
// to the University of California by American Telephone and Telegraph
// Co. or Unix System Laboratories, Inc. and are reproduced herein with
// the permission of UNIX System Laboratories, Inc.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)signal.h	8.4 (Berkeley) 5/4/95

// This header file defines various signal-related types.  We also keep
// the macros to manipulate sigset_t here, to encapsulate knowledge of
// its internals.

//	$NetBSD: featuretest.h,v 1.10 2013/04/26 18:29:06 christos Exp $

// Written by Klaus Klein <kleink@NetBSD.org>, February 2, 1998.
// Public domain.
//
// NOTE: Do not protect this header against multiple inclusion.  Doing
// so can have subtle side-effects due to header file inclusion order
// and testing of e.g. _POSIX_SOURCE vs. _POSIX_C_SOURCE.  Instead,
// protect each CPP macro that we want to supply.

// Feature-test macros are defined by several standards, and allow an
// application to specify what symbols they want the system headers to
// expose, and hence what standard they want them to conform to.
// There are two classes of feature-test macros.  The first class
// specify complete standards, and if one of these is defined, header
// files will try to conform to the relevant standard.  They are:
//
// ANSI macros:
// _ANSI_SOURCE			ANSI C89
//
// POSIX macros:
// _POSIX_SOURCE == 1		IEEE Std 1003.1 (version?)
// _POSIX_C_SOURCE == 1		IEEE Std 1003.1-1990
// _POSIX_C_SOURCE == 2		IEEE Std 1003.2-1992
// _POSIX_C_SOURCE == 199309L	IEEE Std 1003.1b-1993
// _POSIX_C_SOURCE == 199506L	ISO/IEC 9945-1:1996
// _POSIX_C_SOURCE == 200112L	IEEE Std 1003.1-2001
// _POSIX_C_SOURCE == 200809L   IEEE Std 1003.1-2008
//
// X/Open macros:
// _XOPEN_SOURCE		System Interfaces and Headers, Issue 4, Ver 2
// _XOPEN_SOURCE_EXTENDED == 1	XSH4.2 UNIX extensions
// _XOPEN_SOURCE == 500		System Interfaces and Headers, Issue 5
// _XOPEN_SOURCE == 520		Networking Services (XNS), Issue 5.2
// _XOPEN_SOURCE == 600		IEEE Std 1003.1-2001, XSI option
// _XOPEN_SOURCE == 700		IEEE Std 1003.1-2008, XSI option
//
// NetBSD macros:
// _NETBSD_SOURCE == 1		Make all NetBSD features available.
//
// If more than one of these "major" feature-test macros is defined,
// then the set of facilities provided (and namespace used) is the
// union of that specified by the relevant standards, and in case of
// conflict, the earlier standard in the above list has precedence (so
// if both _POSIX_C_SOURCE and _NETBSD_SOURCE are defined, the version
// of rename() that's used is the POSIX one).  If none of the "major"
// feature-test macros is defined, _NETBSD_SOURCE is assumed.
//
// There are also "minor" feature-test macros, which enable extra
// functionality in addition to some base standard.  They should be
// defined along with one of the "major" macros.  The "minor" macros
// are:
//
// _REENTRANT
// _ISOC99_SOURCE
// _ISOC11_SOURCE
// _LARGEFILE_SOURCE		Large File Support
//		<http://ftp.sas.com/standards/large.file/x_open.20Mar96.html>

//	$NetBSD: int_types.h,v 1.7 2014/07/25 21:43:13 joerg Exp $

// -
// Copyright (c) 1990 The Regents of the University of California.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	from: @(#)types.h	7.5 (Berkeley) 3/9/91

//	$NetBSD: common_int_types.h,v 1.1 2014/07/25 21:43:13 joerg Exp $

// -
// Copyright (c) 2014 The NetBSD Foundation, Inc.
// All rights reserved.
//
// This code is derived from software contributed to The NetBSD Foundation
// by Joerg Sonnenberger.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE NETBSD FOUNDATION, INC. AND CONTRIBUTORS
// ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED
// TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
// PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE FOUNDATION OR CONTRIBUTORS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

// 7.18.1 Integer types

// 7.18.1.1 Exact-width integer types

type X__int8_t = int8     /* common_int_types.h:45:27 */
type X__uint8_t = uint8   /* common_int_types.h:46:27 */
type X__int16_t = int16   /* common_int_types.h:47:27 */
type X__uint16_t = uint16 /* common_int_types.h:48:27 */
type X__int32_t = int32   /* common_int_types.h:49:27 */
type X__uint32_t = uint32 /* common_int_types.h:50:27 */
type X__int64_t = int64   /* common_int_types.h:51:27 */
type X__uint64_t = uint64 /* common_int_types.h:52:27 */

// 7.18.1.4 Integer types capable of holding object pointers

type X__intptr_t = int64   /* common_int_types.h:58:27 */
type X__uintptr_t = uint64 /* common_int_types.h:59:26 */

type Sigset_t = struct{ F__bits [4]X__uint32_t } /* sigtypes.h:62:3 */

// Macro for manipulating signal masks.

type Sigaltstack = struct {
	Fss_sp       uintptr
	Fss_size     Size_t
	Fss_flags    int32
	F__ccgo_pad1 [4]byte
} /* sigtypes.h:108:9 */

// Macro for manipulating signal masks.

type Stack_t = Sigaltstack /* sigtypes.h:116:3 */

var _ int8 /* gen.c:2:13: */
