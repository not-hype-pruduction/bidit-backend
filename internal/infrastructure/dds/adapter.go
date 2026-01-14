package dds

/*
#cgo CXXFLAGS: -std=c++11 -I${SRCDIR}/dds -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR}/dds -ldds
#include "wrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/not-hype-pruduction/bridge-backend/internal/domain/dds"
)

type ddsAdapter struct{}

func NewDDSAdapter() *ddsAdapter {
	numThreads := runtime.GOMAXPROCS(0)

	C.SetDDSMaxThreads(C.int(numThreads))

	return &ddsAdapter{}
}

func (a *ddsAdapter) CalculateTable(ctx context.Context, pbn string) (dds.DDTable, error) {
	cPbn := C.CString(pbn)
	defer C.free(unsafe.Pointer(cPbn))

	var res C.DDTableResult
	ret := C.CalcTablePBN(cPbn, &res)

	if ret != 1 { // RETURN_NO_FAULT в DDS обычно 1
		return dds.DDTable{}, fmt.Errorf("DDS error code: %d", ret)
	}

	suits := []string{"NT", "S", "H", "D", "C"}
	table := dds.DDTable{Results: make(map[string][4]int)}

	for i, suit := range suits {
		var r [4]int
		for hand := range 4 {
			r[hand] = int(res.res[i][hand])
		}
		table.Results[suit] = r
	}

	return table, nil
}
