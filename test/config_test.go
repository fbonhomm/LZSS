package test

import "github.com/fbonhomm/Simple-LZSS/source"

var Lzss = source.LZSS{}

// SIMPLE
var DecSimpleStringWithoutDuplicate = []byte(
	"simple test")                                                       	// length 11
var ComSimpleStringWithoutDuplicate = []byte{
	231, 166, 109, 11, 151, 109, 89, 144, 116, 203, 206, 165, 3}         	// length 13

var DecSimpleStringWithDuplicate = []byte(
	"simple test simple test")                                              // length 23
var ComSimpleStringWithDuplicate = []byte{
	231, 166, 109, 11, 151, 109, 89, 144, 116, 203, 206, 165, 11, 2, 0, 16} // length 16


// MEDIUM
var DecMediumString = []byte(
	`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus sed tempor magna. Curabitur at lobortis sem.
	Aliquam pretium, nunc ut consectetur venenatis, enim augue rutrum nibh, vitae iaculis est augue ut est.
	Praesent vehicula lacinia enim in faucibus. Maecenas quis porttitor nisi.In dolor orci, auctor ut cursus vitae, dignissim vitae ante.
	Nulla nec lorem commodo, vehicula massa a, vulputate mi.Pellentesque nibh nibh, bibendum quis vestibulum sit amet, vulputate convallis turpis.
	Integer non ornare purus.Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Maecenas in tortor sed enim condimentum rhoncus a eu orci.
	Vivamus odio lorem, tincidunt eget nisi ac, venenatis interdum neque. Aliquam quis neque a dui efficitur pellentesque vitae eu augue.
	Nulla at tempus magna. In augue orci, vehicula non imperdiet a, sagittis vitae massa. Duis ultricies ante nisi, eget dignissim lorem lacinia sit amet.
	Nullam luctus, diam eget elementum bibendum, ex odio rhoncus est, eu congue orci enim non tortor. Vestibulum non diam at lorem tincidunt rutrum sit amet sit amet leo.
	Mauris lorem risus, fermentum vitae eros ut, mollis faucibus sapien. Nullam sed mauris mi. Vivamus sit amet metus pharetra, ultricies ligula a, pellentesque nibh.
	Mauris eget tortor massa. Phasellus et quam vitae erat posuere pulvinar at et nunc. Fusce est erat, aliquam hendrerit congue eu, egestas eget turpis.
	Nunc vel eros ac lectus interdum ullamcorper nec id dui. Phasellus ut velit euismod, malesuada lacus quis, porta nisi. Sed ultricies lorem id lorem iaculis, a sodales mi semper.
	Proin feugiat justo ut urna commodo, rutrum mattis sem tincidunt. Donec condimentum a urna sit amet blandit.
	Sed lorem leo, ullamcorper sit amet feugiat ac, elementum eu lorem. Suspendisse tristique interdum ex, ut laoreet lacus gravida vel. Nullam varius cursus volutpat.
	Aenean pellentesque orci aliquet posuere pellentesque. Vestibulum in enim sed sapien tincidunt mollis nec et nunc metus.`) // length 2012
var ComMediumString = []byte{
	153,190,149,91,182,45,200,180,112,231,214,109,11,146,236,91,182,111,229,130,156,155,150,46,200,176,109,203,210,101,
	9,114,236,91,183,115,203,142,165,91,150,110,93,185,32,195,146,77,11,55,237,220,177,105,221,158,5,89,150,109,90,186,
	46,65,90,77,107,55,108,219,186,115,65,206,45,75,22,36,221,178,109,225,60,0,108,27,246,172,219,176,46,65,14,173,43,
	55,172,216,52,36,32,233,130,100,251,86,236,91,185,116,211,252,64,108,235,82,97,194,160,108,211,198,173,27,182,45,72,
	184,114,203,210,77,91,183,45,75,144,110,235,186,29,11,178,46,157,13,208,118,203,186,45,235,54,76,25,64,150,32,203,
	186,77,219,22,100,216,186,103,235,150,5,41,183,46,93,57,9,0,221,166,21,139,150,37,72,187,105,233,134,45,11,50,109,
	216,177,117,217,152,1,44,59,135,2,64,41,136,66,0,197,0,216,0,66,149,27,182,236,220,178,110,233,76,2,68,155,230,23,
	96,88,144,108,195,184,0,76,27,198,19,102,90,183,32,205,134,173,59,54,173,216,186,115,93,130,108,26,182,236,152,75,
	128,115,65,198,173,99,6,0,9,32,93,186,105,233,60,0,116,155,118,110,90,151,73,221,44,0,125,43,119,108,90,22,163,0,
	199,96,68,40,132,88,87,238,28,31,0,183,32,89,130,36,155,246,12,71,192,57,161,0,110,97,134,117,75,183,12,27,64,167,
	117,217,156,3,116,91,118,204,23,0,1,136,99,223,182,109,251,150,236,155,90,0,226,80,219,134,157,59,55,44,200,48,181,
	0,235,178,133,91,151,110,88,186,101,65,182,77,235,18,106,89,182,108,186,1,150,157,27,231,20,64,44,4,88,160,98,211,
	138,45,235,150,76,2,128,134,144,118,138,1,248,3,100,147,0,64,2,142,98,28,14,128,118,195,178,9,134,64,2,96,5,0,27,
	192,36,89,1,207,150,209,8,248,214,77,74,64,183,97,229,150,5,9,183,174,28,128,128,86,124,227,96,69,40,0,10,87,110,
	218,54,102,0,238,1,165,132,4,201,38,38,128,15,192,178,116,65,214,101,75,87,110,218,177,101,32,34,206,173,91,230,63,
	224,216,186,98,211,178,121,7,48,5,177,236,142,129,56,247,0,48,2,96,68,8,130,240,9,130,3,32,217,180,109,186,1,36,0,
	148,139,246,173,219,49,62,0,224,1,214,73,9,193,6,32,7,234,91,178,105,223,160,197,100,9,146,110,90,183,99,211,146,
	173,227,13,64,62,32,157,141,144,32,195,142,177,139,96,9,196,30,32,221,178,114,104,35,186,45,3,26,208,37,136,55,40,
	13,33,182,69,60,10,36,91,55,45,200,178,102,205,166,29,131,5,18,14,102,164,91,24,146,2,72,65,112,133,233,2,80,4,193,
	7,0,19,144,144,0,164,48,84,98,132,131,185,143,152,182,45,156,106,129,105,122,4,28,6,156,27,246,108,154,138,0,63,65,
	0,99,186,4,73,196,33,2,144,204,52,70,2,38,68,112,68,232,4,176,43,1,41,78,90,12,117,40,18,80,184,162,128,14,184,17,
	97,20,129,14,128,174,200,178,108,203,4,10,121,141,210,9,16,207,168,12,68,41,197,0,58,1,214,217,128,240,51,197,19,
	198,125,4,54,25,55,0,124,227,220,71,80,30,209,5,66,0,2,172,28,85,32,18,80,34,192,130,100,91,246,13,27,192,166,97,
	235,202,49,3,144,22,195,140,128,241,68,179,101,229,4,10,145,24,42,247,13,223,128,12,128,109,223,132,135,64,17,59,55,
	44,220,180,101,221,186,4,177,59,36,8,98,27,50,26,148,1,112,224,48,81,185,0,224,3,16,46,26,255,128,116,229,28,6,240,
	155,203,54,237,25,115,8,135,1,10,70,197,66,0,163,163,43,130,77,136,220,80,168,104,195,206,193,12,160,35,67,14,68,32,
	149,46,0,73,82,225,214,101,107,55,77,127,0,46,8,188,18,2,1,186,4,105,180,238,220,49,22,3,140,1,228,84,112,9,0,7,20,
	141,108,64,185,101,229,76,0,0,32,90,182,78,228,200,185,116,20,2,56,212,56,143,234,36,16,162,221,178,108,12,105,134,
	57,11,0,2,192,7,128,180,24,222,145,99,223,202,193,26,64,22,50,45,25,123,9,31,149,10,1,140,22,152,0,88,198,33,128,92,
	0,71,130,97,217,150,157,91,55,44,25,116,16,62,0,26,34,196,19,24,3,48,22,160,35,34,200,41,66,0,190,9,5,192,36,35,98,
	197,23,134,75,64,144,115,223,146,161,175,48,25,0,13,2,214,128,106,144,111,238,33,150,173,123,54,77,23,64,181,117,231,
	210,125,67,8,177,14,127,0,183,56,170,64,219,134,57,27,145,6,224,85,206,221,192,55,101,17,248,36,135,57,179,177,68,86,
	44,219,48,126,2,168,33,84,88,16,29,90,166,46,192,111,165,37,50,87,86,136,37,80,207,145,20,144,22,98,194,96,221,185,
	112,100,3,48,5,44,11,66,72,192,111,0,208,8,210,98,203,226,233,38,224,5,48,76,0,0,45,10,254,37,207,202,13,107,55,141,
	126,129,209,2,166,100,237,134,149,155,198,7,128,78,212,183,108,235,210,133,27,134,26,98,144,74,128,97,221,16,204,138,
	17,177,86,162,71,128,79,25,204,64,27,164,238,1,124,130,8,130,0,74,227,85,16,36,145,178,8,82,69,156,105,186, 00} // length 1215

// LONG
var DecLongString = []byte(
	`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam vitae orci nec eros tincidunt iaculis.
	Morbi dapibus fringilla fringilla. Etiam et diam risus. Donec at hendrerit turpis. Integer vulputate mi a posuere pellentesque.
	Ut lacinia nulla a velit iaculis, non finibus libero auctor. Maecenas tempor sagittis porttitor.
	Fusce enim enim, aliquet eget consequat sed, elementum ut neque. Aenean in tincidunt eros.
	Nunc gravida, nisi ac malesuada rhoncus, nunc nibh condimentum tellus, nec rutrum sapien ex vel ipsum.
	Morbi vestibulum viverra sem vitae euismod. Suspendisse potenti. Fusce congue eu turpis nec aliquam.
	Proin sagittis nisl vel nunc sollicitudin, eget dapibus ipsum sagittis. Integer risus leo, rutrum non eros ut, accumsan sagittis ex.
	Pellentesque molestie ac ante rhoncus viverra. Pellentesque metus nibh, tincidunt ac lectus eget, placerat finibus risus.
	Donec quis ligula lobortis ligula commodo maximus sit amet vel neque.In feugiat eu ligula at pretium.
	Pellentesque placerat condimentum scelerisque. Suspendisse ut leo pretium, convallis dui nec, ultrices dolor.
	Pellentesque tincidunt, orci interdum hendrerit ornare, mi ex bibendum nisi, at laoreet lectus felis sit amet ligula.
	Nullam fringilla lorem ut ipsum ultricies dictum. Curabitur mattis pharetra purus vel fringilla. Vivamus elementum ac est non varius.
	Nulla tortor sem, pellentesque id mi quis, iaculis faucibus enim. Phasellus pellentesque erat eget nibh aliquet dapibus.
	Nunc eleifend mauris nunc, in maximus est semper et. In ut purus dictum, commodo libero vitae, imperdiet eros.
	In eget augue ac enim tincidunt varius. Sed et condimentum tellus. Sed eget purus urna.
	In fermentum ligula odio, non aliquet ex sagittis pellentesque. Fusce eget dui tincidunt lacus laoreet bibendum id tincidunt ligula.
	Donec eu tristique dui, nec ornare odio. Phasellus convallis lorem a dictum euismod. Sed eget diam eu purus luctus maximus.
	Cras id velit posuere, rhoncus est in, feugiat velit. Morbi placerat vehicula egestas. In malesuada mauris eget urna facilisis lacinia.
	Aliquam at dolor a massa rutrum tincidunt.In sit amet enim posuere, sollicitudin purus sed, ultricies libero. Morbi vitae ex tellus.
	Cras odio felis, pharetra vel vestibulum eu, pharetra at justo. Aliquam dolor orci, tristique eu volutpat a, congue et lorem.
	Fusce volutpat posuere sem quis rhoncus. Integer dictum eleifend ante, sed rutrum lectus ornare a.
	Morbi nec neque ac velit faucibus rhoncus. Sed consequat, tellus nec interdum pellentesque, ligula metus finibus quam, in mollis nisi urna ac ipsum.
	Donec condimentum venenatis facilisis. Sed at porta ante.Integer pulvinar eros eu augue consequat, vel condimentum lorem dictum.
	Pellentesque scelerisque sed lacus a faucibus. Donec dapibus justo non sagittis fermentum. Pellentesque bibendum urna in urna fringilla finibus.
	Nulla quis sagittis risus. Donec hendrerit non justo quis porttitor. Maecenas ultricies nunc id ligula placerat, vel ullamcorper felis scelerisque.
	Nullam aliquet est id fringilla vulputate.Mauris id efficitur est. Cras neque risus, pretium at massa rutrum, luctus egestas eros.
	Pellentesque non dictum turpis. Praesent pharetra erat in orci scelerisque eleifend. Nulla non lorem molestie, egestas ex a, tristique arcu.
	Sed suscipit tempus urna, eu rutrum arcu iaculis semper. Nunc sed ipsum arcu. Sed sit amet tincidunt nisl, eu tristique mauris.
	Donec mi mi, hendrerit id velit ornare, sollicitudin placerat mi. Quisque in augue eget diam mattis suscipit.
	Donec viverra diam ut nulla tempus luctus. Suspendisse ultrices nisi et neque lacinia, sit amet tristique ligula dignissim.
	Cras ultrices pretium nisl. Vestibulum volutpat dapibus nisi, quis accumsan magna scelerisque eu.Sed vitae lorem diam.
	Curabitur feugiat ornare justo, et mollis risus ultricies et. Nulla molestie tellus nec tortor aliquet, quis aliquet lacus viverra.
	Aliquam interdum, neque dictum luctus mattis, nibh augue molestie magna, vel consectetur enim augue quis enim. Ut eu porta arcu.
	In hac habitasse platea dictumst. Quisque nibh est, tempus placerat lectus eu, faucibus dictum ante.
	Maecenas dolor augue, congue in placerat eu, facilisis in tellus. Phasellus vel lorem nibh. Sed cursus nulla ac lorem iaculis vulputate.
	Donec lorem metus, imperdiet eget est quis, tempus consequat est. Nullam sed enim id sem facilisis hendrerit non ut nisl.
	Nam elit lacus, malesuada sit amet elit id, pulvinar tincidunt sem. Etiam in orci ut est viverra tempor non ut felis.
	Phasellus sollicitudin in lacus et fringilla. Nullam pretium semper rhoncus. Etiam vel auctor eros.
	Nunc maximus, risus nec laoreet mollis, urna mi aliquet metus, pulvinar suscipit est metus eget risus.
	Suspendisse commodo mi placerat ante porta, ut posuere tellus tincidunt. Morbi porta venenatis facilisis.
	Sed molestie est at est accumsan rutrum.Duis porta iaculis est non pulvinar. Quisque vitae gravida sem. Nunc semper eget urna non venenatis.
	Curabitur lacinia tempus enim at efficitur. Phasellus luctus diam leo nullam.`) // length 5044
var ComLongString = []byte{
	153,190,149,91,182,45,200,180,112,231,214,109,11,146,236,91,182,111,229,130,156,155,150,46,200,176,109,203,210,101,9,114,
	236,91,183,115,203,142,165,91,150,110,93,185,32,195,146,77,11,55,237,220,177,105,221,158,5,89,150,109,90,186,46,65,58,173,
	203,150,109,216,182,32,237,166,165,27,182,44,200,183,114,199,166,5,233,182,236,88,144,101,229,190,157,11,146,110,90,183,
	99,211,146,173,235,150,46,200,180,97,199,214,101,155,118,174,75,133,9,155,190,149,43,54,45,72,178,97,225,166,21,91,119,46,
	72,179,114,94,0,166,217,1,80,7,215,37,200,162,116,211,244,0,44,75,23,36,153,70,136,114,211,206,173,59,215,37,72,162,111,
	150,32,134,165,11,18,109,89,183,100,229,150,149,51,1,128,4,192,10,64,151,32,147,186,165,91,246,108,89,185,32,237,214,101,
	11,183,46,221,176,116,203,130,108,155,22,100,88,144,112,223,206,173,91,86,110,89,144,112,203,178,101,91,134,23,224,220,
	184,117,203,148,1,172,74,23,36,219,48,46,0,211,134,5,233,70,7,194,51,64,59,51,16,186,160,178,4,233,246,173,91,144,102,222,
	1,200,65,100,155,86,204,19,64,144,97,235,142,165,251,86,174,75,144,77,195,150,29,91,214,109,24,41,128,101,219,194,121,128,
	24,246,108,90,186,116,211,206,137,6,40,71,39,64,72,68,133,9,141,214,157,59,182,44,200,178,110,211,60,2,144,10,200,18,100,
	88,182,105,198,1,210,5,225,11,144,206,6,204,184,117,76,1,206,45,75,150,133,12,192,178,109,188,1,36,0,172,75,71,9,192,56,
	68,144,65,203,186,45,27,214,45,200,180,110,166,0,61,65,40,3,232,180,174,219,177,32,207,202,13,107,55,45,217,48,9,1,211,206,
	113,6,56,22,100,219,176,108,203,206,173,27,150,108,88,144,114,209,190,117,59,182,14,66,4,212,8,17,1,209,108,64,36,155,102,
	47,40,157,109,0,204,33,152,0,202,173,75,87,78,2,192,57,111,0,203,186,5,89,22,207,62,132,2,24,101,96,237,150,157,75,39,14,
	32,155,4,128,118,211,218,45,43,87,110,24,185,0,62,80,203,214,77,59,183,237,91,178,46,65,78,173,59,23,142,42,192,180,115,
	231,96,3,124,75,199,27,96,90,151,32,150,98,112,0,60,91,7,70,130,44,16,37,16,92,33,195,44,72,132,42,247,205,100,0,155,48,
	181,1,217,8,136,128,14,57,247,45,91,182,105,199,166,165,91,151,108,90,55,118,1,202,34,184,65,49,128,97,19,37,23,78,38,72,
	144,108,203,190,101,17,31,101,33,194,19,200,58,25,0,195,142,29,91,183,237,28,200,0,54,97,203,226,145,147,176,13,24,100,
	219,55,191,1,233,166,45,147,27,49,12,47,0,227,40,31,82,93,130,216,151,53,2,64,80,162,88,180,44,166,0,233,6,100,3,2,96,14,
	0,89,68,150,32,225,176,3,44,43,199,20,192,67,148,76,152,10,19,128,194,140,51,35,192,34,224,217,58,125,0,217,190,21,3,20,
	192,111,238,216,183,109,108,4,190,225,13,136,55,109,155,57,0,18,80,16,101,24,134,76,210,16,176,108,221,179,105,76,1,116,
	9,8,28,98,10,16,174,220,178,116,211,40,136,216,23,213,53,70,59,82,83,64,54,174,0,196,65,232,136,18,12,176,44,64,122,138,
	6,72,187,97,40,5,206,5,73,182,78,18,72,150,32,235,178,165,43,55,237,216,50,48,4,26,32,224,75,163,2,203,162,8,102,26,94,
	128,114,201,36,0,72,133,251,86,174,219,176,114,203,178,96,134,0,32,80,12,70,64,55,110,20,106,35,104,5,80,135,16,0,32,18,
	0,212,80,51,51,0,76,7,5,206,32,13,177,3,194,14,142,3,128,1,0,130,1,12,96,236,208,76,19,68,49,141,212,192,182,46,65,14,173,
	43,55,172,24,76,129,114,120,3,232,196,68,243,71,144,78,137,64,184,117,229,120,76,32,16,0,8,216,106,90,187,97,72,39,224,197,
	209,13,96,33,64,33,164,221,176,114,211,196,141,200,1,9,34,36,64,72,128,19,9,91,3,248,37,167,37,195,12,129,111,34,139,46,
	168,32,205,120,4,136,3,17,21,1,101,32,218,176,115,206,67,108,213,10,27,65,22,162,58,4,87,20,55,32,164,81,240,2,166,53,163,
	10,128,55,96,157,76,0,207,33,230,10,186,249,28,147,83,1,69,32,28,96,128,101,108,0,236,14,16,12,240,80,163,157,134,6,4,205,
	25,23,65,128,64,204,21,128,175,72,54,77,88,132,208,24,187,3,200,66,134,173,147,37,193,166,130,84,4,42,56,65,69,65,78,45,
	75,6,18,130,118,60,45,35,101,17,30,106,214,233,35,16,99,211,108,89,57,123,65,2,135,190,37,155,166,89,192,66,4,174,176,120,
	106,2,109,67,122,18,2,42,99,134,0,21,28,118,0,20,17,58,137,40,146,73,43,192,109,104,151,10,40,140,46,9,153,0,233,120,69,
	136,161,192,30,131,143,6,168,9,197,74,40,116,168,105,134,97,39,18,35,8,207,10,37,136,46,1,15,53,217,124,4,152,3,0,58,164,
	12,224,80,57,44,1,210,10,240,195,144,6,226,44,64,99,72,78,5,78,9,190,83,248,65,144,4,72,3,209,53,214,110,89,180,105,192,0,
	134,25,139,56,151,110,24,92,16,188,129,164,139,144,133,48,180,192,87,112,108,26,49,0,126,19,218,97,148,1,12,194,38,163,20,
	2,3,12,249,139,115,231,20,7,152,143,65,5,70,119,2,234,152,168,16,158,103,36,37,57,20,153,11,49,78,231,34,6,241,21,22,161,
	120,220,172,56,30,9,181,160,74,99,107,192,65,149,1,9,21,130,203,214,113,69,100,10,80,109,221,185,116,114,14,100,96,97,128,
	97,4,33,104,64,199,153,46,129,118,223,178,173,75,23,142,41,0,217,0,86,82,212,1,4,64,64,10,3,144,70,26,10,138,4,190,17,199,
	65,114,225,96,93,81,174,66,49,97,11,65,151,1,249,32,80,67,94,142,252,82,64,131,177,4,161,48,68,110,4,126,24,125,101,124,
	210,112,89,72,11,19,52,0,178,198,18,132,52,50,218,144,89,0,78,185,153,241,16,197,77,66,119,141,73,9,130,34,211,252,223,
	208,13,112,32,5,20,134,118,100,59,141,1,221,134,217,137,80,128,70,203,134,74,74,185,116,240,1,240,2,208,29,208,11,19,110,
	93,182,118,211,248,17,112,130,209,37,128,199,70,108,162,1,9,218,145,168,105,180,19,201,23,246,63,136,44,133,185,141,1,12,
	127,53,60,161,184,65,193,70,161,16,194,38,12,164,25,148,97,138,100,0,80,76,6,0,192,113,7,7,34,198,81,25,190,17,54,97,50,
	65,165,194,89,136,112,173,131,111,196,79,156,146,56,227,116,160,67,164,21,8,28,210,53,37,79,135,14,200,177,111,229,196,87,
	88,37,2,64,40,152,12,87,24,216,139,100,234,0,13,131,117,217,100,186,70,90,193,178,102,205,88,138,152,47,56,199,6,0,228,9,
	198,28,153,32,182,6,80,207,40,133,192,130,137,59,1,224,13,239,35,159,81,246,69,45,68,192,58,50,11,21,170,220,176,101,231,
	120,3,40,168,19,54,98,50,128,17,136,83,77,201,85,110,128,40,132,160,166,129,48,197,83,68,14,79,60,8,25,56,238,134,149,59,
	182,78,25,0,46,11,155,0,199,168,0,136,5,0,19,96,208,136,93,192,58,241,81,152,57,208,149,113,175,113,3,225,81,66,151,1,3,
	24,204,44,182,76,160,142,161,130,67,40,97,157,5,199,29,233,26,115,99,154,1,182,177,36,144,10,167,242,12,31,149,46,92,94,
	83,219,52,9,140,50,35,160,172,68,170,129,30,19,108,119,254,137,92,115,153,27,3,34,133,18,132,96,132,165,26,225,60,254,110,
	40,16,227,161,82,27,66,181,66,99,10,118,32,91,8,76,218,116,156,9,28,74,54,237,153,218,128,115,16,11,56,30,217,33,99,193,
	133,80,130,72,1,11,65,1,89,220,160,76,210,224,27,65,46,182,109,216,51,2,8,8,89,215,117,193,101,1,8,198,162,22,56,17,245,
	116,124,167,188,28,193,70,97,23,224,131,65,127,142,98,25,227,116,234,43,40,213,64,24,195,155,8,170,8,87,16,62,47,93,65,
	204,109,132,76,185,192,147,70,133,61,130,200,137,172,27,127,103,116,34,204,70,120,173,208,99,2,97,12,163,139,60,53,31,96,
	162,66,240,216,224,27,97,88,35,29,128,221,133,35,29,205,60,118,7,162,209,13,136,6,159,66,11,2,244,145,116,203,88,93,1,224,
	96,217,165,58,132,133,0,221,12,226,44,186,198,52,77,241,69,240,86,38,232,136,139,138,33,0,1,81,236,143,32,88,136,70,67,212,
	137,247,199,236,149,155,8,16,70,74,234,72,70,248,162,161,21,129,37,105,93,57,63,18,208,128,236,10,248,33,161,77,197,103,
	207,177,148,247,37,124,18,194,74,97,66,64,163,160,73,163,128,201,33,156,137,21,210,67,218,48,84,129,64,35,64,18,161,223,
	206,142,46,49,8,51,29,250,0,32,89,88,179,48,91,66,105,192,32,24,129,59,76,22,211,136,0,24,72,115,3,231,23,160,75,144,69,
	233,8,88,64,216,98,0,64,119,66,39,148,170,13,197,3,20,159,72,140,160,134,176,62,70,214,20,48,128,149,26,225,0,170,6,209,
	11,1,210,2,65,36,235,12,224,143,7,8,134,3,176,38,184,16,241,29,81,25,132,23,64,44,136,75,23,126,66,172,111,52,8,144,174,
	16,37,128,111,74,159,19,10,52,137,106,196,37,112,195,248,1,129,16,197,51,146,71,152,163,66,134,32,102,79,132,6,128,180,
	115,113,164,177,206,67,149,196,19,103,247,42,95,12,14,88,1,51,37,225,191,135,228,30,156,154,46,137,128,97,33,64,144,38,
	164,68,74,207,21,56,61,95,53,207,202,13,107,55,13,130,135,22,15,203,73,80,101,184,146,49,15,241,108,37,106,149,39,145,179,
	44,22,60,124,0,180,97,201,143,36,62,164,121,132,28,4,22,149,109,93, 00} // 2296
