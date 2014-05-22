package gothumb

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

var sourceInput []byte = []byte{
	255, 216, 255, 219, 0, 67, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 255, 219, 0, 67, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 255, 194, 0, 17, 8, 0, 16, 0, 16, 3, 1, 17, 0,
	2, 17, 1, 3, 17, 1, 255, 196, 0, 22, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 9, 7, 8, 255, 196, 0, 23, 1, 0, 3, 1, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 9, 10, 7, 255, 218, 0, 12, 3, 1, 0,
	2, 16, 3, 16, 0, 0, 1, 197, 226, 155, 252, 74, 5, 184, 87, 6, 237,
	49, 172, 66, 208, 30, 169, 255, 196, 0, 20, 16, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0, 8, 1, 1, 0, 1, 5, 2, 31,
	255, 196, 0, 20, 17, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	32, 255, 218, 0, 8, 1, 3, 1, 1, 63, 1, 31, 255, 196, 0, 20, 17, 1,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0, 8, 1,
	2, 1, 1, 63, 1, 31, 255, 196, 0, 20, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0, 8, 1, 1, 0, 6, 63, 2, 31, 255,
	196, 0, 20, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32,
	255, 218, 0, 8, 1, 1, 0, 1, 63, 33, 31, 255, 218, 0, 12, 3, 1, 0, 2,
	0, 3, 0, 0, 0, 16, 0, 15, 255, 196, 0, 20, 17, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0, 8, 1, 3, 1, 1, 63, 16,
	31, 255, 196, 0, 20, 17, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 32, 255, 218, 0, 8, 1, 2, 1, 1, 63, 16, 31, 255, 196, 0, 20,
	16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0,
	8, 1, 1, 0, 1, 63, 16, 31, 255, 217,
}
var thumbnailOutput []byte = []byte{
	255, 216, 255, 219, 0, 132, 0, 8, 6, 6, 7, 6, 5, 8, 7, 7, 7, 9, 9, 8,
	10, 12, 20, 13, 12, 11, 11, 12, 25, 18, 19, 15, 20, 29, 26, 31, 30,
	29, 26, 28, 28, 32, 36, 46, 39, 32, 34, 44, 35, 28, 28, 40, 55, 41,
	44, 48, 49, 52, 52, 52, 31, 39, 57, 61, 56, 50, 60, 46, 51, 52, 50, 1,
	9, 9, 9, 12, 11, 12, 24, 13, 13, 24, 50, 33, 28, 33, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 255, 192, 0, 17, 8, 0,
	4, 0, 4, 3, 1, 34, 0, 2, 17, 1, 3, 17, 1, 255, 196, 1, 162, 0, 0, 1,
	5, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 16, 0, 2, 1, 3, 3, 2, 4, 3, 5, 5, 4, 4, 0, 0, 1, 125, 1, 2,
	3, 0, 4, 17, 5, 18, 33, 49, 65, 6, 19, 81, 97, 7, 34, 113, 20, 50,
	129, 145, 161, 8, 35, 66, 177, 193, 21, 82, 209, 240, 36, 51, 98, 114,
	130, 9, 10, 22, 23, 24, 25, 26, 37, 38, 39, 40, 41, 42, 52, 53, 54,
	55, 56, 57, 58, 67, 68, 69, 70, 71, 72, 73, 74, 83, 84, 85, 86, 87,
	88, 89, 90, 99, 100, 101, 102, 103, 104, 105, 106, 115, 116, 117, 118,
	119, 120, 121, 122, 131, 132, 133, 134, 135, 136, 137, 138, 146, 147,
	148, 149, 150, 151, 152, 153, 154, 162, 163, 164, 165, 166, 167, 168,
	169, 170, 178, 179, 180, 181, 182, 183, 184, 185, 186, 194, 195, 196,
	197, 198, 199, 200, 201, 202, 210, 211, 212, 213, 214, 215, 216, 217,
	218, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 241, 242, 243,
	244, 245, 246, 247, 248, 249, 250, 1, 0, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 17, 0, 2, 1, 2,
	4, 4, 3, 4, 7, 5, 4, 4, 0, 1, 2, 119, 0, 1, 2, 3, 17, 4, 5, 33, 49, 6,
	18, 65, 81, 7, 97, 113, 19, 34, 50, 129, 8, 20, 66, 145, 161, 177,
	193, 9, 35, 51, 82, 240, 21, 98, 114, 209, 10, 22, 36, 52, 225, 37,
	241, 23, 24, 25, 26, 38, 39, 40, 41, 42, 53, 54, 55, 56, 57, 58, 67,
	68, 69, 70, 71, 72, 73, 74, 83, 84, 85, 86, 87, 88, 89, 90, 99, 100,
	101, 102, 103, 104, 105, 106, 115, 116, 117, 118, 119, 120, 121, 122,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 162, 163, 164, 165, 166, 167, 168, 169, 170, 178,
	179, 180, 181, 182, 183, 184, 185, 186, 194, 195, 196, 197, 198, 199,
	200, 201, 202, 210, 211, 212, 213, 214, 215, 216, 217, 218, 226, 227,
	228, 229, 230, 231, 232, 233, 234, 242, 243, 244, 245, 246, 247, 248,
	249, 250, 255, 218, 0, 12, 3, 1, 0, 2, 17, 3, 17, 0, 63, 0, 242, 83,
	227, 95, 16, 216, 159, 179, 91, 106, 27, 33, 79, 186, 190, 76, 103,
	25, 231, 169, 95, 83, 73, 255, 0, 11, 3, 197, 31, 244, 19, 255, 0,
	201, 120, 191, 248, 154, 193, 188, 255, 0, 143, 183, 252, 63, 149, 65,
	90, 102, 88, 12, 44, 113, 181, 163, 26, 81, 73, 74, 93, 23, 119, 228,
	114, 75, 7, 135, 172, 221, 74, 148, 227, 41, 75, 86, 218, 77, 182,
	247, 109, 245, 111, 171, 63, 255, 217,
}

func TestThumbnail(t *testing.T) {
	reader := bytes.NewReader(sourceInput)

	out, err := Thumbnail(reader, 4, 75)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	content, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(content, thumbnailOutput) {
		t.Errorf("Thumbnail failed")
	}
}