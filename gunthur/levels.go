package gunthur

func loadLevels() map[int]*Scene {
	levels := make(map[int]*Scene)

	level1 := NewScene()
	level1.AddComponent(NewPlayer())
	level1.AddComponent(NewEnemy())
	levels[1] = level1

	return levels
}
