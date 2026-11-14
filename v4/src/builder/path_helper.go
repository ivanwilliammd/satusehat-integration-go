package builder

// setPath meniru set('a/b', v) pada Builder PHP: segmen dipisah "/" membentuk objek bertingkat.
// Dipakai builder yang elemennya bersarang tanpa datatype tersendiri, misalnya collection pada
// Specimen. Nilai di segmen terakhir ditimpa, bukan digabung.
func setPath(data map[string]interface{}, path string, v interface{}) {
	segmen := splitPath(path)
	cur := data
	for _, s := range segmen[:len(segmen)-1] {
		next, _ := cur[s].(map[string]interface{})
		if next == nil {
			next = map[string]interface{}{}
			cur[s] = next
		}
		cur = next
	}
	cur[segmen[len(segmen)-1]] = v
}

// pushPath seperti setPath, tetapi menambahkan ke list di segmen terakhir.
func pushPath(data map[string]interface{}, path string, v interface{}) {
	segmen := splitPath(path)
	cur := data
	for _, s := range segmen[:len(segmen)-1] {
		next, _ := cur[s].(map[string]interface{})
		if next == nil {
			next = map[string]interface{}{}
			cur[s] = next
		}
		cur = next
	}
	kunci := segmen[len(segmen)-1]
	list, _ := cur[kunci].([]interface{})
	cur[kunci] = append(list, v)
}

func splitPath(path string) []string {
	out := []string{}
	mulai := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			out = append(out, path[mulai:i])
			mulai = i + 1
		}
	}
	return append(out, path[mulai:])
}
