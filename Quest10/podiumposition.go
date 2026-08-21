package Quest10

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium)-1; i++ {
		for j := i + 1; j < len(podium); j++ {

			numI := int(podium[i][0][0] - '0')
			numJ := int(podium[j][0][0] - '0')

			if numI > numJ {
				podium[i], podium[j] = podium[j], podium[i]
			}
		}
	}
	return podium
}