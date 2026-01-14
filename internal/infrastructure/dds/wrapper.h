#ifndef DDS_WRAPPER_H
#define DDS_WRAPPER_H

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    int res[5][4]; // [strain: NT,S,H,D,C][direction: N,E,S,W]
} DDTableResult;

// PBN format: "N:QJ6.K652.J85.T98 873.J97.AT764.Q4 K5.AQT83.KQ9.A73 AT942.4.32.KJ652"
int CalcTablePBN(char* pbn_deal, DDTableResult* out);

void SetDDSMaxThreads(int n);

#ifdef __cplusplus
}
#endif

#endif
