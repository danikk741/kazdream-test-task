package custom_map

func hash(k []byte) int {
	var res int
	switch k[0] {
	case 97:
		res = 0
	case 98:
		res = 0
	case 99:
		res = 1
	case 100:
		res = 1
	case 101:
		res = 2
	case 102:
		res = 2
	case 103:
		res = 3
	case 104:
		res = 3
	case 105:
		res = 4
	case 106:
		res = 4
	case 107:
		res = 5
	case 108:
		res = 5
	case 109:
		res = 6
	case 110:
		res = 6
	case 111:
		res = 7
	case 112:
		res = 7
	case 113:
		res = 8
	case 114:
		res = 8
	case 115:
		res = 9
	case 116:
		res = 9
	case 117:
		res = 10
	case 118:
		res = 10
	case 119:
		res = 11
	case 120:
		res = 11
	case 121:
		res = 12
	case 122:
		res = 12
	}
	return res
}
