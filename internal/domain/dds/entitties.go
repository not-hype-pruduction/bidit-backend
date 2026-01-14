package dds

type DDTable struct {
	// Ключи: масть (NT, S, H, D, C), Значение: карта игроков [N, E, S, W]
	Results map[string][4]int
}
