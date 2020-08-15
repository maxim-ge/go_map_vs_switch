package go_map_vs_switch

import (
	"testing"
)

func Benchmark_gmp_UnpredictableLookupSwitchAdd128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch randInputs[i%len(randInputs)] % 128 {
		case 0:
			n += i
		case 1:
			n += i
		case 2:
			n += i
		case 3:
			n += i
		case 4:
			n += i
		case 5:
			n += i
		case 6:
			n += i
		case 7:
			n += i
		case 8:
			n += i
		case 9:
			n += i
		case 10:
			n += i
		case 11:
			n += i
		case 12:
			n += i
		case 13:
			n += i
		case 14:
			n += i
		case 15:
			n += i
		case 16:
			n += i
		case 17:
			n += i
		case 18:
			n += i
		case 19:
			n += i
		case 20:
			n += i
		case 21:
			n += i
		case 22:
			n += i
		case 23:
			n += i
		case 24:
			n += i
		case 25:
			n += i
		case 26:
			n += i
		case 27:
			n += i
		case 28:
			n += i
		case 29:
			n += i
		case 30:
			n += i
		case 31:
			n += i
		case 32:
			n += i
		case 33:
			n += i
		case 34:
			n += i
		case 35:
			n += i
		case 36:
			n += i
		case 37:
			n += i
		case 38:
			n += i
		case 39:
			n += i
		case 40:
			n += i
		case 41:
			n += i
		case 42:
			n += i
		case 43:
			n += i
		case 44:
			n += i
		case 45:
			n += i
		case 46:
			n += i
		case 47:
			n += i
		case 48:
			n += i
		case 49:
			n += i
		case 50:
			n += i
		case 51:
			n += i
		case 52:
			n += i
		case 53:
			n += i
		case 54:
			n += i
		case 55:
			n += i
		case 56:
			n += i
		case 57:
			n += i
		case 58:
			n += i
		case 59:
			n += i
		case 60:
			n += i
		case 61:
			n += i
		case 62:
			n += i
		case 63:
			n += i
		case 64:
			n += i
		case 65:
			n += i
		case 66:
			n += i
		case 67:
			n += i
		case 68:
			n += i
		case 69:
			n += i
		case 70:
			n += i
		case 71:
			n += i
		case 72:
			n += i
		case 73:
			n += i
		case 74:
			n += i
		case 75:
			n += i
		case 76:
			n += i
		case 77:
			n += i
		case 78:
			n += i
		case 79:
			n += i
		case 80:
			n += i
		case 81:
			n += i
		case 82:
			n += i
		case 83:
			n += i
		case 84:
			n += i
		case 85:
			n += i
		case 86:
			n += i
		case 87:
			n += i
		case 88:
			n += i
		case 89:
			n += i
		case 90:
			n += i
		case 91:
			n += i
		case 92:
			n += i
		case 93:
			n += i
		case 94:
			n += i
		case 95:
			n += i
		case 96:
			n += i
		case 97:
			n += i
		case 98:
			n += i
		case 99:
			n += i
		case 100:
			n += i
		case 101:
			n += i
		case 102:
			n += i
		case 103:
			n += i
		case 104:
			n += i
		case 105:
			n += i
		case 106:
			n += i
		case 107:
			n += i
		case 108:
			n += i
		case 109:
			n += i
		case 110:
			n += i
		case 111:
			n += i
		case 112:
			n += i
		case 113:
			n += i
		case 114:
			n += i
		case 115:
			n += i
		case 116:
			n += i
		case 117:
			n += i
		case 118:
			n += i
		case 119:
			n += i
		case 120:
			n += i
		case 121:
			n += i
		case 122:
			n += i
		case 123:
			n += i
		case 124:
			n += i
		case 125:
			n += i
		case 126:
			n += i
		case 127:
			n += i
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}

func Benchmark_gmp_PredictableLookupSwitchAdd128(b *testing.B) {
	var n int

	for i := 0; i < b.N; i++ {
		switch ascInputs[i%len(randInputs)] % 128 {
		case 0:
			n += i
		case 1:
			n += i
		case 2:
			n += i
		case 3:
			n += i
		case 4:
			n += i
		case 5:
			n += i
		case 6:
			n += i
		case 7:
			n += i
		case 8:
			n += i
		case 9:
			n += i
		case 10:
			n += i
		case 11:
			n += i
		case 12:
			n += i
		case 13:
			n += i
		case 14:
			n += i
		case 15:
			n += i
		case 16:
			n += i
		case 17:
			n += i
		case 18:
			n += i
		case 19:
			n += i
		case 20:
			n += i
		case 21:
			n += i
		case 22:
			n += i
		case 23:
			n += i
		case 24:
			n += i
		case 25:
			n += i
		case 26:
			n += i
		case 27:
			n += i
		case 28:
			n += i
		case 29:
			n += i
		case 30:
			n += i
		case 31:
			n += i
		case 32:
			n += i
		case 33:
			n += i
		case 34:
			n += i
		case 35:
			n += i
		case 36:
			n += i
		case 37:
			n += i
		case 38:
			n += i
		case 39:
			n += i
		case 40:
			n += i
		case 41:
			n += i
		case 42:
			n += i
		case 43:
			n += i
		case 44:
			n += i
		case 45:
			n += i
		case 46:
			n += i
		case 47:
			n += i
		case 48:
			n += i
		case 49:
			n += i
		case 50:
			n += i
		case 51:
			n += i
		case 52:
			n += i
		case 53:
			n += i
		case 54:
			n += i
		case 55:
			n += i
		case 56:
			n += i
		case 57:
			n += i
		case 58:
			n += i
		case 59:
			n += i
		case 60:
			n += i
		case 61:
			n += i
		case 62:
			n += i
		case 63:
			n += i
		case 64:
			n += i
		case 65:
			n += i
		case 66:
			n += i
		case 67:
			n += i
		case 68:
			n += i
		case 69:
			n += i
		case 70:
			n += i
		case 71:
			n += i
		case 72:
			n += i
		case 73:
			n += i
		case 74:
			n += i
		case 75:
			n += i
		case 76:
			n += i
		case 77:
			n += i
		case 78:
			n += i
		case 79:
			n += i
		case 80:
			n += i
		case 81:
			n += i
		case 82:
			n += i
		case 83:
			n += i
		case 84:
			n += i
		case 85:
			n += i
		case 86:
			n += i
		case 87:
			n += i
		case 88:
			n += i
		case 89:
			n += i
		case 90:
			n += i
		case 91:
			n += i
		case 92:
			n += i
		case 93:
			n += i
		case 94:
			n += i
		case 95:
			n += i
		case 96:
			n += i
		case 97:
			n += i
		case 98:
			n += i
		case 99:
			n += i
		case 100:
			n += i
		case 101:
			n += i
		case 102:
			n += i
		case 103:
			n += i
		case 104:
			n += i
		case 105:
			n += i
		case 106:
			n += i
		case 107:
			n += i
		case 108:
			n += i
		case 109:
			n += i
		case 110:
			n += i
		case 111:
			n += i
		case 112:
			n += i
		case 113:
			n += i
		case 114:
			n += i
		case 115:
			n += i
		case 116:
			n += i
		case 117:
			n += i
		case 118:
			n += i
		case 119:
			n += i
		case 120:
			n += i
		case 121:
			n += i
		case 122:
			n += i
		case 123:
			n += i
		case 124:
			n += i
		case 125:
			n += i
		case 126:
			n += i
		case 127:
			n += i
		}
	}

	// n will never be < 0, but checking n should ensure that the entire benchmark loop can't be optimized away.
	if n < 0 {
		b.Fatal("can't happen")
	}
}