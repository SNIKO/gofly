package airports

const csvdb = `
1;Goroka;Goroka;Papua New Guinea;GKA;AYGA;-6.081689;145.391881;5282;10;U;Pacific/Port_Moresby
2;Madang;Madang;Papua New Guinea;MAG;AYMD;-5.207083;145.7887;20;10;U;Pacific/Port_Moresby
3;Mount Hagen;Mount Hagen;Papua New Guinea;HGU;AYMH;-5.826789;144.295861;5388;10;U;Pacific/Port_Moresby
4;Nadzab;Nadzab;Papua New Guinea;LAE;AYNZ;-6.569828;146.726242;239;10;U;Pacific/Port_Moresby
5;Port Moresby Jacksons Intl;Port Moresby;Papua New Guinea;POM;AYPY;-9.443383;147.22005;146;10;U;Pacific/Port_Moresby
6;Wewak Intl;Wewak;Papua New Guinea;WWK;AYWK;-3.583828;143.669186;19;10;U;Pacific/Port_Moresby
7;Narsarsuaq;Narssarssuaq;Greenland;UAK;BGBW;61.160517;-45.425978;112;-3;E;America/Godthab
8;Nuuk;Godthaab;Greenland;GOH;BGGH;64.190922;-51.678064;283;-3;E;America/Godthab
9;Sondre Stromfjord;Sondrestrom;Greenland;SFJ;BGSF;67.016969;-50.689325;165;-3;E;America/Godthab
10;Thule Air Base;Thule;Greenland;THU;BGTL;76.531203;-68.703161;251;-4;E;America/Thule
11;Akureyri;Akureyri;Iceland;AEY;BIAR;65.659994;-18.072703;6;0;N;Atlantic/Reykjavik
12;Egilsstadir;Egilsstadir;Iceland;EGS;BIEG;65.283333;-14.401389;76;0;N;Atlantic/Reykjavik
13;Hornafjordur;Hofn;Iceland;HFN;BIHN;64.295556;-15.227222;24;0;N;Atlantic/Reykjavik
14;Husavik;Husavik;Iceland;HZK;BIHU;65.952328;-17.425978;48;0;N;Atlantic/Reykjavik
15;Isafjordur;Isafjordur;Iceland;IFJ;BIIS;66.058056;-23.135278;8;0;N;Atlantic/Reykjavik
16;Keflavik International Airport;Keflavik;Iceland;KEF;BIKF;63.985;-22.605556;171;0;N;Atlantic/Reykjavik
17;Patreksfjordur;Patreksfjordur;Iceland;PFJ;BIPA;65.555833;-23.965;11;0;N;Atlantic/Reykjavik
18;Reykjavik;Reykjavik;Iceland;RKV;BIRK;64.13;-21.940556;48;0;N;Atlantic/Reykjavik
19;Siglufjordur;Siglufjordur;Iceland;SIJ;BISI;66.133333;-18.916667;10;0;N;Atlantic/Reykjavik
20;Vestmannaeyjar;Vestmannaeyjar;Iceland;VEY;BIVM;63.424303;-20.278875;326;0;N;Atlantic/Reykjavik
21;Sault Ste Marie;Sault Sainte Marie;Canada;YAM;CYAM;46.485001;-84.509445;630;-5;A;America/Toronto
22;Winnipeg St Andrews;Winnipeg;Canada;YAV;CYAV;50.056389;-97.0325;760;-6;A;America/Winnipeg
23;Shearwater;Halifax;Canada;YAW;CYAW;44.639721;-63.499444;167;-4;A;America/Halifax
24;St Anthony;St. Anthony;Canada;YAY;CYAY;51.391944;-56.083056;108;-3.5;A;America/St_Johns
25;Tofino;Tofino;Canada;YAZ;CYAZ;49.082222;-125.7725;80;-8;A;America/Vancouver
26;Kugaaruk;Pelly Bay;Canada;YBB;CYBB;68.534444;-89.808056;56;-7;A;America/Edmonton
27;Baie Comeau;Baie Comeau;Canada;YBC;CYBC;49.1325;-68.204444;71;-5;A;America/Toronto
28;Bagotville;Bagotville;Canada;YBG;CYBG;48.330555;-70.996391;522;-5;A;America/Toronto
29;Baker Lake;Baker Lake;Canada;YBK;CYBK;64.298889;-96.077778;59;-6;A;America/Winnipeg
30;Campbell River;Campbell River;Canada;YBL;CYBL;49.950832;-125.270833;346;-8;A;America/Vancouver
31;Brandon Muni;Brandon;Canada;YBR;CYBR;49.91;-99.951944;1343;-6;A;America/Winnipeg
32;Cambridge Bay;Cambridge Bay;Canada;YCB;CYCB;69.108055;-105.138333;90;-7;A;America/Edmonton
33;Nanaimo;Nanaimo;Canada;YCD;CYCD;49.052333;-123.870167;93;-8;A;America/Vancouver
34;Castlegar;Castlegar;Canada;YCG;CYCG;49.296389;-117.6325;1624;-8;A;America/Vancouver
35;Miramichi;Chatham;Canada;YCH;CYCH;47.007778;-65.449167;108;-4;A;America/Halifax
36;Charlo;Charlo;Canada;YCL;CYCL;47.990833;-66.330278;132;-4;A;America/Halifax
37;Kugluktuk;Coppermine;Canada;YCO;CYCO;67.816667;-115.143889;74;-7;A;America/Edmonton
38;Coronation;Coronation;Canada;YCT;CYCT;52.075001;-111.445278;2595;-7;A;America/Edmonton
39;Chilliwack;Chilliwack;Canada;YCW;CYCW;49.152779;-121.93889;32;-8;A;America/Vancouver
40;Clyde River;Clyde River;Canada;YCY;CYCY;70.486111;-68.516667;87;-5;A;America/Toronto
41;Fairmont Hot Springs;Coral Harbour;Canada;YZS;CYCZ;64.193333;-83.359444;2661;-5;A;America/Coral_Harbour
42;Dawson City;Dawson;Canada;YDA;CYDA;64.043056;-139.127778;1215;-8;A;America/Vancouver
43;Burwash;Burwash;Canada;YDB;CYDB;61.371111;-139.040556;2647;-8;A;America/Vancouver
44;Princeton;Princeton;Canada;YDC;CYDC;49.468056;-120.511389;2298;-8;A;America/Vancouver
45;Deer Lake;Deer Lake;Canada;YDF;CYDF;49.210833;-57.391388;72;-3.5;A;America/St_Johns
46;Dease Lake;Dease Lake;Canada;YDL;CYDL;58.422222;-130.032222;2600;-8;A;America/Vancouver
47;Dauphin Barker;Dauphin;Canada;YDN;CYDN;51.100834;-100.0525;999;-6;A;America/Winnipeg
48;Dawson Creek;Dawson Creek;Canada;YDQ;CYDQ;55.742333;-120.183;2148;-7;A;America/Dawson_Creek
49;Edmonton Intl;Edmonton;Canada;YEG;CYEG;53.309723;-113.579722;2373;-7;A;America/Edmonton
50;Arviat;Eskimo Point;Canada;YEK;CYEK;61.094166;-94.070833;32;-6;A;America/Winnipeg
51;Estevan;Estevan;Canada;YEN;CYEN;49.210278;-102.965833;1905;-6;N;America/Regina
52;Edson;Edson;Canada;YET;CYET;53.578888;-116.465;3041;-7;A;America/Edmonton
53;Eureka;Eureka;Canada;YEU;CYEU;79.994722;-85.814167;256;-6;A;America/Winnipeg
54;Inuvik Mike Zubko;Inuvik;Canada;YEV;CYEV;68.304167;-133.482778;224;-7;A;America/Edmonton
55;Iqaluit;Iqaluit;Canada;YFB;CYFB;63.75639;-68.555832;110;-5;A;America/Toronto
56;Fredericton;Fredericton;Canada;YFC;CYFC;45.868889;-66.537222;68;-4;A;America/Halifax
57;Forestville;Forestville;Canada;;CYFE;48.746111;-69.097222;293;-5;A;America/Toronto
58;Flin Flon;Flin Flon;Canada;YFO;CYFO;54.678055;-101.681667;997;-6;A;America/Winnipeg
59;Fort Resolution;Fort Resolution;Canada;YFR;CYFR;61.180832;-113.689722;526;-7;A;America/Edmonton
60;Fort Simpson;Fort Simpson;Canada;YFS;CYFS;61.760153;-121.236525;555;-7;A;America/Edmonton
61;Kingston;Kingston;Canada;YGK;CYGK;44.225277;-76.596944;305;-5;A;America/Toronto
62;La Grande Riviere;La Grande Riviere;Canada;YGL;CYGL;53.625278;-77.704167;639;-5;A;America/Toronto
63;Gaspe;Gaspe;Canada;YGP;CYGP;48.775278;-64.478611;108;-5;A;America/Toronto
64;Geraldton Greenstone Regional;Geraldton;Canada;YGQ;CYGQ;49.778332;-86.939445;1144;-5;A;America/Toronto
65;Iles De La Madeleine;Iles De La Madeleine;Canada;YGR;CYGR;47.424721;-61.778056;35;-5;A;America/Toronto
66;Hudson Bay;Hudson Bay;Canada;YHB;CYHB;52.816666;-102.31139;1175;-6;N;America/Regina
67;Dryden Rgnl;Dryden;Canada;YHD;CYHD;49.831667;-92.744167;1354;-6;A;America/Winnipeg
68;Ulukhaktok Holman;Holman Island;Canada;YHI;CYHI;70.762778;-117.806111;117;-7;A;America/Edmonton
69;Gjoa Haven;Gjoa Haven;Canada;YHK;CYHK;68.635556;-95.849722;152;-7;A;America/Edmonton
70;Hamilton;Hamilton;Canada;YHM;CYHM;43.173611;-79.935;780;-5;A;America/Toronto
71;St Hubert;Montreal;Canada;YHU;CYHU;45.5175;-73.416944;90;-5;A;America/Toronto
72;Hay River;Hay River;Canada;YHY;CYHY;60.839722;-115.782778;543;-7;A;America/Edmonton
73;Halifax Intl;Halifax;Canada;YHZ;CYHZ;44.880833;-63.50861;477;-4;A;America/Halifax
74;Atikokan Muni;Atikokan;Canada;YIB;CYIB;48.773888;-91.638611;1408;-5;A;America/Coral_Harbour
75;Pond Inlet;Pond Inlet;Canada;YIO;CYIO;72.683334;-77.966667;181;-5;A;America/Toronto
76;St Jean;St. Jean;Canada;YJN;CYJN;45.294445;-73.281111;136;-5;A;America/Toronto
77;Stephenville;Stephenville;Canada;YJT;CYJT;48.544167;-58.549999;80;-3.5;A;America/St_Johns
78;Kamloops;Kamloops;Canada;YKA;CYKA;50.702222;-120.444444;1133;-8;A;America/Vancouver
79;Waterloo;Waterloo;Canada;YKF;CYKF;43.460833;-80.378611;1054;-5;A;America/Toronto
80;Schefferville;Schefferville;Canada;YKL;CYKL;54.805278;-66.805278;1709;-5;A;America/Toronto
81;Kindersley;Kindersley;Canada;YKY;CYKY;51.5175;-109.180833;2277;-6;N;America/Regina
82;Buttonville Muni;Toronto;Canada;YKZ;CYKZ;43.862221;-79.37;650;-5;A;America/Toronto
83;Chapleau;Chapleau;Canada;YLD;CYLD;47.82;-83.346667;1470;-5;A;America/Toronto
84;Meadow Lake;Meadow Lake;Canada;YLJ;CYLJ;54.125278;-108.522778;1576;-6;N;America/Regina
85;Lloydminster;Lloydminster;Canada;YLL;CYLL;53.309166;-110.0725;2193;-7;A;America/Edmonton
86;Alert;Alert;Canada;YLT;CYLT;82.517778;-62.280556;100;-5;A;America/Toronto
87;Kelowna;Kelowna;Canada;YLW;CYLW;49.956112;-119.377778;1409;-8;A;America/Vancouver
88;Mayo;Mayo;Canada;YMA;CYMA;63.616389;-135.868333;1653;-8;A;America/Vancouver
89;Moose Jaw Air Vice Marshal C M Mcewen;Moose Jaw;Canada;YMJ;CYMJ;50.330278;-105.559167;1892;-6;N;America/Regina
90;Fort Mcmurray;Fort Mcmurray;Canada;YMM;CYMM;56.653333;-111.221944;1211;-7;A;America/Edmonton
91;Moosonee;Moosonee;Canada;YMO;CYMO;51.291111;-80.607778;30;-5;A;America/Toronto
92;Maniwaki;Maniwaki;Canada;YMW;CYMW;46.272778;-75.990556;656;-5;A;America/Toronto
93;Montreal Intl Mirabel;Montreal;Canada;YMX;CYMX;45.681944;-74.005278;270;-5;A;America/Toronto
94;Natashquan;Natashquan;Canada;YNA;CYNA;50.19;-61.789167;39;-5;A;America/Toronto
95;Gatineau;Gatineau;Canada;YND;CYND;45.521694;-75.563589;211;-5;A;America/Toronto
96;Matagami;Matagami;Canada;YNM;CYNM;49.761667;-77.802778;918;-5;A;America/Toronto
97;Old Crow;Old Crow;Canada;YOC;CYOC;67.570556;-139.839167;824;-8;A;America/Vancouver
98;Cold Lake;Cold Lake;Canada;YOD;CYOD;54.404999;-110.279444;1775;-7;A;America/Edmonton
99;High Level;High Level;Canada;YOJ;CYOJ;58.621389;-117.164722;1110;-7;A;America/Edmonton
100;Ottawa Macdonald Cartier Intl;Ottawa;Canada;YOW;CYOW;45.3225;-75.669167;374;-5;A;America/Toronto
101;Prince Albert Glass Field;Prince Albert;Canada;YPA;CYPA;53.214167;-105.672778;1405;-6;N;America/Regina
102;Peace River;Peace River;Canada;YPE;CYPE;56.226944;-117.447222;1873;-7;A;America/Edmonton
103;Southport;Portage-la-prairie;Canada;YPG;CYPG;49.903056;-98.273889;885;-6;A;America/Winnipeg
104;Pitt Meadows;Pitt Meadows;Canada;;CYPK;49.21611;-122.71;11;-8;A;America/Vancouver
105;Pickle Lake;Pickle Lake;Canada;YPL;CYPL;51.446388;-90.214167;1267;-5;A;America/Coral_Harbour
106;Port Menier;Port Menier;Canada;YPN;CYPN;49.836389;-64.288611;167;-5;A;America/Toronto
107;Peterborough;Peterborough;Canada;YPQ;CYPQ;44.23;-78.363333;628;-5;A;America/Toronto
108;Prince Rupert;Prince Pupert;Canada;YPR;CYPR;54.28611;-130.444722;116;-8;A;America/Vancouver
109;Fort Chipewyan;Fort Chipewyan;Canada;YPY;CYPY;58.767223;-111.117222;761;-7;A;America/Edmonton
110;Muskoka;Muskoka;Canada;YQA;CYQA;44.974722;-79.303333;925;-5;A;America/Toronto
111;Quebec Jean Lesage Intl;Quebec;Canada;YQB;CYQB;46.791111;-71.393333;244;-5;A;America/Toronto
112;Red Deer Regional;Red Deer Industrial;Canada;YQF;CYQF;52.182222;-113.894444;2968;-7;A;America/Edmonton
113;Windsor;Windsor;Canada;YQG;CYQG;42.275556;-82.955556;622;-5;A;America/Toronto
114;Watson Lake;Watson Lake;Canada;YQH;CYQH;60.116389;-128.8225;2255;-8;A;America/Vancouver
115;Kenora;Kenora;Canada;YQK;CYQK;49.788334;-94.363056;1332;-6;A;America/Winnipeg
116;Lethbridge;Lethbridge;Canada;YQL;CYQL;49.630278;-112.799722;3047;-7;A;America/Edmonton
117;Greater Moncton Intl;Moncton;Canada;YQM;CYQM;46.112221;-64.678611;232;-4;A;America/Halifax
119;Comox;Comox;Canada;YQQ;CYQQ;49.710833;-124.886667;84;-8;A;America/Vancouver
120;Regina Intl;Regina;Canada;YQR;CYQR;50.431944;-104.665833;1894;-6;N;America/Regina
121;Thunder Bay;Thunder Bay;Canada;YQT;CYQT;48.371944;-89.323889;653;-5;A;America/Toronto
122;Grande Prairie;Grande Prairie;Canada;YQU;CYQU;55.179722;-118.885;2195;-7;A;America/Edmonton
123;Yorkton Muni;Yorkton;Canada;YQV;CYQV;51.264721;-102.461667;1635;-6;N;America/Regina
124;North Battleford;North Battleford;Canada;YQW;CYQW;52.769167;-108.24361;1799;-6;N;America/Regina
125;Gander Intl;Gander;Canada;YQX;CYQX;48.936944;-54.568056;496;-3.5;A;America/St_Johns
126;Sydney;Sydney;Canada;YQY;CYQY;46.161388;-60.047779;203;-4;A;America/Halifax
127;Quesnel;Quesnel;Canada;YQZ;CYQZ;53.026112;-122.510278;1789;-8;A;America/Vancouver
128;Resolute Bay;Resolute;Canada;YRB;CYRB;74.716944;-94.969444;215;-6;A;America/Winnipeg
129;Riviere Du Loup;Riviere Du Loup;Canada;YRI;CYRI;47.764444;-69.584722;427;-5;A;America/Toronto
130;Roberval;Roberval;Canada;YRJ;CYRJ;48.52;-72.265556;586;-5;A;America/Toronto
131;Rocky Mountain House;Rocky Mountain House;Canada;YRM;CYRM;52.429722;-114.904167;3244;-7;A;America/Edmonton
132;Rankin Inlet;Rankin Inlet;Canada;YRT;CYRT;62.81139;-92.115833;94;-6;A;America/Winnipeg
133;Sudbury;Sudbury;Canada;YSB;CYSB;46.625;-80.798889;1141;-5;A;America/Toronto
134;Sherbrooke;Sherbrooke;Canada;YSC;CYSC;45.438611;-71.691389;792;-5;A;America/Toronto
135;Saint John;St. John;Canada;YSJ;CYSJ;45.316111;-65.890278;357;-4;A;America/Halifax
136;Fort Smith;Fort Smith;Canada;YSM;CYSM;60.020278;-111.961944;671;-7;A;America/Edmonton
137;Nanisivik;Nanisivik;Canada;YSR;CYSR;72.982222;-84.613611;2106;-5;A;America/Toronto
138;Summerside;Summerside;Canada;YSU;CYSU;46.440556;-63.833611;56;-4;A;America/Halifax
139;Sachs Harbour;Sachs Harbour;Canada;YSY;CYSY;71.993889;-125.2425;282;-7;A;America/Edmonton
140;Cape Dorset;Cape Dorset;Canada;YTE;CYTE;64.23;-76.526667;164;-5;A;America/Toronto
141;Thompson;Thompson;Canada;YTH;CYTH;55.801111;-97.864166;729;-6;A;America/Winnipeg
142;Trenton;Trenton;Canada;YTR;CYTR;44.118889;-77.528056;283;-5;A;America/Toronto
143;Timmins;Timmins;Canada;YTS;CYTS;48.569721;-81.376667;967;-5;A;America/Toronto
144;City Centre;Toronto;Canada;YTZ;CYTZ;43.627499;-79.396167;251;-5;A;America/Toronto
145;Tuktoyaktuk;Tuktoyaktuk;Canada;YUB;CYUB;69.433334;-133.026389;15;-7;A;America/Edmonton
146;Pierre Elliott Trudeau Intl;Montreal;Canada;YUL;CYUL;45.470556;-73.740833;118;-5;A;America/Toronto
147;Repulse Bay;Repulse Bay;Canada;YUT;CYUT;66.521389;-86.224722;80;-6;A;America/Winnipeg
148;Hall Beach;Hall Beach;Canada;YUX;CYUX;68.776111;-81.243611;27;-5;A;America/Toronto
149;Rouyn Noranda;Rouyn;Canada;YUY;CYUY;48.206111;-78.835556;988;-5;A;America/Toronto
150;La Ronge;La Ronge;Canada;YVC;CYVC;55.15139;-105.261944;1242;-6;N;America/Regina
151;Vermilion;Vermillion;Canada;YVG;CYVG;53.355833;-110.82389;2025;-7;A;America/Edmonton
152;Qikiqtarjuaq;Broughton Island;Canada;YVM;CYVM;67.545833;-64.031389;21;-5;A;America/Toronto
153;Val D Or;Val D'or;Canada;YVO;CYVO;48.053333;-77.782778;1107;-5;A;America/Toronto
154;Kuujjuaq;Quujjuaq;Canada;YVP;CYVP;58.096111;-68.426944;129;-5;A;America/Toronto
155;Norman Wells;Norman Wells;Canada;YVQ;CYVQ;65.281617;-126.798219;238;-7;A;America/Edmonton
156;Vancouver Intl;Vancouver;Canada;YVR;CYVR;49.193889;-123.184444;14;-8;A;America/Vancouver
157;Buffalo Narrows;Buffalo Narrows;Canada;YVT;CYVT;55.841944;-108.4175;1444;-6;N;America/Regina
158;Wiarton;Wiarton;Canada;YVV;CYVV;44.745834;-81.107222;729;-5;A;America/Toronto
159;Petawawa;Petawawa;Canada;YWA;CYWA;45.952221;-77.319168;427;-5;A;America/Toronto
160;Winnipeg Intl;Winnipeg;Canada;YWG;CYWG;49.910036;-97.239886;783;-6;A;America/Winnipeg
161;Wabush;Wabush;Canada;YWK;CYWK;52.921944;-66.864444;1808;-4;A;America/Halifax
162;Williams Lake;Williams Lake;Canada;YWL;CYWL;52.183056;-122.054167;3085;-8;A;America/Vancouver
163;Wrigley;Wrigley;Canada;YWY;CYWY;63.209444;-123.436667;489;-7;A;America/Edmonton
164;Canadian Rockies Intl;Cranbrook;Canada;YXC;CYXC;49.612222;-115.781944;3084;-7;A;America/Edmonton
165;Edmonton City Centre;Edmonton;Canada;YXD;CYXD;53.5725;-113.520556;2200;-7;A;America/Edmonton
166;Saskatoon J G Diefenbaker Intl;Saskatoon;Canada;YXE;CYXE;52.170834;-106.699722;1653;-6;N;America/Regina
167;Medicine Hat;Medicine Hat;Canada;YXH;CYXH;50.01889;-110.720833;2352;-7;A;America/Edmonton
168;Fort St John;Fort Saint John;Canada;YXJ;CYXJ;56.238056;-120.740278;2280;-7;A;America/Dawson_Creek
169;Sioux Lookout;Sioux Lookout;Canada;YXL;CYXL;50.113889;-91.905278;1258;-6;A;America/Winnipeg
170;Pangnirtung;Pangnirtung;Canada;YXP;CYXP;66.145;-65.713611;75;-5;A;America/Toronto
171;Timiskaming Rgnl;Earlton;Canada;YXR;CYXR;47.695;-79.848889;798;-5;A;America/Toronto
172;Prince George;Prince George;Canada;YXS;CYXS;53.889444;-122.678889;2267;-8;A;America/Vancouver
173;Terrace;Terrace;Canada;YXT;CYXT;54.468508;-128.576219;713;-8;A;America/Vancouver
174;London;London;Canada;YXU;CYXU;43.033056;-81.151111;912;-5;A;America/Toronto
175;Abbotsford;Abbotsford;Canada;YXX;CYXX;49.025278;-122.360556;195;-8;A;America/Vancouver
176;Whitehorse Intl;Whitehorse;Canada;YXY;CYXY;60.709553;-135.067269;2317;-8;A;America/Vancouver
177;North Bay;North Bay;Canada;YYB;CYYB;46.363611;-79.422778;1215;-5;A;America/Toronto
178;Calgary Intl;Calgary;Canada;YYC;CYYC;51.113888;-114.020278;3557;-7;A;America/Edmonton
179;Smithers;Smithers;Canada;YYD;CYYD;54.824722;-127.182778;1712;-8;A;America/Vancouver
180;Fort Nelson;Fort Nelson;Canada;YYE;CYYE;58.836389;-122.596944;1253;-8;A;America/Vancouver
181;Penticton;Penticton;Canada;YYF;CYYF;49.463056;-119.602222;1129;-8;A;America/Vancouver
182;Charlottetown;Charlottetown;Canada;YYG;CYYG;46.290001;-63.121111;160;-4;A;America/Halifax
183;Taloyoak;Spence Bay;Canada;YYH;CYYH;69.546667;-93.576667;92;-7;A;America/Edmonton
184;Victoria Intl;Victoria;Canada;YYJ;CYYJ;48.646944;-123.425833;63;-8;A;America/Vancouver
185;Lynn Lake;Lynn Lake;Canada;YYL;CYYL;56.863888;-101.07611;1170;-6;A;America/Winnipeg
186;Swift Current;Swift Current;Canada;YYN;CYYN;50.291944;-107.690556;2680;-6;N;America/Regina
187;Churchill;Churchill;Canada;YYQ;CYYQ;58.739167;-94.065;94;-6;A;America/Winnipeg
188;Goose Bay;Goose Bay;Canada;YYR;CYYR;53.319168;-60.425833;160;-4;A;America/Halifax
189;St Johns Intl;St. John's;Canada;YYT;CYYT;47.61861;-52.751945;461;-3.5;A;America/St_Johns
190;Kapuskasing;Kapuskasing;Canada;YYU;CYYU;49.413889;-82.4675;743;-5;A;America/Toronto
191;Armstrong;Armstrong;Canada;YYW;CYYW;50.290279;-88.909721;1058;-5;A;America/Toronto
192;Mont Joli;Mont Joli;Canada;YYY;CYYY;48.608612;-68.208056;172;-5;A;America/Toronto
193;Lester B Pearson Intl;Toronto;Canada;YYZ;CYYZ;43.677223;-79.630556;569;-5;A;America/Toronto
194;Downsview;Toronto;Canada;YZD;CYZD;43.7425;-79.465556;652;-5;A;America/Toronto
195;Gore Bay Manitoulin;Gore Bay;Canada;YZE;CYZE;45.885277;-82.567778;635;-5;A;America/Toronto
196;Yellowknife;Yellowknife;Canada;YZF;CYZF;62.462778;-114.440278;675;-7;A;America/Edmonton
197;Slave Lake;Slave Lake;Canada;YZH;CYZH;55.293056;-114.777222;1912;-7;A;America/Edmonton
198;Sandspit;Sandspit;Canada;YZP;CYZP;53.254333;-131.813833;21;-8;A;America/Vancouver
199;Chris Hadfield;Sarnia;Canada;YZR;CYZR;42.999444;-82.308889;594;-5;A;America/Toronto
200;Port Hardy;Port Hardy;Canada;YZT;CYZT;50.680556;-127.366667;71;-8;A;America/Vancouver
201;Whitecourt;Whitecourt;Canada;YZU;CYZU;54.14389;-115.786667;2567;-7;A;America/Edmonton
202;Sept Iles;Sept-iles;Canada;YZV;CYZV;50.223333;-66.265556;180;-5;A;America/Toronto
203;Teslin;Teslin;Canada;YZW;CYZW;60.172779;-132.742778;2313;-8;A;America/Vancouver
204;Greenwood;Greenwood;Canada;YZX;CYZX;44.984444;-64.916944;92;-4;A;America/Halifax
205;Faro;Faro;Canada;ZFA;CZFA;62.2075;-133.375833;2351;-8;A;America/Vancouver
206;Fort Mcpherson;Fort Mcpherson;Canada;ZFM;CZFM;67.4075;-134.860556;116;-7;A;America/Edmonton
207;Blida;Blida;Algeria;;DAAB;36.503613;2.814167;535;1;N;Africa/Algiers
208;Bou Saada;Bou Saada;Algeria;;DAAD;35.3325;4.206389;1506;1;N;Africa/Algiers
209;Soummam;Bejaja;Algeria;BJA;DAAE;36.711997;5.069922;20;1;N;Africa/Algiers
210;Houari Boumediene;Algier;Algeria;ALG;DAAG;36.691014;3.215408;82;1;N;Africa/Algiers
211;Tiska;Djanet;Algeria;DJG;DAAJ;24.292767;9.452444;3176;1;N;Africa/Algiers
212;Boufarik;Boufarik;Algeria;QFD;DAAK;36.545834;2.876111;335;1;N;Africa/Algiers
213;Reggane;Reggan;Algeria;;DAAN;26.710103;0.285647;955;1;N;Africa/Algiers
214;Illizi Takhamalt;Illizi;Algeria;VVZ;DAAP;26.723536;8.622653;1778;1;N;Africa/Algiers
215;Ain Oussera;Ain Oussera;Algeria;;DAAQ;35.525414;2.878714;2132;1;N;Africa/Algiers
216;Tamanrasset;Tamanrasset;Algeria;TMR;DAAT;22.811461;5.451075;4518;1;N;Africa/Algiers
217;Jijel;Jijel;Algeria;GJL;DAAV;36.795136;5.873608;36;1;N;Africa/Algiers
218;Mecheria;Mecheria;Algeria;;DAAY;33.535853;-0.242353;3855;1;N;Africa/Algiers
219;Relizane;Relizane;Algeria;;DAAZ;35.752239;0.626272;282;1;N;Africa/Algiers
220;Annaba;Annaba;Algeria;AAE;DABB;36.822225;7.809167;16;1;N;Africa/Algiers
221;Mohamed Boudiaf Intl;Constantine;Algeria;CZL;DABC;36.276028;6.620386;2265;1;N;Africa/Algiers
222;Cheikh Larbi Tebessi;Tebessa;Algeria;TEE;DABS;35.431611;8.120717;2661;1;N;Africa/Algiers
224;Hassi R Mel;Tilrempt;Algeria;HRM;DAFH;32.930431;3.311542;2540;1;N;Africa/Algiers
225;Bou Chekif;Tiaret;Algeria;TID;DAOB;35.341136;1.463147;3245;1;N;Africa/Algiers
226;Bou Sfer;Bou Sfer;Algeria;;DAOE;35.735389;-0.805389;187;1;N;Africa/Algiers
227;Tindouf;Tindouf;Algeria;TIN;DAOF;27.700372;-8.167103;1453;1;N;Africa/Algiers
228;Ech Cheliff;Ech-cheliff;Algeria;QAS;DAOI;36.212658;1.331775;463;1;N;Africa/Algiers
229;Tafaraoui;Oran;Algeria;TAF;DAOL;35.542444;-0.532278;367;1;N;Africa/Algiers
230;Zenata;Tlemcen;Algeria;TLM;DAON;35.016667;-1.45;814;1;N;Africa/Algiers
231;Es Senia;Oran;Algeria;ORN;DAOO;35.623858;-0.621183;295;1;N;Africa/Algiers
232;Sidi Bel Abbes;Sidi Bel Abbes;Algeria;;DAOS;35.171775;-0.593275;1614;1;N;Africa/Algiers
233;Ghriss;Ghriss;Algeria;MUW;DAOV;35.207725;0.147142;1686;1;N;Africa/Algiers
234;Touat Cheikh Sidi Mohamed Belkebir;Adrar;Algeria;AZR;DAUA;27.837589;-0.186414;919;1;N;Africa/Algiers
235;Biskra;Biskra;Algeria;BSK;DAUB;34.793289;5.738231;289;1;N;Africa/Algiers
236;El Golea;El Golea;Algeria;ELG;DAUE;30.571294;2.859586;1306;1;N;Africa/Algiers
237;Noumerat;Ghardaia;Algeria;GHA;DAUG;32.384106;3.794114;1512;1;N;Africa/Algiers
238;Oued Irara;Hassi Messaoud;Algeria;HME;DAUH;31.672972;6.140444;463;1;N;Africa/Algiers
239;In Salah;In Salah;Algeria;INZ;DAUI;27.251022;2.512017;896;1;N;Africa/Algiers
240;Sidi Mahdi;Touggourt;Algeria;TGR;DAUK;33.067803;6.088672;279;1;N;Africa/Algiers
241;Laghouat;Laghouat;Algeria;LOO;DAUL;33.764383;2.928344;2510;1;N;Africa/Algiers
242;Timimoun;Timimoun;Algeria;TMX;DAUT;29.237119;0.276033;1027;1;N;Africa/Algiers
243;Ouargla;Ouargla;Algeria;OGX;DAUU;31.917223;5.412778;492;1;N;Africa/Algiers
244;In Amenas;Zarzaitine;Algeria;IAM;DAUZ;28.05155;9.642911;1847;1;N;Africa/Algiers
245;Cadjehoun;Cotonou;Benin;COO;DBBB;6.357228;2.384353;19;1;N;Africa/Porto-Novo
246;Ouagadougou;Ouagadougou;Burkina Faso;OUA;DFFD;12.353194;-1.512417;1037;0;N;Africa/Ouagadougou
247;Bobo Dioulasso;Bobo-dioulasso;Burkina Faso;BOY;DFOO;11.160056;-4.330969;1511;0;N;Africa/Ouagadougou
248;Kotoka Intl;Accra;Ghana;ACC;DGAA;5.605186;-0.166786;205;0;N;Africa/Accra
249;Tamale;Tamale;Ghana;TML;DGLE;9.557192;-0.863214;553;0;N;Africa/Accra
250;Wa;Wa;Ghana;;DGLW;10.082664;-2.507694;1060;0;N;Africa/Accra
251;Sunyani;Sunyani;Ghana;NYI;DGSN;7.361828;-2.328756;1014;0;N;Africa/Accra
252;Takoradi;Takoradi;Ghana;TKD;DGTK;4.896056;-1.774756;21;0;N;Africa/Accra
253;Abidjan Felix Houphouet Boigny Intl;Abidjan;Cote d'Ivoire;ABJ;DIAP;5.261386;-3.926294;21;0;N;Africa/Abidjan
254;Bouake;Bouake;Cote d'Ivoire;BYK;DIBK;7.7388;-5.073667;1230;0;N;Africa/Abidjan
255;Daloa;Daloa;Cote d'Ivoire;DJO;DIDL;6.792808;-6.473189;823;0;N;Africa/Abidjan
256;Korhogo;Korhogo;Cote d'Ivoire;HGO;DIKO;9.387183;-5.556664;1214;0;N;Africa/Abidjan
257;Man;Man;Cote d'Ivoire;MJC;DIMN;7.272069;-7.587364;1089;0;N;Africa/Abidjan
258;San Pedro;San Pedro;Cote d'Ivoire;SPY;DISP;4.746717;-6.660817;26;0;N;Africa/Abidjan
259;Yamoussoukro;Yamoussoukro;Cote d'Ivoire;ASK;DIYO;6.903167;-5.365581;699;0;N;Africa/Abidjan
260;Nnamdi Azikiwe Intl;Abuja;Nigeria;ABV;DNAA;9.006792;7.263172;1123;1;N;Africa/Lagos
261;Akure;Akure;Nigeria;AKR;DNAK;7.246739;5.301008;1100;1;N;Africa/Lagos
262;Benin;Benin;Nigeria;BNI;DNBE;6.316981;5.599503;258;1;N;Africa/Lagos
263;Calabar;Calabar;Nigeria;CBQ;DNCA;4.976019;8.347197;210;1;N;Africa/Lagos
264;Enugu;Enugu;Nigeria;ENU;DNEN;6.474272;7.561961;466;1;N;Africa/Lagos
265;Gusau;Gusau;Nigeria;QUS;DNGU;12.171667;6.696111;1520;1;N;Africa/Lagos
266;Ibadan;Ibadan;Nigeria;IBA;DNIB;7.362458;3.978333;725;1;N;Africa/Lagos
267;Ilorin;Ilorin;Nigeria;ILR;DNIL;8.440211;4.493919;1126;1;N;Africa/Lagos
268;Yakubu Gowon;Jos;Nigeria;JOS;DNJO;9.639828;8.86905;4232;1;N;Africa/Lagos
269;Kaduna;Kaduna;Nigeria;KAD;DNKA;10.696025;7.320114;2073;1;N;Africa/Lagos
270;Mallam Aminu Intl;Kano;Nigeria;KAN;DNKN;12.047589;8.524622;1562;1;N;Africa/Lagos
271;Maiduguri;Maiduguri;Nigeria;MIU;DNMA;11.855347;13.08095;1099;1;N;Africa/Lagos
272;Makurdi;Makurdi;Nigeria;MDI;DNMK;7.703883;8.613939;371;1;N;Africa/Lagos
273;Murtala Muhammed;Lagos;Nigeria;LOS;DNMM;6.577369;3.321156;135;1;N;Africa/Lagos
274;Minna New;Minna;Nigeria;MXJ;DNMN;9.652172;6.462256;834;1;N;Africa/Lagos
275;Port Harcourt Intl;Port Hartcourt;Nigeria;PHC;DNPO;5.015494;6.949594;87;1;N;Africa/Lagos
276;Sadiq Abubakar Iii Intl;Sokoto;Nigeria;SKO;DNSO;12.916322;5.207189;1010;1;N;Africa/Lagos
277;Yola;Yola;Nigeria;YOL;DNYO;9.257553;12.430422;599;1;N;Africa/Lagos
278;Zaria;Zaria;Nigeria;ZAR;DNZA;11.130192;7.685806;2170;1;N;Africa/Lagos
279;Maradi;Maradi;Niger;MFQ;DRRM;13.502531;7.126753;1240;1;N;Africa/Niamey
280;Diori Hamani;Niamey;Niger;NIM;DRRN;13.481547;2.183614;732;1;N;Africa/Niamey
281;Tahoua;Tahoua;Niger;THZ;DRRT;14.875658;5.265358;1266;1;N;Africa/Niamey
282;Manu Dayak;Agadez;Niger;AJY;DRZA;16.965997;8.000114;1657;1;N;Africa/Niamey
283;Dirkou;Dirkou;Niger;;DRZD;18.968703;12.86865;1273;1;N;Africa/Niamey
284;Diffa;Diffa;Niger;;DRZF;13.372894;12.626686;994;1;N;Africa/Niamey
285;Zinder;Zinder;Niger;ZND;DRZR;13.778997;8.983761;1516;1;N;Africa/Niamey
286;Habib Bourguiba Intl;Monastir;Tunisia;MIR;DTMB;35.758056;10.754722;9;1;E;Africa/Tunis
287;Carthage;Tunis;Tunisia;TUN;DTTA;36.851033;10.227217;22;1;E;Africa/Tunis
288;Sidi Ahmed Air Base;Bizerte;Tunisia;;DTTB;37.245447;9.791453;20;1;E;Africa/Tunis
289;Remada;Remada;Tunisia;;DTTD;32.306156;10.382108;1004;1;E;Africa/Tunis
290;Gafsa;Gafsa;Tunisia;GAF;DTTF;34.422022;8.822503;1060;1;E;Africa/Tunis
291;Gabes;Gabes;Tunisia;GAE;DTTG;33.876919;10.103333;26;1;E;Africa/Tunis
292;Borj El Amri;Bordj El Amri;Tunisia;;DTTI;36.721339;9.943147;108;1;E;Africa/Tunis
293;Zarzis;Djerba;Tunisia;DJE;DTTJ;33.875031;10.775461;19;1;E;Africa/Tunis
294;El Borma;El Borma;Tunisia;EBM;DTTR;31.704281;9.254619;827;1;E;Africa/Tunis
295;Thyna;Sfax;Tunisia;SFA;DTTX;34.717953;10.690972;85;1;E;Africa/Tunis
296;Nefta;Tozeur;Tunisia;TOE;DTTZ;33.939722;8.110556;287;1;E;Africa/Tunis
297;Niamtougou International;Niatougou;Togo;LRL;DXNG;9.767333;1.09125;1515;0;N;Africa/Lome
298;Gnassingbe Eyadema Intl;Lome;Togo;LFW;DXXX;6.165611;1.254511;72;0;N;Africa/Lome
299;Deurne;Antwerp;Belgium;ANR;EBAW;51.189444;4.460278;39;1;E;Europe/Brussels
300;Beauvechain;Beauvechain;Belgium;;EBBE;50.75861;4.768333;370;1;E;Europe/Brussels
301;Kleine Brogel;Kleine Brogel;Belgium;;EBBL;51.168333;5.47;200;1;E;Europe/Brussels
302;Brussels Natl;Brussels;Belgium;BRU;EBBR;50.901389;4.484444;184;1;E;Europe/Brussels
303;Bertrix;Bertrix;Belgium;;EBBX;49.891667;5.223889;1514;1;E;Europe/Brussels
304;Brussels South;Charleroi;Belgium;CRL;EBCI;50.459197;4.453817;614;1;E;Europe/Brussels
305;Chievres Ab;Chievres;Belgium;;EBCV;50.575833;3.831;193;1;E;Europe/Brussels
306;Koksijde;Koksijde;Belgium;;EBFN;51.090278;2.652778;20;1;E;Europe/Brussels
307;Florennes;Florennes;Belgium;;EBFS;50.243333;4.645833;935;1;E;Europe/Brussels
308;Wevelgem;Kortrijk-vevelgem;Belgium;QKT;EBKT;50.817222;3.204722;64;1;E;Europe/Brussels
309;Liege;Liege;Belgium;LGG;EBLG;50.637417;5.443222;659;1;E;Europe/Brussels
310;Oostende;Ostend;Belgium;OST;EBOS;51.198889;2.862222;13;1;E;Europe/Brussels
311;Zutendaal;Zutendaal;Belgium;;EBSL;50.9475;5.590556;312;1;E;Europe/Brussels
312;Sint Truiden;Sint-truiden;Belgium;;EBST;50.791944;5.201667;246;1;E;Europe/Brussels
313;Saint Hubert Mil;St.-hubert;Belgium;;EBSU;50.035833;5.404167;1930;1;E;Europe/Brussels
314;Ursel;Ursel;Belgium;;EBUL;51.144133;3.474361;95;1;E;Europe/Brussels
315;Weelde;Weelde;Belgium;;EBWE;51.394783;4.960194;105;1;E;Europe/Brussels
316;Zoersel;Zoersel;Belgium;;EBZR;51.264722;4.753333;53;1;E;Europe/Brussels
317;Bautzen;Bautzen;Germany;BBJ;EDAB;51.193531;14.519747;568;1;E;Europe/Berlin
318;Altenburg Nobitz;Altenburg;Germany;AOC;EDAC;50.981817;12.506361;640;1;E;Europe/Berlin
319;Dessau;Dessau;Germany;;EDAD;51.831828;12.184033;187;1;E;Europe/Berlin
320;Eisenhuttenstadt;Eisenhuettenstadt;Germany;;EDAE;52.197333;14.585667;149;1;E;Europe/Berlin
6891;Putnam County Airport;Greencastle;United States;4I7;\N;39.6335556;-86.8138056;842;-5;U;America/New_York
322;Grossenhain;Suhl;Germany;;EDAK;51.308111;13.554973;417;1;E;Europe/Berlin
323;Merseburg;Muehlhausen;Germany;;EDAM;51.363;11.940833;340;1;E;Europe/Berlin
324;Halle Oppin;Halle;Germany;;EDAQ;51.552;12.052667;347;1;E;Europe/Berlin
325;Riesa Gohlis;Riesa;Germany;;EDAU;51.2935;13.356;322;1;E;Europe/Berlin
326;Rechlin Larz;Rechlin-laerz;Germany;;EDAX;53.306417;12.753139;220;1;E;Europe/Berlin
327;Strausberg;Strausberg;Germany;;EDAY;52.579978;13.915683;263;1;E;Europe/Berlin
328;Schonhagen;Schoenhagen;Germany;;EDAZ;52.203258;13.158408;131;1;E;Europe/Berlin
329;Barth;Barth;Germany;BBH;EDBH;54.33754;12.699705;23;1;E;Europe/Berlin
330;Jena Schongleina;Jena;Germany;;EDBJ;50.915161;11.714519;1228;1;E;Europe/Berlin
331;Kyritz;Kyritz;Germany;;EDBK;52.918833;12.425333;130;1;E;Europe/Berlin
332;Magdeburg;Magdeburg;Germany;;EDBM;52.073658;11.626467;268;1;E;Europe/Berlin
333;Rothenburg Gorlitz;Rothenburg/ol;Germany;;EDBR;51.363167;14.95;517;1;E;Europe/Berlin
334;Anklam;Anklam;Germany;;EDCA;53.8327;13.669131;18;1;E;Europe/Berlin
335;Cottbus Drewitz;Cottbus;Germany;;EDCD;51.889475;14.531986;274;1;E;Europe/Berlin
336;Kamenz;Kamenz;Germany;;EDCM;51.29625;14.129;495;1;E;Europe/Berlin
337;Schonefeld;Berlin;Germany;SXF;EDDB;52.380001;13.5225;157;1;E;Europe/Berlin
338;Dresden;Dresden;Germany;DRS;EDDC;51.132767;13.767161;755;1;E;Europe/Berlin
339;Erfurt;Erfurt;Germany;ERF;EDDE;50.979811;10.958106;1036;1;E;Europe/Berlin
340;Frankfurt Main;Frankfurt;Germany;FRA;EDDF;50.026421;8.543125;364;1;E;Europe/Berlin
341;Munster Osnabruck;Munster;Germany;FMO;EDDG;52.134642;7.684831;160;1;E;Europe/Berlin
342;Hamburg;Hamburg;Germany;HAM;EDDH;53.630389;9.988228;53;1;E;Europe/Berlin
343;Tempelhof;Berlin;Germany;THF;EDDI;52.473025;13.403944;167;1;E;Europe/Berlin
344;Koln Bonn;Cologne;Germany;CGN;EDDK;50.865917;7.142744;302;1;E;Europe/Berlin
345;Dusseldorf;Duesseldorf;Germany;DUS;EDDL;51.289453;6.766775;147;1;E;Europe/Berlin
346;Franz Josef Strauss;Munich;Germany;MUC;EDDM;48.353783;11.786086;1487;1;E;Europe/Berlin
347;Nurnberg;Nuernberg;Germany;NUE;EDDN;49.4987;11.066897;1046;1;E;Europe/Berlin
348;Leipzig Halle;Leipzig;Germany;LEJ;EDDP;51.432447;12.241633;465;1;E;Europe/Berlin
349;Saarbrucken;Saarbruecken;Germany;SCN;EDDR;49.214553;7.109508;1058;1;E;Europe/Berlin
350;Stuttgart;Stuttgart;Germany;STR;EDDS;48.689878;9.221964;1276;1;E;Europe/Berlin
351;Tegel;Berlin;Germany;TXL;EDDT;52.559686;13.287711;122;1;E;Europe/Berlin
352;Hannover;Hannover;Germany;HAJ;EDDV;52.461056;9.685078;183;1;E;Europe/Berlin
353;Neuenland;Bremen;Germany;BRE;EDDW;53.0475;8.786667;14;1;E;Europe/Berlin
354;Egelsbach;Egelsbach;Germany;;EDFE;49.960833;8.6415;385;1;E;Europe/Berlin
355;Frankfurt Hahn;Hahn;Germany;HHN;EDFH;49.948672;7.263892;1649;1;E;Europe/Berlin
356;Mannheim City;Mannheim;Germany;MHG;EDFM;49.472706;8.514264;309;1;E;Europe/Berlin
357;Allendorf Eder;Allendorf;Germany;;EDFQ;51.034878;8.680839;1164;1;E;Europe/Berlin
358;Worms;Worms;Germany;;EDFV;49.606511;8.3684;295;1;E;Europe/Berlin
359;Mainz Finthen;Mainz;Germany;;EDFZ;49.968931;8.148336;760;1;E;Europe/Berlin
360;Eisenach Kindel;Eisenach;Germany;;EDGE;50.992797;10.472711;1101;1;E;Europe/Berlin
361;Siegerland;Siegerland;Germany;;EDGS;50.707658;8.082969;1966;1;E;Europe/Berlin
362;Hamburg Finkenwerder;Hamburg;Germany;XFW;EDHI;53.535886;9.837025;22;1;E;Europe/Berlin
363;Kiel Holtenau;Kiel;Germany;KEL;EDHK;54.3795;10.145167;101;1;E;Europe/Berlin
364;Lubeck Blankensee;Luebeck;Germany;LBC;EDHL;53.805367;10.719222;53;1;E;Europe/Berlin
365;Dahlemer Binz;Dahlemer Binz;Germany;;EDKV;50.405888;6.528083;1896;1;E;Europe/Berlin
366;Meinerzhagen;Meinerzhagen;Germany;;EDKZ;51.099445;7.601944;1548;1;E;Europe/Berlin
367;Arnsberg Menden;Arnsberg;Germany;ZCA;EDLA;51.483333;7.899333;794;1;E;Europe/Berlin
368;Essen Mulheim;Essen;Germany;ESS;EDLE;51.402333;6.937333;424;1;E;Europe/Berlin
369;Bielefeld;Bielefeld;Germany;;EDLI;51.964833;8.544833;454;1;E;Europe/Berlin
370;Monchengladbach;Moenchengladbach;Germany;MGL;EDLN;51.230356;6.504494;125;1;E;Europe/Berlin
371;Paderborn Lippstadt;Paderborn;Germany;PAD;EDLP;51.614089;8.616317;699;1;E;Europe/Berlin
372;Stadtlohn Vreden;Stadtlohn;Germany;;EDLS;51.995844;6.840667;157;1;E;Europe/Berlin
373;Dortmund;Dortmund;Germany;DTM;EDLW;51.518314;7.612242;425;1;E;Europe/Berlin
374;Augsburg;Augsburg;Germany;AGB;EDMA;48.425158;10.931764;1515;1;E;Europe/Berlin
375;Biberach An Der Riss;Biberach;Germany;;EDMB;48.111;9.762833;1903;1;E;Europe/Berlin
376;Eggenfelden;Eggenfelden;Germany;;EDME;48.396167;12.723667;1342;1;E;Europe/Berlin
377;Mindelheim Mattsies;Mindelheim;Germany;;EDMN;48.108833;10.524333;1857;1;E;Europe/Berlin
378;Oberpfaffenhofen;Oberpfaffenhofen;Germany;OBF;EDMO;48.081364;11.283067;1947;1;E;Europe/Berlin
379;Straubing;Straubing;Germany;;EDMS;48.90095;12.518186;1054;1;E;Europe/Berlin
380;Vilshofen;Vilshofen;Germany;;EDMV;48.635167;13.195667;991;1;E;Europe/Berlin
381;Leutkirch Unterzeil;Leutkirch;Germany;;EDNL;47.859117;10.014572;2099;1;E;Europe/Berlin
382;Friedrichshafen;Friedrichshafen;Germany;FDH;EDNY;47.671317;9.511486;1367;1;E;Europe/Berlin
383;Schwerin Parchim;Parchim;Germany;SZW;EDOP;53.426997;11.783436;166;1;E;Europe/Berlin
384;Stendal Borstel;Stendal;Germany;ZSN;EDOV;52.627778;11.818333;184;1;E;Europe/Berlin
385;Aalen Heidenheim Elchingen;Aalen-heidenheim;Germany;;EDPA;48.777833;10.264667;1916;1;E;Europe/Berlin
386;Bayreuth;Bayreuth;Germany;BYU;EDQD;49.984428;11.638569;1601;1;E;Europe/Berlin
387;Burg Feuerstein;Burg Feuerstein;Germany;;EDQE;49.793833;11.133167;1674;1;E;Europe/Berlin
388;Hof Plauen;Hof;Germany;HOQ;EDQM;50.288836;11.854919;1960;1;E;Europe/Berlin
389;Hassfurt Schweinfurt;Hassfurt;Germany;;EDQT;50.018;10.5295;718;1;E;Europe/Berlin
390;Koblenz Winningen;Koblenz;Germany;ZNV;EDRK;50.3255;7.528667;640;1;E;Europe/Berlin
391;Trier Fohren;Trier;Germany;ZQF;EDRT;49.8635;6.788167;665;1;E;Europe/Berlin
392;Speyer;Speyer;Germany;ZQC;EDRY;49.302776;8.451195;312;1;E;Europe/Berlin
393;Zweibrucken;Zweibruecken;Germany;;EDRZ;49.209525;7.400647;1133;1;E;Europe/Berlin
394;Donaueschingen Villingen;Donaueschingen;Germany;ZQL;EDTD;47.973331;8.522223;2231;1;E;Europe/Berlin
395;Freiburg;Freiburg;Germany;;EDTF;48.022653;7.832583;799;1;E;Europe/Berlin
396;Mengen Hohentengen;Mengen;Germany;;EDTM;48.053833;9.372833;1820;1;E;Europe/Berlin
397;Schwabisch Hall;Schwaebisch Hall;Germany;;EDTY;49.118317;9.783956;1311;1;E;Europe/Berlin
398;Finsterwalde Schacksdorf;Soest;Germany;;EDUS;51.6075;13.738;399;1;E;Europe/Berlin
399;Braunschweig Wolfsburg;Braunschweig;Germany;BWE;EDVE;52.319167;10.556111;295;1;E;Europe/Berlin
400;Kassel Calden;Kassel;Germany;KSF;EDVK;51.408394;9.377631;908;1;E;Europe/Berlin
401;Hildesheim;Hildesheim;Germany;;EDVM;52.179833;9.945667;293;1;E;Europe/Berlin
402;Bremerhaven;Bremerhaven;Germany;BRV;EDWB;53.507081;8.572878;11;1;E;Europe/Berlin
403;Emden;Emden;Germany;EME;EDWE;53.391186;7.227408;2;1;E;Europe/Berlin
404;Leer Papenburg;Leer;Germany;;EDWF;53.271592;7.442344;3;1;E;Europe/Berlin
405;Wilhelmshaven Mariensiel;Wilhelmshaven;Germany;WVN;EDWI;53.504833;8.053333;16;1;E;Europe/Berlin
406;Borkum;Borkum;Germany;BMK;EDWR;53.5955;6.709167;3;1;E;Europe/Berlin
407;Norderney;Norderney;Germany;NRD;EDWY;53.706822;7.230247;6;1;E;Europe/Berlin
408;Flensburg Schaferhaus;Flensburg;Germany;FLF;EDXF;54.771772;9.378214;130;1;E;Europe/Berlin
409;Rendsburg Schachtholm;Rendsburg;Germany;;EDXR;54.220928;9.600803;23;1;E;Europe/Berlin
410;Westerland Sylt;Westerland;Germany;GWT;EDXW;54.91325;8.340472;51;1;E;Europe/Berlin
411;Amari;Armari Air Force Base;Estonia;;EEEI;59.260286;24.208467;65;2;E;Europe/Tallinn
412;Kardla;Kardla;Estonia;KDL;EEKA;58.990756;22.830733;18;2;E;Europe/Tallinn
413;Kuressaare;Kuressaare;Estonia;URE;EEKE;58.229883;22.509494;14;2;E;Europe/Tallinn
414;Parnu;Parnu;Estonia;EPU;EEPU;58.419044;24.472819;47;2;E;Europe/Tallinn
415;Tallinn;Tallinn-ulemiste International;Estonia;TLL;EETN;59.413317;24.832844;131;2;E;Europe/Tallinn
416;Tartu;Tartu;Estonia;TAY;EETU;58.307461;26.690428;219;2;E;Europe/Tallinn
417;Enontekio;Enontekio;Finland;ENF;EFET;68.362586;23.424322;1005;2;E;Europe/Helsinki
418;Eura;Eura;Finland;;EFEU;61.116112;22.201389;259;2;E;Europe/Helsinki
419;Halli;Halli;Finland;KEV;EFHA;61.85605;24.7866;479;2;E;Europe/Helsinki
420;Helsinki Malmi;Helsinki;Finland;HEM;EFHF;60.254558;25.042828;57;2;E;Europe/Helsinki
421;Helsinki Vantaa;Helsinki;Finland;HEL;EFHK;60.317222;24.963333;179;2;E;Europe/Helsinki
422;Hameenkyro;Hameenkyro;Finland;;EFHM;61.689656;23.073744;449;2;E;Europe/Helsinki
423;Hanko;Hanko;Finland;;EFHN;59.848864;23.083583;20;2;E;Europe/Helsinki
424;Hyvinkaa;Hyvinkaa;Finland;HYV;EFHV;60.654444;24.881111;430;2;E;Europe/Helsinki
425;Kiikala;Kikala;Finland;;EFIK;60.462502;23.6525;381;2;E;Europe/Helsinki
426;Immola;Immola;Finland;;EFIM;61.249172;28.903711;338;2;E;Europe/Helsinki
427;Kitee;Kitee;Finland;;EFIT;62.166111;30.073611;364;2;E;Europe/Helsinki
428;Ivalo;Ivalo;Finland;IVL;EFIV;68.607269;27.405328;481;2;E;Europe/Helsinki
429;Joensuu;Joensuu;Finland;JOE;EFJO;62.662906;29.60755;398;2;E;Europe/Helsinki
430;Jyvaskyla;Jyvaskyla;Finland;JYV;EFJY;62.399453;25.678253;459;2;E;Europe/Helsinki
431;Kauhava;Kauhava;Finland;KAU;EFKA;63.127078;23.051442;151;2;E;Europe/Helsinki
432;Kemi Tornio;Kemi;Finland;KEM;EFKE;65.781889;24.5991;61;2;E;Europe/Helsinki
433;Kajaani;Kajaani;Finland;KAJ;EFKI;64.285472;27.692414;483;2;E;Europe/Helsinki
434;Kauhajoki;Kauhajoki;Finland;;EFKJ;62.462502;22.393055;407;2;E;Europe/Helsinki
435;Kruunupyy;Kruunupyy;Finland;KOK;EFKK;63.721172;23.143131;84;2;E;Europe/Helsinki
436;Kemijarvi;Kemijarvi;Finland;;EFKM;66.712883;27.156786;692;2;E;Europe/Helsinki
437;Kuusamo;Kuusamo;Finland;KAO;EFKS;65.987575;29.239381;866;2;E;Europe/Helsinki
438;Kittila;Kittila;Finland;KTT;EFKT;67.701022;24.84685;644;2;E;Europe/Helsinki
439;Kuopio;Kuopio;Finland;KUO;EFKU;63.00715;27.797756;323;2;E;Europe/Helsinki
440;Lahti Vesivehmaa;Vesivehmaa;Finland;;EFLA;61.144158;25.693508;502;2;E;Europe/Helsinki
441;Lappeenranta;Lappeenranta;Finland;LPP;EFLP;61.044553;28.144397;349;2;E;Europe/Helsinki
442;Mariehamn;Mariehamn;Finland;MHQ;EFMA;60.122203;19.898156;17;2;E;Europe/Mariehamn
443;Menkijarvi;Menkijarvi;Finland;;EFME;62.946667;23.518889;331;2;E;Europe/Helsinki
444;Mikkeli;Mikkeli;Finland;MIK;EFMI;61.6866;27.201794;329;2;E;Europe/Helsinki
445;Nummela;Nummela;Finland;;EFNU;60.333889;24.296389;367;2;E;Europe/Helsinki
446;Oulu;Oulu;Finland;OUL;EFOU;64.930061;25.354564;47;2;E;Europe/Helsinki
447;Piikajarvi;Piikajarvi;Finland;;EFPI;61.245558;22.193356;148;2;E;Europe/Helsinki
448;Pori;Pori;Finland;POR;EFPO;61.461686;21.799983;44;2;E;Europe/Helsinki
449;Pudasjarvi;Pudasjarvi;Finland;;EFPU;65.402222;26.946944;397;2;E;Europe/Helsinki
450;Pyhasalmi;Pyhasalmi;Finland;;EFPY;63.731917;25.926306;528;2;E;Europe/Helsinki
451;Raahe Pattijoki;Pattijoki;Finland;;EFRH;64.688056;24.695833;118;2;E;Europe/Helsinki
452;Rantasalmi;Rantasalmi;Finland;;EFRN;62.065481;28.356494;292;2;E;Europe/Helsinki
453;Rovaniemi;Rovaniemi;Finland;RVN;EFRO;66.564822;25.830411;642;2;E;Europe/Helsinki
454;Rayskala;Rayskala;Finland;;EFRY;60.744722;24.107778;407;2;E;Europe/Helsinki
455;Savonlinna;Savonlinna;Finland;SVL;EFSA;61.943064;28.945136;311;2;E;Europe/Helsinki
456;Selanpaa;Selanpaa;Finland;;EFSE;61.062389;26.798861;417;2;E;Europe/Helsinki
457;Sodankyla;Sodankyla;Finland;SOT;EFSO;67.395033;26.619133;602;2;E;Europe/Helsinki
458;Tampere Pirkkala;Tampere;Finland;TMP;EFTP;61.414147;23.604392;390;2;E;Europe/Helsinki
459;Teisko;Teisko;Finland;;EFTS;61.77335;24.027006;515;2;E;Europe/Helsinki
460;Turku;Turku;Finland;TKU;EFTU;60.514142;22.262808;161;2;E;Europe/Helsinki
461;Utti;Utti;Finland;QVY;EFUT;60.896394;26.938353;339;2;E;Europe/Helsinki
462;Vaasa;Vaasa;Finland;VAA;EFVA;63.05065;21.762175;19;2;E;Europe/Helsinki
463;Varkaus;Varkaus;Finland;VRK;EFVR;62.171111;27.868611;286;2;E;Europe/Helsinki
464;Ylivieska;Ylivieska-raudaskyla;Finland;;EFYL;64.060547;24.715953;252;2;E;Europe/Helsinki
465;Belfast Intl;Belfast;United Kingdom;BFS;EGAA;54.6575;-6.215833;268;0;E;Europe/London
466;St Angelo;Enniskillen;United Kingdom;ENK;EGAB;54.398889;-7.651667;155;0;E;Europe/London
467;Belfast City;Belfast;United Kingdom;BHD;EGAC;54.618056;-5.8725;15;0;E;Europe/London
468;City of Derry;Londonderry;United Kingdom;LDY;EGAE;55.042778;-7.161111;22;0;E;Europe/London
469;Birmingham;Birmingham;United Kingdom;BHX;EGBB;52.453856;-1.748028;327;0;E;Europe/London
470;Coventry;Coventry;United Kingdom;CVT;EGBE;52.369722;-1.479722;267;0;E;Europe/London
471;Leicester;Leicester;United Kingdom;;EGBG;52.607778;-1.031944;469;0;E;Europe/London
472;Gloucestershire;Golouchestershire;United Kingdom;GLO;EGBJ;51.894167;-2.167222;101;0;E;Europe/London
474;Wolverhampton;Halfpenny Green;United Kingdom;;EGBO;52.5175;-2.259444;283;0;E;Europe/London
475;Kemble;Pailton;United Kingdom;;EGBP;51.668056;-2.056944;433;0;E;Europe/London
476;Turweston;Turweston;United Kingdom;;EGBT;52.040833;-1.095556;448;0;E;Europe/London
477;Wellesbourne Mountford;Wellesbourne;United Kingdom;;EGBW;52.192222;-1.614444;159;0;E;Europe/London
478;Manchester;Manchester;United Kingdom;MAN;EGCC;53.353744;-2.27495;257;0;E;Europe/London
479;Manchester Woodford;Woodfort;United Kingdom;;EGCD;53.338056;-2.148889;295;0;E;Europe/London
480;Chivenor;Chivenor;United Kingdom;;EGDC;51.087167;-4.150339;27;0;E;Europe/London
481;St Mawgan;Newquai;United Kingdom;NQY;EGDG;50.440558;-4.995408;390;0;E;Europe/London
482;Lyneham;Lyneham;United Kingdom;LYE;EGDL;51.505144;-1.993428;513;0;E;Europe/London
483;Boscombe Down;Boscombe Down;United Kingdom;;EGDM;51.152189;-1.747414;407;0;E;Europe/London
484;Culdrose;Culdrose;United Kingdom;;EGDR;50.086092;-5.255711;267;0;E;Europe/London
485;St Athan;St. Athan;United Kingdom;;EGDX;51.404811;-3.43575;163;0;E;Europe/London
486;Yeovilton;Yeovilton;United Kingdom;YEO;EGDY;51.009358;-2.638819;75;0;E;Europe/London
487;Haverfordwest;Haverfordwest;United Kingdom;;EGFE;51.833056;-4.961111;159;0;E;Europe/London
488;Cardiff;Cardiff;United Kingdom;CWL;EGFF;51.396667;-3.343333;220;0;E;Europe/London
489;Swansea;Swansea;United Kingdom;SWS;EGFH;51.605333;-4.067833;299;0;E;Europe/London
490;Bristol;Bristol;United Kingdom;BRS;EGGD;51.382669;-2.719089;622;0;E;Europe/London
491;Liverpool;Liverpool;United Kingdom;LPL;EGGP;53.333611;-2.849722;80;0;E;Europe/London
492;Luton;London;United Kingdom;LTN;EGGW;51.874722;-0.368333;526;0;E;Europe/London
493;Plymouth;Plymouth;United Kingdom;PLH;EGHD;50.422778;-4.105833;476;0;E;Europe/London
494;Bournemouth;Bournemouth;United Kingdom;BOH;EGHH;50.78;-1.8425;38;0;E;Europe/London
495;Southampton;Southampton;United Kingdom;SOU;EGHI;50.950261;-1.356803;44;0;E;Europe/London
496;Lasham;Lasham;United Kingdom;QLA;EGHL;51.187167;-1.0335;618;0;E;Europe/London
497;Alderney;Alderney;Guernsey;ACI;EGJA;49.706111;-2.214722;290;0;E;Europe/Guernsey
498;Guernsey;Guernsey;Guernsey;GCI;EGJB;49.434956;-2.601969;336;0;E;Europe/Guernsey
499;Jersey;Jersey;Jersey;JER;EGJJ;49.207947;-2.195508;277;0;E;Europe/Jersey
500;Shoreham;Shoreham By Sea;United Kingdom;ESH;EGKA;50.835556;-0.297222;7;0;E;Europe/London
501;Biggin Hill;Biggin Hill;United Kingdom;BQH;EGKB;51.330833;0.0325;598;0;E;Europe/London
502;Gatwick;London;United Kingdom;LGW;EGKK;51.148056;-0.190278;202;0;E;Europe/London
503;City;London;United Kingdom;LCY;EGLC;51.505278;0.055278;19;0;E;Europe/London
504;Farnborough;Farnborough;United Kingdom;FAB;EGLF;51.275833;-0.776333;238;0;E;Europe/London
505;Chalgrove;Chalsgrove;United Kingdom;;EGLJ;51.676111;-1.080833;240;0;E;Europe/London
506;Blackbushe;Blackbushe;United Kingdom;BBS;EGLK;51.323889;-0.8475;325;0;E;Europe/London
507;Heathrow;London;United Kingdom;LHR;EGLL;51.4775;-0.461389;83;0;E;Europe/London
508;Southend;Southend;United Kingdom;SEN;EGMC;51.571389;0.695556;49;0;E;Europe/London
509;Lydd;Lydd;United Kingdom;LYX;EGMD;50.956111;0.939167;13;0;E;Europe/London
510;Manston;Manston;United Kingdom;MSE;EGMH;51.342222;1.346111;178;0;E;Europe/London
511;Brough;Brough;United Kingdom;;EGNB;53.719667;-0.566333;12;0;E;Europe/London
512;Carlisle;Carlisle;United Kingdom;CAX;EGNC;54.9375;-2.809167;190;0;E;Europe/London
513;Gamston;Repton;United Kingdom;;EGNE;53.280556;-0.951389;91;0;E;Europe/London
514;Blackpool;Blackpool;United Kingdom;BLK;EGNH;53.771667;-3.028611;34;0;E;Europe/London
515;Humberside;Humberside;United Kingdom;HUY;EGNJ;53.574444;-0.350833;121;0;E;Europe/London
516;Walney Island;Barrow Island;United Kingdom;BWF;EGNL;54.131167;-3.263667;173;0;E;Europe/London
517;Leeds Bradford;Leeds;United Kingdom;LBA;EGNM;53.865897;-1.660569;681;0;E;Europe/London
518;Warton;Warton;United Kingdom;;EGNO;53.745097;-2.883061;55;0;E;Europe/London
519;Hawarden;Hawarden;United Kingdom;CEG;EGNR;53.178056;-2.977778;45;0;E;Europe/London
520;Isle Of Man;Isle Of Man;Isle of Man;IOM;EGNS;54.083333;-4.623889;52;0;E;Europe/Isle_of_Man
521;Newcastle;Newcastle;United Kingdom;NCL;EGNT;55.0375;-1.691667;266;0;E;Europe/London
522;Durham Tees Valley Airport;Teesside;United Kingdom;MME;EGNV;54.509189;-1.429406;120;0;E;Europe/London
523;Nottingham East Midlands;East Midlands;United Kingdom;EMA;EGNX;52.831111;-1.328056;306;0;E;Europe/London
524;Llanbedr;Llanbedr;United Kingdom;;EGOD;52.811744;-4.123575;30;0;E;Europe/London
525;Ternhill;Ternhill;United Kingdom;;EGOE;52.871164;-2.533561;272;0;E;Europe/London
526;Shawbury;Shawbury;United Kingdom;;EGOS;52.798169;-2.668042;249;0;E;Europe/London
528;Woodvale;Woodvale;United Kingdom;;EGOW;53.581575;-3.055522;37;0;E;Europe/London
529;Kirkwall;Kirkwall;United Kingdom;KOI;EGPA;58.957778;-2.905;50;0;E;Europe/London
530;Sumburgh;Sumburgh;United Kingdom;LSI;EGPB;59.878889;-1.295556;20;0;E;Europe/London
531;Wick;Wick;United Kingdom;WIC;EGPC;58.458889;-3.093056;126;0;E;Europe/London
532;Dyce;Aberdeen;United Kingdom;ABZ;EGPD;57.201944;-2.197778;215;0;E;Europe/London
533;Inverness;Inverness;United Kingdom;INV;EGPE;57.5425;-4.0475;31;0;E;Europe/London
534;Glasgow;Glasgow;United Kingdom;GLA;EGPF;55.871944;-4.433056;26;0;E;Europe/London
535;Edinburgh;Edinburgh;United Kingdom;EDI;EGPH;55.95;-3.3725;135;0;E;Europe/London
536;Islay;Islay;United Kingdom;ILY;EGPI;55.681944;-6.256667;56;0;E;Europe/London
537;Prestwick;Prestwick;United Kingdom;PIK;EGPK;55.509444;-4.586667;65;0;E;Europe/London
538;Benbecula;Benbecula;United Kingdom;BEB;EGPL;57.481111;-7.362778;19;0;E;Europe/London
539;Scatsta;Scatsta;United Kingdom;SDZ;EGPM;60.432778;-1.296111;81;0;E;Europe/London
540;Dundee;Dundee;United Kingdom;DND;EGPN;56.452499;-3.025833;17;0;E;Europe/London
541;Stornoway;Stornoway;United Kingdom;SYY;EGPO;58.215556;-6.331111;26;0;E;Europe/London
542;Tiree;Tiree;United Kingdom;TRE;EGPU;56.499167;-6.869167;38;0;E;Europe/London
543;Leuchars;Leuchars;United Kingdom;ADX;EGQL;56.372889;-2.868444;38;0;E;Europe/London
544;Lossiemouth;Lossiemouth;United Kingdom;LMO;EGQS;57.705214;-3.339169;42;0;E;Europe/London
545;Cambridge;Cambridge;United Kingdom;CBG;EGSC;52.205;0.175;47;0;E;Europe/London
546;Conington;Peterborough;United Kingdom;;EGSF;52.468056;-0.251111;26;0;E;Europe/London
547;Norwich;Norwich;United Kingdom;NWI;EGSH;52.675833;1.282778;117;0;E;Europe/London
548;Stansted;London;United Kingdom;STN;EGSS;51.885;0.235;348;0;E;Europe/London
549;North Weald;North Weald;United Kingdom;;EGSX;51.721667;0.154167;321;0;E;Europe/London
550;Sheffield City;Fowlmere;United Kingdom;;EGSY;53.394256;-1.388486;231;0;E;Europe/London
551;Cranfield;Cranfield;United Kingdom;;EGTC;52.072222;-0.616667;358;0;E;Europe/London
552;Exeter;Exeter;United Kingdom;EXT;EGTE;50.734444;-3.413889;102;0;E;Europe/London
553;Bristol Filton;Bristol;United Kingdom;FZO;EGTG;51.519444;-2.590833;226;0;E;Europe/London
554;Kidlington;Oxford;United Kingdom;OXF;EGTK;51.836944;-1.32;270;0;E;Europe/London
555;Benson;Benson;United Kingdom;;EGUB;51.616389;-1.095833;226;0;E;Europe/London
556;Lakenheath;Lakenheath;United Kingdom;;EGUL;52.409333;0.561;32;0;E;Europe/London
557;Mildenhall;Mildenhall;United Kingdom;MHZ;EGUN;52.361933;0.486406;33;0;E;Europe/London
558;Wattisham;Wattisham;United Kingdom;;EGUW;52.127283;0.956264;284;0;E;Europe/London
559;Wyton;Wyton;United Kingdom;;EGUY;52.357167;-0.107833;135;0;E;Europe/London
560;Fairford;Fairford;United Kingdom;FFD;EGVA;51.682167;-1.790028;286;0;E;Europe/London
561;Brize Norton;Brize Norton;United Kingdom;BZZ;EGVN;51.749964;-1.583617;288;0;E;Europe/London
562;Odiham;Odiham;United Kingdom;ODH;EGVO;51.234139;-0.942825;405;0;E;Europe/London
563;Cosford;Cosford;United Kingdom;;EGWC;52.640028;-2.305578;272;0;E;Europe/London
564;Northolt;Northolt;United Kingdom;NHT;EGWU;51.553;-0.418167;124;0;E;Europe/London
565;Coningsby;Coningsby;United Kingdom;QCY;EGXC;53.093014;-0.166014;25;0;E;Europe/London
566;Dishforth;Dishforth;United Kingdom;;EGXD;54.137186;-1.420253;117;0;E;Europe/London
567;Leeming;Leeming;United Kingdom;;EGXE;54.292383;-1.5354;132;0;E;Europe/London
568;Church Fenton;Church Fenton;United Kingdom;;EGXG;53.834333;-1.1955;29;0;E;Europe/London
569;Honington;Honington;United Kingdom;BEQ;EGXH;52.342611;0.772939;174;0;E;Europe/London
570;Cottesmore;Cottesmore;United Kingdom;;EGXJ;52.735711;-0.648769;461;0;E;Europe/London
571;Scampton;Scampton;United Kingdom;;EGXP;53.307778;-0.550833;202;0;E;Europe/London
572;Wittering;Wittering;United Kingdom;;EGXT;52.612558;-0.476453;273;0;E;Europe/London
573;Linton On Ouse;Linton-on-ouse;United Kingdom;;EGXU;54.048911;-1.252747;53;0;E;Europe/London
574;Waddington;Waddington;United Kingdom;WTN;EGXW;53.166167;-0.523811;231;0;E;Europe/London
575;Topcliffe;Topcliffe;United Kingdom;;EGXZ;54.205522;-1.382094;92;0;E;Europe/London
576;Cranwell;Cranwell;United Kingdom;;EGYD;53.03035;-0.483242;218;0;E;Europe/London
577;Barkston Heath;Barkston Heath;United Kingdom;;EGYE;52.962225;-0.561625;367;0;E;Europe/London
578;Marham;Marham;United Kingdom;KNF;EGYM;52.648353;0.550692;75;0;E;Europe/London
579;Mount Pleasant;Mount Pleasant;Falkland Islands;MPN;EGYP;-51.822777;-58.447222;244;-3;U;Atlantic/Stanley
580;Schiphol;Amsterdam;Netherlands;AMS;EHAM;52.308613;4.763889;-11;1;E;Europe/Amsterdam
581;Budel;Weert;Netherlands;;EHBD;51.25528;5.601389;114;1;E;Europe/Amsterdam
582;Maastricht;Maastricht;Netherlands;MST;EHBK;50.911658;5.770144;375;1;E;Europe/Amsterdam
583;Deelen;Deelen;Netherlands;;EHDL;52.060556;5.873056;158;1;E;Europe/Amsterdam
584;Drachten;Drachten;Netherlands;;EHDR;53.119167;6.129722;14;1;E;Europe/Amsterdam
585;Eindhoven;Eindhoven;Netherlands;EIN;EHEH;51.450139;5.374528;74;1;E;Europe/Amsterdam
586;Eelde;Groningen;Netherlands;GRQ;EHGG;53.11972;6.579444;17;1;E;Europe/Amsterdam
587;Gilze Rijen;Gilze-rijen;Netherlands;;EHGR;51.567389;4.931833;49;1;E;Europe/Amsterdam
588;De Kooy;De Kooy;Netherlands;DHR;EHKD;52.923353;4.780625;3;1;E;Europe/Amsterdam
589;Lelystad;Lelystad;Netherlands;;EHLE;52.460278;5.527222;-13;1;E;Europe/Amsterdam
590;Leeuwarden;Leeuwarden;Netherlands;LWR;EHLW;53.228611;5.760556;3;1;E;Europe/Amsterdam
591;Rotterdam;Rotterdam;Netherlands;RTM;EHRD;51.956944;4.437222;-15;1;E;Europe/Amsterdam
592;Soesterberg;Soesterberg;Netherlands;UTC;EHSB;52.1273;5.27619;66;1;E;Europe/Amsterdam
593;Twenthe;Enschede;Netherlands;ENS;EHTW;52.27;6.874167;114;1;E;Europe/Amsterdam
594;Valkenburg;Valkenburg;Netherlands;LID;EHVB;52.166139;4.417944;0;1;E;Europe/Amsterdam
595;Woensdrecht;Woensdrecht;Netherlands;WOE;EHWO;51.449092;4.342031;63;1;E;Europe/Amsterdam
596;Cork;Cork;Ireland;ORK;EICK;51.841269;-8.491111;502;0;E;Europe/Dublin
597;Galway;Galway;Ireland;GWY;EICM;53.300175;-8.941592;81;0;E;Europe/Dublin
599;Dublin;Dublin;Ireland;DUB;EIDW;53.421333;-6.270075;242;0;E;Europe/Dublin
600;Ireland West Knock;Connaught;Ireland;NOC;EIKN;53.910297;-8.818492;665;0;E;Europe/Dublin
601;Kerry;Kerry;Ireland;KIR;EIKY;52.180878;-9.523783;112;0;E;Europe/Dublin
602;Casement;Casement;Ireland;;EIME;53.301667;-6.451333;319;0;E;Europe/Dublin
603;Shannon;Shannon;Ireland;SNN;EINN;52.701978;-8.924817;46;0;E;Europe/Dublin
604;Sligo;Sligo;Ireland;SXL;EISG;54.280214;-8.599206;11;0;E;Europe/Dublin
605;Waterford;Waterford;Ireland;WAT;EIWF;52.1872;-7.086964;119;0;E;Europe/Dublin
607;Aarhus;Aarhus;Denmark;AAR;EKAH;56.300017;10.619008;82;1;E;Europe/Copenhagen
608;Billund;Billund;Denmark;BLL;EKBI;55.740322;9.151778;247;1;E;Europe/Copenhagen
609;Kastrup;Copenhagen;Denmark;CPH;EKCH;55.617917;12.655972;17;1;E;Europe/Copenhagen
610;Esbjerg;Esbjerg;Denmark;EBJ;EKEB;55.525942;8.553403;97;1;E;Europe/Copenhagen
611;Gronholt Hillerod;Gronholt;Denmark;;EKGH;55.941387;12.382222;97;1;E;Europe/Copenhagen
612;Karup;Karup;Denmark;KRP;EKKA;56.297458;9.124628;170;1;E;Europe/Copenhagen
613;Laeso;Laeso;Denmark;;EKLS;57.277228;11.000083;25;1;E;Europe/Copenhagen
614;Lolland Falster Maribo;Maribo;Denmark;;EKMB;54.699344;11.440117;16;1;E;Europe/Copenhagen
615;Odense;Odense;Denmark;ODE;EKOD;55.476664;10.330933;56;1;E;Europe/Copenhagen
616;Krusa Padborg;Krusa-padborg;Denmark;;EKPB;54.870306;9.279014;88;1;E;Europe/Copenhagen
617;Roskilde;Copenhagen;Denmark;RKE;EKRK;55.585567;12.131428;146;1;E;Europe/Copenhagen
618;Bornholm Ronne;Ronne;Denmark;RNN;EKRN;55.063267;14.759558;52;1;E;Europe/Copenhagen
619;Sonderborg;Soenderborg;Denmark;SGD;EKSB;54.964367;9.791731;24;1;E;Europe/Copenhagen
621;Skrydstrup;Skrydstrup;Denmark;SKS;EKSP;55.225553;9.263931;141;1;E;Europe/Copenhagen
622;Skive;Skive;Denmark;;EKSV;56.550208;9.172983;74;1;E;Europe/Copenhagen
623;Thisted;Thisted;Denmark;TED;EKTS;57.0688;8.705225;23;1;E;Europe/Copenhagen
624;Kolding Vamdrup;Kolding;Denmark;;EKVD;55.436283;9.330925;143;1;E;Europe/Copenhagen
625;Vagar;Vagar;Faroe Islands;FAE;EKVG;62.063628;-7.277219;280;0;E;Atlantic/Faeroe
626;Aars;Vesthimmerland;Denmark;;EKVH;56.846944;9.458611;119;1;E;Europe/Copenhagen
627;Stauning;Stauning;Denmark;STA;EKVJ;55.990122;8.353906;17;1;E;Europe/Copenhagen
628;Aalborg;Aalborg;Denmark;AAL;EKYT;57.092789;9.849164;10;1;E;Europe/Copenhagen
629;Luxembourg;Luxemburg;Luxembourg;LUX;ELLX;49.626575;6.211517;1234;1;E;Europe/Luxembourg
630;Vigra;Alesund;Norway;AES;ENAL;62.560372;6.110164;69;1;E;Europe/Oslo
631;Andenes;Andoya;Norway;ANX;ENAN;69.2925;16.144167;43;1;E;Europe/Oslo
632;Alta;Alta;Norway;ALF;ENAT;69.976111;23.371667;9;1;E;Europe/Oslo
633;Bomoen;Voss;Norway;;ENBM;60.63885;6.501497;300;1;E;Europe/Oslo
634;Bronnoy;Bronnoysund;Norway;BNN;ENBN;65.461111;12.2175;25;1;E;Europe/Oslo
635;Bodo;Bodo;Norway;BOO;ENBO;67.269167;14.365278;42;1;E;Europe/Oslo
636;Flesland;Bergen;Norway;BGO;ENBR;60.293386;5.218142;170;1;E;Europe/Oslo
637;Batsfjord;Batsfjord;Norway;BJF;ENBS;70.600278;29.6925;490;1;E;Europe/Oslo
638;Kjevik;Kristiansand;Norway;KRS;ENCN;58.204214;8.085369;57;1;E;Europe/Oslo
639;Dagali;Geilo;Norway;;ENDI;60.416667;8.512778;2618;1;E;Europe/Oslo
640;Bardufoss;Bardufoss;Norway;BDU;ENDU;69.055758;18.540356;252;1;E;Europe/Oslo
641;Evenes;Harstad/Narvik;Norway;EVE;ENEV;68.4913;16.678108;84;1;E;Europe/Oslo
642;Leirin;Fagernes;Norway;VDB;ENFG;61.015556;9.288056;2697;1;E;Europe/Oslo
643;Floro;Floro;Norway;FRO;ENFL;61.583611;5.024722;37;1;E;Europe/Oslo
644;Gardermoen;Oslo;Norway;OSL;ENGM;60.193917;11.100361;681;1;E;Europe/Oslo
645;Karmoy;Haugesund;Norway;HAU;ENHD;59.345267;5.208364;86;1;E;Europe/Oslo
646;Hasvik;Hasvik;Norway;HAA;ENHK;70.486675;22.139744;21;1;E;Europe/Oslo
647;Kvernberget;Kristiansund;Norway;KSU;ENKB;63.111781;7.824522;204;1;E;Europe/Oslo
648;Kjeller;Kjeller;Norway;;ENKJ;59.969336;11.036089;354;1;E;Europe/Oslo
649;Hoybuktmoen;Kirkenes;Norway;KKN;ENKR;69.725781;29.891295;283;1;E;Europe/Oslo
650;Lista;Farsund;Norway;FAN;ENLI;58.099486;6.62605;29;1;E;Europe/Oslo
651;Aro;Molde;Norway;MOL;ENML;62.744722;7.2625;10;1;E;Europe/Oslo
652;Kjaerstad;Mosjoen;Norway;MJF;ENMS;65.783997;13.214914;237;1;E;Europe/Oslo
653;Banak;Lakselv;Norway;LKL;ENNA;70.068814;24.973489;25;1;E;Europe/Oslo
654;Notodden;Notodden;Norway;NTB;ENNO;59.565683;9.212222;63;1;E;Europe/Oslo
655;Orland;Orland;Norway;OLA;ENOL;63.698908;9.604003;28;1;E;Europe/Oslo
656;Roros;Roros;Norway;RRS;ENRO;62.578411;11.342347;2054;1;E;Europe/Oslo
657;Moss;Rygge;Norway;RYG;ENRY;59.378933;10.785389;174;1;E;Europe/Oslo
658;Longyear;Svalbard;Norway;LYR;ENSB;78.246111;15.465556;88;1;E;Arctic/Longyearbyen
659;Geiteryggen;Skien;Norway;SKE;ENSN;59.185;9.566944;463;1;E;Europe/Oslo
660;Sorstokken;Stord;Norway;SRP;ENSO;59.791925;5.34085;160;1;E;Europe/Oslo
662;Stokka;Sandnessjoen;Norway;SSJ;ENST;65.956828;12.468944;56;1;E;Europe/Oslo
663;Langnes;Tromso;Norway;TOS;ENTC;69.683333;18.918919;31;1;E;Europe/Oslo
664;Torp;Sandefjord;Norway;TRF;ENTO;59.186703;10.258628;286;1;E;Europe/Oslo
665;Vaernes;Trondheim;Norway;TRD;ENVA;63.457556;10.92425;56;1;E;Europe/Oslo
666;Sola;Stavanger;Norway;SVG;ENZV;58.876778;5.637856;29;1;E;Europe/Oslo
667;Babice;Warsaw;Poland;;EPBC;52.268494;20.911047;352;1;E;Europe/Warsaw
668;Lech Walesa;Gdansk;Poland;GDN;EPGD;54.377569;18.466222;489;1;E;Europe/Warsaw
669;Balice;Krakow;Poland;KRK;EPKK;50.077731;19.784836;791;1;E;Europe/Warsaw
670;Muchowiec;Katowice;Poland;;EPKM;50.238147;19.034181;909;1;E;Europe/Warsaw
671;Pyrzowice;Katowice;Poland;KTW;EPKT;50.474253;19.080019;995;1;E;Europe/Warsaw
673;Mielec;Mielec;Poland;;EPML;50.322275;21.462131;548;1;E;Europe/Warsaw
674;Lawica;Poznan;Poland;POZ;EPPO;52.421031;16.826325;308;1;E;Europe/Warsaw
675;Jasionka;Rzeszow;Poland;RZE;EPRZ;50.109958;22.019;675;1;E;Europe/Warsaw
676;Goleniow;Szczecin;Poland;SZZ;EPSC;53.584731;14.902206;154;1;E;Europe/Warsaw
677;Redzikowo;Slupsk;Poland;OSP;EPSK;54.478889;17.1075;217;1;E;Europe/Warsaw
678;Swidwin;Shapaja;Poland;;EPSN;53.790639;15.82625;385;1;E;Europe/Warsaw
679;Okecie;Warsaw;Poland;WAW;EPWA;52.16575;20.967122;362;1;E;Europe/Warsaw
680;Strachowice;Wroclaw;Poland;WRO;EPWR;51.102683;16.885836;404;1;E;Europe/Warsaw
681;Babimost;Zielona Gora;Poland;IEG;EPZG;52.138517;15.798556;194;1;E;Europe/Warsaw
682;Malmen;Linkoeping;Sweden;;ESCF;58.402278;15.525683;308;1;E;Europe/Stockholm
683;Bravalla;Norrkoeping;Sweden;;ESCK;58.610867;16.103592;90;1;E;Europe/Stockholm
684;Uppsala;Uppsala;Sweden;;ESCM;59.897328;17.588581;68;1;E;Europe/Stockholm
685;Ronneby;Ronneby;Sweden;RNB;ESDF;56.266667;15.265;191;1;E;Europe/Stockholm
686;Rada;Rada;Sweden;;ESFR;58.498136;13.053231;230;1;E;Europe/Stockholm
687;Landvetter;Gothenborg;Sweden;GOT;ESGG;57.662836;12.279819;506;1;E;Europe/Stockholm
688;Jonkoping;Joenkoeping;Sweden;JKG;ESGJ;57.757594;14.068731;741;1;E;Europe/Stockholm
689;Falkoping;Falkoping;Sweden;;ESGK;58.169794;13.587847;785;1;E;Europe/Stockholm
690;Lidkoping;Lidkoping;Sweden;LDK;ESGL;58.465522;13.174414;200;1;E;Europe/Stockholm
691;Save;Gothenborg;Sweden;GSE;ESGP;57.774722;11.870372;59;1;E;Europe/Stockholm
692;Skovde;Skovde;Sweden;KVB;ESGR;58.4564;13.972672;324;1;E;Europe/Stockholm
693;Trollhattan Vanersborg;Trollhattan;Sweden;THN;ESGT;58.318056;12.345;137;1;E;Europe/Stockholm
694;Karlsborg;Karlsborg;Sweden;;ESIA;58.513842;14.507119;308;1;E;Europe/Stockholm
695;Satenas;Satenas;Sweden;;ESIB;58.426445;12.714389;181;1;E;Europe/Stockholm
696;Barkarby;Stockholm;Sweden;;ESKB;59.418694;17.890694;50;1;E;Europe/Stockholm
697;Karlskoga;Karlskoga;Sweden;KSK;ESKK;59.345867;14.495922;400;1;E;Europe/Stockholm
698;Mora;Mora;Sweden;MXX;ESKM;60.957908;14.511383;634;1;E;Europe/Stockholm
699;Skavsta;Stockholm;Sweden;NYO;ESKN;58.788636;16.912189;140;1;E;Europe/Stockholm
700;Arvika;Arvika;Sweden;;ESKV;59.675856;12.639442;237;1;E;Europe/Stockholm
701;Emmaboda;Emmaboda;Sweden;;ESMA;56.610761;15.604761;442;1;E;Europe/Stockholm
702;Feringe;Ljungby;Sweden;;ESMG;56.950278;13.921667;538;1;E;Europe/Stockholm
703;Kristianstad;Kristianstad;Sweden;KID;ESMK;55.921686;14.085536;76;1;E;Europe/Stockholm
704;Landskrona;Landskrona;Sweden;JLD;ESML;55.944444;12.869444;194;1;E;Europe/Stockholm
705;Oskarshamn;Oskarshamn;Sweden;OSK;ESMO;57.350453;16.497972;96;1;E;Europe/Stockholm
706;Anderstorp;Anderstorp;Sweden;;ESMP;57.264167;13.599439;507;1;E;Europe/Stockholm
707;Kalmar;Kalkmar;Sweden;KLR;ESMQ;56.685531;16.287578;17;1;E;Europe/Stockholm
708;Sturup;Malmoe;Sweden;MMX;ESMS;55.530193;13.371639;236;1;E;Europe/Stockholm
709;Halmstad;Halmstad;Sweden;HAD;ESMT;56.691128;12.820211;101;1;E;Europe/Stockholm
710;Hagshult;Hagshult;Sweden;;ESMV;57.292222;14.137222;556;1;E;Europe/Stockholm
711;Kronoberg;Vaxjo;Sweden;VXO;ESMX;56.929144;14.727994;610;1;E;Europe/Stockholm
712;Hallviken;Hallviken;Sweden;;ESNA;63.738333;15.458333;1119;1;E;Europe/Stockholm
713;Hedlanda;Hede;Sweden;;ESNC;62.408889;13.747222;1460;1;E;Europe/Stockholm
714;Sveg;Sveg;Sweden;EVG;ESND;62.047811;14.42295;1178;1;E;Europe/Stockholm
715;Gallivare;Gallivare;Sweden;GEV;ESNG;67.132408;20.814636;1027;1;E;Europe/Stockholm
716;Hudiksvall;Hudiksvall;Sweden;HUV;ESNH;61.768092;17.080719;95;1;E;Europe/Stockholm
717;Jokkmokk;Jokkmokk;Sweden;;ESNJ;66.496236;20.147181;904;1;E;Europe/Stockholm
718;Kramfors Solleftea;Kramfors;Sweden;KRF;ESNK;63.048597;17.768856;34;1;E;Europe/Stockholm
719;Lycksele;Lycksele;Sweden;LYC;ESNL;64.548322;18.716219;705;1;E;Europe/Stockholm
720;Optand;Optand;Sweden;;ESNM;63.128611;14.802778;1236;1;E;Europe/Stockholm
721;Sundsvall Harnosand;Sundsvall;Sweden;SDL;ESNN;62.528125;17.443928;16;1;E;Europe/Stockholm
722;Ornskoldsvik;Ornskoldsvik;Sweden;OER;ESNO;63.408339;18.990039;354;1;E;Europe/Stockholm
723;Pitea;Pitea;Sweden;;ESNP;65.398333;21.260833;43;1;E;Europe/Stockholm
724;Kiruna;Kiruna;Sweden;KRN;ESNQ;67.821986;20.336764;1508;1;E;Europe/Stockholm
725;Orsa;Orsa;Sweden;;ESNR;61.190033;14.712567;683;1;E;Europe/Stockholm
726;Skelleftea;Skelleftea;Sweden;SFT;ESNS;64.624772;21.076892;157;1;E;Europe/Stockholm
727;Sattna;Sattna;Sweden;;ESNT;62.481369;17.002917;886;1;E;Europe/Stockholm
728;Umea;Umea;Sweden;UME;ESNU;63.791828;20.282758;24;1;E;Europe/Stockholm
729;Vilhelmina;Vilhelmina;Sweden;VHM;ESNV;64.579083;16.833575;1140;1;E;Europe/Stockholm
730;Arvidsjaur;Arvidsjaur;Sweden;AJR;ESNX;65.590278;19.281944;1245;1;E;Europe/Stockholm
731;Orebro;Orebro;Sweden;ORB;ESOE;59.223733;15.037956;188;1;E;Europe/Stockholm
733;Vasteras;Vasteras;Sweden;VST;ESOW;59.589444;16.633611;21;1;E;Europe/Stockholm
734;Kallax;Lulea;Sweden;LLA;ESPA;65.543758;22.121989;65;1;E;Europe/Stockholm
735;Vidsel;Vidsel;Sweden;;ESPE;65.875325;20.149917;597;1;E;Europe/Stockholm
736;Arboga;Arboga;Sweden;;ESQO;59.386585;15.924055;33;1;E;Europe/Stockholm
737;Arlanda;Stockholm;Sweden;ARN;ESSA;59.651944;17.918611;137;1;E;Europe/Stockholm
738;Bromma;Stockholm;Sweden;BMA;ESSB;59.354372;17.94165;47;1;E;Europe/Stockholm
739;Borlange;Borlange;Sweden;BLE;ESSD;60.422017;15.515211;503;1;E;Europe/Stockholm
740;Hultsfred;Hultsfred;Sweden;HLF;ESSF;57.525833;15.823333;366;1;E;Europe/Stockholm
741;Gavle;Gavle;Sweden;GVX;ESSK;60.593333;16.951389;224;1;E;Europe/Stockholm
742;Saab;Linkoeping;Sweden;LPI;ESSL;58.40615;15.680508;172;1;E;Europe/Stockholm
743;Kungsangen;Norrkoeping;Sweden;NRK;ESSP;58.586253;16.250622;32;1;E;Europe/Stockholm
745;Eskilstuna;Eskilstuna;Sweden;;ESSU;59.351078;16.7084;139;1;E;Europe/Stockholm
746;Visby;Visby;Sweden;VBY;ESSV;57.662797;18.346211;164;1;E;Europe/Stockholm
748;Kalixfors;Kalixfors;Sweden;;ESUK;67.764789;20.257228;1549;1;E;Europe/Stockholm
750;Spangdahlem Ab;Spangdahlem;Germany;SPM;ETAD;49.972667;6.6925;1197;1;E;Europe/Berlin
751;Ramstein Ab;Ramstein;Germany;RMS;ETAR;49.436911;7.600283;776;1;E;Europe/Berlin
752;Bamberg Aaf;Bamberg;Germany;;ETEJ;49.920433;10.914233;823;1;E;Europe/Berlin
753;Giebelstadt Aaf;Giebelstadt;Germany;GHF;ETEU;49.648131;9.966494;980;1;E;Europe/Berlin
754;Buckeburg;Brueckeburg;Germany;;ETHB;52.2785;9.082167;230;1;E;Europe/Berlin
755;Celle;Celle;Germany;ZCN;ETHC;52.5912;10.022133;129;1;E;Europe/Berlin
756;Rheine Bentlage;Rheine-brentlange;Germany;;ETHE;52.291167;7.387;129;1;E;Europe/Berlin
757;Fritzlar;Fritzlar;Germany;;ETHF;51.1145;9.285833;566;1;E;Europe/Berlin
758;Laupheim;Laupheim;Germany;;ETHL;48.220297;9.910019;1766;1;E;Europe/Berlin
759;Mendig;Mendig;Germany;;ETHM;50.366;7.315333;597;1;E;Europe/Berlin
760;Niederstetten;Niederstetten;Germany;;ETHN;49.391833;9.958167;1536;1;E;Europe/Berlin
761;Roth;Roth;Germany;;ETHR;49.2175;11.100167;1268;1;E;Europe/Berlin
762;Fassberg;Fassberg;Germany;;ETHS;52.919406;10.197528;245;1;E;Europe/Berlin
763;Grafenwohr Aaf;Grafenwoehr;Germany;;ETIC;49.698686;11.940175;1363;1;E;Europe/Berlin
764;Hanau Aaf;Hanau;Germany;ZNF;ETID;50.169189;8.961586;368;1;E;Europe/Berlin
765;Hohenfels Aaf;Hohenfels;Germany;;ETIH;49.218056;11.836111;1455;1;E;Europe/Berlin
766;Kitzingen Aaf;Kitzingen;Germany;;ETIN;49.743057;10.200556;689;1;E;Europe/Berlin
767;Nordholz;Nordholz;Germany;;ETMN;53.767667;8.6585;74;1;E;Europe/Berlin
768;Diepholz;Diepholz;Germany;;ETND;52.585514;8.341014;127;1;E;Europe/Berlin
769;Geilenkirchen;Geilenkirchen;Germany;GKE;ETNG;50.960817;6.042422;296;1;E;Europe/Berlin
770;Hohn;Hohn;Germany;;ETNH;54.312167;9.538167;39;1;E;Europe/Berlin
771;Jever;Jever;Germany;;ETNJ;53.5335;7.888667;24;1;E;Europe/Berlin
772;Laage;Laage;Germany;RLG;ETNL;53.918167;12.278333;138;1;E;Europe/Berlin
773;Norvenich;Noervenich;Germany;;ETNN;50.831167;6.658167;386;1;E;Europe/Berlin
774;Schleswig;Schleswig;Germany;;ETNS;54.459333;9.516333;70;1;E;Europe/Berlin
775;Wittmundhafen;Wittmundhafen;Germany;;ETNT;53.547833;7.667333;26;1;E;Europe/Berlin
776;Neubrandenburg;Neubrandenburg;Germany;;ETNU;53.602167;13.306;228;1;E;Europe/Berlin
777;Wunstorf;Wunstorf;Germany;;ETNW;52.457333;9.427167;187;1;E;Europe/Berlin
778;Vilseck Aaf;Vilseck;Germany;;ETOI;49.63361;11.767222;1353;1;E;Europe/Berlin
779;Coleman Aaf;Coleman;Germany;;ETOR;49.563569;8.463392;309;1;E;Europe/Berlin
780;Wiesbaden Aaf;Wiesbaden;Germany;;ETOU;50.049819;8.325397;461;1;E;Europe/Berlin
781;Landsberg Lech;Landsberg;Germany;;ETSA;48.0705;10.906;2044;1;E;Europe/Berlin
782;Buchel;Buechel;Germany;;ETSB;50.173833;7.063333;1568;1;E;Europe/Berlin
783;Erding;Erding;Germany;;ETSE;48.322333;11.948667;1515;1;E;Europe/Berlin
784;Furstenfeldbruck;Fuerstenfeldbruck;Germany;FEL;ETSF;48.205667;11.267;1703;1;E;Europe/Berlin
785;Holzdorf;Holzdorf;Germany;;ETSH;51.767833;13.167667;265;1;E;Europe/Berlin
786;Ingolstadt Manching;Ingolstadt;Germany;;ETSI;48.715667;11.534;1202;1;E;Europe/Berlin
787;Lechfeld;Lechfeld;Germany;;ETSL;48.1855;10.861167;1822;1;E;Europe/Berlin
788;Neuburg;Neuburg;Germany;;ETSN;48.711;11.2115;1249;1;E;Europe/Berlin
789;Gutersloh;Guetersloh;Germany;GUT;ETUO;51.922833;8.306333;236;1;E;Europe/Berlin
790;Alexander Bay;Alexander Bay;South Africa;ALJ;FAAB;-28.575001;16.533333;98;2;U;Africa/Johannesburg
791;Aggeneys;Aggeneys;South Africa;AGZ;FAAG;-29.281767;18.813869;2648;2;U;Africa/Johannesburg
792;Brakpan;Brakpan;South Africa;;FABB;-26.23865;28.301769;5300;2;U;Africa/Johannesburg
793;Bhisho;Bisho;South Africa;BIY;FABE;-32.89715;27.279111;1950;2;U;Africa/Johannesburg
794;Bloemfontein Intl;Bloemfontein;South Africa;BFN;FABL;-29.092722;26.302444;4458;2;U;Africa/Johannesburg
795;Bethlehem;Bethlehem;South Africa;;FABM;-28.248392;28.336125;5561;2;U;Africa/Johannesburg
796;Bothaville;Bothaville;South Africa;;FABO;-27.366769;26.629194;4236;2;U;Africa/Johannesburg
797;Cape Town Intl;Cape Town;South Africa;CPT;FACT;-33.964806;18.601667;151;2;U;Africa/Johannesburg
798;Calvinia;Calvinia;South Africa;;FACV;-31.500278;19.725897;3250;2;U;Africa/Johannesburg
799;Durban Intl;Durban;South Africa;DUR;FADN;-29.970089;30.950519;33;2;U;Africa/Johannesburg
800;East London;East London;South Africa;ELS;FAEL;-33.035569;27.825939;435;2;U;Africa/Johannesburg
801;Ermelo;Ermelo;South Africa;;FAEO;-26.495644;29.979764;5700;2;U;Africa/Johannesburg
802;Ficksburg Sentraoes;Ficksburg;South Africa;;FAFB;-28.823119;27.9089;5315;2;U;Africa/Johannesburg
803;Grand Central;Johannesburg;South Africa;GCJ;FAGC;-25.986267;28.140061;5325;2;U;Africa/Johannesburg
804;George;George;South Africa;GRJ;FAGG;-34.005553;22.378889;648;2;U;Africa/Johannesburg
806;Graaff Reinet;Graaff Reinet;South Africa;;FAGR;-32.193611;24.541389;2604;2;U;Africa/Johannesburg
807;Grahamstown;Grahamstown;South Africa;;FAGT;-33.284721;26.498083;2135;2;U;Africa/Johannesburg
808;Greytown;Greytown;South Africa;;FAGY;-29.122011;30.586706;3531;2;U;Africa/Johannesburg
809;Harmony;Harmony;South Africa;;FAHA;-28.078694;26.861178;4399;2;U;Africa/Johannesburg
810;Harrismith;Harrismith;South Africa;;FAHR;-28.235072;29.106206;5585;2;U;Africa/Johannesburg
811;Hoedspruit Afb;Hoedspruit;South Africa;HDS;FAHS;-24.368642;31.048744;1743;2;U;Africa/Johannesburg
812;Gariep Dam;Hendrik Verwoerddam;South Africa;;FAHV;-30.562164;25.528286;4176;2;U;Africa/Johannesburg
813;Johannesburg Intl;Johannesburg;South Africa;JNB;FAJS;-26.139166;28.246;5558;2;U;Africa/Johannesburg
814;P C Pelser;Klerksdorp;South Africa;;FAKD;-26.871064;26.718003;4444;2;U;Africa/Johannesburg
815;Kimberley;Kimberley;South Africa;KIM;FAKM;-28.802834;24.765167;3950;2;U;Africa/Johannesburg
816;Krugersdorp;Krugersdorp;South Africa;;FAKR;-26.080978;27.725667;5499;2;U;Africa/Johannesburg
817;Kroonstad;Kroonstad;South Africa;;FAKS;-27.660617;27.315761;4700;2;U;Africa/Johannesburg
818;Johan Pienaar;Kuruman;South Africa;;FAKU;-27.456667;23.411388;4382;2;U;Africa/Johannesburg
819;Kleinsee;Kleinsee;South Africa;KLZ;FAKZ;-29.688403;17.094006;270;2;U;Africa/Johannesburg
820;Lanseria;Johannesburg;South Africa;HLA;FALA;-25.938514;27.926133;4517;2;U;Africa/Johannesburg
821;Lichtenburg;Lichtenburg;South Africa;;FALI;-26.175672;26.184575;4875;2;U;Africa/Johannesburg
822;Makhado Afb;Lambertsbaai;South Africa;;FALM;-23.159911;29.696544;3069;2;U;Africa/Johannesburg
823;Langebaanweg;Langebaanweg;South Africa;;FALW;-32.968889;18.160278;108;2;U;Africa/Johannesburg
824;Ladysmith;Ladysmith;South Africa;LAY;FALY;-28.581667;29.749722;3548;2;U;Africa/Johannesburg
825;Middelburg;Middelburg;South Africa;;FAMB;-25.684775;29.440158;4886;2;U;Africa/Johannesburg
827;Margate;Margate;South Africa;MGH;FAMG;-30.857408;30.343019;495;2;U;Africa/Johannesburg
828;Marble Hall;Marble Hall;South Africa;;FAMI;-24.989114;29.283122;2980;2;U;Africa/Johannesburg
829;Majuba Power Station;Majuba Power Station;South Africa;;FAMJ;-27.079253;29.778528;5600;2;U;Africa/Johannesburg
6813;Susse;Kangia;Greenland;;\N;69.2225;-49.9383;100;-3;U;America/Godthab
831;Malelane;Malalane;South Africa;;FAMN;-25.473603;31.565828;1153;2;U;Africa/Johannesburg
832;Messina;Musina;South Africa;MEZ;FAMS;-25.704458;26.908978;4251;2;U;Africa/Johannesburg
833;Mkuzi;Mkuze;South Africa;;FAMU;-27.626086;32.044275;400;2;U;Africa/Johannesburg
834;Newcastle;Newcastle;South Africa;NCS;FANC;-27.770586;29.976894;4074;2;U;Africa/Johannesburg
835;Nylstroom;Nylstroom;South Africa;;FANY;-24.686056;28.434944;3900;2;U;Africa/Johannesburg
836;Overberg;Overberg;South Africa;;FAOB;-34.554861;20.250681;52;2;U;Africa/Johannesburg
837;Oudtshoorn;Oudtshoorn;South Africa;DUH;FAOH;-33.606967;22.188978;1063;2;U;Africa/Johannesburg
838;Port Elizabeth Intl;Port Elizabeth;South Africa;PLZ;FAPE;-33.984919;25.617275;226;2;U;Africa/Johannesburg
839;Plettenberg Bay;Plettenberg Bay;South Africa;;FAPG;-34.090279;23.327778;465;2;U;Africa/Johannesburg
840;Phalaborwa;Phalaborwa;South Africa;PHW;FAPH;-23.937166;31.15539;1432;2;U;Africa/Johannesburg
841;Polokwane International;Polokwane;South Africa;PTG;FAPI;-23.926089;29.484422;4354;2;U;Africa/Johannesburg
842;Port St Johns;Port Saint Johns;South Africa;;FAPJ;-31.605886;29.519786;1227;2;U;Africa/Johannesburg
843;Pietermaritzburg;Pietermaritzburg;South Africa;PZB;FAPM;-29.648975;30.398667;2423;2;U;Africa/Johannesburg
844;Pilanesberg Intl;Pilanesberg;South Africa;NTY;FAPN;-25.333822;27.173358;3412;2;U;Africa/Johannesburg
845;Polokwane Intl;Potgietersrus;South Africa;;FAPP;-23.845306;29.458611;4076;2;U;Africa/Johannesburg
846;Potchefstroom;Potchefstroom;South Africa;;FAPS;-26.670994;27.0819;4520;2;U;Africa/Johannesburg
847;Parys;Parys;South Africa;;FAPY;-26.889344;27.503417;4740;2;U;Africa/Johannesburg
848;Queenstown;Queenstown;South Africa;UTW;FAQT;-31.920197;26.882206;3637;2;U;Africa/Johannesburg
849;Richards Bay;Richard's Bay;South Africa;RCB;FARB;-28.741039;32.092111;109;2;U;Africa/Johannesburg
850;Rustenburg;Rustenburg;South Africa;;FARG;-25.6443;27.271119;3700;2;U;Africa/Johannesburg
851;Robertson;Robertson;South Africa;;FARS;-33.812181;19.902828;640;2;U;Africa/Johannesburg
852;Springbok;Springbok;South Africa;SBU;FASB;-29.689333;17.939611;2690;2;U;Africa/Johannesburg
853;Secunda;Secunda;South Africa;;FASC;-26.524083;29.170144;5250;2;U;Africa/Johannesburg
854;Saldanha Vredenburg;Saldanha;South Africa;;FASD;-32.964067;17.969331;50;2;U;Africa/Johannesburg
855;Springs;Springs;South Africa;;FASI;-26.248411;28.397508;5340;2;U;Africa/Johannesburg
856;Swartkop;Swartkop;South Africa;;FASK;-25.809717;28.164631;4780;2;U;Africa/Johannesburg
857;Sishen;Sishen;South Africa;SIS;FASS;-27.648606;22.999278;3848;2;U;Africa/Johannesburg
858;Hendrik Swellengrebel;Swellendam;South Africa;;FASX;-34.048222;20.474611;407;2;U;Africa/Johannesburg
859;Skukuza;Skukuza;South Africa;SZK;FASZ;-24.960944;31.588731;1020;2;U;Africa/Johannesburg
860;Tommys Fld;Tommy's Field;South Africa;;FATF;-28.260028;22.993178;4360;2;U;Africa/Johannesburg
861;New Tempe;Bloemfontein;South Africa;;FATP;-29.032928;26.157564;4547;2;U;Africa/Johannesburg
862;Tutuka Power Station;Tutuka;South Africa;;FATT;-26.776556;29.338778;5313;2;U;Africa/Johannesburg
863;Tzaneen;Tzaneen;South Africa;LTA;FATZ;-23.824417;30.329306;1914;2;U;Africa/Johannesburg
864;Prince Mangosuthu Buthelezi;Ulundi;South Africa;ULD;FAUL;-28.320586;31.416519;1720;2;U;Africa/Johannesburg
865;Upington;Upington;South Africa;UTN;FAUP;-28.399097;21.260239;2782;2;U;Africa/Johannesburg
866;Mthatha;Umtata;South Africa;UTT;FAUT;-31.547903;28.674289;2400;2;U;Africa/Johannesburg
867;Vryburg;Vryburg;South Africa;VRU;FAVB;-26.982408;24.728756;3920;2;U;Africa/Johannesburg
868;Virginia;Durban;South Africa;VIR;FAVG;-29.770606;31.058406;20;2;U;Africa/Johannesburg
869;Vredendal;Vredendal;South Africa;;FAVR;-31.640961;18.544789;330;2;U;Africa/Johannesburg
870;Vereeniging;Vereeniging;South Africa;;FAVV;-26.566372;27.960756;4846;2;U;Africa/Johannesburg
871;Wonderboom;Pretoria;South Africa;PRY;FAWB;-25.653858;28.224231;4095;2;U;Africa/Johannesburg
872;Witbank;Witbank;South Africa;;FAWI;-25.832294;29.192019;5078;2;U;Africa/Johannesburg
873;Waterkloof Afb;Waterkloof;South Africa;;FAWK;-25.83;28.2225;4940;2;U;Africa/Johannesburg
874;Welkom;Welkom;South Africa;WEL;FAWM;-27.998;26.669586;4399;2;U;Africa/Johannesburg
875;Ysterplaat;Ysterplaat;South Africa;;FAYP;-33.900244;18.498297;52;2;U;Africa/Johannesburg
876;Zeerust;Zeerust;South Africa;;FAZR;-25.598972;26.042333;4258;2;U;Africa/Johannesburg
877;Francistown;Francistown;Botswana;FRW;FBFT;-21.159597;27.474525;3283;2;U;Africa/Gaborone
878;Jwaneng;Jwaneng;Botswana;JWA;FBJW;-24.602333;24.690971;3900;2;U;Africa/Gaborone
879;Kasane;Kasane;Botswana;BBK;FBKE;-17.832875;25.1624;3289;2;U;Africa/Gaborone
880;Maun;Maun;Botswana;MUB;FBMN;-19.972564;23.431086;3093;2;U;Africa/Gaborone
881;Sir Seretse Khama Intl;Gaberone;Botswana;GBE;FBSK;-24.555225;25.918208;3299;2;U;Africa/Gaborone
882;Selebi Phikwe;Selebi-phikwe;Botswana;PKW;FBSP;-22.05835;27.828767;2925;2;U;Africa/Gaborone
883;Maya Maya;Brazzaville;Congo (Brazzaville);BZV;FCBB;-4.2517;15.253031;1048;1;N;Africa/Brazzaville
884;Owando;Owando;Congo (Kinshasa);FTX;FCOO;-0.53135;15.950094;1214;1;N;Africa/Brazzaville
885;Ouesso;Ouesso;Congo (Kinshasa);OUE;FCOU;1.615994;16.037917;1158;1;N;Africa/Brazzaville
886;Pointe Noire;Pointe-noire;Congo (Brazzaville);PNR;FCPP;-4.816028;11.886597;55;1;N;Africa/Brazzaville
887;Matsapha;Manzini;Swaziland;MTS;FDMS;-26.529022;31.307519;2075;2;U;Africa/Mbabane
888;Bangui M Poko;Bangui;Central African Republic;BGF;FEFF;4.398475;18.518786;1208;1;N;Africa/Bangui
889;Berberati;Berberati;Central African Republic;BBT;FEFT;4.221583;15.786369;1929;1;N;Africa/Bangui
890;Bata;Bata;Equatorial Guinea;BSG;FGBT;1.905469;9.805681;13;1;N;Africa/Malabo
891;Malabo;Malabo;Equatorial Guinea;SSG;FGSL;3.755267;8.708717;76;1;N;Africa/Malabo
892;Ascension Aux Af;Wide Awake;Saint Helena;;FHAW;-7.969597;-14.393664;278;0;N;Atlantic/St_Helena
893;Sir Seewoosagur Ramgoolam Intl;Plaisance;Mauritius;MRU;FIMP;-20.430235;57.6836;186;4;N;Indian/Mauritius
894;Plaine Corail;Rodriguez Island;Mauritius;RRG;FIMR;-19.757658;63.360983;95;4;N;Indian/Mauritius
895;Diego Garcia Nsf;Diego Garcia Island;British Indian Ocean Territory;;FJDG;-7.313267;72.411089;9;6;U;Indian/Chagos
896;Tiko;Tiko;Cameroon;TKC;FKKC;4.089192;9.360528;151;1;N;Africa/Douala
897;Douala;Douala;Cameroon;DLA;FKKD;4.006081;9.719481;33;1;N;Africa/Douala
898;Salak;Maroua;Cameroon;MVR;FKKL;10.451392;14.257361;1390;1;N;Africa/Douala
899;Foumban Nkounja;Foumban;Cameroon;FOM;FKKM;5.636919;10.750817;3963;1;N;Africa/Douala
900;Ngaoundere;N'gaoundere;Cameroon;NGE;FKKN;7.357011;13.559242;3655;1;N;Africa/Douala
901;Garoua;Garoua;Cameroon;GOU;FKKR;9.335892;13.370103;794;1;N;Africa/Douala
902;Bafoussam;Bafoussam;Cameroon;BFX;FKKU;5.536919;10.354583;4347;1;N;Africa/Douala
903;Bamenda;Bamenda;Cameroon;BPC;FKKV;6.039239;10.122639;4065;1;N;Africa/Douala
904;Yaounde Ville;Yaounde;Cameroon;YAO;FKKY;3.836039;11.523461;2464;1;N;Africa/Douala
905;Kasompe;Kasompe;Zambia;;FLKE;-12.572778;27.89395;4636;2;U;Africa/Lusaka
906;Livingstone;Livingstone;Zambia;LVI;FLLI;-17.821756;25.822692;3302;2;U;Africa/Lusaka
907;Lusaka Intl;Lusaka;Zambia;LUN;FLLS;-15.330817;28.452628;3779;2;U;Africa/Lusaka
908;Mfuwe;Mfuwe;Zambia;MFU;FLMF;-13.258878;31.936581;1853;2;U;Africa/Lusaka
909;Mongu;Mongu;Zambia;;FLMG;-15.254542;23.162306;3488;2;U;Africa/Lusaka
910;Ndola;Ndola;Zambia;NLA;FLND;-12.998139;28.664944;4167;2;U;Africa/Lusaka
911;Southdowns;Southdowns;Zambia;KIW;FLSO;-12.900469;28.149858;4145;2;U;Africa/Lusaka
912;Prince Said Ibrahim;Moroni;Comoros;HAH;FMCH;-11.533661;43.27185;93;3;U;Indian/Comoro
913;Bandaressalam;Moheli;Comoros;NWA;FMCI;-12.298108;43.7664;46;3;U;Indian/Comoro
914;Ouani;Anjouan;Comoros;AJN;FMCV;-12.131667;44.430279;62;3;U;Indian/Comoro
915;Dzaoudzi Pamandzi;Dzaoudzi;Mayotte;DZA;FMCZ;-12.804722;45.281113;23;3;U;Indian/Mayotte
916;St Denis Gillot;St.-denis;Reunion;RUN;FMEE;-20.8871;55.510308;66;4;U;Indian/Reunion
917;St Pierre Pierrefonds;St.-pierre;Reunion;ZSE;FMEP;-21.320039;55.423581;56;4;U;Indian/Reunion
918;Ivato;Antananarivo;Madagascar;TNR;FMMI;-18.79695;47.478806;4198;3;U;Indian/Antananarivo
919;Miandrivazo;Miandrivazo;Madagascar;ZVA;FMMN;-19.562778;45.450832;203;3;U;Indian/Antananarivo
920;Sainte Marie;Sainte Marie;Madagascar;SMS;FMMS;-17.093889;49.815834;7;3;U;Indian/Antananarivo
921;Toamasina;Toamasina;Madagascar;TMM;FMMT;-18.109517;49.392536;22;3;U;Indian/Antananarivo
922;Morondava;Morondava;Madagascar;MOQ;FMMV;-20.28475;44.317614;30;3;U;Indian/Antananarivo
923;Arrachart;Antsiranana;Madagascar;DIE;FMNA;-12.3494;49.291747;374;3;U;Indian/Antananarivo
924;Avaratra;Mananara;Madagascar;WMR;FMNC;-16.1639;49.773753;9;3;U;Indian/Antananarivo
925;Andapa;Andapa;Madagascar;ZWA;FMND;-14.651667;49.620556;1552;3;U;Indian/Antananarivo
926;Ambilobe;Ambilobe;Madagascar;AMB;FMNE;-13.188431;48.987978;72;3;U;Indian/Antananarivo
927;Antsirabato;Antalaha;Madagascar;ANM;FMNH;-14.999411;50.320233;20;3;U;Indian/Antananarivo
928;Analalava;Analalava;Madagascar;HVA;FMNL;-14.629694;47.763783;345;3;U;Indian/Antananarivo
929;Philibert Tsiranana;Mahajanga;Madagascar;MJN;FMNM;-15.667144;46.351828;87;3;U;Indian/Antananarivo
930;Fascene;Nosy-be;Madagascar;NOS;FMNN;-13.312067;48.314822;36;3;U;Indian/Antananarivo
931;Besalampy;Besalampy;Madagascar;BPY;FMNQ;-16.741945;44.481388;125;3;U;Indian/Antananarivo
932;Maroantsetra;Maroantsetra;Madagascar;WMN;FMNR;-15.436666;49.688332;13;3;U;Indian/Antananarivo
933;Sambava;Sambava;Madagascar;SVB;FMNS;-14.278611;50.174721;20;3;U;Indian/Antananarivo
934;Vohimarina;Vohemar;Madagascar;VOH;FMNV;-13.375834;50.002777;19;3;U;Indian/Antananarivo
935;Ambalabe;Antsohihy;Madagascar;WAI;FMNW;-14.89875;47.993894;92;3;U;Indian/Antananarivo
936;Ampampamena;Ampampamena;Madagascar;;FMNZ;-13.484814;48.632739;49;3;U;Indian/Antananarivo
937;Tolagnaro;Tolagnaro;Madagascar;FTU;FMSD;-25.038056;46.956111;29;3;U;Indian/Antananarivo
938;Fianarantsoa;Fianarantsoa;Madagascar;WFI;FMSF;-21.441558;47.111736;3658;3;U;Indian/Antananarivo
939;Farafangana;Farafangana;Madagascar;RVA;FMSG;-22.805286;47.820614;26;3;U;Indian/Antananarivo
940;Manakara;Manakara;Madagascar;WVK;FMSK;-22.119722;48.021667;33;3;U;Indian/Antananarivo
941;Mananjary;Mananjary;Madagascar;MNJ;FMSM;-21.201772;48.358317;20;3;U;Indian/Antananarivo
942;Morombe;Morombe;Madagascar;MXM;FMSR;-21.753867;43.375533;16;3;U;Indian/Antananarivo
943;Toliara;Toliara;Madagascar;TLE;FMST;-23.383369;43.728453;29;3;U;Indian/Antananarivo
944;Mbanza Congo;M'banza-congo;Angola;SSY;FNBC;-6.269897;14.247025;1860;1;N;Africa/Luanda
945;Benguela;Benguela;Angola;BUG;FNBG;-12.609025;13.403711;118;1;N;Africa/Luanda
946;Cabinda;Cabinda;Angola;CAB;FNCA;-5.596992;12.188353;66;1;N;Africa/Luanda
6814;Culebra Airport;Culebra Island;Puerto Rico;CPX;TJCP;18.3127;-65.3034;12;-4;U;America/Puerto_Rico
948;Huambo;Huambo;Angola;NOV;FNHU;-12.808878;15.760547;5587;1;N;Africa/Luanda
949;Kuito;Kuito;Angola;SVP;FNKU;-12.404633;16.947414;5618;1;N;Africa/Luanda
950;Lobito;Lobito;Angola;;FNLB;-12.371233;13.536625;10;1;N;Africa/Luanda
951;Luanda 4 De Fevereiro;Luanda;Angola;LAD;FNLU;-8.858375;13.231178;243;1;N;Africa/Luanda
952;Malanje;Malanje;Angola;MEG;FNMA;-9.525086;16.312406;3868;1;N;Africa/Luanda
953;Menongue;Menongue;Angola;SPP;FNME;-14.657583;17.719833;4469;1;N;Africa/Luanda
955;Negage;Negage;Angola;GXG;FNNG;-7.754506;15.287728;4105;1;N;Africa/Luanda
956;Porto Amboim;Porto Amboim;Angola;PBN;FNPA;-10.721956;13.765528;16;1;N;Africa/Luanda
957;Saurimo;Saurimo;Angola;VHC;FNSA;-9.689067;20.431875;3584;1;N;Africa/Luanda
958;Soyo;Soyo;Angola;SZA;FNSO;-6.141086;12.371764;15;1;N;Africa/Luanda
959;Lubango;Lubango;Angola;SDD;FNUB;-14.924733;13.575022;5778;1;N;Africa/Luanda
960;Luena;Luena;Angola;LUO;FNUE;-11.768086;19.897672;4360;1;N;Africa/Luanda
961;Uige;Uige;Angola;UGO;FNUG;-7.603067;15.027822;2720;1;N;Africa/Luanda
962;Xangongo;Xangongo;Angola;XGN;FNXA;-16.755417;14.965344;3635;1;N;Africa/Luanda
963;Oyem;Oyem;Gabon;OYE;FOGO;1.543108;11.581361;2158;1;N;Africa/Libreville
964;Okondja;Okondja;Gabon;OKN;FOGQ;-0.665214;13.673133;1325;1;N;Africa/Libreville
965;Lambarene;Lambarene;Gabon;LBQ;FOGR;-0.704389;10.245722;82;1;N;Africa/Libreville
966;Bitam;Bitam;Gabon;BMM;FOOB;2.075639;11.493195;1969;1;N;Africa/Libreville
967;Port Gentil;Port Gentil;Gabon;POG;FOOG;-0.711739;8.754383;13;1;N;Africa/Libreville
968;Omboue Hopital;Omboue Hospial;Gabon;OMB;FOOH;-1.574733;9.262694;33;1;N;Africa/Libreville
969;Makokou;Makokou;Gabon;MKU;FOOK;0.579211;12.890908;1726;1;N;Africa/Libreville
970;Leon M Ba;Libreville;Gabon;LBV;FOOL;0.4586;9.412283;39;1;N;Africa/Libreville
971;Mvengue;Franceville;Gabon;MVB;FOON;-1.656156;13.438036;1450;1;N;Africa/Libreville
972;Principe;Principe;Sao Tome and Principe;PCP;FPPR;1.662936;7.411742;591;0;N;Africa/Sao_Tome
973;Sao Tome Intl;Sao Tome;Sao Tome and Principe;TMS;FPST;0.378175;6.712153;33;0;N;Africa/Sao_Tome
974;Beira;Beira;Mozambique;BEW;FQBR;-19.796419;34.907556;33;2;U;Africa/Maputo
976;Inhambane;Inhambane;Mozambique;INH;FQIN;-23.876431;35.408544;30;2;U;Africa/Maputo
977;Lichinga;Lichinga;Mozambique;VXC;FQLC;-13.273986;35.266262;4505;2;U;Africa/Maputo
978;Lumbo;Lumbo;Mozambique;;FQLU;-15.033058;40.671728;33;2;U;Africa/Maputo
979;Maputo;Maputo;Mozambique;MPM;FQMA;-25.920836;32.572606;145;2;U;Africa/Maputo
980;Mueda;Mueda;Mozambique;;FQMD;-11.672922;39.563142;2789;2;U;Africa/Maputo
981;Mocimboa Da Praia;Mocimboa Da Praia;Mozambique;MZB;FQMP;-11.361789;40.354875;89;2;U;Africa/Maputo
982;Marrupa;Marrupa;Mozambique;;FQMR;-13.225053;37.552067;2480;2;U;Africa/Maputo
983;Nacala;Nacala;Mozambique;MNC;FQNC;-14.488233;40.71225;410;2;U;Africa/Maputo
984;Nampula;Nampula;Mozambique;APL;FQNP;-15.105611;39.2818;1444;2;U;Africa/Maputo
985;Pemba;Pemba;Mozambique;POL;FQPB;-12.986753;40.522492;331;2;U;Africa/Maputo
986;Quelimane;Quelimane;Mozambique;UEL;FQQL;-17.8555;36.869106;36;2;U;Africa/Maputo
987;Songo;Songo;Mozambique;;FQSG;-15.602694;32.773189;2904;2;U;Africa/Maputo
988;Tete Chingodzi;Tete;Mozambique;TET;FQTT;-16.104817;33.640181;525;2;U;Africa/Maputo
989;Ulongwe;Ulongwe;Mozambique;;FQUG;-14.704617;34.352369;4265;2;U;Africa/Maputo
990;Vilankulo;Vilankulu;Mozambique;VNX;FQVL;-22.018431;35.313297;46;2;U;Africa/Maputo
991;Alphonse;Alphonse;Seychelles;;FSAL;-7.004783;52.726247;10;4;U;Indian/Mahe
992;Desroches;Desroches;Seychelles;DES;FSDR;-5.696697;53.655844;10;4;U;Indian/Mahe
993;Farquhar;Farquhar;Seychelles;;FSFA;-10.109611;51.176119;10;4;U;Indian/Mahe
994;Seychelles Intl;Mahe;Seychelles;SEZ;FSIA;-4.674342;55.521839;10;4;U;Indian/Mahe
995;Praslin;Praslin;Seychelles;PRI;FSPP;-4.319292;55.691417;10;4;U;Indian/Mahe
996;Coetivy;Coetivy;Seychelles;;FSSC;-7.134567;56.278239;10;4;U;Indian/Mahe
997;Abeche;Abeche;Chad;AEH;FTTC;13.847;20.844333;1788;1;N;Africa/Ndjamena
998;Moundou;Moundou;Chad;MQQ;FTTD;8.624406;16.071419;1407;1;N;Africa/Ndjamena
999;Ndjamena Hassan Djamous;N'djamena;Chad;NDJ;FTTJ;12.133689;15.034019;968;1;N;Africa/Ndjamena
1000;Faya Largeau;Faya-largeau;Chad;FYT;FTTY;17.917053;19.111078;771;1;N;Africa/Ndjamena
1001;J M Nkomo Intl;Bulawayo;Zimbabwe;BUQ;FVBU;-20.017431;28.617869;4359;2;U;Africa/Harare
1002;Charles Prince;Harare;Zimbabwe;;FVCP;-17.751561;30.924706;4845;2;U;Africa/Harare
1003;Buffalo Range;Chiredzi;Zimbabwe;BFO;FVCZ;-21.008083;31.57855;1421;2;U;Africa/Harare
1004;Victoria Falls Intl;Victoria Falls;Zimbabwe;VFA;FVFA;-18.095881;25.839006;3490;2;U;Africa/Harare
1005;Harare Intl;Harare;Zimbabwe;HRE;FVHA;-17.931806;31.092847;4887;2;U;Africa/Harare
1006;Kariba Intl;Kariba;Zimbabwe;KAB;FVKB;-16.519761;28.884981;1706;2;U;Africa/Harare
1007;Mutoko;Mutoko;Zimbabwe;;FVMT;-17.431917;32.184502;3950;2;U;Africa/Harare
1008;Mutare;Mutare;Zimbabwe;;FVMU;-18.997499;32.627224;3410;2;U;Africa/Harare
1009;Masvingo Intl;Masvingo;Zimbabwe;MVZ;FVMV;-20.055333;30.859111;3595;2;U;Africa/Harare
1010;Zvishavane;Zvishavane;Zimbabwe;;FVSH;-20.289497;30.088228;3012;2;U;Africa/Harare
1011;Gweru Thornhill;Gwert;Zimbabwe;GWE;FVTL;-19.436394;29.861911;4680;2;U;Africa/Harare
1012;Hwange National Park;Hwange National Park;Zimbabwe;WKM;FVWN;-18.629872;27.021042;3543;2;U;Africa/Harare
1013;Chileka Intl;Blantyre;Malawi;BLZ;FWCL;-15.679053;34.974014;2555;2;U;Africa/Blantyre
1014;Karonga;Karonga;Malawi;KGJ;FWKA;-9.953569;33.893022;1765;2;U;Africa/Blantyre
1015;Kasungu;Kasungu;Malawi;;FWKG;-13.014631;33.468597;3470;2;U;Africa/Blantyre
1016;Kamuzu Intl;Lilongwe;Malawi;LLW;FWKI;-13.789378;33.781;4035;2;U;Africa/Blantyre
1017;Mzuzu;Mzuzu;Malawi;ZZU;FWUU;-11.44475;34.011776;4115;2;U;Africa/Blantyre
1018;Moshoeshoe I Intl;Maseru;Lesotho;MSU;FXMM;-29.462256;27.552503;5348;2;U;Africa/Maseru
1019;Mejametalana;Maseru;Lesotho;;FXMU;-29.304053;27.503458;5105;2;U;Africa/Maseru
1020;Ndjili Intl;Kinshasa;Congo (Kinshasa);FIH;FZAA;-4.38575;15.444569;1027;1;N;Africa/Kinshasa
1021;Ndolo;Kinshasa;Congo (Kinshasa);NLO;FZAB;-4.326689;15.327342;915;1;N;Africa/Kinshasa
1022;Muanda;Muanda;Congo (Kinshasa);MNB;FZAG;-5.930858;12.351789;89;1;N;Africa/Kinshasa
1023;Kitona Base;Kitona Base;Congo (Kinshasa);;FZAI;-5.918056;12.447694;394;1;N;Africa/Kinshasa
1024;Bandundu;Bandoundu;Congo (Kinshasa);FDU;FZBO;-3.311319;17.381683;1053;1;N;Africa/Kinshasa
1025;Kikwit;Kikwit;Congo (Kinshasa);KKW;FZCA;-5.035767;18.785631;1572;1;N;Africa/Kinshasa
1026;Mbandaka;Mbandaka;Congo (Kinshasa);MDK;FZEA;0.0226;18.288744;1040;1;N;Africa/Kinshasa
1027;Gbadolite;Gbadolite;Congo (Kinshasa);BDT;FZFD;4.253206;20.975283;1509;1;N;Africa/Kinshasa
1028;Gemena;Gemena;Congo (Kinshasa);GMA;FZFK;3.235369;19.771258;1378;1;N;Africa/Kinshasa
1029;Kotakoli;Kotakoli;Congo (Kinshasa);;FZFP;4.157639;21.650917;1801;1;N;Africa/Kinshasa
1030;Lisala;Lisala;Congo (Kinshasa);LIQ;FZGA;2.170658;21.496906;1509;1;N;Africa/Kinshasa
1031;Kisangani Simisini;Kisangani;Congo (Kinshasa);FKI;FZIA;0.5175;25.155014;1289;2;U;Africa/Lubumbashi
1032;Matari;Isiro;Congo (Kinshasa);IRP;FZJH;2.827606;27.588253;2438;2;U;Africa/Lubumbashi
1033;Bunia;Bunia;Congo (Kinshasa);BUX;FZKA;1.565719;30.220833;4045;2;U;Africa/Lubumbashi
1034;Buta Zega;Buta Zega;Congo (Kinshasa);;FZKJ;2.818347;24.793706;1378;2;U;Africa/Lubumbashi
1035;Bukavu Kavumu;Bukavu/kavumu;Congo (Kinshasa);BKY;FZMA;-2.308978;28.808803;5643;2;U;Africa/Lubumbashi
1036;Goma;Goma;Congo (Kinshasa);GOM;FZNA;-1.670814;29.238464;5089;2;U;Africa/Kigali
1037;Kindu;Kindu;Congo (Kinshasa);KND;FZOA;-2.919178;25.915361;1630;2;U;Africa/Lubumbashi
1038;Lubumbashi Intl;Lubumashi;Congo (Kinshasa);FBM;FZQA;-11.591333;27.530889;4295;2;U;Africa/Lubumbashi
1039;Kolwezi;Kolwezi;Congo (Kinshasa);KWZ;FZQM;-10.765886;25.505714;5007;2;U;Africa/Lubumbashi
1040;Kalemie;Kalemie;Congo (Kinshasa);FMI;FZRF;-5.875556;29.25;2569;2;U;Africa/Lubumbashi
1041;Kamina Base;Kamina Base;Congo (Kinshasa);KMN;FZSA;-8.642025;25.252897;3543;2;U;Africa/Lubumbashi
1042;Kananga;Kananga;Congo (Kinshasa);KGA;FZUA;-5.900055;22.469166;2139;2;U;Africa/Lubumbashi
1043;Mbuji Mayi;Mbuji-mayi;Congo (Kinshasa);MJM;FZWA;-6.121236;23.569008;2221;2;U;Africa/Lubumbashi
1044;Senou;Bamako;Mali;BKO;GABS;12.533544;-7.949944;1247;0;N;Africa/Bamako
1045;Gao;Gao;Mali;GAQ;GAGO;16.248433;-0.005456;870;0;N;Africa/Bamako
1046;Kayes Dag Dag;Kayes;Mali;KYS;GAKY;14.481233;-11.404397;164;0;N;Africa/Bamako
1047;Ambodedjo;Mopti;Mali;MZI;GAMB;14.512803;-4.079561;906;0;N;Africa/Bamako
1048;Tombouctou;Tombouctou;Mali;TOM;GATB;16.730458;-3.007583;863;0;N;Africa/Bamako
1049;Tessalit;Tessalit;Mali;;GATS;20.242983;0.977308;1621;0;N;Africa/Bamako
1050;Banjul Intl;Banjul;Gambia;BJL;GBYD;13.337961;-16.652206;95;0;N;Africa/Banjul
1051;Fuerteventura;Fuerteventura;Spain;FUE;GCFV;28.452717;-13.863761;83;0;E;Atlantic/Canary
1052;Hierro;Hierro;Spain;VDE;GCHI;27.814847;-17.887056;103;0;E;Atlantic/Canary
1053;La Palma;Santa Cruz De La Palma;Spain;SPC;GCLA;28.626478;-17.755611;107;0;E;Atlantic/Canary
1054;Gran Canaria;Gran Canaria;Spain;LPA;GCLP;27.931886;-15.386586;78;0;E;Atlantic/Canary
1055;Lanzarote;Las Palmas;Spain;ACE;GCRR;28.945464;-13.605225;47;0;E;Atlantic/Canary
1056;Tenerife Sur;Tenerife;Spain;TFS;GCTS;28.044475;-16.572489;209;0;E;Atlantic/Canary
1057;Tenerife Norte;Tenerife;Spain;TFN;GCXO;28.482653;-16.341536;2073;0;E;Atlantic/Canary
1058;Melilla;Melilla;Spain;MLN;GEML;35.279817;-2.956256;156;1;E;Europe/Madrid
1059;Freetown Lungi;Freetown;Sierra Leone;FNA;GFLL;8.616444;-13.195489;84;0;N;Africa/Freetown
1060;Cufar;Cufar;Guinea-Bissau;;GGCF;11.287917;-15.1805;65;0;N;Africa/Bissau
1062;Monrovia Spriggs Payne;Monrovia;Liberia;MLW;GLMR;6.289061;-10.758722;25;0;N;Africa/Monrovia
1063;Monrovia Roberts Intl;Monrovia;Liberia;ROB;GLRB;6.233789;-10.362311;31;0;N;Africa/Monrovia
1064;Inezgane;Agadir;Morocco;AGA;GMAA;30.381353;-9.546311;89;0;N;Africa/Casablanca
1065;Plage Blanche;Tan Tan;Morocco;TTA;GMAT;28.448194;-11.161347;653;0;N;Africa/Casablanca
1066;Saiss;Fes;Morocco;FEZ;GMFF;33.927261;-4.977958;1900;0;N;Africa/Casablanca
1067;Ifrane;Ifrane;Morocco;;GMFI;33.505319;-5.152903;5459;0;N;Africa/Casablanca
1068;Moulay Ali Cherif;Er-rachidia;Morocco;ERH;GMFK;31.9475;-4.398333;3428;0;N;Africa/Casablanca
1069;Bassatine;Meknes;Morocco;MEK;GMFM;33.879067;-5.515125;1890;0;N;Africa/Casablanca
1070;Angads;Oujda;Morocco;OUD;GMFO;34.78715;-1.923986;1535;0;N;Africa/Casablanca
1071;Ben Slimane;Ben Slimane;Morocco;;GMMB;33.655422;-7.22145;627;0;N;Africa/Casablanca
1072;Sale;Rabat;Morocco;RBA;GMME;34.051467;-6.751519;276;0;N;Africa/Casablanca
1074;Mohammed V Intl;Casablanca;Morocco;CMN;GMMN;33.367467;-7.589967;656;0;N;Africa/Casablanca
1075;Menara;Marrakech;Morocco;RAK;GMMX;31.606886;-8.0363;1545;0;N;Africa/Casablanca
1076;Kenitra;Kentira;Morocco;NNA;GMMY;34.298864;-6.595878;16;0;N;Africa/Casablanca
1077;Ouarzazate;Ouarzazate;Morocco;OZZ;GMMZ;30.939053;-6.909431;3782;0;N;Africa/Casablanca
1078;Cherif El Idrissi;Al Hociema;Morocco;AHU;GMTA;35.177103;-3.839525;95;0;N;Africa/Casablanca
1079;Saniat Rmel;Tetouan;Morocco;TTU;GMTN;35.594333;-5.320019;10;0;N;Africa/Casablanca
1080;Ibn Batouta;Tanger;Morocco;TNG;GMTT;35.726917;-5.916889;62;0;N;Africa/Casablanca
1081;Ziguinchor;Ziguinchor;Senegal;ZIG;GOGG;12.555617;-16.281783;75;0;N;Africa/Dakar
1082;Cap Skiring;Cap Skiring;Senegal;CSK;GOGS;12.4102;-16.746125;52;0;N;Africa/Dakar
1083;Kaolack;Kaolack;Senegal;KLC;GOOK;14.146881;-16.051297;26;0;N;Africa/Dakar
1084;Leopold Sedar Senghor Intl;Dakar;Senegal;DKR;GOOY;14.739708;-17.490225;85;0;N;Africa/Dakar
1085;Saint Louis;St. Louis;Senegal;XLS;GOSS;16.050761;-16.463172;9;0;N;Africa/Dakar
1086;Bakel;Bakel;Senegal;BXE;GOTB;14.847256;-12.468264;98;0;N;Africa/Dakar
1087;Kedougou;Kedougou;Senegal;KGG;GOTK;12.572292;-12.220333;584;0;N;Africa/Dakar
1088;Tambacounda;Tambacounda;Senegal;TUD;GOTT;13.736817;-13.653122;161;0;N;Africa/Dakar
1089;Aioun El Atrouss;Aioun El Atrouss;Mauritania;IEO;GQNA;16.711294;-9.637883;951;0;N;Africa/Nouakchott
1090;Tidjikja;Tidjikja;Mauritania;TIY;GQND;18.570103;-11.423547;1316;0;N;Africa/Nouakchott
1091;Kiffa;Kiffa;Mauritania;KFA;GQNF;16.589983;-11.406208;423;0;N;Africa/Nouakchott
1092;Nema;Nema;Mauritania;EMN;GQNI;16.622;-7.316567;758;0;N;Africa/Nouakchott
1093;Kaedi;Kaedi;Mauritania;KED;GQNK;16.159547;-13.507617;75;0;N;Africa/Nouakchott
1094;Nouakchott;Nouakschott;Mauritania;NKC;GQNN;18.097856;-15.947956;7;0;N;Africa/Nouakchott
1095;Selibady;Selibabi;Mauritania;SEY;GQNS;15.179692;-12.207272;262;0;N;Africa/Nouakchott
1096;Atar;Atar;Mauritania;ATR;GQPA;20.506828;-13.043194;758;0;N;Africa/Nouakchott
1097;Nouadhibou;Nouadhibou;Mauritania;NDB;GQPP;20.933067;-17.029956;16;0;N;Africa/Nouakchott
1098;Bir Moghrein;Bir Moghrein;Mauritania;;GQPT;25.236697;-11.588697;1194;0;N;Africa/Nouakchott
1099;Fria;Fira;Guinea;FIG;GUFA;10.350556;-13.569167;499;0;N;Africa/Conakry
1100;Faranah;Faranah;Guinea;FAA;GUFH;10.035467;-10.769825;1476;0;N;Africa/Conakry
1101;Labe;Labe;Guinea;LEK;GULB;11.326058;-12.28685;3396;0;N;Africa/Conakry
1102;Amilcar Cabral Intl;Amilcar Cabral;Cape Verde;SID;GVAC;16.741389;-22.949444;177;-1;U;Atlantic/Cape_Verde
1103;Rabil;Boa Vista;Cape Verde;BVC;GVBA;16.136531;-22.888897;69;-1;U;Atlantic/Cape_Verde
1104;Maio;Maio;Cape Verde;MMO;GVMA;15.155928;-23.213703;36;-1;U;Atlantic/Cape_Verde
1105;Preguica;Sao Nocolau Island;Cape Verde;SNE;GVSN;16.588356;-24.284656;669;-1;U;Atlantic/Cape_Verde
1106;Sao Pedro;Sao Vicente Island;Cape Verde;VXE;GVSV;16.833689;-25.054661;66;-1;U;Atlantic/Cape_Verde
1107;Bole Intl;Addis Ababa;Ethiopia;ADD;HAAB;8.977889;38.799319;7656;3;U;Africa/Addis_Ababa
1108;Lideta;Addis Ababa;Ethiopia;;HAAL;9.003681;38.726019;7749;3;U;Africa/Addis_Ababa
1109;Arba Minch;Arba Minch;Ethiopia;AMH;HAAM;6.039389;37.590453;3895;3;U;Africa/Addis_Ababa
1110;Axum;Axum;Ethiopia;AXU;HAAX;14.14675;38.772833;6915;3;U;Africa/Addis_Ababa
1111;Bahir Dar;Bahar Dar;Ethiopia;BJR;HABD;11.608075;37.321644;5976;3;U;Africa/Addis_Ababa
1112;Dire Dawa Intl;Dire Dawa;Ethiopia;DIR;HADR;9.6247;41.854203;3829;3;U;Africa/Addis_Ababa
1113;Gambella;Gambella;Ethiopia;GMB;HAGM;8.128764;34.563131;1771;3;U;Africa/Addis_Ababa
1114;Gondar;Gondar;Ethiopia;GDQ;HAGN;12.5199;37.434047;6541;3;U;Africa/Addis_Ababa
6811;South Ari Atoll;Paradies Island;Maldives;;ARIA;3.4833;72.8369;0;5;U;Indian/Maldives
1116;Jimma;Jimma;Ethiopia;JIM;HAJM;7.666094;36.816639;5587;3;U;Africa/Addis_Ababa
1117;Lalibella;Lalibella;Ethiopia;LLI;HALL;11.975014;38.979969;6423;3;U;Africa/Addis_Ababa
1118;Makale;Makale;Ethiopia;MQX;HAMK;13.467367;39.533464;7406;3;U;Africa/Addis_Ababa
1119;Asosa;Asosa;Ethiopia;ASO;HASO;10.01855;34.586253;5120;3;U;Africa/Addis_Ababa
1120;Bujumbura Intl;Bujumbura;Burundi;BJM;HBBA;-3.324019;29.318519;2582;2;U;Africa/Bujumbura
1121;Egal Intl;Hargeisa;Somalia;HGA;HCMH;9.518167;44.088758;4423;3;U;Africa/Mogadishu
1122;Berbera;Berbera;Somalia;BBO;HCMI;10.389167;44.941106;30;3;U;Africa/Mogadishu
1123;Kisimayu;Kismayu;Somalia;KMU;HCMK;-0.377353;42.459233;49;3;U;Africa/Mogadishu
6890;Dowagiac Municipal Airport;Dowagiac;United States;C91;\N;41.9929342;-86.1280125;748;-5;U;America/New_York
1126;Alexandria Intl;Alexandria;Egypt;ALY;HEAX;31.183903;29.948889;-6;2;U;Africa/Cairo
1127;Abu Simbel;Abu Simbel;Egypt;ABS;HEBL;22.375953;31.611722;616;2;U;Africa/Cairo
1128;Cairo Intl;Cairo;Egypt;CAI;HECA;30.121944;31.405556;382;2;U;Africa/Cairo
1129;Cairo West;Cairo;Egypt;;HECW;30.116362;30.915445;550;2;U;Africa/Cairo
1130;Hurghada Intl;Hurghada;Egypt;HRG;HEGN;27.178317;33.799436;52;2;U;Africa/Cairo
1131;El Gora;El Gorah;Egypt;EGR;HEGR;31.068975;34.129194;324;2;U;Africa/Cairo
1132;Luxor Intl;Luxor;Egypt;LXR;HELX;25.671028;32.706583;294;2;U;Africa/Cairo
1133;Mersa Matruh;Mersa-matruh;Egypt;MUH;HEMM;31.325356;27.221689;94;2;U;Africa/Cairo
1134;Port Said;Port Said;Egypt;PSD;HEPS;31.279444;32.24;8;2;U;Africa/Cairo
1135;St Catherine Intl;St. Catherine;Egypt;SKV;HESC;28.685278;34.0625;4368;2;U;Africa/Cairo
1136;Aswan Intl;Aswan;Egypt;ASW;HESN;23.964356;32.819975;662;2;U;Africa/Cairo
1137;El Tor;El-tor;Egypt;ELT;HETR;28.209028;33.645517;115;2;U;Africa/Cairo
1138;Eldoret Intl;Eldoret;Kenya;EDL;HKEL;0.404458;35.238928;6941;3;U;Africa/Nairobi
1139;Kakamega;Kakamega;Kenya;;HKKG;0.271342;34.787297;5020;3;U;Africa/Nairobi
1140;Kisumu;Kisumu;Kenya;KIS;HKKI;-0.086139;34.728892;3796;3;U;Africa/Nairobi
1141;Kitale;Kitale;Kenya;KTL;HKKT;0.971989;34.958556;6070;3;U;Africa/Nairobi
6889;Cambridge Municipal Airport;Cambridge;United States;CDI;\N;39.9750278;-81.5775833;799;-5;U;America/New_York
1143;Lodwar;Lodwar;Kenya;LOK;HKLO;3.121967;35.608692;1715;3;U;Africa/Nairobi
1144;Lamu Manda;Lamu;Kenya;LAU;HKLU;-2.252417;40.913097;20;3;U;Africa/Nairobi
1145;Mombasa Moi Intl;Mombasa;Kenya;MBA;HKMO;-4.034833;39.59425;200;3;U;Africa/Nairobi
1146;Naivasha;Naivasha;Kenya;;HKNV;-0.787953;36.433528;6380;3;U;Africa/Nairobi
1147;Nairobi Wilson;Nairobi;Kenya;WIL;HKNW;-1.321719;36.814833;5546;3;U;Africa/Nairobi
1148;Eastleigh;Nairobi;Kenya;;HKRE;-1.277267;36.862339;5380;3;U;Africa/Nairobi
1149;Wajir;Wajir;Kenya;WJR;HKWJ;1.733239;40.091606;770;3;U;Africa/Nairobi
1150;Bu Attifel;Buattifel;Libya;;HLFL;28.795372;22.080939;161;2;N;Africa/Tripoli
1151;Warehouse 59e;Giallo;Libya;;HLGL;28.638458;21.438042;325;2;N;Africa/Tripoli
1152;Ghat;Ghat;Libya;GHT;HLGT;25.145564;10.142647;2296;2;N;Africa/Tripoli
1153;Kufra;Kufra;Libya;AKF;HLKF;24.178728;23.313958;1367;2;N;Africa/Tripoli
1154;Benina;Benghazi;Libya;BEN;HLLB;32.096786;20.269472;433;2;N;Africa/Tripoli
1156;Sebha;Sebha;Libya;SEB;HLLS;26.986964;14.472525;1427;2;N;Africa/Tripoli
1157;Tripoli Intl;Tripoli;Libya;TIP;HLLT;32.663544;13.159011;263;2;N;Africa/Tripoli
1158;Marsa Brega;Marsa Brega;Libya;;HLMB;30.378139;19.576444;50;2;N;Africa/Tripoli
1159;Ras Lanuf Oil;Ras Lanouf V 40;Libya;;HLNF;30.500013;18.527161;42;2;N;Africa/Tripoli
1160;Hon;Hon;Libya;;HLON;29.110106;15.965567;919;2;N;Africa/Tripoli
1161;Dahra;Dahra;Libya;;HLRA;29.472567;17.934936;1050;2;N;Africa/Tripoli
1162;Ghadames East;Ghadames;Libya;LTD;HLTD;30.151695;9.715305;1122;2;N;Africa/Tripoli
1163;Zella 74;Zella 74;Libya;;HLZA;28.589878;17.293858;1085;2;N;Africa/Tripoli
1164;Gisenyi;Gisenyi;Rwanda;GYI;HRYG;-1.677203;29.258875;5082;2;U;Africa/Kigali
1165;Kigali Intl;Kigali;Rwanda;KGL;HRYR;-1.968628;30.13945;4859;2;U;Africa/Kigali
1166;Kamembe;Kamembe;Rwanda;KME;HRZA;-2.462242;28.90795;5192;2;U;Africa/Kigali
1167;Dongola;Dongola;Sudan;DOG;HSDN;19.153867;30.430094;773;3;U;Africa/Khartoum
1168;Damazin;Damazin;Sudan;;HSDZ;11.785889;34.336656;1582;3;U;Africa/Khartoum
1169;El Fashir;El Fasher;Sudan;ELF;HSFS;13.614892;25.32465;2393;3;U;Africa/Khartoum
1170;Kassala;Kassala;Sudan;KSL;HSKA;15.387494;36.328842;1671;3;U;Africa/Khartoum
1171;Kadugli;Kadugli;Sudan;;HSLI;11.138028;29.701122;1848;3;U;Africa/Khartoum
1172;El Obeid;El Obeid;Sudan;EBD;HSOB;13.153219;30.232675;1927;3;U;Africa/Khartoum
1173;Juba;Juba;South Sudan;JUB;HSSJ;4.872006;31.601117;1513;3;U;Africa/Juba
1174;Malakal;Malakal;Sudan;MAK;HSSM;9.558969;31.652242;1291;3;U;Africa/Juba
1175;Khartoum;Khartoum;Sudan;KRT;HSSS;15.589497;32.553161;1265;3;U;Africa/Khartoum
1176;Arusha;Arusha;Tanzania;ARK;HTAR;-3.367794;36.633333;4550;3;U;Africa/Dar_es_Salaam
1177;Mwalimu Julius K Nyerere Intl;Dar Es Salaam;Tanzania;DAR;HTDA;-6.878111;39.202625;182;3;U;Africa/Dar_es_Salaam
1178;Dodoma;Dodoma;Tanzania;DOD;HTDO;-6.170436;35.752578;3637;3;U;Africa/Dar_es_Salaam
1179;Iringa;Iringa;Tanzania;IRI;HTIR;-7.668633;35.752114;4678;3;U;Africa/Dar_es_Salaam
1180;Kilimanjaro Intl;Kilimanjaro;Tanzania;JRO;HTKJ;-3.429406;37.074461;2932;3;U;Africa/Dar_es_Salaam
1181;Lake Manyara;Lake Manyara;Tanzania;LKY;HTLM;-3.376306;35.818278;4150;3;N;Africa/Dar_es_Salaam
1182;Mtwara;Mtwara;Tanzania;MYW;HTMT;-10.339058;40.181781;371;3;U;Africa/Dar_es_Salaam
1183;Mwanza;Mwanza;Tanzania;MWZ;HTMW;-2.444486;32.932667;3763;3;U;Africa/Dar_es_Salaam
1184;Pemba;Pemba;Tanzania;PMA;HTPE;-5.257264;39.811417;80;3;U;Africa/Dar_es_Salaam
1185;Tanga;Tanga;Tanzania;TGT;HTTG;-5.092358;39.071158;129;3;U;Africa/Dar_es_Salaam
1186;Zanzibar;Zanzibar;Tanzania;ZNZ;HTZA;-6.222025;39.224886;54;3;U;Africa/Dar_es_Salaam
1187;Entebbe Intl;Entebbe;Uganda;EBB;HUEN;0.042386;32.443503;3782;3;U;Africa/Kampala
6888;Dusseldorf Hauptbahnhof;Dusseldorf;Germany;QDU;\N;51.220278;6.792778;125;1;E;Europe/Berlin
1189;Soroti;Soroti;Uganda;SRT;HUSO;1.727578;33.622861;3641;3;U;Africa/Kampala
1190;Tirana Rinas;Tirana;Albania;TIA;LATI;41.414742;19.720561;126;1;E;Europe/Tirane
1191;Burgas;Bourgas;Bulgaria;BOJ;LBBG;42.569583;27.515236;135;2;E;Europe/Sofia
1192;Gorna Oryahovitsa;Gorna Orechovica;Bulgaria;GOZ;LBGO;43.151444;25.712889;285;2;E;Europe/Sofia
1193;Plovdiv;Plovdiv;Bulgaria;PDV;LBPD;42.067806;24.850833;597;2;E;Europe/Sofia
1194;Sofia;Sofia;Bulgaria;SOF;LBSF;42.695194;23.406167;1742;2;E;Europe/Sofia
1195;Stara Zagora;Stara Zagora;Bulgaria;;LBSZ;42.376667;25.655195;558;2;E;Europe/Sofia
1196;Varna;Varna;Bulgaria;VAR;LBWN;43.232072;27.825106;230;2;E;Europe/Sofia
1197;Larnaca;Larnaca;Cyprus;LCA;LCLK;34.875117;33.62485;8;2;E;Asia/Nicosia
1198;Pafos Intl;Paphos;Cyprus;PFO;LCPH;34.718039;32.485731;41;2;E;Asia/Nicosia
1199;Akrotiri;Akrotiri;Cyprus;AKT;LCRA;34.590416;32.987861;76;0;E;Europe/London
1200;Dubrovnik;Dubrovnik;Croatia;DBV;LDDU;42.561353;18.268244;527;1;E;Europe/Zagreb
1201;Cepin;Cepin;Croatia;;LDOC;45.542169;18.636233;299;1;E;Europe/Zagreb
1202;Osijek;Osijek;Croatia;OSI;LDOS;45.462667;18.810156;290;1;E;Europe/Zagreb
1203;Pula;Pula;Croatia;PUY;LDPL;44.893533;13.922192;274;1;E;Europe/Zagreb
1204;Grobnicko Polje;Grobnik;Croatia;;LDRG;45.379483;14.503756;951;1;E;Europe/Zagreb
1205;Rijeka;Rijeka;Croatia;RJK;LDRI;45.216889;14.570267;278;1;E;Europe/Zagreb
1206;Split;Split;Croatia;SPU;LDSP;43.538944;16.297964;79;1;E;Europe/Zagreb
1207;Varazdin;Varazdin;Croatia;;LDVA;46.294724;16.38125;548;1;E;Europe/Zagreb
1208;Zagreb;Zagreb;Croatia;ZAG;LDZA;45.742931;16.068778;353;1;E;Europe/Zagreb
1209;Zadar;Zadar;Croatia;ZAD;LDZD;44.108269;15.346697;289;1;E;Europe/Zagreb
1210;Udbina;Udbina;Croatia;;LDZU;44.55761;15.774361;2462;1;E;Europe/Zagreb
1211;Albacete;Albacete;Spain;;LEAB;38.948528;-1.863517;2302;1;E;Europe/Madrid
1212;Alicante;Alicante;Spain;ALC;LEAL;38.282169;-0.558156;142;1;E;Europe/Madrid
1213;Almeria;Almeria;Spain;LEI;LEAM;36.843936;-2.370097;70;1;E;Europe/Madrid
1214;Asturias;Aviles;Spain;OVD;LEAS;43.563567;-6.034622;416;1;E;Europe/Madrid
1215;Cordoba;Cordoba;Spain;ODB;LEBA;37.842006;-4.848878;297;1;E;Europe/Madrid
1216;Bilbao;Bilbao;Spain;BIO;LEBB;43.301097;-2.910608;138;1;E;Europe/Madrid
1218;Barcelona;Barcelona;Spain;BCN;LEBL;41.297078;2.078464;12;1;E;Europe/Madrid
1219;Talavera La Real;Badajoz;Spain;BJZ;LEBZ;38.89125;-6.821333;609;1;E;Europe/Madrid
1220;A Coruna;La Coruna;Spain;LCG;LECO;43.302061;-8.377256;326;1;E;Europe/Madrid
1221;Armilla;Granada;Spain;;LEGA;37.133222;-3.635694;2297;1;E;Europe/Madrid
1222;Girona;Gerona;Spain;GRO;LEGE;41.900969;2.760547;468;1;E;Europe/Madrid
1223;Granada;Granada;Spain;GRX;LEGR;37.188731;-3.777356;1860;1;E;Europe/Madrid
1224;Getafe;Madrid;Spain;;LEGT;40.294139;-3.723833;2029;1;E;Europe/Madrid
1225;Ibiza;Ibiza;Spain;IBZ;LEIB;38.872858;1.373117;20;1;E;Europe/Madrid
1226;Jerez;Jerez;Spain;XRY;LEJR;36.744622;-6.060111;93;1;E;Europe/Madrid
1227;Murcia San Javier;Murcia;Spain;MJV;LELC;37.774972;-0.812389;11;1;E;Europe/Madrid
6887;Alexion;Porto Heli;Greece;PKH;LGHL;37.298792;23.148986;2224;2;E;Europe/Athens
1229;Barajas;Madrid;Spain;MAD;LEMD;40.493556;-3.566764;2000;1;E;Europe/Madrid
1230;Malaga;Malaga;Spain;AGP;LEMG;36.6749;-4.499106;52;1;E;Europe/Madrid
1231;Menorca;Menorca;Spain;MAH;LEMH;39.862597;4.218647;298;1;E;Europe/Madrid
1232;Moron Ab;Sevilla;Spain;OZP;LEMO;37.174917;-5.615944;285;1;E;Europe/Madrid
1233;Ocana;Ocana;Spain;;LEOC;39.9375;-3.503333;2405;1;E;Europe/Madrid
1234;Pamplona;Pamplona;Spain;PNA;LEPP;42.770039;-1.646331;1504;1;E;Europe/Madrid
1235;Alcantarilla;Murcia;Spain;;LERI;37.951111;-1.230319;250;1;E;Europe/Madrid
1236;Reus;Reus;Spain;REU;LERS;41.147392;1.167172;234;1;E;Europe/Madrid
1237;Rota Ns;Rota;Spain;;LERT;36.645211;-6.349458;86;1;E;Europe/Madrid
1238;Salamanca;Salamanca;Spain;SLM;LESA;40.952117;-5.501986;2595;1;E;Europe/Madrid
1239;Son Bonet;Son Bonet;Spain;;LESB;39.598889;2.702778;135;1;E;Europe/Madrid
1240;Palma De Mallorca;Palma De Mallorca;Spain;;LESJ;39.551675;2.738808;27;1;E;Europe/Madrid
1241;San Luis;San Luis;Spain;;LESL;39.862222;4.258333;197;1;E;Europe/Madrid
1242;San Sebastian;San Sebastian;Spain;EAS;LESO;43.356519;-1.790611;15;1;E;Europe/Madrid
1243;Santiago;Santiago;Spain;SCQ;LEST;42.896333;-8.415144;1213;1;E;Europe/Madrid
1244;Seo De Urgel;Seo De Urgel;Spain;LEU;LESU;42.338611;1.409167;2625;1;E;Europe/Madrid
1245;Torrejon;Madrid;Spain;TOJ;LETO;40.496747;-3.445872;2026;1;E;Europe/Madrid
1246;Valencia;Valencia;Spain;VLC;LEVC;39.489314;-0.481625;225;1;E;Europe/Madrid
1247;Valladolid;Valladolid;Spain;VLL;LEVD;41.706111;-4.851944;2775;1;E;Europe/Madrid
1248;Cuatro Vientos;Madrid;Spain;;LEVS;40.370678;-3.785144;2267;1;E;Europe/Madrid
1249;Vitoria;Vitoria;Spain;VIT;LEVT;42.882836;-2.724469;1682;1;E;Europe/Madrid
1250;Vigo;Vigo;Spain;VGO;LEVX;42.2318;-8.626775;855;1;E;Europe/Madrid
1251;Santander;Santander;Spain;SDR;LEXJ;43.427064;-3.820006;16;1;E;Europe/Madrid
1252;Zaragoza Ab;Zaragoza;Spain;ZAZ;LEZG;41.666242;-1.041553;863;1;E;Europe/Madrid
1253;Sevilla;Sevilla;Spain;SVQ;LEZL;37.418;-5.893106;111;1;E;Europe/Madrid
1254;Calais Dunkerque;Calais;France;CQF;LFAC;50.962097;1.954764;12;1;E;Europe/Paris
1255;Peronne St Quentin;Peronne;France;;LFAG;49.868547;3.029578;295;1;E;Europe/Paris
1256;Les Loges;Nangis;France;;LFAI;48.596219;3.006786;428;1;E;Europe/Paris
1257;Couterne;Bagnole-de-l'orne;France;;LFAO;48.545836;-0.387444;718;1;E;Europe/Paris
1258;Bray;Albert;France;;LFAQ;49.971531;2.697661;364;1;E;Europe/Paris
1259;Le Touquet Paris Plage;Le Tourquet;France;LTQ;LFAT;50.517385;1.620587;36;1;E;Europe/Paris
1260;Denain;Valenciennes;France;;LFAV;50.325808;3.461264;177;1;E;Europe/Paris
1261;Glisy;Amiens;France;;LFAY;49.873019;2.387075;208;1;E;Europe/Paris
1262;La Garenne;Agen;France;AGF;LFBA;44.174721;0.590556;204;1;E;Europe/Paris
1263;Cazaux;Cazaux;France;;LFBC;44.533333;-1.125;84;1;E;Europe/Paris
1264;Merignac;Bordeaux;France;BOD;LFBD;44.828335;-0.715556;162;1;E;Europe/Paris
1265;Roumaniere;Bergerac;France;EGC;LFBE;44.825279;0.518611;171;1;E;Europe/Paris
1266;Francazal;Toulouse;France;;LFBF;43.545555;1.3675;535;1;E;Europe/Paris
1267;Chateaubernard;Cognac;France;CNG;LFBG;45.658333;-0.3175;102;1;E;Europe/Paris
1268;Biard;Poitiers;France;PIS;LFBI;46.587745;0.306666;423;1;E;Europe/Paris
1269;Montlucon Gueret;Montlucon-gueret;France;;LFBK;46.222644;2.363964;1497;1;E;Europe/Paris
1270;Bellegarde;Limoges;France;LIG;LFBL;45.862778;1.179444;1300;1;E;Europe/Paris
1271;Mont De Marsan;Mont-de-marsan;France;;LFBM;43.911667;-0.5075;203;1;E;Europe/Paris
1272;Souche;Niort;France;NIT;LFBN;46.311303;-0.401503;203;1;E;Europe/Paris
1273;Blagnac;Toulouse;France;TLS;LFBO;43.629075;1.363819;499;1;E;Europe/Paris
1274;Pau Pyrenees;Pau;France;PUF;LFBP;43.38;-0.418611;616;1;E;Europe/Paris
1275;Lherm;La Rochelle;France;;LFBR;43.448891;1.263333;622;1;E;Europe/Paris
1276;Lourdes;Tarbes;France;LDE;LFBT;43.178675;-0.006439;1260;1;E;Europe/Paris
1277;Brie Champniers;Angouleme;France;ANG;LFBU;45.729247;0.221456;436;1;E;Europe/Paris
1278;La Roche;Brive;France;BVE;LFSL;45.150833;1.469167;379;1;E;Europe/Paris
1279;Bassillac;Perigueux;France;PGX;LFBX;45.198055;0.815556;328;1;E;Europe/Paris
1280;Anglet;Biarritz-bayonne;France;BIQ;LFBZ;43.468419;-1.523325;245;1;E;Europe/Paris
1281;Lalbenque;Cahors;France;;LFCC;44.351387;1.475278;912;1;E;Europe/Paris
1282;Antichan;St.-girons;France;;LFCG;43.007764;1.10315;1368;1;E;Europe/Paris
1283;La Teste De Buch;Arcachon;France;XAC;LFCH;44.59639;-1.110833;49;1;E;Europe/Paris
1284;Le Sequestre;Albi;France;LBI;LFCI;43.913887;2.113056;564;1;E;Europe/Paris
1285;Mazamet;Castres;France;DCM;LFCK;43.55625;2.289183;788;1;E;Europe/Paris
1286;Lasbordes;Toulouse;France;;LFCL;43.586113;1.499167;459;1;E;Europe/Paris
1287;Larzac;Millau;France;;LFCM;43.989342;3.183;2606;1;E;Europe/Paris
1288;Montdragon;Graulhet;France;;LFCQ;43.771111;2.010833;581;1;E;Europe/Paris
1289;Marcillac;Rodez;France;RDZ;LFCR;44.407869;2.482672;1910;1;E;Europe/Paris
1290;Thalamy;Ussel;France;;LFCU;45.534721;2.423889;2428;1;E;Europe/Paris
1291;Villeneuve Sur Lot;Villeneuve-sur-lot;France;;LFCW;44.396946;0.758889;190;1;E;Europe/Paris
1292;Medis;Royan;France;RYN;LFCY;45.628056;-0.9725;72;1;E;Europe/Paris
1293;Mimizan;Mimizan;France;;LFCZ;44.146111;-1.174444;164;1;E;Europe/Paris
1294;Aire Sur L Adour;Aire-sur-l'adour;France;;LFDA;43.709444;-0.245278;259;1;E;Europe/Paris
1295;Montauban;Montauban;France;;LFDB;44.025694;1.378042;351;1;E;Europe/Paris
1296;Lamothe;Auch;France;;LFDH;43.687778;0.601667;411;1;E;Europe/Paris
1297;Artigues De Lussac;Libourne;France;;LFDI;44.982498;-0.134722;157;1;E;Europe/Paris
1298;Les Pujols;Pamiers;France;;LFDJ;43.090556;1.695833;1115;1;E;Europe/Paris
1299;Virazeil;Marmande;France;;LFDM;44.498919;0.200514;105;1;E;Europe/Paris
1300;St Agnant;Rochefort;France;RCO;LFDN;45.887779;-0.983056;60;1;E;Europe/Paris
1301;Pontivy;Pontivy;France;;LFED;48.056789;-2.921244;407;1;E;Europe/Paris
1302;Aubigny Sur Nere;Aubigny-sur-nere;France;;LFEH;47.474167;2.393055;630;1;E;Europe/Paris
1303;Scaer;Guiscriff-scaer;France;;LFES;48.052508;-3.664717;574;1;E;Europe/Paris
1305;Ancenis;Ancenis;France;;LFFI;47.408056;-1.1775;111;1;E;Europe/Paris
1306;Brienne Le Chateau;Brienne-le Chateau;France;;LFFN;48.429764;4.482222;381;1;E;Europe/Paris
1307;Houssen;Colmar;France;CMR;LFGA;48.109853;7.359011;628;1;E;Europe/Paris
1308;Challanges;Beaune;France;;LFGF;47.005886;4.893425;656;1;E;Europe/Paris
1309;Tavaux;Dole;France;DLE;LFGJ;47.039014;5.42725;645;1;E;Europe/Paris
1310;Joigny;Joigny;France;;LFGK;47.992222;3.392222;732;1;E;Europe/Paris
1311;Le Rozelier;Verdun;France;;LFGW;49.122383;5.469047;1230;1;E;Europe/Paris
1312;Ardeche Meridionale;Aubenas-vals-lanas;France;OBS;LFHO;44.544236;4.372192;923;1;E;Europe/Paris
1313;Loudes;Le Puy;France;LPY;LFHP;45.080689;3.762889;2731;1;E;Europe/Paris
1314;Coltines;St.-flour;France;;LFHQ;45.076389;2.993611;3218;1;E;Europe/Paris
1315;Ceyzeriat;Bourg;France;XBK;LFHS;46.20089;5.292028;857;1;E;Europe/Paris
1316;Tarare;Vilefrance;France;XVF;LFHV;45.901947;4.634906;1076;1;E;Europe/Paris
1317;Montbeugny;Moulins;France;XMU;LFHY;46.534581;3.423725;915;1;E;Europe/Paris
1318;Belmont;St.-afrique-belmont;France;;LFIF;43.823334;2.745278;1686;1;E;Europe/Paris
1319;Cassagnes Begonhes;Cassagnes-beghones;France;;LFIG;44.177776;2.515;2024;1;E;Europe/Paris
1320;Metz Nancy Lorraine;Metz;France;ETZ;LFJL;48.982142;6.251319;870;1;E;Europe/Paris
1321;Poretta;Bastia;France;BIA;LFKB;42.552664;9.483731;26;1;E;Europe/Paris
1322;Saint Catherine;Calvi;France;CLY;LFKC;42.530753;8.793189;209;1;E;Europe/Paris
1323;Sud Corse;Figari;France;FSC;LFKF;41.500557;9.097777;87;1;E;Europe/Paris
1324;Campo Dell Oro;Ajaccio;France;AJA;LFKJ;41.923637;8.802917;18;1;E;Europe/Paris
1325;Propriano;Propriano;France;;LFKO;41.660558;8.889747;13;1;E;Europe/Paris
1326;Solenzara;Solenzara;France;SOZ;LFKS;41.924416;9.406;28;1;E;Europe/Paris
1327;Corte;Corte;France;;LFKT;42.29361;9.193055;1132;1;E;Europe/Paris
1328;Branches;Auxerre;France;AUF;LFLA;47.850193;3.497111;523;1;E;Europe/Paris
1329;Aix Les Bains;Chambery;France;CMF;LFLB;45.63805;5.880225;779;1;E;Europe/Paris
1330;Auvergne;Clermont-Ferrand;France;CFE;LFLC;45.786661;3.169169;1090;1;E;Europe/Paris
1331;Bourges;Bourges;France;BOU;LFLD;47.058056;2.370278;529;1;E;Europe/Paris
1332;Challes Les Eaux;Chambery;France;;LFLE;45.56105;5.975761;973;1;E;Europe/Paris
1333;Champforgeuil;Chalon;France;XCD;LFLH;46.826108;4.817633;623;1;E;Europe/Paris
1334;Annemasse;Annemasse;France;QNJ;LFLI;46.191972;6.268386;1620;1;E;Europe/Paris
1335;Saint Exupery;Lyon;France;LYS;LFLL;45.726387;5.090833;821;1;E;Europe/Paris
1336;Charnay;Macon;France;QNX;LFLM;46.295103;4.795767;728;1;E;Europe/Paris
1337;Saint Yan;St.-yan;France;;LFLN;46.412536;4.013264;796;1;E;Europe/Paris
1338;Renaison;Roanne;France;RNE;LFLO;46.058334;4.001389;1106;1;E;Europe/Paris
1339;Meythet;Annecy;France;NCY;LFLP;45.929233;6.098764;1521;1;E;Europe/Paris
1340;Saint Geoirs;Grenoble;France;GNB;LFLS;45.362944;5.329375;1302;1;E;Europe/Paris
1341;Domerat;Montlucon;France;MCU;LFLT;46.352525;2.570486;771;1;E;Europe/Paris
1342;Chabeuil;Valence;France;VAF;LFLU;44.921594;4.9699;525;1;E;Europe/Paris
1343;Charmeil;Vichy;France;VHY;LFLV;46.169689;3.403736;817;1;E;Europe/Paris
1344;Aurillac;Aurillac;France;AUR;LFLW;44.891388;2.421944;2096;1;E;Europe/Paris
1345;Deols;Chateauroux;France;CHR;LFLX;46.862194;1.730667;529;1;E;Europe/Paris
1346;Bron;Lyon;France;LYN;LFLY;45.727172;4.944275;659;1;E;Europe/Paris
1347;Aix Les Milles;Aix-les-milles;France;QXB;LFMA;43.505554;5.367778;367;1;E;Europe/Paris
1348;Le Cannet;Le Luc;France;;LFMC;43.384672;6.387139;265;1;E;Europe/Paris
1349;Mandelieu;Cannes;France;CEQ;LFMD;43.54205;6.953478;13;1;E;Europe/Paris
1350;Boutheon;St-Etienne;France;EBU;LFMH;45.540554;4.296389;1325;1;E;Europe/Paris
1351;Le Tube;Istres;France;;LFMI;43.522736;4.923844;82;1;E;Europe/Paris
1352;Salvaza;Carcassonne;France;CCF;LFMK;43.215978;2.306317;433;1;E;Europe/Paris
1353;Provence;Marseille;France;MRS;LFML;43.435555;5.213611;74;1;E;Europe/Paris
1354;Cote D\\'Azur;Nice;France;NCE;LFMN;43.658411;7.215872;12;1;E;Europe/Paris
1355;Caritat;Orange;France;;LFMO;44.140481;4.866717;197;1;E;Europe/Paris
1356;Rivesaltes;Perpignan;France;PGF;LFMP;42.740442;2.870667;144;1;E;Europe/Paris
1357;Le Castellet;Le Castellet;France;CTT;LFMQ;43.252506;5.785189;1391;1;E;Europe/Paris
1358;Deaux;Ales;France;;LFMS;44.069656;4.142122;668;1;E;Europe/Paris
1359;Mediterranee;Montpellier;France;MPL;LFMT;43.576194;3.963014;17;1;E;Europe/Paris
1360;Vias;Beziers;France;BZR;LFMU;43.323522;3.353903;56;1;E;Europe/Paris
1361;Caumont;Avignon;France;AVN;LFMV;43.9073;4.901831;124;1;E;Europe/Paris
1362;Salon;Salon;France;;LFMY;43.606415;5.10925;195;1;E;Europe/Paris
1363;Lezignan Corbieres;Lezignan-corbieres;France;;LFMZ;43.175835;2.734167;207;1;E;Europe/Paris
1364;Brenoux;Mende;France;MEN;LFNB;44.502108;3.532819;3362;1;E;Europe/Paris
1365;Carpentras;Carpentras;France;;LFNH;44.029847;5.078058;394;1;E;Europe/Paris
1366;Avord;Avord;France;;LFOA;47.053333;2.6325;580;1;E;Europe/Paris
1367;Tille;Beauvais;France;BVA;LFOB;49.454444;2.112778;359;1;E;Europe/Paris
1368;Chateaudun;Chateaudun;France;;LFOC;48.058142;1.376625;433;1;E;Europe/Paris
1369;St Florent;Saumur;France;;LFOD;47.256839;-0.115142;269;1;E;Europe/Paris
1370;Fauville;Evreux;France;;LFOE;49.028669;1.219864;464;1;E;Europe/Paris
1371;Octeville;Le Havre;France;LEH;LFOH;49.533889;0.088056;312;1;E;Europe/Paris
1372;Abbeville;Abbeville;France;;LFOI;50.143492;1.831892;220;1;E;Europe/Paris
1373;Bricy;Orleans;France;ORE;LFOJ;47.987778;1.760556;412;1;E;Europe/Paris
1374;Vatry;Chalons;France;XCR;LFOK;48.776072;4.184492;587;1;E;Europe/Paris
1375;Vallee De Seine;Rouen;France;URO;LFOP;49.384172;1.1748;512;1;E;Europe/Paris
1376;Val De Loire;Tours;France;TUF;LFOT;47.432222;0.727606;357;1;E;Europe/Paris
1377;Le Pontreau;Cholet;France;CET;LFOU;47.082136;-0.877064;443;1;E;Europe/Paris
1378;Entrammes;Laval;France;LVA;LFOV;48.031361;-0.742986;330;1;E;Europe/Paris
1379;St Denis De L Hotel;Orleans;France;;LFOZ;47.896946;2.163333;396;1;E;Europe/Paris
1380;Le Bourget;Paris;France;LBG;LFPB;48.969444;2.441389;218;1;E;Europe/Paris
1381;Creil;Creil;France;CSF;LFPC;49.253547;2.519139;291;1;E;Europe/Paris
1382;Charles De Gaulle;Paris;France;CDG;LFPG;49.012779;2.55;392;1;E;Europe/Paris
1383;Voisins;Coulommiers;France;;LFPK;48.837653;3.016117;470;1;E;Europe/Paris
1384;Villaroche;Melun;France;;LFPM;48.604725;2.671119;302;1;E;Europe/Paris
1385;Toussus Le Noble;Toussous-le-noble;France;TNF;LFPN;48.751922;2.106189;538;1;E;Europe/Paris
1386;Orly;Paris;France;ORY;LFPO;48.725278;2.359444;291;1;E;Europe/Paris
1387;Cormeilles En Vexin;Pontoise;France;POX;LFPT;49.096647;2.040833;325;1;E;Europe/Paris
1388;Velizy;Villacoublay;France;;LFPV;48.774406;2.201536;584;1;E;Europe/Paris
1389;Prunay;Reims;France;;LFQA;49.208689;4.156581;313;1;E;Europe/Paris
1390;Barberey;Troyes;France;QYR;LFQB;48.322136;4.016703;388;1;E;Europe/Paris
1391;Croismare;Luneville;France;;LFQC;48.593275;6.543456;790;1;E;Europe/Paris
1392;Rouvres;Etain;France;;LFQE;49.226917;5.672219;770;1;E;Europe/Paris
1393;Bellevue;Autun;France;;LFQF;46.967283;4.260572;997;1;E;Europe/Paris
1394;Fourchambault;Nevers;France;NVS;LFQG;47.002625;3.113333;602;1;E;Europe/Paris
1395;Epinoy;Cambrai;France;;LFQI;50.221814;3.154236;257;1;E;Europe/Paris
1396;Elesmes;Maubeuge;France;;LFQJ;50.310467;4.033119;452;1;E;Europe/Paris
1397;La Veze;Besancon-la-veze;France;;LFQM;47.206567;6.083681;1271;1;E;Europe/Paris
1398;Bourscheid;Phalsbourg;France;;LFQP;48.76625;7.200519;1017;1;E;Europe/Paris
1399;Lesquin;Lille;France;LIL;LFQQ;50.561942;3.089444;157;1;E;Europe/Paris
1400;Calonne;Merville;France;;LFQT;50.618394;2.642242;61;1;E;Europe/Paris
1401;Charleville Mezieres;Charleville;France;;LFQV;49.783942;4.647078;492;1;E;Europe/Paris
1402;Frotey;Vesoul-frotey;France;;LFQW;47.637611;6.203922;1249;1;E;Europe/Paris
1403;Guipavas;Brest;France;BES;LFRB;48.447911;-4.418539;325;1;E;Europe/Paris
1404;Maupertus;Cherbourg;France;CER;LFRC;49.650106;-1.470281;459;1;E;Europe/Paris
1405;Pleurtuit;Dinard;France;DNR;LFRD;48.587683;-2.079958;219;1;E;Europe/Paris
1406;Escoublac;La Baule;France;;LFRE;47.289444;-2.346389;105;1;E;Europe/Paris
1407;Granville;Granville;France;GFR;LFRF;48.883057;-1.564167;45;1;E;Europe/Paris
1408;St Gatien;Deauville;France;DOL;LFRG;49.365339;0.154306;479;1;E;Europe/Paris
1409;Lann Bihoue;Lorient;France;LRT;LFRH;47.760555;-3.44;160;1;E;Europe/Paris
1410;Les Ajoncs;La Roche-sur-yon;France;EDM;LFRI;46.701944;-1.378625;299;1;E;Europe/Paris
1411;Landivisiau;Landivisiau;France;;LFRJ;48.530258;-4.151642;348;1;E;Europe/Paris
1412;Carpiquet;Caen;France;CFR;LFRK;49.173333;-0.45;256;1;E;Europe/Paris
1413;Poulmic;Lanvedoc;France;;LFRL;48.281703;-4.445017;287;1;E;Europe/Paris
1414;Arnage;Le Mans;France;LME;LFRM;47.948611;0.201667;194;1;E;Europe/Paris
1415;St Jacques;Rennes;France;RNS;LFRN;48.069508;-1.734794;124;1;E;Europe/Paris
1416;Lannion;Lannion;France;LAI;LFRO;48.754378;-3.471656;290;1;E;Europe/Paris
1417;Pluguffan;Quimper;France;UIP;LFRQ;47.974981;-4.167786;297;1;E;Europe/Paris
1418;Nantes Atlantique;Nantes;France;NTE;LFRS;47.153189;-1.610725;90;1;E;Europe/Paris
1419;Armor;St.-brieuc Armor;France;SBK;LFRT;48.537777;-2.854445;453;1;E;Europe/Paris
1420;Ploujean;Morlaix;France;MXN;LFRU;48.603222;-3.815783;272;1;E;Europe/Paris
1421;Meucon;Vannes;France;VNE;LFRV;47.723303;-2.718561;440;1;E;Europe/Paris
1422;Montoir;St.-nazaire;France;SNR;LFRZ;47.312189;-2.149181;13;1;E;Europe/Paris
1423;Bale Mulhouse;Mulhouse;France;MLH;LFSB;47.589583;7.529914;885;1;E;Europe/Paris
1424;Meyenheim;Colmar;France;;LFSC;47.921978;7.399669;693;1;E;Europe/Paris
1425;Longvic;Dijon;France;DIJ;LFSD;47.26889;5.09;726;1;E;Europe/Paris
1426;Frescaty;Metz;France;MZM;LFSF;49.071667;6.131667;629;1;E;Europe/Paris
1427;Mirecourt;Epinal;France;EPL;LFSG;48.324961;6.069983;1084;1;E;Europe/Paris
1428;Haguenau;Haguenau;France;;LFSH;48.794331;7.817613;491;1;E;Europe/Paris
1429;Robinson;St.-dizier;France;;LFSI;48.636008;4.899417;458;1;E;Europe/Paris
1430;Courcelles;Montbeliard;France;;LFSM;47.487;6.790536;1041;1;E;Europe/Paris
1431;Essey;Nancy;France;ENC;LFSN;48.692069;6.230458;751;1;E;Europe/Paris
1432;Ochey;Nancy;France;;LFSO;48.583056;5.955;1106;1;E;Europe/Paris
1433;Pontarlier;Pontarlier;France;;LFSP;46.903958;6.327367;2683;1;E;Europe/Paris
1434;Champagne;Reims;France;RHE;LFSR;49.31;4.05;312;1;E;Europe/Paris
1435;Entzheim;Strasbourg;France;SXB;LFST;48.538319;7.628233;505;1;E;Europe/Paris
1436;Saint Sauveur;Luxeuil;France;;LFSX;47.783131;6.364056;913;1;E;Europe/Paris
1437;Pierrefeu;Cuers;France;;LFTF;43.247803;6.126697;266;1;E;Europe/Paris
1438;Le Palyvestre;Hyeres;France;TLN;LFTH;43.0973;6.14603;7;1;E;Europe/Paris
1439;Garons;Nimes;France;FNI;LFTW;43.757444;4.416347;309;1;E;Europe/Paris
1440;Miquelon;Miquelon;Saint Pierre and Miquelon;MQC;LFVM;47.095472;-56.380278;10;-3;U;America/Miquelon
1441;St Pierre;St.-pierre;Saint Pierre and Miquelon;FSP;LFVP;46.762904;-56.173088;27;-3;U;America/Miquelon
1442;Amberieu;Amberieu;France;;LFXA;45.987335;5.328445;823;1;E;Europe/Paris
1443;Damblain;Damblain;France;;LFYD;48.086325;5.664058;1280;1;E;Europe/Paris
1444;Andravida;Andravida;Greece;PYR;LGAD;37.920708;21.292583;55;2;E;Europe/Athens
1445;Agrinion;Agrinion;Greece;AGQ;LGAG;38.602022;21.351208;154;2;E;Europe/Athens
1446;Dimokritos;Alexandroupolis;Greece;AXD;LGAL;40.855869;25.956264;24;2;E;Europe/Athens
1447;Alexandria;Alexandria;Greece;;LGAX;40.651128;22.488739;27;2;E;Europe/Athens
1448;Nea Anchialos;Nea Anghialos;Greece;VOL;LGBL;39.219619;22.794339;83;2;E;Europe/Athens
1449;Elefsis;Elefsis;Greece;;LGEL;38.063831;23.556011;143;2;E;Europe/Athens
1450;Chios;Chios;Greece;JKH;LGHI;38.343175;26.140572;15;2;E;Europe/Athens
1451;Ioannina;Ioannina;Greece;IOA;LGIO;39.696388;20.8225;1558;2;E;Europe/Athens
1452;Nikos Kazantzakis;Heraklion;Greece;HER;LGIR;35.339719;25.180297;115;2;E;Europe/Athens
1453;Aristotelis;Kastoria;Greece;KSO;LGKA;40.446294;21.282186;2167;2;E;Europe/Athens
1454;Kithira;Kithira;Greece;KIT;LGKC;36.274258;23.016978;1045;2;E;Europe/Athens
1455;Kefallinia;Keffallinia;Greece;EFL;LGKF;38.120069;20.500481;59;2;E;Europe/Athens
1456;Kalamata;Kalamata;Greece;KLX;LGKL;37.068319;22.025525;26;2;E;Europe/Athens
1457;Amigdhaleon;Kavala;Greece;;LGKM;40.972775;24.341417;203;2;E;Europe/Athens
1458;Kos;Kos;Greece;KGS;LGKO;36.793335;27.091667;412;2;E;Europe/Athens
1459;Karpathos;Karpathos;Greece;AOK;LGKP;35.421408;27.146008;66;2;E;Europe/Athens
1460;Ioannis Kapodistrias Intl;Kerkyra/corfu;Greece;CFU;LGKR;39.601944;19.911667;6;2;E;Europe/Athens
1461;Kasos;Kasos;Greece;KSJ;LGKS;35.421358;26.910047;35;2;E;Europe/Athens
1462;Megas Alexandros Intl;Kavala;Greece;KVA;LGKV;40.913306;24.619223;18;2;E;Europe/Athens
1463;Filippos;Kozani;Greece;KZI;LGKZ;40.28611;21.840834;2059;2;E;Europe/Athens
1464;Leros;Leros;Greece;LRS;LGLE;37.184903;26.800289;39;2;E;Europe/Athens
1465;Limnos;Limnos;Greece;LXS;LGLM;39.917072;25.236308;14;2;E;Europe/Athens
1466;Larisa;Larissa;Greece;LRA;LGLR;39.650253;22.4655;241;2;E;Europe/Athens
1467;Megara;Megara;Greece;;LGMG;37.981114;23.365422;12;2;E;Europe/Athens
1468;Mikonos;Mykonos;Greece;JMK;LGMK;37.435128;25.348103;405;2;E;Europe/Athens
1469;Mitilini;Mytilini;Greece;MJT;LGMT;39.056667;26.598333;60;2;E;Europe/Athens
1470;Aktio;Preveza;Greece;PVK;LGPZ;38.925467;20.765311;11;2;E;Europe/Athens
1471;Maritsa;Rhodos;Greece;;LGRD;36.383056;28.108889;204;2;E;Europe/Athens
1472;Rhodes Diagoras;Rhodos;Greece;RHO;LGRP;36.405419;28.086192;17;2;E;Europe/Athens
1473;Araxos;Patras;Greece;GPA;LGRX;38.151111;21.425556;46;2;E;Europe/Athens
1474;Souda;Chania;Greece;CHQ;LGSA;35.531747;24.149678;490;2;E;Europe/Athens
1475;Alexandros Papadiamantis;Skiathos;Greece;JSI;LGSK;39.1771;23.503675;54;2;E;Europe/Athens
1476;Samos;Samos;Greece;SMI;LGSM;37.689999;26.911667;19;2;E;Europe/Athens
6886;Ashford;Lympne;United Kingdom;LYM;EGMK;51.083333;1.016667;351;0;E;Europe/London
1478;Sparti;Sparti;Greece;;LGSP;36.973892;22.526292;500;2;E;Europe/Athens
1479;Santorini;Thira;Greece;JTR;LGSR;36.399169;25.479333;127;2;E;Europe/Athens
1480;Sitia;Sitia;Greece;JSH;LGST;35.216108;26.101325;376;2;E;Europe/Athens
1481;Stefanovikion;Stefanovikion;Greece;;LGSV;39.48;22.767222;146;2;E;Europe/Athens
1482;Skiros;Skiros;Greece;SKU;LGSY;38.967553;24.487228;44;2;E;Europe/Athens
1483;Tanagra;Tanagra;Greece;;LGTG;38.339847;23.564958;485;2;E;Europe/Athens
1484;Kasteli;Kasteli;Greece;;LGTL;35.192019;25.327;1180;2;E;Europe/Athens
1485;Tripolis;Tripolis;Greece;;LGTP;37.530567;22.403633;2113;2;E;Europe/Athens
1486;Makedonia;Thessaloniki;Greece;SKG;LGTS;40.519725;22.97095;22;2;E;Europe/Athens
1487;Tatoi;Dekelia;Greece;;LGTT;38.108928;23.783836;785;2;E;Europe/Athens
1488;Dionysios Solomos;Zakynthos;Greece;ZTH;LGZA;37.750853;20.88425;15;2;E;Europe/Athens
1489;Ferihegy;Budapest;Hungary;BUD;LHBP;47.436933;19.255592;495;1;E;Europe/Budapest
1490;Debrecen;Debrecen;Hungary;DEB;LHDC;47.488917;21.615333;359;1;E;Europe/Budapest
1491;Kecskemet;Kecskemet;Hungary;;LHKE;46.9175;19.749222;376;1;E;Europe/Budapest
1492;Nyiregyhaza;Nyirregyhaza;Hungary;;LHNY;47.983892;21.692317;338;1;E;Europe/Budapest
1493;Ocseny;Ocseny;Hungary;;LHOY;46.303889;18.769167;295;1;E;Europe/Budapest
6885;Door County Cherryland Airport;Sturgeon Bay;United States;SUE;\N;44.8436667;-87.4215556;725;-6;U;America/Chicago
1496;Szentkiralyszabadja;Azentkilyszabadja;Hungary;;LHSA;47.077861;17.968444;919;1;E;Europe/Budapest
1498;Szolnok;Szolnok;Hungary;;LHSN;47.122861;20.2355;322;1;E;Europe/Budapest
1499;Amendola;Amendola;Italy;;LIBA;41.541392;15.718083;183;1;E;Europe/Rome
1500;Crotone;Crotone;Italy;CRV;LIBC;38.997225;17.080169;521;1;E;Europe/Rome
1501;Bari;Bari;Italy;BRI;LIBD;41.138856;16.760594;177;1;E;Europe/Rome
1502;Gino Lisa;Foggia;Italy;FOG;LIBF;41.432917;15.535028;266;1;E;Europe/Rome
1503;Grottaglie;Grottaglie;Italy;TAR;LIBG;40.517514;17.403212;215;1;E;Europe/Rome
1504;Lecce;Lecce;Italy;LCC;LIBN;40.239228;18.133325;156;1;E;Europe/Rome
1505;Pescara;Pescara;Italy;PSR;LIBP;42.431656;14.181067;48;1;E;Europe/Rome
1506;Casale;Brindisi;Italy;BDS;LIBR;40.657633;17.947033;47;1;E;Europe/Rome
1507;Gioia Del Colle;Gioia Del Colle;Italy;;LIBV;40.767833;16.933333;1187;1;E;Europe/Rome
1508;Lamezia Terme;Lamezia;Italy;SUF;LICA;38.905394;16.242269;39;1;E;Europe/Rome
1509;Catania Fontanarossa;Catania;Italy;CTA;LICC;37.466781;15.0664;39;1;E;Europe/Rome
1510;Lampedusa;Lampedusa;Italy;LMP;LICD;35.497914;12.618083;70;1;E;Europe/Rome
1511;Pantelleria;Pantelleria;Italy;PNL;LICG;36.816519;11.968864;635;1;E;Europe/Rome
1512;Palermo;Palermo;Italy;PMO;LICJ;38.175958;13.091019;65;1;E;Europe/Rome
1513;Boccadifalco;Palermo;Italy;;LICP;38.110833;13.313333;345;1;E;Europe/Rome
1514;Reggio Calabria;Reggio Calabria;Italy;REG;LICR;38.071206;15.651556;96;1;E;Europe/Rome
1515;Trapani Birgi;Trapani;Italy;TPS;LICT;37.911403;12.487961;24;1;E;Europe/Rome
1516;Sigonella;Sigonella;Italy;NSY;LICZ;37.401664;14.922358;79;1;E;Europe/Rome
1517;Alghero;Alghero;Italy;AHO;LIEA;40.632133;8.290772;87;1;E;Europe/Rome
1518;Decimomannu;Decimomannu;Italy;DCI;LIED;39.354222;8.972481;100;1;E;Europe/Rome
1519;Elmas;Cagliari;Italy;CAG;LIEE;39.251469;9.054283;13;1;E;Europe/Rome
1520;Olbia Costa Smeralda;Olbia;Italy;OLB;LIEO;40.898661;9.517628;37;1;E;Europe/Rome
1521;Tortoli;Tortoli;Italy;TTB;LIET;39.918761;9.682981;23;1;E;Europe/Rome
1522;Aeritalia;Turin;Italy;;LIMA;45.086353;7.603375;943;1;E;Europe/Rome
1523;Bresso;Milano;Italy;;LIMB;45.542222;9.203333;484;1;E;Europe/Rome
1524;Malpensa;Milano;Italy;MXP;LIMC;45.630606;8.728111;767;1;E;Europe/Rome
1525;Bergamo Orio Al Serio;Bergamo;Italy;BGY;LIME;45.673889;9.704166;782;1;E;Europe/Rome
1526;Torino;Torino;Italy;TRN;LIMF;45.200761;7.649631;989;1;E;Europe/Rome
1527;Albenga;Albenga;Italy;ALL;LIMG;44.050608;8.127428;148;1;E;Europe/Rome
1528;Genova Sestri;Genoa;Italy;GOA;LIMJ;44.413333;8.8375;13;1;E;Europe/Rome
1529;Linate;Milan;Italy;LIN;LIML;45.445103;9.276739;353;1;E;Europe/Rome
1530;Cameri;Cameri;Italy;;LIMN;45.529592;8.669225;586;1;E;Europe/Rome
1531;Parma;Parma;Italy;PMF;LIMP;44.824483;10.296367;161;1;E;Europe/Rome
1532;Piacenza;Piacenza;Italy;QPZ;LIMS;44.913055;9.723333;456;1;E;Europe/Rome
6884;Shoestring Aviation Airfield;Stewartstown;United States;0P2;\N;39.7948244;-76.6471914;1000;-5;U;America/New_York
1534;Levaldigi;Cuneo;Italy;CUF;LIMZ;44.547019;7.623217;1267;1;E;Europe/Rome
1535;Aviano Ab;Aviano;Italy;AVB;LIPA;46.031889;12.596472;410;1;E;Europe/Rome
1536;Bolzano;Bolzano;Italy;BZO;LIPB;46.460194;11.326383;789;1;E;Europe/Rome
1537;Cervia;Cervia;Italy;;LIPC;44.224175;12.307203;18;1;E;Europe/Rome
1538;Bologna;Bologna;Italy;BLQ;LIPE;44.535444;11.288667;123;1;E;Europe/Rome
1539;Treviso;Treviso;Italy;TSF;LIPH;45.6484;12.194422;59;1;E;Europe/Rome
1540;Rivolto;Rivolto;Italy;;LIPI;45.97875;13.049331;179;1;E;Europe/Rome
1541;Forli;Forli;Italy;FRL;LIPK;44.194753;12.070094;97;1;E;Europe/Rome
1542;Ghedi;Ghedi;Italy;;LIPL;45.432167;10.267667;333;1;E;Europe/Rome
1543;Verona Boscomantico;Verona;Italy;;LIPN;45.472017;10.927919;286;1;E;Europe/Rome
1544;Montichiari;Brescia;Italy;VBS;LIPO;45.428889;10.330556;356;1;E;Europe/Rome
1545;Ronchi Dei Legionari;Ronchi De Legionari;Italy;TRS;LIPQ;45.8275;13.472222;37;1;E;Europe/Rome
1546;Rimini;Rimini;Italy;RMI;LIPR;44.020292;12.611747;41;1;E;Europe/Rome
1547;Istrana;Treviso;Italy;;LIPS;45.684867;12.082881;137;1;E;Europe/Rome
1548;Vicenza;Vicenza;Italy;VIC;LIPT;45.573411;11.52955;128;1;E;Europe/Rome
1549;Padova;Padova;Italy;QPA;LIPU;45.395767;11.847903;44;1;E;Europe/Rome
1550;Villafranca;Villafranca;Italy;VRN;LIPX;45.395706;10.888533;239;1;E;Europe/Rome
1551;Venezia Tessera;Venice;Italy;VCE;LIPZ;45.505278;12.351944;7;1;E;Europe/Rome
1552;Ampugnano;Siena;Italy;SAY;LIQS;43.256286;11.255036;634;1;E;Europe/Rome
1553;Ciampino;Rome;Italy;CIA;LIRA;41.799361;12.594936;427;1;E;Europe/Rome
1554;Pratica Di Mare;Pratica Di Mare;Italy;;LIRE;41.654522;12.445183;41;1;E;Europe/Rome
1555;Fiumicino;Rome;Italy;FCO;LIRF;41.804475;12.250797;15;1;E;Europe/Rome
1556;Guidonia;Guidonia;Italy;;LIRG;41.990278;12.740833;289;1;E;Europe/Rome
1558;Marina Di Campo;Marina Di Campo;Italy;EBA;LIRJ;42.760277;10.239445;31;1;E;Europe/Rome
1559;Latina;Latina;Italy;QLT;LIRL;41.542436;12.909019;93;1;E;Europe/Rome
1560;Grazzanise;Grazzanise;Italy;;LIRM;41.060833;14.081944;29;1;E;Europe/Rome
1561;Capodichino;Naples;Italy;NAP;LIRN;40.886033;14.290781;294;1;E;Europe/Rome
1562;Pisa;Pisa;Italy;PSA;LIRP;43.683917;10.39275;6;1;E;Europe/Rome
1563;Firenze;Florence;Italy;FLR;LIRQ;43.809953;11.2051;142;1;E;Europe/Rome
1564;Grosseto;Grosseto;Italy;GRS;LIRS;42.759747;11.071897;15;1;E;Europe/Rome
1565;Urbe;Rome;Italy;;LIRU;41.951946;12.498889;55;1;E;Europe/Rome
1566;Viterbo;Viterbo;Italy;;LIRV;42.430183;12.064156;990;1;E;Europe/Rome
1567;Perugia;Perugia;Italy;PEG;LIRZ;43.095906;12.513222;693;1;E;Europe/Rome
1568;Cerklje;Cerklje;Slovenia;;LJCE;45.899971;15.530194;510;1;E;Europe/Ljubljana
1569;Ljubljana;Ljubljana;Slovenia;LJU;LJLJ;46.223686;14.457611;1273;1;E;Europe/Ljubljana
1570;Maribor;Maribor;Slovenia;MBX;LJMB;46.479861;15.686131;876;1;E;Europe/Ljubljana
1571;Portoroz;Portoroz;Slovenia;POW;LJPZ;45.473353;13.614978;7;1;E;Europe/Ljubljana
1572;Slovenj Gradec;Slovenj Gradec;Slovenia;;LJSG;46.471975;15.116997;1645;1;E;Europe/Ljubljana
1573;Ceske Budejovice;Ceske Budejovice;Czech Republic;;LKCS;48.946381;14.427464;1417;1;E;Europe/Prague
1574;Caslav;Caslav;Czech Republic;;LKCV;49.939653;15.381808;794;1;E;Europe/Prague
1575;Hradec Kralove;Hradec Kralove;Czech Republic;;LKHK;50.2532;15.845228;791;1;E;Europe/Prague
1576;Horovice;Horovice;Czech Republic;;LKHV;49.848111;13.893506;1214;1;E;Europe/Prague
1577;Kbely;Praha;Czech Republic;;LKKB;50.121367;14.543642;939;1;E;Europe/Prague
1578;Kunovice;Kunovice;Czech Republic;;LKKU;49.029444;17.439722;581;1;E;Europe/Prague
1579;Karlovy Vary;Karlovy Vary;Czech Republic;KLV;LKKV;50.202978;12.914983;1989;1;E;Europe/Prague
1580;Plzen Line;Line;Czech Republic;;LKLN;49.675172;13.274617;1188;1;E;Europe/Prague
1581;Mnichovo Hradiste;Mnichovo Hradiste;Czech Republic;;LKMH;50.540211;15.006592;800;1;E;Europe/Prague
1582;Mosnov;Ostrava;Czech Republic;OSR;LKMT;49.696292;18.111053;844;1;E;Europe/Prague
1583;Namest;Namest;Czech Republic;;LKNA;49.165833;16.124925;1548;1;E;Europe/Prague
1584;Pardubice;Pardubice;Czech Republic;PED;LKPD;50.013419;15.738647;741;1;E;Europe/Prague
1585;Pribram;Pribram;Czech Republic;;LKPM;49.720078;14.100567;1529;1;E;Europe/Prague
1586;Prerov;Prerov;Czech Republic;PRV;LKPO;49.425833;17.404722;676;1;E;Europe/Prague
1587;Ruzyne;Prague;Czech Republic;PRG;LKPR;50.100833;14.26;1247;1;E;Europe/Prague
1588;Turany;Brno;Czech Republic;BRQ;LKTB;49.151269;16.694433;778;1;E;Europe/Prague
1589;Vodochody;Vodochody;Czech Republic;;LKVO;50.216581;14.395806;919;1;E;Europe/Prague
1590;Ben Gurion;Tel-aviv;Israel;TLV;LLBG;32.011389;34.886667;135;2;U;Asia/Jerusalem
1591;Teyman;Beer-sheba;Israel;BEV;LLBS;31.287003;34.722953;656;2;U;Asia/Jerusalem
1592;Tel Nov;Tel-nof;Israel;;LLEK;31.839472;34.821844;193;2;U;Asia/Jerusalem
1593;Eyn Shemer;Eyn-shemer;Israel;;LLES;32.440814;35.007661;95;2;U;Asia/Jerusalem
1594;Eilat;Elat;Israel;ETH;LLET;29.561281;34.960081;42;2;U;Asia/Jerusalem
1595;En Yahav;Eyn-yahav;Israel;;LLEY;30.621656;35.203325;-164;2;U;Asia/Jerusalem
1596;Haifa;Haifa;Israel;HFA;LLHA;32.809444;35.043056;28;2;U;Asia/Jerusalem
1597;Hatzor;Haztor;Israel;;LLHS;31.7625;34.727222;148;2;U;Asia/Jerusalem
1598;Mahanaim I Ben Yaakov;Rosh Pina;Israel;RPN;LLIB;32.981047;35.571908;922;2;U;Asia/Jerusalem
1599;Megiddo;Megido Airstrip;Israel;;LLMG;32.597139;35.228803;200;2;U;Asia/Jerusalem
1600;I Bar Yehuda;Metzada;Israel;;LLMZ;31.328169;35.388608;-1266;2;U;Asia/Jerusalem
1601;Nevatim Ab;Nevatim;Israel;;LLNV;31.208347;35.0123;1330;2;U;Asia/Jerusalem
1602;Ovda;Ovda;Israel;VDA;LLOV;29.94025;34.93585;1492;2;U;Asia/Jerusalem
1603;Ramat David;Ramat David;Israel;;LLRD;32.665142;35.179458;185;2;U;Asia/Jerusalem
1604;Ramon;Ramon;Israel;;LLRM;30.776061;34.666769;2126;2;U;Asia/Jerusalem
1605;Sde Dov;Tel-aviv;Israel;SDV;LLSD;32.114661;34.782239;43;2;U;Asia/Jerusalem
1606;Luqa;Malta;Malta;MLA;LMML;35.857497;14.4775;300;1;E;Europe/Malta
1607;Wiener Neustadt East;Wiener Neustadt Ost;Austria;;LOAN;47.843333;16.260139;896;1;E;Europe/Vienna
1608;Wels;Wels;Austria;;LOLW;48.183304;14.040861;1043;1;E;Europe/Vienna
1609;Graz;Graz;Austria;GRZ;LOWG;46.991067;15.439628;1115;1;E;Europe/Vienna
1610;Innsbruck;Innsbruck;Austria;INN;LOWI;47.260219;11.343964;1906;1;E;Europe/Vienna
1611;Linz;Linz;Austria;LNZ;LOWL;48.233219;14.187511;978;1;E;Europe/Vienna
1612;Salzburg;Salzburg;Austria;SZG;LOWS;47.793304;13.004333;1411;1;E;Europe/Vienna
1613;Schwechat;Vienna;Austria;VIE;LOWW;48.110278;16.569722;600;1;E;Europe/Vienna
1614;Klagenfurt;Klagenfurt;Austria;;LOXK;46.642514;14.337739;1470;1;E;Europe/Vienna
1615;Zeltweg;Zeltweg;Austria;;LOXZ;47.202751;14.744223;2220;1;E;Europe/Vienna
1616;Alverca;Alverca;Portugal;;LPAR;38.883278;-9.030097;11;0;E;Europe/Lisbon
1617;Santa Maria;Santa Maria (island);Portugal;SMA;LPAZ;36.97139;-25.170639;308;-1;E;Atlantic/Azores
1618;Braganca;Braganca;Portugal;BGC;LPBG;41.8578;-6.707125;2241;0;E;Europe/Lisbon
1619;Beja;Beja (madeira);Portugal;;LPBJ;38.078903;-7.932397;636;0;E;Europe/Lisbon
1620;Braga;Braga;Portugal;;LPBR;41.587058;-8.445139;247;0;E;Europe/Lisbon
1621;Coimbra;Coimba;Portugal;;LPCO;40.157223;-8.47;587;0;E;Europe/Lisbon
1622;Cascais;Cascais;Portugal;;LPCS;38.725025;-9.355231;326;0;E;Europe/Lisbon
1623;Covilha;Covilha;Portugal;;LPCV;40.264772;-7.479958;1572;0;E;Europe/Lisbon
1624;Evora;Evora;Portugal;;LPEV;38.533472;-7.889639;807;0;E;Europe/Lisbon
1625;Flores;Flores;Portugal;FLW;LPFL;39.455272;-31.131361;112;-1;E;Atlantic/Azores
1626;Faro;Faro;Portugal;FAO;LPFR;37.014425;-7.965911;24;0;E;Europe/Lisbon
1627;Graciosa;Graciosa Island;Portugal;GRW;LPGR;39.092169;-28.029847;86;-1;E;Atlantic/Azores
1628;Horta;Horta;Portugal;HOR;LPHR;38.519894;-28.715872;118;-1;E;Atlantic/Azores
1629;Lajes;Lajes (terceira Island);Portugal;TER;LPLA;38.761842;-27.090797;180;-1;E;Atlantic/Azores
1630;Monte Real;Monte Real;Portugal;;LPMR;39.831244;-8.887261;187;0;E;Europe/Lisbon
1631;Montijo;Montijo;Portugal;;LPMT;38.703861;-9.035916;46;0;E;Europe/Lisbon
1632;Ovar;Ovar;Portugal;;LPOV;40.915867;-8.645919;56;0;E;Europe/Lisbon
1633;Ponta Delgada;Ponta Delgada;Portugal;PDL;LPPD;37.741184;-25.69787;259;-1;E;Atlantic/Azores
1634;Pico;Pico;Portugal;PIX;LPPI;38.554333;-28.441333;109;-1;E;Atlantic/Azores
1635;Portimao;Portimao;Portugal;;LPPM;37.149331;-8.583961;5;0;E;Europe/Lisbon
1636;Porto;Porto;Portugal;OPO;LPPR;41.248055;-8.681389;228;0;E;Europe/Lisbon
1637;Porto Santo;Porto Santo;Portugal;PXO;LPPS;33.073386;-16.349975;341;0;E;Europe/Lisbon
1638;Lisboa;Lisbon;Portugal;LIS;LPPT;38.781311;-9.135919;374;0;E;Europe/Lisbon
1639;Sao Jorge;Sao Jorge Island;Portugal;SJZ;LPSJ;38.6655;-28.175817;311;-1;E;Atlantic/Azores
1640;Sintra;Sintra;Portugal;;LPST;38.831053;-9.339553;440;0;E;Europe/Lisbon
1641;Tancos;Tancos;Portugal;;LPTN;39.47514;-8.364583;266;0;E;Europe/Lisbon
1642;Vila Real;Vila Real;Portugal;VRL;LPVR;41.274334;-7.720472;1805;0;E;Europe/Lisbon
1643;Viseu;Viseu;Portugal;;LPVZ;40.725539;-7.888992;2060;0;E;Europe/Lisbon
6883;Eastern Oregon Regional Airport;Pendleton;United States;PDT;KPDT;45.695;-118.841389;1497;-8;A;America/Los_Angeles
1645;Mostar;Mostar;Bosnia and Herzegovina;OMO;LQMO;43.2829;17.845878;156;1;E;Europe/Sarajevo
1646;Sarajevo;Sarajevo;Bosnia and Herzegovina;SJJ;LQSA;43.824583;18.331467;1708;1;E;Europe/Sarajevo
1647;Arad;Arad;Romania;ARW;LRAR;46.17655;21.262022;352;2;E;Europe/Bucharest
1648;Bacau;Bacau;Romania;BCM;LRBC;46.521946;26.910278;607;2;E;Europe/Bucharest
1649;Tautii Magheraus;Baia Mare;Romania;BAY;LRBM;47.658389;23.470022;605;2;E;Europe/Bucharest
1650;Aurel Vlaicu;Bucharest;Romania;BBU;LRBS;44.503194;26.102111;297;2;E;Europe/Bucharest
1651;Mihail Kogalniceanu;Constanta;Romania;CND;LRCK;44.362222;28.488333;353;2;E;Europe/Bucharest
1652;Cluj Napoca;Cluj-napoca;Romania;CLJ;LRCL;46.785167;23.686167;1036;2;E;Europe/Bucharest
1653;Caransebes;Caransebes;Romania;CSB;LRCS;45.42;22.253333;866;2;E;Europe/Bucharest
1654;Craiova;Craiova;Romania;CRA;LRCV;44.318139;23.888611;626;2;E;Europe/Bucharest
1655;Iasi;Iasi;Romania;IAS;LRIA;47.178492;27.620631;397;2;E;Europe/Bucharest
1656;Oradea;Oradea;Romania;OMR;LROD;47.025278;21.9025;465;2;E;Europe/Bucharest
1657;Henri Coanda;Bucharest;Romania;OTP;LROP;44.572161;26.102178;314;2;E;Europe/Bucharest
1658;Sibiu;Sibiu;Romania;SBZ;LRSB;45.785597;24.091342;1496;2;E;Europe/Bucharest
1659;Satu Mare;Satu Mare;Romania;SUJ;LRSM;47.703275;22.8857;405;2;E;Europe/Bucharest
1660;Stefan Cel Mare;Suceava;Romania;SCV;LRSV;47.6875;26.354056;1375;2;E;Europe/Bucharest
1661;Cataloi;Tulcea;Romania;TCE;LRTC;45.062486;28.714311;200;2;E;Europe/Bucharest
1662;Transilvania Targu Mures;Tirgu Mures;Romania;TGM;LRTM;46.467714;24.412525;963;2;E;Europe/Bucharest
1663;Traian Vuia;Timisoara;Romania;TSR;LRTR;45.809861;21.337861;348;2;E;Europe/Bucharest
1664;Les Eplatures;Les Eplatures;Switzerland;;LSGC;47.08385;6.792836;3368;1;E;Europe/Zurich
1665;Geneve Cointrin;Geneva;Switzerland;GVA;LSGG;46.238064;6.10895;1411;1;E;Europe/Paris
1666;Saanen;Saanen;Switzerland;;LSGK;46.487499;7.250834;3307;1;E;Europe/Zurich
1667;Sion;Sion;Switzerland;SIR;LSGS;46.219592;7.326764;1585;1;E;Europe/Zurich
1668;Alpnach;Alpnach;Switzerland;;LSMA;46.943889;8.284167;1460;1;E;Europe/Zurich
1669;Dubendorf;Dubendorf;Switzerland;;LSMD;47.398644;8.648231;1470;1;E;Europe/Zurich
1670;Emmen;Emmen;Switzerland;;LSME;47.092369;8.305117;1400;1;E;Europe/Zurich
1671;Mollis;Mollis;Switzerland;;LSMF;47.078872;9.064831;1485;1;E;Europe/Zurich
1672;Meiringen;Meiringen;Switzerland;;LSMM;46.743333;8.11;1895;1;E;Europe/Zurich
1673;Payerne;Payerne;Switzerland;;LSMP;46.843208;6.915058;1465;1;E;Europe/Zurich
1674;Buochs;Buochs;Switzerland;;LSMU;46.974914;8.39915;1475;1;E;Europe/Zurich
1675;Lugano;Lugano;Switzerland;LUG;LSZA;46.004275;8.910578;915;1;E;Europe/Zurich
1676;Bern Belp;Bern;Switzerland;BRN;LSZB;46.9141;7.497153;1674;1;E;Europe/Zurich
1677;Grenchen;Grenchen;Switzerland;;LSZG;47.181628;7.417189;1411;1;E;Europe/Zurich
1678;Zurich;Zurich;Switzerland;ZRH;LSZH;47.464722;8.549167;1416;1;E;Europe/Zurich
1679;St Gallen Altenrhein;Altenrhein;Switzerland;ACH;LSZR;47.485033;9.560775;1306;1;E;Europe/Zurich
1680;Samedan;Samedan;Switzerland;SMV;LSZS;46.534075;9.884106;5600;1;E;Europe/Zurich
1681;Guvercinlik;Ankara;Turkey;;LTAB;39.93495;32.740775;2694;2;E;Europe/Istanbul
1682;Esenboga;Ankara;Turkey;ESB;LTAC;40.128082;32.995083;3125;2;E;Europe/Istanbul
1683;Etimesgut;Ankara;Turkey;ANK;LTAD;39.949831;32.688622;2653;2;E;Europe/Istanbul
1684;Akinci;Ankara;Turkey;;LTAE;40.078919;32.565625;2767;2;E;Europe/Istanbul
1685;Adana;Adana;Turkey;ADA;LTAF;36.982166;35.280388;65;2;E;Europe/Istanbul
1686;Incirlik Ab;Adana;Turkey;;LTAG;37.0021;35.425894;238;2;E;Europe/Istanbul
1687;Afyon;Afyon;Turkey;AFY;LTAH;38.726425;30.601114;3310;2;E;Europe/Istanbul
1688;Antalya;Antalya;Turkey;AYT;LTAI;36.898731;30.800461;177;2;E;Europe/Istanbul
1689;Oguzeli;Gaziantep;Turkey;GZT;LTAJ;36.947183;37.478683;2315;2;E;Europe/Istanbul
1690;Iskenderun;Iskenderun;Turkey;;LTAK;36.573472;36.154;25;2;E;Europe/Istanbul
1691;Konya;Konya;Turkey;KYA;LTAN;37.979;32.561861;3381;2;E;Europe/Istanbul
1692;Tulga;Malatya;Turkey;;LTAO;38.35375;38.253944;3016;2;E;Europe/Istanbul
1693;Merzifon;Merzifon;Turkey;MZH;LTAP;40.829375;35.521992;1758;2;E;Europe/Istanbul
1694;Sivas;Sivas;Turkey;VAS;LTAR;39.813828;36.903497;5236;2;E;Europe/Istanbul
1695;Erhac;Malatya;Turkey;MLX;LTAT;38.435347;38.091006;2828;2;E;Europe/Istanbul
1696;Erkilet;Kayseri;Turkey;ASR;LTAU;38.770386;35.495428;3463;2;E;Europe/Istanbul
1697;Sivrihisar;Sivrihisar;Turkey;;LTAV;39.451469;31.365308;3185;2;E;Europe/Istanbul
1698;Tokat;Tokat;Turkey;;LTAW;40.306281;36.371089;1831;2;E;Europe/Istanbul
1699;Cardak;Denizli;Turkey;DNZ;LTAY;37.785567;29.701297;2795;2;E;Europe/Istanbul
1701;Ataturk;Istanbul;Turkey;IST;LTBA;40.976922;28.814606;163;2;E;Europe/Istanbul
1702;Balikesir;Balikesir;Turkey;BZI;LTBF;39.619258;27.925958;340;2;E;Europe/Istanbul
1703;Bandirma;Bandirma;Turkey;BDM;LTBG;40.317972;27.977694;170;2;E;Europe/Istanbul
6882;Tyonek Airport;Tyonek;United States;TYE;\N;61.076667;-151.138056;110;-9;A;America/Anchorage
1705;Eskisehir;Eskisehir;Turkey;ESK;LTBI;39.784138;30.582111;2581;2;E;Europe/Istanbul
1706;Adnan Menderes;Izmir;Turkey;ADB;LTBJ;38.292392;27.156953;412;2;E;Europe/Istanbul
1707;Gaziemir;Izmir;Turkey;;LTBK;38.319106;27.159353;433;2;E;Europe/Istanbul
1708;Cigli;Izmir;Turkey;IGL;LTBL;38.513022;27.010053;16;2;E;Europe/Istanbul
1709;Isparta;Isparta;Turkey;;LTBM;37.785369;30.581817;3250;2;E;Europe/Istanbul
1710;Kutahya;Kutahya;Turkey;;LTBN;39.426708;30.016872;3026;2;E;Europe/Istanbul
1712;Yalova;Yalova;Turkey;;LTBP;40.684353;29.375728;6;2;E;Europe/Istanbul
1713;Topel;Topel;Turkey;;LTBQ;40.735028;30.083336;182;2;E;Europe/Istanbul
1715;Dalaman;Dalaman;Turkey;DLM;LTBS;36.713056;28.7925;20;2;E;Europe/Istanbul
1716;Akhisar;Akhisar;Turkey;;LTBT;38.808887;27.83386;263;2;E;Europe/Istanbul
6881;Riverton Regional;Riverton WY;United States;RIW;KRIW;43.064167;-108.459722;5525;-7;A;America/Denver
1718;Imsik;Bodrum;Turkey;BXN;LTBV;37.140144;27.669717;202;2;E;Europe/Istanbul
1719;Samandira;Istanbul;Turkey;;LTBX;40.992975;29.216539;400;2;E;Europe/Istanbul
1721;Elazig;Elazig;Turkey;EZS;LTCA;38.606925;39.291417;2927;2;E;Europe/Istanbul
1722;Diyarbakir;Diyabakir;Turkey;DIY;LTCC;37.893897;40.201019;2251;2;E;Europe/Istanbul
1723;Erzincan;Erzincan;Turkey;ERC;LTCD;39.710203;39.527003;3783;2;E;Europe/Istanbul
1724;Erzurum;Erzurum;Turkey;ERZ;LTCE;39.956501;41.170166;5763;2;E;Europe/Istanbul
1726;Trabzon;Trabzon;Turkey;TZX;LTCG;40.995108;39.789728;104;2;E;Europe/Istanbul
6880;Montrose Regional Airport;Montrose CO;United States;MTJ;KMTJ;38.509794;-107.894242;5759;-7;A;America/Denver
1728;Van;Van;Turkey;VAN;LTCI;38.468219;43.3323;5480;2;E;Europe/Istanbul
1729;Batman;Batman;Turkey;BAL;LTCJ;37.928969;41.116583;1822;2;E;Europe/Istanbul
1731;Siirt;Siirt;Turkey;;LTCL;37.978886;41.840436;2001;2;E;Europe/Istanbul
1732;Kaklic;Izmir;Turkey;;LTFA;38.517608;26.977406;13;2;E;Europe/Istanbul
1733;Efes;Izmir;Turkey;;LTFB;37.950747;27.329014;10;2;E;Europe/Istanbul
1734;Balti Intl;Saltsy;Moldova;;LUBL;47.838114;27.781475;758;2;E;Europe/Chisinau
1735;Chisinau Intl;Chisinau;Moldova;KIV;LUKK;46.927744;28.930978;399;2;E;Europe/Chisinau
1736;Ohrid;Ohrid;Macedonia;OHD;LWOH;41.179956;20.742325;2313;1;E;Europe/Skopje
1737;Skopje;Skopje;Macedonia;SKP;LWSK;41.961622;21.621381;781;1;E;Europe/Skopje
1738;Gibraltar;Gibraltar;Gibraltar;GIB;LXGB;36.151219;-5.349664;15;1;N;Europe/Gibraltar
1739;Beograd;Belgrade;Serbia;BEG;LYBE;44.818444;20.309139;335;1;E;Europe/Belgrade
1740;Nis;Nis;Serbia;INI;LYNI;43.337289;21.853722;648;1;E;Europe/Belgrade
1741;Podgorica;Podgorica;Montenegro;TGD;LYPG;42.359392;19.251894;141;1;E;Europe/Podgorica
1742;Pristina;Pristina;Serbia;PRN;LYPR;42.572778;21.035833;1789;1;E;Europe/Belgrade
1743;Tivat;Tivat;Montenegro;TIV;LYTV;42.404664;18.723286;20;1;E;Europe/Podgorica
1744;Vrsac;Vrsac;Serbia;;LYVR;45.146889;21.309889;276;1;E;Europe/Belgrade
1745;M R Stefanik;Bratislava;Slovakia;BTS;LZIB;48.170167;17.212667;436;1;E;Europe/Bratislava
1746;Kosice;Kosice;Slovakia;KSC;LZKZ;48.663055;21.241112;755;1;E;Europe/Bratislava
1747;Malacky;Malacky;Slovakia;;LZMC;48.402028;17.118389;679;1;E;Europe/Bratislava
1748;Piestany;Piestany;Slovakia;PZY;LZPP;48.625247;17.828444;545;1;E;Europe/Bratislava
1749;Sliac;Sliac;Slovakia;SLD;LZSL;48.637839;19.134108;1043;1;E;Europe/Bratislava
1750;Trencin;Trencin;Slovakia;;LZTN;48.865003;17.99225;676;1;E;Europe/Bratislava
1751;Tatry;Poprad;Slovakia;TAT;LZTT;49.073594;20.241142;2356;1;E;Europe/Bratislava
6879;Clow International Airport;Bolingbrook;United States;1CS;\N;41.6959744;-88.1292306;670;-6;U;America/Chicago
1753;North Caicos;North Caicos;Turks and Caicos Islands;NCA;MBNC;21.917475;-71.939561;10;-5;U;America/Grand_Turk
1754;Providenciales;Providenciales;Turks and Caicos Islands;PLS;MBPV;21.773625;-72.265886;15;-5;U;America/Grand_Turk
1755;South Caicos;South Caicos;Turks and Caicos Islands;XSC;MBSC;21.515739;-71.528528;6;-5;U;America/Grand_Turk
1756;Arroyo Barril Intl;Samana;Dominican Republic;EPS;MDAB;19.198586;-69.429772;57;-4;U;America/Santo_Domingo
1757;Maria Montez Intl;Barahona;Dominican Republic;BRX;MDBH;18.251464;-71.1204;10;-4;U;America/Santo_Domingo
1758;Cabo Rojo;Cabo Rojo;Dominican Republic;;MDCR;17.928981;-71.644769;262;-4;U;America/Santo_Domingo
1759;Casa De Campo Intl;La Romana;Dominican Republic;LRM;MDLR;18.450711;-68.911833;240;-4;U;America/Santo_Domingo
1760;Punta Cana Intl;Punta Cana;Dominican Republic;PUJ;MDPC;18.567367;-68.363431;47;-4;U;America/Santo_Domingo
1761;Gregorio Luperon Intl;Puerto Plata;Dominican Republic;POP;MDPP;19.7579;-70.570033;15;-4;U;America/Santo_Domingo
1762;Las Americas Intl;Santo Domingo;Dominican Republic;SDQ;MDSD;18.429664;-69.668925;59;-4;U;America/Santo_Domingo
1763;San Isidro Ab;San Isidoro;Dominican Republic;;MDSI;18.503706;-69.761706;111;-4;U;America/Santo_Domingo
1764;Cibao Intl;Santiago;Dominican Republic;STI;MDST;19.406092;-70.604689;565;-4;U;America/Santo_Domingo
1765;Bananera;Bananera;Guatemala;;MGBN;15.473528;-88.837222;130;-6;U;America/Guatemala
1766;Coban;Coban;Guatemala;CBV;MGCB;15.468958;-90.406742;4339;-6;U;America/Guatemala
1767;La Aurora;Guatemala City;Guatemala;GUA;MGGT;14.583272;-90.527475;4952;-6;U;America/Guatemala
1769;Retalhuleu;Retalhuleu;Guatemala;;MGRT;14.521017;-91.697256;656;-6;U;America/Guatemala
1770;San Jose;San Jose;Guatemala;;MGSJ;13.936192;-90.835833;29;-6;U;America/Guatemala
1771;Goloson Intl;La Ceiba;Honduras;LCE;MHLC;15.742481;-86.853036;49;-6;U;America/Tegucigalpa
1772;La Mesa Intl;San Pedro Sula;Honduras;SAP;MHLM;15.452639;-87.923556;91;-6;U;America/Tegucigalpa
1773;Guanaja;Guanaja;Honduras;GJA;MHNJ;16.445367;-85.906611;49;-6;U;America/Tegucigalpa
1774;Juan Manuel Galvez Intl;Roatan;Honduras;RTB;MHRO;16.316814;-86.522961;18;-6;U;America/Tegucigalpa
1775;Tela;Tela;Honduras;TEA;MHTE;15.775864;-87.475847;7;-6;U;America/Tegucigalpa
1776;Toncontin Intl;Tegucigalpa;Honduras;TGU;MHTG;14.060883;-87.217197;3294;-6;U;America/Tegucigalpa
1777;Trujillo;Trujillo;Honduras;;MHTJ;15.926847;-85.93825;3;-6;U;America/Tegucigalpa
1778;Boscobel;Ocho Rios;Jamaica;OCJ;MKBS;18.404247;-76.969017;90;-5;U;America/Jamaica
1779;Norman Manley Intl;Kingston;Jamaica;KIN;MKJP;17.935667;-76.7875;10;-5;U;America/Jamaica
1780;Sangster Intl;Montego Bay;Jamaica;MBJ;MKJS;18.503717;-77.913358;4;-5;U;America/Jamaica
1781;Ken Jones;Port Antonio;Jamaica;POT;MKKJ;18.198806;-76.534528;20;-5;U;America/Jamaica
1782;Tinson Pen;Kingston;Jamaica;KTP;MKTP;17.988558;-76.823761;16;-5;U;America/Jamaica
1783;General Juan N Alvarez Intl;Acapulco;Mexico;ACA;MMAA;16.757061;-99.753953;16;-6;S;America/Mexico_City
1784;Del Norte Intl;Monterrey;Mexico;NTR;MMAN;25.865572;-100.237239;1476;-6;S;America/Mexico_City
1785;Jesus Teran Intl;Aguascalientes;Mexico;AGU;MMAS;21.705558;-102.317858;6112;-6;S;America/Mexico_City
1786;Bahias De Huatulco Intl;Huatulco;Mexico;HUX;MMBT;15.775317;-96.262572;464;-6;S;America/Mexico_City
1787;General Mariano Matamoros;Cuernavaca;Mexico;CVJ;MMCB;18.834764;-99.2613;4277;-6;S;America/Mexico_City
1788;Ciudad Acuna Intl New;Ciudad Acuna;Mexico;;MMCC;29.333917;-101.100892;1410;-6;S;America/Mexico_City
1789;Ciudad Del Carmen Intl;Ciudad Del Carmen;Mexico;CME;MMCE;18.653739;-91.799017;10;-6;S;America/Mexico_City
1790;Nuevo Casas Grandes;Nuevo Casas Grandes;Mexico;;MMCG;30.397439;-107.874969;4850;-7;S;America/Mazatlan
1791;Chilpancingo;Chilpancingo;Mexico;;MMCH;17.573767;-99.514339;4199;-6;S;America/Mexico_City
1792;Culiacan Intl;Culiacan;Mexico;CUL;MMCL;24.764547;-107.474717;108;-7;S;America/Mazatlan
1793;Chetumal Intl;Chetumal;Mexico;CTM;MMCM;18.504667;-88.326847;39;-6;S;America/Mexico_City
1794;Ciudad Obregon Intl;Ciudad Obregon;Mexico;CEN;MMCN;27.392639;-109.833111;205;-7;S;America/Hermosillo
1795;Ingeniero Alberto Acuna Ongay Intl;Campeche;Mexico;CPE;MMCP;19.816794;-90.500314;34;-6;S;America/Mexico_City
1796;Abraham Gonzalez Intl;Ciudad Juarez;Mexico;CJS;MMCS;31.636133;-106.428667;3904;-7;S;America/Mazatlan
1797;General R Fierro Villalobos Intl;Chihuahua;Mexico;CUU;MMCU;28.702875;-105.964567;4462;-7;S;America/Mazatlan
1798;General Pedro Jose Mendez Intl;Ciudad Victoria;Mexico;CVM;MMCV;23.703336;-98.956486;761;-6;S;America/Mexico_City
6878;Kenosha Regional Airport;Kenosha;United States;ENW;\N;42.5956944;-87.9278056;742;-6;A;America/Chicago
1800;Cozumel Intl;Cozumel;Mexico;CZM;MMCZ;20.522403;-86.925644;15;-6;S;America/Mexico_City
1801;Durango Intl;Durango;Mexico;DGO;MMDO;24.124194;-104.528014;6104;-6;S;America/Mexico_City
1802;Tepic;Tepic;Mexico;TPQ;MMEP;21.419453;-104.842581;3020;-7;S;America/Mazatlan
1803;Ensenada;Ensenada;Mexico;ESE;MMES;31.795281;-116.602772;66;-8;S;America/Tijuana
1804;Don Miguel Hidalgo Y Costilla Intl;Guadalajara;Mexico;GDL;MMGL;20.5218;-103.311167;5016;-6;S;America/Mexico_City
1805;General Jose Maria Yanez Intl;Guaymas;Mexico;GYM;MMGM;27.968983;-110.925169;59;-7;S;America/Hermosillo
1806;Tehuacan;Tehuacan;Mexico;TCN;MMHC;18.497189;-97.419942;5509;-6;S;America/Mexico_City
1807;General Ignacio P Garcia Intl;Hermosillo;Mexico;HMO;MMHO;29.095858;-111.047858;627;-7;S;America/Hermosillo
1808;Colima;Colima;Mexico;CLQ;MMIA;19.277011;-103.577397;2467;-6;S;America/Mexico_City
1809;Isla Mujeres;Isla Mujeres;Mexico;ISJ;MMIM;21.245033;-86.739967;7;-6;S;America/Mexico_City
1810;Plan De Guadalupe Intl;Saltillo;Mexico;SLW;MMIO;25.549497;-100.928669;4778;-6;S;America/Mexico_City
1811;Ixtepec;Iztepec;Mexico;;MMIT;16.449336;-95.093697;164;-6;S;America/Mexico_City
1813;Lazaro Cardenas;Lazard Cardenas;Mexico;LZC;MMLC;18.001731;-102.220525;39;-6;S;America/Mexico_City
1814;Valle Del Fuerte Intl;Los Mochis;Mexico;LMM;MMLM;25.685194;-109.080806;16;-7;S;America/Mazatlan
1815;Guanajuato Intl;Del Bajio;Mexico;BJX;MMLO;20.993464;-101.480847;5956;-6;S;America/Mexico_City
1816;General Manuel Marquez De Leon Intl;La Paz;Mexico;LAP;MMLP;24.072694;-110.362475;69;-7;S;America/Mazatlan
1817;Loreto Intl;Loreto;Mexico;LTO;MMLT;25.989194;-111.348361;34;-7;S;America/Mazatlan
1818;General Servando Canales Intl;Matamoros;Mexico;MAM;MMMA;25.769894;-97.525311;25;-6;S;America/Mexico_City
1819;Licenciado Manuel Crescencio Rejon Int;Merida;Mexico;MID;MMMD;20.936981;-89.657672;38;-6;S;America/Mexico_City
1820;General Rodolfo Sanchez Taboada Intl;Mexicali;Mexico;MXL;MMML;32.630634;-115.241637;74;-8;S;America/Tijuana
1821;General Francisco J Mujica Intl;Morelia;Mexico;MLM;MMMM;19.849944;-101.0255;6033;-6;S;America/Mexico_City
1822;Minatitlan;Minatitlan;Mexico;MTT;MMMT;18.103419;-94.580681;36;-6;S;America/Mexico_City
1823;Monclova Intl;Monclova;Mexico;LOV;MMMV;26.955733;-101.470136;1864;-6;S;America/Mexico_City
1824;Licenciado Benito Juarez Intl;Mexico City;Mexico;MEX;MMMX;19.436303;-99.072097;7316;-6;S;America/Mexico_City
1825;General Mariano Escobedo Intl;Monterrey;Mexico;MTY;MMMY;25.778489;-100.106878;1278;-6;S;America/Mexico_City
1826;General Rafael Buelna Intl;Mazatlan;Mexico;MZT;MMMZ;23.161356;-106.266072;38;-7;S;America/Mazatlan
1827;Nogales Intl;Nogales;Mexico;NOG;MMNG;31.226083;-110.975831;3990;-7;S;America/Hermosillo
1828;Quetzalcoatl Intl;Nuevo Laredo;Mexico;NLD;MMNL;27.443918;-99.57046;484;-6;S;America/Mexico_City
1829;Xoxocotlan Intl;Oaxaca;Mexico;OAX;MMOX;16.999906;-96.726639;4989;-6;S;America/Mexico_City
1830;Tajin;Poza Rico;Mexico;PAZ;MMPA;20.602671;-97.460839;497;-6;S;America/Mexico_City
1831;Hermanos Serdan Intl;Puebla;Mexico;PBC;MMPB;19.158144;-98.371447;7361;-6;S;America/Mexico_City
1832;Ingeniero Juan Guillermo Villasana;Pachuca;Mexico;PCA;MMPC;20.0772;-98.782814;7600;-6;S;America/Mexico_City
1833;Puerto Penasco;Punta Penasco;Mexico;PPE;MMPE;31.351878;-113.525728;30;-7;S;America/Hermosillo
1834;Piedras Negras Intl;Piedras Negras;Mexico;PDS;MMPG;28.627394;-100.535211;901;-6;S;America/Mexico_City
1835;Licenciado Y Gen Ignacio Lopez Rayon;Uruapan;Mexico;UPN;MMPN;19.396692;-102.039056;5258;-6;S;America/Mexico_City
1836;Licenciado Gustavo Diaz Ordaz Intl;Puerto Vallarta;Mexico;PVR;MMPR;20.680083;-105.254167;23;-6;S;America/Mexico_City
1837;Puerto Escondido Intl;Puerto Escondido;Mexico;PXM;MMPS;15.876861;-97.089117;294;-6;S;America/Mexico_City
1838;Queretaro Intercontinental;Queretaro;Mexico;QRO;MMQT;20.617289;-100.185658;6296;-6;S;America/Mexico_City
1839;General Lucio Blanco Intl;Reynosa;Mexico;REX;MMRX;26.008908;-98.228513;139;-6;S;America/Mexico_City
1840;Los Cabos Intl;San Jose Del Cabo;Mexico;SJD;MMSD;23.15185;-109.721044;374;-7;S;America/Mazatlan
1841;San Felipe Intl;San Filipe;Mexico;;MMSF;30.930222;-114.808639;148;-8;S;America/Tijuana
1842;Ponciano Arriaga Intl;San Luis Potosi;Mexico;SLP;MMSP;22.254303;-100.930806;6035;-6;S;America/Mexico_City
1843;Tlaxcala;Tlaxcala;Mexico;TXA;MMTA;19.537964;-98.173467;8229;-6;S;America/Mexico_City
1844;General Div P A Angel H Corzo Molina;Tuxtla Gutierrez;Mexico;;MMTB;16.739919;-93.173297;1909;-6;S;America/Mexico_City
1845;Torreon Intl;Torreon;Mexico;TRC;MMTC;25.568278;-103.410583;3688;-6;S;America/Mexico_City
1846;Angel Albino Corzo;Tuxtla Gutierrez;Mexico;TGZ;MMTG;16.561822;-93.026081;1491;-6;S;America/Mexico_City
1847;General Abelardo L Rodriguez Intl;Tijuana;Mexico;TIJ;MMTJ;32.541064;-116.970158;489;-8;S;America/Tijuana
1848;General Francisco Javier Mina Intl;Tampico;Mexico;TAM;MMTM;22.29645;-97.865931;80;-6;S;America/Mexico_City
1849;Tamuin;Tamuin;Mexico;TSL;MMTN;22.038292;-98.806503;164;-6;S;America/Mexico_City
1850;Licenciado Adolfo Lopez Mateos Intl;Toluca;Mexico;TLC;MMTO;19.337072;-99.566008;8466;-6;S;America/Mexico_City
1851;Tapachula Intl;Tapachula;Mexico;TAP;MMTP;14.794339;-92.370025;97;-6;S;America/Mexico_City
1852;Cancun Intl;Cancun;Mexico;CUN;MMUN;21.036528;-86.877083;20;-6;S;America/Mexico_City
1853;C P A Carlos Rovirosa Intl;Villahermosa;Mexico;VSA;MMVA;17.997;-92.817361;46;-6;S;America/Mexico_City
1854;General Heriberto Jara Intl;Vera Cruz;Mexico;VER;MMVR;19.145931;-96.187267;90;-6;S;America/Mexico_City
1855;General Leobardo C Ruiz Intl;Zacatecas;Mexico;ZCL;MMZC;22.897112;-102.68689;7141;-6;S;America/Mexico_City
1856;Ixtapa Zihuatanejo Intl;Zihuatanejo;Mexico;ZIH;MMZH;17.601569;-101.460536;26;-6;S;America/Mexico_City
1857;Zamora;Zamora;Mexico;ZMM;MMZM;20.045036;-102.275955;5141;-6;S;America/Mexico_City
1858;Playa De Oro Intl;Manzanillo;Mexico;ZLO;MMZO;19.144778;-104.558631;30;-6;S;America/Mexico_City
1859;Zapopan;Zapopan;Mexico;;MMZP;20.755833;-103.465278;5333;-6;S;America/Mexico_City
1860;Bluefields;Bluefields;Nicaragua;BEF;MNBL;11.990961;-83.774086;41;-6;U;America/Managua
1861;Los Brasiles;Los Brasiles;Nicaragua;;MNBR;12.190044;-86.353872;262;-6;U;America/Managua
1862;Leon;Leon;Nicaragua;;MNLN;12.428028;-86.902361;328;-6;U;America/Managua
1863;Managua Intl;Managua;Nicaragua;MGA;MNMG;12.141494;-86.168178;194;-6;U;America/Managua
1864;Puerto Cabezas;Puerto Cabezas;Nicaragua;PUZ;MNPC;14.047194;-83.386722;69;-6;U;America/Managua
1865;Bocas Del Toro Intl;Bocas Del Toro;Panama;BOC;MPBO;9.340853;-82.250842;10;-5;U;America/Panama
1866;Cap Manuel Nino Intl;Changuinola;Panama;CHX;MPCH;9.458636;-82.516806;19;-5;U;America/Panama
1867;Enrique Malek Intl;David;Panama;DAV;MPDA;8.391003;-82.434992;89;-5;U;America/Panama
1868;Howard;Howard;Panama;HOW;MPHO;8.914794;-79.599633;52;-5;U;America/Panama
1869;Marcos A Gelabert Intl;Panama;Panama;PAC;MPMG;8.973339;-79.555583;31;-5;U;America/Panama
1870;Ruben Cantu;Santiago;Panama;;MPSA;8.085597;-80.945253;272;-5;U;America/Panama
1871;Tocumen Intl;Panama City;Panama;PTY;MPTO;9.071364;-79.383453;135;-5;U;America/Panama
1872;Buenos Aires;Buenos Aires;Costa Rica;;MRBA;9.163606;-83.329872;1214;-6;U;America/Costa_Rica
6877;North Las Vegas Airport;Las Vegas;United States;VGT;\N;36.2106944;-115.1944444;2205;-8;U;America/Los_Angeles
1874;Coto 47;Coto 47;Costa Rica;OTR;MRCC;8.601556;-82.968614;26;-6;U;America/Costa_Rica
1875;Chacarita;Chacarita;Costa Rica;;MRCH;9.981408;-84.772736;7;-6;U;America/Costa_Rica
6876;Brown County Airport;Georgetown;United States;;KGEO;38.8819456;-83.8827367;958;-5;U;America/New_York
1877;El Carmen De Siquirres;El Carmen;Costa Rica;;MREC;10.202033;-83.472167;56;-6;U;America/Costa_Rica
1878;Nuevo Palmar Sur;Finca 10;Costa Rica;;MRFI;8.916347;-83.507267;49;-6;U;America/Costa_Rica
1879;Golfito;Golfito;Costa Rica;GLF;MRGF;8.653775;-83.180544;49;-6;U;America/Costa_Rica
1880;Guapiles;Guapiles;Costa Rica;;MRGP;10.217194;-83.797003;883;-6;U;America/Costa_Rica
1881;Daniel Oduber Quiros Intl;Liberia;Costa Rica;LIR;MRLB;10.593289;-85.544408;270;-6;U;America/Costa_Rica
1882;Los Chiles;Los Chiles;Costa Rica;;MRLC;11.035277;-84.706108;131;-6;U;America/Costa_Rica
1883;Limon Intl;Limon;Costa Rica;LIO;MRLM;9.957961;-83.022006;7;-6;U;America/Costa_Rica
1884;Nosara;Nosara Beach;Costa Rica;NOB;MRNS;9.97649;-85.653;33;-6;U;America/Costa_Rica
1885;Juan Santamaria Intl;San Jose;Costa Rica;SJO;MROC;9.993861;-84.208806;3021;-6;U;America/Costa_Rica
1886;Pandora;Pandora;Costa Rica;;MRPD;9.732169;-82.983214;98;-6;U;America/Costa_Rica
1887;Palmar Sur;Palmar Sur;Costa Rica;PMZ;MRPM;8.951025;-83.468583;49;-6;U;America/Costa_Rica
1889;La Managua;Quepos;Costa Rica;XQP;MRQP;9.443164;-84.129772;85;-6;U;America/Costa_Rica
1890;Santa Clara De Guapiles;Santa Clara;Costa Rica;;MRSG;10.288278;-83.713519;262;-6;U;America/Costa_Rica
1891;San Vito De Java;San Vito De Jaba;Costa Rica;;MRSV;8.826111;-82.958885;3228;-6;U;America/Costa_Rica
1892;El Salvador Intl;San Salvador;El Salvador;SAL;MSLP;13.440947;-89.055728;101;-6;U;America/El_Salvador
1893;Ilopango Intl;San Salvador;El Salvador;;MSSS;13.699492;-89.119861;2021;-6;U;America/El_Salvador
1894;Cayes;Cayes;Haiti;;MTCA;18.271103;-73.788289;203;-5;U;America/Port-au-Prince
1895;Cap Haitien Intl;Cap Haitien;Haiti;CAP;MTCH;19.732989;-72.194739;10;-5;U;America/Port-au-Prince
1896;Jacmel;Jacmel;Haiti;;MTJA;18.241083;-72.5185;167;-5;U;America/Port-au-Prince
1897;Toussaint Louverture Intl;Port-au-prince;Haiti;PAP;MTPP;18.58005;-72.292542;122;-5;U;America/Port-au-Prince
1898;Gustavo Rizo;Baracoa Playa;Cuba;BCA;MUBA;20.365317;-74.506206;26;-5;U;America/Havana
1899;Carlos Manuel De Cespedes;Bayamo;Cuba;BYM;MUBY;20.396331;-76.621494;203;-5;U;America/Havana
1900;Maximo Gomez;Ciego De Avila;Cuba;AVI;MUCA;22.027053;-78.789617;335;-5;U;America/Havana
1901;Jardines Del Rey;Cunagua;Cuba;;MUCC;22.460986;-78.328422;13;-5;U;America/Havana
1902;Jaime Gonzalez;Cienfuegos;Cuba;CFG;MUCF;22.15;-80.414167;102;-5;U;America/Havana
1903;Vilo Acuna Intl;Cayo Largo del Sur;Cuba;CYO;MUCL;21.616453;-81.545989;10;-5;U;America/Havana
1904;Ignacio Agramonte Intl;Camaguey;Cuba;CMW;MUCM;21.420428;-77.847433;413;-5;U;America/Havana
1905;Antonio Maceo Intl;Santiago De Cuba;Cuba;SCU;MUCU;19.969769;-75.835414;249;-5;U;America/Havana
1906;Florida;Florida;Cuba;;MUFL;21.499722;-78.202778;197;-5;U;America/Havana
1907;Guantanamo Bay Ns;Guantanamo;Cuba;;MUGM;19.906458;-75.207056;56;-5;U;\N
1908;Mariana Grajales;Guantanamo;Cuba;GAO;MUGT;20.085419;-75.158328;56;-5;U;America/Havana
1909;Jose Marti Intl;Havana;Cuba;HAV;MUHA;22.989153;-82.409086;210;-5;U;America/Havana
1910;Frank Pais Intl;Holguin;Cuba;HOG;MUHG;20.785589;-76.315108;361;-5;U;America/Havana
1911;La Coloma;La Coloma;Cuba;LCL;MULM;22.336261;-83.642111;131;-5;U;America/Havana
1912;Orestes Acosta;Moa;Cuba;MOA;MUMO;20.654114;-74.922114;16;-5;U;America/Havana
1913;Sierra Maestra;Manzanillo;Cuba;MZO;MUMZ;20.288172;-77.0893;112;-5;U;America/Havana
1914;Rafael Cabrera;Nueva Gerona;Cuba;GER;MUNG;21.834689;-82.783819;79;-5;U;America/Havana
1915;Playa Baracoa;Baracoa Playa;Cuba;;MUPB;23.032778;-82.579444;102;-5;U;America/Havana
1916;Pinar Del Rio;Pinar Del Rio Norte;Cuba;;MUPR;22.421356;-83.678428;131;-5;U;America/Havana
1917;San Antonio De Los Banos;San Antonio De Banos;Cuba;;MUSA;22.871529;-82.509308;164;-5;U;America/Havana
1918;Abel Santamaria;Santa Clara;Cuba;SNU;MUSC;22.492192;-79.943611;338;-5;U;America/Havana
1919;Santa Lucia;Santa Lucia;Cuba;;MUSL;21.509453;-77.020375;13;-5;U;America/Havana
1920;Siguanea;Siguanea;Cuba;;MUSN;21.642525;-82.955114;39;-5;U;America/Havana
1921;Sancti Spiritus;Sancti Spiritus;Cuba;;MUSS;21.969969;-79.442692;295;-5;U;America/Havana
6875;Grand Geneva Resort Airport;Lake Geneva;United States;C02;\N;42.6149167;-88.3895833;835;-6;U;America/Chicago
1923;Juan Gualberto Gomez Intl;Varadero;Cuba;VRA;MUVR;23.034445;-81.435278;210;-5;U;America/Havana
1924;Hermanos Ameijeiras;Las Tunas;Cuba;VTU;MUVT;20.987642;-76.9358;328;-5;U;America/Havana
1925;Gerrard Smith Intl;Cayman Barac;Cayman Islands;CYB;MWCB;19.686981;-79.882789;8;-5;U;America/Cayman
1926;Owen Roberts Intl;Georgetown;Cayman Islands;GCM;MWCR;19.292778;-81.35775;8;-5;U;America/Cayman
1927;Clarence A Bain;Clarence Bain;Bahamas;;MYAB;24.287664;-77.684614;19;-5;U;America/Nassau
1928;Fresh Creek;Andros Town;Bahamas;ASD;MYAF;24.698283;-77.795611;5;-5;U;America/Nassau
1930;Marsh Harbour;Marsh Harbor;Bahamas;MHH;MYAM;26.511406;-77.083472;6;-5;U;America/Nassau
1931;San Andros;San Andros;Bahamas;SAQ;MYAN;25.053814;-78.048997;5;-5;U;America/Nassau
1932;Spring Point;Spring Point;Bahamas;AXP;MYAP;22.441828;-73.970858;11;-5;U;America/Nassau
1933;Sandy Point;Sandy Point;Bahamas;;MYAS;26.004639;-77.395483;8;-5;U;America/Nassau
1934;Treasure Cay;Treasure Cay;Bahamas;TCB;MYAT;26.745336;-77.391269;8;-5;U;America/Nassau
1935;Chub Cay;Chub Cay;Bahamas;CCZ;MYBC;25.417108;-77.88085;5;-5;U;America/Nassau
1936;Great Harbour Cay;Bullocks Harbour;Bahamas;;MYBG;25.738331;-77.840114;18;-5;U;America/Nassau
1937;South Bimini;Alice Town;Bahamas;BIM;MYBS;25.699881;-79.264656;10;-5;U;America/Nassau
6874;De Kalb Taylor Municipal Airport;De Kalb;United States;DKB;\N;41.9338342;-88.7056864;914;-6;U;America/Chicago
1941;Exuma Intl;Great Exuma;Bahamas;GGT;MYEF;23.562631;-75.877958;9;-5;U;America/Nassau
1942;George Town;George Town;Bahamas;;MYEG;23.466667;-75.78167;5;-5;U;America/Nassau
1943;North Eleuthera;North Eleuthera;Bahamas;ELH;MYEH;25.474861;-76.683489;13;-5;U;America/Nassau
1944;Governors Harbour;Governor's Harbor;Bahamas;GHB;MYEM;25.284706;-76.331011;26;-5;U;America/Nassau
1945;Normans Cay;Norman's Cay;Bahamas;;MYEN;24.594258;-76.820214;8;-5;U;America/Nassau
1946;Rock Sound;Rock Sound;Bahamas;RSD;MYER;24.8917;-76.177739;10;-5;U;America/Nassau
1947;Staniel Cay;Staniel Cay;Bahamas;;MYES;24.169083;-76.439056;5;-5;U;America/Nassau
1948;Grand Bahama Intl;Freeport;Bahamas;FPO;MYGF;26.558686;-78.695553;7;-5;U;America/Nassau
1949;Matthew Town;Matthew Town;Bahamas;IGA;MYIG;20.975;-73.666862;8;-5;U;America/Nassau
1950;Deadmans Cay;Dead Man's Cay;Bahamas;LGI;MYLD;23.179014;-75.093597;9;-5;U;America/Nassau
1951;Stella Maris;Stella Maris;Bahamas;SML;MYLS;23.581444;-75.270475;10;-5;U;America/Nassau
1952;Mayaguana;Mayaguana;Bahamas;MYG;MYMM;22.379528;-73.0135;11;-5;U;America/Nassau
1953;Lynden Pindling Intl;Nassau;Bahamas;NAS;MYNN;25.038958;-77.466231;16;-5;U;America/Nassau
1954;Duncan Town;Duncan Town;Bahamas;;MYRD;22.181803;-75.729456;6;-5;U;America/Nassau
1955;New Port Nelson;Port Nelson;Bahamas;;MYRP;23.684378;-74.836186;15;-5;U;America/Nassau
1956;San Salvador;Cockburn Town;Bahamas;ZSA;MYSM;24.063275;-74.523967;24;-5;U;America/Nassau
1957;Philip S W Goldson Intl;Belize City;Belize;BZE;MZBZ;17.539144;-88.308203;15;-6;U;America/Belize
1958;Aitutaki;Aitutaki;Cook Islands;AIT;NCAI;-18.830922;-159.764233;14;-10;U;Pacific/Rarotonga
1959;Rarotonga Intl;Avarua;Cook Islands;RAR;NCRG;-21.202739;-159.805556;19;-10;U;Pacific/Rarotonga
1960;Nadi Intl;Nandi;Fiji;NAN;NFFN;-17.755392;177.443378;59;12;U;Pacific/Fiji
1961;Nausori Intl;Nausori;Fiji;SUV;NFNA;-18.043267;178.559228;17;12;U;Pacific/Fiji
1963;Fua Amotu Intl;Tongatapu;Tonga;TBU;NFTF;-21.241214;-175.149644;126;13;U;Pacific/Tongatapu
1964;Vavau Intl;Vava'u;Tonga;VAV;NFTV;-18.585336;-173.961717;236;13;U;Pacific/Tongatapu
1965;Bonriki Intl;Tarawa;Kiribati;TRW;NGTA;1.381636;173.147036;9;12;U;Pacific/Tarawa
1966;Tabiteuea North;Tabiteuea North;Kiribati;TBF;NGTE;-1.224469;174.775614;7;12;U;\N
6809;Maitland Airport;Maitland;Australia;MTL;YMND;-32.7033;151.488;75;10;O;Australia/Sydney
1968;Wallis;Wallis;Wallis and Futuna;WLS;NLWW;-13.238281;-176.199228;79;12;U;Pacific/Wallis
1969;Faleolo Intl;Faleolo;Samoa;APW;NSFA;-13.829969;-172.008336;58;13;U;Pacific/Apia
1970;Pago Pago Intl;Pago Pago;American Samoa;PPG;NSTU;-14.331;-170.7105;32;-11;U;Pacific/Pago_Pago
1971;Rurutu;Rurutu;French Polynesia;RUR;NTAR;-22.434069;-151.360614;18;-10;U;\N
1972;Tubuai;Tubuai;French Polynesia;TUB;NTAT;-23.365353;-149.524072;7;-10;U;Pacific/Tahiti
1973;Anaa;Anaa;French Polynesia;AAA;NTGA;-17.352606;-145.509956;10;-10;U;Pacific/Tahiti
1974;Fangatau;Fangatau;French Polynesia;;NTGB;-15.819928;-140.88715;9;-10;U;Pacific/Tahiti
1975;Tikehau;Tikehau;French Polynesia;TIH;NTGC;-15.119617;-148.230697;6;-10;U;Pacific/Tahiti
1976;Reao;Reao;French Polynesia;REA;NTGE;-18.465861;-136.439706;12;-10;U;\N
1977;Fakarava;Fakarava;French Polynesia;FAV;NTGF;-16.05415;-145.656994;13;-10;U;Pacific/Tahiti
1978;Manihi;Manihi;French Polynesia;XMH;NTGI;-14.436764;-146.070056;14;-10;U;Pacific/Tahiti
1979;Totegegie;Totegegie;French Polynesia;GMR;NTGJ;-23.079861;-134.890333;7;-9;U;Pacific/Gambier
1980;Kaukura;Kaukura Atoll;French Polynesia;KKR;NTGK;-15.663333;-146.884769;11;-10;U;Pacific/Tahiti
1981;Makemo;Makemo;French Polynesia;MKP;NTGM;-16.583919;-143.658369;3;-10;U;Pacific/Tahiti
1982;Puka Puka;Puka Puka;French Polynesia;PKP;NTGP;-14.809458;-138.812811;5;-10;U;Pacific/Tahiti
1983;Takapoto;Takapoto;French Polynesia;TKP;NTGT;-14.709544;-145.245814;12;-10;U;Pacific/Tahiti
1984;Arutua;Arutua;French Polynesia;AXR;NTGU;-15.248289;-146.616708;9;-10;U;Pacific/Tahiti
1985;Mataiva;Mataiva;French Polynesia;MVT;NTGV;-14.868055;-148.717225;11;-10;U;Pacific/Tahiti
1986;Takaroa;Takaroa;French Polynesia;TKX;NTKR;-14.455781;-145.024542;13;-10;U;Pacific/Tahiti
1987;Nuku Hiva;Nuku Hiva;French Polynesia;NHV;NTMD;-8.795603;-140.228789;220;-9.5;U;Pacific/Marquesas
6873;Purude University Airport;Lafayette;United States;LAF;\N;40.4123056;-86.9368889;606;-5;A;America/New_York
1989;Bora Bora;Bora Bora;French Polynesia;BOB;NTTB;-16.444378;-151.751286;10;-10;U;Pacific/Tahiti
1990;Rangiroa;Rangiroa;French Polynesia;RGI;NTTG;-14.954283;-147.6608;10;-10;U;Pacific/Tahiti
1991;Huahine;Huahine Island;French Polynesia;HUH;NTTH;-16.687242;-151.021667;7;-10;U;Pacific/Tahiti
1992;Moorea;Moorea;French Polynesia;MOZ;NTTM;-17.489972;-149.761869;9;-10;U;Pacific/Tahiti
1993;Hao;Hao Island;French Polynesia;HOI;NTTO;-18.074814;-140.945886;10;-10;U;\N
1994;Maupiti;Maupiti;French Polynesia;MAU;NTTP;-16.426486;-152.243669;15;-10;U;Pacific/Tahiti
1995;Raiatea;Raiatea Island;French Polynesia;RFP;NTTR;-16.722861;-151.465856;3;-10;U;Pacific/Tahiti
1997;Port Vila Bauerfield;Port-vila;Vanuatu;VLI;NVVV;-17.699325;168.319794;70;11;U;Pacific/Efate
1998;Kone;Kone;New Caledonia;KNQ;NWWD;-21.053428;164.837806;23;11;U;Pacific/Noumea
1999;Koumac;Koumac;New Caledonia;KOC;NWWK;-20.546314;164.255625;42;11;U;Pacific/Noumea
2000;Lifou;Lifou;New Caledonia;LIF;NWWL;-20.7748;167.239864;92;11;U;Pacific/Noumea
2001;Magenta;Noumea;New Caledonia;GEA;NWWM;-22.258278;166.472806;10;11;U;Pacific/Noumea
2002;Mare;Mare;New Caledonia;MEE;NWWR;-21.481678;168.037508;141;11;U;Pacific/Noumea
2003;Touho;Touho;New Caledonia;TOU;NWWU;-20.790028;165.259486;10;11;U;Pacific/Noumea
2004;Ouvea;Ouvea;New Caledonia;UVE;NWWV;-20.640556;166.572778;23;11;U;Pacific/Noumea
2005;La Tontouta;Noumea;New Caledonia;NOU;NWWW;-22.014553;166.212972;52;11;U;Pacific/Noumea
2006;Auckland Intl;Auckland;New Zealand;AKL;NZAA;-37.008056;174.791667;23;12;Z;Pacific/Auckland
2007;Taupo;Taupo;New Zealand;TUO;NZAP;-38.739723;176.084444;1335;12;Z;Pacific/Auckland
2008;Ardmore;Ardmore;New Zealand;AMZ;NZAR;-37.029722;174.973333;111;12;Z;Pacific/Auckland
2009;Christchurch Intl;Christchurch;New Zealand;CHC;NZCH;-43.489358;172.532225;123;12;Z;Pacific/Auckland
2010;Chatham Islands;Chatham Island;New Zealand;CHT;NZCI;-43.81;-176.457222;43;12.75;Z;Pacific/Chatham
2011;Dunedin;Dunedin;New Zealand;DUD;NZDN;-45.928055;170.198333;4;12;Z;Pacific/Auckland
2012;Gisborne;Gisborne;New Zealand;GIS;NZGS;-38.663333;177.978333;15;12;Z;Pacific/Auckland
2013;Glentanner;Glentanner;New Zealand;MON;NZGT;-43.906666;170.128333;1824;12;Z;Pacific/Auckland
2014;Hokitika;Hokitika;New Zealand;HKK;NZHK;-42.713611;170.985278;146;12;Z;Pacific/Auckland
2015;Hamilton;Hamilton;New Zealand;HLZ;NZHN;-37.866661;175.332056;172;12;Z;Pacific/Auckland
2016;Hastings;Hastings;New Zealand;;NZHS;-39.646667;176.766944;72;12;Z;Pacific/Auckland
2017;Kerikeri;Kerikeri;New Zealand;KKE;NZKK;-35.262779;173.911944;492;12;Z;Pacific/Auckland
2018;Kaitaia;Kaitaia;New Zealand;KAT;NZKT;-35.07;173.285278;270;12;Z;Pacific/Auckland
2019;Alexandra;Alexandra;New Zealand;ALR;NZLX;-45.211666;169.373333;752;12;Z;Pacific/Auckland
2020;Mount Cook;Mount Cook;New Zealand;GTN;NZMC;-43.764999;170.133333;2153;12;Z;Pacific/Auckland
2021;Manapouri;Manapouri;New Zealand;TEU;NZMO;-45.533056;167.65;687;12;Z;Pacific/Auckland
2022;Masterton;Masterton;New Zealand;MRO;NZMS;-40.973333;175.633611;364;12;Z;Pacific/Auckland
2023;New Plymouth;New Plymouth;New Zealand;NPL;NZNP;-39.008611;174.179167;97;12;Z;Pacific/Auckland
2024;Nelson;Nelson;New Zealand;NSN;NZNS;-41.298333;173.221111;17;12;Z;Pacific/Auckland
2025;Invercargill;Invercargill;New Zealand;IVC;NZNV;-46.412408;168.312992;5;12;Z;Pacific/Auckland
2026;Ohakea;Ohakea;New Zealand;;NZOH;-40.206039;175.387808;164;12;Z;Pacific/Auckland
2027;Oamaru;Oamaru;New Zealand;OAM;NZOU;-44.97;171.081667;99;12;Z;Pacific/Auckland
2028;Palmerston North;Palmerston North;New Zealand;PMR;NZPM;-40.320556;175.616944;151;12;Z;Pacific/Auckland
2029;Paraparaumu;Paraparaumu;New Zealand;PPQ;NZPP;-40.904722;174.989167;22;12;Z;Pacific/Auckland
2030;Queenstown;Queenstown International;New Zealand;ZQN;NZQN;-45.021111;168.739167;1171;12;Z;Pacific/Auckland
2031;Rotorua;Rotorua;New Zealand;ROT;NZRO;-38.109167;176.317222;935;12;Z;Pacific/Auckland
2032;Waiouru;Waiouru;New Zealand;;NZRU;-39.446389;175.658333;2686;12;Z;Pacific/Auckland
2033;South Pole Station;Stephen's Island;Antarctica;;NZSP;-89.999997;0;9300;12;U;Antarctica/South_Pole
2034;Tauranga;Tauranga;New Zealand;TRG;NZTG;-37.671944;176.19611;13;12;Z;Pacific/Auckland
2035;Timaru;Timaru;New Zealand;TIU;NZTU;-44.302778;171.225278;89;12;Z;Pacific/Auckland
2036;Pukaki;Pukaki;New Zealand;;NZUK;-44.235;170.118333;1575;12;Z;Pacific/Auckland
2037;Woodbourne;Woodbourne;New Zealand;BHE;NZWB;-41.518333;173.870278;109;12;Z;Pacific/Auckland
2038;Mcmurdo Station;Weydon;Antarctica;;NZWD;-77.867358;167.056572;68;12;N;Antarctica/South_Pole
2039;Wanaka;Wanaka;New Zealand;WKA;NZWF;-44.722222;169.245556;1142;12;Z;Pacific/Auckland
2040;Wigram;Wigram;New Zealand;;NZWG;-43.551111;172.552778;74;12;Z;Pacific/Auckland
2041;Whakatane;Whakatane;New Zealand;WHK;NZWK;-37.920556;176.914167;20;12;Z;Pacific/Auckland
2042;Wellington Intl;Wellington;New Zealand;WLG;NZWN;-41.327221;174.805278;41;12;Z;Pacific/Auckland
2043;Wairoa;Wairoa;New Zealand;;NZWO;-39.006944;177.406667;42;12;Z;Pacific/Auckland
2044;Whenuapai;Whenuapai;New Zealand;;NZWP;-36.787777;174.630278;100;12;Z;Pacific/Auckland
2045;Whangarei;Whangarei;New Zealand;WRE;NZWR;-35.768333;174.365;133;12;Z;Pacific/Auckland
2046;Westport;Westport;New Zealand;WSZ;NZWS;-41.738056;171.580833;13;12;Z;Pacific/Auckland
2047;Wanganui;Wanganui;New Zealand;WAG;NZWU;-39.962222;175.025278;27;12;Z;Pacific/Auckland
2048;Herat;Herat;Afghanistan;HEA;OAHR;34.210017;62.2283;3206;4.5;U;Asia/Kabul
2049;Jalalabad;Jalalabad;Afghanistan;JAA;OAJL;34.399842;70.498625;1814;4.5;U;Asia/Kabul
2050;Kabul Intl;Kabul;Afghanistan;KBL;OAKB;34.565853;69.212328;5877;4.5;U;Asia/Kabul
2051;Kandahar;Kandahar;Afghanistan;KDH;OAKN;31.505756;65.847822;3337;4.5;U;Asia/Kabul
2052;Maimana;Maimama;Afghanistan;MMZ;OAMN;35.930789;64.760917;2743;4.5;U;Asia/Kabul
2053;Mazar I Sharif;Mazar-i-sharif;Afghanistan;MZR;OAMS;36.706914;67.209678;1284;4.5;U;Asia/Kabul
2054;Shindand;Shindand;Afghanistan;;OASD;33.391331;62.260975;3773;4.5;U;Asia/Kabul
2055;Sheberghan;Sheberghan;Afghanistan;;OASG;36.750783;65.913164;1053;4.5;U;Asia/Kabul
2056;Konduz;Kunduz;Afghanistan;UND;OAUZ;36.665111;68.910833;1457;4.5;U;Asia/Kabul
2057;Bahrain Intl;Bahrain;Bahrain;BAH;OBBI;26.270834;50.63361;6;3;U;Asia/Bahrain
2058;Shaikh Isa;Bahrain;Bahrain;;OBBS;25.918362;50.590557;136;3;U;Asia/Bahrain
2059;Abha;Abha;Saudi Arabia;AHB;OEAB;18.240367;42.656625;6858;3;U;Asia/Riyadh
2060;Al Ahsa;Al-ahsa;Saudi Arabia;HOF;OEAH;25.285306;49.485189;588;3;U;Asia/Riyadh
2061;Al Baha;El-baha;Saudi Arabia;ABT;OEBA;20.296139;41.634277;5486;3;U;Asia/Riyadh
2062;Bisha;Bisha;Saudi Arabia;BHH;OEBH;19.98435;42.620881;3887;3;U;Asia/Riyadh
2063;Abqaiq;Abqaiq;Saudi Arabia;;OEBQ;25.911281;49.591231;229;3;U;Asia/Riyadh
2064;King Fahd Intl;Dammam;Saudi Arabia;DMM;OEDF;26.471161;49.79789;72;3;U;Asia/Riyadh
2065;King Abdulaziz Ab;Dhahran;Saudi Arabia;DHA;OEDR;26.265417;50.152027;84;3;U;Asia/Riyadh
2066;King Abdullah Bin Abdulaziz;Gizan;Saudi Arabia;GIZ;OEGN;16.901111;42.585833;20;3;U;Asia/Riyadh
2067;Gassim;Gassim;Saudi Arabia;ELQ;OEGS;26.302822;43.773911;2126;3;U;Asia/Riyadh
2068;Guriat;Guriat;Saudi Arabia;URY;OEGT;31.411942;37.279469;1672;3;U;Asia/Riyadh
2069;Hail;Hail;Saudi Arabia;HAS;OEHL;27.437917;41.686292;3331;3;U;Asia/Riyadh
2070;Jubail;Jubail;Saudi Arabia;;OEJB;27.039028;49.405083;26;3;U;Asia/Riyadh
2071;King Faisal Naval Base;Jeddah;Saudi Arabia;;OEJF;21.3481;39.173033;7;3;U;Asia/Riyadh
2072;King Abdulaziz Intl;Jeddah;Saudi Arabia;JED;OEJN;21.679564;39.156536;48;3;U;Asia/Riyadh
2073;King Khaled Military City;King Khalid Mil.city;Saudi Arabia;HBT;OEKK;27.900917;45.528194;1352;3;U;Asia/Riyadh
2074;Prince Mohammad Bin Abdulaziz;Madinah;Saudi Arabia;MED;OEMA;24.553422;39.705061;2151;3;U;Asia/Riyadh
2075;Nejran;Nejran;Saudi Arabia;EAM;OENG;17.611436;44.419169;3982;3;U;Asia/Riyadh
2076;Qaisumah;Hafr Al-batin;Saudi Arabia;AQI;OEPA;28.335192;46.125069;1174;3;U;Asia/Riyadh
2077;Pump Station 3;Petroline 3;Saudi Arabia;;OEPC;25.174547;47.488431;1740;3;U;Asia/Riyadh
2078;Pump Station 6;Petroline 6;Saudi Arabia;;OEPF;24.710333;44.964527;2530;3;U;Asia/Riyadh
2079;Pump Station 10;Petroline 10;Saudi Arabia;;OEPJ;24.107339;41.036047;2840;3;U;Asia/Riyadh
2080;Rabigh;Rabigh;Saudi Arabia;;OERB;22.702608;39.069842;22;3;U;Asia/Riyadh
2081;Rafha;Rafha;Saudi Arabia;RAH;OERF;29.626419;43.490614;1474;3;U;Asia/Riyadh
2082;King Khaled Intl;Riyadh;Saudi Arabia;RUH;OERK;24.95764;46.698776;2049;3;U;Asia/Riyadh
2083;Ras Mishab;Rash Mishab;Saudi Arabia;;OERM;28.079584;48.610973;13;3;U;Asia/Riyadh
2084;Arar;Arar;Saudi Arabia;RAE;OERR;30.906589;41.138217;1813;3;U;Asia/Riyadh
2085;Ras Tanura;Ras Tanura;Saudi Arabia;;OERT;26.723108;50.030814;6;3;U;Asia/Riyadh
2086;Sharurah;Sharurah;Saudi Arabia;SHW;OESH;17.466875;47.121431;2363;3;U;Asia/Riyadh
6808;Mudgee Airport;Mudgee;Australia;DGE;YMDG;-32.5625;149.611;1545;10;O;Australia/Sydney
2088;Sulayel;Sulayel;Saudi Arabia;SLF;OESL;20.464744;45.619644;2021;3;U;Asia/Riyadh
2089;Tabuk;Tabuk;Saudi Arabia;TUU;OETB;28.365417;36.618889;2551;3;U;Asia/Riyadh
2090;Taif;Taif;Saudi Arabia;TIF;OETF;21.483418;40.544334;4848;3;U;Asia/Riyadh
2091;Thumamah;Thumamah;Saudi Arabia;;OETH;25.212986;46.640978;1870;3;U;Asia/Riyadh
2092;Ras Tanajib;Ras Tanajib;Saudi Arabia;;OETN;27.867844;48.76915;30;3;U;Asia/Riyadh
2093;Turaif;Turaif;Saudi Arabia;TUI;OETR;31.692683;38.7312;2803;3;U;Asia/Riyadh
6872;Miami University Airport;Oxford;United States;OXD;\N;39.5022607;-84.7843814;1040;-5;U;America/New_York
2095;Wejh;Wejh;Saudi Arabia;EJH;OEWJ;26.198553;36.476381;66;3;U;Asia/Riyadh
2096;Yenbo;Yenbo;Saudi Arabia;YNB;OEYN;24.144244;38.06335;26;3;U;Asia/Riyadh
2097;Abadan;Abadan;Iran;ABD;OIAA;30.371111;48.228333;19;3.5;E;Asia/Tehran
2098;Dezful;Dezful;Iran;;OIAD;32.434444;48.39764;474;3.5;E;Asia/Tehran
2099;Aghajari;Aghajari;Iran;;OIAG;30.74445;49.677183;88;3.5;E;Asia/Tehran
2100;Gachsaran;Gachsaran;Iran;;OIAH;30.337567;50.827964;2414;3.5;E;Asia/Tehran
2101;Shahid Asyaee;Masjed Soleiman;Iran;QMJ;OIAI;32.002372;49.270364;1206;3.5;E;Asia/Tehran
2102;Omidiyeh;Omidyeh;Iran;;OIAJ;30.835167;49.534916;85;3.5;E;Asia/Tehran
2103;Mahshahr;Bandar Mahshahr;Iran;MRX;OIAM;30.556192;49.151879;8;3.5;E;Asia/Tehran
2104;Ahwaz;Ahwaz;Iran;AWZ;OIAW;31.337431;48.76195;66;3.5;E;Asia/Tehran
2105;Abumusa Island;Abumusa I.;Iran;;OIBA;25.875742;55.032994;23;3.5;E;Asia/Tehran
2106;Bushehr;Bushehr;Iran;BUZ;OIBB;28.944811;50.834637;68;3.5;E;Asia/Tehran
2107;Bastak;Bastak;Iran;;OIBH;27.212678;54.318592;1350;3.5;E;Asia/Tehran
2108;Asaloyeh;Golbandi;Iran;;OIBI;27.481425;52.615483;15;3.5;E;Asia/Tehran
2109;Kish Island;Kish Island;Iran;KIH;OIBK;26.526156;53.980211;101;3.5;E;Asia/Tehran
2110;Bandar Lengeh;Bandar Lengeh;Iran;BDH;OIBL;26.532;54.824847;67;3.5;E;Asia/Tehran
2111;Khark Island;Khark Island;Iran;;OIBQ;29.260278;50.323889;17;3.5;E;Asia/Tehran
2112;Sirri Island;Siri Island;Iran;;OIBS;25.908869;54.5394;43;3.5;E;Asia/Tehran
2113;Lavan Island;Lavan Island;Iran;;OIBV;26.8103;53.356289;76;3.5;E;Asia/Tehran
2114;Shahid Ashrafi Esfahani;Bakhtaran;Iran;KSH;OICC;34.345853;47.158128;4301;3.5;E;Asia/Tehran
2117;Sanandaj;Sanandaj;Iran;SDG;OICS;35.245856;47.009247;4522;3.5;E;Asia/Tehran
2118;Hesa;Daran;Iran;;OIFE;32.928889;51.561111;5256;3.5;E;Asia/Tehran
2119;Shahid Vatan Pour Air Base;Esfahan;Iran;;OIFH;32.567039;51.691594;5310;3.5;E;Asia/Tehran
2120;Kashan;Kashan;Iran;;OIFK;33.895333;51.577044;3465;3.5;E;Asia/Tehran
2121;Esfahan Shahid Beheshti Intl;Esfahan;Iran;;OIFM;32.750836;51.861267;5059;3.5;E;Asia/Tehran
2122;Badr Air Base;Sepah;Iran;;OIFP;32.621108;51.697017;5242;3.5;E;Asia/Tehran
2123;Rasht;Rasht;Iran;RAS;OIGG;37.325314;49.605817;-40;3.5;E;Asia/Tehran
6871;Delaware County Airport;Muncie;United States;MIE;\N;40.2424722;-85.39575;937;-5;U;America/New_York
2125;Arak;Arak;Iran;;OIHR;34.138147;49.847292;5440;3.5;E;Asia/Tehran
2126;Ghazvin Azadi;Abe-ali;Iran;;OIIA;35.952097;50.450778;3769;3.5;E;Asia/Tehran
2127;Kushke Nosrat;Kushke Nosrat;Iran;;OIIC;34.98395;50.806219;3008;3.5;E;Asia/Tehran
2128;Doshan Tappeh;Teheran;Iran;;OIID;35.702983;51.475131;4046;3.5;E;Asia/Tehran
2129;Imam Khomeini Intl;Abyek;Iran;;OIIE;35.416111;51.152222;3305;3.5;E;Asia/Tehran
2130;Ghale Morghi;Teheran;Iran;;OIIG;35.644806;51.380695;3627;3.5;E;Asia/Tehran
2131;Mehrabad Intl;Teheran;Iran;THR;OIII;35.689167;51.313416;3962;3.5;E;Asia/Tehran
2132;Ghazvin;Ghazvin;Iran;;OIIK;36.240061;50.047153;4184;3.5;E;Asia/Tehran
2133;Naja;Khoram Dareh;Iran;;OIIM;35.776286;50.881014;4040;3.5;E;Asia/Tehran
2134;Bandar Abbass Intl;Bandar Abbas;Iran;BND;OIKB;27.218317;56.37785;22;3.5;E;Asia/Tehran
2135;Jiroft;Jiroft;Iran;;OIKJ;28.726925;57.670269;2663;3.5;E;Asia/Tehran
2136;Kerman;Kerman;Iran;KER;OIKK;30.274444;56.951111;5741;3.5;E;Asia/Tehran
2138;Havadarya;Bandar Abbas;Iran;;OIKP;27.158251;56.172461;19;3.5;E;Asia/Tehran
2139;Dayrestan;Gheshm I.;Iran;;OIKQ;26.754639;55.902353;45;3.5;E;Asia/Tehran
2141;Sirjan;Sirjan;Iran;;OIKY;29.550933;55.672708;5846;3.5;E;Asia/Tehran
2142;Birjand;Birjand;Iran;XBJ;OIMB;32.898056;59.266111;4952;3.5;E;Asia/Tehran
2143;Sarakhs;Sarakhs;Iran;;OIMC;36.501178;61.064903;945;3.5;E;Asia/Tehran
2144;Shahroud;Emam Shahr;Iran;;OIMJ;36.425094;55.104833;4197;3.5;E;Asia/Tehran
6870;Whitsunday Airstrip;Airlie Beach;Australia;WSY;YWHI;-20.276;148.755;60;10;O;Australia/Brisbane
2147;Tabas;Tabas;Iran;;OIMT;33.66775;56.892675;2312;3.5;E;Asia/Tehran
2148;Kalaleh;Kalaleh;Iran;;OINE;37.383272;55.452008;425;3.5;E;Asia/Tehran
2151;Ramsar;Ramsar;Iran;RZR;OINR;36.909908;50.679589;-70;3.5;E;Asia/Tehran
6869;Aeropuerto de Rafaela;Rafaela;Argentina;RAF;\N;-31.282072;-61.501594;100;-3;N;America/Cordoba
2153;Fasa;Fasa;Iran;;OISF;28.891756;53.723339;4261;3.5;E;Asia/Tehran
2154;Jahrom;Jahrom;Iran;;OISJ;28.586675;53.579144;3358;3.5;E;Asia/Tehran
2156;Lamerd;Lamerd;Iran;;OISR;27.372744;53.188794;1337;3.5;E;Asia/Tehran
2157;Shiraz Shahid Dastghaib Intl;Shiraz;Iran;SYZ;OISS;29.539242;52.589786;4920;3.5;E;Asia/Tehran
2158;Khoy;Khoy;Iran;;OITK;38.427453;44.973575;3981;3.5;E;Asia/Tehran
6868;Sabadell Airport;Sabadell;Spain;QSA;LELL;41.5209;2.10508;0;1;E;Europe/Madrid
2162;Tabriz Intl;Tabriz;Iran;TBZ;OITT;38.133889;46.235;4459;3.5;E;Asia/Tehran
2163;Zanjan;Zanjan;Iran;;OITZ;36.77365;48.359422;5382;3.5;E;Asia/Tehran
2164;Yazd Shahid Sadooghi;Yazd;Iran;AZD;OIYY;31.904908;54.276503;4054;3.5;E;Asia/Tehran
2165;Zabol;Zabol;Iran;;OIZB;31.098333;61.543889;1628;3.5;E;Asia/Tehran
2166;Chah Bahar;Chah Bahar;Iran;ZBR;OIZC;25.44335;60.382114;43;3.5;E;Asia/Tehran
2167;Zahedan Intl;Zahedan;Iran;ZAH;OIZH;29.475686;60.906189;4564;3.5;E;Asia/Tehran
2168;Iran Shahr;Iran Shahr;Iran;;OIZI;27.236117;60.720039;2040;3.5;E;Asia/Tehran
2169;Saravan;Saravan;Iran;;OIZS;27.419261;62.315789;3930;3.5;E;Asia/Tehran
2170;Queen Alia Intl;Amman;Jordan;AMM;OJAI;31.722556;35.993214;2395;2;E;Asia/Amman
2171;Marka Intl;Amman;Jordan;ADJ;OJAM;31.972703;35.991569;2555;2;E;Asia/Amman
2172;Aqaba King Hussein Intl;Aqaba;Jordan;AQJ;OJAQ;29.611619;35.018067;175;2;E;Asia/Amman
2173;Prince Hasan;Hotel Four;Jordan;;OJHF;32.160747;37.149383;2220;2;E;Asia/Amman
2174;Jerusalem;Jerusalem;West Bank;;OJJR;31.864722;35.219167;2485;2;U;Asia/Gaza
2175;King Hussein;Mafraq;Jordan;OMF;OJMF;32.356353;36.259181;2240;2;E;Asia/Amman
2176;Kuwait Intl;Kuwait;Kuwait;KWI;OKBK;29.226567;47.968928;206;3;U;Asia/Kuwait
2177;Rafic Hariri Intl;Beirut;Lebanon;BEY;OLBA;33.820931;35.488389;87;2;E;Asia/Beirut
2178;Rene Mouawad;Kleiat;Lebanon;;OLKA;34.589333;36.011322;75;2;E;Asia/Beirut
2179;Abu Dhabi Intl;Abu Dhabi;United Arab Emirates;AUH;OMAA;24.432972;54.651138;88;4;U;Asia/Dubai
2180;Bateen;Abu Dhabi;United Arab Emirates;AZI;OMAD;24.428333;54.458084;16;4;U;Asia/Dubai
2181;Al Hamra Aux;Al Hamra;United Arab Emirates;;OMAH;24.073981;52.463644;50;4;U;Asia/Dubai
2182;Jebel Dhana;Jebel Dhana;United Arab Emirates;;OMAJ;24.187428;52.614;43;4;U;Asia/Dubai
6867;Reykjahlid Airport;Myvatn;Iceland;MVA;BIRL;65.6558;-16.9181;1030;0;N;Atlantic/Reykjavik
2184;Al Dhafra;Abu Dhabi;United Arab Emirates;;OMAM;24.248249;54.547722;77;4;U;Asia/Dubai
2185;Arzanah;Arzana;United Arab Emirates;;OMAR;24.780528;52.559944;15;4;U;Asia/Dubai
2186;Das Island;Das Island;United Arab Emirates;;OMAS;25.146219;52.873711;12;4;U;Asia/Dubai
2187;Zirku;Zirku;United Arab Emirates;;OMAZ;24.861542;53.077833;14;4;U;Asia/Dubai
2188;Dubai Intl;Dubai;United Arab Emirates;DXB;OMDB;25.252778;55.364444;62;4;U;Asia/Dubai
2189;Fujairah Intl;Fujeirah;United Arab Emirates;FJR;OMFJ;25.112225;56.323964;152;4;U;Asia/Dubai
2190;Ras Al Khaimah Intl;Ras Al Khaimah;United Arab Emirates;RKT;OMRK;25.613483;55.938817;102;4;U;Asia/Dubai
2191;Sharjah Intl;Sharjah;United Arab Emirates;SHJ;OMSJ;25.328575;55.51715;111;4;U;Asia/Dubai
2192;Khasab;Khasab;Oman;KHS;OOKB;26.170986;56.240569;100;4;U;Asia/Muscat
2193;Masirah;Masirah;Oman;MSH;OOMA;20.675434;58.890467;64;4;U;Asia/Muscat
2194;Seeb Intl;Muscat;Oman;MCT;OOMS;23.593278;58.284444;48;4;U;Asia/Muscat
2195;Salalah;Salalah;Oman;SLL;OOSA;17.038719;54.091297;73;4;U;Asia/Muscat
2196;Thumrait;Thumrait;Oman;TTH;OOTH;17.666;54.024612;1570;4;U;Asia/Muscat
2197;Bhagatanwala;Bhagtanwala;Pakistan;;OPBG;32.056108;72.9484;600;5;N;Asia/Karachi
6866;Spray View Airport;Spray View;Zimbabwe;;FVSV;-17.917;25.817;3210;2;U;Africa/Harare
6865;Lyon Part-Dieu Railway;Lyon;France;XYD;\N;46;5;821;1;U;Europe/Paris
2202;Faisalabad Intl;Faisalabad;Pakistan;LYP;OPFA;31.365014;72.994842;591;5;N;Asia/Karachi
2203;Gwadar;Gwadar;Pakistan;GWD;OPGD;25.233308;62.329494;36;5;N;Asia/Karachi
2204;Gilgit;Gilgit;Pakistan;GIL;OPGT;35.918786;74.333644;4796;5;N;Asia/Karachi
2205;Shahbaz Ab;Jacobsbad;Pakistan;;OPJA;28.284167;68.449711;185;5;N;Asia/Karachi
2206;Jinnah Intl;Karachi;Pakistan;KHI;OPKC;24.906547;67.160797;100;5;N;Asia/Karachi
2207;Allama Iqbal Intl;Lahore;Pakistan;LHE;OPLA;31.521564;74.403594;712;5;N;Asia/Karachi
2208;Walton;Lahore;Pakistan;;OPLH;31.494794;74.346239;679;5;N;Asia/Karachi
2209;Mangla;Mangla;Pakistan;;OPMA;33.050086;73.638381;902;5;N;Asia/Karachi
2210;Muzaffarabad;Muzaffarabad;Pakistan;MFG;OPMF;34.339022;73.508639;2691;5;N;Asia/Karachi
2211;Mianwali;Mianwali;Pakistan;;OPMI;32.563089;71.570681;690;5;N;Asia/Karachi
2212;Moenjodaro;Moenjodaro;Pakistan;MJD;OPMJ;27.335156;68.143053;154;5;N;Asia/Karachi
2213;Masroor;Karachi;Pakistan;;OPMR;24.893564;66.938753;35;5;N;Asia/Karachi
2214;Multan Intl;Multan;Pakistan;MUX;OPMT;30.203222;71.419111;403;5;N;Asia/Karachi
2215;Nawabshah;Nawabshah;Pakistan;WNS;OPNH;26.219442;68.390053;95;5;N;Asia/Karachi
2216;Okara;Okara;Pakistan;;OPOK;30.741025;73.357736;568;5;N;Asia/Karachi
2217;Panjgur;Panjgur;Pakistan;PJG;OPPG;26.954547;64.132517;3289;5;N;Asia/Karachi
2218;Pasni;Pasni;Pakistan;PSI;OPPI;25.290487;63.345083;33;5;N;Asia/Karachi
2219;Peshawar Intl;Peshawar;Pakistan;PEW;OPPS;33.993911;71.514581;1158;5;N;Asia/Karachi
2220;Qasim;Qasim;Pakistan;;OPQS;33.560153;73.033217;1581;5;N;Asia/Karachi
2221;Quetta;Quetta;Pakistan;UET;OPQT;30.251369;66.937764;5267;5;N;Asia/Karachi
2222;Sheikh Zayed;Rahim Yar Khan;Pakistan;RYK;OPRK;28.3839;70.279572;271;5;N;Asia/Karachi
2223;Chaklala;Islamabad;Pakistan;ISB;OPRN;33.616653;73.099233;1668;5;N;Asia/Karachi
2224;Risalpur;Risalpur;Pakistan;;OPRS;34.081112;71.972583;1050;5;N;Asia/Karachi
2225;Rawalakot;Rawala Kot;Pakistan;RAZ;OPRT;33.849658;73.798147;5479;5;N;Asia/Karachi
6864;Point Roberts Airpark;Point Roberts;United States;1RL;K1RL;48.9797222;-123.0788889;10;-8;A;America/Los_Angeles
2227;Sukkur;Sukkur;Pakistan;SKZ;OPSK;27.721989;68.791683;196;5;N;Asia/Karachi
2228;Saidu Sharif;Saidu Sharif;Pakistan;SDT;OPSS;34.813556;72.352814;3183;5;N;Asia/Karachi
2229;Sui;Sui;Pakistan;SUL;OPSU;28.645142;69.176917;763;5;N;Asia/Karachi
2230;Talhar;Talhar;Pakistan;BDN;OPTH;24.841519;68.838408;28;5;N;Asia/Karachi
2232;Wana;Wana;Pakistan;;OPWN;32.304664;69.570433;4550;5;N;Asia/Karachi
2233;Zhob;Zhob;Pakistan;PZH;OPZB;31.358381;69.463606;4728;5;N;Asia/Karachi
2234;Basrah Intl;Basrah;Iraq;BSR;ORMM;30.549069;47.662142;11;3;U;Asia/Baghdad
2235;Aleppo Intl;Aleppo;Syria;ALP;OSAP;36.180675;37.224358;1276;2;E;Asia/Damascus
2236;Damascus Intl;Damascus;Syria;DAM;OSDI;33.410636;36.514267;2020;2;E;Asia/Damascus
2237;Deir Zzor;Deire Zor;Syria;DEZ;OSDZ;35.285374;40.175961;700;2;E;Asia/Damascus
2239;Bassel Al Assad Intl;Latakia;Syria;LTK;OSLK;35.401094;35.948681;157;2;E;Asia/Damascus
2240;Palmyra;Palmyra;Syria;PMS;OSPR;34.557361;38.316889;1322;2;E;Asia/Damascus
2241;Doha Intl;Doha;Qatar;DOH;OTBD;25.261125;51.565056;35;3;U;Asia/Qatar
2242;Canton;Canton Island;Kiribati;CIS;PCIS;-2.768122;-171.710394;9;13;U;Pacific/Enderbury
2243;Rota Intl;Rota;Northern Mariana Islands;ROP;PGRO;14.174308;145.242536;607;10;U;Pacific/Saipan
2244;Francisco C Ada Saipan Intl;Saipan;Northern Mariana Islands;SPN;PGSN;15.119003;145.729356;215;10;U;Pacific/Saipan
2245;Andersen Afb;Andersen;Guam;UAM;PGUA;13.583953;144.930025;627;10;U;Pacific/Guam
2246;Guam Intl;Agana;Guam;GUM;PGUM;13.48345;144.795983;298;10;U;Pacific/Guam
2247;Tinian Intl;West Tinian;Northern Mariana Islands;TIQ;PGWT;14.999203;145.61935;271;10;U;Pacific/Saipan
6807;Scone Airport;Scone;Australia;NSO;YSCO;-32.0372;150.832;745;10;O;Australia/Sydney
2249;Marshall Islands Intl;Majuro;Marshall Islands;MAJ;PKMJ;7.064758;171.272022;6;12;U;Pacific/Majuro
2250;Dyess Aaf;Kwajalein;Marshall Islands;;PKRO;9.396886;167.470869;9;12;U;Pacific/Majuro
2251;Bucholz Aaf;Kwajalein;Marshall Islands;KWA;PKWA;8.720122;167.731661;9;12;U;Pacific/Majuro
2252;Cassidy Intl;Kiritimati;Kiribati;CXI;PLCH;1.986161;-157.349778;5;-12;U;\N
2253;Midway Atoll;Midway;Midway Islands;MDY;PMDY;28.201725;-177.380636;13;-11;U;Pacific/Midway
2254;Chuuk Intl;Chuuk;Micronesia;TKK;PTKK;7.461869;151.843006;11;10;U;Pacific/Truk
2255;Pohnpei Intl;Pohnpei;Micronesia;PNI;PTPN;6.9851;158.208989;10;11;U;Pacific/Ponape
2256;Babelthuap;Babelthuap;Palau;ROR;PTRO;7.367303;134.544278;176;9;U;Pacific/Palau
2257;Kosrae;Kosrae;Micronesia;KSA;PTSA;5.356975;162.958386;11;11;U;Pacific/Kosrae
2258;Yap Intl;Yap;Micronesia;YAP;PTYA;9.498911;138.082497;91;10;U;Pacific/Truk
2259;Shang Yi;Kinmen;Taiwan;KNH;RCBS;24.427892;118.359197;93;8;U;Asia/Taipei
2260;Pingtung South;Pingtung;Taiwan;PIF;RCDC;22.672367;120.461728;78;8;U;Asia/Taipei
2261;Longtan;Longtang;Taiwan;;RCDI;24.855136;121.237636;790;8;U;Asia/Taipei
2262;Fengnin;Fengnin;Taiwan;TTT;RCFN;22.754986;121.101681;143;8;U;Asia/Taipei
2263;Lyudao;Green Island;Taiwan;GNI;RCGI;22.673853;121.466481;28;8;U;Asia/Taipei
2264;Kaohsiung Intl;Kaohsiung;Taiwan;KHH;RCKH;22.577094;120.350006;31;8;U;Asia/Taipei
2265;Chiayi;Chiayi;Taiwan;CYI;RCKU;23.461779;120.39283;85;8;U;Asia/Taipei
2267;Lanyu;Lanyu;Taiwan;KYD;RCLY;22.028842;121.533642;44;8;U;Asia/Taipei
2268;Ching Chuang Kang;Taichung;Taiwan;RMQ;RCMQ;24.264668;120.62058;663;8;N;Asia/Taipei
6806;Cessnock Airport;Cessnock;Australia;CES;YCNK;-32.7875;151.342;211;10;O;Australia/Sydney
2270;Tainan;Tainan;Taiwan;TNN;RCNN;22.950361;120.205778;63;8;U;Asia/Taipei
2271;Hsinchu;Hsinchu;Taiwan;;RCPO;24.818033;120.939394;26;8;U;Asia/Taipei
2272;Magong;Makung;Taiwan;MZG;RCQC;23.568669;119.628311;103;8;U;Asia/Taipei
2273;Jhihhang;Taitung;Taiwan;;RCQS;22.793117;121.181975;128;8;U;Asia/Taipei
2274;Pingtung North;Pingtung;Taiwan;;RCSQ;22.700239;120.482;97;8;U;Asia/Taipei
2275;Sungshan;Taipei;Taiwan;TSA;RCSS;25.069722;121.5525;18;8;U;Asia/Taipei
2276;Taoyuan Intl;Taipei;Taiwan;TPE;RCTP;25.077731;121.232822;106;8;U;Asia/Taipei
2277;Wang An;Wang An;Taiwan;WOT;RCWA;23.370833;119.494444;115;8;U;Asia/Taipei
2278;Hualien;Hualien;Taiwan;HUN;RCYU;24.023725;121.616906;52;8;U;Asia/Taipei
2279;Narita Intl;Tokyo;Japan;NRT;RJAA;35.764722;140.386389;141;9;U;Asia/Tokyo
2280;Matsumoto;Matsumoto;Japan;MMJ;RJAF;36.166758;137.922669;2182;9;U;Asia/Tokyo
2281;Hyakuri;Ibaraki;Japan;IBR;RJAH;36.181083;140.415444;105;9;U;Asia/Tokyo
2282;Minami Torishima;Minami Tori Shima;Japan;;RJAM;24.289697;153.979119;22;9;U;Asia/Tokyo
2283;Iwo Jima;Iwojima;Japan;IWO;RJAW;24.784;141.322722;384;9;U;Asia/Tokyo
2284;Nanki Shirahama;Nanki-shirahama;Japan;SHM;RJBD;33.662222;135.364444;298;9;U;Asia/Tokyo
2285;Kohnan;Kohnan;Japan;;RJBK;34.590836;133.933225;3;9;U;Asia/Tokyo
2286;Obihiro;Obihiro;Japan;OBO;RJCB;42.733333;143.217222;505;9;U;Asia/Tokyo
2287;New Chitose;Sapporo;Japan;CTS;RJCC;42.7752;141.692283;82;9;U;Asia/Tokyo
2288;Hakodate;Hakodate;Japan;HKD;RJCH;41.77;140.821944;151;9;U;Asia/Tokyo
2289;Chitose;Chitose;Japan;SPK;RJCJ;42.794475;141.666447;87;9;U;Asia/Tokyo
2290;Memanbetsu;Memanbetsu;Japan;MMB;RJCM;43.880606;144.164053;135;9;U;Asia/Tokyo
2291;Nakashibetsu;Nakashibetsu;Japan;SHB;RJCN;43.5775;144.96;234;9;U;Asia/Tokyo
2293;Tokachi;Tokachi;Japan;;RJCT;42.890544;143.158475;281;9;U;Asia/Tokyo
2294;Wakkanai;Wakkanai;Japan;WKJ;RJCW;45.404167;141.800833;30;9;U;Asia/Tokyo
2295;Iki;Iki;Japan;IKI;RJDB;33.749027;129.785417;41;9;U;Asia/Tokyo
2296;Yamaguchi Ube;Yamaguchi;Japan;UBJ;RJDC;33.93;131.278611;23;9;U;Asia/Tokyo
2297;Tsushima;Tsushima;Japan;TSJ;RJDT;34.284889;129.33055;213;9;U;Asia/Tokyo
2298;Monbetsu;Monbetsu;Japan;MBE;RJEB;44.303914;143.404028;80;9;U;Asia/Tokyo
2299;Asahikawa;Asahikawa;Japan;AKJ;RJEC;43.670833;142.4475;721;9;U;Asia/Tokyo
2300;Okushiri;Okushiri;Japan;OIR;RJEO;42.071667;139.432911;161;9;U;Asia/Tokyo
2301;Rishiri;Rishiri Island;Japan;RIS;RJER;45.242006;141.186431;112;9;U;Asia/Tokyo
2302;Ashiya;Ashiya;Japan;;RJFA;33.883083;130.653;98;9;U;Asia/Tokyo
2303;Yakushima;Yakushima;Japan;KUM;RJFC;30.385569;130.659017;124;9;U;Asia/Tokyo
2304;Fukue;Fukue;Japan;FUJ;RJFE;32.666269;128.832808;273;9;U;Asia/Tokyo
2305;Fukuoka;Fukuoka;Japan;FUK;RJFF;33.585942;130.450686;32;9;U;Asia/Tokyo
2306;New Tanegashima;Tanegashima;Japan;TNE;RJFG;30.605067;130.991231;768;9;U;Asia/Tokyo
2307;Kagoshima;Kagoshima;Japan;KOJ;RJFK;31.803397;130.719408;906;9;U;Asia/Tokyo
2308;Miyazaki;Miyazaki;Japan;KMI;RJFM;31.877222;131.448611;20;9;U;Asia/Tokyo
2309;Nyutabaru;Nyutabaru;Japan;;RJFN;32.083611;131.451389;259;9;U;Asia/Tokyo
2310;Oita;Oita;Japan;OIT;RJFO;33.479444;131.737222;19;9;U;Asia/Tokyo
2311;New Kitakyushu;Kitakyushu;Japan;KKJ;RJFR;33.845942;131.034689;21;9;U;Asia/Tokyo
2312;Kumamoto;Kumamoto;Japan;KMJ;RJFT;32.837319;130.85505;642;9;U;Asia/Tokyo
2313;Nagasaki;Nagasaki;Japan;NGS;RJFU;32.916944;129.913611;15;9;U;Asia/Tokyo
2314;Kanoya;Kanoya;Japan;;RJFY;31.367608;130.845456;214;9;U;Asia/Tokyo
2315;Tsuiki;Tsuiki;Japan;;RJFZ;33.685;131.040278;55;9;U;Asia/Tokyo
2316;Amami;Amami;Japan;ASJ;RJKA;28.430633;129.712542;27;9;U;Asia/Tokyo
2317;Okierabu;Okierabu;Japan;;RJKB;27.425522;128.700903;101;9;U;Asia/Tokyo
2318;Tokunoshima;Tokunoshima;Japan;TKN;RJKN;27.836381;128.881253;17;9;U;Asia/Tokyo
2319;Fukui;Fukui;Japan;;RJNF;36.142847;136.223922;19;9;U;Asia/Tokyo
2320;Gifu;Gifu;Japan;;RJNG;35.394078;136.869667;128;9;U;Asia/Tokyo
2321;Hamamatsu;Hamamatsu;Japan;;RJNH;34.750239;137.703083;150;9;U;Asia/Tokyo
2322;Komatsu;Kanazawa;Japan;KMQ;RJNK;36.394611;136.406544;36;9;U;Asia/Tokyo
2323;Oki;Oki Island;Japan;OKI;RJNO;36.181125;133.324844;311;9;U;Asia/Tokyo
2324;Toyama;Toyama;Japan;TOY;RJNT;36.648333;137.1875;95;9;U;Asia/Tokyo
2325;Shizuhama;Yaizu;Japan;;RJNY;34.812778;138.298056;23;9;U;Asia/Tokyo
2326;Hiroshima;Hiroshima;Japan;HIJ;RJOA;34.436111;132.919444;1088;9;U;Asia/Tokyo
2327;Okayama;Okayama;Japan;OKJ;RJOB;34.756944;133.855278;806;9;U;Asia/Tokyo
2328;Izumo;Izumo;Japan;IZO;RJOC;35.413611;132.89;15;9;U;Asia/Tokyo
2329;Hofu;Hofu;Japan;;RJOF;34.034667;131.549194;7;9;U;Asia/Tokyo
2330;Yonago Kitaro;Miho;Japan;YGJ;RJOH;35.492222;133.236389;20;9;U;Asia/Tokyo
2331;Iwakuni Mcas;Iwakuni;Japan;;RJOI;34.14386;132.23575;7;9;U;Asia/Tokyo
2332;Kochi;Kochi;Japan;KCZ;RJOK;33.546111;133.669444;42;9;U;Asia/Tokyo
2333;Matsuyama;Matsuyama;Japan;MYJ;RJOM;33.827222;132.699722;25;9;U;Asia/Tokyo
2334;Osaka Intl;Osaka;Japan;ITM;RJOO;34.785528;135.438222;50;9;U;Asia/Tokyo
2335;Tottori;Tottori;Japan;TTJ;RJOR;35.530069;134.166553;65;9;U;Asia/Tokyo
2336;Tokushima;Tokushima;Japan;TKS;RJOS;34.132808;134.606639;26;9;U;Asia/Tokyo
2337;Takamatsu;Takamatsu;Japan;TAK;RJOT;34.214167;134.015556;607;9;U;Asia/Tokyo
2338;Yao;Osaka;Japan;;RJOY;34.596311;135.602944;39;9;U;Asia/Tokyo
2339;Ozuki;Ozuki;Japan;;RJOZ;34.045322;131.052144;13;9;U;Asia/Tokyo
2340;Aomori;Aomori;Japan;AOJ;RJSA;40.734722;140.690833;664;9;U;Asia/Tokyo
2341;Yamagata;Yamagata;Japan;GAJ;RJSC;38.411894;140.371331;353;9;U;Asia/Tokyo
2342;Sado;Sado;Japan;;RJSD;38.060181;138.413928;88;9;U;Asia/Tokyo
2343;Hachinohe;Hachinoe;Japan;;RJSH;40.556447;141.466325;152;9;U;Asia/Tokyo
2344;Hanamaki;Hanamaki;Japan;HNA;RJSI;39.428611;141.135278;297;9;U;Asia/Tokyo
2345;Akita;Akita;Japan;AXT;RJSK;39.615556;140.218611;313;9;U;Asia/Tokyo
2346;Misawa Ab;Misawa;Japan;MSJ;RJSM;40.703222;141.368364;119;9;U;Asia/Tokyo
2347;Sendai;Sendai;Japan;SDJ;RJSS;38.139722;140.916944;15;9;U;Asia/Tokyo
2348;Matsushima;Matsushima;Japan;;RJST;38.404919;141.219572;7;9;U;Asia/Tokyo
6862;Sazena;Sazena;Czech Republic;LKS;LKSZ;50.1929;14.1532;233;1;E;Europe/Prague
2350;Atsugi Naf;Atsugi;Japan;;RJTA;35.454611;139.450167;205;9;U;Asia/Tokyo
2351;Tateyama;Tateyama;Japan;;RJTE;34.987053;139.829208;10;9;U;Asia/Tokyo
2352;Hachijojima;Hachijojima;Japan;HAC;RJTH;33.115;139.785833;303;9;U;Asia/Tokyo
2353;Iruma;Iruma;Japan;;RJTJ;35.841944;139.410556;295;9;U;Asia/Tokyo
2354;Kisarazu;Kisarazu;Japan;;RJTK;35.398272;139.909936;10;9;U;Asia/Tokyo
2355;Shimofusa;Shimofusa;Japan;;RJTL;35.798944;140.011111;98;9;U;Asia/Tokyo
2356;Oshima;Oshima;Japan;OIM;RJTO;34.782033;139.360306;130;9;U;Asia/Tokyo
2358;Kastner Aaf;Zama;Japan;;RJTR;35.513769;139.393675;360;9;U;Asia/Tokyo
2359;Tokyo Intl;Tokyo;Japan;HND;RJTT;35.552258;139.779694;35;9;U;Asia/Tokyo
2360;Yokota Ab;Yokota;Japan;OKO;RJTY;35.748492;139.348483;463;9;U;Asia/Tokyo
2361;Gwangju;Kwangju;South Korea;KWJ;RKJJ;35.126389;126.808889;39;9;U;Asia/Seoul
2364;Jeon Ju;Jhunju;South Korea;;RKJU;35.878436;127.11955;96;9;U;Asia/Seoul
2365;Yeosu;Yeosu;South Korea;RSU;RKJY;34.842328;127.61685;53;9;U;Asia/Seoul
2366;Sokcho;Sokch'o;South Korea;SHO;RKND;38.142614;128.598556;92;9;U;Asia/Seoul
2367;Gangneung;Kangnung;South Korea;KAG;RKNN;37.753561;128.943625;35;9;U;Asia/Seoul
6861;Macomb Municipal Airport;Macomb;United States;MQB;\N;40.5200833;-90.6523889;707;-6;U;America/Chicago
2370;Jeju Intl;Cheju;South Korea;CJU;RKPC;33.511306;126.493028;118;9;U;Asia/Seoul
2371;Jinhae;Chinhae;South Korea;;RKPE;35.141175;128.695792;8;9;U;Asia/Seoul
2372;Gimhae Intl;Busan;South Korea;PUS;RKPK;35.179528;128.938222;6;9;U;Asia/Seoul
6860;Bungle Bungle Airport;Bungle Bungle;Australia;;YBUU;-17.5453;128.307;0;8;U;Australia/Perth
2374;Ulsan;Ulsan;South Korea;USN;RKPU;35.593494;129.351722;45;9;U;Asia/Seoul
2375;A 511;Pyongtaek;South Korea;;RKSG;36.962214;127.031072;51;9;U;Asia/Seoul
2376;Seoul Ab;Seoul East;South Korea;SSN;RKSM;37.445833;127.113889;92;9;U;Asia/Seoul
2377;Osan Ab;Osan;South Korea;OSN;RKSO;37.090617;127.029594;40;9;U;Asia/Seoul
2378;Gimpo;Seoul;South Korea;GMP;RKSS;37.558311;126.790586;58;9;U;Asia/Seoul
2379;Suwon;Suwon;South Korea;;RKSW;37.239406;127.007053;88;9;U;Asia/Seoul
2380;Pohang;Pohang;South Korea;KPO;RKTH;35.987858;129.420486;70;9;U;Asia/Seoul
2381;Daegu Ab;Taegu;South Korea;TAE;RKTN;35.894108;128.658856;116;9;U;Asia/Seoul
2383;Yecheon;Yechon;South Korea;YEC;RKTY;36.631933;128.35485;354;9;U;Asia/Seoul
2384;Naha;Okinawa;Japan;OKA;ROAH;26.195814;127.645869;12;9;N;Asia/Tokyo
2385;Ie Shima Aux Ab;Iejima;Japan;;RODE;26.728775;127.761775;184;9;U;Asia/Tokyo
2386;Kadena Ab;Kadena;Japan;DNA;RODN;26.355612;127.767633;143;9;U;Asia/Tokyo
2387;Ishigaki;Ishigaki;Japan;ISG;ROIG;24.344525;124.186983;93;9;U;Asia/Tokyo
2388;Kumejima;Kumejima;Japan;UEO;ROKJ;26.363506;126.713806;23;9;U;Asia/Tokyo
2389;Minami Daito;Minami Daito;Japan;MMD;ROMD;25.846533;131.263494;167;9;U;Asia/Tokyo
2390;Miyako;Miyako;Japan;MMY;ROMY;24.782833;125.295111;150;9;U;Asia/Tokyo
2391;Kitadaito;Kitadaito;Japan;KTD;RORK;25.944722;131.326944;80;9;U;Asia/Tokyo
2392;Shimojishima;Shimojishima;Japan;SHI;RORS;24.826667;125.144722;54;9;U;Asia/Tokyo
2393;Tarama;Tarama;Japan;;RORT;24.653889;124.675278;36;9;U;Asia/Tokyo
2394;Yoron;Yoron;Japan;RNJ;RORY;27.043964;128.401517;52;9;U;Asia/Tokyo
2395;Futenma Mcas;Futema;Japan;;ROTM;26.274275;127.756494;247;9;U;Asia/Tokyo
2396;Yonaguni;Yonaguni Jima;Japan;OGN;ROYN;24.466944;122.977778;70;9;U;Asia/Tokyo
2397;Ninoy Aquino Intl;Manila;Philippines;MNL;RPLL;14.508647;121.019581;75;8;N;Asia/Manila
6859;Mary Airport;Mary;Turkmenistan;MYP;UTAM;37.6194;61.8967;728;5;U;Asia/Ashgabat
2399;Cotabato;Cotabato;Philippines;CBO;RPMC;7.165242;124.209619;189;8;N;Asia/Manila
2400;Cagayan De Oro;Ladag;Philippines;CGY;RPML;8.415619;124.611219;601;8;N;Asia/Manila
2401;Pagadian;Pagadian;Philippines;PAG;RPMP;7.827197;123.458294;5;8;N;Asia/Manila
2402;Tambler;Romblon;Philippines;;RPMR;6.058003;125.096033;505;8;N;Asia/Manila
6858;Crocodile Camp;Tsavo East;Kenya;;\N;-3.070875;39.2194305555556;400;3;U;Africa/Nairobi
2404;Zamboanga Intl;Zamboanga;Philippines;ZAM;RPMZ;6.922419;122.059633;33;8;N;Asia/Manila
2405;Baguio;Baguio;Philippines;BAG;RPUB;16.375103;120.619636;4251;8;N;Asia/Manila
2406;Daet;Daet;Philippines;;RPUD;14.129167;122.980181;10;8;N;Asia/Manila
2407;Basa Ab;Floridablanca;Philippines;;RPUF;14.986527;120.4925;151;8;N;Asia/Manila
2408;Lingayen;Lingayen;Philippines;;RPUG;16.034786;120.241106;7;8;N;Asia/Manila
2409;San Jose;San Jose;Philippines;;RPUH;12.361517;121.046639;14;8;N;Asia/Manila
2410;Fernando Ab;Lipa;Philippines;;RPUL;13.955017;121.124925;1220;8;N;Asia/Manila
2411;Mamburao;Mamburao;Philippines;;RPUM;13.208092;120.60535;13;8;N;Asia/Manila
2414;Vigan;Vigan;Philippines;;RPUQ;17.555331;120.355797;16;8;N;Asia/Manila
2415;Baler;Baler;Philippines;;RPUR;15.729836;121.500133;108;8;N;Asia/Manila
6857;Kimana;Kimana;Kenya;;\N;-2.8;37.5333333333333;3900;3;U;Africa/Nairobi
6856;Solio Ranch Airport;Solio;Kenya;;\N;-0.245167;36.8849;6300;3;U;Africa/Nairobi
6855;Kalundborg Flyveplads;Kalundborg;Denmark;;EKKL;55.7001;11.2549;13;1;E;Europe/Copenhagen
2421;Bagabag;Bagabag;Philippines;;RPUZ;16.619194;121.252319;820;8;N;Asia/Manila
2422;Daniel Z Romualdez;Tacloban;Philippines;TAC;RPVA;11.227628;125.027758;10;8;N;Asia/Manila
2423;Bacolod;Bacolod;Philippines;BCD;RPVB;10.642511;122.929617;25;8;N;Asia/Manila
6854;Gardermoen Airport;Oslo;Norway;GEN;\N;60.1939;11.1004;681;1;E;Europe/Oslo
2425;Dumaguete;Dumaguete;Philippines;DGT;RPVD;9.333714;123.300472;15;8;N;Asia/Manila
2426;Caticlan;Caticlan;Philippines;;RPVE;11.924503;121.95405;7;8;N;Asia/Manila
2428;Guiuan;Guiuan;Philippines;;RPVG;11.035544;125.741594;7;8;N;Asia/Manila
2429;Iloilo;Iloilo;Philippines;ILO;RPVI;10.713044;122.545297;27;8;N;Asia/Manila
2430;Kalibo;Kalibo;Philippines;KLO;RPVK;11.679431;122.376294;14;8;N;Asia/Manila
2431;Mactan Cebu Intl;Masbate;Philippines;;RPVM;10.307542;123.979439;31;8;N;Asia/Manila
6853;Wausau Downtown Airport;Wausau;United States;AUW;\N;44.9262845;-89.6270018;1201;-6;U;America/Chicago
2433;Puerto Princesa;Puerto Princesa;Philippines;PPS;RPVP;9.742119;118.758731;71;8;N;Asia/Manila
2435;Antique;San Jose;Philippines;SJI;RPVS;10.766044;121.933439;23;8;N;Asia/Manila
2436;Comodoro Pierrestegui;Concordia;Argentina;COC;SAAC;-31.296944;-57.996631;112;-3;N;America/Cordoba
2437;Gualeguaychu;Gualeguaychu;Argentina;GHU;SAAG;-33.010278;-58.613056;75;-3;N;America/Cordoba
2438;Junin;Junin;Argentina;;SAAJ;-34.545889;-60.930556;262;-3;N;America/Cordoba
2439;General Urquiza;Parana;Argentina;PRA;SAAP;-31.794778;-60.480361;243;-3;N;America/Cordoba
2440;Rosario;Rosario;Argentina;ROS;SAAR;-32.903611;-60.785;85;-3;N;America/Cordoba
2441;Sauce Viejo;Santa Fe;Argentina;SFN;SAAV;-31.711666;-60.811668;56;-3;N;America/Cordoba
2442;Aeroparque Jorge Newbery;Buenos Aires;Argentina;AEP;SABE;-34.559175;-58.415606;18;-3;N;America/Cordoba
2443;Ambrosio L V Taravella;Cordoba;Argentina;COR;SACO;-31.323619;-64.207953;1604;-3;N;America/Cordoba
2444;Chamical;Gobernador Gordillo;Argentina;;SACT;-30.345278;-66.29361;1503;-3;N;America/Cordoba
2445;San Fernando;San Fernando;Argentina;;SADF;-34.453189;-58.589644;10;-3;N;America/Cordoba
2446;Mariano Moreno;Jose C. Paz;Argentina;;SADJ;-34.56065;-58.789564;105;-3;N;America/Cordoba
2447;La Plata;La Plata;Argentina;LPG;SADL;-34.972222;-57.894694;72;-3;N;America/Cordoba
2448;Moron;Moron;Argentina;;SADM;-34.676317;-58.642756;95;-3;N;America/Cordoba
2449;El Palomar;El Palomar;Argentina;;SADP;-34.609939;-58.612592;59;-3;N;America/Cordoba
2450;Chos Malal;Chosmadal;Argentina;;SAHC;-37.444692;-70.222469;2788;-3;N;America/Cordoba
2451;General Roca;Fuerte Gral Roca;Argentina;;SAHR;-39.000672;-67.620514;853;-3;N;America/Cordoba
2452;El Plumerillo;Mendoza;Argentina;MDZ;SAME;-32.831717;-68.792856;2310;-3;N;America/Cordoba
2453;Malargue;Malargue;Argentina;LGS;SAMM;-35.493597;-69.574267;4683;-3;N;America/Cordoba
2454;San Rafael;San Rafael;Argentina;AFA;SAMR;-34.588314;-68.403854;2470;-3;N;America/Cordoba
2455;Catamarca;Catamarca;Argentina;CTC;SANC;-28.593214;-65.750925;1522;-3;N;America/Cordoba
2456;Santiago Del Estero;Santiago Del Estero;Argentina;SDE;SANE;-27.765617;-64.310122;656;-3;N;America/Cordoba
2457;Tinogasta;Tinogasta;Argentina;;SANI;-28.03775;-67.580314;3968;-3;N;America/Cordoba
2458;La Rioja;La Rioja;Argentina;IRJ;SANL;-29.381636;-66.795839;1436;-3;N;America/Cordoba
2459;Chilecito;Chilecito;Argentina;;SANO;-29.223888;-67.438889;3099;-3;N;America/Cordoba
2460;Teniente Benjamin Matienzo;Tucuman;Argentina;TUC;SANT;-26.840861;-65.104944;1495;-3;N;America/Cordoba
2461;San Juan;San Julian;Argentina;UAQ;SANU;-31.571472;-68.418194;1959;-3;N;America/Cordoba
2462;Rio Cuarto Area De Material;Rio Cuarto;Argentina;RCU;SAOC;-33.085128;-64.261314;1381;-3;N;America/Cordoba
2463;Villa Dolores;Villa Dolores;Argentina;VDR;SAOD;-31.945183;-65.146283;1915;-3;N;America/Cordoba
2464;Laboulaye;Laboulaye;Argentina;;SAOL;-34.135444;-63.36225;449;-3;N;America/Cordoba
2465;Marcos Juarez;Marcos Juarez;Argentina;;SAOM;-32.683639;-62.157792;361;-3;N;America/Cordoba
2466;Villa Reynolds;Villa Reynolds;Argentina;;SAOR;-33.725144;-65.378086;1591;-3;N;America/Cordoba
2467;San Luis;San Luis;Argentina;LUQ;SAOU;-33.273192;-66.356422;2329;-3;N;America/Cordoba
2468;Corrientes;Corrientes;Argentina;CNQ;SARC;-27.445503;-58.761864;203;-3;N;America/Cordoba
2469;Resistencia;Resistencia;Argentina;RES;SARE;-27.449986;-59.056125;173;-3;N;America/Cordoba
2470;Formosa;Formosa;Argentina;FMA;SARF;-26.212722;-58.228111;194;-3;N;America/Cordoba
2471;Cataratas Del Iguazu;Iguazu Falls;Argentina;IGR;SARI;-25.737281;-54.473444;916;-3;N;America/Cordoba
2472;Paso De Los Libres;Paso De Los Libres;Argentina;AOL;SARL;-29.689425;-57.152078;230;-3;N;America/Cordoba
2473;Monte Caseros;Monte Caseros;Argentina;;SARM;-30.271922;-57.640231;174;-3;N;America/Cordoba
2474;Posadas;Posadas;Argentina;PSS;SARP;-27.385839;-55.970728;430;-3;N;America/Cordoba
2475;Termal;Presidencia R.s.pena;Argentina;;SARS;-26.756519;-60.493103;307;-3;N;America/Cordoba
2476;Salta;Salta;Argentina;SLA;SASA;-24.855978;-65.486169;4088;-3;N;America/Cordoba
2477;Jujuy;Jujuy;Argentina;JUJ;SASJ;-24.392778;-65.097778;3019;-3;N;America/Cordoba
2478;Oran;Oran;Argentina;ORA;SASO;-23.152779;-64.32917;1168;-3;N;America/Cordoba
2479;La Quiaca;La Quiaca;Argentina;;SASQ;-22.150556;-65.5775;11414;-3;N;America/Cordoba
6852;Portage Municipal Airport;Portage;United States;C47;\N;43.5603136;-89.4828607;825;-6;U;America/Chicago
2481;Eldorado;El Dorado;Argentina;;SATD;-26.397499;-54.574722;685;-3;N;America/Cordoba
2482;Goya;Goya;Argentina;;SATG;-29.099472;-59.250583;125;-3;N;America/Cordoba
2483;Obera;Obera;Argentina;;SATO;-27.518156;-55.124156;1125;-3;N;America/Cordoba
2484;Reconquista;Reconquista;Argentina;;SATR;-29.210278;-59.68;161;-3;N;America/Cordoba
2485;Curuzu Cuatia;Curuzu Cuatia;Argentina;;SATU;-29.770555;-57.978889;229;-3;N;America/Cordoba
2486;El Bolson;El Bolson;Argentina;EHL;SAVB;-41.943189;-71.532289;1131;-3;N;America/Cordoba
2487;Comodoro Rivadavia;Comodoro Rivadavia;Argentina;CRD;SAVC;-45.785347;-67.465508;190;-3;N;America/Cordoba
2488;Esquel;Esquel;Argentina;EQS;SAVE;-42.90795;-71.139472;2621;-3;N;America/Cordoba
2490;Almirante Zar;Trelew;Argentina;REL;SAVT;-43.2105;-65.270319;141;-3;N;America/Cordoba
2491;Gobernador Castello;Viedma;Argentina;VDM;SAVV;-40.869222;-63.000389;20;-3;N;America/Cordoba
2492;El Tehuelche;Puerto Madryn;Argentina;PMY;SAVY;-42.759161;-65.102725;426;-3;N;America/Cordoba
2493;Base Marambio;Marambio Base;Antarctica;;SAWB;-64.238335;-56.630833;760;12;U;Antarctica/South_Pole
2494;Puerto Deseado;Puerto Deseado;Argentina;PUD;SAWD;-47.735292;-65.904097;266;-3;N;America/Cordoba
2495;Rio Grande;Rio Grande;Argentina;RGA;SAWE;-53.777667;-67.749389;65;-3;N;America/Cordoba
2496;Rio Gallegos;Rio Gallegos;Argentina;RGL;SAWG;-51.608875;-69.312636;61;-3;N;America/Cordoba
2497;Ushuaia Malvinas Argentinas;Ushuaia;Argentina;USH;SAWH;-54.843278;-68.29575;71;-3;N;America/Cordoba
2498;San Julian;San Julian;Argentina;ULA;SAWJ;-49.306775;-67.802589;190;-3;N;America/Cordoba
2499;Perito Moreno;Perito Moreno;Argentina;PMQ;SAWP;-46.537911;-70.978689;1410;-3;N;America/Cordoba
2500;Santa Cruz;Santa Cruz;Argentina;RZA;SAWU;-50.01655;-68.579197;364;-3;N;America/Cordoba
2501;Comandante Espora;Bahia Blanca;Argentina;BHI;SAZB;-38.724967;-62.169317;246;-3;N;America/Cordoba
2502;Coronel Suarez;Colonel Suarez;Argentina;;SAZC;-37.446111;-61.889297;768;-3;N;America/Cordoba
2503;Olavarria;Olavarria;Argentina;;SAZF;-36.890039;-60.216619;551;-3;N;America/Cordoba
2504;General Pico;General Pico;Argentina;;SAZG;-35.696183;-63.758286;459;-3;N;America/Cordoba
2505;Tres Arroyos;Tres Arroyos;Argentina;;SAZH;-38.386911;-60.329711;400;-3;N;America/Cordoba
2506;Bolivar;Bolivar;Argentina;;SAZI;-36.186594;-61.076367;308;-3;N;America/Cordoba
6851;Gogebic Iron County Airport;Ironwood;United States;IWD;\N;46.5274747;-90.1313967;1230;-6;U;America/Chicago
2508;Mar Del Plata;Mar Del Plata;Argentina;MDQ;SAZM;-37.934167;-57.573333;71;-3;N;America/Cordoba
2509;Presidente Peron;Neuquen;Argentina;NQN;SAZN;-38.949;-68.155711;895;-3;N;America/Cordoba
6850;Mackinac Island Airport;Mackinac Island;United States;MCD;\N;45.8649344;-84.637344;740;-5;U;America/New_York
2511;Comodoro P Zanni;Pehuajo;Argentina;;SAZP;-35.844592;-61.857553;278;-3;N;America/Cordoba
2512;Santa Rosa;Santa Rosa;Argentina;RSA;SAZR;-36.588322;-64.275694;630;-3;N;America/Cordoba
2513;San Carlos De Bariloche;San Carlos De Bariloche;Argentina;BRC;SAZS;-41.151172;-71.157542;2776;-3;N;America/Cordoba
2514;Tandil;Tandil;Argentina;TDL;SAZT;-37.237392;-59.227922;574;-3;N;America/Cordoba
2515;Villa Gesell;Villa Gesell;Argentina;VLG;SAZV;-37.235408;-57.029239;32;-3;N;America/Cordoba
2516;Cutralco;Cutralco;Argentina;;SAZW;-38.939683;-69.264642;2133;-3;N;America/Cordoba
2517;Aviador C Campos;San Martin Des Andes;Argentina;CPC;SAZY;-40.075383;-71.137294;2569;-3;N;America/Cordoba
2518;Conceicao Do Araguaia;Conceicao Do Araguaia;Brazil;CDJ;SBAA;-8.348347;-49.301528;653;-4;S;America/Boa_Vista
2519;Campo Delio Jardim De Mattos;Rio De Janeiro;Brazil;;SBAF;-22.875083;-43.384708;110;-3;S;America/Sao_Paulo
2520;Amapa;Amapa;Brazil;;SBAM;2.077511;-50.858236;45;-3;S;America/Fortaleza
2521;Araraquara;Araracuara;Brazil;AQA;SBAQ;-21.812;-48.133028;2334;-3;S;America/Sao_Paulo
2522;Santa Maria;Aracaju;Brazil;AJU;SBAR;-10.984;-37.070333;23;-3;S;America/Fortaleza
2523;Assis;Assis;Brazil;;SBAS;-22.638564;-50.455914;1850;-3;S;America/Sao_Paulo
2524;Alta Floresta;Alta Floresta;Brazil;AFL;SBAT;-9.866092;-56.106206;947;-4;S;America/Campo_Grande
2525;Aracatuba;Aracatuba;Brazil;ARU;SBAU;-21.141342;-50.424722;1361;-3;S;America/Sao_Paulo
2526;Val De Cans Intl;Belem;Brazil;BEL;SBBE;-1.37925;-48.476292;54;-4;S;America/Boa_Vista
2527;Comandante Gustavo Kraemer;Bage;Brazil;BGX;SBBG;-31.390528;-54.112244;600;-3;S;America/Sao_Paulo
2528;Pampulha Carlos Drummond De Andrade;Belo Horizonte;Brazil;PLU;SBBH;-19.851181;-43.950628;2589;-3;S;America/Sao_Paulo
2529;Bacacheri;Curitiba;Brazil;BFH;SBBI;-25.405078;-49.232036;3057;-3;S;America/Sao_Paulo
2530;Major Brigadeiro Doorgal Borges;Barbacena;Brazil;;SBBQ;-21.267167;-43.761056;3657;-3;S;America/Sao_Paulo
2531;Presidente Juscelino Kubitschek;Brasilia;Brazil;BSB;SBBR;-15.8711;-47.918625;3479;-3;S;America/Sao_Paulo
2532;Bauru;Bauru;Brazil;BAU;SBBU;-22.345042;-49.0538;2025;-3;S;America/Sao_Paulo
2533;Boa Vista;Boa Vista;Brazil;BVB;SBBV;2.846311;-60.690069;276;-4;S;America/Boa_Vista
2534;Barra Do Garcas;Barra Do Garcas;Brazil;;SBBW;-15.861344;-52.388894;1147;-4;S;America/Campo_Grande
2535;Cascavel;Cascavel;Brazil;CAC;SBCA;-25.000339;-53.500764;2473;-3;S;America/Sao_Paulo
2536;Cachimbo;Itaituba;Brazil;;SBCC;-9.333936;-54.965422;1762;-4;S;America/Boa_Vista
2537;Tancredo Neves Intl;Belo Horizonte;Brazil;CNF;SBCF;-19.63375;-43.968856;2715;-3;S;America/Sao_Paulo
2538;Campo Grande;Campo Grande;Brazil;CGR;SBCG;-20.468667;-54.6725;1833;-4;S;America/Campo_Grande
2539;Chapeco;Chapeco;Brazil;XAP;SBCH;-27.134219;-52.656553;2146;-3;S;America/Sao_Paulo
2540;Carolina;Carolina;Brazil;CLN;SBCI;-7.320444;-47.458667;565;-3;S;America/Fortaleza
2541;Forquilhinha;Criciuma;Brazil;CCM;SBCM;-28.725817;-49.424739;93;-3;S;America/Sao_Paulo
2542;Canoas;Porto Alegre;Brazil;;SBCO;-29.945944;-51.144367;26;-3;S;America/Sao_Paulo
2543;Bartolomeu Lisandro;Campos;Brazil;CAW;SBCP;-21.698333;-41.301669;57;-3;S;America/Sao_Paulo
2544;Corumba Intl;Corumba;Brazil;CMG;SBCR;-19.011931;-57.673053;461;-4;S;America/Campo_Grande
2545;Afonso Pena;Curitiba;Brazil;CWB;SBCT;-25.528475;-49.175775;2988;-3;S;America/Sao_Paulo
2546;Caravelas;Caravelas;Brazil;CRQ;SBCV;-17.652283;-39.253069;36;-3;S;America/Fortaleza
2547;Campo Dos Bugres;Caxias Do Sul;Brazil;CXJ;SBCX;-29.197064;-51.187536;2472;-3;S;America/Sao_Paulo
2548;Marechal Rondon;Cuiaba;Brazil;CGB;SBCY;-15.652931;-56.116719;617;-4;S;America/Campo_Grande
2549;Cruzeiro do Sul;Cruzeiro do Sul;Brazil;CZS;SBCZ;-7.599906;-72.769489;637;-5;S;America/Rio_Branco
2550;Presidente Prudente;President Prudente;Brazil;PPB;SBDN;-22.175056;-51.424639;1477;-3;S;America/Sao_Paulo
2551;Eduardo Gomes Intl;Manaus;Brazil;MAO;SBEG;-3.038611;-60.049721;264;-4;S;America/Boa_Vista
2552;Jacareacanga;Jacare-acanga;Brazil;;SBEK;-6.233156;-57.776869;324;-4;S;America/Boa_Vista
2553;Sao Pedro Da Aldeia;Sao Pedro Da Aldeia;Brazil;;SBES;-22.812872;-42.092633;61;-3;S;America/Sao_Paulo
2554;Cataratas Intl;Foz Do Iguacu;Brazil;IGU;SBFI;-25.59615;-54.487206;787;-3;S;America/Sao_Paulo
2555;Hercilio Luz;Florianopolis;Brazil;FLN;SBFL;-27.670489;-48.547181;20;-3;S;America/Sao_Paulo
2556;Fernando De Noronha;Fernando Do Noronha;Brazil;FEN;SBFN;-3.854928;-32.423336;193;-3;S;America/Fortaleza
2557;Fronteira;Fronteira;Brazil;;SBFT;-20.278483;-49.187472;1599;-3;S;America/Sao_Paulo
2558;Furnas;Alpinopolis;Brazil;;SBFU;-20.702817;-46.335264;2413;-3;S;America/Sao_Paulo
2559;Pinto Martins Intl;Fortaleza;Brazil;FOR;SBFZ;-3.776283;-38.532556;82;-3;S;America/Fortaleza
2560;Galeao Antonio Carlos Jobim;Rio De Janeiro;Brazil;GIG;SBGL;-22.808903;-43.243647;28;-3;S;America/Sao_Paulo
2561;Guajara Mirim;Guajara-mirim;Brazil;;SBGM;-10.786375;-65.284792;478;-4;S;America/Boa_Vista
2562;Santa Genoveva;Goiania;Brazil;GYN;SBGO;-16.632033;-49.220686;2450;-3;S;America/Sao_Paulo
2563;Embraer Unidade Gaviao Peixoto;Macae;Brazil;;SBGP;-21.773683;-48.405078;1998;-3;S;America/Sao_Paulo
2564;Guarulhos Gov Andre Franco Montouro;Sao Paulo;Brazil;GRU;SBGR;-23.432075;-46.469511;2459;-3;S;America/Sao_Paulo
2565;Guaratingueta;Guaratingueta;Brazil;;SBGW;-22.791608;-45.204778;1761;-3;S;America/Sao_Paulo
2566;Altamira;Altamira;Brazil;ATM;SBHT;-3.253906;-52.253978;368;-4;S;America/Boa_Vista
2567;Itacoatiara;Itaituba;Brazil;;SBIC;-3.127256;-58.481186;142;-4;S;America/Boa_Vista
2568;Itaituba;Itaituba;Brazil;;SBIH;-4.242342;-56.000669;110;-4;S;America/Boa_Vista
2569;Ilheus;Ilheus;Brazil;IOS;SBIL;-14.815964;-39.033197;15;-3;S;America/Fortaleza
2570;Usiminas;Ipatinga;Brazil;IPN;SBIP;-19.470722;-42.487583;784;-3;S;America/Sao_Paulo
2571;Hidroeletrica;Itumbiara;Brazil;;SBIT;-18.444661;-49.213361;1630;-3;S;America/Sao_Paulo
2572;Prefeito Renato Moreira;Imperatriz;Brazil;IMP;SBIZ;-5.531292;-47.46005;431;-3;S;America/Fortaleza
2573;Julio Cesar;Belem;Brazil;;SBJC;-1.414158;-48.460739;52;-4;S;America/Boa_Vista
2574;Francisco De Assis;Juiz De Fora;Brazil;JDF;SBJF;-21.7915;-43.386778;2989;-3;S;America/Sao_Paulo
2575;Presidente Castro Pinto;Joao Pessoa;Brazil;JPA;SBJP;-7.148381;-34.950681;215;-3;S;America/Fortaleza
2576;Lauro Carneiro De Loyola;Joinville;Brazil;JOI;SBJV;-26.224453;-48.797364;15;-3;S;America/Sao_Paulo
2577;Presidente Joao Suassuna;Campina Grande;Brazil;CPV;SBKG;-7.269917;-35.896364;1646;-3;S;America/Fortaleza
2578;Viracopos;Campinas;Brazil;VCP;SBKP;-23.0075;-47.134444;2170;-3;S;America/Sao_Paulo
2579;Lages;Lajes;Brazil;;SBLJ;-27.782142;-50.281486;3065;-3;S;America/Sao_Paulo
2580;Lins;Lins;Brazil;LIP;SBLN;-21.664039;-49.730519;1575;-3;S;America/Sao_Paulo
2581;Londrina;Londrina;Brazil;LDB;SBLO;-23.333625;-51.130072;1867;-3;S;America/Sao_Paulo
2582;Bom Jesus Da Lapa;Bom Jesus Da Lapa;Brazil;LAZ;SBLP;-13.262086;-43.408114;1454;-3;S;America/Fortaleza
2583;Lagoa Santa;Lagoa Santa;Brazil;;SBLS;-19.661611;-43.896403;2795;-3;S;America/Sao_Paulo
2584;Maraba;Maraba;Brazil;MAB;SBMA;-5.368589;-49.138025;357;-4;S;America/Boa_Vista
2585;Monte Dourado;Almeirim;Brazil;;SBMD;-0.889839;-52.60225;677;-4;S;America/Boa_Vista
2586;Regional De Maringa Silvio Name Junior;Maringa;Brazil;MGF;SBMG;-23.476392;-52.016406;1788;-3;S;America/Sao_Paulo
2587;Mario Ribeiro;Montes Claros;Brazil;MOC;SBMK;-16.706925;-43.8189;2191;-3;S;America/Sao_Paulo
6849;Grand Marais Cook County Airport;Grand Marais;United States;GRM;KCKC;47.8383333;-90.3829444;1799;-6;U;America/Chicago
2589;Ponta Pelada;Manaus;Brazil;;SBMN;-3.146042;-59.9863;267;-4;S;America/Boa_Vista
2590;Zumbi Dos Palmares;Maceio;Brazil;MCZ;SBMO;-9.510808;-35.791678;387;-3;S;America/Fortaleza
2591;Macapa;Macapa;Brazil;MCP;SBMQ;0.050664;-51.072178;56;-3;S;America/Fortaleza
2592;Dix Sept Rosado;Mocord;Brazil;;SBMS;-5.201919;-37.364347;77;-3;S;America/Fortaleza
2593;Marte;Sao Paulo;Brazil;;SBMT;-23.509119;-46.637753;2368;-3;S;America/Sao_Paulo
2594;Manicore;Manicore;Brazil;MNX;SBMY;-5.811381;-61.278319;174;-4;S;America/Boa_Vista
2595;Ministro Victor Konder Intl;Navegantes;Brazil;NVT;SBNF;-26.879999;-48.65139;18;-3;S;America/Sao_Paulo
2596;Santo Angelo;Santo Angelo;Brazil;GEL;SBNM;-28.281683;-54.169139;1056;-3;S;America/Sao_Paulo
2597;Augusto Severo;Natal;Brazil;NAT;SBNT;-5.911417;-35.247717;169;-3;S;America/Fortaleza
2598;Oiapoque;Oioiapoque;Brazil;;SBOI;3.855486;-51.796867;63;-3;S;America/Fortaleza
2599;Salgado Filho;Porto Alegre;Brazil;POA;SBPA;-29.994428;-51.171428;11;-3;S;America/Sao_Paulo
2600;Prefeito Doutor Joao Silva Filho;Parnaiba;Brazil;;SBPB;-2.893747;-41.731961;16;-3;S;America/Fortaleza
2601;Pocos De Caldas;Pocos De Caldas;Brazil;POO;SBPC;-21.843014;-46.567917;4135;-3;S;America/Sao_Paulo
2602;Lauro Kurtz;Passo Fundo;Brazil;PFB;SBPF;-28.243989;-52.326558;2376;-3;S;America/Sao_Paulo
2603;Pelotas;Pelotas;Brazil;PET;SBPK;-31.718353;-52.327689;59;-3;S;America/Sao_Paulo
2604;Senador Nilo Coelho;Petrolina;Brazil;PNZ;SBPL;-9.362411;-40.569097;1263;-3;S;America/Fortaleza
2605;Porto Nacional;Porto Nacional;Brazil;PNB;SBPN;-10.719417;-48.399736;870;-3;S;America/Fortaleza
2606;Ponta Pora;Ponta Pora;Brazil;PMG;SBPP;-22.549639;-55.702614;2156;-4;S;America/Campo_Grande
2607;Governador Jorge Teixeira De Oliveira;Porto Velho;Brazil;PVH;SBPV;-8.709294;-63.902281;294;-4;S;America/Boa_Vista
6848;Porter County Municipal Airport;Valparaiso;United States;NPZ;\N;41.4539722;-87.0070833;770;-6;A;America/Chicago
2609;PlÃ¡cido de Castro;Rio Branco;Brazil;RBR;SBRB;-9.583;-67.4836;633;-5;S;America/Rio_Branco
2610;Guararapes Gilberto Freyre Intl;Recife;Brazil;REC;SBRF;-8.126794;-34.923039;33;-3;S;America/Fortaleza
2611;Rio Grande;Rio Grande;Brazil;RIG;SBRG;-32.082617;-52.166542;27;-3;S;America/Sao_Paulo
2612;Santos Dumont;Rio De Janeiro;Brazil;SDU;SBRJ;-22.910461;-43.163133;11;-3;S;America/Sao_Paulo
2613;Leite Lopes;Ribeirao Preto;Brazil;RAO;SBRP;-21.134167;-47.774189;1802;-3;S;America/Sao_Paulo
2614;Santa Cruz;Rio De Janeiro;Brazil;STU;SBSC;-22.93235;-43.719092;10;-3;S;America/Sao_Paulo
2615;Professor Urbano Ernesto Stumpf;Sao Jose Dos Campos;Brazil;SJK;SBSJ;-23.228172;-45.862739;2120;-3;S;America/Sao_Paulo
2616;Marechal Cunha Machado Intl;Sao Luis;Brazil;SLZ;SBSL;-2.585361;-44.234139;178;-3;S;America/Fortaleza
2618;Congonhas;Sao Paulo;Brazil;CGH;SBSP;-23.626692;-46.655375;2631;-3;S;America/Sao_Paulo
2619;Sao Jose Do Rio Preto;Sao Jose Do Rio Preto;Brazil;SJP;SBSR;-20.816567;-49.406511;1784;-3;S;America/Sao_Paulo
2620;Base Aerea De Santos;Santos;Brazil;SSZ;SBST;-23.925206;-46.2875;10;-3;S;America/Sao_Paulo
2621;Deputado Luis Eduardo Magalhaes;Salvador;Brazil;SSA;SBSV;-12.910994;-38.331044;64;-3;S;America/Fortaleza
2622;Trombetas;Oriximina;Brazil;TMT;SBTB;-1.4896;-56.396803;287;-4;S;America/Boa_Vista
2623;Senador Petronio Portella;Teresina;Brazil;THE;SBTE;-5.059942;-42.823478;219;-3;S;America/Fortaleza
2624;Tefe;Tefe;Brazil;TFF;SBTF;-3.382944;-64.724056;184;-4;S;America/Boa_Vista
2625;Tarauaca;Tarauaca;Brazil;;SBTK;-8.155256;-70.783269;646;-5;S;America/Rio_Branco
2626;Telemaco Borba;Telemaco Borba;Brazil;;SBTL;-24.317775;-50.651592;2610;-3;S;America/Sao_Paulo
2627;Tirios;Obidos Tirios;Brazil;;SBTS;2.223472;-55.946056;1127;-4;S;America/Boa_Vista
2628;Tabatinga;Tabatinga;Brazil;TBT;SBTT;-4.255669;-69.935828;279;-4;S;America/Boa_Vista
2629;Tucurui;Tucurui;Brazil;TUR;SBTU;-3.786008;-49.720267;830;-4;S;America/Boa_Vista
2630;Sao Gabriel Da Cachoeira;Sao Gabriel;Brazil;;SBUA;-0.148419;-66.985589;249;-4;S;America/Boa_Vista
2631;Paulo Afonso;Paulo Alfonso;Brazil;PAV;SBUF;-9.400878;-38.250575;883;-3;S;America/Fortaleza
2632;Rubem Berta;Uruguaiana;Brazil;URG;SBUG;-29.782178;-57.038189;256;-3;S;America/Sao_Paulo
2633;Ten Cel Av Cesar Bombonato;Uberlandia;Brazil;UDI;SBUL;-18.882844;-48.225594;3094;-3;S;America/Sao_Paulo
2634;Urubupunga;Castilho;Brazil;;SBUP;-20.777067;-51.564761;1169;-3;S;America/Sao_Paulo
2635;Uberaba;Uberaba;Brazil;UBA;SBUR;-19.765;-47.964778;2655;-3;S;America/Sao_Paulo
2636;Major Brigadeiro Trompowsky;Varginha;Brazil;VAG;SBVG;-21.590067;-45.473342;3028;-3;S;America/Sao_Paulo
2637;Vilhena;Vilhena;Brazil;BVH;SBVH;-12.694375;-60.098269;2018;-4;S;America/Boa_Vista
2638;Goiabeiras;Vitoria;Brazil;VIX;SBVT;-20.258056;-40.286389;11;-3;S;America/Sao_Paulo
2639;Iauarete;Iauarete;Brazil;;SBYA;0.6075;-69.185837;345;-4;S;America/Boa_Vista
2640;Campo Fontenelle;Piracununga;Brazil;QPS;SBYS;-21.984561;-47.334764;1968;-3;S;America/Sao_Paulo
2641;Chacalluta;Arica;Chile;ARI;SCAR;-18.348531;-70.338742;167;-4;S;America/Santiago
2642;Balmaceda;Balmaceda;Chile;BBA;SCBA;-45.916058;-71.689475;1722;-4;S;America/Santiago
2643;El Bosque;Santiago;Chile;;SCBQ;-33.5618;-70.6884;1844;-4;S;America/Santiago
2644;Chile Chico;Chile Chico;Chile;CCH;SCCC;-46.583341;-71.687405;1070;-4;S;America/Santiago
2645;El Loa;Calama;Chile;CJC;SCCF;-22.498175;-68.903575;7543;-4;S;America/Santiago
2646;General Bernardo O Higgins;Chillan;Chile;;SCCH;-36.582497;-72.031367;495;-4;S;America/Santiago
2647;Carlos Ibanez Del Campo Intl;Punta Arenas;Chile;PUQ;SCCI;-53.002642;-70.854586;139;-4;S;America/Santiago
2648;Teniente Vidal;Coyhaique;Chile;GXQ;SCCY;-45.594211;-72.106133;1020;-4;S;America/Santiago
2649;Diego Aracena Intl;Iquique;Chile;IQQ;SCDA;-20.535222;-70.181275;155;-4;S;America/Santiago
2650;Arturo Merino Benitez Intl;Santiago;Chile;SCL;SCEL;-33.392975;-70.785803;1555;-4;S;America/Santiago
2651;Cerro Moreno Intl;Antofagasta;Chile;ANF;SCFA;-23.444478;-70.4451;455;-4;S;America/Santiago
2652;Capitan Fuentes Martinez;Porvenir;Chile;;SCFM;-53.2537;-70.319228;104;-4;S;America/Santiago
2653;Futaleufu;Futaleufu;Chile;;SCFT;-43.189167;-71.851112;1148;-4;S;America/Santiago
2654;Maria Dolores;Los Angeles;Chile;LSQ;SCGE;-37.401731;-72.425444;374;-4;S;America/Santiago
2655;Guardiamarina Zanartu;Puerto Williams;Chile;;SCGZ;-54.931072;-67.626261;88;-4;S;America/Santiago
2656;Carriel Sur Intl;Concepcion;Chile;CCP;SCIE;-36.77265;-73.063106;26;-4;S;America/Santiago
2657;Mataveri Intl;Easter Island;Chile;IPC;SCIP;-27.164792;-109.421831;227;-6;S;Pacific/Easter
2658;Canal Bajo Carlos Hott Siebert;Osorno;Chile;ZOS;SCJO;-40.611208;-73.061042;187;-4;S;America/Santiago
2659;Vallenar;Vallenar;Chile;;SCLL;-28.596403;-70.755997;1725;-4;S;America/Santiago
2660;De La Independencia;Rancagua;Chile;;SCRG;-34.173694;-70.775694;1446;-4;S;America/Santiago
2661;Teniente Rodolfo Marsh Martin;Isla Rey Jorge;Antarctica;;SCRM;-62.190833;-58.986667;147;12;U;Antarctica/South_Pole
2662;La Florida;La Serena;Chile;LSC;SCSE;-29.916233;-71.199522;481;-4;S;America/Santiago
2663;Eulogio Sanchez;Santiago;Chile;;SCTB;-33.456278;-70.546667;2129;-4;S;America/Santiago
2664;Maquehue;Temuco;Chile;ZCO;SCTC;-38.766819;-72.637097;304;-4;S;America/Santiago
2665;El Tepual Intl;Puerto Montt;Chile;PMC;SCTE;-41.438886;-73.093953;294;-4;S;America/Santiago
2666;Chaiten;Chaiten;Chile;WCH;SCTN;-42.932808;-72.699114;13;-4;S;America/Santiago
2667;Pichoy;Valdivia;Chile;ZAL;SCVD;-39.649956;-73.086111;59;-4;S;America/Santiago
2668;Chachoan;Ambato;Ecuador;ATF;SEAM;-1.212067;-78.574636;8502;-5;U;America/Guayaquil
2669;Hacienda Clementina;Clementia;Ecuador;;SECM;-1.706275;-79.378936;328;-5;U;America/Guayaquil
2670;Francisco De Orellana;Coca;Ecuador;OCC;SECO;-0.462886;-76.986842;834;-5;U;America/Guayaquil
2671;Mariscal Lamar;Cuenca;Ecuador;CUE;SECU;-2.889467;-78.984397;8306;-5;U;America/Guayaquil
2672;Seymour;Galapagos;Ecuador;GPS;SEGS;-0.453758;-90.265914;207;-6;U;Pacific/Galapagos
2673;Jose Joaquin De Olmedo Intl;Guayaquil;Ecuador;GYE;SEGU;-2.157419;-79.883558;19;-5;U;America/Guayaquil
2674;Gualaquiza;Gualaquiza;Ecuador;;SEGZ;-3.423214;-78.566994;2640;-5;U;America/Guayaquil
2675;Atahualpa;Ibarra;Ecuador;;SEIB;0.338419;-78.136422;7304;-5;U;America/Guayaquil
2676;Km 192;Km-192;Ecuador;;SEKK;0.184203;-79.391956;1247;-5;U;America/Guayaquil
2677;Hacienda La Julia;La Julia;Ecuador;;SELJ;-1.704381;-79.552261;50;-5;U;America/Guayaquil
2678;Cotopaxi Intl;Latacunga;Ecuador;;SELT;-0.906833;-78.615756;9205;-5;U;America/Guayaquil
2679;Jose Maria Velasco Ibarra;Macara;Ecuador;;SEMA;-4.378231;-79.941022;1508;-5;U;America/Guayaquil
2680;Coronel E Carvajal;Macas;Ecuador;XMS;SEMC;-2.299167;-78.12075;3452;-5;U;America/Guayaquil
2681;General Manuel Serrano;Machala;Ecuador;MCH;SEMH;-3.268903;-79.961572;11;-5;U;America/Guayaquil
2682;Montalvo;Montalvo;Ecuador;;SEMO;-2.067014;-76.975742;960;-5;U;America/Guayaquil
2683;Eloy Alfaro Intl;Manta;Ecuador;MEC;SEMT;-0.946078;-80.678808;48;-5;U;America/Guayaquil
2684;Maragrosa;Maragrosa;Ecuador;;SEMX;-2.851097;-79.803619;18;-5;U;America/Guayaquil
2685;Amable Calle Gutierrez;Pasaje;Ecuador;;SEPS;-3.319667;-79.769165;22;-5;U;America/Guayaquil
2686;Reales Tamarindos;Portoviejo;Ecuador;PVO;SEPV;-1.041647;-80.472206;130;-5;U;America/Guayaquil
2687;Quevedo;Quevedo;Ecuador;;SEQE;-0.9894;-79.465114;350;-5;U;America/Guayaquil
2688;Mariscal Sucre Intl;Quito;Ecuador;UIO;SEQU;-0.141144;-78.488214;9228;-5;U;America/Guayaquil
2689;Chimborazo;Riobamba;Ecuador;;SERB;-1.653433;-78.656142;9151;-5;U;America/Guayaquil
2690;Coronel Artilleria Victor Larrea;Santa Rosa;Ecuador;;SERO;-3.435161;-79.977817;170;-5;U;America/Guayaquil
2691;General Ulpiano Paez;Salinas;Ecuador;SNC;SESA;-2.204994;-80.988878;18;-5;U;America/Guayaquil
2692;Santo Domingo Los Colorados;Santo Domingo;Ecuador;;SESD;-0.248222;-79.214447;1714;-5;U;America/Guayaquil
2694;Taura;Taura;Ecuador;;SETA;-2.261036;-79.680169;56;-5;U;America/Guayaquil
2695;Mayor Galo Torres;Tena;Ecuador;;SETE;-0.986767;-77.819447;1708;-5;U;America/Guayaquil
2696;Tarapoa;Tarapoa;Ecuador;TPC;SETR;-0.122956;-76.33775;814;-5;U;America/Guayaquil
2697;Teniente Coronel Luis A Mantilla;Tulcan;Ecuador;TUA;SETU;0.809506;-77.708056;9649;-5;U;America/Guayaquil
6847;Silverton;Silverton;United States;;\N;37.812545;-107.662994;9308;-7;U;America/Denver
2699;Silvio Pettirossi Intl;Asuncion;Paraguay;ASU;SGAS;-25.23985;-57.519133;292;-4;S;America/Asuncion
2700;Juan De Ayolas;Ayolas;Paraguay;;SGAY;-27.37065;-56.853944;223;-4;S;America/Asuncion
2701;Teniente Col Carmelo Peralta;Conception;Paraguay;;SGCO;-23.44175;-57.427122;253;-4;S;America/Asuncion
2702;Itaipu;Itaipu;Paraguay;;SGIB;-25.407853;-54.619417;762;-4;S;America/Asuncion
2703;Dr Luis Maria Argana Intl;Mariscal Estigarribia;Paraguay;;SGME;-22.044986;-60.621694;553;-4;S;America/Asuncion
2704;Carlos Miguel Gimenez;Pilar;Paraguay;;SGPI;-26.881467;-58.318036;249;-4;S;America/Asuncion
2705;El Eden;Armenia;Colombia;AXM;SKAR;4.452775;-75.766447;3990;-5;U;America/Bogota
2706;Tres De Mayo;Puerto Asis;Colombia;PUU;SKAS;0.505228;-76.500836;815;-5;U;America/Bogota
2707;Las Flores;El Banco;Colombia;;SKBC;9.045542;-73.974931;111;-5;U;America/Bogota
2708;Palonegro;Bucaramanga;Colombia;BGA;SKBG;7.1265;-73.184778;3897;-5;U;America/Bogota
2709;Eldorado Intl;Bogota;Colombia;BOG;SKBO;4.701594;-74.146947;8361;-5;U;America/Bogota
2710;Ernesto Cortissoz;Barranquilla;Colombia;BAQ;SKBQ;10.889589;-74.780819;98;-5;U;America/Bogota
2711;Jose Celestino Mutis;Bahia Solano;Colombia;BSC;SKBS;6.202917;-77.394675;80;-5;U;America/Bogota
2712;Gerardo Tobar Lopez;Buenaventura;Colombia;BUN;SKBU;3.819628;-76.989767;48;-5;U;America/Bogota
2713;Camilo Daza;Cucuta;Colombia;CUC;SKCC;7.927567;-72.511547;1096;-5;U;America/Bogota
2714;Rafael Nunez;Cartagena;Colombia;CTG;SKCG;10.442381;-75.512961;4;-5;U;America/Bogota
2715;Alfonso Bonilla Aragon Intl;Cali;Colombia;CLO;SKCL;3.543222;-76.381583;3162;-5;U;America/Bogota
2716;La Florida;Tumaco;Colombia;TCO;SKCO;1.814417;-78.749228;8;-5;U;America/Bogota
2717;Las Brujas;Corozal;Colombia;CZU;SKCZ;9.332742;-75.285594;528;-5;U;America/Bogota
2718;Yariguies;Barrancabermeja;Colombia;EJA;SKEJ;7.024331;-73.8068;412;-5;U;America/Bogota
2719;Gustavo Artunduaga Paredes;Florencia;Colombia;FLA;SKFL;1.589189;-75.564372;803;-5;U;America/Bogota
2720;Santiago Vila;Girardot;Colombia;;SKGI;4.276242;-74.796692;900;-5;U;America/Bogota
6846;McCarthy Airport;McCarthy;United States;MXY;\N;61.4370608;-142.90307372;1531;-9;U;America/Anchorage
2722;Juan Casiano;Guapi;Colombia;GPI;SKGP;2.570133;-77.8986;164;-5;U;America/Bogota
2723;Guaymaral;Guaymaral;Colombia;;SKGY;4.812333;-74.064919;8390;-5;U;America/Bogota
2724;Perales;Ibague;Colombia;IBE;SKIB;4.421608;-75.1333;2999;-5;U;America/Bogota
2725;San Luis;Ipiales;Colombia;IPI;SKIP;0.861925;-77.671764;9765;-5;U;America/Bogota
2726;Antonio Roldan Betancourt;Carepa;Colombia;;SKLC;7.811956;-76.716428;46;-5;U;America/Bogota
2727;La Mina;La Mina;Colombia;;SKLM;11.232528;-72.490139;276;-5;U;America/Bogota
2728;Alfredo Vasquez Cobo;Leticia;Colombia;LET;SKLT;-4.193549;-69.943163;277;-5;U;America/Bogota
2729;Olaya Herrera;Medellin;Colombia;EOH;SKMD;6.219958;-75.590519;4940;-5;U;America/Bogota
2730;Baracoa;Magangue;Colombia;MGN;SKMG;9.284739;-74.846092;178;-5;U;America/Bogota
2731;Los Garzones;Monteria;Colombia;MTR;SKMR;8.823744;-75.825831;36;-5;U;America/Bogota
2732;Fabio Alberto Leon Bentley;Mitu;Colombia;MVP;SKMU;1.253664;-70.233878;680;-5;U;America/Bogota
2733;La Nubia;Manizales;Colombia;MZL;SKMZ;5.029597;-75.464708;6871;-5;U;America/Bogota
2734;Benito Salas;Neiva;Colombia;NVA;SKNV;2.95015;-75.294;1464;-5;U;America/Bogota
2735;Aguas Claras;Ocana;Colombia;OCV;SKOC;8.315061;-73.358331;3850;-5;U;America/Bogota
2736;Otu;Otu;Colombia;OTU;SKOT;7.010369;-74.715497;2060;-5;U;America/Bogota
2737;Puerto Bolivar;Puerto Bolivar;Colombia;;SKPB;12.221483;-71.984817;90;-5;U;America/Bogota
2738;Puerto Carreno;Puerto Carreno;Colombia;PCR;SKPC;6.184717;-67.493164;177;-5;U;America/Bogota
2739;Matecana;Pereira;Colombia;PEI;SKPE;4.812675;-75.739519;4416;-5;U;America/Bogota
2740;Pitalito;Pitalito;Colombia;;SKPI;1.857769;-76.085719;4212;-5;U;America/Bogota
2741;Guillermo Leon Valencia;Popayan;Colombia;PPN;SKPP;2.4544;-76.609319;5687;-5;U;America/Bogota
2742;Antonio Narino;Pasto;Colombia;PSO;SKPS;1.396247;-77.291478;5951;-5;U;America/Bogota
2743;El Embrujo;Providencia;Colombia;PVA;SKPV;13.356944;-81.35833;10;-5;U;America/Bogota
2744;Mariquita;Mariquita;Colombia;;SKQU;5.212556;-74.883647;1531;-5;U;America/Bogota
2745;Jose Maria Cordova;Rio Negro;Colombia;MDE;SKRG;6.164536;-75.423119;6955;-5;U;America/Bogota
2746;Almirante Padilla;Rio Hacha;Colombia;RCH;SKRH;11.526222;-72.925958;43;-5;U;America/Bogota
2747;Jorge E Gonzalez Torres;San Jose Del Guaviare;Colombia;SJE;SKSJ;2.579694;-72.639358;605;-5;U;America/Bogota
2748;Simon Bolivar;Santa Marta;Colombia;SMR;SKSM;11.11965;-74.230647;22;-5;U;America/Bogota
2749;Gustavo Rojas Pinilla;San Andres Island;Colombia;ADZ;SKSP;12.583594;-81.711192;19;-5;U;America/Bogota
2750;Eduardo Falla Solano;San Vincente De Caguan;Colombia;SVI;SKSV;2.152175;-74.76635;920;-5;U;America/Bogota
2751;Tame;Tame;Colombia;TME;SKTM;6.451081;-71.760261;1050;-5;U;America/Bogota
2752;Santiago Perez;Arauca;Colombia;AUC;SKUC;7.068881;-70.736925;420;-5;U;America/Bogota
2753;El Carano;Quibdo;Colombia;UIB;SKUI;5.690758;-76.641181;204;-5;U;America/Bogota
2754;Farfan;Tulua;Colombia;ULQ;SKUL;4.088358;-76.235133;3132;-5;U;America/Bogota
2755;Alfonso Lopez Pumarejo;Valledupar;Colombia;VUP;SKVP;10.435042;-73.249506;456;-5;U;America/Bogota
2756;Vanguardia;Villavicencio;Colombia;VVC;SKVV;4.167875;-73.613761;1394;-5;U;America/Bogota
2758;Bermejo;Bermejo;Bolivia;BJO;SLBJ;-22.773336;-64.312881;1250;-4;U;America/La_Paz
2759;Jorge Wilsterman;Cochabamba;Bolivia;CBB;SLCB;-17.421058;-66.177114;8360;-4;U;America/La_Paz
2760;Chimore;Chapacura;Bolivia;;SLCH;-16.990019;-65.141533;1000;-4;U;America/La_Paz
2761;Heroes Del Acre;Cobija;Bolivia;CIJ;SLCO;-11.040436;-68.782972;892;-4;U;America/La_Paz
2762;El Alto Intl;La Paz;Bolivia;LPB;SLLP;-16.513339;-68.192256;13325;-4;U;America/La_Paz
2763;Juan Mendoza;Oruro;Bolivia;;SLOR;-17.962589;-67.076236;12146;-4;U;America/La_Paz
2764;Capitan Nicolas Rojas;Potosi;Bolivia;POI;SLPO;-19.543069;-65.723706;12913;-4;U;America/La_Paz
2765;Tte De Av Salvador Ogaya G;Puerto Suarez;Bolivia;PSZ;SLPS;-18.975281;-57.820586;439;-4;U;America/La_Paz
2766;Santa Ana Del Yacuma;Santa Ana;Bolivia;;SLSA;-13.762208;-65.435158;472;-4;U;America/La_Paz
2767;Juana Azurduy De Padilla;Sucre;Bolivia;SRE;SLSU;-19.007083;-65.288747;9527;-4;U;America/La_Paz
2768;Capitan Oriel Lea Plaza;Tarija;Bolivia;TJA;SLTJ;-21.555736;-64.701325;6084;-4;U;America/La_Paz
2769;Tte Av Jorge Henrich Arauz;Trinidad;Bolivia;TDD;SLTR;-14.818739;-64.918019;509;-4;U;America/La_Paz
2770;Tcnl Rafael Pabon;Villa Montes;Bolivia;;SLVM;-21.255231;-63.405611;1306;-4;U;America/La_Paz
2771;Viru Viru Intl;Santa Cruz;Bolivia;VVI;SLVR;-17.644756;-63.135364;1225;-4;U;America/La_Paz
2772;Yacuiba;Yacuiba;Bolivia;BYC;SLYA;-21.960925;-63.651669;2116;-4;U;America/La_Paz
2773;Johan A Pengel Intl;Zandery;Suriname;PBM;SMJP;5.452831;-55.187783;59;-3;U;America/Paramaribo
2774;Rochambeau;Cayenne;French Guiana;CAY;SOCA;4.819808;-52.360447;26;-3;U;America/Cayenne
2775;St Georges De L Oyapock;St.-georges Oyapock;French Guiana;;SOOG;3.8976;-51.804083;46;-3;U;America/Cayenne
2776;Huancabamba;Huancabamba;Peru;;SPAB;-5.256767;-79.442856;6312;-5;U;America/Lima
2777;Alferez Vladimir Sara Bauer;Andoas;Peru;;SPAS;-2.796128;-76.466617;728;-5;U;America/Lima
2778;Atalaya;Atalaya;Peru;;SPAY;-10.729117;-73.766503;751;-5;U;America/Lima
6845;Seward Airport;Seward;United States;SWD;\N;60.1269383;-149.4188122;22;-9;U;America/Anchorage
2780;Iberia;Iberia;Peru;;SPBR;-11.411578;-69.488711;750;-5;U;America/Lima
2781;Cap Fap David Abenzur Rengifo Intl;Pucallpa;Peru;PCL;SPCL;-8.377939;-74.574297;513;-5;U;America/Lima
2782;Teniente Jaime A De Montreuil Morales;Chimbote;Peru;CHM;SPEO;-9.149614;-78.52385;69;-5;U;America/Lima
2783;Puerto Esperanza;Puerto Esperanza;Peru;;SPEP;-9.768131;-70.706456;725;-5;U;America/Lima
2784;Cesar Torke Podesta;Moquegua;Peru;;SPEQ;-17.178961;-70.930803;4480;-5;U;America/Lima
2785;Capt Jose A Quinones Gonzales Intl;Chiclayo;Peru;CIX;SPHI;-6.787475;-79.828097;97;-5;U;America/Lima
2786;Coronel Fap Alfredo Mendivil Duarte;Ayacucho;Peru;AYP;SPHO;-13.154819;-74.204417;8917;-5;U;America/Lima
2787;Andahuaylas;Andahuaylas;Peru;ANS;SPHY;-13.706408;-73.350378;11300;-5;U;America/Lima
2788;Comandante Fap German Arias Graziani;Anta;Peru;ATA;SPHZ;-9.347444;-77.598392;9097;-5;U;America/Lima
2789;Jorge Chavez Intl;Lima;Peru;LIM;SPIM;-12.021889;-77.114319;113;-5;U;America/Lima
2790;Juanjui;Juanjui;Peru;JJI;SPJI;-7.1691;-76.728561;1148;-5;U;America/Lima
2791;Francisco Carle;Jauja;Peru;;SPJJ;-11.783144;-75.473394;11034;-5;U;America/Lima
2792;Juliaca;Juliaca;Peru;JUL;SPJL;-15.467103;-70.158169;12552;-5;U;America/Lima
6844;Michigan City Municipal Airport;Michigan City;United States;MGC;KMGC;41.7033;-86.8211;500;-6;A;America/Chicago
2794;Ilo;Ilo;Peru;;SPLO;-17.695036;-71.343964;72;-5;U;America/Lima
2795;Las Palmas;Las Palmas;Peru;;SPLP;-12.160708;-76.998942;250;-5;U;America/Lima
2796;Pedro Canga;Tumbes;Peru;TBP;SPME;-3.552528;-80.381356;115;-5;U;America/Lima
2797;Moises Benzaquen Rengifo;Yurimaguas;Peru;YMS;SPMS;-5.893772;-76.118211;587;-5;U;America/Lima
2799;Collique;Collique;Peru;;SPOL;-11.9287;-77.061139;410;-5;U;America/Lima
2800;Chachapoyas;Chachapoyas;Peru;CHH;SPPY;-6.201806;-77.856064;8333;-5;U;America/Lima
2801;Coronel Francisco Secada Vignetta Intl;Iquitos;Peru;IQT;SPQT;-3.784739;-73.308806;306;-5;U;America/Lima
2802;Rodriguez Ballon;Arequipa;Peru;AQP;SPQU;-16.341072;-71.583083;8405;-5;U;America/Lima
2803;San Ramon;San Ramon;Peru;;SPRM;-11.128639;-75.3505;2600;-5;U;America/Lima
2804;Capitan Carlos Martinez De Pinillos;Trujillo;Peru;TRU;SPRU;-8.081411;-79.108761;106;-5;U;America/Lima
2805;Pisco Intl;Pisco;Peru;PIO;SPSO;-13.744864;-76.220284;39;-5;U;America/Lima
2806;Cadete Guillermo Del Castillo Paredes;Tarapoto;Peru;TPP;SPST;-6.508742;-76.373247;869;-5;U;America/Lima
2807;Coronel Carlos Ciriani Santa Rosa Intl;Tacna;Peru;TCQ;SPTN;-18.053333;-70.275833;1538;-5;U;America/Lima
2808;Padre Aldamiz;Puerto Maldonado;Peru;PEM;SPTU;-12.613611;-69.228611;659;-5;U;America/Lima
2809;Capitan Fap Guillermo Concha Iberico;Piura;Peru;PIU;SPUR;-5.20575;-80.616444;120;-5;U;America/Lima
2810;Capitan Montes;Talara;Peru;TYL;SPYL;-4.576639;-81.254139;282;-5;U;America/Lima
6843;Niijima Airport;Niijima;Japan;;RJAN;34.366944;139.268611;133;9;N;Asia/Tokyo
2812;Teniente Alejandro Velasco Astete Intl;Cuzco;Peru;CUZ;SPZO;-13.535722;-71.938781;10860;-5;U;America/Lima
2813;Angel S Adami;Montevideo;Uruguay;;SUAA;-34.789208;-56.264703;174;-3;S;America/Montevideo
2814;Santa Bernardina Intl;Durazno;Uruguay;;SUDU;-33.358867;-56.499172;305;-3;S;America/Montevideo
6842;East Glacier Park Amtrak Station;East Glacier;United States;;\N;48.443965;-113.218556;5171;-7;A;America/Denver
2816;Carrasco Intl;Montevideo;Uruguay;MVD;SUMU;-34.838417;-56.030806;105;-3;S;America/Montevideo
2817;Nueva Hesperides Intl;Salto;Uruguay;STY;SUSO;-31.438481;-57.985294;187;-3;S;America/Montevideo
2818;Oswaldo Guevara Mujica;Acarigua;Venezuela;AGV;SVAC;9.553422;-69.237536;741;-4.5;U;America/Caracas
2819;Anaco;Anaco;Venezuela;AAO;SVAN;9.430225;-64.470725;721;-4.5;U;America/Caracas
2820;San Fernando De Atabapo;San Fernando Deatabapo;Venezuela;;SVAT;4.051819;-67.701072;298;-4.5;U;America/Caracas
2821;General Jose Antonio Anzoategui Intl;Barcelona;Venezuela;BLA;SVBC;10.107139;-64.689161;26;-4.5;U;America/Caracas
2822;Barinas;Barinas;Venezuela;BNS;SVBI;8.619575;-70.220825;666;-4.5;U;America/Caracas
2823;El Libertador Ab;Maracaibo;Venezuela;;SVBL;10.183375;-67.557319;1450;-4.5;U;America/Caracas
2824;Barquisimeto Intl;Barquisimeto;Venezuela;BRM;SVBM;10.042747;-69.358619;2042;-4.5;U;America/Caracas
2826;Ciudad Bolivar;Ciudad Bolivar;Venezuela;CBL;SVCB;8.121898;-63.537353;197;-4.5;U;America/Caracas
2827;Caicara Del Orinoco;Caicara De Orinoco;Venezuela;;SVCD;7.626078;-66.164917;141;-4.5;U;America/Caracas
2828;San Carlos;San Carlos;Venezuela;;SVCJ;9.647722;-68.574656;512;-4.5;U;America/Caracas
2829;Calabozo;Calabozo;Venezuela;;SVCL;8.924656;-67.417094;328;-4.5;U;America/Caracas
2830;Canaima;Canaima;Venezuela;CAJ;SVCN;6.231989;-62.854433;1450;-4.5;U;America/Caracas
2831;Carora;Carora;Venezuela;;SVCO;10.175603;-70.065214;1437;-4.5;U;America/Caracas
2832;General Jose Francisco Bermudez;Carupano;Venezuela;CUP;SVCP;10.660014;-63.261681;33;-4.5;U;America/Caracas
2833;Jose Leonardo Chirinos;Coro;Venezuela;CZE;SVCR;11.414867;-69.681656;52;-4.5;U;America/Caracas
2834;Oscar Machado Zuloaga;Caracas;Venezuela;;SVCS;10.286589;-66.816219;2145;-4.5;U;America/Caracas
2835;Antonio Jose De Sucre;Cumana;Venezuela;CUM;SVCU;10.450333;-64.130472;25;-4.5;U;America/Caracas
2836;Capitan Manuel Rios Guarico Airbase;Carrizal;Venezuela;;SVCZ;9.372167;-66.922989;525;-4.5;U;America/Caracas
2837;El Dorado;El Dorado;Venezuela;;SVED;6.715438;-61.639219;318;-4.5;U;America/Caracas
2838;Elorza;Elorza;Venezuela;;SVEZ;7.059722;-69.496694;295;-4.5;U;America/Caracas
2839;Guasdualito;Guasdualito;Venezuela;;SVGD;7.211081;-70.75645;426;-4.5;U;America/Caracas
2840;Guiria;Guiria;Venezuela;GUI;SVGI;10.574078;-62.312667;42;-4.5;U;America/Caracas
2841;Guanare;Guanare;Venezuela;GUQ;SVGU;9.026944;-69.75515;606;-4.5;U;America/Caracas
2842;Higuerote;Higuerote;Venezuela;;SVHG;10.462453;-66.092758;10;-4.5;U;America/Caracas
2843;Andres Miguel Salazar Marcano;Isla De Coche;Venezuela;;SVIE;10.794406;-63.981589;10;-4.5;U;America/Caracas
2844;Josefa Camejo;Paraguana;Venezuela;LSP;SVJC;11.780775;-70.151497;75;-4.5;U;America/Caracas
2845;San Juan De Los Morros;San Juan De Los Morros;Venezuela;;SVJM;9.906953;-67.379639;1404;-4.5;U;America/Caracas
2846;La Fria;La Fria;Venezuela;LFR;SVLF;8.239167;-72.271028;323;-4.5;U;America/Caracas
2847;La Orchila;La Orchila;Venezuela;;SVLO;11.808822;-66.179214;5;-4.5;U;America/Caracas
2848;La Chinita Intl;Maracaibo;Venezuela;MAR;SVMC;10.558208;-71.727856;235;-4.5;U;America/Caracas
2849;Alberto Carnevalli;Merida;Venezuela;MRD;SVMD;8.582294;-71.161186;5007;-4.5;U;America/Caracas
2850;Del Caribe Intl Gen Santiago Marino;Porlamar;Venezuela;PMV;SVMG;10.912926;-63.967581;74;-4.5;U;America/Caracas
2851;Simon Bolivar Intl;Caracas;Venezuela;CCS;SVMI;10.603117;-66.990583;235;-4.5;U;America/Caracas
2852;Maturin;Maturin;Venezuela;MUN;SVMT;9.749067;-63.1534;224;-4.5;U;America/Caracas
2853;Casique Aramare;Puerto Ayacucho;Venezuela;PYH;SVPA;5.619992;-67.606103;245;-4.5;U;America/Caracas
2854;General Bartolome Salom Intl;Puerto Cabello;Venezuela;PBL;SVPC;10.4805;-68.073025;32;-4.5;U;America/Caracas
2855;Paramillo;San Cristobal;Venezuela;;SVPM;7.801317;-72.202847;3314;-4.5;U;America/Caracas
2856;General Manuel Carlos Piar;Guayana;Venezuela;PZO;SVPR;8.288528;-62.760361;472;-4.5;U;America/Caracas
2857;Palmarito;Palmarito;Venezuela;;SVPT;7.575706;-70.174328;347;-4.5;U;America/Caracas
2858;San Antonio Del Tachira;San Antonio;Venezuela;SVZ;SVSA;7.840831;-72.439742;1312;-4.5;U;America/Caracas
2859;Santa Barbara De Barinas;Santa Barbara;Venezuela;;SVSB;7.803514;-71.165717;590;-4.5;U;America/Caracas
2860;Santa Elena De Uairen;Santa Ana De Uairen;Venezuela;;SVSE;4.554722;-61.144922;2939;-4.5;U;America/Caracas
2861;Mayor Buenaventura Vivas;Santo Domingo;Venezuela;STD;SVSO;7.565111;-72.035125;1083;-4.5;U;America/Caracas
2862;Sub Teniente Nestor Arias;San Felipe;Venezuela;SFH;SVSP;10.278728;-68.755211;761;-4.5;U;America/Caracas
2863;San Fernando De Apure;San Fernando De Apure;Venezuela;SFD;SVSR;7.883317;-67.444025;154;-4.5;U;America/Caracas
2864;San Tome;San Tome;Venezuela;SOM;SVST;8.945147;-64.151083;837;-4.5;U;America/Caracas
2865;Santa Barbara Del Zulia;Santa Barbara;Venezuela;STB;SVSZ;8.974425;-71.943014;32;-4.5;U;America/Caracas
2866;Tucupita;Tucupita;Venezuela;TUV;SVTC;9.088994;-62.094175;105;-4.5;U;America/Caracas
2867;Tumeremo;Tumeremo;Venezuela;;SVTM;7.249381;-61.528933;344;-4.5;U;America/Caracas
2868;Arturo Michelena Intl;Valencia;Venezuela;VLN;SVVA;10.149733;-67.9284;1417;-4.5;U;America/Caracas
6841;Flugplatz Hoexter Holzminden;Hoexter Holzminden;Germany;;EDVI;51.4838;9.2259;855;1;E;Europe/Berlin
2870;Dr Antonio Nicolas Briceno;Valera;Venezuela;VLV;SVVL;9.340797;-70.584089;1893;-4.5;U;America/Caracas
2871;Valle De La Pascua;Valle De La Pascua;Venezuela;VDP;SVVP;9.222028;-65.993583;410;-4.5;U;America/Caracas
2872;Linden;Linden;Guyana;;SYLD;5.965922;-58.270336;180;-4;U;America/Guyana
2873;Lethem;Lethem;Guyana;LTM;SYLT;3.372761;-59.789439;351;-4;U;America/Guyana
2874;V C Bird Intl;Antigua;Antigua and Barbuda;ANU;TAPA;17.136749;-61.792667;62;-4;U;America/Antigua
2875;Grantley Adams Intl;Bridgetown;Barbados;BGI;TBPB;13.074603;-59.492456;169;-4;U;America/Barbados
2876;Canefield;Canefield;Dominica;DCF;TDCF;15.336719;-61.392211;13;-4;U;America/Dominica
2877;Melville Hall;Dominica;Dominica;DOM;TDPD;15.547028;-61.3;73;-4;U;America/Dominica
2878;Le Lamentin;Fort-de-france;Martinique;FDF;TFFF;14.591033;-61.003175;16;-4;U;America/Martinique
2879;Grand Case;St. Martin;Guadeloupe;SFG;TFFG;18.099914;-63.047197;7;-4;U;\N
2881;Le Raizet;Pointe-a-pitre;Guadeloupe;PTP;TFFR;16.265306;-61.531806;36;-4;U;America/Guadeloupe
2882;Point Salines Intl;Point Salines;Grenada;GND;TGPY;12.004247;-61.786192;41;-4;U;America/Grenada
2883;Cyril E King;St. Thomas;Virgin Islands;STT;TIST;18.337306;-64.973361;23;-4;U;America/St_Thomas
2884;Henry E Rohlsen;St. Croix Island;Virgin Islands;STX;TISX;17.701889;-64.798556;74;-4;U;America/St_Thomas
2885;Rafael Hernandez;Aguadilla;Puerto Rico;BQN;TJBQ;18.494861;-67.129444;237;-4;U;America/Puerto_Rico
2886;Diego Jimenez Torres;Fajardo;Puerto Rico;FAJ;TJFA;18.308889;-65.661861;64;-4;U;America/Puerto_Rico
2887;Fernando Luis Ribas Dominicci;San Juan;Puerto Rico;SIG;TJIG;18.456828;-66.098139;10;-4;U;America/Puerto_Rico
2888;Eugenio Maria De Hostos;Mayaguez;Puerto Rico;MAZ;TJMZ;18.255694;-67.148472;28;-4;U;America/Puerto_Rico
2889;Mercedita;Ponce;Puerto Rico;PSE;TJPS;18.008306;-66.563028;29;-4;U;America/Puerto_Rico
2890;Luis Munoz Marin Intl;San Juan;Puerto Rico;SJU;TJSJ;18.439417;-66.001833;9;-4;U;America/Puerto_Rico
2891;Robert L Bradshaw;Basse Terre;Saint Kitts and Nevis;SKB;TKPK;17.311194;-62.718667;170;-4;U;America/St_Kitts
6840;Dinslaken Schwarze-Heide;Dinslaken;Germany;;EDLD;51.3659;6.5155;198;1;E;Europe/Berlin
2893;George F L Charles;Castries;Saint Lucia;SLU;TLPC;14.020228;-60.992936;22;-4;U;America/St_Lucia
2894;Hewanorra Intl;Hewandorra;Saint Lucia;UVF;TLPL;13.733194;-60.952597;14;-4;U;America/St_Lucia
2895;Reina Beatrix Intl;Oranjestad;Aruba;AUA;TNCA;12.501389;-70.015221;60;-4;U;America/Aruba
2896;Flamingo;Kralendijk;Netherlands Antilles;BON;TNCB;12.131044;-68.268511;20;-4;U;America/Curacao
2897;Hato;Willemstad;Netherlands Antilles;CUR;TNCC;12.188853;-68.959803;29;-4;U;America/Curacao
2898;F D Roosevelt;Oranjestad;Netherlands Antilles;EUX;TNCE;17.496492;-62.979439;129;-4;U;America/Curacao
2899;Princess Juliana Intl;Philipsburg;Netherlands Antilles;SXM;TNCM;18.040953;-63.1089;14;-4;U;America/Curacao
2900;Wallblake;The Valley;Anguilla;AXA;TQPF;18.204834;-63.055084;127;-4;U;America/Anguilla
2901;Crown Point;Scarborough;Trinidad and Tobago;TAB;TTCP;11.149658;-60.832194;38;-4;U;America/Port_of_Spain
2902;Piarco;Port-of-spain;Trinidad and Tobago;POS;TTPP;10.595369;-61.337242;58;-4;U;America/Port_of_Spain
2903;Terrance B Lettsome Intl;Tortola;British Virgin Islands;EIS;TUPJ;18.444834;-64.542975;15;-4;U;America/Tortola
6839;Allakaket Airport;Allakaket;United States;AET;PFAL;66.5519;-152.6222;441;-9;A;America/Anchorage
2905;Canouan;Canouan Island;Saint Vincent and the Grenadines;CIW;TVSC;12.699042;-61.342431;11;-4;U;America/St_Vincent
2906;Mustique;Mustique;Saint Vincent and the Grenadines;MQS;TVSM;12.887947;-61.180161;8;-4;U;America/St_Vincent
2907;E T Joshua;Kingstown;Saint Vincent and the Grenadines;SVD;TVSV;13.144306;-61.210861;66;-4;U;America/St_Vincent
2908;Almaty;Alma-ata;Kazakhstan;ALA;UAAA;43.352072;77.040508;2234;6;U;Asia/Qyzylorda
2909;Balkhash;Balkhash;Kazakhstan;;UAAH;46.893333;75.005;1446;6;U;Asia/Qyzylorda
2910;Astana Intl;Tselinograd;Kazakhstan;TSE;UACC;51.022222;71.466944;1165;6;U;Asia/Qyzylorda
2911;Taraz;Dzhambul;Kazakhstan;DMB;UADD;42.853611;71.303611;2184;6;U;Asia/Qyzylorda
2912;Manas;Bishkek;Kyrgyzstan;FRU;UAFM;43.061306;74.477556;2058;6;U;Asia/Bishkek
2913;Osh;Osh;Kyrgyzstan;OSS;UAFO;40.608989;72.793269;2927;6;U;Asia/Bishkek
2914;Shymkent;Chimkent;Kazakhstan;CIT;UAII;42.364167;69.478889;1385;6;U;Asia/Qyzylorda
6803;Yakutat;Yakutat;United States;YAK;PAYA;59.3012;-139.3937;33;-9;A;\N
2916;Uralsk;Uralsk;Kazakhstan;URA;UARR;51.150833;51.543056;125;5;U;Asia/Oral
2917;Pavlodar;Pavlodar;Kazakhstan;PWQ;UASP;52.195;77.073889;410;6;U;Asia/Qyzylorda
2918;Semipalatinsk;Semiplatinsk;Kazakhstan;PLX;UASS;50.3513;80.2344;761;6;U;Asia/Qyzylorda
2919;Aktau;Shevchenko;Kazakhstan;;UATE;43.86005;51.091978;73;5;U;Asia/Oral
2920;Aktyubinsk;Aktyubinsk;Kazakhstan;AKX;UATT;50.245833;57.206667;738;5;U;Asia/Oral
2922;Heydar Aliyev;Baku;Azerbaijan;GYD;UBBB;40.4675;50.046667;10;4;E;Asia/Baku
2923;Yakutsk;Yakutsk;Russia;YKS;UEEE;62.09325;129.770672;325;10;N;Asia/Yakutsk
2925;Mirny;Mirnyj;Russia;MJZ;UERR;62.534689;114.038928;1156;10;N;Asia/Yakutsk
2926;Ignatyevo;Blagoveschensk;Russia;BQS;UHBB;50.425394;127.412478;638;10;N;Asia/Yakutsk
2927;Novy;Khabarovsk;Russia;KHV;UHHH;48.528044;135.188361;244;11;N;Asia/Vladivostok
6838;Sawyer International Airport;Marquette;United States;MQT;KMQT;46.353611;-87.395278;1221;-5;A;America/New_York
2929;Provideniya Bay;Provideniya Bay;Russia;PVS;UHMD;64.378139;-173.243306;71;12;N;Asia/Magadan
2930;Sokol;Magadan;Russia;GDX;UHMM;59.910989;150.720439;574;12;N;Asia/Magadan
2931;Pevek;Pevek;Russia;PWE;UHMP;69.783283;170.597006;11;12;N;\N
2932;Yelizovo;Petropavlovsk;Russia;PKC;UHPP;53.167889;158.453669;131;12;N;Asia/Magadan
2933;Khomutovo;Yuzhno-sakhalinsk;Russia;UUS;UHSS;46.888672;142.717531;59;11;N;Asia/Vladivostok
2934;Knevichi;Vladivostok;Russia;VVO;UHWW;43.398953;132.148017;46;11;N;Asia/Vladivostok
2935;Kadala;Chita;Russia;HTA;UIAA;52.026317;113.305556;2272;10;N;Asia/Yakutsk
2936;Bratsk;Bratsk;Russia;BTK;UIBB;56.370556;101.698331;1610;9;N;Asia/Irkutsk
2937;Irkutsk;Irkutsk;Russia;IKT;UIII;52.268028;104.388975;1675;9;N;Asia/Irkutsk
2938;Mukhino;Ulan-ude;Russia;UUD;UIUU;51.807764;107.437644;1690;9;N;Asia/Irkutsk
2939;Boryspil Intl;Kiev;Ukraine;KBP;UKBB;50.345;30.894722;427;2;E;Europe/Kiev
2940;Donetsk Intl;Donetsk;Ukraine;DOK;UKCC;48.073611;37.739722;791;2;E;Europe/Kiev
2941;Dnipropetrovsk Intl;Dnepropetrovsk;Ukraine;DNK;UKDD;48.357222;35.100556;481;2;E;Europe/Kiev
2942;Simferopol Intl;Simferopol;Ukraine;SIP;UKFF;45.052222;33.975139;639;2;E;Europe/Kiev
2944;Zhuliany Intl;Kiev;Ukraine;IEV;UKKK;50.401694;30.449697;586;2;E;Europe/Kiev
2945;Lviv Intl;Lvov;Ukraine;LWO;UKLL;49.8125;23.956111;1071;2;E;Europe/Kiev
6837;Ford Airport;Iron Mountain;United States;IMT;KIMT;45.8183611;-88.1145556;1182;-6;A;America/Chicago
2947;Odesa Intl;Odessa;Ukraine;ODS;UKOO;46.426767;30.676464;172;2;E;Europe/Kiev
2948;Pulkovo;St. Petersburg;Russia;LED;ULLI;59.800292;30.262503;78;4;N;Europe/Moscow
2949;Murmansk;Murmansk;Russia;MMK;ULMM;68.781672;32.750822;266;4;N;Europe/Moscow
2950;Gomel;Gomel;Belarus;GME;UMGG;52.527022;31.016692;471;3;E;Europe/Minsk
2951;Vitebsk;Vitebsk;Belarus;VTB;UMII;55.1265;30.349639;683;3;E;Europe/Minsk
2952;Khrabrovo;Kaliningrad;Russia;KGD;UMKK;54.89005;20.592633;42;3;N;Europe/Kaliningrad
2953;Minsk 1;Minsk;Belarus;MHP;UMMM;53.864472;27.539683;748;3;E;Europe/Minsk
2954;Minsk 2;Minsk 2;Belarus;MSQ;UMMS;53.882469;28.030731;670;3;E;Europe/Minsk
2955;Abakan;Abakan;Russia;ABA;UNAA;53.74;91.385;831;8;N;Asia/Krasnoyarsk
2956;Barnaul;Barnaul;Russia;BAX;UNBB;53.363775;83.538533;838;7;N;Asia/Omsk
2957;Kemerovo;Kemorovo;Russia;KEJ;UNEE;55.270094;86.107208;863;7;N;Asia/Omsk
2958;Omsk;Omsk;Russia;OMS;UNOO;54.967042;73.310514;311;7;N;Asia/Omsk
2960;Pashkovskiy;Krasnodar;Russia;KRR;URKK;45.034689;39.170539;118;4;N;Europe/Moscow
2961;Uytash;Makhachkala;Russia;MCX;URML;42.816822;47.652294;12;4;N;Europe/Moscow
2962;Mineralnyye Vody;Mineralnye Vody;Russia;MRV;URMM;44.225072;43.081889;1054;4;N;Europe/Moscow
2963;Shpakovskoye;Stavropol;Russia;STW;URMT;45.109165;42.112778;1486;4;N;Europe/Moscow
2964;Rostov Na Donu;Rostov;Russia;ROV;URRR;47.258208;39.818089;280;4;N;Europe/Moscow
2965;Sochi;Sochi;Russia;AER;URSS;43.449928;39.956589;89;4;N;Europe/Moscow
2966;Astrakhan;Astrakhan;Russia;ASF;URWA;46.283333;48.006278;-65;4;N;Europe/Moscow
2967;Gumrak;Volgograd;Russia;VOG;URWW;48.782528;44.345544;482;4;N;Europe/Moscow
2968;Balandino;Chelyabinsk;Russia;CEK;USCC;55.305836;61.503333;769;6;N;Asia/Yekaterinburg
2969;Magnitogorsk;Magnetiogorsk;Russia;MQF;USCM;53.393131;58.755661;1431;6;N;Asia/Yekaterinburg
6834;Great Barrier Island;Claris;New Zealand;GBZ;NZGB;-36.1429;175.2819;10;12;Z;Pacific/Auckland
2972;Nizhnevartovsk;Nizhnevartovsk;Russia;NJC;USNN;60.949272;76.483617;177;6;N;Asia/Yekaterinburg
2973;Bolshoye Savino;Perm;Russia;PEE;USPP;57.914517;56.021214;404;6;N;Asia/Yekaterinburg
2974;Surgut;Surgut;Russia;SGC;USRR;61.343694;73.401842;200;6;N;Asia/Yekaterinburg
2975;Koltsovo;Yekaterinburg;Russia;SVX;USSS;56.743108;60.802728;764;6;N;Asia/Yekaterinburg
2976;Ashgabat;Ashkhabad;Turkmenistan;ASB;UTAA;37.986814;58.360967;692;5;U;Asia/Ashgabat
2977;Turkmenbashi;Krasnovodsk;Turkmenistan;KRW;UTAK;40.063333;53.007222;279;5;U;Asia/Ashgabat
2978;Turkmenabat;Chardzhou;Turkmenistan;;UTAV;39.083333;63.613333;630;5;U;Asia/Ashgabat
2979;Dushanbe;Dushanbe;Tajikistan;DYU;UTDD;38.543333;68.825;2575;5;U;Asia/Dushanbe
2980;Bukhara;Bukhara;Uzbekistan;BHK;UTSB;39.775;64.483333;751;5;U;Asia/Samarkand
2981;Samarkand;Samarkand;Uzbekistan;SKD;UTSS;39.700547;66.983829;2224;5;U;Asia/Samarkand
6833;Al Udeid AB;Doha;Qatar;IUD;OTBH;25.1174;51.3228;130;3;N;Asia/Qatar
2983;Yuzhny;Tashkent;Uzbekistan;TAS;UTTT;41.257861;69.281186;1417;5;U;Asia/Samarkand
2984;Bryansk;Bryansk;Russia;BZK;UUBP;53.214194;34.176447;663;4;N;Europe/Moscow
2985;Sheremetyevo;Moscow;Russia;SVO;UUEE;55.972642;37.414589;622;4;N;Europe/Moscow
2986;Migalovo;Tver;Russia;KLD;UUEM;56.824736;35.757678;469;4;N;Europe/Moscow
2987;Chertovitskoye;Voronezh;Russia;VOZ;UUOO;51.814211;39.229589;514;4;N;Europe/Moscow
2988;Vnukovo;Moscow;Russia;VKO;UUWW;55.591531;37.261486;685;4;N;Europe/Moscow
2989;Syktyvkar;Syktyvkar;Russia;SCW;UUYY;61.64705;50.84505;342;4;N;Europe/Moscow
2990;Kazan;Kazan;Russia;KZN;UWKD;55.606186;49.278728;411;4;N;Europe/Moscow
2991;Orenburg;Orenburg;Russia;REN;UWOO;51.795786;55.456744;387;6;N;Asia/Yekaterinburg
2992;Ufa;Ufa;Russia;UFA;UWUU;54.557511;55.874417;449;6;N;Asia/Yekaterinburg
2993;Kurumoch;Samara;Russia;KBY;UWWW;53.504858;50.164336;477;4;N;Europe/Moscow
2994;Ahmedabad;Ahmedabad;India;AMD;VAAH;23.077242;72.63465;189;5.5;N;Asia/Calcutta
2995;Akola;Akola;India;AKD;VAAK;20.699006;77.058628;999;5.5;N;Asia/Calcutta
2996;Aurangabad;Aurangabad;India;IXU;VAAU;19.862728;75.398114;1911;5.5;N;Asia/Calcutta
2997;Chhatrapati Shivaji Intl;Mumbai;India;BOM;VABB;19.088686;72.867919;37;5.5;N;Asia/Calcutta
2998;Bilaspur;Bilaspur;India;PAB;VABI;21.9884;82.110983;899;5.5;N;Asia/Calcutta
2999;Bhuj;Bhuj;India;BHJ;VABJ;23.287828;69.670147;268;5.5;N;Asia/Calcutta
3000;Belgaum;Belgaum;India;IXG;VABM;15.859286;74.618292;2487;5.5;N;Asia/Calcutta
3001;Vadodara;Baroda;India;BDQ;VABO;22.336164;73.226289;129;5.5;N;Asia/Calcutta
3002;Bhopal;Bhopal;India;BHO;VABP;23.287467;77.337375;1719;5.5;N;Asia/Calcutta
3003;Bhavnagar;Bhaunagar;India;BHU;VABV;21.752206;72.185181;44;5.5;N;Asia/Calcutta
3004;Daman;Daman;India;NMB;VADN;20.434364;72.843206;33;5.5;N;Asia/Calcutta
3005;Deesa;Deesa;India;;VADS;24.267936;72.204433;485;5.5;N;Asia/Calcutta
3006;Guna;Guna;India;;VAGN;24.654681;77.347347;1600;5.5;N;Asia/Calcutta
3007;Goa;Goa;India;GOI;VAGO;15.380833;73.831422;184;5.5;N;Asia/Calcutta
3008;Devi Ahilyabai Holkar;Indore;India;IDR;VAID;22.721786;75.801086;1850;5.5;N;Asia/Calcutta
3009;Jabalpur;Jabalpur;India;JLR;VAJB;23.177817;80.052047;1624;5.5;N;Asia/Calcutta
3010;Jamnagar;Jamnagar;India;JGA;VAJM;22.465522;70.012556;69;5.5;N;Asia/Calcutta
3011;Kandla;Kandla;India;IXY;VAKE;23.112719;70.100289;96;5.5;N;Asia/Calcutta
3012;Khajuraho;Khajuraho;India;HJR;VAKJ;24.817197;79.918597;728;5.5;N;Asia/Calcutta
3013;Kolhapur;Kolhapur;India;KLH;VAKP;16.664658;74.289353;1996;5.5;N;Asia/Calcutta
3014;Keshod;Keshod;India;IXK;VAKS;21.317069;70.270403;167;5.5;N;Asia/Calcutta
3015;Dr Ambedkar Intl;Nagpur;India;NAG;VANP;21.092192;79.047183;1033;5.5;N;Asia/Calcutta
3016;Nasik Road;Nasik Road;India;ISK;VANR;19.963739;73.807644;1959;5.5;N;Asia/Calcutta
3017;Pune;Pune;India;PNQ;VAPO;18.582111;73.919697;1942;5.5;N;Asia/Calcutta
3018;Porbandar;Porbandar;India;PBD;VAPR;21.648675;69.657219;23;5.5;N;Asia/Calcutta
3019;Rajkot;Rajkot;India;RAJ;VARK;22.309183;70.779525;441;5.5;N;Asia/Calcutta
3020;Raipur;Raipur;India;RPR;VARP;21.180406;81.738753;1041;5.5;N;Asia/Calcutta
3021;Sholapur;Sholapur;India;SSE;VASL;17.627958;75.934842;1584;5.5;N;Asia/Calcutta
3022;Surat;Surat;India;STV;VASU;21.114061;72.741792;16;5.5;N;Asia/Calcutta
3023;Udaipur;Udaipur;India;UDR;VAUD;24.617697;73.8961;1684;5.5;N;Asia/Calcutta
3024;Bandaranaike Intl Colombo;Colombo;Sri Lanka;CMB;VCBI;7.180756;79.884117;30;5.5;U;Asia/Colombo
3025;Anuradhapura;Anuradhapura;Sri Lanka;;VCCA;8.301486;80.4279;324;5.5;U;Asia/Colombo
3026;Batticaloa;Batticaloa;Sri Lanka;;VCCB;7.705756;81.678783;20;5.5;U;Asia/Colombo
3027;Colombo Ratmalana;Colombo;Sri Lanka;RML;VCCC;6.821994;79.886208;22;5.5;U;Asia/Colombo
3028;Amparai;Galoya;Sri Lanka;GOY;VCCG;7.337081;81.625881;150;5.5;U;Asia/Colombo
3029;Kankesanturai;Jaffna;Sri Lanka;JAF;VCCJ;9.792331;80.070089;33;5.5;U;Asia/Colombo
3030;China Bay;Trinciomalee;Sri Lanka;TRR;VCCT;8.538514;81.181853;6;5.5;U;Asia/Colombo
6832;Kirkuk AB;Kirkuk;Iraq;KIK;ORKK;35.17;44.3483;1061;3;N;Asia/Baghdad
3033;Kampong Chhnang;Kompong Chnang;Cambodia;;VDKH;12.255236;104.563875;56;7;U;Asia/Phnom_Penh
3034;Phnom Penh Intl;Phnom-penh;Cambodia;PNH;VDPP;11.546556;104.844139;40;7;U;Asia/Phnom_Penh
3035;Siem Reap;Siem-reap;Cambodia;REP;VDSR;13.410666;103.81284;60;7;U;Asia/Phnom_Penh
3036;Stung Treng;Stung Treng;Cambodia;;VDST;13.531897;106.014531;203;7;U;Asia/Phnom_Penh
3037;Along;Along;India;;VEAN;28.175317;94.802036;900;5.5;N;Asia/Calcutta
3038;Agartala;Agartala;India;IXA;VEAT;23.886978;91.24045;46;5.5;N;Asia/Calcutta
3039;Aizawl;Aizwal;India;AJL;VEAZ;23.746603;92.802767;1001;5.5;N;Asia/Calcutta
3040;Bagdogra;Baghdogra;India;IXB;VEBD;26.681206;88.328567;412;5.5;N;Asia/Calcutta
3041;Bokaro;Bokaro;India;;VEBK;23.643489;86.148886;715;5.5;N;Asia/Calcutta
3042;Bhubaneshwar;Bhubaneswar;India;BBI;VEBS;20.244364;85.817781;138;5.5;N;Asia/Calcutta
3043;Netaji Subhash Chandra Bose Intl;Kolkata;India;CCU;VECC;22.654739;88.446722;16;5.5;N;Asia/Calcutta
3044;Cooch Behar;Cooch-behar;India;COH;VECO;26.330508;89.467203;138;5.5;N;Asia/Calcutta
3045;Dhanbad;Dhanbad;India;DBD;VEDB;23.834044;86.425261;847;5.5;N;Asia/Calcutta
6800;Delta County Airport;Escanaba;United States;ESC;KESC;45.722778;-87.093611;609;-5;A;America/New_York
6801;Lauf-Lillinghof;Lillinghof;Germany;;\N;49.6543;11.6539;954;1;E;Europe/Berlin
3048;Gaya;Gaya;India;GAY;VEGY;24.744308;84.951175;380;5.5;N;Asia/Calcutta
3049;Hirakud;Hirakud;India;;VEHK;21.580231;84.005728;658;5.5;N;Asia/Calcutta
3050;Imphal;Imphal;India;IMF;VEIM;24.75995;93.896697;2540;5.5;N;Asia/Calcutta
3051;Jharsuguda;Jharsuguda;India;;VEJH;21.913536;84.050383;751;5.5;N;Asia/Calcutta
3052;Jamshedpur;Jamshedpur;India;IXW;VEJS;22.813211;86.168844;505;5.5;N;Asia/Calcutta
3053;Jorhat;Jorhat;India;JRH;VEJT;26.731528;94.175536;311;5.5;N;Asia/Calcutta
3054;Kailashahar;Kailashahar;India;IXH;VEKR;24.308192;92.007156;79;5.5;N;Asia/Calcutta
3055;Silchar;Silchar;India;IXS;VEKU;24.912928;92.978742;352;5.5;N;Asia/Calcutta
3056;Lilabari;Lilabari;India;IXI;VELR;27.295494;94.09765;330;5.5;N;Asia/Calcutta
3057;Dibrugarh;Mohanbari;India;MOH;VEMN;27.483853;95.016922;362;5.5;N;Asia/Calcutta
3058;Muzaffarpur;Mazuffarpur;India;;VEMZ;26.119089;85.313664;174;5.5;N;Asia/Calcutta
3059;Nawapara;Nawapara;India;;VENP;20.870036;82.519553;1058;5.5;N;Asia/Calcutta
3060;Panagarh;Panagarh;India;;VEPH;23.474336;87.427508;240;5.5;N;Asia/Calcutta
3061;Patna;Patina;India;PAT;VEPT;25.591317;85.087992;170;5.5;N;Asia/Calcutta
3062;Purnea;Purnea;India;;VEPU;25.759594;87.410011;129;5.5;N;Asia/Calcutta
3063;Birsa Munda;Ranchi;India;IXR;VERC;23.31425;85.321675;2148;5.5;N;Asia/Calcutta
3064;Rourkela;Rourkela;India;RRK;VERK;22.25665;84.814567;659;5.5;N;Asia/Calcutta
3065;Utkela;Utkela;India;;VEUK;20.097411;83.183797;680;5.5;N;Asia/Calcutta
3066;Vishakhapatnam;Vishakhapatnam;India;VTZ;VEVZ;17.721167;83.224483;15;5.5;N;Asia/Calcutta
3067;Zero;Zero;India;;VEZO;27.588283;93.828061;5403;5.5;N;Asia/Calcutta
3068;Coxs Bazar;Cox's Bazar;Bangladesh;CXB;VGCB;21.452194;91.963889;12;6;U;Asia/Dhaka
3069;Shah Amanat Intl;Chittagong;Bangladesh;CGP;VGEG;22.249611;91.813286;12;6;U;Asia/Dhaka
3070;Ishurdi;Ishurdi;Bangladesh;IRD;VGIS;24.1525;89.049446;45;6;U;Asia/Dhaka
3071;Jessore;Jessore;Bangladesh;JSR;VGJR;23.1838;89.160833;20;6;U;Asia/Dhaka
3072;Shah Mokhdum;Rajshahi;Bangladesh;RJH;VGRJ;24.437219;88.616511;64;6;U;Asia/Dhaka
3073;Saidpur;Saidpur;Bangladesh;SPD;VGSD;25.759228;88.908869;125;6;U;Asia/Dhaka
3074;Osmany Intl;Sylhet Osmani;Bangladesh;ZYL;VGSY;24.963242;91.866783;50;6;U;Asia/Dhaka
3075;Tejgaon;Dhaka;Bangladesh;;VGTJ;23.778783;90.382689;24;6;U;Asia/Dhaka
3076;Zia Intl;Dhaka;Bangladesh;DAC;VGZR;23.843333;90.397781;30;6;U;Asia/Dhaka
3077;Hong Kong Intl;Hong Kong;Hong Kong;HKG;VHHH;22.308919;113.914603;28;8;U;Asia/Hong_Kong
3078;Sek Kong;Sek Kong;Hong Kong;;VHSK;22.436592;114.080397;50;8;U;Asia/Hong_Kong
3079;Agra;Agra;India;AGR;VIAG;27.155831;77.960892;551;5.5;N;Asia/Calcutta
3080;Allahabad;Allahabad;India;IXD;VIAL;25.440064;81.733872;322;5.5;N;Asia/Calcutta
3081;Amritsar;Amritsar;India;ATQ;VIAR;31.709594;74.797264;756;5.5;N;Asia/Calcutta
3082;Nal;Bikaner;India;;VIBK;28.070606;73.207161;750;5.5;N;Asia/Calcutta
3083;Bakshi Ka Talab;Bakshi Ka Talab;India;;VIBL;26.988339;80.893117;385;5.5;N;Asia/Calcutta
3084;Varanasi;Varanasi;India;VNS;VIBN;25.452358;82.859342;266;5.5;N;Asia/Calcutta
3085;Kullu Manali;Kulu;India;KUU;VIBR;31.876706;77.154367;3573;5.5;N;Asia/Calcutta
3086;Bhatinda;Bhatinda;India;;VIBT;30.270139;74.755772;662;5.5;N;Asia/Calcutta
3087;Bhiwani;Bhiwani;India;;VIBW;28.837039;76.179094;720;5.5;N;Asia/Calcutta
3088;Bareilly;Bareilly;India;;VIBY;28.422061;79.450842;580;5.5;N;Asia/Calcutta
3089;Chandigarh;Chandigarh;India;IXC;VICG;30.673469;76.788542;1012;5.5;N;Asia/Calcutta
3090;Kanpur Chakeri;Kanpur;India;;VICX;26.4043;80.410119;405;5.5;N;Asia/Calcutta
3091;Safdarjung;Delhi;India;;VIDD;28.584511;77.205783;705;5.5;N;Asia/Calcutta
3092;Dehradun;Dehra Dun;India;DED;VIDN;30.189689;78.180256;1831;5.5;N;Asia/Calcutta
3093;Indira Gandhi Intl;Delhi;India;DEL;VIDP;28.5665;77.103088;777;5.5;N;Asia/Calcutta
3094;Gwalior;Gwalior;India;GWL;VIGR;26.293336;78.227753;617;5.5;N;Asia/Calcutta
3095;Hissar;Hissar;India;;VIHR;29.179444;75.755336;700;5.5;N;Asia/Calcutta
3096;Jhansi;Jhansi;India;;VIJN;25.491172;78.558422;801;5.5;N;Asia/Calcutta
3097;Jodhpur;Jodhpur;India;JDH;VIJO;26.251092;73.048869;717;5.5;N;Asia/Calcutta
3098;Jaipur;Jaipur;India;JAI;VIJP;26.824192;75.812161;1263;5.5;N;Asia/Calcutta
3099;Jaisalmer;Jaisalmer;India;JSA;VIJR;26.888653;70.864967;751;5.5;N;Asia/Calcutta
3100;Jammu;Jammu;India;IXJ;VIJU;32.689142;74.837389;1029;5.5;N;Asia/Calcutta
3101;Kanpur;Kanpur;India;KNU;VIKA;26.441444;80.364864;411;5.5;N;Asia/Calcutta
3102;Kota;Kota;India;KTU;VIKO;25.160219;75.845631;896;5.5;N;Asia/Calcutta
3103;Ludhiana;Ludhiaha;India;LUH;VILD;30.854681;75.952592;834;5.5;N;Asia/Calcutta
3104;Leh;Leh;India;IXL;VILH;34.135872;77.546514;10682;5.5;N;Asia/Calcutta
3105;Lucknow;Lucknow;India;LKO;VILK;26.760594;80.889339;410;5.5;N;Asia/Calcutta
3106;Pathankot;Pathankot;India;IXP;VIPK;32.233778;75.634628;1017;5.5;N;Asia/Calcutta
3107;Patiala;Patiala;India;;VIPL;30.314847;76.364469;820;5.5;N;Asia/Calcutta
3108;Pantnagar;Nainital;India;PGH;VIPT;29.033408;79.473744;769;5.5;N;Asia/Calcutta
3109;Fursatganj;Raibarelli;India;;VIRB;26.248489;81.380506;360;5.5;N;Asia/Calcutta
3111;Sarsawa;Saharanpur;India;;VISP;29.993919;77.425311;891;5.5;N;Asia/Calcutta
3112;Srinagar;Srinagar;India;SXR;VISR;33.987139;74.77425;5429;5.5;N;Asia/Calcutta
3113;Satna;Satna;India;TNI;VIST;24.562319;80.854933;1060;5.5;N;Asia/Calcutta
6495;Balkhash Airport;Balkhash;Kazakhstan;BXH;\N;46.8933;75.005;1446;6;E;Asia/Qyzylorda
3115;Luang Phabang Intl;Luang Prabang;Laos;LPQ;VLLB;19.897914;102.160764;955;7;U;Asia/Vientiane
3116;Pakse;Pakse;Laos;PKZ;VLPS;15.132053;105.781417;351;7;U;Asia/Vientiane
3117;Phonesavanh;Phong Savanh;Laos;;VLPV;19.454864;103.218208;3628;7;U;Asia/Vientiane
3118;Savannakhet;Savannakhet;Laos;ZVK;VLSK;16.556594;104.759531;509;7;U;Asia/Vientiane
3119;Sam Neua;Sam Neua;Laos;;VLSN;20.418358;104.066583;3281;7;U;Asia/Vientiane
3120;Wattay Intl;Vientiane;Laos;VTE;VLVT;17.988322;102.563256;564;7;U;Asia/Vientiane
3121;Macau Intl;Macau;Macau;MFM;VMMC;22.149556;113.591558;20;8;U;Asia/Macau
3122;Bhairahawa;Bhairawa;Nepal;BWA;VNBW;27.505703;83.41625;358;5.75;N;Asia/Katmandu
6831;Nastaetten;Nastaetten;Germany;;\N;50.1972;7.8916;600;1;E;Europe/Berlin
3124;Janakpur;Janakpur;Nepal;;VNJP;26.708806;85.922394;256;5.75;N;Asia/Katmandu
3125;Tribhuvan Intl;Kathmandu;Nepal;KTM;VNKT;27.696583;85.3591;4390;5.75;N;Asia/Katmandu
3127;Pokhara;Pokhara;Nepal;PKR;VNPK;28.200881;83.982056;2712;5.75;N;Asia/Katmandu
3128;Simara;Simara;Nepal;SIF;VNSI;27.159456;84.980122;450;5.75;N;Asia/Katmandu
3129;Biratnagar;Biratnagar;Nepal;BIR;VNVT;26.481453;87.264036;236;5.75;N;Asia/Katmandu
3130;Agatti;Agatti Island;India;AGX;VOAT;10.823656;72.176042;14;5.5;N;Asia/Calcutta
3131;Bangalore;Bangalore;India;BLR;VOBL;12.949986;77.668206;2912;5.5;N;Asia/Calcutta
3132;Bellary;Bellary;India;BEP;VOBI;15.162783;76.882775;30;5.5;N;Asia/Calcutta
3133;Bidar;Bidar;India;;VOBR;17.908081;77.487142;2178;5.5;N;Asia/Calcutta
3134;Vijayawada;Vijayawada;India;VGA;VOBZ;16.530433;80.796847;82;5.5;N;Asia/Calcutta
3135;Coimbatore;Coimbatore;India;CJB;VOCB;11.030031;77.043383;1324;5.5;N;Asia/Calcutta
3136;Cochin;Kochi;India;COK;VOCI;10.155556;76.391389;8;5.5;N;Asia/Calcutta
3137;Calicut;Calicut;India;CCJ;VOCL;11.136839;75.9553;342;5.5;N;Asia/Calcutta
3138;Cuddapah;Cuddapah;India;CDP;VOCP;14.509961;78.772833;430;5.5;N;Asia/Calcutta
3139;Carnicobar;Carnicobar;India;;VOCX;9.152508;92.819628;5;5.5;N;Asia/Calcutta
3140;Dundigul;Dundigul;India;;VODG;17.627178;78.403361;2013;5.5;N;Asia/Calcutta
3141;Hyderabad;Hyderabad;India;HYD;VOHY;17.453117;78.467586;1742;5.5;N;Asia/Calcutta
3142;Madurai;Madurai;India;IXM;VOMD;9.834508;78.093378;459;5.5;N;Asia/Calcutta
3143;Mangalore;Mangalore;India;IXE;VOML;12.961267;74.890069;337;5.5;N;Asia/Calcutta
3144;Chennai Intl;Madras;India;MAA;VOMM;12.994414;80.180517;52;5.5;N;Asia/Calcutta
3145;Nagarjuna Sagar;Nagarjunsagar;India;;VONS;16.542653;79.318714;658;5.5;N;Asia/Calcutta
3146;Port Blair;Port Blair;India;IXZ;VOPB;11.641161;92.729744;14;5.5;N;Asia/Calcutta
3147;Pondicherry;Pendicherry;India;PNY;VOPC;11.968722;79.810058;118;5.5;N;Asia/Calcutta
3148;Rajahmundry;Rajahmundry;India;RJA;VORY;17.110361;81.818208;151;5.5;N;Asia/Calcutta
3149;Salem;Salem;India;;VOSM;11.783314;78.065606;1008;5.5;N;Asia/Calcutta
3150;Tanjore;Tanjore;India;;VOTJ;10.722428;79.101567;253;5.5;N;Asia/Calcutta
3151;Tirupati;Tirupeti;India;TIR;VOTP;13.632492;79.543256;350;5.5;N;Asia/Calcutta
3152;Trichy;Tiruchirappalli;India;TRZ;VOTR;10.765364;78.709722;288;5.5;N;Asia/Calcutta
3153;Thiruvananthapuram Intl;Trivandrum;India;TRV;VOTV;8.482122;76.920114;15;5.5;N;Asia/Calcutta
3154;Tambaram;Tambaram;India;;VOTX;12.907214;80.121861;90;5.5;N;Asia/Calcutta
3155;Paro;Thimphu;Bhutan;PBH;VQPR;27.403192;89.424606;7332;6;N;Asia/Thimphu
3156;Male Intl;Male;Maldives;MLE;VRMM;4.191833;73.529128;6;5;U;Indian/Maldives
3157;Don Muang Intl;Bangkok;Thailand;DMK;VTBD;13.912583;100.60675;9;7;U;Asia/Bangkok
3158;Kamphaeng Saen;Nakhon Pathom;Thailand;;VTBK;14.101975;99.917219;30;7;U;Asia/Bangkok
3159;Khok Kathiam;Lop Buri;Thailand;;VTBL;14.874561;100.663367;123;7;U;Asia/Bangkok
3886;Naha;Naha;Indonesia;NAH;WAMH;3.683214;125.528019;16;8;N;Asia/Makassar
3161;U Taphao Intl;Pattaya;Thailand;UTP;VTBU;12.679944;101.005028;42;7;U;Asia/Bangkok
3162;Watthana Nakhon;Prachin Buri;Thailand;;VTBW;13.7688;102.315492;200;7;U;Asia/Bangkok
3163;Lampang;Lampang;Thailand;LPT;VTCL;18.270933;99.504167;811;7;U;Asia/Bangkok
3164;Phrae;Phrae;Thailand;PRH;VTCP;18.132169;100.164664;538;7;U;Asia/Bangkok
3165;Hua Hin;Prachuap Khiri Khan;Thailand;HHQ;VTPH;12.636225;99.951533;62;7;U;Asia/Bangkok
3166;Takhli;Nakhon Sawan;Thailand;;VTPI;15.277306;100.295861;107;7;U;Asia/Bangkok
3167;Sak Long;Phetchabun;Thailand;;VTPL;16.824322;101.251389;500;7;U;Asia/Bangkok
3169;Nakhon Sawan;Nakhon Sawan;Thailand;;VTPN;15.672997;100.136794;113;7;U;Asia/Bangkok
3170;Phitsanulok;Phitsanulok;Thailand;PHS;VTPP;16.782939;100.279122;154;7;U;Asia/Bangkok
3171;Khunan Phumipol;Tak;Thailand;;VTPY;17.234211;99.057911;492;7;U;Asia/Bangkok
3172;Khoun Khan;Satun;Thailand;;VTSA;6.661403;100.080317;18;7;U;Asia/Bangkok
3173;Narathiwat;Narathiwat;Thailand;NAW;VTSC;6.519922;101.7434;16;7;U;Asia/Bangkok
3174;Krabi;Krabi;Thailand;KBV;VTSG;8.095969;98.988764;93;7;U;Asia/Bangkok
3175;Songkhla;Songkhla;Thailand;;VTSH;7.186564;100.608031;12;7;U;Asia/Bangkok
3176;Pattani;Pattani;Thailand;PAN;VTSK;6.785458;101.153569;8;7;U;Asia/Bangkok
3177;Samui;Ko Samui;Thailand;USM;VTSM;9.547794;100.062272;64;7;U;Asia/Bangkok
3178;Cha Ian;Nakhon Si Thammarat;Thailand;;VTSN;8.471147;99.955625;44;7;U;Asia/Bangkok
3179;Phuket Intl;Phuket;Thailand;HKT;VTSP;8.1132;98.316872;82;7;U;Asia/Bangkok
3180;Ranong;Ranong;Thailand;;VTSR;9.777622;98.585483;57;7;U;Asia/Bangkok
3181;Hat Yai Intl;Hat Yai;Thailand;HDY;VTSS;6.933206;100.392975;90;7;U;Asia/Bangkok
3182;Trang;Trang;Thailand;TST;VTST;7.508744;99.616578;67;7;U;Asia/Bangkok
3183;Udon Thani;Udon Thani;Thailand;UTH;VTUD;17.386436;102.788247;579;7;U;Asia/Bangkok
3184;Sakon Nakhon;Sakon Nakhon;Thailand;SNO;VTUI;17.195142;104.118625;529;7;U;Asia/Bangkok
3185;Surin;Surin;Thailand;;VTUJ;14.868264;103.498256;478;7;U;Asia/Bangkok
3186;Loei;Loei;Thailand;LOE;VTUL;17.439133;101.722064;860;7;U;Asia/Bangkok
3187;Khorat;Nakhon Ratchasima;Thailand;;VTUN;14.934514;102.078639;729;7;U;Asia/Bangkok
3188;Rob Muang;Roi Et;Thailand;;VTUR;16.07035;103.6459;459;7;U;Asia/Bangkok
6830;Finsterwalde-Heinrichsruh;Finsterwalde;Germany;;EDAS;51.6377;13.2419;384;1;E;Europe/Berlin
6829;Orchid Beach;Fraser Island;Australia;OKB;KOKB;-24.95841;153.3145;6;10;O;Australia/Brisbane
6828;Mara Lodges;Mara Lodges;Kenya;;\N;-1.183581;35.099931;5600;3;U;Africa/Nairobi
3196;Danang Intl;Danang;Vietnam;DAD;VVDN;16.043917;108.19937;33;7;U;Asia/Saigon
3197;Hanoi Gia Lam;Hanoi;Vietnam;;VVGL;21.040975;105.886011;50;7;U;Asia/Saigon
3198;Kep;Kep;Vietnam;;VVKP;21.394639;106.261083;55;7;U;Asia/Saigon
3199;Noibai Intl;Hanoi;Vietnam;HAN;VVNB;21.221192;105.807178;39;7;U;Asia/Saigon
3200;Nhatrang;Nhatrang;Vietnam;NHA;VVNT;12.227467;109.192322;20;7;U;Asia/Saigon
3201;Phubai;Hue;Vietnam;;VVPB;16.4015;107.702614;48;7;U;Asia/Saigon
6827;Namanga;Namanga;Kenya;;\N;-2.55;36.7833333;4445;3;U;Africa/Dar_es_Salaam
3204;Phu Quoc;Phuquoc;Vietnam;;VVPQ;10.227025;103.967169;23;7;U;Asia/Saigon
3205;Tansonnhat Intl;Ho Chi Minh City;Vietnam;SGN;VVTS;10.818797;106.651856;33;7;U;Asia/Saigon
3207;Ann;Ann;Burma;;VYAN;19.769156;94.026133;74;6.5;U;Asia/Rangoon
3208;Anisakan;Anisakan;Burma;;VYAS;21.955433;96.40605;3000;6.5;U;Asia/Rangoon
3209;Bagan;Bagan;Burma;;VYBG;21.178756;94.930169;312;6.5;U;Asia/Rangoon
6826;Bamburi;Bamburi;Kenya;BMQ;KBMQ;-3.98268888888889;39.7308972222222;50;3;U;Africa/Nairobi
3211;Coco Island;Coco Island;Burma;;VYCI;14.141517;93.368531;20;6;U;\N
3213;Heho;Heho;Burma;HEH;VYHH;20.747036;96.792044;3858;6.5;U;Asia/Rangoon
3214;Hommalinn;Hommalin;Burma;;VYHL;24.899597;94.914033;534;6.5;U;Asia/Rangoon
3215;Kengtung;Kengtung;Burma;KET;VYKG;21.301611;99.635997;2798;6.5;U;Asia/Rangoon
6825;Williamson Country Regional Airport;Marion;United States;MWA;KMWA;37.7549569;-89.0110936;472;-6;A;America/Chicago
3217;Kyaukpyu;Kyaukpyu;Burma;KYP;VYKP;19.426447;93.534836;20;6.5;U;Asia/Rangoon
6824;Lommis Airport;Lommis;Switzerland;;LSZT;47.5244;9.00306;1539;1;E;Europe/Zurich
3220;Lashio;Lashio;Burma;LSH;VYLS;22.977881;97.752183;2450;6.5;U;Asia/Rangoon
3221;Lanywa;Lanywa;Burma;;VYLY;20.940361;94.822617;175;6.5;U;Asia/Rangoon
3222;Mandalay Intl;Mandalay;Burma;MDL;VYMD;21.702156;95.977928;300;6.5;U;Asia/Rangoon
3223;Myeik;Myeik;Burma;MGZ;VYME;12.439797;98.621478;75;6.5;U;Asia/Rangoon
3224;Myitkyina;Myitkyina;Burma;MYT;VYMK;25.383636;97.351919;475;6.5;U;Asia/Rangoon
3226;Momeik;Momeik;Burma;;VYMO;23.092525;96.645272;600;6.5;U;Asia/Rangoon
3227;Mong Hsat;Mong Hsat;Burma;MOG;VYMS;20.516758;99.256825;1875;6.5;U;Asia/Rangoon
3228;Nampong;Nampong;Burma;;VYNP;25.354375;97.29515;459;6.5;U;Asia/Rangoon
3229;Namsang;Namsang;Burma;;VYNS;20.890492;97.735922;3100;6.5;U;Asia/Rangoon
3230;Hpa An;Hpa-an;Burma;;VYPA;16.893714;97.674581;150;6.5;U;Asia/Rangoon
6823;Amlikon Glider Airport;Amlikon;Switzerland;;LSPA;47.5742;9.0475;1371;1;E;Europe/Zurich
3232;Putao;Putao;Burma;PBU;VYPT;27.329922;97.426269;1500;6.5;U;Asia/Rangoon
3233;Pyay;Pyay;Burma;;VYPY;18.824478;95.266003;120;6.5;U;Asia/Rangoon
3234;Shante;Shante;Burma;;VYST;20.941668;95.914497;630;6.5;U;Asia/Rangoon
3235;Sittwe;Sittwe;Burma;AKY;VYSW;20.132708;92.872628;27;6.5;U;Asia/Rangoon
3236;Thandwe;Thandwe;Burma;SNW;VYTD;18.460731;94.300119;20;6.5;U;Asia/Rangoon
3237;Tachileik;Tachilek;Burma;THL;VYTL;20.483831;99.935353;1280;6.5;U;Asia/Rangoon
3238;Taungoo;Taungoo;Burma;;VYTO;19.031275;96.401239;160;6.5;U;Asia/Rangoon
3239;Yangon Intl;Yangon;Burma;RGN;VYYY;16.907305;96.133222;109;6.5;U;Asia/Rangoon
3240;Hasanuddin;Ujung Pandang;Indonesia;UPG;WAAA;-5.061631;119.554042;47;8;N;Asia/Makassar
3241;Frans Kaisiepo;Biak;Indonesia;BIK;WABB;-1.190017;136.107997;46;9;N;Asia/Jayapura
3242;Nabire;Nabire;Indonesia;NBX;WABI;-3.368183;135.496406;20;9;N;Asia/Jayapura
3243;Moses Kilangin;Timika;Indonesia;TIM;WABP;-4.528275;136.887375;103;9;N;Asia/Jayapura
3244;Sentani;Jayapura;Indonesia;DJJ;WAJJ;-2.576953;140.516372;289;9;N;Asia/Jayapura
3245;Wamena;Wamena;Indonesia;WMX;WAJW;-4.102511;138.957372;5085;9;N;Asia/Jayapura
3246;Mopah;Merauke;Indonesia;MKQ;WAKK;-8.520294;140.418453;10;9;N;Asia/Jayapura
3247;Jalaluddin;Gorontalo;Indonesia;GTO;WAMG;0.637119;122.849858;105;8;N;Asia/Makassar
3930;Incheon Intl;Seoul;South Korea;ICN;RKSI;37.469075;126.450517;23;9;U;Asia/Seoul
3249;Mutiara;Palu;Indonesia;PLW;WAML;-0.918542;119.909642;284;8;N;Asia/Makassar
3250;Sam Ratulangi;Manado;Indonesia;MDC;WAMM;1.549447;124.925878;264;8;N;Asia/Makassar
3251;Kasiguncu;Poso;Indonesia;PSJ;WAMP;-1.416753;120.657669;174;8;N;Asia/Makassar
3252;Pitu;Morotai Island;Indonesia;OTI;WAMR;2.045992;128.324708;49;9;N;Asia/Jayapura
3253;Sultan Babullah;Ternate;Indonesia;TTE;WAMT;0.831414;127.381486;49;9;N;Asia/Jayapura
3254;Bubung;Luwuk;Indonesia;LUW;WAMW;-1.038919;122.771906;56;8;N;Asia/Makassar
6821;Mukachevo Air Base;Mukacheve;Ukraine;;\N;48.4;22.6833;390;2;E;Europe/Kiev
3256;Pattimura;Ambon;Indonesia;AMQ;WAPP;-3.710264;128.089136;33;9;N;Asia/Jayapura
3257;Fak Fak;Fak Fak;Indonesia;FKQ;WASF;-2.920192;132.267031;462;9;N;Asia/Jayapura
3258;Kaimana;Kaimana;Indonesia;KNG;WASK;-3.644517;133.695553;19;9;N;Asia/Jayapura
3259;Babo;Babo;Indonesia;BXB;WASO;-2.532242;133.438894;10;9;N;Asia/Jayapura
3260;Rendani;Manokwari;Indonesia;MKW;WASR;-0.891833;134.049183;23;9;N;Asia/Jayapura
3261;Jefman;Sorong;Indonesia;SOQ;WASS;-0.926358;131.121194;10;9;N;Asia/Jayapura
3262;Bintulu;Bintulu;Malaysia;BTU;WBGB;3.12385;113.020472;74;8;N;Asia/Kuala_Lumpur
3263;Kuching Intl;Kuching;Malaysia;KCH;WBGG;1.484697;110.346933;89;8;N;Asia/Kuala_Lumpur
3264;Limbang;Limbang;Malaysia;LMN;WBGJ;4.808303;115.010439;14;8;N;Asia/Kuala_Lumpur
3265;Marudi;Marudi;Malaysia;MUR;WBGM;4.1775;114.321944;103;8;N;Asia/Kuala_Lumpur
3266;Miri;Miri;Malaysia;MYY;WBGR;4.322014;113.986806;59;8;N;Asia/Kuala_Lumpur
3267;Sibu;Sibu;Malaysia;SBW;WBGS;2.261603;111.985322;122;8;N;Asia/Kuala_Lumpur
3268;Lahad Datu;Lahad Datu;Malaysia;LDU;WBKD;5.032247;118.324036;45;8;N;Asia/Kuala_Lumpur
3269;Kota Kinabalu Intl;Kota Kinabalu;Malaysia;BKI;WBKK;5.937208;116.051181;10;8;N;Asia/Kuala_Lumpur
3270;Labuan;Labuan;Malaysia;LBU;WBKL;5.300683;115.250181;101;8;N;Asia/Kuala_Lumpur
3271;Tawau;Tawau;Malaysia;TWU;WBKW;4.313369;118.121953;57;8;N;Asia/Kuala_Lumpur
3272;Brunei Intl;Bandar Seri Begawan;Brunei;BWN;WBSB;4.9442;114.928353;73;8;U;Asia/Brunei
3273;Sultan Syarif Kasim Ii;Pekanbaru;Indonesia;PKU;WIBB;0.460786;101.444539;102;7;N;Asia/Jakarta
3274;Pinang Kampai;Dumai;Indonesia;DUM;WIBD;1.609194;101.433558;55;7;N;Asia/Jakarta
3275;Soekarno Hatta Intl;Jakarta;Indonesia;CGK;WIII;-6.125567;106.655897;34;7;N;Asia/Jakarta
3276;Binaka;Gunung Sitoli;Indonesia;GNS;WIMB;1.166381;97.704681;20;7;N;Asia/Jakarta
3277;Aek Godang;Padang Sidempuan;Indonesia;;WIME;1.400103;99.430453;922;7;N;Asia/Jakarta
3278;Minangkabau;Padang;Indonesia;PDG;WIPT;-0.874989;100.351881;9;7;N;Asia/Jakarta
3279;Polonia;Medan;Indonesia;MES;WIMM;3.558056;98.671722;114;7;N;Asia/Jakarta
3280;Dr Ferdinand Lumban Tobing;Sibolga;Indonesia;;WIMS;1.555944;98.888908;43;7;N;Asia/Jakarta
3281;Nanga Pinoh I;Nangapinoh;Indonesia;;WIOG;-0.348869;111.747606;123;7;N;Asia/Jakarta
3282;Rahadi Usman;Ketapang;Indonesia;KTG;WIOK;-1.816639;109.963483;46;7;N;Asia/Jakarta
6820;Mukachevo;Mukacheve;Ukraine;;\N;48.4;22.683;390;2;E;Europe/Kiev
3284;Supadio;Pontianak;Indonesia;PNK;WIOO;-0.150711;109.403892;10;7;N;Asia/Jakarta
6819;Borovaya Airfield;Minsk;Belarus;;UMMB;53.57;27.39;0;3;E;Europe/Minsk
6798;Weser-Wuemme;Hellwege;Germany;;EDWM;53.054;9.208667;59;1;E;Europe/Berlin
3287;Sultan Thaha;Jambi;Indonesia;DJB;WIPA;-1.638017;103.644378;82;7;N;Asia/Jakarta
3288;Fatmawati Soekarno;Bengkulu;Indonesia;BKS;WIPL;-3.8637;102.339036;50;7;N;Asia/Jakarta
3289;Sultan Mahmud Badaruddin Ii;Palembang;Indonesia;PLM;WIPP;-2.89825;104.699903;49;7;N;Asia/Jakarta
3291;Japura;Rengat;Indonesia;RGT;WIPR;-0.352808;102.334917;62;7;N;Asia/Jakarta
3292;Lhok Sukon;Lhok Sukon;Indonesia;;WITL;5.069506;97.259192;28;7;N;Asia/Jakarta
6818;Balti International Airport;Strymba;Moldova;BZY;\N;47.8381;27.7815;758;2;U;Europe/Chisinau
3294;Sultan Iskandarmuda;Banda Aceh;Indonesia;BTJ;WITT;5.523522;95.420372;65;7;N;Asia/Jakarta
3295;Kluang;Kluang;Malaysia;;WMAP;2.041394;103.307394;150;8;N;Asia/Kuala_Lumpur
3296;Sultan Abdul Halim;Alor Setar;Malaysia;AOR;WMKA;6.189667;100.398183;15;8;N;Asia/Kuala_Lumpur
3297;Butterworth;Butterworth;Malaysia;;WMKB;5.465917;100.391167;11;8;N;Asia/Kuala_Lumpur
3298;Sultan Ismail Petra;Kota Bahru;Malaysia;KBR;WMKC;6.16685;102.293014;16;8;N;Asia/Kuala_Lumpur
3299;Kuantan;Kuantan;Malaysia;KUA;WMKD;3.775389;103.209056;58;8;N;Asia/Kuala_Lumpur
3300;Kerteh;Kerteh;Malaysia;KTE;WMKE;4.537222;103.426756;18;8;N;Asia/Kuala_Lumpur
3301;Simpang;Simpang;Malaysia;;WMKF;3.11225;101.70275;111;8;N;Asia/Kuala_Lumpur
3302;Sultan Azlan Shah;Ipoh;Malaysia;IPH;WMKI;4.567972;101.092194;130;8;N;Asia/Kuala_Lumpur
3303;Sultan Ismail;Johor Bahru;Malaysia;JHB;WMKJ;1.641308;103.669619;135;8;N;Asia/Kuala_Lumpur
3304;Kuala Lumpur Intl;Kuala Lumpur;Malaysia;KUL;WMKK;2.745578;101.709917;69;8;N;Asia/Kuala_Lumpur
3305;Langkawi Intl;Pulau;Malaysia;LGK;WMKL;6.329728;99.728667;29;8;N;Asia/Kuala_Lumpur
3306;Malacca;Malacca;Malaysia;MKZ;WMKM;2.263361;102.251553;35;8;N;Asia/Kuala_Lumpur
3307;Sultan Mahmud;Kuala Terengganu;Malaysia;TGG;WMKN;5.382639;103.10336;21;8;N;Asia/Kuala_Lumpur
3308;Penang Intl;Penang;Malaysia;PEN;WMKP;5.297139;100.276864;11;8;N;Asia/Kuala_Lumpur
3309;Suai;Suai;East Timor;;WPDB;-9.303306;125.286753;96;9;U;Asia/Dili
3310;Presidente Nicolau Lobato Intl;Dili;East Timor;DIL;WPDL;-8.546553;125.524719;25;9;U;Asia/Dili
3311;Cakung;Baucau;East Timor;;WPEC;-8.485547;126.399389;1777;9;U;Asia/Dili
3312;Sembawang;Sembawang;Singapore;;WSAG;1.425264;103.812806;86;8;N;Asia/Singapore
3313;Paya Lebar;Paya Lebar;Singapore;QPG;WSAP;1.360417;103.90953;65;8;N;Asia/Singapore
3314;Tengah;Tengah;Singapore;;WSAT;1.387258;103.708719;50;8;N;Asia/Singapore
3315;Seletar;Singapore;Singapore;XSP;WSSL;1.41695;103.867653;36;8;N;Asia/Singapore
3316;Changi Intl;Singapore;Singapore;SIN;WSSS;1.350189;103.994433;22;8;N;Asia/Singapore
3317;Brisbane Archerfield;Brisbane;Australia;;YBAF;-27.570278;153.008056;63;10;O;Australia/Brisbane
3318;Bamaga Injinoo;Amberley;Australia;ABM;YBAM;-10.950833;142.459444;34;10;O;Australia/Brisbane
3319;Alice Springs;Alice Springs;Australia;ASP;YBAS;-23.806667;133.902222;1789;9.5;N;Australia/Darwin
3320;Brisbane Intl;Brisbane;Australia;BNE;YBBN;-27.384167;153.1175;13;10;N;Australia/Brisbane
3321;Gold Coast;Coolangatta;Australia;OOL;YBCG;-28.164444;153.504722;21;10;N;Australia/Brisbane
3322;Cairns Intl;Cairns;Australia;CNS;YBCS;-16.885833;145.755278;10;10;N;Australia/Brisbane
3323;Charleville;Charlieville;Australia;CTL;YBCV;-26.413334;146.2625;1003;10;O;Australia/Brisbane
3324;Mount Isa;Mount Isa;Australia;ISA;YBMA;-20.663889;139.488611;1121;10;O;Australia/Brisbane
3325;Sunshine Coast;Maroochydore;Australia;MCY;YBMC;-26.603333;153.091111;15;10;O;Australia/Brisbane
3326;Mackay;Mackay;Australia;MKY;YBMK;-21.171667;149.179722;19;10;O;Australia/Brisbane
3328;Proserpine Whitsunday Coast;Prosserpine;Australia;PPP;YBPN;-20.495;148.552222;82;10;O;Australia/Brisbane
3329;Rockhampton;Rockhampton;Australia;ROK;YBRK;-23.381944;150.475278;34;10;O;Australia/Brisbane
3330;Townsville;Townsville;Australia;TSV;YBTL;-19.2525;146.765278;18;10;N;Australia/Brisbane
3331;Weipa;Weipa;Australia;WEI;YBWP;-12.678611;141.925278;63;10;O;Australia/Brisbane
3332;Avalon;Avalon;Australia;AVV;YMAV;-38.039444;144.469444;35;10;O;Australia/Hobart
3333;Albury;Albury;Australia;ABX;YMAY;-36.067778;146.958056;539;10;O;Australia/Sydney
3334;Melbourne Essendon;Melbourne;Australia;MEB;YMEN;-37.728056;144.901944;282;10;O;Australia/Hobart
3335;East Sale;East Sale;Australia;;YMES;-38.098889;147.149444;23;10;O;Australia/Hobart
3336;Hobart;Hobart;Australia;HBA;YMHB;-42.836111;147.510278;13;10;O;Australia/Melbourne
3337;Launceston;Launceston;Australia;LST;YMLT;-41.545278;147.214167;562;10;O;Australia/Melbourne
3338;Melbourne Moorabbin;Melbourne;Australia;MBW;YMMB;-37.975833;145.102222;50;10;O;Australia/Hobart
3339;Melbourne Intl;Melbourne;Australia;MEL;YMML;-37.673333;144.843333;434;10;O;Australia/Hobart
3340;Point Cook;Point Cook;Australia;;YMPC;-37.932222;144.753333;14;10;O;Australia/Hobart
3341;Adelaide Intl;Adelaide;Australia;ADL;YPAD;-34.945;138.530556;20;9.5;O;Australia/Adelaide
3343;Edinburgh;Edinburgh;Australia;;YPED;-34.7025;138.620833;67;9.5;O;Australia/Adelaide
3344;Perth Jandakot;Perth;Australia;JAD;YPJT;-32.0975;115.881111;99;8;O;Australia/Perth
3345;Karratha;Karratha;Australia;KTA;YPKA;-20.712222;116.773333;29;8;O;Australia/Perth
3346;Kalgoorlie Boulder;Kalgoorlie;Australia;KGI;YPKG;-30.789444;121.461667;1203;8;O;Australia/Perth
3347;Kununurra;Kununurra;Australia;KNX;YPKU;-15.778056;128.7075;145;8;O;Australia/Perth
3348;Learmonth;Learmonth;Australia;LEA;YPLM;-22.235556;114.088611;19;8;O;Australia/Perth
3349;Port Hedland Intl;Port Hedland;Australia;PHE;YPPD;-20.377778;118.626389;33;8;O;Australia/Perth
3350;Adelaide Parafield;Adelaide;Australia;;YPPF;-34.793333;138.633056;57;9.5;O;Australia/Adelaide
3351;Perth Intl;Perth;Australia;PER;YPPH;-31.940278;115.966944;67;8;N;Australia/Perth
3352;Woomera;Woomera;Australia;UMR;YPWR;-31.144167;136.816944;548;9.5;O;Australia/Adelaide
3353;Christmas Island;Christmas Island;Christmas Island;XCH;YPXM;-10.450556;105.690278;916;7;U;Indian/Christmas
3354;Sydney Bankstown;Sydney;Australia;BWU;YSBK;-33.924444;150.988333;29;10;O;Australia/Sydney
3355;Canberra;Canberra;Australia;CBR;YSCB;-35.306944;149.195;1886;10;O;Australia/Sydney
3356;Coffs Harbour;Coff's Harbour;Australia;CFS;YSCH;-30.320556;153.116389;18;10;O;Australia/Sydney
3357;Camden;Camden;Australia;CDU;YSCN;-34.040278;150.687222;230;10;O;Australia/Sydney
3358;Dubbo;Dubbo;Australia;DBO;YSDU;-32.216667;148.574722;935;10;O;Australia/Sydney
3359;Norfolk Island Intl;Norfolk Island;Norfolk Island;NLK;YSNF;-29.041625;167.938742;371;11.5;U;Pacific/Norfolk
3360;Richmond;Richmond;Australia;RCM;YSRI;-33.600556;150.780833;67;10;O;Australia/Sydney
3361;Sydney Intl;Sydney;Australia;SYD;YSSY;-33.946111;151.177222;21;10;O;Australia/Sydney
3362;Tamworth;Tamworth;Australia;TMW;YSTW;-31.083889;150.846667;1334;10;O;Australia/Sydney
3363;Wagga Wagga;Wagga Wagga;Australia;WGA;YSWG;-35.165278;147.466389;724;10;O;Australia/Sydney
3364;Capital Intl;Beijing;China;PEK;ZBAA;40.080111;116.584556;116;8;U;Asia/Chongqing
6817;Hongyuan Airfield;Hongyuan;China;;\N;32.800428;102.534785;11500;8;N;Asia/Chongqing
3366;Dongshan;Hailar;China;HLD;ZBLA;49.204997;119.825;2169;8;U;Asia/Chongqing
3368;Binhai;Tianjin;China;TSN;ZBTJ;39.124353;117.346183;10;8;U;Asia/Chongqing
3369;Wusu;Taiyuan;China;TYN;ZBYN;37.746897;112.628428;2575;8;U;Asia/Chongqing
3370;Baiyun Intl;Guangzhou;China;CAN;ZGGG;23.392436;113.298786;50;8;U;Asia/Chongqing
3371;Huanghua;Changcha;China;CSX;ZGHA;28.189158;113.219633;217;8;U;Asia/Chongqing
3372;Liangjiang;Guilin;China;KWL;ZGKL;25.218106;110.039197;570;8;U;Asia/Chongqing
3373;Wuxu;Nanning;China;NNG;ZGNN;22.608267;108.172442;421;8;U;Asia/Chongqing
3374;Baoan Intl;Shenzhen;China;SZX;ZGSZ;22.639258;113.810664;13;8;U;Asia/Chongqing
3375;Xinzheng;Zhengzhou;China;CGO;ZHCC;34.519672;113.840889;495;8;U;Asia/Chongqing
3376;Tianhe;Wuhan;China;WUH;ZHHH;30.783758;114.2081;113;8;U;Asia/Chongqing
3377;Pyongyang Intl;Pyongyang;Korea;FNJ;ZKPY;39.224061;125.67015;117;9;U;Asia/Pyongyang
3378;Zhongchuan;Lanzhou;China;ZGC;ZLLL;36.515242;103.620775;6388;8;U;Asia/Chongqing
3379;Xianyang;Xi'an;China;XIY;ZLXY;34.447119;108.751592;1572;8;U;Asia/Chongqing
3380;Chinggis Khaan Intl;Ulan Bator;Mongolia;ULN;ZMUB;47.843056;106.766639;4364;8;U;Asia/Ulaanbaatar
3381;Gasa;Jinghonggasa;China;;ZPJH;21.973914;100.759611;1815;8;U;Asia/Chongqing
3382;Wujiaba;Kunming;China;KMG;ZPPP;24.992364;102.743536;6217;8;U;Asia/Chongqing
3383;Gaoqi;Xiamen;China;XMN;ZSAM;24.544036;118.127739;59;8;U;Asia/Chongqing
3384;Changbei Intl;Nanchang;China;KHN;ZSCN;28.865;115.9;143;8;U;Asia/Chongqing
3385;Changle;Fuzhou;China;FOC;ZSFZ;25.935064;119.663272;46;8;U;Asia/Chongqing
3386;Xiaoshan;Hangzhou;China;HGH;ZSHC;30.229503;120.434453;23;8;U;Asia/Chongqing
3387;Lishe;Ninbo;China;NGB;ZSNB;29.826683;121.461906;13;8;U;Asia/Chongqing
3388;Lukou;Nanjing;China;NKG;ZSNJ;31.742042;118.862025;49;8;U;Asia/Chongqing
3389;Luogang;Hefei;China;HFE;ZSOF;31.780019;117.298436;108;8;U;Asia/Chongqing
3390;Liuting;Qingdao;China;TAO;ZSQD;36.266108;120.374436;33;8;U;Asia/Chongqing
3391;Hongqiao Intl;Shanghai;China;SHA;ZSSS;31.197875;121.336319;10;8;U;Asia/Chongqing
3392;Laishan;Yantai;China;YNT;ZSYT;37.401667;121.371667;59;8;U;Asia/Chongqing
3393;Jiangbei;Chongqing;China;CKG;ZUCK;29.719217;106.641678;1365;8;U;Asia/Chongqing
3394;Longdongbao;Guiyang;China;KWE;ZUGY;26.538522;106.800703;3736;8;U;Asia/Chongqing
3395;Shuangliu;Chengdu;China;CTU;ZUUU;30.578528;103.947086;1625;8;U;Asia/Chongqing
3396;Qingshan;Xichang;China;XIC;ZUXC;27.989083;102.184361;5112;8;U;Asia/Chongqing
3397;Kashi;Kashi;China;KHG;ZWSH;39.542922;76.019956;4529;8;U;Asia/Chongqing
3398;Hotan;Hotan;China;HTN;ZWTN;37.038522;79.864933;4672;8;U;Asia/Chongqing
3399;Diwopu;Urumqi;China;URC;ZWWW;43.907106;87.474244;2125;8;U;Asia/Chongqing
3400;Taiping;Harbin;China;HRB;ZYHB;45.623403;126.250328;457;8;U;Asia/Chongqing
6797;Hohenems;Hohenems;Austria;HOJ;LOIH;47.385;9.7;412;1;E;Europe/Vienna
3402;Hailang;Mudanjiang;China;;ZYMD;44.524072;129.568972;883;8;U;Asia/Chongqing
3404;Zhoushuizi;Dalian;China;DLC;ZYTL;38.965667;121.5386;107;8;U;Asia/Chongqing
3406;Pudong;Shanghai;China;PVG;ZSPD;31.143378;121.805214;13;8;U;Asia/Chongqing
3407;Pulau Tioman;Tioman;Malaysia;TOD;WMBT;2.818183;104.160019;15;8;N;Asia/Kuala_Lumpur
3408;Subang-Sultan Abdul Aziz Shah Intl;Kuala Lumpur;Malaysia;SZB;WMSA;3.130583;101.549333;90;8;N;Asia/Kuala_Lumpur
3409;Noto;Wajima;Japan;NTQ;RJNW;37.293097;136.961853;718;9;U;Asia/Tokyo
3410;Borg El Arab Intl;Alexandria;Egypt;HBE;HEBA;30.917669;29.696408;177;2;U;Africa/Cairo
3411;Barter Island Lrrs;Barter Island;United States;BTI;PABA;70.133989;-143.581867;2;-9;A;America/Anchorage
3412;Wainwright As;Fort Wainwright;United States;K03;PAWT;70.613378;-159.86035;35;-9;A;America/Anchorage
3413;Cape Lisburne Lrrs;Cape Lisburne;United States;LUR;PALU;68.875133;-166.110022;12;-9;A;America/Anchorage
3414;Point Lay Lrrs;Point Lay;United States;PIZ;PPIZ;69.732875;-163.005342;25;-9;A;America/Anchorage
3415;Hilo Intl;Hilo;United States;ITO;PHTO;19.721375;-155.048469;38;-10;N;Pacific/Honolulu
3416;Executive;Orlando;United States;ORL;KORL;28.545464;-81.332936;113;-5;A;America/New_York
3417;Bettles;Bettles;United States;BTT;PABT;66.913944;-151.529056;644;-9;A;America/Anchorage
3418;Clear;Clear Mews;United States;Z84;PACL;64.301203;-149.120144;552;-9;A;America/Anchorage
3419;Indian Mountain Lrrs;Indian Mountains;United States;UTO;PAIM;65.992794;-153.704289;1220;-9;A;America/Anchorage
3420;Fort Yukon;Fort Yukon;United States;FYU;PFYU;66.571492;-145.250417;433;-9;A;America/Anchorage
3421;Sparrevohn Lrrs;Sparrevohn;United States;SVW;PASV;61.097369;-155.574228;1583;-9;A;America/Anchorage
3422;Bryant Ahp;Fort Richardson;United States;FRN;PAFR;61.266381;-149.653119;378;-9;A;America/Anchorage
3423;Tatalina Lrrs;Tatalina;United States;TLJ;PATL;62.894369;-155.976525;964;-9;A;America/Anchorage
3424;Cape Romanzof Lrrs;Cape Romanzof;United States;CZF;PACZ;61.780297;-166.038747;457;-9;A;America/Anchorage
3425;Laurence G Hanscom Fld;Bedford;United States;BED;KBED;42.469953;-71.289031;133;-5;A;America/New_York
3426;St Paul Island;St. Paul Island;United States;SNP;PASN;57.167333;-170.220444;63;-9;A;America/Anchorage
3427;Cape Newenham Lrrs;Cape Newenham;United States;EHM;PAEH;58.646428;-162.062778;541;-9;A;America/Anchorage
3428;St George;Point Barrow;United States;PBV;PAPB;56.578344;-169.661611;125;-9;A;America/Anchorage
3429;Iliamna;Iliamna;United States;ILI;PAIL;59.754356;-154.910961;186;-9;A;America/Anchorage
3430;Platinum;Port Moller;United States;PTU;PAPM;59.011356;-161.819664;15;-9;A;America/Anchorage
3431;Big Mountain Afs;Big Mountain;United States;BMX;PABM;59.361247;-155.258822;663;-9;A;America/Anchorage
3432;Oscoda Wurtsmith;Oscoda;United States;OSC;KOSC;44.451558;-83.394053;634;-5;A;America/New_York
3433;Marina Muni;Fort Ord;United States;OAR;KOAR;36.681878;-121.762347;134;-8;A;America/Los_Angeles
3434;Sacramento Mather;Sacramento;United States;MHR;KMHR;38.553897;-121.297592;96;-8;A;America/Los_Angeles
3435;Bicycle Lake Aaf;Fort Irwin;United States;BYS;KBYS;35.280531;-116.630031;2350;-8;A;America/Los_Angeles
3436;Twentynine Palms Eaf;Twenty Nine Palms;United States;NXP;KNXP;34.296161;-116.162203;2051;-8;A;America/Los_Angeles
3437;Fort Smith Rgnl;Fort Smith;United States;FSM;KFSM;35.336583;-94.367444;469;-6;A;America/Chicago
3438;Merrill Fld;Anchorage;United States;MRI;PAMR;61.213544;-149.844447;137;-9;A;America/Anchorage
3439;Grants Milan Muni;Grants;United States;GNT;KGNT;35.167286;-107.901989;6537;-7;A;America/Denver
3440;Ponca City Rgnl;Ponca City;United States;PNC;KPNC;36.731958;-97.099781;1007;-6;A;America/Chicago
3441;Hunter Aaf;Hunter Aaf;United States;SVN;KSVN;32.01;-81.145683;42;-5;A;America/New_York
3442;Grand Forks Intl;Grand Forks;United States;GFK;KGFK;47.949256;-97.176111;845;-6;A;America/Chicago
3443;Grider Fld;Pine Bluff;United States;PBF;KPBF;34.173142;-91.935597;206;-6;A;America/Chicago
3444;Whiting Fld Nas North;Milton;United States;NSE;KNSE;30.724167;-87.021944;199;-6;A;America/Chicago
3445;Hana;Hana;United States;HNM;PHHN;20.795636;-156.014439;78;-10;A;Pacific/Honolulu
3446;Ernest A Love Fld;Prescott;United States;PRC;KPRC;34.654472;-112.419583;5045;-7;A;America/Phoenix
3447;Trenton Mercer;Trenton;United States;TTN;KTTN;40.276692;-74.813469;213;-5;A;America/New_York
3448;General Edward Lawrence Logan Intl;Boston;United States;BOS;KBOS;42.364347;-71.005181;19;-5;A;America/New_York
3449;Travis Afb;Fairfield;United States;SUU;KSUU;38.262692;-121.927464;62;-8;A;America/Los_Angeles
3450;Griffiss Afld;Rome;United States;RME;KRME;43.2338;-75.407033;504;-5;A;America/New_York
3451;Wendover;Wendover;United States;ENV;KENV;40.718694;-114.030889;4237;-7;A;America/Denver
3452;Mobile Downtown;Mobile;United States;BFM;KBFM;30.626783;-88.068092;26;-6;A;America/Chicago
3453;Metropolitan Oakland Intl;Oakland;United States;OAK;KOAK;37.721278;-122.220722;9;-8;A;America/Los_Angeles
3454;Eppley Afld;Omaha;United States;OMA;KOMA;41.303167;-95.894069;984;-6;A;America/Chicago
3455;Port Angeles Cgas;Port Angeles;United States;NOW;KNOW;48.141481;-123.414075;13;-8;A;America/Los_Angeles
3456;Kahului;Kahului;United States;OGG;PHOG;20.89865;-156.430458;54;-10;N;Pacific/Honolulu
3457;Wichita Mid Continent;Wichita;United States;ICT;KICT;37.649944;-97.433056;1333;-6;A;America/Chicago
3458;Kansas City Intl;Kansas City;United States;MCI;KMCI;39.297606;-94.713905;1026;-6;A;America/Chicago
3459;Dane Co Rgnl Truax Fld;Madison;United States;MSN;KMSN;43.139858;-89.337514;887;-6;A;America/Chicago
3460;Dillingham;Dillingham;United States;DLG;PADL;59.044667;-158.5055;74;-9;A;America/Anchorage
3461;Boone Co;Harrison;United States;HRO;KHRO;36.261519;-93.154728;1365;-6;A;America/Chicago
3462;Phoenix Sky Harbor Intl;Phoenix;United States;PHX;KPHX;33.434278;-112.011583;1135;-7;N;America/Phoenix
3463;Bangor Intl;Bangor;United States;BGR;KBGR;44.807444;-68.828139;192;-5;A;America/New_York
3464;Fort Lauderdale Executive;Fort Lauderdale;United States;FXE;KFXE;26.197281;-80.170706;13;-5;A;America/New_York
3465;East Texas Rgnl;Longview;United States;GGG;KGGG;32.384014;-94.711486;365;-6;A;America/Chicago
3466;Anderson Rgnl;Andersen;United States;AND;KAND;34.494583;-82.709389;782;-5;A;America/New_York
3467;Spokane Intl;Spokane;United States;GEG;KGEG;47.619861;-117.533833;2376;-8;A;America/Los_Angeles
3468;North Perry;Hollywood;United States;HWO;KHWO;26.001222;-80.240722;8;-5;A;America/New_York
3469;San Francisco Intl;San Francisco;United States;SFO;KSFO;37.618972;-122.374889;13;-8;A;America/Los_Angeles
3470;Cut Bank Muni;Cutbank;United States;CTB;KCTB;48.608353;-112.376144;3854;-7;A;America/Denver
3471;Acadiana Rgnl;Louisiana;United States;ARA;KARA;30.037758;-91.883896;24;-6;A;America/Chicago
3472;Gainesville Rgnl;Gainesville;United States;GNV;KGNV;29.690056;-82.271778;152;-5;A;America/New_York
3473;Memphis Intl;Memphis;United States;MEM;KMEM;35.042417;-89.976667;341;-6;A;America/Chicago
3474;Bisbee Douglas Intl;Douglas;United States;DUG;KDUG;31.469028;-109.603667;4154;-7;A;America/Phoenix
3475;Allen Aaf;Delta Junction;United States;BIG;PABI;63.994547;-145.721642;1291;-9;A;America/Anchorage
3476;Tstc Waco;Waco;United States;CNW;KCNW;31.637831;-97.074139;470;-6;A;America/Chicago
3477;Annette Island;Annette Island;United States;ANN;PANT;55.042436;-131.572233;119;-9;A;America/Anchorage
3478;Caribou Muni;Caribou;United States;CAR;KCAR;46.8715;-68.017917;626;-5;A;America/New_York
3479;Little Rock Afb;Jacksonville;United States;LRF;KLRF;34.916944;-92.149722;311;-6;A;America/Chicago
3480;Redstone Aaf;Redstone;United States;HUA;KHUA;34.678653;-86.684781;685;-6;A;America/Chicago
3481;Pope Field;Fort Bragg;United States;POB;KPOB;35.170883;-79.014472;217;-5;A;America/New_York
3482;Dalhart Muni;Dalhart;United States;DHT;KDHT;36.022586;-102.547278;3991;-6;A;America/Chicago
3483;Laughlin Afb;Del Rio;United States;DLF;KDLF;29.359486;-100.777975;1082;-6;A;America/Chicago
3484;Los Angeles Intl;Los Angeles;United States;LAX;KLAX;33.942536;-118.408075;126;-8;A;America/Los_Angeles
3485;Anniston Metro;Anniston;United States;ANB;KANB;33.588167;-85.858111;612;-6;A;America/Chicago
3486;Cleveland Hopkins Intl;Cleveland;United States;CLE;KCLE;41.411689;-81.849794;791;-5;A;America/New_York
3487;Dover Afb;Dover;United States;DOV;KDOV;39.129539;-75.465958;28;-5;A;America/New_York
3488;Cincinnati Northern Kentucky Intl;Cincinnati;United States;CVG;KCVG;39.048836;-84.667822;896;-5;A;America/New_York
3489;Tipton;Fort Meade;United States;FME;KFME;39.085386;-76.759414;150;-5;A;America/New_York
3490;China Lake Naws;China;United States;NID;KNID;35.685422;-117.692039;2283;-8;A;America/Los_Angeles
3491;Huron Rgnl;Huron;United States;HON;KHON;44.3852;-98.228542;1289;-6;A;America/Chicago
3492;Juneau Intl;Juneau;United States;JNU;PAJN;58.354972;-134.576278;21;-9;A;America/Anchorage
3493;Lafayette Rgnl;Lafayette;United States;LFT;KLFT;30.205278;-91.987611;43;-6;A;America/Chicago
3494;Newark Liberty Intl;Newark;United States;EWR;KEWR;40.6925;-74.168667;18;-5;A;America/New_York
3495;Boise Air Terminal;Boise;United States;BOI;KBOI;43.564361;-116.222861;2871;-7;A;America/Denver
3496;Creech Afb;Indian Springs;United States;INS;KINS;36.587183;-115.673353;3133;-8;A;America/Los_Angeles
3497;Garden City Rgnl;Garden City;United States;GCK;KGCK;37.927528;-100.724417;2891;-6;A;America/Chicago
3498;Minot Intl;Minot;United States;MOT;KMOT;48.259378;-101.280333;1716;-6;A;America/Chicago
3499;Wheeler Aaf;Wahiawa;United States;HHI;PHHI;21.4835;-158.039667;837;-10;A;Pacific/Honolulu
3500;Maxwell Afb;Montgomery;United States;MXF;KMXF;32.382944;-86.365778;171;-6;A;America/Chicago
3501;Robinson Aaf;Robinson;United States;RBM;KRBM;34.850089;-92.300153;587;-6;A;America/Chicago
3502;Dallas Love Fld;Dallas;United States;DAL;KDAL;32.847111;-96.851778;487;-6;A;America/Chicago
3503;Butts Aaf;Fort Carson;United States;FCS;KFCS;38.678394;-104.756581;5838;-7;A;America/Denver
3504;Helena Rgnl;Helena;United States;HLN;KHLN;46.606806;-111.98275;3877;-7;A;America/Denver
3505;Miramar Mcas;Miramar;United States;NKX;KNKX;32.867694;-117.14175;478;-8;A;America/Los_Angeles
3506;Luke Afb;Phoenix;United States;LUF;KLUF;33.535;-112.38306;1085;-7;A;America/Phoenix
3507;Hurlburt Fld;Mary Esther;United States;HRT;KHRT;30.427803;-86.689278;38;-6;A;America/Chicago
3508;Jack Northrop Fld Hawthorne Muni;Hawthorne;United States;HHR;KHHR;33.922839;-118.335186;66;-8;A;America/Los_Angeles
3509;Houlton Intl;Houlton;United States;HUL;KHUL;46.123083;-67.792056;489;-5;A;America/New_York
3510;Vance Afb;Enid;United States;END;KEND;36.339167;-97.9165;1307;-6;A;America/Chicago
3511;Point Mugu Nas;Point Mugu;United States;NTD;KNTD;34.120285;-119.12094;13;-8;A;America/Los_Angeles
3512;Edwards Afb;Edwards Afb;United States;EDW;KEDW;34.905417;-117.883739;2302;-8;A;America/Los_Angeles
3513;Lake Charles Rgnl;Lake Charles;United States;LCH;KLCH;30.126112;-93.223335;15;-6;A;America/Chicago
3514;Kona Intl At Keahole;Kona;United States;KOA;PHKO;19.738767;-156.045631;47;-10;N;Pacific/Honolulu
3515;Myrtle Beach Intl;Myrtle Beach;United States;MYR;KMYR;33.67975;-78.928333;25;-5;A;America/New_York
3516;Lemoore Nas;Lemoore;United States;NLC;KNLC;36.333012;-119.95208;234;-8;A;America/Los_Angeles
3517;Nantucket Mem;Nantucket;United States;ACK;KACK;41.253053;-70.060181;48;-5;A;America/New_York
3518;Felker Aaf;Fort Eustis;United States;FAF;KFAF;37.1325;-76.608841;12;-5;A;America/New_York
3519;Campbell Aaf;Hopkinsville;United States;HOP;KHOP;36.668567;-87.496183;573;-6;A;America/Chicago
3520;Ronald Reagan Washington Natl;Washington;United States;DCA;KDCA;38.852083;-77.037722;15;-5;A;America/New_York
3521;Patuxent River Nas;Patuxent River;United States;NHK;KNHK;38.285981;-76.411781;39;-5;A;America/New_York
3522;Palacios Muni;Palacios;United States;PSX;KPSX;28.727508;-96.250958;14;-6;A;America/Chicago
3523;Arkansas Intl;Blytheville;United States;BYH;KBYH;35.964347;-89.943956;254;-6;A;America/Chicago
3524;Atlantic City Intl;Atlantic City;United States;ACY;KACY;39.457583;-74.577167;75;-5;A;America/New_York
3525;Tinker Afb;Oklahoma City;United States;TIK;KTIK;35.414739;-97.386633;1291;-6;A;America/Chicago
3526;Elizabeth City Cgas Rgnl;Elizabeth City;United States;ECG;KECG;36.260581;-76.174572;12;-5;A;America/New_York
3527;Pueblo Memorial;Pueblo;United States;PUB;KPUB;38.289085;-104.496572;4726;-7;A;America/Denver
3528;Northern Maine Rgnl At Presque Isle;Presque Isle;United States;PQI;KPQI;46.688958;-68.044797;534;-5;A;America/New_York
3529;Kirtland Air Force Base;Kirtland A.f.b.;United States;IKR;KIKR;35.040222;-106.609194;5355;-7;A;America/Denver
3530;Gray Aaf;Fort Lewis;United States;GRF;KGRF;47.079217;-122.580783;302;-8;A;America/Los_Angeles
3531;Kodiak;Kodiak;United States;ADQ;PADQ;57.749967;-152.493856;78;-9;A;America/Anchorage
3532;Upolu;Opolu;United States;UPP;PHUP;20.265256;-155.859989;96;-10;A;Pacific/Honolulu
3533;Fort Lauderdale Hollywood Intl;Fort Lauderdale;United States;FLL;KFLL;26.072583;-80.15275;9;-5;A;America/New_York
3534;Davis Fld;Muskogee;United States;MKO;KMKO;35.656489;-95.366656;612;-6;A;America/Chicago
3535;Falls Intl;International Falls;United States;INL;KINL;48.566186;-93.403067;1185;-6;A;America/Chicago
3536;Salt Lake City Intl;Salt Lake City;United States;SLC;KSLC;40.788389;-111.977772;4227;-7;A;America/Denver
3537;Childress Muni;Childress;United States;CDS;KCDS;34.433781;-100.287992;1954;-6;A;America/Chicago
3538;Keesler Afb;Biloxi;United States;BIX;KBIX;30.410425;-88.924433;33;-6;A;America/Chicago
3539;Lawson Aaf;Fort Benning;United States;LSF;KLSF;32.337322;-84.991283;232;-5;A;America/New_York
3540;Kingsville Nas;Kingsville;United States;NQI;KNQI;27.507223;-97.809723;50;-6;A;America/Chicago
3541;Marshall Aaf;Fort Riley;United States;FRI;KFRI;39.055275;-96.764453;1063;-6;A;America/Chicago
3542;Harrisburg Intl;Harrisburg;United States;MDT;KMDT;40.193494;-76.763403;310;-5;A;America/New_York
3543;Lincoln;Lincoln;United States;LNK;KLNK;40.850971;-96.75925;1219;-6;A;America/Chicago
3544;Capital City;Lansing;United States;LAN;KLAN;42.7787;-84.587357;861;-5;A;America/New_York
3545;Waimea Kohala;Kamuela;United States;MUE;PHMU;20.001328;-155.668108;2671;-10;A;Pacific/Honolulu
3546;Massena Intl Richards Fld;Massena;United States;MSS;KMSS;44.935833;-74.845547;215;-5;A;America/New_York
3547;Hickory Rgnl;Hickory;United States;HKY;KHKY;35.741147;-81.38955;1189;-5;A;America/New_York
3548;Albert Whitted;St. Petersburg;United States;SPG;KSPG;27.765111;-82.626972;7;-5;A;America/New_York
3549;Page Fld;Fort Myers;United States;FMY;KFMY;26.586611;-81.86325;17;-5;A;America/New_York
3550;George Bush Intercontinental;Houston;United States;IAH;KIAH;29.984433;-95.341442;97;-6;A;America/Chicago
3551;Millinocket Muni;Millinocket;United States;MLT;KMLT;45.647836;-68.685561;408;-5;A;America/New_York
3552;Andrews Afb;Camp Springs;United States;ADW;KADW;38.810806;-76.867028;280;-5;A;America/New_York
3553;Smith Reynolds;Winston-salem;United States;INT;KINT;36.133722;-80.222;969;-5;A;America/New_York
3554;Southern California Logistics;Victorville;United States;VCV;KVCV;34.597453;-117.382997;2885;-8;A;America/Los_Angeles
3555;Bob Sikes;Crestview;United States;CEW;KCEW;30.778833;-86.522111;213;-6;A;America/Chicago
3556;Wheeler Sack Aaf;Fort Drum;United States;GTB;KGTB;44.055619;-75.719458;690;-5;A;America/New_York
3557;St Clair Co Intl;Port Huron;United States;PHN;KPHN;42.910957;-82.528862;650;-5;A;America/New_York
3558;Meadows Fld;Bakersfield;United States;BFL;KBFL;35.433598;-119.05677;507;-8;A;America/Los_Angeles
3559;El Paso Intl;El Paso;United States;ELP;KELP;31.80725;-106.377583;3958;-7;A;America/Denver
3560;Valley Intl;Harlingen;United States;HRL;KHRL;26.2285;-97.654389;36;-6;A;America/Chicago
3561;Columbia Metropolitan;Columbia;United States;CAE;KCAE;33.938833;-81.119528;236;-5;A;America/New_York
3562;Davis Monthan Afb;Tucson;United States;DMA;KDMA;32.166467;-110.883144;2704;-7;A;America/Phoenix
3563;Pensacola Nas;Pensacola;United States;NPA;KNPA;30.352656;-87.318647;28;-6;A;America/Chicago
3564;Pensacola Rgnl;Pensacola;United States;PNS;KPNS;30.473425;-87.186611;121;-6;A;America/Chicago
3565;Grand Forks Afb;Red River;United States;RDR;KRDR;47.961098;-97.401194;913;-6;A;America/Chicago
3566;William P Hobby;Houston;United States;HOU;KHOU;29.645419;-95.278889;46;-6;A;America/Chicago
3567;Buckley Afb;Buckley;United States;BKF;KBKF;39.701668;-104.75166;5662;-7;A;America/Denver
3568;Northway;Northway;United States;ORT;PAOR;62.961334;-141.929136;1716;-9;A;America/Anchorage
3569;Palmer Muni;Palmer;United States;PAQ;PAAQ;61.594914;-149.088711;242;-9;A;America/Anchorage
3570;Pittsburgh Intl;Pittsburgh;United States;PIT;KPIT;40.491467;-80.232872;1204;-5;A;America/New_York
3571;Wiley Post Will Rogers Mem;Barrow;United States;BRW;PABR;71.285446;-156.766003;44;-9;A;America/Anchorage
3572;Ellington Fld;Houston;United States;EFD;KEFD;29.607333;-95.15875;32;-6;A;America/Chicago
3573;Whidbey Island Nas;Whidbey Island;United States;NUW;KNUW;48.351803;-122.655906;47;-8;A;America/Los_Angeles
3574;Alice Intl;Alice;United States;ALI;KALI;27.740889;-98.026944;178;-6;A;America/Chicago
3575;Moody Afb;Valdosta;United States;VAD;KVAD;30.967833;-83.193;233;-5;A;America/New_York
3576;Miami Intl;Miami;United States;MIA;KMIA;25.79325;-80.290556;8;-5;A;America/New_York
3577;Seattle Tacoma Intl;Seattle;United States;SEA;KSEA;47.449;-122.309306;433;-8;A;America/Los_Angeles
3578;Lovell Fld;Chattanooga;United States;CHA;KCHA;35.035278;-85.203808;683;-5;A;America/New_York
3579;Igor I Sikorsky Mem;Stratford;United States;BDR;KBDR;41.163472;-73.126167;9;-5;A;America/New_York
3580;Jackson Evers Intl;Jackson;United States;JAN;KJAN;32.311167;-90.075889;346;-6;A;America/Chicago
3581;Scholes Intl At Galveston;Galveston;United States;GLS;KGLS;29.265322;-94.860406;6;-6;A;America/Chicago
3582;Long Beach;Long Beach;United States;LGB;KLGB;33.817722;-118.151611;60;-8;A;America/Los_Angeles
3583;Dillingham;Dillingham;United States;HDH;PHDH;21.579475;-158.197281;14;-10;A;Pacific/Honolulu
3584;Williamsport Rgnl;Williamsport;United States;IPT;KIPT;41.241836;-76.921094;529;-5;A;America/New_York
3585;Indianapolis Intl;Indianapolis;United States;IND;KIND;39.717331;-86.294383;797;-5;A;America/New_York
3586;Whiteman Afb;Knobnoster;United States;SZL;KSZL;38.730306;-93.547864;870;-6;A;America/Chicago
3587;Akron Fulton Intl;Akron;United States;AKC;KAKR;41.0375;-81.466917;1067;-5;A;America/New_York
3588;Greenwood Leflore;Greenwood;United States;GWO;KGWO;33.494328;-90.084706;162;-6;A;America/Chicago
3589;Westchester Co;White Plains;United States;HPN;KHPN;41.066959;-73.707575;439;-5;A;America/New_York
3590;Francis S Gabreski;West Hampton Beach;United States;FOK;KFOK;40.843656;-72.631789;67;-5;A;America/New_York
3591;Jonesboro Muni;Jonesboro;United States;JBR;KJBR;35.831708;-90.646417;262;-6;A;America/Chicago
3592;Tonopah Test Range;Tonopah;United States;TNX;KTNX;37.798836;-116.78075;5549;-8;A;America/Los_Angeles
3593;Palm Beach Co Park;West Palm Beach;United States;LNA;KLNA;26.593;-80.085056;14;-5;A;America/New_York
3594;North Island Nas;San Diego;United States;NZY;KNZY;32.699219;-117.21531;26;-8;A;America/Los_Angeles
3595;Biggs Aaf;El Paso;United States;BIF;KBIF;31.849528;-106.380039;3948;-7;A;America/Denver
3596;Yuma Mcas Yuma Intl;Yuma;United States;YUM;KYUM;32.656578;-114.60598;216;-7;N;America/Phoenix
3597;Cavern City Air Terminal;Carlsbad;United States;CNM;KCNM;32.337472;-104.263278;3295;-7;A;America/Denver
3598;Duluth Intl;Duluth;United States;DLH;KDLH;46.842091;-92.193649;1428;-6;A;America/Chicago
3599;Bethel;Bethel;United States;BET;PABE;60.779778;-161.838;121;-9;A;America/Anchorage
3600;Bowman Fld;Louisville;United States;LOU;KLOU;38.228;-85.663722;546;-5;A;America/New_York
3601;Sierra Vista Muni Libby Aaf;Fort Huachuca;United States;FHU;KFHU;31.588472;-110.344389;4719;-7;A;America/Phoenix
3602;Lihue;Lihue;United States;LIH;PHLI;21.975983;-159.338958;153;-10;N;Pacific/Honolulu
3603;Terre Haute Intl Hulman Fld;Terre Haute;United States;HUF;KHUF;39.451464;-87.307561;589;-5;A;America/New_York
3604;Havre City Co;Havre;United States;HVR;KHVR;48.542983;-109.762342;2590;-7;A;America/Denver
3605;Grant Co Intl;Grant County Airport;United States;MWH;KMWH;47.207708;-119.32019;1185;-8;A;America/Los_Angeles
3606;Edward F Knapp State;Montpelier;United States;MPV;KMPV;44.203503;-72.562328;1165;-5;A;America/New_York
3607;San Nicolas Island Nolf;San Nicolas Island;United States;;KNSI;33.239765;-119.45816;506;-8;A;America/Los_Angeles
3608;Richmond Intl;Richmond;United States;RIC;KRIC;37.505167;-77.319667;167;-5;A;America/New_York
3609;Shreveport Rgnl;Shreveport;United States;SHV;KSHV;32.446629;-93.8256;258;-6;A;America/Chicago
3610;Merle K Mudhole Smith;Cordova;United States;CDV;PACV;60.491778;-145.477556;48;-9;A;America/Anchorage
3611;Norfolk Intl;Norfolk;United States;ORF;KORF;36.894611;-76.201222;26;-5;A;America/New_York
3612;Southeast Texas Rgnl;Beaumont;United States;BPT;KBPT;29.950833;-94.020694;15;-6;A;America/Chicago
3613;Savannah Hilton Head Intl;Savannah;United States;SAV;KSAV;32.127583;-81.202139;51;-5;A;America/New_York
3614;Hill Afb;Ogden;United States;HIF;KHIF;41.123939;-111.973039;4789;-7;A;America/Denver
3615;Nome;Nome;United States;OME;PAOM;64.512203;-165.445247;37;-9;A;America/Anchorage
3616;Scappoose Industrial Airpark;San Luis;United States;SPB;KSPB;45.771028;-122.861833;58;-8;A;America/Los_Angeles
3617;St Petersburg Clearwater Intl;St. Petersburg;United States;PIE;KPIE;27.910167;-82.687389;10;-5;A;America/New_York
3618;Menominee Marinette Twin Co;Macon;United States;MNM;KMNM;45.12665;-87.638443;625;-6;A;America/Chicago
3619;Lone Star Executive;Conroe;United States;CXO;KCXO;30.351833;-95.414467;245;-6;A;America/Chicago
3620;Deadhorse;Deadhorse;United States;SCC;PASC;70.19475;-148.465167;64;-9;A;America/Anchorage
3621;San Antonio Intl;San Antonio;United States;SAT;KSAT;29.533694;-98.469778;809;-6;A;America/Chicago
3622;Greater Rochester Intl;Rochester;United States;ROC;KROC;43.118866;-77.672389;559;-5;A;America/New_York
3623;Patrick Afb;Coco Beach;United States;COF;KCOF;28.234922;-80.610125;8;-5;A;America/New_York
3624;Teterboro;Teterboro;United States;TEB;KTEB;40.850103;-74.060837;9;-5;A;America/New_York
3625;Ellsworth Afb;Rapid City;United States;RCA;KRCA;44.145042;-103.103567;3279;-7;A;America/Denver
3626;Raleigh Durham Intl;Raleigh-durham;United States;RDU;KRDU;35.877639;-78.787472;435;-5;A;America/New_York
3627;James M Cox Dayton Intl;Dayton;United States;DAY;KDAY;39.902375;-84.219375;1009;-5;A;America/New_York
3628;Kenai Muni;Kenai;United States;ENA;PAEN;60.573111;-151.245;99;-9;A;America/Anchorage
3629;Mc Alester Rgnl;Mcalester;United States;MLC;KMLC;34.882403;-95.783463;770;-6;A;America/Chicago
3630;Niagara Falls Intl;Niagara Falls;United States;IAG;KIAG;43.107333;-78.946194;589;-5;A;America/New_York
3631;Coulter Fld;Bryan;United States;CFD;KCFD;30.715694;-96.331361;367;-6;A;America/Chicago
3632;Wright Aaf;Wright;United States;;KLHW;31.889097;-81.562333;45;-5;A;America/New_York
3633;Newport News Williamsburg Intl;Newport News;United States;PHF;KPHF;37.131894;-76.492989;43;-5;A;America/New_York
3634;Esler Rgnl;Alexandria;United States;ESF;KESF;31.394903;-92.295772;112;-6;A;America/Chicago
3635;Altus Afb;Altus;United States;LTS;KLTS;34.667067;-99.266681;1382;-6;A;America/Chicago
3636;Tucson Intl;Tucson;United States;TUS;KTUS;32.116083;-110.941028;2643;-7;N;America/Phoenix
3637;Minot Afb;Minot;United States;MIB;KMIB;48.415572;-101.357661;1668;-6;A;America/Chicago
3638;Beale Afb;Marysville;United States;BAB;KBAB;39.136089;-121.436567;113;-8;A;America/Los_Angeles
3639;Greater Kankakee;Kankakee;United States;IKK;KIKK;41.071389;-87.846278;630;-6;A;America/Chicago
3640;Seymour Johnson Afb;Goldsboro;United States;GSB;KGSB;35.339383;-77.960589;110;-5;A;America/New_York
3641;Theodore Francis Green State;Providence;United States;PVD;KPVD;41.732581;-71.420383;55;-5;A;America/New_York
3642;Salisbury Ocean City Wicomico Rgnl;Salisbury;United States;SBY;KSBY;38.340525;-75.510289;52;-5;A;America/New_York
3643;Rancho Murieta;Rancho Murieta;United States;RIU;KRIU;38.486778;-121.102778;141;-8;A;America/Los_Angeles
3644;Bob Hope;Burbank;United States;BUR;KBUR;34.200667;-118.358667;778;-8;A;America/Los_Angeles
3645;Detroit Metro Wayne Co;Detroit;United States;DTW;KDTW;42.212444;-83.353389;645;-5;A;America/New_York
3646;Tampa Intl;Tampa;United States;TPA;KTPA;27.975472;-82.53325;26;-5;A;America/New_York
3647;Pembina Muni;Pembina;United States;PMB;KPMB;48.942501;-97.240833;795;-6;A;America/Chicago
3648;Polk Aaf;Fort Polk;United States;POE;KPOE;31.044833;-93.191667;330;-6;A;America/Chicago
3649;Eielson Afb;Fairbanks;United States;EIL;PAEI;64.665667;-147.1015;548;-9;A;America/Anchorage
3650;Chisholm Hibbing;Hibbing;United States;HIB;KHIB;47.3866;-92.838994;1353;-6;A;America/Chicago
3651;Angelina Co;Lufkin;United States;LFK;KLFK;31.234014;-94.75;296;-6;A;America/Chicago
3652;Midland Intl;Midland;United States;MAF;KMAF;31.942528;-102.201914;2871;-6;A;America/Chicago
3653;Austin Straubel Intl;Green Bay;United States;GRB;KGRB;44.485072;-88.129589;695;-6;A;America/Chicago
3654;Ardmore Muni;Ardmore;United States;ADM;KADM;34.300833;-97.008889;762;-6;A;America/Chicago
3655;Mc Guire Afb;Wrightstown;United States;WRI;KWRI;40.015556;-74.591667;131;-5;A;America/New_York
3656;Cherry Point Mcas;Cherry Point;United States;NKT;KNKT;34.900872;-76.880733;29;-5;A;America/New_York
3657;Emanuel Co;Santa Barbara;United States;SBO;KSBO;32.609139;-82.369944;327;-5;A;America/New_York
3658;Augusta Rgnl At Bush Fld;Bush Field;United States;AGS;KAGS;33.369944;-81.9645;144;-5;A;America/New_York
3659;Sloulin Fld Intl;Williston;United States;ISN;KISN;48.177939;-103.642347;1982;-6;A;America/Chicago
3660;Adams Fld;Little Rock;United States;LIT;KLIT;34.729444;-92.224306;262;-6;A;America/Chicago
3661;Stewart Intl;Newburgh;United States;SWF;KSWF;41.504094;-74.104839;491;-5;A;America/New_York
3662;Baudette Intl;Baudette;United States;BDE;KBDE;48.728444;-94.612222;1086;-6;A;America/Chicago
3663;Sacramento Executive;Sacramento;United States;SAC;KSAC;38.512524;-121.49347;24;-8;A;America/Los_Angeles
3664;Homer;Homer;United States;HOM;PAHO;59.645556;-151.476583;84;-9;A;America/Anchorage
3665;Waynesville Rgnl Arpt At Forney Fld;Fort Leonardwood;United States;TBN;KTBN;37.741631;-92.140736;1159;-6;A;America/Chicago
3666;Dobbins Arb;Marietta;United States;MGE;KMGE;33.915382;-84.516319;1068;-5;A;America/New_York
3667;Fairchild Afb;Spokane;United States;SKA;KSKA;47.615058;-117.655772;2461;-8;A;America/Los_Angeles
3668;Roscommon Co;Houghton Lake;United States;HTL;KHTL;44.359806;-84.671112;1150;-5;A;America/New_York
3669;Tyndall Afb;Panama City;United States;PAM;KPAM;30.069567;-85.575417;17;-6;A;America/Chicago
3670;Dallas Fort Worth Intl;Dallas-Fort Worth;United States;DFW;KDFW;32.896828;-97.037997;607;-6;A;America/Chicago
3671;Melbourne Intl;Melbourne;United States;MLB;KMLB;28.102753;-80.645258;33;-5;A;America/New_York
3672;Mc Chord Afb;Tacoma;United States;TCM;KTCM;47.137678;-122.476475;322;-8;A;America/Los_Angeles
3673;Austin Bergstrom Intl;Austin;United States;AUS;KAUS;30.194528;-97.669889;542;-6;A;America/Chicago
3674;Rickenbacker Intl;Columbus;United States;LCK;KLCK;39.813786;-82.927822;744;-5;A;America/New_York
3675;Sawyer Intl;Gwinn;United States;;KSAW;46.353625;-87.395353;1221;-5;A;America/New_York
3676;Mc Ghee Tyson;Knoxville;United States;TYS;KTYS;35.810972;-83.994028;981;-5;A;America/New_York
3677;Hood Aaf;Fort Hood;United States;HLR;KHLR;31.138731;-97.714469;924;-6;A;America/Chicago
3678;Lambert St Louis Intl;St. Louis;United States;STL;KSTL;38.748697;-90.370028;618;-6;A;America/Chicago
3679;Millville Muni;Millville;United States;MIV;KMIV;39.367806;-75.072222;85;-5;A;America/New_York
3680;Sheppard Afb Wichita Falls Muni;Wichita Falls;United States;SPS;KSPS;33.988797;-98.491894;1019;-6;A;America/Chicago
3681;Cincinnati Muni Lunken Fld;Cincinnati;United States;LUK;KLUK;39.103333;-84.418611;483;-5;A;America/New_York
3682;Hartsfield Jackson Atlanta Intl;Atlanta;United States;ATL;KATL;33.636719;-84.428067;1026;-5;A;America/New_York
3683;Castle;Merced;United States;MER;KMER;37.380481;-120.568189;189;-8;A;America/Los_Angeles
3684;Mc Clellan Afld;Sacramento;United States;MCC;KMCC;38.667639;-121.400611;75;-8;A;America/Los_Angeles
3685;Gerald R Ford Intl;Grand Rapids;United States;GRR;KGRR;42.880833;-85.522806;794;-5;A;America/New_York
3686;Winkler Co;Wink;United States;INK;KINK;31.779628;-103.201361;2822;-6;A;America/Chicago
3687;Fresno Yosemite Intl;Fresno;United States;FAT;KFAT;36.776194;-119.71814;336;-8;A;America/Los_Angeles
3688;Vero Beach Muni;Vero Beach;United States;VRB;KVRB;27.655556;-80.417944;24;-5;A;America/New_York
3689;Imperial Co;Imperial;United States;IPL;KIPL;32.834219;-115.578744;-54;-8;A;America/Los_Angeles
3690;Nashville Intl;Nashville;United States;BNA;KBNA;36.124472;-86.678194;599;-6;A;America/Chicago
3691;Laredo Intl;Laredo;United States;LRD;KLRD;27.54375;-99.461556;508;-6;A;America/Chicago
3692;Elmendorf Afb;Anchorage;United States;EDF;PAED;61.250986;-149.806503;212;-9;A;America/Anchorage
3693;Ralph Wien Mem;Kotzebue;United States;OTZ;PAOT;66.884678;-162.59855;11;-9;A;America/Anchorage
3694;Altoona Blair Co;Altoona;United States;AOO;KAOO;40.296372;-78.320022;1504;-5;A;America/New_York
3695;Dyess Afb;Abilene;United States;DYS;KDYS;32.420756;-99.8546;1789;-6;A;America/Chicago
3696;South Arkansas Rgnl At Goodwin Fld;El Dorado;United States;ELD;KELD;33.220972;-92.813278;277;-6;A;America/Chicago
3697;La Guardia;New York;United States;LGA;KLGA;40.777245;-73.872608;22;-5;A;America/New_York
3698;Tallahassee Rgnl;Tallahassee;United States;TLH;KTLH;30.396528;-84.350333;81;-5;A;America/New_York
3699;Dupage;West Chicago;United States;DPA;KDPA;41.907778;-88.248611;758;-6;A;America/Chicago
3700;Waco Rgnl;Waco;United States;ACT;KACT;31.611289;-97.230519;516;-6;A;America/Chicago
3701;Augusta State;Augusta;United States;AUG;KAUG;44.320647;-69.797317;352;-5;A;America/New_York
3702;Hillsboro Muni;Hillsboro;United States;INJ;KINJ;32.083486;-97.097228;685;-6;A;America/Chicago
3703;Jacksonville Nas;Jacksonville;United States;NIP;KNIP;30.235834;-81.680556;22;-5;A;America/New_York
3704;Mc Kellar Sipes Rgnl;Jackson;United States;MKL;KMKL;35.599889;-88.915611;434;-6;A;America/Chicago
3705;Molokai;Molokai;United States;MKK;PHMK;21.152886;-157.096256;454;-10;N;Pacific/Honolulu
3706;Godman Aaf;Fort Knox;United States;FTK;KFTK;37.907058;-85.972106;756;-5;A;America/New_York
3707;New River Mcas;Jacksonville;United States;;KNCA;34.708417;-77.439667;26;-5;A;America/New_York
3708;San Angelo Rgnl Mathis Fld;San Angelo;United States;SJT;KSJT;31.35775;-100.496306;1919;-6;A;America/Chicago
3709;Calexico Intl;Calexico;United States;CXL;KCXL;32.669502;-115.51333;4;-8;A;America/Los_Angeles
3710;Chico Muni;Chico;United States;CIC;KCIC;39.795383;-121.858422;238;-8;A;America/Los_Angeles
3711;Burlington Intl;Burlington;United States;BTV;KBTV;44.471861;-73.153278;335;-5;A;America/New_York
3712;Jacksonville Intl;Jacksonville;United States;JAX;KJAX;30.494056;-81.687861;30;-5;A;America/New_York
3713;Durango La Plata Co;Durango;United States;DRO;KDRO;37.151516;-107.75377;6685;-7;A;America/Denver
3714;Washington Dulles Intl;Washington;United States;IAD;KIAD;38.944533;-77.455811;313;-5;A;America/New_York
3715;Easterwood Fld;College Station;United States;CLL;KCLL;30.588583;-96.363833;321;-6;A;America/Chicago
3716;Felts Fld;Spokane;United States;SFF;KSFF;47.682819;-117.322558;1953;-8;A;America/Los_Angeles
3717;General Mitchell Intl;Milwaukee;United States;MKE;KMKE;42.947222;-87.896583;723;-6;A;America/Chicago
3718;Abilene Rgnl;Abilene;United States;ABI;KABI;32.411319;-99.681897;1791;-6;A;America/Chicago
3719;Columbia Rgnl;Columbia;United States;COU;KCOU;38.818094;-92.219631;889;-6;A;America/Chicago
3720;Portland Intl;Portland;United States;PDX;KPDX;45.588722;-122.5975;30;-8;A;America/Los_Angeles
3721;Dade Collier Training And Transition;Miami;United States;TNT;KTNT;25.861806;-80.897;13;-5;A;America/New_York
3722;Palm Beach Intl;West Palm Beach;United States;PBI;KPBI;26.683161;-80.095589;19;-5;A;America/New_York
3723;Fort Worth Meacham Intl;Fort Worth;United States;FTW;KFTW;32.819778;-97.362444;710;-6;A;America/Chicago
3724;Ogdensburg Intl;Ogdensburg;United States;OGS;KOGS;44.681854;-75.4655;297;-5;A;America/New_York
3725;Otis Angb;Falmouth;United States;FMH;KFMH;41.658439;-70.521417;131;-5;A;America/New_York
3726;Boeing Fld King Co Intl;Seattle;United States;BFI;KBFI;47.53;-122.301947;21;-8;A;America/Los_Angeles
3727;Lackland Afb Kelly Fld Annex;San Antonio;United States;SKF;KSKF;29.384228;-98.581108;691;-6;A;America/Chicago
3728;Honolulu Intl;Honolulu;United States;HNL;PHNL;21.318681;-157.922428;13;-10;N;Pacific/Honolulu
3729;Des Moines Intl;Des Moines;United States;DSM;KDSM;41.533972;-93.663083;958;-6;A;America/Chicago
3730;Craven Co Rgnl;New Bern;United States;EWN;KEWN;35.072972;-77.042944;18;-5;A;America/New_York
3731;San Diego Intl;San Diego;United States;SAN;KSAN;32.733556;-117.189667;17;-8;A;America/Los_Angeles
3732;Monroe Rgnl;Monroe;United States;MLU;KMLU;32.510864;-92.037689;79;-6;A;America/Chicago
3733;Shaw Afb;Sumter;United States;SSC;KSSC;33.972719;-80.470564;242;-5;A;America/New_York
3734;Ontario Intl;Ontario;United States;ONT;KONT;34.056;-117.601194;944;-8;A;America/Los_Angeles
3735;Majors;Greenvile;United States;GVT;KGVT;33.067839;-96.065333;535;-6;A;America/Chicago
3736;Roswell Intl Air Center;Roswell;United States;ROW;KROW;33.301556;-104.530556;3671;-7;A;America/Denver
3737;Coleman A Young Muni;Detroit;United States;DET;KDET;42.409195;-83.009861;626;-5;A;America/New_York
3738;Brownsville South Padre Island Intl;Brownsville;United States;BRO;KBRO;25.906833;-97.425861;22;-6;A;America/Chicago
3739;Dothan Rgnl;Dothan;United States;DHN;KDHN;31.321339;-85.449628;401;-6;A;America/Chicago
3740;Cape May Co;Wildwood;United States;WWD;KWWD;39.008507;-74.908275;23;-5;A;America/New_York
3741;Fallon Nas;Fallon;United States;NFL;KNFL;39.416584;-118.70098;3934;-8;A;America/Los_Angeles
3742;Selfridge Angb;Mount Clemens;United States;MTC;KMTC;42.608333;-82.8355;580;-5;A;America/New_York
3743;Four Corners Rgnl;Farmington;United States;FMN;KFMN;36.74125;-108.229944;5506;-7;A;America/Denver
3744;Corpus Christi Intl;Corpus Christi;United States;CRP;KCRP;27.770361;-97.501222;44;-6;A;America/Chicago
3745;Syracuse Hancock Intl;Syracuse;United States;SYR;KSYR;43.111187;-76.106311;421;-5;A;America/New_York
3746;Key West Nas;Key West;United States;NQX;KNQX;24.575834;-81.688889;6;-5;A;America/New_York
3747;Chicago Midway Intl;Chicago;United States;MDW;KMDW;41.785972;-87.752417;620;-6;A;America/Chicago
3748;Norman Y Mineta San Jose Intl;San Jose;United States;SJC;KSJC;37.3626;-121.929022;62;-8;A;America/Los_Angeles
3749;Lea Co Rgnl;Hobbs;United States;HOB;KHOB;32.687528;-103.217028;3661;-7;A;America/Denver
3750;Northeast Philadelphia;Philadelphia;United States;PNE;KPNE;40.081944;-75.010586;121;-5;A;America/New_York
3751;Denver Intl;Denver;United States;DEN;KDEN;39.861656;-104.673178;5431;-7;A;America/Denver
3752;Philadelphia Intl;Philadelphia;United States;PHL;KPHL;39.871944;-75.241139;36;-5;A;America/New_York
3753;Sioux Gateway Col Bud Day Fld;Sioux City;United States;SUX;KSUX;42.402603;-96.384367;1098;-6;A;America/Chicago
3754;Middle Georgia Rgnl;Macon;United States;MCN;KMCN;32.69285;-83.649211;354;-5;A;America/New_York
3755;Truth Or Consequences Muni;Truth Or Consequences;United States;TCS;KTCS;33.236944;-107.27175;4853;-7;A;America/Denver
3756;Palmdale Rgnl Usaf Plt 42;Palmdale;United States;PMD;KPMD;34.629391;-118.08456;2543;-8;A;America/Los_Angeles
3757;Randolph Afb;San Antonio;United States;RND;KRND;29.529675;-98.2789;762;-6;A;America/Chicago
3758;El Centro Naf;El Centro;United States;NJK;KNJK;32.829222;-115.671667;-42;-8;A;America/Los_Angeles
3759;Port Columbus Intl;Columbus;United States;CMH;KCMH;39.997972;-82.891889;815;-5;A;America/New_York
3760;Drake Fld;Fayetteville;United States;FYV;KFYV;36.005094;-94.170059;1251;-6;A;America/Chicago
3761;Henry Post Aaf;Fort Sill;United States;FSI;KFSI;34.649833;-98.402167;1189;-6;A;America/Chicago
3762;Princeton Muni;Princeton;United States;PNM;KPNM;45.559868;-93.608217;979;-6;A;America/Chicago
3763;Wright Patterson Afb;Dayton;United States;FFO;KFFO;39.826111;-84.048332;823;-5;A;America/New_York
3764;Edward G Pitka Sr;Galena;United States;GAL;PAGA;64.736178;-156.937417;152;-9;A;America/Anchorage
3765;Chandler Muni;Chandler;United States;;KCHD;33.269111;-111.811111;1243;-7;A;America/Phoenix
3766;Mineral Wells;Mineral Wells;United States;MWL;KMWL;32.781606;-98.060175;974;-6;A;America/Chicago
3767;Mc Connell Afb;Wichita;United States;IAB;KIAB;37.621853;-97.268192;1371;-6;A;America/Chicago
3768;New Orleans Nas Jrb;New Orleans;United States;NBG;KNBG;29.825333;-90.035;3;-6;A;America/Chicago
3769;Beaufort;Beaufort;United States;BFT;KNBC;32.477411;-80.723161;37;-5;A;America/New_York
3770;Texarkana Rgnl Webb Fld;Texarkana;United States;TXK;KTXK;33.453722;-93.991028;390;-6;A;America/Chicago
3771;Plattsburgh Intl;Plattsburgh;United States;PBG;KPBG;44.650944;-73.468139;234;-5;A;America/New_York
3772;Phillips Aaf;Aberdeen;United States;APG;KAPG;39.466219;-76.168808;57;-5;A;America/New_York
3773;Tucumcari Muni;Tucumcari;United States;TCC;KTCC;35.182777;-103.603186;4065;-7;A;America/Denver
3774;Ted Stevens Anchorage Intl;Anchorage;United States;ANC;PANC;61.174361;-149.996361;152;-9;A;America/Anchorage
3775;Robert Gray Aaf;Killeen;United States;GRK;KGRK;31.06725;-97.828917;1015;-6;A;America/Chicago
3776;Black Rock;Zuni Pueblo;United States;ZUN;KZUN;35.083228;-108.791778;6454;-7;A;America/Denver
3777;Bellingham Intl;Bellingham;United States;BLI;KBLI;48.792694;-122.537528;170;-8;A;America/Los_Angeles
3778;Millington Rgnl Jetport;Millington;United States;NQA;KNQA;35.356667;-89.870278;320;-6;A;America/Chicago
3779;Elkins Randolph Co Jennings Randolph;Elkins;United States;EKN;KEKN;38.889444;-79.857139;1987;-5;A;America/New_York
3780;Hartford Brainard;Hartford;United States;HFD;KHFD;41.736722;-72.649444;18;-5;A;America/New_York
3781;North Central State;Smithfield;United States;SFZ;KSFZ;41.920764;-71.491381;441;-5;A;America/New_York
3782;Mobile Rgnl;Mobile;United States;MOB;KMOB;30.691231;-88.242814;219;-6;A;America/Chicago
3783;Moffett Federal Afld;Mountain View;United States;NUQ;KNUQ;37.416142;-122.049139;32;-8;A;America/Los_Angeles
3784;Santa Fe Muni;Santa Fe;United States;SAF;KSAF;35.617108;-106.089422;6348;-7;A;America/Denver
3785;Barking Sands Pmrf;Barking Sands;United States;BKH;PHBK;22.022833;-159.785;23;-10;A;Pacific/Honolulu
3786;Beauregard Rgnl;Deridder;United States;DRI;KDRI;30.831722;-93.339917;202;-6;A;America/Chicago
3787;Bradshaw Aaf;Bradshaw Field;United States;BSF;PHSF;19.760056;-155.553717;6190;-10;A;Pacific/Honolulu
3788;Nogales Intl;Nogales;United States;OLS;KOLS;31.417722;-110.84789;3955;-7;A;America/Phoenix
3789;Macdill Afb;Tampa;United States;MCF;KMCF;27.849339;-82.521214;14;-5;A;America/New_York
3790;Scott Afb Midamerica;Belleville;United States;BLV;KBLV;38.545178;-89.835183;459;-6;A;America/Chicago
3791;Opa Locka;Miami;United States;OPF;KOPF;25.907;-80.278389;8;-5;A;America/New_York
3792;Del Rio Intl;Del Rio;United States;DRT;KDRT;29.374208;-100.927158;1002;-6;A;America/Chicago
3793;Southwest Florida Intl;Fort Myers;United States;RSW;KRSW;26.536167;-81.755167;30;-5;A;America/New_York
3794;King Salmon;King Salmon;United States;AKN;PAKN;58.676778;-156.649278;68;-9;A;America/Anchorage
3795;Muir Aaf;Muir;United States;MUI;KMUI;40.434811;-76.569411;489;-5;A;America/New_York
3796;Kapalua;Lahania-kapalua;United States;JHM;PHJH;20.962936;-156.673031;256;-10;N;Pacific/Honolulu
3797;John F Kennedy Intl;New York;United States;JFK;KJFK;40.639751;-73.778925;13;-5;A;America/New_York
3798;Homestead Arb;Homestead;United States;HST;KHST;25.48855;-80.383567;6;-5;A;America/New_York
3799;Riverside Muni;Riverside;United States;RAL;KRAL;33.951875;-117.445103;818;-8;A;America/Los_Angeles
3800;Sherman Aaf;Fort Leavenworth;United States;FLV;KFLV;39.368332;-94.914686;772;-6;A;America/Chicago
3801;Wallops Flight Facility;Wallops Island;United States;WAL;KWAL;37.940194;-75.466389;40;-5;A;America/New_York
3802;Holloman Afb;Alamogordo;United States;HMN;KHMN;32.852519;-106.106525;4093;-7;A;America/Denver
3803;Willow Grove Nas Jrb;Willow Grove;United States;NXX;KNXX;40.199833;-75.148167;358;-5;A;America/New_York
3804;Cheyenne Rgnl Jerry Olson Fld;Cheyenne;United States;CYS;KCYS;41.155722;-104.811839;6156;-7;A;America/Denver
3805;Stockton Metropolitan;Stockton;United States;SCK;KSCK;37.894167;-121.238306;33;-8;A;America/Los_Angeles
3806;Charleston Afb Intl;Charleston;United States;CHS;KCHS;32.898647;-80.040528;45;-5;A;America/New_York
3807;Reno Tahoe Intl;Reno;United States;RNO;KRNO;39.499108;-119.768108;4415;-8;A;America/Los_Angeles
3808;Ketchikan Intl;Ketchikan;United States;KTN;PAKT;55.355556;-131.71375;88;-9;A;America/Anchorage
3809;Willow Run;Detroit;United States;YIP;KYIP;42.237928;-83.530408;716;-5;A;America/New_York
3810;Vandenberg Afb;Lompoc;United States;VBG;KVBG;34.737333;-120.584306;369;-8;A;America/Los_Angeles
3811;Birmingham Intl;Birmingham;United States;BHM;KBHM;33.562942;-86.75355;644;-6;A;America/Chicago
3812;Lakehurst Naes;Lakehurst;United States;NEL;KNEL;40.033333;-74.353333;103;-5;A;America/New_York
3813;Eareckson As;Shemya;United States;SYA;PASY;52.712275;174.11362;98;-9;A;America/Anchorage
3814;Nellis Afb;Las Vegas;United States;LSV;KLSV;36.236197;-115.034253;1870;-8;A;America/Los_Angeles
3815;March Arb;Riverside;United States;RIV;KRIV;33.880711;-117.259453;1535;-8;A;America/Los_Angeles
3816;Modesto City Co Harry Sham;Modesto;United States;MOD;KMOD;37.625817;-120.954422;97;-8;A;America/Los_Angeles
3817;Sacramento Intl;Sacramento;United States;SMF;KSMF;38.695417;-121.590778;27;-8;A;America/Los_Angeles
3818;Waukegan Rgnl;Chicago;United States;UGN;KUGN;42.422161;-87.867908;727;-6;A;America/Chicago
3819;City Of Colorado Springs Muni;Colorado Springs;United States;COS;KCOS;38.805805;-104.700778;6187;-7;A;America/Denver
3820;Buffalo Niagara Intl;Buffalo;United States;BUF;KBUF;42.940525;-78.732167;724;-5;A;America/New_York
3821;Griffing Sandusky;Sandusky;United States;SKY;KSKY;41.433361;-82.652333;580;-5;A;America/New_York
3822;Snohomish Co;Everett;United States;PAE;KPAE;47.906342;-122.281564;606;-8;A;America/Los_Angeles
3823;Mountain Home Afb;Mountain Home;United States;MUO;KMUO;43.043603;-115.872431;2996;-7;A;America/Denver
3824;Cedar City Rgnl;Cedar City;United States;CDC;KCDC;37.700967;-113.098847;5622;-7;A;America/Denver
3825;Bradley Intl;Windsor Locks;United States;BDL;KBDL;41.938889;-72.683222;173;-5;A;America/New_York
3826;Mc Allen Miller Intl;Mcallen;United States;MFE;KMFE;26.175833;-98.238611;107;-6;A;America/Chicago
3827;Norfolk Ns;Norfolk;United States;NGU;KNGU;36.937644;-76.289289;15;-5;A;America/New_York
3828;Westover Arb Metropolitan;Chicopee Falls;United States;CEF;KCEF;42.194014;-72.534783;241;-5;A;America/New_York
3829;Lubbock Preston Smith Intl;Lubbock;United States;LBB;KLBB;33.663639;-101.822778;3282;-6;A;America/Chicago
3830;Chicago Ohare Intl;Chicago;United States;ORD;KORD;41.978603;-87.904842;668;-6;A;America/Chicago
3831;Boca Raton;Boca Raton;United States;BCT;KBCT;26.3785;-80.107694;13;-5;A;America/New_York
3832;Fairbanks Intl;Fairbanks;United States;FAI;PAFA;64.815114;-147.856267;434;-9;A;America/Anchorage
3833;Quantico Mcaf;Quantico;United States;NYG;KNYG;38.501683;-77.305333;11;-5;A;America/New_York
3834;Cannon Afb;Clovis;United States;CVS;KCVS;34.382775;-103.322147;4295;-7;A;America/Denver
3835;Kaneohe Bay Mcaf;Kaneohe Bay;United States;NGF;PHNG;21.450453;-157.768;24;-10;A;Pacific/Honolulu
3836;Offutt Afb;Omaha;United States;OFF;KOFF;41.118332;-95.912511;1052;-6;A;America/Chicago
3837;Gulkana;Gulkana;United States;GKN;PAGK;62.154888;-145.456639;1580;-9;A;America/Anchorage
3838;Watertown Intl;Watertown;United States;ART;KART;43.991922;-76.021739;325;-5;A;America/New_York
3839;Palm Springs Intl;Palm Springs;United States;PSP;KPSP;33.829667;-116.506694;477;-8;A;America/Los_Angeles
3840;Rick Husband Amarillo Intl;Amarillo;United States;AMA;KAMA;35.219369;-101.705931;3607;-6;A;America/Chicago
3841;Fort Dodge Rgnl;Fort Dodge;United States;FOD;KFOD;42.5512;-94.191842;1157;-6;A;America/Chicago
3842;Barksdale Afb;Shreveport;United States;BAD;KBAD;32.50182;-93.662674;166;-6;A;America/Chicago
3843;Forbes Fld;Topeka;United States;FOE;KFOE;38.950944;-95.663611;1078;-6;A;America/Chicago
3844;Cotulla Lasalle Co;Cotulla;United States;COT;KCOT;28.456694;-99.220294;474;-6;A;America/Chicago
3845;Wilmington Intl;Wilmington;United States;ILM;KILM;34.270615;-77.902569;32;-5;A;America/New_York
3846;Baton Rouge Metro Ryan Fld;Baton Rouge;United States;BTR;KBTR;30.533167;-91.149639;70;-6;A;America/Chicago
3847;Meridian Nas;Meridian;United States;NMM;KNMM;32.552083;-88.555557;317;-6;A;America/Chicago
3848;Tyler Pounds Rgnl;Tyler;United States;TYR;KTYR;32.354139;-95.402386;544;-6;A;America/Chicago
3849;Baltimore Washington Intl;Baltimore;United States;BWI;KBWI;39.175361;-76.668333;146;-5;A;America/New_York
3850;Hobart Muni;Hobart;United States;HBR;KHBR;34.991308;-99.051353;1564;-6;A;America/Chicago
3851;Lanai;Lanai;United States;LNY;PHNY;20.785611;-156.951419;1308;-10;N;Pacific/Honolulu
3852;Alexandria Intl;Alexandria;United States;AEX;KAEX;31.3274;-92.549833;89;-6;A;America/Chicago
3853;Condron Aaf;White Sands;United States;WSD;KWSD;32.341484;-106.40277;3934;-7;A;America/Denver
3854;Cold Bay;Cold Bay;United States;CDB;PACD;55.206061;-162.725436;96;-9;A;America/Anchorage
3855;Tulsa Intl;Tulsa;United States;TUL;KTUL;36.198389;-95.888111;677;-6;A;America/Chicago
3856;Sitka Rocky Gutierrez;Sitka;United States;SIT;PASI;57.047138;-135.361611;21;-9;A;America/Anchorage
3857;Long Island Mac Arthur;Islip;United States;ISP;KISP;40.79525;-73.100222;99;-5;A;America/New_York
3858;Minneapolis St Paul Intl;Minneapolis;United States;MSP;KMSP;44.881956;-93.221767;841;-6;A;America/Chicago
3859;New Castle;Wilmington;United States;ILG;KILG;39.678722;-75.606528;79;-5;A;America/New_York
3860;Unalaska;Unalaska;United States;DUT;PADU;53.900139;-166.5435;22;-9;A;America/Anchorage
3861;Louis Armstrong New Orleans Intl;New Orleans;United States;MSY;KMSY;29.993389;-90.258028;4;-6;A;America/Chicago
3862;Portland Intl Jetport;Portland;United States;PWM;KPWM;43.646161;-70.309281;77;-5;A;America/New_York
3863;Will Rogers World;Oklahoma City;United States;OKC;KOKC;35.393089;-97.600733;1295;-6;A;America/Chicago
3864;Albany Intl;Albany;United States;ALB;KALB;42.748267;-73.801692;285;-5;A;America/New_York
3865;Valdez Pioneer Fld;Valdez;United States;VDZ;PAVD;61.133949;-146.248342;118;-9;A;America/Anchorage
3866;Langley Afb;Hampton;United States;LFI;KLFI;37.082881;-76.360547;11;-5;A;America/New_York
3867;John Wayne Arpt Orange Co;Santa Ana;United States;SNA;KSNA;33.675667;-117.868222;56;-8;A;America/Los_Angeles
3868;Columbus Afb;Colombus;United States;CBM;KCBM;33.643833;-88.443833;219;-6;A;America/Chicago
3869;Kendall Tamiami Executive;Kendall-tamiami;United States;TMB;KTMB;25.647889;-80.432777;8;-5;A;America/New_York
3870;Oceana Nas;Oceana;United States;NTU;KNTU;36.820703;-76.033542;22;-5;A;America/New_York
3871;Grissom Arb;Peru;United States;GUS;KGUS;40.648094;-86.152119;812;-5;A;America/New_York
3872;Natrona Co Intl;Casper;United States;CPR;KCPR;42.908;-106.464417;5347;-7;A;America/Denver
3873;Eglin Afb;Valparaiso;United States;VPS;KVPS;30.48325;-86.5254;87;-6;A;America/Chicago
3874;Craig Fld;Selma;United States;SEM;KSEM;32.343947;-86.987803;166;-6;A;America/Chicago
3875;Key West Intl;Key West;United States;EYW;KEYW;24.556111;-81.759556;3;-5;A;America/New_York
3876;Charlotte Douglas Intl;Charlotte;United States;CLT;KCLT;35.214;-80.943139;748;-5;A;America/New_York
3877;Mc Carran Intl;Las Vegas;United States;LAS;KLAS;36.080056;-115.15225;2141;-8;A;America/Los_Angeles
3878;Orlando Intl;Orlando;United States;MCO;KMCO;28.429394;-81.308994;96;-5;A;America/New_York
3879;Florence Rgnl;Florence;United States;FLO;KFLO;34.185361;-79.723889;146;-5;A;America/New_York
3880;Great Falls Intl;Great Falls;United States;GTF;KGTF;47.482;-111.370689;3677;-7;A;America/Denver
3881;Youngstown Warren Rgnl;Youngstown;United States;YNG;KYNG;41.260736;-80.679097;1196;-5;A;America/New_York
3882;Ladd Aaf;Fort Wainwright;United States;FBK;PAFB;64.8375;-147.614444;454;-9;A;America/Anchorage
3883;Mc Minnville Muni;Mackminnville;United States;MMV;KMMV;45.194444;-123.135944;163;-8;A;America/Los_Angeles
3884;Robins Afb;Macon;United States;WRB;KWRB;32.640144;-83.59185;294;-5;A;America/New_York
3885;Suvarnabhumi Intl;Bangkok;Thailand;BKK;VTBS;13.681108;100.747283;5;7;U;Asia/Bangkok
3887;Andi Jemma;Masamba;Indonesia;;WAWM;-2.558044;120.324383;164;8;N;Asia/Makassar
3888;Soroako;Soroako;Indonesia;;WAWS;-2.531203;121.357639;1388;8;N;Asia/Makassar
3889;Pongtiku;Makale;Indonesia;;WAWT;-3.044736;119.821536;2884;8;N;Asia/Makassar
3890;Wolter Monginsidi;Kendari;Indonesia;KDI;WAWW;-4.081608;122.418231;538;8;N;Asia/Makassar
3891;Maimun Saleh;Sabang;Indonesia;SBG;WITB;5.874131;95.339672;393;7;N;Asia/Jakarta
3892;Cibeureum;Tasikmalaya;Indonesia;;WICM;-7.346603;108.246092;1148;7;N;Asia/Jakarta
3893;Iswahyudi;Madiun;Indonesia;;WARI;-7.615767;111.434117;361;7;N;Asia/Jakarta
3894;Abdul Rachman Saleh;Malang;Indonesia;MLG;WARA;-7.926556;112.714514;1726;7;N;Asia/Jakarta
3895;Budiarto;Tangerang;Indonesia;;WICB;-6.293169;106.5699;151;7;N;Asia/Jakarta
3896;Husein Sastranegara;Bandung;Indonesia;BDO;WICC;-6.900625;107.576294;2436;7;N;Asia/Jakarta
3897;Penggung;Cirebon;Indonesia;CBN;WICD;-6.756144;108.539672;89;7;N;Asia/Jakarta
3898;Adi Sutjipto;Yogyakarta;Indonesia;JOG;WARJ;-7.788181;110.431758;350;7;N;Asia/Jakarta
3899;Tunggul Wulung;Cilacap;Indonesia;CXP;WIHL;-7.645056;109.033911;69;7;N;Asia/Jakarta
3900;Pondok Cabe;Jakarta;Indonesia;PCB;WIHP;-6.336964;106.764561;200;7;N;Asia/Jakarta
3901;Achmad Yani;Semarang;Indonesia;SRG;WARS;-6.971447;110.374122;13;7;N;Asia/Jakarta
3903;Hang Nadim;Batam;Indonesia;BTH;WIDD;1.121028;104.118753;126;7;N;Asia/Jakarta
3904;H As Hanandjoeddin;Tanjung Pandan;Indonesia;TJQ;WIOD;-2.745722;107.754917;164;7;N;Asia/Jakarta
3905;Depati Amir;Pangkal Pinang;Indonesia;PGK;WIPK;-2.1622;106.139064;109;7;N;Asia/Jakarta
3906;Kijang;Tanjung Pinang;Indonesia;TNJ;WIDN;0.922683;104.532311;52;7;N;Asia/Jakarta
3907;Dabo;Singkep;Indonesia;SIQ;WIDS;-0.479189;104.579283;95;7;N;Asia/Jakarta
3908;Syamsudin Noor;Banjarmasin;Indonesia;BDJ;WAOO;-3.442356;114.762553;66;8;N;Asia/Makassar
3909;Batu Licin;Batu Licin;Indonesia;;WAOC;-3.412408;115.995136;3;8;N;Asia/Makassar
3910;Iskandar;Pangkalan Bun;Indonesia;PKN;WAOI;-2.705197;111.673208;75;7;N;Asia/Jakarta
3911;Tjilik Riwut;Palangkaraya;Indonesia;PKY;WAOP;-2.225128;113.942661;82;7;N;Asia/Jakarta
6822;Militaerlager;S-Chanf;Switzerland;;\N;46.6166;9.9833;5100;1;U;Europe/Zurich
3913;Wai Oti;Maumere;Indonesia;MOF;WATC;-8.640647;122.236889;115;8;N;Asia/Makassar
3914;H Hasan Aroeboesman;Ende;Indonesia;ENE;WATE;-8.849294;121.660644;49;8;N;Asia/Makassar
3915;Satar Tacik;Ruteng;Indonesia;RTG;WATG;-8.597011;120.477061;3510;8;N;Asia/Makassar
3916;El Tari;Kupang;Indonesia;KOE;WATT;-10.171583;123.671136;335;8;N;Asia/Makassar
3917;Mutiara Ii;Labuhan Bajo;Indonesia;LBJ;WATO;-8.486656;119.88905;66;8;N;Asia/Makassar
3919;Sepinggan;Balikpapan;Indonesia;BPN;WALL;-1.268272;116.894478;12;8;N;Asia/Makassar
3920;Juwata;Taraken;Indonesia;TRK;WALR;3.326694;117.565569;20;8;N;Asia/Makassar
3921;Temindung;Samarinda;Indonesia;SRI;WALS;-0.484531;117.157111;33;8;N;Asia/Makassar
3922;Tanjung Santan;Tanjung Santan;Indonesia;;WALT;-0.093215;117.439292;121;8;N;Asia/Makassar
3923;Selaparang;Mataram;Indonesia;AMI;WADA;-8.560708;116.094656;52;8;N;Asia/Makassar
3924;Muhammad Salahuddin;Bima;Indonesia;BMU;WADB;-8.539647;118.687322;3;8;N;Asia/Makassar
6799;Krems Langenlois;Krems;Austria;;LOAG;48.446075;15.631243;1022;1;E;Europe/Vienna
3927;Mau Hau;Waingapu;Indonesia;WGP;WADW;-9.669217;120.302006;33;8;N;Asia/Makassar
3928;Juanda;Surabaya;Indonesia;SUB;WARR;-7.379831;112.786858;9;7;N;Asia/Jakarta
3929;Adi Sumarmo Wiryokusumo;Solo City;Indonesia;SOC;WARQ;-7.516089;110.756892;421;7;N;Asia/Jakarta
3931;Chiang Mai Intl;Chiang Mai;Thailand;CNX;VTCC;18.766847;98.962644;1036;7;U;Asia/Bangkok
3932;Chiang Rai Intl;Chiang Rai;Thailand;CEI;VTCT;19.952342;99.882928;1280;7;U;Asia/Bangkok
3933;Nakhon Si Thammarat;Nakhon Si Thammarat;Thailand;NST;VTSF;8.539617;99.944725;13;7;U;Asia/Bangkok
3940;Bali Ngurah Rai;Denpasar;Indonesia;DPS;WADD;-8.748169;115.167172;14;8;N;Asia/Makassar
3935;Nakhon Ratchasima;Nakhon Ratchasima;Thailand;NAK;VTUQ;14.949497;102.312736;765;7;U;Asia/Bangkok
3936;Nakhon Phanom;Nakhon Phanom;Thailand;KOP;VTUW;17.383794;104.643022;587;7;U;Asia/Bangkok
3937;Ubon Ratchathani;Ubon Ratchathani;Thailand;UBP;VTUU;15.251278;104.870231;406;7;U;Asia/Bangkok
3938;Khon Kaen;Khon Kaen;Thailand;KKC;VTUK;16.466628;102.783661;670;7;U;Asia/Bangkok
3939;Sukhothai;Sukhothai;Thailand;THS;VTPO;17.237992;99.818183;179;7;U;Asia/Bangkok
3941;Eleftherios Venizelos Intl;Athens;Greece;ATH;LGAV;37.936358;23.944467;308;2;E;Europe/Athens
3942;Chubu Centrair Intl;Nagoya;Japan;NGO;RJGG;34.858414;136.805408;15;9;U;Asia/Tokyo
3943;Kobe;Kobe;Japan;UKB;RJBE;34.632778;135.223889;22;9;U;Asia/Tokyo
3944;Pullman-Moscow Rgnl;Pullman;United States;PUW;KPUW;46.743861;-117.109583;2556;-8;A;America/Los_Angeles
3945;Lewiston Nez Perce Co;Lewiston;United States;LWS;KLWS;46.3745;-117.015389;1442;-8;A;America/Los_Angeles
3946;Elmira Corning Rgnl;Elmira;United States;ELM;KELM;42.159889;-76.891611;954;-5;A;America/New_York
3947;Ithaca Tompkins Rgnl;Ithaca;United States;ITH;KITH;42.491028;-76.458444;1099;-5;A;America/New_York
3948;Monterey Peninsula;Monterey;United States;MRY;KMRY;36.587;-121.842944;257;-8;A;America/Los_Angeles
3949;Santa Barbara Muni;Santa Barbara;United States;SBA;KSBA;34.426211;-119.840372;10;-8;A;America/Los_Angeles
3950;Daytona Beach Intl;Daytona Beach;United States;DAB;KDAB;29.179917;-81.058056;34;-5;A;America/New_York
8421;Taoyuan Station;Zhongli;Taiwan;;\N;25.013092;121.215216;250;8;N;Asia/Taipei
3952;Liepaja Intl;Liepaja;Latvia;LPX;EVLA;56.5175;21.096944;16;2;E;Europe/Riga
3953;Riga Intl;Riga;Latvia;RIX;EVRA;56.923611;23.971111;34;2;E;Europe/Riga
3954;Siauliai Intl;Siauliai;Lithuania;SQQ;EYSA;55.893886;23.394975;443;2;E;Europe/Vilnius
3955;Barysiai;Barysiai;Lithuania;HLJ;EYSB;56.070556;23.558056;270;2;E;Europe/Vilnius
3956;Kaunas Intl;Kaunas;Lithuania;KUN;EYKA;54.963919;24.084778;256;2;E;Europe/Vilnius
3957;S. Darius;Kaunas;Lithuania;;EYKS;54.879792;23.881511;246;2;E;Europe/Vilnius
3958;Palanga Intl;Palanga;Lithuania;PLQ;EYPA;55.973228;21.093856;33;2;E;Europe/Vilnius
3959;Vilnius Intl;Vilnius;Lithuania;VNO;EYVI;54.634133;25.285767;646;2;E;Europe/Vilnius
3960;Pajuostis;Panevezys;Lithuania;PNV;EYPP;55.729444;24.460833;197;2;E;Europe/Vilnius
3962;Erebuni;Yerevan;Armenia;;UDYE;40.122114;44.464992;2948;4;E;Asia/Yerevan
3963;Stepanavan;Stepanavan;Armenia;;UDLS;41.04845;44.337172;4836;4;E;Asia/Yerevan
3964;Zvartnots;Yerevan;Armenia;EVN;UDYZ;40.147275;44.395881;2838;4;E;Asia/Yerevan
3965;Gyumri;Gyumri;Armenia;LWN;UDSG;40.750369;43.859342;5000;4;E;Asia/Yerevan
3966;Assab Intl;Assab;Eritrea;ASA;HHSB;13.071783;42.645006;46;3;U;Africa/Asmera
3967;Asmara Intl;Asmara;Eritrea;ASM;HHAS;15.291853;38.910667;7661;3;U;Africa/Asmera
3968;Massawa Intl;Massawa;Eritrea;MSW;HHMS;15.669989;39.370103;194;3;U;Africa/Asmera
3969;Yasser Arafat Intl;Gaza;Palestine;GZA;LVGZ;31.246389;34.276111;320;2;U;Asia/Gaza
3974;Riyan;Mukalla;Yemen;RIY;OYRN;14.662639;49.375028;54;3;U;Asia/Aden
3971;Batumi;Batumi;Georgia;BUS;UGSB;41.610278;41.599694;105;4;N;Asia/Tbilisi
3972;Kopitnari;Kutaisi;Georgia;KUT;UGKO;42.176653;42.482583;223;4;N;Asia/Tbilisi
3973;Tbilisi;Tbilisi;Georgia;TBS;UGTB;41.669167;44.954722;1624;4;N;Asia/Tbilisi
3975;Taiz Intl;Taiz;Yemen;TAI;OYTZ;13.685964;44.139056;4838;3;U;Asia/Aden
3976;Hodeidah Intl;Hodeidah;Yemen;HOD;OYHD;14.753;42.976336;41;3;U;Asia/Aden
3977;Aden Intl;Aden;Yemen;ADE;OYAA;12.829542;45.028792;7;3;U;Asia/Aden
3978;Ataq;Ataq;Yemen;AXK;OYAT;14.551322;46.826183;3735;3;U;Asia/Aden
3979;Al Ghaidah Intl;Al Ghaidah Intl;Yemen;AAY;OYGD;16.191667;52.175;134;3;U;Asia/Aden
3980;Sanaa Intl;Sanaa;Yemen;SAH;OYSN;15.476258;44.219739;7216;3;U;Asia/Aden
3986;Allgau;Memmingen;Germany;FMM;EDJA;47.988758;10.2395;2077;1;E;Europe/Berlin
3982;Beihan;Beihan;Yemen;BHN;OYBN;14.781972;45.720083;3800;3;U;Asia/Aden
6863;RK Heliplex;Panorama;Canada;;\N;50.2735;-116.14;358;-7;A;America/Edmonton
3984;Socotra Intl;Socotra;Yemen;SCT;OYSQ;12.630672;53.905778;146;3;U;Asia/Aden
3985;Al Badie;Al Badie;Yemen;;OYBA;18.719311;50.836911;908;3;U;Asia/Aden
3987;Kapadokya;Nevsehir;Turkey;NAV;LTAZ;38.771867;34.53455;3100;2;E;Europe/Istanbul
3988;Ministro Pistarini;Buenos Aires;Argentina;EZE;SAEZ;-34.822222;-58.535833;67;-3;N;America/Cordoba
3989;Erbil Intl;Erbil;Iraq;EBL;ORER;36.237611;43.963158;1341;3;U;Asia/Baghdad
3990;Emerald;Emerald;Australia;EMD;YEML;-23.5675;148.179167;624;10;O;Australia/Brisbane
3991;Ellinikon International Airport;Athens;Greece;;LGAT;37.5354;23.4346;68;2;E;Europe/Athens
3992;Kansai;Osaka;Japan;KIX;RJBB;34.4347222;135.244167;49;9;U;Asia/Tokyo
3993;Wall Street Heliport;New York;United States;JRB;KJRB;40.701214;-74.009028;7;-5;A;America/New_York
3994;Tagbilaran;Tagbilaran;Philippines;TAG;RPVT;9.66408056;123.853247;38;8;N;Asia/Manila
3995;Ilulissat;Ilulissat;Greenland;JAV;BGJN;69.23444;-51.05111;2;-3;E;America/Godthab
3996;Qasigiannguit;Qasigiannguit;Greenland;JCH;BGCH;68.833336;-51;2;-3;E;America/Godthab
3997;Aasiaat;Aasiaat;Greenland;JEG;BGEM;68.7;-52.75;2;-3;E;America/Godthab
3998;Son Sant Joan;Palma de Mallorca;Spain;PMI;LEPA;39.55361;2.727778;24;1;E;Europe/Madrid
3999;Darwin Intl;Darwin;Australia;DRW;YPDN;-12.4083333;130.87266;103;9.5;N;Australia/Darwin
4000;Surat Thani;Surat Thani;Thailand;URT;VTSB;9.1325;99.135556;286;7;U;Asia/Bangkok
4001;Bagan Intl;Nyuang U;Burma;NYU;VYBR;21.173833266;94.9246666;290;6.5;U;Asia/Rangoon
4002;Godofredo P;Caticlan;Philippines;MPH;RPXE;11.9214999;121.953;20;8;N;Asia/Manila
4003;Bimini North Seaplane Base;Bimini;Bahamas;NSB;\N;25.767;-79.25;0;-5;U;America/Nassau
4004;Talkeetna;Talkeetna;United States;TKA;PATK;62.3205;-150.093694;358;-9;A;America/Anchorage
4005;Xewkija Heliport;Gozo;Malta;GZM;LMMG;36.027222;14.272778;300;1;E;Europe/Malta
4006;Tweed-New Haven Airport;New Haven;United States;HVN;KHVN;41.26375;-72.886806;14;-5;A;America/New_York
4007;Asheville Regional Airport;Asheville;United States;AVL;KAVL;35.436194;-82.541806;2165;-5;A;America/New_York
4008;Piedmont Triad;Greensboro;United States;GSO;KGSO;36.09775;-79.937306;925;-5;A;America/New_York
4009;Sioux Falls;Sioux Falls;United States;FSD;KFSD;43.582014;-96.741914;1429;-6;A;America/Chicago
4010;Ayers Rock;Uluru;Australia;AYQ;YAYE;-25.186111;130.975556;1626;9.5;N;Australia/Darwin
4011;Manchester Regional Airport;Manchester NH;United States;MHT;KMHT;42.932556;-71.435667;266;-5;A;America/New_York
4012;Naples Muni;Naples;United States;APF;KAPF;26.152619;-81.775294;8;-5;A;America/New_York
4013;Redang;Redang;Malaysia;RDN;WMPR;5.76528;103.007;10;8;N;Asia/Kuala_Lumpur
4014;Louisville International Airport;Louisville;United States;SDF;KSDF;38.1740858;-85.7364989;501;-5;A;America/New_York
4015;Charlottesville-Albemarle;Charlottesville VA;United States;CHO;KCHO;38.138639;-78.452861;639;-5;A;America/New_York
4016;Roanoke Regional;Roanoke VA;United States;ROA;KROA;37.325472;-79.975417;1175;-5;A;America/New_York
4017;Blue Grass;Lexington KY;United States;LEX;KLEX;38.0365;-84.605889;979;-5;A;America/New_York
4018;Evansville Regional;Evansville;United States;EVV;KEVV;38.036997;-87.532364;418;-6;A;America/Chicago
4019;Albuquerque International Sunport;Albuquerque;United States;ABQ;KABQ;35.0402222;-106.6091944;5355;-7;A;America/Denver
4020;Gallatin Field;Bozeman;United States;BZN;KBZN;45.777643;-111.160151;4500;-7;A;America/Denver
4021;Billings Logan International Airport;Billings;United States;BIL;KBIL;45.80921;-108.537654;3652;-7;A;America/Denver
4022;Bert Mooney Airport;Butte;United States;BTM;KBTM;45.954806;-112.497472;5550;-7;A;America/Denver
4023;Cherry Capital Airport;Traverse City;United States;TVC;KTVC;44.741445;-85.582235;624;-5;A;America/New_York
4024;Mundo Maya International;Flores;Guatemala;FRS;MGTK;16.913819;-89.866383;427;-6;U;America/Guatemala
4025;Hancock County - Bar Harbor;Bar Harbor;United States;BHB;KBHB;44.4497689;-68.3615653;83;-5;A;America/New_York
4026;Knox County Regional Airport;Rockland;United States;RKD;KRKD;44.0601111;-69.0992303;56;-5;A;America/New_York
4027;Jackson Hole Airport;Jacksn Hole;United States;JAC;KJAC;43.607333333;-110.73775;6451;-7;A;America/Denver
4028;Chicago Rockford International Airport ;Rockford;United States;RFD;KRFD;42.1953611;-89.0972222;742;-6;A;America/Chicago
4029;Domododevo;Moscow;Russia;DME;UUDD;55.408611;37.906111;588;4;N;Europe/Moscow
4030;Phoenix International;Sanya;China;SYX;ZJSY;18.302897;109.412272;92;8;U;Asia/Chongqing
4031;Milford Sound Airport;Milford Sound;New Zealand;MFN;NZMF;-44.67333;167.92333;10;12;Z;Pacific/Auckland
4032;East 34th Street Heliport;New York;United States;TSS;NONE;40.7425;-73.971944;10;-5;A;America/New_York
4033;Lijiang Airport;Lijiang;China;LJG;ZPLJ;26.883333;100.23333;7923;8;U;Asia/Chongqing
4034;Greenville-Spartanburg International;Greenville;United States;GSP;KGSP;34.895556;-82.218889;964;-5;A;America/New_York
4035;Cologne Railway;Cologne;Germany;QKL;\N;50.9425;6.958056;100;1;E;Europe/Berlin
4036;Stuttgart Railway Station;Stuttgart;Germany;ZWS;\N;48.783611;9.181667;200;1;E;Europe/Berlin
4037;Central Illinois Rgnl;Bloomington;United States;BMI;KBMI;40.477111;-88.915917;871;-6;A;America/Chicago
4038;Gulfport-Biloxi;Gulfport;United States;GPT;KGPT;30.407278;-89.070111;28;-6;A;America/Chicago
4039;Kalamazoo;Kalamazoo;United States;AZO;KAZO;42.234875;-85.552058;874;-5;A;America/New_York
4040;Toledo;Toledo;United States;TOL;KTOL;41.586806;-83.807833;684;-5;A;America/New_York
4041;Fort Wayne;Fort Wayne;United States;FWA;KFWA;40.978472;-85.195139;815;-5;A;America/New_York
4042;Decatur;Decatur;United States;DEC;KDEC;39.834564;-88.865689;682;-6;A;America/Chicago
4043;Cedar Rapids;Cedar Rapids;United States;CID;KCID;41.884694;-91.710806;869;-6;A;America/Chicago
4044;La Crosse Municipal;La Crosse;United States;LSE;KLSE;43.878986;-91.256711;654;-6;A;America/Chicago
4045;Central Wisconsin;Wassau;United States;CWA;KCWA;44.772726;-89.646635;840;-6;A;America/Chicago
4046;Peoria Regional;Peoria;United States;PIA;KPIA;40.664203;-89.693258;660;-6;A;America/Chicago
4047;Appleton;Appleton;United States;ATW;KATW;44.257526;-88.507576;680;-6;A;America/Chicago
4048;Rochester;Rochester;United States;RST;KRST;43.908283;-92.500014;1317;-6;A;America/Chicago
4049;Champaign;Champaign;United States;CMI;KCMI;40.03925;-88.278056;754;-6;A;America/Chicago
4050;Manhattan Reigonal;Manhattan;United States;MHK;KMHK;39.140972;-96.670833;1057;-6;A;America/Chicago
4051;Kingscote Airport;Kingscote;Australia;KGC;YKSC;-35.713889;137.521389;24;9.5;O;Australia/Adelaide
4052;Hervey Bay Airport;Hervey Bay;Australia;HVB;YHBA;-25.318889;152.880278;60;10;O;Australia/Brisbane
4053;EuroAirport Basel-Mulhouse-Freiburg;Basel;Switzerland;BSL;\N;47.59;7.529167;885;1;E;Europe/Paris
4054;Dali;Dali;China;DLU;ZPDL;25.649444;100.319444;6978;8;U;Asia/Chongqing
4055;Jinghong;Jinghong;China;JHG;\N;21.973914;100.759611;1815;8;U;Asia/Chongqing
4056;Mulu;Mulu;Malaysia;MZV;\N;4.048333;114.805;80;8;N;Asia/Kuala_Lumpur
4057;Sharm El Sheikh Intl;Sharm El Sheikh;Egypt;SSH;HESH;27.977222;34.394722;143;2;N;Africa/Cairo
4058;Franklin;Franklin;United States;FKL;KFKL;41.377874;-79.860362;1540;-5;A;America/New_York
4059;Jomo Kenyatta International;Nairobi;Kenya;NBO;HKJK;-1.319167;36.9275;5327;3;N;Africa/Nairobi
4060;Seronera;Seronera;Tanzania;SEU;HTSN;-2.458056;34.8225;5080;3;U;Africa/Dar_es_Salaam
4061;El Calafate;El Calafate;Argentina;FTE;SAWC;-50.280322;-72.053103;669;-3;N;America/Cordoba
4062;Armidale;Armidale;Australia;ARM;YARM;-30.528056;151.617222;3556;10;O;Australia/Sydney
4063;Grand Junction Regional;Grand Junction;United States;GJT;KGJT;39.122413;-108.526735;4858;-7;A;America/Denver
4064;St George Muni;Saint George;United States;SGU;KSGU;37.090583;-113.593056;2941;-7;A;America/Denver
4065;David Wayne Hooks Field;Houston;United States;DWH;KDWH;30.063746;-95.554276;152;-6;A;America/Chicago
4066;Port O\\'Connor Airfield;Port O\\'Connor;United States;S46;XS46;28.429977;-96.442859;10;-6;A;America/Chicago
4067;Sarasota Bradenton Intl;Sarasota;United States;SRQ;KSRQ;27.395444;-82.554389;30;-5;A;America/New_York
4071;Van Nuys;Van Nuys;United States;VNY;KVNY;34.209811;-118.489972;802;-8;A;America/Los_Angeles
4069;Bermuda Intl;Bermuda;Bermuda;BDA;TXKF;32.364042;-64.678703;12;-4;A;Atlantic/Bermuda
4070;Arthur Dunn Airpark;Titusville;United States;X21;\N;28.622552;-80.83541;30;-5;A;America/New_York
4072;Quad City Intl;Moline;United States;MLI;KMLI;41.448528;-90.507539;590;-6;A;America/Chicago
4073;Panama City Bay Co Intl;Panama City;United States;PFN;KPFN;30.212083;-85.682806;20;-6;A;America/Chicago
4074;Honiara International;Honiara;Solomon Islands;HIR;AGGH;-9.428;160.054789;28;11;U;Pacific/Guadalcanal
4075;Faa\\'a International;Papeete;French Polynesia;PPT;NTAA;-17.556667;-149.611389;7;-10;U;Pacific/Tahiti
4076;Nauru Intl;Nauru;Nauru;INU;ANYN;-0.547458;166.9191;22;12;U;Pacific/Nauru
4077;Funafuti International;Funafuti;Tuvalu;FUN;NGFU;-8.525;179.196389;9;12;U;Pacific/Funafuti
4078;Tolmachevo;Novosibirsk;Russia;OVB;UNNT;55.012622;82.650656;365;7;N;Asia/Omsk
4079;Orlando;Orlando;United States;DWS;\N;28.398;-81.57;340;-5;A;America/New_York
4080;Stavns;Samsoe;Denmark;;EKSS;55.89;10.62;1;1;E;Europe/Copenhagen
4081;Xieng Khouang;Phon Savan;Laos;XKH;VLXK;19.449997;103.158333;3445;7;U;Asia/Vientiane
4082;Phu Bai;Hue;Vietnam;HUI;\N;16.401499;107.702614;48;7;U;Asia/Saigon
4083;Bismarck Municipal Airport;Bismarck;United States;BIS;KBIS;46.775842;-100.757931;1661;-6;A;America/Chicago
4084;Telluride;Telluride;United States;TEX;KTEX;37.953759;-107.90848;9078;-7;A;America/Denver
4085;Yinchuan;Yinchuan;China;INC;ZLIC;38.481944;106.009167;3600;8;U;Asia/Chongqing
4086;Mae Hong Son;Mae Hong Son;Thailand;HGN;VTCH;19.301667;97.975;929;7;U;Asia/Bangkok
4087;Rapid City Regional Airport;Rapid City;United States;RAP;KRAP;44.045278;-103.057222;3204;-7;A;America/Denver
4088;McClellan-Palomar Airport;Carlsbad;United States;CLD;KCRQ;33.0742;-117.1648;328;-8;A;America/Los_Angeles
4089;Bishop International;Flint;United States;FNT;KFNT;42.965424;-83.743629;782;-5;A;America/New_York
4090;Francisco Bangoy International;Davao;Philippines;DVO;RPMD;7.125522;125.645778;59;8;N;Asia/Manila
4091;Madeira;Funchal;Portugal;FNC;LPMA;32.697889;-16.774453;192;0;E;Europe/Lisbon
4092;Santarem;Santarem;Brazil;STM;SBSN;-2.422431;-54.792789;198;-4;S;America/Boa_Vista
4093;Sihanoukville;Sihanoukville;Cambodia;KOS;VDSV;10.579686;103.636828;33;7;N;Asia/Phnom_Penh
4094;Ekati;Ekati;Canada;YOA;CYOA;64.698889;-110.614722;1540;-7;A;America/Edmonton
4095;Napier;NAPIER;New Zealand;NPE;NZNR;-39.465833;176.87;6;12;Z;Pacific/Auckland
4096;Levuka Airfield;Levuka;Fiji;LEV;NFNB;-17.68333;178.83333;11;12;N;Pacific/Fiji
4097;Lhasa-Gonggar;Lhasa;China;LXA;ZULS;29.297778;90.911944;13136;8;U;Asia/Chongqing
4098;Redding Muni;Redding;United States;RDD;KRDD;40.509;-122.293389;504;-8;A;America/Los_Angeles
4099;Mahlon Sweet Fld;Eugene;United States;EUG;KEUG;44.124583;-123.211972;374;-8;A;America/Los_Angeles
4100;Idaho Falls Rgnl;Idaho Falls;United States;IDA;KIDA;43.514556;-112.07075;4744;-7;A;America/Denver
4101;Rogue Valley Intl Medford;Medford;United States;MFR;KMFR;42.374228;-122.8735;1335;-8;A;America/Los_Angeles
4102;Kaikoura;Kaikoura;New Zealand;KBZ;NZKI;-42.416668;173.68333;19;12;Z;Pacific/Auckland
4103;Roberts Fld;Redmond-Bend;United States;RDM;KRDM;44.254066;-121.149964;3077;-8;A;America/Los_Angeles
4104;Koromiko;Picton;New Zealand;PCN;NZPN;-41.348333;173.955278;60;12;Z;Pacific/Auckland
4105;Windhoek Hosea Kutako International Airport ;Windhoek;Namibia;WDH;FYWV;-22.486667;17.4625;5640;1;S;Africa/Windhoek
4106;Victoria Inner Harbour Airport;Victoria;Canada;YWH;CYWH;48.422778;-123.3875;0;-8;A;America/Vancouver
4107;Vancouver Coal Harbour;Vancouver;Canada;CXH;CAQ3;49.289722;-123.115833;0;-8;A;America/Vancouver
4108;Jinan;Jinan;China;TNA;ZSJN;36.857214;117.215992;76;8;U;Asia/Chongqing
4109;Changzhou;Changzhou;China;CZX;ZSCG;31.941667;119.711667;39;8;U;Asia/Chongqing
4110;Yibin;Yibin;China;YBP;ZUYB;28.800278;104.544444;1000;8;U;Asia/Chongqing
4111;Roschino;Tyumen;Russia;TJM;USTR;57.189567;65.3243;378;6;N;Asia/Yekaterinburg
4112;Akron Canton Regional Airport;Akron;United States;CAK;KCAK;40.9160833;-81.4421944;1228;-5;A;America/New_York
4113;Huntsville International Airport-Carl T Jones Field;Huntsville;United States;HSV;KHSV;34.6371944;-86.7750556;629;-6;A;America/Chicago
4114;Mid-Ohio Valley Regional Airport;PARKERSBURG;United States;PKB;KPKB;39.3451039;-81.4392031;858;-5;A;America/New_York
4115;Montgomery Regional Airport ;MONTGOMERY;United States;MGM;KMGM;32.3006389;-86.3939722;221;-6;A;America/Chicago
4116;Tri-Cities Regional Airport;BRISTOL;United States;TRI;KTRI;36.4752222;-82.4074167;1519;-5;A;America/New_York
4117;Barkley Regional Airport;PADUCAH;United States;PAH;KPAH;37.0602875;-88.7729583;410;-6;A;America/Chicago
4118;Kurumoch;Samara;Russia;KUF;\N;53.5;50.15;0;4;N;Europe/Moscow
4119;Ambouli International Airport;Djibouti;Djibouti;JIB;HDAM;11.5472;43.1594;49;3;U;Africa/Djibouti
4120;Meilan;Haikou;China;HAK;ZJHK;19.934856;110.458961;75;8;U;Asia/Chongqing
4121;Mafia;Mafia Island;Tanzania;MFA;HTMA;-7.913889;39.665;60;3;U;Africa/Dar_es_Salaam
4127;Glacier Park Intl;Kalispell;United States;FCA;KFCA;48.310472;-114.256;2977;-7;A;America/Denver
4123;Mtemere Airstrip;Selous;Tanzania;;;-7.75093364715576;38.2032814025878;60;3;U;Africa/Dar_es_Salaam
4124;Page Municipal Airport;Page;United States;PGA;KPGA;36.9261;-111.4483;4316;-7;A;America/Phoenix
4125;Utila Airport;Utila;Honduras;UII;MHUT;16.091667;-86.8875;10;-6;U;America/Tegucigalpa
4126;Santa Elena Airport;Santa Elena de Uairen;Venezuela;SNV;\N;4.554722;-61.144922;3585;-4.5;U;America/Caracas
4128;Mbs Intl;Saginaw;United States;MBS;KMBS;43.532913;-84.079647;668;-5;A;America/New_York
4129;Greater Binghamton Edwin A Link Fld;Binghamton;United States;BGM;KBGM;42.208689;-75.979839;1636;-5;A;America/New_York
4130;Baghdad International Airport;Baghdad;Iraq;BGW;ORBI;33.262539;44.234578;114;3;U;Asia/Baghdad
4131;Nan;Nan;Thailand;NNT;VTCN;18.807914;100.783419;685;7;U;Asia/Bangkok
4132;Roi Et;Roi Et;Thailand;ROI;VTUV;16.116761;103.773797;451;7;U;Asia/Bangkok
4133;Buri Ram;Buri Ram;Thailand;BFV;VTUO;15.229539;103.253231;590;7;U;Asia/Bangkok
4134;Ranong;Ranong;Thailand;UNN;\N;9.777622;98.585483;57;7;U;Asia/Bangkok
4135;Trat;Trat;Thailand;TDX;VTBO;12.274572;102.318958;105;7;U;Asia/Bangkok
4136;Blythe Airport;Blythe;United States;BLH;KBLH;33.619167;-114.716889;399;-8;A;America/Los_Angeles
4137;Al Asad Airbase;Al Asad;Iraq;;ORAA;33.785608;42.4412;618;3;U;Asia/Baghdad
4138;Al Taqaddum Airbase;Al Taqaddum;Iraq;;ORAT;33.338053;43.597072;275;3;U;Asia/Baghdad
4139;Balad Southeast Airport;Al Bakr;Iraq;;ORBD;33.940194;44.361583;161;3;U;Asia/Baghdad
4140;Diosdado Macapagal International;Angeles City;Philippines;CRK;RPLC;15.185833;120.560278;484;8;N;Asia/Manila
4141;Sandakan;Sandakan;Malaysia;SDK;WBKS;5.900897;118.059486;46;8;N;Asia/Kuala_Lumpur
4142;Luang Namtha;Luang Namtha;Laos;LXG;VLLN;20.960556;101.4025;1968;7;U;Asia/Vientiane
4143;Oudomxay;Muang Xay;Laos;ODY;VLOS;20.6827;101.994;1380;7;N;Asia/Vientiane
4144;Shenyang Taoxian International Airport;Shenyang;China;SHE;ZYTX;41.3824;123.2901;198;8;U;Asia/Chongqing
4145;Dongying Airport;Dongying;China;DOY;ZSDY;37.2716;118.2819;86;8;U;Asia/Chongqing
4146;John A. Osborne Airport;Geralds;Montserrat;MNI;TRPG;16.791389;-62.193333;40;-4;N;America/Montserrat
4147;Petersburg James A. Johnson;Petersburg;United States;PSG;PAPG;56.801667;-132.945278;105;-9;A;America/Anchorage
4148;Luoyang Airport;Luoyang;China;LYA;ZHLY;34.41;112.28;210;8;U;Asia/Chongqing
4149;Xuzhou Guanyin Airport;Xuzhou;China;XUZ;ZSXZ;34.16;117.11;140;8;U;Asia/Chongqing
4150;Esfahan Shahid Beheshti Intl;Isfahan;Iran;IFN;\N;32.750836;51.861267;5059;3.5;E;Asia/Tehran
4151;Magwe;Magwe;Burma;MWQ;VYMW;20.165453;94.941185;275;6.5;U;Asia/Rangoon
4152;Khamti;Khamti;Burma;KHM;VYKI;25.988333;95.674444;6000;6.5;U;Asia/Rangoon
4153;Dalat;Dalat;Vietnam;DLI;VVDL;11.75;108.367;3156;7;U;Asia/Saigon
4154;Dong Hoi;Dong Hoi;Vietnam;VDH;\N;17.515;106.590556;50;7;U;Asia/Saigon
4155;Rach Gia;Rach Gia;Vietnam;VKG;VVRG;9.949676;105.133659;7;7;U;Asia/Saigon
4156;Ca Mau;Ca Mau;Vietnam;CAH;VVCM;9.188049;105.174721;50;7;U;Asia/Saigon
4157;Chu Lai;Chu Lai;Vietnam;VCL;VVCA;15.405944;108.705889;10;7;U;Asia/Saigon
4158;Dong Tac;Tuy Hoa;Vietnam;TBB;VVTH;13.04955;109.333706;20;7;U;Asia/Saigon
4159;Pai;Pai;Thailand;PYY;VTCI;19.372;98.437;1271;7;U;Asia/Bangkok
4160;Brac;Brac;Croatia;BWK;LDSB;43.285719;16.679719;1776;1;E;Europe/Zagreb
4161;Yaounde Nsimalen;Yaounde;Cameroon;NSI;FKYS;3.722556;11.553269;2278;1;N;Africa/Douala
4162;Conakry;Conakry;Guinea;CKY;GUCY;9.576889;-13.611961;72;0;N;Africa/Conakry
4163;Bergrestaurant;Trogen;Switzerland;;\N;46.992094;8.390315;1800;1;E;Europe/Zurich
4164;Uetliberg;Zuerich;Switzerland;;\N;47.36029;8.47805;1800;1;E;Europe/Zurich
4165;Flugplatz Merzbrueck;Aachen;Germany;AAH;EDKA;50.823194;6.186389;570;1;E;Europe/Berlin
4166;Baden Airpark;Karlsruhe/Baden-Baden;Germany;FKB;EDSB;48.7793;8.08048;407;1;E;Europe/Berlin
4167;Orlando Sanford Intl;Sanford;United States;SFB;KSFB;28.777639;-81.237489;55;-5;A;America/New_York
4168;Duong Dong Airport;Phu Quoc;Vietnam;PQC;\N;10.227025;103.967169;23;7;U;Asia/Saigon
4169;John Murtha Johnstown-Cambria County Airport;Johnstown;United States;JST;KJST;40.316111;-78.833889;2284;-5;A;America/New_York
4170;Lukla;Lukla;Nepal;LUA;VNLK;27.687778;86.731389;9100;5.75;N;Asia/Katmandu
4171;Bhojpur;Bhojpur;Nepal;BHP;VNBJ;27.14743;87.050819;4000;5.75;N;Asia/Katmandu
4172;Lamidanda;Lamidanda;Nepal;LDN;VNLD;27.253117;86.670044;4100;5.75;N;Asia/Katmandu
4173;Jomsom;Jomsom;Nepal;JMO;VNJS;28.782222;83.7225;8800;5.75;N;Asia/Katmandu
4174;Manang;Manang;Nepal;NGX;VNMA;28.633;84;11000;5.75;N;Asia/Katmandu
4175;Phaplu;Phaplu;Nepal;PPL;VNPL;27.517;86.6;7918;5.75;N;Asia/Katmandu
4176;Thamkharka;Thamkharka;Nepal;TMK;\N;27.052222;86.861944;6000;5.75;N;Asia/Katmandu
4177;Rumjatar;Rumjatar;Nepal;RUM;VNRT;27.303509;86.55043;4500;5.75;N;Asia/Katmandu
4178;Tulsipur;Dang;Nepal;DNP;VNDG;28.111111;82.294167;2100;5.75;N;Asia/Katmandu
4179;Rukumkot;Rukumkot;Nepal;RUK;VNRK;28.627;82.195;2500;5.75;N;Asia/Katmandu
4180;Jumla;Jumla;Nepal;JUM;VNJL;29.274167;82.193333;7700;5.75;N;Asia/Katmandu
4181;Chaurjhari;Chaurjhari;Nepal;HRJ;VNCJ;28;83.833;4000;5.75;N;Asia/Katmandu
4182;Taplejung;Taplejung;Nepal;TPJ;VNTJ;27.35;84.667;4000;5.75;N;Asia/Katmandu
4183;Tumling Tar;Tumling Tar;Nepal;TMI;VNTR;27.315;87.193333;1700;5.75;N;Asia/Katmandu
4184;Surkhet;Surkhet;Nepal;SKH;VNSK;28.586;81.636;2400;5.75;N;Asia/Katmandu
4185;Simikot;Simikot;Nepal;IMK;VNST;29.971064;81.818932;9246;5.75;N;Asia/Katmandu
4186;Dolpa;Dolpa;Nepal;DOP;VNDP;28.985718;82.819145;8200;5.75;N;Asia/Katmandu
4187;Bajhang;Bajhang;Nepal;BJH;VNBG;29.53896;81.185364;4100;5.75;N;Asia/Katmandu
4188;Dhangarhi;Dhangarhi;Nepal;DHI;VNDH;28.753333;80.581944;690;5.75;N;Asia/Katmandu
4189;Muan;Muan;South Korea;MWX;RKJB;34.991389;126.382778;51;9;U;Asia/Seoul
4190;Astypalaia;Astypalaia;Greece;JTY;LGPL;36.579886;26.375822;165;2;E;Europe/Athens
4191;Ikaria;Ikaria;Greece;JIK;LGIK;37.682717;26.347061;79;2;E;Europe/Athens
4192;Kalymnos Island;Kalymnos;Greece;JKL;LGKY;36.963333;26.940556;771;2;E;Europe/Athens
4193;Milos;Milos;Greece;MLO;LGML;36.696111;24.4775;12;2;E;Europe/Athens
4194;Naxos;Cyclades Islands;Greece;JNX;LGNX;37.080556;25.368056;10;2;E;Europe/Athens
4195;Paros;Paros;Greece;PAS;LGPA;37.010278;25.127778;121;2;E;Europe/Athens
4196;Kastelorizo;Kastelorizo;Greece;KZS;LGKJ;36.127777;29.566656;474;2;E;Europe/Athens
4197;Marsa Alam Intl;Marsa Alam;Egypt;RMF;HEMA;25.557111;34.583711;251;2;U;Africa/Cairo
4198;Niederrhein;Weeze;Germany;NRN;EDLV;51.602222;6.141944;106;1;E;Europe/Berlin
4199;Busuanga;Busuanga;Philippines;USU;RPVV;12.121458;120.100031;148;8;N;Asia/Manila
4200;Butuan;Butuan;Philippines;BXU;RPME;8.951322;125.477972;141;8;N;Asia/Manila
4201;Dipolog;Dipolog;Philippines;DPL;RPMG;8.601261;123.334481;12;8;N;Asia/Manila
4202;Laoag Intl;Laoag;Philippines;LAO;RPLI;18.178092;120.531522;25;8;N;Asia/Manila
4203;Legazpi;Legazpi;Philippines;LGP;RPLP;13.157064;123.746247;66;8;N;Asia/Manila
4204;Ozamis;Ozamis;Philippines;OZC;RPMO;8.178508;123.841731;75;8;N;Asia/Manila
6510;Zweibruecken;Zweibruecken;Germany;ZQW;\N;49.209445;7.401323;1132;1;E;Europe/Berlin
4206;Mactan Cebu Intl;Cebu;Philippines;CEB;RPVM;10.307542;123.979439;31;8;N;Asia/Manila
4207;Sonderlandeplatz Norden-Norddeich;Norden;Germany;NOE;EDWS;53.632221;7.191389;3;1;E;Europe/Berlin
4208;Verkehrslandeplatz Juist;Juist;Germany;JUI;EDWJ;53.681572;7.055731;6;1;E;Europe/Berlin
4209;Aeroporto de Porto Seguro;Porto Seguro;Brazil;BPS;SBPS;-16.438611;-39.080833;167;-3;S;America/Fortaleza
4210;Gounda Airport;Gounda;Central African Republic;GDA;\N;9.272222;21.197222;1509;1;N;Africa/Bangui
4211;Heliport Hotel das Cataratas;Cataratas National Park;Brazil;;\N;-25.684;-54.44024;1000;-3;S;America/Cordoba
4213;Iguatu;Iguatu;Brazil;;SNIG;-6.346639;-39.293777;699;-3;S;America/Fortaleza
4214;Palmas;Palmas;Brazil;PMW;SBPJ;-10.241667;-48.35278;774;-3;S;America/Fortaleza
4215;Caldas Novas;Caldas Novas;Brazil;CLV;SBCN;-17.7267;-48.6114;2247;-3;S;America/Sao_Paulo
4216;Missoula Intl;Missoula;United States;MSO;KMSO;46.916306;-114.090556;3205;-7;A;America/Denver
4217;Blackall;Blackall;Australia;BKQ;YBCK;-24.427778;145.428611;928;10;O;Australia/Brisbane
4218;Bundaberg;Bundaberg;Australia;BDB;YBUD;-24.903889;152.318611;107;10;O;Australia/Brisbane
4219;Grand Canyon National Park Airport;Grand Canyon;United States;GCN;KGCN;35.9523611;-112.1469722;6609;-7;N;America/Phoenix
4220;Sugar Land Regional Airport;Sugar Land;United States;SGR;KSGR;29.62225;-95.6565278;82;-6;A;America/Chicago
4221;Hayman Island Airport;Hayman Island;Australia;HIS;YHYN;-20.066668;148.86667;10;10;O;Australia/Brisbane
4222;Centennial;Denver;United States;APA;KAPA;39.570129;-104.849294;5883;-7;A;America/Denver
4223;Clovis Muni;Clovis;United States;CVN;KCVN;34.425139;-103.079278;4216;-7;A;America/Denver
4224;Fort Stockton Pecos Co;Fort Stockton;United States;FST;KFST;30.915667;-102.916139;3011;-6;A;America/Chicago
4225;Las Vegas Muni;Las Vegas;United States;LVS;KLVS;35.654222;-105.142389;6877;-7;A;America/Denver
4226;West Houston;Houston;United States;IWS;KIWS;29.818194;-95.672611;111;-6;A;America/Chicago
4227;La Junta Muni;La Junta;United States;LHX;KLHX;38.049719;-103.509431;4238;-7;A;America/Denver
4228;Las Cruces Intl;Las Cruces;United States;LRU;KLRU;32.289417;-106.921972;4456;-7;A;America/Denver
4229;Stephens Co;Breckenridge;United States;BKD;KBKD;32.719047;-98.891;1284;-6;A;America/Chicago
4230;Draughon Miller Central Texas Rgnl;Temple;United States;TPL;KTPL;31.1525;-97.407778;682;-6;A;America/Chicago
4231;Ozona Muni;Ozona;United States;OZA;KOZA;30.735281;-101.202972;2381;-6;A;America/Chicago
4232;Hong Kong Kai Tak;Hong Kong;Hong Kong;;VHXX;22.315365;114.195349;1;8;U;Asia/Hong_Kong
6416;Athen Helenikon Airport;Athens;Greece;HEW;\N;37.8933;23.7261;69;2;E;Europe/Athens
6415;Waikoloa Heliport;Waikoloa Village;United States;WKL;HI07;19.9136;-155.864;109;-10;N;Pacific/Honolulu
4235;Kaadedhdhoo;Kaadedhdhoo;Maldives;KDM;\N;0.4880555;72.995556;3;5;U;Indian/Maldives
4236;Aklavik;Aklavik;Canada;LAK;CYKD;68.223333;-135.005833;23;-7;A;America/Edmonton
4237;Deline;Deline;Canada;YWJ;CYWJ;65.1833333;-125.41666667;550;-7;A;America/Edmonton
4238;Tulita;Tulita;Canada;ZFN;CZFN;64.0833333;-125.5833333;320;-7;A;America/Edmonton
4239;Fort Good Hope;Fort Good Hope;Canada;YGH;CYGH;66.26666667;-128.65;215;-7;A;America/Edmonton
4240;Inuvik Town;Inuvik;Canada;;\N;66.3676;-133.7594;200;-7;A;America/Edmonton
4241;INAWR;Inuvik;Canada;;\N;68.3676;-133.7594;200;-7;A;America/Edmonton
4242;Tanna island;Tanna;Vanuatu;TAH;\N;-19.455198;169.22394;3;11;U;Pacific/Efate
6509;Sun Island;Sun Island;Maldives;;\N;3.48596289680953;72.8060746192932;0;5;N;Indian/Maldives
4244;Paulatuk;Paulatuk;Canada;YPC;CYPC;62.35;-124.3333;50;-7;A;America/Edmonton
4246;Nicholson Peninsula;Nicholson Peninsula;Canada;;YUCX;69.917;-128.95;1;-7;A;America/Edmonton
4247;Santa Cruz;Santa Cruz;Bolivia;SRZ;\N;-17.8;-63.166667;1224;-4;U;America/La_Paz
4248;Kulusuk;Kulusuk;Greenland;KUS;\N;65.566667;-37.1166667;112;-3;E;America/Godthab
4249;Juancho E. Yrausquin;Saba;Netherlands Antilles;SAB;TNCS;17.645278;-63.220556;60;-4;U;America/Curacao
4250;Eagle Co Rgnl;Vail;United States;EGE;KEGE;39.642556;-106.917694;6540;-7;A;America/Denver
4251;Stapleton;Denver;United States;;\N;39.779255;-104.88184;5333;-7;A;America/Denver
4252;Skagen;Stokmarknes;Norway;SKN;ENSK;68.580833;15.026111;11;1;E;Europe/Oslo
4253;Cuyahoga County;Richmond Heights;United States;CGF;KCGF;41.565124;-81.4863555;879;-5;A;America/New_York
4254;Mansfield Lahm Regional;Mansfield;United States;MFD;KMFD;40.8214167;-82.5166389;1297;-5;A;America/New_York
4255;Columbus Metropolitan Airport;Columbus;United States;CSG;KCSG;32.5163333;-84.9388611;397;-5;A;America/New_York
4256;Lawton-Fort Sill Regional Airport;Lawton;United States;LAW;KLAW;34.5677144;-98.4166367;1110;-6;A;America/Chicago
4257;Fort Collins Loveland Muni;Fort Collins;United States;FNL;KFNL;40.451828;-105.011336;5016;-7;A;America/Denver
6835;Gouvia Marina;Gouvia;Greece;;\N;39.648682;19.855771;0;2;E;Europe/Athens
6836;Paxos Marina;Paxos;Greece;;\N;39.202676;20.187968;0;2;E;Europe/Athens
4261;Flagstaff Pulliam Airport;Flagstaff;United States;FLG;KFLG;35.140318;-111.6692392;7015;-7;N;America/Phoenix
4262;Lake Tahoe Airport;South Lake Tahoe;United States;TVL;KTVL;38.893889;-119.995278;8544;-8;A;America/Los_Angeles
4263;Magic Valley Regional Airport;Twin Falls;United States;TWF;KTWF;42.481803;-114.487733;4151;-7;A;America/Denver
4264;Monaco;Monaco;Monaco;MCM;\N;43.73333333;7.41666666;0;1;E;Europe/Monaco
4265;Martha\\'s Vineyard;Vineyard Haven MA;United States;MVY;KMVY;41.391667;-70.615278;67;-5;A;America/New_York
4266;Newport State;Newport RI;United States;UUU;\N;41.533056;-71.282222;172;-5;A;America/New_York
4267;Hartness State;Springfield VT;United States;VSF;\N;43.343333;-72.517222;577;-5;A;America/New_York
4268;Concord Municipal;Concord NH;United States;CON;\N;43.20267;-71.50233;342;-5;A;America/New_York
4269;Sanford Regional;Sanford ME;United States;SFM;\N;43.39383;-70.708;244;-5;A;America/New_York
4270;Groton New London;Groton CT;United States;GON;KGON;41.330056;-72.045139;9;-5;A;America/New_York
4271;Saint Cloud Regional Airport;Saint Cloud;United States;STC;KSTC;45.546556;-94.059889;1031;-6;A;America/Chicago
4272;Bagan;Bagan;Burma;BPE;\N;21.17383327;94.9246666;0;6.5;U;Asia/Rangoon
4273;Golden Triangle Regional Airport;Columbus Mississippi;United States;GTR;KGTR;33.450333;-88.591361;264;-6;A;America/Chicago
4274;Nizhny Novgorod;Nizhniy Novgorod;Russia;GOJ;UWGG;56.230119;43.784042;256;4;N;Europe/Moscow
4275;Bowerman Field;Hoquiam;United States;HQM;\N;46.9711944;-123.9365556;18;-8;A;America/Los_Angeles
4276;Erie Intl Tom Ridge Fld;Erie;United States;ERI;KERI;42.082022;-80.176217;733;-5;A;America/New_York
4277;Conrad Maldives Resort;Rangali Island;Maldives;;\N;3.5833;72.7167;0;5;U;Indian/Maldives
4278;Barnstable Muni Boardman Polando Fld;Barnstable;United States;HYA;KHYA;41.669336;-70.280356;55;-5;A;America/New_York
4279;San Pedro;San Pedro;Belize;SPR;MZ10;17.913936;-87.971075;4;-6;U;America/Belize
4280;Sedona;Sedona;United States;SDX;KSEZ;34.848628;-111.788472;4830;-7;A;America/Phoenix
4281;Dry Tortugas;Dry Tortugas;United States;;\N;24.62;-82.87;5;-5;A;America/New_York
4282;Dry Tortugas;Dry Tortugas;United States;;\N;24.62;24.62;5;2;A;Africa/Tripoli
4283;Dry Tortugas;Dry Tortugas;United States;;\N;24.628;82.873;5;5.5;A;Asia/Calcutta
4284;Morgantown Muni Walter L Bill Hart Fld;Morgantown;United States;MGW;KMGW;39.642908;-79.916314;1248;-5;A;America/New_York
4285;Yeager;Charleston;United States;CRW;KCRW;38.373147;-81.593189;981;-5;A;America/New_York
4286;Wilkes Barre Scranton Intl;Scranton;United States;AVP;KAVP;41.338478;-75.723403;962;-5;A;America/New_York
4287;Bemidji Regional Airport;Bemidji;United States;BJI;KBJI;47.510722;-94.934722;1391;-6;A;America/Chicago
4288;Heydar Aliyev;Baku;Azerbaijan;BAK;\N;40.4675;50.046667;10;4;E;Asia/Baku
4289;Thangool;Biloela;Australia;THG;YTNG;-24.493889;150.576111;644;10;O;Australia/Brisbane
4290;Fagali\\'i;Apia;Samoa;FGI;NSFI;-13.84861111;-171.74083333;25;13;U;Pacific/Apia
4291;Ballina Byron Gateway;Ballina Byron Bay;Australia;BNK;YBNA;-28.833889;153.5625;7;10;O;Australia/Sydney
4292;Hector International Airport;Fargo;United States;FAR;KFAR;46.92065;-96.8157639;902;-6;A;America/Chicago
4293;Downtown;Kansas City;United States;MKC;KMKC;39.1275;-94.598889;759;-6;A;America/Chicago
6505;Phoenix-Mesa Gateway;Mesa;United States;AZA;KIWA;33.307833;-111.655;1382;-7;N;America/Phoenix
4295;Ratanakiri;Ratanakiri;Cambodia;RBE;\N;13.73;106.987;0;7;U;Asia/Phnom_Penh
4296;Gillette-Campbell County Airport;Gillette;United States;GCC;KGCC;44.3489167;-105.5393611;4365;-7;A;America/Denver
4297;Tomsk Bogashevo Airport;Tomsk;Russia;TOF;UNTT;56.380278;85.208333;597;7;N;Asia/Omsk
4298;El Toro;Santa Ana;United States;NZJ;KNZJ;33.676132;-117.731164;383;-8;A;America/Los_Angeles
4299;Phetchabun;Phetchabun;Thailand;PHY;VTPB;16.676028;101.195108;450;7;U;Asia/Bangkok
4300;Chumphon;Chumphon;Thailand;CJM;VTSE;10.7112;99.361706;18;7;U;Asia/Bangkok
4301;Jiuzhaigou Huanglong;Jiuzhaigou;China;JZH;ZUJZ;32.857;103.683;11311;8;U;Asia/Chongqing
4302;Wai Sha Airport;Shantou;China;SWA;ZGOW;23.4;116.683;0;8;U;Asia/Chongqing
4303;Les Ailerons;Enghien-moisselles;France;;LFFE;0;0;302;0;E;\N
4304;Cheddi Jagan Intl;Georgetown;Guyana;GEO;SYCJ;6.498553;-58.254119;95;-4;U;America/Guyana
4305;Ciudad del Este;Ciudad del Este;Paraguay;AGT;SGES;-25.4555;-54.843592;846;-4;S;America/Asuncion
4306;Ogle;Georgetown;Guyana;OGL;SYGO;6.806944;-58.104444;10;-4;U;America/Guyana
4307;Kaieteur;Kaieteur;Guyana;KAI;SYKA;5.167;-59.483;750;-4;U;America/Guyana
4308;Dunhuang Airport;Dunhuang;China;DNH;ZLDH;40.094;94.4818;3688;8;U;Asia/Chongqing
4309;Falconara;Ancona;Italy;AOI;LIPY;43.616389;13.362222;10;1;E;Europe/Rome
4310;Samjiyon;Samjiyon;South Korea;;ZZ04;41.542301;128.243571;4410;8;U;Asia/Chongqing
8833;Iosco County;East Tawas;United States;ECA;K6D9;44.311;-83.422;606;-5;A;America/New_York
4312;Copiapo;Copiapo;Chile;CPO;SCHA;-27;-70;1000;-4;S;America/Santiago
4313;Taba Intl;Taba;Egypt;TCP;HETB;29.587778;34.778056;2415;2;U;Africa/Cairo
4314;Edward Bodden Airfield;Little Cayman;Cayman Islands;LYB;MWCL;19.6591666667;-80.09083333;1;-5;U;America/Cayman
4315;Bodrum - Milas;Bodrum;Turkey;BJV;LTFE;37.249;27.667;19;2;E;Europe/Istanbul
4316;7 Novembre;Tabarka;Tunisia;TBJ;DTKA;36.978333;8.876389;0;1;E;Africa/Tunis
4317;Sabiha Gokcen;Istanbul;Turkey;SAW;LTFJ;40.898553;29.309219;312;2;E;Europe/Istanbul
4318;University Park Airport;State College Pennsylvania;United States;SCE;KUNV;40.849278;-77.848694;1239;-5;A;America/New_York
4319;Broome;Broome;Australia;BME;YPBR;-17.8;122.2;56;8;O;Australia/Perth
4320;Newcastle Airport;Newcastle;Australia;NTL;YWLM;-32.78;151.83;31;10;O;Australia/Sydney
4321;Bakki Airport;Bakki;Iceland;;BIBA;63.55805;-20.14833;45;0;N;Atlantic/Reykjavik
4322;Woerthersee International Airport;Klagenfurt;Austria;KLU;LOWK;46.642514;14.337739;452;1;E;Europe/Vienna
4323;General Manuel Carlos Piar;Ciudad Guayana;Venezuela;CGU;\N;8.288527;-62.760361;472;-4.5;U;America/Caracas
4324;Flugplatz Hoepen;Schneverdingen;Germany;;\N;53.1167;9.8;267;1;E;Europe/Berlin
4325;Hammerfest Airport;Hammerfest;Norway;HFT;ENHF;70.679722;23.668889;799;1;E;Europe/Oslo
4326;Valan;Honningsvag;Norway;HVG;ENHV;70.99;25.83;46;1;E;Europe/Oslo
4327;Mehamn;Mehamn;Norway;MEH;ENMR;71.029722;27.826667;41;1;E;Europe/Oslo
4328;Airport;Vadso;Norway;VDS;ENVD;70.065;29.844;127;1;E;Europe/Oslo
4329;Riem;Munich;Germany;;MUCX;48.137778;11.690278;1487;1;E;Europe/Berlin
4330;Imam Khomeini;Tehran;Iran;IKA;OIIE;35.416111;51.152222;3305;3.5;E;Asia/Tehran
4331;Mashhad;Mashhad;Iran;MHD;OIMM;36.234;59.643;3263;3.5;E;Asia/Tehran
4332;Egelsbach;Egelsbach;Germany;QEF;\N;49.9608;8.64361;384;1;E;Europe/Berlin
4333;Ust-Ilimsk;Ust Ilimsk;Russia;UIK;UIBS;58.135;102.566;0;9;N;Asia/Irkutsk
4334;Mudanjiang;Mudanjiang;China;MDG;\N;44.523889;129.568889;883;8;U;Asia/Chongqing
4335;Key Field;Meridian;United States;MEI;KMEI;32.332624;-88.751868;297;-6;A;America/Chicago
4336;Abraham Lincoln Capital;Springfield;United States;SPI;KSPI;39.8441;-89.677889;597;-6;A;America/Chicago
4337;Uzundzhovo;Haskovo;Bulgaria;HKV;LB14;41.976375;25.589817;160;2;E;Europe/Sofia
4338;Cortez Muni;Cortez;United States;CEZ;KCEZ;37.303;-108.628056;5918;-7;A;America/Denver
4339;Yampa Valley;Hayden;United States;HDN;KHDN;40.481181;-107.21766;6602;-7;A;America/Denver
4340;Gallup Muni;Gallup;United States;GUP;KGUP;35.511058;-108.789308;6472;-7;A;America/Denver
4341;Liberal Muni;Liberal;United States;LBL;KLBL;37.044222;-100.95986;2885;-6;A;America/Chicago
4342;Lamar Muni;Lamar;United States;LAA;KLAA;38.069694;-102.6885;3706;-7;A;America/Denver
4343;Renner Fld;Goodland;United States;GLD;KGLD;39.370621;-101.698992;3656;-7;A;America/Denver
4344;Yellowstone Rgnl;Cody;United States;COD;KCOD;44.520194;-109.023806;5102;-7;A;America/Denver
4345;Hovden;Orsta-Volda;Norway;HOV;ENOV;62.18;6.0742;242;1;E;Europe/Oslo
4346;RNAS WATTON;WATTON;United Kingdom;;EGYR;52.33;0.51;173;0;E;Europe/London
4347;ISLES OF SCILLY;ST MARY\\'S;United Kingdom;ISC;\N;49.919;-6.3075;70;0;E;Europe/London
4348;Springfield Branson Natl;Springfield;United States;SGF;KSGF;37.245667;-93.388639;1268;-6;A;America/Chicago
4349;Framnes;Narvik;Norway;NVK;ENNK;68.435833;17.388056;97;1;E;Europe/Oslo
4350;Berlevag;Berlevag;Norway;BVG;ENBV;70.866667;29;43;1;E;Europe/Oslo
4351;Fornebu;Oslo;Norway;FBU;ENFB;59.883333;10.616667;56;1;E;Europe/Oslo
4352;Alykel;Norilsk;Russia;NSK;UOOO;69.311053;87.332183;595;8;N;Asia/Krasnoyarsk
4353;Vityazevo;Anapa;Russia;AAQ;URKA;45.002097;37.347272;174;4;N;Europe/Moscow
4354;Joplin Rgnl;Joplin;United States;JLN;KJLN;37.151814;-94.498269;981;-6;A;America/Chicago
4355;Lehigh Valley Intl;Allentown;United States;ABE;KABE;40.652083;-75.440806;393;-5;A;America/New_York
4356;NW Arkansas Regional;Bentonville;United States;XNA;KXNA;36.2818694;-94.3068111;1287;-6;A;America/Chicago
4357;Atyrau;Atyrau;Kazakhstan;GUW;UATG;47.121944;51.821389;0;5;U;Asia/Oral
4358;Kzyl-Orda;Kzyl-Orda;Kazakhstan;KZO;UAOO;44.709;65.591;0;6;U;Asia/Qyzylorda
4359;South Bend Rgnl;South Bend;United States;SBN;KSBN;41.708661;-86.31725;799;-5;A;America/New_York
4360;Bykovo;Moscow;Russia;BKA;UUBB;55.617222;38.059999;427;4;N;Europe/Moscow
4361;Chintheche;Chintheche;Malawi;;FWCC;-11.8333;34.1667;0;2;U;Africa/Blantyre
4362;Talagi;Arkhangelsk;Russia;ARH;ULAA;64.360281;40.430167;19;4;N;Europe/Moscow
4363;Central;Saratov;Russia;RTW;UWSS;51.334366;46.022952;152;4;N;Europe/Moscow
4364;Novyi Urengoy;Novy Urengoy;Russia;NUX;USMU;66.041811;76.313938;20;6;N;Asia/Yekaterinburg
4365;Noyabrsk;Noyabrsk;Russia;NOJ;USRO;63.110079;75.162243;20;6;N;Asia/Yekaterinburg
6508;Washington Union Station;Washington;United States;ZWU;\N;38.89746;-77.00643;76;-5;A;America/New_York
4367;Aktau;Aktau;Kazakhstan;SCO;UATE;43.86005;51.091978;21;5;U;Asia/Oral
4368;Ukhta;Ukhta;Russia;UCT;\N;63.340297;53.482592;100;4;N;Europe/Moscow
4369;Usinsk;Usinsk;Russia;USK;\N;66.00077;57.221113;20;4;N;Europe/Moscow
4370;Pechora;Pechora;Russia;PEX;\N;65.070387;57.082045;20;4;N;Europe/Moscow
4371;Naryan-Mar;Naryan-Mar;Russia;NNM;ULAM;67.380537;53.051016;20;4;N;Europe/Moscow
4372;Kresty;Pskov;Russia;PKV;ULOO;57.783917;28.395614;154;4;N;Europe/Moscow
4373;Kogalym International;Kogalym;Russia;KGP;USRK;62.18;74.53;220;6;N;Asia/Yekaterinburg
4374;Emelyanovo;Krasnoyarsk;Russia;KJA;UNKL;56.18;92.475;940;8;N;Asia/Krasnoyarsk
4375;Sary-Arka;Karaganda;Kazakhstan;KGF;UAKK;49.670833;73.334444;1765;6;U;Asia/Qyzylorda
4376;Severny;Novosibirsk;Russia;;UNCC;55.1;82.9;560;7;N;Asia/Omsk
4377;Uraj;Uraj;Russia;URJ;USHU;60.1;64.83;186;6;N;Asia/Yekaterinburg
4378;Turkmenabat;Turkmenabat;Turkmenistan;CRZ;\N;39.083333;63.602222;630;5;U;Asia/Ashgabat
4379;Yuzhny;Ivanovo;Russia;IWA;UUBI;56.939444;40.940833;410;4;N;Europe/Moscow
4380;Changchun;Changchun;China;CGQ;ZYCC;43.5412;125.1201;227;8;U;Asia/Chongqing
4381;Niigata;Niigata;Japan;KIJ;RJSN;37.5711;139.0646;1;9;U;Asia/Tokyo
4382;Johnston Atoll;Johnston Island;Johnston Atoll;JON;PJON;16.7286;-169.534;7;-10;U;Pacific/Johnston
4383;Smith Fld;Fort Wayne IN;United States;SMD;KSMD;41.143353;-85.152778;834;-5;A;America/New_York
4384;Arcata;Arcata CA;United States;ACV;KACV;40.978111;-124.108611;221;-8;A;America/Los_Angeles
4385;Camp Mabry Austin City;Austin TX;United States;ATT;KATT;30.31666;-97.7666;0;-6;A;America/Chicago
4386;Albert J Ellis;Jacksonville NC;United States;OAJ;KOAJ;34.829164;-77.612139;94;-5;A;America/New_York
4387;Tuscaloosa Rgnl;Tuscaloosa AL;United States;TCL;KTCL;33.220627;-87.611403;170;-6;A;America/Chicago
4388;Dubuque Rgnl;Dubuque IA;United States;DBQ;KDBQ;42.402;-90.709472;1076;-6;A;America/Chicago
4389;Forde Bringeland;Forde;Norway;FDE;\N;61.392;5.763;1034;1;E;Europe/Oslo
4390;Shun Tak Heliport;Hong Kong;Hong Kong;;VHST;22.289372;114.152153;10;8;N;Asia/Hong_Kong
6092;Poliarny Airport;Yakutia;Russia;PYJ;UERP;66.400431;112.030325;1660;10;N;Asia/Yakutsk
6090;Nakhchivan Airport;Nakhchivan;Azerbaijan;NAJ;UBBN;39.1888;45.4584;2863;4;E;Asia/Baku
6089;Ganja Airport;Ganja;Azerbaijan;KVD;UBBG;40.7377;46.3176;1083;4;E;Asia/Baku
6804;Torp;Sandefjord;Norway;;\N;59.186703;10.258628;286;1;E;Europe/Oslo
6086;Ust Kamenogorsk Airport;Ust Kamenogorsk;Kazakhstan;UKK;UASK;50.0366;82.4942;939;6;U;Asia/Qyzylorda
6084;Petropavlosk South Airport;Petropavlosk;Kazakhstan;PPK;UACP;54.7747;69.1839;453;6;U;Asia/Qyzylorda
6078;Les Bases Airport;Grand Bourg;Guadeloupe;GBJ;TFFM;15.86875;-61.270022;16;-4;U;America/Guadeloupe
6077;St-François Airport;St-François;Guadeloupe;SFC;TFFC;16.2578;-61.2625;10;-4;U;America/Guadeloupe
6074;Codrington Airport;Codrington;Antigua and Barbuda;BBQ;TAPH;17.6358;-61.8286;15;-4;U;America/Antigua
6073;Ji-Paraná Airport;Ji-Paraná;Brazil;JPR;SWJI;-10.8708;-61.8465;598;-4;S;America/Boa_Vista
6071;Escuela Mariscal Sucre Airport;Maracay;Venezuela;MYC;SVBS;10.249978;-67.649419;1338;-4.5;U;America/Caracas
6068;Maria Reiche Neuman Airport;Nazca;Peru;NZA;SPZA;-14.854192;-74.961811;1860;-5;U;America/Lima
6066;Mayor General FAP Armando Revoredo Iglesias Airport;Cajamarca;Peru;CJA;SPJR;-7.139183;-78.4894;8781;-5;U;America/Lima
6064;Gatokae Airport;Gatokae;Solomon Islands;GTA;;-8.75;158.2;0;11;U;Pacific/Guadalcanal
6063;Boorama Airport;Boorama;Somalia;BXX;;9.933;43.15;0;3;U;Africa/Mogadishu
6062;Mucuri Airport;Mucuri;Brazil;MVS;SNMU;-18.0489;-39.8642;276;-3;S;America/Fortaleza
6061;Zorg en Hoop Airport;Paramaribo;Suriname;ORG;SMZO;5.81108;-55.1907;10;-3;U;America/Paramaribo
6059;Reyes Airport;Reyes;Bolivia;REY;SLRY;-14.3044;-67.3534;935;-4;U;America/La_Paz
6057;Puerto Rico Airport;Puerto Rico/Manuripi;Bolivia;PUR;SLPR;-11.1077;-67.5512;597;-4;U;America/La_Paz
6055;El Alcaraván Airport;Yopal;Colombia;EYP;SKYP;5.319114;-72.383975;1028;-5;U;America/Bogota
6047;General Rivadeneira Airport;Esmeraldas;Ecuador;ESM;SETN;0.978519;-79.6266;32;-5;U;America/Guayaquil
6043;Pucón Airport;Pucon;Chile;ZPC;SCPC;-39.2928;-71.9159;853;-4;S;America/Santiago
6041;Toledo Airport;Toledo;Brazil;TOW;SBTD;-24.6863;-53.6975;1843;-3;S;America/Sao_Paulo
6040;Santa Maria Airport;Santa Maria;Brazil;RIA;SBSM;-29.711358;-53.688153;287;-3;S;America/Sao_Paulo
6036;Chapada Diamantina Airport;Lençóis;Brazil;LEC;SBLE;-12.4823;-41.277;1676;-3;S;America/Fortaleza
6805;Goulburn Airport;Goulburn;Australia;GUL;YGLB;-34.8103;149.726;2141;10;O;Australia/Sydney
6034;Orlando Bezerra de Menezes Airport;Juazeiro Do Norte;Brazil;JDO;SBJU;-7.218958;-39.2701;1392;-3;S;America/Fortaleza
6032;Santa Teresita Airport;Santa Teresita;Argentina;SST;SAZL;-36.542317;-56.721756;9;-3;N;America/Cordoba
6031;Gobernador Gregores Airport;Gobernador Gregores;Argentina;GGS;SAWR;-48.7831;-70.15;356;-3;N;America/Cordoba
6029;Antoine De St Exupery Airport;San Antonio Oeste;Argentina;OES;SAVN;-40.7512;-65.0343;85;-3;N;America/Cordoba
6028;Las Heras Airport;Las Heras;Argentina;LHS;SAVH;-46.533056;-68.951111;1082;-3;N;America/Cordoba
6027;General Enrique Mosconi Airport;Tartagal;Argentina;TTG;SAST;-22.619167;-63.793189;1472;-3;N;America/Cordoba
6025;Masbate Airport;Masbate;Philippines;MBT;RPVJ;12.3694;123.629;26;8;N;Asia/Manila
6024;Catarman National Airport;Catarman;Philippines;CRM;RPVF;12.502417;124.635778;6;8;N;Asia/Manila
6013;Jolo Airport;Jolo;Philippines;JOL;RPMJ;6.05367;121.011;118;8;N;Asia/Manila
6012;Camiguin Airport;Camiguin;Philippines;CGM;RPMH;9.25352;124.707;53;8;N;Asia/Manila
6010;Cuyo Airport;Cuyo;Philippines;CYU;RPLO;10.8581;121.069;0;8;N;Asia/Manila
6008;Cheongju International Airport;Chongju;South Korea;CJJ;RKTU;36.7166;127.499119;191;9;U;Asia/Seoul
6007;Sacheon Air Base;Sacheon;South Korea;HIN;RKPS;35.088543;128.07037;25;9;U;Asia/Seoul
6005;Wonju Airport;Wonju;South Korea;WJU;RKNW;37.438081;127.960383;329;9;U;Asia/Seoul
6004;Mokpo Airport;Mokpo;South Korea;MPK;RKJM;34.758906;126.379872;23;9;U;Asia/Seoul
6003;Kunsan Air Base;Kunsan;South Korea;KUV;RKJK;35.903756;126.615906;29;9;U;Asia/Seoul
6002;Miyakejima Airport;Miyakejima;Japan;MYE;RJTQ;34.0736;139.56;67;9;U;Asia/Tokyo
6001;Shonai Airport;Shonai;Japan;SYO;RJSY;38.812222;139.787222;86;9;U;Asia/Tokyo
6000;Odate Noshiro Airport;Odate Noshiro;Japan;ONJ;RJSR;40.1919;140.371;292;9;U;Asia/Tokyo
5999;Fukushima Airport;Fukushima;Japan;FKS;RJSF;37.2274;140.431;1221;9;U;Asia/Tokyo
5998;Iwami Airport;Iwami;Japan;IWJ;RJOW;34.6764;131.79;184;9;U;Asia/Tokyo
5997;Nagoya Airport;Nagoya;Japan;NKM;RJNA;35.255;136.924;52;9;U;Asia/Tokyo
5996;Saga Airport;Saga;Japan;HSG;RJFS;33.1497;130.302;6;9;N;Asia/Tokyo
5995;Okadama Airport;Sapporo;Japan;OKD;RJCO;43.1161;141.38;25;9;U;Asia/Tokyo
5994;Kushiro Airport;Kushiro;Japan;KUH;RJCK;43.041;144.193;327;9;U;Asia/Tokyo
5993;Matsu Beigan Airport;Matsu Islands;Taiwan;MFK;RCMT;26.224153;120.00275;41;8;U;Asia/Taipei
5992;Hengchun Airport;Hengchun;Taiwan;HCN;RCKW;22.041075;120.730208;46;8;U;Asia/Taipei
5991;Matsu Nangan Airport;Matsu Islands;Taiwan;LZN;RCFG;26.1598;119.958;232;8;U;Asia/Taipei
5990;Eniwetok Airport;Eniwetok Atoll;Marshall Islands;ENT;PKMA;11.34075;162.327861;13;12;U;Pacific/Majuro
5989;Kalaupapa Airport;Molokai;United States;LUP;PHLU;21.211;-156.974;24;-10;A;Pacific/Honolulu
5988;El Nido Airport;El Nido;Philippines;ENI;;11.202;119.417;0;8;N;Asia/Manila
5987;Wipim Airport;Wipim;Papua New Guinea;WPM;;-8.78822;142.882;173;10;U;Pacific/Port_Moresby
5986;Baimuru Airport;Baimuru;Papua New Guinea;VMU;;-7.49686;144.82;27;10;U;Pacific/Port_Moresby
5985;Nuku Airport;Nuku;Papua New Guinea;UKU;;-3.667;142.483;750;10;U;Pacific/Port_Moresby
5984;Tufi Airport;Tufi;Papua New Guinea;TFI;;-9.07595;149.32;85;10;U;Pacific/Port_Moresby
5983;Suki Airport;Suki;Papua New Guinea;SKC;;-8.033;141.717;24;10;U;Pacific/Port_Moresby
5982;Balimo Airport;Balimo;Papua New Guinea;OPU;;-8.05;142.933;51;10;U;Pacific/Port_Moresby
5981;Obo Airport;Obo;Papua New Guinea;OBX;;-7.583;141.317;29;10;U;Pacific/Port_Moresby
5980;Losuia Airport;Losuia;Papua New Guinea;LSA;;-8.50582;151.081;27;10;U;Pacific/Port_Moresby
5979;Londolovit Airport;Londolovit;Papua New Guinea;LNV;;-3.04361;152.629;167;10;U;Pacific/Port_Moresby
5978;Lake Murray Airport;Lake Murray;Papua New Guinea;LMY;;-7.00992;141.494;52;10;U;Pacific/Port_Moresby
5977;Kamusi Airport;Kamusi;Papua New Guinea;KUY;;-7.42035;143.122;0;10;U;Pacific/Port_Moresby
5976;Kokoda Airport;Kokoda;Papua New Guinea;KKD;;-8.88468;147.731;1240;10;U;Pacific/Port_Moresby
5975;Kandrian Airport;Kandrian;Papua New Guinea;KDR;;-6.183;149.533;280;10;U;Pacific/Port_Moresby
5974;Jacquinot Bay Airport;Jacquinot Bay;Papua New Guinea;JAQ;;-5.6525;151.507;136;10;U;Pacific/Port_Moresby
5973;Nissan Island Airport;Nissan Island;Papua New Guinea;IIS;;-4.49972;154.226;52;10;U;Pacific/Port_Moresby
5972;Ihu Airport;Ihu;Papua New Guinea;IHU;;-7.89756;145.396;47;10;U;Pacific/Port_Moresby
5971;Gasmata Island Airport;Gasmata Island;Papua New Guinea;GMI;;-6.27111;150.331;23;10;U;Pacific/Port_Moresby
5970;Tadji Airport;Aitape;Papua New Guinea;ATP;;-3.18985;142.43;48;10;U;Pacific/Port_Moresby
5969;Wrangell Airport;Wrangell;United States;WRG;PAWG;56.4843;-132.37;44;-9;A;America/Anchorage
5968;Chevak Airport;Chevak;United States;VAK;PAVA;61.5338;-165.584;75;-9;A;America/Anchorage
5967;Aniak Airport;Aniak;United States;ANI;PANI;61.5816;-159.543;88;-9;A;America/Anchorage
5966;Mountain Village Airport;Mountain Village;United States;MOU;PAMO;62.0954;-163.682;337;-9;A;America/Anchorage
5965;McGrath Airport;Mcgrath;United States;MCG;PAMC;62.9529;-155.606;338;-9;A;America/Anchorage
5964;Kalskag Airport;Kalskag;United States;KLG;PALG;61.5363;-160.341;55;-9;A;America/Anchorage
5963;Haines Airport;Haines;United States;HNS;PAHN;59.2438;-135.524;15;-9;A;America/Anchorage
5962;Holy Cross Airport;Holy Cross;United States;HCR;PAHC;62.1883;-159.775;70;-9;A;America/Anchorage
5961;Skagway Airport;Skagway;United States;SGY;PAGY;59.4601;-135.316;44;-9;A;America/Anchorage
5960;Gustavus Airport;Gustavus;United States;GST;PAGS;58.4253;-135.707;34;-9;A;America/Anchorage
5959;Adak Airport;Adak Island;United States;ADK;PADK;51.878;-176.646;18;-9;A;America/Anchorage
5958;Sambu Airport;Boca de Sábalo;Panama;SAX;;8.017;-78.2;32;-5;U;America/Panama
5957;Contadora Airport;Contadora Island;Panama;OTD;;8.62876;-79.0347;14;-5;U;America/Panama
5956;EL Real Airport;El Real;Panama;ELE;;8.133;-77.7;57;-5;U;America/Panama
5955;Bahia Piña Airport;Bahia Piña;Panama;BFQ;;7.583;-78.2;14;-5;U;America/Panama
5954;Sayun International Airport;Sayun Intl;Yemen;GXF;OYSY;15.966111;48.7883;2097;3;U;Asia/Aden
5953;Kamishly Airport;Kamishly;Syria;KAC;OSKL;37.020625;41.191394;1480;2;E;Asia/Damascus
5952;Sulaymaniyah International Airport;Sulaymaniyah;Iraq;ISU;ORSU;35.5608;45.3147;2494;3;U;Asia/Baghdad
5951;Turbat International Airport;Turbat;Pakistan;TUK;OPTU;25.986369;63.030167;498;5;N;Asia/Karachi
5950;Sehwan Sharif Airport;Sehwan Sharif;Pakistan;SYW;OPSN;26.4731;67.7172;121;5;N;Asia/Karachi
5949;Skardu Airport;Skardu;Pakistan;KDU;OPSD;35.335508;75.536047;7316;5;N;Asia/Karachi
5948;Parachinar Airport;Parachinar;Pakistan;PAJ;OPPC;33.9021;70.0716;5800;5;N;Asia/Karachi
5947;Ormara Airport;Ormara Raik;Pakistan;ORW;OPOR;25.2747;64.586;10;5;N;Asia/Karachi
5946;Khuzdar Airport;Khuzdar;Pakistan;KDD;OPKH;27.7906;66.6473;4012;5;N;Asia/Karachi
5945;Hyderabad Airport;Hyderabad;Pakistan;HDD;OPKD;25.3181;68.3661;130;5;N;Asia/Karachi
5944;Jiwani Airport;Jiwani;Pakistan;JIW;OPJI;25.0678;61.8054;186;5;N;Asia/Karachi
5943;Dera Ismael Khan Airport;Dera Ismael Khan;Pakistan;DSK;OPDI;31.909422;70.896639;594;5;N;Asia/Karachi
5942;Dera Ghazi Khan Airport;Dera Ghazi Khan;Pakistan;DEA;OPDG;29.961011;70.485925;492;5;N;Asia/Karachi
5941;Dalbandin Airport;Dalbandin;Pakistan;DBA;OPDB;28.8783;64.3998;2800;5;N;Asia/Karachi
5940;Chitral Airport;Chitral;Pakistan;CJL;OPCH;35.886592;71.800578;4920;5;N;Asia/Karachi
5939;Bahawalpur Airport;Bahawalpur;Pakistan;BHV;OPBW;29.3481;71.717981;392;5;N;Asia/Karachi
5938;Bannu Airport;Bannu;Pakistan;BNP;OPBN;32.9729;70.5279;1325;5;N;Asia/Karachi
5937;Al Ain International Airport;Al Ain;United Arab Emirates;AAN;OMAL;24.261667;55.609167;869;4;U;Asia/Dubai
5936;Uromiyeh Airport;Uromiyeh;Iran;OMH;OITR;37.6681;45.0687;4343;3.5;E;Asia/Tehran
5935;Ardabil Airport;Ardabil;Iran;ADU;OITL;38.325678;48.424356;4315;3.5;E;Asia/Tehran
5934;Lar Airport;Lar;Iran;LRR;OISL;27.674725;54.383278;2641;3.5;E;Asia/Tehran
5933;Sari Dasht E Naz Airport;Dasht-e-naz;Iran;SRY;OINZ;36.635833;53.193611;35;3.5;E;Asia/Tehran
5932;Noshahr Airport;Noshahr;Iran;NSH;OINN;36.663333;51.464722;-61;3.5;E;Asia/Tehran
5931;Sabzevar National Airport;Sabzevar;Iran;AFZ;OIMS;36.168083;57.595183;3010;3.5;E;Asia/Tehran
5930;Bojnourd Airport;Bojnourd;Iran;BJB;OIMN;37.492958;57.308219;3499;3.5;E;Asia/Tehran
5929;Rafsanjan Airport;Rafsanjan;Iran;RJN;OIKR;30.297714;56.051139;5298;3.5;E;Asia/Tehran
5928;Bam Airport;Bam;Iran;BXR;OIKM;29.084169;58.450042;3231;3.5;E;Asia/Tehran
5927;Khoram Abad Airport;Khorram Abad;Iran;KHD;OICK;33.435378;48.282889;3782;3.5;E;Asia/Tehran
5926;Wadi Al Dawasir Airport;Wadi-al-dawasir;Saudi Arabia;EWD;OEWD;20.504275;45.199556;2062;3;U;Asia/Riyadh
5925;Al-Jawf Domestic Airport;Al-Jawf;Saudi Arabia;AJF;OESK;29.785133;40.100006;2261;3;U;Asia/Riyadh
5924;Dawadmi Domestic Airport;Dawadmi;Saudi Arabia;DWD;OEDW;24.5;44.4;3429;3;U;Asia/Riyadh
6506;Saul Airport;Saul;French Guiana;XAU;SOOS;3.61361;-53.2042;656;-3;U;America/Cayenne
5922;Faizabad Airport;Faizabad;Afghanistan;FBD;OAFZ;37.1211;70.5181;3872;4.5;U;Asia/Kabul
5921;Île des Pins Airport;Île des Pins;New Caledonia;ILP;NWWE;-22.5889;167.456;315;11;U;Pacific/Noumea
5920;Belep Islands Airport;Waala;New Caledonia;BMY;NWWC;-19.7206;163.661;306;11;U;Pacific/Noumea
5919;Tiga Airport;Tiga;New Caledonia;TGJ;NWWA;-21.0961;167.804;128;11;U;Pacific/Noumea
5918;Ipota Airport;Ipota;Vanuatu;IPA;NVVI;-18.8783;169.308;23;11;U;Pacific/Efate
5917;Futuna Airport;Futuna Island;Vanuatu;FTA;NVVF;-19.5164;170.232;0;11;U;Pacific/Efate
5916;Dillon's Bay Airport;Dillon's Bay;Vanuatu;DLY;NVVD;-18.7694;169.001;538;11;U;Pacific/Efate
5915;Aniwa Airport;Aniwa;Vanuatu;AWD;NVVB;-19.24;169.605;69;11;U;Pacific/Efate
5914;Anelghowhat Airport;Anelghowhat;Vanuatu;AUY;NVVA;-20.2492;169.771;7;11;U;Pacific/Efate
5913;North West Santo Airport;Olpoi;Vanuatu;OLZ;NVSZ;-14.8817;166.558;0;11;U;Pacific/Efate
5912;Southwest Bay Airport;Malekula Island;Vanuatu;SWJ;NVSX;-16.495;167.438;0;11;U;Pacific/Efate
5911;Valesdir Airport;Valesdir;Vanuatu;VLS;NVSV;-16.7961;168.177;10;11;U;Pacific/Efate
5910;Uléi Airport;Ambryn Island;Vanuatu;ULB;NVSU;-16.333;168.283;0;11;U;Pacific/Efate
5909;Tongoa Island Airport;Tongoa Island;Vanuatu;TGH;NVST;-16.8911;168.551;443;11;U;Pacific/Efate
5908;Santo Pekoa International Airport;Santo;Vanuatu;SON;NVSS;-15.505033;167.219742;184;11;U;Pacific/Efate
5907;Redcliffe Airport;Redcliffe;Vanuatu;RCL;NVSR;-15.472;167.835;36;11;U;Pacific/Efate
5906;Gaua Island Airport;Gaua Island;Vanuatu;ZGU;NVSQ;-14.2181;167.587;0;11;U;Pacific/Efate
5905;Norsup Airport;Norsup;Vanuatu;NUS;NVSP;-16.0797;167.401;23;11;U;Pacific/Efate
5904;Lonorore Airport;Lonorore;Vanuatu;LNE;NVSO;-15.8656;168.172;43;11;U;Pacific/Efate
5903;Naone Airport;Maewo Island;Vanuatu;MWF;NVSN;-15;168.083;0;11;U;Pacific/Efate
5902;Lamen Bay Airport;Lamen Bay;Vanuatu;LNB;NVSM;-16.5842;168.159;7;11;U;Pacific/Efate
5901;Lamap Airport;Lamap;Vanuatu;LPM;NVSL;-16.454;167.823;7;11;U;Pacific/Efate
5900;Tavie Airport;Paama Island;Vanuatu;PBJ;NVSI;-16.439;168.257;0;11;U;Pacific/Efate
5899;Sara Airport;Pentecost Island;Vanuatu;SSR;NVSH;-15.4708;168.152;0;11;U;Pacific/Efate
5898;Longana Airport;Longana;Vanuatu;LOD;NVSG;-15.3067;167.967;167;11;U;Pacific/Efate
5897;Craig Cove Airport;Craig Cove;Vanuatu;CCV;NVSF;-16.265;167.924;69;11;U;Pacific/Efate
5896;Sangafa Airport;Sangafa;Vanuatu;EAE;NVSE;-17.0903;168.343;7;11;U;Pacific/Efate
5895;Torres Airstrip;Loh/Linua;Vanuatu;TOH;NVSD;-13.328;166.638;0;11;U;Pacific/Efate
5894;Sola Airport;Sola;Vanuatu;SLH;NVSC;-13.8517;167.537;7;11;U;Pacific/Efate
5893;Mota Lava Airport;Ablow;Vanuatu;MTV;NVSA;-13.666;167.712;63;11;U;Pacific/Efate
5892;Ua Huka Airport;Ua Huka;French Polynesia;UAH;NTMU;-8.93611;-139.552;160;-9.5;U;Pacific/Marquesas
5891;Ua Pou Airport;Ua Pou;French Polynesia;UAP;NTMP;-9.35167;-140.078;16;-9.5;U;Pacific/Marquesas
5890;Hiva Oa-Atuona Airport;Hiva-oa;French Polynesia;AUQ;NTMN;-9.768794;-139.011256;1481;-9.5;U;Pacific/Marquesas
5889;Ahe Airport;Ahe;French Polynesia;AHE;NTHE;-14.4281;-146.257;11;-10;U;Pacific/Tahiti
5888;Apataki Airport;Apataki;French Polynesia;APK;NTGD;-15.5736;-146.415;8;-10;U;Pacific/Tahiti
5887;Maota Airport;Savaii Island;Samoa;MXS;NSMA;-13.733;-172.3;0;13;U;Pacific/Apia
5886;Mountain Airport;Mountain;Nepal;MWP;;28;85.333;0;5.75;N;Asia/Katmandu
5885;Pointe Vele Airport;Futuna Island;Wallis and Futuna;FUT;NLWF;-14.3114;-178.066;20;12;U;Pacific/Wallis
5884;Niue International Airport;Alofi;Niue;IUE;NIUE;-19.080028;-169.925639;209;-11;U;Pacific/Niue
5883;Vanua Balavu Airport;Vanua Balavu;Fiji;VBV;NFVB;-17.269;-178.976;76;12;U;Pacific/Fiji
5882;Kuini Lavenia Airport;Niuatoputapu;Tonga;NTT;NFTP;-15.9767;-173.755;0;13;U;Pacific/Tongatapu
5881;Mata'aho Airport;Angaha, Niuafo'ou Island;Tonga;NFO;NFTO;-15.5708;-175.633;160;13;U;Pacific/Tongatapu
5880;Lifuka Island Airport;Lifuka;Tonga;HPA;NFTL;-19.777;-174.341;31;13;U;Pacific/Tongatapu
5879;Kaufana Airport;Eua Island;Tonga;EUA;NFTE;-21.3783;-174.958;325;13;U;Pacific/Tongatapu
5878;Savusavu Airport;Savusavu;Fiji;SVU;NFNS;-16.8028;179.341;17;12;U;Pacific/Fiji
5877;Rotuma Airport;Rotuma;Fiji;RTA;NFNR;-12.4825;177.071;22;12;U;Pacific/Fiji
5876;Koro Island Airport;Koro Island;Fiji;KXF;NFNO;-17.3458;179.422;358;12;U;Pacific/Fiji
5875;Matei Airport;Matei;Fiji;TVU;NFNM;-16.6906;-179.877;60;12;U;Pacific/Fiji
5874;Labasa Airport;Lambasa;Fiji;LBS;NFNL;-16.466749;179.33986;44;12;U;Pacific/Fiji
5873;Lakeba Island Airport;Lakeba Island;Fiji;LKB;NFNK;-18.1992;-178.817;280;12;U;Pacific/Fiji
5872;Ngau Airport;Ngau;Fiji;NGI;NFNG;-18.1156;179.34;50;12;U;Pacific/Fiji
5871;Moala Airport;Moala;Fiji;MFJ;NFMO;-18.5667;179.951;13;12;U;Pacific/Fiji
5870;Mana Island Airport;Mana Island;Fiji;MNF;NFMA;-17.6731;177.098;0;12;U;Pacific/Fiji
5869;Vunisea Airport;Vunisea;Fiji;KDV;NFKD;-19.0581;178.157;6;12;U;Pacific/Fiji
5868;Malolo Lailai Island Airport;Malolo Lailai Island;Fiji;PTF;NFFO;-17.7779;177.197;10;12;U;Pacific/Fiji
5867;Cicia Airport;Cicia;Fiji;ICI;NFCI;-17.7433;-179.342;13;12;U;Pacific/Fiji
5866;Penrhyn Island Airport;Penrhyn Island;Cook Islands;PYE;NCPY;-9.00667;-158.037;8;-10;U;Pacific/Rarotonga
5865;Mitiaro Island Airport;Mitiaro Island;Cook Islands;MOI;NCMR;-19.8425;-157.703;25;-10;U;Pacific/Rarotonga
5864;Mauke Airport;Mauke Island;Cook Islands;MUK;NCMK;-20.1361;-157.345;26;-10;U;Pacific/Rarotonga
5863;Manihiki Island Airport;Manihiki Island;Cook Islands;MHX;NCMH;-10.3767;-161.002;0;-10;U;Pacific/Rarotonga
5862;Mangaia Island Airport;Mangaia Island;Cook Islands;MGS;NCMG;-21.8956;-157.905;45;-10;U;Pacific/Rarotonga
5861;Atiu Island Airport;Atiu Island;Cook Islands;AIU;NCAT;-19.9678;-158.119;36;-10;U;Pacific/Rarotonga
5860;Nassau Paradise Island Airport;Nassau;Bahamas;PID;MYPI;25.083;-77.3;0;-5;U;America/Nassau
5859;Colonel Hill Airport;Colonel Hill;Bahamas;CRI;MYCI;22.745561;-74.182353;5;-5;U;America/Nassau
5858;New Bight Airport;Cat Island;Bahamas;CAT;MYCB;24.315292;-75.452331;5;-5;U;America/Nassau
5857;Arthurs Town Airport;Arthur's Town;Bahamas;ATC;MYCA;24.629417;-75.673775;18;-5;U;America/Nassau
5856;Congo Town Airport;Andros;Bahamas;COX;MYAK;24.158933;-77.589758;15;-5;U;America/Nassau
5855;Long Banga Airport;Long Banga;Malaysia;LBP;;3.18495;115.454;750;8;N;Asia/Kuala_Lumpur
5854;Salina Cruz Naval Air Station;Salina Cruz;Mexico;SCX;MM57;16.2126;-95.2016;75;-6;S;America/Mexico_City
5853;Alberto Delgado Airport;Trinidad;Cuba;TND;MUTD;21.788461;-79.997203;125;-5;U;America/Havana
5852;Cayo Coco Airport;Cayo Coco;Cuba;CCC;MUOC;22.5132;-78.511;6;-5;U;America/Havana
5851;Port-de-Paix Airport;Port-de-Paix;Haiti;PAX;MTPX;19.9336;-72.8486;9;-5;U;America/Port-au-Prince
5850;Jeremie Airport;Jeremie;Haiti;JEE;MTJE;18.6631;-74.1703;147;-5;U;America/Port-au-Prince
5849;Playa Samara Airport;Playa Samara;Costa Rica;PLD;MRSR;10.25;-85.417;10;-6;U;America/Costa_Rica
5848;Tobias Bolanos International Airport;San Jose;Costa Rica;SYQ;MRPV;9.957053;-84.139797;3287;-6;U;America/Costa_Rica
5847;Puerto Jimenez Airport;Puerto Jimenez;Costa Rica;PJM;MRPJ;8.53333;-83.3;7;-6;U;America/Costa_Rica
5845;Islita Airport;Nandayure;Costa Rica;PBP;MRIA;9.85611;-85.3708;7;-6;U;America/Costa_Rica
5844;Cabo Velas Airport;Nicoya;Costa Rica;TNO;MRCV;10.3557;-85.852892;33;-6;U;America/Costa_Rica
5843;Barra del Colorado Airport;Pococi;Costa Rica;BCL;MRBC;10.768736;-83.585614;3;-6;U;America/Costa_Rica
5842;Aerotortuguero Airport;Roxana;Costa Rica;TTQ;MRAO;10.569;-83.5148;82;-6;U;America/Costa_Rica
5841;Captain Ramon Xatruch Airport;La Palma;Panama;PLP;MPLP;8.40667;-78.1417;30;-5;U;America/Panama
5840;Jaqué Airport;Jaqué;Panama;JQE;MPJE;7.51778;-78.1572;29;-5;U;America/Panama
5839;Enrique Adolfo Jimenez Airport;Colón;Panama;ONX;MPEJ;9.35664;-79.8674;25;-5;U;America/Panama
5838;Alonso Valderrama Airport;Chitré;Panama;CTD;MPCE;7.98784;-80.4097;33;-5;U;America/Panama
5837;Lencero Airport;Jalapa;Mexico;JAL;MMJA;19.475083;-96.797506;3127;-6;S;America/Mexico_City
5836;Guerrero Negro Airport;Guerrero Negro;Mexico;GUB;MMGR;28.0261;-114.024;59;-8;S;America/Tijuana
5835;Ciudad Constitución Airport;Ciudad Constitución;Mexico;CUA;MMDA;25.0538;-111.615;213;-7;S;America/Mazatlan
5834;Captain Rogelio Castillo National Airport;Celaya;Mexico;CYW;MMCY;20.545994;-100.88655;5709;-6;S;America/Mexico_City
5833;Mili Island Airport;Mili Island;Marshall Islands;MIJ;MLIP;6.08333;171.733;4;12;U;Pacific/Majuro
5832;Puerto Lempira Airport;Puerto Lempira;Honduras;PEU;MHPL;15.2622;-83.7812;25;-6;U;America/Tegucigalpa
5831;Ahuas Airport;Ahuas;Honduras;AHS;MHAH;15.4722;-84.3522;98;-6;U;America/Tegucigalpa
5830;Wotho Island Airport;Wotho Island;Marshall Islands;WTO;;10.1733;166.003;0;12;U;Pacific/Majuro
5829;Wotje Atoll Airport;Wotje Atoll;Marshall Islands;WTE;N36;9.46667;170.233;4;12;U;Pacific/Majuro
5828;Woja Airport;Majuro Atoll;Marshall Islands;WJA;;7.083;171.133;0;12;U;Pacific/Majuro
5827;Jaluit Airport;Jabor Jaluit Atoll;Marshall Islands;UIT;N55;5.90924;169.637;4;12;U;Pacific/Majuro
5826;Rongelap Island Airport;Rongelap Island;Marshall Islands;RNP;;11.1572;166.887;0;12;U;Pacific/Majuro
5825;Namorik Atoll Airport;Namorik Atoll;Marshall Islands;NDK;3N0;5.63167;168.125;15;12;U;Pacific/Majuro
5824;Majkin Airport;Majkin;Marshall Islands;MJE;;7.833;168.167;0;12;U;Pacific/Majuro
5823;Mejit Atoll Airport;Mejit Atoll;Marshall Islands;MJB;Q30;10.2833;170.883;5;12;U;Pacific/Majuro
5822;Maloelap Island Airport;Maloelap Island;Marshall Islands;MAV;;8.70444;171.23;0;12;U;Pacific/Majuro
5821;Likiep Airport;Likiep Island;Marshall Islands;LIK;;9.82316;169.308;0;12;U;Pacific/Majuro
5820;Kaben Airport;Kaben;Marshall Islands;KBT;;8.90056;170.844;0;12;U;Pacific/Majuro
5819;Jeh Airport;Ailinglapalap Atoll;Marshall Islands;JEJ;;7.56535;168.962;0;12;U;Pacific/Majuro
5818;Jabot Airport;Ailinglapalap Atoll;Marshall Islands;JAT;;7.45235;168.552;0;12;U;Pacific/Majuro
5817;Enyu Airfield;Bikini Atoll;Marshall Islands;BII;;11.5225;165.565;0;12;U;Pacific/Majuro
5816;Aur Island Airport;Aur Atoll;Marshall Islands;AUL;;8.14528;171.173;0;12;U;Pacific/Majuro
5815;Ailuk Airport;Ailuk Island;Marshall Islands;AIM;;10.2168;169.983;0;12;U;Pacific/Majuro
5814;Utirik Airport;Utirik Island;Marshall Islands;UTK;03N;11.222;169.852;4;12;U;Pacific/Majuro
5813;Quezaltenango Airport;Quezaltenango;Guatemala;AAZ;MGQZ;14.8656;-91.502;7779;-6;U;America/Guatemala
5812;Puerto Barrios Airport;Puerto Barrios;Guatemala;PBR;MGPB;15.730878;-88.583767;33;-6;U;America/Guatemala
5811;Dr Joaquin Balaguer International Airport;La Isabela;Dominican Republic;JBQ;MDJB;18.5725;-69.9856;98;-4;U;America/Santo_Domingo
5810;Samaná El Catey International Airport;Samana;Dominican Republic;AZS;MDCY;19.267;-69.742;30;-4;U;America/Santo_Domingo
5809;Salt Cay Airport;Salt Cay;Turks and Caicos Islands;SLX;MBSY;21.333;-71.2;3;-5;U;America/Grand_Turk
5808;Middle Caicos Airport;Middle Caicos;Turks and Caicos Islands;MDS;MBMC;21.833;-71.817;9;-5;U;America/Grand_Turk
5807;JAGS McCartney International Airport;Cockburn Town;Turks and Caicos Islands;GDT;MBGT;21.4445;-71.1423;13;-5;U;America/Grand_Turk
5806;Žilina Airport;Žilina;Slovakia;ILZ;LZZI;49.231528;18.6135;1020;1;E;Europe/Bratislava
5805;Ubari Airport;Ubari;Libya;QUB;;26.5675;12.8231;1387;2;N;Africa/Tripoli
5804;Misratah Airport;Misratah;Libya;MRA;;32.325;15.061;60;2;N;Africa/Tripoli
5803;Samsun-Çarşamba Airport;Samsun;Turkey;SZF;LTFH;41.2545;36.5671;18;2;E;Europe/Istanbul
5802;Balikesir Korfez Airport;Balikesir Korfez;Turkey;EDO;LTFD;39.5546;27.0138;50;2;E;Europe/Istanbul
5801;Isparta Süleyman Demirel Airport;Isparta;Turkey;ISE;LTFC;37.8554;30.3684;2835;2;E;Europe/Istanbul
5800;Adiyaman Airport;Adiyaman;Turkey;ADF;LTCP;37.7314;38.4689;2216;2;E;Europe/Istanbul
5799;Agri Airport;Agri;Turkey;AJI;LTCO;39.6546;43.0271;5462;2;E;Europe/Istanbul
5798;Kahramanmaras Airport;Kahramanmaras;Turkey;KCM;LTCN;37.539;36.9534;1723;2;E;Europe/Istanbul
5797;Sanliurfa Airport;Sanliurfa;Turkey;SFQ;LTCH;37.094261;38.847103;1483;2;E;Europe/Istanbul
5796;Kars Airport;Kars;Turkey;KSY;LTCF;40.562222;43.115002;5889;2;E;Europe/Istanbul
5795;Usak Airport;Usak;Turkey;USQ;LTBO;38.681478;29.471675;2897;2;E;Europe/Istanbul
5794;Banja Luka International Airport;Banja Luka;Bosnia and Herzegovina;BNX;LQBK;44.941444;17.297501;400;1;E;Europe/Sarajevo
5793;Corvo Airport;Corvo;Portugal;CVU;LPCR;39.6715;-31.1136;0;-1;E;Atlantic/Azores
5792;Salerno Pontecagnano Airport;Salerno;Italy;QSR;LIRI;40.6204;14.911294;123;1;E;Europe/Rome
5791;Aosta Airport;Aosta;Italy;AOT;LIMW;45.738456;7.368719;1791;1;E;Europe/Rome
5790;Sármellék International Airport;Sármellék;Hungary;SOB;LHSM;46.686389;17.159056;408;1;N;Europe/Budapest
5789;Győr-Pér International Airport;Győr;Hungary;QGY;LHPR;47.627097;17.808347;424;1;N;Europe/Budapest
5788;Pécs-Pogány Airport;Pécs-Pogány;Hungary;PEV;LHPP;45.990928;18.240983;1000;1;N;Europe/Budapest
5787;Syros Airport;Syros Island;Greece;JSY;LGSO;37.422792;24.950936;236;2;E;Europe/Athens
5786;La Môle Airport;La Môle;France;LTT;LFTZ;43.2054;6.482;59;1;E;Europe/Paris
6810;EuroAirport;Mulhouse;France;EAP;\N;47.589583;7.529914;885;1;E;Europe/Paris
6511;Tekapo;Lake Tekapo;New Zealand;;NZTL;-44;170.28;2;12;Z;Pacific/Auckland
5783;Angers-Loire Airport;Angers/Marcé;France;ANE;LFJR;47.5603;-0.312222;194;1;E;Europe/Paris
5782;Île d'Yeu Airport;Île d'Yeu;France;IDY;LFEY;46.718611;-2.391111;79;1;E;Europe/Paris
5781;Logroño-Agoncillo Airport;Logroño-Agoncillo;Spain;RJL;LELO;42.4542;-2.32083;1161;1;E;Europe/Madrid
5780;Ercan International Airport;Nicosia;Cyprus;ECN;LCEN;35.1547;33.4961;404;2;E;Asia/Nicosia
5779;Yakima Air Terminal McAllister Field;Yakima;United States;YKM;KYKM;46.5682;-120.544;1095;-8;A;America/Los_Angeles
5778;Kiwayu (Mkononi) Airport;Kiwayu;Kenya;KWY;;-1.96056;41.2975;21;3;U;Africa/Nairobi
5777;Worland Municipal Airport;Worland;United States;WRL;KWRL;43.9657;-107.951;4227;-7;A;America/Denver
5776;Valdosta Regional Airport;Valdosta;United States;VLD;KVLD;30.7825;-83.2767;203;-5;A;America/New_York
5775;Victoria Regional Airport;Victoria;United States;VCT;KVCT;28.8526;-96.9185;115;-6;A;America/Chicago
5774;Quincy Regional Baldwin Field;Quincy;United States;UIN;KUIN;39.9427;-91.1946;769;-6;A;America/Chicago
5773;Tupelo Regional Airport;Tupelo;United States;TUP;KTUP;34.2681;-88.7699;346;-6;A;America/Chicago
5772;Santa Maria Pub Cpt G Allan Hancock Airport;Santa Maria;United States;SMX;KSMX;34.8989;-120.457;261;-8;A;America/Los_Angeles
5771;Salina Municipal Airport;Salina;United States;SLN;KSLN;38.791;-97.6522;1288;-6;A;America/Chicago
5770;Adirondack Regional Airport;Saranac Lake;United States;SLK;KSLK;44.3853;-74.2062;1663;-5;A;America/New_York
5769;Sheridan County Airport;Sheridan;United States;SHR;KSHR;44.7692;-106.98;4021;-7;A;America/Denver
5768;San Luis County Regional Airport;San Luis Obispo;United States;SBP;KSBP;35.2368;-120.642;212;-8;A;America/Los_Angeles
5767;Rutland State Airport;Rutland;United States;RUT;KRUT;43.5294;-72.9496;787;-5;A;America/New_York
5766;Rock Springs Sweetwater County Airport;Rock Springs;United States;RKS;KRKS;41.5942;-109.065;6760;-7;A;America/Denver
5765;Rhinelander Oneida County Airport;Rhinelander;United States;RHI;KRHI;45.6312;-89.4675;1624;-6;A;America/Chicago
5764;Reading Regional Carl A Spaatz Field;Reading;United States;RDG;KRDG;40.3785;-75.9652;344;-5;A;America/New_York
5763;Pease International Tradeport;Portsmouth;United States;PSM;KPSM;43.0779;-70.8233;100;-5;A;America/New_York
5762;Pellston Regional Airport of Emmet County Airport;Pellston;United States;PLN;KPLN;45.5709;-84.7967;720;-5;A;America/New_York
5761;Pierre Regional Airport;Pierre;United States;PIR;KPIR;44.3827;-100.286;1742;-6;A;America/Chicago
5760;Pocatello Regional Airport;Pocatello;United States;PIH;KPIH;42.9098;-112.596;4452;-7;A;America/Denver
5759;Hattiesburg Laurel Regional Airport;Hattiesburg/Laurel;United States;PIB;KPIB;31.4671;-89.3371;298;-6;A;America/Chicago
5758;Owensboro Daviess County Airport;Owensboro;United States;OWB;KOWB;37.7401;-87.1668;406;-6;A;America/Chicago
5757;Southwest Oregon Regional Airport;North Bend;United States;OTH;KOTH;43.4171;-124.246;17;-8;A;America/Los_Angeles
5756;Northwest Alabama Regional Airport;Muscle Shoals;United States;MSL;KMSL;34.7453;-87.6102;550;-6;A;America/Chicago
5755;Frank Wiley Field;Miles City;United States;MLS;KMLS;46.428;-105.886;2630;-7;A;America/Denver
5754;Muskegon County Airport;Muskegon;United States;MKG;KMKG;43.1695;-86.2382;628;-5;A;America/New_York
5753;Lynchburg Regional Preston Glenn Field;Lynchburg;United States;LYH;KLYH;37.3267;-79.2004;938;-5;A;America/New_York
5752;Lewistown Municipal Airport;Lewistown;United States;LWT;KLWT;47.0493;-109.467;4170;-7;A;America/Denver
5751;Lancaster Airport;Lancaster;United States;LNS;KLNS;40.1217;-76.2961;403;-5;A;America/New_York
5750;Klamath Falls Airport;Klamath Falls;United States;LMT;KLMT;42.1561;-121.733;4095;-8;A;America/Los_Angeles
5749;Lebanon Municipal Airport;Lebanon;United States;LEB;KLEB;43.6261;-72.3042;603;-5;A;America/New_York
5748;North Platte Regional Airport Lee Bird Field;North Platte;United States;LBF;KLBF;41.1262;-100.684;2776;-6;A;America/Chicago
5747;Arnold Palmer Regional Airport;Latrobe;United States;LBE;KLBE;40.2759;-79.4048;1185;-5;A;America/New_York
5746;Laramie Regional Airport;Laramie;United States;LAR;KLAR;41.3121;-105.675;7284;-7;A;America/Denver
5745;Jamestown Regional Airport;Jamestown;United States;JMS;KJMS;46.9297;-98.6782;1498;-6;A;America/Chicago
5744;Kirksville Regional Airport;Kirksville;United States;IRK;KIRK;40.0935;-92.5449;966;-6;A;America/Chicago
5743;Kili Airport;Kili Island;Marshall Islands;KIO;Q51;5.64452;169.12;5;12;U;Pacific/Majuro
5742;Tri State Milton J Ferguson Field;Huntington;United States;HTS;KHTS;38.3667;-82.558;828;-5;A;America/New_York
5741;Memorial Field;Hot Springs;United States;HOT;KHOT;34.478;-93.0962;540;-6;A;America/Chicago
5740;Central Nebraska Regional Airport;Grand Island;United States;GRI;KGRI;40.9675;-98.3096;1847;-6;A;America/Chicago
5739;Wokal Field Glasgow International Airport;Glasgow;United States;GGW;KGGW;48.2125;-106.615;2296;-7;A;America/Denver
5738;Fayetteville Regional Grannis Field;Fayetteville;United States;FAY;KFAY;34.9912;-78.8803;189;-5;A;America/New_York
5737;New Bedford Regional Airport;New Bedford;United States;EWB;KEWB;41.6761;-70.9569;80;-5;A;America/New_York
5736;Elko Regional Airport;Elko;United States;EKO;KEKO;40.8249;-115.792;5140;-8;A;America/Los_Angeles
5735;Chippewa Valley Regional Airport;Eau Claire;United States;EAU;KEAU;44.8658;-91.4843;913;-6;A;America/Chicago
5734;DuBois Regional Airport;Du Bois;United States;DUJ;KDUJ;41.1783;-78.8987;1817;-5;A;America/New_York
5733;Dodge City Regional Airport;Dodge City;United States;DDC;KDDC;37.7634;-99.9656;2594;-6;A;America/Chicago
5732;Houghton County Memorial Airport;Hancock;United States;CMX;KCMX;47.1684;-88.4891;1095;-5;A;America/New_York
5731;William R Fairchild International Airport;Port Angeles;United States;CLM;KCLM;48.1202;-123.5;291;-8;A;America/Los_Angeles
5730;Harrison Marion Regional Airport;Clarksburg;United States;CKB;KCKB;39.2966;-80.2281;1217;-5;A;America/New_York
5729;Chippewa County International Airport;Sault Ste Marie;United States;CIU;KCIU;46.2508;-84.4724;800;-5;A;America/New_York
5728;Cape Girardeau Regional Airport;Cape Girardeau;United States;CGI;KCGI;37.2253;-89.5708;342;-6;A;America/Chicago
5727;Del Norte County Airport;Crescent City;United States;CEC;KCEC;41.7802;-124.237;57;-8;A;America/Los_Angeles
5726;Southeast Iowa Regional Airport;Burlington;United States;BRL;KBRL;40.7832;-91.1255;698;-6;A;America/Chicago
5725;Brunswick Golden Isles Airport;Brunswick;United States;BQK;KBQK;31.2588;-81.4665;26;-5;A;America/New_York
5724;Raleigh County Memorial Airport;Beckley;United States;BKW;KBKW;37.7873;-81.1242;2504;-5;A;America/New_York
5723;Western Nebraska Regional Airport;Scottsbluff;United States;BFF;KBFF;41.874;-103.596;3967;-7;A;America/Denver
5722;Bradford Regional Airport;Bradford;United States;BFD;KBFD;41.8031;-78.6401;2143;-5;A;America/New_York
5721;Watertown Regional Airport;Watertown;United States;ATY;KATY;44.914;-97.1547;1748;-6;A;America/Chicago
5720;Alpena County Regional Airport;Alpena;United States;APN;KAPN;45.0781;-83.5603;689;-5;A;America/New_York
5719;Walla Walla Regional Airport;Walla Walla;United States;ALW;KALW;46.0949;-118.288;1191;-8;A;America/Los_Angeles
5718;Waterloo Regional Airport;Waterloo;United States;ALO;KALO;42.5571;-92.4003;873;-6;A;America/Chicago
5717;Alamogordo White Sands Regional Airport;Alamogordo;United States;ALM;KALM;32.8399;-105.991;4200;-7;A;America/Denver
5716;Athens Ben Epps Airport;Athens;United States;AHN;KAHN;33.9486;-83.3263;808;-5;A;America/New_York
5715;Southwest Georgia Regional Airport;Albany;United States;ABY;KABY;31.5355;-84.1945;197;-5;A;America/New_York
5714;Aberdeen Regional Airport;Aberdeen;United States;ABR;KABR;45.4491;-98.4218;1302;-6;A;America/Chicago
5713;San Domino Island Heliport;Tremiti Islands;Italy;TQR;;42.1025;15.488;0;1;E;Europe/Rome
5712;Gheshm Airport;Gheshm;Iran;GSM;;26.9487;56.2688;100;3.5;E;Asia/Tehran
5711;Diu Airport;Diu;India;DIU;VA1P;20.7131;70.9211;31;5.5;N;Asia/Calcutta
5710;Gulu Airport;Gulu;Uganda;ULU;HUGU;2.805556;32.271792;3510;3;U;Africa/Kampala
5709;Arua Airport;Arua;Uganda;RUA;HUAR;3.05;30.917;3951;3;U;Africa/Kampala
5708;Tabora Airport;Tabora;Tanzania;TBO;HTTB;-5.07639;32.8333;3868;3;U;Africa/Dar_es_Salaam
5707;Shinyanga Airport;Shinyanga;Tanzania;SHY;HTSY;-3.667;33.417;3800;3;U;Africa/Dar_es_Salaam
5706;Musoma Airport;Musoma;Tanzania;MUZ;HTMU;-1.483;33.8;3806;3;U;Africa/Dar_es_Salaam
5705;Kikwetu Airport;Lindi;Tanzania;LDI;HTLI;-9.85111;39.7578;100;3;U;Africa/Dar_es_Salaam
5704;Kigoma Airport;Kigoma;Tanzania;TKQ;HTKA;-4.883;29.633;2700;3;U;Africa/Dar_es_Salaam
5703;Bukoba Airport;Bukoba;Tanzania;BKZ;HTBU;-1.3;31.8;3745;3;U;Africa/Dar_es_Salaam
5702;Port Sudan New International Airport;Port Sudan;Sudan;PZU;HSPN;19.4336;37.2341;135;3;U;Africa/Khartoum
5701;Nyala Airport;Nyala;Sudan;UYL;HSNN;12.0535;24.9562;2106;3;U;Africa/Khartoum
5700;Atbara Airport;Atbara;Sudan;ATB;HSAT;17.7;33.967;1181;3;U;Africa/Khartoum
5699;La Abraq Airport;Al Bayda';Libya;LAQ;HLLQ;32.788673;21.964333;2157;2;N;Africa/Tripoli
5698;Mitiga Airport;Tripoli;Libya;MJI;HLLM;32.8941;13.276;36;2;N;Africa/Tripoli
5697;Gamal Abdel Nasser Airport;Tobruk;Libya;TOB;HLGN;31.861;23.907;519;2;N;Africa/Tripoli
5696;Gardabya Airport;Sirt;Libya;SRX;HLGD;31.0635;16.595;267;2;N;Africa/Tripoli
5695;Nanyuki Civil Airport;Nanyuki;Kenya;NYK;HKNY;-0.067;37.033;6250;3;U;Africa/Nairobi
5694;Malindi Airport;Malindi;Kenya;MYD;HKML;-3.22931;40.1017;80;3;U;Africa/Nairobi
5693;Lokichoggio Airport;Lokichoggio;Kenya;LKG;HKLK;4.204117;34.348186;2074;3;U;Africa/Nairobi
5692;Amboseli Airport;Amboseli National Park;Kenya;ASV;HKAM;-2.64505;37.2531;3755;3;U;Africa/Nairobi
5691;Asyut International Airport;Asyut;Egypt;ATZ;HEAT;27.046508;31.011983;772;2;U;Africa/Cairo
5690;El Arish International Airport;El Arish;Egypt;AAC;HEAR;31.073333;33.835833;121;2;U;Africa/Cairo
5689;Burao Airport;Burao;Somalia;BUO;HCMV;9.517;45.567;3400;3;U;Africa/Mogadishu
5688;Galcaio Airport;Galcaio;Somalia;GLK;HCMR;6.78083;47.4547;975;3;U;Africa/Mogadishu
5687;Aden Adde International Airport;Mogadishu;Somalia;MGQ;HCMM;2.01444;45.3047;29;3;U;Africa/Mogadishu
5686;Bosaso Airport;Bosaso;Somalia;BSA;HCMF;11.2753;49.1494;3;3;U;Africa/Mogadishu
5685;Alula Airport;Alula;Somalia;ALU;HCMA;11.95;50.733;6;3;U;Africa/Mogadishu
5684;Tippi Airport;Tippi;Ethiopia;TIE;HATP;7.117;35.383;1100;3;U;Africa/Addis_Ababa
5683;Mizan Teferi Airport;Mizan Teferi;Ethiopia;MTF;HAMT;6.967;35.533;4396;3;U;Africa/Addis_Ababa
5682;Kabri Dehar Airport;Kabri Dehar;Ethiopia;ABK;HAKD;6.734;44.253;1800;3;U;Africa/Addis_Ababa
5681;Gore Airport;Gore;Ethiopia;GOR;HAGR;8.167;35.55;6580;3;U;Africa/Addis_Ababa
5680;Gode Airport;Gode;Ethiopia;GDE;HAGO;5.935128;43.578567;834;3;U;Africa/Addis_Ababa
5679;Dembidollo Airport;Dembidollo;Ethiopia;DEM;HADD;8.554;34.858;5200;3;U;Africa/Addis_Ababa
5678;Combolcha Airport;Dessie;Ethiopia;DSE;HADC;11.0825;39.7114;6117;3;U;Africa/Addis_Ababa
5677;Beica Airport;Beica;Ethiopia;BEI;HABE;9.38639;34.5219;5410;3;U;Africa/Addis_Ababa
5676;Baco Airport;Baco;Ethiopia;BCO;HABC;5.78287;36.562;0;3;U;Africa/Addis_Ababa
5675;Sao Filipe Airport;Sao Filipe, Fogo Island;Cape Verde;SFL;GVSF;14.885;-24.48;617;-1;U;Atlantic/Cape_Verde
5674;Praia International Airport;Praia, Santiago Island;Cape Verde;RAI;GVNP;14.9245;-23.4935;230;-1;U;Atlantic/Cape_Verde
5673;El Aroui Airport;El Aroui;Morocco;NDR;GMMW;34.9888;-3.02821;574;0;N;Africa/Casablanca
5672;Hassan I Airport;El Aaiún;Western Sahara;EUN;GMML;27.1517;-13.2192;207;0;N;Africa/El_Aaiun
5671;Mogador Airport;Essadouira;Morocco;ESU;GMMI;31.3975;-9.681667;384;0;N;Africa/Casablanca
5670;Dakhla Airport;Dakhla;Western Sahara;VIL;GMMH;23.7183;-15.932;36;0;N;Africa/El_Aaiun
5669;Smara Airport;Smara;Western Sahara;SMW;GMMA;26.7318;-11.6847;350;0;N;Africa/El_Aaiun
5668;Iginniarfik Heliport;Iginniarfik;Greenland;QFI;;68.7;-52.9667;0;-3;E;America/Godthab
5667;Akunnaaq Heliport;Akunnaaq;Greenland;QCU;;68.75;-52.3333;0;-3;E;America/Godthab
5666;Groennedal Heliport;Groennedal;Greenland;JGR;;61.2333;-48.1;0;-3;E;America/Godthab
5665;Osvaldo Vieira International Airport;Bissau;Guinea-Bissau;OXB;GGOV;11.89485;-15.653681;129;0;N;Africa/Bissau
5664;Kenema Airport;Kenema;Sierra Leone;KEN;GFKE;7.89129;-11.1766;485;0;N;Africa/Freetown
5663;Bo Airport;Bo;Sierra Leone;KBS;GFBO;7.9444;-11.761;328;0;N;Africa/Freetown
5662;Sherbro International Airport;Bonthe;Sierra Leone;BTE;GFBN;7.53242;-12.5189;14;0;N;Africa/Freetown
5661;Ceuta Heliport;Ceuta;Spain;JCU;GECT;35.8969;-5.29908;0;1;E;Europe/Madrid
5660;Lauriston Airport;Carriacou Island;Grenada;CRU;;12.4761;-61.4728;0;-4;U;America/Grenada
5659;La Gomera Airport;La Gomera;Spain;GMZ;GCGM;28.0296;-17.2146;718;0;E;Atlantic/Canary
5658;Papa Stour Airport;Papa Stour;United Kingdom;PSV;;60.3217;-1.69306;0;0;E;Europe/London
5657;Outer Skerries Airport;Outer Skerries;United Kingdom;OUK;;60.417;-0.75;0;0;E;Europe/London
5656;Foula Airport;Foula;United Kingdom;FOA;;60.121;-2.052;0;0;E;Europe/London
5655;Gamba;Gamba;Gabon;GAX;;-2.71016;9.96062;73;1;N;Africa/Libreville
5654;Ilebo Airport;Ilebo;Congo (Kinshasa);PFR;FZVS;-4.333;20.583;1450;2;N;Africa/Lubumbashi
5653;Lodja Airport;Lodja;Congo (Kinshasa);LJA;FZVA;-3.417;23.45;1647;2;N;Africa/Lubumbashi
5652;Tshikapa Airport;Tshikapa;Congo (Kinshasa);TSH;FZUK;-6.43833;20.7947;1595;2;N;Africa/Lubumbashi
5651;Basankusu Airport;Basankusu;Congo (Kinshasa);BSU;FZEN;1.22472;19.7889;1217;1;N;Africa/Kinshasa
5650;Basango Mboliasa Airport;Kiri;Congo (Kinshasa);KRZ;FZBT;-1.435;19.024;1013;1;N;Africa/Kinshasa
5649;Nioki Airport;Nioki;Congo (Kinshasa);NIO;FZBI;-2.7175;17.6847;1043;1;N;Africa/Kinshasa
5648;Inongo Airport;Inongo;Congo (Kinshasa);INO;FZBA;-1.94722;18.2858;1040;1;N;Africa/Kinshasa
5647;Tshimpi Airport;Matadi;Congo (Kinshasa);MAT;FZAM;-5.79961;13.4404;1115;1;N;Africa/Kinshasa
5646;Boma Airport;Boma;Congo (Kinshasa);BOA;FZAJ;-5.854;13.064;26;1;N;Africa/Kinshasa
5645;Eros Airport;Windhoek;Namibia;ERS;FYWE;-22.6122;17.0804;5575;1;S;Africa/Windhoek
5644;Swakopmund Airport;Swakopmund;Namibia;SWP;FYSM;-22.6619;14.5681;207;1;S;Africa/Windhoek
5643;Oranjemund Airport;Oranjemund;Namibia;OMD;FYOG;-28.5847;16.4467;14;1;S;Africa/Windhoek
5642;Ondangwa Airport;Ondangwa;Namibia;OND;FYOA;-17.8782;15.9526;3599;1;S;Africa/Windhoek
5641;Luderitz Airport;Luderitz;Namibia;LUD;FYLZ;-26.6874;15.2429;457;1;S;Africa/Windhoek
5640;Club Makokola Airport;Club Makokola;Malawi;CMK;FWCM;-14.3069;35.1325;1587;2;U;Africa/Blantyre
5639;Sarh Airport;Sarh;Chad;SRH;FTTA;9.14444;18.3744;1198;1;N;Africa/Ndjamena
5638;Croisette Heliport;Cannes;France;JCA;;43.536;7.03736;0;1;E;Europe/Paris
5637;Chimoio Airport;Chimoio;Mozambique;VPY;FQCH;-19.151267;33.428958;2287;2;U;Africa/Maputo
5636;Tchibanga Airport;Tchibanga;Gabon;TCH;FOOT;-2.85;11.017;269;1;N;Africa/Libreville
5635;Mouilla Ville Airport;Mouila;Gabon;MJL;FOGM;-1.84514;11.0567;295;1;N;Africa/Libreville
5634;Koulamoutou Airport;Koulamoutou;Gabon;KOU;FOGK;-1.18461;12.4413;1070;1;N;Africa/Libreville
5633;Namibe Airport;Mocamedes;Angola;MSZ;FNMO;-15.261222;12.146756;210;1;N;Africa/Luanda
5632;Ondjiva Pereira Airport;Ondjiva;Angola;VPE;FNGI;-17.043464;15.683822;3566;1;N;Africa/Luanda
5631;Dundo Airport;Dundo;Angola;DUE;FNDU;-7.40089;20.8185;2451;1;N;Africa/Luanda
5630;Catumbela Airport;Catumbela;Angola;CBT;FNCT;-12.4792;13.4869;0;1;N;Africa/Luanda
5629;Manja Airport;Manja;Madagascar;MJA;FMSJ;-21.417;44.317;787;3;U;Indian/Antananarivo
5628;Mandritsara Airport;Mandritsara;Madagascar;WMA;FMNX;-15.817;48.833;1007;3;U;Indian/Antananarivo
5627;Tsaratanana Airport;Tsaratanana;Madagascar;TTS;FMNT;-16.75;47.617;1073;3;U;Indian/Antananarivo
5626;Mampikony Airport;Mampikony;Madagascar;WMP;FMNP;-16.049;47.622;0;3;U;Indian/Antananarivo
5625;Soalala Airport;Soalala;Madagascar;DWB;FMNO;-16.083;45.367;141;3;U;Indian/Antananarivo
5624;Ambanja Airport;Ambanja;Madagascar;IVA;FMNJ;-13.65;48.467;36;3;U;Indian/Antananarivo
5623;Port Bergé Airport;Port Bergé;Madagascar;WPB;FMNG;-15.583;47.617;213;3;U;Indian/Antananarivo
5622;Ambatondrazaka Airport;Ambatondrazaka;Madagascar;WAM;FMMZ;-17.8;48.433;2513;3;U;Indian/Antananarivo
5621;Tsiroanomandidy Airport;Tsiroanomandidy;Madagascar;WTS;FMMX;-18.75;46.05;2776;3;U;Indian/Antananarivo
5620;Tambohorano Airport;Tambohorano;Madagascar;WTA;FMMU;-17.4761;43.9728;23;3;U;Indian/Antananarivo
5619;Morafenobe Airport;Morafenobe;Madagascar;TVA;FMMR;-17.85;44.917;748;3;U;Indian/Antananarivo
5618;Maintirano Airport;Maintirano;Madagascar;MXT;FMMO;-18.05;44.033;95;3;U;Indian/Antananarivo
5617;Belo sur Tsiribihina Airport;Belo sur Tsiribihina;Madagascar;BMD;FMML;-19.6867;44.5419;154;3;U;Indian/Antananarivo
5616;Ankavandra Airport;Ankavandra;Madagascar;JVA;FMMK;-18.8;45.283;427;3;U;Indian/Antananarivo
5615;Antsalova Airport;Antsalova;Madagascar;WAQ;FMMG;-18.7;44.617;551;3;U;Indian/Antananarivo
5614;Iconi Airport;Moroni;Comoros;YVA;FMCN;-11.7108;43.2439;33;3;U;Indian/Comoro
5613;Solwesi Airport;Solwesi;Zambia;SLI;FLSW;-12.1737;26.3651;4551;2;U;Africa/Lusaka
5612;Chipata Airport;Chipata;Zambia;CIP;FLCP;-13.5583;32.5872;3360;2;U;Africa/Lusaka
5611;Loubomo Airport;Loubomo;Congo (Brazzaville);DIS;FCPL;-4.2;12.7;1079;1;N;Africa/Brazzaville
5610;Limpopo Valley Airport;Tuli Lodge;Botswana;TLD;FBTL;-22.1892;29.1269;1772;2;U;Africa/Gaborone
5609;Shakawe Airport;Shakawe;Botswana;SWX;FBSW;-18.3739;21.8326;3379;2;U;Africa/Gaborone
5608;Orapa Airport;Orapa;Botswana;ORP;FBOR;-21.2667;25.3167;3100;2;U;Africa/Gaborone
5607;Ghanzi Airport;Ghanzi;Botswana;GNZ;FBGZ;-21.6925;21.6581;3730;2;U;Africa/Gaborone
5606;Mmabatho International Airport;Mafeking;South Africa;MBD;FAMM;-25.798444;25.548028;4181;2;U;Africa/Johannesburg
5605;Malamala Airport;Malamala;South Africa;AAM;FAMD;-24.818111;31.544584;1124;2;U;Africa/Johannesburg
5604;Kruger Mpumalanga International Airport;Mpumalanga;South Africa;MQP;FAKN;-25.3832;31.1056;2829;2;U;Africa/Johannesburg
5603;Rand Airport;Johannesburg;South Africa;QRA;FAGM;-26.242506;28.151169;5483;2;U;Africa/Johannesburg
5602;Ventspils International Airport;Ventspils;Latvia;VTS;EVVA;57.357778;21.544167;19;2;E;Europe/Riga
5601;Shire Inda Selassie Airport;Shire Indasilase;Ethiopia;SHC;;14.0781;38.2725;6207;3;U;Africa/Addis_Ababa
5600;Shilavo Airport;Shilavo;Ethiopia;HIL;;6.08333;44.7667;1296;3;U;Africa/Addis_Ababa
5599;Hemavan Airport;Hemavan;Sweden;HMV;ESUT;65.806111;15.082778;1503;1;E;Europe/Stockholm
5598;Storuman Airport;Mohed;Sweden;SQO;ESUD;64.960894;17.696583;915;1;E;Europe/Stockholm
5597;Ängelholm-Helsingborg Airport;Ängelholm;Sweden;AGH;ESTA;56.2961;12.8471;68;1;E;Europe/Stockholm
5596;Torsby Airport;Torsby;Sweden;TYF;ESST;60.157622;12.991269;393;1;E;Europe/Stockholm
5595;Karlstad Airport;Karlstad;Sweden;KSD;ESOK;59.4447;13.3374;352;1;E;Europe/Stockholm
5594;Hagfors Airport;Hagfors;Sweden;HFS;ESOH;60.020064;13.578908;474;1;E;Europe/Stockholm
5593;Östersund Airport;Östersund;Sweden;OSD;ESNZ;63.1944;14.5003;1233;1;E;Europe/Stockholm
5592;Łódź Władysław Reymont Airport;Lodz;Poland;LCJ;EPLL;51.721881;19.398133;604;1;E;Europe/Warsaw
5591;Bydgoszcz Ignacy Jan Paderewski Airport;Bydgoszcz;Poland;BZG;EPBY;53.0968;17.9777;235;1;E;Europe/Warsaw
5590;Værøy Heliport;Værøy;Norway;VRY;ENVR;67.6667;12.6833;36;1;E;Europe/Oslo
5589;Svartnes Airport;Vardø;Norway;VAW;ENSS;70.355392;31.044889;42;1;E;Europe/Oslo
5588;Sorkjosen Airport;Sorkjosen;Norway;SOJ;ENSR;69.7868;20.9594;16;1;E;Europe/Oslo
5587;Svolvær Helle Airport;Svolvær;Norway;SVJ;ENSH;68.2433;14.6692;27;1;E;Europe/Oslo
5586;Sogndal Airport;Sogndal;Norway;SOG;ENSG;61.1561;7.13778;0;1;E;Europe/Oslo
5585;Anda Airport;Sandane;Norway;SDN;ENSD;61.83;6.10583;196;1;E;Europe/Oslo
5584;Røst Airport;Røst;Norway;RET;ENRS;67.5278;12.1033;7;1;E;Europe/Oslo
5583;Ryum Airport;Rørvik;Norway;RVK;ENRM;64.8383;11.1461;14;1;E;Europe/Oslo
5582;Røssvoll Airport;Mo i Rana;Norway;MQN;ENRA;66.3639;14.3014;229;1;E;Europe/Oslo
5581;Namsos Høknesøra Airport;Namsos;Norway;OSY;ENNM;64.4722;11.5786;7;1;E;Europe/Oslo
5580;Leknes Airport;Leknes;Norway;LKN;ENLK;68.1525;13.6094;78;1;E;Europe/Oslo
5579;Sindal Airport;Sindal;Denmark;CNL;EKSN;57.503525;10.229372;92;1;N;Europe/Copenhagen
5578;Weston Airport;Leixlip;Ireland;;EIWT;53.351333;-6.4875;150;0;E;Europe/Dublin
5577;Donegal Airport;Dongloe;Ireland;CFN;EIDL;55.044192;-8.341;30;0;E;Europe/Dublin
6816;Raron Airport;Raron;Switzerland;;LSTA;46.3036;7.82333;2029;1;E;Europe/Zurich
5575;Barra Airport;Barra;United Kingdom;BRR;EGPR;57.0228;-7.44306;5;0;E;Europe/London
5574;Anglesey Airport;Angelsey;United Kingdom;HLY;EGOV;53.248097;-4.535339;37;0;E;Europe/London
5573;Penzance Heliport;Penzance;United Kingdom;PZE;EGHK;50.1281;-5.51845;14;0;E;Europe/London
5572;Land's End / St. Just Airport;Land's End;United Kingdom;LEQ;EGHC;50.1028;-5.67056;401;0;E;Europe/London
5571;Westray Airport;Westray;United Kingdom;WRY;EGEW;59.3503;-2.95;29;0;E;Europe/London
5570;Lerwick / Tingwall Airport;Lerwick;United Kingdom;LWK;EGET;60.1922;-1.24361;43;0;E;Europe/London
5569;Sanday Airport;Sanday;United Kingdom;NDY;EGES;59.2503;-2.57667;68;0;E;Europe/London
5568;Stronsay Airport;Stronsay;United Kingdom;SOY;EGER;59.1553;-2.64139;39;0;E;Europe/London
5567;Papa Westray Airport;Papa Westray;United Kingdom;PPW;EGEP;59.3517;-2.90028;91;0;E;Europe/London
5566;North Ronaldsay Airport;North Ronaldsay;United Kingdom;NRL;EGEN;59.3675;-2.43444;40;0;E;Europe/London
5565;Fair Isle Airport;Fair Isle;United Kingdom;FIE;EGEF;59.5358;-1.62806;223;0;E;Europe/London
5564;Eday Airport;Eday;United Kingdom;EOI;EGED;59.1906;-2.77222;10;0;E;Europe/London
5563;Campbeltown Airport;Campbeltown;United Kingdom;CAL;EGEC;55.4372;-5.68639;42;0;E;Europe/London
5562;Robin Hood Doncaster Sheffield Airport;Doncaster, Sheffield;United Kingdom;DSA;EGCN;53.474722;-1.004444;55;0;E;Europe/London
5561;Nottingham Airport;Nottingham;United Kingdom;NQT;EGBN;52.92;-1.079167;138;0;E;Europe/London
5560;Seinäjoki Airport;Seinäjoki / Ilmajoki;Finland;SJY;EFSI;62.6921;22.8323;302;2;E;Europe/Helsinki
5559;Helgoland-Düne Airport;Helgoland;Germany;HGL;EDXH;54.1853;7.91583;7;1;E;\N
5558;Heide-Büsum Airport;Büsum;Germany;HEI;EDXB;54.1533;8.90167;7;1;E;Europe/Berlin
5557;Heringsdorf Airport;Heringsdorf;Germany;HDF;EDAH;53.878706;14.152347;93;1;E;Europe/Berlin
5556;Kumasi Airport;Kumasi;Ghana;KMS;DGSI;6.71456;-1.59082;942;0;N;Africa/Accra
5555;Guemar Airport;Guemar;Algeria;ELU;DAUO;33.5114;6.77679;203;1;N;Africa/Algiers
5554;Bordj Badji Mokhtar Airport;Bordj Badji Mokhtar;Algeria;BMW;DATM;21.375;0.923889;1303;1;N;Africa/Algiers
5553;Béchar Boudghene Ben Ali Lotfi Airport;Béchar;Algeria;CBH;DAOR;31.6457;-2.26986;2661;1;N;Africa/Algiers
5552;Batna Airport;Batna;Algeria;BLJ;DABT;35.752106;6.308589;2697;1;N;Africa/Algiers
5551;Wollaston Lake Airport;Wollaston Lake;Canada;ZWL;CZWL;58.1069;-103.172;1360;-6;N;America/Regina
5550;Churchill Falls Airport;Churchill Falls;Canada;ZUM;CZUM;53.5619;-64.1064;1442;-4;A;America/Halifax
5549;Shamattawa Airport;Shamattawa;Canada;ZTM;CZTM;55.8656;-92.0814;289;-6;A;America/Winnipeg
5548;Sandy Lake Airport;Sandy Lake;Canada;ZSJ;CZSJ;53.0642;-93.3444;951;-6;A;America/Winnipeg
5547;Round Lake (Weagamow Lake) Airport;Round Lake;Canada;ZRJ;CZRJ;52.9436;-91.3128;974;-6;A;America/Winnipeg
5546;Sachigo Lake Airport;Sachigo Lake;Canada;ZPB;CZPB;53.8911;-92.1964;876;-6;A;America/Winnipeg
5545;Masset Airport;Masset;Canada;ZMT;CZMT;54.0275;-132.125;25;-8;A;America/Vancouver
5544;Muskrat Dam Airport;Muskrat Dam;Canada;MSA;CZMD;53.4414;-91.7628;911;-6;A;America/Winnipeg
5543;Kashechewan Airport;Kashechewan;Canada;ZKE;CZKE;52.2825;-81.6778;35;-5;A;America/Toronto
5542;Swan River Airport;Swan River;Canada;ZJN;CZJN;52.1206;-101.236;1100;-6;A;America/Winnipeg
5541;Gods River Airport;Gods River;Canada;ZGI;CZGI;54.8397;-94.0786;627;-6;A;America/Winnipeg
5540;Fond-Du-Lac Airport;Fond-Du-Lac;Canada;ZFD;CZFD;59.3344;-107.182;814;-6;N;America/Regina
5539;Eastmain River Airport;Eastmain River;Canada;ZEM;CZEM;52.2264;-78.5225;24;-5;A;America/Toronto
5538;Bathurst Airport;Bathurst;Canada;ZBF;CZBF;47.6297;-65.7389;193;-4;A;America/Halifax
5537;Ilford Airport;Ilford;Canada;ILF;CZBD;56.0614;-95.6139;642;-6;A;America/Winnipeg
5536;York Landing Airport;York Landing;Canada;ZAC;CZAC;56.0894;-96.0892;621;-6;A;America/Winnipeg
5535;Salluit Airport;Salluit;Canada;YZG;CYZG;62.1794;-75.6672;743;-5;A;America/Toronto
5534;Whale Cove Airport;Whale Cove;Canada;YXN;CYXN;62.24;-92.5981;40;-6;A;America/Winnipeg
5533;Webequie Airport;Webequie;Canada;YWP;CYWP;52.9597;-87.3689;685;-5;A;America/Toronto
5532;Deer Lake Airport;Deer Lake;Canada;YVZ;CYVZ;52.6558;-94.0614;1092;-6;A;America/Winnipeg
5531;Big Trout Lake Airport;Big Trout Lake;Canada;YTL;CYTL;53.8178;-89.8969;729;-6;A;America/Winnipeg
5530;St. Theresa Point Airport;St. Theresa Point;Canada;YST;CYST;53.8456;-94.8519;773;-6;A;America/Winnipeg
5529;Sanikiluaq Airport;Sanikiluaq;Canada;YSK;CYSK;56.5378;-79.2467;104;-5;A;\N
5528;Stony Rapids Airport;Stony Rapids;Canada;YSF;CYSF;59.2503;-105.841;805;-6;N;America/Regina
5527;Red Lake Airport;Red Lake;Canada;YRL;CYRL;51.0669;-93.7931;1265;-6;A;America/Winnipeg
5526;Rae Lakes Airport;Gamètì;Canada;YRA;CYRA;64.1161;-117.31;723;-7;A;America/Edmonton
5525;Nakina Airport;Nakina;Canada;YQN;CYQN;50.182777;-86.696388;1057;-5;A;America/Toronto
5524;The Pas Airport;The Pas;Canada;YQD;CYQD;53.9714;-101.091;887;-6;A;America/Winnipeg
5523;Powell River Airport;Powell River;Canada;YPW;CYPW;49.8342;-124.5;425;-8;A;America/Vancouver
5522;Peawanuck Airport;Peawanuck;Canada;YPO;CYPO;54.9881;-85.4433;173;-5;A;America/Toronto
5521;Pikangikum Airport;Pikangikum;Canada;YPM;CYPM;51.8197;-93.9733;1114;-6;A;America/Winnipeg
5520;Inukjuak Airport;Inukjuak;Canada;YPH;CYPH;58.4719;-78.0769;83;-5;A;America/Toronto
5519;Oxford House Airport;Oxford House;Canada;YOH;CYOH;54.9333;-95.2789;663;-6;A;America/Winnipeg
5518;Points North Landing Airport;Points North Landing;Canada;YNL;CYNL;58.2767;-104.082;1605;-6;N;America/Regina
5517;Norway House Airport;Norway House;Canada;YNE;CYNE;53.9583;-97.8442;734;-6;A;America/Winnipeg
5516;Wemindji Airport;Wemindji;Canada;YNC;CYNC;53.0106;-78.8311;66;-5;A;America/Toronto
5515;Umiujaq Airport;Umiujaq;Canada;YUD;CYMU;56.5361;-76.5183;250;-5;A;America/Toronto
5514;Chapais Airport;Chibougamau;Canada;YMT;CYMT;49.7719;-74.5281;1270;-5;A;America/Toronto
5513;Mary's Harbour Airport;Mary's Harbour;Canada;YMH;CYMH;52.3028;-55.8472;38;-3.5;A;America/St_Johns
5512;Kangiqsualujjuaq (Georges River) Airport;Kangiqsualujjuaq;Canada;XGR;CYLU;58.7114;-65.9928;215;-5;A;America/Toronto
5511;Lutselk'e Airport;Lutselk'e;Canada;YSG;CYLK;62.4183;-110.682;596;-7;A;America/Edmonton
5510;Lansdowne House Airport;Lansdowne House;Canada;YLH;CYLH;52.1956;-87.9342;834;-5;A;America/Toronto
5509;Kimmirut Airport;Kimmirut;Canada;YLC;CYLC;62.85;-69.8833;175;-5;A;America/Toronto
5508;Aupaluk Airport;Aupaluk;Canada;YPJ;CYLA;59.2967;-69.5997;119;-5;A;America/Toronto
5507;Waskaganish Airport;Waskaganish;Canada;YKQ;CYKQ;51.4733;-78.7583;80;-5;A;America/Toronto
5506;Akulivik Airport;Akulivik;Canada;AKV;CYKO;60.8186;-78.1486;75;-5;A;America/Toronto
5505;Island Lake Airport;Island Lake;Canada;YIV;CYIV;53.8572;-94.6536;770;-6;A;America/Winnipeg
5504;Ivujivik Airport;Ivujivik;Canada;YIK;CYIK;62.4173;-77.9253;126;-5;A;America/Toronto
5503;Chevery Airport;Chevery;Canada;YHR;CYHR;50.4689;-59.6367;39;-4;A;America/Blanc-Sablon
5502;Hopedale Airport;Hopedale;Canada;YHO;CYHO;55.4483;-60.2286;39;-4;A;America/Halifax
5501;Nemiscau Airport;Nemiscau;Canada;YNS;CYHH;51.6911;-76.1356;802;-5;A;America/Toronto
5500;Vancouver Harbour Water Airport;Vancouver;Canada;YHC;CYHC;49.2944;-123.111;0;-8;A;America/Vancouver
5499;Quaqtaq Airport;Quaqtaq;Canada;YQC;CYHA;61.0464;-69.6178;103;-5;A;America/Toronto
5498;Grise Fiord Airport;Grise Fiord;Canada;YGZ;CYGZ;76.4261;-82.9092;146;-5;A;America/Toronto
5497;Gillam Airport;Gillam;Canada;YGX;CYGX;56.3575;-94.7106;476;-6;A;America/Winnipeg
5496;Kuujjuarapik Airport;Kuujjuarapik;Canada;YGW;CYGW;55.2819;-77.7653;34;-5;A;America/Toronto
5495;Igloolik Airport;Igloolik;Canada;YGT;CYGT;69.3647;-81.8161;174;-5;A;America/Toronto
5494;Gods Lake Narrows Airport;Gods Lake Narrows;Canada;YGO;CYGO;54.5589;-94.4914;617;-6;A;America/Winnipeg
5493;Texada Gillies Bay Airport;Texada;Canada;YGB;CYGB;49.6942;-124.518;326;-8;A;America/Vancouver
5492;Makkovik Airport;Makkovik;Canada;YMN;CYFT;55.0769;-59.1864;234;-4;A;America/Halifax
5491;Fort Hope Airport;Fort Hope;Canada;YFH;CYFH;51.5619;-87.9078;899;-5;A;America/Toronto
5490;Fort Albany Airport;Fort Albany;Canada;YFA;CYFA;52.2014;-81.6969;48;-5;A;America/Toronto
5489;Fort Severn Airport;Fort Severn;Canada;YER;CYER;56.0189;-87.6761;48;-5;A;America/Toronto
5488;Nain Airport;Nain;Canada;YDP;CYDP;56.5492;-61.6803;22;-4;A;America/Halifax
5487;Chesterfield Inlet Airport;Chesterfield Inlet;Canada;YCS;CYCS;63.3469;-90.7311;32;-6;A;America/Winnipeg
5486;Cartwright Airport;Cartwright;Canada;YRF;CYCA;53.6828;-57.0419;40;-4;A;America/Halifax
5485;Lourdes De Blanc Sablon Airport;Lourdes-De-Blanc-Sablon;Canada;YBX;CYBX;51.4436;-57.1853;121;-4;A;America/Blanc-Sablon
5484;Uranium City Airport;Uranium City;Canada;YBE;CYBE;59.5614;-108.481;1044;-6;N;America/Regina
5483;Lac Du Bonnet Airport;Lac Du Bonnet;Canada;;CYAX;50.2944;-96.01;850;-6;A;America/Winnipeg
5482;Attawapiskat Airport;Attawapiskat;Canada;YAT;CYAT;52.9275;-82.4319;31;-5;A;America/Toronto
5481;Kangirsuk Airport;Kangirsuk;Canada;YKG;CYAS;60.0272;-69.9992;403;-5;A;America/Toronto
5480;Kasabonika Airport;Kasabonika;Canada;XKS;CYAQ;53.5247;-88.6428;672;-5;A;America/Toronto
5479;Fort Frances Municipal Airport;Fort Frances;Canada;YAG;CYAG;48.6542;-93.4397;1125;-6;A;America/Winnipeg
5478;Cat Lake Airport;Cat Lake;Canada;YAC;CYAC;51.7272;-91.8244;1344;-6;A;America/Winnipeg
5477;Tarapacá Airport;Tarapacá;Colombia;TCD;;-2.867;-69.733;0;-5;U;America/Bogota
5476;Apartadó Airport;Apartadó;Colombia;APO;;7.033;-77.2;0;-5;U;America/Bogota
5475;Nantong Airport;Nantong;China;NTG;;32.0708;120.976;0;8;U;Asia/Chongqing
5474;La Tabatière Airport;La Tabatière;Canada;ZLT;CTU5;50.8308;-58.9756;102;-4;A;America/Blanc-Sablon
5473;Tête-à-la-Baleine Airport;Tête-à-la-Baleine;Canada;ZTB;CTB6;50.6744;-59.3836;107;-4;A;America/Blanc-Sablon
5472;Chisasibi Airport;Chisasibi;Canada;YKU;CSU2;53.8056;-78.9169;43;-5;A;America/Toronto
5471;Poplar Hill Airport;Poplar Hill;Canada;YHP;CPV7;52.1133;-94.2556;1095;-6;A;America/Winnipeg
5470;Ogoki Post Airport;Ogoki Post;Canada;YOG;CNT3;51.6586;-85.9017;594;-5;A;America/Toronto
5469;Kingfisher Lake Airport;Kingfisher Lake;Canada;KIF;CNM5;53.0125;-89.8553;866;-5;A;America/Toronto
5468;Bearskin Lake Airport;Bearskin Lake;Canada;XBE;CNE3;53.9656;-91.0272;800;-6;A;America/Winnipeg
5467;North Spirit Lake Airport;North Spirit Lake;Canada;YNO;CKQ3;52.49;-92.9711;1082;-6;A;America/Winnipeg
5466;Wunnumin Lake Airport;Wunnumin Lake;Canada;WNN;CKL3;52.8939;-89.2892;819;-5;A;America/Toronto
5465;Wapekeka Airport;Angling Lake;Canada;YAX;CKB6;53.8492;-89.5794;712;-6;A;America/Winnipeg
5464;Summer Beaver Airport;Summer Beaver;Canada;SUR;CJV7;52.7086;-88.5419;832;-5;A;America/Toronto
5463;Whatì Airport;Whatì;Canada;YLE;CEM3;63.1317;-117.246;882;-7;A;America/Edmonton
5462;Colville Lake Airport;Colville Lake;Canada;YCK;CEB3;67.0333;-126.083;850;-7;A;America/Edmonton
5461;Rigolet Airport;Rigolet;Canada;YRG;CCZ2;54.1797;-58.4575;180;-4;A;America/Halifax
5460;Port Hope Simpson Airport;Port Hope Simpson;Canada;YHA;CCP4;52.5281;-56.2861;347;-3.5;A;America/St_Johns
5459;St. Lewis (Fox Harbour) Airport;St. Lewis;Canada;YFX;CCK4;52.3728;-55.6739;74;-3.5;A;America/St_Johns
5458;Williams Harbour Airport;Williams Harbour;Canada;YWM;CCA6;52.5669;-55.7847;70;-3.5;A;America/St_Johns
5457;Anahim Lake Airport;Anahim Lake;Canada;YAA;CAJ4;52.4525;-125.303;3635;-8;A;America/Vancouver
5456;Whistler/Green Lake Water Aerodrome;Whistler;Canada;YWS;CAE5;50.1436;-122.949;2100;-8;A;America/Vancouver
5455;Punta Gorda Airport;Punta Gorda;Belize;PND;;16.1024;-88.8083;7;-6;U;America/Belize
5454;Caye Caulker Airport;Caye Caulker;Belize;CUK;;17.7347;-88.0325;1;-6;U;America/Belize
5453;Vopnafjörður Airport;Vopnafjörður;Iceland;VPN;BIVO;65.7206;-14.8506;0;0;N;Atlantic/Reykjavik
5452;Thorshofn Airport;Thorshofn;Iceland;THO;BITN;66.2185;-15.3356;65;0;N;Atlantic/Reykjavik
6945;Drake Bay Airport;Puntarenas;Costa Rica;DRK;MRDK;8.71889;-83.6417;12;-6;N;America/Costa_Rica
5450;Grímsey Airport;Grímsey;Iceland;GRY;BIGR;66.5547;-18.0175;66;0;N;Atlantic/Reykjavik
5449;Qaarsut Airport;Uummannaq;Greenland;JQA;BGUQ;70.7342;-52.6962;289;-3;E;America/Godthab
5448;Upernavik Airport;Upernavik;Greenland;JUV;BGUK;72.7902;-56.1306;414;-3;E;America/Godthab
5447;Sisimiut Airport;Sisimiut;Greenland;JHS;BGSS;66.9513;-53.7293;33;-3;E;America/Godthab
5446;Qaanaaq Airport;Qaanaaq;Greenland;NAQ;BGQQ;77.4886;-69.3887;51;-4;E;America/Thule
5445;Narsaq Heliport;Narsaq;Greenland;JNS;BGNS;60.9167;-46.0586;0;-3;E;America/Godthab
5444;Nanortalik Heliport;Nanortalik;Greenland;JNN;BGNN;60.14;-45.2317;0;-3;E;America/Godthab
5443;Maniitsoq Airport;Maniitsoq;Greenland;JSU;BGMQ;65.4125;-52.9394;91;-3;E;America/Godthab
5442;Qaqortoq Heliport;Qaqortoq;Greenland;JJU;BGJH;60.7158;-46.0294;0;-3;E;America/Godthab
5441;Qeqertarsuaq Heliport;Qeqertarsuaq Airport;Greenland;JGO;BGGN;69.2511;-53.5381;0;-3;E;America/Godthab
5440;Paamiut Heliport;Paamiut;Greenland;JFR;BGFH;61.9922;-49.6625;0;-3;E;America/Godthab
5439;Neerlerit Inaat Airport;Neerlerit Inaat;Greenland;CNP;BGCO;70.7433;-22.6606;45;-1;E;America/Scoresbysund
5438;Alluitsup Paa Heliport;Alluitsup Paa;Greenland;LLU;BGAP;60.4644;-45.5778;89;-3;E;America/Godthab
5437;Wapenamanda Airport;Wapenamanda;Papua New Guinea;WBM;AYWD;-5.6433;143.895;5889;10;U;Pacific/Port_Moresby
5436;Vanimo Airport;Vanimo;Papua New Guinea;VAI;AYVN;-2.69717;141.302;10;10;U;Pacific/Port_Moresby
5435;Tokua Airport;Tokua;Papua New Guinea;RAB;AYTK;-4.34046;152.38;32;10;U;Pacific/Port_Moresby
5434;Tabubil Airport;Tabubil;Papua New Guinea;TBG;AYTB;-5.27861;141.226;1570;10;U;Pacific/Port_Moresby
5433;Tari Airport;Tari;Papua New Guinea;TIZ;AYTA;-5.845;142.948;5500;10;U;Pacific/Port_Moresby
5432;Misima Island Airport;Misima Island;Papua New Guinea;MIS;AYMS;-10.6892;152.838;26;10;U;Pacific/Port_Moresby
5431;Moro Airport;Moro;Papua New Guinea;MXH;AYMR;-6.36333;143.238;2740;10;U;Pacific/Port_Moresby
5430;Momote Airport;Momote;Papua New Guinea;MAS;AYMO;-2.06189;147.424;12;10;U;Pacific/Port_Moresby
5429;Mendi Airport;Mendi;Papua New Guinea;MDU;AYMN;-6.14774;143.657;5680;10;U;Pacific/Port_Moresby
5428;Kavieng Airport;Kavieng;Papua New Guinea;KVG;AYKV;-2.5794;150.808;7;10;U;Pacific/Port_Moresby
5427;Kerema Airport;Kerema;Papua New Guinea;KMA;AYKM;-7.96361;145.771;10;10;U;Pacific/Port_Moresby
5426;Kikori Airport;Kikori;Papua New Guinea;KRI;AYKK;-7.42438;144.25;50;10;U;Pacific/Port_Moresby
5425;Kiunga Airport;Kiunga;Papua New Guinea;UNG;AYKI;-6.12571;141.282;88;10;U;Pacific/Port_Moresby
5424;Kimbe Airport;Hoskins;Papua New Guinea;HKN;AYHK;-5.46217;150.405;66;10;U;Pacific/Port_Moresby
5423;Girua Airport;Girua;Papua New Guinea;PNP;AYGR;-8.80454;148.309;311;10;U;Pacific/Port_Moresby
5422;Gurney Airport;Gurney;Papua New Guinea;GUR;AYGN;-10.3115;150.334;88;10;U;Pacific/Port_Moresby
5421;Daru Airport;Daru;Papua New Guinea;DAU;AYDU;-9.08676;143.208;20;10;U;Pacific/Port_Moresby
5420;Chimbu Airport;Kundiawa;Papua New Guinea;CMU;AYCH;-6.02429;144.971;4974;10;U;Pacific/Port_Moresby
5419;Buka Airport;Buka Island;Papua New Guinea;BUA;AYBK;-5.42232;154.673;11;10;U;Pacific/Port_Moresby
5418;Ramata Airport;Ramata;Solomon Islands;RBV;AGRM;-8.16806;157.643;0;11;U;Pacific/Guadalcanal
5417;Kagau Island Airport;Kagau Island;Solomon Islands;KGE;AGKG;-7.333;157.583;0;11;U;Pacific/Guadalcanal
5416;Suavanao Airport;Suavanao;Solomon Islands;VAO;AGGV;-7.58556;158.731;0;11;U;Pacific/Guadalcanal
5415;Marau Airport;Marau;Solomon Islands;RUS;AGGU;-9.86167;160.825;0;11;U;Pacific/Guadalcanal
5414;Rennell/Tingoa Airport;Rennell Island;Solomon Islands;RNL;AGGR;-11.5339;160.063;0;11;U;Pacific/Guadalcanal
5413;Mono Airport;Stirling Island;Solomon Islands;MNY;AGGO;-7.41694;155.565;0;11;U;Pacific/Guadalcanal
5412;Nusatupe Airport;Gizo;Solomon Islands;GZO;AGGN;-8.09778;156.864;13;11;U;Pacific/Guadalcanal
5411;Munda Airport;Munda;Solomon Islands;MUA;AGGM;-8.32797;157.263;10;11;U;Pacific/Guadalcanal
5410;Santa Cruz/Graciosa Bay/Luova Airport;Santa Cruz/Graciosa Bay/Luova;Solomon Islands;SCZ;AGGL;-10.7203;165.795;18;11;U;Pacific/Guadalcanal
5409;Ngorangora Airport;Kirakira;Solomon Islands;IRA;AGGK;-10.4497;161.898;0;11;U;Pacific/Guadalcanal
5408;Babanakira Airport;Mbambanakira;Solomon Islands;MBU;AGGI;-9.7475;159.839;0;11;U;Pacific/Guadalcanal
5407;Fera/Maringe Airport;Fera Island;Solomon Islands;FRE;AGGF;-8.1075;159.577;0;11;U;Pacific/Guadalcanal
5406;Ballalae Airport;Ballalae;Solomon Islands;BAS;AGGE;-6.967;155.883;5;11;U;Pacific/Guadalcanal
5405;Auki Airport;Auki;Solomon Islands;AKS;AGGA;-8.70257;160.682;5;11;U;Pacific/Guadalcanal
5404;Uru Harbour Airport;Atoifi;Solomon Islands;ATD;AGAT;-8.87333;161.011;0;11;U;Pacific/Guadalcanal
6097;Komsomolsk-on-Amur Airport;Komsomolsk-on-Amur;Russia;KXK;UHKK;50.4094;136.934;92;11;N;Asia/Vladivostok
6096;Moyo Airport;Moyo;Uganda;OYG;;3.633;31.75;0;3;U;Africa/Kampala
6095;Tiksi Airport;Tiksi;Russia;IKS;UEST;71.6977;128.903;26;10;N;Asia/Yakutsk
6094;Cherskiy Airport;Cherskiy;Russia;CYX;UESS;68.7406;161.338;20;12;N;Asia/Magadan
6093;Chokurdakh Airport;Chokurdah;Russia;CKH;UESO;70.6231;147.902;151;12;N;Asia/Magadan
6091;Chulman;Neryungri;Russia;CNN;UELL;56.9139;124.914;2812;10;N;Asia/Yakutsk
6088;Kostanay West Airport;Kostanay;Kazakhstan;KSN;UAUU;53.206944;63.550278;595;6;U;Asia/Qyzylorda
6085;Zhezkazgan Airport;Zhezkazgan;Kazakhstan;DZN;UAKD;47.708333;67.733333;1250;6;U;Asia/Qyzylorda
6083;Kokshetau Airport;Kokshetau;Kazakhstan;KOV;UACK;53.3291;69.5946;900;6;U;Asia/Qyzylorda
6082;Union Island International Airport;Union Island;Saint Vincent and the Grenadines;UNI;TVSU;12.583;-61.417;16;-4;U;America/St_Vincent
6081;J F Mitchell Airport;Bequia;Saint Vincent and the Grenadines;BQU;TVSB;12.988444;-61.262033;15;-4;U;America/St_Vincent
6080;Virgin Gorda Airport;Spanish Town;British Virgin Islands;VIJ;TUPW;18.4464;-64.4275;9;-4;U;America/Tortola
6079;Vance Winkworth Amory International Airport;Charlestown;Saint Kitts and Nevis;NEV;TKPN;17.205678;-62.589869;14;-4;U;America/St_Kitts
6076;Baillif Airport;Basse Terre;Guadeloupe;BBR;TFFB;16.0133;-61.7422;59;-4;U;America/Guadeloupe
6075;La Désirade Airport;Grande Anse;Guadeloupe;DSD;TFFA;16.2969;-61.0844;10;-4;U;America/Guadeloupe
6072;Juan Pablo Pérez Alfonso Airport;El Vigía;Venezuela;VIG;SVVG;8.6241;-71.672819;250;-4.5;U;America/Caracas
6070;El Jaguel / Punta del Este Airport;Maldonado;Uruguay;MDO;SUPE;-34.917;-54.917;66;-3;S;America/Montevideo
6069;Santa Rosa Airport;Santa Rosa;Brazil;SRA;SSZR;-27.9067;-54.5204;984;-3;S;America/Sao_Paulo
6067;Alferez Fap David Figueroa Fernandini Airport;Huánuco;Peru;HUU;SPNC;-9.878811;-76.204797;6070;-5;U;America/Lima
6065;Caballococha Airport;Caballococha;Peru;LHC;SPBC;-3.916858;-70.508225;328;-5;U;America/Lima
6060;Capitán Av. German Quiroga G. Airport;San Borja;Bolivia;SRJ;SLSB;-14.8592;-66.7375;633;-4;U;America/La_Paz
6058;Capitán Av. Selin Zeitun Lopez Airport;Riberalta;Bolivia;RIB;SLRI;-11;-66;462;-4;U;America/La_Paz
6056;Capitán de Av. Emilio Beltrán Airport;Guayaramerín;Bolivia;GYA;SLGY;-10.8206;-65.3456;557;-4;U;America/La_Paz
6054;Obando Airport;Puerto Inírida;Colombia;PDA;SKPD;3.85;-67.91;460;-5;U;America/Bogota
6053;Reyes Murillo Airport;Nuquí;Colombia;NQU;SKNQ;5.7;-77.28;12;-5;U;America/Bogota
6052;La Pedrera Airport;La Pedrera;Colombia;LPD;SKLP;-1.33;-69.58;590;-5;U;America/Bogota
6051;Caucaya Airport;Puerto Leguízamo;Colombia;LQM;SKLG;-0.18;-74.77;573;-5;U;America/Bogota
6050;La Jagua Airport;Garzón;Colombia;GLJ;SKGZ;2.17;-75.67;2620;-5;U;America/Bogota
6049;Santa Ana Airport;Cartago;Colombia;CRC;SKGO;4.758181;-75.955753;2979;-5;U;America/Bogota
6048;Stanley Airport;Stanley;Falkland Islands;PSY;SFAL;-51.685672;-57.777644;75;-3;U;Atlantic/Stanley
6046;Camilo Ponce Enriquez Airport;La Toma (Catamayo);Ecuador;LOH;SETM;-3.99589;-79.3719;4056;-5;U;America/Guayaquil
6045;San Cristóbal Airport;San Cristóbal;Ecuador;SCY;SEST;-0.910206;-89.61745;62;-6;U;Pacific/Galapagos
6044;Sorocaba Airport;Sorocaba;Brazil;SOD;SDCO;-23.478;-47.49;2077;-3;S;America/Sao_Paulo
6042;Ricardo García Posada Airport;El Salvador;Chile;ESR;SCES;-26.3111;-69.7652;5240;-4;S;America/Santiago
6039;Vitória da Conquista Airport;Vitória Da Conquista;Brazil;VDC;SBQV;-14.862761;-40.863106;2998;-3;S;America/Fortaleza
6038;Marília Airport;Marília;Brazil;MII;SBML;-22.196892;-49.9264;2122;-3;S;America/Sao_Paulo
6037;Macaé Airport;Macaé;Brazil;MEA;SBME;-22.343;-41.766;8;-3;S;America/Sao_Paulo
6033;Necochea Airport;Necochea;Argentina;NEC;SAZO;-38.483056;-58.817222;72;-3;N;America/Cordoba
6030;Lago Argentino Airport;El Calafate;Argentina;ING;SAWA;-50.3361;-72.2486;732;-3;N;America/Cordoba
6026;Roxas Airport;Roxas City;Philippines;RXS;RPVR;11.597669;122.751669;10;8;N;Asia/Manila
6023;Calbayog Airport;Calbayog City;Philippines;CYP;RPVC;12.072706;124.545092;12;8;N;Asia/Manila
6022;Virac Airport;Virac;Philippines;VRC;RPUV;13.576439;124.205672;121;8;N;Asia/Manila
6021;Tuguegarao Airport;Tuguegarao;Philippines;TUG;RPUT;17.638311;121.730614;70;8;N;Asia/Manila
6020;San Fernando Airport;San Fernando;Philippines;SFE;RPUS;16.595589;120.303219;13;8;N;Asia/Manila
6019;Basco Airport;Basco;Philippines;BSO;RPUO;20.451322;121.979883;291;8;N;\N
6018;Naga Airport;Naga;Philippines;WNP;RPUN;13.584886;123.270239;142;8;N;Asia/Manila
6017;Tandag Airport;Tandag;Philippines;TDG;RPMW;9.07211;126.171;16;8;N;Asia/Manila
6016;Surigao Airport;Sangley Point;Philippines;SUG;RPMS;9.757567;125.479328;20;8;N;Asia/Manila
6507;Fremont Airport;Fremont;United States;;\N;41.3331;-83.1612;663;-5;A;America/New_York
6014;Sanga Sanga Airport;Sanga Sanga;Philippines;SGS;RPMN;5.04699;119.743;15;8;N;Asia/Manila
6011;General Santos International Airport;General Santos City;Philippines;GES;RPMB;6.106439;125.2353;28;8;N;Asia/Manila
6009;Subic Bay International Airport;Olongapo City;Philippines;SFS;RPLB;14.7944;120.271;64;8;N;Asia/Manila
6006;Yangyang International Airport;Sokcho / Gangneung;South Korea;YNY;RKNY;38.061311;128.669164;241;9;U;Asia/Seoul
6098;Ugolny Airport;Anadyr;Russia;DYR;UHMA;64.73495;177.741483;194;12;N;Asia/Magadan
6099;Okhotsk Airport;Okhotsk;Russia;OHO;UHOO;59.4101;143.057;0;11;N;Asia/Vladivostok
6100;Ujae Atoll Airport;Ujae Atoll;Marshall Islands;UJE;UJAP;8.92806;165.762;0;12;U;Pacific/Majuro
6101;Mariupol International Airport;Mariupol International;Ukraine;MPW;UKCM;47.0761;37.4496;251;2;E;Europe/Kiev
6102;Luhansk International Airport;Lugansk;Ukraine;VSG;UKCW;48.4174;39.3741;636;2;E;Europe/Kiev
6103;Zaporizhzhia International Airport;Zaporozhye;Ukraine;OZH;UKDE;47.867;35.3157;373;2;E;Europe/Kiev
6104;Lozuvatka International Airport;Krivoy Rog;Ukraine;KWG;UKDR;48.0433;33.21;408;2;E;Europe/Kiev
6105;Osnova International Airport;Kharkov;Ukraine;HRK;UKHH;49.924786;36.289986;508;2;E;Europe/Kiev
6106;Ivano Frankivsk International Airport;Ivano-Frankivsk;Ukraine;IFO;UKLI;48.884167;24.686111;919;2;E;Europe/Kiev
6107;Chernivtsi International Airport;Chernovtsk;Ukraine;CWC;UKLN;48.259322;25.980831;826;2;E;Europe/Kiev
6108;Rivne International Airport;Rivne;Ukraine;RWN;UKLR;50.6071;26.1416;755;2;E;Europe/Kiev
6109;Uzhhorod International Airport;Uzhgorod;Ukraine;UDJ;UKLU;48.634278;22.263356;383;2;E;Europe/Kiev
6110;Solovki Airport;Solovetsky Islands;Russia;CSH;ULAS;65.03;35.7333;0;4;N;Europe/Moscow
6111;Cherepovets Airport;Cherepovets;Russia;CEE;ULBC;59.2736;38.0158;377;4;N;Europe/Moscow
6112;Amderma Airport;Amderma;Russia;AMV;ULDD;69.7633;61.5564;13;4;N;Europe/Moscow
6113;Kotlas Airport;Kotlas;Russia;KSZ;ULKK;61.2358;46.6975;184;4;N;Europe/Moscow
6114;Petrozavodsk Airport;Petrozavodsk;Russia;PES;ULPB;61.8852;34.1547;151;4;N;Europe/Moscow
6115;Hrodno Airport;Hrodna;Belarus;GNA;UMMG;53.602;24.0538;443;3;E;Europe/Minsk
6116;Mogilev Airport;Mogilev;Belarus;MVQ;UMOO;53.9549;30.0951;637;3;E;Europe/Minsk
6117;Yeniseysk Airport;Yeniseysk;Russia;EIE;UNII;58.4742;92.1125;253;8;N;Asia/Krasnoyarsk
6118;Kyzyl Airport;Kyzyl;Russia;KYZ;UNKY;51.6694;94.4006;2123;8;N;Asia/Krasnoyarsk
6119;Spichenkovo Airport;Novokuznetsk;Russia;NOZ;UNWW;53.8114;86.8772;1024;7;N;Asia/Omsk
6120;Khatanga Airport;Khatanga;Russia;HTG;UOHH;71.978058;102.490514;95;8;N;Asia/Krasnoyarsk
6121;Igarka Airport;Igarka;Russia;IAA;UOII;67.4372;86.6219;82;8;N;Asia/Krasnoyarsk
6122;Grozny Airport;Grozny;Russia;GRV;URMG;43.2981;45.7841;548;4;N;Europe/Moscow
6123;Nalchik Airport;Nalchik;Russia;NAL;URMN;43.5129;43.6366;1461;4;N;Europe/Moscow
6124;Beslan Airport;Beslan;Russia;OGZ;URMO;43.2051;44.6066;1673;4;N;Europe/Moscow
6125;Elista Airport;Elista;Russia;ESL;URWI;46.3739;44.3309;501;4;N;Europe/Moscow
6126;Aleknagik Airport;Aleknagik;United States;WKK;5A8;59.2826;-158.618;66;-9;A;America/Anchorage
6127;Brookings Regional Airport;Brookings;United States;BKX;BKX;44.3048;-96.8169;1648;-6;A;America/Chicago
6128;Mercer County Airport;Bluefield;United States;BLF;BLF;37.2958;-81.2077;2857;-5;A;America/New_York
6129;Kearney Municipal Airport;Kearney;United States;EAR;EAR;40.727;-99.0068;2131;-6;A;America/Chicago
6130;Mid Delta Regional Airport;Greenville;United States;GLH;GLH;33.4829;-90.9856;131;-6;A;America/Chicago
6131;Laughlin-Bullhead Intl;Bullhead;United States;IFP;IFP;35.1574;-114.56;695;-7;A;America/Phoenix
6132;Kingman Airport;Kingman;United States;IGM;IGM;35.2595;-113.938;3449;-7;A;America/Phoenix
6133;Tri Cities Airport;Pasco;United States;PSC;PSC;46.2647;-119.119;410;-8;A;America/Los_Angeles
6134;Akutan Seaplane Base;Akutan;United States;KQA;KQA;54.1325;-165.785;0;-9;A;America/Anchorage
6135;Grant County Airport;Silver City;United States;SVC;SVC;32.6365;-108.156;5446;-7;A;America/Denver
6136;Lopez Island Airport;Lopez;United States;LPS;S31;48.4839;-122.938;209;-8;A;America/Los_Angeles
6137;Salekhard Airport;Salekhard;Russia;SLY;USDD;66.590753;66.611042;218;6;N;Asia/Yekaterinburg
6138;Khanty Mansiysk Airport;Khanty-Mansiysk;Russia;HMA;USHH;61.028479;69.086067;76;6;N;Asia/Yekaterinburg
6139;Nyagan Airport;Nyagan;Russia;NYA;USHN;62.11;65.615;361;6;N;Asia/Yekaterinburg
6140;Sovetsky Tyumenskaya Airport;Sovetskiy;Russia;OVS;USHS;61.32;63.6044;351;6;N;Asia/Yekaterinburg
6141;Izhevsk Airport;Izhevsk;Russia;IJK;USII;56.8281;53.4575;531;4;N;Europe/Moscow
6142;Pobedilovo Airport;Kirov;Russia;KVX;USKK;58.5033;49.3483;479;4;N;Europe/Moscow
6143;Nadym Airport;Nadym;Russia;NYM;USMM;65.4809;72.6989;49;6;N;Asia/Yekaterinburg
6144;Raduzhny Airport;Raduzhnyi;Russia;RAT;USNR;62.1586;77.3289;250;6;N;Asia/Yekaterinburg
6145;Nefteyugansk Airport;Nefteyugansk;Russia;NFG;USRN;61.1083;72.65;115;6;N;Asia/Yekaterinburg
6146;Kurgan Airport;Kurgan;Russia;KRO;USUU;55.4753;65.4156;240;6;N;Asia/Yekaterinburg
6147;Khudzhand Airport;Khudzhand;Tajikistan;LBD;UTDL;40.2154;69.6947;1450;5;U;Asia/Dushanbe
6148;Andizhan Airport;Andizhan;Uzbekistan;AZN;UTKA;40.7277;72.294;1515;5;U;Asia/Samarkand
6149;Fergana Airport;Fergana;Uzbekistan;FEG;UTKF;40.3588;71.745;1980;5;U;Asia/Samarkand
6150;Namangan Airport;Namangan;Uzbekistan;NMA;UTKN;40.9846;71.5567;1555;5;U;Asia/Samarkand
6151;Nukus Airport;Nukus;Uzbekistan;NCU;UTNN;42.4884;59.6233;246;5;U;Asia/Samarkand
6152;Urgench Airport;Urgench;Uzbekistan;UGC;UTNU;41.5843;60.6417;320;5;U;Asia/Samarkand
6153;Karshi Khanabad Airport;Khanabad;Uzbekistan;KSQ;UTSL;38.8336;65.9215;1365;5;U;Asia/Samarkand
6154;Termez Airport;Termez;Uzbekistan;TMJ;UTST;37.286667;67.31;1027;5;U;Asia/Samarkand
6155;Staroselye Airport;Rybinsk;Russia;RYB;UUBK;58.1042;38.9294;423;4;N;Europe/Moscow
6156;Belgorod International Airport;Belgorod;Russia;EGO;UUOB;50.6438;36.5901;735;4;N;Europe/Moscow
6157;Kursk East Airport;Kursk;Russia;URS;UUOK;51.7506;36.2956;686;4;N;Europe/Moscow
6158;Lipetsk Airport;Lipetsk;Russia;LPK;UUOL;52.7028;39.5378;584;4;N;Europe/Moscow
6159;Vorkuta Airport;Vorkuta;Russia;VKT;UUYW;67.4886;63.9931;604;4;N;Europe/Moscow
6160;Bugulma Airport;Bugulma;Russia;UUA;UWKB;54.64;52.8017;991;4;N;Europe/Moscow
6161;Yoshkar-Ola Airport;Yoshkar-Ola;Russia;JOK;UWKJ;56.7006;47.9047;348;4;N;Europe/Moscow
6162;Cheboksary Airport;Cheboksary;Russia;CSY;UWKS;56.0903;47.3473;558;4;N;Europe/Moscow
6163;Ulyanovsk East Airport;Ulyanovsk;Russia;ULY;UWLW;54.401;48.8027;252;4;N;Europe/Moscow
6164;Orsk Airport;Orsk;Russia;OSW;UWOR;51.0725;58.5956;909;6;N;Asia/Yekaterinburg
6165;Penza Airport;Penza;Russia;PEZ;UWPP;53.1106;45.0211;614;4;N;Europe/Moscow
6166;Saransk Airport;Saransk;Russia;SKX;UWPS;54.1251;45.2123;676;4;N;Europe/Moscow
6167;Balakovo Airport;Balakovo;Russia;BWO;UWSB;51.8583;47.7456;95;4;N;Europe/Moscow
6168;Hubli Airport;Hubli;India;HBX;VAHB;15.3617;75.0849;2171;5.5;N;Asia/Calcutta
6169;Koggala Airport;Koggala;Sri Lanka;KCT;VCCK;5.99368;80.3203;10;5.5;U;Asia/Colombo
6170;Wirawila Airport;Wirawila;Sri Lanka;WRZ;VCCW;6.254494;81.235189;50;5.5;U;Asia/Colombo
6171;Battambang Airport;Battambang;Cambodia;BBM;VDBG;13.0956;103.224;59;7;U;Asia/Phnom_Penh
6172;Shillong Airport;Shillong;India;SHL;VEBI;25.7036;91.9787;2910;5.5;N;Asia/Calcutta
6173;Lokpriya Gopinath Bordoloi International Airport;Guwahati;India;GAU;VEGT;26.106092;91.585939;162;5.5;N;Asia/Calcutta
6174;Dimapur Airport;Dimapur;India;DMU;VEMR;25.8839;93.7711;487;5.5;N;Asia/Calcutta
6175;Tezpur Airport;Tezpur;India;TEZ;VETZ;26.7091;92.7847;240;5.5;N;Asia/Calcutta
6176;Barisal Airport;Barisal;Bangladesh;BZL;VGBR;22.801;90.3012;23;6;U;Asia/Dhaka
6177;Ban Huoeisay Airport;Huay Xai;Laos;OUI;VLHS;20.2573;100.437;1380;7;N;Asia/Vientiane
6178;Kontum Airport;Kontum;Vietnam;KON;;14.35;108.017;1804;7;U;Asia/Saigon
6179;Bharatpur Airport;Bharatpur;Nepal;BHR;VNBP;27.6781;84.4294;600;5.75;N;Asia/Katmandu
6180;Chandragadhi Airport;Chandragarhi;Nepal;BDP;VNCG;26.570822;88.079578;300;5.75;N;Asia/Katmandu
6181;Meghauli Airport;Meghauli;Nepal;MEY;VNMG;27.583;84.233;600;5.75;N;Asia/Katmandu
6182;Nepalgunj Airport;Nepalgunj;Nepal;KEP;VNNG;28.103633;81.667006;540;5.75;N;Asia/Katmandu
6183;Gan Island Airport;Gan Island;Maldives;GAN;VRMG;-0.693342;73.1556;6;5;U;Indian/Maldives
6184;Hanimaadhoo Airport;Haa Dhaalu Atoll;Maldives;HAQ;VRMH;6.74423;73.1705;4;5;U;Indian/Maldives
6185;Kadhdhoo Airport;Laamu Atoll;Maldives;KDO;VRMK;1.85917;73.5219;4;5;U;Indian/Maldives
6186;Mae Sot Airport;Tak;Thailand;MAQ;VTPM;16.699856;98.545056;690;7;U;Asia/Bangkok
6187;Buon Ma Thuot Airport;Buonmethuot;Vietnam;BMV;VVBM;12.668311;108.120272;1729;7;U;Asia/Saigon
6188;Cat Bi International Airport;Haiphong;Vietnam;HPH;VVCI;20.819386;106.724989;6;7;U;Asia/Saigon
6189;Cam Ranh Airport;Nha Trang;Vietnam;CXR;VVCR;11.998153;109.219372;40;7;U;Asia/Saigon
6190;Co Ong Airport;Conson;Vietnam;VCS;VVCS;8.731831;106.632589;20;7;U;Asia/Saigon
6191;Trà Nóc Airport;Can Tho;Vietnam;VCA;VVCT;10.085119;105.711922;9;7;U;Asia/Saigon
6192;Dien Bien Phu Airport;Dienbienphu;Vietnam;DIN;VVDB;21.397481;103.007831;1611;7;U;Asia/Saigon
6193;Phu Cat Airport;Phucat;Vietnam;UIH;VVPC;13.954986;109.042267;80;7;U;Asia/Saigon
6194;Pleiku Airport;Pleiku;Vietnam;PXU;VVPK;14.004522;108.017158;2434;7;U;Asia/Saigon
6195;Vinh Airport;Vinh;Vietnam;VII;VVVH;18.737569;105.670764;23;7;U;Asia/Saigon
6196;Banmaw Airport;Banmaw;Burma;BMO;VYBM;24.269033;97.246153;370;6.5;U;Asia/Rangoon
6197;Dawei Airport;Dawei;Burma;TVY;VYDW;14.103886;98.203636;84;6.5;U;Asia/Rangoon
6198;Kawthoung Airport;Kawthoung;Burma;KAW;VYKT;10.049258;98.538006;180;6.5;U;Asia/Rangoon
6199;Loikaw Airport;Loikaw;Burma;LIW;VYLK;19.691494;97.214825;2940;6.5;U;Asia/Rangoon
6200;Mawlamyine Airport;Mawlamyine;Burma;MNU;VYMM;16.444747;97.660669;52;6.5;U;Asia/Rangoon
6201;Pathein Airport;Pathein;Burma;BSX;VYPN;16.815233;94.779911;20;6.5;U;Asia/Rangoon
6202;Pakhokku Airport;Pakhokku;Burma;PKK;VYPU;21.3333;95.1;151;6.5;U;Asia/Rangoon
6203;Sumbawa Besar Airport;Sumbawa Island;Indonesia;SWQ;WADS;-8.489039;117.412119;16;8;N;Asia/Makassar
6204;Tambolaka Airport;Waikabubak-Sumba Island;Indonesia;TMC;WADT;-9.409717;119.244494;161;8;N;Asia/Makassar
6205;Bokondini Airport;Bokondini-Papua Island;Indonesia;BUI;WAJB;-3.58365;138.533;9104;9;N;Asia/Jayapura
6206;Senggeh Airport;Senggeh-Papua Island;Indonesia;SEH;WAJS;-3.43333;140.817;914;9;N;Asia/Jayapura
6207;Tanjung Harapan Airport;Tanjung Selor-Borneo Island;Indonesia;TJS;WALG;2.83641;117.374;10;8;N;Asia/Makassar
6208;Datadawai Airport;Datadawai-Borneo Island;Indonesia;DTD;WALJ;0.717;116.483;229;8;N;Asia/Makassar
6209;Barau(Kalimaru) Airport;Tanjung Redep-Borneo Island;Indonesia;BEJ;WALK;2.155497;117.432256;59;8;N;Asia/Makassar
6210;Warukin Airport;Tanjung-Borneo Island;Indonesia;TJG;WAON;-2.21656;115.436;197;8;N;Asia/Makassar
6211;Sampit(Hasan) Airport;Sampit-Borneo Island;Indonesia;SMQ;WAOS;-2.499194;112.974992;50;7;N;Asia/Jakarta
6212;Dumatubun Airport;Langgur-Kei Islands;Indonesia;LUV;WAPL;-5.661619;132.731431;10;9;N;Asia/Jayapura
6213;Mali Airport;Alor Island;Indonesia;ARD;WATM;-8.13234;124.597;10;8;N;Asia/Makassar
6214;Belaga Airport;Belaga;Malaysia;BLG;WBGC;2.65;113.767;200;8;N;Asia/Kuala_Lumpur
6215;Long Lellang Airport;Long Datih;Malaysia;LGL;WBGF;3.421;115.154;1400;8;N;Asia/Kuala_Lumpur
6216;Long Seridan Airport;Long Seridan;Malaysia;ODN;WBGI;3.967;115.05;607;8;N;Asia/Kuala_Lumpur
6217;Mukah Airport;Mukah;Malaysia;MKM;WBGK;2.90639;112.08;13;8;N;Asia/Kuala_Lumpur
6218;Bakalalan Airport;Bakalalan;Malaysia;BKM;WBGQ;3.974;115.618;2900;8;N;Asia/Kuala_Lumpur
6219;Lawas Airport;Lawas;Malaysia;LWY;WBGW;4.84917;115.408;5;8;N;Asia/Kuala_Lumpur
6220;Bario Airport;Bario;Malaysia;BBN;WBGZ;3.73389;115.479;3350;8;N;Asia/Kuala_Lumpur
6221;Tomanggong Airport;Tomanggong;Malaysia;TMG;WBKM;5.4;118.65;26;8;N;Asia/Kuala_Lumpur
6222;Kudat Airport;Kudat;Malaysia;KUD;WBKT;6.9225;116.836;10;8;N;Asia/Kuala_Lumpur
6223;Radin Inten II (Branti) Airport;Bandar Lampung-Sumatra Island;Indonesia;TKG;WICT;-5.242339;105.178939;282;7;N;Asia/Jakarta
6224;Halim Perdanakusuma International Airport;Jakarta;Indonesia;HLP;WIHH;-6.26661;106.891;84;7;N;Asia/Jakarta
6225;Ranai Airport;Ranai-Natuna Besar Island;Indonesia;NTX;WION;3.908714;108.387897;7;7;N;Asia/Jakarta
6226;Pangsuma Airport;Putussibau-Borneo Island;Indonesia;PSU;WIOP;0.835578;112.937144;297;7;N;Asia/Jakarta
6227;Susilo Airport;Sintang-Borneo Island;Indonesia;SQG;WIOS;0.063619;111.473428;98;7;N;Asia/Jakarta
6228;Pendopo Airport;Talang Gudang-Sumatra Island;Indonesia;PDO;WIPQ;-3.286069;103.8796;184;7;N;Asia/Jakarta
6229;Malikus Saleh Airport;Lhok Seumawe-Sumatra Island;Indonesia;LSW;WITM;5.226681;96.950342;90;7;N;Asia/Jakarta
6230;Pulau Pangkor Airport;Pangkor Island;Malaysia;PKG;WMPA;4.24472;100.553;19;8;N;Asia/Kuala_Lumpur
6231;Stagen Airport;Laut Island;Indonesia;KBU;WRBK;-3.29472;116.165;4;8;N;Asia/Makassar
6232;Long Bawan Airport;Long Bawan-Borneo Island;Indonesia;LBW;WRLB;3.867;115.683;3590;8;N;Asia/Makassar
6233;Nunukan Airport;Nunukan-Nunukan Island;Indonesia;NNX;WRLF;4.13653;117.667;52;8;N;Asia/Makassar
6234;Long Apung Airport;Long Apung-Borneo Island;Indonesia;LPU;WRLP;0.583;115.6;627;8;N;Asia/Makassar
6235;Albany Airport;Albany;Australia;ALH;YABA;-34.9433;117.809;233;8;O;Australia/Perth
6236;Argyle Airport;Argyle;Australia;GYL;YARG;-16.6369;128.451;522;8;O;Australia/Perth
6237;Aurukun Airport;Aurukun;Australia;AUU;YAUR;-13.3539;141.721;31;10;O;Australia/Brisbane
6238;Barcaldine Airport;Barcaldine;Australia;BCI;YBAR;-23.5653;145.307;878;10;O;Australia/Brisbane
6239;Badu Island Airport;Badu Island;Australia;BDD;YBAU;-10.15;141.175;14;9;O;\N
6240;Birdsville Airport;Birdsville;Australia;BVI;YBDV;-25.8975;139.348;159;10;O;Australia/Brisbane
6241;Broken Hill Airport;Broken Hill;Australia;BHQ;YBHI;-32.0014;141.472;958;9.5;O;Australia/Adelaide
6242;Hamilton Island Airport;Hamilton Island;Australia;HTI;YBHM;-20.3581;148.952;15;10;O;Australia/Brisbane
6243;Bedourie Airport;Bedourie;Australia;BEU;YBIE;-24.3461;139.46;300;10;O;Australia/Brisbane
6244;Bourke Airport;Bourke;Australia;BRK;YBKE;-30.0392;145.952;352;10;O;Australia/Sydney
6245;Burketown Airport;Burketown;Australia;BUC;YBKT;-17.7486;139.534;21;10;O;Australia/Brisbane
6246;Boigu Airport;Boigu;Australia;GIC;YBOI;-9.23278;142.218;23;10;O;\N
6247;Oakey Airport;Oakey;Australia;OKY;YBOK;-27.411389;151.735278;1335;10;O;Australia/Brisbane
6248;Boulia Airport;Boulia;Australia;BQL;YBOU;-22.9133;139.9;542;10;O;Australia/Brisbane
6249;Bathurst Airport;Bathurst;Australia;BHS;YBTH;-33.4094;149.652;2435;10;O;Australia/Sydney
6250;Blackwater Airport;Blackwater;Australia;BLT;YBTR;-23.6031;148.807;657;10;O;Australia/Brisbane
6251;Carnarvon Airport;Carnarvon;Australia;CVQ;YCAR;-24.8806;113.672;13;8;O;Australia/Perth
6252;Cobar Airport;Cobar;Australia;CAZ;YCBA;-31.5383;145.794;724;10;O;Australia/Sydney
6253;Coober Pedy Airport;Coober Pedy;Australia;CPD;YCBP;-29.04;134.721;740;9.5;O;Australia/Adelaide
6254;Coconut Island Airport;Coconut Island;Australia;CNC;YCCT;-10.05;143.07;3;10;O;\N
6255;Cloncurry Airport;Cloncurry;Australia;CNJ;YCCY;-20.6686;140.504;616;10;O;Australia/Brisbane
6256;Ceduna Airport;Ceduna;Australia;CED;YCDU;-32.1306;133.71;77;9.5;O;Australia/Adelaide
6257;Cooktown Airport;Cooktown;Australia;CTN;YCKN;-15.4447;145.184;26;10;O;Australia/Brisbane
6258;Cunnamulla Airport;Cunnamulla;Australia;CMA;YCMU;-28.03;145.622;630;10;O;Australia/Brisbane
6259;Coonamble Airport;Coonamble;Australia;CNB;YCNM;-30.9833;148.376;604;10;O;Australia/Sydney
6260;Coen Airport;Coen;Australia;CUQ;YCOE;-13.7608;143.114;532;10;O;Australia/Brisbane
6261;Cooma Snowy Mountains Airport;Cooma;Australia;OOM;YCOM;-36.3006;148.974;3088;10;O;Australia/Sydney
6262;Doomadgee Airport;Doomadgee;Australia;DMD;YDMG;-17.9403;138.822;153;10;O;Australia/Brisbane
6263;Darnley Island Airport;Darnley Island;Australia;NLF;YDNI;-9.58333;143.767;0;10;O;\N
6264;Devonport Airport;Devonport;Australia;DPO;YDPO;-41.1697;146.43;33;10;O;Australia/Melbourne
6265;Elcho Island Airport;Elcho Island;Australia;ELC;YELD;-12.0194;135.571;101;9.5;O;Australia/Darwin
6266;Esperance Airport;Esperance;Australia;EPR;YESP;-33.6844;121.823;470;8;O;Australia/Perth
6267;Flinders Island Airport;Flinders Island;Australia;FLS;YFLI;-40.0917;147.993;10;10;O;Australia/Melbourne
6268;Geraldton Airport;Geraldton;Australia;GET;YGEL;-28.7961;114.707;121;8;O;Australia/Perth
6269;Gladstone Airport;Gladstone;Australia;GLT;YGLA;-23.8697;151.223;64;10;O;Australia/Brisbane
6270;Groote Eylandt Airport;Groote Eylandt;Australia;GTE;YGTE;-13.975;136.46;53;9.5;O;Australia/Darwin
6271;Griffith Airport;Griffith;Australia;GFF;YGTH;-34.2508;146.067;439;10;O;Australia/Sydney
6272;Horn Island Airport;Horn Island;Australia;HID;YHID;-10.5864;142.29;43;10;O;Australia/Brisbane
6273;Hooker Creek Airport;Hooker Creek;Australia;HOK;YHOO;-18.3367;130.638;320;9.5;O;Australia/Darwin
6274;Mount Hotham Airport;Mount Hotham;Australia;MHU;YHOT;-37.0475;147.334;4260;10;O;Australia/Hobart
6275;Hughenden Airport;Hughenden;Australia;HGD;YHUG;-20.815;144.225;1043;10;O;Australia/Brisbane
6276;Julia Creek Airport;Julia Creek;Australia;JCK;YJLC;-20.6683;141.723;404;10;O;Australia/Brisbane
6277;Kalbarri Airport;Kalbarri;Australia;KAX;YKBR;-27.69;114.262;157;8;O;Australia/Perth
6278;King Island Airport;King Island;Australia;KNS;YKII;-39.8775;143.878;132;10;O;Australia/Melbourne
6279;Kalkgurung Airport;Kalkgurung;Australia;KFG;YKKG;-17.4319;130.808;646;9.5;O;Australia/Darwin
6280;Karumba Airport;Karumba;Australia;KRB;YKMB;-17.4567;140.83;5;10;O;Australia/Brisbane
6281;Kowanyama Airport;Kowanyama;Australia;KWM;YKOW;-15.4856;141.751;35;10;O;Australia/Brisbane
6282;Kubin Airport;Kubin;Australia;KUG;YKUB;-10.225;142.218;15;10;O;Australia/Brisbane
6283;Leonora Airport;Leonora;Australia;LNO;YLEO;-28.8781;121.315;1217;8;O;Australia/Perth
6284;Lake Evella Airport;Lake Evella;Australia;LEL;YLEV;-12.4989;135.806;256;9.5;O;Australia/Darwin
6285;Lord Howe Island Airport;Lord Howe Island;Australia;LDH;YLHI;-31.5383;159.077;5;10.5;O;\N
6286;Lockhart River Airport;Lockhart River;Australia;IRG;YLHR;-12.7869;143.305;77;10;O;Australia/Brisbane
6287;Lismore Airport;Lismore;Australia;LSY;YLIS;-28.8303;153.26;35;10;O;Australia/Sydney
6288;Lightning Ridge Airport;Lightning Ridge;Australia;LHG;YLRD;-29.4567;147.984;540;10;O;Australia/Sydney
6289;Longreach Airport;Longreach;Australia;LRE;YLRE;-23.4342;144.28;627;10;O;Australia/Brisbane
6290;Leinster Airport;Leinster;Australia;LER;YLST;-27.8433;120.703;1631;8;O;Australia/Perth
6291;Laverton Airport;Laverton;Australia;LVO;YLTN;-28.6136;122.424;1530;8;O;Australia/Perth
6292;Mabuiag Island Airport;Mabuiag Island;Australia;UBB;YMAA;-9.95;142.183;0;10;O;Australia/Brisbane
6293;Meekatharra Airport;Meekatharra;Australia;MKR;YMEK;-26.6117;118.548;1713;8;O;Australia/Perth
6294;Merimbula Airport;Merimbula;Australia;MIM;YMER;-36.9086;149.901;7;10;O;Australia/Sydney
6295;Milingimbi Airport;Milingimbi;Australia;MGT;YMGB;-12.0944;134.894;53;9.5;O;Australia/Darwin
6296;Maningrida Airport;Maningrida;Australia;MNG;YMGD;-12.0561;134.234;123;9.5;O;Australia/Darwin
6297;McArthur River Mine Airport;McArthur River Mine;Australia;MCV;YMHU;-16.4425;136.084;131;9.5;O;Australia/Darwin
6298;Mildura Airport;Mildura;Australia;MQL;YMIA;-34.2292;142.086;167;10;O;Australia/Hobart
6299;Mount Magnet Airport;Mount Magnet;Australia;MMG;YMOG;-28.1161;117.842;1354;8;O;Australia/Perth
6300;Moree Airport;Moree;Australia;MRZ;YMOR;-29.4989;149.845;701;10;O;Australia/Sydney
6301;Moranbah Airport;Moranbah;Australia;MOV;YMRB;-22.0578;148.077;770;10;O;Australia/Brisbane
6302;Moruya Airport;Moruya;Australia;MYA;YMRY;-35.8978;150.144;14;10;O;Australia/Sydney
6303;Mount Gambier Airport;Mount Gambier;Australia;MGB;YMTG;-37.7456;140.785;212;9.5;O;Australia/Adelaide
6304;Mornington Island Airport;Mornington Island;Australia;ONG;YMTI;-16.6625;139.178;33;10;O;Australia/Brisbane
6305;Murray Island Airport;Murray Island;Australia;MYI;YMUI;-9.91667;144.055;0;10;O;\N
6306;Maryborough Airport;Maryborough;Australia;MBH;YMYB;-25.5133;152.715;38;10;O;Australia/Brisbane
6307;Narrandera Airport;Narrandera;Australia;NRA;YNAR;-34.7022;146.512;474;10;O;Australia/Sydney
6308;Narrabri Airport;Narrabri;Australia;NAA;YNBR;-30.3192;149.827;788;10;O;Australia/Sydney
6309;Normanton Airport;Normanton;Australia;NTN;YNTN;-17.6836;141.07;73;10;O;Australia/Brisbane
6310;Newman Airport;Newman;Australia;ZNE;YNWN;-23.4178;119.803;1724;8;O;Australia/Perth
6311;Olympic Dam Airport;Olympic Dam;Australia;OLP;YOLD;-30.485;136.877;343;9.5;O;Australia/Adelaide
6312;Port Augusta Airport;Argyle;Australia;PUG;YPAG;-32.506944;137.716667;56;9.5;O;Australia/Adelaide
6313;Palm Island Airport;Palm Island;Australia;PMK;YPAM;-18.7553;146.581;28;10;O;Australia/Brisbane
6314;Paraburdoo Airport;Paraburdoo;Australia;PBO;YPBO;-23.1711;117.745;1406;8;O;Australia/Perth
6315;Cocos Keeling Island Airport;Cocos Keeling Island;Cocos (Keeling) Islands;CCK;YPCC;-12.1883;96.8339;10;6.5;U;Indian/Cocos
6316;Gove Airport;Gove;Australia;GOV;YPGV;-12.2694;136.818;192;9.5;O;Australia/Darwin
6317;Parkes Airport;Parkes;Australia;PKE;YPKS;-33.1314;148.239;1069;10;O;Australia/Sydney
6318;Port Lincoln Airport;Port Lincoln;Australia;PLO;YPLC;-34.6053;135.88;36;9.5;O;Australia/Adelaide
6319;Pormpuraaw Airport;Pormpuraaw;Australia;EDR;YPMP;-14.8967;141.609;10;10;O;Australia/Brisbane
6320;Port Macquarie Airport;Port Macquarie;Australia;PQQ;YPMQ;-31.4358;152.863;12;10;O;Australia/Sydney
6321;Portland Airport;Portland;Australia;PTJ;YPOD;-38.3181;141.471;265;10;O;Australia/Hobart
6322;Quilpie Airport;Quilpie;Australia;ULP;YQLP;-26.6122;144.253;655;10;O;Australia/Brisbane
6323;Ramingining Airport;Ramingining;Australia;RAM;YRNG;-12.3564;134.898;206;9.5;O;Australia/Darwin
6324;Roma Airport;Roma;Australia;RMA;YROM;-26.545;148.775;1032;10;O;Australia/Brisbane
6325;St George Airport;St George;Australia;SGO;YSGE;-28.0497;148.595;656;10;O;Australia/Brisbane
6326;Shark Bay Airport;Shark Bay;Australia;MJK;YSHK;-25.8939;113.577;111;8;O;Australia/Perth
6327;Saibai Island Airport;Saibai Island;Australia;SBR;YSII;-9.37833;142.625;0;10;O;\N
6328;Strahan Airport;Strahan;Australia;SRN;YSRN;-42.155;145.292;20;10;O;Australia/Melbourne
6329;Thargomindah Airport;Thargomindah;Australia;XTG;YTGM;-27.9864;143.811;433;10;O;Australia/Brisbane
6330;Tennant Creek Airport;Tennant Creek;Australia;TCA;YTNK;-19.6344;134.183;1236;9.5;O;Australia/Darwin
6331;Victoria River Downs Airport;Victoria River Downs;Australia;VCD;YVRD;-16.4033;131.002;89;9.5;O;Australia/Darwin
6332;Warraber Island Airport;Sue Islet;Australia;SYU;YWBS;-10.2083;142.825;3;10;O;Australia/Brisbane
6333;Windorah Airport;Windorah;Australia;WNR;YWDH;-25.4131;142.667;452;10;O;Australia/Brisbane
6334;Whyalla Airport;Whyalla;Australia;WYA;YWHA;-33.0589;137.514;41;9.5;O;Australia/Adelaide
6335;Wiluna Airport;Wiluna;Australia;WUN;YWLU;-26.6292;120.221;1649;8;O;Australia/Perth
6336;Wollongong Airport;Wollongong;Australia;WOL;YWOL;-34.5611;150.789;31;10;O;Australia/Sydney
6337;Winton Airport;Winton;Australia;WIN;YWTN;-22.3636;143.086;638;10;O;Australia/Brisbane
6338;Wynyard Airport;Burnie;Australia;BWT;YWYY;-40.9989;145.731;62;10;O;Australia/Melbourne
6339;Yorke Island Airport;Yorke Island;Australia;OKR;YYKI;-9.75703;143.411;0;10;O;\N
6340;Yam Island Airport;Yam Island;Australia;XMY;YYMI;-9.90111;142.776;0;10;O;\N
6341;Beijing Nanyuan Airport;Beijing;China;NAY;ZBBB;39.7825;116.387778;0;8;U;Asia/Chongqing
6342;Chifeng Airport;Chifeng;China;CIF;ZBCF;42.235;118.908;0;8;U;Asia/Chongqing
6343;Changzhi Airport;Changzhi;China;CIH;ZBCZ;36.2475;113.126;0;8;U;Asia/Chongqing
6344;Datong Airport;Datong;China;DAT;ZBDT;40.0603;113.482;3442;8;U;Asia/Chongqing
6345;Baita Airport;Hohhot;China;HET;ZBHH;40.851422;111.824103;3556;8;U;Asia/Chongqing
6346;Baotou Airport;Baotou;China;BAV;ZBOW;40.56;109.997;3321;8;U;Asia/Chongqing
6347;Shijiazhuang Daguocun International Airport;Shijiazhuang;China;SJW;ZBSJ;38.280686;114.6973;233;8;U;Asia/Chongqing
6348;Tongliao Airport;Tongliao;China;TGO;ZBTL;43.5567;122.2;0;8;U;Asia/Chongqing
6349;Ulanhot Airport;Ulanhot;China;HLH;ZBUL;46.083;122.017;0;8;U;Asia/Chongqing
6350;Xilinhot Airport;Xilinhot;China;XIL;ZBXH;43.9156;115.964;0;8;U;Asia/Chongqing
6351;Beihai Airport;Beihai;China;BHY;ZGBH;21.5394;109.294;0;8;U;Asia/Chongqing
6352;Changde Airport;Changde;China;CGD;ZGCD;28.9189;111.64;0;8;U;Asia/Chongqing
6353;Dayong Airport;Dayong;China;DYG;ZGDY;29.1028;110.443;692;8;U;Asia/Chongqing
6354;Meixian Airport;Meixian;China;MXZ;ZGMX;24.35;116.133;0;8;U;Asia/Chongqing
6355;Zhuhai Airport;Zhuhai;China;ZUH;ZGSD;22.0064;113.376;0;8;U;Asia/Chongqing
6356;Bailian Airport;Liuzhou;China;LZH;ZGZH;24.2075;109.391;295;8;U;Asia/Chongqing
6357;Zhanjiang Airport;Zhanjiang;China;ZHA;ZGZJ;21.2144;110.358;0;8;U;Asia/Chongqing
6358;Enshi Airport;Enshi;China;ENH;ZHES;30.3203;109.485;0;8;U;Asia/Chongqing
6359;Nanyang Airport;Nanyang;China;NNY;ZHNY;32.9808;112.615;0;8;U;Asia/Chongqing
6360;Xiangfan Airport;Xiangfan;China;XFN;ZHXF;32.1506;112.291;0;8;U;Asia/Chongqing
6361;Yichang Airport;Yichang;China;YIH;ZHYC;30.671;111.441;0;8;U;Asia/Chongqing
6362;Ankang Airport;Ankang;China;AKA;ZLAK;32.7081;108.931;0;8;U;Asia/Chongqing
6363;Golmud Airport;Golmud;China;GOQ;ZLGM;34.633;98.867;0;8;U;Asia/Chongqing
6364;Hanzhong Airport;Hanzhong;China;HZG;ZLHZ;33.0636;107.008;0;8;U;Asia/Chongqing
6365;Qingyang Airport;Qingyang;China;IQN;ZLQY;35.7997;107.603;0;8;U;Asia/Chongqing
6366;Xining Caojiabu Airport;Xining;China;XNN;ZLXN;36.5275;102.043;0;8;U;Asia/Chongqing
6367;Yan'an Airport;Yan'an;China;ENY;ZLYA;36.6369;109.554;0;8;U;Asia/Chongqing
6368;Yulin Airport;Yulin;China;UYN;ZLYL;38.2692;109.731;0;8;U;Asia/Chongqing
6369;Arvaikheer Airport;Arvaikheer;Mongolia;AVK;ZMAH;46.2503;102.802;5932;8;U;Asia/Ulaanbaatar
6370;Altai Airport;Altai;Mongolia;LTI;ZMAT;46.3764;96.2211;7260;8;U;Asia/Ulaanbaatar
6371;Bayankhongor Airport;Bayankhongor;Mongolia;BYN;ZMBH;46.1633;100.704;6085;8;U;Asia/Ulaanbaatar
6372;Dalanzadgad Airport;Dalanzadgad;Mongolia;DLZ;ZMDZ;43.5917;104.43;4787;8;U;Asia/Ulaanbaatar
6373;Khovd Airport;Khovd;Mongolia;HVD;ZMKD;47.9541;91.6282;4898;7;U;Asia/Hovd
6374;Muren Airport;Muren;Mongolia;MXV;ZMMN;49.6633;100.099;4272;8;U;Asia/Ulaanbaatar
6375;Diqing Airport;Shangri-La;China;DIG;ZPDQ;27.7936;99.6772;0;8;U;Asia/Chongqing
6376;Mangshi Airport;Luxi;China;LUM;ZPLX;24.4011;98.5317;2890;8;U;Asia/Chongqing
6377;Simao Airport;Simao;China;SYM;ZPSM;22.7933;100.959;0;8;U;Asia/Chongqing
6378;Zhaotong Airport;Zhaotong;China;ZAT;ZPZT;27.3256;103.755;0;8;U;Asia/Chongqing
6379;Ganzhou Airport;Ganzhou;China;KOW;ZSGZ;25.8258;114.912;0;8;U;Asia/Chongqing
6380;Jingdezhen Airport;Jingdezhen;China;JDZ;ZSJD;29.3386;117.176;112;8;U;Asia/Chongqing
6381;Jiujiang Lushan Airport;Jiujiang;China;JIU;ZSJJ;29.733;115.983;0;8;U;Asia/Chongqing
6382;Quzhou Airport;Quzhou;China;JUZ;ZSJU;28.9658;118.899;0;8;U;Asia/Chongqing
6383;Lianyungang Airport;Lianyungang;China;LYG;ZSLG;34.55;119.25;0;8;U;Asia/Chongqing
6384;Huangyan Luqiao Airport;Huangyan;China;HYN;ZSLQ;28.5622;121.429;0;8;U;Asia/Chongqing
6385;Shubuling Airport;Linyi;China;LYI;ZSLY;35.0461;118.412;0;8;U;Asia/Chongqing
6386;Quanzhou Airport;Quanzhou;China;JJN;ZSQZ;24.7964;118.59;0;8;U;Asia/Chongqing
6387;Tunxi International Airport;Huangshan;China;TXN;ZSTX;29.7333;118.256;0;8;U;Asia/Chongqing
6388;Weifang Airport;Weifang;China;WEF;ZSWF;36.6467;119.119;0;8;U;Asia/Chongqing
6389;Weihai Airport;Weihai;China;WEH;ZSWH;37.1871;122.229;145;8;U;Asia/Chongqing
6390;Wuxi Airport;Wuxi;China;WUX;ZSWX;31.4944;120.429;0;8;U;Asia/Chongqing
6391;Nanping Wuyishan Airport;Wuyishan;China;WUS;ZSWY;27.7019;118.001;614;8;U;Asia/Chongqing
6392;Wenzhou Yongqiang Airport;Wenzhou;China;WNZ;ZSWZ;27.9122;120.852;0;8;U;Asia/Chongqing
6393;Yancheng Airport;Yancheng;China;YNZ;ZSYN;33.3856;120.125;0;8;U;Asia/Chongqing
6394;Yiwu Airport;Yiwu;China;YIW;ZSYW;29.3447;120.032;0;8;U;Asia/Chongqing
6395;Zhoushan Airport;Zhoushan;China;HSN;ZSZS;29.9342;122.362;0;8;U;Asia/Chongqing
6396;Qamdo Bangda Airport;Bangda;China;BPX;ZUBD;30.5536;97.1083;14219;8;U;Asia/Chongqing
6397;Dachuan Airport;Dazhou;China;DAX;ZUDX;31.3;107.5;0;8;U;Asia/Chongqing
6398;Guangyuan Airport;Guangyuan;China;GYS;ZUGU;32.3911;105.702;0;8;U;Asia/Chongqing
6399;Luzhou Airport;Luzhou;China;LZO;ZULZ;28.8522;105.393;0;8;U;Asia/Chongqing
6400;Mianyang Airport;Mianyang;China;MIG;ZUMY;31.4281;104.741;0;8;U;Asia/Chongqing
6401;Nanchong Airport;Nanchong;China;NAO;ZUNC;30.754;106.062;0;8;U;Asia/Chongqing
6402;Nyingchi Airport;Nyingchi;China;LZY;ZUNZ;29.3033;94.3353;9675;8;U;Asia/Chongqing
6403;Wanxian Airport;Wanxian;China;WXN;ZUWX;30.8361;108.406;0;8;U;Asia/Chongqing
6404;Aksu Airport;Aksu;China;AKU;ZWAK;41.2625;80.2917;0;8;U;Asia/Chongqing
6405;Qiemo Airport;Qiemo;China;IQM;ZWCM;38.1494;85.5328;4108;8;U;Asia/Chongqing
6406;Kuqa Airport;Kuqa;China;KCA;ZWKC;41.7181;82.9869;3524;8;U;Asia/Chongqing
6407;Korla Airport;Korla;China;KRL;ZWKL;41.6978;86.1289;0;8;U;Asia/Chongqing
6408;Karamay Airport;Karamay;China;KRY;ZWKM;45.617;84.883;0;8;U;Asia/Chongqing
6409;Yining Airport;Yining;China;YIN;ZWYN;43.9558;81.3303;0;8;U;Asia/Chongqing
6410;Heihe Airport;Heihe;China;HEK;ZYHE;50.25;127.3;8530;8;U;Asia/Chongqing
6411;Jiamusi Airport;Jiamusi;China;JMU;ZYJM;46.843394;130.465389;262;8;U;Asia/Chongqing
6412;Jinzhou Airport;Jinzhou;China;JNZ;ZYJZ;41.1014;121.062;0;8;U;Asia/Chongqing
6413;Qiqihar Sanjiazi Airport;Qiqihar;China;NDG;ZYQQ;47.239628;123.918131;477;8;U;Asia/Chongqing
6414;Yanji Airport;Yanji;China;YNJ;ZYYJ;42.8828;129.451258;624;8;U;Asia/Chongqing
6417;Valletta Sea Plane Terminal;Valletta;Malta;;\N;35.888863;14.508417;0;1;E;Europe/Malta
6418;Gozo Sea Plane Terminal;Mgarr;Malta;;\N;36.026389;14.298333;0;1;E;Europe/Malta
6419;Mount Keith;Mount Keith;Australia;WME;YMNE;-27.286389;120.554722;1792;8;O;Australia/Perth
6420;Gran Roque Airport;Los Roques;Venezuela;LRV;SVRS;11.95;-66.67;17;-4.5;S;America/Caracas
6421;Inishmore Airport;Inis Mor;Ireland;IOR;EIIM;53.1067;-9.65361;24;0;U;Europe/Dublin
6422;Connemara Regional Airport;Indreabhan;Ireland;NNR;EICA;53.2303;-9.46778;0;0;U;Europe/Dublin
6423;Guettin MecklenburgVorpommern Germany;Ruegen;Germany;GTI;EDCG;54.383333;13.325278;69;1;U;Europe/Berlin
6424;Berezovo;Berezovo;Russia;NBB;USHB;63.9241;65.0487;98;6;N;Asia/Yekaterinburg
6425;Szczecin-Dabie;Szczecin;Poland;;EPSD;53.471138;14.63775;3;1;E;Europe/Warsaw
6426;Worcester Regional Airport;Worcester;United States;ORH;KORH;42.2673;-71.8757;1009;-5;A;America/New_York
6427;Anqing Airport;Anqing;China;AQG;ZSAQ;30.5822;117.0502;0;8;N;Asia/Chongqing
6428;Jing Gang Shan Airport;Ji An;China;JGS;\N;26.8997;114.7375;0;8;N;Asia/Chongqing
6429;Shanhaiguan Airport;Qinhuangdao;China;SHP;ZBSH;39.9681;119.731;0;8;N;Asia/Chongqing
6430;Zhangxiao;Yuncheng;China;YCU;ZBYC;35.018;110.993;0;8;N;Asia/Chongqing
6431;Lanzhou Airport;Lanzhou;China;LHW;ZLAN;36.117;103.617;0;8;N;Asia/Chongqing
6432;Jiayuguan Airport;Jiayuguan;China;JGN;ZLJQ;39.8569;98.3414;0;8;N;Asia/Chongqing
6433;Dandong;Dandong;China;DDG;ZYDD;40.0255;124.2866;0;8;N;Asia/Chongqing
6434;Ordos Ejin Horo;Dongsheng;China;DSN;ZBDS;39.85;110.033;0;8;N;Asia/Chongqing
6435;Panzhihua;Panzhihua;China;PZI;ZUZH;26.54;101.799;9186;8;N;Asia/Chongqing
6436;Grytviken;Grytviken;South Georgia and the Islands;;\N;-54.16534;-36.30288;0;-2;N;Atlantic/South_Georgia
6437;South Shetland;South Shetland;Antarctica;;\N;-68;-58;0;-4;N;\N
6438;South Shetland;South Shetland;Antarctica;;\N;-63.37;-62.83;0;12;N;Antarctica/South_Pole
6439;New Rochelle Amtrak Station;New Rochelle;United States;;\N;40.913923;-73.782008;66;-5;A;America/New_York
8838;Anoka County Blaine Airport;Anoka;United States;;KANE;45.1448889;-93.2101944;912;-6;U;America/Chicago
6441;New Haven Rail Station;New Haven;United States;ZVE;\N;41.298669;-72.925992;7;-5;A;America/New_York
6442;Chicago Union Station;Chicago;United States;;\N;41.8791966;-87.6388507;581;-6;A;America/Chicago
6443;Dibrugarh Airport;Dibrugarh;India;DIB;\N;27.4839;95.0169;362;5.5;N;Asia/Calcutta
6444;Doha Free Zone Airport;Doha;Qatar;XOZ;\N;24.8333;50.9166;0;3;N;Asia/Qatar
6445;Bremerton National;Bremerton;United States;PWT;KPWT;47.490244;-122.764814;444;-8;A;America/Los_Angeles
6446;Spencer Muni;Spencer;United States;SPW;KSPW;43.165527;-95.202805;1339;-6;A;America/Chicago
6447;Jefferson City Memorial Airport;Jefferson City;United States;JEF;KJEF;38.5912;-92.1561;549;-6;A;America/Chicago
6448;Grand Canyon West Airport;Grand Canyon West;United States;GCW;\N;35.5925;-113.4859;4825;-7;U;America/Phoenix
6449;Boulder City Municipal Airport;Boulder City;United States;BLD;KBVU;35.5651;-114.514;2201;-7;U;America/Phoenix
6450;Tannheim;Tannheim;Germany;;EDMT;48.011736;10.100903;1902;1;E;Europe/Berlin
6451;Glenview Amtrak Station;Glenview;United States;;\N;42.074197;-87.805346;630;-6;A;America/Chicago
6452;Baltimore Penn Station;Baltimore;United States;;\N;39.307408;-76.615552;62;-5;A;America/New_York
6453;Summit Camp;Ice Cap;Greenland;;\N;72.579187;-38.4572;11000;0;U;America/Danmarkshavn
6454;Unst Airport;Unst;United Kingdom;UNT;EGPW;60.7472;-0.85385;0;0;E;Europe/London
6455;Pagerungan;Pagerungan;Indonesia;;WA19;-6.956608;115.931239;20;7;N;Asia/Jakarta
6456;Provincetown Muni;Provincetown;United States;PVC;KPVC;42.071945;-70.22139;9;-5;A;America/New_York
6457;Kenmore Air Harbor Seaplane Base;Seattle;United States;LKE;KW55;47.629;-122.339;14;-8;A;America/Los_Angeles
6458;Seria - Anduki;Seria;Brunei;;WBAK;4.6333;114.3833;5;8;N;Asia/Brunei
6459;Magas;Nazran;Russia;%u0;%u04;43.323815;45.017761;1060;4;N;Europe/Moscow
6460;Saint Barthelemy;Gustavia;France;SBH;TFFJ;17.9023;-62.8324;50;-4;E;\N
6461;Morro de Sao Paulo;Morro de Sao Paulo;Brazil;;\N;-13.2314;-38.5432;10;3;S;\N
6462;Morro de Sao Paulo;Morro de Sao Paulo;Brazil;;\N;-13.386571;-38.909122;10;-3;S;America/Fortaleza
6463;Belize City Municipal Airport;Belize;Belize;TZA;\N;17.5344;-88.298;15;-6;N;America/Belize
6464;Kostroma - Sokerkino;Kostroma;Russia;;UUBA;57.47;41.01;453;4;N;Europe/Moscow
6465;Sukhumi Dranda;Sukhumi;Georgia;SUI;UGSS;42.87;41.12;0;4;N;Asia/Tbilisi
6466;Tambow;Tambow;Russia;TBW;UUOT;52.81;41.48;126;4;N;Europe/Moscow
6467;Oban Airport;North Connel;United Kingdom;OBN;EGEO;56.464;-5.4;23;0;E;Europe/London
6468;Vilamendhoo;Vilamendhoo;Maldives;;\N;73.0191098670891;3.60965475494316;0;1;U;Arctic/Longyearbyen
6469;Vilamendhoo;Vilamendhoo;Maldives;;\N;3.60965475494316;73.0191098670891;0;5;U;Indian/Maldives
6470;Vilamendhoo;Vilamendhoo;Maldives;;\N;3.484;73.802;0;5;U;Indian/Maldives
6471;Vilamendhoo;Vilamendhoo;Maldives;;\N;73.801;3.801;0;5;U;\N
6472;Aaa;Aaa;Maldives;;\N;73.9;3.222;0;5;U;\N
6473;Vilamendhoo;Vilamendhoo;Maldives;;\N;0;0;0;5;U;\N
6474;Vilamendhoo;Vilamendhoo;Maldives;;\N;3.8;73.8;0;5;U;Indian/Maldives
6475;Sharya;Sharya;Russia;;\N;58.2321;45.3314;453;4;N;Europe/Moscow
6476;Mt. Fuji Shizuoka Airport;Shizuoka;Japan;FSZ;RJNS;34.796111;138.189444;433;9;N;Asia/Tokyo
6477;Erechim Airport;Erechim;Brazil;ERM;SSER;-27.6619;-52.2683;2498;-3;S;America/Sao_Paulo
6478;La Cote;Prangins;Switzerland;;LSGP;46.413056;6.258611;1350;1;E;Europe/Zurich
6479;Courchevel Airport;Courcheval;France;CVF;LFLJ;45.3967;6.63472;6588;1;E;Europe/Paris
6481;Fullerton Municipal Airport;Fullerton;United States;FUL;KFUL;33.521925;-117.584722;96;-8;A;America/Los_Angeles
6482;Concord Rgnl;Concord;United States;;KJQF;35.387775;-80.709136;705;-5;A;America/New_York
6483;Sandown;Isle Of Wight;United Kingdom;;EGHN;50.677776;-1.317803;18;0;E;Europe/London
6484;Fort William Heliport;Fort William;United Kingdom;FWM;\N;56.816666;-5.116667;0;0;E;Europe/London
6485;Navoi Airport;Navoi;Uzbekistan;NVI;UTSA;40.1172;65.1708;0;5;E;Asia/Samarkand
6486;La Defense Heliport;Paris;France;JPU;\N;48.86667;2.333333;0;1;E;Europe/Paris
6487;Andernos-Les-Bains;Andernos-Les-Bains;France;;LFCD;44.81111;-1.186023;0;1;E;Europe/Paris
6488;Ronda Airport;Ronda;Spain;RRA;\N;36.75;-5.166667;2500;1;E;Europe/Madrid
6489;Bienenfarm Airport;Nauen;Germany;;EDOI;52.6617;12.7458;131;1;E;Europe/Berlin
8131;Nguma Island Lodge Airstrip;Etsha 13;Botswana;;\N;-18.9538;22.3668;0;2;U;Africa/Gaborone
6491;Champion-7;Champion-7;Brunei;;\N;5.21;114.7419;12;8;N;\N
6492;Ain Arnat Airport;Setif;Algeria;QSF;DAAS;36.1781;5.32449;3360;1;U;Africa/Algiers
6493;La Rochelle-Ile de Re;La Rochelle;France;LRH;LFBH;46.1792;-1.19528;74;1;E;Europe/Paris
6494;Friedman Mem;Hailey;United States;SUN;KSUN;43.504444;-114.296194;5320;-7;A;America/Denver
6497;Yverdon-Les-Bains;Yverdon-Les-Bains;Switzerland;;LSGY;46.86667;6.7;1427;1;E;Europe/Zurich
6498;Portsmouth Airport;Portsmouth;United Kingdom;PME;\N;50.8;-1.083333;0;0;E;Europe/London
6499;Aleksandrowice;Bielsko-Biala;Poland;;EPBA;49.84997;19.020193;860;1;E;Europe/Warsaw
6500;Mason City Municipal;Mason City;United States;MCW;KMCW;43.2247;-93.4067;1243;-6;A;America/Chicago
6501;Salar de Uyuni;Salar de Uyuni;Bolivia;;\N;-20.330706;-67.046881;12000;-4;S;America/La_Paz
6502;Isla Pescado;Isla Pescado;Bolivia;;\N;-20.240932;-67.62538;12020;-4;S;America/La_Paz
6503;Toro Toro;Toro Toro;Bolivia;;\N;-18.133114;-65.768781;9007;-4;S;America/La_Paz
6512;QTHRL;Queenstown;New Zealand;;\N;-44.8898;168.6762;1;12;Z;Pacific/Auckland
6513;QTHRS;Queenstown;New Zealand;;\N;-44.9486;168.7068;1;12;Z;Pacific/Auckland
6514;Helirafting Start;Queenstown;New Zealand;;\N;-44.9486;168.7068;1300;12;Z;Pacific/Auckland
6515;Helirafting Landung;Queenstown;New Zealand;;\N;-44.8898;168.6762;1300;12;Z;Pacific/Auckland
6517;Jomo Kenyatta;Nairobi;Kenya;;\N;-1.319242;36.927775;1625;3;U;Africa/Nairobi
6518;DWEST;DWEST;Germany;;\N;51.025;7.1837;300;1;E;Europe/Berlin
6519;Niederoeblarn;Niederoeblarn;Austria;;LOGO;47.478;14.008;2142;1;E;Europe/Vienna
6520;Bad Voeslau;Bad Voeslau;Austria;;LOAV;47.965;16.259;765;1;E;Europe/Vienna
6521;Arekuna Camp;Arekuna;Venezuela;;\N;6.290501;-62.534499;0;-4.5;S;America/Caracas
6522;Uetersen;Uetersen;Germany;;EDHE;53.646667;9.704167;7;1;E;Europe/Berlin
6523;HLP1 HS-16;Salzwedel;Germany;;\N;52.8333;11.2027;111;1;U;Europe/Berlin
6524;HLP2 HS-16;Nordhausen;Germany;;\N;51.4888;10.7833;278;1;U;Europe/Berlin
6525;HLP HQ GT;Paetz;Germany;;\N;52.2361;13.6667;98;1;U;Europe/Berlin
6944;Jining Airport ;Jining;China;JNG;\N;35.417;116.533;1540;8;U;Asia/Chongqing
6943;Heilongjiang Mohe Airport;Mohe County;China;OHE;\N;52.915;122.427;1500;8;U;Asia/Chongqing
6942;Daqing Saertu Airport;Daqing;China;DAQ;\N;46.583333;125;1020;8;U;Asia/Chongqing
6941;Byron Airport;Byron;United States;;\N;37.8284;-121.626;79;-8;A;America/Los_Angeles
6940;Tunoshna;Yaroslavl;Russia;IAR;UUDL;57.560666676667;40.157369454444;270;4;N;Europe/Moscow
6939;Belaya;Sredniiy;Russia;;UIIB;52.91500001;103.57500001;1396;9;N;Asia/Irkutsk
6938;Borisoglebskoe;Kazan;Russia;;\N;55.86067;49.10099;384;4;N;Europe/Moscow
6937;Dzemgi;Komsomolsk-on-Amur;Russia;;UHKD;50.605;137.08167;82;11;N;Asia/Vladivostok
6936;Khabarovsk-MVL;Khabarovsk;Russia;;UHHT;48.528;135.179;213;11;N;Asia/Vladivostok
6935;Irkutsk-2;Irkutsk;Russia;;UIIR;52.3678;104.183;13411;9;N;Asia/Irkutsk
6934;Zhukovski;Ramenskoe;Russia;;UUBW;55.589;38.154;335;4;N;Europe/Moscow
6933;Gelendzhik;Gelendzhik;Russia;;URKG;44.55;38.083;100;4;N;Europe/Moscow
6932;Taganrog-Juzhnyi;Taganrog;Russia;;URRT;47.117;38.513;100;4;N;\N
6931;Neuchatel Airport;Neuchatel;Switzerland;QNC;LSGN;46.9575;6.86472;1427;1;E;Europe/Zurich
6930;Locarno Airport;Locarno;Switzerland;ZJI;LSZL;46.1608;8.87861;650;1;E;Europe/Zurich
6929;Speck-Fehraltorf Airport;Fehraltorf;Switzerland;;LSZK;47.3764;8.7575;1748;1;E;Europe/Zurich
6928;Lausanne-la Blecherette Airport;Lausanne;Switzerland;;LSGL;46.5453;6.61667;2041;1;E;Europe/Zurich
6927;Triengen Airport;Triengen;Switzerland;;LSPN;47.2267;8.07806;1594;1;E;Europe/Zurich
6926;Rimatara;Rimatara;French Polynesia;RMT;NTAM;-22.637253;-152.805192;56;-10;U;\N
6925;Ust-Ilimsk;Ust-Ilimsk;Russia;;\N;58.116667;102.4666667;1244;9;N;Asia/Irkutsk
6924;Ust-Kut;Ust-Kut;Russia;UKX;UITT;56.85;105.7167;2033;9;N;Asia/Irkutsk
6923;Kavalerovo;Kavalerovo;Russia;;\N;44.270556;135.054722;675;11;N;Asia/Vladivostok
6922;Kirensk;Kirensk;Russia;;UIKK;57.7667;108.05;780;9;N;Asia/Irkutsk
6921;Fortman Airport;St. Marys;United States;1OH;\N;40.5553253;-84.3866186;885;-5;U;America/New_York
6920;Bellona;Bellona;Solomon Islands;BN1;\N;15.981666666667;-11.3;3;0;U;Africa/Nouakchott
6919;Bellona;Bellona;Solomon Islands;BNY;\N;-11.302;159.8;3;11;U;Pacific/Guadalcanal
6918;Ringi Cove Airport;Ringi Cove;Solomon Islands;RIN;AGRC;-8.12639;157.143;0;11;U;Pacific/Guadalcanal
6917;Antonio Juarbe Pol Airport;Arecibo;Puerto Rico;ARE;TJAB;18.451111;-66.675556;23;-4;N;America/Puerto_Rico
6916;Pangborn Field;Wenatchee;United States;EAT;KEAT;47.398;-120.206;1249;-8;A;America/Los_Angeles
6915;Bendigo Airport;Bendigo;Australia;;YBDG;-36.7394;144.33;705;10;U;Australia/Hobart
6913;Aeroporto Prefeito Octavio de Almeida Neves;Sao Joao del Rei;Brazil;JDR;\N;-21.0864;-44.2258;3120;-3;S;America/Sao_Paulo
6914;RAAF Pearce;Perth;Australia;;YPEA;-31.66778;116.015;149;8;O;Australia/Perth
6911;Wangerooge Airport;Wangerooge;Germany;AGE;EDWG;53.7828;7.91389;7;1;E;Europe/Berlin
6910;Harle Airport;Harlesiel;Germany;;EDXP;53.7067;7.82028;7;1;E;Europe/Berlin
6909;Wittman Regional Airport;Oshkosh;United States;OSH;KOSH;44.024983;-88.551336;790;-6;A;America/Chicago
6908;Brest;Brest;Belarus;BQT;UMBB;52.06;23.53;468;1;E;Europe/Warsaw
6907;Ternopol;Ternopol;Ukraine;TNL;UKLT;49.31;25.42;1073;2;E;Europe/Kiev
6906;Chernigov;Chernigov;Ukraine;CEJ;UKRR;51.24;31.09;446;2;E;Europe/Kiev
6905;Lutsk;Lutsk;Ukraine;UKC;UKLC;50.6833;25.4833;774;2;U;Europe/Kiev
6904;Southwest Michigan Regional Airport;Benton Harbor;United States;BEH;\N;42.1285833;-86.4285;643;-5;A;America/New_York
6903;Waukesha County Airport;Waukesha;United States;UES;\N;43.0410278;-88.2370556;911;-6;A;America/Chicago
6902;Ronda;Ronda;Spain;;\N;36.73333;-5.166667;1000;1;U;Europe/Madrid
6901;Thurles;Thurles;Ireland;;\N;52.67888;-7.814369;500;0;U;Europe/Dublin
6900;Limerick;Limerick;Ireland;;\N;52.659;-8.624;500;0;U;Europe/Dublin
6899;Nowra Airport;Nowra;Australia;NOA;YSNW;-34.9489;150.537;400;10;O;Australia/Sydney
6898;RAAF Williams Laverton Base;Laverton;Australia;;YLVT;-37.8636;144.746;18;10;O;Australia/Hobart
6897;Tindal Airport;Katherine;Australia;KTR;YPTN;-14.5211;132.378;443;9.5;O;Australia/Darwin
6896;Amberley;Amberley;Australia;;YAMB;-27.64;152.712;91;10;O;Australia/Brisbane
6895;Geiranger;Geiranger;Norway;;GEIR;62.1;7.2;0;1;E;Europe/Oslo
6894;Zell am See;Zell am See;Austria;;LOWZ;47.2933;12.7883;2470;1;E;Europe/Vienna
6893;Galt Field Airport;Greenwood;United States;10C;\N;42.4028889;-88.3751111;875;-6;U;America/Chicago
6892;Everglades Airpark;Everglades;United States;X01;\N;25.8488611;-81.3902778;5;-5;U;America/New_York
6795;Choibalsan Airport;Choibalsan;Mongolia;COQ;ZMCD;48.1357;114.646;2457;8;U;Asia/Ulaanbaatar
6794;Taree Airport;Taree;Australia;TRO;YTRE;-31.8886;152.514;38;10;U;Australia/Sydney
6793;Orange Airport;Orange;Australia;OAG;YORG;-33.3817;149.133;3115;10;U;Australia/Sydney
6792;Grafton Airport;Grafton;Australia;GFN;YGFN;-29.7594;153.03;110;10;U;Australia/Sydney
6791;Marinduque Airport;Gasan;Philippines;MRQ;RPUW;13.360967;121.825583;32;8;N;Asia/Manila
6790;Hamadan Airport;Hamadan;Iran;HDM;OIHH;34.869167;48.5525;5755;3.5;E;Asia/Tehran
6789;St Augustin Airport;St-Augustin;Canada;YIF;CYIF;51.2117;-58.6583;20;-4;A;America/Blanc-Sablon
6788;Vieques Airport;Vieques Island;Puerto Rico;VQS;TJCG;18.1158;-65.4227;19;-4;U;America/Puerto_Rico
6787;Kalay Airport;Kalemyo;Myanmar;KMV;VYKL;23.188811;94.051094;499;6.5;U;Asia/Rangoon
6786;Terre-de-Haut Airport;Les Saintes;Guadeloupe;LSS;TFFS;15.8644;-61.5806;46;-4;U;America/Guadeloupe
6785;Yenisehir Airport;Yenisehir;Turkey;YEI;LTBR;40.255208;29.562569;764;2;E;Europe/Istanbul
6784;Tekirdağ Çorlu Airport;Çorlu;Turkey;TEQ;LTBU;41.13825;27.919094;574;2;E;Europe/Istanbul
6783;Sinop Airport;Sinop;Turkey;SIC;LTCM;42.0158;35.0664;20;2;E;Europe/Istanbul
6782;Mus Airport;Mus;Turkey;MSR;LTCK;38.747769;41.661236;4157;2;E;Europe/Istanbul
6781;Canakkale Airport;Canakkale;Turkey;CKZ;LTBH;40.137722;26.426777;23;2;E;Europe/Istanbul
6780;Anadolu Airport;Eskissehir;Turkey;AOE;LTBY;39.809858;30.519378;2588;2;E;Europe/Istanbul
6779;Katima Mulilo Airport;Mpacha;Namibia;MPA;FYKM;-17.6344;24.1767;3144;1;S;Africa/Windhoek
6778;Walvis Bay Airport;Walvis Bay;Namibia;WVB;FYWB;-22.9799;14.6453;299;1;S;Africa/Windhoek
6777;Capitan Corbeta C A Curbelo International Airport;Punta del Este;Uruguay;PDP;SULS;-34.855139;-55.094278;95;-3;S;America/Montevideo
6776;Sialkot Airport;Sialkot;Pakistan;SKT;OPST;32.5356;74.3639;810;5;N;Asia/Karachi
6775;Bonaventure Airport;Bonaventure;Canada;YVB;CYVB;48.0711;-65.4603;123;-5;A;America/Toronto
6774;Brus Laguna Airport;Brus Laguna;Honduras;BHG;MHBL;15.7631;-84.5436;2;-6;U;America/Tegucigalpa
6773;Samburu South Airport;Samburu South;Kenya;UAS;HKSB;0.530583;37.5342;3295;3;U;Africa/Nairobi
6772;Chaoyang Airport;Chaoyang;China;CHG;ZYCY;41.5381;120.435;0;8;U;Asia/Chongqing
6771;Walaha Airport;Walaha;Vanuatu;WLH;NVSW;-15.412;167.691;151;11;U;Pacific/Efate
6770;Tanjung Manis Airport;Tanjung Manis;Malaysia;TGC;WBTM;2.17784;111.202;15;8;N;Asia/Kuala_Lumpur
6769;Long Akah Airport;Long Akah;Malaysia;LKH;WBGL;3.3;114.783;289;8;N;Asia/Kuala_Lumpur
6768;Geneina Airport;Geneina;Sudan;EGN;HSGN;13.4817;22.4672;2650;3;U;Africa/Khartoum
6767;Togiak Airport;Togiak Village;United States;TOG;PATG;59.0528;-160.397;21;-9;A;America/Anchorage
6766;Port Heiden Airport;Port Heiden;United States;PTH;PAPH;56.9591;-158.633;95;-9;A;America/Anchorage
6765;King Cove Airport;King Cove;United States;KVC;PAVC;55.1163;-162.266;155;-9;A;America/Anchorage
6764;New Stuyahok Airport;New Stuyahok;United States;KNW;PANW;59.4499;-157.328;302;-9;A;America/Anchorage
6763;Igiugig Airport;Igiugig;United States;IGG;PAIG;59.324;-155.902;90;-9;A;America/Anchorage
6762;Shimla Airport;Shimla;India;SLV;VISM;31.081803;77.067967;5072;5.5;N;Asia/Calcutta
6761;Nanded Airport;Nanded;India;NDC;VAND;19.1833;77.3167;1250;5.5;N;Asia/Calcutta
6760;Kangra Airport;Kangra;India;DHM;VIGG;32.1651;76.2634;2525;5.5;N;Asia/Calcutta
6759;Shahre Kord Airport;Shahre Kord;Iran;CQD;OIFS;32.2972;50.8422;6723;3.5;E;Asia/Tehran
6758;Sege Airport;Sege;Solomon Islands;EGM;AGGS;-8.57889;157.876;0;11;U;Pacific/Guadalcanal
6757;Burgos Airport;Burgos;Spain;RGS;LEBG;42.357628;-3.620764;2945;1;E;Europe/Madrid
6756;Leon Airport;Leon;Spain;LEN;LELN;42.589;-5.655556;3006;1;E;Europe/Madrid
6755;Deering Airport;Deering;United States;DRG;PADE;66.0696;-162.766;21;-9;A;America/Anchorage
6754;Sugraly Airport;Zarafshan;Uzbekistan;AFS;UTSN;41.6139;64.2332;1396;5;U;Asia/Samarkand
6753;Mardin Airport;Mardin;Turkey;MQM;LTCR;37.2233;40.6317;1729;2;E;Europe/Istanbul
6752;Tacheng Airport;Tacheng;China;TCG;ZWTC;46.6725;83.3408;0;8;U;Asia/Chongqing
6751;Tocache Airport;Tocache;Peru;;SPCH;-8.183;-76.517;1500;-5;U;America/Lima
6750;Nueva Loja Airport;Lago Agrio;Ecuador;LGQ;SELA;0.093056;-76.8675;980;-5;U;America/Guayaquil
6749;Parsabade Moghan Airport;Parsabad;Iran;PFQ;OITP;39.603606;47.8815;251;3.5;E;Asia/Tehran
6748;Ilam Airport;Ilam;Iran;IIL;OICI;33.586606;46.404842;4404;3.5;E;Asia/Tehran
6747;Gorgan Airport;Gorgan;Iran;GBT;OING;36.909381;54.401339;-24;3.5;E;Asia/Tehran
6746;Sahand Airport;Maragheh;Iran;ACP;OITM;37.348017;46.127903;4396;3.5;E;Asia/Tehran
6745;Romblon Airport;Romblon;Philippines;TBH;RPVU;12.311;122.085;10;8;N;Asia/Manila
6744;Changzhoudao Airport;Wuzhou;China;WUZ;ZGWZ;23.4567;111.248;89;8;U;Asia/Chongqing
6743;Hami Airport;Hami;China;HMI;ZWHM;42.8414;93.6692;0;8;U;Asia/Chongqing
6742;Sand Point Airport;Sand Point;United States;SDP;PASD;55.315;-160.523;21;-9;A;America/Anchorage
6741;Gorakhpur Airport;Gorakhpur;India;GOP;VEGK;26.739708;83.449708;259;5.5;N;Asia/Calcutta
6740;Araracuara Airport;Araracuara;Colombia;ACR;SKAC;-0.58;-72.41;1250;-5;U;America/Bogota
6739;Hagerstown Regional Richard A Henson Field;Hagerstown;United States;HGR;KHGR;39.7079;-77.7295;704;-5;A;America/New_York
6738;Bella Coola Airport;Bella Coola;Canada;QBC;CYBD;52.3875;-126.596;117;-8;A;America/Vancouver
6737;Pajala Airport;Pajala;Sweden;PJA;ESUP;67.2456;23.0689;542;1;E;Europe/Stockholm
6736;Port Clarence Coast Guard Station;Port Clarence;United States;KPC;PAPC;65.2537;-166.859;10;-9;A;America/Anchorage
6735;Governador Valadares Airport;Governador Valadares;Brazil;GVR;SBGV;-18.8952;-41.9822;561;-3;S;America/Sao_Paulo
6734;Kirovsk-Apatity Airport;Apatity;Russia;KVK;ULMK;67.4633;33.5883;515;4;N;Europe/Moscow
6733;Cauayan Airport;Cauayan;Philippines;CYZ;RPUY;16.929861;121.753036;200;8;N;Asia/Manila
6732;Tambor Airport;Nicoya;Costa Rica;TMU;MRTR;9.73852;-85.0138;33;-6;U;America/Costa_Rica
6731;Arenal Airport;La Fortuna/San Carlos;Costa Rica;FON;MRAN;10.478;-84.6345;342;-6;U;America/Costa_Rica
6730;Imo Airport;Imo;Nigeria;QOW;DNIM;5.42706;7.20603;373;1;N;Africa/Lagos
6729;Arctic Village Airport;Arctic Village;United States;ARC;PARC;68.1147;-145.579;2092;-9;A;America/Anchorage
6728;Tasiujaq Airport;Tasiujaq;Canada;YTQ;CYTQ;58.6678;-69.9558;122;-5;A;America/Toronto
6727;Puvirnituq Airport;Puvirnituq;Canada;YPX;CYPX;60.0506;-77.2869;74;-5;A;America/Toronto
6726;Ormoc Airport;Ormoc City;Philippines;OMC;RPVO;11.057997;124.565322;83;8;N;Asia/Manila
6725;Noatak Airport;Noatak;United States;WTK;PAWN;67.5661;-162.975;88;-9;A;America/Anchorage
6724;Savoonga Airport;Savoonga;United States;SVA;PASA;63.6864;-170.493;53;-9;A;America/Anchorage
6723;Shishmaref Airport;Shishmaref;United States;SHH;PASH;66.2496;-166.089;12;-9;A;America/Anchorage
6722;Ruby Airport;Ruby;United States;RBY;PARY;64.7272;-155.47;653;-9;A;America/Anchorage
6721;Point Hope Airport;Point Hope;United States;PHO;PPHO;68.3488;-166.799;12;-9;A;America/Anchorage
6720;Mekoryuk Airport;Mekoryuk;United States;MYU;PAMY;60.3714;-166.271;48;-9;A;America/Anchorage
6719;Kivalina Airport;Kivalina;United States;KVL;PAVL;67.7362;-164.563;13;-9;A;America/Anchorage
6718;St Marys Airport;St Mary's;United States;KSM;PASM;62.0605;-163.302;311;-9;A;America/Anchorage
6717;Kaltag Airport;Kaltag;United States;KAL;PAKV;64.3191;-158.741;187;-9;A;America/Anchorage
6716;Hooper Bay Airport;Hooper Bay;United States;HPB;PAHP;61.5239;-166.147;7;-9;A;America/Anchorage
6715;Gambell Airport;Gambell;United States;GAM;PAGM;63.7668;-171.733;27;-9;A;America/Anchorage
6714;Atqasuk Edward Burnell Sr Memorial Airport;Atqasuk;United States;ATK;PATQ;70.4673;-157.436;96;-9;A;America/Anchorage
6713;Anvik Airport;Anvik;United States;ANV;PANV;62.6467;-160.191;309;-9;A;America/Anchorage
6712;Anaktuvuk Pass Airport;Anaktuvuk Pass;United States;AKP;PAKP;68.1336;-151.743;2103;-9;A;America/Anchorage
6946;Altay Airport;Altay;China;AAT;ZWAT;47.866667;88.116667;2491;8;U;Asia/Chongqing
6947;Tuzla;Null;Bosnia and Herzegovina;;LQTZ;44.458656;18.724783;784;1;E;Europe/Sarajevo
6948;Fort Worth NAS;Dallas;United States;;KNFW;32.769167;-97.441528;650;-6;A;America/Chicago
6949;Naypyidaw;Naypyidaw;Burma;ELA;VYEL;19.6235;96.201028;302;6.5;U;Asia/Rangoon
6950;Kyauktu;Kyauktu;Burma;;VYXG;21.406667;94.130278;1345;6.5;U;Asia/Rangoon
6951;Jan Mayensfield;Jan Mayen;Norway;ZXB;ENJA;70.961111;-8.575833;33;1;E;Arctic/Longyearbyen
6952;Bokepyin;Bokepyin;Burma;;VYBP;11.267;98.767;100;6.5;U;Asia/Rangoon
6953;Huanghua Intl;Changsha;China;HHA;\N;28.189158;113.219633;217;8;U;Asia/Chongqing
6954;Manzhouli;Manzhouli;China;NZH;\N;49.566667;117.329444;0;8;U;Asia/Chongqing
6955;Wuhai;Wuhai;China;WUA;ZBUH;39.794444;106.799444;0;8;U;Asia/Chongqing
6956;Gary Chicago International Airport;Gary;United States;GYY;KGYY;41.6163;-87.4128;591;-6;A;America/Chicago
6957;Brainerd Lakes Rgnl;Brainerd;United States;BRD;KBRD;46.398308;-94.138078;1226;-6;U;America/Chicago
6958;Greenbrier Valley Airport;Lewisburg;United States;LWB;KLWB;37.858333;-80.399444;2302;-5;U;America/New_York
6959;Pitt-Greenville Airport;Greenville;United States;PGV;KPGV;35.635278;-77.385278;27;-5;A;America/New_York
6960;Chefornak Airport;Chefornak;United States;CYF;PACK;60.149167;-164.285556;40;-9;A;America/Anchorage
6961;Oxnard - Ventura County;Oxnard;United States;OXR;KOXR;34.200833;-119.207222;15;-8;A;America/Los_Angeles
6962;Branson LLC;Branson;United States;BKG;KBBG;36.531994;-93.200556;1302;-6;A;America/Chicago
6963;Tongren;Tongren;China;TEN;ZUTR;27.884;109.31;0;8;U;Asia/Chongqing
6964;Jinggangshan;Jian;China;KNC;ZSJA;26.8575;114.737222;0;8;U;Asia/Chongqing
6965;Penn Station;Baltimore;United States;ZBP;\N;39.307222;-76.615556;66;-5;A;America/New_York
6966;Penn Station;New York;United States;ZYP;\N;40.7505;-73.9935;35;-5;A;America/New_York
6967;Niau;Niau;French Polynesia;NIU;NTKN;-16.119037;-146.368406;49;-10;U;Pacific/Tahiti
6968;Stratton ANGB - Schenectady County Airpor;Scotia NY;United States;SCH;KSCH;42.85245555;-73.928866666;378;-5;A;America/New_York
6969;Begishevo;Nizhnekamsk;Russia;NBC;UWKE;55.34;52.06;50;4;N;Europe/Moscow
6970;Bogovarovo;Bogovarovo;Russia;;\N;58.8535;47.0114;100;4;N;Europe/Moscow
6971;Summit Camp;Greenland Ice Cap;Greenland;;GSUM;72.5795841916667;-38.4591850305556;10552;0;N;America/Danmarkshavn
6972;Warri Airport;Osubi;Nigeria;QRW;DNSU;5.31;5.45;50;1;U;Africa/Lagos
6973;Volkel;Volkel;Netherlands;;EHVK;51.656389;5.708611;72;1;E;Europe/Amsterdam
6974;Sayak Airport;Siargao;Philippines;;RPNS;9.859097;126.014017;10;8;N;Asia/Manila
6975;Langeoog Airport;Langeoog;Germany;LGO;EDWL;53.7425;7.49778;7;1;E;Europe/Berlin
6976;Fane Airport;Fane;Papua New Guinea;FNE;\N;-8.55;147.083;0;10;U;Pacific/Port_Moresby
6977;Itokama Airport;Itokama;Papua New Guinea;ITK;\N;-9.2;148.25;0;10;U;Pacific/Port_Moresby
6978;Ononge Airport;Ononge;Papua New Guinea;ONB;\N;-8.583;147.2;0;10;U;Pacific/Port_Moresby
6979;Tapini Airport;Tapini;Papua New Guinea;TPI;\N;-8.35;146.983;0;10;U;Pacific/Port_Moresby
6980;Wanigela Airport;Wanigela;Papua New Guinea;AGL;\N;-9.333;149.15;0;10;U;Pacific/Port_Moresby
6981;Woitape Airport;Woitape;Papua New Guinea;WTP;\N;-8.533;147.25;0;10;U;Pacific/Port_Moresby
6982;Awaba Airport;Awaba;Papua New Guinea;AWB;\N;-8.014;142.75;0;10;U;Pacific/Port_Moresby
6983;Telefomin Airport;Telefomin;Papua New Guinea;TFM;\N;-5.117;141.633;0;10;U;Pacific/Port_Moresby
6984;Kappelen;Biel;Switzerland;;LSZP;47.089167;7.29;1437;1;E;Europe/Zurich
6985;Nelspruit Airport;Nelspruit;South Africa;NLP;FANS;-25.5;30.9138;2875;2;U;Africa/Johannesburg
6986;Cherkassy;Cherkassy;Ukraine;CKC;UKKE;49.416666;32.1333;114;2;E;Europe/Kiev
6987;Gotska Sandon Heliport;Gotland;Sweden;;\N;58.392393;19.193142;0;1;E;Europe/Stockholm
6988;Lauterhorn;Faro;Sweden;;\N;57.9521;19.0812;0;1;E;Europe/Stockholm
6989;St. Augustine Airport;St. Augustine Airport;United States;UST;KSGJ;29.959167;-81.339722;10;-5;A;America/New_York
6990;Mykolaiv International Airport;Nikolayev;Ukraine;NLV;UKON;47.0579;31.9198;184;2;E;Europe/Kiev
6991;Ramechhap;Ramechhap;Nepal;RHP;VNRC;27.394005;86.06144;1555;5.75;U;Asia/Katmandu
6992;Charles M Schulz Sonoma Co;Santa Rosa;United States;STS;KSTS;38.508978;-122.81288;125;-8;A;America/Los_Angeles
6993;Kissimmee Gateway Airport;Kissimmee;United States;ISM;KISM;28.289806;-81.437083;82;-5;A;America/New_York
6994;Lake City Municipal Airport;Lake City;United States;LCQ;KLCQ;30.181944;-82.576944;201;-5;A;America/New_York
6995;DeLand Municipal Airport;DeLand;United States;;KDED;29.066944;-81.283889;79;-5;A;America/New_York
6996;Haller Airpark Airport;Green Cove Springs;United States;;7FL4;29.903021;-81.685923;75;-5;A;America/New_York
6997;Santa Lucia PNP Airstrip;Santa Lucia;Peru;;SLPA;-8.337041;-76.385733;50;-5;U;America/Lima
6998;Logan-Cache;Logan;United States;LGU;KLGU;41.791;-111.852;4457;-7;A;America/Denver
6999;Brigham City;Brigham City;United States;BMC;KBMC;41.552;-112.062;4229;-7;A;America/Denver
7000;Malad City;Malad City;United States;MLD;KMLD;42.17;-112.289;4503;-7;A;America/Denver
7001;Aspen Pitkin County Sardy Field;Aspen;United States;ASE;KASE;39.2232;-106.869;7820;-7;A;America/Denver
7002;Hilton Head;Hilton Head;United States;HHH;KHHH;32.216;-80.752;10;-5;U;America/New_York
7003;Barataevka;Ulyanovsk;Russia;ULV;UWLL;54.268299;48.2267;463;4;N;Europe/Moscow
7004;Horog;Horog;Tajikistan;;UTOD;37.5;71.5;2000;4.5;U;Asia/Kabul
7005;Sabi Sabi Airport;Sabi Sabi;South Africa;GSS;\N;-24.9415;31.4446;1276;2;U;Africa/Johannesburg
7006;Philadelphia 30th St Station;Philadelphia;United States;ZFV;\N;39.9557;-75.182;0;-5;A;America/New_York
7007;KBWD;Brownwood;United States;BWD;\N;31.7936111;-98.9565;1387;-6;A;America/Chicago
7008;Mexia - Limestone County Airport;Mexia;United States;LXY;\N;31.6411783;-96.5144594;544;-6;A;America/Chicago
7009;Kerrville Municipal Airport;Kerrville;United States;ERV;KERV;29.9766667;-99.0854722;1617;-6;A;America/Chicago
7010;Birrfeld;Birrfeld;Switzerland;;LSZF;47.4436;8.23361;1300;1;E;Europe/Zurich
7011;Sussex Co;Georgetown;United States;GED;KGED;38.689194;-75.358889;50;-5;A;America/New_York
7012;Seal Cove Seaplane Base;Prince Rupert;Canada;ZSW;CZSW;54.3333;-130.283;0;-8;A;America/Vancouver
7013;Great Bend Municipal;Great Bend;United States;GBN;KGBD;38.344167;-98.859167;1887;-6;A;America/Chicago
7014;Hays Regional Airport;Hays;United States;HYS;KHYS;38.8422;-99.2732;1998;-6;A;America/Chicago
7015;Spirit Of St Louis;Null;United States;SUS;KSUS;38.662119;-90.652044;463;-6;A;America/Chicago
7016;Ely Municipal;Ely;United States;LYU;KELO;47.824444;-91.830833;1456;-6;A;America/Chicago
7017;Grand Rapids Itasca County;Grand Rapids MN;United States;GPZ;KGPZ;47.211111;-93.509722;413;-6;A;America/Chicago
7018;Thief River Falls;Thief River Falls;United States;TVF;KTVF;48.065556;-96.185;1116;-6;A;America/Chicago
7019;Eagle River;Eagle River;United States;EGV;KEGV;45.932333;-89.268283;1642;-6;A;America/Chicago
7020;Lakeland;Minocqua - Woodruff;United States;ARV;KARV;45.927778;-89.730833;1629;-6;A;America/Chicago
7021;Ankeny Regl Airport;Ankeny;United States;IKV;KIKV;41.691389;-93.566389;910;-6;A;America/Chicago
7022;Berens River;Berens River;Canada;YBV;CYBV;52.358889;-97.018333;728;-6;A;America/Winnipeg
7023;Corpus Christi NAS;Corpus Christi;United States;NGP;KNGP;27.692701;-97.290376;18;-6;A;America/Chicago
7024;Seaplane Base;Port Simpson;Canada;YPI;CYPI;54.566667;-130.433333;0;-8;A;America/Vancouver
7025;Avalon;Catalina Island;United States;AVX;KAVX;33.405;-118.415833;1602;-8;A;America/Los_Angeles
7026;Mojave;Mojave;United States;MHV;KMHV;35.059364;-118.151856;2791;-8;A;America/Los_Angeles
7027;Air Base;Interlaken;Switzerland;ZIN;LSMI;46.6766;7.87908;1893;1;E;Europe/Zurich
7028;Kenmore Air Harbor Inc Seaplane Base;Kenmore;United States;KEH;\N;47.7548;-122.259;14;-8;A;America/Los_Angeles
7029;Municipal;Corozal;Belize;CZH;\N;18.3822;-88.4119;40;-6;N;America/Belize
7030;Inisheer;Inisheer;Ireland;INQ;EIIR;53.0647;-9.5109;40;0;E;Europe/Dublin
7031;Winterlandeplatz;Maennlichen;Switzerland;;\N;46.610833;7.9425;7297;1;E;Europe/Zurich
7032;Paketzentrum;Ostermundigen;Switzerland;;\N;46.963;7.482;1854;1;E;Europe/Zurich
7033;Strezhevoy;Strezhevoy;Russia;SWT;UNSS;60.716667;77.65;164;7;N;Asia/Omsk
7034;Cashel;Cashel;Ireland;;\N;52.5158333;-7.8855556;440;0;U;Europe/Dublin
7035;Hutchinson Municipal Airport;Hutchinson;United States;HUT;KHUT;38.0655;-97.8606;1543;-6;A;America/Chicago
7036;Bagram AFB;Kabul;Afghanistan;BPM;OAIX;34.5646;69.1554;4895;4.5;U;Asia/Kabul
7037;Al Kharj AFB;Al Kharj;Saudi Arabia;;OEPS;24.0346;47.345;1651;3;N;Asia/Riyadh
7038;Eagle County Airport;Eagle;United States;EGA;\N;39.6427611;-106.9159347;6548;-7;U;America/Denver
7039;Oak Lawn Train Station;Oak Lawn;United States;;\N;41.7198242969588;-87.7488327026367;500;-6;U;America/Chicago
7040;Wrigleyville;Chicago;United States;;\N;41.948958067066;-87.6587533950806;500;-6;U;America/Chicago
7041;Gelendzik;Gelendzik;Russia;GDZ;\N;44.566666666667;38.016666666667;50;4;N;Europe/Moscow
7042;Rosecrans Mem;Rosecrans;United States;STJ;KSTJ;39.771944;-94.909706;826;-6;A;America/Chicago
7043;Hartford Union Station;Hartford;United States;ZRT;\N;41.76888;-72.6815;0;-5;A;America/New_York
7044;Stamford Amtrak Station;Stamford;United States;ZTF;\N;41.046937;-73.541493;0;-5;A;America/New_York
7045;Newark Penn Station;Newark;United States;ZRP;\N;40.734722;-74.164167;0;-5;A;America/New_York
7046;Papa Airport;Papa;Hungary;;LHPA;47.3636;17.5008;466;1;E;Europe/Budapest
7047;Cuxhaven Airport;Cuxhaven;Germany;NDZ;KNDZ;53.768612;8.644722;75;1;E;Europe/Berlin
7048;Volk Fld;Camp Douglas;United States;VOK;KVOK;43.938956;-90.253433;912;-6;A;America/Chicago
7049;BFT County Airport;Beauford;United States;BFT;KBFT;32.41083;-80.635;500;-5;A;America/New_York
7050;Adana-Incirlik Airbase;Adana;Turkey;UAB;KUAB;37.00028;35.41833;500;2;E;Europe/Istanbul
7051;Gunnison - Crested Butte;Gunnison;United States;GUC;KGUC;38.533889;-106.933056;7678;-7;A;America/Denver
7052;Xi\\'An Xiguan;Xi\\'AN;China;SIA;ZLSN;34.3767;109.12;0;8;U;Asia/Chongqing
7053;Zamperini Field Airport;Torrance;United States;TOA;KTOA;33.803392;-118.339611;101;-8;A;America/Los_Angeles
7054;Manistee County-Blacker Airport;Manistee;United States;MBL;KMBL;44.2725;-86.246944;621;-5;A;America/New_York
7055;Hickam Air Force Base;Honolulu;United States;;PHIK;21.318681;-157.922427;13;-10;N;Pacific/Honolulu
7056;Charlotte County-Punta Gorda Airport;Punta Gorda;United States;PGD;KPGD;26.919722;-81.990556;26;-5;A;America/New_York
7057;Grand Canyon Heliport;Grand Canyon;United States;JGC;\N;35.96666666;-112.13333333;2500;-7;A;America/Phoenix
7058;Northern Aroostook Regional Airport;Frenchville;United States;WFK;KFVE;47.285556;-68.312778;988;-5;A;America/New_York
7059;Chautauqua County-Jamestown;Jamestown;United States;JHW;KJHW;42.153333;-79.258056;525;-5;A;America/New_York
7060;Riviere Rouge - Mont-Tremblant International Inc. Airport;Mont-Tremblant;Canada;YTM;CYFJ;46.409444;-74.78;827;-5;A;America/Toronto
7061;Lake Cumberland Regional Airport;Somerset;United States;SME;KSME;37.053611;-84.615556;927;-5;A;America/New_York
7062;Shenandoah Valley Regional Airport;Weyers Cave;United States;SHD;KSHD;38.263889;-78.896389;1201;-5;A;America/New_York
7063;Devils Lake Regional Airport;Devils Lake;United States;DVL;KDVL;48.114444;-98.908611;1445;-6;A;America/Chicago
7064;Dickinson Theodore Roosevelt Regional Airport;Dickinson;United States;DIK;KDIK;46.7975;-102.801944;2592;-7;A;America/Denver
7065;Sidney-Richland Municipal Airport;Sidney;United States;SDY;KSDY;47.706944;-104.1925;1984;-7;A;America/Denver
7066;Chadron Municipal Airport;Chadron;United States;CDR;KCDR;42.8375;-103.095556;3297;-7;A;America/Denver
7067;Alliance Municipal Airport;Alliance;United States;AIA;KAIA;42.053333;-102.803889;3931;-7;A;America/Denver
7068;McCook Regional Airport;McCook;United States;MCK;KMCK;40.206389;-100.592222;2583;-6;A;America/Chicago
7069;Florida Keys Marathon Airport;Marathon;United States;MTH;KMTH;24.726111;-81.051389;7;-5;A;America/New_York
7070;Dawson Community Airport;Glendive;United States;GDV;KGDV;47.138611;-104.807222;749;-7;A;America/Denver
7071;LM Clayton Airport;Wolf Point;United States;OLF;KOLF;48.094444;-105.575;1986;-7;A;America/Denver
7072;Yellowstone Airport;West Yellowstone;United States;WYS;KWYS;44.688333;-111.1175;6644;-7;A;America/Denver
7073;San Luis Valley Regional Airport;Alamosa;United States;ALS;KALS;37.435;-105.866667;7539;-7;A;America/Denver
7074;Canyonlands Field;Moab;United States;CNY;KCNY;38.755;-109.754722;4555;-7;A;America/Denver
7075;Ely Airport;Ely;United States;ELY;KELY;39.299722;-114.841944;6259;-8;A;America/Los_Angeles
7076;Vernal Regional Airport;Vernal;United States;VEL;KVEL;40.440833;-109.51;5278;-7;A;America/Denver
7077;Sierra Blanca Regional Airport;Ruidoso;United States;SRR;KSRR;33.46285;-105.534751;6814;-7;A;America/Denver
7078;Show Low Regional Airport;Show Low;United States;SOW;KSOW;34.265556;-110.005556;6415;-7;A;America/Phoenix
7079;McCall Municipal Airport;McCall;United States;MYL;KMYL;44.889722;-116.101389;5021;-7;A;America/Denver
7080;Lemhi County Airport;Salmon;United States;SMN;KSMN;45.123889;-113.881389;4043;-7;A;America/Denver
7081;Mammoth Yosemite Airport;Mammoth Lakes;United States;MMH;KMMH;37.624049;-118.837772;7128;-8;A;America/Los_Angeles
7082;Friday Harbor Airport;Friday Harbor;United States;FRD;KFHR;48.521944;-123.024444;113;-8;A;America/Los_Angeles
7083;Orcas Island Airport;Eastsound;United States;ESD;KORS;48.708056;-122.910556;31;-8;A;America/Los_Angeles
7084;Anacortes Airport;Anacortes;United States;OTS;\N;48.498889;-122.6625;241;-8;A;America/Los_Angeles
7085;Astoria Regional Airport;Astoria;United States;AST;KAST;46.157972;-123.878694;15;-8;A;America/Los_Angeles
7086;Newport Municipal Airport;Newport;United States;ONP;KNOP;44.580361;-124.057917;160;-8;A;America/Los_Angeles
7087;Emmonak Airport;Emmonak;United States;EMK;PAEM;62.786111;-164.490833;13;-9;A;America/Anchorage
7088;Unalakleet Airport;Unalakleet;United States;UNK;PAUN;63.888333;-160.798889;21;-9;A;America/Anchorage
7089;Ugnu-Kuparuk Airport;Kuparuk;United States;UUK;PAKU;70.330833;-149.5975;67;-9;A;America/Anchorage
7090;Shageluk Airport;Shageluk;United States;SHX;PAHX;62.692222;-159.569167;79;-9;A;America/Anchorage
7091;Chuathbaluk Airport;Chuathbaluk;United States;CHU;PACH;61.579167;-159.215556;243;-9;A;America/Anchorage
7092;Nuiqsut Airport;Nuiqsut;United States;NUI;PAQT;70.21;-151.005556;38;-9;A;America/Anchorage
7093;Eek Airport;Eek;United States;EEK;PAEE;60.213611;-162.043889;15;-9;A;America/Anchorage
7094;Kasigluk Airport;Kasigluk;United States;KUK;PFKA;60.873333;-162.524444;40;-9;A;America/Anchorage
7095;Kwethluk Airport;Kwethluk;United States;KWT;PFKW;60.790278;-161.443611;30;-9;A;America/Anchorage
7096;Kwigillingok Airport;Kwigillingok;United States;KWK;PAGG;59.876389;-163.168611;18;-9;A;America/Anchorage
7097;Marshall Don Hunter Sr. Airport;Marshall;United States;MLL;PADM;61.8646418;-162.026111;103;-9;A;America/Anchorage
7098;Russian Mission Airport;Russian Mission;United States;RSH;PARS;61.775;-161.319444;51;-9;A;America/Anchorage
7099;Tuntutuliak Airport;Tuntutuliak;United States;WTL;\N;60.335278;-162.666944;16;-9;A;America/Anchorage
7100;Ekwok Airport;Ekwok;United States;KEK;\N;59.356944;-157.471111;135;-9;A;America/Anchorage
7101;Koliganek Airport;Koliganek;United States;KGK;PAJZ;59.726667;-157.259444;269;-9;A;America/Anchorage
7102;Levelock Airport;Levelock;United States;KLL;\N;59.128056;-156.858611;39;-9;A;America/Anchorage
7103;Manokotak Airport;Manokotak;United States;KMO;PAMB;58.990278;-159.05;51;-9;A;America/Anchorage
7104;Twin Hills Airport;Twin Hills;United States;TWA;\N;59.074444;-160.275;82;-9;A;America/Anchorage
7105;Chalkyitsik Airport;Chalkyitsik;United States;CIK;PACI;66.645;-143.74;544;-9;A;America/Anchorage
7106;Eagle Airport;Eagle;United States;EAA;PAEG;64.778056;-141.149722;908;-9;A;America/Anchorage
7107;Hughes Airport;Hughes;United States;HUS;PAHU;66.039167;-154.264722;299;-9;A;America/Anchorage
7108;Huslia Airport;Huslia;United States;HSL;PAHL;65.697778;-156.351389;213;-9;A;America/Anchorage
7109;Livingood Airport;Livingood;United States;LIV;\N;65.531111;-148.541111;696;-9;A;America/Anchorage
7110;Minto Airport;Minto;United States;MNT;\N;65.143611;-149.37;460;-9;A;America/Anchorage
7111;Nulato Airport;Nulato;United States;NUL;PANU;64.729444;-158.074167;399;-9;A;America/Anchorage
7112;Rampart Airport;Rampart;United States;RMP;\N;65.507778;-150.140833;302;-9;A;America/Anchorage
7113;Tanana Airport;Tanana;United States;TAL;\N;65.179556;-152.075833;207;-9;A;America/Anchorage
7114;Venetie Airport;Venetie;United States;VEE;PAVE;67.008611;-146.366389;574;-9;A;America/Anchorage
7115;Beaver Airport;Beaver;United States;WBQ;PAWB;66.362222;-147.406667;359;-9;A;America/Anchorage
7116;Central Airport;Central;United States;CEM;PACE;65.573889;-144.780833;937;-9;A;America/Anchorage
7117;Shungnak Airport;Shungnak;United States;SHG;PAGH;66.888056;-157.1625;197;-9;A;America/Anchorage
7118;Birch Creek Airport;Brich Creek;United States;KBC;\N;66.256708;-145.815319;450;-9;A;America/Anchorage
7119;Coldfoot Airport;Coldfoot;United States;CXF;\N;67.251389;-150.176111;1014;-9;A;America/Anchorage
7120;Inyokern Airport;Inyokern;United States;IYK;KIYK;35.658889;-117.829444;2455;-8;A;America/Los_Angeles
7121;Visalia Municipal Airport;Visalia;United States;VIS;KVIS;36.318611;-119.392778;295;-8;A;America/Los_Angeles
7122;Merced Municipal Airport;Merced;United States;MCE;KMCE;37.284722;-120.513889;156;-8;A;America/Los_Angeles
7123;Laguna de Los Patos International Airport;Colonia;Uruguay;CYR;SUCA;-34.4564;-57.7706;66;-3;U;America/Montevideo
7124;Camelo;Camelo;Uruguay;;SULO;-33.963882;-58.288612;10;-3;U;America/Montevideo
7125;Amarais Airport;Campinas;Brazil;CPQ;SDAM;-22.8592;-47.1082;2008;-3;S;America/Sao_Paulo
7126;Phoenix Goodyear;Goodyear;United States;;KGYR;33.423725;-112.374456;968;-7;N;America/Phoenix
7127;Park City;Park City;United States;;\N;40.659444;-111.4997222;7000;-7;U;America/Denver
7128;Toowoomba;Toowoomba;Australia;TWB;YTWB;-27.542778;151.916389;2086;10;N;Australia/Brisbane
7129;Ballera;Ballera;Australia;;YLLE;-27.408333;141.808333;385;10;N;Australia/Brisbane
7130;Gatton;Gatton;Australia;;YGAT;-27.5583000183105;152.341995239258;400;10;N;Australia/Brisbane
7131;Arkalyk Airport;Arkalyk;Kazakhstan;AYK;UAUR;50.2395;66.941;1152;6;U;Asia/Qyzylorda
7132;Hamburger Hafen;Hamburg;Germany;;ZZ06;53.542369;9.981592;0;1;E;Europe/Berlin
7133;Flugplatz Fehmarn-Neujellingsdorf;Neujellingsdorf;Germany;;ZZ05;54.455802;11.109273;0;1;E;Europe/Berlin
7134;Nabern Teck;Kirchheim-Teck;Germany;;EDTN;48.3676;9.2863;1215;1;E;Europe/Berlin
7135;Angoon Seaplane Base;Angoon;United States;AGN;PAGN;57.503611;-134.585;0;-9;A;America/Anchorage
7136;Elfin Cove Seaplane Base;Elfin Cove;United States;ELV;PAEL;58.195278;-136.3475;0;-9;A;America/Anchorage
7137;Tenakee Seaplane Base;Tenakee Springs;United States;TKE;\N;57.779722;-135.218333;0;-9;A;America/Anchorage
7138;Pelican Seaplane Base;Pelican;United States;PEC;\N;57.955278;-136.236389;0;-9;A;America/Anchorage
7139;Chatham Seaplane Base;Sitka;United States;CYM;\N;57.515;-134.946111;0;-9;A;America/Anchorage
7140;Funter Bay Seaplane Base;Funter Bay;United States;FNR;PANR;58.254444;-134.897778;0;-9;A;America/Anchorage
7141;Excursion Inlet Seaplane Base;Excursion Inlet;United States;EXI;\N;58.420556;-135.449167;0;-9;A;America/Anchorage
7142;Hoonah Airport;Hoonah;United States;HNH;PAOH;58.096111;-135.409722;19;-9;A;America/Anchorage
7143;Kake Airport;Kake;United States;AFE;PAFE;56.961389;-133.910278;172;-9;A;America/Anchorage
7144;Craig Seaplane Base;Craig;United States;CGA;\N;55.478889;-133.147778;0;-9;A;America/Anchorage
7145;Hollis Seaplane Base;Hollis;United States;HYL;\N;55.481667;-132.646111;0;-9;A;America/Anchorage
7146;Metlakatla Seaplane Base;Metakatla;United States;MTM;PAMM;55.131111;-131.578056;0;-9;A;America/Anchorage
7147;Thorne Bay Seaplane Base;Thorne Bay;United States;KTB;\N;55.688056;-132.536667;0;-9;A;America/Anchorage
7148;Hydaburg Seaplane Base;Hydaburg;United States;HYG;PAHY;55.206389;-132.828333;0;-9;A;America/Anchorage
7149;Hyder Seaplane Base;Hyder;United States;WHD;\N;55.903333;-130.006667;0;-8;A;America/Vancouver
7150;Point Baker Seaplane Base;Point Baker;United States;KPB;\N;56.351944;-133.6225;0;-9;A;America/Anchorage
7151;Port Protection Seaplane Base;Port Protection;United States;PPV;\N;56.328889;-133.61;0;-9;A;America/Anchorage
7152;North Whale Seaplane Base;North Whale Pass;United States;WWP;\N;56.116389;-133.121667;0;-9;A;America/Anchorage
7153;Chignik Lake Airport;Chignik Lake;United States;KCQ;\N;56.255;-158.775278;50;-9;A;America/Anchorage
7154;Egegik Airport;Egegik;United States;EGX;PAII;58.185556;-157.375556;92;-9;A;America/Anchorage
7155;Chignik Lagoon Airport;Chignik Lagoon;United States;KCL;\N;56.311111;-158.534167;25;-9;A;America/Anchorage
7156;Chignik Bay Seaplane Base;Chignik;United States;KBW;\N;56.295556;-158.401398;0;-9;A;America/Anchorage
7157;Perryville Airport;Perryville;United States;KPV;PAPE;55.906667;-159.160833;29;-9;A;America/Anchorage
7158;Pilot Point Airport;Pilot Point;United States;PIP;PAPN;57.585393;-157.571944;57;-9;A;America/Anchorage
7159;South Naknek Airport;South Naknek;United States;WSN;PFWS;58.702222;-157.0025;162;-9;A;America/Anchorage
7160;Akhiok Airport;Akhiok;United States;AKK;PAKH;56.938611;-154.1825;44;-9;A;America/Anchorage
7161;Karuluk Airport;Karluk;United States;KYK;PAKY;57.566944;-154.450278;137;-9;A;America/Anchorage
7162;Larsen Bay Airport;Larsen Bay;United States;KLN;PALB;57.535;-153.976667;87;-9;A;America/Anchorage
7163;Old Harbor Airport;Old Harbor;United States;OLH;\N;57.218056;-153.269722;55;-9;A;America/Anchorage
7164;Ouzinkie Airport;Ouzinkie;United States;KOZ;\N;57.922876;-152.500511;55;-9;A;America/Anchorage
7165;Port Lions Airport;Port Lions;United States;ORI;\N;57.885278;-152.846111;52;-9;A;America/Anchorage
7166;Alitak Seaplane Base;Lazy Bay;United States;ALZ;\N;56.899444;-154.247778;0;-9;A;America/Anchorage
7167;Amook Bay Seaplane Base;Amook Bay;United States;AOS;\N;57.471389;-153.815278;0;-9;A;America/Anchorage
7168;Kitoi Bay Seaplane Base;Kitoi Bay;United States;KKB;\N;58.190833;-152.370556;0;-9;A;America/Anchorage
7169;Moser Bay Seaplane Base;Moser Bay;United States;KMY;\N;57.025556;-154.145833;0;-9;A;America/Anchorage
7170;Olga Bay Seaplane Base;Olga Bay;United States;KOY;\N;57.161389;-154.229722;0;-9;A;America/Anchorage
7171;Port Bailey Seaplane Base;Port Bailey;United States;KPY;\N;57.93;-153.040556;0;-9;A;America/Anchorage
7172;Port Williams Seaplane Base;Port Williams;United States;KPR;\N;58.49;-152.582222;0;-9;A;America/Anchorage
7173;Seal Bay Seaplane Base;Seal Bay;United States;SYB;\N;58.166667;-152.5;0;-9;A;America/Anchorage
7174;San Juan - Uganik Seaplane Base;San Juan;United States;WSJ;\N;57.730278;-153.320556;0;-9;A;America/Anchorage
7175;West Point Village Seaplane Base;West Point;United States;KWP;\N;57.77;-153.548889;0;-9;A;America/Anchorage
7176;Zachar Bay Seaplane Base;Zachar Bay;United States;KZB;\N;57.55;-153.75;0;-9;A;America/Anchorage
7177;Ambler Airport;Ambler;United States;ABL;PAFM;67.106389;-157.8575;334;-9;A;America/Anchorage
7178;Buckland Airport;Buckland;United States;BKC;PABL;65.981667;-161.149167;31;-9;A;America/Anchorage
7179;Bob Baker Memorial Airport;Kiana;United States;IAN;PAIK;66.975833;-160.436667;166;-9;A;America/Anchorage
7180;Kobuk Airport;Kobuk;United States;OBU;PAOB;66.912222;-156.897222;137;-9;A;America/Anchorage
7181;Robert Curtis Memorial Airport;Noorvik;United States;ORV;PFNO;66.8175;-161.022222;55;-9;A;America/Anchorage
7182;Selawik Airport;Selawik;United States;WLK;PASK;66.6;-159.985833;17;-9;A;America/Anchorage
7183;Brevig Mission Airport;Brevig Mission;United States;KTS;PFKT;65.331389;-166.465833;35;-9;A;America/Anchorage
7184;Elim Airport;Elim;United States;ELI;PFEL;64.615;-162.270556;162;-9;A;America/Anchorage
7185;Golovin Airport;Golovin;United States;GLV;PAGL;64.550556;-163.007222;59;-9;A;America/Anchorage
7186;Teller Airport;Teller;United States;TLA;PATE;65.240278;-166.339444;294;-9;A;America/Anchorage
7187;Wales Airport;Wales;United States;WAA;PAIW;65.6225;-168.095;22;-9;A;America/Anchorage
7188;White Mountain Airport;White Mountain;United States;WMO;PAWM;64.689167;-163.412778;267;-9;A;America/Anchorage
7189;Council Airport;Council;United States;CIL;\N;64.897778;-163.703333;85;-9;A;America/Anchorage
7190;Koyuk Alfred Adams Airport;Koyuk;United States;KKA;PAKK;64.939444;-161.154167;154;-9;A;America/Anchorage
7191;St. Michael Airport;St. Michael;United States;SMK;PAMK;63.49;-162.110278;98;-9;A;America/Anchorage
7192;Shaktoolik Airport;Shaktoolik;United States;SKK;PFSH;64.371111;-161.223889;24;-9;A;America/Anchorage
7193;Stebbins Airport;Stebbins;United States;WBB;\N;63.515833;-162.278056;14;-9;A;America/Anchorage
7194;Tin City LRRS Airport;Tin City;United States;TNC;PATC;65.563056;-167.921667;269;-9;A;America/Anchorage
7195;Atka Airport;Atka;United States;AKB;PAAK;52.220278;-174.206389;56;-9;A;America/Anchorage
7196;Nikolski Air Station;Nikolski;United States;IKO;PAKO;52.941667;-168.848889;77;-9;A;America/Anchorage
7197;Icy Bay Airport;Icy Bay;United States;ICY;\N;59.968889;-141.661667;50;-9;A;America/Anchorage
7198;Yakataga Airport;Yakataga;United States;CYT;PACY;60.081901;-142.493611;12;-9;A;America/Anchorage
7199;Alakanuk Airport;Alakanuk;United States;AUK;PAUK;62.68;-164.66;10;-9;A;America/Anchorage
7200;Sheldon Point Airport;Nunam Iqua;United States;SXP;\N;62.520556;-164.847778;12;-9;A;America/Anchorage
7201;Kipnuk Airport;Kipnuk;United States;KPN;PAKI;59.933056;-164.030556;11;-9;A;America/Anchorage
7202;False Pass Airport;False Pass;United States;KFP;PAKF;54.8475;-163.410278;20;-9;A;America/Anchorage
7203;Nelson Lagoon;Nelson Lagoon;United States;NLG;PAOU;56.0075;-161.160278;14;-9;A;America/Anchorage
7204;Port Moller Airport;Cold Bay;United States;PML;PAAL;56.006111;-160.560833;20;-9;A;America/Anchorage
7205;Klawock Airport;Klawock;United States;KLW;PAKW;55.579167;-133.076111;80;-9;A;America/Anchorage
7206;Quinhagak Airport;Quinhagak;United States;KWN;PAQH;59.755;-161.845278;42;-9;A;America/Anchorage
7207;Kotlik Airport;Kotlik;United States;KOT;PFKO;63.030556;-163.532778;15;-9;A;America/Anchorage
7208;Koyukuk Airport;Koyukuk;United States;KYU;PFKU;64.875833;-157.730556;149;-9;A;America/Anchorage
7209;Scammon Bay Airport;Scammon Bay;United States;SCM;PACM;61.845278;-165.571389;14;-9;A;America/Anchorage
7210;Nondalton Airport;Nondalton;United States;NNL;PANO;59.966944;-154.851667;262;-9;A;America/Anchorage
7211;Pedro Bay Airport;Pedro Bay;United States;PDB;\N;59.782222;-154.1325;45;-9;A;America/Anchorage
7212;Nunapitchuk Airport;Nunapitchuk;United States;NUP;\N;60.905833;-162.439167;12;-9;A;America/Anchorage
7213;Kongiganak Airport;Kongiganak;United States;KKH;PADY;59.960833;-162.881111;30;-9;A;America/Anchorage
7214;Nikolai Airport;Nikolai;United States;NIB;PAFS;63.010833;-154.383889;427;-9;A;America/Anchorage
7215;Takotna Airport;Takotna;United States;TCT;\N;62.971944;-156.082778;825;-9;A;America/Anchorage
7216;Pilot Station Airport;Pilot Station;United States;PQS;\N;61.934444;-162.899444;305;-9;A;America/Anchorage
7217;Akiak Airport;Akiak;United States;AKI;PFAK;60.902778;-161.230556;30;-9;A;America/Anchorage
7218;Tuluksak Airport;Tuluksak;United States;TLT;\N;61.096944;-160.969444;30;-9;A;America/Anchorage
7219;Grayling Airport;Grayling;United States;KGX;\N;62.894444;-160.065;99;-9;A;America/Anchorage
7220;Wainwright Airport;Wainwright;United States;AIN;PAWI;70.638056;-159.994722;41;-9;A;America/Anchorage
7221;ZAPALA;ZAPALA;Argentina;APZ;SAHZ;-38.9755;-70.113581;3330;-3;N;America/Cordoba
7222;Rincon de los Sauces;Rincon de los Sauces;Argentina;RDS;\N;-37.390617;-68.904211;1969;-3;N;America/Cordoba
7223;Colonia Sarmiento;Colonia Sarmiento;Argentina;OLN;\N;-45.6;-69.083333;849;-3;N;America/Cordoba
7224;Grytvyken;Grytvyken;South Georgia and the Islands;;\N;-54.276667;-36.511667;538;-2;S;Atlantic/South_Georgia
7225;Rio Turbio;Rio Turbio;Argentina;RYO;\N;-51.533333;-72.3;1158;-3;N;America/Cordoba
7226;Puerto Natales;Puerto Natales;Chile;PNT;SCNT;-51.733333;-72.516667;9;-4;S;America/Santiago
7227;Caleta Olivia;Caleta Olivia;Argentina;CVI;\N;-46.4333;-67.5333;124;-3;N;America/Cordoba
7228;Fitz Roy;El Chalten;Argentina;;\N;-47.033333;-67.25;757;-3;N;America/Cordoba
7229;Sierra Grande;Sierra Grande;Argentina;SGV;SAVS;-41.592;-65.341;820;-3;N;America/Cordoba
7230;Ingeniero Jacobacci;Ingeniero Jacobacci;Argentina;IGB;SAVJ;-41.3;-69.5833;2936;-3;N;America/Cordoba
7231;Lago Posadas;Lago Posadas;Argentina;;\N;-47.533333;-71.75;748;-3;N;America/Cordoba
7232;El Chalten;El Chalten;Argentina;ELX;\N;-49.328889;-72.93;1279;-3;N;America/Cordoba
7233;Chenega Bay Airport;Chenega;United States;NCN;PFCB;60.077222;-147.991944;72;-9;A;America/Anchorage
7234;Chisana Airport;Chisana;United States;CZN;\N;62.071111;-142.048333;1011;-9;A;America/Anchorage
7235;Tok Junction Airport;Tok;United States;6K8;PFTO;63.329444;-142.953611;1639;-9;A;America/Anchorage
7236;Circle City Airport;Circle;United States;IRC;PACR;65.827778;-144.076111;613;-9;A;America/Anchorage
7237;Coffman Cove Seaplane Base;Coffman Cove;United States;KCC;\N;56.014722;-132.833889;0;-9;A;America/Anchorage
7238;Crooked Creek Airport;Crooked Creek;United States;CKD;\N;61.867778;-158.135;178;-9;A;America/Anchorage
7239;Red Devil Airport;Red Devil;United States;RDV;\N;61.788056;-157.350278;174;-9;A;America/Anchorage
7240;Sleetmute Airport;Sleetmute;United States;SLQ;PASL;61.700566;-157.165833;190;-9;A;America/Anchorage
7241;Stony River 2 Airport;Stony River;United States;SRV;\N;61.7875;-156.591111;230;-9;A;America/Anchorage
7242;Healy River Airport;Healy;United States;HKB;PAHV;63.8675;-148.968889;1263;-9;A;America/Anchorage
7243;Kake Seaplane Base;Kake;United States;KAE;\N;56.973056;-133.945556;0;-9;A;America/Anchorage
7244;Klawock Seaplane Base;Klawock;United States;AQC;PAQC;55.554658;-133.101693;0;-9;A;America/Anchorage
7245;Minchumina Airport;Lake Minchumina;United States;MHM;PAMH;63.886111;-152.301944;678;-9;A;America/Anchorage
7246;Manley Hot Springs Airport;Manley Hot Springs;United States;MLY;PAML;64.9975;-150.644167;270;-9;A;America/Anchorage
7247;St. George Airport;St. George;United States;STG;\N;56.577222;-169.663611;125;-9;A;America/Anchorage
7248;Tatitlek Airport;Tatitlek;United States;TEK;\N;60.8725;-146.691111;62;-9;A;America/Anchorage
7249;Ketchikan harbor Seaplane Base;Ketchikan;United States;WFB;\N;55.344444;-131.663333;0;-9;A;America/Anchorage
7250;Fox Harbour Airport;Fox Harbour;Canada;;\N;45.87;-63.461111;62;-4;A;America/Halifax
7251;Natuashish Airport;Natuashish;Canada;;CNH2;55.913889;-61.184444;33;-4;A;America/Halifax
7252;Postville Airport;Postville;Canada;YSO;CCD4;54.910278;-59.785278;223;-4;A;America/Halifax
7253;Kangiqsujuaq - Wakeham Bay Airport;Kangiqsujuaq;Canada;YWB;CYKG;61.588611;-71.929444;501;-5;A;America/Toronto
7254;Alma Airport;Alma;Canada;YTF;CYTF;48.508611;-71.641389;449;-5;A;America/Toronto
7255;Havre Saint-Pierre Airport;Havre-Saint-Pierre;Canada;YGV;CYGV;50.281944;-63.611389;124;-5;A;America/Toronto
7256;Rimouski Airport;Rimouski;Canada;YXK;CYXK;48.478056;-68.496944;82;-5;A;America/Toronto
7257;Vakarufahli Island;Vakarufahli Island;Maldives;;\N;3.578742;72.902075;0;5;U;Indian/Maldives
7258;Vakarufalhi Island;Vakarufalhi Island;Maldives;;\N;3.578742;72.902075;0;5;N;Indian/Maldives
7259;Tadoule Lake Airport;Tadoule Lake;Canada;XTL;CYBQ;58.706111;-98.512222;923;-6;A;America/Winnipeg
7260;Lac Brochet Airport;Lac Brochet;Canada;XLB;CZWH;58.614167;-101.468889;1212;-6;A;America/Winnipeg
7261;South Indian Lake Airport;South Indian Lake;Canada;XSI;CZSN;56.792778;-98.907222;951;-6;A;America/Winnipeg
7262;Brochet Airport;Brochet;Canada;YBT;CYBT;57.889444;-101.679167;1131;-6;A;America/Winnipeg
7263;Little Grand Rapids Airport;Little Grand Rapids;Canada;ZGR;CZGR;52.045;-95.466111;1008;-6;A;America/Winnipeg
7264;Cross Lake - Charlie Sinclair Memorial Airport;Cross Lake;Canada;YCR;CYCR;54.610833;-97.760278;707;-6;A;America/Winnipeg
7265;Red Sucker Lake Airport;Red Sucker Lake;Canada;YRS;CYRS;54.167222;-93.557222;729;-6;A;America/Winnipeg
7266;Rainbow Lake Airport;Rainbow Lake;Canada;YOP;CYOP;58.491389;-119.407778;1756;-7;A;America/Edmonton
7267;Bonnyville Airport;Bonnyville;Canada;YBY;CYBF;54.304722;-110.741111;1839;-7;A;America/Edmonton
7268;Nanaimo Harbour Water Airport;Nanaimo;Canada;ZNA;CAC8;49.183333;-123.95;0;-8;A;America/Vancouver
7269;Ganges Water Aerodrome;Ganges;Canada;;CAX6;48.85;-123.5;0;-8;A;America/Vancouver
7270;Bedwell Harbour Water Aerdrome;Bedwell Harbour;Canada;;\N;48.75;-123.233333;0;-8;A;America/Vancouver
7271;Qualicum Beach Airport;Qualicum Beach;Canada;XQU;\N;49.337222;-124.393889;191;-8;A;America/Vancouver
7272;Fort St. James - Perison Airport;Fort St. James;Canada;YJM;CYJM;54.397222;-124.262778;2364;-8;A;America/Vancouver
7273;Boundary Bay Airport;Boundary Bay;Canada;YDT;CZBB;49.073889;-123.0075;6;-8;A;America/Vancouver
7274;Langley Regional Airport;Langley Township;Canada;;CYNJ;49.101111;-122.630556;34;-8;A;America/Vancouver
7275;Bella Bella Airport;Bella Bella;Canada;ZEL;CYJQ;52.139722;-128.063611;162;-8;A;America/Vancouver
7276;Sechelt Aerodrome;Sechelt-Gibsons;Canada;YHS;\N;49.460556;-123.718611;250;-8;A;America/Vancouver
7277;Wekweeti Airport;Wekweeti;Canada;;CFJ2;64.190833;-114.076667;1206;-7;A;America/Edmonton
7278;Campo Cuatro Milpas Airport;Guasave;Mexico;;MM52;25.651944;-108.538056;72;-7;A;America/Mazatlan
7279;Isla de Cedros Airport;Cedros;Mexico;;MMCD;28.0375;-115.189444;98;-6;A;America/Mexico_City
7280;Cabo San Lucas International Airport;Cabo San Lucas;Mexico;;MMSL;22.9475;-109.937081;459;-7;A;America/Mazatlan
7281;Bahia Tortugas Airfield;Bahia Tortugas;Mexico;;\N;27.705278;-114.911111;102;-7;A;America/Mazatlan
7282;Palo Verde Airport;San Bruno;Mexico;PVP;\N;27.093056;-112.098889;55;-7;A;America/Mazatlan
7283;Ziyaraifushi;Ziyaraifushi;Maldives;;\N;4.531313889;73.373325;0;5;U;Indian/Maldives
7284;Brussels Gare du Midi;Brussels;Belgium;ZYR;\N;50.8;4.4;180;1;E;Europe/Brussels
7285;Caye Chapel Airport;Caye Chapel;Belize;CYC;\N;17.683611;-88.045;10;-6;U;America/Belize
7286;Big Creek Airport;Big Creek;Belize;BGK;\N;16.516667;-88.416667;16;-6;U;America/Belize
7287;Dangriga Airport;Dangriga;Belize;DGA;\N;16.966667;-88.216667;10;-6;U;America/Belize
7288;Placencia Airport;Placencia;Belize;PLJ;\N;16.536944;-88.361667;42;-6;U;America/Belize
7289;Sartaneja Airport;Sarteneja;Belize;SJX;\N;18.355556;-88.130833;10;-6;U;America/Belize
7290;Huehuetenango Airport;Huehuetenango;Guatemala;;\N;15.314722;-91.476111;1901;-6;U;America/Guatemala
7291;Corn Island Airport;Corn Island;Nicaragua;RNI;MNCI;12.168611;-83.0675;26;-6;U;America/Managua
7292;Bonanza Airport;Bonanza;Nicaragua;BZA;MNBZ;14.041667;-84.630556;597;-6;U;America/Managua
7293;Rosita Airport;Rosita;Nicaragua;RFS;MNRT;13.897222;-84.404722;207;-6;U;America/Managua
7294;Siuna Airport;Siuna;Nicaragua;SIU;MNSI;13.716667;-84.776944;480;-6;U;America/Managua
7295;Waspam Airport;Waspam;Nicaragua;WSP;MNWP;14.737778;-83.975833;115;-6;U;America/Managua
7296;San Carols Airport;San Carlos;Nicaragua;;MNSC;11.133333;-84.783333;131;-6;U;America/Managua
7297;Carrillo Airport;Carrillo;Costa Rica;RIK;MRCR;9.866667;-85.483333;10;-6;U;America/Costa_Rica
7298;Fussen;Fussen;Germany;FUS;FUSS;47.585;10.6866;800;1;U;Europe/Berlin
7299;John A. Osborne Airport;Montserrat;Montserrat;;\N;16.791389;-62.193333;550;-4;U;America/Montserrat
7300;Monte Plata Batley Juan Sanchez Field;Monte Plata;Dominican Republic;;MDJS;18.81;-69.79;1424;-4;U;America/Santo_Domingo
7301;Constanza Airport;Constanza;Dominican Republic;COZ;MDCZ;18.907626;-70.719852;3931;-4;U;America/Santo_Domingo
7302;Negril Aerodrome;Negril;Jamaica;NEG;MKNG;18.34;-78.335556;9;-5;U;America/Jamaica
7303;Bochum Railway;Bochum;Germany;EBO;\N;51.478506;7.222781;45;1;U;Europe/Berlin
7304;Fliegerhost ;Kaufbeuren;Germany;KFB;\N;47.874;10.6294;680;1;U;Europe/Berlin
7305;Munich Railway;Munich;Germany;ZMU;\N;48.1408;11.555;500;1;U;Europe/Berlin
7306;Nuernberg Railway;Nuernberg;Germany;ZAQ;\N;49.446389;11.081944;312;1;U;Europe/Berlin
7307;Jose Aponte de la Torre Airport;Ceiba;Puerto Rico;RVR;TJRV;18.245278;-65.643333;38;-4;A;America/Puerto_Rico
7308;Aeropuerto Internacional Valle del Conlara;Merlo;Argentina;RLO;\N;-32.349803;-65.179932;100;-3;N;America/Cordoba
7309;Charlotte Amalie Harbor;Charlotte Amalie;Virgin Islands;;VI22;18.338611;-64.940833;0;-4;A;America/St_Thomas
7310;Christiansted Harbor Seaplane Base;Christiansted;Virgin Islands;SSB;\N;17.747222;-64.705;0;-4;A;America/St_Thomas
7311;Alto Rio Senguer Airport;Alto Rio Senguer;Argentina;ARR;SAVR;-45.016667;-70.816667;2145;-3;N;America/Cordoba
7312;Jose de San Martin Airport;Jose de San Martin;Argentina;JSM;SAWS;-44.016667;-70.466667;4359;-3;N;America/Cordoba
7313;Uyuni Airport;Uyuni;Bolivia;UYU;SLUY;-20.466667;-66.833333;11990;-4;U;America/La_Paz
7314;Augsburg Railway;Augsburg;Germany;ZAU;\N;48.3655;10.8863;500;1;U;Europe/Berlin
7315;Mannheim Railway;Mannheim;Germany;ZMA;\N;49.479633;8.469858;200;1;E;Europe/Berlin
7316;Essen Railway;Essen;Germany;ZES;\N;51.451389;7.013889;200;1;U;Europe/Berlin
7317;Rurrenabaque Airport;Rerrenabaque;Bolivia;RBQ;SLRQ;-14.4275;-67.498056;898;-4;U;America/La_Paz
7318;Lancaster Amtrak Station;Lancaster;United States;;\N;40.05;-76.31;345;-5;A;America/New_York
7319;Ardmore Amtrak Station;Ardmore;United States;;\N;40.01;-75.29;360;-5;A;America/New_York
7320;Abaiang Atoll Airport;Abaiang Atoll;Kiribati;ABF;NGAB;1.8;173.04;1;12;U;Pacific/Tarawa
7321;Metropark Amtrak Station;Iselin;United States;;\N;40.568;-74.3275;1;-5;A;America/New_York
7322;St. Louis Downtown Airport;East St. Louis;United States;CPS;\N;38.5707244;-90.1562211;413;-6;U;America/Chicago
7323;Afobaka Airstrip;Afobaka;Suriname;;SMAF;4.9982;-54.9919;10;-3;U;America/Paramaribo
7324;Alalapadu Airstrip;Alapadu;Suriname;;SMDU;2.5232;-56.3241;10;-3;U;America/Paramaribo
7325;Albina Airstrip;Albina;Suriname;ABN;SMBN;5.516667;-54.05;10;-3;U;America/Paramaribo
7326;Lawa Anapaike Airstrip;Anapaike;Suriname;;SMLA;3.416667;-54.033333;10;-3;U;America/Paramaribo
7327;Apetina Airstrip;Apetina;Suriname;;SMPT;3.5017;-55.0614;10;-3;U;America/Paramaribo
7328;Botopassi Airstrip;Botopasi;Suriname;BTO;\N;4.233333;-55.45;10;-3;U;America/Paramaribo
7329;Djoemoe Airstrip;Djoemoe;Suriname;DOE;\N;4.016667;-55.483333;10;-3;U;America/Paramaribo
7330;Drietabbetje Airstrip;Drietabbetje;Suriname;DRJ;SMDA;4.11667;-54.66667;10;-3;U;America/Paramaribo
7331;Kabalebo Airstrip;Kabalebo;Suriname;;SMKA;4.406;-57.223;10;-3;U;America/Paramaribo
7332;Kayser Airstrip;Kayser;Suriname;;SMKE;3.1;-56.483;10;-3;U;America/Paramaribo
7333;Kwamelasemoetoe Airstrip;Kwamelasemoetoe;Suriname;;SMSM;2.33333;-56.783333;10;-3;U;America/Paramaribo
7334;Moengo Airstrip;Moengo;Suriname;;SMMO;5.616667;-54.4;10;-3;U;America/Paramaribo
7335;Majoor Henry Fernandes Airport;Nieuw Nickerie;Suriname;ICK;SMNI;5.955556;-57.039444;9;-3;U;America/Paramaribo
7336;Vincent Fayks Airport;Paloemeu;Suriname;OEM;SMPA;5.811111;-55.190833;10;-3;U;America/Paramaribo
7337;Sarakreek Airstrip;Sarakreek;Suriname;;SMSK;4.9;-55.083333;10;-3;U;America/Paramaribo
7338;Sipaliwini Airstrip;Sipaliwini;Suriname;;SMSI;1.96605;-56.0035;10;-3;U;America/Paramaribo
7339;Stoelmans Eiland Airstrip;Stoelmans Eiland;Suriname;SMZ;SMST;4.35;-54.41667;10;-3;U;America/Paramaribo
7340;Totness Airstrip;Totness;Suriname;TOT;SMCO;5.865833;-56.3275;10;-3;U;America/Paramaribo
7341;Wageningen Airstrip;Wageningen;Suriname;AGI;SMWA;5.76667;-56.63333;10;-3;U;America/Paramaribo
7342;Kaieteur International Airport;Kaieteur Falls;Guyana;KIA;PSKA;5.163333;-59.483333;95;-4;U;America/Guyana
7343;Codela Airport;Guapiles;Costa Rica;CSC;MRCA;10.414;-85.0917;328;-6;U;America/Costa_Rica
7344;Newport News Amtrak Station;Newport News;United States;;\N;37.0228;-76.4519;0;-5;A;America/New_York
7345;Portland Union Station;Portland;United States;;\N;45.529;-122.6768;0;-8;A;America/Los_Angeles
7346;Orinduik Airport;Orinduik;Guyana;ORJ;SYOR;4.7;-60.016667;10;-4;U;America/Guyana
7347;Annai Airport;Annai;Guyana;NAI;SYAN;3.95;-59.133333;10;-4;U;America/Guyana
7348;Apoteri Airport;Apoteri;Guyana;;SYAP;4.033333;-58.583333;10;-4;U;America/Guyana
7349;Imbaimadai Airport;Imbaimadai;Guyana;IMB;SYIB;5.69252;-60.28198;10;-4;U;America/Guyana
7350;Kamarang Airport;Kamarang;Guyana;KAR;SYKM;5.865278;-60.614167;10;-4;U;America/Guyana
7351;Mabaruma Airport;Mabaruma;Guyana;USI;SYMB;8.2;-59.783333;10;-4;U;America/Guyana
7352;Mahdia Airport;Mahdia;Guyana;MHA;SYMD;5.266667;-59.15;10;-4;U;America/Guyana
7353;Dr. Augusto Roberto Fuster International Airport;Pedro Juan Caballero;Paraguay;PJC;SGPJ;-22.641389;-55.829722;1873;-4;U;America/Asuncion
7354;Alcides Fernandez Airport;Acandi;Colombia;ACD;SKAD;8.516667;-77.3;50;-5;U;America/Bogota
7355;Los Colonizadores Airport;Saravena;Colombia;RVE;SKSA;6.916667;-71.9;10;-5;U;America/Bogota
7356;La Chorrera Airport;La Chorrera;Colombia;LCR;\N;-0.733333;-73.016667;10;-5;U;America/Bogota
7357;Batagay Airport;Batagay;Russia;;UEBB;67.648;134.695;696;11;N;Asia/Vladivostok
7358;La Macarena;La Macarena;Colombia;LMC;\N;2.179167;-73.7875;10;-5;U;America/Bogota
7359;Villa Garzon Airport;Villa Garzon;Colombia;VGZ;SKVG;0.978889;-76.605556;1248;-5;U;America/Bogota
7360;El Bagre Airport;El Bagre;Colombia;EBG;SKEB;7.596389;-74.808889;180;-5;U;America/Bogota
7361;Juan H. White;Caucasia;Colombia;CAQ;SKCU;7.968333;-75.198333;174;-5;U;America/Bogota
7362;Mandinga Airport;Condoto;Colombia;COG;SKCD;5.071667;-76.676389;10;-5;U;America/Bogota
7363;Golfo de Morrosquillo Airport;Tolu;Colombia;TLU;SKTL;9.511944;-75.586389;10;-5;U;America/Bogota
7364;Cabo Frio International Airport;Cabo Frio;Brazil;CFB;SBCB;-22.921667;-42.074167;14;-3;S;America/Sao_Paulo
7365;Westport Amtrak Station;Westport;United States;;\N;44.1871;-73.4519;1;-5;A;America/New_York
7366;Trenton Amtrak Station;Trenton;United States;;\N;40.21889;-74.75417;1;-5;A;America/New_York
7367;Sinop Airport;Sinop;Brazil;OPS;SWSI;-11.885;-55.586;10;-4;S;America/Campo_Grande
7368;Gurupi Airport;Gurupi;Brazil;GRP;SWGI;-11.728889;-49.068889;10;-3;S;America/Fortaleza
7369;Campo Alegre Airport;Santana do Araguaia;Brazil;CMP;SNKE;-9.505;-50.625;525;-4;S;America/Boa_Vista
7370;Breves Airport;Breves;Brazil;BVS;SNVS;-1.681944;-50.48;131;-4;S;America/Boa_Vista
7371;Soure Airport;Soure;Brazil;SFK;SNSW;-0.716944;-48.522778;33;-4;S;America/Boa_Vista
7372;Julio Belem Airport;Parintins;Brazil;PIN;SWPI;-2.627778;-56.735833;500;-4;S;America/Boa_Vista
7373;Barreiras Airport;Barreiras;Brazil;BRA;SNBR;-12.083333;-45;1356;-3;S;America/Fortaleza
7374;Confresa Airport;Santa Terezinha;Brazil;STZ;SWST;-10.47;-50.502778;650;-4;S;America/Campo_Grande
7375;Minacu Airport;Minacu;Brazil;MQH;SBMC;-13.526944;-48.220556;1053;-3;S;America/Sao_Paulo
7376;Araguaina Airport;Araguaina;Brazil;AUX;SWGN;-7.228333;-48.240833;771;-3;S;America/Fortaleza
7377;Novo Aripuana Airport;Novo Aripuana;Brazil;NVP;SWNA;-5.121389;-60.380556;500;-4;S;America/Boa_Vista
7378;Bom Futuro Airport;Lucas do Rio Verde;Brazil;LVR;SWFE;-13.05;-55.910833;500;-4;S;America/Campo_Grande
7379;Franca Airport;Franca;Brazil;FRC;SIMK;-20.538611;-47.400833;1040;-3;S;America/Sao_Paulo
7380;Dourados Airport;Dourados;Brazil;DOU;SSDO;-22.220833;-54.805833;1433;-4;S;America/Campo_Grande
7381;Labrea Airport;Labrea;Brazil;LBR;SWLB;-7.258889;-64.797778;225;-4;S;America/Boa_Vista
7382;Rondonopolis Airport;Rondonopolis;Brazil;ROO;SWRD;-16.466667;-54.633333;500;-4;S;America/Campo_Grande
7383;Tancredo Thomaz de Faria Airport;Guarapuava;Brazil;GPB;SBGU;-25.383333;-51.45;200;-3;S;America/Sao_Paulo
7384;Joacaba Airport;Joacaba;Brazil;JCB;SSJA;-27.172778;-51.500833;1713;-3;S;America/Sao_Paulo
7385;North Philadelphia Amtrak Station;Philadelphia;United States;;\N;39.9969556;-75.1487722;1;-5;A;America/New_York
7386;Aberdeen Railway Station;Aberdeen;United Kingdom;;\N;57.1436;-2.0985;1;0;U;Europe/London
7387;Glasgow Railway Station;Glasgow;United Kingdom;;\N;55.8622;-4.2512;0;0;U;Europe/London
7388;Edinburgh Waverly Station;Edinburgh;United Kingdom;ZXE;\N;55.952;-3.189;0;0;U;Europe/London
7389;Newcastle Railway Station;Newcastle Upon Tyne;United Kingdom;;\N;54.9686;-1.6171;0;0;U;Europe/London
7390;Leeds Railway Station;Leeds;United Kingdom;;\N;53.794;-1.547;0;0;U;Europe/London
7391;Manchester Picadilly Station;Manchester;United Kingdom;;\N;53.477;-2.23;0;0;U;Europe/London
7392;Liverpool Railway Station;Liverpool;United Kingdom;;\N;53.405;-2.979;0;0;U;Europe/London
7393;London Euston Railway Station;London;United Kingdom;;\N;51.5284;-0.1331;0;0;U;Europe/London
7394;General leite de Castro Airport;Rio Verde;Brazil;RVD;SWLC;-17.790278;-50.918333;2244;-3;S;America/Sao_Paulo
7395;Araxa Airport;Araxa;Brazil;AAX;SBAX;-19.563056;-46.960278;3276;-3;S;America/Sao_Paulo
7396;Maues Airport;Maues;Brazil;MBZ;SWMW;-3.383611;-57.718611;500;-4;S;America/Boa_Vista
7397;Borba Airport;Borba;Brazil;RBB;SWBR;-4.387778;-59.593889;500;-4;S;America/Boa_Vista
7398;Coari Airport;Coari;Brazil;CIZ;SWKO;-4.085;-63.140833;33;-4;S;America/Boa_Vista
7399;Barcelos Airport;Barcelos;Brazil;BAZ;SWBC;-0.975;-62.923889;500;-4;S;America/Boa_Vista
7400;Herbert Glacier;Juneau;United States;;\N;58.571385;-134.607754;0;-9;A;America/Anchorage
7401;Seattle Cruise Terminal;Seattle;United States;;\N;47.615884;-122.330017;0;-8;A;America/Los_Angeles
7402;Juneau Cruise Pier;Juneau;United States;;\N;58.335451;-134.414978;0;-9;A;America/Anchorage
7403;Skagway Cruise Pier;Skagway;United States;;\N;59.470896;-135.31723;0;-9;A;America/Anchorage
7404;Ketchikan Cruise Pier;Ketchikan ;United States;;\N;55.346327;-131.644707;0;-9;A;America/Anchorage
7405;Victoria Cruise Pier;Victoria;Canada;;\N;48.429144;-123.367023;0;-8;A;America/Vancouver
7406;Diamantino Airport;Diamantino;Brazil;DMT;SWDM;-14.408889;-56.445833;1837;-4;S;America/Campo_Grande
7407;Guanambi Airport;Guanambi;Brazil;GNM;SNGI;-14.216667;-42.783333;500;-3;S;America/Fortaleza
7408;Tsletsi Airport;Djelfa;Algeria;QDJ;DAFI;34.6657;3.351;10;1;U;Africa/Algiers
7409;Nzagi Airport;Nzagi;Angola;;FNZG;-7.716944;21.358056;2431;1;U;Africa/Luanda
7410;Catoca Airport;Catoca;Angola;;\N;-9.433017;20.311189;500;1;U;Africa/Luanda
7411;Lucapa Airport;Lucapa;Angola;LBZ;FNLK;-8.443056;20.732222;500;1;U;Africa/Luanda
7412;Kapanda Airport;Kapanda;Angola;KNP;FNCP;-9.771944;15.456111;500;1;U;Africa/Luanda
7413;Am Timan Airport;Am Timan;Chad;AMC;FTTN;11.034;20.274;1421;1;U;Africa/Ndjamena
7414;Sharq Al-Owainat Airport;Sharq Al-Owainat;Egypt;GSQ;HEOW;22.5806;28.7207;859;2;U;Africa/Cairo
7415;Eastern WV Regional Airport;Martinsburg;United States;MRB;KMRB;39.2407;-77.591;554;-5;A;America/New_York
7416;Awasa Airport;Awasa;Ethiopia;AWA;HALA;7.067;38.5;5149;3;U;Africa/Addis_Ababa
7417;Jijiga Airport;Jijiga;Ethiopia;JIJ;HAJJ;9.359722;42.7875;2000;3;U;Africa/Addis_Ababa
7418;Mekane Salam Airport;Mekane Selam;Ethiopia;MKS;HAMA;10.633333;38.783333;3000;3;U;Africa/Addis_Ababa
7419;Debre Marqos;Debre Marqos;Ethiopia;DBM;HADM;10.316667;37.733333;7440;3;U;Africa/Addis_Ababa
7420;Debre Tabor Airport;Debre Tabor;Ethiopia;DBT;HADT;11.966667;38;7740;3;U;Africa/Addis_Ababa
7421;Harar Meda Airport;Debre Zeyit;Ethiopia;QHR;HAHM;8.7163;39.0059;8850;3;U;Africa/Addis_Ababa
7422;Robe Airport;Goba;Ethiopia;GOB;HAGB;6.733333;44.266667;4000;3;U;Africa/Addis_Ababa
7423;Mayumba Airport;Mayumba;Gabon;MYB;FOOY;-3.416667;10.65;500;1;U;Africa/Libreville
7424;Mara Serena Airport;Masai Mara;Kenya;MRE;\N;-1.406111;35.008056;500;3;U;Africa/Nairobi
7425;Lewa Airport;Lewa;Kenya;;\N;0.1955;37.4699;2000;3;U;Africa/Nairobi
7426;Mulika Lodge Airport;Meru National Park;Kenya;;HKMK;0.230278;38.170556;2000;3;U;Africa/Nairobi
7427;Rumbek Airport;Rumbek;Sudan;RBX;HSMK;6.825;29.669;2000;3;U;Africa/Juba
7428;Yei Airport;Yei;Sudan;;HSYE;4.083;30.65;1000;3;U;Africa/Juba
7429;Cape Palmas Airport;Greenville;Liberia;CPA;GLCP;4.37902;-7.69695;500;0;U;Africa/Monrovia
7430;Ambatomainty Airport;Ambatomainty;Madagascar;AMY;\N;-17.1667;47.1833;500;3;U;Indian/Antananarivo
7431;Kyoto;Kyoto;Japan;UKY;\N;35.016667;135.766667;262;9;N;Asia/Tokyo
7432;Ecuvillens Airport;Ecuvillens;Switzerland;;LSGE;46.754997;7.076111;2293;1;E;Europe/Zurich
7433;Andermatt;Andermatt;Switzerland;;\N;46.63889;8.58889;4747;1;E;Europe/Zurich
7434;Wohlen Airfield;Wohlen bei Bern;Switzerland;;\N;46.96694;7.35806;1795;1;E;Europe/Zurich
7435;Bazaruto Island Airport;Bazaruto Island;Mozambique;BZB;\N;-21.542778;35.473056;500;2;U;Africa/Maputo
7436;Benguera Island Airport;Benguera Island;Mozambique;BCW;\N;-21.849167;35.436944;50;2;U;Africa/Maputo
7437;Inhaca Airport;Inhaca;Mozambique;;FQIA;-26;32.916667;500;2;U;Africa/Maputo
7438;Indigo Bay Lodge Airport;Indigo Bay Lodge;Mozambique;IBL;\N;-21.707222;35.452222;50;2;U;Africa/Maputo
7439;Gombe Lawanti International Airport;Gombe;Nigeria;;\N;9.2575;12.430278;500;1;U;Africa/Lagos
7440;Akwa Ibom International Airport;Uyo;Nigeria;;\N;5.05;7.933333;500;1;U;Africa/Lagos
7441;Katsina Airport;Katsina;Nigeria;;\N;13.007778;7.660278;50;1;U;Africa/Lagos
7442;Ouro Sogui Airport;Matam;Senegal;MAX;GOSM;15.593611;-13.322778;85;0;U;Africa/Dakar
7443;Bird Island Airport;Bird Island;Seychelles;BDI;FSSB;-3.721389;55.208611;10;4;U;\N
7444;K50 Airport;Mogadishu;Somalia;;\N;1.999167;44.974167;500;3;U;Africa/Mogadishu
7445;El Daein;El Daein;Sudan;;\N;14;32.316667;1000;3;U;Africa/Khartoum
7446;Wadi Halfa Airport;Wadi Halfa;Sudan;WHF;HSSW;21.800278;31.516389;1500;3;U;Africa/Khartoum
7447;Enfidha - Zine El Abidine Ben Ali International Airport;Enfidha;Tunisia;NBE;DTNZ;36.075833;10.438611;500;1;U;Africa/Tunis
7448;Kidepo Airport;Kidepo;Uganda;;HUKD;3.719167;33.754167;3904;3;U;Africa/Kampala
7449;Kitgum Airport;Kitgum;Uganda;;HUKT;3.278611;32.89;2493;3;U;Africa/Kampala
7450;Pakuba Airport;Pakuba;Uganda;PAF;HUPA;2.3275;31.5;2297;3;U;Africa/Kampala
7451;Svea Airport;Sveagruva;Svalbard;;ENSA;77.9;16.683333;29;1;U;Arctic/Longyearbyen
7452;Ny-Alesund Airport;Ny-Alesund;Svalbard;;ENAS;78.9275;11.874167;50;1;U;Arctic/Longyearbyen
7453;Hatay Airport;Hatay;Turkey;HTY;LTDA;36.362778;36.282222;25;2;U;Europe/Istanbul
7454;Kihnu Airfield;Kihnu;Estonia;;EEKU;58.148;24.003;50;2;U;Europe/Tallinn
7455;Ruhnu Airfield;Ruhnu;Estonia;;EERU;57.784;23.266;50;2;U;Europe/Tallinn
7456;Raivavae Airport;Raivavae;French Polynesia;RVV;NTAV;-23.87;-147.67;25;-10;N;Pacific/Tahiti
7457;Foshan;Foshan;China;FUO;\N;23.133333;113.28333;8;8;N;Asia/Chongqing
7458;Huizhou;Huizhou;China;HUZ;\N;23.083332;114.36667;23;8;N;Asia/Chongqing
7459;Lleida-Alguaire Airport;Lleida;Spain;ILD;LEDA;41.727778;0.535833;1148;1;E;Europe/Madrid
7460;Aeropuerto Capitan Fuentes Martinez;Porvenir;Chile;WPR;\N;-53.2537;-70.319228;104;-4;U;America/Santiago
7461;Ouessant Airport;Ouessant;France;;LFEC;48.458056;-5.095556;25;1;U;Europe/Paris
7462;Central Railway Station;Montreal;Canada;YMY;\N;45.499722;-73.566111;0;-5;A;America/Toronto
7463;Union Station;Toronto;Canada;YBZ;\N;43.645278;-79.380556;0;-5;A;America/Toronto
7464;Bildudalur Airport;Bildudalur;Iceland;BIU;BIBD;65.641389;-23.546111;26;0;N;Atlantic/Reykjavik
7465;Gjogur Airport;Gjogur;Iceland;GJR;BIGJ;65.995278;-21.326944;98;0;N;Atlantic/Reykjavik
7466;Saudarkrokur;Saudarkrokur;Iceland;SAK;BIKR;65.731667;-19.572778;8;0;N;Atlantic/Reykjavik
7467;Selfoss Airport;Selfoss;Iceland;;BISF;63.929167;-21.037778;45;0;N;Atlantic/Reykjavik
7468;Inishmaan Aerodrome;Inishmaan;Ireland;IIA;EIMN;53.091944;-9.57;13;0;U;Europe/Dublin
7469;Taldykorgan Airport;Taldykorgan;Kazakhstan;TDK;UAAT;45.016667;78.366667;1969;6;U;Asia/Qyzylorda
7470;Olgii Airport;Olgii;Mongolia;ULG;ZMUL;48.991667;89.919722;5610;7;U;Asia/Hovd
7471;Lille;Lille;France;XDB;\N;50.563333;3.08805;1;1;U;Europe/Paris
7472;Qurghonteppa International Airport;Kurgan Tyube;Tajikistan;;UTDT;37.862222;68.862778;5000;5;U;Asia/Dushanbe
7473;Vologda Airport;Vologda;Russia;VGD;ULWW;59.281667;39.946667;387;4;N;Europe/Moscow
7474;Severo-Evensk Airport;Evensk;Russia;;UHMW;61.921667;159.23;0;12;N;Asia/Magadan
7475;Olenyok Airport;Olenyok;Russia;;UERO;68.514722;112.48;847;10;N;Asia/Yakutsk
7476;Saskylakh Airport;Saskylakh;Russia;;UERS;71.927778;114.08;500;10;N;Asia/Yakutsk
7477;Lensk Airport;Lensk;Russia;;\N;60.719444;114.931944;700;10;N;Asia/Yakutsk
7478;Burevestnik Airport;Iturup Island;Russia;BVV;\N;44.92;147.621667;79;11;N;Asia/Vladivostok
7479;Okha Airport;Okha;Russia;OHH;\N;53.583333;142.933333;100;11;N;Asia/Vladivostok
7480;Leshukonskoye Airport;Arkhangelsk;Russia;LDG;ULAL;64.895833;45.722778;220;4;N;Europe/Moscow
7481;Nizhneangarsk Airport;Nizhneangarsk;Russia;;UIUN;55.801667;109.586667;1545;9;N;Asia/Irkutsk
7482;Taksimo Airport;Taksimo;Russia;;\N;56.361667;114.93;1500;9;N;Asia/Irkutsk
7483;Vanavara Airport;Vanavara;Russia;;UNIW;60.355;102.31;892;8;N;Asia/Krasnoyarsk
7484;Aykhal Airport;Aykhal;Russia;;UERA;65.959167;111.546389;1499;10;N;Asia/Yakutsk
7485;Uktus Airport;Yekaterinburg;Russia;;USSK;56.701667;60.79;643;6;N;Asia/Yekaterinburg
7486;Baykit Airport;Baykit;Russia;;UNIB;61.676667;96.355;853;8;N;Asia/Krasnoyarsk
7487;Biysk Airport;Biysk;Russia;;UNBI;52.466667;85.35;620;7;N;Asia/Omsk
7488;Huesca-Pirineos Airport;Huesca;Spain;HSK;LEHC;42.080833;-0.323333;1768;1;U;Europe/Madrid
7489;Ciudad Real Central Airport;Ciudad Real;Spain;CQM;LERL;38.856389;-3.97;636;1;U;Europe/Madrid
7490;Al Najaf International Airport;Najaf;Iraq;NJF;ORNI;31.991667;44.404167;500;3;U;Asia/Baghdad
7491;Hilversum Railway Station;Hilversum;Netherlands;QYI;\N;52.226389;5.181667;3;1;E;Europe/Amsterdam
7492;Colonsay Airport;Colonsay;United Kingdom;CSA;EGEY;56.0575;-6.243056;44;0;U;Europe/London
7493;Coll Airport;Coll;United Kingdom;COL;\N;56.633333;-6.557222;50;0;U;Europe/London
7494;Rock Hill York Co Bryant Airport;Rock Hill;United States;RKH;KUZA;34.9878;-81.0572;667;-5;A;America/New_York
7495;Allegheny County Airport;Pittsburgh;United States;AGC;KAGC;40.3544;-79.9302;1252;-5;A;America/New_York
7496;Cecil Field;Jacksonville;United States;NZC;KVQQ;30.2187;-81.8767;81;-5;A;America/New_York
7497;Fulton County Airport Brown Field;Atlanta;United States;FTY;KFTY;33.7791;-84.5214;841;-5;A;America/New_York
7498;Tresco Heliport;Tresco;United Kingdom;TSO;EGHT;49.945556;-6.331389;20;0;U;Europe/London
7499;Tarin Kowt Airport;Tarin Kowt;Afghanistan;TII;OATN;32.605278;65.864167;3500;4.5;U;Asia/Kabul
7500;Zaranj Airport;Zaranj;Afghanistan;ZAJ;OAZJ;30.969167;61.866944;1581;4.5;U;Asia/Kabul
7501;Chaghcharan Airport;Chaghcharan;Afghanistan;CCN;OACC;34.526667;65.271667;7383;4.5;U;Asia/Kabul
7502;Four Corners;Four Corners;United States;;\N;36.998976;-109.045172;1000;-7;U;America/Denver
7503;Fuyang Airport;Fuyang;China;FUG;ZSFY;32.9;115.816667;500;8;U;Asia/Chongqing
7504;Longyan Airport;Longyan;China;LCX;\N;25.674167;116.746389;300;8;U;Asia/Chongqing
7505;Baoshan Airport;Baoshan;China;BSD;ZPBS;25.053333;99.168333;500;8;U;Asia/Chongqing
7506;Xingyi Airport;Xingyi;China;ACX;\N;25.0882;104.9587;500;8;U;Asia/Chongqing
7507;Macau Ferry Pier;Macau;Macau;XZM;\N;22.197075;113.558911;0;8;U;Asia/Macau
7508;Liping Airport;Liping;China;HZH;\N;26.206;109.039;500;8;U;Asia/Chongqing
7509;Ocean Isle Beach Airport;Ocean Isle Beach;United States;60J;\N;33.9085056;-78.4366722;32;-5;U;America/New_York
7510;Stepanakert;Stepanakert;Azerbaijan;;UB13;39.901439;46.787031;2001;4;E;Asia/Baku
7511;Ohio State University Airport;Columbus;United States;OSU;KOSU;40.0798;-83.073;905;-5;U;America/New_York
7922;Rio Sidra;Rio Sidra;Panama;RSI;\N;8.966667;-80.333336;10;-5;U;America/Panama
7513;Addison;Addison;United States;ADS;KADS;32.9685594;-96.8364478;644;-6;A;America/Chicago
7514;Destin;Destin;United States;DTS;KDTS;30.4000611;-86.4714772;23;-6;A;America/Chicago
7515;Fort Jefferson;Fort Jefferson - Dry Tortugas;United States;RBN;\N;24.61667;-82.86667;0;-5;A;America/New_York
7516;Chernobayevka Airport;Kherson;Ukraine;KHE;UKOH;46.6758;32.5064;148;2;E;Europe/Kiev
7517;Ryans Creek Aerodrome;Stewart Island;New Zealand;SZS;NZRC;-46.899693;168.101592;100;12;U;Pacific/Auckland
7518;Assumption Island;Assumption Island;Seychelles;;FSAS;-9.7422;46.5067;10;4;U;Indian/Mahe
7519;Zhijiang Airport;Zhijiang;China;HJJ;\N;27.441389;109.699722;1000;8;U;Asia/Chongqing
7520;Aldabra;Assumption Island;Seychelles;;\N;-9.74;46.51;100;4;U;Indian/Mahe
7521;Yarmouth Airport;Yarmouth;Canada;YQI;CYQI;43.8269;-66.0881;141;-4;A;America/Halifax
7522;Kinston Regional Jetport;Kinston;United States;ISO;KISO;35.331389;-77.608889;94;-5;A;America/New_York
7523;First Flight Airport;Kill Devil Hills;United States;FFA;KFFA;36.02;-75.67;13;-5;A;America/New_York
7524;Pebble Island Airstrip;Pebble Island Settlement;Falkland Islands;;\N;-51.31;-59.61;8;-3;S;Atlantic/Stanley
7525;Sea Lion Island Landing Strip;Sea Lion Island;Falkland Islands;;\N;-52.4282;-59.0777;0;-3;S;Atlantic/Stanley
7526;Lively Settlement Airstrip;Lively Island;Falkland Islands;;\N;-52.03;-58.47;0;-3;S;Atlantic/Stanley
7527;Lincang Airport;Lincang;China;LNJ;ZPLC;23.738333;100.025;1500;8;U;Asia/Chongqing
7528;Wenshan Airport;Wenshan;China;WNH;\N;23.375833;104.243056;1000;8;U;Asia/Chongqing
7529;Ponta Pelada Airport;Manaus;Brazil;PLL;\N;-3.14604;-59.9863;267;-4;S;America/Boa_Vista
7530;Sao Gabriel da Cachoeira Airport;Sao Gabriel da Cachoeira;Brazil;SJL;\N;-0.148056;-66.9858;251;-4;S;America/Boa_Vista
7531;Maturaca;Maturaca;Brazil;;SWMK;0.628269;-66.115128;354;-4;S;America/Boa_Vista
7532;Carajas Airport;Parauapebas;Brazil;CKS;SBCJ;-6.11781;-50.0035;2064;-4;S;America/Boa_Vista
7533;Centro de Lancamento de Alcantara;Alcantara;Brazil;;SNCW;-2.373;-44.396389;148;-3;S;America/Fortaleza
7534;Cachimbo;Itaituba;Brazil;ITB;\N;-4.2446;-56.00384;0;-4;S;America/Boa_Vista
7535;Latur Airport;Latur;India;LTU;\N;18.411944;76.465;1584;5.5;U;Asia/Calcutta
7536;Matak Airport;Anambas Islands;Indonesia;MWK;WIOM;3.348119;106.25805;10;7;N;\N
7537;Mainz Finthen;Mainz;Germany;QFZ;\N;0;0;300;1;E;\N
7538;Wuerzburg;Wuerzburg;Germany;;EDFW;51.93749;8.38257;400;1;E;Europe/Berlin
7539;Mainz;Mainz;Germany;QMZ;\N;50.00829;8.27356;360;1;E;Europe/Berlin
7540;Berlin Gatow;Berlin;Germany;GWW;\N;52.475;13.139;160;1;E;Europe/Berlin
7541;Rheine Bentlage;Rheine;Germany;ZPQ;\N;52.27656;7.43843;120;1;E;Europe/Berlin
7542;Sao Jacinto;Aveiro;Portugal;;LPAV;40.72222;-8.816667;26;0;E;Europe/Lisbon
7543;Tana Toraja Airport;Toraja;Indonesia;TTR;\N;-3.416667;119.916664;0;8;N;Asia/Makassar
7544;Hopsten Air Base;Hopsten;Germany;;ETNP;52.3387;7.54133;423;1;E;Europe/Berlin
7545;Jose Aponte De La Torre Airport;Ceiba;Puerto Rico;;TJNR;18.2453;-65.6434;38;-4;S;America/Puerto_Rico
7546;Persian Gulf Airport;Khalije Fars;Iran;PGU;OIBP;27.379444;52.7375;27;3.5;U;Asia/Tehran
7547;Yasuj Airport;Yasuj;Iran;YES;OISY;30.700556;51.545;5939;3.5;U;Asia/Tehran
7548;Mosul International Airport;Mosul;Iraq;OSB;ORBM;36.305833;43.1475;719;3;U;Asia/Baghdad
7549;Tajima Airport;Toyooka;Japan;TJH;RJBT;35.512778;134.786944;578;9;U;Asia/Tokyo
7550;Amakusa Airfield;Amakusa;Japan;AXJ;RJDA;32.482222;130.158889;340;9;U;Asia/Tokyo
7551;Kikai Airport;Kikai;Japan;KKX;RJKI;28.321389;129.928056;15;9;U;Asia/Tokyo
7552;Aguni Airport;Aguni;Japan;AGJ;RORA;26.592778;127.240278;38;9;U;Asia/Tokyo
7553;Chongjin Airport;Chongjin;North Korea;;\N;41.801389;129.855;500;9;U;Asia/Pyongyang
7554;Haeju Airport;Haeju;North Korea;HAE;\N;38.00543;125.77863;131;9;U;Asia/Pyongyang
7555;Layang Layang Airport;Layang Layang Atoll;Malaysia;LAC;\N;7.372222;113.841667;5;8;N;\N
7556;Donoi Airport;Uliastai;Mongolia;;ZMDN;47.712778;96.524167;5724;8;U;Asia/Ulaanbaatar
7557;Bulgan Airport;Bulgan;Mongolia;UGA;ZMBN;48.854167;103.484167;3873;8;U;Asia/Ulaanbaatar
7558;Ulaangom Airport;Ulaangom;Mongolia;ULO;ZMUG;49.973333;92.079722;3500;7;U;Asia/Hovd
7559;Borongan Airport;Borongan;Philippines;BPR;RPVW;11.674167;125.478611;7;8;N;Asia/Manila
7560;Lubang Community Airport;Lubang;Philippines;LBX;RPLU;13.855833;121.105833;25;8;N;Asia/Manila
7561;Bentota Airport;Bentota;Sri Lanka;BJT;\N;6.416667;79.983333;100;5.5;U;Asia/Colombo
7562;Dickwella Airport;Dickwella;Sri Lanka;DIW;\N;5.966667;80.683333;100;5.5;U;Asia/Colombo
7563;Kulob Airport;Kulyab;Tajikistan;TJU;UTDK;37.981667;69.799444;2293;5;U;Asia/Dushanbe
7564;Cimei Airport;Cimei;Taiwan;CMJ;RCCM;23.266667;119.666667;63;8;U;Asia/Taipei
7565;Dasoguz Airport;Dasoguz;Turkmenistan;TAZ;UTAT;41.764722;59.833056;500;5;U;Asia/Ashgabat
7566;Barrow Island Airport;Barrow Island;Australia;BWB;YBWX;-20.798;115.406;25;8;U;Australia/Perth
7567;Morawa Airport;Morawa;Australia;MWB;\N;-29.211;116.009;899;8;U;Australia/Perth
7568;Exmouth Airport;Exmouth;Australia;EXM;\N;-21.933;114.128;50;8;U;Australia/Perth
7569;Derby Airport;Derby;Australia;DRB;YDBY;-17.39;123.68;26;8;U;Australia/Perth
7570;Walgett Airport;Walgett;Australia;WGE;YWLG;-30.0318;148.1222;439;10;U;Australia/Sydney
7571;Bathurst Island Airport;Bathurst Island;Australia;BRT;YBTI;-11.769167;130.619722;67;9.5;U;Australia/Darwin
7572;Dunk Island Airport;Dunk Island;Australia;DKI;YDKI;-17.939722;146.141944;38;10;U;Australia/Brisbane
7573;Lizard Island Airport;Lizard Island;Australia;LZR;YLZI;-14.673056;145.454444;40;10;O;Australia/Brisbane
7574;Hamilton Airport;Hamilton;Australia;HLT;YHML;-37.648889;142.065278;803;10;O;Australia/Hobart
7575;Halls Creek Airport;Halls Creek;Australia;HCQ;YHLC;-18.233889;127.669722;1346;8;U;Australia/Perth
7576;Fitzroy Crossing Airport;Fitzroy Crossing;Australia;FIZ;YFTZ;-18.178;125.591;374;8;U;Australia/Perth
7577;Ravensthorpe Airport;Ravensthorpe;Australia;RVT;YNRV;-33.797222;120.208056;209;8;U;Australia/Perth
7578;Wilkins Runway;Budd Coast;Antarctica;;YWKS;-66.689444;111.485833;2395;12;U;Antarctica/South_Pole
7579;Provo Municipal Airport;Provo;United States;PVU;KPVU;40.21805555;-111.72222222;4497;-7;A;America/Denver
7580;Steamboat Springs Airport-Bob Adams Field;Steamboat Springs;United States;SBS;KSBS;40.51625;-106.8663056;6882;-7;A;America/Denver
7581;Delta Municipal Airport;Delta;United States;DTA;KDTA;39.3806386;-112.5077147;4759;-7;A;America/Denver
7582;Richfield Minicipal Airport;Richfield;United States;RIF;KRIF;38.7364361;-112.0989444;5301;-7;A;America/Denver
7583;Carbon County Regional-Buck Davis Field;Price;United States;PUC;KPUC;39.609722;-110.75278;5957;-7;A;America/Denver
7584;Los Alamos Airport;Los Alamos;United States;LAM;KLAM;35.8798019;-106.2694153;7171;-7;A;America/Denver
7585;Borrego Valley Airport;Borrego Springs;United States;BXS;KBXS;33.2590278;-116.3209722;520;-8;A;America/Los_Angeles
7586;Lake Havasu City Airport;Lake Havasu City;United States;HII;KHII;34.5711111;-114.3582778;783;-7;N;America/Phoenix
7587;Winslow-Lindbergh Regional Airport;Winslow;United States;INW;KINW;35.0219167;-110.7225278;4941;-7;N;America/Phoenix
7588;Douglas Municipal Airport;Douglas;United States;DGL;KDGL;31.3426028;-109.5064544;4173;-7;N;America/Phoenix
7589;Marakei Airport;Marakei;Kiribati;MZK;NGMK;2.050278;173.266667;10;12;U;Pacific/Tarawa
7590;Abemama Atoll Airport;Abemama;Kiribati;AEA;NGTB;0.490833;173.828611;8;12;U;Pacific/Tarawa
7591;Aranuka Airport;Buariki;Kiribati;AAK;NGUK;0.185278;173.636389;6;10;U;\N
7592;Kuria Airport;Kuria;Kiribati;KUC;NGKT;0.228611;173.410556;10;10;U;\N
7593;Arorae Island Airport;Arorae;Kiribati;AIS;NGTR;-2.633333;179.816667;6;10;U;\N
7594;Tamana Airport;Tamana;Kiribati;TMN;NGTM;-2.5;175.983333;10;12;U;Pacific/Tarawa
7595;Beru Island Airport;Beru Island;Kiribati;BEZ;NGBR;-1.254722;176.007222;10;10;U;\N
7596;Nikunau Airport;Nikunau;Kiribati;NIG;NGNU;-1.35;176.45;10;12;U;Pacific/Tarawa
7597;Butaritari Atoll Airport;Butaritari;Kiribati;BBG;NGTU;3.086521;172.811465;5;12;U;Pacific/Tarawa
7598;Makin Airport;Makin;Kiribati;MTK;NGMN;3.383333;173;10;12;U;Pacific/Tarawa
7599;Maiana Airport;Maiana;Kiribati;MNK;NGMA;0.933333;173;10;10;U;\N
7600;Nonouti Airport;Nonouti;Kiribati;NON;NGTO;-0.616667;174.366667;10;10;U;\N
7601;Tabiteuea South Airport;Tabiteuea;Kiribati;TSU;NGTS;-1.35;174.8;10;10;U;\N
7602;Bosset Airport;Bosset;Papua New Guinea;BOT;\N;-7.240833;141.092333;60;10;U;Pacific/Port_Moresby
7603;Ine Airport;Ine;Marshall Islands;IMI;\N;7.016667;171.483333;4;12;U;Pacific/Majuro
7604;Tinak Airport;Tinak;Marshall Islands;TIC;\N;7.133333;171.916667;4;12;U;Pacific/Majuro
7605;Ebon Airport;Ebon;Marshall Islands;;\N;4.598889;168.753056;5;12;U;Pacific/Majuro
7606;Elenak Airport;Elenak;Marshall Islands;EAL;\N;9.083333;167.333333;10;12;U;Pacific/Majuro
7607;Lae Airport;Lae;Marshall Islands;LML;\N;8.921667;166.265556;10;12;U;Pacific/Majuro
7608;Airok Airport;Airok;Marshall Islands;AIC;\N;7.1;171.233333;10;12;U;Pacific/Majuro
7609;Enejit Airport;Enejit;Marshall Islands;EJT;\N;6.040278;171.984444;10;12;U;Pacific/Majuro
7610;Whitianga Airport;Whitianga;New Zealand;WTZ;NZWT;-36.835833;175.700278;500;12;U;Pacific/Auckland
7611;Takaka Aerodrome;Takaka;New Zealand;KTF;NZTK;-40.85;172.8;500;12;U;Pacific/Auckland
7612;Peleliu Airfield;Peleliu;Palau;C23;\N;6.998333;134.232778;9;9;U;Pacific/Palau
7613;Angaur Airstrip;Angaur;Palau;;\N;6.906389;134.145;20;9;U;Pacific/Palau
7614;Asau Airport;Savai\\'i;Samoa;AAU;\N;-13.505;-172.627778;6;13;U;Pacific/Apia
7615;Afutara Airport;Afutara;Solomon Islands;AFT;AGAF;-9.183056;160.949722;5;11;U;Pacific/Guadalcanal
7616;Ulawa Airport;Ulawa;Solomon Islands;RNA;AGAR;-9.854722;161.979167;5;11;U;Pacific/Guadalcanal
7617;Choiseul Bay Airport;Choiseul Bay;Solomon Islands;CHY;AGGC;-6.711944;156.396111;5;11;U;Pacific/Guadalcanal
7618;Santa Ana Airport;Santa Ana;Solomon Islands;NNB;AGGT;-10.848056;162.454167;6;11;U;Pacific/Guadalcanal
7619;Yandina Airport;Yandina;Solomon Islands;XYA;AGGY;-9.086111;159.218889;6;11;U;Pacific/Guadalcanal
7620;Batuna Airport;Batuna;Solomon Islands;BPF;AGBT;-8.63;158;5;11;U;Pacific/Guadalcanal
7621;Bartow Municipal Airport;Bartow;United States;BOW;KBOW;27.9434;-81.7834;125;-5;U;America/New_York
7622;Sokerkino;Kostroma;Russia;KMW;UUBD;57.7961;41.0204;446;4;N;Europe/Moscow
7623;Yuryevets;Yuryevets;Russia;;\N;57.316667;43.1;0;4;N;Europe/Moscow
7624;Ozerny;Ozerny;Russia;;\N;52.752919;111.849;0;9;N;Asia/Irkutsk
7625;Enkheluk;Enkheluk;Russia;;\N;52.500967;107.018852;0;9;N;Asia/Irkutsk
7626;Khakusy;Khakusy;Russia;;\N;55.365;109.812;0;9;N;Asia/Irkutsk
7627;Fitiuta Airport;Fiti\\'uta;American Samoa;FTI;NSFQ;-14.216111;-169.423611;110;-11;U;Pacific/Pago_Pago
7628;Ofu Airport;Ofu;American Samoa;OFU;\N;-14.184444;-169.67;9;-11;U;Pacific/Pago_Pago
7629;Livermore Municipal;Livermore;United States;LVK;KLVK;37.41362;-121.49133;400;-8;A;America/Los_Angeles
7630;MariposaYosemite;Mariposa;United States;MPI;KMPI;37.3039;-120.0222;2454;-8;A;America/Los_Angeles
8419;Jungseok;Seogwipo;South Korea;;RKPD;33.399561;126.711567;1171;9;U;Asia/Seoul
7633;Grootfontein;Grootfontein;Namibia;GFY;FYGF;-19.602167;18.122667;4636;1;S;Africa/Windhoek
7634;Rundu;Rundu;Namibia;NDU;FYRU;-17.956461;19.719439;3627;1;S;Africa/Windhoek
7635;Beppu Airport;Beppu;Japan;BPU;\N;33.3;131.5333;197;9;N;Asia/Tokyo
7636;Heron Island;Heron Island;Australia;HRN;\N;-23.441667;151.911389;0;10;O;Australia/Brisbane
7637;Lady Elliot Island;Lady Elliot Island;Australia;LYT;\N;-24.113333;152.715278;0;10;O;Australia/Brisbane
7638;Orpheus Island;Orpheus Island;Australia;ORS;\N;-18.634;146.5;0;10;O;Australia/Brisbane
7639;Paddington Station;London;United Kingdom;QQP;\N;51.515833;-0.176111;0;0;E;Europe/London
7640;Liskeard Station;Liskeard;United States;;\N;50.44708;-4.46753;0;0;E;Europe/London
7641;Port du Bloscon;Roscoff;France;;\N;48.721816;-3.96395;0;1;E;Europe/Paris
7642;Tasiilaq;Angmagssalik;Greenland;AGM;BGAM;65.612222;-37.618333;24;-3;U;America/Godthab
7643;Neets Bay;Ketchikan;United States;;\N;55.778889;-131.601389;0;-9;U;America/Anchorage
7644;Fraser Railroad Station;Fraser BC;Canada;;\N;59.715517;-135.046692;2000;-8;U;America/Vancouver
7645;Carcross;Carcross YT;Canada;;\N;60.17463;-134.698088;2100;-8;U;America/Vancouver
7646;Jacqueline Cochran Regional Airport;Palm Springs;United States;TRM;KTRM;33.626666;-116.1596667;0;-8;A;America/Los_Angeles
7647;Santa Monica Municipal Airport;Santa Monica;United States;SMO;KSMO;34.0158333;-118.4513056;177;-8;A;America/Los_Angeles
7648;Bermuda Dunes Airport;Palm Springs;United States;UDD;KUDD;33.7484375;-116.2748133;73;-8;A;America/Los_Angeles
7649;Scottsdale Airport;Scottsdale;United States;ZSY;KSDL;33.6228889;-111.9105278;1519;-7;A;America/Phoenix
7650;Olympia Regional Airpor;Olympia;United States;OLM;KOLM;46.9694044;-122.9025447;209;-8;A;America/Los_Angeles
7651;Yolo County Airport;Davis-Woodland-Winters;United States;DWA;KDWA;38.5793889;-121.8569444;100;-8;A;America/Los_Angeles
7652;Garfield County Regional Airport;Rifle;United States;RIL;KRIL;39.5263056;-107.7269444;5548;-7;A;America/Denver
7653;Shively Field Airport;SARATOGA;United States;SAA;KSAA;41.4448594;-106.8235264;7012;-7;A;America/Denver
7654;Dekalb-Peachtree Airport;Atlanta;United States;PDK;KPDK;33.8756111;-84.3019722;1003;-5;A;America/New_York
7655;Monroe County Airport;Bloomington;United States;BMG;KBMG;39.1460208;-86.6166805;846;-5;A;America/New_York
7656;Witham Field Airport;Stuart;United States;SUA;KSUA;27.1816996;-80.221294;16;-5;A;America/New_York
7657;Morristown Municipal Airport;Morristown;United States;MMU;KMMU;40.79935;-74.4148747;187;-5;A;America/New_York
7658;Napa County Airport;Napa;United States;APC;KAPC;38.2131944;-122.2806944;35;-8;A;America/Los_Angeles
7659;Brown Field Municipal Airport;San Diego;United States;SDM;KSDM;32.5722722;-116.9801611;526;-8;A;America/Los_Angeles
7660;Wangen-Lachen;Wangen-Lachen;Switzerland;;LSPV;47.204805;8.866834;1335;1;U;Europe/Zurich
7661;Pahokee Airport;Pahokee;United States;PHK;\N;26.789;-80.692;18;-5;A;America/New_York
7662;Venice Municipal;Venice;United States;;KVNC;27.42439;-82.262364;18;-5;A;America/New_York
7663;Pahokee;Pahokee;United States;;KPHK;26.471041;-80.413591;15;-5;A;America/New_York
7664;Kineshma;Kineshma;Russia;KIE;\N;57.45000001;42.15000001;420;4;N;Europe/Moscow
7665;Nezhitino;Nezhitino;Russia;NEZ;\N;57.479444454444;43.29750001;400;4;N;Europe/Moscow
7666;Glasgow Buchanan Bus Station;Glasgow;United Kingdom;;\N;55.8652;-4.25033;0;0;E;Europe/London
7667;London Victoria Bus Station;London;United Kingdom;;\N;51.494999;-0.144643;0;0;E;Europe/London
7668;Machu Pichu Airport;Machu Pichu;Peru;MFT;\N;-13.1167;-72.5667;9500;-5;S;America/Lima
7669;Panama City-NW Florida Bea.;Panama City;United States;ECP;KECP;30.3417;-85.7973;69;-6;A;America/Chicago
7670;San Bernardino International Airport;San Bernardino;United States;SBD;KSBD;34.0953521;-117.2348722;1159;-8;A;America/Los_Angeles
7671;Valenca Airport;Valenca;Brazil;VAL;SNVB;-13.2965;-38.9924;21;-3;S;America/Fortaleza
7672;Dix Sept Rosado Airport;Mossoro;Brazil;MVF;SBMW;-5.20192;-37.3643;76;-3;S;America/Fortaleza
7673;Caruaru Airport;Caruaru;Brazil;CAU;SNRU;-8.28239;-36.0135;1891;-3;S;America/Fortaleza
7674;Wake Island Afld;Wake island;Wake Island;AWK;PWAK;19.282067;166.636444;14;12;U;Pacific/Wake
7675;Aeroclube de Nova Iguacu;Nova Iguacu;Brazil;QNV;SDNY;-22.7453;-43.4603;164;-3;S;America/Sao_Paulo
7676;Gare du Nord;Paris;France;XPG;\N;48.880931;2.355323;423;1;E;Europe/Paris
7677;Gare Montparnasse;Paris;France;XGB;\N;48.84;2.318611;423;1;E;Europe/Paris
7678;Saint-Pierre-des-Corps;Tours;France;XSH;\N;47.385626;0.723347;159;1;E;Europe/Paris
7679;Darsena Norte;Buenos Aires;Argentina;;\N;-34.5978;-58.367841;46;-3;N;America/Cordoba
7680;Casa Central;Montevideo;Uruguay;;\N;-34.90278;-56.213781;105;-3;S;America/Montevideo
7681;Puerto Franco;Colonia del Sacramento;Uruguay;;\N;-34.4741;-57.843427;89;-3;S;America/Montevideo
7682;Tres Cruces;Montevideo;Uruguay;;\N;-34.893891;-56.166589;141;-3;S;America/Montevideo
7683;San Carlos Airport;San Carlos;United States;SQL;KSQL;37.511944;-122.249444;5;-8;A;America/Los_Angeles
7684;Courtelary ;Courtelary ;Switzerland;;LSZJ;47.18611;7.15833;2250;1;E;Europe/Zurich
7685;Koszalin - Zegrze Pomorskie Airport;Koszalin;Poland;OSZ;EPKO;54.041;16.266;111;1;E;Europe/Warsaw
7686;Ntswi Island;Okavango Delta;Botswana;;FBCO;-19.531109;23.09092;3000;2;U;Africa/Gaborone
7687;Dujiangyan;Dujiangyan;China;;\N;30.998584;103.619673;2385;8;N;Asia/Chongqing
7688;Maiwa;Maiwa;China;;\N;33.058044;102.908775;11500;8;N;Asia/Chongqing
7689;Lelystad Airport;Lelystad;Netherlands;LEY;\N;52.4603;5.52722;12;1;E;Europe/Amsterdam
7690;Rocky Mount Wilson Regional Airport;Rocky Mount;United States;RWI;KRWI;35.8563;-77.8919;159;-5;A;America/New_York
7691;Whittier Airport;Whittier;United States;;PAWR;60.7772125;-148.7215775;30;-9;U;America/Anchorage
7692;Soldotna Airport;Soldotna;United States;SXQ;\N;60.4749583;-151.0382389;113;-9;U;America/Anchorage
7693;Gillespie;El Cajon;United States;SEE;KSEE;32.8262222;-116.9724444;388;-8;A;America/Los_Angeles
7694;San Clemente Island Nalf;San Clemente Island;United States;;KNUC;33.022744;-118.588489;184;-8;A;America/Los_Angeles
7695;Cotopaxi International Airport;Latacunga;Ecuador;LTX;\N;-0.5425;-78.3657;9205;-5;S;America/Guayaquil
7696;London St Pancras;London;United Kingdom;STP;\N;51.53;-0.125;0;0;E;Europe/London
7697;Amsterdam Centraal;Amsterdam;Netherlands;ZYA;\N;52.378333;4.9;0;1;E;Europe/Amsterdam
7698;Mammy Yoko Heliport;Freetown;Sierra Leone;JMY;\N;8.490278;-13.289722;0;0;N;Africa/Freetown
7699;Shearwater Heliport;Victoria Falls;Zimbabwe;;\N;-17.911443;25.828167;3937;2;U;Africa/Harare
7700;Phan Rang Airport;Phan Rang;Vietnam;PHA;VVPR;11.6335;108.952;101;7;N;Asia/Saigon
7701;Na-San Airport;Son-La;Vietnam;SQH;VVNS;21.217;104.033;2133;7;N;Asia/Saigon
7702;Truckee-Tahoe Airport;Truckee;United States;TKF;KTRK;39.3200422;-120.1395628;5900;-8;A;America/Los_Angeles
7703;Frejus Saint Raphael;Frejus;France;FRJ;LFTU;43.416667;6.733333;7;1;E;Europe/Paris
7704;Geelong Airport;Geelong;Australia;GEX;YGLG;-38.225;144.333;43;10;O;Australia/Hobart
7705;Detroit Amtrak Station;Detroit;United States;;\N;42.368097;-83.072397;630;-5;A;America/New_York
7706;Ventura Amtrak;Ventura;United States;;\N;34.27833;-119.29222;59;-8;A;America/Los_Angeles
7707;Berlin Hauptbahnhof;Berlin;Germany;QPP;\N;52.52493;13.36963;110;1;E;Europe/Berlin
7708;Amsterdam Centraal;Amsterdam;Netherlands;ZYA;\N;52.7873;4.90074;0;1;E;Europe/Amsterdam
8169;Moultrie Municipal Airport;Moultrie;United States;MGR;\N;31.0849167;-83.80325;294;-5;A;America/New_York
7710;Mezen;Mezen;Russia;;ULAE;65.878333;44.215;46;4;N;Europe/Moscow
7711;Vaskovo;Arkhangelsk;Russia;;ULAH;64.441667;40.421667;8;4;N;Europe/Moscow
7712;Cobb County Airport-Mc Collum Field;Atlanta;United States;RYY;KRYY;34.0131569;-84.5970556;1041;-5;A;America/New_York
7713;Oneonta Municipal Airport;Oneonta;United States;ONH;\N;42.524722;-75.064444;1763;-5;A;America/New_York
7714;Tulln ;Tulln;Austria;;LOXT;48.318961;16.114275;575;1;U;Europe/Vienna
7715;Wideawake Field;Georgetown Acension Island Santa Helena;United Kingdom;ASI;\N;-7.969597;-14.393664;278;0;U;Atlantic/St_Helena
7716;Dell Flight Strip;Dell;United States;4U9;K4U9;44.7357483;-112.7200133;6007;-7;A;America/Denver
7717;Mission Field Airport;Livingston-Montana;United States;LVM;KLVM;45.6993889;-110.4483056;4660;-7;A;America/Denver
7718;Kota Kinabalu Airport;Kota Kinabalu;Malaysia;ZWR;\N;11.1111;11.1111;1;1;N;Africa/Lagos
7719;Valetta Waterfront;Valetta;Malta;;\N;35.889867;14.508584;0;1;E;Europe/Malta
7720;Big Timber Airport;Big Timber;United States;6S0;K6S0;45.8063889;-109.9811111;4492;-7;A;America/Denver
7721;Tulip City Airport;Holland;United States;BIV;KBIV;42.7427778;-86.1078333;698;-5;A;America/New_York
7722;London Heliport;London;United Kingdom;;EGLW;51.47;-0.177833;18;0;E;Europe/London
7723;San Nicolo Airport;Venice;Italy;;LIPV;45.4283;12.3881;13;1;E;Europe/Rome
7724;Tallinn Linnahall Heliport;Tallinn;Estonia;;EECL;59.4486;24.7532;23;2;E;Europe/Tallinn
7725;Hernesaari Heliport;Helsinki;Finland;HEN;EFHE;60.1478;24.9244;7;2;E;Europe/Helsinki
7726;Linkenheim Airport;Linkenheim;Germany;;EDRI;49.1417;8.39472;325;1;E;Europe/Berlin
7727;Monument Valley Airport;Monument Valley;United States;;UT25;37.0167;-110.201;5192;-7;U;America/Denver
7728;Hilversum Airport;Hilversum;Netherlands;;EHHV;52.1919;5.14694;3;1;E;Europe/Amsterdam
7729;West 30th St. Heliport;New York;United States;JRA;KJRA;40.7545;-74.0071;7;-5;A;America/New_York
7730;Texel Airport;Texel;Netherlands;;EHTX;53.1153;4.83361;2;1;E;Europe/Amsterdam
7731;La Cerdanya;Das i Fontanals de Cerdanya;Spain;;LECD;42.3864;1.8667;3586;1;E;Europe/Madrid
7732;Lakeland Linder Regional Airport;Lakeland;United States;LAL;KLAL;27.9889167;-82.0185556;142;-5;A;America/New_York
7733;Valetta Grand Harbour;Valetta;Malta;;\N;14.507392;35.888695;0;3;E;Africa/Khartoum
7734;Mgarr Seaplane Base;Mgarr;Malta;;\N;36.0245;14.298359;0;1;E;Europe/Malta
7735;Bangkok;Bangkok;Thailand;;\N;13.9126;100.607;9;7;U;Asia/Bangkok
7737;Jebel Ali Seaplane Base;Jebel Ali Golf Resort;United Arab Emirates;;\N;24.988951;55.023624;0;4;U;Asia/Dubai
7736;Soneva Fushi;Baa Atoll;Maldives;;\N;5.113033;73.080288;0;5;U;Indian/Maldives
7738;Stary Oskol;Stary Oskol;Russia;;UUOS;51.33;37.76833;791;4;N;Europe/Moscow
7739;Savona Cruise Terminal;Savona;Italy;;\N;44.309654;8.486305;0;1;E;Europe/Rome
7740;Barcelona Cruise Terminal;Barcelona;Spain;;\N;41.358901;2.178447;0;1;E;Europe/Madrid
7741;Casablanca Harbor;Casablanca;Morocco;;\N;33.604737;-7.611337;0;0;U;Africa/Casablanca
7742;Lanzarote Arrecife Cruise Terminal;Arrecife Lanzarote;Spain;;\N;28.967298;-13.527528;0;0;E;Atlantic/Canary
7743;Tenerife Cruise Terminal;Santa Cruz de Tenerife;Spain;;\N;28.470068;-16.242471;0;0;E;Atlantic/Canary
7744;Funchal Cruise Terminal;Funchal Madeira;Portugal;;\N;32.641868;-16.911478;0;0;E;Europe/Lisbon
7745;Malaga Cruise Terminal;Malaga;Spain;;\N;36.702989;-4.41395;0;1;E;Europe/Madrid
7746;Ponta Delgada Cruise Terminal;Ponta Delgada Acores;Portugal;;\N;37.738409;-25.662539;0;-1;U;Atlantic/Azores
7747;Vigo Cruise Terminal;Vigo;Spain;;\N;42.241537;-8.728799;0;1;E;Europe/Madrid
7748;Fort Lauderdale Cruise Terminal;Fort Lauderdale;United States;;\N;26.0882;-80.115373;0;-5;A;America/New_York
7749;Southampton Cruise Terminal;Southampton;United Kingdom;;\N;50.900976;-1.413975;0;0;E;Europe/London
7750;Miami Cruise Terminal;Miami;United States;;\N;25.779701;-80.177279;0;-5;A;America/New_York
7751;Nassau Cruise Terminal;Nassau;Bahamas;;\N;25.081158;-77.341207;0;-5;A;America/Nassau
7752;Lisbon Cruise Terminal;Lisbon;Portugal;;\N;38.712606;-9.122483;0;0;E;Europe/Lisbon
7753;Cadiz Cruise Terminal;Cadiz;Spain;;\N;36.534821;-6.290649;0;1;E;Europe/Madrid
7754;Marseille Cruise Terminal;Marseille;France;;\N;43.343969;5.333025;0;1;E;Europe/Paris
7755;Los Angeles San Pedro Cruise Terminal;Los Angeles;United States;;\N;33.747198;-118.276856;0;-8;A;America/Los_Angeles
7756;Kailua Kona Harbor;Kailua Kona Hawaii;United States;;\N;19.635393;-155.998328;0;-10;A;Pacific/Honolulu
7757;Kauai Cruise Terminal Nawiliwili;Nawiliwili Kauai Hawaii;United States;;\N;21.954549;-159.355981;0;-10;A;Pacific/Honolulu
7758;Hilo Cruise Terminal Hawaii;Hilo Hawaii;United States;;\N;19.730967;-155.054094;0;-10;A;Pacific/Honolulu
7759;Honolulu Cruise Terminal Oahu;Honolulu Oahu;United States;;\N;21.301185;-157.865671;0;-10;A;Pacific/Honolulu
7760;Lahaina Harbor Maui Hawaii;Lahaina Maui Hawaii;United States;;\N;20.876456;-156.683235;0;-10;A;Pacific/Honolulu
7761;Ensenada Cruise Terminal;Ensenada;Mexico;;\N;31.855945;-116.624186;0;-8;A;America/Tijuana
7762;Nanaimo Train Station;Nanaimo Station British Columbia;Canada;;\N;49.163978;-123.942502;92;-8;A;America/Vancouver
7763;Victoria Rail Station British Columbia;Victoria Station British Columbia;Canada;;\N;48.428339;-123.370618;75;-8;A;America/Vancouver
7764;Vancouver Cruise Terminal;Vancouver BC;Canada;;\N;49.289012;-123.111463;0;-8;A;America/Vancouver
7765;Seattle Cruise Terminal;Seattle WA;United States;;\N;47.611241;-122.350026;0;-8;A;America/Los_Angeles
7766;Syangboche;Syangboche;Nepal;SYH;VNSB;27.811187;86.712661;12309;5.75;N;Asia/Katmandu
7767;Idlewild Intl;New York;United States;IDL;KIDL;40.639751;-73.778924;13;-5;A;America/New_York
7768;Cheremshanka;Krasnoyarsk;Russia;;UNKM;56.179;92.54265;1000;8;N;Asia/Krasnoyarsk
7769;French Valley Airport;Murrieta-Temecula;United States;RBK;KF70;33.5741791;-117.1284732;1350;-8;A;America/Los_Angeles
7770;Anchorage Main Station;Anchorage Alaska;United States;;\N;61.221749;-149.890852;102;-9;A;America/Anchorage
7771;Seward Train Station;Seward Alaska;United States;;\N;60.121586;-149.439983;22;-9;A;America/Anchorage
7772;Fenosu;Oristano;Italy;FNU;LIER;39.895;8.6383;36;1;E;Europe/Rome
7773;White Waltham Airfield;Maidenhead;United Kingdom;;EGLM;51.501;-0.774;133;0;E;Europe/London
7774;Mysore Airport;Mysore;India;MYQ;VOMY;12.3072;76.6497;2349;5.5;U;Asia/Calcutta
7775;Erie-Ottawa Regional Airport;Port Clinton;United States;PCW;KPCW;41.5162703;-82.8694868;590;-5;U;America/New_York
7776;Dayton-Wright Brothers Airport;Dayton;United States;MGY;KMGY;39.5889722;-84.2248611;957;-5;U;America/New_York
7777;Richmond Municipal Airport;Richmond;United States;RID;KRID;39.7561006;-84.8427175;1140;-5;U;America/New_York
7778;Findlay Airport;Findley;United States;FDY;KFDY;41.0120278;-83.6686111;819;-5;U;America/New_York
7779;Niagara Falls Station;Niagara Falls;Canada;;NIAG;43.108802;-79.06361;0;-5;A;America/New_York
7860;Deputatsky;Deputatsky;Russia;;UEVD;69.293;139.5357;951;11;N;Asia/Vladivostok
7780;Burlington Airpark;Burlington;Canada;;CZBA;43.4425;-79.850833;602;-5;A;America/Toronto
7781;Baganara Island;Baganara;Guyana;;\N;6.343791;-58.591204;38;-4;N;America/Guyana
7782;Penneshaw Airport;Penneshaw;Australia;PEA;YPSH;-34.75;137.933;0;9.5;O;Australia/Adelaide
7783;Matamanoa Helipad;Matamanoa;Fiji;;\N;-17.638056;177.066944;10;12;U;Pacific/Fiji
7784;Kaufbeuren BF;Kaufbeuren;Germany;KFX;KAUF;47.885;10.6294;1200;1;E;Europe/Berlin
7785;Munich HBF;Munich;Germany;MUQ;MUNI;48.1408;11.555;1200;1;E;Europe/Berlin
7786;Nurnberg HBF;Nurnberg;Germany;NUR;NURN;49.446;11.081944;1200;1;E;Europe/Berlin
7787;Ebenhofen BF;Ebenhofen;Germany;EBE;EBEN;47.824;10.623;1200;1;E;Europe/Berlin
7788;Augsburg HBF;Augsburg;Germany;AUB;AUGS;48.3655;10.886;1200;1;E;Europe/Berlin
7789;Biessenhofen BF;Biessenhofen;Germany;BIE;BIES;47.851;10.634;1200;1;E;Europe/Berlin
7790;Buchloe BF;Buchloe;Germany;BUH;BUCH;48.0411;10.715;1200;1;E;Europe/Berlin
7791;Fussen BF;Fussen;Germany;FUX;FUSN;47.585;10.6866;1200;1;E;Europe/Berlin
7792;Kempten HBF;Kempten;Germany;KEX;KEMP;47.724;10.311;1200;1;E;Europe/Berlin
7793;Aiome;Aiome;Papua New Guinea;AIE;\N;-5.8;144.44;350;10;U;Pacific/Port_Moresby
7794;Simbai;Simbai;Papua New Guinea;SIM;\N;-4;144;5450;10;U;Pacific/Port_Moresby
7795;Ambunti;Ambunti;Papua New Guinea;AUJ;\N;-4.15;142.51;160;10;U;Pacific/Port_Moresby
7796;Gunns Camp;Okavango;Botswana;;\N;-19.52666;23.14944;0;2;U;Africa/Gaborone
7797;Sossusvlei;Sossusvlei;Namibia;;\N;-24.489444;15.815278;0;1;S;Africa/Windhoek
7798;Marktoberdorf BF;Marktoberdorf;Germany;OAL;MARK;47.78;10.627;1200;1;E;Europe/Berlin
7799;Marktoberdorf Schule;Marktoberdorf;Germany;MOS;MARO;47.777;10.623;1200;1;E;Europe/Berlin
7800;Essen HBF;Essen;Germany;ESX;ESSE;51.451389;7.0138;1000;1;E;Europe/Berlin
7801;Bochum HBF;Bochum;Germany;BOX;BOCH;51.478506;7.2222;1000;1;E;Europe/Berlin
7802;Koln HBF;Koln;Germany;KOX;KOLN;50.9425;6.958056;1000;1;E;Europe/Berlin
7803;Mannheim HBF;Mannheim;Germany;;\N;49.471;8.469858;1000;1;E;Europe/Berlin
7804;Booker;Wycombe;United Kingdom;;EGTB;51.611667;-0.808056;520;0;E;Europe/London
7805;Bembridge;Bembridge;United Kingdom;BBP;EGHJ;50.677778;-1.109444;53;0;E;Europe/London
7806;Kings County Municipal Airport;Waterville;Canada;;CCW3;45.0519;-64.6517;119;-4;A;America/Halifax
7807;Chojna Air Base;Chojna;Poland;;\N;52.9394;14.4217;187;1;E;Europe/Warsaw
7808;Leuterschach BF;Leuterschach;Germany;LES;LEUT;47.75;10.601;1200;1;E;Europe/Berlin
7809;Black Hills Airport-Clyde Ice Field;Spearfish-South Dakota;United States;SPF;KSPF;44.4811407;-103.7860053;3931;-7;A;America/Denver
7810;Knokke-Heist Westkapelle Heliport;Knokke;Belgium;KNO;EBKW;51.3222;3.29306;0;1;U;Europe/Brussels
7811;Redcliffe;Rothwell;Australia;;YRED;-27.206944;153.072778;75;10;U;Australia/Brisbane
7812;Gdynia;Gdynia;Poland;QYD;EPOK;54.5797;18.5172;144;1;E;Europe/Warsaw
7813;Malbork;Malbork;Poland;;EPMB;54.026944;19.134167;16;1;E;Europe/Warsaw
7814;Lask;Lask;Poland;;EPLK;51.551667;19.179058;633;1;E;Europe/Warsaw
7815;Miroslawiec;Miroslawiec;Poland;;EPMI;53.395072;16.082814;459;1;E;Europe/Warsaw
7816;Krzesiny;Poznan;Poland;;EPKS;52.331719;16.966428;265;1;E;Europe/Warsaw
7817;Olive Branch Muni;Olive Branch;United States;OLV;KOLV;34.876944;-89.783333;350;-6;A;America/Chicago
7818;Vina Del Mar;Vina del Mar;Chile;;SCVM;-32.949611;-71.478583;461;-4;S;America/Santiago
7819;Bangkok Intl closed;Bangkok;Thailand;;BKKX;13.91111;100.60611;5;7;U;Asia/Bangkok
7820;Brampton Airport;Brampton;Canada;;CNC3;43.7603;-79.875;935;-5;A;America/Toronto
7821;Zonguldak;Zonguldak;Turkey;ONQ;LTAS;41.506111;32.088611;44;2;E;Europe/Istanbul
7822;Rocky Mountain Metropolitan Airport;Broomfield-CO;United States;BJC;KBJC;39.90888888;-105.11722222;5670;-7;A;America/Denver
7823;Minot Train Station;Minot;United States;;\N;48.2361;-101.2987;0;-6;A;America/Chicago
7824;Havre Train Station;Havre;United States;;\N;48.55457;-109.67836;0;-7;A;America/Denver
7825;Wishram Train Station;Wishram;United States;;\N;45.6576;-120.9664;0;-8;A;America/Los_Angeles
7826;McNary Field;Salem;United States;SLE;KSLE;44.9095;-123.003;214;-8;A;America/Los_Angeles
7827;Tunica Municipal Airport;Tunica;United States;UTM;KUTA;34.681;-90.3467;194;-6;A;America/Chicago
7828;Batken Airport;Batken;Kyrgyzstan;;UA30;40.0427;70.8381;3474;6;E;Asia/Bishkek
7829;Sun Island;Sun Island;Maldives;;\N;4.2868;73.554894;1;5;U;Indian/Maldives
7830;Kasaba Bay Airport;Kasaba Bay;Zambia;ZKB;FLKY;-8.525;30.663;2780;2;U;Africa/Lusaka
7831;Lindau HBF;Lindau;Germany;LND;LIND;47.5489;9.688;1000;1;E;Europe/Berlin
7832;Guardiamarina Zanartu Airport;Puerto Williams;Chile;WPU;\N;-54.9311;-67.6263;88;-4;U;America/Santiago
7833;Volkel AB;Volkel;Netherlands;UDE;\N;51.656389;5.708611;72;1;E;Europe/Amsterdam
7834;Hoogeveen Airport;Hoogeveen;Netherlands;;EHHO;52.7308;6.51611;40;1;E;Europe/Amsterdam
7835;Teuge Airport;Deventer;Netherlands;;EHTE;52.2447;6.04667;17;1;E;Europe/Amsterdam
7836;Midden-Zeeland Airport;Middelburg;Netherlands;;EHMZ;51.5122;3.73111;6;1;E;Europe/Amsterdam
7837;Ameland Airport;Ameland;Netherlands;;EHAL;53.4517;5.67722;11;1;E;Europe/Amsterdam
7838;Saint-Cyr-l-Ecole Airport;Saint-Cyr;France;;LFPZ;48.810278;2.073332;373;1;E;Europe/Paris
7839;Lawrence J Timmerman Airport;Milwaukee;United States;MWC;KMWC;43.1103889;-88.0344167;745;-6;A;America/Chicago
7840;Southern Wisconsin Regional Airport;Janesville;United States;JVL;KJVL;42.62025;-89.0415556;808;-6;A;America/Chicago
7841;Mantsonyane Airport;Mantsonyane;Lesotho;;FXMN;-29.5461;28.271;7100;2;U;Africa/Maseru
7842;Hatfield;Hatfield;United Kingdom;HTF;\N;51.765;-0.24833;254;0;U;Europe/London
7843;Burswood Park Helipad;Perth;Australia;;\N;-31.952;115.859;25;8;O;Australia/Perth
7844;Toronto Coach Terminal;Toronto;Canada;;TRTO;43.655996;-79.38416;0;-5;A;America/Toronto
7845;Montreal Central Bus Station;Montreal;Canada;;MTRL;45.51527887;-73.561427593;0;-5;A;America/Toronto
7846;Arlington Municipal;Arlington;United States;GKY;KGKY;32.6638611;-97.0942778;628;-6;A;America/Chicago
7847;Gwinnett County Airport-Briscoe Field;Lawrenceville;United States;LZU;KLZU;33.9780761;-83.9623772;1061;-5;A;America/New_York
7848;Bowling Green-Warren County Regional Airport;Bowling Green;United States;BWG;KBWG;36.9645278;-86.4196667;547;-6;A;America/Chicago
7849;Richard Lloyd Jones Jr Airport;Tulsa;United States;RVS;KRVS;36.0396111;-95.9846389;638;-6;A;America/Chicago
7850;Emeryville Amtrak Station;Emeryville;United States;;\N;37.8405149;-122.29134;23;-8;A;America/Los_Angeles
7851;Bakersfield Amtrak Station;Bakersfield;United States;;\N;35.372159;-119.008393;408;-8;A;America/Los_Angeles
7852;Krymsk;Novorossiysk;Russia;NOI;\N;44.4016;37.7779;50;4;N;\N
7853;Minhad HB;Minhad AB;United Arab Emirates;NHD;OMDM;25.02694;55.36611;165;4;U;Asia/Dubai
7854;Kirovograd;Kirovograd;Ukraine;KGO;UKKG;48.54;32.29;510;2;E;Europe/Kiev
7855;Roitzschjora Airport;Roitzschjora;Germany;;EDAW;51.5778;12.4944;289;1;U;Europe/Berlin
7856;Alalamain Intl.;Dabaa City;Egypt;DBB;HEAL;30.15;28.0833;144;2;N;Africa/Cairo
7857;Bryce Canyon;Bryce Canyon;United States;BCE;KBCE;37.706444;-112.145806;7590;-7;A;America/Denver
7858;Heidelberg;Heidelberg;Germany;HDB;EDIU;49.393333;8.6525;0;1;E;Europe/Berlin
7859;Burlington-Alamance Regional Airport;Burlington;United States;BUY;KBUY;36.0485433;-79.4748892;617;-5;A;America/New_York
7861;Chkalovsky Airport;Shchyolkovo;Russia;CKL;UUMU;55.878333;38.061667;152;4;N;Europe/Moscow
7862;Tengchong Tuofeng Airport;Tengchong;China;TCZ;ZUTC;24.938651;98.483591;5500;8;N;Asia/Chongqing
7863;Belbek Sevastopol International Airport;Sevastopol;Ukraine;UKS;UKFB;44.691431;33.57567;344;2;U;Europe/Kiev
8420;Berlin Brandenburg Willy Brandt;Berlin;Germany;BER;EDDB;52.366667;13.503333;154;1;E;Europe/Berlin
7865;Paradise Island Seaplane Base;Nassau;Bahamas;WZY;\N;25.0872;-77.3239;0;-5;U;America/Nassau
7866;Selous Siswandu;Selous National Park;Tanzania;;\N;-7.70515;38.158644;0;3;U;Africa/Dar_es_Salaam
7867;De Peel Air Base;Deurne;Netherlands;;EHDP;51.5173;5.85572;98;1;E;Europe/Amsterdam
7868;Camp Bastion ;Camp Bastion;Afghanistan;;OAZI;31.865556;64.195278;2808;4.5;N;Asia/Kabul
7869;New Century AirCenter Airport;Olathe;United States;JCI;KIXD;38.8309167;-94.8903056;1087;-6;A;America/Chicago
7870;Easton-Newnam Field Airport;Easton;United States;ESN;KESN;38.8041667;-76.069;72;-5;A;America/New_York
7871;Stafsberg Airport;Hamar;Norway;HMR;ENHA;60.8181;11.068;713;1;E;Europe/Oslo
7872;Ringebu Airport;Frya;Norway;;ENRI;61.5472;10.0611;571;1;E;Europe/Oslo
7873;Starmoen;Elverum;Norway;;ENSM;60.88;11.6731;659;1;E;Europe/Oslo
7874;Ferry County Airport;Republic;United States;R49;\N;48.7182058;-118.6564714;2522;-8;A;America/Los_Angeles
7875;Yuba County Airport;Yuba City;United States;MYV;KMYV;39.0553;-121.3411;62;-8;A;America/Los_Angeles
7876;Basel SBB;Basel;Switzerland;;\N;47.547531;7.58977;800;1;E;Europe/Zurich
7877;Stockholm Cruise Port;Stockholm;Sweden;STO;\N;59.3233;18.081;0;1;E;Europe/Stockholm
7878;Helsingborg Cruise Port;Helsingborg;Sweden;JHE;\N;56.0419;12.6912;0;1;E;Europe/Stockholm
7879;Phillip Island Cruise Port;Phillip Island;Australia;;YPID;-38.4468;145.2378;0;10;O;Australia/Hobart
7880;Halliburton Field Airport;Duncan;United States;DUC;KDUC;34.4713056;-97.9598611;1114;-6;A;America/Chicago
7881;Port Authority Bus Terminal;New York;United States;;NYPA;40.75616;-73.9906;0;-5;A;America/New_York
7882;Filitheyo;N.Nilandhoo Atoll;Maldives;;\N;3.207801;73.037308;0;5;N;Indian/Maldives
7883;Helsinki Cruise Port;Helsinki;Finland;;HELC;60.163056;24.969167;6;2;E;Europe/Helsinki
7884;Chinle Municipal Airport;Chinle;United States;E91;\N;36.1108806;-109.5754222;5547;-7;N;America/Denver
7885;Garner Field;Uvalde;United States;UVA;KUVA;29.215429;-99.748962;942;-6;A;America/Chicago
7886;Lewis University Airport;Lockport;United States;LOT;KLOT;41.606326;-88.083003;680;-6;A;America/Chicago
7887;Frasca Field;Urbana;United States;C16;\N;40.144979;-88.200197;735;-6;A;America/Chicago
7888;Buchanan Field Airport;Concord;United States;CCR;KCCR;37.9896667;-122.0568889;26;-8;A;America/Los_Angeles
7889;Key Largo;Ocean Reef Club Airport;United States;OCA;07FA;25.325393;-80.274775;8;-5;A;America/New_York
7890;Denver Union Station;Denver;United States;;\N;39.753187;-105.000093;0;-7;A;America/Denver
7891;DELETE;DELETE;United States;;\N;-1.1111;-1.1111;0;0;A;\N
7892;Polygone;Strasbourg Neudorf;France;;LFGC;48.5548;7.7778;130;1;E;Europe/Paris
7893;Nannhausen;Nannhausen;Germany;;EDRN;50;7.591667;1296;1;E;Europe/Berlin
7894;Yushu Batang;Yushu;China;YUS;ZLYS;32.825;97.125;13000;8;U;Asia/Chongqing
7895;Playon Chico;Playon Chico;Panama;PYC;\N;9.303333;-78.236111;10;-5;U;America/Panama
7896;Ustupo;Ustupo;Panama;UTU;\N;9.137778;-77.933611;10;-5;U;America/Panama
7897;Mamitupo;Mamitupo;Panama;MPU;\N;9.186667;-77.984167;10;-5;U;America/Panama
7898;Huai An Lianshui Airport;Huai An;China;HIA;ZSSH;33.7772;119.1478;23;8;N;Asia/Chongqing
7899;Mureck;Mureck;Austria;;\N;46.70778;15.76941;764;1;E;Europe/Vienna
7900;Guessing;Guessing;Austria;;\N;47.05932;16.32449;820;1;E;Europe/Vienna
7901;El Porvenir;El Porvenir;Panama;PVE;\N;9.559167;-78.971111;5;-5;U;America/Panama
7902;Oshawa Airport;Oshawa;Canada;YOO;CYOO;43.9228;-78.895;459;-5;A;America/Toronto
7903;Marl-Loemuehle Airport;Recklinghausen;Germany;;EDLM;51.6472;7.16333;240;1;U;Europe/Berlin
7904;Farila Air Base;Farila;Sweden;;ESNF;62.03;15.752;660;1;E;Europe/Stockholm
7905;Lahr Airport;Lahr;Germany;LHA;EDTL;48.3693;7.82772;511;1;E;Europe/Berlin
7906;Monywa Airport;Monywa;Burma;;VYMY;22.233;95.117;298;6.5;N;Asia/Rangoon
7907;Ohio University Airport;Athens;United States;;KUNI;39.2118928;-82.2292554;766;-5;U;America/New_York
7908;Springfield-Beckly Municipal Airport;Springfield;United States;SGH;KSGH;39.8402778;-83.8401667;1051;-5;A;America/New_York
7909;Sun Island Airport;South Aari Atoll;Maldives;MSI;SIAM;3.295;72.4814;2;5;N;Indian/Maldives
7910;Fes Sefrou Airport;Fes;Morocco;;GMFU;34.0081;-4.96556;1539;0;U;Africa/Casablanca
7911;Herrera International Airport;Santo Domingo;Dominican Republic;HEX;MDHE;18.475;-69.975;22;-4;U;America/Santo_Domingo
7912;Cooinda;Cooinda;Australia;CDA;YCOO;-12.9033;132.532;43;9.5;U;Australia/Darwin
7913;Jabiru;Jabiru;Australia;JAB;YJAB;-12.6571;132.893;52;9.5;U;Australia/Darwin
7914;Pitt Island Airport;Pitt Island;New Zealand;;PITT;-44.300278;-176.220556;12;12.75;U;Pacific/Chatham
7915;Plattling Bahnhof;Plattling;Germany;;\N;48.780147;12.863573;1050;1;E;Europe/Berlin
7916;Osterhofen Bahnhof;Osterhofen Niederbayern;Germany;;\N;48.7;13.016667;1043;1;E;Europe/Berlin
7917;Passau Hbf;Passau;Germany;;\N;48.574167;13.450833;1000;1;E;Europe/Berlin
7918;Regensburg-Oberhub Airport;Regensburg;Germany;;EDNR;49.1419;12.0819;1299;1;E;Europe/Berlin
7919;Regensburg HBF;Regensburg;Germany;RGB;REGE;49.022;12.1111;1200;1;E;Europe/Berlin
7920;Treuchtlingen BF;Treuchtlingen;Germany;TLG;TREU;49.04;11.081944;1200;1;E;Europe/Berlin
7921;Rivera Intl.;Rivera;Uruguay;;\N;-31.1127;55.4611;200;-3;U;\N
7923;Miami Seaplane Base;Miami;United States;MPB;\N;25.7783;-80.1703;0;-5;A;America/New_York
7924;Hastings Airport;Freetown;Sierra Leone;HGS;GFHA;8.39713;-13.1291;60;0;N;Africa/Freetown
7925;Philip Billard Muni;Topeka;United States;TOP;KTOP;39.068657;-95.622482;881;-6;A;America/Chicago
7926;Grumeti Airstrip;Serengeti National Park;Tanzania;;\N;-2.157488;34.221232;5080;3;N;Africa/Dar_es_Salaam
7927;Emporia Municipal Airport;Emporia;United States;EMP;\N;38.3321;-96.1912;1208;-6;A;America/Chicago
7928;Benson Airstrip;Uvalde;United States;;2XS8;29.229405;-99.823947;929;-6;A;America/Chicago
7929;Rough River State Park;Null;United States;;K2I3;37.609778;-86.506925;577;-6;A;America/Chicago
7930;Smyrna Airport;Smyrna;United States;;KMQY;36.009;-86.5201;543;-6;A;America/Chicago
7931;Franklin County Airport;Sewanee;United States;;KUOS;35.2051;-85.8981;1953;-6;A;America/Chicago
7932;Gunsa;Shiquanhe;China;NGQ;ZUAL;32.10027;80.052778;13780;8;N;Asia/Chongqing
7933;Magdeburg-Cochstedt;Cochstedt;Germany;CSO;EDBC;51.855833;11.418333;596;1;E;Europe/Berlin
7934;Wurzburg HBF;Wurzburg;Germany;WZB;WURZ;49.7999;9.95555;1000;1;E;Europe/Berlin
7935;Collin County Regional Airport at Mc Kinney;DALLAS;United States;TKI;KTKI;33.1779444;-96.5905278;585;-6;A;America/Chicago
7936;Chicago Executive;Chicago-Wheeling;United States;PWK;KPWK;42.1142897;-87.9015376;647;-6;A;America/Chicago
7937;Ol Kiombo Airstrip;Masai Mara National Park;Kenya;;\N;-1.409595;35.110803;0;3;U;Africa/Nairobi
7938;Kelso Longview;Kelso;United States;KLS;KKLS;46.118;-122.898389;20;-8;A;America/Los_Angeles
7939;Letiste Benesov;Benesov;Czech Republic;;LKBE;49.4427;14.3841;1322;1;U;Europe/Prague
7940;Put-in-Bay Airport;Put-in-Bay;United States;3W2;\N;41.3521;-82.497;595;-5;A;America/New_York
7941;Bougouni;Bougouni;Mali;;GABG;11.45;-7.517;300;0;N;Africa/Bamako
7942;Glenwood Springs Train Station;Glenwood Springs;United States;;\N;39.548;-107.3233;0;-7;A;America/Denver
7943;Grand Junction Train Station;Grand Junction;United States;;\N;39.0646;-108.5705;0;-7;A;America/Denver
7944;Reno Amtrak Station;Reno;United States;;\N;39.5287;-119.8116;0;-8;A;America/Los_Angeles
7945;Sacramento Amtrak Station;Sacramento;United States;;\N;38.584791;-121.500517;0;-8;A;America/Los_Angeles
7946;Tureia Airport;Tureia;French Polynesia;ZTA;NTGY;-20.7897;-138.57;12;-10;N;\N
7947;Ice Runway-McMurdo Sstation;Ross Island;Antarctica;;NZIR;-77.8477778;166.6683333;1;12;N;Antarctica/South_Pole
7948;Keekorok;Keekorok;Kenya;;HKKE;-1.586418;35.259036;5536;3;N;Africa/Nairobi
7949;Ol Kiombo;Mara Intrepids;Kenya;;\N;-1.409569;35.110788;5006;3;N;Africa/Nairobi
7950;Ol Kiombo;Mara Intrepids;Kenya;;\N;-1.409569;35.110788;5006;3;N;Africa/Nairobi
7951;Kichwa Tembo;Kichwa Tembo;Kenya;;\N;-1.264132;35.022719;5174;3;N;Africa/Nairobi
7952;Busan;Busan;South Korea;;\N;35.179444;129.075556;0;9;U;Asia/Seoul
7953;Busan;Busan;South Korea;;\N;35.179444;129.075556;0;9;U;Asia/Seoul
7954;Gare de Marne-la-Vallee;Chessy;France;;\N;48.869722;2.782778;250;1;E;Europe/Paris
7955;Jefferson County Intl;Port Townsend;United States;TWD;\N;48.0314;-122.4838;108;-8;A;America/Los_Angeles
7956;Lynden Airport;Lynden;United States;38W;\N;48.9558961;-122.4581183;106;-8;A;America/Los_Angeles
7957;Jefferson County Intl;Port Townsend;United States;0S9;\N;48.0538086;-122.8106436;108;-8;A;America/Los_Angeles
7958;Savut;Savut;Botswana;;\N;1;1;1;1;U;\N
7959;Xakan;Xakan;Botswana;;\N;1;1;1;1;N;\N
7960;Xakan;Xakan;Botswana;;\N;1;1;1;1;N;\N
7961;Xugan;Xugan;Botswana;;\N;1;1;1;0;N;\N
7962;Puerto Obaldia;Puerto Obaldia;Panama;PUE;MPOA;8.68333;-77.5333;223;-5;U;America/Panama
7963;Kerch Intl;Kerch;Ukraine;KHC;UKFK;45.372869;36.402761;152;2;E;Europe/Kiev
7964;Khao Sok National Park;Surat Thani;Thailand;;\N;8.936667;98.530278;1200;7;U;Asia/Bangkok
7965;Khao Sok National Park;Surat Thani;Thailand;;\N;8.936667;98.530278;1200;7;U;Asia/Bangkok
7966;Boston South Station;Boston;United States;;BOST;42.35;-71.0558;0;-5;A;America/New_York
7967;Tarifa;Tarifa;Spain;;\N;36.070781;-5.602764;20;1;E;Europe/Madrid
7968;Train Station;Ottawa;Canada;XDS;\N;45.4164;-75.6517;374;-5;A;America/Toronto
7969;Train Station;Belleville;Canada;XVV;\N;44.1793;-77.3747;315;-5;A;America/Toronto
7970;Train Station;Edmonton;Canada;XZL;\N;53.5789;-113.5307;2208;-7;A;America/Edmonton
7971;Train Station;Richmond;United States;ZRD;\N;37.5343;-77.42945;26;-5;A;America/New_York
7972;Sentral;Kuala Lumpur;Malaysia;XKL;\N;3.134;101.686;126;8;N;Asia/Kuala_Lumpur
7973;Train Station;Churchill;Canada;XAD;\N;58.76775;-94.17425;10;-6;A;America/Winnipeg
7974;Train Station;Winnipeg;Canada;XEF;\N;49.8889;-97.1342;751;-6;A;America/Winnipeg
7975;Kingston VIA Station;Kingston;Canada;;\N;44.2572;-76.53715;282;-5;A;America/Toronto
7976;Ukunda Airport;Ukunda;Kenya;UKA;HKUK;-4.29694;39.5714;0;3;U;Africa/Nairobi
7977;Whitehaven Beach;Whitsunday Island;Australia;;\N;-20.28866;149.044966;0;10;U;Australia/Brisbane
7978;Wilmington Airborne Airpark;Wilmington;United States;ILN;KILN;39.42792;-83.792118;1077;-5;U;America/New_York
7979;Marana Regional;Tucson;United States;AVW;KAVQ;32.409556;-111.218388;2031;-7;U;America/Phoenix
7980;Casa Grande Municipal Airport;Casa Grande;United States;CGZ;KCGZ;32.954889;-111.766832;1464;-7;U;America/Phoenix
7981;Mobile;Mobile;United States;;1AZ0;33.111944;-112.269166;1261;-7;U;America/Phoenix
7982;Buckeye Municipal Airport;Buckeye;United States;BXK;KBXK;33.420417;-112.68618;1033;-7;U;America/Phoenix
7983;Gila Bend Municipal Airport;Gila Bend;United States;E63;KE63;32.960169;-112.673636;789;-7;U;America/Phoenix
7984;McMinn Co;Athens;United States;MMI;KMMI;35.39919;-84.56177;874;-5;N;America/New_York
7985;Sterling Municipal Airport;Sterling;United States;STK;KSTK;40.6153136;-103.2648454;4040;-7;A;America/Denver
7986;Rawlins Municipal Airport-Harvey Field;Rawlins;United States;RWL;KFWL;41.8055975;-107.19994;6813;-7;A;America/Denver
7987;Mackenzie Airport;Mackenzie British Columbia;Canada;YZY;CYZY;55.29944;-123.08333;2264;-8;A;America/Vancouver
7988;JALOU;Jalu;Libya;;\N;29.142222;21.380556;0;2;U;Africa/Tripoli
7989;JALOU;Jalu;Libya;;\N;29.142222;21.380556;0;2;U;Africa/Tripoli
7990;Caldwell Essex County Airport;Caldwell;United States;CDW;KCDW;40.8752222;-74.2813611;172;-5;A;America/New_York
7991;Lee C Fine Memorial Airport;Kaiser Lake Ozark;United States;AIZ;KAIZ;38.096035;-92.5494875;869;-6;A;America/Chicago
7992;Yuryevets;Yuryevets;Russia;;\N;57.3106;43.1003;269;4;N;Europe/Moscow
7993;Big Bear City;Big Bear;United States;L35;\N;34.2637778;-116.8560278;6725;-8;A;America/Los_Angeles
7994;KIEV  INTERNATIONAL AIRPORT;KIEV;Ukraine;KIP;KIEV;50.1403;30.1808;250;2;E;Europe/Kiev
7995;Los Angeles Union Station;Los Angeles;United States;;\N;34.056111;-118.234167;300;-8;A;America/Los_Angeles
7996;Bamberg BF;Bamberg;Germany;BAM;BAMB;49.911;10.9;800;1;E;Europe/Berlin
7997;Ingolstadt BF;Ingolstadt;Germany;IGS;INGS;48.7777;11.422;888;1;E;Europe/Berlin
7998;Thomasville Regional Airport;Thomasville;United States;TVI;KTVI;30.9017921;-83.8811285;264;-5;A;America/New_York
7999;Henderson Executive Airport;Henderson;United States;HSH;KHND;35.972778;-115.134444;1881;-8;A;America/Los_Angeles
8000;Gostomel Antonov;Kiev;Ukraine;GML;UKKM;50.603611;30.191944;517;2;E;Europe/Kiev
8001;Zuoying Station;Kaohsiung;Taiwan;;\N;22.687389;120.307481;31;8;N;Asia/Taipei
8002;Taipei Station;Taipei;Taiwan;;\N;25.047778;121.517222;30;8;N;Asia/Taipei
8003;Donghae;Donghae;South Korea;;\N;37.489;129.124;0;9;N;Asia/Seoul
8004;Sakaiminato Port;Sakaiminato;Japan;;\N;35.539;133.264;0;9;N;Asia/Tokyo
8005;Henry Tift Myers Airport;Tifton;United States;TMA;KTMA;31.4289814;-83.488545;355;-5;A;America/New_York
8006;Landshut Ellermuehle;Landshut;Germany;;EDML;48.5133333333333;12.035;1200;1;E;Europe/Berlin
8007;Itzehoe-Hungriger Wolf Airport;Itzehoe;Germany;;EDHF;53.9944;9.57861;89;1;E;Europe/Berlin
8008;Savannah Cruise Terminal;Savannah;United States;;\N;32.08558;-81.097559;0;-5;U;America/New_York
8009;CHARLESTON CRUISE TERMINAL;CHARLESTON;United States;;\N;32.781081;-79.92364;0;-5;U;America/New_York
8010;NEW YORK CRUISE TERMINAL PIER 92;NEW YORK;United States;;\N;40.767681;-73.999107;0;-5;U;America/New_York
8011;Navy Dockyard Hamilton Bermuda;HAMILTON;Bermuda;;\N;32.327037;-64.830626;0;-4;U;Atlantic/Bermuda
8012;Tokyo Station;Tokyo;Japan;;\N;35.672498;139.753218;16;9;N;Asia/Tokyo
8013;Shin-Osaka Station;Osaka;Japan;;\N;34.738932;135.500093;61;9;N;Asia/Tokyo
8014;Hiroshima Station;Hiroshima;Japan;;\N;34.3973853173792;132.475978195923;28;9;N;Asia/Tokyo
8015;Frankfurt-Main Hauptbahnhof;Frankfurt;Germany;ZRB;\N;50.1070257990375;8.66276050515751;372;1;E;Europe/Berlin
8016;Amsterdam Amstel Station;Amsterdam;Netherlands;;\N;52.3466365904832;4.91781401731873;21;1;E;Europe/Amsterdam
8017;Florenc Central Bus Station;Prague;Czech Republic;;\N;50.0909636578274;14.4394929551697;617;1;E;Europe/Prague
8018;Wien Sudbahnhof;Vienna;Austria;;\N;48.1866480486625;16.3815008466187;651;1;E;Europe/Vienna
8019;Praha hlavni nadrazi;Prague;Czech Republic;XYG;\N;50.0826892098189;14.4350297593689;691;1;E;Europe/Prague
8020;Wien Westbahnhof;Vienna;Austria;;\N;48.1965754250116;16.3374911611023;675;1;E;Europe/Vienna
8021;Venezia Santa Lucia;Venice;Italy;;\N;45.4417738709625;12.3199978080597;8;1;E;Europe/Rome
8022;Firenze Santa Maria Novella;Florence;Italy;;\N;43.7766651171349;11.2480175237426;159;1;E;Europe/Rome
8023;Roma Termini;Rome;Italy;;\N;41.9004811137748;12.5020006853333;189;1;E;Europe/Rome
8024;St Pancras Railway Station;London;United Kingdom;QQS;\N;51.532519492138;-0.12630037301642;80;0;E;Europe/London
8025;Chitabe Airstrip;Chitabe;Botswana;;\N;-19.4687;23.38271;3000;2;U;Africa/Gaborone
8026;Fish River Canyon Lodge Airstrip;Fish River Canyon;Namibia;;\N;-27.659288;17.837624;3077;1;S;Africa/Windhoek
8027;Sossusvlei Desert Lodge Airstrip;Sossusvlei;Namibia;;\N;-24.802812;15.891713;2950;1;S;Africa/Windhoek
8028;Ulusaba Airstrip;Ulusaba;Namibia;ULX;\N;-24.782766;31.353929;1470;2;S;Africa/Johannesburg
8029;RADOM;RADOM;Poland;QXR;EPRA;51.231978;21.124183;479;1;E;Europe/Warsaw
8030;Deer Valley Municipal Airport;Phoenix ;United States;DVT;KDVT;33.4117;112.457;1478;8;A;Asia/Chongqing
8031;Calgary Springbank Airport;Calgary;Canada;YBW;CYBW;51.1031;-114.374;3939;-7;A;America/Edmonton
8032;Golden Airport;Golden;Canada;YGE;CYGE;51.2992;-116.982;2575;-7;A;America/Edmonton
8033;Revelstoke Airport;Revelstoke;Canada;YRV;CYRV;50.9667;-118.183;1459;-8;A;America/Vancouver
8034;Republic Airport;Farmingdale;United States;;KFRG;40.7288;-73.4134;82;-5;A;America/New_York
8035;Allstedt;Allstedt;Germany;;EDBT;51.3806;11.4467;932;1;E;Europe/Berlin
8036;BWKAM;Amberg;Germany;;\N;49.43;11.8;500;1;E;Europe/Berlin
8037;Glindbruchkippe;Peine;Germany;;KIPP;52.3237;10.182;230;1;E;Europe/Berlin
8038;General Freire;Curico;Chile;;SCIC;-34.9667;-71.2164;722;-4;S;America/Santiago
8039;Peine-Eddesse Airport;Peine;Germany;;EDVP;52.4025;10.2289;249;1;E;Europe/Berlin
8040;Seattle King Street Station;Seattle;United States;;\N;47.5985;-122.3299;7;-8;A;America/Los_Angeles
8041;Tula;Tula;Russia;TYA;\N;54.142208;37.355619;590;4;N;Europe/Moscow
8042;Hondo Municipal Airport;Hondo;United States;HDO;KHDO;29.3591;-99.1775;930;-6;A;America/Chicago
8043;Zhongwei Xiangshan Airport;Zhongwei;China;ZHY;ZLZW;37.5728;105.1544;4305;8;N;Asia/Chongqing
8044;Milwaukee Airport Station;Milwaukee;United States;;\N;42.940556;-87.924722;730;-6;A;America/Chicago
8045;Springfield IL Amtrak;Springfield;United States;;\N;39.8023;-89.6515;593;-6;A;America/Chicago
8046;St. Louis Amtrak - old;St. Louis;United States;;\N;38.6241;-90.2056;450;-6;A;America/Chicago
8047;Camarillo Amtrak;Camarillo;United States;;\N;34.2164;-119.0335;140;-8;A;America/Los_Angeles
8048;Seattle Pier 52;Seattle;United States;;\N;47.6005;-122.3388;7;-8;A;America/Los_Angeles
8049;Bremerton Terminal;Bremerton;United States;;\N;47.5619;-122.625;7;-8;A;America/Los_Angeles
8050;McKinley National Park Airport;McKinley Park;United States;MCL;PAIN;63.732757;-148.91129;1720;-9;A;America/Anchorage
8051;Lake Hood Seaplane Base;Anchorage;United States;LHD;PALH;61.1866382;-149.9653918;71;-9;A;America/Anchorage
8052;Prospect Creek Airport;Prospect Creek;United States;PPC;PAPR;66.814167;-150.643611;1095;-9;A;America/Anchorage
8053;Khwai River Lodge;Khwai River;Botswana;KHW;FBKR;-19.149166;23.7875;3000;2;N;Africa/Gaborone
8054;Spremberg-Welzow Airport;Welzow;Germany;;EDCY;51.5756;14.1369;374;1;E;Europe/Berlin
8055;Taichung Airport;Taichung;Taiwan;TXG;RCLG;24.1863;120.654;369;8;N;Asia/Taipei
8056;Columbia County;Hudson NY;United States;HCC;\N;42.2913;-73.7103;198;-5;A;America/New_York
8057;Sasovo;Sasovo;Russia;;\N;54.3539;41.9739;365;4;N;Europe/Moscow
8058;Hong Kong Macau Ferry Terminal;Hong Kong;Hong Kong;;\N;22.289372;114.152153;0;8;N;Asia/Hong_Kong
8059;Hong Kong Macau Ferry Terminal;Hong Kong;Hong Kong;;\N;22.289372;114.152153;0;8;N;Asia/Hong_Kong
8060;Macau Taipa Ferry Terminal;Macau;Macau;;\N;22.1633;113.57404;0;8;N;Asia/Macau
8061;Macau Taipa Ferry Terminal;Macau;Macau;;\N;22.1633;113.57404;0;8;N;Asia/Macau
8062;Wheeling Ohio County Airport;Wheeling;United States;HLG;\N;40.175;-80.6463;1195;-5;A;America/New_York
8063;Fitzgerald Municipal Airport;Fitzgerald;United States;FZG;KFZG;31.6839046;-83.2709036;365;-5;A;America/New_York
8064;Perry-Foley Airport;Perry;United States;40J;\N;30.0692778;-83.5805833;44;-5;A;America/New_York
8065;Cairo-Grady County Airport;Cairo;United States;70J;\N;30.8879767;-84.1547353;265;-5;A;America/New_York
8066;Pacific Station;Vancouver;Canada;;\N;49.2739;-123.0986;0;-8;A;America/Vancouver
8067;Tamar Mepe Airport;Mestia;Georgia;;\N;43.06848;42.75424;100;4;N;Asia/Tbilisi
8068;Samsun ;Samsun;Turkey;SSX;\N;41.1641;36.1814;521;2;U;Europe/Istanbul
8069;Ye;Ye;Burma;XYE;VYYE;15.25;97.85;25;6.5;U;Asia/Rangoon
8070;Isla San Felix;Isla San Felix;Chile;;SCFX;-26.293864;-80.096214;165;-4;U;America/Santiago
8071;Kingston Train Station;Kingston;Canada;XEG;KGST;44.256944;-76.536943;282;-5;A;America/Toronto
8072;Dorval Railway Station;Dorval;Canada;XAX;\N;45.448611;-73.74111;0;-5;A;America/Toronto
8073;Brockville Megabus Stop;Brockville;Canada;;BRKM;44.6003;-75.7039;0;-5;A;America/Toronto
8074;Brockville VIA Station;Brockville;Canada;;BRKV;44.5919;-75.693343;0;-5;A;America/Toronto
8075;Shaybah airport;Shaybah;Saudi Arabia;;OESB;22.5147;53.9597;262;3;N;Asia/Riyadh
8076;Dubai Al Maktoum;Dubai;United Arab Emirates;DWC;OMDW;24.55056;55.103174;170;4;U;Asia/Dubai
8077;Aransas County Airport;Rockport;United States;RKP;KRKP;28.0862222;-97.0436944;24;-6;A;America/Chicago
8078;Bandanaira Airport;Bandanaira-Naira Island;Indonesia;NDA;\N;-4.53333;129.9;37;9;N;\N
8079;Megeve Airport;Verdun;France;MVV;LFHM;45.8208;6.65222;4823;1;U;Europe/Paris
8080;Meribel Airport;Ajaccio;France;MFX;LFKX;45.4069;6.58056;5636;1;U;Europe/Paris
8081;Oldbury;Oldbury;United Kingdom;;\N;52.55;-2.0166;499;0;E;Europe/London
8082;Tianyang;Baise;China;AEB;\N;23.72;106.96;485;8;N;Asia/Chongqing
8083;Okaukuejo Airport;Okaukuejo;Namibia;OKF;FYOO;-19.1492;15.9119;3911;1;S;Africa/Windhoek
8084;Mokuti Lodge Airport;Mokuti Lodge;Namibia;OKU;FYMO;-18.8128;17.0594;3665;1;S;Africa/Windhoek
8085;Rotenburg Wuemme Airport;Rotenburg Wuemme;Germany;;EDXQ;53.1283;9.34861;98;1;E;Europe/Berlin
8086;Wipperfuerth-Neye Airport;Wipperfuerth;Germany;;EDKN;51.1242;7.37361;863;1;E;Europe/Berlin
8087;Osnabrueck-Atterheide Airport;Osnabrueck;Germany;;EDWO;52.2864;7.96972;285;1;E;Europe/Berlin
8088;Ballenstedt Airport;Ballenstedt;Germany;;EDCB;51.7458;11.2297;535;1;E;Europe/Berlin
8089;Hartenholm Airport;Hasenmoor;Germany;;EDHM;53.915;10.0356;108;1;E;Europe/Berlin
8090;Ganderkesee Atlas Airfield Airport;Ganderkesee;Germany;;EDWQ;53.0361;8.50556;95;1;E;Europe/Berlin
8091;Nienburg-Holzbalge Airport;Nienburg Weser;Germany;;EDXI;52.7097;9.1625;82;1;E;Europe/Berlin
8092;Damme Airport;Damme;Germany;;EDWC;52.4875;8.18556;151;1;E;Europe/Berlin
8093;Borkenberge Airport;Duelmen;Germany;;EDLB;51.78;7.288;157;1;E;Europe/Berlin
8094;Obermehler-Schlotheim Airport;Obermehler;Germany;;EDCO;51.2678;10.6347;909;1;E;Europe/Berlin
8095;Hodenhagen Airport;Hodenhagen;Germany;;EDVH;52.7611;9.60556;79;1;E;Europe/Berlin
8096;Grube Airport;Grube;Germany;;EDHB;54.2444;11.0247;7;1;E;Europe/Berlin
8097;Toender Airport;Toender;Denmark;;EKTD;54.9297;8.84057;0;1;E;Europe/Copenhagen
8098;Celle-Arloh Airport;Celle;Germany;;EDVC;52.6872;10.1114;207;1;E;Europe/Berlin
8099;Uelzen Airport;Uelzen;Germany;;EDVU;52.9839;10.465;246;1;E;Europe/Berlin
8100;Hamm-Lippewiesen Airport;Hamm;Germany;;EDLH;51.6897;7.81611;190;1;E;Europe/Berlin
8101;Luesse Airport;Luesse;Germany;;EDOJ;52.1411;12.6647;217;1;E;Europe/Berlin
8102;Porta Westfalica Airport;Bad Oeynhausen;Germany;;EDVY;52.2208;8.85917;148;1;E;Europe/Berlin
8103;Brilon Hochsauerlandkreis Airport;Brilon;Germany;;EDKO;51.4025;8.64167;1509;1;E;Europe/Berlin
8104;Hameln Pyrmont Airport;Bad Pyrmont;Germany;;EDVW;51.9667;9.29167;1178;1;E;Europe/Berlin
8105;Nordholz Spieka Airport;Cuxhaven;Germany;;EDXN;53.7672;8.64361;72;1;E;Europe/Berlin
8106;Koethen Airport;Koethen;Germany;;EDCK;51.7211;11.9528;305;1;E;Europe/Berlin
8107;St. Michaelisdonn Airport;Sankt Michaelisdonn;Germany;;EDXM;53.9781;9.14472;125;1;E;Europe/Berlin
8108;Salzgitter Druette Airport;Salzgitter;Germany;;EDVS;52.1544;10.4267;328;1;E;Europe/Berlin
8109;Karlshoefen Airport;Karlshoefen;Germany;;EDWK;53.3328;9.02833;20;1;E;Europe/Berlin
8110;Oldenburg Hatten Airport;Oldenburg;Germany;;EDWH;53.0689;8.31361;26;1;E;Europe/Berlin
8111;Rinteln Airport;Rinteln;Germany;;EDVR;52.1753;9.05333;180;1;E;Europe/Berlin
8112;Muenster-Telgte Airport;Muenster;Germany;;EDLT;51.9444;7.77389;177;1;E;Europe/Berlin
8113;St. Peter-Ording Airport;Sankt Peter-Ording;Germany;PSH;EDXO;54.3089;8.68694;7;1;E;Europe/Berlin
8114;Luechow-Rehbeck Airport;Luechow;Germany;;EDHC;53.0161;11.1444;49;1;E;Europe/Berlin
8115;Klietz-Scharlibbe Airport;Scharlibbe;Germany;;EDCL;52.7094;12.0733;95;1;E;Europe/Berlin
8116;Burg Airport;Burg;Germany;;EDBG;52.2417;11.8561;174;1;E;Europe/Berlin
8117;Crisp County Cordele Airport;Cordele;United States;CKF;KCKF;31.9888333;-83.7739167;310;-5;A;America/New_York
8118;Ormond Beach municipal Airport;Ormond Beach;United States;OMN;KOMN;29.1804;-81.06497;28;-5;A;America/New_York
8119;Bad Neuenahr-Ahrweiler Airport;Bad Neuenahr;Germany;;EDRA;50.5578;7.13639;673;1;E;Europe/Berlin
8120;Bad Duerkheim Airport;Bad Duerkheim;Germany;;EDRF;49.4731;8.19639;351;1;E;Europe/Berlin
8121;Portland Troutdale;Troutdale;United States;TTD;KTTD;45.54937;-122.401253;39;-8;A;America/Los_Angeles
8122;Portland Hillsboro;Hillsboro;United States;HIO;KHIO;45.540394;-122.949825;204;-8;A;America/Los_Angeles
8123;One Police Plaza Heliport;New York;United States;;NK39;40.7126;-73.9996;244;-5;A;America/New_York
8124;Leverkusen Airport;Leverkusen;Germany;;EDKL;51.0153;7.00556;157;1;E;Europe/Berlin
8125;Suwannee County Airport;Live Oak;United States;24J;\N;30.300125;-83.0246944;104;-5;A;America/New_York
8126;Wershofen Eifel;Wershofen;Germany;;EDRV;50.451667;6.784444;1560;1;E;Europe/Berlin
8127;FOB Salerno;Khost;Afghanistan;KHT;OAKS;33.3334;69.952;3756;4.5;N;Asia/Kabul
8128;TCO;Tengiz;Kazakhstan;;\N;46.303056;53.4275;0;5;U;Asia/Oral
8129;TCO;Tengiz;Kazakhstan;;\N;46.303056;53.4275;0;5;U;Asia/Oral
8130;NAYPYITAW;NAYPYITAW;Burma;NYT;VYNT;19.374;96.121;294;6.5;N;Asia/Rangoon
8132;Maramba Aerodrome;Livingstone;Zambia;;\N;-17.8846;25.8468;0;2;U;Africa/Lusaka
8133;Bend Municipal Airport;Bend;United States;;KBDN;44.0945556;-121.2002222;3460;-8;A;America/Los_Angeles
8134;Christmas Valley Airport;Christmas Valley;United States;;K62S;43.2365314;-120.6660967;4317;-8;A;America/Los_Angeles
8135;Burns Municipal Airport;Burns;United States;;KBNO;43.5919167;-118.95554444;4159;-8;A;America/Los_Angeles
8136;Prineville Airport;Prineville;United States;;KS39;44.2869939;-120.9038328;3250;-8;A;America/Los_Angeles
8137;Red Bluff Municipal Airport;Red Bluff;United States;;KRBL;40.1506944;-122.2523056;352;-8;A;America/Los_Angeles
8138;Gnoss Field Airport;Novato;United States;;KDVO;38.1436111;-122.5561;2;-8;A;America/Los_Angeles
8139;Lake County Airport;Lakeview;United States;;KLKV;42.1611111;-120.3990833;4733;-8;A;America/Los_Angeles
8140;Tillamook Airport;Tillamook;United States;;KTMK;45.4182419;-123.8143839;36;-8;A;America/Los_Angeles
8141;Ontario Municipal Airport;Ontario;United States;;KONO;44.0193611;-117.0130278;2193;-7;A;America/Denver
8142;Columbia Gorge Regional - The Dalles Municipal Airport;The Dalles;United States;;KDLS;45.6185556;-121.1673333;247;-8;A;America/Los_Angeles
8143;Montgomery County Airpark;Gaithersburg;United States;GAI;KGAI;39.1006;-77.09576;0;-5;A;America/New_York
8144;Cincinnati Union Terminal;Cincinnati;United States;;\N;39.109961;-84.537074;497;-5;A;America/New_York
8145;FOB Sharana;Sharan;Afghanistan;;\N;33.12777215;68.8369847;7400;4.5;N;Asia/Kabul
8146;Sharona;Sharona;Afghanistan;AZ3;OASA;33.1277215;68.8369847;7400;4.5;N;Asia/Kabul
8147;Pembroke Airport;Pembroke;Canada;YTA;CYTA;45.864399;-77.251701;529;-5;A;America/Toronto
8148;Tsumeb Airport;Tsumeb;Namibia;TSB;FYTM;-19.1542;17.4357;4353;1;U;Africa/Windhoek
8149;Suffield Heliport;Suffield;Canada;YSD;CYSD;50.2666701;-111.182999;2525;-7;A;America/Edmonton
8150;Field 21;Wainwright;Canada;;\N;52.830601;-111.100998;2260;-7;A;America/Edmonton
8151;Field 21;Wainwright;Canada;;\N;52.830601;-111.100998;2260;-7;A;America/Edmonton
8152;Aeroporto Blumenau;BLUMENAU;Brazil;BNU;SSBL;-26.834239;-49.091696;70;-3;S;America/Sao_Paulo
8153;Aeroclub Mures;Targu Mures;Romania;;LRMS;46.3201;24.3157;1000;2;E;Europe/Bucharest
8154;Aeroclub Sibiu;Sibiu;Romania;;\N;45.4649;24.053;1451;2;E;Europe/Bucharest
8155;Aeroclub Sibiu;Sibiu;Romania;SIB;\N;45.4649;24.053;1451;2;E;Europe/Bucharest
8156;Bolshoe Gryzlovo;Stupino;Russia;;UUDG;54.786667;37.649167;797;4;N;Europe/Moscow
8157;Aerodrom Cioca;Timisoara;Romania;;\N;46.436;24.4445;280;2;E;Europe/Bucharest
8158;Aeroclub Cioca;Timisoara;Romania;CIO;LRT2;45.471009;21.111967;280;2;E;Europe/Bucharest
8159;Crocodile Camp Air Strip;Tsavo East;Kenya;;\N;-3.0764;39.2428;0;3;U;Africa/Nairobi
8160;Crocodile Camp Air Strip;Tsavo East;Kenya;;\N;-3.0764;39.2428;0;3;U;Africa/Nairobi
8161;Cleveland Clinic;Cleveland;United States;;CCLN;41.50225;-81.62233;0;-5;A;America/New_York
8162;Charlevoix Municipal Airport;Charelvoix;United States;CVX;KCVX;45.3047778;-85.2753333;669;-5;A;America/New_York
8163;Quincy Municipal Airport;Quincy;United States;2J9;\N;30.5978708;-84.557425;225;-5;A;America/New_York
8164;Pferdsfeld;Pferdsfeld;Germany;;ETSP;49.857778;7.604167;492;1;E;Europe/Berlin
8165;Mykines Heliport;Mykines;Faroe Islands;;EKMS;62.1021;-7.6459;112;0;E;Atlantic/Faeroe
8166;V-A Waterfront;Cape Town;South Africa;;\N;-33.901161;18.425896;0;2;N;Africa/Johannesburg
8167;Steenberg Helipad;Cape Town;South Africa;;\N;-34.070582;18.424067;50;2;N;Africa/Johannesburg
8168;Steenberg Helipad;Cape Town;South Africa;;\N;-34.070582;18.424067;50;2;N;Africa/Johannesburg
8170;Roche Harbor Seaplane Base;Roche Harbor;United States;RCE;\N;48.608056;-123.159722;0;-8;A;America/Los_Angeles
8171;Blakely Island Airport;Blakely Island;United States;BYW;\N;48.56025;-122.80243;0;-8;A;America/Los_Angeles
8172;Rosario Seaplane Base;Rosario;United States;RSJ;\N;48.645556;-122.868056;0;-8;A;America/Los_Angeles
8173;Westsound Seaplane Base;Westsound;United States;WSX;\N;48.617778;-122.952778;0;-8;A;America/Los_Angeles
8174;Friday Harbor Seaplane Base;Friday Harbor;United States;FBS;\N;48.537222;-123.009722;0;-8;A;America/Los_Angeles
8175;Big Bay Water Aerodrome;Big Bay;Canada;YRR;\N;50.4;-125.133333;0;-8;A;America/Vancouver
8176;Furnace Creek;Death Valley National Park;United States;L06;\N;36.273;-116.515;0;-8;A;America/Los_Angeles
8177;Cornwall Regional Airport;Cornwall;Canada;YCC;CYCC;45.092778;-74.567778;175;-5;A;America/Toronto
8178;Seppe;Bosschenhoofd;Netherlands;;EHSE;51.3317;4.339;30;1;E;Europe/Brussels
8179;St.Stephan;St.Stephan;Switzerland;;LSTS;46.497442;7.412572;3304;1;E;Europe/Zurich
8180;Zona da Mata Regional Airport;Juiz de Fora;Brazil;IZA;SDZY;-21.5130558014;-43.1730575562;1348;-3;S;America/Sao_Paulo
8181;Aeroclub Ghimbav;Brasov;Romania;;LRBG;45.4153;25.3137;1743;2;E;Europe/Bucharest
8182;Flagler County Airport;Flagler;United States;XFL;KXFL;29.2821;-81.1212;33;-5;A;America/New_York
8183;Aeroclub Deva;Deva;Romania;DVA;\N;45.5153;22.5813;615;2;U;Europe/Bucharest
8184;Aeroclub Cluj;Dezmir;Romania;DZM;\N;46.4643;23.4258;1019;2;U;Europe/Bucharest
8185;Antarctica;Antarctica;Antarctica;;\N;-66;90;8000;12;U;Antarctica/South_Pole
8186;Antarctica;Antarctica;Antarctica;;\N;-66;90;8000;12;N;Antarctica/South_Pole
8187;Morrisville Stowe State Airport;Morrisville;United States;MVL;KMVL;44.535;-72.614;732;-5;A;America/New_York
8188;Dallas Executive Airport;Dallas;United States;RBD;KRBD;32.680833;-96.868333;201;-6;A;America/Chicago
8189;Como Water AD;Como;Italy;;LILY;45.4853;9.0411;663;1;E;Europe/Rome
8190;Superficie Cielo Blu;Vellezzo Bellini;Italy;;\N;45.283855;9.111979;350;1;E;Europe/Rome
8191;Welke Airport;Beaver Island;United States;6Y8;\N;45.721111;-85.520278;664;-5;A;America/New_York
8192;Krainiy;Baikonur;Kazakhstan;;UAOL;45.6183;63.2144;328;6;U;Asia/Qyzylorda
8193;MOW;Moscow;Russia;MOW;\N;55.7557;37.6176;0;4;N;Europe/Moscow
8194;Westerly State Airport;Washington County;United States;WST;KWST;41.349722;-71.803333;81;-5;A;America/New_York
8195;Block Island State Airport;Block Island;United States;BID;KBID;41.168056;-71.577778;108;-5;A;America/New_York
8196;Atmautluak Airport;Atmautluak;United States;;\N;60.866667;-162.273056;18;-9;A;America/Anchorage
8197;Atmautluak Airport;Atmautluak;United States;;\N;60.866667;-162.273056;18;-9;A;America/Anchorage
8198;Atmautluak Airport;Atmautluak;United States;369;\N;60.866667;-162.273056;18;-9;A;America/Anchorage
8199;Nightmute Airport;Nightmute;United States;NME;PAGT;60.471111;-164.700833;4;-9;A;America/Anchorage
8200;Toksook Bay Airport;Toksook Bay;United States;OOK;PAOO;60.541389;-165.087222;59;-9;A;America/Anchorage
8201;Tununak Airport;Tununak;United States;TNK;\N;60.575556;-165.271667;14;-9;A;America/Anchorage
8202;Goodnews Airport;Goodnews Bay;United States;GNU;\N;59.1175;-161.5775;15;-9;A;America/Anchorage
8203;Newtok Airport;Newtok;United States;WWT;\N;60.939167;-164.641111;25;-9;U;America/Anchorage
8204;Achutupo Airport;Achutupo;Panama;ACU;\N;9.2;-77.98;0;-5;U;America/Panama
8205;Tubuala Airport;Tubuala;Panama;TUW;\N;9.52;-79.03;0;-5;U;America/Panama
8206;Garachine Airport;Garachine;Panama;GHE;\N;8.06;-78.36;0;-5;U;America/Panama
8207;Mulatupo Airport;Mulatupo;Panama;MPP;\N;8.95;-77.75;0;-5;U;America/Panama
8208;Ittoqqortoormiit Heliport;Ittoqqortoormiit;Greenland;OBY;BGSC;70.485278;-21.966667;238;-1;U;America/Scoresbysund
8209;Vinnitsa;Vinnitsa;Ukraine;VIN;UKWW;49.2433;28.6063;900;2;E;Europe/Kiev
8210;Gent Sint-Pieters;Gent;Belgium;;\N;51.0352;3.7097;130;1;E;Europe/Brussels
8211;Brugge;Bruges;Belgium;;\N;51.1972;3.2172;120;1;E;Europe/Brussels
8212;Brugge;Bruges;Belgium;;\N;51.1972;3.2172;130;1;E;Europe/Brussels
8213;Shin-yokohama-eki;Yokohama;Japan;;\N;35.5075;139.6175;50;9;N;Asia/Tokyo
8214;Angers St Laud;Angers;France;QXG;\N;47.464714;-0.556405;300;1;E;Europe/Paris
8215;Decatur County Industrial Air Park;Bainbridge;United States;BGE;KBGE;30.9715981;-84.6369278;141;-5;A;America/New_York
8216;La Romaine Airport;La Romaine;Canada;;CTT5;50.257222;-60.669167;90;-4;A;America/Blanc-Sablon
8217;Kegaska Airport;Kegaska;Canada;ZKG;CTK6;50.195833;-61.265833;32;-4;A;America/Blanc-Sablon
8218;Black Tickle Airport;Black Tickle;Canada;YBI;CCE4;53.47;-55.7875;57;-4;A;America/Halifax
8219;Silver Springs Airport;Silver Springs;United States;SPZ;KSPZ;39.4030278;-119.2511944;4269;-8;A;America/Los_Angeles
8220;Whiteman Airport;Los Angeles;United States;WHP;KWHP;34.2593253;-118.4134331;1003;-8;A;America/Los_Angeles
8221;Madera Municipal Airport;Madera;United States;MAE;KMAE;36.9886111;-120.1124444;255;-8;A;America/Los_Angeles
8222;Mountain Home Municipal Airport;Mountain Home;United States;U76;\N;43.1316111;-115.7305671;3167;-7;A;America/Denver
8223;Trail Airport;Trail;Canada;YZZ;CAD4;49.055556;-117.609167;1427;-8;A;America/Vancouver
8224;Victoria Airport Water Aerodrome;Patricia Bay;Canada;;CAP5;48.65;-123.45;0;-8;A;America/Vancouver
8225;Arctic Bay Airport;Arctic Bay;Canada;YAB;CJX7;73.006389;-85.047222;72;-6;A;America/Winnipeg
8226;Hope Bay Aerodrome;Hope Bay;Canada;;CHB3;68.163889;-106.614444;125;-7;A;America/Edmonton
8227;Hector Silva Airstrip;Belmopan;Belize;BCV;\N;17.269444;-88.776111;50;-6;U;America/Belize
8228;Grand-Santi Airport;Grand-Santi;French Guiana;;SOGS;4.283333;-54.381111;186;-3;S;America/Cayenne
8229;Maripasoula Airport;Maripasoula;French Guiana;MPY;SOOA;3.6575;-54.037222;377;-3;S;America/Cayenne
8230;Saint-Laurent-du-Maroni Airport;Saint-Laurent-du-Maroni;French Guiana;LDX;SOOM;5.483056;-54.034444;17;-3;S;America/Cayenne
8231;Cayana Airstrip;Cayana;Suriname;AAJ;\N;3.898611;-55.577778;50;-3;S;America/Paramaribo
8232;Laduani Airstrip;Laduani;Suriname;LDO;\N;4.3754;-55.4075;50;-3;S;America/Paramaribo
8233;Kanas Airport;Burqin;China;KJI;ZWKN;48.221111;86.998056;3900;8;U;Asia/Chongqing
8234;Capurgana Airport;Capurgana;Colombia;CPB;SKCA;8.616667;-77.333333;100;-5;S;America/Bogota
8235;Sohag International;Sohag;Egypt;HMB;HEMK;26.341189;31.742983;859;2;E;Africa/Cairo
8236;Rivera International Airport;Rivera;Uruguay;RVY;SURV;-30.974444;-55.476111;712;-3;S;America/Montevideo
8237;Patos de Minas Airport;Patos de Minas;Brazil;POJ;SNPD;-18.672778;-46.491111;851;-3;S;America/Sao_Paulo
8238;Bauru-Arealva;Bauru;Brazil;JTC;SJTC;-22.157778;-49.068333;594;-3;S;America/Sao_Paulo
8239;Ourilandia do Norte Airport;Ourilandia do Norte;Brazil;OIA;SDOW;-6.763056;-51.05;229;-4;S;America/Boa_Vista
8240;Redencao Airport;Redencao;Brazil;RDC;SNDC;-8.033333;-49.98;669;-4;S;America/Boa_Vista
8241;Sao Felix do Xingu Airport;Sao Felix do Xingu;Brazil;SXX;SNFX;-6.641389;-51.952222;656;-4;S;America/Boa_Vista
8242;Bonito Airport;Bointo;Brazil;BYO;SJDB;-21.247222;-56.4525;1180;-4;S;America/Campo_Grande
8243;Sao Felix do Araguaia Airport;Sao Felix do Araguaia;Brazil;SXO;SWFX;-11.6325;-50.689444;650;-4;S;America/Campo_Grande
8244;Carlos Alberto da Costa Neves Airport;Cacador;Brazil;CFC;SBCD;-26.788333;-50.939722;3376;-3;S;America/Sao_Paulo
8245;Carauari Airport;Carauari;Brazil;CAF;SWCA;-4.871389;-66.8975;354;-4;S;America/Boa_Vista
8246;Porto Urucu;Porto Urucu;Brazil;;SWUY;-4.884167;-65.355556;210;-4;S;America/Boa_Vista
8247;Amaury Feitosa Tomaz Airport;Eirunepe;Brazil;ERN;SWEI;-6.639444;-69.879722;412;-4;S;America/Boa_Vista
8248;Concordia Airport;Concordia;Brazil;CCI;SSCK;-27.180556;-52.052778;2461;-3;S;America/Sao_Paulo
8249;Paulo Abdala Airport;Francisco Beltrao;Brazil;FBE;SSFB;-26.059167;-53.063333;2100;-3;S;America/Sao_Paulo
8250;Confresa Airport;Confresa;Brazil;CFO;SJHG;-10.633611;-51.567222;250;-4;S;America/Campo_Grande
8251;Jackson County Airport;Jefferson;United States;19A;\N;34.1758638;-83.5615972;951;-5;U;America/New_York
8252;Apalachicola Regional Airport;Apalachicola;United States;AAF;KAAF;29.7276066;-85.0274416;20;-5;A;America/New_York
8253;Orlando de Carvalho Airport;Umuarama;Brazil;UMU;SSUM;-23.798611;-53.313611;1558;-3;S;America/Sao_Paulo
8254;Diamantina Airport;Diamantina;Brazil;DTI;SNDT;-18.231944;-43.650278;4446;-3;S;America/Sao_Paulo
8255;Fonte Boa Airport;Fonte Boa;Brazil;FBA;SWOB;-2.5325;-66.083333;207;-4;S;America/Boa_Vista
8256;Senadora Eunice Micheles Airport;Sao Paulo de Olivenca;Brazil;OLC;SDCG;-3.465556;-68.918889;335;-4;S;America/Boa_Vista
8257;Humaita Airport;Humaita;Brazil;HUW;SWHT;-7.532222;-63.072222;230;-4;S;America/Boa_Vista
8258;Tapuruquara Airport;Santa Isabel do Rio Negro;Brazil;IRZ;SWTP;-0.416944;-65.033889;223;-4;S;America/Boa_Vista
8259;Oriximina Airport;Oriximina;Brazil;ORX;SNOX;-1.714167;-55.836111;262;-4;S;America/Boa_Vista
8260;Una-Comandatuba Airport;Una;Brazil;UNA;SBTC;-15.355278;-38.998889;20;-3;S;America/Fortaleza
8261;Algeciras Heliport;Algeciras;Spain;;\N;36.128889;-5.441111;98;1;E;Europe/Madrid
8262;Algeciras Heliport;Algeciras;Spain;;\N;36.128889;-5.441111;98;1;E;Europe/Madrid
8263;Kai Tak International Airport ;Hong Kong;Hong Kong;;\N;22.3177;114.202;30;8;U;Asia/Hong_Kong
8264;Kai Tak International Airport ;Hong Kong;Hong Kong;;\N;22.3177;114.202;30;8;U;Asia/Hong_Kong
8265;Telfer Airport;Telfer;Australia;TEF;YTEF;-21.713057;122.21222;970;8;N;Australia/Perth
8266;Gazipasa Airport;Alanya;Turkey;GZP;LTFG;36.2993;32.3014;92;2;E;Europe/Istanbul
8267;FOB Shank;Shank;Afghanistan;;OASH;33.921369;69.07813;6614;4.5;N;Asia/Kabul
8268;St. Augustine Airport;St. Augustine;United States;SGJ;\N;29.95925;-81.3397222;10;-5;A;America/New_York
8269;Londolozi;Londolozi;South Africa;LDZ;\N;-24.8333;31.5;3510;2;N;Africa/Johannesburg
8270;Singita Sabi Sands;Sabi Sands;South Africa;;\N;-24.84;31.41;3500;2;N;Africa/Johannesburg
8271;Singita Sabi Sands;Sabi Sands;South Africa;;\N;-24.84;31.41;3500;2;N;Africa/Johannesburg
8272;Singita Sabi Sands;Sabi Sands;South Africa;INY;\N;-24.815;31.42;1200;2;N;Africa/Johannesburg
8273;Douglas Municipal Airport;Douglas;United States;DQH;KDQH;31.4767385;-82.8605664;257;-5;A;America/New_York
8274;St Lucie County International Airport;Fort Pierce;United States;FRP;KFPR;27.497472;-80.372638;23;-5;A;America/New_York
8275;Alexandria;Alexandria;United States;ALX;ALX_;38.806346;-77.0621;1;-5;A;America/New_York
8276;Taunton Municipal Airport - King Field;Taunton;United States;TAN;KTAN;41.8744017;-71.0166453;43;-5;A;America/New_York
8277;Plymouth Municipal Airport;Plymouth;United States;PYM;KPYM;41.9090278;-70.7287778;148;-5;A;America/New_York
8278;Quonset State Airport;North Kingstown;United States;OQU;KOQU;41.5971389;-71.4121389;18;-5;A;America/New_York
8279;Mansfield Municipal;Mansfield;United States;1B9;\N;42.0001331;-71.1967714;122;-5;A;America/New_York
8280;Norwood Memorial Airport;Norwood;United States;OWD;KOWD;42.1905278;-71.1729444;49;-5;A;America/New_York
8281;Barnes Municipal;Westfield;United States;BAF;KBAF;42.1579472;-72.715875;270;-5;A;America/New_York
8282;Windham Airport;Willimantic;United States;IJD;KIJD;41.7440278;-72.1802222;246;-5;A;America/New_York
8283;Orange County Airport;Montgomery;United States;MGJ;KMGJ;41.5099884;-74.2646444;364;-5;A;America/New_York
8284;Capital City Airport;Harrisburg;United States;CXY;KCXY;40.2171389;-76.8513611;347;-5;A;America/New_York
8285;Marshfield Municipal Airport;Marshfield;United States;GHG;KGHG;42.09825;-70.6721389;9;-5;A;America/New_York
8286;Danbury Municipal Airport;Danbury;United States;DXR;KDXR;41.3715353;-73.4821906;458;-5;A;America/New_York
8287;Boire Field Airport;Nashua;United States;ASH;KASH;42.78175;-71.5147778;199;-5;A;America/New_York
8288;Lawrence Municipal Airport;Lawrence;United States;LWM;KLWM;42.7171944;-71.1234167;148;-5;A;America/New_York
8289;Waterbury-Oxford Airport;Oxford;United States;OXC;KOXC;41.4785556;-73.13525;726;-5;A;America/New_York
8290;Fitchburg Municipal Airport;Fitchburg;United States;FIT;KFIT;42.5541111;-71.7589722;348;-5;A;America/New_York
8291;Stockmar Airport;Villa Rica;United States;;20GA;33.7565556;-84.8850625;1110;-5;A;America/New_York
8292;Cartersville Airport;Cartersville;United States;VPC;KVPC;34.1231475;-84.8487067;759;-5;A;America/New_York
8293;Centre-Piedmont-Cherokee County Regional Airport;Centre;United States;PYP;KPYP;34.0899167;-85.6100833;595;-6;A;America/Chicago
8294;Richard B Russell Airport;Rome;United States;RMG;KRMG;34.3507778;-85.1586667;664;-5;A;America/New_York
8295;Northeast Alabama Regional Airport;Gadsden;United States;GAD;KGAD;33.9726489;-86.0890834;569;-6;A;America/Chicago
8296;Knoxville Downtown Island Airport;Knoxville;United States;DKX;KDKX;35.9638333;-83.8736667;833;-5;A;America/New_York
8297;Barrow County Airport;Winder;United States;WDR;KWDR;33.9828611;-83.6674167;943;-5;A;America/New_York
8298;Plantation Airpark;Sylvania;United States;JYL;KJYL;32.6452778;-81.5971111;188;-5;A;America/New_York
8299;Dalton Municipal Airport;Dalton;United States;DNN;KDNN;34.7229444;-84.87025;709;-5;A;America/New_York
8300;West Georgia Regional Airport - O V Gray Field;Carrollton;United States;CTJ;KCTJ;33.6316964;-85.1522641;1165;-5;A;America/New_York
8301;Isbell Field Airport;Fort Payne;United States;4A9;\N;34.4736944;-85.7213889;877;-6;A;America/Chicago
8302;LaGrange-Callaway Airport;LaGrange;United States;LGC;KLGC;33.0088611;-85.0726111;694;-5;A;America/New_York
8303;Baldwin County Airport;Milledgeville;United States;MLJ;KMLJ;33.154225;-83.2414139;385;-5;A;America/New_York
8304;Polk County Airport - Cornelius Moore Field;Cedartown;United States;4A4;\N;34.0186944;-85.1464722;974;-5;A;America/New_York
8305;Harris County Airport;Pine Mountain;United States;PIM;KPIM;32.8406944;-84.8824444;902;-5;A;America/New_York
8306;Atlanta Regional Airport - Falcon Field;Atlanta;United States;FFC;KFFC;33.35725;-84.5718333;808;-5;A;America/New_York
8307;Covington Municipal Airport;Covington;United States;9A1;\N;33.6323083;-83.8474472;809;-5;A;America/New_York
8308;Lee Gilmer Memorial Airport;Gainesville;United States;GVL;KGVL;34.272627;-83.8302233;1276;-5;A;America/New_York
8309;Cherokee County Airport;Canton;United States;;KCNI;34.3122175;-84.4221556;1219;-5;A;America/New_York
8310;DeFuniak Springs Airport;DeFuniak Springs;United States;54J;\N;30.7311111;-86.1537778;289;-6;A;America/Chicago
8311;Barwick Lafayette Airport;LaFayette;United States;9A5;\N;34.6884792;-85.2903319;777;-5;A;America/New_York
8312;Harry Clever Field Airport;New Philadelpha;United States;PHD;KPHD;40.4701667;-81.4199444;894;-5;A;America/New_York
8313;Darlington County Jetport;Darlington;United States;UDG;KUDG;34.4493845;-79.8900608;192;-5;A;America/New_York
8314;Hilton Head Airport;Hilton Head Island;United States;HXD;KHXD;32.2243611;-80.6974722;19;-5;A;America/New_York
8315;Gilmer County Airport;Ellijay;United States;49A;\N;34.6282222;-84.5265833;1486;-5;A;America/New_York
8316;Elizabethton Municipal Airport;Elizabethton;United States;0A9;\N;36.3712222;-82.1734167;1593;-5;A;America/New_York
8317;Moton Field Municipal Airport;Tuskegee;United States;06A;\N;32.4605722;-85.6800278;264;-6;A;America/Chicago
8318;Daniel Field Airport;Augusta;United States;DNL;KDNL;33.4665028;-82.0393808;423;-5;A;America/New_York
8319;Foothills Regional Airport;Morganton;United States;MRN;KMRN;35.8202336;-81.6115119;1270;-5;A;America/New_York
8320;Pike County Airport - Hatcher Field;Pikeville;United States;PBX;KPBX;37.5617639;-82.5663889;1473;-5;A;America/New_York
8321;Mallards Landing Airport;Locust Grove;United States;;GA04;33.3656703;-84.1651964;837;-5;A;America/New_York
8322;Toccoa RG Letourneau Field Airport;Toccoa;United States;TOC;KTOC;34.5928117;-83.2963717;996;-5;A;America/New_York
8323;Shaftesbury-Compton Abbas Aerodrome;Shaftesbury;United Kingdom;;EGHA;50.9672;-2.1536;811;0;U;Europe/London
8324;SCC4;MFO;Egypt;;\N;30.524819;34.584025;2037;2;U;Asia/Jerusalem
8325;SCC4;MFO;Egypt;;\N;30.524819;34.584025;2037;2;E;Asia/Jerusalem
8326;SCC4;MFO;Egypt;;\N;30.524819;34.584025;2037;2;E;Asia/Jerusalem
8327;SCC4;MFO;Egypt;;\N;30.524819;34.584025;2037;2;E;Asia/Jerusalem
8328;Poltava;Poltava;Ukraine;PLV;UKHP;49.34261;34.23551;333;2;U;Europe/Kiev
8329;Ferry Port;Calais;France;;\N;50.967536;1.849308;6;1;E;Europe/Paris
8330;Ferry Port;Dover;United Kingdom;;\N;51.126876;1.339747;6;0;E;Europe/London
8331;Aweil Airport;Aweil;Sudan;;HSAW;8.77;27.3975;1394;3;U;Africa/Juba
8332;Wau Airport;Wau;Sudan;WUU;HSWW;7.725;27.98;1420;3;U;Africa/Juba
8333;Humera Airport;Humera;Ethiopia;HUE;HAHU;14.3;36.616667;1500;3;U;Africa/Asmera
8334;Moyale Airport;Moyale;Kenya;OYL;HKMY;3.465;39.105;2790;3;U;Africa/Nairobi
8335;Zagora Airport;Zagora;Morocco;;GMAZ;30.330556;-5.838056;1500;0;U;Africa/Casablanca
8336;Yengema Airport;Yengema;Sierra Leone;WYE;GFYE;8.615308;-11.047192;1300;0;U;Africa/Freetown
8337;Gbangbatok Airport;Gbangbatok;Sierra Leone;GBK;GFGK;7.767;-12.383;500;0;U;Africa/Freetown
8338;Hong Kong China Ferry Terminal;Hong Kong;Hong Kong;;\N;22.299204;114.166164;0;8;N;Asia/Hong_Kong
8339;Tampa Executive Airport;Tampa;United States;VDF;\N;28.0139772;-82.3452778;22;-5;A;America/New_York
8340;UOTT;Turukhansk;Russia;;UOTT;65.46;87.53;125;8;N;Asia/Krasnoyarsk
8341;Podkamennaya Tunguska;Bor;Russia;;UNIP;61.34;89.56;213;8;N;Asia/Krasnoyarsk
8342;Fort Worth Alliance Airport;Fort Worth;United States;AFW;KAFW;32.987778;-97.318889;722;-6;A;America/Chicago
8343;East Troy Municipal Airport;East Troy;United States;57C;K57C;42.7971667;-88.3726111;860;-6;A;America/Chicago
8344;Kolpashevo Airport;Kolpashevo;Russia;;UNLL;58.3333;82.9333;243;7;N;Asia/Omsk
8345;Montgomery Field;San Diego;United States;MYF;NULL;32.4759;117.759;17;8;A;Asia/Chongqing
8346;Gimpo International Airpot;Seoul;South Korea;SEL;\N;37.558311;126.790586;58;9;N;Asia/Seoul
8347;Ali Al Salem Air Base;Kuwait;Kuwait;;OKAS;29.346742;47.520753;472;3;A;Asia/Kuwait
8348;Renmark;Renmark;Australia;RMK;YREN;-34.118;140.404;115;9.5;O;Australia/Adelaide
8349;Leigh Creek;Leigh Creek;Australia;;YLEC;-30.598333;138.425833;856;9.5;O;Australia/Adelaide
8350;Warburton;Warburton Community;Australia;;YWBR;-26.128333;126.583333;1500;8;O;Australia/Perth
8351;Cunderdin;Cunderdin;Australia;;YCUN;-31.622223;117.216667;705;8;O;Australia/Perth
8352;Rottnest Island;Rottnest Island;Australia;;YRTI;-32.006667;115.539722;12;8;O;Australia/Perth
8353;Forrest;Forrest;Australia;;YFRT;-30.838056;128.115;511;8;O;Australia/Perth
8354;Ballarat;Ballarat;Australia;;YBLT;-37.511667;143.791111;1433;10;O;Australia/Hobart
8355;Keewaywin;Keewaywin;Canada;KEW;CPV8;52.991111;-92.836389;988;-6;A;America/Winnipeg
8356;Marathon;Marathon;Canada;YSP;CYSP;48.755278;-86.344444;1035;-5;A;America/Toronto
8357;Rene Fontaine;Hearst;Canada;YHF;CYHF;49.713889;-83.686944;826;-5;A;America/Toronto
8358;Hornepayne;Hornepayne;Canada;YHN;CYHN;49.193056;-84.758889;1099;-5;A;America/Toronto
8359;Kirkland Lake;Kirkland Lake;Canada;YKX;CYKX;48.210278;-79.981389;1157;-5;A;America/Toronto
8360;Manitouwadge;Manitouwadge;Canada;YMG;CYMG;49.083889;-85.860556;1198;-5;A;America/Toronto
8361;Wawa;Wawa;Canada;YXZ;CYXZ;47.966944;-84.786389;944;-5;A;America/Toronto
8362;Manitoulin East;Manitowaning;Canada;YEM;CYEM;45.8425;-81.8575;869;-5;A;America/Toronto
8363;Slate Falls;Slate Falls;Canada;;CKD9;51.13;-91.665556;1355;-6;A;America/Winnipeg
8364;Collingwood;Collingwood;Canada;;CNY3;44.449167;-80.158333;730;-5;A;America/Toronto
8365;Brantford;Brantford;Canada;YFD;CYFD;43.131389;-80.3425;815;-5;A;America/Toronto
8366;Lawrence Municipal;Lawrence;United States;LWC;KLWC;39.009167;-95.2175;833;-6;A;America/Chicago
8367;Wellington Municipal;Wellington;United States;EGT;KEGT;37.39411;-97.423225;1277;-6;A;America/Chicago
8368;Isla San Jose;Isla San Jose;Panama;;\N;8.252;-79.0668;150;-5;N;America/Panama
8369;Griffin-Spalding County Airport;Griffin;United States;6A2;\N;33.2269722;-84.2749444;958;-5;A;America/New_York
8370;Dieu Merci;Saint-Elie;French Guiana;;\N;4.7828;-53.2611;253;-3;N;America/Cayenne
8371;DEGRA;REMIRE-MONTJOLY;French Guiana;;\N;4.8708;52.2783;33;-3;N;\N
8372;Herapel;Morsbach;France;;\N;49.1606;6.8649;840;1;E;Europe/Paris
8373;Degrad des Cannes;REMIRE-MONTJOLY;French Guiana;;\N;4.8708;-52.2783;33;-3;N;America/Cayenne
8374;Elysee;Elysee;French Guiana;;\N;4.75;-54.05;100;-3;N;America/Cayenne
8375;Citron;Citron;French Guiana;;\N;4.7301;-53.9516;100;-3;N;America/Cayenne
8376;Yaou;Maripasoula;French Guiana;;\N;3.7243;-53.9747;300;-3;N;America/Cayenne
8377;Helicojyp;REMIRE-MONTJOLY;French Guiana;;\N;4.8936;-52.2922;39;-3;N;America/Cayenne
8378;Nicosia International Airport;Nicosia;Cyprus;NIC;\N;35.15111;33.27222;0;2;E;Asia/Nicosia
8379;Pompano Beach Airpark;Pompano Beach;United States;PMP;KPMP;26.2471389;-80.1110556;19;-5;A;America/New_York
8380;Mallacoota Airport;Mallacoota;Australia;XMC;YMCO;-37.598301;149.720001;31;10;U;Australia/Hobart
8381;Prince Abdul Majeed Airport;Al-Ula;Saudi Arabia;ULH;OEAO;26.636718;37.908018;2047;3;N;Asia/Riyadh
8382;Pogapa Airstrip;Pogapa;Indonesia;;\N;-3.7518;136.842481;6060;9;N;Asia/Jayapura
8383;Pogapa Airstrip;Pogapa;Indonesia;;\N;-3.7518;136.842481;6060;9;N;Asia/Jayapura
8384;Greyhound Station;Washington DC;United States;;DCBS;38.90339;-77.00583;0;-5;A;America/New_York
8385;Roma Street Railway Station;Brisbane;Australia;;\N;-27.4666667;-153.0166667;100;10;N;\N
8386;Roma Street Railway Station;Brisbane;Australia;;\N;-27.4666667;153.0166667;100;10;N;Australia/Brisbane
8387;Rockhampton Railway Station;Rockhampton;Australia;;\N;-23.378941;150.512323;100;10;N;Australia/Brisbane
8388;Sydney Terminal Railway Station;Sydney;Australia;;\N;-33.88373;151.20592;100;10;O;Australia/Sydney
8389;Melbourne Southern Cross Railway Station;Melbourne;Australia;;\N;-37.8181602;144.9533883;111;10;O;Australia/Hobart
8390;Nambour Railway Station;Nambour;Australia;;\N;-26.617353;152.973785;400;10;N;Australia/Brisbane
8391;Bundaberg Railway Station;Bundaberg;Australia;;\N;-24.867837;152.349416;100;10;U;Australia/Brisbane
8392;Hung Hom Railway Station;Hong Kong;Hong Kong;;\N;22.308813;114.169624;100;8;N;Asia/Hong_Kong
8393;Hong Kong Railway Station;Hong Kong;Hong Kong;;\N;22.2846291082599;114.158227443695;100;8;N;Asia/Hong_Kong
8394;Shelby County Airport;Alabaster;United States;EET;KEET;33.1777778;-86.7832222;586;-6;A;America/Chicago
8395;Yuendumu ;Yuendumu ;Australia;YUE;YYND;-22.254167;131.781944;2205;9.5;O;Australia/Darwin
8396;Gare de Strasbourg;Strasbourg;France;XWG;\N;48.585068;7.734547;475;1;E;Europe/Paris
8397;Sky Ranch at Carefree;Carefree;United States;;18AZ;33.8180947;-111.8979242;2568;-7;N;America/Phoenix
8398;St Peter and St Paul Fortress;St. Petersburg;Russia;;\N;59.951443;30.313831;30;4;N;Europe/Moscow
8399;Guangzhou South Railway Station;Guangzhou;China;GZS;\N;22.990081;113.270631;1;8;U;Asia/Chongqing
8400;Oddballs Camp Private Airstrip;Okavango Oddballs Camp;Botswana;;\N;-19.531281;23.091847;3133;2;N;Africa/Gaborone
8401;Lombok International Airport;Praya;Indonesia;LOP;WADL;-8.7573222;116.276675;52;8;N;Asia/Makassar
8402;Marmul;Marmul;Oman;OMM;OONR;18.133333;55.266666;900;4;N;Asia/Muscat
8403;One Hundred Mile House Airport;One Hundred Mile House;Canada;;CAV3;51.3833;-121.1825;3055;-8;A;America/Vancouver
8404;South Cariboo Regional Airport;108 Mile Ranch;Canada;ZML;CZML;51.4412;-121.1958;3129;-8;A;America/Vancouver
8405;Glasgow City Heliport;Glasgow;United Kingdom;;EGEG;55.5141;-4.1749;10;0;E;Europe/London
8406;Yarram;Yarram;Australia;;YYRM;-38.56;146.75;200;10;O;Australia/Hobart
8407;Hebei Handan Airport;Handan;China;HDG;ZBHD;36.524;114.43;0;8;N;Asia/Chongqing
8408;Jade Mountain Helipad;Soufriere;Saint Lucia;;\N;13.868921;-61.073596;341;-4;N;America/St_Lucia
8409;Indianapolis Metropolitan Airport;Indianapolis;United States;UMP;KUMP;39.935278;-86.045;811;-5;A;America/New_York
8410;London-Corbin Airport-MaGee Field;London;United States;LOZ;KLOZ;37.0868889;-84.0773889;1212;-5;A;America/New_York
8411;Camden Station;Baltimore;United States;;\N;39.2837181;-76.621634;49;-5;A;America/New_York
8412;Fredericksburg Amtrak Station;Fredericksburg;United States;FBG;\N;38.298416;-77.456875;130;-5;A;America/New_York
8413;Pamalican Airstrip;Cuyo;Philippines;;\N;11.354499;120.730991;0;8;N;Asia/Manila
8414;Warsaw Modlin;Warsaw;Poland;WMI;EPMO;52.451111;20.651667;341;1;E;Europe/Warsaw
8415;Buffalo Bus Terminal;Buffalo;United States;;BBUF;42.883512;-78.87205;0;-5;A;America/New_York
8416;Twelve Apostles Heliport;Princetown;Australia;;\N;-38.665833;143.104444;165;10;U;Australia/Hobart
8417;Jixi Airport;Jixi;China;JXA;ZYJX;45.30611;130.99667;0;8;N;Asia/Chongqing
8418;King St Station;Seattle;United States;;SEAT;47.5985;-122.3299;7;-8;A;America/Los_Angeles
8422;Palatka Amtrak Station;Palatka;United States;;\N;29.649178;-81.640456;16;-5;A;America/New_York
8423;Gimli Industrial Park Airport;Gimli;Canada;YGM;CYGM;50.628056;-97.043333;753;-6;A;America/Winnipeg
8424;Matheson Island Airport;Matheson Island;Canada;;\N;51.732222;-96.934444;725;-6;A;America/Winnipeg
8425;Matheson Island Airport;Matheson Island;Canada;;\N;51.732222;-96.934444;725;-6;A;America/Winnipeg
8426;Matheson Island Airport;Matheson Island;Canada;;CJT2;51.732222;-96.934444;725;-6;A;America/Winnipeg
8427;Tura;Tura;Russia;;UNIT;64.33;100.43;400;8;N;Asia/Krasnoyarsk
8428;Beloyarsky;Beloyarsky;Russia;EYK;USHQ;63.683056;66.683056;25;6;E;Asia/Yekaterinburg
8429;Ypenburg;Ypenburg;Netherlands;;EHYB;52.040376;4.365771;0;1;E;Europe/Amsterdam
8430;John H. Batten Airport;Racine;United States;RAC;KRAC;42.7605;-87.8152;674;-6;A;America/Chicago
8431;Taytay Sandoval;Taytay;Philippines;RZP;RPSD;11.05167;119.5183;75;8;U;Asia/Manila
8432;Gera Leumnitz;Gera;Germany;;EDAJ;50.8814;12.1322;100;1;E;Europe/Berlin
8433;Gera Leumnitz;Gera;Germany;;\N;50.880023;12.132411;100;1;E;Europe/Berlin
8434;Mira Flores;Mira Flores;Colombia;;\N;1.444549;-71.949921;300;-5;S;America/Bogota
8435;Kasimovo;Saint-Petersburg;Russia;;XLLN;60.200059;30.334926;230;4;N;Europe/Moscow
8436;Otaru;Otaru;Japan;QOT;\N;43.1894444;141.0022222;0;9;N;Asia/Tokyo
8437;Noboribetsu;Noboribetsu;Japan;;\N;42.46867;141.03833;361;9;N;Asia/Tokyo
8438;Terlet;Arnhem;Netherlands;;EHTL;52.057222;5.924444;276;1;E;Europe/Amsterdam
8439;Schaffen Diest;Schaffen;Belgium;;\N;51.002846;5.060649;100;1;E;Europe/Brussels
8440;Blomberg 3GGW;Blomberg;Germany;;\N;51.930401;9.094285;200;1;E;Europe/Berlin
8441;Buckeburg;Buckeburg;Germany;;\N;52.279764;9.079843;200;1;E;Europe/Berlin
8442;Shigatse Peace Airport;Shigatse;China;RKZ;ZURK;29.351667;89.306944;12408;8;N;Asia/Chongqing
8443;Redlands Municipal Airport;Redlands;United States;REI;KREI;34.08525;-117.146388;1574;-8;A;America/Los_Angeles
8444;Chemehuevi Valley;Chemehuevi Valley;United States;49X;\N;34.528889;-114.431971;638;-8;A;America/Los_Angeles
8445;Flabob Airport;Riverside;United States;RIR;KRIR;33.988778;-117.409971;764;-8;A;America/Los_Angeles
8446;Tacoma Narrows Airport;Tacoma;United States;TIW;KTIW;47.267944;-122.57811;294;-8;A;America/Los_Angeles
8447;Tampa North Aero Park;Tampa;United States;X39;\N;28.221278;-82.374555;68;-5;A;America/New_York
8448;Athuruga;Athuruga;Maldives;;\N;3;72;0;5;N;Indian/Maldives
8449;Athuruga;Athuruga;Maldives;;\N;3;72;0;5;U;Indian/Maldives
8450;Athuruga;Athuruga;Maldives;;\N;3;72;0;5;U;Indian/Maldives
8451;STROE KAZERNE;STROE;Netherlands;;\N;52.198138;5.705802;10;1;E;Europe/Amsterdam
8452;\\'t Harde;Oldebroek;Netherlands;;\N;52.406503;5.898492;10;1;E;Europe/Amsterdam
8453;WERL KAZERNE;WERL;Germany;;\N;51.566907;7.915326;100;1;E;Europe/Berlin
8454;SZEGED;SZEGED;Hungary;;\N;46.252761;20.096519;100;1;E;Europe/Budapest
8455;LAKE WANNSEE;BERLIN;Germany;;\N;52.437517;13.175561;0;1;E;Europe/Berlin
8456;Tenkiller Lake Airpark;COOKSON;United States;;\N;35.705541;-94.933605;150;-6;U;America/Chicago
8457;HOEVENEN;HOEVENEN;Belgium;;\N;51.304111;4.388766;50;1;E;Europe/Brussels
8458;GRAS;HARREVELD;Netherlands;;\N;51.981127;6.516252;50;1;E;Europe/Amsterdam
8459;Muenster-Telgte;TELGTE;Germany;;EDLO;51.944556;7.773292;100;1;E;Europe/Berlin
8460;Jack Edwards Airport;Gulf Shores;United States;JKA;KJKA;30.2896389;-87.6717778;17;-6;A;America/Chicago
8461;General Villamil Airport;Isabela;Ecuador;;SEII;-0.942499999;-90.95305555;35;-6;N;Pacific/Galapagos
8462;Balzers Heliport;Balzers;Switzerland;;LSXB;47.076667;9.5;1585;1;E;Europe/Vaduz
8463;Vieste Heliport;Vieste;Italy;VIF;\N;41.885277;16.176945;46;1;E;Europe/Rome
8464;Alpha;Cork;Ireland;;\N;51.400377;-7.901464;100;0;U;\N
8465;Gogld;Island of Gogland;Russia;;\N;60.011662;27.008236;10;3;U;\N
8466;HIBER;Offshore;Mexico;;\N;19.398611;-92.016111;100;-6;U;\N
8467;Punitz-Guessing;Punitz-Guessing;Austria;;LOGG;47.14833;16.31833;950;1;E;Europe/Vienna
8468;Khmeinitskiy;Khmeinitskiy;Ukraine;HMJ;UKLH;49.366;26.933;351;2;E;Europe/Kiev
8469;Hiroshima-Nishi;Hiroshima;Japan;HIW;RJBH;34.367;132.408;9;9;N;Asia/Tokyo
8470;Annapolis;Annapolis;United States;;ANNA;38.972944;-76.501158;0;-5;A;America/New_York
8471;High River Regional Airport;High River;Canada;;\N;50.3333;-113.833;3431;-7;A;America/Edmonton
8472;Yalata;Yalata;Australia;;YYTA;-31.29;131.51;100;9.5;O;Australia/Adelaide
8473;Hazleton Municipal;Hazleton;United States;HZL;KHZL;40.989167;-76.0025;1603;-5;A;America/New_York
8474;Greater Cumberland Rgnl.;Cumberland;United States;CBE;KCBE;39.615278;-78.760556;775;-5;A;America/New_York
8475;Sugar Loaf Shores Airport;Key West;United States;;7FA1;24.6487544;-81.5798083;4;-5;A;America/New_York
8476;Wyndham;Wyndham;Australia;;YWYM;-15.2857;128.0722;100;8;U;Australia/Perth
8477;Bob Quinn Lake;Bob Quinn Lake;Canada;YBO;CBW4;56.58;-130.5;50;-8;U;America/Vancouver
8478;Songo Songo;Songo Songo Island;Tanzania;;\N;-8.528056;39.506944;1;3;U;Africa/Dar_es_Salaam
8479;Jongomero Camp;Ruaha National Park;Tanzania;;\N;-7.895;34.580853;0;3;U;Africa/Dar_es_Salaam
8480;Msembe Airstrip;Msembe;Tanzania;;\N;-7.686111;35.926944;0;3;U;Africa/Dar_es_Salaam
8481;Fox Glacier Airstrip;Fox Glacier;New Zealand;FGL;\N;-43.4;170.033;475;12;Z;Pacific/Auckland
8482;Grabtsevo;Kaluga;Russia;KLF;UUBS;54.55;36.367;600;4;E;Europe/Moscow
8483;Reykjahlid;Reykjahlid;Iceland;;\N;65.6;-17;100;0;U;Atlantic/Reykjavik
8484;Central Station;Stockholm;Sweden;XEV;\N;59.33;18.058056;10;1;E;Europe/Stockholm
8485;Central Station;Uppsala;Sweden;QYX;\N;59.858333;17.646111;23;1;E;Europe/Stockholm
8486;Centraal;Rotterdam;Netherlands;QRH;\N;51.924444;4.469444;7;1;E;Europe/Amsterdam
8487;Holesovice;Praha;Czech Republic;XYJ;\N;50.11;14.439444;616;1;E;Europe/Prague
8488;Hauptbahnhof;Bonn;Germany;BNJ;\N;50.731944;7.096944;204;1;E;Europe/Berlin
8489;Central;Copenhagen;Denmark;ZGH;\N;55.672778;12.564444;16;1;E;Europe/Copenhagen
8490;Hauptbahnhof;Salzburg;Austria;ZSB;\N;47.813056;13.046667;1391;1;E;Europe/Vienna
8491;Centraal;Antwerp;Belgium;ZWE;\N;51.217222;4.421111;51;1;E;Europe/Brussels
8492;Ellisras;Lephalale;South Africa;ELL;FAEA;-23.666667;27.75;2700;2;U;Africa/Johannesburg
8493;Tri-County Regional Airport;Lone Rock;United States;LNR;KLNR;43.211667;-90.181667;718;-6;A;America/Chicago
8494;Price County Airport;Phillips;United States;;KPBH;45.705;-90.403056;1466;-6;A;America/Chicago
8495;Municipal;Monroe;United States;;KEFT;42.616389;-89.591944;1079;-6;A;America/Chicago
8496;Regional Airport;Joliet;United States;JOT;KJOT;41.517778;-88.175556;582;-6;A;America/Chicago
8497;Illinois Valley Regional;Peru;United States;VYS;KVYS;41.351944;-89.153056;654;-6;A;America/Chicago
8498;Dirk Hartog Island;Dirk Hartog Island;Australia;;YDHD;-26.00937;113.160008;15;8;O;Australia/Perth
8499;Reynolds Field;Jackson;United States;JXN;KJXN;42.260556;-84.460556;1001;-5;A;America/New_York
8500;Flugplatz;Furstenwalde;Germany;;EDAL;52.393889;14.099722;141;1;E;Europe/Berlin
8501;Flugplatz Finow;Eberswalde;Germany;;EDAV;52.827222;13.693611;121;1;E;Europe/Berlin
8502;Joseph A. Hardy Airport;Connellsville;United States;;KVVS;39.959167;-79.657222;1267;-5;A;America/New_York
8503;Bedford County;Bedford;United States;;KHMZ;40.086111;-78.513333;1162;-5;A;America/New_York
8504;Wings Field;Philadelphia;United States;BBX;KLOM;40.1375;-75.265;302;-5;A;America/New_York
8505;County;Okeechobee;United States;OBE;KOBE;27.265833;-80.851111;34;-5;A;America/New_York
8506;Regional - Hendricks AAF;Sebring;United States;SEF;KSEF;27.456389;-81.342222;63;-5;A;America/New_York
8507;Executive;Avon Park;United States;AVO;KAVO;27.591389;-81.528889;160;-5;A;America/New_York
8508;Gilbert Airport;Winter Haven;United States;GIF;KGIF;28.062778;-81.753333;145;-5;A;America/New_York
8509;Municipal Airport;Zephyrhills;United States;ZPH;KZPH;28.228056;-82.155833;90;-5;A;America/New_York
8510;International Airport;Ocala;United States;OCF;KOCF;29.1725;-82.224167;89;-5;A;America/New_York
8511;Jesup-Wayne County Airport;Jesup;United States;JES;KJES;31.553889;-81.8825;107;-5;A;America/New_York
8512;Madison GA Municipal Airport;Madison;United States;52A;K52A;33.612125;-83.4604444;694;-5;A;America/New_York
8513;Coweta County Airport;Newnan;United States;CCO;KCCO;33.3115656;-84.7697554;970;-5;A;America/New_York
8514;McDuffie County Airport;Thomson;United States;HQU;KHQU;33.5297315;-82.5169509;501;-5;A;America/New_York
8515;Municipal Airport;Aiken;United States;AIK;KAIK;33.6493889;-81.6850278;529;-5;A;America/New_York
8516;Woodward Field;Camden;United States;CDN;KCDN;34.2835833;-80.5648611;302;-5;A;America/New_York
8517;Municipal Airport;Lumberton;United States;LBT;KLBT;34.6098056;-79.0595556;125;-5;A;America/New_York
8518;Ridgeland Airport;Ridgeland;United States;3J1;\N;32.4934167;-80.99175;79;-5;A;America/New_York
8519;Moore County Airport;Pinehurst-Southern Pines;United States;SOP;KSOP;35.2376111;-79.3887958;455;-5;A;America/New_York
8520;Richmond County Airport;Rockingham;United States;RCZ;KRCZ;34.8913056;-79.7596111;358;-5;A;America/New_York
8521;Bamberg County Airport;Bamberg;United States;99N;\N;33.3045278;-81.1084167;231;-5;A;America/New_York
8522;Richland Airport;Richland Center;United States;93C;\N;43.2833575;-90.2982819;742;-6;A;America/Chicago
8523;Municipal Airport;Viroqua;United States;Y51;\N;43.5793603;-90.8964742;1292;-6;A;America/Chicago
8524;Baraboo Wisconsin Dells Airport;Baraboo;United States;DLL;KDLL;43.5217843;-89.7709266;979;-6;A;America/Chicago
8525;Foster Field;Apple River;United States;7A4;\N;42.4664444;-90.1693889;990;-6;A;America/Chicago
8526;Regional Airport;Statesville;United States;SVH;KSVH;35.7649958;-80.9538958;968;-5;A;America/New_York
8527;Sylvania Airport;Sturtevant;United States;C89;\N;42.70325;-87.9589722;785;-6;A;America/Chicago
8528;Municipal Airport;Burlington;United States;BUU;KBUU;42.6907171;-88.3046825;780;-6;A;America/Chicago
8529;Stroudsburg-Pocono Airport;East Stroudsburg;United States;N53;\N;41.0358717;-75.1606789;480;-5;A;America/New_York
8530;Spring Hill Airport;Sterling;United States;70N;\N;41.3473569;-75.4158972;1729;-5;A;America/New_York
8531;Randall Airport;Middletown;United States;06N;\N;41.431912;-74.3915611;523;-5;A;America/New_York
8532;William T. Piper Mem.;Lock Haven;United States;LHV;KLHV;41.1357778;-77.4223056;556;-5;A;America/New_York
8533;Grove City Airport;Grove City;United States;29D;\N;41.1460278;-80.16775;1371;-5;A;America/New_York
8534;Lansdowne Airport;Youngstown;United States;04G;\N;41.1304722;-80.6195833;1044;-5;A;America/New_York
8535;Wadsworth Municipal;Wadsworth;United States;3G3;\N;41.0031572;-81.7564401;974;-5;A;America/New_York
8536;Ashland County Airport;Ashland;United States;3G4;\N;40.9029722;-82.2556389;1206;-5;A;America/New_York
8537;Pittsburgh-Monroeville Airport;Monroeville;United States;4G0;\N;40.4526389;-79.7749167;1187;-5;A;America/New_York
8538;Zelienople Municipal Airport;Zelienople;United States;;KPJC;40.8019722;-80.1608611;898;-5;A;America/New_York
8539;Somerset County Airport;Somerset;United States;2G9;\N;40.0388708;-79.0149951;2275;-5;A;America/New_York
8540;Youngstown Elser Metro Airport;Youngstown;United States;4G4;\N;40.9617953;-80.6773264;1070;-5;A;America/New_York
8541;Braceville Airport;Braceville;United States;41N;\N;41.2111675;-80.9692572;900;-5;A;America/New_York
8542;Lorain County Regional Airport;Lorain-Elyria;United States;LPR;KLPR;41.3442778;-82.1776389;793;-5;A;America/New_York
8543;Germack Airport;Geneva;United States;7D9;\N;41.7778322;-80.9039797;820;-5;A;America/New_York
8544;Burke Lakefront Airport;Cleveland;United States;BKL;KBKL;41.5175;-81.683333;583;-5;A;America/New_York
8545;Chautauqua County-Dunkirk Airport;Dunkirk;United States;DKK;KDKK;42.4933353;-79.2720417;693;-5;A;America/New_York
8546;Hamburg Inc Airport;Hamburg;United States;4G2;\N;42.7008925;-78.9147569;751;-5;A;America/New_York
8547;Trenton-Robbinsville Airport;Trenton;United States;N87;\N;40.2139444;-74.6017778;118;-5;A;America/New_York
8548;South Jersey Regional Airport;Mount Holly;United States;VAY;KVAY;39.9428889;-74.84575;53;-5;A;America/New_York
8549;Spitfire Aerodrome;Pedricktown;United States;7N7;\N;39.7355633;-75.3977211;40;-5;A;America/New_York
8550;Linden Airport;Linden;United States;LDJ;KLDJ;40.6174472;-74.2445942;23;-5;A;America/New_York
8551;Morgantown Airport;Morgantown;United States;O03;\N;40.1570414;-75.8704892;600;-5;A;America/New_York
8552;Harford County Airport;Churchville;United States;0W3;\N;39.5668378;-76.2024028;409;-5;A;America/New_York
8553;Tri-State Steuben County Airport;Angola;United States;ANQ;KANQ;41.6396983;-85.0834933;995;-5;A;America/New_York
8554;Plymouth Municipal Airport;Plymouth;United States;C65;\N;41.3651307;-86.3002574;800;-5;A;America/New_York
8555;Warsaw Municipal Airport;Warsaw;United States;;KASW;41.2746944;-85.8400556;850;-5;A;America/New_York
8556;Van Wert County Airport;Van Wert;United States;VNW;KVNW;40.8638295;-84.606358;787;-5;A;America/New_York
8557;Port Bucyrus-Crawford County Airport;Bucyrus;United States;17G;\N;40.7815556;-82.9748056;1003;-5;A;America/New_York
8558;Lake Wales Municipal Airport;Lake Wales;United States;X07;\N;27.8938056;-81.6203889;127;-5;A;America/New_York
8559;Brooks Field Airport;Marshall;United States;RMY;KRMY;42.2511932;-84.9554443;941;-5;A;America/New_York
8560;Genesee County Airport;Batavia;United States;GVQ;KGVQ;43.03175;-78.1696667;914;-5;A;America/New_York
8561;Finger Lakes Regional Airport;Seneca Falls;United States;0G7;\N;42.8835647;-76.7812318;492;-5;A;America/New_York
8562;Stormville Airport;Stormville;United States;N69;\N;41.5769708;-73.7323514;358;-5;A;America/New_York
8563;Robertson Field;Plainville;United States;4B8;\N;41.6893333;-72.8646944;202;-5;A;America/New_York
8564;Williams County Airport;Bryan;United States;0G6;\N;41.4673056;-84.5067778;730;-5;A;America/New_York
8565;Clearwater Air Park;Clearwater;United States;CLW;KCLW;27.9764722;-82.7586667;70;-5;A;America/New_York
8566;South Lakeland Airport;Lakeland;United States;X49;\N;27.9333581;-82.0439739;110;-5;A;America/New_York
8567;Scatsta Airport;Scatsta;United Kingdom;SCS;\N;60.433;-1.3;0;0;U;Europe/London
8568;North Sea;Buchan Alpha;United Kingdom;;\N;57.9039;0.0319;0;0;U;\N
8569;North Sea;Buchan Alpha;United Kingdom;;\N;57.9039;0.0319;0;0;U;\N
8570;North Sea;Claymore;United Kingdom;;CLAA;58.4494;-0.2536;0;0;U;\N
8571;North Sea;FPSO Anasuria;United Kingdom;;ANAS;57.1524;0.0483;0;0;U;\N
8572;North Sea;BP Bruce;United Kingdom;;BRU1;58.44;2.33;0;0;U;\N
8573;North Sea;Buchan Alpha;United Kingdom;;BUCA;57.9039;0.0319;0;0;U;\N
8574;North Sea;Buzzard;United Kingdom;;BUZZ;56.8;0.766;0;0;U;\N
8575;North Sea;Clyde;United Kingdom;;CLY1;56.45;2.2833;0;0;U;\N
8576;North Sea;Cormorant Alpha;United Kingdom;;COA1;61.2333;-1.15;0;0;U;\N
8577;North Sea;Forties Alpha;United Kingdom;;FA08;57.7167;-1.0167;0;0;U;\N
8578;North Sea;Forties Charlie;United Kingdom;;FC08;57.7167;-1.0167;0;0;U;\N
8579;North Sea;Fulmar;United Kingdom;;FUA1;56.48;2.13;0;0;U;\N
8580;North Sea;Gannet Alpha;United Kingdom;;GAA1;57.3;1.2;0;0;U;\N
8581;North Sea;Montrose Alpha;United Kingdom;;MOA1;57.4506;1.3883;0;0;U;\N
8582;North Sea;Tern;United Kingdom;;TEA1;61.3333;0.8333;0;0;U;\N
8583;Atlantic Ocean;FPSO Xikomba;Angola;;XKSO;-6.0333;11.0166;0;0;U;\N
8584;Zoo;Berlin;Germany;QWC;\N;52.507222;13.3325;120;1;E;Europe/Berlin
8585;Frankfurt Oder Hbf;Frankfurt Oder;Germany;ZFR;\N;52.336944;14.547222;177;1;E;Europe/Berlin
8586;Gare de Lyon;Paris;France;PLY;\N;48.844722;2.373611;129;1;E;Europe/Paris
8587;Gare de LEst;Paris;France;XHP;\N;48.876944;2.359167;149;1;E;Europe/Paris
8588;All Airports;Paris;France;PAR;\N;48.856389;2.352222;107;1;E;Europe/Paris
8589;Ostbahnhof;Berlin;Germany;BHF;\N;52.51;13.434722;135;1;E;Europe/Berlin
8590;All Airports;London;United Kingdom;LON;\N;51.508056;-0.127778;66;0;E;Europe/London
8591;All Airports;New York;United States;NYC;\N;40.714167;-74.005833;31;-5;A;America/New_York
8592;All Airports;Chicago;United States;CHI;\N;41.883611;-87.631667;596;-6;A;America/Chicago
8593;Meigs Field;Chicago;United States;CGX;KCGX;41.860278;-87.609722;585;-6;A;America/Chicago
8594;All Airports;Tokyo;Japan;TYO;\N;35.689444;139.691667;123;9;U;Asia/Tokyo
8595;All Airports;Beijing;China;BJS;\N;39.904167;116.407222;171;8;U;Asia/Chongqing
8596;All Airports;Milan;Italy;MIL;\N;45.463611;9.188056;402;1;E;Europe/Rome
8597;All Airports;Washington;United States;WAS;\N;38.889444;-77.035278;25;-5;A;America/New_York
8598;All Airports;Montreal;Canada;YMQ;\N;45.508611;-73.553889;53;-5;A;America/Toronto
8599;All Airports;Toronto;Canada;YTO;\N;43.653056;-79.383056;302;-5;A;America/Toronto
8600;Longhua Airport;Shanghai;China;;ZSSL;31.166944;121.453611;14;8;U;Asia/Chongqing
8601;Pickens County Airport;Jasper;United States;JZP;KJZP;34.4534722;-84.4572222;1535;-5;A;America/New_York
8602;Gesundbrunnen;Berlin;Germany;BGS;\N;52.548611;13.389444;150;1;E;Europe/Berlin
8603;Saarmund Airport;Saarmund;Germany;;EDCS;52.3089943;13.1006727;160;1;E;Europe/Berlin
8604;Grand Strand Airport;North Myrtle Beach;United States;CRE;\N;33.81175;-78.7239444;31;-5;A;America/New_York
8605;Sun Moon Lake Airport;Sun Moon Lake;Taiwan;SMT;\N;23.8833;120.933;2405;8;U;Asia/Taipei
8606;Lansing Municipal;Lansing;United States;IGQ;KIGQ;41.5349167;-87.5295278;620;-6;A;America/Chicago
8607;Bloyer Field;Tomah;United States;Y72;\N;43.9762222;-90.4806111;966;-6;A;America/Chicago
8608;Ramona Airport;Ramona;United States;RNM;KRNM;33.0391667;-116.91525;1395;-8;A;America/Los_Angeles
8609;Dole-Pont Sur Yonne Airport;Pont Sur Yonne;France;;LFGO;48.288333;3.246389;223;1;E;Europe/Paris
8610;Dole-St Florentin Cheu Airport;St Florentin Cheu;France;;LFGP;47.981389;3.776389;345;1;E;Europe/Paris
8611;Guiscriff-Saulieu Liernais Airport;Saulieu;France;;LFEW;47.238611;4.263056;1699;1;E;Europe/Paris
8612;Olten Airport;Olten;Switzerland;;LSPO;47.343611;7.8875;1370;1;E;Europe/Zurich
8613;Buochs Airport;Buochs;Switzerland;BXO;LSZC;46.9725;8.398611;1464;1;E;Europe/Zurich
8614;Ambri Airport;Quinto;Switzerland;;LSPM;46.513056;8.689444;3214;1;E;Europe/Zurich
8615;Lodrino Airport;Lodrino;Switzerland;;LSML;46.29499;8.9916;850;1;E;Europe/Zurich
8616;Roudnice Airport;Roudnice nad Lebem;Czech Republic;;LKRO;50.407222;14.232778;719;1;E;Europe/Prague
8617;Usti Nad Labem Airport;Usti Nad Labem;Czech Republic;;LKUL;50.698889;13.970833;765;1;E;Europe/Prague
8618;Mauterndorf Airport;Mauterndorf;Austria;;LOSM;47.131667;13.695278;3629;1;E;Europe/Vienna
8619;Notsch Im Gailtal Airport;Notsch Im Gailtal;Austria;;LOKN;46.580833;13.629167;1810;1;E;Europe/Vienna
8620;Forchheim Airport;Karlsruhe;Germany;;EDTK;49.097778;8.453056;400;1;E;Europe/Berlin
8621;Bergstrasse Airport;Weinheim;Germany;;EDGZ;49.567222;8.611389;319;1;E;Europe/Berlin
8622;Creil-Meaux Esbly Airport;Meaux;France;;LFPE;48.928611;2.833056;210;1;E;Europe/Paris
8623;Belleau Airport;Chateau-Thierry;France;;LFFH;49.066944;3.356111;719;1;E;Europe/Paris
8624;Branch County Memorial Airport;Coldwater;United States;OEB;KOEB;41.9335691;-85.0522935;959;-5;A;America/New_York
8625;Wilkes-Barre Wyoming Valley Airport;Wilkes-Barre;United States;WBW;KWBW;41.2973074;-75.8522405;543;-5;A;America/New_York
8626;Lost Nation Municipal Airport;Willoughby;United States;LNN;KLNN;41.6840278;-81.38975;626;-5;A;America/New_York
8627;Taoyuan Air Base;Taoyuan;Taiwan;;RCGM;25.05660057;121.2442627;143;8;U;Asia/Taipei
8628;Uummannaq Heliport;Uummannaq;Greenland;UMD;BGUM;70.4047;52.0702;5;-4;E;\N
8629;Civitavecchia;Civitavecchia;Italy;;\N;42.6;11.48;0;1;E;Europe/Rome
8630;Bayannur;Bayannur;China;RLK;\N;40.926389;107.738889;3389;8;U;Asia/Chongqing
8631;Capital City Airport;Frankfort;United States;FFT;KFFT;38.1819722;-84.9061389;812;-5;A;America/New_York
8632;Lewiston Maine;Lewiston;United States;LEW;KLEW;44.0484728;-70.2835075;288;-5;A;America/New_York
8633;Inowroclaw Inowr Airport;Inowroclaw;Poland;;EPIR;52.829444;18.330556;259;1;E;Europe/Warsaw
8634;Pruszcz Gdansk Airport;Pruszcz Gdansk;Poland;;EPPR;54.248056;18.671667;21;1;E;Europe/Warsaw
8635;Florence;Florence;United States;6S2;\N;43.9828168;-124.1113687;51;-8;A;America/Los_Angeles
8636;Martin Campbell Field Airport;Copperhead;United States;1A3;\N;35.0158056;-84.3468333;1789;-5;A;America/New_York
8637;Sudkreuz;Berlin;Germany;;BPAP;52.475556;13.364444;136;1;E;Europe/Berlin
8638;Naval Air Station;Glenview;United States;NBU;KNBU;42.090556;-87.8225;653;-6;A;America/Chicago
8639;Bloyer Field;Tomah;United States;;KY72;43.975278;-90.483333;965;-6;A;America/Chicago
8640;Yongning Air Base;Beijing;China;;CN-0;40.50310135;116.1080017;1677;8;U;Asia/Chongqing
8641;Marco Islands;Marco Island Airport;United States;MRK;KMKY;25.9950278;-81.6725278;5;-5;A;America/New_York
8642;Boipeba;Boipeba Island;Brazil;;\N;-13.565589;-38.939241;5;-3;S;America/Fortaleza
8643;Boipeba;Boipeba Island;Brazil;;\N;-13.565589;-38.939241;5;-3;S;America/Fortaleza
8644;Drummond Island Airport;Drummond Island;United States;DRM;KDRM;46.0093114;-83.7439342;668;-5;A;America/New_York
8645;Garland Airport;Lewiston;United States;8M8;\N;44.8065278;-84.2761944;1218;-5;A;America/New_York
8646;Gladwin Zettel Memorial Airport;Gladwin;United States;GDW;KGDW;43.9705893;-84.47502;774;-5;A;America/New_York
8647;Lowell City Airport;Lowell;United States;24C;\N;42.95392;-85.3439058;681;-5;A;America/New_York
8648;South Haven Area Regional Airport;South Haven;United States;LWA;KLWA;42.3511944;-86.2556389;666;-5;A;America/New_York
8649;Schaumburg Regional;Schaumburg;United States;06C;\N;41.9893408;-88.1012428;801;-6;A;America/Chicago
8650;Khulna Seaplane Landing Site;Khulna;Bangladesh;;\N;22.464897;89.351261;0;6;N;Asia/Dhaka
8651;Khulna Seaplane Landing Site;Khulna;Bangladesh;;\N;22.464897;89.351261;0;6;N;Asia/Dhaka
8652;Khulna Seaplane Landing Site;Khulna;Bangladesh;;\N;22.464897;89.351261;0;6;N;Asia/Dhaka
8653;Marshfield Municipal Airport;Marshfield;United States;MFI;KMFI;44.6368797;-90.1893267;1278;-6;A;America/Chicago
8654;Alexander Field South Wood County Airport;Wisconsin Rapids;United States;ISW;KISW;44.3606456;-89.8381412;1021;-6;A;America/Chicago
8655;Clinton Municipal;Clinton;United States;CWI;KCWI;41.83075;-90.3289722;708;-6;A;America/Chicago
8656;Wiescheid;Langenfeld;Germany;;\N;51.141663;6.987305;1000;1;E;Europe/Berlin
8657;Wadi Rum;Wadi Rum Desert;Jordan;;\N;29.682087;35.470734;1000;2;U;Asia/Amman
8658;Beverly Municipal Airport;Beverly;United States;BVY;KBVY;42.584141;-70.9161444;107;-5;A;America/New_York
8659;Priob\\'e;Priob\\'e;Russia;;\N;62.5461;65.6261;200;6;N;Asia/Yekaterinburg
8660;Nagaur Airport;Nagaur;India;;VI73;27.20829964;73.71140289;3853;5.5;U;Asia/Calcutta
8661;Ostafyevo International Airport;Moscow;Russia;;UUMO;55.511667;37.507222;568;4;N;Europe/Moscow
8662;Oakdale Airport;Oakdale;United States;O27;\N;37.7563333;-120.8001944;237;-8;A;America/Los_Angeles
8663;Trois Rivieres Airport;Trois Rivieres;Canada;YRQ;CYRQ;46.351667;-72.680556;199;-5;A;America/Toronto
8664;Poplar Bluff Municipal Airport;Poplar Bluff;United States;POF;KPOF;36.77394444;-90.3248611;331;-6;A;America/Chicago
8665;Somerset Airport;Somerville;United States;;KSMQ;40.6259908;-74.6702433;105;-5;A;America/New_York
8666;Eastport Municipal Airport;Eastport;United States;EPM;KEPM;44.9101111;-67.0126944;45;-5;A;America/New_York
8667;Keokuk Municipal Airport;Keokuk;United States;EOK;KEOK;40.4599078;-91.4285011;671;-6;A;America/Chicago
8668;Banks Airport;Swans Island;United States;ME5;\N;44.1653889;-68.4281667;100;-5;A;America/New_York
8669;Perth Scone Airport;Perth;United Kingdom;PSL;EGPT;56.439722;-3.371389;392;0;E;Europe/London
8670;Caernarfon Airport;Caernarfon;United Kingdom;;EGCK;53.104167;-4.340278;14;0;E;Europe/London
8671;Grefrath-Niershorst Airport;Grefrath;Germany;;EDLF;51.333889;6.359444;105;1;E;Europe/Berlin
8672;Offenburg;Offenburg;Germany;;\N;48.4711;7.9411;535;1;E;Europe/Berlin
8673;Oberwolfach;Oberwolfach;Germany;;\N;48.3167;8.2167;1060;1;E;Europe/Berlin
8674;Karlsruhe Hbf;Karlsruhe;Germany;;\N;49.0135;8.4044;377;1;E;Europe/Berlin
8675;Lichtenwalde;Lichtenwalde;Germany;;\N;50.8833;13.0167;984;1;E;Europe/Berlin
8676;Karlsruhe;Karlsruhe;Germany;;\N;49.0135;8.4044;377;1;E;Europe/Berlin
8677;Holzhau;Holzhau;Germany;;\N;50.7333334;13.6;1969;1;E;Europe/Berlin
8678;Porto;Porto;Portugal;;\N;41.1667;8.5833;341;1;E;Europe/Rome
8679;Lisboa-Oriente;Lisboa;Portugal;;\N;38.424972;9.82179;7;0;E;\N
8680;Bad Herrenalb;Bad Herrenalb;Germany;;\N;48.7976817;8.4361503;1198;1;E;Europe/Berlin
8681;St. Paul Downtown Airport - Holman Field;St. Paul;United States;;KSTP;44.9346225;-93.0603424;705;-6;A;America/Chicago
8682;Soderhamn Airport;Soderhamn;Sweden;SOO;ESCL;61.261667;17.098056;94;1;E;Europe/Stockholm
8683;Newcastle Airfield;Newcastle;Ireland;;EINC;53.073056;-6.039722;14;0;E;Europe/Dublin
8684;Saravane Airport;Saravane;Laos;VNA;VLSV;15.709444;106.411667;612;7;U;Asia/Vientiane
8685;Lisboa-Oriente;Lisboa;Portugal;;\N;38.7;9.1833;6;0;E;\N
8686;Lisboa-Oriente;Lisboa-Oriente;Portugal;;\N;38.7;-9.1833;6;0;E;Europe/Lisbon
8687;Zurich;Zurich;Switzerland;;\N;47.369;8.538;1339;1;E;Europe/Zurich
8688;Bagdad Airport;Bagdad;United States;E51;\N;34.5958528;-113.170195;4183;-7;A;America/Phoenix
8689;Segeletz Airport;Segeletz;Germany;;EDAI;52.826667;12.541944;141;1;E;Europe/Berlin
8690;Fuentemilanos Airport;Segovia;Spain;;LEFM;40.889233;-4.239478;3284;1;E;Europe/Madrid
8691;Akeno Aero;Akeno;Japan;;RJOE;34.533333;136.672222;12;9;U;Asia/Tokyo
8692;Ust Kamchatsk Airport;Ust Kamchatsk;Russia;;UHPK;56.238556;162.688851;200;12;N;Asia/Magadan
8693;Oconomowoc Airport;Oconomowoc;United States;;0WI8;43.1388947;-88.4723206;885;-6;A;America/Chicago
8694;Kozyrevsk Airport;Kozyrevsk;Russia;;UHPO;56.09;159.876663;331;12;N;Asia/Magadan
8695;Dikson Airport;Dikson;Russia;DKS;UODD;73.517807;80.379669;47;8;N;Asia/Krasnoyarsk
8696;Beverley Airport;Mine Site;Australia;;YBEE;-30.19033959;139.5554942;381;9.5;O;Australia/Adelaide
8697;Bogande Airport;Bogande;Canada;XBG;\N;48.329498;-70.995697;531;-5;A;America/Toronto
8698;Bantry Aerodrome;Bantry;Ireland;BYT;EIBN;51.668598;-9.48417;7;0;E;Europe/Dublin
8699;Reeroe Airport;Caherciveen;Ireland;CHE;\N;51.933334;-10.233333;314;0;E;Europe/Dublin
8700;Aliwal North;Aliwal North;South Africa;;FAAN;-30.683333;26.733333;4361;2;U;Africa/Johannesburg
8701;Alkantpan Airport;Alkantpan;South Africa;;FACO;-29.906389;22.316667;3585;2;U;Africa/Johannesburg
8702;Alldays Airport;Alldays;South Africa;ADY;FAAL;-22.678889;29.055278;2676;2;U;Africa/Johannesburg
8703;Barberton Airport;Barberton;South Africa;;FABR;-25.7175;30.975;2260;2;U;Africa/Johannesburg
8704;Ocean Ridge Airport;Gualala;United States;E55;\N;38.8016111;-123.5306389;940;-8;A;America/Los_Angeles
8705;Kent State Airport;Kent;United States;1G3;\N;41.1513889;-81.4151111;1134;-5;A;America/New_York
8706;Bulter County Regional Airport;Hamilton;United States;;KHAO;39.36375;-84.5219444;633;-5;A;America/New_York
8707;Bungoma Airport;Bungoma;Kenya;;HKBU;0.576389;34.551389;4752;3;U;Africa/Nairobi
8708;Bura East;Bura;Kenya;;HKBR;-1.185;39.813889;295;3;U;Africa/Nairobi
8709;Busia;Busia;Kenya;;HKBA;0.272;34.748;3989;3;U;Africa/Nairobi
8710;Embu;Embu;Kenya;;HKEM;-0.568889;37.492222;4131;3;U;Africa/Nairobi
8711;Garba Tula;Garba Tula;Kenya;;HKGT;0.3112;38.3056;1764;3;U;Africa/Nairobi
8712;Garissa;Garissa;Kenya;GAS;HKGA;-0.468611;39.649444;469;3;U;Africa/Nairobi
8713;Hola;Hola;Kenya;HOA;HKHO;-1.52;40.00389;220;3;U;Africa/Nairobi
8714;Homa Bay;Homa Bay;Kenya;;HKHB;-0.591668;34.48;4231;3;U;Africa/Nairobi
8715;Isiolo;Isiolo;Kenya;;HKIS;0.34361;37.587778;3688;3;U;Africa/Nairobi
8716;Kalokol;Kalokol;Kenya;KLK;HYFG;3.49;35.84;1296;3;U;Africa/Nairobi
8717;Kericho;Kericho;Kenya;KEY;HKKR;-0.385;35.245;6453;3;U;Africa/Nairobi
8718;Kilaguni;Kilaguni;Kenya;ILU;HKKL;-2.9;38.073889;2832;3;U;Africa/Nairobi
8719;Kerio Valley;Kimwarer;Kenya;KRV;\N;0.1912;35.3954;7095;3;U;Africa/Nairobi
8720;Andavadoaka;Andavadoaka;Madagascar;DVD;\N;-22.111111;43.270556;27;3;U;Indian/Antananarivo
8721;Antsirabe;Antsirabe;Madagascar;ATJ;FMME;-19.836944;47.065278;4913;3;U;Indian/Antananarivo
8722;Bekily;Bekily;Madagascar;OVA;FMSL;-24.236111;45.305;1259;3;U;Indian/Antananarivo
8723;Ust-Tsylma;Ust-Tsylma;Russia;;UUYX;65.434722232222;52.20000001;262;4;N;Europe/Moscow
8724;Atlantic City Rail Terminal;Atlantic City NJ;United States;ZRA;\N;39.3665;-74.442;8;-5;A;America/New_York
8725;Springfield Amtrak Station;Springfield MA;United States;ZSF;\N;42.106;-72.593054;65;-5;A;America/New_York
8726;Amherst Amtrak Station AMM;Amherst MA;United States;XZK;\N;42.375;-72.511389;258;-5;A;America/New_York
8727;Cape May Ferry Terminal;Cape May NJ;United States;;\N;38.9696286;-74.960918;0;-5;A;America/New_York
8728;Lewes Ferry Terminal;Lewes DE;United States;;\N;38.781655;-75.119534;0;-5;A;America/New_York
8729;Bar Harbor Yarmouth Ferry Terminal;Bar Harbor ME;United States;;\N;44.39917;-68.225;0;-5;A;America/New_York
8730;Yarmouth Ferry Terminal;Yarmouth NS;Canada;;\N;43.833227;-66.122653;0;-4;A;America/Halifax
8731;Gare de Montreux Railway Station;Montreux;Switzerland;;\N;46.435833;6.910278;1280;1;E;Europe/Zurich
8732;London - Kings Cross;London;United Kingdom;QQK;\N;51.5326;0.1233;72;0;E;Europe/London
8733;Stevenage Railway Station;Stevenage;United Kingdom;XVJ;\N;51.902;0.207;262;0;E;Europe/London
8734;Peterborough Railway Station;Peterborough;United Kingdom;XVH;\N;52.5748;-0.2502;91;0;E;Europe/London
8735;Pomalaa;Pomalaa;Indonesia;PUM;\N;-4.183333;121.61667;0;8;N;Asia/Makassar
8736;St. Ignace Ferry Dock;St. Ignace MI;United States;;\N;45.866405;-84.720254;587;-5;A;America/New_York
8737;Mackinac Island Dock;Mackinac Island MI;United States;;\N;45.848801;-84.615927;889;-5;A;America/New_York
8738;YUR'EVETS;YUR'EVETS;Russia;;UUIC;57.338582;43.0949;413;4;E;Europe/Moscow
8739;Whittlesford Parkway Rail Station;Whittlesford;United Kingdom;WLF;\N;52.104;0.166;95;0;E;Europe/London
8740;Gorno-Altaysk Airport;Gorno-Altaysk;Russia;RGK;UNBG;51.966667;85.833333;965;7;U;Asia/Omsk
8741;Gryzliny - Lansk;Gryzliny;Poland;;EPGR;53.3629;20.204;520;1;E;Europe/Warsaw
8742;Fond Du Lac County Airport;Fond du Lac;United States;FLD;KFLD;43.7711667;-88.4884167;808;-6;A;America/Chicago
8743;Waupaca Municipal Airport;Waupaca;United States;PCZ;KPCZ;44.33325;-89.0197778;840;-6;A;America/Chicago
8744;Stevens Point Municipal Airport;Stevens Point;United States;STE;KSTE;44.5451356;-89.5302844;1110;-6;A;America/Chicago
8745;Mys Shmidta Airport;Mys Shmidta;Russia;;UHMI;68.868301;179.373001;20;12;N;Asia/Magadan
8746;Moma Airport;Honuu;Russia;;UEMA;66.450859;143.261551;656;12;N;Asia/Magadan
8747;Luce County Airport;Newberry;United States;ERY;KERY;46.311199;-85.457298;869;-5;A;America/New_York
8748;Forest Lake Airport;Forest Lake;United States;25D;\N;45.2477456;-92.9943853;925;-6;A;America/Chicago
8749;Hannover Messe-Heliport;Hannover;Germany;ZVM;\N;52.329035;9.822136;152;1;E;Europe/Berlin
8750;Peenemunde Airfield;Peenemunde;Germany;PEF;EDCP;54.157778;13.772778;7;1;E;Europe/Berlin
8751;Goraszka;Goraszka;Poland;;EPGO;52.1104;21.1652;110;1;E;Europe/Warsaw
8752;Pangandaran-Java Island-Nusawiru Airport;Nusawiru;Indonesia;;WI1A;-7.72039;108.489998;16;7;N;Asia/Jakarta
8753;Pom Pom Camp;Pom Pom;Botswana;;\N;-19.35;22.5;965;2;U;Africa/Gaborone
8754;Moremi Crossing;Moremi;Botswana;;\N;-19.31;23.08;965;2;U;Africa/Gaborone
8755;New Delhi Train Station;New Delhi;India;;\N;28.64197;77.2206;708;5.5;N;Asia/Calcutta
8756;Agra Cantonment Railway Station;Agra;India;;\N;27.1578;77.98993;571;5.5;N;Asia/Calcutta
8757;Urumqi Railway Station;Urumqi;China;;\N;43.7809;87.585;2844;8;N;Asia/Chongqing
8758;Turpan Railway Station;Turpan;China;;\N;43.151;88.875;98;8;N;Asia/Chongqing
8759;Galion Municipal Airport;Galion;United States;GQQ;KGQQ;40.7533889;-82.7238056;1224;-5;A;America/New_York
8760;Musiara Airstrip;Musiara;Kenya;;HKMZ;-1.29861;35.063611;5200;3;U;Africa/Nairobi
8761;Tiputini;Tiputini;Ecuador;TPN;SETI;-0.766667;-75.53333;997;-5;S;America/Guayaquil
8762;Shell Mera;Pastaza;Ecuador;PTZ;SESM;-2;-77;7076;-5;S;America/Guayaquil
8763;Clarksville-Montgomery County Regional Airport;Clarksville;United States;CKV;KCKV;36.621944;-87.415;550;-6;A;America/Chicago
8764;Oban Rail Station;Oban;United Kingdom;;\N;56.4124;-5.4702;0;0;E;Europe/London
8765;Memmingen Rail Station;Memmingen;Germany;;\N;47.9878;10.1811;2077;1;E;Europe/Berlin
8766;Ulm Railway Station;Ulm;Germany;;\N;48.4;9.9833;1641;1;E;Europe/Berlin
8767;Aachen HBF;Aachen;Germany;;\N;50.7667;6.1;873;1;E;Europe/Berlin
8768;Liege-Guillemins Railway Station;Liege;Belgium;XHN;\N;50.6333;5.56667;659;1;E;Europe/Brussels
8769;Bruxelles-Central;Brussels;Belgium;;FBCL;50.8411;4.3564;43;1;E;Europe/Brussels
8770;Wyk auf Foehr;Wyk;Germany;OHR;\N;54.411;8.3145;24;1;E;Europe/Berlin
8771;Lompoc Airport;Lompoc;United States;LPC;KLPC;34.6656;-120.4675;88;-8;A;America/Los_Angeles
8772;Chester County G O Carlson Airport;Coatesville;United States;CTH;KMQS;39.9789;-75.8654;660;-5;A;America/New_York
8773;Bost Airport;Lashkar Gah;Afghanistan;BST;OABT;31.558889;64.364167;2464;4.5;U;Asia/Kabul
8774;Lankaran International Airport;Lankaran;Azerbaijan;LLK;UBBL;38.746389;48.817778;30;4;U;Asia/Baku
8775;Qabala Airport;Qabala;Azerbaijan;GBB;UBBQ;40.826667;47.7125;100;4;E;Asia/Baku
8776;Zaqatala International Airport;Zaqatala;Azerbaijan;ZTU;UBBY;41.562222;46.667222;100;4;E;Asia/Baku
8777;Lake Placid Airport;Lake Placid;United States;LKP;KLKP;44.264444;-73.961944;1747;-5;A;America/New_York
8778;Long Lake;Long Lake;United States;NY9;\N;43.9750617;-74.42044;1629;-5;A;America/New_York
8779;Tyopliy Klyuch;Khandyga;Russia;;UEMH;62.471;136.505;932;10;N;Asia/Yakutsk
8780;Magan;Yakutsk;Russia;;UEMM;62.0626;129.3334;577;10;N;Asia/Yakutsk
8781;Zhoubai;Qianjiang;China;JIQ;\N;29.515;108.83;2461;8;N;Asia/Chongqing
8782;Mendeleevo;Yuzhno-Kurilsk;Russia;DEE;YXCM;43.5739;145.4107;50;12;N;\N
8783;Wishram Amtrak Station;Wishram;United States;WIH;\N;45.6576;-120.9664;1;-8;A;America/Los_Angeles
8784;Lima Allen County Airport;Lima;United States;AOH;KAOH;40.707478;-84.0270781;975;-5;A;America/New_York
8785;Sondok Airport;Hamhung;North Korea;DSO;\N;39.747333;127.474611;0;9;U;Asia/Pyongyang
8786;Samjiyon Airport;Samjiyon;North Korea;YJS;\N;41.907167;128.409889;0;9;U;Asia/Pyongyang
8787;Geiranger;Geiranger Fjord;Norway;;\N;62.474976458;6.1528587341;0;1;E;Europe/Oslo
8788;Gangtok Helipad;Gangtok;India;;\N;27.3521;88.6051;1500;5.5;N;Asia/Calcutta
8789;McKinnon Airport;Brunswick;United States;SSI;KSSI;31.1519722;-81.3910556;19;-5;A;America/New_York
8790;Beaver Falls;Beaver Falls;United States;BFP;KBFP;40.7724722;-80.3914444;1253;-5;A;America/New_York
8791;Seaplane Base;Winterhaven;United States;F57;\N;28.0575;-81.7628056;140;-5;A;America/New_York
8792;Georgetown County Airport;Georgetown;United States;GGE;KGGE;33.3114018;-79.3203139;40;-5;A;America/New_York
8793;Hardwick Field Airport;Cleveland;United States;HDI;KHDI;35.2199994;-84.832369;874;-5;A;America/New_York
8794;Mark Anton Airport;Dayton;United States;2A0;\N;35.48625;-84.9310833;718;-5;A;America/New_York
8795;Jefferson County Airpark;Steubenville;United States;2G2;\N;40.3602179;-80.7008742;1196;-5;A;America/New_York
8796;Renton;Renton;United States;RNT;KRNT;47.4931389;-122.21575;32;-8;A;America/Los_Angeles
8963;Bromont Airport;Bromont;Canada;ZBM;CZBM;45.290833;-72.741944;374;-5;A;America/Toronto
8798;Brackett Field;La Verne;United States;POC;KPOC;34.0916667;-117.7817778;1014;-8;A;America/Los_Angeles
8799;Jekyll Island Airport;Jekyll Island;United States;09J;\N;31.0744722;-81.4277778;11;-5;A;America/New_York
8800;CedarKey;Cedar Key;United States;CDK;KCDK;29.1342222;-83.0504722;11;-5;A;America/New_York
8801;Cross City;Cross City;United States;CTY;KCTY;29.6355278;-83.10475;42;-5;A;America/New_York
8802;Clemson;Clemson;United States;CEU;KCEU;34.6722222;-82.8858889;891;-5;A;America/New_York
8803;Heber City Municipal Airport;Heber;United States;36U;\N;40.4818056;-111.4288056;5637;-7;A;America/Denver
8804;Beech Factory Airport;Wichita;United States;BEC;KBEC;37.6939167;-97.2149167;1409;-6;A;America/Chicago
8805;Cherokee County Airport;Canton;United States;47A;\N;34.3122175;-84.4221556;1219;-5;A;America/New_York
8806;Fernandina Beach Municipal Airport;Fernandina Beach;United States;55J;\N;30.6118333;-81.4611944;16;-5;A;America/New_York
8807;Padang Tabing;Padang;Indonesia;;WIMG;-0.875;100.3519;3;7;U;Asia/Jakarta
8808;Tom B. David Field Airport;Calhoun;United States;;KCZL;34.45545;-84.9391622;651;-5;A;America/New_York
8809;Habersham County Airport;Cornelia;United States;;KAJR;34.4998483;-83.5566701;1448;-5;A;America/New_York
8810;Hamburg Hbf;Hamburg;Germany;ZMB;\N;53.552776;10.006683;30;1;E;Europe/Berlin
8811;Georgetown Municipal Airport;Georgetown;United States;GTU;KGTU;30.678809;-97.6793837;790;-6;A;America/Chicago
8812;Old Rhinebeck Airport;Rhinebeck;United States;;NY94;41.969833;-73.864555;323;-5;U;America/New_York
8813;Duxford Aerodrome;Duxford;United Kingdom;QFO;EGSU;52.09083;0.13194;125;0;E;Europe/London
8814;BLABU;Blankenburg;Germany;;\N;51.7946;11.0183;551;1;E;Europe/Berlin
8815;Sidney Muni Airport;Sidney;United States;SNY;KSNY;41.10167;-102.985;4313;-7;A;America/Denver
8816;Manantali Bengassi Airport ;Bengassi;Mali;;GA46;13.25561;-10.50432;538;0;N;Africa/Bamako
8817;Schofields Aerodrome;Schofields;Australia;;\N;-33.717389;150.871396;21;10;O;Australia/Sydney
8818;Luray Caverns Airport;Luray;United States;;KLUA;38.666944;-78.500556;903;-5;A;America/New_York
8819;Eagle's Nest Airport;Waynesboro;United States;W13;\N;38.076944;-78.944167;1437;-5;A;America/New_York
8820;Great Keppel Island;Great Keppel Island;Australia;GKL;YGKL;-23.186;150.943;250;10;U;Australia/Brisbane
8821;Roper Bar;Roper Bar;Australia;RPB;YRRB;-14.73457;134.5252;70;9.5;U;Australia/Darwin
8822;Mount Garnet;Mount Garnet;Australia;;YMRT;-17.70601;145.14889;2159;10;U;Australia/Brisbane
8823;Innisfail;Innisfail;Australia;IFL;YIFL;-17.55908;146.01123;46;10;U;Australia/Brisbane
8824;Kalaeloa;Kapolei;United States;;PHJR;21.307222;-158.070278;30;-10;U;Pacific/Honolulu
8825;Bamyan Airport;Bamyan;Afghanistan;BIN;OABN;34.816667;67.816667;2550;4.5;N;Asia/Kabul
8826;Changbaishan Airport;Baishan;China;NBS;\N;42.088056;127.548889;2762;8;N;Asia/Chongqing
8827;SFO Helicopter;Sausalito;United States;;\N;37.87856;-122.512742;13;-8;A;America/Los_Angeles
8828;Lewa Airport;Lewa Downs;Kenya;;\N;0.1927;37.4725;6250;3;U;Africa/Nairobi
8829;Huaorani Ecolodge;Huaorani;Ecuador;;\N;-1.0102;-77.1501;1000;-5;U;America/Guayaquil
8830;Chongjin Airport;Chongjin;North Korea;RGO;ZZ07;41.4297;129.6488;0;9;U;Asia/Pyongyang
8831;Moomba;Moomba;Australia;MOO;YOOM;-28.1;140.2;143;9.5;O;Australia/Adelaide
8832;Lublin;Lublin;Poland;LUZ;EPLB;51.239333;22.714083;203;1;E;Europe/Warsaw
8834;Madison County Executive Airport;Huntsville;United States;;KMDQ;34.861;-86.558;756;-6;A;America/Chicago
8835;Leesburg Executive Airport;Leesburg;United States;JYO;KJYO;39.078;-77.558;389;-5;A;America/New_York
8836;CNC4;Guelph;Canada;;\N;43.563889;-80.196111;1100;-5;A;America/Toronto
8837;CNC4;Geulph;Canada;;\N;43.563889;-80.196111;1100;-5;A;America/Toronto
8839;Robinson Crusoe Airport;San Juan Bautista;Chile;;SCIR;-33.665001;-78.929703;433;-4;S;America/Santiago
8840;Maamigili Airport;Maamigili;Maldives;VAM;VRMV;3.4702;72.8344;6;5;U;Indian/Maldives
8841;Hilton Iru fushi;Maldives Hilton Iru fushi;Maldives;IRU;IRUF;5.7435;73.3249;6;5;U;Indian/Maldives
8842;Dhigurah Centara Grand Maldives;Dhigurah;Maldives;DHG;DHGU;3.5939;72.8834;6;5;U;Indian/Maldives
8843;Beijing Railway Station;Beijing;China;;\N;39.5408;116.2516;143;8;N;Asia/Chongqing
8844;Chengde Railway Station;Chengde;China;;\N;40.965755;117.954019;1083;8;N;Asia/Chongqing
8845;Yongzhou Lingling Airport;Yongzhou;China;LLF;\N;26.345556;111.612222;428;8;U;Asia/Chongqing
8846;Losinj Airport;Mali Losinj;Croatia;LSZ;LDLO;44.3357;14.2335;154;1;E;\N
8847;Onslow ;Onslow;Australia;ONS;YOLW;-21.667;115.116;23;8;O;Australia/Perth
8848;Theodore;Theodore;Australia;TDR;YTDR;-24.986737;150.093112;560;10;N;Australia/Brisbane
8849;RUDNIKI ;RUDNIKI;Poland;CZW;EPRU;50.884722;19.202222;262;1;E;Europe/Warsaw
8850;Williamson-Sodus Airport;Williamson;United States;SDC;KSDC;43.2347904;-77.119444;424;-5;A;America/New_York
8851;Clear Lake Metroport;Clear Lake City;United States;CLC;\N;29.5569;-95.137497;35;-6;A;America/Chicago
8852;Gilberto Lavaque;Cafayate;Argentina;CFX;SASC;-26.06245;-65.932095;5523;-3;S;America/Cordoba
8853;Woking;Fairoaks;United Kingdom;;EGTF;51.348056;-0.558611;80;0;U;Europe/London
8854;Boulder Municipal;Boulder;United States;WBU;KBDU;40.039444;-105.225833;5288;-7;A;America/Denver
8855;Neustadt-Glewe;Neustadt-Glewe;Germany;;EDAN;53.364779;11.607352;115;1;E;Europe/Berlin
8856;Berchtesgarden BKS;Bischofswiesen;Germany;;\N;47.62313;12.972214;2297;1;E;Europe/Berlin
8857;Hohe Duene;Warnemuende;Germany;;ETMW;54.181438;12.132082;10;1;E;Europe/Berlin
8858;Sannvhe;Tangshan;China;TVS;ZBSN;39.717444;118.002389;25;8;N;Asia/Chongqing
8859;GOETSENHOVEN AB;GOETSENHOVEN;Belgium;;EBTN;50.7817;4.95778;246;1;E;Europe/Brussels
8860;Wildenrath AB Closed;Wildenrath;Germany;;EDUW;51.113721;6.210245;200;1;E;Europe/Berlin
8861;Bremgarten Airport;Bremgarten;Germany;;EDTG;47.902779;7.617778;300;1;E;Europe/Berlin
8862;Bitburg AB closed;Birburg;Germany;;EDRB;49.945278;6.565;1200;1;U;Europe/Berlin
8863;Toul AB closed;Toul;France;;\N;48.779999;5.98;936;1;E;Europe/Paris
8864;Palo Alto Airport of Santa Clara County;Palo Alto;United States;PAO;KPAO;37.4611111;-122.1150556;7;-8;A;America/Los_Angeles
8865;Dubai Cruise Terminal;Dubai;United Arab Emirates;;\N;25.274193;55.28311;0;4;U;Asia/Dubai
8866;Ust-Nera Airport;Ust-Nera;Russia;;UEMT;64.568056;143.236111;3000;11;U;Asia/Vladivostok
8867;Vostice;Vysoke Myto;Czech Republic;;LKVM;49.5537;16.1109;988;1;E;Europe/Prague
8868;Spa La Sauveniere;Spa;Belgium;;EBSP;50.4824981689453;5.91029977798462;1581;1;E;Europe/Brussels
8869;Mesa Falcon Field;Mesa;United States;FFZ;KFFZ;33.4608001708984;-111.727996826172;1394;-7;A;America/Phoenix
8870;Coolidge Municipal Airport;Cooldige;United States;P08;KP08;32.9359016418457;-111.427001953125;1574;-7;A;America/Phoenix
8871;Cottonwood Airport;Cottonwood;United States;P52;KP52;34.7299995422363;-112.035003662109;3550;-7;A;America/Phoenix
8872;Suarlee Airport;Namur;Belgium;;EBNM;50.4879989624023;4.76891994476318;594;1;E;Europe/Brussels
8873;Hasselt Airport;Hasselt;Belgium;;EBZH;50.9700012207031;5.37507009506226;141;1;E;Europe/Brussels
8874;Phoenix Regional Airport;Phoenix;United States;A39;KA39;32.99169921875;-111.920997619629;1300;-7;A;America/Phoenix
8875;Wickenburg Municipal Airport;Wickenburg;United States;E25;KE25;33.96889877;-112.7990036;2377;-7;A;America/Phoenix
8876;Yangzhou Taizhou Airport;Yangzhou;China;YTY;ZSYA;32.5617;119.715;7;8;N;Asia/Chongqing
8877;Oakland Co. Intl;Pontiac;United States;PTK;KPTK;42.667;-83.35;980;-5;A;America/New_York
8878;Fuzhou Railway Station;Fuzhou;China;;\N;26.11351;119.319998;39;8;N;Asia/Chongqing
8879;Xiamen Railway Station;Xiamen;China;;\N;24.4679;118.116;255;8;N;Asia/Chongqing
8880;Kissidougou;Kissidougou;Guinea;KSI;GUKU;9.160556;-10.124443;550;0;N;Africa/Conakry
8881;Pecs;Pecs;Hungary;QPJ;\N;46.070833;18.233056;1000;1;E;Europe/Budapest
8882;Neuwied;Neuwied;Germany;;\N;50.432008;7.472785;300;1;E;Europe/Berlin
8883;Castelnuovo di Garfagnana;Castelnuovo di Garfagnana;Italy;;\N;44.116464;10.409954;1100;1;E;Europe/Rome
8884;Portovenere;Portovenere;Italy;;\N;44.0519359;9.8353078;10;1;E;Europe/Rome
8885;Amsterdam Zuid WTC;Amsterdam;Netherlands;;\N;52.338841;4.871471;10;1;E;Europe/Amsterdam
8886;Budapest Keteli;Budapest;Hungary;;\N;47.500497;19.085484;100;1;E;Europe/Budapest
8887;Zagreb Glavni Kolod.;Zagreb;Croatia;;\N;45.804289;15.979121;500;1;E;Europe/Zagreb
8888;Dillant Hopkins Airport;Keene;United States;EEN;\N;72.270833;42.898333;149;-5;A;\N
8889;Tianshui Airport;Tianshui;China;THQ;ZLTS;34.333;105.514;9186;8;N;Asia/Chongqing
8890;Kawama Airport;Kawama;Cuba;;MUKW;23.123642;-81.301805;16;-5;U;America/Havana
8891;Yumuri Valley;Yumuri Valley;Cuba;;\N;23.1;-81.9;0;-5;U;America/Havana
8892;Kooddoo;Kooddoo;Maldives;GKK;VRMO;0.7308;73.433;10;5;U;Indian/Maldives
8893;Glasgow Industrial;Glasgow;United States;;07MT;48.4211069;-106.5277039;2762;-7;A;America/Denver
8894;Rochester Airport;Rochester;United Kingdom;RCS;EGTO;51.351944;0.502778;426;0;E;Europe/London
8895;Den Haag Centraal;Den Haag;Netherlands;;\N;52.081667;4.329167;0;1;E;Europe/Amsterdam
8896;Paris Nord;Paris;France;;\N;48.880931;2.355323;0;1;E;Europe/Paris
8897;Venlo railway station;Venlo;Netherlands;;\N;51.364722;6.171389;0;1;E;Europe/Amsterdam
8898;Monchengladbach Central Station;Monchengladbach;Germany;;\N;51.196;6.446;0;1;E;Europe/Berlin
8899;Utrecht Centraal;Utrecht ;Netherlands;;\N;52.089167;5.109722;0;1;E;Europe/Amsterdam
8900;Rensselaer Rail Station;Albany;United States;;\N;42.641389;-73.741111;324;-5;A;America/New_York
8901;Amtrak Station;Tacoma;United States;;\N;47.242733;-122.42041;50;-8;A;America/Los_Angeles
8902;Amtrak Station;Orlando;United States;;\N;28.5259;-81.3813;82;-5;A;America/New_York
8903;Denali Rail;Healy;United States;;\N;63.7306;-148.913733333333;1730;-9;A;America/Anchorage
8904;Albian Aerodrome;Wood Buffalo;Canada;JHL;\N;57.223889;-111.418889;1048;-7;A;America/Edmonton
8905;Monroe Reqional Airport;Charlotte;United States;EQY;\N;35.01833;-80.62001;679;-5;A;America/New_York
8906;Station Aare;Aare;Sweden;;\N;63.398779;13.075956;1240;1;E;Europe/Stockholm
8907;Villach Hauptbahnhof;Villach;Austria;;\N;46.618333;13.848333;1644;1;E;Europe/Vienna
8908;Krakow Glowny;Krakow;Poland;;\N;50.065556;19.947222;719;1;E;Europe/Warsaw
8909;Wien Meidling;Vienna;Austria;;\N;48.175;16.335278;738;1;E;Europe/Vienna
8910;Roma Termini;Rome;Italy;;\N;41.900833;12.501944;197;1;E;Europe/Rome
8911;Napoli Centrale;Naples;Italy;;\N;40.8525;14.271944;39;1;E;Europe/Rome
8912;Euston Station;London;United Kingdom;;\N;51.5284;-0.1331;89;0;E;Europe/London
8913;Kankan;Kankan;Guinea;KNN;GUXD;10.448333;-9.227499;8694;0;U;Africa/Conakry
8914;Termal;Rio Hondo;Argentina;RHD;SANH;-27.509946;-64.936666;1181;-3;S;America/Cordoba
8915;Guangzhou Railway Station;Guangzhou;China;;\N;23.151343;113.252073;19;8;N;Asia/Chongqing
8916;Shenzhen West Railway Station;Shenzhen;China;;\N;22.532043;113.908729;3;8;N;Asia/Chongqing
8917;Shenzhen North Railway Station;Shenzhen;China;;\N;22.607058;114.027203;229;8;N;Asia/Chongqing
8918;Keetmanshoop;Keetmanshoop;Namibia;KMP;FYKT;-26.5397;18.1114;3506;1;U;Africa/Windhoek
8919;Shekou Ferry Terminal;Shenzhen;China;;\N;22.476555;113.912855;2;8;N;Asia/Chongqing
8920;Zhuhai-Jiuzhou Port;Zhuhai;China;;\N;22.233333;113.583333;2;8;N;Asia/Chongqing
8921;Kangding Airport;Kangding;China;KGT;ZUKD;30.1575;101.734722;14042;8;N;Asia/Chongqing
8922;Train Station;Aranyaprathet;Thailand;;\N;13.4058;102.311;150;7;N;Asia/Bangkok
8923;Train Station;Nha Trang;Vietnam;;\N;12.15;109.11;50;7;N;Asia/Saigon
8924;Veliky Ustyug;Veliky Ustyug;Russia;VUS;ULWU;60.788333;46.26;331;4;N;Europe/Moscow
8925;Iowa City Municipal Airport;Iowa City;United States;IOW;KIOW;41.639244;-91.546503;668;-6;A;America/Chicago
8926;Turpan;Turpan;China;TLQ;ZWTP;42.942328;89.185877;20;8;U;Asia/Chongqing
8927;Lorenzo;Morro de Sao Paulo;Brazil;;SNCL;-13.389444445;-38.91;3;-3;S;America/Fortaleza
8928;Windom Municipal Airport;Windom;United States;MWM;KMWM;43.9134017;-95.1094083;1410;-6;A;America/Chicago
8929;Sasakwa Airstrip;Sasakwa;Tanzania;;\N;-2.0714;34.5068;4000;3;N;Africa/Dar_es_Salaam
8930;Yabuli;Yabuli Town;China;;\N;44.4653;128.27;1200;8;N;Asia/Chongqing
8931;Longview Ranch Airport;Longview;United States;;OG39;44.6608611;-119.6526153;2080;-8;A;America/Los_Angeles
8932;Rothera Research Station;Rothera Research Station;United Kingdom;;EGAR;-67.34;-68.08;200;12;U;Antarctica/South_Pole
8933;Sortavala;Sortavala;Russia;;ULPW;61.736099;30.673599;14;4;N;Europe/Moscow
8934;Valaam;Valaam;Russia;;\N;61.37115;30.88977;15;4;N;Europe/Moscow
8935;Lee Airport;Annapolis;United States;ANP;KANP;38.942778;-76.568333;34;-5;A;America/New_York
8936;Songshan Train Station;Taipei;Taiwan;;\N;25.049214;121.577961;0;8;N;Asia/Taipei
8937;Kaohsiung Station;Kaohsiung;Taiwan;;\N;22.638432;120.301877;39;8;N;Asia/Taipei
8938;Hualien Station;Hualien;Taiwan;;\N;23.992722;121.60135;45;8;N;Asia/Taipei
8939;El Merk;El Merk;Algeria;;\N;30.300039;8.160833;0;1;U;Africa/Algiers
8940;Budapest Keleti pu.;Budapest;Hungary;;\N;-47.500278;-160.916111;0;1;E;\N
8941;Bratislava hl. st.;Bratislava;Slovakia;;\N;-48.158263;17.105839;430;1;E;\N
8942;Ndutu;Ndutu;Tanzania;DUU;HTND;-3.0243398;34.98450589999993;5400;3;U;Africa/Dar_es_Salaam
8943;Cuamba;Cuamba;Mozambique;FXO;FQCB;-14.8175;36.528333;1931;2;N;Africa/Maputo
8944;Bodaibo;Bodaibo;Russia;ODO;UIKB;57.5137;114.1348;940;9;N;Asia/Irkutsk
8945;Berdoba Airport;Berdoba;Chad;;\N;16.072;22.8466;2953;1;N;Africa/Ndjamena
8946;Zhytomyr;Zhytomyr;Ukraine;ZTR;UKKV;50.270556;28.738611;0;2;E;Europe/Kiev
8947;Mozyr;Mozyr;Belarus;;UMGM;51.981944;29.165;0;3;N;Europe/Minsk
8948;Paluknys;Paluknys;Lithuania;;EYVP;54.483056;24.989722;0;2;E;Europe/Vilnius
8949;Mattala Rajapaksa Intl.;Mattala;Sri Lanka;HRI;VCRI;6.284467;81.124128;157;5.5;N;Asia/Colombo
8950;Pecos Municipal Airport;Pecos;United States;PEQ;KPEQ;31.3823889;-103.5107222;2613;-6;A;America/Chicago
8951;Hattiesburg Bobby L. Chain Municipal Airport;Hattiesburg;United States;HBG;KHBG;31.2649444;-89.2528889;151;-6;A;America/Chicago
8952;Botucatu;Botucatu;Brazil;QCJ;SDBK;-22.938;-48.467;3012;-3;S;America/Sao_Paulo
8953;Anapolis Military Airbase;Anapolis;Brazil;;SBAN;-16.23;-48.966;3731;-3;S;America/Sao_Paulo
8954;Sao Carlos TAM;Sao Carlos;Brazil;QSC;SDSC;-21.875;-47.904;2649;-3;S;America/Sao_Paulo
8955;Byron Airport;Byron Bay;Australia;;\N;28.383504;153.365447;10;10;O;\N
8956;Grindelwald;Grindelwald;Switzerland;;\N;46.35106;7.57408;6000;1;U;Europe/Zurich
8957;Grindelwald;Grindelwald;Switzerland;;\N;46.3506;7.57408;6000;1;U;Europe/Zurich
8958;Chan Gurney;Yankton;United States;YKN;KYKN;42.8711;-97.3969;1200;-6;A;America/Chicago
8959;Selfield Airport;Selma Alabama;United States;SES;KSES;32.4404;-86.9522;131;-6;A;America/Chicago
8960;Linkwasha Airfield;Hwange National Park;Zimbabwe;;\N;-19.123333;27.201389;3543;2;N;Africa/Harare
8961;Sir Bani Yas Island;Sir Bani Yas Island;United Arab Emirates;XSB;\N;24.2856083;52.5783472;5;4;U;Asia/Dubai
8962;Dalma Airport;Dalma Island;United Arab Emirates;ZDY;\N;24.5033833;52.3360528;8;4;U;Asia/Dubai
8964;Playa del Carmen Airport;Playa del Carmen;Mexico;PCM;\N;20.6225;-87.082221;1;-6;U;America/Mexico_City
8965;Ellough;Beccles;United Kingdom;;EGSM;52.438333;1.606944;80;0;E;Europe/London
8966;Kratie Airport;Kratie;Cambodia;KTI;VDKT;12.4905;106.055;3610;7;N;Asia/Phnom_Penh
8967;Caldera Airport;Caldera;Chile;;SCCL;-27.078052;-70.795305;180;-4;U;America/Santiago
8968;San Pedro de Atacama Airport;San Pedro de Atacama;Chile;;SCPE;-22.918888;-68.162777;7959;-4;U;America/Santiago
8969;Copacabana Airport;Copacabana;Bolivia;;SLCC;-16.183326;-69.087501;12591;-4;U;America/La_Paz
8970;Havelock Island Seaport;Havelock Island;India;;\N;12.042636;92.982574;0;5.5;N;Asia/Calcutta
8971;Guyuan;Guyuan;China;GYU;ZLGY;36.078611;106.216667;5577;8;U;Asia/Chongqing
8972;Brawdy RAF Airport;Brawdy;United Kingdom;;EGDA;51.880462;-5.12148;355;0;E;Europe/London
8973;Gaios Bay;Paxi;Greece;;\N;39.195944;20.195961;0;2;E;Europe/Athens
8974;Changhai;Changhai;China;CNI;ZYCH;39.266389;122.666944;1;8;N;Asia/Chongqing
8975;Redhill Aerodrome;Redhill;United Kingdom;KRH;EGKR;51.213612;-0.13861;221;0;E;Europe/London
8976;Bahnhof;Bad Gastein;Austria;;\N;47.1;13.0167;3287;1;E;Europe/Vienna
8977;Acquafredda Station;Maratea;Italy;;\N;40;15.7167;1000;1;E;Europe/Rome
8978;Victoria Station;London;United Kingdom;;\N;51.4964;-0.14391;48;0;E;Europe/London
8979;Jiagedaqi Airport;Jiagedaqi District;China;JGD;\N;50.375;124.117;1205;8;N;Asia/Chongqing
8980;Chinchilla;Chinchilla;Australia;CCL;YCCA;-26.769444;150.616667;1030;10;O;Australia/Brisbane
8981;Frazier Lake Airpark;Hollister;United States;1C9;\N;54.013333333333335;-124.76833333333333;152;-8;A;America/Vancouver
8982;Hayward Executive Airport;Hayward;United States;HWD;KHWD;37.65888888888889;-122.12166666666666;52;-8;A;America/Los_Angeles
8983;Motueka;Motueka;New Zealand;MZP;NZMK;-41.123299;172.988998;39;12;Z;Pacific/Auckland
8984;Shute Harbour;Shute Harbour;Australia;JHQ;YSHR;-20.278299;148.757004;12;10;O;Australia/Brisbane
8985;Enstone;Enstone;United Kingdom;;EGTN;51.928167;-1.4285;550;0;U;Europe/London
8986;Portoferrario Cruise Terminal;Portoferrario;Italy;;\N;42.49;10.19;13;1;E;Europe/Rome
8987;Pittsburgh Amtrak;Pittsburgh;United States;;\N;40.4406;-79.9961;1370;-5;A;America/New_York
8988;Raton Amtrak;Raton;United States;;\N;36.9033;-104.4386;8800;-7;A;America/Denver
8989;Ann Arbor Municipal Airport;Ann Arbor;United States;ARB;KARB;42.132274;-83.444418;839;-5;A;America/New_York
8990;Myrdal Station;Myrdal;Norway;;\N;60.4418;7.0723;2844;1;E;Europe/Oslo
8991;Flam Station;Flam;Norway;;\N;60.5017;7.0714;6;1;E;Europe/Oslo
8992;Leikanger Terminal;Leikanger;Norway;;\N;61.1351;6.4723;6;1;E;Europe/Oslo
8993;North Sea;Gryphon FPSO;United Kingdom;;GRY1;59.35;1.57;0;0;U;\N
8994;Shepparton;Shepparton;Australia;SHT;YSHT;-36.428902;145.393005;374;10;O;Australia/Hobart
8995;Temora;Temora;Australia;TEM;YTEM;-34.4248;147.5104;921;10;O;Australia/Sydney
8996;Gayndah;Gayndah;Australia;GAH;YGAY;-25.61528;151.62083;369;10;N;Australia/Brisbane
8997;Popondetta;Popondetta;Papua New Guinea;EIA;\N;-8.19999980927;147.850006104;311;10;U;Pacific/Port_Moresby
8998;Wilcannia;Wilcannia;Australia;WIO;YWCA;-31.5174999237;143.378997803;250;10;O;Australia/Sydney
8999;Bollards Lagoon;Bollards Lagoon Station;Australia;;\N;-28.9832259;140.852844;250;9.5;O;Australia/Adelaide
9000;Delamere Range;Delamere Range;Australia;;\N;-15.7467;131.919998;200;9.5;U;Australia/Darwin
9001;Ivanhoe;Ivanhoe;Australia;;YIVO;-32.883301;144.309998;331;10;O;Australia/Sydney
9002;Menindee;Menindee;Australia;;YMED;-32.22;142.2418;164;10;O;Australia/Sydney
9003;Naduri;Naduri;Papua New Guinea;;\N;-8.93;147.73;900;10;U;Pacific/Port_Moresby
9004;Pooncarie;Pooncarie;Australia;;YPCE;-33.366699;142.587977;184;10;O;Australia/Sydney
9005;Wiawera Station;Wiawera station;Australia;;\N;-32.29515;140.4145;180;9.5;O;Australia/Adelaide
9006;Tilpa;Tilpa;Australia;;YTLP;-30.556;144.251;280;10;O;Australia/Sydney
9007;Yarramba Station;Yarramba Station;Australia;;\N;-31.66136;140.6277;180;9.5;O;Australia/Adelaide
9008;Reichenbach;Reichenbach im Kandertal;Switzerland;;LSGR;46.61298;7.67773;2300;1;E;Europe/Zurich
9009;Takayama Station;Takayama;Japan;;\N;36.1406;137.2519;2200;9;N;Asia/Tokyo
9010;Himeji Station;Himeji;Japan;;\N;34.49395;134.412706;141;9;U;Asia/Tokyo
9011;Miyajimaguchi Station;Miyajimaguchi;Japan;;\N;34.3122;132.3022;6;9;U;Asia/Tokyo
9012;Miyajima Port;Miyajima;Japan;;\N;34.302043;132.32213;0;9;U;Asia/Tokyo
9013;Brown Bluff Port;Brown Bluff;Antarctica;;\N;-63.32;-56.54;0;12;S;Antarctica/South_Pole
9014;Devil Island Port;Devil Island;Antarctica;;\N;-63.48;-57.17;0;12;S;Antarctica/South_Pole
9015;Gourdin Island Port;Gourdin Island;Antarctica;;\N;-63.12;-57.18;0;12;S;Antarctica/South_Pole
9016;Cuverville Island Port;Cuverville Island;Antarctica;;\N;-64.41;-62.38;0;12;S;Antarctica/South_Pole
9017;Neko Harbor;Neko Harbor;Antarctica;;\N;-64.5;-62.33;0;12;S;Antarctica/South_Pole
9018;Port Lockroy;Port Lockroy;Antarctica;;\N;-64.4931;-63.294;0;12;S;Antarctica/South_Pole
9019;Walker Bay;Walker Bay;Antarctica;;\N;-62.38;-60.42;0;12;S;Antarctica/South_Pole
9020;Whalers Bay;Deception Island;Antarctica;;\N;-62.5837;-60.39;0;12;S;Antarctica/South_Pole
9021;Labadee Port;Labadee;Haiti;;\N;19.4711;-72.1444;0;-5;A;America/Port-au-Prince
9022;Lagos Lagoon;Lekki;Nigeria;;\N;6.465907;3.532115;0;1;N;Africa/Lagos
9023;Escravos;Escravos;Nigeria;;\N;5.617203;5.187993;0;1;U;Africa/Lagos
9024;Volosovo;Volosovo;Russia;;XUMW;55.0683;37.4517;577;4;N;Europe/Moscow
9025;Bijie Feixiong Airport;Bijie;China;BFJ;\N;27.253;105.426;5250;8;N;Asia/Chongqing
9026;Lensk;Lensk;Russia;ULK;UERL;60.723;114.825;843;10;U;Asia/Yakutsk
9027;Fairbanks Train Depot;Fairbanks;United States;;\N;64.5109;-147.4445;446;-9;A;America/Anchorage
9028;Denali Train Depot;Denali;United States;;\N;63.43083;-150.31111;3733;-9;A;America/Anchorage
9029;Talkeetna Train Depot;Talkeetna;United States;;\N;62.18934;-150.06177;348;-9;A;America/Anchorage
9030;Anchorage Rail Depot;Anchorage;United States;;\N;61.1318;-149.5326;102;-9;A;America/Anchorage
9031;Whittier Port;Whittier;United States;;\N;60.774174;-148.677649;0;-9;A;America/Anchorage
9032;Qingdao Railway Station;Qingdao;China;;\N;36.063497;120.312542;45;8;N;Asia/Chongqing
9033;Beijing South Railway Station;Beijing;China;;\N;39.5157;116.2235;144;8;N;Asia/Chongqing
9034;Shanghai Railway Station;Shanghai;China;;\N;31.251552;121.450446;16;8;N;Asia/Chongqing
9035;Suzhou Railway Station;Suzhou;China;;\N;31.195;120.3623;9;8;N;Asia/Chongqing
9036;Biluty;Khamar-Daban;Russia;;\N;51.994;102.099;3000;9;N;Asia/Irkutsk
9037;Shumak;Khamar-Daban;Russia;;\N;51.964;101.867;3600;9;N;Asia/Irkutsk
9038;Shumak;Khamar-Daban;Russia;;\N;51.964;101.897;3300;9;N;Asia/Irkutsk
9039;Kitoi;Khamar-Daban;Russia;;\N;52.111;101.986;2000;9;N;Asia/Irkutsk
9040;Kitoi;Khamar-Daban;Russia;;\N;52.111;101.986;2000;9;N;Asia/Irkutsk
9041;Angasolka;Baikal;Russia;;\N;51.758;103.849;1500;9;N;Asia/Irkutsk
9042;Kavalerovo;Kavalerovo;Russia;;UHWK;44.273;135.029;660;11;N;Asia/Vladivostok
9043;Igdir;Igdir;Turkey;IGD;LTCT;39.983056;43.866389;3100;2;E;Europe/Istanbul
9044;Sanliurfa GAP;Sanliurfa;Turkey;GNY;LTCS;37.45;38.9;2700;2;E;Europe/Istanbul
9045;Zafer;Kutahya;Turkey;KZR;LTBZ;39.111389;30.13;3327;2;E;Europe/Istanbul
9046;Barentsburg;Barentsburg;Svalbard;;\N;78.1006;14.1998;20;1;U;Arctic/Longyearbyen
9047;Piramida;Piramida;Svalbard;;\N;78.6528;16.3392;20;1;U;Arctic/Longyearbyen
9048;Velikiye Luki;Velikiye Luki;Russia;VLU;ULOL;56.381667;30.61;328;4;N;Europe/Moscow
9049;Boston Back Bay Station;Boston;United States;ZTY;\N;42.3478;-71.075;20;-5;A;America/New_York
9050;Burlington GO Station;Burlington;Canada;;\N;43.3408282;-79.8093881;243;-5;A;America/Toronto
9051;Niagra Falls Railway Station;Niagra Falls;Canada;;\N;43.1088;-79.0634;521;-5;A;America/New_York
9052;Niagara Falls Station;Niagara Falls;United States;;\N;43.1135;-79.0318;0;-5;A;America/New_York
9053;Yelahanka AFB;Bangalore;India;YLK;VOYK;13.136;77.607;2912;5.5;N;Asia/Calcutta
9054;Yeltsovka;Novosibirsk;Russia;;UNNE;55.077;82.997;617;7;N;Asia/Omsk
9055;Veligandu Resort;Veligandu;Maldives;;\N;4.1753;73.0037;0;5;U;Indian/Maldives
9056;Kuramathi;Kuramathi;Maldives;;\N;4.1525;73;0;5;U;Indian/Maldives
9057;Olkiombo Airstrip;Masai Mara;Kenya;;HKOK;-1.409569;35.110788;4997;3;U;Africa/Nairobi
9058;Shanghai South Railway Station;Shanghai;China;;\N;31.154425;121.42416;13;8;N;Asia/Chongqing
9059;Suzhou North Railway Station;Suzhou;China;;\N;31.423646;120.639045;6;8;N;Asia/Chongqing
9060;Ulithi;Ulithi;Micronesia;ULI;\N;10.016699791;139.800003052;10;10;O;Pacific/Truk
9061;Mar-Kuel;Mar-Kuel;Russia;;UHHK;57.2712;132.314;200;11;N;Asia/Vladivostok
9062;Kodinsk;Kodinsk;Russia;;UNKI;58.285;99.055;200;8;N;Asia/Krasnoyarsk
9063;Nanjing Railway Station;Nanjing;China;;\N;32.088583;118.792189;29;8;N;Asia/Chongqing
9064;Balkanabat;Balkanabat;Turkmenistan;BKN;\N;39.481389;54.362778;7;5;N;Asia/Ashgabat
9065;Belmont Airport;Lake Macquarie;Australia;BEO;YPEC;-33.066667;151.648333;5;10;O;Australia/Sydney
9066;Clayton County Tara Field;Hampton;United States;4A7;K4A7;33.389099;-84.332397;874;-5;A;America/New_York
9067;Stapleton International Airport;Denver;United States;;DENX;39.779255;-104.88184;5333;-7;A;America/Denver
9068;Brampton Island;Brampton Island;Australia;BMP;YBPI;-20.804444402073734;149.2794443842832;5;10;O;Australia/Brisbane
9069;Shanghai Maglev Long Yang Road Station;Shanghai;China;;\N;31.202854;121.558335;16;8;U;Asia/Chongqing
9070;Nuernberg Port;Nuernberg;Germany;;\N;49.392222;11.061944;990;1;E;Europe/Berlin
9071;NAS Alameda;Alameda;United States;NGZ;KNGZ;37.7861;-122.3186;10;-8;U;America/Los_Angeles
9072;Stein;Stein;Germany;;\N;49.406324;11.003054;984;1;E;Europe/Berlin
9073;Dippenricht Heliport;Dippenricht;Germany;;\N;49.253972;11.363216;1400;1;E;Europe/Berlin
9074;Lagoon Cove water aerodrome;Lagoon Cove;Canada;;\N;50.6;-126.3167;0;-8;A;America/Vancouver
9075;Endelage West;Endelage;Denmark;;EKEL;55.75616;10.250716;0;1;E;Europe/Copenhagen
9076;Flugplatz St. Johann in Tirol;St. Johann in Tirol;Austria;;LOIJ;47.517222;12.435833;2198;1;E;Europe/Vienna
9077;Echo Bay water aerodrome;Echo Bay;Canada;;\N;50.45;-126.29;0;-8;A;America/Vancouver
9078;Kingcome water aerodrome;Kingcome;Canada;;\N;50.58;-128.11;0;-8;A;America/Vancouver
9079;Thalmaessing Waizenhofen;Thalmaessing;Germany;;EDPW;49.064444;11.208056;1893;1;E;Europe/Berlin
9080;Tcentralny;Taganrog;Russia;;XRRC;47.246;38.84;183;4;N;Europe/Moscow
9081;Buffalo Exchange Street Station;Buffalo;United States;;\N;42.8783;-78.8738;600;-5;A;America/New_York
9082;Wilmington Amtrak Station;Wilmington;United States;ZWI;\N;39.736667;-75.551667;0;-5;A;America/New_York
9083;Vagaru Island;Viceroy Resort;Maldives;;\N;6.089305;73.203857;5;5;N;Indian/Maldives
9084;Fort McMurray - Mildred Lake Airport;Fort McMurray;Canada;NML;\N;57.055599;-111.573997;1046;-7;A;America/Edmonton
9085;Genoa Cruise Terminal;Genova;Italy;;\N;44.40725;8.876301;0;1;E;Europe/Rome
9086;Suzhou Industrial Park Railway Station;Suzhou;China;;\N;31.344444;120.706667;7;8;N;Asia/Chongqing
9087;Elkhart Municipal;Elkhart;United States;EKI;\N;41.719444;-86.003333;778;-5;A;America/New_York
9088;Cochrane;Cochrane;Canada;YCN;CYCN;49.106667;-81.015278;852;-5;A;America/Toronto
9089;Aeroporto Estadual Arthur Siqueira;Braganca Paulista;Brazil;BJP;SBBP;-22.979167;-46.5375;3;-3;S;America/Sao_Paulo
9090;Brusselton;Brusselton;Australia;BQB;YBLN;-33.687222;115.400278;55;8;O;Australia/Perth
9091;Srednekolymsk;Srednekolymsk;Russia;;UESK;67.2748;153.4258;98;12;N;Asia/Magadan
9092;Garowe - International;Garowe;Somalia;GGR;\N;8.24;48.29;553;3;N;Africa/Mogadishu
9093;Salt Lake City Intermodal Hub;Salt Lake City;United States;;SLCR;40.762778;-111.908333;0;-7;A;America/Denver
9094;Thorny Bush Game Lodge Airport;Hoedspruit;South Africa;;\N;-24.414722;31.165001;1900;2;U;Africa/Johannesburg
9095;Inverell;Inverell;Australia;IVR;YIVL;-29.888333;151.144167;2667;10;O;Australia/Sydney
9096;Glen Innes;Glen Innes;Australia;GLI;YGLI;-29.675;151.69;3433;10;O;Australia/Sydney
9097;Hangzhou Railway Station;Hangzhou;China;;\N;30.246;120.1784;39;8;N;Asia/Chongqing
9098;Sorrento Valley Station;San Diego;United States;;\N;32.903;-117.225;150;-8;A;America/Los_Angeles
9099;San Diego Union Station;San Diego;United States;;\N;32.716944;-117.168611;50;-8;A;America/Los_Angeles
9100;Petersdorf;Ansbach;Germany;;EDQF;49.216;10.4026;1370;1;E;Europe/Berlin
9101;Venice-Venezia Mestre;Venice;Italy;;\N;45.48189;12.231923;0;1;E;Europe/Rome
9102;Immokalee ;Immokalee ;United States;IMM;KIMM;26.433889;-81.401389;37;-5;A;America/New_York
9103;Maranggo;Sulawesi Tenggara;Indonesia;;WA44;-5.76457;123.917;169;8;N;Asia/Makassar
9104;Rancho San Simeon Airport;Cambria;United States;;66CA;35.607748;-121.110192;320;-8;A;America/Los_Angeles
9105;Port Canaveral;Brevard County;United States;;\N;28.4113919;-80.6081066;0;-5;A;America/New_York
9106;Cococay;Ragged Island;Bahamas;;\N;22.148056;-75.693333;0;-5;A;America/Nassau
9107;Yichun Mingyueshan Airport;Yichun;China;YIC;ZSYC;27.8033;114.3081;131;8;U;Asia/Chongqing
9108;Costa maya port;Costa maya;Mexico;;\N;18.714;-87.709;0;-6;U;America/Mexico_City
9109;Costa maya port;Costa maya;Mexico;;\N;18.714;-87.709;0;-6;U;America/Mexico_City
9110;Fajardo;Fajardo;Puerto Rico;;\N;18.32;-65.65;0;-4;U;America/Puerto_Rico
9111;Billings field;Essex;Canada;;\N;42.199;-82.465;650;-5;U;America/Toronto
9112;Hiberico Offshore Platform;Cuidad del Carmen;Mexico;;\N;19.398611;-92.016111;0;-6;A;\N
9113;Alpha offshore platform;Cork;Ireland;;\N;51.400377;-7.901464;0;0;E;\N
9114;Peterhof Helicopter Landing area;Saint-Petersburg;Russia;;\N;59.890529;29.894609;0;4;U;Europe/Moscow
9115;Gogland Helicopter Landing;Gogland Island;Russia;;\N;60.011662;27.008236;0;4;U;\N
9116;New Carrollton Rail Station;New Carrollton;United States;ZRZ;\N;38.948;-76.8719;39;-5;A;America/New_York
9117;Ship Channel Cay;Ship Channel Cay;Bahamas;;\N;24.819302;-76.824331;0;-5;A;America/Nassau
9118;Walkersville Train Depot;Walkersville;United States;;\N;39.48618333333334;-77.35475;308;-5;A;America/New_York
9119;Seattle Pier 69 - Victoria Clipper;Seattle;United States;;\N;47.613801;-122.354019;7;-8;A;America/Los_Angeles
9120;Dinwiddie County Airport;Petersburg;United States;PTB;KPTB;37.18375;-77.507388;193;-5;A;America/New_York
9121;Mount Aso helipad;Aso;Japan;;\N;32.884308;131.064638;1350;9;N;Asia/Tokyo
9122;Mount Aso helipad;Aso;Japan;;\N;32.884308;131.064638;1350;9;N;Asia/Tokyo
9123;Kasongo Lunda;Kasongo;Congo (Kinshasa);KGN;FZOK;-6.583333;16.816668;0;1;U;Africa/Kinshasa
9124;McMurdo Station Pegasus Field;McMurdo Station;Antarctica;;NZPG;-77.963333;166.524444;18;12;Z;Antarctica/South_Pole
9125;Klatovy;Klatovy;Czech Republic;;LKKT;49.418327;13.321944;1299;1;E;Europe/Prague
9126;Sheboygan County Memorial Airport;Sheboygan;United States;SBM;KSBM;43.7696;-87.851402;755;-6;A;America/Chicago
9127;Dave Forest Airport;Cloudbreak;Australia;KFE;YFDF;-22.291945;119.437225;1555;8;N;Australia/Perth
9128;Ephraim-Gibraltar Airport;Ephraim;United States;3D2;\N;45.1354167;-87.1879444;773;-6;A;America/Chicago
9129;Kangel Danda Airport;Kangel Danda;Nepal;;VNKL;27.41667;86.64167;6880;5.75;N;Asia/Katmandu
9130;Man Maya Airport;Khanidanda;Nepal;;VNKD;27.19028;86.77167;4435;5.75;N;Asia/Katmandu
9131;Talcha Airport;Rara National Park;Nepal;;VNT1;29.43;81;0;5.75;N;Asia/Katmandu
9132;Bajura Airport;Bajura;Nepal;BJU;VNBR;29.50611;81.66944;4606;5.75;N;Asia/Katmandu
9133;Walkersville Turnaround;Frederick;United States;;\N;39.4449;-77.4012;302;-5;A;America/New_York
9134;Airport Chara;Chara;Russia;;UIAR;56.912765;118.271007;671;10;N;Asia/Yakutsk
9135;Tho Xuan Airport;Thanh Hoa;Vietnam;THD;\N;19.90167;105.46778;0;7;N;Asia/Saigon
9136;Herning;Herning;Denmark;;EKHG;56.197224;9.147222;136;1;E;Europe/Copenhagen
9137;Sonderlandeplatz Schleissheim;Schleissheim;Germany;;EDNX;48.239167;11.561389;1594;1;E;Europe/Berlin
9138;Pinal Airpark;Marana;United States;MZJ;KMZJ;32.509722;-111.325278;1893;-7;N;America/Phoenix
9139;Glendale Municipal Airport;Glendale;United States;GEU;KGEU;33.5269167;-112.2951389;1071;-7;N;America/Phoenix
9140;Safford Regional Airport;Safford;United States;SAD;KSAD;32.8533333;-109.6350833;3179;-7;N;America/Phoenix
9141;Verden Scharnhorst;Verden;Germany;;EDWV;52.965278;9.282778;144;1;E;Europe/Berlin
9142;Mara North;Masai Mara;Kenya;;\N;-1.142636;35.139856;5468;3;N;Africa/Nairobi
9143;Fort Lauderdale-Hollywood International Airport Station at Dania Beach;Dania Beach;United States;;\N;26.061933;-80.16565;9;-5;A;America/New_York
9144;Deerfield Beach Station;Deerfield Beach;United States;;\N;26.316944;-80.122222;13;-5;A;America/New_York
9145;Lecheo;Korinthos;Greece;;\N;37.92592;22.860228;60;2;E;Europe/Athens
9146;Solomon Airport;Solomon;Australia;SLJ;YSOL;-22.255;117.85;2008;8;N;Australia/Perth
9147;Bruck-Nittenau;Nittenau;Germany;;EDNM;49.2225;12.2967;1161;1;E;Europe/Berlin
9148;Kerama Airport;Kerama;Japan;KJP;ROKR;26.168333;127.293333;156;9;N;Asia/Tokyo
9149;Americana;Americana;Brazil;;SDAI;-22.7558;-47.269402;2085;-3;S;America/Sao_Paulo
9150;Brussels Gare Centrale;Brussels;Belgium;;\N;50.84548;4.357174;43;1;E;Europe/Brussels
9151;Comino Airport;Comino;Malta;JCO;\N;36.016;14.333;201;1;E;Europe/Malta
9152;Udine stazione;Udine;Italy;;\N;46.04;13.14;1;1;E;Europe/Rome
9153;Bratislava Hlavna Stanica;Bratislava;Slovakia;;\N;48.08;17.06;1;1;E;Europe/Vienna
9154;Depati Parbo Airport;Sungai Penuh;Indonesia;;\N;1.881071;101.442261;2000;7;U;Asia/Jakarta
9155;Kolosovka;Kolosovka;Russia;;\N;56.2752;73.3641;100;7;U;Asia/Omsk
9156;Ekibastuz airport;Ekibastuz;Kazakhstan;;UASB;51.5723;75.2613;776;6;N;Asia/Qyzylorda
9157;Znamenka airport;Znamenskoe;Russia;;\N;57.1523;73.7512;432;7;E;Asia/Omsk
9158;Tevriz airport;Tevriz;Russia;;\N;57.5243;72.4653;456;7;U;Asia/Omsk
9159;Besobe airport;Bestobe;Kazakhstan;;\N;52.4986;73.1434;645;6;N;Asia/Qyzylorda
9160;Stepnogorsk airport;Stepnogorsk;Kazakhstan;;\N;52.3575;71.7865;432;6;N;Asia/Qyzylorda
9161;Bekdash airport;Karabogaz;Turkmenistan;;\N;41.5397;52.5711;12;5;U;Asia/Ashgabat
9162;Adamovka;Adamovka;Russia;;UWOD;51.4929;59.92752;931;6;N;Asia/Yekaterinburg
9163;Kvarkeno;Kvarkeno;Russia;;UWOH;52.078333;59.683333;305;6;U;Asia/Yekaterinburg
9164;Sikeston Memorial Municipal;Sikeston;United States;SIK;KSIK;36.8988889;-89.56175;315;-6;A;America/Chicago
9165;Nynashamn Ferry Port;Nynashamn;Sweden;NYN;\N;58.9;17.95;0;1;E;Europe/Stockholm
9166;Rodby Port;Rodby;Denmark;ROD;\N;54.65;11.35;0;1;E;Europe/Copenhagen
9167;Puttgarden;Puttgarden;Germany;QUA;\N;54.5;11.2167;0;1;E;Europe/Berlin
9168;Sonevafushi;Sonevafushi;Maldives;;\N;5.112336;73.07805;0;5;U;Indian/Maldives
9169;Royal Island Resort;Royal Island Resort;Maldives;;\N;5.162911;73.053578;0;5;U;Indian/Maldives
9170;Royal Island Resort;Royal Island Resort;Maldives;;\N;5.162911;73.053578;0;5;U;Indian/Maldives
9171;Port of Belfast;Belfast;United Kingdom;BE2;\N;54.6178;-5.9017;0;0;E;Europe/London
9172;Stranraer Ferry Port;Stranraer;United Kingdom;SR2;\N;54.902;-5.027;0;0;E;Europe/London
9173;Oceano County Airport;Oceano;United States;L52;KL52;35.10147222;-120.62236111;14;-8;U;America/Los_Angeles
9174;Eupen - Rail;Eupen;Belgium;;\N;50.3743;6.0201;1000;1;E;Europe/Brussels
9175;Tetiaroa Airport;Tetiaroa;French Polynesia;TTI;NTTE;-17.0132999;-149.5870056;0;-10;N;Pacific/Tahiti
9176;Osnabrueck Hbf;Osnabrueck;Germany;;\N;52.272778;8.061111;207;1;E;Europe/Berlin
9177;Hannover Hbf;Hannover;Germany;;\N;52.366667;9.716667;180;1;E;Europe/Berlin
9178;Floyd Bennett Memorial Airport;Queensbury;United States;GFL;KGFL;43.341222;-73.610305;328;-5;A;America/New_York
9179;Saratoga County Airport;Ballston Spa;United States;5B2;K5B2;43.050722;-73.861638;433;-5;A;America/New_York
9180;Burrello-Mechanicville Airport;Mechanicville;United States;K27;\N;42.893133;-73.66845;195;-5;A;America/New_York
9181;Tampa Padang;Mamuju;Indonesia;MJU;WAWJ;-2.583333;119.033333;2;8;N;Asia/Makassar
9182;Crystal River;Crystal River;United States;CGC;KCGC;28.8676111;-82.5741111;9;-5;A;America/New_York
9183;Martin State;Baltimore;United States;MTN;KMTN;39.3256667;-76.4137778;22;-5;A;America/New_York
9184;Lincoln Regional Airport Karl Harder Field;Lincoln;United States;LHM;KLHM;38.9091667;-121.3513333;121;-8;A;America/Los_Angeles
9185;Fostoria Metropolitan Airport;Fostoria;United States;FZI;KFZI;41.1908333;-83.3930833;752;-5;A;America/New_York
9186;Eastern Slopes Regional;Fryeburg;United States;IZG;KIZG;43.9911389;-70.9478889;455;-5;A;America/New_York
9187;Coral Creek;Placida;United States;;FA54;26.8537417;-82.2526028;8;-5;A;America/New_York
9188;Lakefront;New Orleans;United States;NEW;KNEW;30.0424167;-90.02825;7;-6;A;America/Chicago
9189;Pappy Boyington;Coeur d'Alene;United States;COE;KCOE;47.7743056;-116.8195833;2320;-8;A;America/Los_Angeles
9190;Beaumont Municipal;Beaumont;United States;BMT;KBMT;30.0702044;-94.2150967;32;-6;A;America/Chicago
9191;Vermilion Regional;Danville;United States;DNV;KDNV;40.1996944;-87.5955278;696;-6;A;America/Chicago
9192;Varberg Railway Station;Varberg;Sweden;;\N;57.107118;12.2520907;0;1;E;Europe/Stockholm
9193;Coonabarabran;Coonabarabran;Australia;COJ;YCBB;-31.337221;149.27194;2117;10;O;Australia/Sydney
9194;Toronto Union Station;Toronto;Canada;;\N;43.645278;-79.380556;249;-5;A;America/Toronto
9195;Windsor Station;Windsor;Canada;;\N;42.3254;-83.0092;623;-5;A;America/Toronto
9196;Dearborn Amtrak Station;Dearborn;United States;;\N;42.312222;-83.198611;591;-5;A;America/New_York
9197;Merritt Island Airport;Cocoa;United States;COI;\N;28.3416111;-80.6854722;6;-5;A;America/New_York
9198;Valkaria Municipal;Valkaria;United States;X59;\N;27.9608611;-80.5583333;26;-5;A;America/New_York
9199;Space Coast Reg'l Airport;Titusville;United States;TIX;\N;28.5148;-80.7992278;34;-5;A;America/New_York
9200;Sebastian Municipal;Sebastian;United States;X26;\N;27.81325;-80.4955833;21;-5;A;America/New_York
9201;Terrace Bay;Terrace Bay;Namibia;TCY;\N;-19.9705556;13.0244444;0;1;N;Africa/Windhoek
9202;Nxabega Airstrip;Nxabega;Botswana;;FB57;-19.5147;22.77665;3100;2;U;Africa/Gaborone
9203;Kelani River-Peliyagoda Waterdrome;Colombo;Sri Lanka;KEZ;\N;6.967463700288899;79.88197231301456;28;5.5;N;Asia/Colombo
9204;Polgolla Reservoir;Kandy;Sri Lanka;KDZ;\N;7.325125009685146;80.64211571225314;1472;5.5;N;Asia/Colombo
9205;Bumi;Bumi Hills;Zimbabwe;;FVBM;-16.81685727384918;28.348887920401467;1633;2;N;Africa/Harare
9206;Priv;Lagos;Portugal;;\N;37.096744;-8.663535;0;0;E;Europe/Lisbon
9207;Warnervale Airport;Warnervale Airport;Australia;;YWVA;-33.240833;151.427778;25;10;U;Australia/Sydney
9208;Boat;Hati;Haiti;;\N;19.15;-73;0;-5;N;America/Port-au-Prince
9209;Boat;Hati;Haiti;;\N;19.1;-73;0;-5;N;America/Port-au-Prince
9210;Cartitas;Hati;Haiti;;\N;18.55;-72.3;50;-5;N;America/Port-au-Prince
9211;Boat;Hati;Haiti;;\N;19;-73;0;-5;N;America/Port-au-Prince
9212;Hospital;Port Au Prince;Haiti;;\N;18.5501;-72.3001;50;-5;N;America/Port-au-Prince
9213;Clinic;Saint Marc;Haiti;;\N;19.1001;-72.7001;30;-5;N;America/Port-au-Prince
9214;Village;Dodard;Haiti;;\N;19.0661;-72.3001;2639;-5;N;America/Port-au-Prince
9215;Bouarfa International;Bouarfa;Morocco;;GMFB;32.518889;-1.9775;3564;0;N;Africa/Casablanca
9216;Beijing West Railway Station;Beijing;China;;\N;39.8955463;116.3201722;167;8;N;Asia/Chongqing
9217;Hung Hom Railway Station;Hong Kong;Hong Kong;;\N;22.303067;114.18152699999996;39;8;N;Asia/Hong_Kong
9218;NYERI;NYERI;Kenya;NYE;HKNI;-0.34;36.91;5830;3;U;Africa/Nairobi
9219;Angaga Seaplane Dock;Alifu Dhaalu Atoll;Maldives;;\N;3.652859;72.824086;2;5;N;Indian/Maldives
9220;Madoogali Seaplane Dock;Alifu Atoll;Maldives;;\N;4.0956;72.753439;2;5;N;Indian/Maldives
9221;Effingham Memorial Airport;Effingham;United States;1H2;KEFH;39.07;-88.534;585;-6;A;America/Chicago
9222;Stoke;Stoke-on-Trent;United Kingdom;;\N;53;2.18333;500;0;E;\N
9223;Tobermory Dock;Tobermory;Canada;;\N;45.2578121;-81.6637838;1335;-5;A;America/Toronto
9224;South Baymouth Docks;South Baymouth;Canada;;\N;45.5564479;-82.0071535;1335;-5;A;America/Toronto
9225;Andrau Airport;Houston;United States;AAP;KAAP;29.43;-95.35;80;-6;A;America/Chicago
9226;Flying Cloud Airport;Eden Prairie;United States;FCM;KFCM;44.4938;-93.2726;906;-6;A;America/Chicago
9227;Likoma Island Airport;Likoma Island;Malawi;LIX;FWLK;-12.083;34.733001;1600;2;N;Africa/Blantyre
9228;Johnson County Airport;Olathe;United States;OJC;KOJC;38.5051;-94.4415;1096;-6;A;America/Chicago
9229;Sigiriya Airport;Sigiriya;Sri Lanka;GIU;VCCS;7.57212;80.43412;630;5.5;N;Asia/Colombo
9230;Westbahnhoff;Vienna;Austria;XWW;\N;48.196667;16.337778;600;1;E;Europe/Vienna
9231;York Mills GO Bus Terminal;York Mills;Canada;;\N;43.745;-79.406389;449;-5;A;America/Toronto
9232;Neumuenster;Neumuenster;Germany;EUM;EDHN;54.079069;9.941719;69;1;E;Europe/Berlin
9233;Tak;Tak;Thailand;TKT;VTPT;16.89611;99.25361;478;7;N;Asia/Bangkok
9234;Marshall Aiport;Marshall;United States;;\N;44.4469;95.7883;354;8;A;Asia/Ulaanbaatar
9235;Meedhuffushi Seaplane dock;Meedhuffushi Island;Maldives;;\N;3.003899;73.000696;0;5;U;Indian/Maldives
9236;Lake Simcoe;Barrie-Orillia;Canada;YLS;CYLS;44.487778;-79.559444;972;-5;A;America/Toronto
9237;Huronia;Midland;Canada;YEE;CYEE;44.684722;-79.929167;773;-5;A;America/Toronto
9238;Markham;Markham;Canada;NU8;CNU8;43.935278;-79.263333;807;-5;A;America/Toronto
9239;Stanhope;Haliburton;Canada;ND4;CND4;45.110833;-78.64;1066;-5;A;America/Toronto
9240;Lindsay;Lindsay;Canada;NF4;CNF4;44.364722;-78.783889;895;-5;A;America/Toronto
9241;Niagara District;Saint Catherines;Canada;YCM;CYSN;43.191667;-79.171111;321;-5;A;America/Toronto
9242;Apopka;Orlando;United States;X04;\N;28.707222;-81.581667;143;-5;A;America/New_York
9243;Edenvale;Edenvale;Canada;;CNV8;44.438889;-79.965278;718;-5;A;America/Toronto
9244;Orillia;Orillia;Canada;;CNJ4;44.683333;-79.316667;725;-5;A;America/Toronto
9245;Holland Landing;Holland Landing;Canada;;CLA4;44.089444;-79.495;855;-5;A;America/Toronto
9246;Kent;Chatham;Canada;XCM;CNZ3;42.306111;-82.081667;645;-5;A;America/Toronto
9247;Parry Sound;Parry Sound;Canada;YPD;CNK4;45.253889;-79.827778;830;-5;A;America/Toronto
9248;Saugeen;Hanover;Canada;;CYHS;44.158056;-81.063333;940;-5;A;America/Toronto
9249;Brandywine Airport;West Goshen Township;United States;OQN;KOQN;39.5924;-75.3455;466;-5;A;America/New_York
9250;Batoka;Livingston;Zambia;;\N;-17.53;25.5;3235;2;U;Africa/Lusaka
9251;Manassas;Manassas;United States;MNZ;KHEF;38.721389;-77.515556;192;-5;A;America/New_York
9252;Texas Gulf Coast Regional;Angleton;United States;;KLBX;29.1086389;-95.4620833;25;-6;A;America/Chicago
9253;Bubovice;Bubovice;Czech Republic;;LKBU;49.974422;14.180801;1385;1;E;Europe/Prague
9254;Shenzhen Railway Station;Shenzhen;China;;\N;22.5344;114.1121;67;8;N;Asia/Chongqing
9255;Rakkestad Flyplass;Rakkestad;Norway;;ENRK;59.3975;11.3469;381;1;E;Europe/Oslo
9256;Johnstown Amtrak;Johnstown;United States;;\N;40.3297222;-78.9222222;1142;-5;A;America/New_York
9257;Harvey Field S43;Snohomish WA;United States;;\N;47.904451;-122.102529;22;-8;A;America/Los_Angeles
9258;Urgup;Urgup;Turkey;;\N;38.606239;34.860276;4000;2;E;Europe/Istanbul
9259;Dolo Ado;Dolo Ado;Ethiopia;;\N;4.175858;42.034083;630;3;U;Africa/Addis_Ababa
9260;AYaou;Yaou;French Guiana;;\N;3.724286;-53.974698;0;-3;N;America/Cayenne
9261;Camop;Camopi;French Guiana;;\N;3.17004;-52.337801;0;-3;N;America/Cayenne
9262;CApat;Croisee d Apatou;French Guiana;;\N;5.04305;-54.0425;0;-3;N;America/Cayenne
9263;Citro;Citron;French Guiana;;\N;4.730072;-53.95158;0;-3;N;America/Cayenne
9264;DieuM;Dieu Merci;French Guiana;;\N;4.782964;-53.260539;0;-3;N;America/Cayenne
9265;Dorli;Dorlin;French Guiana;;\N;3.7353;-53.5287;0;-3;N;America/Cayenne
9266;Elyse;Elysee;French Guiana;;\N;4.75;-54.05;0;-3;N;America/Cayenne
9267;Herap;Morsbach;France;;\N;49.160592;6.864892;0;1;E;Europe/Paris
9268;OSMSE;Saint Elie;French Guiana;;\N;4.81444;-53.28833;0;-3;N;America/Cayenne
9269;Prevo;Remire;French Guiana;;\N;4.89333;-52.29222;0;-3;N;America/Cayenne
9270;Sisli Belediyesi Evlendirme Dairesi Heliport;Istanbul;Turkey;;\N;41.0410805;28.9923138;100;2;E;Europe/Istanbul
9271;Kadikoy Ispark Heliport;Istanbul;Turkey;;\N;40.989305;29.018833;100;2;E;Europe/Istanbul
9272;Bingol;Bingol;Turkey;BGG;LTCU;38.86111;40.5925;3490;2;E;Europe/Istanbul
9273;Uzunyazi;Kastamonu;Turkey;KFS;LTAL;41.316944;33.796111;3520;2;E;Europe/Istanbul
9274;Okawango Lodge;Okawango;Botswana;;\N;-19.25;23.09;3100;2;N;Africa/Gaborone
9275;Kai Tak;Hong Kong;Hong Kong;;HKGX;22.317883;114.202108;50;8;U;Asia/Hong_Kong
9276;Elstree;Elstree;United Kingdom;;EGTR;51.4807;-0.00093;250;0;E;Europe/London
9277;Sandtoft Airfield;Sandtoft;United Kingdom;;EGCF;53.559444;-0.858332;11;0;E;Europe/London
9278;Bourne Park;Hurtsbourne Tarrant;United Kingdom;;\N;51.264449;-1.461975;550;0;E;Europe/London
9279;Hite Airport;Hanksville;United States;UT3;\N;37.8916547;-110.3840289;3840;-7;A;America/Denver
9280;Dieppe Ferry Port;Dieppe;France;;\N;49.933616;1.088039;0;1;E;Europe/Paris
9281;Newhaven Ferry Port;Newhaven;United Kingdom;;\N;50.790385;0.054551;0;0;E;Europe/London
9282;Ollantaytambo Train Station;Ollantaytambo;Peru;;\N;-13.263155;-72.270291;9160;-5;N;America/Lima
9283;Aguas Calientes Train Station;Aguas Calientes;Peru;;\N;-13.1556;-72.5238;6693;-5;N;America/Lima
9284;Poroy Train Station;Cusco;Peru;;\N;-13.4944;-72.0421;11710;-5;N;America/Lima
9285;Aberdeen Harbour;Aberdeen;United Kingdom;;\N;57.13967;-2.07467;0;0;E;Europe/London
9286;Lerwick Harbour;Lerwick;United Kingdom;;\N;60.15;-1.13333;0;0;E;Europe/London
9287;Torshavn Harbour;Torshavn;Faroe Islands;;\N;62.0057;-6.7717;0;0;E;Atlantic/Faeroe
9288;Seydisfjordur;Seydisfjordur;Iceland;;\N;65.26016;-14.00905;0;0;E;Atlantic/Reykjavik
9289;Queen Street Station;Glasgow;United Kingdom;GLQ;\N;55.8622;-4.2512;50;0;E;Europe/London
9290;Torit;Torit;South Sudan;;HSTR;4.411111;32.578889;2018;3;N;Africa/Juba
9291;Falmouth Port;Falmouth;Jamaica;;\N;18.492762;-77.649675;0;-5;N;America/Jamaica
9292;Cozumel Port;Cozumel;Mexico;;\N;20.510439;-86.94925;0;-6;A;America/Mexico_City
9293;Philipsburg Port;Philipsburg;Netherlands Antilles;;\N;18.012868;-63.045242;0;-4;N;America/Curacao
9294;Basseterre Port;Basseterre;Saint Kitts and Nevis;;\N;17.290464;-62.709473;0;-4;N;America/St_Kitts
9295;Charlotte Amalie Port;Charlotte Amalie;Virgin Islands;;\N;18.335431;-64.926824;0;-4;N;America/St_Thomas
9296;San Juan Port;San Juan;Puerto Rico;;\N;18.458992;-66.097693;0;-4;N;America/Puerto_Rico
9297;Santa Barbara Train Station;Santa Barbara;United States;;\N;34.413611;-119.691667;8;-8;A;America/Los_Angeles
9298;Eugene Amtrak Station;Eugene;United States;;\N;44.055062;-123.092372;426;-8;A;America/Los_Angeles
9299;Lincolnville GO Station;Lincolnville;Canada;;\N;43.995278;-79.233333;873;-5;A;America/Toronto
9300;Mormugao-Goa Port;Mormugao;India;;\N;15.4;73.8;0;5.5;N;Asia/Calcutta
9301;Cochin Port;Cochin;India;;\N;9.58;76.14;0;5.5;N;Asia/Calcutta
9302;Penang Cruise Terminal;Penang;Malaysia;;\N;5.418828;100.345804;0;8;N;Asia/Kuala_Lumpur
9303;Port Klang-Kuala Lumpur Cruise Terminal;Port Klang;Malaysia;;\N;2.984629;101.336839;0;8;N;Asia/Kuala_Lumpur
9304;Singapore Cruise Terminal;Singapore;Singapore;;\N;1.264185;103.819821;0;8;N;Asia/Singapore
9305;Shelby County Airport;Shelbyville;United States;2H0;K2H0;39.410556;-88.845556;550;-6;A;America/Chicago
9306;Illinois Terminal;Champaign;United States;;\N;40.1155;-88.2411;732;-6;A;America/Chicago
9307;Mattoon Amtrak;Mattoon;United States;;\N;39.4828;-88.3759;733;-6;A;America/Chicago
9308;Dawu;Lvliang;China;LLV;\N;37.683333;111.142778;3020;8;N;Asia/Chongqing
9309;Delft Central;Delft;Netherlands;;\N;52.006667;4.356667;0;1;E;Europe/Amsterdam
9310;Yading Daocheng;Daocheng;China;DCY;ZUDC;29.3231;100.0533;14472;8;N;Asia/Chongqing
9311;Gannan;Xiahe city;China;GXH;ZLXH;34.4909;102.3719;10466;8;N;Asia/Chongqing
9312;Komandoo;Komandoo;Maldives;;\N;5.494003;73.424183;0;5;N;Indian/Maldives
9313;Komandoo;Komandoo;Maldives;;\N;5.494003;73.424183;0;5;N;Indian/Maldives
9314;Meck Island Airstrip;Meck Island;Marshall Islands;;\N;9;167.726944;10;12;Z;Pacific/Majuro
9315;Sormovo;Nizhny Novgorod;Russia;;XUDS;56.1906;43.4742;210;4;N;Europe/Moscow
9316;Grand Canyon West Airport;Peach Springs;United States;1G4;\N;35.899904;-113.815674;4813;-7;A;America/Phoenix
9317;Vangso;Vangso;Sweden;;ESSZ;59.13333;17.25;10;1;E;Europe/Stockholm
9318;Yaroslavsky Station;Moscow;Russia;;\N;55.776;37.658;400;4;N;Europe/Moscow
9319;Railway Station;Perm;Russia;;\N;58.0203;56.2525;558;6;N;Asia/Yekaterinburg
9320;Railway Station;Novosibirsk;Russia;;\N;55.0188;82.9339;531;7;N;Asia/Omsk
9321;Railway Station;Ulan Bator;Mongolia;;\N;47.92;106.92;4300;8;N;Asia/Ulaanbaatar
9322;Main Station;Beijing;China;;\N;39.9;116.421;130;8;N;Asia/Chongqing
9323;Railway Station;Irkutsk;Russia;;\N;52.3122;104.2958;1305;9;N;Asia/Irkutsk
9324;Kuredu ;Kuredu;Maldives;;\N;5.3244;73.3756;10;5;U;Indian/Maldives
9325;Heliport NUS;Umea;Sweden;;\N;63.817;20.2981;100;1;E;Europe/Stockholm
9326;Heliport SU;Sundsvall;Sweden;;\N;62.4077;17.3024;10;1;E;Europe/Stockholm
9327;Comiso;Comiso;Italy;CIY;LICB;37;14.6144;756;1;E;Europe/Rome
9328;Leroo La Tau;Leroo La Tau;Botswana;;\N;-20.422617;24.517433;1000;2;U;Africa/Gaborone
9329;Xaxanaka;Xaxanaka;Botswana;;\N;-19.198083;23.431383;1000;2;U;Africa/Gaborone
9330;Xugana;Xugana;Botswana;;\N;-19.068;23.10015;1000;2;U;Africa/Gaborone
9331;Nxamaseri;Nxamaseri;Botswana;;\N;-18.597283;21.999267;1000;2;U;Africa/Gaborone
9332;Kwando;Kwando;Botswana;;\N;-18.208717;23.396733;1000;2;U;Africa/Gaborone
9333;Boston North Station;Boston;United States;;\N;42.366235;-71.061122;141;-5;A;America/New_York
9334;Rockport Station;Rockport;United States;;\N;42.6558328;-70.6266667;77;-5;A;America/New_York
9335;Salem Station;Salem;United States;;\N;42.5249995;-70.8958333;26;-5;A;America/New_York
9336;Salem Ferry Dock;Salem;United States;;\N;42.522421;-70.881607;26;-5;A;America/New_York
9337;Long Wharf;Boston;United States;;\N;42.36024;-71.048105;16;-5;A;America/New_York
9338;Joliet Union Station;Joliet;United States;;\N;41.5244;-88.0795;532;-6;A;America/Chicago
9339;Laraway Road;New Lenox;United States;;\N;41.485;-87.9596;703;-6;A;America/Chicago
9340;Lasalle Station;Chicago;United States;;\N;41.875;-87.632;592;-6;A;America/Chicago
9341;Waterloo International;London;United Kingdom;;\N;51.5031;-0.1147;10;0;E;Europe/London
9342;Norfolk Station;Norfolk;United States;;\N;36.843233;-76.275952;16;-5;A;America/New_York
9343;London Railway Station;London;Canada;;\N;42.9819;-81.2464;0;-5;A;America/Toronto
9344;Gloucester Station;Gloucester;United States;;\N;42.61666666666667;-70.66833333333334;50;-5;A;America/New_York
9345;Kitchener Bus Terminal;Kitchener;Canada;;\N;43.44968;-80.49222;1083;-5;A;America/Toronto
9346;Gare du Palais;Quebec;Canada;;\N;46.8174;-71.2139;13;-5;A;America/Toronto
9347;Johnstown Amtrak;Johnstown;United States;;\N;40.3296;-78.9219;1168;-5;A;America/New_York
9348;Enumclaw Airport;Enumclaw;United States;;WA77;47.1956569;-122.0220561;738;-8;A;America/Los_Angeles
9349;Rock Airport;Tarentum;United States;9G1;\N;40.6035463;-79.8261189;1063;-5;A;America/New_York
9350;Grand Central Terminal;New York;United States;;\N;40.752726;-73.977229;70;-5;A;America/New_York
9351;Tremont;New York;United States;;\N;40.847301;-73.89955;200;-5;A;America/New_York
9352;Plum Island Airport;Newburyport;United States;2B2;\N;42.7953611;-70.8394444;11;-5;A;America/New_York
9353;Brooks Camp;Brooks Camp;United States;;\N;58.556539;-155.777311;100;-9;A;America/Anchorage
9354;Kulik Lake Airport;Kulik Lake;United States;LKK;\N;58.96591;-155.108089;1000;-9;U;America/Anchorage
9355;Central Station;Glasgow;United Kingdom;ZGG;\N;55.858;-4.258;25;0;E;Europe/London
9356;Euston Station;London;United Kingdom;QQU;\N;51.5284;-0.1331;89;0;E;Europe/London
9357;Waterloo International;London;United Kingdom;QQW;\N;51.5031;-0.1147;10;0;E;Europe/London
9358;Belorussky Station;Moscow;Russia;;\N;55.7764;37.5803;400;4;N;Europe/Moscow
9359;Railway Station;Omsk;Russia;;\N;54.9833;73.3667;285;7;N;Asia/Omsk
9360;Railway Station;Krasnoyarsk;Russia;;\N;56.0167;93.0667;446;8;N;Asia/Krasnoyarsk
9361;Railway Station;Naushki;Russia;;\N;50.3827;106.1056;1975;9;N;Asia/Irkutsk
9362;Railway Station;Erlian;China;;\N;43.65;111.9833;3159;8;N;Asia/Chongqing
9363;Guangzhoudong Railway Station;Guangzhou;China;;\N;23.149554;113.32479;0;8;U;Asia/Chongqing
9364;Domodedovo;Moscow;Russia;;\N;55.442;37.763;207;4;U;Europe/Moscow
9365;Levashovo;St. Petersburg;Russia;;XLLV;60.08738;30.19152;59;4;E;Europe/Moscow
9366;Kettle Falls Lodge;Minaki;Canada;;\N;50.349812;-94.583456;180;-6;A;America/Winnipeg
9367;Severobajkalsk Port;Severobajkalsk;Russia;;\N;55.4209;109.3225;500;9;U;Asia/Irkutsk
9368;Vanino-Port;Vanino;Russia;;\N;49.093;140.285;20;11;U;Asia/Vladivostok
9369;Holmsk-Port;Holmsk;Russia;;\N;47.040278;142.043056;29;11;U;Asia/Vladivostok
9370;Nogliki;Nogliki;Russia;;\N;51.816667;143.116667;10;11;U;Asia/Vladivostok
9371;Markovo Airport;Markovo;Russia;KVM;UHMO;64.665278;170.414167;89;12;E;Asia/Magadan
9372;Seymchan Airport;Seymchan;Russia;;UHMS;62.919167;152.421667;679;12;E;Asia/Magadan
9373;Zyryanka West Airport;Zyryanka;Russia;ZKP;UESU;65.736389;150.700278;118;12;E;Asia/Magadan
9374;Susuman Airport;Susuman;Russia;;UHMH;62.767222;148.146111;2129;12;E;Asia/Magadan
9375;OYMYAKON;OYMYAKON;Russia;;UEMY;63.45;142.783333;2431;11;E;Asia/Vladivostok
9376;Ust-Maya Airport;Ust-Maya;Russia;UMS;UEMU;60.358056;134.437778;561;10;E;Asia/Yakutsk
9377;Aldan Airport;Aldan;Russia;ADH;UEEA;58.603056;125.407222;2241;10;E;Asia/Yakutsk
9378;Olekminsk Airport;Olekminsk;Russia;;UEMO;60.399167;120.464167;426;10;E;Asia/Yakutsk
9379;Vitim Airport;Vitim;Russia;;UERT;59.457778;112.562778;610;10;E;Asia/Yakutsk
9380;Pellworm;Pellworm;Germany;;EDHP;54.536431;8.67994;10;1;E;Europe/Berlin
9381;Tynda;Tynda;Russia;;\N;55.1512;124.716667;350;10;U;Asia/Yakutsk
9382;Tommot;Tommot;Russia;;\N;58.966667;126.266667;280;10;U;Asia/Yakutsk
9383;Etretat;Etretat;France;;\N;49.701;0.205;102;1;E;Europe/Paris
9384;Le Havre;Le Havre;France;;\N;49.4901;0.1001;13;1;E;Europe/Paris
9385;Weimar;Weimar;Germany;;\N;50.983333;11.316667;208;1;E;Europe/Berlin
9386;Nalati;Xinyuan;China;NLT;ZWNL;43.433056;83.380278;0;8;U;Asia/Chongqing
9387;Port Alsworth Airport;Port alsworth;United States;PTA;PALJ;60.208281;-154.306586;253;-9;A;America/Anchorage
9388;Fontaine Airport;Belfort;France;BOR;LFSQ;47.6556015;7.0108299;3663;1;E;Europe/Paris
9389;Pampulha;Belo Horizonte;Brazil;BHZ;\N;-19.8517;-43.9508;8708;-3;S;America/Sao_Paulo
9390;Fairfield County Airport;Winnsboro;United States;FDW;KFDW;34.315472;-81.108806;577;-5;A;America/New_York
9391;Obock;Obock;Djibouti;OBC;HDOB;11.968333;43.278611;69;3;N;Africa/Djibouti
9392;Tadjoura;Tadjoura;Djibouti;TDJ;HDTJ;11.783;42.917;246;3;N;Africa/Djibouti
9393;Santa Cruz des Quiche Airport;Santa Cruz des Quiche;Guatemala;AQB;MGQC;15.0122004;-91.1505966;6631;-6;S;America/Guatemala
9394;Nordfjordur Airport;Nordfjordur;Iceland;NOR;BINF;65.131897;-13.7463999;3648;0;E;Atlantic/Reykjavik
9395;Bursa Airport;Bursa;Turkey;BTZ;LTBE;40.2332993;29.0091991;772;2;U;Europe/Istanbul
9396;Skyhaven Airport;Rochester;United States;DAW;KDAW;43.2840556;-70.9292778;321;-5;A;America/New_York
9397;Waris Airport;Waris-Papua Island;Indonesia;WAR;WAJR;-3.22243;140.9790039;1224;9;U;Asia/Jayapura
9398;Port Oceanic Airport;Port Oceanic;United States;PRL;\N;60.2083333;-147.8194444;142;-9;A;America/Anchorage
9399;Newton City-County Airport;Newton;United States;EWK;KEWK;38.0570785;-97.2752278;1533;-6;A;America/Chicago
9400;La Ferte Alais Airport;La Ferte Alais;France;;LFFQ;48.497799;2.34333;453;1;E;Europe/Paris
9401;Angguruk airstrip;Angguruk;Indonesia;;\N;-4.2001;139.4329;4200;9;U;Asia/Jayapura
9402;Bairnsdale Airport;Bairnsdale;Australia;BSJ;YBNS;-37.8875;147.567778;165;10;O;Australia/Hobart
9403;Bolton Field;Columbus;United States;TZR;KTZR;39.901111;-83.136944;905;-5;A;America/New_York
9404;Ocean Shores Municipal;Ocean Shores;United States;W04;KW04;46.5995;-124.0854;15;-8;A;America/Los_Angeles
9405;Packwood;Packwood;United States;55S;K55S;46.3625;-121.4067;1057;-8;A;America/Los_Angeles
9406;Fort Bridger;Fort Bridger;United States;FBR;KFBR;41.236;-110.2436;7038;-7;A;America/Denver
9407;Prosser;Prosser;United States;S40;KS40;46.128;-119.4773;705;-8;A;America/Los_Angeles
9408;Chehalis-Centralia;Chehalis;United States;CLS;KCLS;46.4062;-122.5897;177;-8;A;America/Los_Angeles
9409;Desert Aire;Mattawa;United States;M94;KM94;46.4124;-119.5518;586;-8;A;America/Los_Angeles
9410;Lebanon State;Lebanon;United States;S30;KS30;44.3179;-122.5577;344;-8;A;America/Los_Angeles
9411;Evanston-Uinta CO Burns Fld;Evanston;United States;EVW;KEVW;41.1649;-111.0208;7143;-7;A;America/Denver
9412;Sabetha Municipal;Sabetha;United States;K83;KK83;39.5425;-95.4677;1330;-6;A;America/Chicago
9413;Mount Pleasant Regional-Faison Field;Mount Pleasant;United States;LRO;KLRO;32.5387;-79.4697;12;-5;A;\N
9414;Jimmy Carter Regional;Americus;United States;ACJ;KACJ;32.0665;-84.1133;468;-5;A;America/New_York
9415;Weedon Field;Eufala;United States;EUF;KEUF;31.5708;-85.0774;285;-6;A;America/Chicago
9416;Saluda County;Saluda;United States;6J4;K6J4;33.5561;-81.4768;539;-5;A;America/New_York
9417;Dare County Regional;Manteo;United States;MQI;KMQI;35.5514;-75.4173;13;-5;A;America/New_York
9418;Auburn University Regional;Auburn;United States;AUO;KAUO;32.3691;-85.2604;777;-6;A;America/Chicago
9419;Tri-Cities;Endicott;United States;CZG;KCZG;42.0471;-76.0578;833;-5;A;America/New_York
9420;Isla Culebrita;Culebrita;Puerto Rico;;\N;18.31206;-65.228405;10;-4;U;America/Puerto_Rico
9421;Culebrita;Isla Culebrita;Puerto Rico;;\N;18.31206;-65.228405;10;-4;U;America/Puerto_Rico
9422;Apollo Bay;Apollo Bay;Australia;;\N;-38.758827;143.676624;10;10;U;Australia/Hobart
9423;Zabrat;Baku;Azerbaijan;;ZABR;40.483;49.967;0;4;U;Asia/Baku
9424;Chirag;Caspian Sea;Azerbaijan;;CHIR;40.283;51.283;0;4;U;\N
9425;Kendal Glider Port;Kendall;United States;;\N;25.603889;-80.585556;10;-5;U;America/New_York
9426;Lipa Noi Pier;Lipa Noi;Thailand;;\N;9.481772;99.926558;10;7;N;Asia/Bangkok
9427;Don Sak Pier;Don Sak;Thailand;;\N;9.321146;99.735693;10;7;N;Asia/Bangkok
9428;Bessemer;Bessemer;United States;EKY;KEKY;33.1876;-86.5558;700;-6;A;America/Chicago
9429;Colorado Springs East;Ellicott;United States;A50;KA50;38.87;-104.41;6145;-7;A;America/Denver
9430;Olds Didsbury Airport;Didsbury;Canada;;\N;51.710833;-114.106111;3360;-7;A;America/Edmonton
9431;Sundre Airport;Sundre;Canada;;\N;51.774722;-114.680833;3663;-7;A;America/Edmonton
9432;Buffalo-Depew Station;Depew;United States;;\N;42.9074;-78.7266;676;-5;A;America/New_York
9433;South Bend Station;South Bend;United States;;\N;41.678;-86.2876;722;-5;A;America/New_York
9434;Carroll Avenue NICTD Station;Michigan City;United States;;\N;41.7133;-86.8677;630;-6;A;America/Chicago
9435;Chicago Millennium Station;Chicago;United States;;\N;41.8837;-87.623;620;-6;A;America/Chicago
9436;St. Louis Gateway Transportation Center;St. Louis;United States;;\N;38.6239;-90.2039;476;-6;A;America/Chicago
9437;Kansas City Union Station;Kansas City;United States;;\N;39.0865;-94.586;791;-6;A;America/Chicago
9438;Alvarado Transportation Center;Albuquerque;United States;;\N;35.0826;-106.6474;4970;-7;A;America/Denver
9439;San Francisco 4th and King St. Station;San Francisco;United States;;\N;37.7763884;-122.3944444;15;-8;A;America/Los_Angeles
9440;Oakland Jack London Square Station;Oakland;United States;;\N;37.7935;-122.2719;13;-8;A;America/Los_Angeles
9441;Downtown Mountain View Station;Mountain View;United States;;\N;37.3943934;-122.075872;15;-8;A;America/Los_Angeles
9442;San Mateo Station;San Mateo;United States;;\N;37.5683328;-122.3241667;15;-8;A;America/Los_Angeles
9443;Sunnyvale Station;Sunnyvale;United States;;\N;37.3786106;-122.0308333;15;-8;A;America/Los_Angeles
9444;Millbrae Intermodal Terminal;San Mateo;United States;;\N;37.6002772;-122.3866667;15;-8;A;America/Los_Angeles
9445;Cleveland Greyhound Station;Cleveland;United States;;\N;41.5034;-81.6822;679;-5;A;America/New_York
9446;Issyk-Kul International Airport;Tamchy;Kyrgyzstan;ИКУ;UCFL;42.3527;76.4278;5426;6;N;Asia/Bishkek
9447;Crystal Airport;Crystal;United States;MIC;KMIC;45.0343;-93.2114;869;-6;A;America/Chicago
9448;Clarke CO;Quitman;United States;23M;K23M;32.0517;-88.4434;320;-6;A;America/Chicago
9449;Ellaidhoo;Ellaidhoo;Maldives;;\N;4.005694444444444;72.9486111111111;0;5;N;Indian/Maldives
9450;W H Barron Field;Dublin;United States;DBN;KDBN;32.3388;-82.591;311;-5;A;America/New_York
9451;Port Authority;New York;United States;;\N;40.756667;-73.991111;33;-5;A;America/New_York
9452;Bus Terminal;Allentown;United States;;ATWN;40.605604;-75.462384;338;-5;A;America/New_York
9453;Binghamton Bus Terminal;Binghamton;United States;;BING;42.101438;-75.910431;850;-5;A;America/New_York
9454;William F. Walsh Regional Transportation Center;Syracuse;United States;;SYTC;43.0767;-76.1691;380;-5;A;America/New_York
9455;Union Station;Utica;United States;UCA;\N;43.104167;-75.223333;456;-5;A;America/New_York
9456;Clinton Bus Stop;Clinton;United States;;CLTN;43.048889;-75.380278;604;-5;A;America/New_York
9457;Heuston Station;Dublin;Ireland;;\N;53.3463;-6.2927;19;0;E;Europe/Dublin
9458;Killarney Station;Killarney;Ireland;;\N;52.5089;-9.5015;144;0;U;Europe/Dublin
9459;Susquehanna Trailways Terminal;Williamsport;United States;;WSPT;41.239395;-77.000245;518;-5;A;America/New_York
9460;Martz Trailways Bus Terminal;Wilkes-Barre;United States;;WKBR;41.243958;-75.882365;525;-5;A;America/New_York
9461;Hazelton Bus Station;Hazelton;United States;;HZLT;40.954055;-75.976102;1689;-5;A;America/New_York
9462;Lewisburg Bus Stop;Lewisburg;United States;;LWSB;40.963889;-76.888056;456;-5;A;America/New_York
9463;Madison NJT Station;Madison;United States;;MDSN;40.757028;-74.415194;266;-5;A;America/New_York
9464;Maplewood NJT Station;Maplewood;United States;;MPLD;40.731111;-74.275556;115;-5;A;America/New_York
9465;Harrisburg Transportation Center;Harrisburg;United States;;HBRG;40.26223;-76.878709;320;-5;A;America/New_York
9466;Metro-North Station;Poughkeepsie;United States;;PKPS;41.706516;-73.937774;180;-5;A;America/New_York
9467;Port Everglades;Fort Lauderdale;United States;;FTLD;26.087679;-80.115942;0;-5;A;America/New_York
9468;Crown Bay;Charlotte Amalie;Virgin Islands;;STCB;18.33482;-64.952187;0;-4;N;America/St_Thomas
9469;Cruise Ship Dock;Roseau;Dominica;;RSAU;15.296416;-61.38781;0;-4;N;America/Dominica
9470;Grenada Port Authority Cruise Ship Terminal;St. George's;Grenada;;GREN;12.050773;-61.754982;0;-4;N;America/Grenada
9471;Martinique Cruise Terminal;Fort de France;Martinique;;MTQE;14.601432;-61.073111;0;-4;N;America/Martinique
9472;Dr AC Wathey Cruise Pier;Philipsburg;Netherlands Antilles;;PHBG;18.011475;-63.046401;0;-4;N;America/Curacao
9473;Half Moon Cay;Little San Salvador Island;Bahamas;;HMCY;24.574135;-75.955868;0;-5;U;America/Nassau
9474;Pukarua Airport;Pukarua;French Polynesia;PUK;NTGQ;-18.3166666;-137.016666;6;-10;N;\N
9475;Kabale;Kabale;Uganda;;HUKB;-1.223333;29.96;5850;3;N;Africa/Kampala
9476;Cruz Bay Ferry Dock;Cruz Bay;Virgin Islands;;CBFD;18.331263;-64.795408;0;-4;A;America/St_Thomas
9477;Red Hook Bay Ferry Dock;Red Hook;Virgin Islands;;RHVI;18.326479;-64.849361;0;-4;A;America/St_Thomas
9478;Gare Routiere de Fort-de-France;Fort-de-France;Martinique;;FDFM;14.602932;-61.070873;0;-4;U;America/Martinique
9479;Saint Pierre;Saint Pierre;Martinique;;SPMQ;14.746061;-61.176809;0;-4;U;America/Martinique
9480;Philipsburg Bus Stop;Philipsburg;Netherlands Antilles;;PHSM;18.0263;-63.049759;0;-4;N;America/Curacao
9481;Marigot Bus Stop;Marigot;France;;MRGT;18.066039;-63.084385;0;-4;N;\N
9482;Summit NJT Station;Summit;United States;;SMMT;40.716556;-74.35775;374;-5;A;America/New_York
9483;Batoka Airfield;Dambwa - Livingstone;Zambia;;\N;-17.88464;25.84697;2695;2;N;Africa/Lusaka
9484;Santa Justa;Sevilla;Spain;;SVSJ;37.392051;-5.975135;23;1;E;Europe/Madrid
9485;Cordoba Station;Cordoba;Spain;;CRDB;37.888573;-4.789177;390;1;E;Europe/Madrid
9486;Granada Station;Granada;Spain;;GRND;37.184055;-3.60916;2421;1;E;Europe/Madrid
9487;Atocha;Madrid;Spain;;MDAT;40.4075;-3.691667;2188;1;E;Europe/Madrid
9488;Corvallis Muni;Corvallis;United States;CVO;KCVO;44.5067;-123.2915;250;-8;A;America/Los_Angeles
9489;Helipad of Viraj Group;Boisar;India;;\N;19.779;72.805;50;5.5;U;Asia/Calcutta
9490;Heydar Aliyev;Caspian Sea;Azerbaijan;;HEYD;39.86218;50.3959;0;4;E;\N
9491;Phantom Lake;Phantom Lake;Canada;;\N;49.859473;-123.494509;3386;-8;A;America/Vancouver
9492;Tuzla;Tuzla;Romania;;LRTZ;43.9841995;28.6096992;78;2;U;Europe/Bucharest
9493;Toledo Railway Station;Toledo;Spain;;TLDO;39.862363;-4.011083;1736;1;E;Europe/Madrid
9494;Chamartin Station;Madrid;Spain;;MADC;40.472272;-3.68215;2188;1;E;Europe/Madrid
9495;Abando;Bilbao;Spain;;BLBO;43.25947;-2.929017;62;1;E;Europe/Madrid
9496;Aeroport;Barcelona;Spain;;ARPT;41.288436;2.072409;39;1;E;Europe/Madrid
9497;Passeig de Gracia;Barcelona;Spain;;PGRC;41.392186;2.16482;39;1;U;Europe/Madrid
9498;Sants;Barcelona;Spain;;BSNT;41.379359;2.140474;39;1;E;Europe/Madrid
9499;Railway Station;Figueres;Spain;;FGRS;42.265065;2.968822;128;1;E;Europe/Madrid
9500;Martz Trailways;Scranton;United States;;SCRT;41.410253;-75.670297;745;-5;A;America/New_York
9501;New Brunswick Station;New Brunswick;United States;;NBWK;40.496432;-74.446447;62;-5;A;America/New_York
9502;Trenton Transit Center;Trenton;United States;;TRTC;40.218839;-74.754279;49;-5;A;America/New_York
9503;Holmesburg Jct Station;Philadelphia;United States;;HOLM;40.032679;-75.023718;39;-5;U;America/New_York
9504;White Plains;White Plains;United States;;WPLN;41.032549;-73.775069;213;-5;A;America/New_York
9505;Southeast;Southeast;United States;;STHT;41.413179;-73.623551;338;-5;A;America/New_York
9506;B Street Cruise Terminal;San Diego;United States;;SANB;32.716861;-117.174296;0;-8;A;America/Los_Angeles
9507;Terminal Maritima;Puerto Vallarta;Mexico;;PRTV;20.657616;-105.241082;0;-6;A;America/Mexico_City
9508;Bahia San Lucas;Cabo San Lucas;Mexico;;CABO;22.880854;-109.902117;0;-7;A;America/Mazatlan
9509;Ferry Terminal;Anacortes;United States;;ACRT;48.50728;-122.677596;0;-8;A;America/Los_Angeles
9510;Ferry Dock;Friday Harbor;United States;;FRID;48.535673;-123.014044;0;-8;A;America/Los_Angeles
9511;Morristown Station;Morristown;United States;;MRST;40.796957;-74.473834;315;-5;A;America/New_York
9512;Port of Seattle;Seattle;United States;;PSEA;47.609178;-122.350605;0;-8;A;America/Los_Angeles
9513;Ferry Terminal;Whittier;United States;;WHIT;60.776566;-148.683367;0;-9;A;America/Anchorage
9514;Ferry Terminal;Valdez;United States;;VLDZ;61.123976;-146.365309;0;-9;A;America/Anchorage
9515;Gare Routiere;Saint-Louis;Reunion;;STLS;-21.289604;55.407269;103;4;E;Indian/Reunion
9516;Bus Stop;La Riviere;Reunion;;LRIV;-21.272574;55.437706;190;4;E;Indian/Reunion
9517;Bus Stop;Cilaos;Reunion;;CILS;-21.135518;55.472087;4000;4;E;Indian/Reunion
9518;Bus Stop;Hell-Bourg;Reunion;;HLBG;-21.06402;55.518453;3000;4;E;Indian/Reunion
9519;Bus Stop;Salazie;Reunion;;SLZI;-21.027485;55.538955;3000;4;E;Indian/Reunion
9520;Gare;St Andre;Reunion;;ANDR;-20.963272;55.652508;100;4;E;Indian/Reunion
9521;L'Ocean Bus Terminal;Saint Denis;Reunion;;SDNS;-20.877683;55.457393;0;4;E;Indian/Reunion
9522;Gare Routiere;Saint Pierre;Reunion;;STPR;-21.33424;55.471331;0;4;E;Indian/Reunion
9523;Baie Ste Anne;Praslin;Seychelles;;PLIN;-4.347119;55.765688;0;4;N;Indian/Mahe
9524;La Passe;La Digue;Seychelles;;LADG;-4.347969;55.829069;0;4;N;Indian/Mahe
9525;Ferry Dock;Bainbridge Island;United States;;BAIN;47.622237;-122.509362;0;-8;A;America/Los_Angeles
9526;Bus;Phnom Penh;Cambodia;;PNMP;11.561716;104.914276;0;7;N;Asia/Phnom_Penh
9527;Bus;Siem Reap;Cambodia;;SMRP;13.361002;103.859543;0;7;N;Asia/Phnom_Penh
9528;Bus;Sihanoukville;Cambodia;;SNKV;10.60722;103.524886;0;7;N;Asia/Phnom_Penh
9529;Bus;Kampot;Cambodia;;KMPT;10.614922;104.177724;0;7;N;Asia/Phnom_Penh
9530;Bus;Kep;Cambodia;;KEPC;10.480174;104.294228;0;7;N;Asia/Phnom_Penh
9531;Main Station;Taipei;Taiwan;;TPEI;25.046176;121.517532;0;8;N;Asia/Taipei
9532;Prominent Hill;Prominent Hill;Australia;PXH;YPMH;-29.716667;135.521667;734;9.5;O;Australia/Adelaide
9533;Chatsworth Station;Chatsworth;United States;CWT;\N;34.256944;-118.598889;978;-8;A;America/Los_Angeles
9534;Algerciras Port;Algerciras;Spain;;\N;36.136;-5.435;0;1;E;Europe/Madrid
9535;Ganges Water Aerodrome;Ganges;Canada;YGG;\N;48.85;-123.5;0;-8;A;America/Vancouver
9536;Pender Harbour Water Aerodrome;Pender Harbour;Canada;YPT;\N;49.616667;-124.016667;0;-8;A;America/Vancouver
9537;Mansons Landing Water Aerodrome;Mansons Landing;Canada;YMU;\N;50.066667;-124.983333;0;-8;A;America/Vancouver
9538;Port McNeill Airport;Port McNeill;Canada;YMP;\N;50.575556;-127.028611;225;-8;A;America/Vancouver
9539;Sullivan Bay Water Aerodrome;Sullivan Bay;Canada;YTG;\N;50.883333;-126.833333;0;-8;A;America/Vancouver
9540;Deer Harbor Seaplane;Deer Harbor;United States;DHB;\N;48.618397;-123.00596;0;-8;A;America/Los_Angeles
9541;San Diego Old Town Transit Center;San Diego;United States;OLT;\N;32.7552;-117.1995;0;-8;A;America/Los_Angeles
`