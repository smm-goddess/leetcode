package find_longest_chain

import (
	"fmt"
	"testing"
)

func TestFindLongestChain(t *testing.T) {
	//data := [][]int{
	//	{3, 4},
	//	{2, 3},
	//	{1, 2},
	//}
	//if findLongestChain(data) != 2 {
	//	t.Error("error")
	//}
	data := [][]int{
		{-324, 606}, {939, 980}, {-829, -399}, {789, 819}, {-26, 468}, {-816, -506}, {-157, 417}, {670, 740}, {708, 832}, {-597, 209},
		{-498, 933}, {-298, 308}, {826, 925}, {914, 961}, {-984, 994}, {-306, 438}, {244, 721}, {-726, -87}, {-499, -418}, {590, 997},
		{-357, 387}, {795, 926}, {389, 731}, {-196, -124}, {966, 983}, {329, 938}, {339, 897}, {-736, 670}, {-214, -133}, {-379, 950},
		{-265, 575}, {-93, 43}, {-581, -478}, {-335, 182}, {214, 463}, {197, 250}, {679, 903}, {220, 826}, {246, 493}, {350, 989},
		{82, 763}, {870, 943}, {475, 654}, {618, 978}, {98, 588}, {-559, -85}, {763, 786}, {-323, -291}, {-628, 649}, {146, 968},
		{-269, 140}, {-142, 306}, {-130, 880}, {344, 636}, {335, 809}, {882, 960}, {895, 971}, {660, 905}, {-659, -61}, {936, 953},
		{512, 684}, {2, 295}, {-32, 116}, {366, 980}, {293, 971}, {-313, 432}, {637, 660}, {846, 979}, {327, 552}, {873, 947},
		{-813, 277}, {-336, 831}, {-189, 386}, {-231, 619}, {60, 544}, {161, 869}, {-693, 470}, {482, 546}, {-922, 720}, {-282, 346},
		{439, 465}, {959, 979}, {-169, -47}, {-320, 412}, {90, 158}, {797, 892}, {-857, 750}, {499, 965}, {741, 896}, {116, 414},
		{106, 646}, {747, 815}, {-895, -659}, {321, 568}, {-525, 373}, {849, 909}, {228, 903}, {-469, 42}, {-237, 135}, {-486, 987},
		{-560, -298}, {-495, 952}, {995, 997}, {-555, 723}, {124, 574}, {581, 653}, {-122, 234}, {684, 929}, {-60, 826}, {476, 832},
		{201, 488}, {65, 538}, {443, 450}, {609, 757}, {187, 965}, {44, 736}, {570, 815}, {591, 782}, {-856, -146}, {-529, -476},
		{-561, -424}, {213, 601}, {-595, 422}, {632, 910}, {-689, 410}, {723, 790}, {686, 827}, {115, 921}, {54, 177}, {173, 730},
		{522, 847}, {-510, -447}, {10, 31}, {852, 887}, {-587, 185}, {202, 547}, {-16, 203}, {-580, 845}, {-954, -943}, {-802, -306},
		{656, 917}, {562, 829}, {90, 622}, {448, 575}, {74, 445}, {705, 994}, {-577, -199}, {-734, -506}, {-73, 263}, {-196, 324},
		{-776, -196}, {938, 975}, {-468, 337}, {809, 960}, {895, 935}, {-316, 931}, {-564, -221}, {453, 863}, {159, 872}, {-554, 141},
		{256, 496}, {937, 980}, {875, 878}, {-290, 351}, {362, 461}, {-340, 325}, {-116, -61}, {-799, -599}, {459, 957}, {-644, -481},
		{430, 811}, {304, 838}, {-495, -277}, {-550, 772}, {762, 938}, {-836, 818}, {353, 725}, {-120, 440}, {-355, 244}, {-557, 562},
		{-305, -158}, {370, 883}, {661, 879}, {-276, 651}, {807, 899}, {-268, -181}, {-436, 167}, {-786, 400}, {3, 755}, {327, 677},
		{266, 952}, {910, 996}, {-601, -559}, {-256, -161}, {-461, -283}, {-788, 593}, {-880, 12}, {286, 586}, {252, 615}, {456, 604},
		{162, 682}, {848, 931}, {434, 967}, {-326, 17}, {-266, 262}, {-961, -110}, {487, 889}, {-850, 715}, {-998, 454}, {127, 596},
		{633, 844}, {148, 212}, {-286, 537}, {-362, 67}, {323, 732}, {858, 929}, {-910, -844}, {309, 691}, {-862, 493}, {925, 943},
		{476, 882}, {421, 954}, {-612, -360}, {750, 956}, {-534, 725}, {-759, -508}, {620, 902}, {-661, 471}, {-403, 619}, {169, 488},
		{-125, 142}, {-706, -243}, {-876, -681}, {-161, 235}, {-969, 522}, {-741, -408}, {893, 911}, {990, 999}, {-629, 853}, {477, 909},
		{818, 969}, {463, 697}, {433, 829}, {136, 435}, {-329, 944}, {-710, -432}, {-286, 209}, {-451, 539}, {-146, -121}, {522, 769},
		{247, 296}, {-704, 140}, {-485, 87}, {209, 551}, {287, 466}, {606, 922}, {-199, 840}, {-822, -583}, {-407, -271}, {392, 908},
		{551, 681}, {708, 973}, {-734, -241}, {959, 999}, {-831, -50}, {420, 434}, {262, 731}, {880, 988}, {216, 945}, {-234, 830},
		{960, 975}, {-431, -40}, {-305, 17}, {-859, 211}, {-427, -227}, {-518, 498}, {-673, 628}, {343, 533}, {-658, 973}, {-472, 71},
		{697, 794}, {497, 655}, {675, 932}, {-624, -198}, {135, 296}, {62, 405}, {443, 638}, {-337, -328}, {-662, -358}, {-648, 753},
		{372, 557}, {-716, 644}, {857, 886}, {777, 956}, {-133, -40}, {742, 935}, {122, 145}, {-401, 35}, {-540, -368}, {-816, 78},
		{-697, -505}, {418, 514}, {-561, 358}, {-184, 172}, {-304, -185}, {856, 947}, {949, 997}, {-397, 817}, {258, 984}, {912, 992},
		{-996, -975}, {-435, -363}, {880, 912}, {512, 739}, {-925, -155}, {968, 987}, {709, 899}, {487, 948}, {229, 504}, {-372, 200},
		{463, 890}, {-847, 348}, {347, 531}, {-313, 861}, {-383, 647}, {727, 774}, {714, 793}, {-143, 672}, {-669, 511}, {-50, 936},
		{364, 408}, {-653, -112}, {-394, 339}, {963, 978}, {-746, 399}, {19, 711}, {-772, -166}, {668, 947}, {-542, 878}, {-426, -39},
		{-491, 848}, {927, 986}, {685, 757}, {-113, 463}, {951, 968}, {-520, -22}, {-144, 456}, {-983, 969}, {-153, 293}, {-938, -792},
		{-447, -256}, {69, 655}, {-589, -364}, {315, 860}, {790, 885}, {-60, 158}, {441, 745}, {394, 917}, {-583, 836}, {271, 617},
		{432, 887}, {253, 985}, {-383, -272}, {446, 868}, {-830, -151}, {-250, 95}, {-741, -412}, {-569, -66}, {-798, 965}, {831, 911},
		{382, 522}, {765, 795}, {-80, 418}, {-725, -313}, {862, 911}, {765, 814}, {-15, 728}, {735, 782}, {685, 986}, {289, 592},
		{372, 723}, {-70, 233}, {463, 687}, {-360, 520}, {781, 864}, {842, 969}, {-754, 906}, {808, 886}, {312, 970}, {-803, 753},
		{-514, -460}, {-706, 569}, {246, 315}, {10, 869}, {-210, 45}, {-304, 163}, {-638, 6}, {16, 519}, {-15, 290}, {748, 953},
		{760, 872}, {513, 514}, {517, 567}, {64, 940}, {923, 925}, {-882, 466}, {-92, 249}, {31, 493}, {-308, 311}, {925, 962},
		{-921, -590}, {-342, 945}, {539, 750}, {270, 966}, {-441, 788}, {633, 999}, {-99, 268}, {-88, 398}, {-114, 223}, {502, 575},
		{-513, -305}, {-404, 423}, {-593, -108}, {964, 993}, {281, 943}, {-575, -178}, {-802, 385}, {569, 619}, {774, 804},
		{444, 446}, {-381, 534}, {-529, -15}, {-653, -157}, {-534, -313}, {-782, 388}, {-655, 795}, {994, 996}, {372, 842}, {239, 681},
		{884, 998}, {-43, 163}, {403, 848}, {-368, -126}, {260, 549}, {234, 593}, {-555, 150}, {-576, -378}, {214, 343}, {-617, 699},
		{236, 612}, {464, 646}, {351, 569}, {-911, -585}, {-738, 834}, {634, 680}, {401, 722}, {996, 999}, {654, 827}, {815, 956},
		{765, 818}, {-738, 680}, {570, 605}, {-795, 680}, {257, 580}, {236, 799}, {-399, 515}, {-784, -572}, {-963, 813}, {227, 470},
		{2, 297}, {-487, -50}, {-788, -499}, {634, 948}, {-87, 370}, {836, 864}, {-514, -97}, {762, 779}, {-197, 501}, {491, 508},
		{787, 878}, {-200, 761}, {122, 458}, {363, 372}, {-988, 464}, {360, 546}, {-312, -204}, {959, 999}, {-201, 974}, {-326, 245},
		{245, 640}, {-215, 583}, {14, 480}, {-916, 548}, {-228, 553}, {-401, 361}, {290, 463}, {799, 996}, {-796, -569}, {-990, 290},
		{593, 690}, {605, 819}, {610, 873}, {534, 895}, {-312, 572}, {-648, 331}, {-540, 557}, {-784, 132}, {-646, 150}, {84, 721},
		{-949, 159}, {-703, -637}, {-249, -54}, {-572, 722}, {776, 986}, {228, 862}, {843, 928}, {-488, 267}, {172, 515}, {-570, -148},
		{-649, -184}, {-332, 746}, {881, 969}, {-61, 199}, {404, 934}, {750, 825}, {-793, -156}, {-252, 346}, {-1000, 643}, {191, 522},
		{-873, 447}, {-730, -283}, {744, 967}, {204, 251}, {-163, 773}, {92, 992}, {869, 984}, {929, 995}, {606, 787}, {181, 453},
		{194, 571}, {-745, 251}, {-342, 21}, {497, 986}, {-935, -573}, {754, 936}, {-352, -306}, {839, 968}, {185, 558}, {-107, -50},
		{708, 879}, {515, 873}, {-195, 989}, {-727, 676}, {142, 946}, {-952, 468}, {-705, -405}, {-630, -181}, {-692, 814}, {-337, 232},
		{558, 776}, {818, 861}, {-68, 3}, {576, 883}, {851, 933}, {821, 892}, {-694, 209}, {642, 762}, {-844, -395}, {504, 610},
		{-246, 214}, {798, 865}, {832, 883}, {-212, 96}, {-79, 659}, {690, 705}, {119, 473}, {-724, 24}, {127, 866}, {-512, 692},
		{-706, 412}, {598, 812}, {892, 977}, {501, 897}, {95, 282}, {128, 318}, {73, 263}, {-290, 510}, {-102, 173}, {541, 837},
		{-612, 285}, {-397, 86}, {-792, 820}, {757, 890}, {-392, 4}, {171, 531}, {-41, 657}, {-666, 23}, {460, 816}, {-597, 18},
		{250, 839}, {39, 666}, {560, 664}, {663, 883}, {-449, 8}, {123, 963}, {-31, 410}, {688, 855}, {190, 809}, {-428, -276},
		{849, 869}, {275, 492}, {-63, 40}, {-888, 435}, {-477, 415}, {537, 753}, {244, 986}, {-291, -106}, {-473, 95}, {-382, -203},
		{652, 928}, {-855, 526}, {665, 771}, {-481, 839}, {-873, 56}, {-64, 642}, {-931, 269}, {457, 580}, {383, 833}, {-7, 600},
		{890, 980}, {768, 843}, {-184, 630}, {237, 548}, {-254, -208}, {-357, 173}, {-333, -213}, {884, 989}, {-700, -644},
		{-562, 559}, {290, 1000}, {-686, 443}, {585, 668}, {-432, -423}, {852, 932}, {210, 690}, {-844, -495}, {-341, 344},
		{-319, 249}, {-273, 683}, {-87, 170}, {-442, 64}, {778, 873}, {33, 641}, {319, 906}, {963, 993}, {763, 771}, {734, 863},
		{-465, -2}, {-718, 601}, {-824, 985}, {357, 491}, {476, 927}, {-921, -414}, {-740, 237}, {15, 449}, {796, 807}, {400, 766},
		{-479, -70}, {-732, 459}, {-460, 44}, {-513, -44}, {-158, 729}, {547, 590}, {-572, 468}, {-499, 280}, {592, 626},
		{430, 550}, {-437, 891}, {-158, 370}, {435, 474}, {-345, -17}, {-899, -273}, {329, 627}, {-493, 957}, {-774, -354},
		{376, 674}, {-259, -28}, {433, 713}, {871, 872}, {419, 656}, {426, 838}, {-469, -194}, {-890, 708}, {-202, 351}, {435, 770},
		{916, 971}, {-282, 652}, {100, 164}, {-313, 50}, {-90, -34}, {776, 871}, {166, 239}, {275, 689}, {-660, 458}, {495, 806},
		{-846, 265}, {-623, -120}, {461, 894}, {874, 962}, {-160, 388}, {-343, -321}, {-968, 3}, {-714, -123}, {-873, 815},
		{217, 775}, {-614, 134}, {482, 582}, {-681, -284}, {700, 786}, {323, 584}, {-581, 905}, {560, 972}, {888, 915}, {-343, 986},
		{229, 278}, {563, 797}, {663, 953}, {-823, 512}, {111, 540}, {188, 484}, {-510, -85}, {-815, -498}, {260, 979},
		{857, 1000}, {-192, 54}, {-13, 439}, {-817, 709}, {345, 575}, {-795, 230}, {-396, 414}, {-504, -467}, {208, 556}, {-36, 160},
		{-565, 512}, {942, 980}, {736, 950}, {369, 438}, {593, 970}, {-188, 514}, {-419, 528}, {771, 891}, {-98, 614}, {148, 909},
		{31, 85}, {374, 957}, {298, 415}, {-128, 888}, {195, 554}, {-423, -264}, {-423, 891}, {-48, 532}, {122, 926}, {538, 849},
		{345, 695}, {513, 718}, {-408, 466}, {703, 802}, {-448, 36}, {662, 714}, {40, 719}, {-710, 613}, {258, 487}, {-929, 86},
		{716, 720}, {-636, 80}, {-237, 518}, {552, 582}, {-729, -644}, {-900, -736}, {885, 911}, {268, 848}, {16, 86}, {141, 478},
		{-742, -646}, {571, 846}, {-879, -366}, {-4, 239}, {434, 597}, {713, 751}, {802, 940}, {-390, -139}, {162, 413},
		{742, 796}, {243, 624}, {-872, 978}, {-369, 659}, {-255, 243}, {45, 132}, {-386, 335}, {669, 981}, {-800, 70}, {-829, -494},
		{-243, 560}, {-620, 454}, {-184, 75}, {471, 578}, {-917, 26}, {-799, -545}, {711, 839}, {-385, 585}, {-400, 685},
		{-452, 959}, {396, 608}, {-517, -90}, {985, 1000}, {513, 612}, {710, 968}, {104, 399}, {436, 485}, {-629, 913}, {-364, 170},
		{-139, 947}, {265, 452}, {306, 936}, {-90, 593}, {769, 773}, {-938, -305}, {136, 830}, {89, 121}, {769, 839}, {451, 736},
		{-151, -61}, {421, 711}, {303, 472}, {42, 910}, {143, 245}, {-958, 706}, {-447, -100}, {-679, 548}, {146, 345}, {105, 224},
		{699, 827}, {903, 922}, {241, 684}, {414, 611}, {-880, 657}, {967, 989}, {-783, -333}, {936, 972}, {225, 329}, {75, 897},
		{252, 368}, {-404, 859}, {-668, 462}, {239, 992}, {-129, 442}, {-247, 555}, {-611, 764}, {-971, 833}, {517, 826}, {-618, 56},
		{303, 894}, {160, 797}, {-821, -460}, {-119, 760}, {-352, -137}, {-473, -435}, {-117, -3}, {888, 936}, {235, 392}, {-216, -156}, {783, 950},
	}
	if findLongestChain(data) != 38 {
		fmt.Println("38 error")
	}
}
