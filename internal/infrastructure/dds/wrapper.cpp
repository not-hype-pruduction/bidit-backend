#include "wrapper.h"
#include "include/dll.h"
#include "string.h"
#include <cstring>

int CalcTablePBN(char *pbn_deal, DDTableResult *out) {
    ddTableDealPBN tableDealPBN;
    strcpy(tableDealPBN.cards, pbn_deal);

    ddTableResults results;
    int res = CalcDDtablePBN(tableDealPBN, &results);

    if (res == 1) {
        for (int suit = 0; suit < 5; suit++) {
            for (int hand = 0; hand < 4; hand++) {
                out->res[suit][hand] = results.resTable[suit][hand];
            }
        }
    }

    return res;
}

void SetDDSMaxThreads(int max_threads) {
    SetMaxThreads(max_threads);
}
