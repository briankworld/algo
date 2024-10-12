func minAreaRect(points [][]int) int {
    s := make(map[string]struct{})
	minArea := math.MaxInt32

	for _, p := range points {
		key := fmt.Sprintf("%d,%d", p[0], p[1])
		s[key] = struct{}
	}

	for i:=0; i <len(points); i++ {
		for j := i+1; j < len(points); j++ {
			x1, y1 = points[i][0], points[i][1]
			x2, y2 = points[j][0], points[j][1]

			if x1 != x2 && y1 != y2 {
				if _, ok1 := s[fmt.Sprintf("%d,%d", x1, y2)]; ok1 &&
				  _, ok2 := s[fmt.Sprintf("%d,%d", x2, y1)]; ok2 {
					area := abs(x1-x2) * abs(y1-y2)
					minArea = min(minArea, area)
				}
			}
		}
	}

	if minArea == math.MaxInt32 {
		return 0
	}
	return minArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}