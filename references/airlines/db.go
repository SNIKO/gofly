package airlines

const csvdb = `
1,Private flight,\N,-,N/A,,,Y
2,135 Airways,\N,,GNL,GENERAL,United States,N
3,1Time Airline,\N,1T,RNX,NEXTIME,South Africa,Y
4,2 Sqn No 1 Elementary Flying Training School,\N,,WYT,,United Kingdom,N
5,213 Flight Unit,\N,,TFU,,Russia,N
6,223 Flight Unit State Airline,\N,,CHD,CHKALOVSK-AVIA,Russia,N
7,224th Flight Unit,\N,,TTF,CARGO UNIT,Russia,N
8,247 Jet Ltd,\N,,TWF,CLOUD RUNNER,United Kingdom,N
9,3D Aviation,\N,,SEC,SECUREX,United States,N
10,40-Mile Air,\N,Q5,MLA,MILE-AIR,United States,Y
11,4D Air,\N,,QRT,QUARTET,Thailand,N
12,611897 Alberta Limited,\N,,THD,DONUT,Canada,N
13,Ansett Australia,\N,AN,AAA,ANSETT,Australia,Y
14,Abacus International,\N,1B,,,Singapore,Y
15,Abelag Aviation,\N,W9,AAB,ABG,Belgium,N
16,Army Air Corps,\N,,AAC,ARMYAIR,United Kingdom,N
17,Aero Aviation Centre Ltd.,\N,,AAD,SUNRISE,Canada,N
18,Aero Servicios Ejecutivos Internacionales,\N,,SII,ASEISA,Mexico,N
19,Aero Biniza,\N,,BZS,BINIZA,Mexico,N
20,Aero Albatros,\N,,ABM,ALBATROS ESPANA,Spain,N
21,Aigle Azur,\N,ZI,AAF,AIGLE AZUR,France,Y
22,Aloha Airlines,\N,AQ,AAH,ALOHA,United States,Y
23,Alaska Island Air,\N,,AAK,ALASKA ISLAND,United States,N
24,American Airlines,\N,AA,AAL,AMERICAN,United States,Y
25,Aviation Management Corporation,\N,,AAM,AM CORP,United States,N
26,Atlantis Airlines (USA),\N,,AAO,ATLANTIS AIR,United States,N
27,Aerovista Airlines,\N,,AAP,AEROVISTA GROUP,United Arab Emirates,N
28,Asiana Airlines,\N,OZ,AAR,ASIANA,Republic of Korea,Y
29,Askari Aviation,\N,4K,AAS,AL-AAS,Pakistan,Y
30,Australia Asia Airlines,\N,,AAU,AUSTASIA,Australia,N
31,Astro Air International,\N,,AAV,ASTRO-PHIL,United States,N
32,Afriqiyah Airways,\N,8U,AAW,AFRIQIYAH,Libya,Y
33,Afrinat International Airlines,\N,Q9,AFU,,Gambia,N
34,Afric'air Express,\N,,AAX,AFREX,Ivory Coast,N
35,Allegiant Air,\N,G4,AAY,ALLEGIANT,United States,Y
36,Angus Aviation,\N,,AAZ,ANGUS,Canada,N
37,Artem-Avia,\N,,ABA,ARTEM-AVIA,Ukraine,N
38,African Business and Transportations,\N,,ABB,AFRICAN BUSINESS-,Democratic Republic of the Congo,N
39,Aban Air,\N,K5,ABE,ABAN,Iran,n
40,Aerial Oy,\N,,ABF,SKYWINGS,Finland,N
41,Abakan-Avia,\N,,ABG,ABAKAN-AVIA,Russia,N
42,ABSA - Aerolinhas Brasileiras,\N,M3,TUS,ABSA Cargo,Brazil,Y
43,Abaet,\N,,ABJ,Abaet,Brazil,N
44,Alberta Citylink,\N,,ABK,ALBERTA CITYLINK,Canada,N
45,APSA Colombia,\N,,ABO,AEROEXPRESO,Colombia,N
46,Aerovias Bueno,\N,,ABU,AEROBUENO,Colombia,N
47,Aerocenter,\N,,ACR, Escuela de Formacion de Pioltos Privados de Avion,AEROCENTER,N
48,Antrak Air,\N,04,ABV,ANTRAK,Ghana,N
49,Airborne Express,\N,GB,ABX,ABEX,United States,N
50,ABX Air,\N,GB,ABX,ABEX,United States,N
51,ATA Brasil,\N,,ABZ,ATA-BRAZIL,Brazil,N
52,Avcard Services,\N,,ACC,,United Kingdom,N
53,Academy Airlines,\N,,ACD,ACADEMY,United States,N
54,Aero Comondu,\N,,ACO,AERO COMONDU,Mexico,N
55,Astral Aviation,\N,8V,ACP,ASTRAL CARGO,Kenya,Y
56,Air Cess,\N,,ACS,,Liberia,N
57,Air Aurora,\N,,AAI,BOREALIS,United States,N
58,Air Cargo Transportation System,\N,,ACU,AFRISPIRIT,Kenya,N
59,Air Charter Service,\N,,ACV,,United Kingdom,N
60,Aero Asia International,\N,E4,RSO,AERO ASIA,Pakistan,N
61,Air Charters,\N,,ACX,PARAIR,Canada,N
62,Air Togo,\N,YT,TGA,AIR TOGO,Togo,N
63,Air Somalia,\N,,RSM,AIR SOMALIA,Somali Republic,N
64,Atlas Cargo Lines,\N,,ACY,ATLAS CARGOLINES,Morocco,N
65,Airservices Australia,\N,,ADA,AUSCAL,Australia,N
66,Advance Leasing Company,\N,4G,,,United States,N
67,Aztec Worldwide Airlines,\N,7A,,,United States,N
68,Air Tindi,\N,8T,,,Canadian Territories,Y
69,Antonov Airlines,\N,,ADB,ANTONOV BUREAU,Ukraine,N
70,Air Atlantic Dominicana,\N,,ADC,ATLAN-DOMINICAN,Dominican Republic,N
71,Advanced Air Co.,\N,,ADD,,Japan,N
72,Ada Air,\N,ZY,ADE,ADA AIR,Albania,Y
73,Aerea Flying Training Organization,\N,,ADG,AEREA TRAINING,Spain,N
74,Audeli Air,\N,,ADI,AUDELI,Spain,N
75,Abicar,\N,,ADJ,ABICAR,Ivory Coast,N
76,ADC Airlines,\N,Z7,ADK,ADCO,Nigeria,N
77,Aero Dynamics,\N,,ADL,COTSWOLD,United Kingdom,N
78,Aerolineas Dominicanas,\N,,ADM,DOMINAIR,Dominican Republic,N
79,Aerodienst GmbH,\N,,ADN,AERODIENST,Germany,N
80,Aerodiplomatic,\N,,ADP,AERODIPLOMATIC,Mexico,N
81,Aerodyne,\N,,ADY,AERODYNE,United States,N
82,Avion Taxi,\N,,ADQ,AIR DATA,Canada,N
83,Adria Airways,\N,JP,ADR,ADRIA,Slovenia,Y
84,Aviones de Sonora,\N,,ADS,SONORAV,Mexico,N
85,Air Dorval,\N,,ADT,AIR DORVAL,Canada,N
86,Airdeal Oy,\N,,ADU,AIRDEAL,Finland,N
87,Advance Air Charters,\N,,ADV,ADVANCE,Canada,N
88,Air Andaman,\N,,ADW,AIR ANDAMAN,Thailand,N
89,Anderson Aviation,\N,,ADX,ANDAX,United States,N
90,Air Europa,\N,UX,AEA,EUROPA,Spain,Y
91,Air Southwest Ltd.,\N,,ASW,AIRSOUTHWEST,Canada,N
92,Air Special,\N,,ASX,AIRSPEC,Czech Republic,N
93,Aero Benin,\N,EM,AEB,AEROBEN,Benin,Y
94,Aerocesar,\N,,AEC, Aerovias Del Cesar,AEROCESAR,N
95,Aerotrans Airlines,\N,,AED,,Russia,N
96,Aegean Airlines,\N,A3,AEE,AEGEAN,Greece,Y
97,Aerofumigaciones Sam,\N,,AEG,FUMIGACIONES SAM,Chile,N
98,Aeroexpreso Interamericano,\N,,AEI,INTERAM,Colombia,N
99,Air Express,\N,,AEJ,KHAKI EXPRESS,Tanzania,N
100,Aerocon,\N,,AEK,AEROCON,Bolivia,N
101,Aero Madrid,\N,,AEM,AEROMADRID,Spain,N
102,Aeroenlaces Nacionales,\N,,AEN,,Mexico,N
103,Aeroservicios Ejecutivos Del Occidente,\N,,AEO,AERO OCCIDENTE,Mexico,N
104,Aerotec,\N,,AEP,AEROTEC,Spain,N
105,Air Atlantique,\N,KI,AAG,ATLANTIC,United Kingdom,N
106,Air Europe,\N,PE,AEL,AIR EUROPE,Italy,Y
107,Air Alma,\N,,AAJ,AIR ALMA,Canada,N
108,Air Express,\N,,AEQ,LUNA,Sweden,N
109,Alaska Central Express,\N,KO,AER,ACE AIR,United States,Y
110,ACES Colombia,\N,,AES,ACES,Colombia,Y
111,Aeronautical Radio of Thailand,\N,,AET,AEROTHAI,Thailand,N
112,Astraeus,\N,5W,AEU,FLYSTAR,United Kingdom,Y
113,Aeroventas,\N,,AEV,AEROVENTAS,Mexico,N
114,Aerosvit Airlines,\N,VV,AEW,AEROSVIT,Ukraine,Y
115,Airway Express,\N,,AEX,AVCO,United States,N
116,Air Italy,\N,I9,AEY,AIR ITALY,Italy,Y
117,Aerial Transit Company,\N,,AEZ,AERIAL TRANZ,United States,N
118,Alfa Air,\N,,AFA,BLUE ALFA,Czech Republic,N
119,American Falcon,\N,WK,AFB,AMERICAN FALCON,Argentina,N
120,Alliance Airlines,\N,QQ,UTY,UNITY,Australia,Y
121,Air Universal,\N,,UVS,UNI-LEONE,Sierra Leone,N
122,Auvia Air,\N,,UVT,AUVIA,Indonesia,N
123,African West Air,\N,,AFC,AFRICAN WEST,Senegal,N
124,Airfast Indonesia,\N,,AFE,AIRFAST,Indonesia,N
125,Ariana Afghan Airlines,\N,FG,AFG,ARIANA,Afghanistan,Y
126,Air Fecteau,\N,,AFH,FECTO,Canada,N
127,Africaone,\N,,AFI,AFRICAWORLD,Uganda,N
128,Alliance,\N,,AFJ,JAMBO,Uganda,N
129,Africa Air Links,\N,,AFK,AFRICA LINKS,Sierra Leone,N
130,Aeroflot Russian Airlines,\N,SU,AFL,AEROFLOT,Russia,Y
131,Aero Empresa Mexicana,\N,,AFO,AERO EMPRESA,Mexico,N
132,Air Bosna,\N,JA,BON,AIR BOSNA,Bosnia and Herzegovina,Y
133,Air Bravo,\N,,BRF,AIR BRAVO,Uganda,N
134,Air Brasd'or,\N,,BRL,BRASD'OR,Canada,N
135,Air 500,\N,,BRM,BOOMERANG,Canada,N
136,Alba Servizi Aerotrasporti,\N,,AFQ,ALBA,Italy,N
137,Air France,\N,AF,AFR,AIRFRANS,France,Y
138,Air Partner,\N,,ACG,AIR PARTNER,United Kingdom,N
139,Air Caledonie International,\N,SB,ACI,AIRCALIN,France,Y
140,Air Caledonia,\N,,ACM,WEST CAL,Canada,N
141,Air Guam,\N,,AGM,AIR GUAM,United States,N
142,Air Gabon,\N,GN,AGN,GOLF NOVEMBER,Gabon,N
143,Air Data,\N,,AFS,,United Kingdom,N
144,Air Afrique Vacancies,\N,,AFV,AFRIQUE VACANCE,Ivory Coast,N
145,Air Cargo America,\N,,MVM,PEGASUS,United States,N
146,Air Salone,\N,2O,,,Sierra Leone,Y
147,Air-Angol,\N,,NGO,AIR ANGOL,Angola,N
148,Air Nigeria,\N,,NGP,REGAL EAGLE,Nigeria,N
149,Air Cargo Carriers,\N,2Q,SNC,NIGHT CARGO,United States,Y
150,Air Samarkand,\N,,SND,ARSAM,Uzbekistan,N
151,Air Senegal International,\N,V7,SNG,AIR SENEGAL,Senegal,N
152,Air Sandy,\N,,SNY,AIR SANDY,Canada,N
153,Air Namibia,\N,SW,NMB,NAMIBIA,Namibia,Y
154,Air Intersalonika,\N,,NSK,INTERSALONIKA,Greece,N
155,Air Anatolia,\N,,NTL,AIR ANATOLIA,Turkey,N
156,Air Saigon,\N,,SGA,AIR SAIGON,Vietnam,N
157,Afrique Regional Airways,\N,,AFW,AFRAIR,Ivory Coast,N
158,Airfreight Express,\N,,AFX, ,United Kingdom,N
159,Africa Chartered Services,\N,,AFY,AFRICA CHARTERED,Nigeria,N
160,Africa Freight Services,\N,,AFZ,AFREIGHT,Zambia,N
161,Aeronaves Del Centro,\N,,AGA,,Venezuela,N
162,Air Service Gabon,\N,G8,AGB,,Gabon,N
163,Arab Agricultural Aviation Company,\N,,AGC,AGRICO,Egypt,N
164,Atlantic Gulf Airlines,\N,,AGF,ATLANTIC GULF,United States,N
165,Aerolitoral,\N,5D,SLI,COSTERA,Mexico,Y
166,Algoma Airways,\N,,AGG,ALGOMA,Canada,N
167,Altagna,\N,,AGH,ALTAGNA,France,N
168,Angola Air Charter,\N,,AGO,ANGOLA CHARTER,Angola,N
169,Angola Air Charter,\N,,AGO,ANGOLA CHARTER,Angola,N
170,AERFI Group,\N,,AGP,AIR TARA,Ireland,N
171,Aerogala,\N,,AGQ,GALASERVICE,Chile,N
172,Amadeus Global Travel Distribution,\N,1A,AGT,AMADEUS,Spain,N
173,Angara Airlines,\N,,AGU,SARMA,Russia,N
174,Air Glaciers,\N,7T,AGV,AIR GLACIERS,Switzerland,Y
175,Aero Gambia,\N,,AGW,AERO GAMBIA,Gambia,N
176,Aviogenex,\N,,AGX,GENEX,Serbia,Y
177,Atlantic Coast Airlines,\N,,BLR,BLUE RIDGE,United States,N
178,Aero Barloz,\N,,BLZ,AEROLOZ,Mexico,N
179,Aeroper,\N,PL,PLI,Aeroperu,Peru,Y
180,Atlas Blue,\N,8A,BMM,ATLAS BLUE,Morocco,Y
181,Aero Banobras,\N,,BNB,AEROBANOBRAS,Mexico,N
182,Aero Flight Service,\N,,AGY,FLIGHT GROUP,United States,N
183,Agrolet-Mci,\N,,AGZ,AGROLET,Slovakia,N
184,Air Alpha Greenland,\N,GD,AHA,AIR ALPHA,Denmark,N
185,Azal Avia Cargo,\N,,AHC,AZALAVIACARGO,Azerbaijan,N
186,Airport Helicopter Basel,\N,,AHE, Muller & Co.,AIRPORT HELICOPTER,N
187,Aeroservices Corporate,\N,,CJE,BIRD JET,France,N
188,Aspen Helicopters,\N,,AHF,ASPEN,United States,N
189,Aerochago Airlines,\N,,AHG,AEROCHAGO,Dominican Republic,N
190,Airplanes Holdings,\N,,AHH,AIRHOLD,Ireland,N
191,Air Hong Kong,\N,LD,AHK,AIR HONG KONG,Hong Kong,N
192,Aerochiapas,\N,,AHP,AEROCHIAPAS,Mexico,N
193,Air Adriatic,\N,,AHR,ADRIATIC,Croatia,N
194,Air Viggi San Raffaele,\N,,AHS,AIRSAR,Italy,N
195,ABC Air Hungary,\N,,AHU,ABC HUNGARY,Hungary,N
196,Aeromist-Kharkiv,\N,HT,AHW,AEROMIST,Ukraine,N
197,Azerbaijan Airlines,\N,J2,AHY,AZAL,Azerbaijan,Y
198,Avies,\N,U3,AIA,AVIES,Estonia,Y
199,Airbus Industrie,\N,AP,AIB,AIRBUS INDUSTRIE,France,N
200,Alpine Air Chile,\N,,AIH,ALPINE CHILE,Chile,N
201,Air Integra,\N,,AII,INTEGRA,Canada,N
202,ABC Aerolineas,\N,,AIJ,ABC AEROLINEAS,Mexico,N
203,African Airlines International Limited,\N,,AIK,AFRICAN AIRLINES,Kenya,N
204,African International Airways,\N,,AIN,FLY CARGO,Swaziland,N
205,Alpine Air Express,\N,5A,AIP,ALPINE AIR,United States,N
206,Alicante Internacional Airlines,\N,,AIU,ALIA,Spain,N
207,Aba Air,\N,,ABP,BAIR,Czech Republic,N
208,Airblue,\N,ED,ABQ,PAKBLUE,Pakistan,Y
209,Airmark Aviation,\N,,THM,THAI AIRMARK,Thailand,N
210,Airlift International,\N,,AIR,AIRLIFT,United States,Y
211,Airest,\N,,AIT, ,Estonia,N
212,Air Baffin,\N,,BFF,AIR BAFFIN,Canada,N
213,Air Bandama,\N,,BDM,BANDAMA,Ivory Coast,N
214,Air Berlin,\N,AB,BER,AIR BERLIN,Germany,Y
215,Air Brousse,\N,,ABT,AIR BROUSSE,Canada,N
216,Air Contractors,\N,AG,ABR,CONTRACT,Ireland,N
217,Air Illinois,\N,,AIL,AIR ILLINOIS,United States,N
218,Air India Limited,\N,AI,AIC,AIRINDIA,India,Y
219,Air Inter Gabon,\N,,AIG, ,Gabon,N
220,Air Bourbon,\N,ZB,BUB,BOURBON,Reunion,Y
221,Air Atlanta Icelandic,\N,CC,ABD,ATLANTA,Iceland,Y
222,Air Inuit,\N,,AIE,AIR INUIT,Canada,N
223,Air Sureste,\N,,AIS,SURESTE,Spain,N
224,Air Srpska,\N,RB,SBK,Air Srpska,Bosnia and Herzegovina,N
225,Air Tahiti Nui,\N,TN,THT,TAHITI AIRLINES,France,Y
226,Airvias S/A Linhas Aereas,\N,,AIV,AIRVIAS,Brazil,N
227,Aero Services Executive,\N,W4,BES,BIRD EXPRESS,France,N
228,Atlantic Island Airways,\N,,AIW,TARTAN,Canada,N
229,Aircruising Australia,\N,,AIX,CRUISER,Australia,N
230,Aircrew Check and Training Australia,\N,,AIY,AIRCREW,Australia,N
231,Arkia Israel Airlines,\N,IZ,AIZ,ARKIA,Israel,Y
232,A J Services,\N,,AJA,AYJAY SERVICES,United Kingdom,N
233,Aero JBR,\N,,AJB,AERO JBR,Mexico,N
234,Aero Jet Express,\N,,AJE,JET EXPRESS,Mexico,N
235,Avia Consult Flugbetriebs,\N,,AJF,AVIACONSULT,Austria,N
236,Ameristar Jet Charter,\N,,AJI,AMERISTAR,United States,N
237,A2 Jet Leasing,\N,,AJJ,ATLANTIC JET,United States,N
238,Allied Air,\N,,AJK,BAMBI,Nigeria,N
239,Air Jamaica,\N,JM,AJM,JAMAICA,Jamaica,Y
240,Air One,\N,AP,ADH,HERON,Italy,Y
241,Air Sahara,\N,S2,RSH,SAHARA,India,Y
242,Air Malta,\N,KM,AMC,AIR MALTA,Malta,Y
243,Aeroejecutivo,\N,,AJO,AEROEXO,Mexico,N
244,Aero Jets Corporativos,\N,,AJP,AEROJETS,Mexico,N
245,Aeroejecutivos Colombia,\N,,AJS,AEROEJECUTIVOS,Colombia,N
246,Amerijet International,\N,M6,AJT,AMERIJET,United States,N
247,Air Jetsul,\N,,AJU,AIRJETSUL,Portugal,N
248,ANA & JP Express,\N,,AJV,AYJAY CARGO,Japan,N
249,Alpha Jet International,\N,,AJW,ALPHAJET,United States,N
250,Air Japan,\N,NQ,AJX,AIR JAPAN,Japan,Y
251,Ajet,\N,,AJY,AYJET,Cyprus,N
252,Air Korea Co. Ltd.,\N,,AKA,,Republic of Korea,Y
253,Air Livonia,\N,,LIV,LIVONIA,Estonia,N
254,Air BC,\N,,ABL,AIRCOACH,Canada,N
255,Air Fret Senegal,\N,,ABN, ,Senegal,N
256,Air Jamahiriya Company,\N,,LJA,AIR JAMAHIRIYA,Libya,N
257,Aktjubavia,\N,,AKB,KARAB,Kazakhstan,N
258,Arca Aerovias Colombianas Ltda.,\N,,AKC,ARCA,Colombia,N
259,Anikay Air Company,\N,,AKF,ANIKAY,Kyrgyzstan,N
260,Akhal,\N,,AKH,AKHAL,Turkmenistan,N
261,Aeromilenio,\N,,MNI,AEROMIL,Mexico,N
262,Aklak Air,\N,,AKK,AKLAK,Canada,N
263,Air Kiribati,\N,4A,AKL,,Kiribati,Y
264,Alkan Air,\N,,AKN,ALKAN AIR,Canada,N
265,Angkor Airways,\N,,AKW,ANGKORWAYS,Cambodia,N
266,Air Nippon Network Co. Ltd.,\N,EH,AKX,ALFA WING,Japan,N
267,AK Navigator LLC,\N,,AKZ,ABSOLUTE,Kazakhstan,N
268,Aero Albatros,\N,,ALB,ALBATROS,Mexico,N
269,Albion Aviation,\N,,ALD,ALBION,United Kingdom,N
270,Aeroalas Colombia,\N,,ALE,AEROALAS,Colombia,N
271,Allied Command Europe (Mobile Force),\N,,ALF,ACEFORCE,Belgium,N
272,Aerotaxis Albatros,\N,,BTS,AEROLINEAS ALBATROS,Mexico,N
273,American Flyers,\N,,FYS,AMERICAN FLYERS,United States,N
274,Aero Coach Aviation,\N,,DFA,AERO COACH,United States,N
275,Atlantic Southeast Airlines,\N,,CAA,Chandler,United States,N
276,ACM Air Charter,\N,,BVR,BAVARIAN,Germany,N
277,Air Logistics,\N,,ALG,AIRLOG,United States,N
278,Aerovallarta,\N,,ALL,VALLARTA,Mexico,N
279,Air ALM,\N,,ALM,ANTILLEAN,Netherlands Antilles,N
280,Air Lincoln,\N,,ALN,CHICAGO LINCOLN,United States,N
281,America West Airlines,\N,HP,AWE,CACTUS,United States,Y
282,Air Wisconsin,\N,ZW,AWI,AIR WISCONSIN,United States,Y
283,Aerotransporte de Carga Union,\N,,TNO,AEROUNION,Mexico,N
284,Aero Taxis Cessna,\N,,TND,TAXIS CESSNA,Mexico,N
285,Arizona Express Airlines,\N,,TMP,TEMPE,United States,N
286,Tatarstan Airlines,\N,U9,TAK,TATARSTAN,Russia,Y
287,Allegheny Commuter Airlines,\N,,ALO,ALLEGHENY,United States,Y
288,Alpliner AG,\N,,ALP,ALPINER,Switzerland,N
289,Altair Aviation (1986),\N,,ALQ,ALTAIR,Canada,N
290,Air Luxor STP,\N,,ALU,LUXORJET,Sao Tome and Principe,N
291,Alas de Venezuela,\N,,ALV,ALVEN,Venezuela,N
292,Alas Nacionales,\N,,ALW, S.A.,ALNACIONAL,N
293,Alyeska Air Service,\N,,ALY,ALYESKA,United States,N
294,Alta Flights (Charters) Ltd.,\N,,ALZ, ,Canada,N
295,Air Sunshine,\N,,RSI,AIR SUNSHINE,United States,Y
296,ATMA,\N,,AMA,ADIK,Kazakhstan,N
297,Aerolineas Medellin,\N,,AMD,AEROLINEAS MEDELLIN,Colombia,N
298,Ameriflight,\N,,AMF,AMFLIGHT,United States,N
299,Air Libert,\N,VD,,,France,Y
300,Air Lithuania,\N,TT,KLA,KAUNAS,Lithuania,N
301,Air Minas Linhas A,\N,,AMG,AIR MINAS,Brazil,N
302,Alan Mann Helicopters Ltd.,\N,,AMH,MANN,United Kingdom,N
303,Air Maldives,\N,,AMI,AIR MALDIVES,Maldives,N
304,Aviation Amos,\N,,AMJ,AVIATION AMOS,Canada,N
305,Amerer Air,\N,,AMK,AMER AIR,Austria,N
306,Air Malawi,\N,QM,AML,MALAWI,Malawi,Y
307,Aeroputul International Marculesti,\N,,AMM,AEROM,Moldova,N
308,Air Montenegro,\N,,AMN,MONTENEGRO,Montenegro,N
309,Air Montreal (Air Holdings Inc.),\N,,AMO,AIR MONTREAL,Canada,N
310,Aero Transporte S.A. (ATSA),\N,,AMP,ATSA,Peru,N
311,Aeromedicare Ltd.,\N,,AMQ,LIFELINE,United Kingdom,N
312,Air Sicilia,\N,BM,,,Italy,Y
313,Air Specialties Corporation,\N,,AMR,AIR AM,United States,N
314,Air Muskoka,\N,,AMS,AIR MUSKOKA,Canada,N
315,ATA Airlines,\N,,AMT,AMTRAN,United States,Y
316,Air Macau,\N,NX,AMU,AIR MACAO,Macao,Y
317,AMC Airlines,\N,,AMV,,Egypt,Y
318,Air Midwest,\N,ZV,AMW,AIR MIDWEST,United States,N
319,Air Seychelles,\N,HM,SEY,SEYCHELLES,Seychelles,Y
320,Air Sofia,\N,,SFB,AIR SOFIA,Bulgaria,N
321,AeroMéxico,\N,AM,AMX,AEROMEXICO,Mexico,Y
322,Air Ambar,\N,,AMY,AIR AMBAR,Dominican Republic,N
323,Amiya Airline,\N,,AMZ,AMIYA AIR,Nigeria,N
324,All Nippon Airways,ANA All Nippon Airways,NH,ANA,ALL NIPPON,Japan,Y
325,Air Navigation And Trading Co. Ltd.,\N,,ANB,AIR NAV,United Kingdom,N
326,Anglo Cargo,\N,,ANC,ANGLO,United Kingdom,N
327,Air Nostrum,\N,YW,ANE,AIR NOSTRUM,Spain,Y
328,Air Niugini,\N,PX,ANG,NUIGINI,Papua New Guinea,Y
329,Air Arabia,\N,G9,ABY,ARABIA,United Arab Emirates,Y
330,Air Canada,\N,AC,ACA,AIR CANADA,Canada,Y
331,Airport Bratsk,\N,,BRP,AEROBRA,Russia,N
332,Air Wings,\N,,BSB,ARBAS,Moldova,N
333,Air Baltic,\N,BT,BTI,AIRBALTIC,Latvia,Y
334,Alajnihah for Air Transport,\N,,ANH,ALAJNIHAH,Libya,N
335,Air Atlantic (Nig) Limited,\N,,ANI,NIGALANTIC,Nigeria,N
336,Air Nippon,\N,EL,ANK,ANK AIR,Japan,Y
337,Antares Airtransport,\N,,ANM, Maintenance & Service GmbH,ANTARES,N
338,Airnorth,\N,TL,ANO,TOPEND,Australia,Y
339,Aerol,\N,,ANQ,ANTIOQUIA,Colombia,N
340,Andes Lineas Aereas,\N,,ANS,AEROANDES,Argentina,N
341,Air North Charter - Canada,\N,4N,ANT,AIR NORTH,Canada,Y
342,Air Nevada,\N,,ANV,AIR NEVADA,United States,N
343,AOM French Airlines,\N,IW,AOM,French Lines,France,N
344,Aviacion Del Noroeste,\N,,ANW, S.A. de C.V.,AVINOR,N
345,Air New Zealand,\N,NZ,ANZ,NEW ZEALAND,New Zealand,Y
346,Avia Jaynar,\N,,SAP,TOBOL,Kazakhstan,N
347,Aero Servicios Empresariales,\N,,EMS,SERVIEMPRESARIAL,Mexico,N
348,Alcon Servicios Aereos,\N,,AOA, S.A. de C.V.,ALCON,N
349,AVCOM,\N,J6,AOC,AERO AVCOM,Russia,N
350,Aero Vodochody,\N,,AOD,AERO CZECH,Czech Republic,N
351,Air One Executive,\N,,AOE, ,Italy,N
352,Atair Pty Ltd.,\N,,AOF,ATAIR,South Africa,N
353,Aero VIP,\N,2D,AOG,AVIP,Argentina,N
354,Aeromundo Ejecutivo,\N,,MUN,AEROMUNDO,Mexico,N
355,Aerolinea Muri,\N,,MUR,MURI,Mexico,N
356,Astoria,\N,,AOI, Inc.,ASTORIA,N
357,Aero Rent JSC,\N,,NRO,AEROMASTER,Russia,N
358,Aeronord-Grup,\N,,NRP,AERONORD,Moldova,N
359,Aeroatlantico Colombia,\N,,AOK,,Colombia,N
360,Angkor Airlines,\N,,AOL,ANGKOR AIR,Cambodia,N
361,Aero Entreprise,\N,,AON,AERO ENTERPRISE,France,N
362,As,\N,,AOO, Opened Joint Stock Company,COMPANY AS,N
363,Aeropiloto,\N,,AOP,AEROPILOTO,Portugal,N
364,Aeroenlaces Nacionales,\N,,VIV,AEROENLACES,Mexico,N
365,Aerovis Airlines,\N,,VIZ,AEROVIZ,Ukraine,N
366,Avjet International (FZE),\N,,VJE,,United Arab Emirates,N
367,Aerovista Gulf Express,\N,,VGF,VISTA GULF,United Arab Emirates,N
368,Air Vegas,\N,6V,VGA,AIR VEGAS,United States,N
369,Almaver,\N,,VER,ALMAVER,Mexico,N
370,Afro International Ent. Limited,\N,,AOR,INTER-AFRO,Nigeria,N
371,Alitalia Express,\N,XM,SMX,ALIEXPRESS,Italy,Y
372,Asia Overnight Express,\N,OE,AOT,ASIA OVERNIGHT,Philippines,N
373,Air Tractor,\N,,AOU,AIR TRACTOR,Croatia,N
374,Aero Vision,\N,,AOV,AEROVISION,France,N
375,Aerotaxi Del Valle,\N,,AOX,AEROVALLE,Colombia,N
376,Air Park Aviation Ltd.,\N,,APA,CAN-AM,Canada,N
377,Airpac Airlines,\N,,APC, Inc.,AIRPAC,N
378,Aeroservicios Monterrey,\N,,SVM,SERVIMONTE,Mexico,N
379,Amapola Flyg AB,\N,,APF,AMAPOLA,Sweden,N
380,Air People International,\N,,APG,AIR PEOPLE,Thailand,N
381,Alpha Aviation,\N,,APH, Inc.,AIRFLIGHT,N
382,ASA Pesada,\N,,API, Lda.,ASA PESADA,N
383,Air Print,\N,,APJ, S.A.,AIR PRINT,N
384,Atlantic Airlines,\N,,BJK,BLACKJACK,United States,N
385,Aerotransporte Petrolero,\N,,PET,AEROPETRO,Colombia,N
386,Aero Flight,\N,GV,ARF,Aero Fox,Germany,Y
387,ACM Aviation,\N,,BJT,BAY JET,United States,N
388,Aircompany Barcol,\N,,BKL,BARCOL,Russia,N
389,All Charter Limited,\N,,BLA,ALL CHARTER,United Kingdom,N
390,Appalachian Flying Service,\N,,APL, Inc.,APPALACHIAN,N
391,Airpac,\N,,APM, Inc.,ALASKA PACIFIC,N
392,Aeropro,\N,,APO,AEROPRO,Canada,N
393,Aerolineas Pacifico Atlantico,\N,,APP, S.A. (Apair),AEROPERLAS,N
394,Aspen Aviation,\N,,APQ,ASPEN BASE,United States,N
395,Aeropuma,\N,,APU, S.A.,AEROPUMA,N
396,Air Plan International,\N,,APV,AIR PLAN,Democratic Republic of the Congo,N
397,Arrow Air,\N,JW,APW,BIG A,United States,Y
398,Apex Air Cargo,\N,,APX,PARCEL EXPRESS,United States,N
399,APA Internacional,\N,,APY,APA INTERNACIONAL,Dominican Republic,N
400,Aeroatlas,\N,,AQA, S.A.,ATCO,N
401,Aquila Air Ltd.,\N,,AQL,AQUILA,Canada,N
402,Air Queensland,\N,,AQN,BUSHAIR,Australia,N
403,Aluminum Company Of America,\N,,AQO,ALCOA SHUTTLE,United States,N
404,Aviones de Renta de Quintana Roo,\N,,AQT, S.A. de C.V.,AVIOQUINTANA,N
405,AirQuarius Aviation,\N,,AQU,QUARIUS,South Africa,N
406,Aerodyne Charter Company,\N,,AQZ,QUANZA,United States,N
407,Arik Air,\N,W3,ARA,ARIK AIR,Nigeria,N
408,Avia Air N.V.,\N,,ARB,AVIAIR,Aruba,N
409,Air Routing International Corp.,\N,,ARC, ,United States,N
410,Aerocondor,\N,2B,ARD,AEROCONDOR,Portugal,Y
411,Aires,\N,4C,ARE, Aerovias de Integracion Regional, S.A.,Y
412,Aerolineas Argentinas,\N,AR,ARG,ARGENTINA,Argentina,Y
413,Arrowhead Airways,\N,,ARH,ARROWHEAD,United States,N
414,Aero Vics,\N,,ARI,AEROVICS,Mexico,N
415,Aerojet de Costa Rica,\N,,ARJ, S.A.,,N
416,Antillana De Navegacion Aerea,\N,,SUN,,Dominican Republic,N
417,Aeroservicios De San Luis,\N,,SUO,SERVICIO SANLUIS,Mexico,N
418,Aerosuper,\N,,SUP,AEROSUPER,Mexico,N
419,Aero Link Air Services S.L.,\N,,ARK,LINK SERVICE,Spain,N
420,Airlec - Air Aquitaine Transport,\N,,ARL,AIRLEC,France,N
421,Aeromarket Express,\N,,ARM,AMEX,Spain,N
422,Arrow Aviation Ltd.,\N,,ARO,ARROW,Canada,N
423,RPA-Aviataxi Ltd.,\N,,ARP,,Russia,N
424,Air Klaipeda,\N,,KLD,AIR KLAIPEDA,Lithuania,N
425,Armstrong Air,\N,,ARQ, Inc.,ARMSTRONG,N
426,Air Armenia,\N,QN,ARR,AIR ARMENIA,Armenia,N
427,Aeromet Servicios,\N,,ARS,METSERVICE,Chile,N
428,Aerotal Aerolineas Territoriales de Colombia Ltda,\N,,ART,AEROTAL,Colombia,N
429,Aravco Ltd.,\N,,ARV,ARAVCO,United Kingdom,N
430,Aria,\N,,ARW,ARIABIRD,France,N
431,Air Xpress,\N,,ARX, Inc.,AIREX,N
432,Airline Alania,\N,,OST,ALANIA,Russia,N
433,Air Tchad,\N,,HTT,HOTEL TANGO,Chad,N
434,Aerloineas de Techuacan,\N,,HUC,LINEAS TEHUACAN,Mexico,N
435,Aerotransportes Huitzilin,\N,,HUT,AEROHUITZILIN,Mexico,N
436,Aero Transportes Del Humaya,\N,,HUY,AERO HUMAYA,Mexico,N
437,Argosy Airways,\N,,ARY,GOSEY,United States,N
438,Air Resorts,\N,,ARZ,AIR RESORTS,United States,N
439,Alaska Airlines,\N,AS,ASA, Inc.,ALASKA,Y
440,Air-Spray 1967 Ltd.,\N,,ASB,AIR SPRAY,Canada,N
441,Air Star Corporation,\N,,ASC,AIR STAR,Canada,N
442,Air Sinai,\N,4D,ASD,AIR SINAI,Egypt,Y
443,Airstars,\N,PL,ASE,MOROZOV,Russia,N
444,Air Schefferville,\N,,ASF, Inc.,SCHEFF,N
445,African Star Airways (PTY) Ltd.,\N,,ASG,AFRICAN STAR,South Africa,N
446,Aerosun International,\N,,ASI, Inc.,AEROSUN,N
447,Air Satellite,\N,,ASJ,SATELLITE,Canada,N
448,Awesome Flight Services (PTY) Ltd.,\N,,ASM,AWESOME,South Africa,N
449,Air and Sea Transport,\N,,ASN, ,Russia,N
450,Aero Slovakia,\N,,ASO,AERO NITRA,Slovakia,N
451,Airsprint,\N,,ASP,AIRSPRINT,Canada,N
452,Atlantic Southeast Airlines,\N,EV,ASQ,ACEY,United States,Y
453,All Star Airlines,\N,,ASR, Inc.,ALL STAR,N
454,Air Class,\N,,ASS, S.A. de C.V.,AIR CLASS,N
455,Aerlineas Del Oeste,\N,,AST,AEROESTE,Mexico,N
456,Aviones Are,\N,,NRE,AVIONES ARE,Mexico,N
457,Air Nepal International,\N,,NPL,AIR NEPAL,Nepal,N
458,Air Napier,\N,,NPR,,New Zealand,N
459,Air Taxi & Cargo,\N,,WAM,TAXI CARGO,Sudan,N
460,Arrow Panama,\N,,WAP,ARROW PANAMA,Panama,N
461,Astravia-Bissau Air Transports Ltd.,\N,,ASV,ASTRAVIA,Guinea-Bissau,N
462,Astrakhan Airlines,\N,OB,ASZ,AIR ASTRAKHAN,Russia,Y
463,Air Transport Association,\N,,ATA,,United States,N
464,Atlantair Ltd.,\N,,ATB,STARLITE,Canada,N
465,Air Tanzania,\N,TC,ATC,TANZANIA,Tanzania,Y
466,Aerotours Dominicana,\N,,ATD,AEROTOURS,Dominican Republic,N
467,Atlantis Transportation Services,\N,,ATE, Ltd.,ATLANTIS CANADA,N
468,Aerotrans Airline,\N,,ATG,BACHYT,Kazakhstan,N
469,Air Charter World,\N,,XAC,,United States,N
470,Air Burkina,\N,2J,VBW,BURKINA,Burkina Faso,Y
471,Air Travel Corp.,\N,,ATH,AIR TRAVEL,United States,N
472,Aero-Tropics Air Services,\N,HC,ATI,,Australia,N
473,Air Traffic GmbH,\N,,ATJ,SNOOPY,Germany,N
474,AeroTACA,\N,,ATK,AEROTACA,Colombia,N
475,Aerocentral,\N,,ATL,CENTRALMEX,Mexico,N
476,Airlines Of Tasmania,\N,FO,ATM,AIRTAS,Australia,Y
477,Air Saint Pierre,\N,PJ,SPM,,France,Y
478,Air Transport International,\N,8C,ATN,AIR TRANSPORT,United States,N
479,ASTRAL Colombia - Aerotransportes Especiales Ltda.,\N,,ATP,ASTRAL,Colombia,N
480,Air Transport Schiphol,\N,,ATQ,MULTI,Netherlands,N
481,Aeroferinco,\N,,FEO,FERINCO,Mexico,N
482,Aero Taxis Y Servicios Alfe,\N,,FES,AERO ALFE,Mexico,N
483,Avialesookhrana,\N,,FFA,AVIALESOOKHRANA,Russia,N
484,Africair Service,\N,,FFB,FOXTROT FOXTROT,Senegal,N
485,Atlas Airlines,\N,,ATR,ATLAS-AIR,United States,N
486,Air Transport Service,\N,,ATS,,Democratic Republic of the Congo,N
487,Aer Turas,\N,,ATT,AERTURAS,Ireland,N
488,Atlant Aerobatics Ltd.,\N,,ATU,ATLANT HUNGARY,Hungary,N
489,Avanti Air,\N,,ATV,AVANTI AIR,Germany,N
490,Aero Trades (Western) Ltd.,\N,,ATW,AERO TRADES,Canada,N
491,Austrian Airlines,\N,OS,AUA,AUSTRIAN,Austria,Y
492,Air Southwest,\N,,WOW,SWALLOW,United Kingdom,Y
493,Augsburg Airways,\N,IQ,AUB,AUGSBURG-AIR,Germany,Y
494,Air Corporate,\N,,CPV,AIRCORPORATE,Italy,N
495,Aviastar-Tu,\N,,TUP,TUPOLEVAIR,Russia,N
496,AirBridge Cargo,\N,RU,ABW,AIRBRIDGE CARGO,Russia,N
497,ATUR,\N,,TUR,,Ecuador,Y
498,ATESA Aerotaxis Ecuatorianos,\N,,TXU,ATESA,Ecuador,N
499,Austin Express,\N,,TXX,COWBOY,United States,N
500,Audi Air,\N,,AUD, Inc.,AUDI AIR,N
501,Augusta Air Luftfahrtunternehmen,\N,,AUF,AUGUSTA,Germany,N
502,Abu Dhabi Amiri Flight,\N,MO,AUH,SULTAN,United Arab Emirates,Y
503,Aeroflot-Nord,\N,5N,AUL,DVINA,Russia,Y
504,Air Atlantic Uruguay,\N,,AUM,ATLAMUR,Uruguay,N
505,Aviones Unidos,\N,,AUN,AVIONES UNIDOS,Mexico,N
506,Avia Business Group,\N,,AUP,,Russia,N
507,Aero Servicios Expecializados,\N,,SVE,AEROESPECIAL,Mexico,N
508,Aurigny Air Services,\N,GR,AUR,AYLINE,United Kingdom,Y
509,Aus-Air,\N,NO,AUS,,Australia,N
510,Austral Lineas Aereas,\N,AU,AUT,AUSTRAL,Argentina,Y
511,Aurora Aviation,\N,,AUU, Inc.,AURORA AIR,N
512,Air Uganda International Ltd.,\N,,AUX,,Uganda,N
513,Aerolineas Uruguayas,\N,,AUY, S.A.,AUSA,N
514,Australian Airlines,\N,AO,AUZ,AUSTRALIAN,Australia,N
515,Avianca - Aerovias Nacionales de Colombia,\N,AV,AVA, S.A.,AVIANCA,Y
516,Aviation Beaufort,\N,,AVB,BEAUPAIR,United Kingdom,N
517,Alamo Aviacion,\N,,AVD, S.L.,ALAMO,N
518,Aviair Aviation Ltd.,\N,,AVF,CARIBOO,Canada,N
519,Air Falcon,\N,,AVG,DJIBOUTI FALCON,Djibouti,N
520,AV8 Helicopters,\N,,AVH,KENT HELI,United Kingdom,N
521,Avia Traffic Company,\N,,AVJ,ATOMIC,Kyrgyzstan,N
522,AV8 Helicopters,\N,,AVK,AVIATE-COPTER,South Africa,N
523,Aviacion Ejecutiva Mexicana,\N,,AVM, S.A.,AVEMEX,N
524,Air Vanuatu,\N,NF,AVN,AIR VAN,Vanuatu,Y
525,Aviation at Work,\N,,AVO,AVIATION WORK,South Africa,N
526,Avcorp Registrations,\N,,AVP,AVCORP,United Kingdom,N
527,Alfa Aerospace,\N,,LFP,ALFA-SPACE,Australia,N
528,Atlantic Airfreight Aviation,\N,,LFR,LANFREIGHT,Sao Tome and Principe,N
529,Aerolift Company,\N,,LFT,LIFTCO,Sierra Leone,N
530,Aviation Services,\N,,AVQ, Inc.,AQUILINE,N
531,Active Aero Charter,\N,,AVR, Inc.,ACTIVE AERO,N
532,Avialsa T-35,\N,,AVS,AVIALSA,Spain,N
533,Asia Avia Airlines,\N,,AVT,ASIAVIA,Indonesia,N
534,Avia Sud A,\N,,AVU,AVIASUD,France,N
535,Airvantage Incorporated,\N,,AVV,AIRVANTAGE,United States,N
536,Aviator Airways,\N,,AVW,AVIATOR,Greece,N
537,Aviapaslauga,\N,,AVX,PASLAUGA,Lithuania,N
538,ASL,\N,,XXX,,Belgium,N
539,Air Yugoslavia,\N,,YRG,YUGAIR,Serbia,N
540,Airlink Zambia,\N,K8,ZAK,,Zambia,N
541,Agence Nationale des Aerodromes et de la Meteorologie,\N,,ZZM,,Ivory Coast,N
542,Airbus Transport International,\N,,BGA,BELOUGA,France,N
543,Air Bangladesh,\N,B9,BGD,AIR BANGLA,Bangladesh,Y
544,Aviodetachment-28,\N,,BGF,BULGARIAN,Bulgaria,N
545,Aero BG,\N,,BGG,AERO BG,Mexico,N
546,Aerotaxis De La Bahia,\N,,BHC,BAHIA,Mexico,N
547,Air Mediterranee,\N,DR,BIE,MEDITERRANEE,France,Y
548,Aviaservice,\N,,BIV,AVIASERVICE,Georgia,N
549,Aerolineas De El Salvador,\N,,SZA,AESA,El Salvador,N
550,Swaziland Airlink,\N,,SZL,SWAZILINK,Swaziland,N
551,Air Moorea,\N,,TAH,AIR MOOREA,France,Y
552,Aerovaradero,\N,,AVY, S.A.,AEROVARADERO,N
553,Air Valencia,\N,,AVZ,AIR VALENCIA,Spain,N
554,Airways International,\N,,AWB, Inc.,AIRNAT,N
555,Airwork,\N,,AWK,AIRWORK,New Zealand,N
556,Australian Wetleasing,\N,,AWL,AUSSIEWORLD,Australia,N
557,Air Niamey,\N,,AWN,AIR NIAMEY,Niger,N
558,Awood Air Ltd.,\N,,AWO,AWOOD AIR,Canada,N
559,Arctic Wings And Rotors Ltd.,\N,,AWR,ARCTIC WINGS,Canada,N
560,Auo Airclub AIST-M,\N,,ISM,STORK,Russia,N
561,Arab Wings,\N,,AWS,ARAB WINGS,Jordan,N
562,Air West,\N,,AWT,AIR WEST,Canada,N
563,Aeroline GmbH,\N,7E,AWU,SYLT-AIR,Germany,Y
564,Airwave Transport,\N,,AWV, Inc.,AIRWAVE,N
565,Air Wales,\N,6G,AWW,RED DRAGON,United Kingdom,Y
566,Aeroway,\N,,AWY, S.L.,AEROWEE,N
567,Air Caraïbes,\N,TX,FWI,FRENCH WEST,France,Y
568,AirWest,\N,,AWZ, ,Sudan,N
569,Air India Express,\N,IX,AXB,EXPRESS INDIA,India,Y
570,Air Express,\N,,AXD,AIR SUDEX,Sudan,N
571,Asian Express Airlines,\N,HJ,AXF,FREIGHTEXPRESS,Australia,N
572,Aeromexhaga,\N,,AXH,AEROMEXHAGA,Mexico,N
573,Aeron International Airlines,\N,,AXI, Inc.,AIR FREIGHTER,N
574,African Express Airways,\N,,AXK,EXPRESS JET,Kenya,N
575,Air Exel,\N,XT,AXL,EXEL COMMUTER,Netherlands,Y
576,AirAsia,Air Asia,AK,AXM,ASIAN EXPRESS,Malaysia,Y
577,Alexandair,\N,,AXN,ALEXANDROS,Greece,N
578,Aeromax,\N,,AXP,AEROMAX SPAIN,Spain,N
579,Action Airlines (Action Air Charter),\N,,AXQ,ACTION AIR,United States,N
580,Alberni Airways,\N,,BNI,ALBERNI,Canada,N
581,Aerolineas Bonanza,\N,,BNZ,AERO BONANZA,Mexico,N
582,Aerobona,\N,,BOC,AEROBONA,Mexico,N
583,Aboitiz Air,\N,,BOI,ABAIR,Philippines,N
584,Axel Rent,\N,,AXR, S.A.,RENTAXEL,N
585,Altus Airlines,\N,,AXS,ALTUS,United States,N
586,Aviaxess,\N,,AXV,AXAVIA,France,N
587,Avioimpex A.D.p.o.,\N,,AXX,IMPEX,Macedonia,N
588,Axis Airways,\N,6V,AXY,AXIS,France,N
589,Aladia Airlines,\N,,AYD,AIRLINES ALADIA,Mexico,N
590,Aykavia Aircompany,\N,,AYK,,Armenia,N
591,Airman,\N,,AYM, S.L.,AIRMAN,N
592,Atlantic Airlines,\N,,AYN, S.A.,ATLANTIC NICARAGUA,N
593,Awsaj Aviation Services,\N,,AYS,,Libya,N
594,Ayeet Aviation & Tourism,\N,,AYT,AYEET,Israel,N
595,Atlant-Soyuz Airlines,\N,3G,AYZ,ATLANT-SOYUZ,Russia,Y
596,Alitalia,\N,AZ,AZA,ALITALIA,Italy,Y
597,Azamat,\N,,AZB,TUMARA,Kazakhstan,N
598,Arcus-Air Logistic,\N,ZE,AZE,ARCUS AIR,Germany,N
599,Air Zermatt AG,\N,,AZF,AIR ZERMATT,Switzerland,N
600,Azalhelikopter,\N,,AZK,AZALHELICOPTER,Azerbaijan,N
601,Africa One,\N,,AZL,SKY AFRICA,Zambia,N
602,Aerocozumel,\N,,AZM,AEROCOZUMEL,Mexico,N
603,Amaszonas,\N,Z8,AZN,,Bolivia,Y
604,Arizona Pacific Airways,\N,,AZP,ARIZONA PACIFIC,United States,N
605,Aviacon Zitotrans Air Company,\N,,AZS,ZITOTRANS,Russia,N
606,Azimut,\N,,AZT, S.A.,AZIMUT,N
607,Azov Avia Airlines,\N,,AZV,AZOV AVIA,Ukraine,N
608,Air Zimbabwe,\N,UM,AZW,AIR ZIMBABWE,Zimbabwe,Y
609,Aero Jomacha,\N,,MHC,AERO JOMACHA,Mexico,N
610,Air Memphis,\N,,MHS,AIR MEMPHIS,Egypt,N
611,Air Memphis,\N,,MHU,MEPHIS UGANDA,Uganda,N
612,Air Marrakech Service,\N,,MKH,AIR MARRAKECH,Morocco,N
613,Air Max Africa,\N,,AZX,AZIMA,Gabon,N
614,Arizona Airways,\N,,AZY, Inc.,ARIZAIR,N
615,Azza Transport,\N,,AZZ,AZZA TRANSPORT,Sudan,N
616,Air Continental Inc,\N,,NAR,NIGHT AIR,United States,N
617,Antanik-Air,\N,,NAU,ANTANIK,Ukraine,N
618,Air Newark,\N,,NER,NEWAIR,United States,N
619,Aircraft Support and Services,\N,,NFF,,Lebanon,N
620,Aerobanana,\N,,OBA,AEROBANANA,Mexico,N
621,Amako Airlines,\N,,OBK,AMAKO AIR,Nigeria,N
622,Aserca Airlines,\N,R7,OCA,AROSCA,Venezuela,Y
623,Afrique CArgo Service Senegal,\N,,NFS,,Senegal,N
624,Angoservice,\N,,NGC,ANGOSERVICE,Angola,N
625,Angel Airlines,\N,,NGE,ANGEL AIR,Thailand,N
626,Angel Flight America,\N,,NGF,ANGEL FLIGHT,United States,N
627,Air Atonabee,\N,,OUL,CITY EXPRESS,Canada,N
628,Aero Nova,\N,,OVA,AERONOVA,Spain,N
629,Amira Air,\N,,XPE,EXPERT,Austria,N
630,Aero Express Intercontinental,\N,,XSS,INTER EXPRESS,Mexico,N
631,Advance Aviation Services,\N,,XTJ,,United States,N
632,Air Libya Tibesti,\N,,TLR,AIR LIBYA,Libya,N
633,Aerovic,\N,,OVC,,Ecuador,N
634,Airventure,\N,,RVE,AIRVENTURE,Belgium,N
635,Aero Servicios,\N,,RVI,AERO SERVICIOS,Mexico,N
636,Airvallee,\N,,RVL,AIR VALLEE,Italy,N
637,Aeromover,\N,,OVE,AEROMOVER,Mexico,N
638,Aerovias Ejecutivas,\N,,OVI,VIAS EJECUTIVAS,Mexico,N
639,Aero Servicio Pity,\N,,PTD,PITY,Mexico,N
640,Aero Copter,\N,,PTE,,Mexico,N
641,Rossiya-Russian Airlines,Pulkovo Aviation Enterprise,FV,SDM,PULKOVO,Russia,Y
642,Air Pal,\N,,PLL,AIRPAL,Spain,N
643,Air Pullmantur,\N,,PLM,PULLMANTUR,Spain,N
644,Aviaexpress,\N,RX,AEH,Avex,Hungary,N
645,Aviones Para Servirle,\N,,PSG,SERVIAVIONES,Mexico,N
646,Avio Sluzba,\N,,SLU,AVIO SLUZBA,Serbia,N
647,Air Scorpio,\N,,SCU,SCORPIO UNIVERS,Bulgaria,N
648,Air Spirit,\N,,SIP,AIR SPIRIT,United States,N
649,Alatau Airlines,\N,,BMV,OLIGA,Kazakhstan,N
650,Aviateca,\N,,GUG,AVIATECA,Guatemala,N
651,Aroostook Aviation,\N,,PXX,PINE STATE,United States,N
652,Aeropycsa,\N,,PYC,AEROPYCSA,Mexico,N
653,Association of Private Pilots of Kazakhstan,\N,,PVK,BORIS,Kazakhstan,N
654,The Amiri Flight,\N,,BAH,BAHRAIN,Bahrain,N
655,Aero Services,\N,,BAS,AEROSERV,Barbados,N
656,Ababeel Aviation,\N,,BBE,BABEL AIR,Sudan,N
657,Air Bashkortostan,\N,,BBT,AGYDAL,Russia,N
658,Ambulance Air Africa,\N,,MCY,MERCY,South Africa,N
659,American Eagle Airlines,\N,MQ,EGF,EAGLE FLIGHT,United States,Y
660,Aeropuelche,\N,,PUE,PUELCHE,Chile,N
661,Aeroput,\N,,PUT,,Serbia,N
662,Azzurra Air,\N,,AZI,AZZURRA,Italy,N
663,Airshop,\N,FF,,,Netherlands,N
664,African Transport Trading and Investment Company,\N,ML,,,Sudan,N
665,AD Aviation,\N,,VUE,FLIGHTVUE,United Kingdom,Y
666,Aero Costa Taxi Aereo,\N,,XCT,AEROCOSTAXI,Mexico,N
667,Aerovitro,\N,,VRO,AEROVITRO,Mexico,N
668,Aerotaxi Villa Rica,\N,,VRI,VILLARICA,Mexico,N
669,Aerovega,\N,,VEG,AEROVEGA,Mexico,N
670,Aerovilla,\N,,VVG,AEROVILLA,Colombia,N
671,Aerolineas Villaverde,\N,,VLR,VILLAVERDE,Mexico,N
672,Aero Air,\N,,WIL,WILLIAMETTE,United States,N
673,Aero Ejecutivos,\N,,VEJ,VENEJECUTIV,Venezuela,N
674,Aero Industries Inc,\N,,WAB,WABASH,United States,N
675,Aero Servicios Vanguardia,\N,,VNG,VANGUARDIA,Mexico,N
676,Aero Taxi Los Valles,\N,,VAD,VALLES,Spain,N
677,Aero Vilamoura,\N,,VMR,AERO VILAMOURA,Portugal,N
678,Aero Virel,\N,,VLS,VIREL,Mexico,N
679,Aeronautical Radio Inc,\N,,XAA,,United States,N
680,Aerovuelox,\N,,VUO,AEROVUELOX,Mexico,N
681,Aeronaves TSM,\N,,VTM,AERONAVES TSM,Mexico,N
682,Air Ivoire,\N,VU,VUN,AIRIVOIRE,Ivory Coast,Y
683,Air Botswana,\N,BP,BOT,BOTSWANA,Botswana,Y
684,Air-Rep,\N,,XPR,,United States,N
685,Air Excel,\N,,XLL,TINGA-TINGA,Tanzania,N
686,Air Evans,\N,,VAE,AIR-EVANS,Spain,N
687,Air Sorel,\N,,WHY,AIR SOREL,Canada,N
688,Air Net Private Charter,\N,,WDR,WIND RIDER,United States,N
689,Air Executive Charter,\N,,XEC,,Germany,N
690,Air Foyle,\N,GS,UPA,FOYLE,United Kingdom,Y
691,Air Midwest,\N,,VTY,VICTORY,Nigeria,N
692,Air Tahiti,\N,VT,VTA,AIR TAHITI,French Polynesia,Y
693,Air Urga,\N,3N,URG,URGA,Ukraine,N
694,Air Vardar,\N,,VDR,VARDAR,Macedonia,N
695,Air VIA,\N,VL,VIM,,Bulgaria,Y
696,Air Walser,\N,,WLR,AIRWALSER,Italy,N
697,Aircompany Rosavia,\N,,URA,ROSAVIA,Ukraine,N
698,Aircraft Performance Group,\N,,XLB,,United States,N
699,Airwaves Airlink,\N,,WLA,AIRLIMITED,Zambia,N
700,Airways Corporation of New Zealand,\N,,XFX,AIRCORP,New Zealand,N
701,Airways,\N,,WAY,GARONNE,France,N
702,Airwings oy,\N,,WGS,AIRWINGS,Finland,N
703,Airkenya,\N,,XAK,SUNEXPRESS,Kenya,N
704,Air-Lift Associates,\N,,WPK,WOLFPACK,United States,N
705,Airtrans Ltd,\N,,VAB,,Russia,N
706,ARP 410 Airlines,\N,,URP,AIR-ARP,Ukraine,N
707,Auckland Regional Rescue Helicopter Trust,\N,,WPR,WESTPAC RESCUE,New Zealand,N
708,Aurora Airlines,\N,,URR,AIR AURORA,Slovenia,N
709,Austro Aereo,\N,,UST,AUSTRO AEREO,Ecuador,N
710,Aviation Partners,\N,,WLT,WINGLET,United States,N
711,Avialift Vladivostok,\N,,VLV,VLADLIFT,Russia,N
712,Aviacion Comercial de America,\N,,VME,AVIAMERICA,Mexico,N
713,Aviast Air,\N,,VVA,IALSI,Russia,N
714,Aviation North,\N,,WLV,WOLVERINE,United States,N
715,Africa West,\N,FK,WTA,WEST TOGO,Togo,Y
716,Avient Air Zambia,\N,,VNT,AVIENT,Zambia,N
717,Aviazur,\N,,VZR,IAZUR,France,N
718,Aviaprad,\N,,VID,AVIAPRAD,Russia,N
719,Avirex,\N,G2,VXG,AVIREX-GABON,Gabon,N
720,Aviaexpress Aircompany,\N,,VXX,EXPRESSAVIA,Ukraine,N
721,AMR Services Corporation,\N,,XAM,ALLIANCE,United States,N
722,Airline Operations Services,\N,,XAO,,United States,N
723,Airlines 400,\N,,VAZ,REMONT AIR,Russia,N
724,ATRAN Cargo Airlines,\N,V8,VAS,ATRAN,Russian Federation,Y
725,Ameravia,\N,,VAM,AMERAVIA,Uruguay,N
726,Angkor Air,\N,,VAV,AIR ANGKOR,Cambodia,N
727,AVB-2004 Ltd,\N,,VBC,AIR VICTOR,Bulgaria,N
728,ASECNA,\N,,XKX,,France,N
729,AT and T Aviation Division,\N,,XAT,,United States,N
730,Avalair,\N,,VAI,AIR AVALAIR,Serbia,N
731,Averitt Air Charter,\N,,VRT,AVERITT,United States,N
732,Avensa,\N,VE,AVE,Avensa,Venezuela,N
733,Avolar Aerolineas,\N,V5,VLI,AEROVOLAR,Mexico,N
734,Avstar Aviation,\N,,VSA,STARBIRD,South Africa,N
735,AVESCA,\N,,VSC,AVESCA,Colombia,N
736,Aviaprom Enterprises,\N,,XAV,AVIAPROM,Russia,N
737,Avia Trans Air Transport,\N,,VTT,VIATRANSPORT,Sudan,N
738,Aviostart AS,\N,,VSR,AVIOSTART,Bulgaria,N
739,Aviacao Transportes Aereos e Cargas,\N,,VTG,ATACARGO,Angola,N
740,Assistance Aeroportuaire de L'Aeroport de Paris,\N,,XJA,,France,N
741,American Flight Service Systems,\N,,XFS,,United States,N
742,AASANA,\N,,XXV,,Bolivia,N
743,Australian air Express,\N,,XME,AUS-CARGO,Australia,N
744,AMS Group,\N,,XMG,,Russia,N
745,Air Caraibes Atlantique,\N,,CAJ,CAR LINE,France,N
746,Air China Cargo,\N,,CAO,AIRCHINA FREIGHT,China,N
747,Aerovias Caribe,\N,,CBE,AEROCARIBE,Mexico,N
748,Aerotaxi del Cabo,\N,,CBO,TAXI CABO,Mexico,N
749,Air Columbus,\N,,CBS,AIR COLUMBUS,Ukraine,N
750,Aero Cabo,\N,,CBV,CABOAEREO,Mexico,N
751,Air China,\N,CA,CCA,AIR CHINA,China,Y
752,Aerocardal,\N,,CDA,CARDAL,Chile,N
753,Aero Condor Peru,\N,Q6,CDP,CONDOR-PERU,Peru,Y
754,Aerotrans,\N,,CDU,,Russia,N
755,Airline Skol,\N,,CDV,SKOL,Russia,N
756,Aerofan,\N,,CFF,AEROFAN,Spain,N
757,ACEF,\N,,CFM,ACEF,Portugal,N
758,Africa One,\N,,CFR,,Democratic Republic of the Congo,N
759,Aero Calafia,\N,,CFV,CALAFIA,Mexico,N
760,Air Cargo Belize,\N,,CGB,CARGO BELIZE,Belize,N
761,Aero Clube Do Algarve,\N,,CGV,CLUBE ALGARVE,Portugal,N
762,Air Great Wall,\N,,CGW,CHANGCHENG,China,N
763,Aircompany Chaika,\N,,CHJ,AIR CHAIKA,Ukraine,N
764,Air Charter Services,\N,,CHR,ZAIRE CHARTER,Democratic Republic of the Congo,N
765,Air Charter Professionals,\N,,CHV,CHARTAIR,United States,N
766,Asia Continental Airlines,\N,,CID,ACID,Kazakhstan,N
767,Arctic Circle Air Service,\N,5F,CIR,AIR ARCTIC,United States,N
768,Aviation Charter Services,\N,,CKL,CIRCLE CITY,United States,N
769,Aerovias Castillo,\N,,CLL,AEROCASTILLO,Mexico,N
770,Aero Club De Portugal,\N,,CLP,CLUB PORTUGAL,Portugal,N
771,Air Care Alliance,\N,,CMF,COMPASSION,United States,N
772,Aero Continente Dominicana,\N,,CND,CONDOMINICANA,Dominican Republic,N
773,Air Toronto,\N,,CNE,CONNECTOR,Canada,N
774,Aquila Air,\N,,CNH,CHENANGO,United States,N
775,Air Consul,\N,,CNU,AIR CONSUL,Spain,N
776,Allcanada Express,\N,,CNX,CANEX,Canada,N
777,ATS Private Company,\N,,CPF,TECHSERVICE,Ukraine,N
778,Air Corridor,\N,QC,CRD,AIR CORRIDOR,Mozambique,N
779,Air Central,\N,NV,CRF,AIR CENTRAL,Japan,N
780,Air Cruzal,\N,,CRJ,AIR CRUZAL,Angola,N
781,Aerotransportes Corporativos,\N,,CRP,AEROTRANSCORP,Mexico,N
782,Air Creebec,\N,,CRQ,CREE,Canada,N
783,Aero Charter and Transport,\N,,CTA,CHAR-TRAN,United States,N
784,Air Tenglong,\N,,CTE,TENGLONG,China,N
785,Aerlineas Centauro,\N,,CTR,CENTAURO,Mexico,N
786,Aerocuahonte,\N,,CUO,CUAHONTE,Mexico,N
787,Air Chathams,\N,CV,CVA,CHATHAM,New Zealand,Y
788,Air Marshall Islands,\N,CW,CWM,AIR MARSHALLS,Marshall Islands,Y
789,Australian Customs Service,\N,,CWP,COASTWATCH,Australia,N
790,Air One Cityliner,\N,,CYL,CITYLINER,Italy,N
791,Air Transport,\N,,CYO,COYOTE,United States,N
792,Access Air,\N,ZA,CYD,CYCLONE,United States,Y
793,Aerocheyenne,\N,,CYE,AEROCHEYENNE,Mexico,N
794,Air Algerie,\N,AH,DAH,AIR ALGERIE,Algeria,Y
795,Aerovias DAP,\N,,DAP,DAP,Chile,N
796,Air Alpha,\N,,DBA,DOUBLE-A,United States,N
797,Air Niagara Express,\N,,DBD,AIR NIAGARA,Canada,N
798,Aviation Defense Service,\N,,DEF,TIRPA,France,N
799,Air Service Groningen,\N,,DEG,DEGGER,Netherlands,N
800,Adam Air,\N,KI,DHI,ADAM SKY,Indonesia,Y
801,Astar Air Cargo,\N,ER,DHL,DAHL,United States,N
802,Archer Aviation,\N,,DHM,ARCHER,United Kingdom,N
803,Aeromedica,\N,,DIC,AEROMEDICA,Mexico,N
804,Aerodin,\N,,DIN,AERODIN,Mexico,N
805,Antinea Airlines,\N,HO,DJA,ANTINEA,Algeria,N
806,Air Djibouti,\N,,DJU,AIR DJIB,Djibouti,N
807,Air Dolomiti,\N,EN,DLA,DOLOMOTI,Italy,Y
808,Aero Modelo,\N,,DLS,AEROMODELO,Mexico,N
809,Aerolineas Del Sur,\N,,DLU,DEL SUR,Chile,N
810,Aerodinamica de Monterrey,\N,,DMC,DINAMICAMONT,Mexico,N
811,Aeroservicios Dinamicos,\N,,DMI,AERODINAMICO,Mexico,N
812,Aerotaxis Dosmil,\N,,DML,,Mexico,N
813,Aerodespachos De El Salvador,\N,,DNA,AERODESPACHOS,El Salvador,N
814,Aeroynamics Malaga,\N,,DNC,FLYINGOLIVE,Spain,N
815,Aerodynamics Incorporated,\N,,DNJ,DYNAJET,United States,N
816,Aeroflot-Don,\N,D9,DNV,DONAVIA,Russia,Y
817,Air Madrid,\N,NM,DRD,ALADA AIR,Spain,Y
818,Airways Flight Training,\N,,DRM,DARTMOOR,United Kingdom,N
819,Aeronaves Del Noreste,\N,,DRO,AERONORESTE,Mexico,N
820,Addis Air Cargo Services,\N,,DSC,ADDIS CARGO,Ethiopia,N
821,Aero Algarve,\N,,DSK,SKYBANNER,Portugal,N
822,Aex Air,\N,,DST,DESERT,United States,N
823,Aero Davinci International,\N,,DVI,AERO DAVINCI,Mexico,N
824,Aero Dynamics,\N,,DYN,AERO DYNAMIC,United States,N
825,Aeroservicios Ecuatorianos,\N,,EAE,AECA,Ecuador,N
826,Aero-Pyrenees,\N,,EAP,AERO-PYRENEES,France,N
827,Air Transport,\N,,EAT,TRANS EUROPE,Slovakia,N
828,Aero Airlines,\N,EE,EAY,REVAL,Estonia,N
829,Aero Ejecutivo De Baja California,\N,,EBC,CALIXJET,Mexico,N
830,Air City,\N,4F,ECE,AIRCITY,Germany,N
831,Aero Ejecutivos RCG,\N,,ECG,EJECTUIVOS RCG,Mexico,N
832,Aeronautica Castellana,\N,,ECL,AERO CASTELLANA,Spain,N
833,Aerolineas Comerciales,\N,,ECM,AERO COMERCIALES,Mexico,N
834,Aerolineas Nacionales Del Ecuador,\N,,EDA,ANDES,Ecuador,N
835,Air Este,\N,,EET,AESTE,Spain,N
836,Air Mana,\N,,EFC,FLIGHT TAXI,France,N
837,Aer Lingus,\N,EI,EIN,SHAMROCK,Ireland,Y
838,Aeroservicios Ejecutivos Corporativos,\N,,EJP,EJECCORPORATIVOS,Mexico,N
839,Alpi Eagles,\N,E8,ELG,ALPI EAGLES,Italy,N
840,Arrendoora y Transportadora Aerea,\N,,END,ARRENDADORA,Mexico,N
841,Eliossola,\N,,EOS,ELIOSSOLA,Italy,N
842,Shenzhen Donghai Airlines,\N,,EPA,DONGHAI AIR,China,N
843,Aeronaves Del Noroeste,\N,,ENW,AERONOR,Spain,N
844,Airailes,\N,,EOL,EOLE,France,N
845,Aero Ermes,\N,,EOM,AERO ERMES,Mexico,N
846,Aero Transportes Empresariales,\N,,EPL,EMPRESARIALES,Mexico,N
847,Epps Air Service,\N,,EPS,EPPS AIR,United States,N
848,Aero Empresarial,\N,,EPE,AEROEMPRESARIAL,Mexico,N
849,Air S,\N,KY,EQL,EQUATORIAL,Air S,N
850,Avianergo,\N,,ERG,AVIANERGO,Russia,N
851,Aero Servicios Regiomontanos,\N,,ERI,ASERGIO,Mexico,N
852,Aerosec,\N,,ERK,AEROSEC,Chile,N
853,Aeromaan,\N,,ERM,EOMAAN,Mexico,N
854,Aerosaba,\N,,ESB,AEROSABA,Mexico,N
855,Avitat,\N,,ESO,,United Kingdom,N
856,Aerolineas Ejecutivas Del Sureste,\N,,ESU,ALESUR,Mexico,N
857,Aeronautica Le Esperanza,\N,,ESZ,ESPERANZA,Mexico,N
858,ATTICO,\N,,ETC,TRANATTICO,Sudan,N
859,Aero Siete,\N,,ETE,AEROSIETE,Mexico,N
860,Air Atlanta Europe,\N,,EUK,SNOWBIRD,United Kingdom,N
861,Air Evex,\N,,EVE,SUNBEAM,Germany,N
862,Aeronautical Academy of Europe,\N,,EVR,DIANA,Portugal,N
863,Air Exchange,\N,,EXG,EXCHANGE,United States,N
864,Atlantic Helicopters,\N,,FAC,FAROECOPTER,Denmark,N
865,Argentine Air Force,\N,,FAG,FUAER,Argentina,N
866,Air Fiji,\N,PC,FAJ,FIJIAIR,Fiji,N
867,AF-Air International,\N,,FAN,FANBIRD,Angola,N
868,Aviation Data Systems,\N,,FBW,,United States,N
869,Air Carriers,\N,,FCI,FAST CHECK,United States,N
870,Aerofrisco,\N,,FCO,AEROFRISCO,Mexico,N
871,Alfa 4,\N,,FCU,,Mexico,N
872,African Airlines,\N,,FDA,AIR SANKORE,Mali,N
873,African Medical and Research Foundation,\N,,FDS,FLYDOC,Kenya,N
874,Aero Freight,\N,,FGT,FREIAERO,Mexico,N
875,Aerosafin,\N,,FIC,AEROSAFIN,Mexico,N
876,Air Finland,\N,OF,FIF,AIR FINLAND,Finland,Y
877,Aerodata Flight Inspection,\N,,FII,FLIGHT CHECKER,Germany,N
878,Airfix Aviation,\N,,FIX,AIRFIX,Finland,Y
879,Air Pacific,\N,FJ,FJI,PACIFIC,Fiji,Y
880,Air Falcon,\N,,FLD,,Pakistan,N
881,Atlantic Airways,\N,RC,FLI,FAROELINE,Faroe Islands,Y
882,Air Florida,\N,QH,FLZ,AIR FLORIDA,United States,Y
883,Air Fret De Mauritanie,\N,,FMT,,Mauritania,N
884,Avio Nord,\N,,FNM,,Italy,N
885,Aeroflota Del Noroeste,\N,,FNO,RIAZOR,Spain,N
886,Aero Fenix,\N,,FNX,AERO FENIX,France,N
887,Ford Motor Company,\N,,FRD,FORD,United States,N
888,African Company Airlines,\N,,FPY,AFRICOMPANY,Democratic Republic of the Congo,N
889,Afrijet Airlines,\N,,FRJ,AFRIJET,Nigeria,N
890,Afrika Aviation Handlers,\N,,FRK,AFRIFAST,Kenya,N
891,Afrique Chart'air,\N,,FRQ,CHARTER AFRIQUE,Cameroon,N
892,Aerofreight Airlines,\N,,FRT,,Russia,N
893,Aereos Limited,\N,,FST,FAST TRACK,United Kingdom,N
894,Air Affaires Tchad,\N,,FTC,AFFAIRES TCHAD,Chad,N
895,ABC Bedarsflug,\N,,FTY,FLY TYROL,Austria,N
896,Air Iceland,\N,NY,FXI,FAXI,Iceland,Y
897,Air Philippines,\N,2P,GAP,ORIENT PACIFIC,Philippines,Y
898,Aerogaucho,\N,,GAU,AEROGAUCHO,Uruguay,N
899,Aero Business Charter,\N,,GBJ,GLOBAL JET,Germany,N
900,Atlantic Airlines,\N,,GBN,ATLANTIC GABON,Gabon,N
901,Aeronor,\N,,GCF,AEROCARTO,Spain,N
902,Aerogem Cargo,\N,,GCK,AEROGEM,Ghana,N
903,Aerovias del Golfo,\N,,GFO,AEROVIAS GOLFO,Mexico,N
904,Aeronautica,\N,,GGL,GIRA GLOBO,Angola,N
905,Air Georgian,\N,ZX,GGN,GEORGIAN,Canada,N
906,Aviance,\N,,GHL,HANDLING,United Kingdom,N
907,Air Ghana,\N,,GHN,AIR GHANA,Ghana,N
908,African International Transport,\N,,GIL,AFRICAN TRANSPORT,Guinea,N
909,Air Guinee Express,\N,2U,GIP,FUTURE EXPRESS,Guinea,Y
910,Africa Airlines,\N,,GIZ,AFRILENS,Guinea,N
911,Air Gemini,\N,,GLL,TWINS,Angola,N
912,Aero Charter,\N,,GLT,GASLIGHT,United States,N
913,Aguilas Mayas Internacional,\N,,GME,MAYAN EAGLES,Guatemala,N
914,Aerotaxis Guamuchil,\N,,GMM,AEROGUAMUCHIL,Mexico,N
915,Aeroservicios Gama,\N,,GMS,SERVICIOS GAMA,Mexico,N
916,Amber Air,\N,0A,GNT,GINTA,Lithuania,N
917,Alberta Government,\N,,GOA,ALBERTA,Canada,N
918,Air Scotland,\N,,GRE,GREECE AIRWAYS,Greece,N
919,Air Georgia,\N,DA,GRG,AIR GEORGIA,Georgia,N
920,Air Cargo Center,\N,,GRI,,Sao Tome and Principe,N
921,Air Greenland,\N,GL,GRL,GREENLAND,Denmark,Y
922,Allegro,\N,LL,GRO,ALLEGRO,Mexico,N
923,Agroar - Trabalhos Aereos,\N,,GRR,AGROAR,Portugal,N
924,Aircompany Grodno,\N,,GRX,GRODNO,Belarus,N
925,Airlift Alaska,\N,,GSP,GREEN SPEED,United States,N
926,Agrocentr-Avia,\N,,GSV,AGRAV,Kazakhstan,N
927,Altin Havayolu Tasimaciligi Turizm Ve Ticaret,\N,,GTC,GOLDEN WINGS,Turkey,N
928,Atlas Air,\N,5Y,GTI,GIANT,United States,Y
929,Aerotaxi Grupo Tampico,\N,,GTP,GRUPOTAMPICO,Mexico,N
930,Aerotaxis de Aguascalientes,\N,,GUA,AGUASCALIENTES,Mexico,N
931,Air Guyane,\N,GG,GUY,GREEN BIRD,French Guiana,Y
932,Air Victoria Georgia,\N,,GVI,IRINA,Georgia,N
933,Air D'Ayiti,\N,H9,HAD,HAITI AVIA,Haiti,N
934,Air Comores International,\N,GG,HAH,AIR COMORES,Comoros,N
935,Air Taxi,\N,,HAT,TAXI BIRD,France,N
936,Aerohein,\N,,HEI,AEROHEIN,Chile,N
937,Atlantic Air Lift,\N,,HGH,HIGHER,France,N
938,Atlantic Airlines de Honduras,\N,,HHA,ATLANTIC HONDURAS,Honduras,N
939,Aviacion Ejecutiva De Hildago,\N,,HID,EJECUTIVA HIDALGO,Mexico,N
940,Air Haiti,\N,,HJA,AIRHAITI,Haiti,N
941,Al Rais Cargo,\N,,HJT,AL-RAIS CARGO,United Arab Emirates,N
942,Air-Invest,\N,,HKH,HAWKHUNGARY,Hungary,N
943,Air Tahoma,\N,,HMA,TAHOMA,United States,N
944,Air Nova,\N,,HMT,HAMILTON,Canada,N
945,Aero Homex,\N,,HOM,AERO HOMEX,Mexico,N
946,Almiron Aviation,\N,,HPO,ALMIRON,Uganda,N
947,Avinor,\N,,HQO,,Norway,N
948,Airlink Airways,\N,,HYR,HIGHFLYER,Ireland,N
949,Air Horizon,\N,8C,HZT,HORIZON TOGO,Togo,N
950,AC Insat-Aero,\N,,IAE,,Russia,N
951,Air Inter Cameroun,\N,,ICM,INTER-CAMEROUN,Cameroon,N
952,Air Lift,\N,,IFI,HELLAS LIFT,Greece,N
953,Aero Survey,\N,,IKM,GHANA SURVEY,Ghana,N
954,Aero Airline,\N,,ILK,ILEK,Kazakhstan,N
955,Airtime Charters,\N,,IME,AIRTIME,United Kingdom,N
956,Aerotaxis Cimarron,\N,,IMN,TAXI CIMARRON,Mexico,N
957,Aero International,\N,,INA,AERO-NACIONAL,Mexico,N
958,Aeroingenieria,\N,,ING,AEROINGE,Chile,N
959,Aeroservicios Intergrados de Norte,\N,,INO,INTENOR,Mexico,N
960,Airpull Aviation,\N,,IPL,IPULL,Spain,N
961,Arvand Airlines,\N,,IRD,ARVAND,Iran,N
962,Atlas Aviation Group,\N,,IRH,ATLAS AVIA,Iran,N
963,Aram Airline,\N,,IRW,ARAM,Iran,N
964,Aria Tour,\N,,IRX,ARIA,Iran,N
965,Aerotaxi,\N,,ITE,AEROTAXI,Czech Republic,N
966,Avita-Servicos Aereos,\N,,ITF,AIR AVITA,Angola,N
967,Aero Citro,\N,,ITO,AERO CITRO,Mexico,N
968,Air Executive,\N,,IVE,COMPANY EXEC,Spain,N
969,Aviainvest,\N,,IWS,,Russia,N
970,Air Bagan,\N,W9,JAB,AIR BAGAN,Myanmar,Y
971,Aerojal,\N,,JAD,AEROJAL,Mexico,N
972,Airlink,\N,,JAR,AIRLINK,Austria,N
973,Ambjek Air Services,\N,,JEE,AMBJEK AIR,Nigeria,N
974,Avfinity,\N,,UTX,,United States,N
975,Alexandair,\N,,JMR,ALEXANDAIR,Canada,N
976,Air Jamaica Express,\N,,JMX,JAMAICA EXPRESS,Jamaica,N
977,Air Swift Aviation,\N,,JOA,,Australia,N
978,Aerojobeni,\N,,JOB,JOBENI,Mexico,N
979,Atyrau Air Ways,\N,IP,JOL,EDIL,Kazakhstan,N
980,Aerosmith Aviation,\N,,JPR,JASPER,United States,N
981,Arrendamiento de Aviones Jets,\N,,JTS,AVIONESJETS,Mexico,N
982,Aero Juarez,\N,,JUA,JUAREZ,Mexico,N
983,Air Canada Jazz,\N,QK,JZA,JAZZ,Canada,Y
984,Asia Aero Survey and Consulting Engineers,\N,,KAA,AASCO,Republic of Korea,N
985,Air Kirovograd,\N,,KAD,AIR KIROVOGRAD,Ukraine,N
986,Air Mach,\N,,KAM,ICO-AIR,Italy,N
987,Air Kufra,\N,,KAV,AIRKUFRA,Libya,N
988,Arkhabay,\N,,KEK,ARKHABAY,Kazakhstan,N
989,Aero Charter Krifka,\N,,KFK,KRIFKA AIR,Austria,N
990,Air Kraft Mir,\N,,KFT,AIR KRAFT MIR,Uzbekistan,N
991,Air Concorde,\N,,KGD,CONCORDE AIR,Bulgaria,N
992,Alexandria Airlines,\N,,KHH,,Egypt,N
993,Afit,\N,,KIE,TWEETY,Germany,N
994,Air South,\N,,KKB,KHAKI BLUE,United States,N
995,Atlasjet,\N,KK,KKK,ATLASJET,Turkey,Y
996,Air Mali International,\N,,KLB,TRANS MALI,Mali,N
997,Aerokaluz,\N,,KLZ,AEROKALUZ,Mexico,N
998,Air Koryo,\N,JS,KOR,AIR KORYO,Democratic People's Republic of Korea,Y
999,Araiavia,\N,,KOY,ALEKS,Kazakhstan,N
1000,AeroSucre,\N,,KRE,AEROSUCRE,Colombia,N
1001,Air Kokshetau,\N,,KRT,KOKTA,Kazakhstan,N
1002,Air Kissari,\N,,KSI,KISSARI,Angola,N
1003,Aeronavigaciya,\N,,KTN,AERONAVIGACIYA,Ukraine,N
1004,Alliance Avia,\N,,KVR,KAVAIR,Kazakhstan,N
1005,Av Atlantic,\N,,KYC,DOLPHIN,United States,N
1006,Air Astana,\N,KC,KZR,ASTANALINE,Kazakhstan,Y
1007,Aerovias De Lagos,\N,,LAG,AEROLAGOS,Mexico,N
1008,Albanian Airlines,\N,LV,LBC,ALBANIAN,Albania,Y
1009,Albisa,\N,,LBI,ALBISA,Mexico,N
1010,Albatros Airways,\N,,LBW,ALBANWAYS,Albania,N
1011,Aerolineas Aereas Ejecutivas De Durango,\N,,LDG,DURANGO,Mexico,N
1012,Aerologic,\N,,LDL,,Russia,N
1013,Al-Donas Airlines,\N,,LDN,ALDONAS AIR,Nigeria,N
1014,Aero Lider,\N,,LDR,AEROLIDER,Mexico,N
1015,Aleem,\N,,LEM,,Egypt,N
1016,Aerolineas Ejecutivas,\N,,LET,MEXEJECUTIV,Mexico,N
1017,Air Alfa,\N,,LFA,,Turkey,Y
1018,Aero Control Air,\N,,LFC,LIFE FLIGHT CANADA,Canada,N
1019,Aerolaguna,\N,,LGN,AEROLAGUNA,Mexico,N
1020,Al Ahram Aviation,\N,,LHR,AL AHRAM,Egypt,N
1021,Alidaunia,\N,D4,LID,ALIDA,Italy,N
1022,Al-Dawood Air,\N,,LIE,AL-DAWOOD AIR,Nigeria,N
1023,American Aviation,\N,,LKP,LAKE POWELL,United States,N
1024,Airlink Solutions,\N,,LKS,AIRLIN,Spain,N
1025,Air Solutions,\N,,LKY,LUCKY,United States,N
1026,Alliance Air,\N,CD,LLR,ALLIED,India,N
1027,Aerolima,\N,,LMA,AEROLIMA,Mexico,N
1028,Alamia Air,\N,,LML,ALAMIA AIR,Libya,N
1029,Air Plus Argentina,\N,,LMP,AIR FLIGHT,Argentina,N
1030,Almaty Aviation,\N,,LMT,ALMATY,Kazakhstan,N
1031,Aerolineas Mexicanas J S,\N,,LMX,LINEAS MEXICANAS,Mexico,N
1032,Air Almaty,\N,,LMY,AGLEB,Kazakhstan,N
1033,Air Almaty ZK,\N,,LMZ,ALUNK,Kazakhstan,N
1034,Aerolane,\N,XL,LNE,LAN ECUADOR,Ecuador,Y
1035,Airlink,\N,,LNK,LINK,South Africa,N
1036,Aerolineas Internacionales,\N,,LNT,LINEAINT,Mexico,N
1037,Alok Air,\N,,LOK,ALOK AIR,Sudan,N
1038,Air Saint Louis,\N,,LOU,AIR SAINTLOUIS,Senegal,N
1039,Alpine Aviation,\N,,LPC,NETSTAR,South Africa,N
1040,Air Alps Aviation,\N,A6,LPV,ALPAV,Austria,N
1041,Alrosa-Avia,\N,,LRO,ALROSA,Russia,N
1042,Al Rida Airways,\N,,LRW,AL RIDA,Mauritania,N
1043,Aurela,\N,,LSK,AURELA,Lithuania,N
1044,Aerobusinessservice,\N,,LSM,,Russia,N
1045,Alsair,\N,,LSR,ALSAIR,France,N
1046,Aerotaxis Latinoamericanos,\N,,LTI,LATINO,Mexico,N
1047,Alninati Aeronautics,\N,,LUC,STEF,Switzerland,N
1048,Atlantis European Airways,\N,TD,LUR,,Armenia,Y
1049,Aliven,\N,,LVN,ALIVEN,Italy,N
1050,Aviavilsa,\N,,LVR,AVIAVILSA,Lithuania,N
1051,Air Luxor GB,\N,L8,LXG,LUXOR GOLF,Guinea-Bissau,N
1052,Air Luxor,\N,LK,LXR,AIRLUXOR,Portugal,Y
1053,Apatas Air,\N,,LYT,APATAS,Lithuania,N
1054,Air Ban,\N,,LZP,DOC AIR,Bulgaria,N
1055,Air Lazur,\N,,LZR,LAZUR BEE-GEE,Bulgaria,N
1056,Aerodromo De La Mancha,\N,,MAM,AEROMAN,Spain,N
1057,Air Mauritius,\N,MK,MAU,AIRMAURITIUS,Mauritius,Y
1058,Avag Air,\N,,MBA,AVAG AIR,Austria,N
1059,Air Manas,\N,,MBB,AIR MANAS,Kyrgyzstan,N
1060,Airjet Exploracao Aerea De Carga,\N,,MBC,MABECO,Angola,N
1061,Aeriantur-M,\N,,MBV,AEREM,Moldova,N
1062,Air Mercia,\N,,MCB,WESTMID,United Kingdom,N
1063,Air Medical,\N,,MCD,AIR MED,United Kingdom,N
1064,Aerolineas Marcos,\N,,MCO,MARCOS,Mexico,N
1065,Atlantic Aero and Mid-Atlantic Freight,\N,,MDC,NIGHT SHIP,United States,N
1066,Air Madagascar,\N,MD,MDG,AIR MADAGASCAR,Madagascar,Y
1067,Aerosud Charter,\N,,MDX,MEDAIR,South Africa,N
1068,Air Meridan,\N,,MEF,EMPENNAGE,Nigeria,N
1069,Aero McFly,\N,,MFL,MCFLY,Mexico,N
1070,Asia Pacific Airlines,\N,,MGE,MAGELLAN,United States,N
1071,Aeromagar,\N,,MGS,AEROMAGAR,Mexico,N
1072,Aero Premier De Mexico,\N,,MIE,AEROPREMIER,Mexico,N
1073,Air Moldova,\N,9U,MLD,AIR MOLDOVA,Moldova,Y
1074,Amal Airlines,\N,,MLF,AMAL,Djibouti,N
1075,Air Mali,\N,,MLI,AIR MALI,Mali,N
1076,Air Madeleine,\N,,MLN,AIR MADELEINE,Canada,N
1077,Aermarche,\N,,MMC,AERMARCHE,Italy,N
1078,Air Alsie,\N,,MMD,MERMAID,Denmark,N
1079,Aviation Company Meridian,\N,,MMM,AVIAMERIDIAN,Russia,N
1080,AMP Incorporated,\N,,MMP,AMP-INC,United States,N
1081,Airmax,\N,,MMX,PERUMAX,Peru,N
1082,Aerolineas Amanecer,\N,,MNE,AEROAMANECER,Mexico,N
1083,Aero Mongolia,\N,M0,MNG,AERO MONGOLIA,Mongolia,N
1084,Air Monarch Cargo,\N,,MOC,MONARCH CARGO,Mexico,N
1085,Aeropublicitaria De Angola,\N,,MOP,PUBLICITARIA,Angola,N
1086,Aerolineas De Morelia,\N,,MOR,AEROMORELIA,Mexico,N
1087,Air Plus Comet,\N,A7,MPD,RED COMET,Spain,Y
1088,Aeromexpress,\N,QO,MPX,AEROMEXPRESS,Mexico,N
1089,Air ITM,\N,,MQT,MUSKETEER,France,N
1090,Aeromorelos,\N,,MRL,AEROMORELOS,Mexico,N
1091,Aerocharter,\N,,MRM,MARITIME,Canada,N
1092,Abas,\N,,MRP,ABAS,Czech Republic,N
1093,Air Mauritanie,\N,MR,MRT,MIKE ROMEO,Mauritania,N
1094,Air Cairo,\N,,MSC,,Egypt,N
1095,Air Sport,\N,,MSK,AIR SPORT,Bulgaria,N
1096,Aeromas,\N,,MSM,AEROMAS EXPRESS,Uruguay,N
1097,Aerol,\N,,MSO,MESO AMERICANAS,Mexico,N
1098,Aero-Kamov,\N,,MSV,AERAFKAM,Russia,N
1099,Aerotaxis Metropolitanos,\N,,MTB,AEROMETROPOLIS,Mexico,N
1100,Aeromet Linea Aerea,\N,,MTE,AEROMET,Chile,N
1101,Air Metack,\N,,MTK,AIRMETACK,Angola,N
1102,Air Montegomery,\N,,MTY,MONTY,United Kingdom,N
1103,Aerotaxi Mexicano,\N,,MXO,MAXAERO,Mexico,N
1104,Aero Yaqui Mayo,\N,,MYS,AERO YAQUI,Mexico,N
1105,AVC Airlines,\N,,MZK,,Japan,N
1106,Aerovias Montes Azules,\N,,MZL,MONTES AZULES,Mexico,N
1107,Aeroland Airways,\N,3S,,,Greece,N
1109,Astair,\N,8D,,,Russian Federation,Y
1110,Albarka Air,\N,F4,NBK,AL-AIR,Nigeria,N
1111,ACA-Ancargo Air Sociedade de Transporte de Carga Lda,\N,,NCL,ANCARGO AIR,Angola,N
1112,Aero Servicios de Nuevo Laredo,\N,,NEL,AEROLAREDO,Mexico,N
1113,Angoavia,\N,,NGV,ANGOAVIA,Angola,N
1114,Aeroni,\N,,NID,AERONI,Mexico,N
1115,Aeroejecutiva Nieto,\N,,NIE,AERONIETO,Mexico,N
1116,Aero Contractors,\N,AJ,NIG,AEROLINE,Nigeria,Y
1117,Aerokuzbass,\N,,NKZ,NOVOKUZNETSK,Russia,N
1118,Atlantic Airlines,\N,,NPT,NEPTUNE,United Kingdom,N
1119,Atlantic Richfield Company,\N,,NRS,NORTH SLOPE,United States,N
1120,Aerolineas Sosa,\N,,NSO,SOSA,Honduras,N
1121,Aero Norte,\N,,NTD,,Mexico,N
1122,Air Inter Ivoire,\N,,NTV,INTER-IVOIRE,Ivory Coast,N
1123,Aeroservicios De Nuevo Leon,\N,,NUL,SERVICIOS NUEVOLEON,Mexico,N
1124,Avial NV Aviation Company,\N,,NVI,NEW AVIAL,Russia,N
1125,Airwing,\N,,NWG,NORWING,Norway,N
1126,Air Next,\N,,NXA,BLUE-DOLPHIN,Japan,N
1127,Arkhangelsk 2 Aviation Division,\N,,OAO,DVINA,Russia,N
1128,Aerogisa,\N,,OGI,AEROGISA,Mexico,N
1129,Aerolineas Olve,\N,,OLV,OLVE,Mexico,N
1130,Aeromega,\N,,OMG,OMEGA,United Kingdom,N
1131,Air One Nine,\N,,ONR,EDER,Libya,N
1132,Air Ontario,\N,,ONT,ONTARIO,Canada,N
1133,Aerocorp,\N,,ORP,CORPSA,Mexico,N
1134,Action Air,\N,,ORS,AVIATION SERVICE,Italy,N
1135,Aerosan,\N,,OSN,AEROSAN,Mexico,N
1136,Aviapartner Limited Company,\N,,OSO,,Russia,N
1137,Aliparma,\N,,PAJ,ALIPARMA,Italy,N
1138,Air Parabet,\N,,PBT,PARABET,Bangladesh,N
1139,Air Burundi,\N,8Y,PBU,AIR-BURUNDI,Burundi,N
1140,Aeropostal Cargo de Mexico,\N,,PCG,POSTAL CARGO,Mexico,N
1141,Air Pack Express,\N,,PCK,AIRPACK EXPRESS,Spain,N
1142,Air Palace,\N,,PCS,AIR PALACE,Mexico,N
1143,Aeropelican Air Services,\N,OT,PEL,PELICAN,Australia,Y
1144,Aerolineas Chihuahua,\N,,PFI,PACIFICO CHIHUAHUA,Mexico,N
1145,Air Cargo Express International,\N,,PFT,PROFREIGHT,United States,N
1146,Al Farana Airline,\N,,PHR,PHARAOH,Egypt,N
1147,Ave.com,\N,,PHW,PHOENIX SHARJAH,United Arab Emirates,N
1148,Air South West,\N,,PIE,PIRATE,United Kingdom,N
1149,Aeroservicios California Pacifico,\N,,PIF,AEROCALPA,Mexico,N
1150,AST Pakistan Airways,\N,,PKA,PAKISTAN AIRWAY,Pakistan,N
1151,Aero Personal,\N,,PNL,AEROPERSONAL,Mexico,N
1152,Aero Servicios Platinum,\N,,PNU,AERO PLATINUM,Mexico,N
1153,Apoyo Aereo,\N,,POY,APOYO AEREO,Mexico,N
1154,Aerotransportes Privados,\N,,PRI,AEROPRIV,Mexico,N
1155,Atlantic Coast Jet,\N,,PRT,PATRIOT,United States,N
1156,Air Paradise International,\N,AD,PRZ,RADISAIR,Indonesia,N
1157,Aeroservicio Sipse,\N,,PSE,SIPSE,Mexico,N
1158,Aeroservicios Corporativos De San Luis,\N,,PSL,CORSAN,Mexico,N
1159,Aerotransportes Privados,\N,,PVA,TRANSPRIVADO,Mexico,N
1160,Aereo Taxi Paraza,\N,,PZA,AEREO PARAZA,Mexico,N
1161,Air Class Lineas Aereas,\N,QD,QCL,ACLA,Uruguay,N
1162,Aviation Consultancy Office,\N,,QEA,,Australia,N
1163,Aero Taxi,\N,,QAT,AIR QUASAR,Canada,N
1164,Aero Taxi Aviation,\N,,QKC,QUAKER CITY,United States,N
1165,Aviation Quebec Labrador,\N,,QLA,QUEBEC LABRADOR,Canada,N
1166,African Safari Airways,\N,QS,QSC,ZEBRA,Kenya,N
1167,Aero Quimmco,\N,,QUI,QUIMMCO,Mexico,N
1168,Alada,\N,,RAD,AIR ALADA,Angola,N
1169,Aerotur-Air,\N,,RAI,DIASA,Kazakhstan,N
1170,Air Center Helicopters,\N,,RAP,RAPTOR,United States,N
1171,Aur Rum Benin,\N,,RBE,RUM BENIN,Benin,N
1172,Aeroserivios Del Bajio,\N,,RBJ,AEROBAJIO,Mexico,N
1173,Airbus France,\N,4Y,RBU,AIRBUS FRANCE,France,N
1174,Air Roberval,\N,,RBV,AIR ROBERVAL,Canada,N
1175,Arubaexel,\N,,RBX,ARUBA,Aruba,N
1176,Air Charters Eelde,\N,,RCC,RACER,Netherlands,N
1177,Aerocer,\N,,RCE,AEROCER,Mexico,N
1178,Aeroflot-Cargo,\N,,RCF,AEROFLOT-CARGO,Russia,N
1179,Air Mobility Command,\N,MC,RCH,REACH,United States,N
1180,Air Cassai,\N,,RCI,AIR CASSAI,Angola,N
1181,Aero Renta De Coahuila,\N,,RCO,AEROCOAHUILA,Mexico,N
1182,Aerocorp,\N,,RCP,AEROCORPSA,Mexico,N
1183,Aerolineas Regionales,\N,,RCQ,REGIONAL CARGO,Mexico,N
1184,Atlantic S.L.,\N,,RCU,AIR COURIER,Spain,N
1185,Air Service Center,\N,,RCX,SERVICE CENTER,Italy,N
1186,Adygeya Airlines,\N,,RDD,ADLINES,Russia,N
1187,Air Ada,\N,,RDM,AIR ADA,Mauritania,N
1188,Aer Arann,\N,RE,REA,AER ARANN,Ireland,Y
1189,Aero-Rent,\N,,REN,AERORENT,Mexico,N
1190,Australian Maritime Safety Authority,\N,,RES,RESCUE,Australia,N
1191,Air Austral,\N,UU,REU,REUNION,France,Y
1192,Aero-Rey,\N,,REY,AEROREY,Mexico,N
1193,Aero Africa,\N,,RFC,AERO AFRICA,Swaziland,N
1194,Aerotransportes Rafilher,\N,,RFD,RAFHILER,Mexico,N
1195,Argo,\N,,RGO,ARGOS,Dominican Republic,N
1196,Airbourne School of Flying,\N,,RGT,AIRGOAT,United Kingdom,N
1197,Air Archipels,\N,,RHL,ARCHIPELS,France,N
1198,Aviation Ministry of the Interior of the Russian Federation,\N,,RIF,INTERMIN AVIA,Russian Federation,N
1199,Aeris Gestion,\N,,RIS,AERIS,Spain,N
1200,Asian Spirit,\N,6K,RIT,ASIAN SPIRIT,Philippines,Y
1201,Aeroservicios Jet,\N,,RJS,ASERJET,Mexico,N
1202,Air Afrique,\N,RK,RKA,AIRAFRIC,Ivory Coast,Y
1203,Airlinair,\N,A5,RLA,AIRLINAIR,France,Y
1204,Air Nelson,\N,,RLK,,New Zealand,N
1205,Air Leone,\N,,RLL,AEROLEONE,Sierra Leone,N
1206,Aero Lanka,\N,QL,RLN,AERO LANKA,Sri Lanka,Y
1207,Air Alize,\N,,RLZ,ALIZE,France,N
1208,Air Amder,\N,,RMD,AIR AMDER,Mauritania,N
1209,Armenian Airlines,\N,,RME,ARMENIAN,Armenia,N
1210,Armenian International Airways,\N,MV,RML,,Armenia,N
1211,Arm-Aero,\N,,RMO,ARM-AERO,Armenia,N
1212,Air Max,\N,,RMX,AEROMAX,Bulgaria,N
1213,Air Salone,\N,20,RNE,AIR SALONE,Sierra Leone,Y
1214,Aeronem Air Cargo,\N,,RNM,AEROMNEM,Ecuador,N
1215,Air Cargo Masters,\N,,RNR,RUNNER,United States,N
1216,Armavia,\N,U8,RNV,ARMAVIA,Armenia,Y
1217,Aeroeste,\N,,ROE,ESTE-BOLIVIA,Bolivia,N
1218,Aero Gen,\N,,ROH,AEROGEN,Mexico,N
1219,Avior Airlines,\N,,ROI,AVIOR,Venezuela,N
1220,Aeroel Airways,\N,,ROL,AEROEL,Israel,N
1221,Aeromar,\N,BQ,ROM,BRAVO QUEBEC,Dominican Republic,N
1222,Aeroitalia,\N,,ROO,AEROITALIA,Italy,N
1223,Aerodan,\N,,ROD,AERODAN,Mexico,N
1224,AeroRep,\N,P5,RPB,AEROREPUBLICA,Colombia,Y
1225,Aerolineas Del Pacifico,\N,,RPC,AEROPACSA,Ecuador,N
1226,Aero Roca,\N,,RRC,AEROROCA,Mexico,N
1227,Aerotransportes Internacionales De Torreon,\N,,RRE,AERO TORREON,Mexico,N
1228,Acvila Air-Romanian Carrier,\N,,RRM,AIR ROMANIA,Romania,N
1229,Aerolineas Ejecutivas Tarascas,\N,,RSC,TARASCAS,Mexico,N
1230,Aero-Service,\N,BF,RSR,CONGOSERV,Republic of the Congo,Y
1231,Aerosur,\N,5L,RSU,AEROSUR,Bolivia,Y
1232,Aeronorte,\N,,RTE,LUZAVIA,Portugal,Y
1233,Artis,\N,,RTH,ARTHELICO,France,N
1234,Arhabaev Tourism Airlines,\N,,RTO,ARTOAIR,Kazakhstan,N
1235,Air Turquoise,\N,,RTQ,TURQUOISE,France,N
1236,Aerotucan,\N,,RTU,AEROTUCAN,Mexico,N
1237,Air Anastasia,\N,,RUD,ANASTASIA,Ukraine,N
1238,Air Rum,\N,,RUM,AIR RUM,Sierra Leone,N
1239,ACT Havayollari,\N,,RUN,CARGO TURK,Turkey,N
1240,Air VIP,\N,,RVP,AEROVIP,Portugal,N
1241,Aircompany Veteran,\N,,RVT,AIR-VET,Armenia,N
1242,Alliance Express Rwanda,\N,,RWB,,Rwanda,N
1243,Arrow Ecuador Arrowec,\N,,RWC,ARROWEC,Ecuador,N
1244,Aerogroup 98 Limited,\N,,RWY,TYNWALD,United Kingdom,N
1245,Aeroxtra,\N,,RXT,AERO-EXTRA,Mexico,N
1246,Express Air Charter,\N,,RXX,REX AIR,Canada,N
1247,Air Whitsunday,\N,,RWS,,Australia,N
1248,Aero Zambia,\N,,RZL,AERO ZAMBIA,Zambia,N
1249,Aero Zano,\N,,RZN,ZANO,Mexico,N
1250,Anoka Air Charter,\N,,RZZ,RED ZONE,United States,N
1251,Aerosaab,\N,,SBH,AEROSAAB,Mexico,N
1252,Associated Aviation,\N,,SCD,ASSOCIATED,Nigeria,N
1253,American Jet International,\N,,SCM,SCREAMER,United States,N
1254,Air Santo Domingo,\N,EX,SDO,AERO DOMINGO,Dominican Republic,N
1255,Aero Sudpacifico,\N,,SDP,SUDPACIFICO,Mexico,N
1256,Aero Servicios Ejecutivas Del Pacifico,\N,,SEF,SERVIPACIFICO,Mexico,N
1257,Aero California,\N,JR,SER,AEROCALIFORNIA,Mexico,N
1258,A-Safar Air Services,\N,,SFM,AIR SAFAR,Nigeria,N
1259,Aerosegovia,\N,,SGV,SEGOVIA,Nicaragua,N
1260,Airshare Holdings,\N,,SHH,AIRSHARE,United Kingdom,N
1261,Aerosiyusa,\N,,SIY,SIYUSA,Mexico,N
1262,Aero Silza,\N,,SIZ,AEROSILZA,Mexico,N
1263,Air San Juan,\N,,SJN,SAN JUAN,United States,N
1264,Aero-North Aviation Services,\N,,SKP,SKIPPER,Canada,N
1265,Aero Sami,\N,,SMI,SAMI,Mexico,N
1266,Avient Aviation,\N,Z3,SMJ,AVAVIA,Zimbabwe,Y
1267,Aerolineas Sol,\N,,SOD,ALSOL,Mexico,N
1268,Air Soleil,\N,,SOE,AIR SOLEIL,Mauritania,N
1269,Aero Soga,\N,,SOG,AEROSOGA,Guinea-Bissau,N
1270,Airspeed Aviation,\N,,SPD,SPEEDLINE,Canada,N
1271,Air Service,\N,M3,SPJ,AIR SKOPJE,Macedonia,N
1272,Aeroservicios Ejecutivos Del Pacifico,\N,,SPO,EJECTUIV PACIFICO,Mexico,N
1273,Asian Aerospace Service,\N,,SPY,THAI SPACE,Thailand,N
1274,Airworld,\N,,SPZ,SPEED SERVICE,South Africa,N
1275,Alsaqer Aviation,\N,,SQR,ALSAQER AVIATION,Libya,N
1276,Air Safaris and Services,\N,,SRI,AIRSAFARI,New Zealand,N
1277,Aero Servicio Corporativo,\N,,SRV,SERVICORP,Mexico,N
1278,Air Sultan,\N,,SSL,SIERRA SULTAN,Sierra Leone,N
1279,Aero 1 Pro-Jet,\N,,SSM,RAPID,Canada,N
1280,Airquarius Air Charter,\N,,SSN,SUNSTREAM,South Africa,N
1281,Aeropac,\N,,STK,SAT PAK,United States,N
1282,Air St. Thomas,\N,,STT,PARADISE,United States,N
1283,Aerolineas Del Sureste,\N,,SUE,AEROSURESTE,Mexico,N
1284,Aerial Surveys (1980) Limited,\N,,SUY,SURVEY,New Zealand,N
1285,Air Slovakia,\N,GM,SVK,SLOVAKIA,Slovakia,N
1286,Adler Aviation,\N,,SWH,SHOCKWAVE,Canada,N
1287,Aircompany Yakutia,\N,R3,SYL,AIR YAKUTIA,Russia,Y
1288,Aerosud Aviation,\N,,SYT,SKYTRACK,South Africa,N
1289,Aeroservicios de La Costa,\N,,TAA,AERO COSTA,Mexico,N
1290,Aeromar,\N,VW,TAO,TRANS-AEROMAR,Mexico,Y
1291,Aerotrebol,\N,,TBL,AEROTREBOL,Mexico,N
1292,Aero Taxi de Los Cabos,\N,,TBO,AERO CABOS,Mexico,N
1293,Air Turks and Caicos,\N,JY,TCI,KERRMONT,Turks and Caicos Islands,N
1294,Aerotranscolombina de Carga,\N,,TCO,TRANSCOLOMBIA,Colombia,N
1295,Air Cargo Express,\N,,TDG,TURBO DOG,United States,N
1296,Atlas Helicopters,\N,,TDT,TRIDENT,United Kingdom,N
1297,Air Today,\N,,TDY,AIR TODAY,United States,N
1298,Aero Servicios Azteca,\N,,TED,AEROAZTECA,Mexico,N
1299,Arkefly,\N,OR,TFL,ARKEFLY,Netherlands,Y
1300,Antair,\N,,TIR,ANTAIR,Mexico,N
1301,Atlantique Air Assistance,\N,,TLB,TRIPLE-A,France,N
1302,Aero Taxi Autlan,\N,,TLD,AEREO AUTLAN,Mexico,N
1303,Aero Util,\N,,TLE,AEROUTIL,Mexico,N
1304,Aero Toluca Internactional,\N,,TLU,AEROTOLUCA,Mexico,N
1305,Aero Taxi del Centro de Mexico,\N,,TME,TAXICENTRO,Mexico,N
1306,Aerotropical,\N,,TOC,TROPICMEX,Mexico,N
1307,Air Tomisko,\N,,TOH,TOMISKO CARGO,Serbia,N
1308,Airlines PNG,\N,CG,TOK,BALUS,Papua New Guinea,Y
1309,Aero Tanala,\N,,TON,AEROTANALA,Mexico,N
1310,Aero Tropical,\N,,TPB,AERO TROPICAL,Angola,N
1311,Air Cal,\N,TY,TPC,AIRCAL,France,N
1312,Air Horizon,\N,,TPK,TCHAD-HORIZON,Chad,N
1313,Aero Taxi del Potosi,\N,,TPO,TAXI-POTOSI,Mexico,N
1314,Aeroturquesa,\N,,TQS,AEROTURQUESA,Mexico,N
1315,Airmark Aviation,\N,,TRH,TRANSTAR,United States,N
1316,AirTran Airways,\N,FL,TRS,CITRUS,United States,Y
1317,Air Transat,\N,TS,TSC,TRANSAT,Canada,Y
1318,airtransse,\N,,TSQ,AIRTRA,Japan,N
1319,Aerolineas Turisticas del Caribe,\N,,TTB,AERO TURISTICAS,Mexico,N
1320,Avcenter,\N,,TTE,TETON,United States,N
1321,Air Tungaru,\N,,TUN,TUNGARU,Kiribati,N
1322,Avialeasing Aviation Company,\N,EC,TWN,TWINARROW,Uzbekistan,Y
1323,Aerotaxis del Noroeste,\N,,TXD,TAXI OESTE,Mexico,N
1324,Aerotaxis Alfe,\N,,TXF,ALFE,Mexico,N
1325,Aereotaxis,\N,,TXI,AEREOTAXIS,Mexico,N
1326,Tyrolean Airways,\N,VO,TYR,TYROLEAN,Austria,Y
1327,Aero Tomza,\N,,TZA,AERO TOMZA,Mexico,N
1328,Air Zambezi,\N,,TZT,ZAMBEZI,Zimbabwe,N
1329,Afra Airlines,\N,,UAG,AFRALINE,Ghana,N
1330,Aerostar Airlines,\N,,UAR,AEROSTAR,Ukraine,N
1331,Air Division of the Eastern Kazakhstan Region,\N,,UCK,GALETA,Kazakhstan,N
1332,Aero-Charter Ukraine,\N,DW,UCR,CHARTER UKRAINE,Ukraine,N
1333,Air LA,\N,,UED,AIR L-A,United States,N
1334,Air Ukraine,\N,6U,UKR,AIR UKRAINE,Ukraine,N
1335,Air Umbria,\N,,UMB,AIR UMBRIA,Italy,N
1336,Atuneros Unidos de California,\N,,UND,ATUNEROS UNIDOS,Mexico,N
1337,Aeroclub Flaps,\N,,FLP,AEROCLUB FLAPS,Spain,N
1338,Aerolineas Galapagos (Aerogal),\N,2K,GLG,AEROGAL,Ecuador,Y
1339,Air Castle Corporation / Global Airways,\N,,GLB,GLO-AIR,United States,N
1340,Alrosa Mirny Air Enterprise,\N,6R,DRU,MIRNY,Russia,Y
1341,Bearing Supplies Limited,\N,,PVO,PROVOST,United Kingdom,N
1342,Balkan Agro Aviation,\N,,BAA,BALKAN AGRO,Bulgaria,N
1343,Bahrain Air BSC (Closed),\N,,BAB,AWAL,Bahrain,N
1344,BAC Leasing Limited,\N,,BAC, ,United Kingdom,N
1345,BAE Systems,\N,,BAE,FELIX,United Kingdom,N
1346,Belgian Air Force,\N,,BAF,BELGIAN AIRFORCE,Belgium,N
1347,Baker Aviation,\N,8Q,BAJ,BAKER AVIATION,United States,N
1348,Blackhawk Airways,\N,,BAK,BLACKHAWK,United States,N
1349,Britannia Airways,\N,,BAL,BRITANNIA,United Kingdom,N
1350,Business Air Services,\N,,BAM,BUSINESS AIR,Canada,N
1351,British Antarctic Survey,\N,,BAN,PENGUIN,United Kingdom,N
1352,Bradly Air (Charter) Services,\N,,BAR,BRADLEY,Canada,N
1353,Bissau Airlines,\N,,BAU,AIR BISSAU,Guinea-Bissau,N
1354,Bay Aviation Ltd,\N,,BAV,BAY AIR,Bangladesh,N
1355,British Airways,\N,BA,BAW,SPEEDBIRD,United Kingdom,Y
1356,Best Aero Handling Ltd,\N,,BAX,,Russia,N
1357,Bayon Airlines,\N,,BAY,BAYON,Cambodia,N
1358,Bannert Air,\N,,BBA,BANAIR,Austria,N
1359,Biman Bangladesh Airlines,\N,BG,BBC,BANGLADESH,Bangladesh,Y
1360,Bluebird Cargo,\N,BF,BBD,BLUE CARGO,Iceland,N
1361,Beibars CJSC,\N,,BBS,BEIBARS,Kazakhstan,N
1362,Bravo Airlines,\N,,BBV,BRAVO EUROPE,Spain,N
1363,Blue Bird Aviation,\N,,BBZ,COBRA,Kenya,N
1364,BACH Flugbetriebsges,\N,B4,BCF,BACH,Austria,N
1365,Blue Islands,\N,,BCI,BLUE ISLAND,United Kingdom,N
1366,British Caribbean Airways,\N,,BCL,,United Kingdom,N
1367,British Charter,\N,,BCR,BACKER,United Kingdom,N
1368,BCT Aviation,\N,,BCT,BOBCAT,United Kingdom,N
1369,Business Aviation Center,\N,,BCV,BUSINESS AVIATION,Ukraine,N
1370,Blue Dart Aviation,\N,BZ,BDA,BLUE DART,India,N
1371,B&H Airlines,\N,JA,BON,Air Bosna,Bosnia and Herzegovina,N
1372,Bissau Discovery Flying Club,\N,,BDF,BISSAU DISCOVERY,Guinea-Bissau,N
1373,Belgian Army,\N,,AYB,BELGIAN ARMY,Belgium,N
1374,Badr Airlines,\N,,BDR,BADR AIR,Sudan,N
1375,Best Aviation Ltd,\N,,BEA,BEST AIR,Bangladesh,N
1376,Belgorod Aviation Enterprise,\N,,BED,BELOGORYE,Russia,N
1377,Bar Harbor Airlines,\N,,AJC,BAR HARBOR,United States,N
1378,Balear Express,\N,,BEF,BALEAR EXPRESS,Spain,N
1379,Bel Air Helicopters,\N,,BEH,BLUECOPTER,Denmark,N
1380,Berkut Air,\N,,BEK,BERKUT,Kazakhstan,N
1381,BETA - Brazilian Express Transportes Aereos,\N,,BET,BETA CARGO,Brazil,N
1382,Basler Flight Service,\N,,BFC,BASLER,United States,N
1383,Bear Flight,\N,,BFG,BEARFLIGHT,Sweden,N
1384,Buffalo Airways,\N,J4,BFL,BUFFALO,Canada,N
1385,Bombardier,\N,,BFO,BOMBARDIER,Canada,N
1386,Burkina Airlines,\N,,BFR,BURKLINES,Burkina Faso,N
1387,Business Flight Sweden,\N,,BFS,BUSINESS FLIGHT,Sweden,N
1388,Bahrain Defence Force,\N,,BFW,SUMMAN,Bahrain,N
1389,BH Air,\N,,BGH,BALKAN HOLIDAYS,Bulgaria,N
1390,British Gulf International,\N,,BGI,BRITISH GULF,S,N
1391,British Gulf International-Fez,\N,,BGK,GULF INTER,Kyrgyzstan,N
1392,Benin Golf Air,\N,A8,BGL,BENIN GOLF,Benin,N
1393,Bugulma Air Enterprise,\N,,BGM,BUGAVIA,Russia,N
1394,Budget Air Bangladesh,\N,,BGR,BUDGET AIR,Bangladesh,N
1395,Bergen Air Transport,\N,,BGT,BERGEN AIR,Norway,N
1396,Buddha Air,\N,,BHA,BUDDHA AIR,Nepal,N
1397,Balkh Airlines,\N,,BHI,SHARIF,Afghanistan,N
1398,Bristow Helicopters,\N,,BHL,BRISTOW,United Kingdom,N
1399,Bristow Helicopters Nigeria,\N,,BHN,BRISTOW HELICOPTERS,Nigeria,N
1400,Bhoja Airlines,\N,,BHO,BHOJA,Pakistan,N
1401,Belair Airlines,\N,4T,BHP,BELAIR,Switzerland,Y
1402,Bighorn Airways,\N,,BHR,BIGHORN AIR,United States,N
1403,Bahamasair,\N,UP,BHS,BAHAMAS,Bahamas,Y
1404,Bright Air,\N,,BHT,BRIGHTAIR,Netherlands,N
1405,Bringer Air Cargo Taxi Aereo,\N,E6,,,Brazil,N
1406,Balkan Bulgarian Airlines,\N,LZ,,,,Y
1407,BA Connect,\N,TH,,,United Kingdom,N
1408,Bosphorus European Airways,\N,,BHY,BOSPHORUS,Turkey,N
1409,Binair,\N,,BID,BINAIR,Germany,N
1410,Big Island Air,\N,,BIG,BIG ISLE,United States,N
1411,British International Helicopters,\N,BS,BIH,BRINTEL,United Kingdom,Y
1412,Billund Air Center,\N,,BIL,BILAIR,Denmark,N
1413,Boise Interagency Fire Center,\N,,BIN,BISON-AIR,United States,N
1414,Bird Leasing,\N,,BIR,BIRD AIR,United States,N
1415,Bizjet Ltd,\N,,BIZ,BIZZ,United Kingdom,N
1416,Baja Air,\N,,BJA,BAJA AIR,Mexico,N
1417,Baltic Jet Aircompany,\N,,BJC,BALTIC JET,Latvia,N
1418,Business Jet Solutions,\N,,BJS,SOLUTION,United States,N
1419,Bankair,\N,B4,BKA,BANKAIR,United States,N
1420,BF-Lento OY,\N,,BKF,BAKERFLIGHT,Finland,Y
1421,Barken International,\N,,BKJ,BARKEN JET,United States,N
1422,Bangkok Airways,\N,PG,BKP,BANGKOK AIR,Thailand,Y
1423,Bukovyna,\N,,BKV,BUKOVYNA,Ukraine,N
1424,Blue Bird Aviation,\N,,BLB,BLUEBIRD SUDAN,Sudan,N
1425,Bellesavia,\N,,BLC,BELLESAVIA,Belarus,N
1426,Blue Line,\N,,BLE,BLUE BERRY,France,N
1427,Blue1,\N,KF,BLF,BLUEFIN,Finland,Y
1428,Belgavia,\N,,BLG,BELGAVIA,Belgium,N
1429,Blue Horizon Travel Club,\N,,BLH,BLUE HORIZON,United States,N
1430,Blue Jet,\N,,BLJ,BLUEWAY,Spain,N
1431,Baltic Airlines,\N,,BLL,BALTIC AIRLINES,Russia,Y
1432,Blue Airways,\N,,BLM,BLUE ARMENIA,Armenia,N
1433,Bali International Air Service,\N,,BLN,BIAR,Indonesia,N
1434,Bearskin Lake Air Service,\N,JV,BLS,BEARSKIN,Canada,Y
1435,Baltic Aviation,\N,,BLT,BALTAIR,United States,N
1436,Bellview Airlines,\N,B3,BLV,BELLVIEW AIRLINES,Nigeria,Y
1437,bmi,bmi British Midland,BD,BMA,MIDLAND,United Kingdom,Y
1438,British Medical Charter,\N,,BMD,BRITISH MEDICAL,United Kingdom,N
1439,Briggs Marine Environmental Services,\N,,BME,BRIGGS,United Kingdom,N
1440,Bristow Masayu Helicopters,\N,,BMH,MASAYU,Indonesia,N
1441,bmibaby,\N,WW,BMI,BABY,United Kingdom,Y
1442,Bemidji Airlines,\N,CH,BMJ,BEMIDJI,United States,Y
1443,Bismillah Airlines,\N,5Z,BML,BISMILLAH,Bangladesh,N
1444,Bowman Aviation,\N,,BMN,BOWMAN,United States,N
1445,British Midland Regional,\N,,BMR,,United Kingdom,Y
1446,BMW,\N,,BMW,BMW-FLIGHT,Germany,N
1447,Banco de Mexico,\N,,BMX,BANXICO,Mexico,N
1448,Bond Offshore Helicopters,\N,,BND,BOND,United Kingdom,N
1449,Benina Air,\N,,BNE,BENINA AIR,Libya,N
1450,BN Group Limited,\N,,BNG,VECTIS,United Kingdom,N
1451,Blue Nile Ethiopia Trading,\N,,BNL,NILE TRADING,Ethiopia,N
1452,Bonair Aviation,\N,,BNR,BONAIR,Canada,N
1453,Bancstar - Valley National Corporation,\N,,BNS,BANCSTAR,United States,N
1454,Bentiu Air Transport,\N,,BNT,BENTIU AIR,Sudan,N
1455,Benane Aviation Corporation,\N,,BNV,BENANE,Mauritania,N
1456,British North West Airlines,\N,,BNW,BRITISH NORTH,United Kingdom,N
1457,Boniair,\N,,BOA,KUMANOVO,Macedonia,N
1458,Bond Air Services,\N,,BOD,UGABOND,Uganda,N
1459,Boeing,\N,,BOE,BOEING,United States,N
1460,Bordaire,\N,,BOF,BORDAIR,Canada,N
1461,Bookajet Limited,\N,,BOO,BOOKAJET,United Kingdom,N
1462,Bouraq Indonesia Airlines,\N,BO,BOU,BOURAQ,Indonesia,N
1463,Blue Panorama Airlines,\N,BV,BPA,BLUE PANOROMA,Italy,Y
1464,Berkhut ZK,\N,,BPK,VENERA,Kazakhstan,N
1465,Bundepolizei-Fliegertruppe,\N,,BPO,PIROL,Germany,N
1466,Budapest Aircraft Services/Manx2,\N,,BPS,BASE,Hungary,Y
1467,Bonus Aviation,\N,,BPT,BONUS,United Kingdom,N
1468,British Petroleum Exploration,\N,,BPX,,Colombia,N
1469,BRA-Transportes Aereos,\N,7R,BRB,BRA-TRANSPAEREOS,Brazil,N
1470,Brock Air Services,\N,,BRD,BROCK AIR,Canada,N
1471,Breeze Ltd,\N,,BRE,AVIABREEZE,Ukraine,N
1472,Bering Air,\N,8E,BRG,BERING AIR,United States,Y
1473,Briansk State Air Enterprise,\N,,BRK,BRIANSK-AVIA,Russia,N
1474,Branson Airlines,\N,,BRN,BRANSON,United States,N
1475,BASE Regional Airlines,\N,,BRO,COASTRIDER,Netherlands,N
1476,Brazilian Air Force,\N,,BRS,BRAZILIAN AIR FORCE,Brazil,Y
1477,British Regional Airlines,\N,,BRT,BRITISH,United Kingdom,N
1478,Belavia Belarusian Airlines,\N,B2,BRU,BELARUS AVIA,Belarus,Y
1479,Bravo Air Congo,\N,K6,BRV,BRAVO,Democratic Republic of the Congo,N
1480,Bright Aviation Services,\N,,BRW,BRIGHT SERVICES,Bulgaria,N
1481,Braniff International Airways,\N,BN,BNF,Braniff,United States,N
1482,Buffalo Express Airlines,\N,,BRX,BUFF EXPRESS,United States,N
1483,Burundayavia,\N,,BRY,BURAIR,Kazakhstan,N
1484,Bistair - Fez,\N,,BSC,BIG SHOT,Kyrgyzstan,N
1485,Blue Star Airlines,\N,,BSD,AIRLINES STAR,Mexico,N
1486,Brasair Transportes Aereos,\N,,BSI,BRASAIR,Brazil,N
1487,Blue Swann Aviation,\N,,BSJ,BLUE SWANN,United Kingdom,N
1488,Blue Sky Aviation,\N,,BSM,,Lebanon,N
1489,Bissau Aero Transporte,\N,,BSS,BISSAU AIRSYSTEM,Guinea-Bissau,N
1490,Best Air,\N,,BST,TUNCA,Turkey,N
1491,Blue Sky Airways,\N,,BSW,SKY BLUE,Czech Republic,N
1492,Big Sky Airlines,\N,GQ,BSY,BIG SKY,United States,N
1493,BAL Bashkirian Airlines,\N,V9,BTC,BASHKIRIAN,Russia,N
1494,Baltijas Helicopters,\N,,BTH,BALTIJAS HELICOPTERS,Latvia,N
1495,Baltyka,\N,,BTK,BALTYKA,Ukraine,N
1496,Baltia Air Lines,\N,,BTL,BALTIA FLIGHT,United States,N
1497,Botir-Avia,\N,,BTR,BOTIR-AVIA,Kyrgyzstan,N
1498,Business Express,\N,,GAA,BIZEX,United States,N
1499,BT-Slavuta,\N,,BTT,BEETEE-SLAVUTA,Ukraine,N
1500,Metro Batavia,\N,7P,BTV,BATAVIA,Indonesia,Y
1501,Bulgarian Air Charter,\N,,BUC,BULGARIAN CHARTER,Bulgaria,N
1502,Blue Airlines,\N,,BUL,BLUE AIRLINES,Democratic Republic of the Congo,N
1503,Buryat Airlines Aircompany,\N,,BUN,BURAL,Russia,N
1504,Buzz Stansted,\N,,BUZ,BUZZ,United Kingdom,N
1505,Buffalo Airways,\N,,BVA,BUFFALO AIR,United States,N
1506,Bulgarian Aeronautical Centre,\N,,BVC,BULGARIAN WINGS,Bulgaria,N
1507,Baron Aviation Services,\N,,BVN,SHOW-ME,United States,N
1508,Berjaya Air,\N,J8,BVT,BERJAYA,Malaysia,Y
1509,Bellview Airlines,\N,,BVU, Sierra Leone,,N
1510,Blue Wings,\N,QW,BWG,BLUE WINGS,Germany,Y
1511,Blue Wing Airlines,\N,,BWI,BLUE TAIL,Suriname,N
1512,BWA (2002) Ltd,\N,,BWL,BRITWORLD,United Kingdom,N
1513,Bahrain Executive Air Services,\N,,BXA,BEXAIR,Bahrain,N
1514,Bar XH Air,\N,,BXH,PALLISER,Canada,N
1515,Brussels International Airlines,\N,,BXI,XENIA,Belgium,N
1516,BAX Global,\N,8W,,,,N
1517,Redding Aero Enterprises,\N,,BXR,BOXER,United States,N
1518,Berry Aviation,\N,,BYA,BERRY,United States,N
1519,Bylina Joint-Stock Company,\N,,BYL,BYLINA,Russia,N
1520,Berytos Airlines,\N,,BYR,,Lebanon,N
1521,Bayu Indonesia Air,\N,BM,BYE,BAYU,Indonesia,N
1522,Bizair Fluggesellschaft,\N,,BZA,BERLIN BEAR,Germany,N
1523,Brit Air,\N,DB,BZH,BRITAIR,France,Y
1524,Butane Buzzard Aviation Corporation,\N,,BZZ,BUZZARD,United Kingdom,N
1525,Business Flight Salzburg,\N,,AUJ,AUSTROJET,Austria,N
1526,BKS Air (Rivaflecha),\N,,CKM,COSMOS,Spain,N
1527,Bristol Flying Centre,\N,,CLF,CLIFTON,United Kingdom,N
1528,Barnes Olsen Aeroleasing,\N,,CLN,SEELINE,United Kingdom,N
1529,Baltimore Air Transport,\N,,CPJ,CORPJET,United States,N
1530,Boston-Maine Airways,\N,E9,CXS,CLIPPER CONNECTION,United States,N
1531,Brussels Airlines,SN Brussels Airlines,SN,DAT,BEE-LINE,Belgium,Y
1532,Baltimore Airways,\N,,EAH,EASTERN,United States,N
1533,Bond Aviation,\N,,EBA,BOND AVIATION,Italy,N
1534,Brazilian Army Aviation,\N,,EXB,BRAZILIAN ARMY,Brazil,N
1535,Business Express Delivery,\N,,EXP,EXPRESS AIR,Canada,N
1536,Bel Limited,\N,,FOS,,Russia,N
1537,Bangkok Aviation Center,\N,,HAW,THAI HAWK,Thailand,N
1538,Benair,\N,,HAX,SCOOP,Norway,N
1539,Binter Canarias,\N,NT,IBB,,Spain,Y
1540,Bonyad Airlines,\N,,IRJ,BONYAD AIR,Iran,N
1541,Burundaiavia,\N,,IVR,RERUN,Kazakhstan,N
1542,Blue Air,\N,0B,JOR,BLUE TRANSPORT,Romania,Y
1543,British Mediterranean Airways,\N,KJ,LAJ,BEE MED,United Kingdom,Y
1544,Belle Air,\N,,LBY,ALBAN-BELLE,Albania,N
1545,Blom Geomatics,\N,,LED,SWEEPER,Norway,N
1546,Benin Littoral Airways,\N,,LTL,LITTORAL,Benin,N
1547,Bombardier Business Jet Solutions,\N,,LXJ,FLEXJET,United States,N
1548,Bulgaria Air,\N,FB,LZB,FLYING BULGARIA,Bulgaria,Y
1549,Brazilian Navy Aviation,\N,,MBR,BRAZILIAN NAVY,Brazil,N
1550,Barents AirLink,\N,8N,NKF,NORDFLIGHT,Sweden,Y
1551,Belgian Navy,\N,,NYB,BELGIAN NAVY,Belgium,N
1552,Bakoji Airlines Services,\N,,OGJ,BAKO AIR,Nigeria,N
1553,Benders Air,\N,,PEB,PALEMA,Sweden,N
1554,Balmoral Central Contracts,\N,,PNT,PORTNET,South Africa,N
1555,BGB Air,\N,,POI,BOJBAN,Kazakhstan,N
1556,Butte Aviation,\N,,PPS,PIPESTONE,United States,N
1557,Bond Air Services,\N,,RHD,RED HEAD,United Kingdom,N
1558,Business Airfreight,\N,,RLR,RATTLER,United States,N
1559,BAC Express Airlines,\N,,RPX,RAPEX,United Kingdom,N
1560,Boscombe Down DERA (Formation),\N,,RRS,BLACKBOX,United Kingdom,N
1561,Business Jet Sweden,\N,,SCJ,SCANJET,Sweden,N
1562,British Airways Shuttle,\N,,SHT,SHUTTLE,United Kingdom,N
1563,British Sky Broadcasting,\N,,SKH,SKYNEWS,United Kingdom,N
1564,Bell Helicopter Textron,\N,,TXB,TEXTRON,Canada,N
1565,Buzzaway Limited,\N,,UKA,UKAY,United Kingdom,N
1566,Biz Jet Charter,\N,,VLX,AVOLAR,United States,N
1567,Blue Chip Jet,\N,,VOL,BLUE SPEED,Sweden,N
1568,BAE Systems,\N,,WFD,AVRO,United Kingdom,N
1569,BAE Systems,\N,,WTN,TARNISH,United Kingdom,N
1570,Baseops International,\N,,XBO,,United States,N
1571,Bureau Veritas,\N,,XDA,,France,N
1572,British Airways,\N,,XMS,SANTA,United Kingdom,N
1573,Boskovic Air Charters Limited,\N,,ZBA,BOSKY,Kenya,N
1574,C Air Jet Airlines,\N,,SRJ,SYRJET,Syrian Arab Republic,N
1575,C N Air,\N,,ORO,CAPRI,Spain,N
1576,C and M Aviation,\N,,TIP,TRANSPAC,United States,N
1577,C&M Airways,\N,,RWG,RED WING,United States,N
1578,C.S.P.,\N,,RMU, Societe,AIR-MAUR,N
1579,CABI,\N,,CBI,CABI,Ukraine,N
1580,CAI,\N,,CPI,AIRCAI,Italy,N
1581,CAL Cargo Air Lines,\N,5C,ICL,CAL,Israel,Y
1582,CAM Air Management,\N,,CMR,CAMEO,United Kingdom,N
1583,CATA L,\N,,CTZ,CATA,Argentina,N
1584,CCF Manager Airline,\N,,CCF,TOMCAT,Germany,N
1585,CEDTA (Compania Ecuatoriana De Transportes Aereos),\N,,CED,CEDTA,Ecuador,N
1586,CHC Denmark,\N,,HBI,HELIBIRD,Denmark,N
1587,CHC Helicopters,\N,,HEM,HEMS,Australia,N
1588,CHC Airways,\N,AW,SCH,,Netherlands,N
1589,CHC Helikopter Service,\N,,HKS,HELIBUS,Norway,N
1590,CI-Tours,\N,,VCI,CI-TOURS,Ivory Coast,N
1591,CKC Services,\N,,CKC,,United States,N
1592,CM Stair,\N,,CMZ,CEE-EM STAIRS,Mauritania,N
1593,CNET,\N,,CNT,KNET,France,N
1594,COAPA AIR,\N,,OAP,COAPA,Mexico,N
1595,COMAV,\N,,PDR,SPEEDSTER,Namibia,N
1596,CRI Helicopters Mexico,\N,,CRH,HELI-MEX,Mexico,N
1597,CSA Air,\N,,IRO,IRON AIR,United States,N
1598,CSE Aviation,\N,,CSE,OXFORD,United Kingdom,N
1599,CTK Network Aviation,\N,,CTQ,CITYLINK,Ghana,N
1600,Cabair College of Air Training,\N,,CBR,CABAIR,United Kingdom,N
1601,Cabo Verde Express,\N,,CVE,KABEX,Cape Verde,N
1602,Caernarfon Airworld,\N,,CWD,AMBASSADOR,United Kingdom,N
1603,Cairo Air Transport Company,\N,,CCE,,Egypt,N
1604,Cal Gulf Aviation,\N,,CGC,CAL-GULF,Sao Tome and Principe,N
1605,Cal-West Aviation,\N,,REZ,CAL AIR,United States,N
1606,California Air Shuttle,\N,,CSL,CALIFORNIA SHUTTLE,United States,N
1607,Calima Aviacion,\N,XG,CLI,CALIMA,Spain,Y
1608,Calm Air,\N,MO,CAV,CALM AIR,Canada,N
1609,Camai Air,\N,R9,CAM,AIR CAMAI,United States,N
1610,Cambodia Airlines,\N,,KHM,CAMBODIA AIRLINES,Cambodia,N
1611,Cameroon Airlines,\N,UY,UYC,CAM-AIR,Cameroon,N
1612,Campania Helicopteros De Transporte,\N,,HSO,HELIASTURIAS,Spain,N
1613,CanJet,\N,C6,CJA,CANJET,Canada,N
1614,Canada Jet Charters,\N,,PIL,PINNACLE,Canada,N
1615,Canadian Airlines,\N,CP,CDN,CANADIAN,Canada,Y
1616,Canadian Coast Guard,\N,,CTG,CANADIAN COAST GUARD,Canada,N
1617,Canadian Eagle Aviation,\N,,HIA,HAIDA,Canada,N
1618,Canadian Forces,\N,,CFC,CANFORCE,Canada,N
1619,Canadian Global Air Ambulance,\N,,BZD,BLIZZARD,Canada,N
1620,Canadian Helicopters,\N,,CDN,CANADIAN,Canada,N
1621,Canadian Interagency Forest Fire Centre,\N,,TKR,TANKER,Canada,N
1622,Canadian National Telecommunications,\N,,XNC,,Canada,N
1623,Canadian North,\N,5T,MPE,EMPRESS,Canada,Y
1624,Canadian Regional Airlines,\N,,CDR,CANADIAN REGIONAL,Canada,N
1625,Canadian Warplane Heritage Museum,\N,,CWH,WARPLANE HERITAGE,Canada,N
1626,Canadian Western Airlines,\N,W2,CWA,CANADIAN WESTERN,Canada,N
1627,Canair,\N,,CWW,CANAIR,China,N
1628,Cancun Air,\N,,CUI,CAN-AIR,Mexico,N
1629,Cape Air,\N,9K,KAP,CAIR,United States,Y
1630,Cape Air Transport,\N,,CTO,,Australia,N
1631,Cape Central Airways,\N,,SEM,SEMO,United States,N
1632,Cape Smythe Air,\N,,CMY,CAPE SMYTHE AIR,United States,N
1633,Capital Air Service,\N,,CPX,CAPAIR,United States,N
1634,Capital Airlines,\N,,CPD,CAPITAL DELTA,Kenya,N
1635,Capital Airlines Limited,\N,,NCP,CAPITAL SHUTTLE,Nigeria,N
1636,Capital Cargo International Airlines,\N,PT,CCI,CAPPY,United States,N
1637,Capital City Air Carriers,\N,,CCQ,CAP CITY,United States,N
1638,Capital Trading Aviation,\N,,EGL,PRESTIGE,United Kingdom,N
1639,Capitol Air Express,\N,,CEX,CAPITOL EXPRESS,United States,N
1640,Capitol Wings Airline,\N,,CWZ,CAPWINGS,United States,N
1641,Caravan Air,\N,,VAN,CAMEL,Mauritania,N
1642,Cardiff Wales Flying Club,\N,,CWN,CAMBRIAN,United Kingdom,N
1643,Cardinal/Air Virginia,\N,,FVA,AIR VIRGINIA,United States,N
1644,Cardolaar,\N,,GOL,CARGOLAAR,Namibia,N
1645,Cards Air Services,\N,,CDI,CARDS,Philippines,N
1646,Care Flight,\N,,CFH,CARE FLIGHT,Australia,N
1647,Carga Aerea Dominicana,\N,,CDM,CARGA AEREA,Dominican Republic,N
1648,Carga Express Internacional,\N,,EST,CARGAINTER,Mexico,N
1649,Cargo 360,\N,GG,GGC,LONG-HAUL,United States,N
1650,Cargo Express,\N,,MCX,MAURICARGO,Mauritania,N
1651,Cargo Ivoire,\N,,CRV,CARGOIV,Ivory Coast,N
1652,Cargo Link (Caribbean),\N,,CLM,CARGO LINK,Barbados,N
1653,Cargo Three,\N,,CTW,THIRD CARGO,Panama,N
1654,Cargoitalia,\N,2G,CRG,WHITE PELICAN,Italy,N
1655,Cargojet Airways,\N,W8,CJT,CARGOJET,Canada,N
1656,Cargolux,\N,CV,CLX,CARGOLUX,Luxembourg,N
1657,Cargoman,\N,,CGM,HOTEL CHARLIE,Oman,N
1658,Carib Aviation,\N,,DEL,RED TAIL,Antigua and Barbuda,N
1659,Carib Express,\N,,BCB,WAVEBIRD,Barbados,N
1660,Caribair,\N,,CBC,CARIBAIR,Dominican Republic,N
1661,Caribbean Air Cargo,\N,,DCC,CARICARGO,Barbados,N
1662,Caribbean Air Transport,\N,,CLT,CARIBBEAN,Netherlands,N
1663,Caribbean Airlines,\N,BW,BWA,CARIBBEAN AIRLINES,Trinidad and Tobago,Y
1664,Caribbean Airways,\N,,IQQ,CARIBJET,Barbados,N
1665,Caribbean Express,\N,,TLC,CARIB-X,United States,N
1666,Caribbean Star Airlines,\N,8B,GFI,CARIB STAR,Antigua and Barbuda,N
1667,Caribintair,\N,,CRT,CARIBINTAIR,Haiti,N
1668,Carill Aviation,\N,,CVG,CARILL,United Kingdom,N
1669,Carpatair,\N,V3,KRP,CARPATAIR,Romania,Y
1670,Carranza,\N,,CRR,CARRANZA,Chile,N
1671,Carroll Air Service,\N,,ULS,ULSTER,United States,N
1672,Casement Aviation,\N,,CMT,CASEMENT,United States,N
1673,Casino Airline,\N,,CSO,CASAIR,United States,N
1674,Casper Air Service,\N,,CSP,CASPER AIR,United States,N
1675,Caspian Airlines,\N,RV,CPN,CASPIAN,Iran,Y
1676,Castle Aviation,\N,,CSJ,CASTLE,United States,N
1677,Cat Aviation,\N,,CAZ,EUROCAT,Switzerland,N
1678,Catalina Flying Boats,\N,,CBT,CATALINA AIR,United States,N
1679,Catex,\N,,TEX,CATEX,France,N
1680,Cathay Pacific,\N,CX,CPA,CATHAY,Hong Kong SAR of China,Y
1681,Caverton Helicopters,\N,,CJR,CAVERTON AIR,Nigeria,N
1682,Cayman Airways,\N,KX,CAY,CAYMAN,Cayman Islands,Y
1683,Cebu Pacific,\N,5J,CEB,CEBU AIR,Philippines,Y
1684,Cecil Aviation,\N,,CIL,CECIL,United Kingdom,N
1685,Cega Aviation,\N,,CEG,CEGA,United Kingdom,N
1686,Celtic Airways,\N,,CEC,CELTAIR,United Kingdom,N
1687,Celtic West,\N,,CWE,CELTIC,United Kingdom,N
1688,Centavia,\N,7N,CNA,,Serbia,N
1689,Cente D'Essais En Vol,\N,,CEV,CENTEV,France,N
1690,Centennial Airlines,\N,,CNL,WYO-AIR,United States,N
1691,Centennial Flight Centre,\N,,CNS,CENTENNIAL,Canada,N
1692,Center Vol,\N,,CVO,CENTERVOL,Spain,N
1693,Center-South,\N,,CTS,CENTER-SOUTH,Russia,N
1694,Centrafrican Airlines,\N,,CET,CENTRAFRICAIN,Central African Republic,N
1695,Central Air Express,\N,,CAX,CENTRAL EXPRESS,Democratic Republic of Congo,N
1696,Central Airlines,\N,,CTL,CENTRAL COMMUTER,United States,N
1697,Central Airways,\N,,CNY,CENTRAL LEONE,Sierra Leone,N
1698,Central American Airlines,\N,,ACN,AEROCENTRO,Nicaragua,N
1699,Central Aviation,\N,,YOG,YOGAN AIR,United States,N
1700,Central Connect Airlines,\N,,CCG,,Czech Republic,Y
1701,Central De Discos De Reynosa,\N,,DRN,DISCOS REYNOSA,Mexico,N
1702,Central European Airlines,\N,,CMA,EUROCENTRAL,Czech Republic,N
1703,Central Flying Service,\N,,CHA,CHARTER CENTRAL,United States,N
1704,Central Mongolia Airways,\N,,CEM,CENTRAL MONGOLIA,Mongolia,N
1705,Central Mountain Air,\N,,GLR,GLACIER,Canada,N
1706,Central Queensland Aviation College,\N,,CQC,,Australia,N
1707,Central Skyport,\N,,CSI,SKYPORT,United States,N
1708,Centralwings,\N,C0,CLW,CENTRALWINGS,Poland,Y
1709,Centre Airlines,\N,,DTV,DUTCH VALLEY,United States,N
1710,Centre of Applied Geodynamica,\N,,CGS,GEO CENTRE,Russia,N
1711,Centre-Avia,\N,J7,CVC,AVIACENTRE,Russia,N
1712,Centro De Helicopteros Corporativos,\N,,CCV,HELICORPORATIVO,Mexico,N
1713,Centro de Formacion Aeronautica de Canarias,\N,,ACF,FORCAN,Spain,N
1714,Centurion Air Cargo,\N,WE,CWC,CHALLENGE CARGO,United States,N
1715,Century Aviation,\N,,URY,CENTURY AVIA,Mexico,N
1716,Century Aviation International,\N,,HAI,,Canada,N
1717,Certified Air Dispatch,\N,,XAD,,United States,N
1718,Cetraca Aviation Service,\N,,CER,CETRACA,Democratic Republic of the Congo,N
1719,Chabahar Airlines,\N,,IRU,CHABAHAR,Iran,N
1720,Chalair Aviation,\N,,CLG,CHALLAIR,France,N
1721,Chalk's Ocean Airways,\N,OP,CHK,CHALKS,United States,N
1722,Challenge Air Transport,\N,,CLS,AIRISTO,Germany,N
1723,Challenge Aviation,\N,,CHS,CHALLENGE AVIATION,Australia,N
1724,Challenge International Airlines,\N,,OFF,CHALLENGE AIR,United States,N
1725,AC Challenge Aero,\N,,CHG,SKY CHALLENGER,Ukraine,N
1726,Champagne Airlines,\N,,CPH,CHAMPAGNE,France,N
1727,Champion Air,\N,MG,CCP,CHAMPION AIR,United States,N
1728,Chanchangi Airlines,\N,,NCH,CHANCHANGI,Nigeria,N
1729,Changan Airlines,\N,2Z,CGN,CHANGAN,China,N
1730,Channel Island Aviation,\N,,CHN,CHANNEL,United States,N
1731,Chantilly Air,\N,,WML,MARLIN,United States,N
1732,Chaparral Airlines,\N,,CPL,CHAPARRAL,United States,N
1733,Chari Aviation Services,\N,S8,CSU,CHARI SERVICE,Chad,N
1734,Charlan Air Charter,\N,,CAH,CHARLAN,South Africa,N
1735,Charlie Hammonds Flying Service,\N,,HMD,HAMMOND,United States,N
1736,Charlotte Air National Guard,\N,,CGD,,United States,N
1737,Charter Air,\N,,CHW,CHARTER WIEN,Austria,Y
1738,Chartright Air,\N,,HRT,CHARTRIGHT,Canada,N
1739,Chautauqua Airlines,\N,RP,CHQ,CHAUTAUQUA,United States,Y
1740,Cheboksary Airenterprise JSC,\N,,CBB,CHEBAIR,Russia,N
1741,Chemech Aviation,\N,,CHM,,Pakistan,N
1742,Cherline,\N,,CHZ,CHERL,Russia,N
1743,Chernomor-Avia,\N,,CMK,CHERAVIA,Russia,N
1744,Cherokee Express,\N,,CBM,BLUE MAX,United States,N
1745,Cherry Air,\N,,CCY,CHERRY,United States,N
1746,Chesapeake Air Service,\N,,CAB,CHESAPEAKE AIR,United States,N
1747,Chevron U.S.A,\N,,CVR,CHEVRON,United States,N
1748,Cheyenne Airways,\N,,CYA,CHEYENNE AIR,United States,N
1749,Chicago Air,\N,,CGO,WILD ONION,United States,N
1750,Chicago Express,\N,C8,WDY,WINDY CITY,United States,N
1751,Chief Rat Flight Services,\N,,RAT,RIVERRAT,South Africa,N
1752,Chilchota Taxi Aereo,\N,,CCH,CHILCHOTA,Mexico,N
1753,Chilcotin Caribou Aviation,\N,,DES,CHILCOTIN,Canada,N
1754,Chilliwack Aviation,\N,,CAD,CHILLIWACKAIR,Canada,N
1755,Chim-Nir Aviation,\N,,ETN,CHMINIR,Israel,N
1756,China Airlines,\N,CI,CAL,DYNASTY,Taiwan,Y
1757,China Cargo Airlines,\N,CK,CKK,CARGO KING,China,N
1758,China Eastern Airlines,\N,MU,CES,CHINA EASTERN,China,Y
1759,China Express Airlines,\N,,HXA,CHINA EXPRESS,China,N
1760,China Flying Dragon Aviation,\N,,CFA,FEILONG,China,N
1761,China General Aviation Corporation,\N,,CTH,TONGHANG,China,N
1762,China National Aviation Corporation,\N,,CAG,CHINA NATIONAL,China,N
1763,China Northern Airlines,\N,CJ,CBF,CHINA NORTHERN,China,N
1764,China Northwest Airlines,\N,WH,CNW,CHINA NORTHWEST,China,N
1765,China Ocean Helicopter Corporation,\N,,CHC,CHINA HELICOPTER,China,N
1766,China Postal Airlines,\N,8Y,CYZ,CHINA POST,China,N
1767,China Southern Airlines,\N,CZ,CSN,CHINA SOUTHERN,China,Y
1768,China Southwest Airlines,\N,,CXN,CHINA SOUTHWEST,China,N
1769,China United Airlines,\N,HR,CUA,LIANHANG,China,Y
1770,China Xinhua Airlines,\N,XO,CXH,XINHUA,China,N
1771,Yunnan Airlines,\N,3Q,CYH,YUNNAN,China,Y
1772,Chinguetti Airlines,\N,,CGU,CHINGUETTI,Mauritania,N
1773,Chipola Aviation,\N,,CEP,CHIPOLA,United States,N
1774,Chippewa Air Commuter,\N,,CPW,CHIPPEWA-AIR,United States,N
1775,Chitaavia,\N,X7,CHF,CHITA,Russia,N
1776,Christman Air System,\N,,CAS,CHRISTMAN,United States,N
1777,Christophorus Flugrettungsverein,\N,,OEC,CHRISTOPHORUS,Austria,N
1778,Chrome Air Services,\N,,CHO,CHROME AIR,United States,N
1779,Church Aircraft,\N,,CHU,CHURCHAIR,United States,N
1780,Cielos Airlines,\N,A2,CIU,CIELOS,Peru,N
1781,Cimber Air,\N,QI,CIM,CIMBER,Denmark,Y
1782,Cirrus,\N,,RRU,HELICIRRUS,Chile,N
1783,Cirrus Air,\N,,NTS,NITE STAR,United States,N
1784,Cirrus Airlines,\N,C9,RUS,CIRRUS AIR,Germany,Y
1785,Cirrus Middle East,\N,,JTI,,Lebanon,N
1786,Citationshares,\N,,FIV,FIVE STAR,United States,N
1787,Citibank,\N,,XCX,,United States,N
1788,Citic General Aviation,\N,,HZX,ZHONGXIN,China,N
1789,City Airline,\N,CF,SDR,SWEDESTAR,Sweden,Y
1790,City Connexion Airlines,\N,G3,CIX,CONNEXION,Burundi,Y
1791,City of Bangor,\N,,XBG,,United States,N
1792,CityJet,\N,WX,BCY,CITY-IRELAND,Ireland,Y
1793,Cityair (Chester) Limited,\N,,CAQ,AIR CHESTER,United Kingdom,N
1794,Cityfly,\N,,CII,CITYFLY,Italy,N
1795,BA CityFlyer,\N,CJ,CFE,FLYER,United Kingdom,Y
1796,Cityline Hungary,\N,,CNB,CITYHUN,Hungary,N
1797,Citylink Airlines,\N,,HSR,HOOSIER,United States,N
1798,Civair Airways,\N,,CIW,CIVFLIGHT,South Africa,N
1799,Civil Air Patrol South Carolina Wing,\N,,BKR,BOX KAR,United States,N
1800,Civil Air Transport,\N,CT,CAT,Mandarin,Taiwan,N
1801,Civil Aviation Authority,\N,,CIA,CALIMERA,Slovakia,N
1802,Civil Aviation Authority of New Zealand,\N,,CIV,CIVAIR,New Zealand,N
1803,Civil Aviation Inspectorate of the Czech Republic,\N,,CBA,CALIBRA,Czech Republic,N
1804,Claessens International Limited,\N,,FMC,CLAESSENS,United Kingdom,N
1805,Clark Aviation,\N,,CLK,CLARKAIR,United States,N
1806,Clasair,\N,,CSF,CALEDONIAN,United Kingdom,N
1807,Clay Lacy Aviation,\N,,CLY,CLAY-LACY,United States,N
1808,Click Airways,\N,,CGK,CLICK AIR,Kyrgyzstan,Y
1809,Cloud 9 Air Charters,\N,,CLZ,CLOUDLINE,South Africa,N
1810,Clowes Estates Limited,\N,,CLD,CLOWES,United Kingdom,N
1811,Club 328,\N,,SDJ,SPACEJET,United Kingdom,N
1812,Club Air,\N,6P,ISG,CLUBAIR,Italy,N
1813,Coast Air,\N,BX,CST,COAST CENTER,Norway,N
1814,Coastal Air,\N,DQ,,U.S. Virgin Islands,United States,Y
1815,Coastal Air Transport,\N,,TCL,TRANS COASTAL,United States,N
1816,Coastal Airways,\N,,CNG,SID-AIR,United States,N
1817,Coastal Travels,\N,,CSV,COASTAL TRAVEL,Tanzania,N
1818,Cohlmia Aviation,\N,,CHL,COHLMIA,United States,N
1819,Colaereos,\N,,OLR,COLAEREOS,Ecuador,N
1820,Colemill Enterprises,\N,,CLE,COLEMILL,United States,N
1821,Colgan Air,\N,9L,CJC,COLGAN,United States,Y
1822,Colibri Aviation,\N,,CAE,HUMMINGBIRD,Canada,N
1823,Colt International,\N,,CCX,,United States,N
1824,Columbia Airlines,\N,,COL,COLAIR,Canada,N
1825,Columbia Helicopters,\N,,WCO,COLUMBIA HELI,United States,N
1826,Columbus Air Transport,\N,,KLR,KAY-LER,United States,N
1827,Colvin Aviation,\N,,GHP,GRASSHOPPER EX,United States,N
1828,Comair,\N,OH,COM,COMAIR,United States,Y
1829,Comair,\N,MN,CAW,COMMERCIAL,South Africa,Y
1830,Comed Group,\N,,CDE,COMEX,United Kingdom,N
1831,Comeravia,\N,,CVV,COMERAVIA,Venezuela,N
1832,Comercial Aerea,\N,,CRS,COMERCIAL AEREA,Mexico,N
1833,Comet Airlines,\N,,CMG,SUNSPY,Nigeria,N
1834,Comfort Air,\N,,FYN,FLYNN,Germany,N
1835,Comlux Aviation,\N,,CLA,COMLUX,Switzerland,N
1836,Commair Aviation,\N,,CMH,COMMODORE,United Kingdom,N
1837,Commandement Du Transport Aerien Militaire Francais,\N,,CTM,COTAM,France,N
1838,Commander Air Charter,\N,,CML,COMMANDAIR,Canada,N
1839,Commander Mexicana,\N,,CRM,COMMANDERMEX,Mexico,N
1840,Commercial Aviation,\N,,CMS,ACCESS,Canada,N
1841,Commodore Aviation,\N,,GAR,,Australia,N
1842,Commonwealth Jet Service,\N,,CJS,COMMONWEALTH,United States,N
1843,CommutAir,\N,C5,UCA,COMMUTAIR,United States,Y
1844,Comores Airlines,\N,KR,CWK,CONTICOM,Comoros,Y
1845,Compagnia Generale Ripreseaeree,\N,,CGR,COMPRIP,Italy,N
1846,Compagnie Aerienne du Mali,\N,,CMM,CAMALI,Mali,N
1847,Compagnie Mauritanienne Des Transports,\N,,CPM,,Mauritania,N
1848,Compagnie de Bauxites de Guinee,\N,,GIC,CEBEGE,Guinea,N
1849,Compania Aerea de Valencia,\N,,AIF,,Spain,N
1850,Compania Aerotecnicas Fotograficas,\N,,ATF,AEROTECNICAS,Spain,N
1851,Compania Boliviana de Transporte Aereo Privado Aerosur,\N,,ASU, S.A.,ASUR,N
1852,Compania De Actividades Y Servicios De Aviacion,\N,,LCT,STELLAIR,Spain,N
1853,Compania Ejecutiva,\N,,EJV,EJECUTIVA,Mexico,N
1854,Compania Helicopteros Del Sureste,\N,,HSE,HELISURESTE,Spain,N
1855,Compania Mexicana De Aeroplanos,\N,,MDR,AEROPLANOS,Mexico,N
1856,Compania Mexicargo,\N,GJ,MXC,MEXICARGO,Mexico,N
1857,Compania Transportes Aereos Del Sur,\N,,HSS,TAS HELICOPTEROS,Spain,N
1858,Compania de Servicios Aereos Tavisa,\N,,TAV,TAVISA,Spain,N
1859,Company Flight,\N,,CYF,COMPANY FLIGHT,Denmark,N
1860,Compass Airlines,\N,CP,CPZ,Compass Rose,United States,Y
1861,Compass International Airways,\N,,CPS,COMPASS,United Kingdom,N
1862,Compuflight Operations Service,\N,,XCO,,United States,N
1863,Compuserve Incorporated,\N,,XCS,,United States,N
1864,Conair Aviation,\N,,CRC,CONAIR-CANADA,Canada,N
1865,Concordavia,\N,,COD,CONCORDAVIA,Ukraine,N
1866,Condor Aero Services,\N,,CNR,CONAERO,United States,N
1867,Condor,\N,,CIB,CONDOR BERLIN,Germany,N
1868,Condor Flugdienst,\N,DE,CFG,CONDOR,Germany,Y
1869,Confort Air,\N,,COF,CONFORT,Canada,N
1870,Congo Air,\N,,CAK,,Bahamas,N
1871,Congressional Air,\N,,CGA,CONGRESSIONAL,United States,N
1872,Conifair Aviation,\N,,ROY,,Canada,N
1873,Connectair Charters,\N,,BSN,BASTION,Canada,N
1874,Conquest Airlines,\N,,CAC,CONQUEST AIR,United States,N
1875,Conroe Aviation Services,\N,,CXO,CONROE AIR,United States,N
1876,Consorcio Aviaxsa,\N,6A,CHP,AVIACSA,Mexico,Y
1877,Consorcio Helitec,\N,,VCH,CONSORCIO HELITEC,Venezuela,N
1878,Constanta Airline,\N,,UZA,CONSTANTA,Ukraine,N
1879,Contact Air,Contactair,C3,KIS,CONTACTAIR,Germany,Y
1880,Contel ASC,\N,,XCL,,United States,N
1881,Continental Airlines,\N,CO,COA,CONTINENTAL,United States,Y
1882,Continental Airways,\N,PC,PVV,CONTAIR,Russia,N
1883,Continental Express,\N,CO,,JETLINK,United States,Y
1884,Continental Micronesia,\N,CS,CMI,AIR MIKE,United States,Y
1885,Continental Oil,\N,,CON,CONOCO,United States,N
1886,Conviasa,\N,V0,VCV,CONVIASA,Venezuela,Y
1887,Cook Inlet Aviation,\N,,CKA,COOK-AIR,United States,N
1888,Cooper Aerial Surveys,\N,,SVY,SURVEYOR,United Kingdom,N
1889,Copa Airlines,\N,CM,CMP,COPA,Panama,Y
1890,Copenhagen Airtaxi,\N,,CAT,AIRCAT,Denmark,N
1891,Copper State Air Service,\N,,COP,COPPER STATE,United States,N
1892,Copterline,\N,,AAQ,COPTERLINE,Finland,Y
1893,Coptrade Air Transport,\N,,CCW,COPTRADE AIR,Sudan,N
1894,Corendon Airlines,\N,,CAI,CORENDON,Turkey,Y
1895,Coronado Aerolineas,\N,,CRA,CORAL,Colombia,N
1896,Corpac Canada,\N,,CPB,PENTA,Canada,N
1897,Corporacion Aereo Cencor,\N,,CNC,CENCOR,Mexico,N
1898,Corporacion Aeroangeles,\N,,CPG,CORPORANG,Mexico,N
1899,Corporacion Paraguaya De Aeronautica,\N,,CGY,,Paraguay,N
1900,Corporate Air,\N,,CPT,AIR SPUR,United States,N
1901,Corporate Air,\N,,CPR,CORPAIR,United States,N
1902,Corporate Aircraft Company,\N,,CPO,MOKAN,United States,N
1903,Corporate Airlink,\N,,COO,CORPORATE,Canada,N
1904,Corporate Aviation Services,\N,,CKE,CHECKMATE,United States,N
1905,Corporate Flight International,\N,,VHT,VEGAS HEAT,United States,N
1906,Corporate Flight Management,\N,,VTE,VOLUNTEER,United States,N
1907,Corporate Jets,\N,,CJI,SEA JET,United States,N
1908,Corsairfly,\N,SS,CRL,CORSAIR,France,Y
1909,Corse-Mediterranee,\N,XK,CCM,CORSICA,France,Y
1910,Cosmic Air,\N,F5,COZ,COSMIC AIR,Nepal,N
1911,Cougar Helicopters,\N,,CHI,COUGAR,Canada,N
1912,Coulson Flying Service,\N,,MGB,MOCKINGBIRD,United Kingdom,N
1913,Country Connection Airlines,\N,,NSW,,Australia,N
1914,Country International Airlines,\N,,CIK,COUNTRY AIR,Kyrgyzstan,N
1915,Courier Services,\N,,CSD,DELIVERY,United States,N
1916,Court Helicopters,\N,,CUT,COURT AIR,South Africa,N
1917,Coval Air,\N,,CVL,COVAL,Canada,N
1918,Cowi,\N,,COW,COWI,Netherlands,N
1919,Coyne Aviation,\N,,COY,COYNE AIR,United Kingdom,N
1920,Cranfield University,\N,,CFD,AERONAUT,United Kingdom,N
1921,Cree Airways,\N,,CRE,CREE AIR,Canada,N
1922,Crelam,\N,,ELM,CRELAM,Mexico,N
1923,Crest Aviation,\N,,CAN,CREST,United Kingdom,Y
1924,Crimea Universal Avia,\N,,KRM,TRANS UNIVERSAL,Ukraine,N
1925,Croatia Airlines,\N,OU,CTN,CROATIA,Croatia,Y
1926,Croatian Air Force,\N,,HRZ,CROATIAN AIRFORCE,Croatia,N
1927,Cross Aviation,\N,,CRX,CROSSAIR,United Kingdom,N
1928,Crossair Europe,\N,QE,ECC,Cigogne,Switzerland,N
1929,Crow Executive Air,\N,,CWX,CROW EXPRESS,United States,N
1930,Crown Air Systems,\N,,CKR,CROWN AIR,United States,N
1931,Crown Airways,\N,,CRO,CROWN AIRWAYS,United States,Y
1932,Crownair,\N,,CRW,REGAL,Canada,N
1933,Cruiser Linhas Aereas,\N,,VCR,VOE CRUISER,Brazil,N
1934,Cryderman Air Service,\N,,CTY,CENTURY,United States,N
1935,Crystal Shamrock,\N,,CYT,CRYSTAL-AIR,United States,N
1936,Cubana de Aviación,\N,CU,CUB,CUBANA,Cuba,Y
1937,Cumberland Airways (Nicholson Air Service),\N,,CBL,CUMBERLAND,United States,N
1938,Custom Air Transport,\N,,CTT,CATT,United States,N
1939,Cygnus Air,\N,,RGN,CYGNUS AIR,Spain,N
1940,Cyprair Tours,\N,,CYC,CYPRAIR,Cyprus,N
1941,Cypress Airlines,\N,,CYS,SKYBIRD,Canada,N
1942,Cyprus Airways,\N,CY,CYP,CYPRUS,Cyprus,Y
1943,Cyprus Turkish Airlines,\N,YK,,,Turkey,Y
1944,Czech Air Force,\N,,CEF,CZECH AIR FORCE,Czech Republic,N
1945,Czech Air Handling,\N,,AHD,AIRHANDLING,Czech Republic,N
1946,Czech Airlines,CSA Czech Airlines,OK,CSA,CSA-LINES,Czech Republic,Y
1947,Czech Government Flying Service,\N,,CIE,CZECHIA,Czech Republic,N
1948,D & D Aviation,\N,,DDA,DUSTY,United States,N
1949,D&K Aviation,\N,,DNK,DIRECT JET,United States,N
1950,DAP Helicopteros,\N,,DHE,HELIDAP,Chile,N
1951,DFS UK Limited,\N,,VLF,VOLANTE,United Kingdom,N
1952,DAS Air Cargo,\N,WD,DSR,DAIRAIR,Uganda,N
1953,DAS Airlines,\N,,RKC,DAS CONGO,Democratic Republic of the Congo,N
1954,DAT Danish Air Transport,\N,DX,DTR,DANISH,Denmark,Y
1955,DAT Enterprise Limited,\N,,ENT,DATENT,United Kingdom,N
1956,DERA Boscombe Down,\N,,BDN,GAUNTLET,United Kingdom,N
1957,DESNA,\N,,DSN,DESNA,Ukraine,N
1958,DETA Air,\N,,DET,SAMAL,Kazakhstan,N
1959,DGO Jet,\N,,DGO,DGO JET,Mexico,N
1960,DHL Aero Expreso,\N,,DAE,YELLOW,Panama,N
1961,DHL Air,\N,,DHK,WORLD EXPRESS,United Kingdom,N
1962,DHL Aviation,\N,,DHV,WORLDSTAR,South Africa,N
1963,DHL International,\N,ES,DHX,DILMUN,Bahrain,N
1964,DHL de Guatemala,\N,L3,JOS,,Guatemala,N
1965,DSWA,\N,,RSK,REDSKIN,United States,N
1966,Daallo Airlines,\N,D3,DAO,DALO AIRLINES,Djibouti,Y
1967,Dagestan Airlines,\N,N2,DAG,DAGAL,Russia,N
1968,Dahla Airlines,\N,,DHA,,Democratic Republic of Congo,N
1969,Daimler Chrysler Aviation,\N,,DCS,TWIN STAR,Germany,N
1970,Daimler-Chrysler,\N,,DCX,DAIMLER,United States,N
1971,Daka,\N,,DKA,,Kazakhstan,N
1972,Dala Air Services,\N,,DLR,DALA AIR,Nigeria,N
1973,Dalavia,\N,H8,KHB,DALAVIA,Russia,Y
1974,Dallas Express Airlines,\N,,DXP,DALLAS EXPRESS,United States,N
1975,Damascene Airways,\N,,DAS,AIRDAM,Syrian Arab Republic,N
1976,Danbury Airways,\N,,DSA,DANBURY AIRWAYS,United States,N
1977,Dancopter,\N,,DOP,DANCOPTER,Denmark,N
1978,Danish Air Force,\N,,DAF,DANISH AIRFORCE,Denmark,N
1979,Danish Army,\N,,DAR,DANISH ARMY,Denmark,N
1980,Danish Navy,\N,,DNY,DANISH NAVY,Denmark,N
1981,Danu Oro Transportas,\N,,DNU,DANU,Lithuania,N
1982,Darta,\N,,DRT,DARTA,France,N
1983,Darwin Airline,\N,0D,DWT,DARWIN,Switzerland,Y
1984,Dasab Airlines,\N,,DSQ,DASAB AIR,Uganda,N
1985,Dash Air Charter,\N,,DSH,DASH CHARTER,United States,N
1986,Dash Aviation,\N,,GOB,PILGRIM,United Kingdom,N
1987,Dasnair,\N,,DGX,DASNA,Switzerland,N
1988,Dassault Aviation,\N,,DAB,,France,N
1989,Dassault Falcon Jet Corporation,\N,,CVF,CLOVERLEAF,United States,N
1990,Dassault Falcon Service,\N,,DSO,DASSAULT,France,N
1991,Data International,\N,,DTN,DATA AIR,Sudan,N
1992,Date Transformation Corp,\N,,XDT,,United States,N
1993,Dauair,\N,D5,DAU,DAUAIR,Germany,N
1994,David Crawshaw Consultants Limited,\N,,DCO,,United Kingdom,N
1995,Dawn Air,\N,,DWN,DAWN AIR,United States,N
1996,DayJet,\N,,DJS,DAYJET,United States,N
1997,Daya Aviation,\N,,DAY,DAYA,Sri Lanka,N
1998,De Havilland,\N,,DHC,DEHAVILLAND,Canada,N
1999,Deadalos Flugtbetriebs,\N,,IAY,IASON,Austria,N
2000,Decatur Aviation,\N,,DAA,DECUR,United States,N
2001,Deccan Aviation,\N,,DKN,DECCAN,India,N
2002,Deccan Aviation (Lanka),\N,,DLK,DEKKANLANKA,Sri Lanka,N
2003,Deer Jet,\N,,DER,DEER JET,China,N
2004,Deere and Company,\N,,JDC,JOHN DEERE,United States,N
2005,Delaware Skyways,\N,,DWR,DELAWARE,United States,N
2006,Delta Aerotaxi,\N,,DEA,JET SERVICE,Italy,Y
2007,Delta Air Charter,\N,,SNO,SNOWBALL,Canada,N
2008,Delta Air Elite,\N,,ELJ,ELITE JET,United States,N
2009,Delta Air Lines,\N,DL,DAL,DELTA,United States,Y
2010,Delta Engineering Aviation,\N,,KMB,KEMBLEJET,United Kingdom,N
2011,Delta Express International,\N,,DLI,DELTA EXPRESS,Ukraine,N
2012,Delta State University,\N,,DSU,DELTA STATE,United States,N
2013,Denim Air,\N,,DNM,DENIM,Netherlands,Y
2014,Denver Express,\N,,FEC,FALCON EXPRESS,United States,N
2015,Denver Jet,\N,,DJT,DENVER JET,United States,N
2016,Departmento De Agricultura De La Generalitat De Cataluna,\N,,FGC,FORESTALES,Spain,N
2017,Deraya Air Taxi,\N,,DRY,DERAYA,Indonesia,N
2018,Des R Cargo Express,\N,,DRX,,Mauritania,N
2019,Desarrollo Milaz,\N,,MIZ,MILAZ,Mexico,N
2020,Destiny Air Services,\N,,DTY,DESTINY,Sierra Leone,N
2021,Deutsche Bahn,\N,2A,,,Germany,Y
2022,Deutsche Rettungsflugwacht,\N,1I,AMB,CIVIL AIR AMBULANCE,Germany,N
2023,Deutsches Zentrum fur Luft-und Raumfahrt EV,\N,,LFO,LUFO,Germany,N
2024,Di Air,\N,,DIS,DI AIR,Serbia,N
2025,Diamond Aviation,\N,,SPK,SPARKLE,United States,N
2026,Didier Rousset Buy,\N,,DRB,DIDIER,Chile,N
2027,Digital Equipment Corporation,\N,,DGT,DIGITAL,United States,N
2028,Dinar,\N,D7,RDN,AERO DINAR,Argentina,N
2029,Diplomatic Freight Services,\N,,DIP,DIPFREIGHT,United Kingdom,N
2030,Direccion General de Aviacion Civil y Telecomunicasciones,\N,,ENA,ENA,Spain,N
2031,Direct Air,\N,,DIA,BLUE SKY,United States,N
2032,Direct Air trading as Midway Connection,\N,,XAP,MID-TOWN,United States,N
2033,Direct Flight,\N,,DCT,,United Kingdom,N
2034,Sky Express,\N,,SXP,EXPRESS SKY,Poland,N
2035,Dirgantara Air Service,\N,AW,DIR,DIRGANTARA,Indonesia,N
2036,Discover Air,\N,,DCV,DISCOVER,United States,N
2037,Discovery Airways,\N,DH,DVA,DISCOVERY AIRWAYS,United States,N
2038,Dispatch Services,\N,,XDS,,United States,N
2039,Dix Aviation,\N,,DIX,DIX FLIGHT,Germany,N
2040,Dixie Airways,\N,,DEE,TACAIR,United States,N
2041,Djibouti Airlines,\N,D8,DJB,DJIBOUTI AIR,Djibouti,Y
2042,Dniproavia,\N,,UDN,DNIEPRO,Ukraine,Y
2043,Flying Dolphin Airlines,\N,,FDN,FLYING DOLPHIN,United Arab Emirates,N
2044,Dolphin Express Airlines,\N,,IXX,ISLAND EXPRESS,United States,N
2045,Dome Petroleum,\N,,DPL,DOME,Canada,N
2046,Dominguez Toledo (Grupo Mayoral),\N,,MYO,MAYORAL,Spain,N
2047,Dominicana de Aviaci,\N,DO,DOA,DOMINICANA,Dominican Republic,Y
2048,Domodedovo Airlines,\N,E3,DMO,DOMODEDOVO,Russia,Y
2049,Don Avia,\N,,DVB,DONSEBAI,Kazakhstan,N
2050,Donair Flying Club,\N,,DON,DONAIR,United Kingdom,N
2051,DonbassAero,\N,5D,UDC,DONBASS AERO,Ukraine,Y
2052,Dorado Air,\N,,DAD,DORADO AIR,Dominican Republic,N
2053,Dornier,\N,,DOR,DORNIER,Germany,N
2054,Dornier Aviation Nigeria,\N,,DAV,DANA AIR,Nigeria,N
2055,Dos Mundos,\N,,DOM,DOS MUNDOS,Dominican Republic,N
2056,Dragonair,\N,KA,HDA, Hong Kong Dragon Airlines,DRAGON,Y
2057,Dreamcatcher Airways,\N,,DCA,DREAM CATCHER,United Kingdom,N
2058,Druk Air,\N,KB,DRK,ROYAL BHUTAN,Bhutan,Y
2059,Drummond Island Air,\N,,DRE,MICHIGAN,United States,N
2060,Dubai Airwing,\N,,DUB,DUBAI,United Arab Emirates,N
2061,Dubrovnik Air,\N,,DBK,SEAGULL,Croatia,Y
2062,Ducair,\N,,DUK,LION KING,Luxembourg,N
2063,Duchess of Britany (Jersey) Limited,\N,,DBJ,DUCHESS,United Kingdom,N
2064,UK Royal/HRH Duke of York,\N,,LPD,LEOPARD,United Kingdom,N
2065,Dun'Air,\N,,DUN,DUNAIR,Mauritania,N
2066,Duncan Aviation,\N,,PHD,PANHANDLE,United States,N
2067,Dunyaya Bakis Hava Tasimaciligi,\N,,VVF,WORLDFOCUS,Turkey,N
2068,Duo Airways,\N,,DUO,FLY DUO,United Kingdom,N
2069,Durango Jet,\N,,DJE,DURANGO JET,Mexico,N
2070,Dutch Antilles Express,\N,,DNL,DUTCH ANTILLES,Netherlands Antilles,Y
2071,Dutch Caribbean Express,\N,,DCE,DUTCH CARIBBEAN,Netherlands Antilles,N
2072,Dutchbird,\N,,DBR,DUTCHBIRD,Netherlands,N
2073,Dwyer Aircraft Services,\N,,DFS,DWYAIR,United States,N
2074,Dynair Services,\N,,XDY,,United States,N
2075,Dynamair Aviation,\N,,DNR,DYNAMAIR,Canada,N
2076,Dynamic Air,\N,,DYE,DYNAMIC,Netherlands,N
2077,dba,\N,DI,BAG,SPEEDWAY,Germany,Y
2078,E H Darby Aviation,\N,,EHD,PLATINUM AIR,United States,N
2079,Electronic Data Systems,\N,1C,,,Switzerland,N
2080,Electronic Data Systems,\N,,1Y,,,N
2081,EAA Escola De Aviacao Aerocondor,\N,,EAD,AERO-ESCOLA,Portugal,N
2082,Executive Airlines Services,\N,,EXW,ECHOLINE,Nigeria,N
2083,EFAOS- Agencia De Viagens e Turismo,\N,,EFS,EFAOS,Angola,N
2084,EFD Eisele Flugdienst,\N,,EFD,EVER FLIGHT,Germany,N
2085,EFS-Flugservice,\N,,FSD,FLUGSERVICE,Germany,N
2086,EIS Aircraft,\N,,EIS,COOL,Germany,N
2087,EPAG Groupe Air France,\N,,IAG,EPAG,France,N
2088,ESI Eliservizi Italiani,\N,,ESI,ELISERVIZI,Italy,N
2089,EU Airways,\N,,EUY,EUROAIRWAYS,Ireland,N
2090,EUjet,\N,VE,EUJ,UNION JET,Ireland,N
2091,EVA Air,\N,BR,EVA,EVA,Taiwan,Y
2092,Eagle Aero,\N,,ICR,ICARUS FLIGHTS,United States,N
2093,Eagle Air,\N,,EGR,EAGLE SIERRA,Sierra Leone,N
2094,Eagle Air,\N,H7,,,Uganda,Y
2095,Eagle Air Company,\N,,EGX,THAI EAGLE,Thailand,N
2096,Eagle Air Iceland,\N,,FEI,ARCTIC EAGLE,Iceland,N
2097,Eagle Aviation,\N,,EGU,AFRICAN EAGLE,Uganda,N
2098,Eagle Aviation,\N,,GYP,GYPSY,United Kingdom,N
2099,Eagle Aviation France,\N,,EGN,FRENCH EAGLE,France,N
2100,Eagle International,\N,,SEG,SEN-EAGLE,Senegal,N
2101,Eagle Jet Charter,\N,,EGJ,EAGLE JET,United States,N
2102,Eaglemed (Ballard Aviation),\N,,EMD,EAGLEMED,United States,N
2103,Earth Airlines Services,\N,,ERX,EARTH AIR,Nigeria,N
2104,East African,\N,QU,UGX,CRANE,Uganda,Y
2105,East African Safari Air,\N,S9,HSA,DUMA,Kenya,N
2106,East African Safari Air Express,\N,,EXZ,TWIGA,Kenya,N
2107,East Asia Airlines,\N,,EMU,,Macao,N
2108,East Coast Airways,\N,,ECT,EASTWAY,South Africa,N
2109,East Coast Jets,\N,,ECJ,EASTCOAST JET,United States,N
2110,East Hampton Aire,\N,,EHA,AIRE HAMPTON,United States,N
2111,East Kansas City Aviation,\N,,EKC,BLUE GOOSE,United States,N
2112,East Midlands Helicopters,\N,,CTK,COSTOCK,United Kingdom,N
2113,East Star Airlines,\N,,DXH,EAST STAR,China,N
2114,East-West Airlines,\N,,EWA,EASTWEST,Australia,N
2115,Eastern Air,\N,,EAZ,EASAIR,Zambia,N
2116,Eastern Air Executive,\N,,EAX,EASTEX,United Kingdom,N
2117,Eastern Airways,\N,T3,EZE,EASTFLIGHT,United Kingdom,Y
2118,Eastern Australia Airlines,\N,,EAQ,,Australia,N
2119,Eastern Carolina Aviation,\N,,ECI,EASTERN CAROLINA,United States,N
2120,Eastern Executive Air Charter,\N,,GNS,GENESIS,United Kingdom,N
2121,Eastern Express,\N,,LIS,LARISA,Kazakhstan,N
2122,Eastern Metro Express,\N,,EME,EMAIR,United States,N
2123,Eastern Pacific Aviation,\N,,EPB,EAST PAC,Canada,N
2124,Eastern Sky Jets,\N,,ESJ,EASTERN SKYJETS,United Arab Emirates,N
2125,Eastland Air,\N,DK,ELA,,Australia,Y
2126,Eastwind Airlines,\N,W9,SGR,STINGER,United States,N
2127,Easy Link Aviation Services,\N,,FYE,FLYME,Nigeria,N
2128,Eckles Aircraft,\N,,CMN,CIMMARON AIRE,United States,N
2129,Eclipse Aviation,\N,,EJT,ECLIPSE JET,United States,N
2130,Eco Air,\N,,ECQ,SKYBRIDGE,Nigeria,N
2131,Ecoair,\N,,DEI,,Algeria,N
2132,Ecomex Air Cargo,\N,,ECX,AIR ECOMEX,Angola,N
2133,Ecotour,\N,,ECD,ECOTOUR,Mexico,N
2134,Ecoturistica de Xcalak,\N,,XCC,XCALAK,Mexico,N
2135,Ecuatoguineana De Aviacion (EGA),\N,,ECV,EQUATOGUINEA,Equatorial Guinea,N
2136,Ecuatorial Cargo,\N,,EQC,ECUA-CARGO,Equatorial Guinea,N
2137,Ecuavia,\N,,ECU,ECUAVIA,Ecuador,Y
2138,Edelweiss Air,\N,WK,EDW,EDELWEISS,Switzerland,Y
2139,Edgartown Air,\N,,SLO,SLOW,United States,N
2140,Edinburgh Air Charter,\N,,EDC,SALTIRE,United Kingdom,N
2141,Edwards Jet Center of Montana,\N,,EDJ,EDWARDS,United States,N
2142,Efata Papua Airlines,\N,,EIJ,EFATA,Indonesia,N
2143,Egyptair,\N,MS,MSR,EGYPTAIR,Egypt,Y
2144,Egyptair Cargo,\N,,MSX,EGYPTAIR CARGO,Egypt,N
2145,Egyptian Air Force,\N,,EGY,,Egypt,N
2146,Egyptian Aviation,\N,,EJX,,Egypt,N
2147,Egyptian Aviation Company,\N,,EMA,,Egypt,N
2148,Ei Air Exports,\N,,EIX,AIR EXPORTS,Ireland,N
2149,Eirjet,\N,,EIR,EIRJET,Ireland,N
2150,El Al Israel Airlines,\N,LY,ELY,ELAL,Israel,Y
2151,El Caminante Taxi Aereo,\N,,CMX,EL CAMINANTE,Mexico,N
2152,El Quilada International Aviation,\N,,GLQ,QUILADA,Sudan,N
2153,El Sal Air,\N,,ELS,EL SAL,El Salvador,N
2154,El Sol De America,\N,,ESC,SOLAMERICA,Venezuela,N
2155,El-Buraq Air Transport,\N,UZ,BRQ,BURAQAIR,Libya,Y
2156,Elan Express,\N,,ELX,ELAN,United States,N
2157,Elbe Air Transport,\N,,LBR,MOTION,Germany,N
2158,Elbrus-Avia Air Enterprise,\N,,NLK,ELAVIA,Russia,N
2159,Eldinder Aviation,\N,,DND,DINDER,Sudan,N
2160,Elicar,\N,,PDV,ELICAR,Italy,N
2161,Elidolomiti,\N,,EDO,ELIDOLOMITI,Italy,N
2162,Elieuro,\N,,ELB,ELILOBARDIA,Italy,N
2163,Elifriulia,\N,,EFG,ELIFRIULIA,Italy,N
2164,Elilario Italia,\N,,ELH,LARIO,Italy,N
2165,Elilombarda,\N,,EOA,LOMBARDA,Italy,N
2166,Elimediterranea,\N,,MEE,ELIMEDITERRANEA,Italy,N
2167,Elios,\N,,VUL,ELIOS,Italy,N
2168,Elipiu',\N,,IEP,ELIPIU,Italy,N
2169,Elisra Airlines,\N,,RSA,ESRA,Sudan,N
2170,Elite Air,\N,,EAI,ELAIR,Togo,N
2171,Elite Jets,\N,,EJD,ELITE DUBAI,United Arab Emirates,N
2172,Elitellina,\N,,FGS,ELITELLINA,Italy,N
2173,Elliott Aviation,\N,,ELT,ELLIOT,United States,N
2174,Elmagal Aviation Services,\N,,MGG,ELMAGAL,Sudan,N
2175,Elrom Aviation and Investments,\N,,ELR,,Israel,N
2176,Embassy Airlines,\N,,EAM,EMBASSY AIR,Nigeria,N
2177,Embassy Freight Company,\N,,EFT,EMBASSY FREIGHT,United States,N
2178,Empresa Brasileira De Aeronautica,\N,,EMB,EMBRAER,Brazil,N
2179,Embry-Riddle Aeronautical University,\N,,XSL,SATSLAB,United States,N
2180,Emerald Airways,\N,,JEM,GEMSTONE,United Kingdom,N
2181,Emery Worldwide Airlines,\N,,EWW,EMERY,United States,N
2182,Emetebe,\N,,EMT,EMETEBE,Ecuador,N
2183,Emirates,Emirates Airlines,EK,UAE,EMIRATES,United Arab Emirates,Y
2184,Emoyeni Air Charter,\N,,SBC,SABIAN AIR,South Africa,N
2185,Empire Air Service,\N,,EMP,EMPIRE,United States,N
2186,Empire Airlines,\N,EM,CFS,EMPIRE AIR,United States,N
2187,Empire Aviation Services,\N,,MPR,,Nigeria,N
2188,Empire Test Pilots' School,\N,,ETP,TESTER,United Kingdom,N
2189,Empresa (Aero Uruguay),\N,,AUO, S.A.,UNIFORM OSCAR,N
2190,Empresa Aero-Servicios Parrague,\N,,PRG,ASPAR,Chile,N
2191,Empresa Aerocaribbean,\N,,CRN,AEROCARIBBEAN,Cuba,N
2192,Empresa Aviacion Interamericana,\N,,VNA,EBBA,Uruguay,N
2193,Empresa Ecuatoriana De Aviacion,\N,EU,EEA,ECUATORIANA,Ecuador,Y
2194,Empresa Nacional De Servicios Aereos,\N,,CNI,SERAER,Cuba,N
2195,Empresa Venezolana,\N,,VNE,VENEZOLANA,Venezuela,N
2196,Empresa de Aviacion Aerogaviota,\N,,GTV,GAVIOTA,Cuba,N
2197,Empressa Brasileira de Infra-Estrutura Aeroportuaria-Infraero,\N,,XLT,INFRAERO,Brazil,N
2198,Endecots,\N,,ENC,ENDECOTS,Ecuador,N
2199,Enimex,\N,,ENI,ENIMEX,Estonia,N
2200,Enkor JSC,\N,G8,ENK,ENKOR,Russia]],N
2201,Enrique Gleisner Vivanco,\N,,EGV,GLEISNER,Chile,N
2202,Ensenada Vuelos Especiales,\N,,ESE,ENSENADA ESPECIAL,Mexico,N
2203,Entergy Services,\N,,ENS,ENTERGY SHUTTLE,United States,N
2204,Enterprise World Airways,\N,,EWS,WORLD ENTERPRISE,Democratic Republic of the Congo,N
2205,Eos Airlines,\N,E0,ESS,NEW DAWN,United States,N
2206,Equaflight Service,\N,,EKA,EQUAFLIGHT,Republic of the Congo,N
2207,Equatair Air Services (Zambia),\N,,EQZ,ZAMBIA CARGO,Zambia,N
2208,Equatorial Airlines,\N,,EQT,,Equatorial Guinea,N
2209,Era Helicopters,\N,,ERH,ERAH,United States,N
2210,Eram Air,\N,,IRY,ERAM AIR,Iran,N
2211,Erfoto,\N,,ERF,ERFOTO,Portugal,N
2212,Erie Airways,\N,,ERE,AIR ERIE,United States,N
2213,Eritrean Airlines,\N,B8,ERT,ERITREAN,Eritrea,Y
2214,Escuela De Pilotos Are Aviacion,\N,,CTV,ARE AVIACION,Spain,N
2215,Espace Aviation Services,\N,,EPC,ESPACE,Democratic Republic of the Congo,N
2216,Esso Resources Canada,\N,,ERC,ESSO,Canada,N
2217,Estafeta Carga Aerea,\N,E7,ESF,,Mexico,N
2218,Estonian Air,\N,OV,ELL,ESTONIAN,Estonia,Y
2219,Estrellas Del Aire,\N,,ETA,ESTRELLAS,Mexico,N
2220,Ethiopian Airlines,\N,ET,ETH,ETHIOPIAN,Ethiopia,Y
2221,Eti 2000,\N,,MJM,ELCO ETI,Italy,N
2222,Etihad Airways,\N,EY,ETD,ETIHAD,United Arab Emirates,Y
2223,Etram Air Wing,\N,,ETM,ETRAM,Angola,N
2224,Euraviation,\N,,EVN,EURAVIATION,Italy,N
2225,Euro Continental AIE,\N,,ECN,EURO CONTINENTAL,Spain,N
2226,Euro Exec Express,\N,RZ,,,Sweden,Y
2227,Euro Sun,\N,,ESN,EURO SUN,Turkey,N
2228,Euro-Asia Air,\N,,EAK,EAKAZ,Kazakhstan,N
2229,Euro-Asia Air International,\N,,KZE,KAZEUR,Kazakhstan,N
2230,EuroAtlantic Airways,\N,MM,MMZ,EUROATLANTIC,Portugal,N
2231,EuroJet Aviation,\N,,GOJ,GOJET,United Kingdom,N
2232,Euroair,\N,,EUP,EUROSTAR,Greece,N
2233,Euroamerican Air,\N,,EUU,EUROAMERICAN,Uruguay,N
2234,Euroceltic Airways,\N,,ECY,ECHELON,United Kingdom,N
2235,Eurocontrol,\N,,EUC,,Belgium,N
2236,Eurocopter,\N,,ECF,EUROCOPTER,France,N
2237,Eurocypria Airlines,\N,UI,ECA,EUROCYPRIA,Cyprus,Y
2238,Eurofly,\N,,EEZ,E-FLY,Italy,N
2239,Eurofly Service,\N,GJ,EEU,EUROFLY,Italy,Y
2240,Euroguineana de Aviacion,\N,,EUG,EUROGUINEA,Equatorial Guinea,N
2241,Eurojet Italia,\N,,ERJ,JET ITALIA,Italy,N
2242,Eurojet Limited,\N,,JLN,JET LINE,Malta,N
2243,Eurojet Romania,\N,,RDP,JET-ARROW,Romania,N
2244,Eurojet Servis,\N,,EJS,EEJAY SERVICE,Czech Republic,N
2245,Eurolot,\N,K2,ELO,EUROLOT,Poland,Y
2246,Euromanx Airways,\N,3W,EMX,EUROMANX,Austria,N
2247,Europe Air Lines,\N,,GED,LANGUEDOC,France,N
2248,Europe Airpost,\N,,FPO,FRENCH POST,France,N
2249,European 2000 Airlines,\N,,EUT,FIESTA,Malta,N
2250,European Aeronautical Group UK,\N,,EAG,,United Kingdom,N
2251,European Air Express,\N,EA,EAL,STAR WING,Germany,Y
2252,European Air Transport,\N,QY,BCS,EUROTRANS,Belgium,N
2253,European Aviation Air Charter,\N,E7,EAF,EUROCHARTER,United Kingdom,N
2254,European Business Jets,\N,,EBJ,,United Kingdom,N
2255,European Coastal Airlines,\N,,ECB,COASTAL CLIPPER,Croatia,N
2256,European Executive,\N,,ETV,EURO EXEC,United Kingdom,N
2257,European Executive Express,\N,,EXC,ECHO EXPRESS,Sweden,N
2258,Eurosense,\N,,EBG,EUROSENSE,Bulgaria,N
2259,Euroskylink,\N,,ESX,CATFISH,United Kingdom,N
2260,Eurowings,\N,EW,EWG,EUROWINGS,Germany,Y
2261,Evergreen International Airlines,\N,EZ,EIA,EVERGREEN,United States,Y
2262,Everts Air Alaska/Everts Air Cargo,\N,,VTS,EVERTS,United States,N
2263,Examiner Training Agency,\N,,EMN,AGENCY,United Kingdom,N
2264,Excel Airways,\N,JN,XLA,EXPO,United Kingdom,Y
2265,Excel Charter,\N,,XEL,HELI EXCEL,United Kingdom,Y
2266,Excellent Air,\N,,GZA,EXCELLENT AIR,Germany,N
2267,Execair Aviation,\N,MB,EXA,CANADIAN EXECAIRE,Canada,N
2268,Execujet Charter,\N,,VCN,AVCON,Switzerland,N
2269,Execujet Middle East,\N,,EJO,MIDJET,United Arab Emirates,N
2270,Execujet Scandinavia,\N,,VMP,VAMPIRE,Denmark,N
2271,Executive Aerospace,\N,,EAS,AEROSPACE,South Africa,N
2272,Executive Air,\N,,LFL,LIFE FLIGHT,Zimbabwe,N
2273,Executive Air Charter,\N,,EAC,EXECAIR,United States,N
2274,Executive Air Fleet,\N,,XAF,,United States,N
2275,Executive Aircraft Charter and Charter Services,\N,,ECS,ECHO,Ireland,N
2276,Executive Aircraft Services,\N,,XAH,,United Kingdom,N
2277,Executive Airlines,\N,OW,EXK,EXECUTIVE EAGLE,United States,N
2278,Executive Airlines,\N,,EXU,SACAIR,Spain,N
2279,Executive Aviation Services,\N,,JTR,JESTER,United Kingdom,N
2280,Executive Flight,\N,,EXE,EXEC,United States,N
2281,Executive Flight Operations Ontario Government,\N,,TRI,TRILLIUM,Canada,N
2282,Executive Jet Charter,\N,,EXJ,,United Kingdom,N
2283,Executive Jet Management,\N,,EJM,JET SPEED,United States,N
2284,Executive Turbine Aviation,\N,,TEA,TRAVELMAX,South Africa,N
2285,Eximflight,\N,,EXF,EXIMFLIGHT,Mexico,N
2286,Exin,\N,,EXN,EXIN,Poland,N
2287,Expertos En Carga,\N,,EXR,EXPERTOS ENCARGA,Mexico,N
2288,Expo Aviation,\N,8D,EXV,EXPOAVIA,Sri Lanka,N
2289,Express Air,\N,,FXA,EFFEX,United States,N
2290,Express International Cargo,\N,,EIC,EXCARGO,S,N
2291,Express Line Aircompany,\N,,XPL,EXPRESSLINE,United States,N
2292,Express Net Airlines,\N,,XNA,EXPRESSNET,United States,N
2293,Express One International,\N,EO,LHN,LONGHORN,United States,Y
2294,Express Tours,\N,,XTO,EXPRESS TOURS,Canada,N
2295,ExpressJet,\N,XE,BTA,JET LINK,United States,Y
2296,Exxavia Limited,\N,,JTM,SKYMAN,Ireland,N
2297,easyJet,EasyJet Airline,U2,EZY,EASY,United Kingdom,Y
2298,NetJets,\N,,EJA,,United States,N
2299,F.S. Air Service,\N,,EYE,SOCKEYE,United States,N
2300,FAI Airservice,\N,,IFA,RED ANGLE,Germany,N
2301,FINFO Flight Inspection Aircraft,\N,,FLC,FLIGHT CHECK,United States,N
2302,FLM Aviation Mohrdieck,\N,,FKI,KIEL AIR,Germany,N
2303,FLTPLAN,\N,,DCM,DOT COM,United States,N
2304,FLowair Aviation,\N,,FLW,QUICKFLOW,France,N
2305,FMG Verkehrsfliegerschule Flughafen Paderborn-Lippstadt,\N,,FMG,HUSKY,Germany,N
2306,FR Aviation,\N,,FRA,RUSHTON,United Kingdom,N
2307,FSB Flugservice & Development,\N,,FSB,SEABIRD,Germany,N
2308,FSH Luftfahrtunternehmen,\N,,LEJ,LEIPZIG FAIR,Germany,N
2309,Fab Air,\N,,FBA,FAB AIR,Kyrgyzstan,N
2310,Facts Air,\N,,FCS,MEXFACTS,Mexico,N
2311,Fair Aviation,\N,,FAV,FAIRAVIA,South Africa,N
2312,Fair Wind Air Charter,\N,,FWD,FAIR WIND,United Arab Emirates,N
2313,Fairlines,\N,,FLS,FAIRLINE,Netherlands,N
2314,Fairoaks Flight Centre,\N,,FFC,FAIROAKS,United Kingdom,N
2315,Fairways Corporation,\N,,FWY,FAIRWAYS,United States,N
2316,Falcon Air,\N,,FCN,FALCON,Sweden,N
2317,Falcon Air,\N,,FAR,FALCAIR,Slovenia,N
2318,Falcon Air Express,\N,,FAO,PANTHER,United States,N
2319,Falcon Airline,\N,,FAU,FALCON AIRLINE,Nigeria,N
2320,Falcon Aviation,\N,IH,,,Sweden,N
2321,Falcon Aviation Services,\N,,FVS,FALCON AVIATION,United Arab Emirates,N
2322,Falcon Jet Centre,\N,,FJC,FALCONJET,United Kingdom,N
2323,Falwell Aviation,\N,,FAW,FALWELL,United States,N
2324,Far Eastern Air Transport,\N,EF,EFA,Far Eastern,Taiwan,Y
2325,Farmingdale State University,\N,,FDL,FARMINGDALE STATE,United States,N
2326,Farnair Hungary,\N,,FAH,BLUE STRIP,Hungary,N
2327,Farnair Netherlands,\N,,FRN,FARNED,Netherlands,N
2328,Farnair Switzerland,\N,,FAT,FARNER,Switzerland,N
2329,Farnas Aviation Services,\N,,RAF,FARNAS,Sudan,N
2330,Faroecopter,\N,,HBL,HELIBLUE,Denmark,N
2331,Faroejet,\N,F6,RCK,ROCKROSE,Faroe Islands,N
2332,Farwest Airlines,\N,,FRW,FARWEST,United States,N
2333,Faso Airways,\N,F3,FSW,FASO,Burkina Faso,N
2334,Fast Helicopters,\N,,FHL,FINDON,United Kingdom,N
2335,Fayban Air Services,\N,,FAY,FAYBAN AIR,Nigeria,N
2336,Fayetteville Flying Service and Scheduled Skyways System,\N,,SKM,SKYTEM,United States,N
2337,Federal Air,\N,,FDR,FEDAIR,South Africa,N
2338,Federal Airlines,\N,,FLL,FEDERAL AIRLINES,Sweden,N
2339,Federal Armed Forces,\N,,DCN,DIPLOMATIC CLEARANCE,Germany,N
2340,Federal Armored Service,\N,,FRM,FEDARM,United States,N
2341,Federal Aviation Administration,\N,,NHK,NIGHTHAWK,United States,N
2342,Federal Express,\N,FX,FDX,FEDEX,United States,N
2343,Feniks Airline,\N,,FNK,AURIKA,Kazakhstan,N
2344,Feria Aviacion,\N,,FER,FERIA,Spain,N
2345,Fika Salaama Airlines,\N,N8,HGK,SALAAMA,Uganda,N
2346,Finalair Congo,\N,4S,FNC,FINALAIR CONGO,Republic of the Congo,N
2347,Financial Airxpress,\N,,FAK,FACTS,United States,N
2348,Fine Airlines,\N,,FBF,FINE AIR,United States,N
2349,Finist'air,\N,,FTR,FINISTAIR,France,N
2350,Finnair,\N,AY,FIN,FINNAIR,Finland,Y
2351,Finncomm Airlines,\N,FC,WBA,WESTBIRD,Finland,Y
2352,Finnish Air Force,\N,,FNF,FINNFORCE,Finland,N
2353,Firefly,\N,FY,FFM,FIREFLY,Malaysia,Y
2354,First Air,\N,7F,FAB,,Canada,Y
2355,First Air Transport,\N,,JRF,,Japan,N
2356,First Cambodia Airlines,\N,,FCC,FIRST CAMBODIA,Cambodia,N
2357,First Choice Airways,\N,DP,FCA,JETSET,United Kingdom,Y
2358,First City Air,\N,,MBL,FIRST CITY,United Kingdom,N
2359,First Flying Squadron,\N,,GGA,JAWJA,United States,N
2360,First Line Air,\N,,FIR,FIRSTLINE AIR,Sierra Leone,N
2361,First Sabre,\N,,FTS,FIRST SABRE,Mexico,N
2362,Fischer Air,\N,8F,FFR,FISCHER,Czech Republic,N
2363,Fischer Air Polska,\N,,FFP,FLYING FISH,Poland,N
2364,Flagship Express Services,\N,,FSX,FLAG,United States,N
2365,Flair Airlines,\N,,FLE,FLAIR,Canada,N
2366,Flamenco Airways,\N,,WAF,FLAMENCO,United States,N
2367,Flamingo Air,\N,,FMR,FLAMINGO AIR,United States,N
2368,Flamingo Air-Line,\N,,FLN,ILIAS,Kazakhstan,N
2369,Flash Airlines,\N,,FSH,FLASH,Egypt,N
2370,Fleet Requirements Air Direction Unit,\N,,BWY,BROADWAY,United Kingdom,N
2371,Fleetair,\N,,FLR,FLEETAIR,South Africa,N
2372,Flexair,\N,,FXY,FLEXY,Netherlands,N
2373,Flexflight,\N,,FXT,,Denmark,N
2374,Flight Alaska,\N,,TUD,TUNDRA,United States,N
2375,Flight Calibration Service,\N,,FCK,NAV CHECKER,Germany,N
2376,Flight Centre Victoria,\N,,FCV,NAVAIR,Canada,N
2377,Flight Corporation,\N,,FCP,FLIGHTCORP,New Zealand,N
2378,Flight Dispatch Services,\N,,FDP,,Poland,N
2379,Flight Express,\N,,FLX,FLIGHT EXPRESS,United States,N
2380,Flight Inspection Center of the General Administration of Civil Aviation in China,\N,,CFI,CHINA JET,China,N
2381,Flight Inspections and Systems,\N,,LTS,SPECAIR,Russia,N
2382,Flight International,\N,,IVJ,INVADER JACK,United States,N
2383,Flight Line,\N,,MIT,MATCO,United States,N
2384,Flight Ops International,\N,,FOI,,United States,N
2385,Flight Options,\N,,OPT,OPTIONS,United States,N
2386,Flight Precision Limited,\N,,CLB,CALIBRATOR,United Kingdom,N
2387,Flight Safety Limited,\N,,FSL,FLIGHTSAFETY,United Kingdom,N
2388,Flight Support Sweden,\N,,FSU,,Sweden,N
2389,Flight Trac,\N,,CCK,CABLE CHECK,United States,N
2390,Flight Training Europe,\N,,AYR,CYGNET,Spain,N
2391,Flight West Airlines,\N,,FWQ,,Australia,N
2392,Flight-Ops International,\N,,KLO,KLONDIKE,Canada,N
2393,Flightcraft,\N,,CSK,CASCADE,United States,N
2394,Flightexec,\N,,FEX,FLIGHTEXEC,Canada,N
2395,Flightline,\N,B5,FLT,FLIGHTLINE,United Kingdom,Y
2396,Flightpass Limited,\N,,FPS,FLIGHTPASS,United Kingdom,N
2397,Flightstar Corporation,\N,,FSR,FLIGHTSTAR,United States,N
2398,Flightworks,\N,,KDZ,KUDZU,United States,N
2399,Flint Aviation Services,\N,,FAZ,FLINT AIR,United States,N
2400,Florida Air,\N,,OJY,OHJAY,United States,N
2401,Florida Coastal Airlines,\N,PA,FCL,FLORIDA COASTAL,United States,N
2402,Florida Department of Agriculture,\N,,FFS,FORESTRY,United States,N
2403,Florida Jet Service,\N,,FJS,FLORIDAJET,United States,N
2404,Florida West International Airways,\N,RF,FWL,FLO WEST,United States,Y
2405,Flugdienst Fehlhaber,\N,,FFG,WITCHCRAFT,Germany,N
2406,Flugschule Basel,\N,,FLU,YELLOW FLYER,Switzerland,N
2407,Flugschule Eichenberger,\N,,EZB,EICHENBURGER,Switzerland,N
2408,Flugwerkzeuge Aviation Software,\N,,FWZ,,Austria,N
2409,Fly Air,\N,F2,FLM,FLY WORLD,Turkey,N
2410,Fly CI Limited,\N,,FCT,DEALER,United Kingdom,N
2411,Fly Europa Limited,\N,,FEE,FLY EURO,United Kingdom,N
2412,Fly Excellent,\N,,FXL,FLY EXCELLENT,Sweden,N
2413,Fly International Airways,\N,,NVJ,NOUVINTER,Tunisia,N
2414,Fly Line,\N,,FIL,FLYLINE,Spain,N
2415,Fly Me Sweden,\N,SH,FLY,FLYBIRD,Sweden,N
2416,Fly Wex,\N,,IAD,FLYWEX,Italy,N
2417,AirAsia X,FlyAsianXpress,D7,XAX,XANADU,Malaysia,Y
2418,FlyLal,\N,TE,LIL,LITHUANIA AIR,Lithuania,Y
2419,FlyNordic,\N,LF,NDC,NORDIC,Sweden,Y
2420,Flybaboo,\N,F7,BBO,BABOO,Switzerland,Y
2421,Flybe,\N,BE,BEE,JERSEY,United Kingdom,Y
2422,Flycolumbia,\N,,FCE,FLYCOLUMBIA,Spain,N
2423,Flycom,\N,,FLO,FLYCOM,Slovenia,N
2424,Flygaktiebolaget Gota Vingar,\N,,GVG,BLUECRAFT,Sweden,N
2425,Flyglobespan,\N,B4,GSM,GLOBESPAN,United Kingdom,Y
2426,Flygprestanda,\N,,FPA,,Sweden,N
2427,Flygtransporter I Nykoping,\N,,ETS,EXTRANS,Sweden,N
2428,Flyguppdraget Backamo,\N,,INU,INSTRUCTOR,Sweden,N
2429,Flyhy Cargo Airlines,\N,,FYH,FLY HIGH,Thailand,Y
2430,Flying Carpet Company,\N,,FCR,FLYING CARPET,Lebanon,N
2431,Flying Service,\N,,FYG,FLYING GROUP,Belgium,N
2432,Flying-Research Aerogeophysical Center,\N,,FGP,FLYING CENTER,Russia,N
2433,Flylink Express,\N,,FLK,FLYLINK,Spain,N
2434,Flyteam Aviation,\N,,FTM,FLYTEAM,United Kingdom,N
2435,Focus Air,\N,,FKS,FOCUS,United States,N
2436,Fokker,\N,,FOP,,Netherlands,N
2437,Fonnafly,\N,,NOF,FONNA,Norway,N
2438,Ford Motor Company,\N,,FOB,FORDAIR,United Kingdom,N
2439,Formosa Airlines,\N,VY,FOS,,Taiwan,Y
2440,Formula One Management,\N,,FOR,FORMULA,United Kingdom,N
2441,Forth and Clyde Helicopter Services,\N,,FHS,HELISCOT,United Kingdom,N
2442,Fortunair Canada,\N,,FXC,AIR FUTURE,Canada,N
2443,Forward Air International Airlines,\N,BN,,,United States,N
2444,Foster Aviation,\N,,FSA,FOSTER-AIR,United States,N
2445,Foster Yeoman,\N,,JFY,YEOMAN,United Kingdom,N
2446,Fotografia F3,\N,,FTE,FOTOGRAFIA,Spain,N
2447,Four Island Air,\N,,FIA,ISLANDAIR,Antigua and Barbuda,N
2448,Four Star Aviation / Four Star Cargo,\N,HK,FSC,FOUR STAR,United States,N
2449,Four Winds Aviation,\N,,WDS,WINDS,United States,N
2450,Foxair,\N,,FXR,WILDFOX,Italy,N
2451,France Douanes,\N,,FDO,FRENCH CUSTOM,France,N
2452,Free Bird Airlines,\N,,FHY,FREE BIRD,Turkey,N
2453,Freedom Air,\N,,FOM,FREE AIR,New Zealand,N
2454,Freedom Air,\N,FP,FRE,FREEDOM,United States,Y
2455,Freedom Air Services,\N,,FFF,INTER FREEDOM,Nigeria,N
2456,Freedom Airlines,\N,,FRL,FREEDOM AIR,United States,Y
2457,Freedom Airways,\N,,FAS,FREEDOM AIRWAYS,Cyprus,N
2458,Freeway Air,\N,,FWC,FREEWAY,Netherlands,N
2459,Freight Runners Express,\N,,FRG,FREIGHT RUNNERS,United States,N
2460,Force Aerienne Francaise,\N,,FAF,FRENCH AIR FORCE,France,N
2461,Aviation Legere De L'Armee De Terre,\N,,FMY,FRENCH ARMY,France,N
2462,France Marine Nationale,\N,,FNY,FRENCH NAVY,France,N
2463,Fresh Air,\N,,FRR,FRESH AIR,Nigeria,N
2464,Freshaer,\N,,FAE,WILDGOOSE,United Kingdom,N
2465,Friendship Air Alaska,\N,,FAL,FRIENDSHIP,United States,N
2466,Friendship Airlines,\N,,FLF,FRIEND AIR,Uganda,N
2467,Froggy Corporate Aviation,\N,,FGY,,Australia,N
2468,Frontier Airlines,\N,F9,FFT,FRONTIER FLIGHT,United States,Y
2469,Frontier Commuter,\N,,ITR,OUT BACK,United States,N
2470,Frontier Flying Service,\N,2F,FTA,FRONTIER-AIR,United States,Y
2471,Frontier Guard,\N,,FNG,FINNGUARD,Finland,N
2472,Fujairah Aviation Centre,\N,,FUJ,FUJAIRAH,United Arab Emirates,N
2473,Fujian Airlines,\N,,CFJ,FUJIAN,China,N
2474,Full Express,\N,,GAX,GRAND AIRE,United States,N
2475,Fumigacion Aerea Andaluza,\N,,FAM,FAASA,Spain,N
2476,Fun Flying Thai Air Service,\N,,FFY,FUN FLYING,Thailand,N
2477,Fundació Rego,\N,,ROG,REGO,Spain,N
2478,Funtshi Aviation Service,\N,,FUN,FUNTSHI,Democratic Republic of the Congo,N
2479,Futura Gael,\N,,FGL,Applewood,Ireland,N
2480,Futura International Airways,\N,FH,FUA,FUTURA,Spain,N
2481,G & L Aviation,\N,,GML,GEEANDEL,South Africa,N
2482,G5 Executive,\N,,EXH,BATMAN,Germany,N
2483,GAK/Mitchell Aero,\N,,MTA,GAK AVIATION,United States,N
2484,GATSA,\N,,GGS,GATSA,Mexico,N
2485,GB Airlink,\N,,GBX,ISLAND TIGER,United States,N
2486,GB Airways,\N,GT,GBL,GEEBEE AIRWAYS,United Kingdom,Y
2487,GCS Air Service,\N,,GCS,GALION,United States,N
2488,GEC Marconi Avionics,\N,,FFU,FERRANTI,United Kingdom,N
2489,GECAS,\N,,GCC,GECAS,Ireland,N
2490,GENSA,\N,,GEN,GENSA-BRASIL,Brazil,N
2491,GETRA,\N,,GET,GETRA,Equatorial Guinea,N
2492,GFW Aviation,\N,,GFW,,Australia,N
2493,GH Stansted Limited,\N,,GHI,,United Kingdom,N
2494,GM Helicopters,\N,,GMG,GEE-EM HELICOPTERS,Latvia,N
2495,GP Express Airlines,\N,,GPE,REGIONAL EXPRESS,United States,N
2496,GPM Aeroservicio,\N,,GPR,GPM AEROSERVICIO,Mexico,N
2497,GR-Avia,\N,,GIB,GRAVIA,Guinea,N
2498,GST Aero Aircompany,\N,,BMK,MURAT,Kazakhstan,N
2499,GTA Air,\N,,GTX,BIG-DEE,United States,N
2500,Ga-Ma Helicoptere,\N,,GAH,GAMHELICO,France,N
2501,Gabon Express,\N,,GBE,GABEX,Gabon,N
2502,Gabon-Air-Transport,\N,,GRT,,Gabon,N
2503,Gacela Air Cargo,\N,,GIG,GACELA AIR,Mexico,N
2504,Gail Force Express,\N,,GFC,GAIL FORCE,United States,N
2505,Gain Jet Aviation,\N,,GNJ,HERCULES JET,Greece,N
2506,Galair International,\N,,SWF,GALAIR,United Kingdom,N
2507,Galaircervis,\N,,GLS,GALS,Ukraine,N
2508,Galaxy Air,\N,7O,GAL,GALAXY,Kyrgyzstan,N
2509,Galaxy Airlines,\N,,GXY,GALAX,Japan,N
2510,Galena Air Service,\N,,GAS,GALENA AIR SERVICE,United States,N
2511,Galileo International,\N,1G,,,United States,N
2512,Gama Aviation,\N,,GMA,GAMA,United Kingdom,N
2513,Gambia International Airlines,\N,GC,GNR,GAMBIA INTERNATIONAL,Gambia,N
2514,Gambia New Millennium Air,\N,,NML,NEWMILL,Gambia,N
2515,Gamisa Aviacion,\N,,GMJ,GAMISA,Spain,N
2516,Gandalf Airlines,\N,G7,GNF,Gandalf,Italy,N
2517,Gander Aviation,\N,,GAN,GANAIR,Canada,N
2518,Garden State Airlines,\N,,GSA,GARDEN STATE,United States,N
2519,Garrison Aviation,\N,,AHM,AIR HURON,Canada,N
2520,Garuda Indonesia,\N,GA,GIA,INDONESIA,Indonesia,Y
2521,Gatari Hutama Air Services,\N,,GHS,GATARI,Indonesia,N
2522,Gauteng Air Cargo,\N,,EGO,GAUTENG,South Africa,N
2523,Gavina,\N,,GVN,GAVINA,Spain,N
2524,Gazpromavia,\N,4G,GZP,GAZPROMAVIA,Russia,Y
2525,Geesair,\N,,GEE,GEESAIR,Canada,N
2526,Gemini Air Cargo,\N,GR,GCO,GEMINI,United States,N
2527,Gendall Air,\N,,GAB,GENDALL,Canada,N
2528,Gendarmerie Belge,\N,,GDB,BELGIAN GENERMERIE,Belgium,N
2529,Gendarmie Nationale,\N,,FGN,FRANCE GENDARME,France,N
2530,General Aerospace,\N,,SWK,SKYWALKER,Canada,N
2531,General Airways,\N,,GWS,GENAIR,South Africa,N
2532,General Aviation,\N,,GNZ,GONZO,Poland,N
2533,General Aviation Flying Services,\N,,GTH,GOTHAM,United States,N
2534,General Aviation Terminal,\N,,XGA,,Canada,N
2535,General Motors,\N,,GMC,GENERAL MOTORS,United States,N
2536,Genex,\N,,GNX,,Belarus,N
2537,Geographic Air Surveys,\N,,GSL,SURVEY-CANADA,Canada,N
2538,Georgian Airways,\N,A9,TGZ,TAMAZI,Georgia,Y
2539,Georgian Aviation Federation,\N,,FGA,GEORGIA FED,Georgia,N
2540,Georgian Cargo Airlines Africa,\N,,GGF,GEORGIAN AFRICA,Senegal,N
2541,Georgian National Airlines,\N,QB,GFG,NATIONAL,Georgia,Y
2542,DLR,\N,,GPL,GERMAN POLAR,Germany,N
2543,German Air Force,\N,,GAF,GERMAN AIR FORCE,Germany,N
2544,German Army,\N,,GAM,GERMAN ARMY,Germany,N
2545,German Navy,\N,,GNY,GERMAN NAVY,Germany,N
2546,German Sky Airlines,\N,,GHY,GERMAN SKY,Germany,N
2547,Germania,\N,ST,GMI,GERMANIA,Germany,Y
2548,Germanwings,\N,4U,GWI,GERMAN WINGS,Germany,Y
2549,Gerry's Dnata,\N,,GDN,,Pakistan,N
2550,Gesekkschaft Fur Flugzieldarstellung,\N,,GFD,KITE,Germany,N
2551,Gestair,\N,GP,GES,GESTAIR,Spain,N
2552,Gestar,\N,,GTR,STAR GESTAR,Chile,N
2553,Gestion Aerea Ajecutiva,\N,,GJT,BANJET,Spain,N
2554,Ghadames Air Transport,\N,,GHT,,Libya,N
2555,Ghana Airways,\N,,GHA,GHANA,Ghana,N
2556,Ghana International Airlines,\N,G0,GHB,GHANA AIRLINES,Ghana,Y
2557,Gibson Aviation,\N,,NTC,NIGHT CHASE,United States,N
2558,Global Air Charter,\N,,RPS,RESPONSE,United States,N
2559,Global Air Operations,\N,,GAG,GLOBAL AIRGROUP,Comoros,N
2560,Global Air Services Nigeria,\N,,GBS,GLOBAL SERVE,Nigeria,N
2561,Global Aircargo,\N,,GLC,,Bahrain,N
2562,Global Airways,\N,,BSP,,Democratic Republic of Congo,N
2563,Global Aviation Operations,\N,,GBB,GLOBE,South Africa,N
2564,Global Aviation and Services Group,\N,,GAK,AVIAGROUP,Libya,N
2565,Global Georgian Airways,\N,,GGZ,GLOBAL GEORGIAN,Georgia,N
2566,Global Jet Austria,\N,,GLJ,GLOBAL AUSTRIA,Austria,N
2567,Global Jet Corporation,\N,,NSM,THUNDERCLOUD,United States,N
2568,Global Jet Luxembourg,\N,,SVW,SILVER ARROWS,Luxembourg,N
2569,Global Sky Aircharter,\N,,GSK,GLOBAL SKY,United States,N
2570,Global Supply Systems,\N,,GSS,JET LIFT,United Kingdom,N
2571,Global System,\N,,XGS,,United States,N
2572,Global Weather Dynamics,\N,,XGW,,United States,N
2573,Global Wings,\N,,GLW,,Japan,N
2574,Globe Jet,\N,,GJA,,Lebanon,N
2575,Go Air,\N,G8,GOW,GOAIR,India,Y
2576,Go One Airways,\N,GK,,,United Kingdom,N
2577,GoJet Airlines,\N,G7,GJS,GATEWAY,United States,Y
2578,Gobierno De Guinea Ecuatorial,\N,,GGE,,Equatorial Guinea,N
2579,Gof-Air,\N,,GOF,GOF-AIR,Mexico,N
2580,Gofir,\N,,GOI,SWISS HAWK,Switzerland,N
2581,Gol Transportes Aéreos,\N,G3,GLO,GOL TRANSPORTE,Brazil,Y
2582,Gold Belt Air Transport,\N,,GBT,GOLD BELT,Canada,N
2583,GoldAir,\N,,GDA,AIR PARTNER,United Kingdom,N
2584,Goldeck-Flug,\N,,GDK,GOLDECK FLUG,Austria,N
2585,Golden Air,\N,DC,GAO,GOLDEN,Sweden,Y
2586,Golden Airlines,\N,,GDD,GOLDEN AIRLINES,United States,N
2587,Golden Pacific Airlines,\N,,GPA,GOLDEN PAC,United States,N
2588,Golden Rule Airlines,\N,,GRS,GOLDEN RULE,Kyrgyzstan,N
2589,Golden Star Air Cargo,\N,,GLD,GOLDEN STAR,Sudan,N
2590,Goldfields Air Services,\N,,GOS,,Australia,N
2591,Golfe Air Quebec,\N,,GAQ,GOLFAIR,Canada,N
2592,Goliaf Air,\N,,GLE,GOLIAF AIR,Sao Tome and Principe,N
2593,Gomel Airlines,\N,,GOM,GOMEL,Belarus,N
2594,Goodridge (UK) Limited,\N,,RDR,RED STAR,United Kingdom,N
2595,Gorkha Airlines,\N,G1,IKA,GORKHA AIRLINES,Nepal,N
2596,Gorlitsa Airlines,\N,,GOR,GORLITSA,Ukraine,N
2597,Government Flying Service,\N,,HKG,HONGKONG GOVERNMENT,Hong Kong SAR of China,N
2598,Government of Zambia Communications Flight,\N,,GRZ,COM FLIGHT,Zambia,N
2599,Grampian Flight Centre,\N,,HLD,GRANITE,United Kingdom,N
2600,Granada Aviacion,\N,,GAV,GRANAVI,Spain,N
2601,Grand Aire Express,\N,,GAE,GRAND EXPRESS,United States,N
2602,Grand Airways,\N,,GND,GRAND VEGAS,United States,N
2603,Grand Canyon Airlines,\N,,CVU,CANYON VIEW,United States,N
2604,Grant Aviation,\N,GS,GUN,HOOT,United States,N
2605,Grantex Aviation,\N,,LMK,LANDMARK,United Kingdom,N
2606,Great American Airways,\N,,GRA,GREAT AMERICAN,United States,N
2607,Great Lakes Airlines,\N,ZK,GLA,LAKES AIR,United States,Y
2608,Great Lakes Airways (Uganda),\N,,GLU,LAKES CARGO,Uganda,N
2609,Great Plains Airlines,\N,,GRP,GREAT PLAINS,United States,N
2610,Great Wall Airlines,\N,IJ,GWL,GREAT WALL,China,N
2611,Great Western Air,\N,,GWA,G-W AIR,United States,N
2612,Hellenic Air Force,\N,,HAF,HELLENIC AIR FORCE,Greece,N
2613,Greek Navy,\N,,HNA,HELLENIC NAVY,Greece,N
2614,Griffin Aviation,\N,,GFF,GRIFFIN AIR,Cyprus,N
2615,Grixona,\N,,GXA,GRIXONA,Moldova,N
2616,Grizodubova Air Company,\N,,GZD,GRIZODUBOVA AIR,Russia,N
2617,Grossmann Air Service,\N,,HTG,GROSSMANN,Austria,N
2618,Grossmann Jet Service,\N,,GSJ,GROSSJET,Czech Republic,N
2619,Ground Handling Service de Mexico,\N,,GHV,GROUND HANDLING,Mexico,N
2620,Grup Air-Med,\N,,GPM,GRUPOMED,Spain,N
2621,Grupo De Aviacion Ejecutiva,\N,,EJC,GRUPOEJECUTIVA,Mexico,N
2622,Grupo TACA,TACA,TA,TAT,TACA-COSTARICA,Costa Rica,Y
2623,Grupo Vuelos Mediterraneo,\N,,VMM,VUELOS MED,Spain,N
2624,Grupoaereo Monterrey,\N,,GMT,GRUPOMONTERREY,Mexico,N
2625,Guard Systems,\N,,GSY,GUARD AIR,Norway,N
2626,Guine Bissaur Airlines,\N,G6,BSR,BISSAU AIRLINES,Guinea-Bissau,N
2627,Guinea Airways,\N,,GIJ,GUINEA AIRWAYS,Guinea,N
2628,Guinea Cargo,\N,,GNC,GUINEA CARGO,Equatorial Guinea,N
2629,Guinee Airlines,\N,J9,GIF,GUINEE AIRLINES,Guinea,N
2630,Guinea Ecuatorial Airlines,\N,,GEA,GEASA,Equatorial Guinea,N
2631,Guinee Paramount Airlines,\N,,GIQ,GUIPAR,Guinea,N
2632,Guizhou Airlines,\N,,CGH,GUIZHOU,China,N
2633,Guja,\N,,GUS,GUJA,Mexico,N
2634,Gujarat Airways,\N,G8,GUJ,GUJARATAIR,India,N
2635,Gulf & Caribbean Cargo / Contract Air Cargo,\N,,TSU,TRANSAUTO,United States,N
2636,Gulf African Airlines - Gambia,\N,,GUF,GULF AFRICAN,,N
2637,Gulf Air,\N,,GFA,GULF AIR,Oman,Y
2638,Gulf Air Bahrain,\N,GF,GBA,GULF BAHRAIN,Bahrain,Y
2639,Gulf Air Inc,\N,,GAT,GULF TRANS,United States,N
2640,Gulf Central Airlines,\N,,GCN,GULF CENTRAL,United States,N
2641,Gulf Flite Center,\N,,SFY,SKY FLITE,United States,N
2642,Gulf Pearl Air Lines,\N,,GPC,AIR GULFPEARL,Libya,N
2643,Gulfstream Aerospace,\N,,GLF,GULFSTREAM TEST,United States,N
2644,Gulfstream Airlines,\N,,GFS,GULFSTAR,United States,N
2645,Gulfstream International Airlines,\N,,GFT,GULF FLIGHT,United States,Y
2646,Gull Air,\N,,GUL,GULL-AIR,United States,N
2647,Gum Air,\N,,GUM,GUM AIR,Suriname,N
2648,Guneydogu Havacilik Isletmesi,\N,,GDH,RISING SUN,Turkey,N
2649,Guyana Airways 2000,\N,GY,,,,N
2650,Gwyn Aviation,\N,,GWN,GWYN,United Kingdom,N
2651,HC Airlines,\N,,HLA,HEAVYLIFT,United Kingdom,N
2652,HPM Investments,\N,,HWD,FLITEWISE,United Kingdom,N
2653,HT Helikoptertransport,\N,,KTR,COPTER TRANS,Sweden,N
2654,HTA Helicopteros,\N,,AHT,HELIAPRA,Portugal,N
2656,Hadison Aviation,\N,,FMS,HADI,United States,N
2657,Hageland Aviation Services,\N,H6,HAG,HAGELAND,United States,Y
2658,Hagondale Limited,\N,,POW,AIRNET,United Kingdom,N
2659,Hahn Air,\N,HR,HHN,ROOSTER,Germany,N
2660,Hainan Airlines,\N,HU,CHH,HAINAN,China,Y
2661,Hainan Phoenix Information Systems,\N,1R,,,China,N
2662,Haiti Air Freight,\N,,HLS,,Haiti,N
2663,Haiti Ambassador Airlines,\N,2T,HAM,,Haiti,Y
2664,Haiti International Air,\N,,HTI,HAITI INTERNATIONAL,Haiti,N
2665,Haiti International Airline,\N,,HRB,HAITI AIRLINE,Haiti,N
2666,Haiti National Airlines (HANA),\N,,HNR,HANAIR,Haiti,N
2667,Haiti North Airline,\N,,HTN,,Haiti,N
2668,Haiti Trans Air,\N,,HTC,HAITI TRANSAIR,Haiti,N
2669,Haitian Aviation Line,\N,,HBC,HALISA,Haiti,N
2670,Hajvairy Airlines,\N,,HAJ,HAJVAIRY,Pakistan,N
2671,Hak Air,\N,,HKL,HAK AIRLINE,Nigeria,N
2672,Hala Air,\N,,HLH,HALA AIR,Sudan,N
2673,Halcyon Air/Bissau Airways,\N,,HCN,HALCYON,Guinea-Bissau,N
2674,Hamburg International,\N,4R,HHI,HAMBURG JET,Germany,Y
2675,Hamlin Jet,\N,,HJL,BIZJET,United Kingdom,N
2676,Hamra Air,\N,,HMM,HAMRA,United Arab Emirates,N
2677,Hand D Aviation,\N,,WVA,WABASH VALLEY,United States,N
2678,Hangar 8,\N,,HGR,HANG,United Kingdom,N
2679,Hangard Aviation,\N,,HGD,HANGARD,Mongolia,N
2680,Hansung Airlines,\N,,HAN,HANSUNG AIR,Republic of Korea,N
2681,TUIfly,\N,X3,HLX,YELLOW CAB,Germany,Y
2682,Hapagfly,\N,HF,HLF,HAPAG LLOYD,Germany,Y
2683,Harbor Airlines,\N,HB,HAR,HARBOR,United States,N
2684,Harmony Airways,\N,HQ,HMY,HARMONY,Canada,N
2685,Haughey Air,\N,,NBR,NORBROOK,United Kingdom,N
2686,Haverfordwest Air Charter Services,\N,,PYN,POYSTON,United Kingdom,N
2687,Havilah Air Services,\N,,HAV,HAVILAH,Nigeria,N
2688,Hawaiian Airlines,\N,HA,HAL,HAWAIIAN,United States,Y
2689,Hawaiian Pacific Airlines,\N,HP,,,United States,N
2690,Hawk Air,\N,,HKR,AIR HAW,Argentina,N
2691,Hawk De Mexico,\N,,HMX,HAWK MEXICO,Mexico,N
2692,Hawkair,\N,BH,,,Canada,Y
2693,Hawkaire,\N,,HKI,HAWKEYE,United States,N
2694,Hazelton Airlines,\N,,HZL,HAZELTON,Australia,N
2695,Heavylift Cargo Airlines,\N,HN,HVY,HEAVY CARGO,Australia,N
2696,Heavylift International,\N,,HVL,HEAVYLIFT INTERNATIONAL,United Arab Emirates,N
2697,Helcopteros De Cataluna,\N,,HDC,HELICATALUNA,Spain,N
2698,Helenair (Barbados),\N,,HCB,HELEN,Barbados,N
2699,Helenair Corporation,\N,,HCL,HELENCORP,Saint Lucia,N
2700,Helenia Helicopter Service,\N,,HHP,HELENIA,Denmark,N
2701,Heli Air Services,\N,,HLR,HELI BULGARIA,Bulgaria,N
2702,Heli Ambulance Team,\N,,ALJ,ALPIN HELI,Austria,N
2703,Heli Bernina,\N,,HEB,HELIBERNINA,Switzerland,N
2704,Heli France,\N,8H,HFR,HELIFRANCE,France,Y
2705,Heli Hungary,\N,,HYH,HELIHUNGARY,Hungary,N
2706,Heli Medwest De Mexico,\N,,HLM,HELIMIDWEST,Mexico,N
2707,Heli Securite,\N,,HLI,HELI SAINT-TROPEZ,France,N
2708,Heli Trip,\N,,HTP,HELI TRIP,Mexico,N
2709,Heli Union Heli Prestations,\N,,HLU,HELI UNION,France,N
2710,Heli-Air-Monaco,\N,,MCM,HELI AIR,Monaco,N
2711,Heli-Holland,\N,,HHE,HELI HOLLAND,Netherlands,N
2712,Heli-Iberica,\N,,HRA,ERICA,Spain,N
2713,Heli-Iberica Fotogrametria,\N,,HIF,HIFSA,Spain,N
2714,Heli-Inter Guyane,\N,,HIG,INTER GUYANNE,France,N
2715,Heli-Link,\N,,HLK,HELI-LINK,Switzerland,N
2716,Heliamerica De Mexico,\N,,HMC,HELIAMERICA,Mexico,N
2717,Heliavia-Transporte Aereo,\N,,HEA,HELIAVIA,Portugal,N
2718,Heliaviation Limited,\N,,CDY,CADDY,United Kingdom,N
2719,Helibravo Aviacao,\N,,HIB,HELIBRAVO,Portugal,N
2720,Helicap,\N,,HLC,HELICAP,France,N
2721,Helicentre Coventry,\N,,COV,HELICENTRE,United Kingdom,N
2722,Helicol,\N,,HEL,HELICOL,Colombia,N
2723,Helicopter,\N,,HCP,HELI CZECH,Czech Republic,N
2724,Helicopter & Aviation Services,\N,,JKY,JOCKEY,United Kingdom,N
2725,Helicopter Training & Hire,\N,,MVK,MAVRIK,United Kingdom,N
2726,Helicopteros Aero Personal,\N,,HAP,HELIPERSONAL,Mexico,N
2727,Helicopteros Agroforestal,\N,,HAA,AGROFORESTAL,Chile,N
2728,Helicopteros Internacionales,\N,,HNT,HELICOP INTER,Mexico,N
2729,Helicopteros Y Vehiculos Nacionales Aereos,\N,,HEN,HELINAC,Mexico,N
2730,Helicsa,\N,,HHH,HELICSA,Spain,N
2731,Helijet,\N,JB,JBA,HELIJET,Canada,Y
2732,Helikopterdrift,\N,,HDR,HELIDRIFT,Norway,N
2733,Helikopterservice Euro Air,\N,,SCO,SWEDCOPTER,Sweden,N
2734,Heliocean,\N,,OCE,HELIOCEAN,France,N
2735,Helios Airways,\N,ZU,HCY,HELIOS,Cyprus,N
2736,Helipistas,\N,,HLP,HELIPISTAS,Spain,N
2737,Heliportugal,\N,,HPL,HELIPORTUGAL,Portugal,N
2738,Heliservicio Campeche,\N,,HEC,HELICAMPECHE,Mexico,N
2739,Helisul,\N,,HSU,HELIS,Portugal,N
2740,Heliswiss,\N,,HSI,HELISWISS,Switzerland,N
2741,Helitafe,\N,,HLT,HELITAFE,Mexico,N
2742,Helitalia,\N,,HIT,HELITALIA,Italy,N
2743,Helitaxi Ofavi,\N,,OFA,OFAVI,Mexico,N
2744,Helitrans,\N,,HTA,SCANBIRD,Norway,N
2745,Helitrans Air Service,\N,,HTS,HELITRANS,United States,N
2746,Heliworks,\N,,HLW,HELIWORKS,Chile,N
2747,Hellas Jet,\N,T4,HEJ,HELLAS JET,Greece,Y
2748,Hello,\N,HW,FHE,FLYHELLO,Switzerland,Y
2749,Helog,\N,,HLG,HELOG,Switzerland,N
2750,Helvetic Airways,\N,2L,OAW,HELVETIC,Switzerland,Y
2751,Hemus Air,\N,DU,HMS,HEMUS AIR,Bulgaria,N
2752,Henebury Aviation,\N,,HAC,,Australia,N
2753,Heritiage Flight (Valley Air Services),\N,,SSH,SNOWSHOE,United States,N
2754,Herman's Markair Express,\N,,MRX,SPEEDMARK,United States,N
2755,Herritage Aviation Developments,\N,,HED,FLAPJACK,United Kingdom,N
2756,Hewa Bora Airways,\N,EO,ALX,ALLCONGO,Democratic Republic of the Congo,N
2757,Hex'Air,\N,UD,HER,HEX AIRLINE,France,Y
2758,Hi-Jet Helicopter Services,\N,,HHS,HIJET,Suriname,N
2759,Hi Fly,\N,5K,HFY,SKY FLYER,Portugal,N
2760,High-Line Airways,\N,,HLB,HIGH-LINE,Canada,N
2761,Highland Airways,\N,,HWY,HIWAY,United Kingdom,Y
2762,Hispaniola Airways,\N,,HIS,HISPANIOLA,Dominican Republic,N
2763,Hogan Air,\N,,HGA,HOGAN AIR,United States,N
2764,Hokkaido Air System,\N,,NTH,NORTH AIR,Japan,N
2765,Hokkaido International Airlines,\N,HD,ADO,AIR DO,Japan,Y
2766,Hokuriki-Koukuu Company,\N,,ABH,,Japan,N
2767,Hola Airlines,\N,H5,HOA,HOLA,Spain,N
2768,Holding International Group,\N,,HIN,HOLDING GROUP,Mexico,N
2769,Holiday Airlines,\N,,HOL,HOLIDAY,United States,N
2770,Holstenair Lubeck,\N,,HTR,HOLSTEN,Germany,N
2771,Homac Aviation,\N,,HMV,HOMAC,Luxembourg,N
2772,Honduras Airlines,\N,,HAS,HONDURAS AIR,Honduras,N
2773,Hong Kong Airlines,\N,HX,CRK,BAUHINIA,Hong Kong SAR of China,Y
2774,Hong Kong Express Airways,\N,UO,HKE,HONGKONG SHUTTLE,Hong Kong SAR of China,Y
2775,Honiara Cargo Express,\N,,HEX,HONIARA CARGO,Solomon Islands,N
2776,Hop-A-Jet,\N,,HPJ,HOPA-JET,United States,N
2777,Hope Air,\N,HH,,HOPE AIR,Canada,N
2778,Horizon Air,Horizon Airlines,QX,QXE,HORIZON AIR,United States,Y
2779,Horizon Air Service,\N,,KOK,KOKO,United States,N
2780,Horizon Air for Transport and Training,\N,,HSM,ALOFUKAIR,Libya,N
2781,Horizon Air-Taxi,\N,,HOR,HORIZON,Switzerland,N
2782,Horizon Airlines,\N,BN,HZA,,Australia,Y
2783,Horizon Plus,\N,,HPS,HORIZON PLUS,Bangladesh,N
2784,Horizons Unlimited,\N,,HUD,HUD,United States,N
2785,Horizontes Aereos,\N,,HOZ,HORIZONTES AEREOS,Mexico,N
2786,Hoteles Dinamicos,\N,,HDI,DINAMICOS,Mexico,N
2787,Houston Helicopters,\N,,HHO,HOUSTON HELI,United States,N
2788,Houston Jet Services,\N,,GGV,GREGG AIR,Austria,N
2789,Hozu-Avia,\N,,OZU,HOZAVIA,Kazakhstan,N
2790,Hub Airlines,\N,,HUB,HUB,United States,N
2791,Huessler Air Service,\N,,HUS,HUESSLER,United States,N
2792,Hughes Aircraft Company,\N,,GMH,HUGHES EXPRESS,United States,N
2793,Hummingbird Helicopter Service,\N,,WHR,WHIRLEYBIRD,United States,N
2794,Hunair Hungarian Airlines,\N,,HUV,SILVER EAGLE,Hungary,N
2795,Hungarian Air Force,\N,,HUF,HUNGARIAN AIRFORCE,Hungary,N
2796,Hyack Air,\N,,HYA,HYACK,Canada,N
2797,Hydro Air Flight Operations,\N,,HYC,HYDRO CARGO,South Africa,N
2798,Hydro-Québec,\N,,HYD,HYDRO,Canada,N
2799,Héli Sécurité Helicopter Airlines,\N,H4,,,France,N
2800,International Air Carrier Association,\N,,ITC,,Belgium,N
2801,IBC Airways,\N,II,CSQ,CHASQUI,United States,N
2802,IBL Aviation,\N,,IBL,CATOVAIR,Mauritius,N
2803,IBM Euroflight Operations,\N,,BBL,BLUE,Switzerland,N
2804,ICAO,\N,,YYY,,,N
2805,ICAR Airlines,\N,C3,,,Ukraine,N
2806,ICC Canada,\N,,CIC,AIR TRADER,Canada,N
2807,IDG Technology Air,\N,,IDG,INDIGO,Czech Republic,N
2808,IFL Group,\N,,IFL,EIFEL,United States,N
2809,II Lione Alato Arl,\N,,RDE,FLIGHT RED,United Kingdom,N
2810,IJM International Jet Management,\N,,IJM,JET MANAGEMENT,Austria,N
2811,IKI International Airlines,\N,,IKK,IKIAIR,Japan,N
2812,IKON FTO,\N,,IKN,IKON,Germany,N
2813,IMP Aviation Services,\N,,BLU,BLUENOSE,Canada,N
2814,IMP Group Aviation Services,\N,,XGG,,Canada,N
2815,INFINI Travel Information,\N,1F,,,Japan,N
2816,IPEC Aviation,\N,,IPA,IPEC,Australia,N
2817,IPM Europe,\N,,IPM,SHIPEX,United Kingdom,N
2818,IRS Airlines,\N,,LVB,SILVERBIRD,Nigeria,N
2819,ISD Avia,\N,,ISD,ISDAVIA,Ukraine,N
2820,ITA Software,\N,1U,,,United States,N
2821,IVV Femida,\N,,FDF,,Russia,N
2822,Iberia Airlines,\N,IB,IBE,IBERIA,Spain,Y
2823,Ibertour Servicios Aereos,\N,,IBR,IBERTOUR,Spain,N
2824,Ibertrans Aerea,\N,,IBT,IBERTRANS,Spain,N
2825,Iberworld,\N,TY,IWD,,Spain,Y
2826,Ibex Airlines,\N,FW,IBX,IBEX,Japan,Y
2827,Ibicenca Air,\N,,IBC,IBICENCA,Spain,N
2828,Ibk-Petra,\N,,AKI, ,Sudan,N
2829,Icar Air,\N,,RAC,TUZLA AIR,Bosnia and Herzegovina,Y
2830,Icare Franche Compte,\N,,FRC,FRANCHE COMPTE,France,N
2831,Icaro,\N,,ICD,ICARO,Ecuador,N
2832,Icaro,\N,,ICA,ICARFLY,Italy,N
2833,Icarus,\N,,IUS,ICARUS,Italy,N
2834,Icejet,\N,,ICJ,ICEJET,Iceland,N
2835,Icelandair,\N,FI,ICE,ICEAIR,Iceland,Y
2836,Icelandic Coast Guard,\N,,ICG,ICELAND COAST,Iceland,N
2837,Ikaros DK,\N,,IKR,IKAROS,Denmark,N
2838,Il Ciocco International Travel Service,\N,,CIO,CIOCCO,Italy,N
2839,Il-Avia,\N,,ILV,ILAVIA,Russia,N
2840,Ildefonso Redriguez,\N,,IDL,ILDEFONSO,Mexico,N
2841,Iliamna Air Taxi,\N,,IAR,ILIAMNA AIR,United States,N
2842,Ilpo Aruba Cargo,\N,,ILP,,Aruba,N
2843,Ilyich-Avia,\N,,ILL,ILYICHAVIA,Ukraine,N
2844,Imaer,\N,,IMR,IMAER,Portugal,N
2845,Imair Airlines,\N,IK,ITX,IMPROTEX,Azerbaijan,Y
2846,Imperial Airways,\N,,PNX,PHOENIX,United States,N
2847,Imtrec Aviation,\N,,IMT,IMTREC,Cambodia,N
2848,Independence Air,\N,DH,IDE,INDEPENDENCE AIR,United States,N
2849,Independent Air Freighters,\N,,IDP,INDEPENDENT,Australia,N
2850,IndiGo Airlines,\N,6E,IGO,IFLY,India,Y
2851,India International Airways,\N,,IIL,INDIA INTER,India,N
2852,Indian Air Force,\N,,IFC,INDIAN AIRFORCE,India,N
2853,Indian Airlines,\N,IC,IAC,INDAIR,India,Y
2854,Indicator Company,\N,,IDR,INDICATOR,Hungary,N
2855,Indigo,\N,I9,IBU,INDIGO BLUE,United States,Y
2856,Indonesia Air Transport,\N,,IDA,INTRA,Indonesia,N
2857,Indonesia AirAsia,\N,QZ,AWQ,WAGON AIR,Indonesia,Y
2858,Indonesian Airlines,\N,IO,IAA,INDO LINES,Indonesia,Y
2859,Industri Pesawat Terbang Nusantara,\N,,IPN,NUSANTARA,Indonesia,N
2860,Industrias Titan,\N,,ITN,TITANLUX,Spain,N
2861,Infinit Air,\N,,FFI,INFINIT,Spain,N
2862,Inflite The Jet Centre,\N,,INS,,United Kingdom,N
2863,Innotech Aviation,\N,,IVA,INNOTECH,Canada,N
2864,Insel Air International,\N,,INC,INSELAIR,Netherlands Antilles,N
2865,Instituto Cartografico de Cataluna,\N,,ICC,CARTO,Spain,N
2866,Intair,\N,,INT,INTAIRCO,Canada,N
2867,Intal Avia,\N,,INL,INTAL AVIA,Kyrgyzstan,N
2868,Intavia Limited,\N,,FFL,,United Kingdom,N
2869,InteliJet Airways,\N,IJ,IJT,INTELIJET,Colombia,N
2870,Intensive Air,\N,,XRA,INTENSIVE,South Africa,N
2871,Inter Air,\N,,ITW,INTER WINGS,Bulgaria,N
2872,Inter Express,\N,,INX,INTER-EURO,Turkey,N
2873,Inter Islands Airlines,\N,H4,IIN,,Cape Verde,N
2874,Inter RCA,\N,,CAR,QUEBEC ROMEO,Central African Republic,N
2875,Inter Tropic Airlines,\N,,NTT,INTER-TROPIC,Sierra Leone,N
2876,Inter-Air,\N,,ITA,CAFEX,United States,N
2877,Inter-Canadian,\N,,ICN,INTER-CANADIAN,Canada,N
2878,Inter-Island Air,\N,,UGL,UGLY VAN,United States,N
2879,Inter-Mountain Airways,\N,,IMA,INTER-MOUNTAIN,United States,N
2880,Inter-State Aviation,\N,,ITS,INTER-STATE,United States,N
2881,Interair South Africa,\N,D6,ILN,INLINE,South Africa,Y
2882,Interaire,\N,,NTE,INTERMEX,Mexico,N
2883,Interavia Airlines,\N,ZA,SUW,ASTAIR,Russia,Y
2884,Interaviatrans,\N,,IVT,INTERAVIA,Ukraine,N
2885,Intercontinental de Aviaci,\N,RS,ICT,CONTAVIA,Colombia,N
2886,Intercopters,\N,,ICP,CHOPER,Spain,N
2887,Interflight,\N,,IFT,INTERFLIGHT,United Kingdom,N
2888,Interflight (Learjet),\N,,IJT,,United Kingdom,N
2889,Interfly,\N,,RFL,INFLY,Italy,N
2890,Interfreight Forwarding,\N,,IFF,INTERFREIGHT,Sudan,N
2891,Interguide Air,\N,,IGN,DIVINE AIR,Nigeria,N
2892,Interisland Airlines,\N,,ISN,TRI-BIRD,Philippines,N
2893,Interisland Airways Limited,\N,,IWY,ISLANDWAYS,Turks and Caicos Islands,N
2894,Interjet,\N,,MTF,INTERJET,Italy,N
2895,Interjet Helicopters,\N,,IHE,INTERCOPTER,Greece,N
2896,Interlink Airlines,\N,ID,ITK,INTERLINK,South Africa,Y
2897,International Air Cargo Corporation,\N,,IAK,AIR CARGO EGYPT,Egypt,N
2898,International Air Corporation,\N,,EXX,EXPRESS INTERNATIONAL,United States,N
2899,International Air Service,\N,,IAS,STARFLEET,United States,N
2900,International Air Services,\N,,IAX,INTERAIR SERVICES,Liberia,N
2901,International Business Air,\N,6I,IBZ,INTERBIZ,Sweden,N
2902,International Business Aircraft,\N,,IBY,CENTRAL STAGE,United States,N
2903,International Charter Services,\N,,ICS,INTERSERVI,Mexico,N
2904,International Charter Xpress,\N,,ICX,INTEX,United States,N
2905,International Committee of the Red Cross,\N,,RED,RED CROSS,Switzerland,N
2906,International Company for Transport,\N,,IIG, Trade and Public Works,ALDAWLYH AIR,N
2907,International Flight Training Academy,\N,,IFX,IFTA,United States,N
2908,International Jet Aviation Services,\N,,IJA,I-JET,United States,N
2909,International Jet Charter,\N,,HSP,HOSPITAL,United States,N
2910,International SOS WIndhoek,\N,,RSQ,SKYMEDIC,Namibia,N
2911,International Sabilisation Assistance Force,\N,,ISF,,United Kingdom,N
2912,International Security Assistance Force,\N,,THN,ATHENA,Canada,N
2913,International Trans-Air,\N,,ITH,INTRANS NIGERIA,Nigeria,N
2914,Interport Corporation,\N,,IPT,INTERPORT,United States,N
2915,Intersky,\N,,IKY,GENERAL SKY,Bulgaria,N
2916,Intersky,\N,3L,ISK,INTERSKY,Austria,Y
2917,Interstate Airline,\N,I4,FWA,FREEWAYAIR,Netherlands,N
2918,Intervuelos,\N,,ITU,INTERLOS,Mexico,N
2919,Inversija,\N,,INV,INVER,Latvia,N
2920,Iona National Airways,\N,,IND,IONA,Ireland,N
2921,Iowa Airways,\N,,IOA,IOWA AIR,United States,N
2922,Iran Air,\N,IR,IRA,IRANAIR,Iran,Y
2923,Iran Aseman Airlines,\N,EP,IRC,,Iran,Y
2924,Iranair Tours,\N,,IRB,,Iran,N
2925,Naft Air Lines,\N,,IRG,NAFT,Iran,N
2926,Iraqi Airways,\N,IA,IAW,IRAQI,Iraq,Y
2927,Irbis Air,\N,,BIS,IRBIS,Kazakhstan,N
2928,Irish Air Corps,\N,,IRL,IRISH,Ireland,N
2929,Irish Air Transport,\N,,RDK,IRISH TRANS,Ireland,N
2930,Irish Aviation Authority,\N,,XMR,AUTHORITY,Ireland,N
2931,Irtysh Airlines,\N,,IRT,IRTYSH AIRLINES,Uzbekistan,N
2932,Irving Oil,\N,,XIA,,Canada,N
2933,Irving Oil Transport,\N,,KCE,KACEY,Canada,N
2934,Island Air,\N,,ISI,ISLANDMEX,Mexico,N
2935,Island Air Charters,\N,,ILF,ISLAND FLIGHT,United States,N
2936,Island Air Express,\N,,XYZ,RAINBIRD,United States,N
2937,Island Airlines,\N,IS,,,United States,Y
2938,Island Aviation,\N,,SOY,SORIANO,Philippines,N
2939,Island Aviation Services,\N,,DQA,,Maldives,N
2940,Island Aviation and Travel,\N,,IOM,ISLE AVIA,United Kingdom,N
2941,Island Express,\N,2S,SDY,SANDY ISLE,United States,N
2942,Cargo Plus Aviation,\N,8L,CGP,,United Arab Emirates,Y
2943,Island Helicopters,\N,,MTP,METROCOPTER,United States,N
2944,Island Link,\N,,ILC,,Japan,N
2945,Islandair Jersey,\N,,IAJ,JARLAND,United Kingdom,N
2946,Islands Nationair,\N,CN,,,Papua New Guinea,N
2947,Icebird Airline,\N,,ICB,ICEBIRD,Iceland,N
2948,Islas Airways,\N,IF,ISW,PINTADERA,Spain,Y
2949,Isle Grande Flying School,\N,,IGS,ISLA GRANDE,United States,N
2950,Islena De Inversiones,\N,WC,ISV,,Honduras,Y
2951,Isles of Scilly Skybus,\N,,IOS,SCILLONIA,United Kingdom,N
2952,Israel Aircraft Industries,\N,,IAI,ISRAEL AIRCRAFT,Israel,N
2953,Israeli Air Force,\N,,IAF,,Israel,N
2954,Israir,\N,6H,ISR,ISRAIR,Israel,Y
2955,Istanbul Airlines,\N,,IST,ISTANBUL,Turkey,N
2956,Itali Airlines,\N,9X,ACL,ITALI,Italy,N
2957,Italy First,\N,,IFS,RIVIERA,Italy,N
2958,Itek Air,\N,GI,IKA,ITEK-AIR,Kyrgyzstan,Y
2959,Ivoire Aero Services,\N,,IVS,IVOIRE AERO,Ivory Coast,N
2960,Ivoire Airways,\N,,IVW,IVOIRAIRWAYS,Ivory Coast,N
2961,Ivoire Jet Express,\N,,IJE,IVOIRE JET,Ivory Coast,N
2962,Iwamoto Crane Co Ltd,\N,,OIC,,Japan,N
2963,Ixair,\N,,IXR,X-BIRD,France,N
2964,Izair,\N,H9,IZM,IZMIR,Turkey,N
2965,Izhavia,\N,,IZA,IZHAVIA,Russia,N
2966,J C Bamford (Excavators),\N,,JCB,JAYSEEBEE,United Kingdom,N
2967,J P Hunt Air Carriers,\N,,RFX,REFLEX,United States,N
2968,J-Air,\N,,JLJ,J AIR,Japan,N
2969,JAL Express,\N,JC,JEX,JANEX,Japan,Y
2970,JALways,\N,JO,JAZ,JALWAYS,Japan,Y
2971,JDAviation,\N,,JDA,JAY DEE,United Kingdom,N
2972,JDP Lux,\N,,JDP,RED PELICAN,Luxembourg,N
2973,JHM Cargo Expreso,\N,,JHM,,Costa Rica,N
2974,JM Family Aviation,\N,,TQM,TACOMA,United States,N
2975,JMC Airlines,\N,,JMC,JAYEMMSEE,United Kingdom,N
2976,JSC Transport Automated Information Systems,\N,1M,,,Russia,N
2977,JS Air,\N,,JSJ,JS CHARTER,Pakistan,N
2978,JS Aviation,\N,,JES,JAY-ESS AVIATION,Mexico,N
2979,Jackson Air Services,\N,,JCK,JACKSON,Canada,N
2980,Jade Cargo International,\N,,JAE,JADE CARGO,China,N
2981,Jamahiriya Airways,\N,,JAW,JAW,Libya,N
2982,Jambo Africa Airlines,\N,,JMB,JAMBOAFRICA,Democratic Republic of Congo,N
2983,Jana-Arka,\N,,JAK,YANZAR,Kazakhstan,N
2984,Janair,\N,,JAX,JANAIR,United States,N
2985,Japan Air Commuter,\N,,JAC,COMMUTER,Japan,N
2986,Japan Aircraft Service,\N,,JSV,,Japan,N
2987,Japan Airlines,JAL Japan Airlines,JL,JAL,JAPANAIR,Japan,Y
2988,Japan Airlines Domestic,\N,JL,JAL,J-BIRD,Japan,Y
2989,Japan Asia Airways,\N,EG,JAA,ASIA,Japan,Y
2990,Japan Transocean Air,\N,NU,JTA,JAI OCEAN,Japan,Y
2991,Jat Airways,\N,JU,JAT,JAT,Serbia,Y
2992,Jatayu Airlines,\N,VJ,JTY,JATAYU,Indonesia,N
2993,Jazeera Airways,\N,J9,JZR,JAZEERA,Kuwait,Y
2994,Jeju Air,\N,7C,JJA,JEJU AIR,Republic of Korea,Y
2995,Jenney Beechcraft,\N,,JNY,JENAIR,United States,N
2996,Jeppesen Data Plan,\N,,XLD,,United States,N
2997,Jeppesen UK,\N,,JPN,JETPLAN,United Kingdom,N
2998,Jet Air,\N,,JEA,JETA,Poland,N
2999,Jet Air Group,\N,,JSI,SISTEMA,Russia,N
3000,Jet Airways,\N,9W,JAI,JET AIRWAYS,India,Y
3001,Jet Airways,\N,QJ,,,United States,Y
3002,Jet Aspen Air Lines,\N,,JTX,JET ASPEN,United States,N
3003,Jet Aviation,\N,,PJS,JETAVIATION,Switzerland,N
3004,Jet Aviation Business Jets,\N,,BZF,BIZFLEET,United States,N
3005,Jet Center Flight Training,\N,,JCF,JET CENTER,Spain,N
3006,Jet Charter,\N,,JCT,JET CHARTER,United States,N
3007,Jetclub,\N,0J,,,Switzerland,N
3008,Jet Connection,\N,,JCX,JET CONNECT,Germany,N
3009,Jet Courier Service,\N,,DWW,DON JUAN,United States,N
3010,Jet East International,\N,,JED,JET EAST,United States,N
3011,Jet Executive International Charter,\N,,JEI,JET EXECUTIVE,Germany,N
3012,Jet Fighter Flights,\N,,RZA,RAZOR,Australia,N
3013,Jet Freighters,\N,,CFT,CASPER FREIGHT,United States,N
3014,Jet G&D Aviation,\N,,JGD,JET GEE-AND-DEE,Israel,N
3015,Jet Line International,\N,,MJL,MOLDJET,Moldova,N
3016,Jet Link,\N,,JEK,JET OPS,Israel,N
3017,Jet Link Aviation,\N,,HTL,HEARTLAND,United States,N
3018,Jet Norte,\N,,JNR,JET NORTE,Mexico,N
3019,Jet Rent,\N,,JRN,JET RENT,Mexico,N
3020,Jet Service,\N,,JDI,JEDI,Poland,N
3021,Jetstar Asia Airways,\N,3K,JSA,JETSTAR ASIA,Singapore,Y
3022,Jet Stream,\N,,JSM,JET STREAM,Moldova,N
3023,Jet Trans Aviation,\N,,JTC,JETRANS,Ghana,N
3024,Jet-2000,\N,,JTT,MOSCOW JET,Russia,N
3025,Jet-Ops,\N,,OPS,OPS-JET,United Arab Emirates,N
3026,Jet2.com,\N,LS,EXS,CHANNEX,United Kingdom,Y
3027,Jet4You,,8J,JFU,ARGAN,Morocco,Y
3028,JetAfrica Swaziland,\N,,OSW,BEVO,Swaziland,N
3029,JetBlue Airways,\N,B6,JBU,JETBLUE,United States,Y
3030,JetConnect,\N,,QNZ,QANTAS JETCONNECT,New Zealand,N
3031,JetMagic,\N,,JMG,JET MAGIC,Ireland,N
3032,Jetairfly,\N,JF,JAF,BEAUTY,Belgium,Y
3033,Jetall Holdings,\N,,JTL,FIREFLY,Canada,N
3034,Jetalliance,\N,,JAG,JETALLIANCE,Austria,N
3035,Jetclub,\N,0J,JCS,JETCLUB,Switzerland,N
3036,Jetcorp,\N,,UEJ,JETCORP,United States,N
3037,Jetcraft Aviation,\N,,JCC,JETCRAFT,Australia,N
3038,Jetex Aviation,\N,,JXA,,Lebanon,N
3039,Jetflite,\N,,JEF,JETFLITE,Finland,Y
3040,Jetfly Airlines,\N,,JFL,LINEFLYER,Austria,N
3041,Jetgo International,\N,,JIC,JIC-JET,Thailand,N
3042,Jetlink Express,\N,,JLX,KEN JET,Kenya,N
3043,Jetlink Holland,\N,,JLH,,Netherlands,N
3044,Jetnova de Aviacion Ejecutiva,\N,,JNV,JETNOVA,Spain,N
3045,Jetpro,\N,,JPO,JETPRO,Mexico,N
3046,Jetran Air,\N,,MDJ,JETRAN AIR,Romania,N
3047,Jetrider International,\N,,JRI,JETRIDER,United Kingdom,N
3048,Jets Ejecutivos,\N,,JEJ,MEXJETS,Mexico,N
3049,Jets Personales,\N,,JEP,JET PERSONALES,Spain,N
3050,Jets Y Servicios Ejecutivos,\N,,JSE,SERVIJETS,Mexico,N
3051,JetsGo,\N,SG,JGO,JETSGO,Canada,N
3052,Jetstar Airways,\N,JQ,JST,JETSTAR,Australia,Y
3053,Jetstream Air,\N,,JSH,STREAM-AIR,Hungary,N
3054,Jetstream Executive Travel,\N,,JXT,VANNIN,United Kingdom,N
3055,Jett Paqueteria,\N,,JPQ,JETT PAQUETERIA,Mexico,N
3056,Jett8 Airlines Cargo,\N,JX,JEC,,Singapore,N
3057,Jettime,\N,,JTG,JETTIME,Denmark,N
3058,Jettrain Corporation,\N,,JTN,JETTRAIN,United States,N
3059,Jetways of Iowa,\N,,JWY,JETWAYS,United States,N
3060,Jetx Airlines,\N,GX,JXX,JETBIRD,Iceland,N
3061,Jibair,\N,,JIB,JIBAIRLINE,Djibouti,N
3062,Jigsaw Project,\N,,JSW,JIGSAW,United Kingdom,N
3063,Jim Hankins Air Service,\N,,HKN,HANKINS,United States,N
3064,Jim Ratliff Air Service,\N,,RAS,SHANHIL,United States,N
3065,Joanas Avialinijos,\N,,JDG,LADYBLUE,Lithuania,N
3066,Job Air,\N,,JBR,JOBAIR,Czech Republic,N
3067,Johnson Air,\N,,JHN,AIR JOHNSON,United States,N
3068,Johnsons Air,\N,,JON,JOHNSONSAIR,Ghana,N
3069,Johnston Airways,\N,,JMJ,JOHNSTON,United States,N
3070,Joint Military Commission,\N,,JMM,JOICOMAR,Sudan,N
3071,Jomartaxi Aereo,\N,,JMT,JOMARTAXI,Mexico,N
3072,Jonsson,\N,,ODI, H Air Taxi,ODINN,N
3073,Jordan Aviation,\N,R5,JAV,JORDAN AVIATION,Jordan,N
3074,Jordan International Air Cargo,\N,,JCI,,Jordan,N
3075,Jorvik,\N,,JVK,ISLANDIC,Iceland,N
3076,Ju-Air,\N,,JUR,JUNKERS,Switzerland,N
3077,Juanda Flying School,\N,,JFS,JAEMCO,Indonesia,N
3078,Juba Cargo Services & Aviation Company,\N,,JUC,JUBA CARGO,Sudan,N
3079,Jubba Airways,\N,,JUB,JUBBA,Somali Republic,N
3080,Jubilee Airways,\N,,DKE,DUKE,United Kingdom,N
3081,Juneyao Airlines,\N,HO,DKH,JUNEYAO AIRLINES,China,Y
3082,Justair Scandinavia,\N,,MEY,MELODY,Sweden,N
3083,Justice Prisoner and Alien Transportation System,\N,,DOJ,JUSTICE,United States,N
3084,K D Air Corporation,\N,,KDC,KAY DEE,Canada,N
3085,K S Avia,\N,,KSA,SKY CAMEL,Latvia,N
3086,K-Mile Air,\N,,KMI,KAY-MILE AIR,Thailand,N
3087,KD Avia,\N,KD,KNI,KALININGRAD AIR,Russia,Y
3088,KLM Cityhopper,\N,WA,KLC,CITY,Netherlands,Y
3089,KLM Helicopter,\N,,KLH,KLM HELI,Netherlands,N
3090,KLM Royal Dutch Airlines,\N,KL,KLM,KLM,Netherlands,Y
3091,Kabo Air,\N,N2,QNK,KABO,Nigeria,N
3092,Kahama Mining Corporation,\N,,KMC,KAHAMA,Tanzania,N
3093,Kaiser Air,\N,,KAI,KAISER,United States,N
3094,Kalitta Air,\N,K4,CKS,CONNIE,United States,N
3095,Kalitta Charters,\N,,KFS,KALITTA,United States,N
3096,Kallat El Saker Air Company,\N,,KES,KALLAT EL SKER,Libya,N
3097,Kam Air,\N,RQ,KMF,KAMGAR,Afghanistan,Y
3098,Kampuchea Airlines,\N,E2,KMP,KAMPUCHEA,Cambodia,N
3099,Kanaf-Arkia Airlines,\N,,KIZ,,Israel,N
3100,Kanfey Ha'emek Aviation,\N,,KHE,KANFEY HAEMEK,Israel,N
3101,Kansas State University,\N,,KSU,K-STATE,United States,N
3102,Karat,\N,V2,AKT,AVIAKARAT,Russia,N
3103,Karibu Airways Company,\N,,KRB,KARIBU AIR,Tanzania,N
3104,Karlog Air Charter,\N,,KLG,KARLOG,Denmark,N
3105,Karthago Airlines,\N,,KAJ,KARTHAGO,Tunisia,N
3106,Kartika Airlines,\N,,KAE,KARTIKA,Indonesia,N
3107,Kata Transportation,\N,,KTV,KATAVIA,Sudan,N
3108,Katekavia,\N,,KTK,KATEKAVIA,Russia,N
3109,Kato Airline,\N,,KAT,KATO-AIR,Norway,N
3110,Kavminvodyavia,\N,KV,MVD,AIR MINVODY,Russia,Y
3111,Kavouras Inc,\N,,XKA,,United States,N
3112,Kaz Agros Avia,\N,,KRN,ANTOL,Kazakhstan,N
3113,Kaz Air West,\N,,KAW,KAZWEST,Kazakhstan,N
3114,Kazan Aviation Production Association,\N,,KAO,KAZAVAIA,Russia,N
3115,Kazan Helicopters,\N,,KPH,KAMA,Russia,N
3116,Kazavia,\N,,KKA,KAKAIR,Kazakhstan,N
3117,Kazaviaspas,\N,,KZS,SPAKAZ,Kazakhstan,N
3118,Keenair Charter -,\N,,JFK,KEENAIR,United Kingdom,N
3119,Kelix Air,\N,,KLX,KELIX,Nigeria,N
3120,Kelner Airways,\N,,FKL,KELNER,Canada,N
3121,Kelowna Flightcraft Air Charter,\N,,KFA,FLIGHTCRAFT,Canada,N
3122,Kendell Airlines,\N,,KDA,KENDELL,Australia,Y
3123,Kenmore Air,\N,M5,KEN,KENMORE,United States,Y
3124,Kenn Borek Air,\N,,KBA,BOREK AIR,Canada,N
3125,Kent Aviation,\N,,KAH,DEKAIR,Canada,N
3126,Kenya Airways,\N,KQ,KQA,KENYA,Kenya,Y
3127,Kevis,\N,,KVS,KEVIS,Kazakhstan,N
3128,Key Airlines,\N,,KEY,KEY AIR,United States,N
3129,Key Lime Air,\N,,LYM,KEY LIME,United States,N
3130,Keystone Aerial Surveys,\N,,FTP,FOOTPRINT,United States,N
3131,Keystone Air Services,\N,BZ,KEE,KEYSTONE,Canada,N
3132,Khalifa Airways,\N,,KZW,KHALIFA AIR,Algeria,N
3133,Kharkov Aircraft Manufacturing Company,\N,,WKH,WEST-KHARKOV,Ukraine,N
3134,Khazar,\N,,KHR,KHAZAR,Turkmenistan,N
3135,Khoezestan Photros Air Lines,\N,,KHP,PHOTROS AIR,Iran,N
3136,Khoriv-Avia,\N,,KRV,KHORIV-AVIA,Ukraine,N
3137,Khors Aircompany,\N,,KHO,AIRCOMPANY KHORS,Ukraine,N
3138,Khyber Afghan Airlines,\N,,KHY,KHYBER,Afghanistan,N
3139,Kiev Aviation Plant,\N,,UAK,AVIATION PLANT,Ukraine,N
3140,King Aviation,\N,,KNG,KING,United Kingdom,N
3141,Kingfisher Air Services,\N,,BEZ,SEA BREEZE,United States,N
3142,Kingfisher Airlines,\N,IT,KFR,KINGFISHER,India,Y
3143,Knighthawk Air Express,\N,,KNX,KNIGHT FLIGHT,Canada,N
3144,Kingston Air Services,\N,,KAS,KINGSTON AIR,Canada,N
3145,Kinnarps,\N,,KIP,KINNARPS,Sweden,N
3146,Kinshasa Airways,\N,,KNS,KINSHASA AIRWAYS,Democratic Republic of the Congo,N
3147,Kirov Air Enterprise,\N,,KTA,VYATKA-AVIA,Russia,N
3148,Kish Air,\N,Y9,IRK,KISHAIR,Iran,Y
3149,Kitty Hawk Aircargo,\N,,KHA,AIR KITTYHAWK,United States,N
3150,Kitty Hawk Airways,\N,,KHC,CARGO HAWK,United States,N
3151,Kiwi International Air Lines,\N,KP,KIA,KIWI AIR,United States,N
3152,Knight Air,\N,,KNA,KNIGHTAIR,Canada,N
3153,Knighthawk Express,\N,,KHX,RIZZ,United States,N
3154,Knights Airlines,\N,,KGT,KNIGHT-LINER,Nigeria,N
3155,Koanda Avacion,\N,,KOA,KOANDA,Spain,N
3156,Koda International,\N,,OYE,KODA AIR,Nigeria,N
3157,Kogalymavia Air Company,\N,7K,KGL,KOGALYM,Russia,Y
3158,Kom Activity,\N,,KOM,COMJET,Netherlands,N
3159,Komiaviatrans State Air Enterprise,\N,,KMA,KOMI AVIA,Russia,N
3160,Komiinteravia,,8J,KMV,KOMIINTER,Russia,N
3161,Komsomolsk-on-Amur Air Enterprise,\N,,KNM,KNAAPO,Russia,N
3162,Koob-Corp - 96 KFT,\N,,KOB,AUTOFLEX,Hungary,N
3163,Korean Air,\N,KE,KAL,KOREANAIR,Republic of Korea,Y
3164,Kosmas Air,\N,,KMG,KOSMAS CARGO,Serbia,N
3165,Kosmos,\N,,KSM,KOSMOS,Russia,Y
3166,Kosova Airlines,\N,,KOS,KOSOVA,Serbia,N
3167,Kovar Air,\N,,WOK,WOKAIR,Czech Republic,N
3168,Krasnojarsky Airlines,\N,7B,KJC,KRASNOJARSKY AIR,Russia,Y
3169,Kremenchuk Flight College,\N,,KFC,KREMENCHUK,Ukraine,N
3170,Krimaviamontag,\N,,KRG,AVIAMONTAG,Ukraine,N
3171,Kroonk Air Agency,\N,,KRO,KROONK,Ukraine,N
3172,Krylo Airlines,\N,K9,KRI,Krylo,Russia,N
3173,Krym,\N,,KYM,CRIMEA AIR,Ukraine,N
3174,Krystel Air Charter,\N,,OPC,OPTIC,United Kingdom,N
3175,Kuban Airlines,\N,GW,KIL,AIR KUBAN,Russia,Y
3176,Kunpeng Airlines,\N,VD,KPA,KUNPENG,China,N
3177,Kurzemes Avio,\N,,KZA,,Russia,N
3178,Kustbevakningen,\N,,KBV,SWECOAST,Sweden,N
3179,Kuwait Airways,\N,KU,KAC,KUWAITI,Kuwait,Y
3180,Kuzu Airlines Cargo,\N,GO,KZU,KUZU CARGO,Turkey,Y
3181,Kvadro Aero,\N,,QVR,PEGASO,Kyrgyzstan,N
3182,Kwena Air,\N,,KWN,KWENA,South Africa,N
3183,Kyrgyz Airlines,\N,N5,KGZ,BERMET,Kyrgyzstan,N
3184,Kyrgyz Trans Avia,\N,,KTC,DINARA,Kyrgyzstan,N
3185,Kyrgyzstan,\N,QH,LYN,ALTYN AVIA,Kyrgyzstan,N
3186,Kyrgyzstan Airlines,\N,R8,KGA,KYRGYZ,Kyrgyzstan,N
3187,Kyrgyzstan Department of Aviation,\N,,DAM,FLIGHT RESCUE,Kyrgyzstan,N
3188,Kyrgz General Aviation,\N,,KGB,KEMIN,Kyrgyzstan,N
3189,Kibris T,\N,KY,KYV,AIRKIBRIS,Turkey,N
3190,L A Helicopter,\N,,LAH,STAR SHIP,United States,N
3191,L J Aviation,\N,,LJY,ELJAY,United States,N
3192,L R Airlines,\N,,LRB,LADY RACINE,Czech Republic,N
3193,L&L Flygbildteknik,\N,,PHO,PHOTOFLIGHT,Sweden,N
3194,L'Express,\N,,LEX,LEX,United States,N
3195,L-3 Communications Flight Internation Aviation,\N,,FNT,FLIGHT INTERNATIONAL,United States,N
3196,L.A.B. Flying Service,\N,JF,LAB,LAB,United States,N
3197,LACSA,\N,LR,LRC,LACSA,Costa Rica,Y
3198,LADE - Lineas Aereas Del Estado,\N,,LDE,LADE,Argentina,N
3199,LAI - Linea Aerea IAACA,\N,KG,BNX,AIR BARINAS,Venezuela,N
3200,LAN Airlines,\N,LA,LAN,LAN,Chile,Y
3201,LAN Argentina,\N,4M,DSM,LAN AR,Argentina,Y
3202,LAN Cargo,\N,,LCO,LAN CARGO,Chile,N
3203,LAN Dominica,\N,,LNC,LANCANA,Dominican Republic,N
3204,LAN Express,\N,LU,LXP,LANEX,Chile,Y
3205,LAN Peru,\N,LP,LPE,LANPERU,Peru,Y
3206,LANSA,\N,,LSA,INTERNACIONAL,Dominican Republic,N
3207,LAP Colombia - Lineas Aereas Petroleras,\N,,APT, S.A.,LAP,N
3208,LASTP,\N,,OTN,LASTP,S,N
3209,LC Busre,\N,,LCB,BUSRE,Peru,N
3210,LOT Polish Airlines,\N,LO,LOT,POLLOT,Poland,Y
3211,LTE International Airways,\N,XO,LTE,FUN JET,Spain,Y
3212,LTU Austria,\N,L3,LTO,BILLA TRANSPORT,Austria,Y
3213,LTU International,\N,LT,LTU,LTU,Germany,Y
3214,LTV Jet Fleet Corporation,\N,,JFC,JET-FLEET,United States,N
3215,LUKoil-Avia,\N,,LUK,LUKOIL,Russia,N
3216,La Ronge Aviation Services,\N,,ASK,AIR SASK,Canada,N
3217,La Valenciana Taxi Aereo,\N,,LVT,TAXIVALENCIANA,Mexico,N
3218,Labcorp,\N,,SKQ,SKYLAB,United States,N
3219,Labrador Airways,\N,,LAL,LAB AIR,Canada,N
3220,Lagun Air,\N,N6,JEV,,Spain,N
3221,Lake Havasu Air Service,\N,,HCA,HAVASU,United States,N
3222,Lakeland Aviation,\N,,LKL,LAKELAND,United States,N
3223,Laker Airways,\N,,LKR,LAKER,United States,N
3224,Laker Airways (Bahamas),\N,,LBH,LAKER BAHAMAS,United States,N
3225,Lamra,\N,,LMR,LAMAIR,Sudan,N
3226,Lanaes Aereas Trans Costa Rica,\N,,TCR,TICOS,Costa Rica,N
3227,Landsflug,\N,,ISL,ISLANDIA,Iceland,N
3228,Langtry Flying Group,\N,,PAP,PROFLIGHT,United Kingdom,N
3229,Lankair,\N,IK,LKN,Lankair,Sri Lanka,N
3230,Lanza Air,\N,,LZA,AEROLANZA,Spain,N
3231,Lanzarote Aerocargo,\N,,LZT,BARAKA,Spain,N
3232,Lao Air Company,\N,,LLL,LAVIE,Lao Peoples Democratic Republic,N
3233,Lao Airlines,\N,QV,LAO,LAO,Lao Peoples Democratic Republic,Y
3234,Lao Capricorn Air,\N,,LKA,NAKLOA,Lao Peoples Democratic Republic,N
3235,Laoag International Airlines,\N,L7,LPN,LAOAG AIR,Philippines,N
3236,Laredo Air,\N,,LRD,LAREDO AIR,United States,N
3237,LatCharter,\N,,LTC,LATCHARTER,Latvia,Y
3238,Latvian Air Force,\N,,LAF,LATVIAN AIRFORCE,Latvia,N
3239,Lauda Air,\N,NG,LDA,LAUDA AIR,Austria,Y
3240,Lauda Air Italy,\N,,LDI,LAUDA ITALY,Italy,N
3241,Laughlin Express,\N,,LEP,LAUGHLIN EXPRESS,United States,N
3242,Laus,\N,,LSU,LAUS AIR,Croatia,N
3243,Lawrence Aviation,\N,,LAR,LAWRENCE,United States,N
3244,Layang-Layang Aerospace,\N,,LAY,LAYANG,Malaysia,N
3245,Lease-a-Plane International,\N,,LPL,LEASE-A-PLANE,United States,N
3246,Lebanese Air Transport,\N,LQ,LAQ,LAT,Lebanon,N
3247,Lebanese Air Transport (Charter),\N,,LAT,LEBANESE AIR,Lebanon,N
3248,Lebanon Airport Development Corporation,\N,,LAD,LADCO-AIR,United States,N
3250,Leconte Airlines,\N,,LCA,LECONTE,United States,N
3251,Leeward Islands Air Transport,\N,LI,LIA,LIAT,Antigua and Barbuda,Y
3252,Legend Airlines,\N,,LGD,LEGENDARY,United States,N
3253,Leisure Air,\N,,LWD,LEISURE WORLD,United States,N
3254,Lentini Aviation,\N,,LEN,LENTINI,United States,N
3255,Leo-Air,\N,,LOR,LEO CHARTER,South Africa,N
3256,Leonsa De Aviacion,\N,,LEL,LEONAVIA,Spain,N
3257,Libyan Airlines,\N,,LYW,LIBYAN AIRWAYS,Libya,N
3258,Libyan Arab Airlines,\N,LN,LAA,LIBAIR,Libya,Y
3259,Libyan Arab Air Cargo,\N,,LCR,LIBAC,Libya,N
3260,Lid Air,\N,,LIQ,,Sweden,N
3261,Lignes Aeriennes Congolaises,\N,,LCG,CONGOLAISE,Democratic Republic of the Congo,N
3262,Lignes Aeriennes Du Tchad,\N,,LKD,LATCHAD,Chad,N
3263,Lignes Mauritaniennes Air Express,\N,,LME,LIMAIR EXPRESS,Mauritania,N
3264,Lignes Nationales Aeriennes - Linacongo,\N,,GCB,LINACONGO,Republic of the Congo,N
3265,Lincoln Air National Guard,\N,,GDQ,,United States,N
3266,Lincoln Airlines,\N,,LRT,,Australia,N
3267,Lindsay Aviation,\N,,LSY,LINDSAY AIR,United States,N
3268,Linea Aerea Costa Norte,\N,,NOT,COSTA NORTE,Chile,N
3269,Linea Aerea Mexicana de Carga,\N,,LMC,LINEAS DECARGA,Mexico,N
3270,Linea Aerea SAPSA,\N,L7,LNP,SAPSA,Chile,N
3271,Linea Aerea de Fumig Aguas Negras,\N,,NEG,AGUAS NEGRAS,Chile,N
3272,Linea Aerea de Servicio Ejecutivo Regional,\N,8z,LER,LASER,Venezuela,N
3273,Linea De Aeroservicios,\N,,LSE,,Chile,N
3274,Linea Turistica Aerotuy,\N,LD,TUY,AEREOTUY,Venezuela,N
3275,Lineas Aereas Alaire S.L.,\N,,ALR,AEROLAIRE,Spain,N
3276,Lineas Aereas Azteca,\N,ZE,LCD,LINEAS AZTECA,Mexico,N
3277,Lineas Aereas Canedo LAC,\N,,LCN,CANEDO,Bolivia,N
3278,Lineas Aereas Comerciales,\N,,LCM,LINEAS COMERCIALES,Mexico,N
3279,Lineas Aereas Ejectuivas De Durango,\N,,EDD,LINEAS DURANGO,Mexico,N
3280,Lineas Aereas Eldorado,\N,,EDR,ELDORADRO,Colombia,N
3281,Lineas Aereas Federales,\N,,FED,FEDERALES,Argentina,N
3282,Lineas Aereas Monarca,\N,,LMN,LINEAS MONARCA,Mexico,N
3283,Lineas Aereas San Jose,\N,,LIJ,LINEAS JOSE,Mexico,N
3284,Lineas Aereas del Humaya,\N,,UMA,HUMAYA,Mexico,N
3285,Linex,\N,,LEC,LECA,Central African Republic,N
3286,Linhas Aereas Santomenses,\N,,SMS,SANTOMENSES,S,N
3287,Linhas A,\N,LM,LAM,MOZAMBIQUE,Mozambique,Y
3288,Link Airways of Australia,\N,,LAW,,Australia,N
3289,Lion Air Services,\N,,WGT,WORLDGATE,United Kingdom,N
3290,Lion Mentari Airlines,\N,JT,LNI,LION INTER,Indonesia,Y
3291,Lions-Air,\N,,LEU,LIONSAIR,Switzerland,N
3292,Lithuanian Air Force,\N,,LYF,LITHUANIAN AIRFORCE,Lithuania,N
3293,Little Red Air Service,\N,,LRA,LITTLE RED,Canada,N
3294,Livingston,\N,LM,LVG,LIVINGSTON,Italy,N
3295,Lloyd Aereo Boliviano,\N,LB,LLB,LLOYDAEREO,Bolivia,N
3296,Lnair Air Services,\N,,LNA,ELNAIR,Spain,N
3297,Lockheed Air Terminal,\N,,XLG,,United States,N
3298,Lockeed Aircraft Corporation,\N,,LAC,LOCKHEED,United States,N
3299,Lockheed DUATS,\N,,XDD,,United States,N
3300,Lockheed Martin Aeronautics,\N,,CBD,CATBIRD,United States,N
3301,Lockheed Martin Aeronautics Company,\N,,LNG,LIGHTNING,United States,N
3302,Logan Air,\N,,LOG,LOGAN,United Kingdom,N
3303,Lom Praha Flying School,\N,,CLV,AEROTRAINING,Czech Republic,N
3304,Lomas Helicopters,\N,,LMS,LOMAS,United Kingdom,N
3305,London City Airport Jet Centre,\N,,LCY,LONDON CITY,United Kingdom,N
3306,London Executive Aviation,\N,,LNX,LONEX,United Kingdom,N
3307,London Flight Centre (Stansted),\N,,LOV,LOVEAIR,United Kingdom,N
3308,London Helicopter Centres,\N,,LHC,MUSTANG,United Kingdom,N
3309,Lone Star Airlines,\N,,LSS,LONE STAR,United States,N
3310,Long Island Airlines,\N,,ORA,LONG ISLAND,United States,N
3311,Longtail Aviation,\N,,LGT,LONGTAIL,Bermuda,N
3312,Lorraine Aviation,\N,,LRR,LORRAINE,France,N
3313,Los Cedros Aviacion,\N,,LSC,CEDROS,Chile,N
3314,Lotus Air,\N,,TAS,LOTUS FLOWER,Egypt,N
3315,Luchtvaartmaatschappij Twente,\N,,LTW,TWENTAIR,Netherlands,N
3316,Lucky Air,\N,,LKE,LUCKY AIR,China,N
3317,Luft Carago,\N,,LUT,LUGO,South Africa,N
3318,Luftfahrt-Vermietungs-Dienst,\N,,LVD,AIR SANTE,Austria,N
3319,Luftfahrtgesellschaft Walter,\N,HE,LGW,WALTER,Germany,Y
3320,Lufthansa,\N,LH,DLH,LUFTHANSA,Germany,Y
3321,Lufthansa Cargo,\N,GEC,GEC,LUFTHANSA CARGO,Germany,Y
3322,Lufthansa CityLine,\N,CL,CLH,HANSALINE,Germany,Y
3323,Lufthansa Systems,\N,L1,,,Germany,N
3324,Lufthansa Technik,\N,,LHT,LUFTHANSA TECHNIK,Germany,N
3325,Lufttaxi Fluggesellschaft,,DV,LTF,Garfield,Germany,N
3326,Lufttransport,\N,L5,LTR,LUFT TRANSPORT,Norway,Y
3327,Luhansk,\N,,LHS,ENTERPRISE LUHANSK,Ukraine,N
3328,Lund University School of Aviation,\N,,UNY,UNIVERSITY,Sweden,N
3329,Luxair,\N,LG,LGL,LUXAIR,Luxembourg,Y
3330,Luxaviation,\N,,LXA,RED LION,Luxembourg,N
3331,Luxembourg Air Rescue,\N,,LUV,LUX RESCUE,Luxembourg,N
3332,Luxflight Executive,\N,,LFE,LUX EXPRESS,Luxembourg,N
3333,Luxor Air,\N,,LXO,,Egypt,N
3334,Luzair,\N,,LUZ,LISBON JET,Portugal,N
3335,Lviv Airlines,\N,5V,UKW,UKRAINE WEST,Ukraine,N
3336,Lydd Air,\N,,LYD,LYDDAIR,United Kingdom,N
3337,Lynch Flying Service,\N,,LCH,LYNCH AIR,United States,N
3338,Lynden Air Cargo,\N,L2,LYC,LYNDEN,United States,N
3339,Lynx Air International,\N,,LXF,LYNX FLIGHT,United States,N
3340,Lynx Aviation,\N,,LYX,LYNX AIR,Pakistan,N
3342,L,\N,MJ,LPR,LAPA,Argentina,Y
3343,L,\N,,LAU,SURAMERICANO,Colombia,N
3344,M & N Aviation,\N,,JNH,JONAH,United States,N
3345,MAC Fotografica,\N,,MCF,MAC FOTO,Spain,N
3346,MANAG'AIR,\N,,MRG,MANAG'AIR,France,N
3347,MAP-Management and Planung,\N,,MPJ,MAPJET,Austria,N
3348,MAS Airways,\N,,TFG,TRAFALGAR,United Kingdom,N
3349,MasAir,\N,M7,MAA,MAS CARGA,Mexico,Y
3350,MAT Macedonian Airlines,\N,IN,MAK,MAKAVIO,Macedonia,Y
3351,MCC Aviation,\N,,MCC,DISCOVERY,South Africa,N
3352,MG Aviacion,\N,,MGA,MAG AVACION,Spain,N
3353,MIA Airlines,\N,,JLA,SALLINE,Romania,N
3354,MIAT Mongolian Airlines,\N,OM,MGL,MONGOL AIR,Mongolia,Y
3355,MIT Airlines,\N,,MNC,MUNCIE,Canada,N
3356,MK Airline,\N,,MKA,KRUGER-AIR,Ghana,N
3357,MNG Airlines,\N,MB,MNB,BLACK SEA,Turkey,Y
3358,MSR Flug-Charter,\N,,EBF,SKYRUNNER,Germany,N
3359,MTC Aviacion,\N,,MCV,MTC AVIACION,Mexico,N
3360,Mac Aviation,\N,,MAQ,MAC AVIATION,Spain,N
3361,Mac Dan Aviation Corporation,\N,,MCN,MAC DAN,United States,N
3362,MacKnight Airlines,\N,,MTD,,Australia,N
3363,Macair Airlines,\N,CC,MCK,,Australia,Y
3364,Macedonian Airlines,\N,,MCS,MACAIR,Greece,N
3365,Madina Air,\N,,MDH,MADINA AIR,Libya,N
3366,Maersk,\N,DM,,,Denmark,Y
3367,Magic Blue Airlines,\N,,MJB,MAGIC BLUE,Netherlands,N
3368,Magna Air,\N,,MGR,MAGNA AIR,Austria,N
3369,Mahalo Air,\N,,MLH,MAHALO,United States,N
3370,Mahan Air,\N,W5,IRM,MAHAN AIR,Iran,Y
3371,Mahfooz Aviation,\N,M2,MZS,MAHFOOZ,Gambia,N
3372,Maine Aviation,\N,,MAT,MAINE-AV,United States,N
3373,Majestic Airlines,\N,,MAJ,MAGIC AIR,United States,N
3374,Mak Air,\N,,AKM,MAKAIR,Kazakhstan,N
3375,Malagasy Airlines,\N,,MLG,,Madagascar,N
3376,Malawi Express,\N,,MLX,MALAWI EXPRESS,Malawi,N
3377,Malaya Aviatsia Dona,\N,,MKK,AEROKEY,Russia,N
3378,Malaysia Airlines,\N,MH,MAS,MALAYSIAN,Malaysia,Y
3379,Mali Air,\N,,MAE,MALI AIREXPRESS,Austria,N
3380,Mali Air Express,\N,,VXP,AVION EXPRESS,Mali,N
3381,Mali Airways,\N,,MTZ,MALI AIRWAYS,Mali,N
3382,Malila Airlift,\N,,MLC,MALILA,Democratic Republic of the Congo,N
3383,Mall Airways,\N,,MLS,MALL-AIRWAYS,United States,N
3384,Malmo Aviation,\N,,SCW,SCANWING,Sweden,Y
3385,Malmoe Air Taxi,\N,,LOD,LOGIC,Sweden,N
3386,Malmö Aviation,\N,TF,SCW,Scanwings,Sweden,Y
3387,Malta Air Charter,\N,R5,MAC,MALTA CHARTER,Malta,Y
3388,Malta Wings,\N,,MWS,MALTA WINGS,Malta,N
3389,Malév,\N,MA,MAH,MALEV,Hungary,Y
3390,Manaf International Airways,\N,,MLB,MANAF,Burundi,N
3391,Mandala Airlines,\N,RI,MDL,MANDALA,Indonesia,Y
3392,Mandarin Airlines,\N,AE,MDA,Mandarin,Taiwan,Y
3393,Mango,\N,JE,MNO,TULCA,South Africa,Y
3394,Manhattan Air,\N,,MHN,MANHATTAN,United Kingdom,N
3395,Manitoulin Air Services,\N,,MTO,MANITOULIN,Canada,N
3396,Mann Air,\N,,MNR,TEEMOL,United Kingdom,N
3397,Mannion Air Charter,\N,,MAN,MANNION,United States,N
3398,Mantrust Asahi Airways,\N,,MTS,MANTRUST,Indonesia,N
3399,Manx Airlines,\N,,MNX,MANX,United Kingdom,N
3400,Maple Air Services,\N,,MAD,MAPLE AIR,Canada,N
3401,March Helicopters,\N,,MAR,MARCH,United Kingdom,N
3402,Marcopolo Airways,\N,,MCP,MARCOPOLO,Afghanistan,N
3403,Marghi Air,\N,,MGI,MARGHI,Nigeria,N
3404,Markair,\N,,MRK,MARKAIR,United States,N
3405,Markoss Aviation,\N,,MKO,GOSHAWK,United Kingdom,N
3406,Mars RK,\N,6V,MRW,AVIAMARS,Ukraine,N
3407,Marshall Aerospace,\N,,MCE,MARSHALL,United Kingdom,N
3408,Marsland Aviation,\N,M7,MSL,MARSLANDAIR,Sudan,N
3409,Martin Aviation Services,\N,,XMA,,United States,N
3410,Martin-Baker,\N,,MBE,MARTIN,United Kingdom,N
3411,Martinair,\N,MP,MPH,MARTINAIR,Netherlands,Y
3412,Martinaire,\N,,MRA,MARTEX,United States,N
3413,Martyn Fiddler Associates,\N,,MFA,SEAHORSE,United Kingdom,N
3414,Marvin Limited,\N,,MVN,MARVIN,United Kingdom,N
3415,Maryland State Police,\N,,TRP,TROOPER,United States,N
3416,Massachusetts Institute of Technology,\N,,MTH,RESEARCH,United States,N
3417,Massey University School of Aviation,\N,,MSY,MASSEY,New Zealand,N
3418,Master Airways,\N,,MSW,MASTER AIRWAYS,Serbia,N
3419,Master Planner,\N,,MPL,,United States,N
3420,Masterjet,\N,,LMJ,MASTERJET,Portugal,N
3421,Mastertop Linhas Aereas,\N,Q4,,,Brazil,N
3422,Mauria,\N,,MIA,MAURIA,Mauritania,N
3423,Mauritanienne Aerienne Et Navale,\N,,MNV,NAVALE,Mauritania,N
3424,Mauritanienne Air Fret,\N,,MRF,MAUR-FRET,Mauritania,N
3425,Mauritanienne Airways,\N,,MWY,MAURITANIENNE,Mauritania,N
3426,Mauritanienne De Transport Aerien,\N,,MDE,MAURI-TRANS,Mauritania,N
3427,Maverick Airways,\N,,MVR,MAV-AIR,United States,N
3428,Mavial Magadan Airlines,\N,H5,MVL,Mavial,Russia,N
3429,Max Avia,\N,,MAI,MAX AVIA,Kyrgyzstan,N
3430,Max Sea Food,\N,,MSF,MAXESA,El Salvador,N
3431,Max-Aviation,\N,,MAX,MAX AVIATION,Canada,N
3432,Maxair,\N,8M,MXL,MAXAIR,Sweden,Y
3433,Maximus Air Cargo,\N,,MXU,CARGO MAX,United Arab Emirates,N
3434,Maxjet Airways,\N,MY,MXJ,MAX-JET,United States,N
3435,Maxsus-Avia,\N,,MXS,MAXSUS-AVIA,Uzbekistan,N
3436,May Air Xpress,\N,,MXP,BEECHNUT,United States,N
3437,Maya Island Air,\N,MW,MYD,MYLAND,Belize,Y
3438,Mayair,\N,,MYI,MAYAIR,Mexico,N
3439,Mbach Air,\N,,MBS,MBACHI AIR,Malawi,N
3440,McAlpine Helicopters,\N,,MCH,MACLINE,United Kingdom,N
3441,McCall Aviation,\N,,MKL,MCCALL,United States,N
3442,McDonnell Douglas,\N,,DAC,DACO,United States,N
3443,McNeely Charter Services,\N,,MDS,MID-SOUTH,United States,N
3444,Med-Trans of Florida,\N,,MEK,MED-TRANS,United States,N
3445,Medavia Company,\N,,MDM,MEDAVIA,Malta,N
3446,Medical Air Rescue Services,\N,,MRZ,MARS,Zimbabwe,N
3447,Medical Aviation Services,\N,,MCL,MEDIC,United Kingdom,N
3448,Mediterranean Air Freight,\N,,MDF,MED-FREIGHT,Greece,N
3449,Mediterranean Airways,\N,,MDY,,Egypt,N
3450,Medjet International,\N,,MEJ,MEDJET,United States,N
3451,Mega,\N,,MGK,MEGLA,Kazakhstan,N
3452,Mega Linhas Aereas,\N,,MEL,MEGA AIR,Brazil,N
3453,Mekong Airlines,\N,,MKN,MEKONG AIRLINES,Cambodia,N
3454,Menajet,\N,IM,MNJ,MENAJET,Lebanon,N
3456,Merchant Express Aviation,\N,,MXX,MERCHANT,Nigeria,N
3457,Mercury Aircourier Service,\N,,MEC,MERCAIR,United States,N
3458,Meridian,\N,,POV,AIR POLTAVA,Ukraine,N
3459,Meridian Air Cargo,\N,,MRD,MERIDIAN,United States,N
3460,Meridian Airlines,\N,,MHL,HASSIMAIR,Nigeria,N
3461,Meridian Aviation,\N,,DSL,DIESEL,United Kingdom,N
3462,Meridian Limited,\N,,MEM,MERIDIAN CHERRY,Ukraine,N
3463,Meridiana,\N,IG,ISS,MERAIR,Italy,Y
3464,Merlin Airways,\N,,MEI,AVALON,United States,N
3465,Merpati Nusantara Airlines,\N,MZ,MNA,MERPATI,Indonesia,Y
3466,Mesa Airlines,\N,YV,ASH,AIR SHUTTLE,United States,Y
3467,Mesaba Airlines,\N,XJ,MES,MESABA,United States,Y
3468,Meta Linhas A,\N,,MSQ,META,Brazil,N
3469,Meteorological Research Flight,\N,,MET,METMAN,United Kingdom,N
3470,Methow Aviation,\N,,MER,METHOW,United States,N
3471,Metro Business Aviation,\N,,MVI,,United Kingdom,N
3472,Metro Express,\N,,MEX,EAGLE EXPRESS,United States,N
3473,Metroflight,\N,,MTR,METRO,United States,N
3474,Metrojet,\N,,MTJ,METROJET,Hong Kong SAR of China,N
3475,Metropix UK,\N,,PIX,METROPIX,United Kingdom,N
3476,Metroplis,\N,,MPS,METRO REGIONAL,Netherlands,N
3477,Mex Blue,\N,,MXB,MEX BLUE,Mexico,N
3478,Mex-Jet,\N,,MJT,MEJETS,Mexico,N
3479,Mexicana de Aviaci,\N,MX,MXA,MEXICANA,Mexico,Y
3480,Mexico Transportes Aereos,\N,,MXT,TRANSMEX,Mexico,N
3481,Miami Air Charter,\N,,HUR,HURRICANE CHARTER,United States,N
3482,Miami Air International,\N,GL,BSK,BISCAYNE,United States,N
3483,Miami Valley Aviation,\N,,OWL,NIGHT OWL,United States,N
3484,Miapet-Avia,\N,,MPT,MIAPET,Armenia,N
3485,Michelin Air Services,\N,,BIB,,France,N
3486,Micromatter Technology Solutions,\N,,WIZ,WIZARD,United Kingdom,N
3487,Mid Airlines,\N,,NYL,NILE,Sudan,N
3488,Mid-Pacific Airlines,\N,,MPA,MID PAC,United States,N
3489,Midamerica Jet,\N,,MJR,MAJOR,United States,N
3490,Middle East Airlines,\N,ME,MEA,CEDAR JET,Lebanon,Y
3491,Midland Airport Services,\N,,MID,,United Kingdom,N
3492,Midline Air Freight,\N,,MFR,MIDLINE FREIGHT,United States,N
3493,Midstate Airlines,\N,,MIS,MIDSTATE,United States,N
3494,Midway Airlines,\N,JI,MDW,MIDWAY,United States,Y
3495,Midway Express,\N,,FLA,PALM,United States,N
3496,Midwest Air Freighters,\N,,FAX,FAIRFAX,United States,N
3497,Midwest Airlines,\N,YX,MEP,,United States,Y
3498,Midwest Airlines (Egypt),\N,MY,MWA,,Egypt,Y
3499,Midwest Aviation,\N,,NIT,NIGHTTRAIN,United States,N
3500,Midwest Aviation Division,\N,,MWT,MIDWEST,United States,N
3501,Midwest Helicopters De Mexico,\N,,HTE,HELICOPTERSMEXICO,Mexico,N
3502,Millardair,\N,,MAB,MILLARDAIR,Canada,N
3503,Millen Corporation,\N,,RJM,MILLEN,United Kingdom,N
3504,Millennium Air,\N,,MLK,NIGERJET,Nigeria,N
3505,Miller Flying Services,\N,,MFS,MILLER TIME,United States,N
3506,Million Air,\N,,OXO,MILL AIR,United States,N
3507,Mimino,\N,,MIM,MIMINO,Russia,N
3508,Mina Airline Company,\N,,NAB,,Egypt,N
3509,Minair,\N,,OMR,ORMINE,Central African Republic,N
3510,Minebea Technologies,\N,,EBE,MINEBEA,United States,N
3511,Mines Air Services Zambia,\N,,MAZ,MINES,Zambia,N
3512,Miniliner,\N,,MNL,MINILINER,Italy,N
3513,Ministic Air,\N,,MNS,MINISTIC,Canada,N
3514,Ministry of Agriculture,\N,,WDG, Fisheries and Food,WATCHDOG,N
3515,Minsk Aircraft Overhaul Plant,\N,,LIR,LISLINE,Belarus,N
3516,Miramichi Air Service,\N,,MIR,MIRAMICHI,Canada,N
3517,Miras,\N,,MIF,MIRAS,Kazakhstan,N
3518,Misr Overseas Airways,\N,,MOS,,Egypt,N
3519,Mission Aviation Fellowship,\N,,MAF,MISSI,Indonesia,N
3520,Missionair,\N,,MSN,MISIONAIR,Spain,N
3521,Missions Gouvernemtales Francaises,\N,,MRN,MARIANNE,France,N
3522,Mississippi State University,\N,,BDG,BULLDOG,United States,N
3523,Mississippi Valley Airways,\N,,MVA,VALAIR,United States,N
3524,Mistral Air,\N,,MSA,AIRMERCI,Italy,N
3525,Mobil Oil,\N,,MBO,MOBIL,Canada,N
3526,Mocambique Expresso,\N,,MXE,MOZAMBIQUE EXPRESS,Mozambique,N
3527,Mofaz Air,\N,,MFZ,MOFAZ AIR,Malaysia,N
3528,Moldaeroservice,\N,,MLE,MOLDAERO,Moldova,N
3529,Moldavian Airlines,\N,2M,MDV,MOLDAVIAN,Moldova,Y
3530,Moldova,\N,,MVG,MOLDOVA-STATE,Moldova,N
3531,Mombasa Air Safari,\N,,RRV,SKYROVER,Kenya,N
3532,Monarch Airlines,\N,ZB,MON,MONARCH,United Kingdom,Y
3533,Monarch Airlines,\N,,MNH,MONARCH AIR,United States,N
3534,Myway Airlines,\N,8I,,,Italy,Y
3535,Moncton Flying Club,\N,,MFC,EAST WIND,Canada,N
3536,Monde Air Charters,\N,,MDB,MONDEAIR CARGO,United Kingdom,N
3537,Monerrey Air Taxi,\N,,MTI,MONTERREY AIR,Mexico,N
3538,Monky Aerotaxis,\N,,MKY,MONKY,Mexico,N
3539,Montenegro Airlines,\N,YM,MGX,MONTAIR,Montenegro,Y
3540,Montserrat Airways,\N,,MNT,MONTSERRAT,Montserrat,N
3541,Mooney Aircraft Corporation,\N,,MNY,MOONEY FLIGHT,United States,N
3542,Morningstar Air Express,\N,,MAL,MORNINGSTAR,Canada,Y
3543,Morris Air Service,\N,,MSS,WASATCH,United States,N
3544,Morrison Flying Service,\N,,MRO,MORRISON,United States,N
3545,Moskovia Airlines,\N,3R,GAI,GROMOV AIRLINE,Russia,Y
3546,Mosphil Aero,\N,,MPI,MOSPHIL,Philippines,N
3547,Motor Sich,\N,M9,MSI,MOTOR SICH,Ukraine,Y
3548,Mount Cook Airlines,\N,NM,NZM,MOUNTCOOK,New Zealand,N
3549,Mountain Air Cargo,\N,,MTN,MOUNTAIN,United States,N
3550,Mountain Air Company,\N,N4,MTC,MOUNTAIN LEONE,Sierra Leone,N
3551,Mountain Air Express,\N,,PKP,PIKES PEAK,United States,N
3552,Mountain Air Service,\N,,BRR,MOUNTAIN AIR,United States,N
3553,Mountain Bird,\N,,MBI,MOUNTAIN BIRD,United States,N
3554,Mountain High Aviation,\N,,MHA,MOUNTAIN HIGH,United States,N
3555,Mountain Pacific Air,\N,,MPC,MOUNTAIN PACIFIC,Canada,N
3556,Mountain Valley Air Service,\N,,MTV,MOUNTAIN VALLEY,United States,N
3557,Mowhawk Airlines,\N,,MOW,MOHAWK AIR,United States,N
3558,Mudan Airlines,\N,,MDN,,Somali Republic,N
3559,Mudanjiang General Aviation,\N,,CMJ,MUDANJIANG,China,N
3560,Multi Taxi,\N,,MTX,MULTITAXI,Mexico,N
3561,Multi-Aero,\N,,WBR,WEBER,United States,N
3562,Multiflight,\N,,MFT,YORKAIR,,N
3563,Murmansk Aircompany,\N,,MNZ,MURMAN AIR,Russia,N
3564,Murray Air,\N,,MUA,MURRAY AIR,United States,N
3565,Musrata Air Transport,\N,,MMR,MUSRATA AIR,Libya,N
3566,Mustique Airways,\N,,MAW,MUSTIQUE,Barbados,N
3567,My Way Airlines,\N,,MYW,FRANKY,Italy,N
3568,MyTravel Airways,\N,VZ,MYT,KESTREL,United Kingdom,Y
3569,Myanma Airways,\N,UB,UBA,UNIONAIR,Myanmar,Y
3570,Myanmar Airways International,\N,8M,MMM,assignment postponed,Myanmar,Y
3571,Myflug,\N,,MYA,MYFLUG,Iceland,Y
3572,Mytravel Airways,\N,,VKG,VIKING,Denmark,N
3573,NEL Cargo,\N,,NLG,NELCARGO,Ivory Coast,N
3574,NHT Lineas Aereas,\N,,NHG,HELGA,Brazil,N
3575,NZ Warbirds Association,\N,,WAR,WARBIRDS,New Zealand,N
3576,Nacoia Lda,\N,,ANL,AIR NACOIA,Angola,N
3577,Nada Air Service,\N,,NHZ,NADA AIR,Chad,N
3578,Compangnie Nationale Naganagani,\N,,BFN,,Burkina Faso,N
3579,Nahanni Air Services Ltd,\N,,NAH,NAHANNI,Canada,N
3580,Nakheel Aviation,\N,,NKL,NAKHEEL,United Arab Emirates,N
3581,Namibia Commerical Aviation,\N,,MRE,MED RESCUE,Namibia,N
3582,Namibian Defence Force,\N,,NDF,NAMIBIAN AIR FORCE,Namibia,N
3583,Nanjing Airlines,\N,,CNJ,NINGHANG,China,N
3584,Nantucket Airlines,\N,DV,ACK,ACK AIR,United States,N
3585,Nanyah Aviation,\N,,NYA,NANYAH,Israel,N
3586,Napier Air Service Inc,\N,,NAP,NAPIER,United States,N
3587,Nas Air,\N,,NCM,AIR BANE,Angola,N
3588,Nas Air,\N,P9,,,Mali,N
3589,Nasair,\N,UE,NAS,NASAIRWAYS,Eritrea,Y
3590,Nashville Jet Charters,\N,,NJC,NASHVILLE JET,United States,N
3591,Natalco Air Lines,\N,,NCO,NATALCO,S,N
3592,Natioanl Air Traffic Controllers Association,\N,,NTK,NATCA,United States,N
3593,National Air Charter,\N,,NSR,NASAIR,Indonesia,N
3594,National Air Traffic Services,\N,,RFI,SHERLOCK,United Kingdom,N
3595,National Airlines,\N,,NAN,NATION AIR,United States,N
3596,National Airlines,\N,N4,NCN,,Chile,N
3597,National Airlines,\N,N7,ROK,RED ROCK,United States,N
3598,National Airlines,\N,NA,NAL,NATIONAL,United States,N
3599,National Airlines,\N,,KUS,KUSWAG,South Africa,N
3600,National Airways Cameroon,\N,9O,,,Cameroon,N
3601,National Airways Corporation,\N,,LFI,AEROMED,South Africa,N
3602,National Aviation Company,\N,,GTY,,Egypt,N
3603,National Aviation Consultants,\N,,TNC,NATCOM,Canada,N
3604,National Express,\N,,NXT,NATIONAL FREIGHT,United States,N
3605,National Grid,\N,,GRD,GRID,United Kingdom,N
3606,National Jet Express,\N,,JTE,JETEX,Australia,N
3607,National Jet Service,\N,,AND,AIR INDIANA,United States,N
3608,National Jet Systems,\N,NC,NJS,NATIONAL JET,Australia,Y
3609,National Oceanic and Atmospheric Administration,\N,,NAA,NOAA,United States,N
3610,National Overseas Airlines Company,\N,,NOL,NAT AIRLINE,Egypt,N
3611,Nationale Luchtvaartschool,\N,,NLS,PANDER,Netherlands,N
3612,Nations Air Express Inc,\N,,NAE,NATIONS EXPRESS,United States,N
3613,Nationwide Airlines,\N,CE,NTW,NATIONWIDE,South Africa,Y
3614,Nationwide Airlines (Zambia),\N,,NWZ,ZAMNAT,Zambia,N
3615,Natural Environment Research Council,\N,,EVM,SCIENCE,United Kingdom,N
3616,Natureair,\N,,NRR,NATUREAIR,Costa Rica,N
3617,Naturelink Charter,\N,,NRK,NATURELINK,South Africa,N
3618,Nauru Air Corporation,\N,ON,RON,AIR NAURU,Nauru,Y
3619,NAV CANADA,\N,,NVC,NAV CAN,Canada,N
3620,Nav Flight Planning,\N,,NAV,NAV DISPATCH,Czech Republic,N
3621,Navegacao Aerea De Portugal,\N,,NVP,,Portugal,N
3622,Navegacion Servicios Aereos Canarios SA,\N,,NAY,NAYSA,Spain,N
3623,Navid,\N,,IRI,NAVID,Iran,N
3624,Naviera Mexicana,\N,,NVM,NAVIERA,Mexico,N
3625,Navigator Airlines,\N,,NVL,NAVLINES,Armenia,N
3626,Navinc Airlines Services,\N,,XNV,,United States,N
3627,Navitaire,\N,1N,,,United States,N
3628,Navtech System Support,\N,,XNS,,Canada,N
3629,Nayzak Air Transport,\N,,NZA,,Libya,N
3630,State of Nebraska,\N,,NEB,NEBRASKA,United States,N
3631,Necon Air,\N,,NEC,NECON AIR,Nepal,N
3632,Nederlandse Kustwacht,\N,,NCG,NETHERLANDS COASTGUARD,Netherlands,N
3633,Nefteyugansk Aviation Division,\N,,NFT,NEFTEAVIA,Russia,N
3634,Neiltown Air,\N,,NLA,NEILTOWN AIR,Canada,N
3635,Nelair Charters,\N,,NLC,NELAIR,South Africa,N
3636,Nelson Aviation College,\N,,CGE,COLLEGE,New Zealand,N
3637,Nepal Airlines,\N,RA,RNA,ROYAL NEPAL,Nepal,Y
3638,Neos,\N,NO,NOS,MOONFLOWER,Italy,N
3639,Neosiam Airways,\N,,TOX,SKY KINGDOM,Thailand,N
3640,Neric,\N,,NSL,NERICAIR,United Kingdom,N
3641,NetJets,\N,1I,EJA,EXECJET,United States,Y
3642,Network Aviation Services,\N,,NET,NETWORK,Nigeria,N
3643,New England Air Express,\N,,NEZ,ENGAIR,United States,N
3644,New England Airlines,\N,EJ,NEA,NEW ENGLAND,United States,Y
3645,New Heights 291,\N,,NHT,NEWHEIGHTS,South Africa,N
3646,New World Jet Corporation,\N,,NWD,NEW WORLD,United States,N
3647,New York Helicopter,\N,,NYH,NEW YORK,United States,N
3648,New York State Police,\N,,GRY,GRAY RIDER,United States,N
3649,New Zealand Air Defence Force,\N,,KRC,KIWI RESCUE,New Zealand,N
3650,Newair,\N,,HVA,HAVEN-AIR,United States,N
3651,Newfoundland Labrador Air Transport,\N,,NLT,NALAIR,Canada,N
3652,NextJet,\N,2N,NTJ,NEXTJET,Sweden,Y
3653,Nextflight Aviation,\N,,NXF,NEXTFLIGHT,United States,N
3654,Nexus Aviation,\N,,NXS,NEXUS AVIATION,Nigeria,N
3655,Nicaraguense De Aviacion,\N,,NIS,NICA,Nicaragua,N
3656,Nicon Airways,\N,,NCN,NICON AIRWAYS,Nigeria,N
3657,Nigeria Airways,\N,,NGA,NIGERIA,Nigeria,N
3658,Nigerian Air Force,\N,,NGR,NIGERIAN AIRFORCE,Nigeria,N
3659,Nigerian Global,\N,,NGX,AIR GLOBAL,Nigeria,N
3660,Night Express,\N,,EXT,EXECUTIVE,Germany,N
3661,Niki,\N,HG,NLY,FLYNIKI,Austria,Y
3662,Nikolaev-Air,\N,,NKV,AIR NIKOLAEV,Ukraine,N
3663,Nile Safaris Aviation,\N,,NSA,NILE SAFARIS,Sudan,N
3664,Nile Valley Aviation Company,\N,,NVA,,Egypt,N
3665,Nile Wings Aviation Services,\N,,NLW,NILE WINGS,Sudan,N
3666,Nimbus Aviation,\N,,NBS,NIMBUS,United Kingdom,N
3667,Nippon Cargo Airlines,\N,KZ,NCA,NIPPON CARGO,Japan,N
3668,Nizhnevartovskavia,\N,,NVK,VARTOSKAVIA,Russia,N
3669,Search and Rescue 202,\N,,SRG,,United Kingdom,N
3670,Search and Rescue 22,\N,,SRD,,United Kingdom,N
3671,No. 32 (The Royal) Squadron RAF,\N,,NOH,NORTHOLT,United Kingdom,N
3672,84 Squadron Royal Air Force @ RAF Akrotiri,\N,,AKG,GRIFTER,United Kingdom,N
3673,Nobil Air,\N,,NBL,NOBIL AIR,Moldova,N
3674,Nok Air,\N,DD,NOK,NOK AIR,Thailand,Y
3675,Nolinor Aviation,\N,,NRL,NOLINOR,Canada,N
3676,Nomad Aviation,\N,,NMD,NOMAD AIR,Namibia,N
3677,Norcopter,\N,,NOC,,Norway,N
3678,Nord-Flyg,\N,,NEF,NORDEX,Sweden,N
3679,Nordeste Linhas Aereas Regionais,\N,JH,NES,NORDESTE,Brazil,N
3680,Nordic Regional,\N,6N,NRD,NORTH RIDER,Sweden,N
3681,Nordic Solutions,\N,,NVD,NORDVIND,Lithuania,N
3682,Nordstree (Australia),\N,,NDS,,Australia,N
3683,Norestair,\N,,NRT,NORESTAIR,Spain,N
3684,Norfolk County Flight College,\N,,NCF,COUNTY,United Kingdom,Y
3685,Norontair,\N,,NOA,NORONTAIR,Canada,N
3686,Norrlandsflyg,\N,,HMF,LIFEGUARD SWEDEN,Sweden,N
3687,Norse Air Charter,\N,,NRX,NORSE AIR,South Africa,N
3688,Norsk Flyrjeneste,\N,,NIR,NORSEMAN,Norway,N
3689,Norsk Helikopter,\N,,NOR,NORSKE,Norway,N
3690,Norsk Luftambulanse,\N,,DOC,HELIDOC,Norway,N
3691,Nortavia,\N,,RTV,TIC-TAC,Portugal,N
3692,North Adria Aviation,\N,,NAI,NORTH-ADRIA,Croatia,N
3693,North American Airlines,\N,,NTM,NORTHAM,Canada,Y
3694,North American Charters,\N,,HMR,HAMMER,Canada,Y
3695,North American Jet Charter Group,\N,,NAJ,JET GROUP,United States,N
3696,North Atlantic Air Inc,\N,,NAT,MASS AIR,United States,N
3697,North Atlantic Cargo,\N,,NFC,NORTH ATLANTIC,Norway,N
3698,North British Airlines,\N,,NBN,TEESAIR,United Kingdom,N
3699,North Caribou Flying Service Ltd,\N,,NCB,NORTH CARIBOU,Canada,N
3700,North Coast Air Services Ltd,\N,,NCC,NORTH COAST,Canada,N
3701,North Coast Aviation,\N,N9,,,Papua New Guinea,N
3702,North Flying,\N,M3,NFA,NORTH FLYING,Denmark,N
3703,North Sea Airways,\N,,NRC,NORTH SEA,Netherlands,N
3704,North Star Air Cargo,\N,,SBX,SKY BOX,United States,N
3705,North Vancouver Airlines,\N,,NRV,NORVAN,Canada,N
3706,North West Airlines,\N,,NWW,HALANT,Australia,N
3707,North West Geomatics,\N,,PTO,PHOTO,Canada,N
3708,North-East Airlines,\N,,NEN,NORTHEAST SWAN,Nigeria,N
3709,North-West Air Transport Company - Vyborg,\N,,VBG,VYBORG AIR,Russia,N
3710,North-Wright Airways,\N,HW,NWL,NORTHWRIGHT,Canada,N
3711,Northafrican Air Transport,\N,,NLL,NORTHAFRICAN AIR,Libya,N
3712,Northaire Feight Lines,\N,,NFL,GREAT LAKES,United States,N
3713,Northamptonshire School of Flying,\N,,NSF,NORTON,United Kingdom,N
3714,Northcoast Executive Airlines,\N,,NCE,TOP HAT,United States,N
3715,Northeast Airlines,\N,,NEE,NORTHEAST,United States,N
3716,Northeast Aviation,\N,,NPX,NORTHEAST EXPRESS,United States,N
3717,Northern Air Cargo,\N,NC,NAC,YUKON,United States,N
3718,Northern Airlines Sanya,\N,,BYC,BEIYA,China,N
3719,Northern Airways,\N,,NDA,NORTHERN DAKOTA,United States,N
3720,Northern Aviation Service,\N,,CMU,LANNA AIR,Thailand,N
3721,Northern Dene Airways,\N,U7,,,Canada,Y
3722,Northern Executive Aviation,\N,,NEX,NEATAX,United Kingdom,N
3723,Northern Illinois Commuter,\N,,NIC,ILLINOIS COMMUTER,United States,N
3724,Northern Jet Management,\N,,NTX,NORTAX,United States,N
3725,Northern Thunderbird Air,\N,,NTA,THUNDERBIRD,Canada,N
3726,Northland Aviation,\N,,KOE,KOKEE,United States,N
3727,Northstar Aviation,\N,,NSS,NORTHSTAR,United States,N
3728,Northumbria Helicopters,\N,,NHL,NORTHUMBRIA,United Kingdom,N
3729,Northway Aviation Ltd,\N,,NAL,NORTHWAY,Canada,N
3730,Northwest Aero Associates,\N,,NWE,,United States,N
3731,Northwest Airlines,\N,NW,NWA,NORTHWEST,United States,Y
3732,Northwest Regional Airlines,\N,FY,NWR,,Australia,N
3733,Northwest Territorial Airways,\N,,NWT,TERRITORIAL,Canada,N
3734,Northwestern Air,\N,J3,PLR,POLARIS,Canada,Y
3735,Northwinds Northern,\N,,NWN,NORTHWINDS,Canada,N
3736,Nortland Air Manitoba,\N,,NAM,MANITOBA,Canada,N
3737,Norwegian Air Shuttle,\N,DY,NAX,NOR SHUTTLE,Norway,Y
3738,Norwegian Aviation College,\N,,TFN,SPRIT,Norway,Y
3739,Notams International,\N,,XNT,,United States,N
3740,Nouvel Air Tunisie,\N,BJ,LBT,NOUVELAIR,Tunisia,Y
3741,Nova Airline,\N,M4,NOV,NOVANILE,Sudan,N
3742,Nova Scotia Department of Lands and Forests,\N,,PTR,PATROL,Canada,N
3743,Novair,\N,1I,NVR,NAVIGATOR,Sweden,Y
3744,Novogorod Air Enterprise,\N,,NVG,SADKO AVIA,Russia,N
3745,Novosibirsk Aircraft Repairing Plant,\N,,NSP,NARPAIR,Russia,N
3746,Novosibirsk Aviaenterprise,\N,,NBE,NAKAIR,Russia,N
3747,Novosibirsk Aviation Production Association,\N,,NPO,NOVSIB,Russia,N
3748,Noy Aviation,\N,,NOY,NOY AVIATION,Israel,N
3749,Nuevo Continente,\N,N6,ACQ,AERO CONTINENTE,Peru,N
3750,Nuevo Horizonte Internacional,\N,,NHR,NUEVO HORIZONTE,Mexico,N
3751,Nunasi-Central Airlines,\N,,NUN,NUNASI,Canada,N
3752,Nurman Avia Indopura,\N,,NIN,NURVINDO,Indonesia,N
3753,Nyasa Express,\N,,NYS,NYASA,Malawi,N
3754,Nas Air,\N,XY,KNE,NAS EXPRESS,Saudi Arabia,Y
3755,O Air,\N,,OCN,O-BIRD,France,N
3756,O'Connor Airlines,\N,UQ,OCM,OCONNOR,Australia,N
3757,OAG,\N,CR,,,United Kingdom,N
3758,OSACOM,\N,,JPA,J-PAT,United States,N
3759,Oasis Hong Kong Airlines,\N,O8,OHK,OASIS,Hong Kong,Y
3760,Ocean Air,\N,,BCN,BLUE OCEAN,Mauritania,Y
3761,Ocean Airlines,\N,VC,VCX,OCEANCARGO,Italy,N
3762,Ocean Sky (UK),\N,,OCS,OCEANSKY,United Kingdom,N
3763,Ocean Wings Commuter Service,\N,,TUK,TUCKERNUCK,United States,N
3764,Oceanair,\N,O6,ONE,OCEANAIR,Brazil,Y
3765,Oceanic Airlines,\N,O2,,,Guinea,Y
3766,Odessa Airlines,\N,,ODS,ODESSA AIR,Ukraine,N
3767,Odyssey International,\N,,ODY,ODYSSEY,Canada,N
3768,Office Federal De'Aviation Civile,\N,,FOC,FOCA,Switzerland,N
3769,Ogooue Air Cargo,\N,,GBO,,Gabon,N
3770,Okada Airlines,\N,,OKJ,OKADA AIR,Nigeria,N
3771,Okapi Airways,\N,,OKP,OKAPI,Democratic Republic of Congo,N
3772,Okay Airways,\N,,OKA,OKAYJET,China,N
3773,Oklahoma Department of Public Safety,\N,,OKL,OKLAHOMA,United States,N
3774,Olimex Aerotaxi,\N,,OLX,OLIMEX,Czech Republic,N
3775,Olimp Air,\N,,KVK,PONTA,Kazakhstan,N
3776,Olympic Airlines,\N,OA,OAL,OLYMPIC,Greece,Y
3777,Olympic Aviation,\N,,OLY,OLAVIA,Greece,N
3778,Oman Air,\N,WY,OMA,OMAN AIR,Oman,Y
3779,Oman Royal Flight,\N,,ORF,OMAN,Oman,N
3780,Omni - Aviacao e Tecnologia,\N,,OAV,OMNI,Portugal,N
3781,Omni Air International,\N,OY,OAE,OMNI-EXPRESS,United States,Y
3782,Omniflys,\N,,OMF,OMNIFLYS,Mexico,N
3783,Omskavia Airline,\N,N3,OMS,OMSK,Russia,N
3784,On Air Limited,\N,,ORL,ON AIR,Canada,N
3785,One Two Go Airlines,\N,,OTG,THAI EXPRESS,Thailand,Y
3786,Onetime Airlines Zambia,\N,,OTM,ZEDTIME,Zambia,N
3787,Ontario Ministry of Health,\N,,MED,MEDICAL,Canada,N
3788,Onur Air,\N,8Q,OHY,ONUR AIR,Turkey,Y
3789,Opal Air,\N,,OPA,,Australia,N
3790,Open Sky Aviation,\N,,OSA,,Lebanon,N
3791,OpenSkies,\N,,BOS,MISTRAL,United Kingdom,N
3792,Operadora Turistica Aurora,\N,,ORR,TURISTICA AURORA,Mexico,N
3793,Operadora de Lineas Ejecutivas,\N,,OLE,OPERADORA,Mexico,N
3794,Operadora de Transportes Aereos,\N,,OTP,OPERADORA AEREO,Mexico,N
3795,Operadora de Veulos Ejectutivos,\N,,OPV,OPERADORA DE VUELOS,Mexico,N
3796,Operation Enduring Freedom,\N,,LLO,APOLLO,Canada,N
3797,Operational Aviation Services,\N,,OAX,,Australia,N
3798,Orange Air Services,\N,,ORD,ORANGE SERVICES,Sierra Leone,N
3799,Orange Air Sierra Leone,\N,,ORJ,ORANGE SIERRA,Sierra Leone,N
3800,Orange Aviation,\N,,ORE,ORANGE AVIATION,Israel,N
3801,Orbit Express Airlines,\N,,ORX,OREX,Turkey,N
3802,Orca Air,\N,,ORK,ORCA TAXI,Egypt,N
3803,Orebro Aviation,\N,,BUE,BLUELIGHT,Sweden,N
3804,Orel State Air Enterprise,\N,,ORM,ORPRISE,Russia,N
3805,Orenburg Airlines,\N,R2,ORB,ORENBURG,Russia,Y
3806,Organizacion De Transportes Aereos,\N,,OTA,ORGANIZACION,Mexico,N
3807,Organizacoes Mambra,\N,,OML,MAMBRA,Angola,N
3808,Orient Air,\N,,OVV,ORIENTSYR,Syrian Arab Republic,N
3809,Orient Airlines,\N,,OTR,ORIENTROC,Sudan,N
3810,Orient Airways,\N,,ORN,ORIENT LINER,Pakistan,N
3811,Orient Thai Airlines,\N,OX,OEA,ORIENT THAI,Thailand,Y
3812,Oriental Air Bridge,\N,,NGK,ORIENTAL BRIDGE,Japan,N
3813,Oriental Airlines,\N,,OAC,ORIENTAL AIR,Nigeria,N
3814,Origin Pacific Airways,\N,QO,OGN,ORIGIN,New Zealand,Y
3815,Orion Air Charter,\N,,OED,ORION CHARTER,South Africa,N
3816,Orion-x,\N,,OIX,ORIONIX,Russia,N
3817,Orlan-2000,\N,,KOV,ORLAN,Kazakhstan,N
3818,Ornage Aircraft Leasing,\N,,RNG,ORANGE,Netherlands,N
3819,Orscom Tourist Installations Company,\N,,OAD,ORSCOM,Egypt,N
3820,Osh Avia,\N,,OSH,OSH AVIA,Kyrgyzstan,N
3821,Ostend Air College,\N,,OCO,AIR COLLEGE,Belgium,N
3822,Ostfriesische Lufttransport,\N,OL,OLT,OLTRA,Germany,Y
3823,Oulun Tilauslento,\N,,FNL,FINN FLIGHT,Finland,N
3824,Our Airline,\N,ON,RON,OUR AIRLINE,Nauru,N
3825,Out Of The Blue Air Safaris,\N,,OOT,OOTBAS,South Africa,N
3826,Overland Airways,\N,OJ,OLA,OVERLAND,Nigeria,Y
3827,Oxaero,\N,,OXE,OXOE,United Kingdom,N
3828,Oxford Air Services,\N,,WDK,WOODSTOCK,United Kingdom,N
3829,Oxley Aviation,\N,,OAA,,Australia,N
3830,Ozark Air Lines,\N,OZ,OZR,OZARK,United States,N
3831,Ozjet Airlines,\N,O7,OZJ,AUSJET,Australia,Y
3832,P & P Floss Pick Manufacturers,\N,,KTL,KNOTTSBERRY,South Africa,N
3833,PAC Air,\N,,PCR,PACAIR,United States,N
3834,PAN Air,\N,PV,PNR,SKYJET,Spain,Y
3835,PB Air,\N,9Q,PBA,PEEBEE AIR,Thailand,Y
3836,PDQ Air Charter,\N,,PDQ,DISPATCH,United States,N
3837,PHH Aviation System,\N,,XAS,,United States,N
3838,PLM Dollar Group,\N,,PDG,OSPREY,United Kingdom,N
3839,PLUNA,\N,PU,PUA,PLUNA,Uruguay,Y
3840,PMTair,\N,U4,PMT,MULTITRADE,Cambodia,Y
3841,PRT Aviation,\N,,PRP,PRONTO,Spain,N
3842,PSA Airlines,\N,,JIA,BLUE STREAK,United States,N
3843,PTL Luftfahrtunternehemen,\N,,KST,KING STAR,Germany,N
3844,Paccair,\N,,WIS,WISCAIR,United States,N
3845,Pace Airlines,\N,Y5,PCE,PACE,United States,N
3846,Pacific Air Boats,\N,,PAB,AIR BOATS,Canada,N
3847,Pacific Air Charter,\N,,PRC,PACIFIC CHARTER,United States,N
3848,Pacific Air Express,\N,,PCF,PACIFIC EXPRESS,United States,N
3849,Pacific Air Transport,\N,,PXP,PAK EXPRESS,United States,N
3850,Jetstar Pacific,Pacific Airlines,BL,PIC,PACIFIC AIRLINES,Vietnam,Y
3851,Pacific Alaska Airlines,\N,,PAK,PACIFIC ALASKA,United States,N
3852,Pacific Aviation (Australia),\N,,PCV,PACAV,Australia,N
3853,Pacific Aviation (United States),\N,,PCX,,United States,N
3854,Pacific Blue,\N,DJ,PBN,BLUEBIRD,New Zealand,N
3855,Pacific Coast Airlines,\N,,PQA,SAGE BRUSH,United States,N
3856,Pacific Coastal Airline,\N,8P,PCO,PASCO,Canada,Y
3857,Pacific East Asia Cargo Airlines,\N,Q8,PEC,PAC-EAST CARGO,Philippines,Y
3858,Pacific Flight Services,\N,,PFA,PACIFIC SING,Singapore,N
3859,Pacific International Airlines,\N,,PIN,ROAD RUNNERS,United States,N
3860,Pacific Island Aviation,\N,,PSA,PACIFIC ISLE,United States,Y
3861,Pacific Jet,\N,,PCJ,PACIFIC JET,United States,N
3862,Pacific Pearl Airways,\N,,PPM,PACIFIC PEARL,Philippines,N
3863,Pacific Rim Airways,\N,,PAR,PACRIM,Australia,N
3864,Pacific Southwest Airlines,\N,PS,PSX,SMILEY,United States,N
3865,Pacific Wings,\N,LW,NMI,TSUNAMI,United States,Y
3866,Pacificair,\N,GX,,,Philippines,N
3867,Pacificair Airlines,\N,,PFR,PACIFIC WEST,United States,N
3868,Package Express,\N,,RCY,RACE CITY,United States,N
3869,Paisajes Espanoles,\N,,PAE,PAISAJES,Spain,N
3870,Pak West Airlines,\N,,PKW,PLATINUM WEST,United States,N
3871,Pakistan International Airlines,PIA Pakistan International,PK,PIA,PAKISTAN,Pakistan,Y
3872,Pakker Avio,\N,,PKR,PAKKER AVIO,Estonia,N
3873,Pal Aerolineas,\N,,LPA,LINEASPAL,Mexico,N
3874,Palau Asia Pacific Airlines,\N,,PPC,PALAU ASIAPAC,Palau,N
3875,Palau Trans Pacific Airline,\N,GP,PTP,TRANS PACIFIC,Palau,N
3876,Palestinian Airlines,\N,PF,PNW,PALESTINIAN,Egypt,N
3877,Palmer Aviation,\N,,JSP,PALMER,United Kingdom,N
3878,Pamir Airways,\N,NR,PIR,PAMIR,Afghanistan,N
3879,Pan African Air Services,\N,,PFN,PANAFRICAN,Sierra Leone,N
3880,Pan African Airways,\N,,ODM,,Kenya,N
3881,Pan Air,\N,,PAX,PANNEX,United States,N
3882,Pan Am Weather Systems,\N,,XPA,,United States,N
3883,Pan Am World Airways,\N,,PWD,,Dominican Republic,N
3885,Pan American Airways,\N,PA,PAA,,United States,N
3886,Pan American World Airways,\N,PA,PAA,CLIPPER,United States,N
3887,Pan Europeenne Air Service,\N,,PEA,,France,N
3888,Pan Havacilik Ve Ticaret,\N,,PHT,PANANK,Turkey,N
3889,Pan Malaysian Air Transport,\N,,PMA,PAN MALAYSIA,Malaysia,N
3890,Pan-Air,\N,,PNC,PANAIRSA,Mexico,N
3891,Panafrican Airways,\N,PQ,PNF,PANWAYS,Ivory Coast,N
3892,Panagra Airways,\N,,PGI,PANAGRA,United States,N
3893,Panamedia,\N,,PEI,PANAMEDIA,Spain,N
3894,Panavia,\N,,PVI,,Panama,N
3895,Panh,\N,,PNH,KUBAN LIK,Russia,N
3896,Pannon Air Service,\N,,PHU,PANNON,Hungary,N
3897,Panorama,\N,,PNM,PANORAMA,Spain,N
3898,Panorama Air Tour,\N,,PAH,LANI,United States,N
3899,Panorama Flight Service,\N,,AFD,AIRFED,United States,N
3900,Pantanal Linhas Aéreas,\N,P8,PTN,PANTANAL,Brazil,N
3901,Papair Terminal,\N,,HMP,PAPAIR TERMINAL,Haiti,N
3902,Paradise Airways,\N,,PAI,SEA RAY,United States,N
3903,Paradise Island Airways,\N,,PDI,PARADISE ISLAND,United States,N
3904,Paragon Air Express,\N,,PGX,PARAGON EXPRESS,United States,N
3905,Paragon Global Flight Support,\N,,PGF,,United Kingdom,N
3906,Paramount Airlines,\N,,PRR,PARAMOUNT,Sierra Leone,N
3907,Paramount Airways,\N,I7,PMW,PARAWAY,India,Y
3908,Parcel Express,\N,,APE,AIR PARCEL,United States,N
3909,Pariz Air,\N,,IRE,PARIZAIR,Iran,N
3910,Pars Aviation Service,\N,,PRA,PARSAVIA,Iran,N
3911,Parsa,\N,,PST,TURISMO REGIONAL,Panama,N
3912,Parsons Airways Northern,\N,,FAP,,Canada,N
3913,Pascan Aviation,\N,,PSC,PASCAN,Canada,N
3914,Passaredo Transportes Aereos,\N,,PTB,PASSAREDO,Brazil,Y
3915,Patria Cargas Aereas,\N,,PTC,PATRIA,Argentina,N
3916,Patriot Aviation Limited,\N,,BYT,BYTE,United Kingdom,N
3917,Patterson Aviation Company,\N,,ETL,ENTEL,United States,N
3918,Pawan Hans,\N,,PHE,PAWAN HANS,India,N
3919,Payam Air,\N,,IRP,PAYAMAIR,Iran,N
3920,Peach Air,\N,,KGC,GOLDCREST,United Kingdom,N
3921,Pearl Air,\N,,PRL,PEARL LINE,Pakistan,N
3922,Pearl Air Services,\N,,PBY,PEARL SERVICES,Uganda,N
3923,Pearl Airways,\N,HP,HPA,PEARL AIRWAYS,Haiti,N
3924,Peau Vava%Gʻ%@u,\N,,PVU,PEAU,Tonga,N
3925,Pecotox Air,\N,,PXA,PECOTOX,Moldova,N
3926,Pegasus Airlines,\N,PC,PGT,SUNTURK,Turkey,Y
3927,Pegasus Aviation,\N,,PEV,PEGAVIATION,Greece,N
3928,Pegasus Hava Tasimaciligi,\N,1I,,,Turkey,N
3929,Pegasus Helicopters,\N,,HAK,HELIFALCON,Norway,N
3930,Pelican Air Services,\N,,PDF,PELICAN AIRWAYS,South Africa,N
3931,Pelican Express,\N,,PEX,PELICAN EXPRESS,United States,N
3932,Pelita Air Service,\N,,PAS,PELITA,Indonesia,N
3933,Pem-Air,\N,,PEM,PEM-AIR,Canada,N
3934,Pen-Avia,\N,,PDY,PENDLEY,United Kingdom,N
3935,Peninsula Airways,\N,KS,PEN,PENINSULA,United States,Y
3936,Peninter Aerea,\N,,PNE,PENINTER,Mexico,N
3937,Penya De L'Aire,\N,,PCA,PENA DEL AIRE,Spain,N
3938,Peran,\N,,CVT,CVETA,Kazakhstan,N
3939,Perforadora Central,\N,,PCC,PERFORADORA CENTRAL,Mexico,N
3940,Perimeter Aviation,\N,,PAG,PERIMETER,Canada,N
3941,Perm Airlines,\N,P9,PGP,PERM AIR,Russia,N
3942,Personas Y Pasquetes Por Air,\N,,PPQ,PERSONSPAQ,Mexico,N
3943,Peruvian Air Force,\N,,FPR,,Peru,N
3944,Peruvian Navy,\N,,INP,,Peru,N
3945,Petroleos Mexicanos,\N,,PMX,PEMEX,Mexico,N
3946,Petroleum Helicopters,\N,,PHM,PETROLEUM,United States,N
3947,Petroleum Helicopters de Colombia,\N,,PHC,HELICOPTERS,Colombia,N
3948,Petropavlovsk-Kamchatsk Air Enterprise,\N,,PTK,PETROKAM,Russia,N
3949,Petty Transport,\N,,PTY,PETTY,United States,N
3950,Phenix Aviation,\N,,PHV,NEW BIRD,France,N
3951,Phetchabun Airline,\N,,PMY,PHETCHABUN AIR,Thailand,N
3952,Philippine Airlines,\N,PR,PAL,PHILIPPINE,Philippines,Y
3953,Philips Aviation Services,\N,,PHI,PHILAIR,Netherlands,N
3954,Phillips Air,\N,,BCH,BEACHBALL,United States,N
3955,Phillips Alaska,\N,,PDD,PADA,United States,N
3956,Phillips Michigan City Flying Service,\N,,PHL,PHILLIPS,United States,N
3957,Phoebus Apollo Aviation,\N,,PHB,PHOEBUS,South Africa,N
3958,Phoebus Apolloa Zambia,\N,,KZM,CARZAM,Zambia,N
3959,Phoenix Air Group,\N,,PHA,GRAY BIRD,United States,N
3960,Phoenix Air Lines,\N,,PHN,PHOENIX BRASIL,Brazil,N
3961,Phoenix Air Service,\N,,PAM,PHOENIX,Germany,N
3962,Phoenix Air Transport,\N,,PPG,PAPAGO,United States,N
3963,Phoenix Airline Services,\N,,WDY,WINDYCITY,United States,N
3964,Phoenix Airways,\N,HP,,,Switzerland,N
3965,Phoenix Avia,\N,,PHY,PHOENIX ARMENIA,Armenia,N
3966,Phoenix Aviation,\N,,PHG,PHOENIX GROUP,Kyrgyzstan,N
3967,Phoenix Flight Operations,\N,,XPX,,United States,N
3968,Phuket Air,\N,9R,VAP,PHUKET AIR,Thailand,N
3969,Piedmont Airlines (1948-1989),\N,PI,PDT,PIEDMONT,United States,Y
3970,Pilatus Flugzeugwerke,\N,,PCH,PILATUS WINGS,Switzerland,N
3971,Pilatus PC-12 Center De Mexico,\N,,PLU,PILATUS MEXICO,Mexico,N
3972,Pimichikamac Air,\N,,MKS,MIKISEW,Canada,N
3973,Pineapple Air,\N,,PNP,PINEAPPLE AIR,Bahamas,N
3974,Pinframat,\N,,PIM,PINFRAMAT,Angola,N
3975,Pinnacle Air Group,\N,,PCL,PINNACLE GROUP,United States,N
3976,Pinnacle Airlines,\N,9E,FLG,FLAGSHIP,United States,Y
3977,Pioneer Airlines,\N,,PIO,PIONEER,United States,N
3978,Pioneers Limited,\N,,PER,,Pakistan,N
3979,Pirinair Express,\N,,PRN,PRINAIR EXPRESS,Spain,N
3980,Planar,\N,,PLN,PLANAR,Angola,N
3981,Planemaster Services,\N,,PMS,PLANEMASTER,United States,N
3982,Planet Airways,\N,,PLZ,PLANET,United States,N
3983,Players Air,\N,,PYZ,PLAYERS AIR,United States,N
3984,Ploizeihubschrauberstaffel Hamburg,\N,,LIB,LIBELLE,Germany,N
3985,Plymouth School of Flying,\N,,PSF,LIZARD,United Kingdom,N
3986,Pocono Air Lines,\N,,POC,POCONO,United States,N
3987,Podilia-Avia,\N,,PDA,PODILIA,Ukraine,N
3988,Point Afrique Niger,\N,,PAZ,POINTAIR NIGER,Niger,N
3989,Point Airlines,\N,,RMI,POINT AIRLINE,Nigeria,N
3990,Pointair Burkina,\N,,PAW,POINTAIR BURKINA,Burkina Faso,N
3991,Points of Call Airlines,\N,,PTS,POINTSCALL,Canada,N
3992,Polar Air Cargo,\N,PO,PAC,POLAR,United States,N
3993,Polar Airlines de Mexico,\N,,PMO,POLAR MEXICO,Mexico,N
3994,Polestar Aviation,\N,,PSR,POLESTAR,United Kingdom,N
3995,Polet,\N,,POT,POLET,Russia,Y
3996,Police Aux Frontiers,\N,,POF,AIRPOL,France,N
3997,Police Aviation Services,\N,,PLC,SPECIAL,United Kingdom,N
3998,Polish Air Force,\N,,PLF,POLISH AIRFORCE,Poland,N
3999,Polish Navy,\N,,PNY,POLISH NAVY,Poland,N
4000,Polizeifliegerstaffel Nordrhein-Westfalen,\N,,NRW,HUMMEL,Germany,N
4001,Polizeihibschrauberstaffel Neidersachsen,\N,,PPH,POLICE PHOENIX,Germany,N
4002,Polizeihuberschrauberstaffel Sachsen-Anhalt,\N,,PIK,POLICE IKARUS,Germany,N
4003,Polizeihubschrafterstaffel Rheinland-Pfalz,\N,,SRP,SPERBER,Germany,N
4004,Polizeihubschrauberstaffel Baden-Wurtemberg,\N,,PBW,BUSSARD,Germany,N
4005,Polizeihubschrauberstaffel Bayern,\N,,EDL,POLICE EDELWEISS,Germany,N
4006,Polizeihubschrauberstaffel Brandenburg,\N,,PBB,ADEBAR,Germany,N
4007,Polizeihubschrauberstaffel Hessen,\N,,PHH,IBIS,Germany,N
4008,Polizeihubschrauberstaffel Mecklenburg-Vorpommern,\N,,PMV,POLICE MERLIN,Germany,N
4009,Polizeihubschrauberstaffel Sachsen,\N,,PHS,PASSAT,Germany,N
4010,Polizeihubschrauberstaffel Thuringen,\N,,HBT,HABICHT,Germany,N
4011,Polo Aviation,\N,,CUK,CHUKKA,United Kingdom,N
4012,Polynesian Air-Ways,\N,,PLA,POLYAIR,United States,N
4013,Polynesian Airlines,\N,PH,PAO,POLYNESIAN,Samoa,Y
4014,Polynesian Blue,\N,DJ,PLB,POLYBLUE,New Zealand,N
4015,Polyot Sirena,\N,1U,,,Russia,N
4016,Pond Air Express,\N,,PND,POND AIR,United States,N
4017,Pont International Airline Services,\N,,PSI,PONT,Suriname,N
4018,Pool Aviation,\N,,PLX,POOLEX,United Kingdom,N
4019,Port Townsend Airways,\N,,PTQ,TOWNSEND,United States,N
4020,Porteadora De Cosola,\N,,POR,PORTEADORA,Mexico,N
4021,Porter Airlines,\N,PD,POE,PORTER AIR,Canada,Y
4022,Portugalia,\N,NI,PGA,PORTUGALIA,Portugal,Y
4023,Portuguese Air Force,\N,,AFP,PORTUGUESE AIR FORCE,Portugal,N
4024,Portuguese Army,\N,,POA,PORTUGUESE ARMY,Portugal,N
4025,Portuguese Navy,\N,,PON,PORTUGUESE NAVY,Portugal,N
4026,Potomac Air,\N,BK,PDC,DISTRICT,United States,Y
4027,Potosina Del Aire,\N,,PSN,POTOSINA,Mexico,N
4028,Powell Air,\N,,PWL,POWELL AIR,Canada,N
4029,Prairie Flying Service,\N,,PFS,PRAIRIE,United States,N
4030,Pratt and Whitney Canada,\N,,PWC,PRATT,Canada,N
4031,Precision Air,\N,PW,PRF,PRECISION AIR,Tanzania,Y
4032,Precision Airlines,\N,,PRE,PRECISION,United States,N
4033,Premiair,\N,,BAT,BALLISTIC,Luxembourg,N
4034,Premiair Aviation Services,\N,,PGL,PREMIERE,United Kingdom,N
4035,Premiair Fliyng Club,\N,,PME,ADUR,United Kingdom,N
4036,Premium Air Shuttle,\N,,EMI,BLUE SHUTTLE,Nigeria,N
4037,Premium Aviation,\N,,PMU,PREMIUM,Germany,N
4038,Presidence Du Faso,\N,,BFA,,Burkina Faso,N
4039,Presidencia de La Republica de Guinea Ecuatorial,\N,,ONM,,Equatorial Guinea,N
4040,President Airlines,null,TO,PSD,,Cambodia,N
4041,Presidential Aviation,\N,,PRD,PRESIDENTIAL,United States,N
4042,Priester Aviation,\N,,PWA,PRIESTER,United States,N
4043,Primair,\N,,PMM,PRIMAVIA,Russia,N
4044,Primaris Airlines,\N,FE,WCP,WHITECAP,United States,N
4045,Primas Courier,\N,,PMC,PRIMAC,United States,N
4046,Primavia Limited,\N,,CRY,CARRIERS,United Kingdom,N
4047,Prime Air,\N,,PRM,PRIME AIR,United States,N
4048,Prime Aviation,\N,,PKZ,PRAVI,Kazakhstan,N
4049,Prince Edward Air,\N,,CME,COMET,Canada,N
4050,Princely Jets,\N,,PJP,PRINCELY JETS,Pakistan,N
4051,Princess Air,\N,8Q,,,,N
4052,Princeton Aviation Corporation,\N,,PCN,PRINCETON,United States,N
4053,Priority Air Charter,\N,,PRY,PRIORITY AIR,United States,N
4054,Priority Air Transport,\N,,PAT,PAT,United States,N
4055,Priority Aviation Company,\N,,BCK,BANKCHECK,United States,N
4056,Privatair,\N,,PTI,PRIVATAIR,Switzerland,Y
4057,Private Jet Expeditions,\N,,PJE,PEE JAY,United States,N
4058,Private Jet Management,\N,,PJA,PRIVATE FLIGHT,United States,N
4059,Private Wings Flugcharter,\N,8W,PWF,PRIVATE WINGS,Germany,N
4060,Privilege Style L,\N,,PVG,PRIVILEGE,Spain,N
4061,Pro Air,\N,,PRH,PROHAWK,United States,N
4062,Pro Air Service,\N,,PSZ,POP-AIR,United States,N
4063,Probiz Guinee,\N,,GIY,PROBIZ,Guinea,N
4064,Professional Express Courier Service,\N,,PAD,AIR PROFESSIONAL,United States,N
4065,Professione VOlare,\N,,PVL,VOLARE,Italy,N
4066,Proflight Commuter Services,\N,P0,,,Zambia,Y
4067,Promotora Industria Totolapa,\N,,PTT,TOTOLAPA,Mexico,N
4068,Propair,\N,,PRO,PROPAIR,Canada,N
4069,Propflight Air Services,\N,,PFZ,PROFLIGHT-ZAMBIA,Zambia,N
4070,Propheter Aviation,\N,,PPA,AIR PROP,United States,N
4071,Proteus Helicopteres,\N,,PTH,PROTEUS,France,N
4072,Providence Airline,\N,,PTL,PLANTATION,United States,N
4073,Providence Aviation Services,\N,,AWD,,Pakistan,N
4074,Provincial Airlines,\N,,SPR,SPEEDAIR,Canada,N
4075,Provincial Express,\N,,PRV,PROVINCIAL,Canada,N
4076,Pskovavia,\N,,PSW,PSKOVAVIA,Russia,N
4077,Psudiklat Perhubungan Udara/PLP,\N,,UDA,UDARA,Indonesia,N
4078,Ptarmigan Airways,\N,,PTA,PTARMIGAN,Canada,N
4079,Publiservicios Aereos,\N,,PSP,PUBLISERVICIOS,Mexico,N
4080,Publivoo,\N,,PUV,PUBLIVOO,Portugal,N
4081,Puerto Rico National Guard,\N,,PNG,,United States,N
4082,Puerto Vallarta Taxi Aereo,\N,,TXV,TAXIVALLARTA,Mexico,N
4083,Pulkovo Aircraft Services,\N,,PGH,,Russia,N
4084,Puma Linhas Aereas,\N,,PLY,PUMA BRASIL,Brazil,N
4085,Puntavia Air Services,\N,,PTV,PUNTAVIA,Djibouti,N
4086,Punto Fa,\N,,MGO,MANGO,Spain,N
4087,Pyramid Air Lines,\N,,PYR,PYAIR,Egypt,N
4088,Qanot Sharq,\N,,QNT,QANAT SHARQ,Uzbekistan,N
4089,Qantas,Qantas Airways,QF,QFA,QANTAS,Australia,Y
4090,Qatar Air Cargo,\N,,QAC,QATAR CARGO,Qatar,N
4091,Qatar Airways,\N,QR,QTR,QATARI,Qatar,Y
4092,Qatar Amiri Flight,\N,,QAF,AMIRI,Qatar,N
4093,Qeshm Air,\N,,IRQ,QESHM AIR,Iran,N
4094,Quantex Environmental,\N,,QTX,AIR QUANTEX,Canada,N
4095,Quebec Government Air Service,\N,,QUE,QUEBEC,Canada,N
4096,Queen Air,\N,,QNA,QUEEN AIR,Dominican Republic,N
4097,Quest Diagnostics,\N,,LBQ,LABQUEST,United States,N
4098,Quick Air Jet Charter,\N,,QAJ,DAGOBERT,Germany,N
4099,Quick Airways Holland,\N,,QAH,QUICK,Netherlands,N
4100,Quisqueya Airlines,\N,,QAS,QUISQUEYA,Haiti,N
4101,Qurinea Air Service,\N,,QAQ,QURINEA AIR,Libya,N
4102,Qwest Commuter Corporation,\N,,QCC,QWEST AIR,United States,N
4103,Qwestair,\N,,QWA,,Australia,N
4104,Qwila Air,\N,,QWL,Q-CHARTER,South Africa,N
4105,RA Jet Aeroservicios,\N,,RJT,RA JET,Mexico,N
4106,RACSA,\N,R6,,,Guatemala,Y
4107,RAF Barkston Heath,\N,,BKH,,United Kingdom,N
4108,Church Fenton Flying Training Unit,\N,,CFN,CHURCH FENTON,United Kingdom,N
4109,Coltishall Flying Training Unit,\N,,COH,COLT,United Kingdom,N
4110,Coningsby Flying Training Unit,\N,,CBY,TYPHOON,United Kingdom,N
4111,Cottesmore Flying Training Unit,\N,,COT,COTTESMORE,United Kingdom,N
4112,Cranwell Flying Training Unit,\N,,CWL,CRANWELL,United Kingdom,N
4113,Kinloss Flying Training Unit,\N,,KIN,KINLOSS,United Kingdom,Y
4114,Leeming Flying Training Unit,\N,,LEE,JAVELIN,United Kingdom,N
4115,RAF Leuchars,\N,,LCS,LEUCHARS,United Kingdom,N
4116,Linton-on-Ouse Flying Training Unit,\N,,LOP,LINTON ON OUSE,United Kingdom,N
4117,Lossiemouth Flying Training Unit,\N,,LOS,LOSSIE,United Kingdom,N
4118,Marham Flying Training Unit,\N,,MRH,MARHAM,United Kingdom,N
4119,Northwood Headquarters (RAF,\N,,NWO,,United Kingdom,N
4120,RAF Scampton,\N,,SMZ,SCAMPTON,United Kingdom,N
4121,RAF St Athan,\N,,STN,SAINT ATHAN,UNited Kingdom,N
4122,RAF St Mawgan Search and Rescue,\N,,SMG,,United Kingdom,N
4123,RAF Topcliffe Flying Training Unit,\N,,TOF,TOPCLIFFE,United Kingdom,N
4124,RAF Valley Flying Training Unit,\N,,VYT,ANGLESEY,United Kingdom,N
4125,RAF Valley SAR Training Unit,\N,,VLL,,United Kingdom,N
4126,Waddington FTU,\N,,WAD,VULCAN,United Kingdom,N
4127,Wittering FTU,\N,,WIT,STRIKER,United Kingdom,N
4128,RAF-Avia,\N,,MTL,MITAVIA,Latvia,N
4129,RAK Airways,\N,,RKM,RAKAIR,United Arab Emirates,N
4130,RWL Luftfahrtgesellschaft,\N,,RWL,RHEINTRAINER,Germany,N
4131,Rabbit-Air,\N,,RBB,RABBIT,Switzerland,N
4132,Race Cargo Airlines,\N,,ACE,Fastcargo,Ghana,N
4133,Rader Aviation,\N,,GBR,GREENBRIER AIR,United States,N
4134,Radixx Solutions International,\N,1D,,,United States,N
4135,Raji Airlines,\N,,RAJ,RAJI,Pakistan,N
4136,Raleigh Flying Service,\N,,RFA,RALEIGH SERVICE,United States,N
4137,Ram Air Freight,\N,,REX,RAM EXPRESS,United States,N
4138,Ram Aircraft Corporation,\N,,RMT,RAM FLIGHT,United States,N
4139,Ramp 66,\N,,PPK,PELICAN,United States,N
4140,Rangemile Limited,\N,,RGM,RANGEMILE,United Kingdom,N
4141,Raslan Air Service,\N,,MWR,RASLAN,Egypt,N
4142,Rath Aviaton,\N,,RAQ,RATH AVIATION,Austria,N
4143,Ratkhan Air,\N,,CSM,LORRY,Kazakhstan,N
4144,Raven Air,\N,,RVR,RAVEN,United Kingdom,N
4145,Raven Air,\N,,RVN,RAVEN U-S,United States,N
4146,Ray Aviation,\N,,REI,RAY AVIATION,Israel,N
4147,Raya Jet,\N,,RYT,,Jordan,N
4148,Raytheon Aircraft Company,\N,,RTN,RAYTHEON,United States,N
4149,Raytheon Corporate Jets,\N,,RCJ,NEWPIN,United Kingdom,N
4150,Raytheon Travel Air,\N,,KSS,KANSAS,United States,N
4151,Real Aero Club De Baleares,\N,,RCB,BALEARES,Spain,N
4152,Real Aero Club de Reus-Costa Dorado,\N,,CDT,AEROREUS,Spain,N
4153,Real Aeroclub De Ternerife,\N,,RCD,AEROCLUB,Spain,N
4154,Real Aviation,\N,,RLV,REAL,Ghana,N
4155,Rebus,\N,,REB,REBUS,Bulgaria,N
4156,Red Aviation,\N,,PSH,PASSION,United Kingdom,N
4157,Red Baron Aviation,\N,,RBN,RED BARON,United States,N
4158,Red Devils Parachute Display Team,\N,,DEV,RED DEVILS,United Kingdom,N
4159,Red Sea Aviation,\N,,RDV,RED AVIATION,Egypt,N
4160,Red Sky Ventures,\N,,RSV,RED SKY,Namibia,N
4161,Red Star,\N,,STR,STARLINE,United Arab Emirates,N
4162,Redhill Aviation,\N,8L,RHC,REDAIR,United Kingdom,N
4163,Reed Aviation,\N,,RAV,REED AVIATION,United Kingdom,N
4164,Reef Air,\N,,REF,REEF AIR,New Zealand,N
4165,Reem Air,\N,V4,REK,REEM AIR,Kyrgyzstan,N
4166,Reeve Aleutian Airways,\N,,RVV,REEVE,United States,N
4167,Regal Bahamas International Airways,\N,,RBH,CALYPSO,Bahamas,N
4168,Regency Airlines,\N,,RGY,REGENCY,United States,N
4169,Regent Air,\N,,RAH,REGENT,Canada,N
4170,Regio Air,\N,,RAG,GERMAN LINK,Germany,N
4171,Region Air,\N,,RGR,REGIONAIR,Canada,N
4172,Regional 1,\N,,TSH,TRANSCANADA,Canada,N
4173,Regional Air,\N,,RIL,,Mauritania,N
4174,Regional Air Express,\N,,REW,REGIONAL WINGS,Germany,N
4175,Regional Air Lines,\N,,RGL,MAROC REGIONAL,Morocco,N
4176,Regional Air Services,\N,,REG,REGIONAL SERVICES,Tanzania,N
4177,Regional Airlines,\N,FN,,,Morocco,Y
4178,Regional Express,\N,ZL,RXA,REX,Australia,Y
4179,Regional Geodata Air,\N,,JJM,GEODATA,Spain,N
4180,RegionsAir,\N,3C,CEA,CORP-X,United States,N
4181,Reliance Aviation,\N,,REL,RELIANCE AIR,United States,N
4182,Reliant Airlines,\N,,RLT,RELIANT,United States,N
4183,Relief Transport Services,\N,,RTS,RELIEF,United Kingdom,N
4184,Renan,\N,,RAN,RENAN,Moldova,N
4185,Reno Air,\N,QQ,ROA,RENO AIR,United States,N
4186,Renown Aviation,\N,,RGS,RENOWN,United States,N
4187,Republic Airlines,\N,RW,RPA,BRICKYARD,United States,Y
4188,Republic Express Airlines,\N,RH,RPH,PUBLIC EXPRESS,Indonesia,Y
4189,Republicair,\N,,RBC,REPUBLICAIR,Mexico,N
4190,Resort Air,\N,,RST,RESORT AIR,United States,N
4191,Rhoades Aviation,\N,,RDS,RHOADES EXPRESS,United States,N
4192,Riau Airlines,\N,,RIU,RIAU AIR,Indonesia,N
4193,Rich International Airways,\N,,RIA,RICHAIR,United States,N
4194,Richards Aviation,\N,,RVC,RIVER CITY,United States,N
4195,Richardson's Airway,\N,,RIC,RICHARDSON,United States,N
4196,Richland Aviation,\N,,RCA,RICHLAND,United States,N
4197,Rick Lucas Helicopters,\N,,HPR,HELIPRO,New Zealand,N
4198,Rico Linhas A,\N,C7,RLE,RICO,Brazil,N
4199,Ridder Avia,\N,,RID,AKRID,Kazakhstan,N
4200,Riga Airclub,\N,,RAK,SPORT CLUB,Latvia,N
4201,Rijnmond Air Services,\N,,RAZ,RIJNMOND,Netherland,N
4202,Rikspolisstyrelsen,\N,,POL,,Sweden,N
4203,Rimrock Airlines,\N,,RIM,RIMROCK,United States,N
4204,Rio Air Express,\N,,SKA,RIO EXPRESS,Brazil,N
4205,Rio Airways,\N,,REO,RIO,United States,N
4206,Rio Grande Air,\N,E2,GRN,GRANDE,United States,N
4207,Rio Sul Servi,\N,SL,RSL,RIO SUL,Brazil,N
4208,River Ministries Air Charter,\N,,RVM,RIVER,South Africa,N
4209,River State Government of Nigeria,\N,,RGP,GARDEN CITY,Nigeria,N
4210,Rivne Universal Avia,\N,,UNR,RIVNE UNIVERSAL,Ukraine,N
4211,Roadair Lines,\N,,RDL,ROADAIR,Canada,N
4212,Robinton Aero,\N,,RBT,ROBIN,Dominican Republic,N
4213,Roblex Aviation,\N,,ROX,ROBLEX,United States,N
4214,Rockwell Collins Avionics,\N,,RKW,ROCKWELL,United States,N
4215,Rocky Mountain Airlines,\N,,ROC,,Canada,N
4216,Rocky Mountain Airways,\N,,RMA,ROCKY MOUNTAIN,United States,N
4217,Rocky Mountain Holdings,\N,,LIF,LIFECARE,United States,N
4218,Rodze Air,\N,,RDZ,RODZE AIR,Nigeria,N
4219,Rog-Air,\N,,FAD,AIR FRONTIER,Canada,N
4220,Rollright Aviation,\N,,RRZ,ROLLRIGHT,United Kingdom,N
4221,Rolls-Royce Limited,\N,,RRL,MERLIN,United Kingdom,N
4222,Rolls-Royce plc,\N,,BTU,ROLLS,United Kingdom,N
4223,Romanian Air Force,\N,,ROF,ROMAF,Romania,N
4224,Romavia,\N,,RMV,AEROMAVIA,Romania,N
4225,Ronso,\N,,RNS,RONSO,Mexico,N
4226,Roraima Airways,\N,,ROR,RORAIMA,Guyana,N
4227,Rosneft-Baltika,\N,,RNB,ROSBALT,Russia,N
4228,Ross Aviation,\N,,NRG,ENERGY,United States,N
4229,Rossair,\N,,RFS,,Australia,N
4230,Rossair,\N,,RSS,ROSS CHARTER,South Africa,N
4231,Rossair Europe,\N,,ROS,CATCHER,Netherlands,N
4232,Rossiya,\N,R4,,,Russia,Y
4233,Roswell Airlines,\N,,RAL,ROSWELL,United States,N
4234,Air Rarotonga,\N,GZ,RAR,,Cook Islands,Y
4235,Rotatur,\N,,RTR,ROTATUR,Brazil,N
4236,Rotormotion,\N,,RKT,ROCKET,United Kingdom,N
4237,Rotterdam Jet Center,\N,,JCR,ROTTERDAM JETCENTER,Netherlands,N
4238,Rover Airways International,\N,,ROV,ROVERAIR,United States,N
4239,Rovos Air,\N,,VOS,ROVOS,South Africa,N
4240,Royal Air Cargo,\N,,RCG,ROYAL CARGO,South Africa,N
4241,Royal Air Force,\N,RR,RFR,RAFAIR,United Kingdom,N
4242,Royal Air Force of Oman,\N,RS,MJN,MAJAN,Oman,N
4243,Royal Air Force,\N,,ACW,AIR CADET,United Kingdom,N
4244,Royal Air Force,\N,,RRR,ASCOT,United Kingdom,N
4245,Royal Air Force,\N,,RRF,KITTY,United Kingdom,N
4246,Royal Air Force,\N,,SHF,VORTEX,United Kingdom,N
4247,Royal Air Freight,\N,,RAX,AIR ROYAL,United States,N
4248,Royal Air Maroc,\N,AT,RAM,ROYALAIR MAROC,Morocco,Y
4249,Royal Airlines,\N,R0,RPK,ROYAL PAKISTAN,Pakistan,N
4250,Royal American Airways,\N,,RLM,ROYAL AMERICAN,United States,N
4251,Royal Aruban Airline,\N,V5,RYL,ROYAL ARUBAN,Aruba,N
4252,Royal Australian Air Force,\N,,ASY,AUSSIE,Australia,N
4253,Royal Aviation Express,\N,,RXP,ROY EXPRESS,Canada,N
4254,Royal Bahrain Airlines,\N,,RYB,ROYAL BAHRAIN,Bahrain,N
4255,Royal Brunei Airlines,\N,BI,RBA,BRUNEI,Brunei,Y
4256,Royal Daisy Airlines,\N,,KDR,DARLINES,Uganda,N
4257,Royal Ghanaian Airlines,\N,,RGA,ROYAL GHANA,Ghana,N
4258,Royal Jet,\N,,ROJ,ROYALJET,United Arab Emirates,N
4259,Royal Jordanian,\N,RJ,RJA,JORDANIAN,Jordan,Y
4260,Royal Jordanian Air Force,\N,,RJZ,JORDAN AIR FORCE,Jordan,N
4261,Royal Khmer Airlines,\N,RK,RKH,KHMER AIR,Cambodia,N
4262,Royal Malaysian Air Force,\N,,RMF,ANGKASA,Malaysia,N
4263,Royal Navy,\N,,NVY,NAVY,United Kingdom,N
4264,Royal Nepal Airlines,\N,RA,RNA,ROYAL NEPAL,Nepal,Y
4265,Royal Netherland Navy,\N,,NRN,NETHERLANDS NAVY,Netherlands,N
4266,Royal Netherlands Air Force,\N,,NAF,NETHERLANDS AIR FORCE,Netherlands,N
4267,Royal New Zealand Air Force,\N,,KIW,KIWI,New Zealand,N
4268,Royal Norwegian Air Force,\N,,NOW,NORWEGIAN,Norway,N
4269,Royal Oman Police,\N,,ROP,,Oman,N
4270,Royal Phnom Penh Airways,\N,,PPW,PHNOM-PENH AIR,Cambodia,Y
4271,Royal Rwanda Airlines,\N,,RRA,ROYAL RWANDA,Rwanda,N
4272,Royal Saudi Air Force,\N,,RSF,ARSAF,Saudi Arabia,N
4273,Royal Sky,\N,,RYS,ROYAL SKY,Thailand,N
4274,Royal Swazi National Airways,\N,,RSN,SWAZI NATIONAL,Swaziland,N
4275,Royal Tongan Airlines,\N,WR,HRH,TONGA ROYAL,Tonga,N
4276,Royal West Airlines,\N,,RWE,ROYAL WEST,United States,N
4277,Rubystar,\N,,RSB,RUBYSTAR,Belarus,N
4278,Rumugu Air & Space Nigeria,\N,,RMG,RUMUGU AIR,Nigeria,N
4279,Rusaero,\N,,RUR,,Russia,N
4280,Rusaero,\N,,KLE,,Russia,N
4281,Rusair JSAC,\N,,CGI,CGI-RUSAIR,Russia,N
4282,Rusich-T,\N,,RUH,,Russia,N
4283,Rusline,\N,,RLU,RUSLINE AIR,Russia,Y
4284,Russian Aircraft Corporation-MiG,\N,,MIG,MIG AVIA,Russia,N
4285,Russian Federation Air Force,\N,,RFF,RUSSIAN AIRFORCE,Russia,N
4286,Russian Sky Airlines,\N,P7,ESL,RADUGA,Russia,N
4287,Rusuertol,\N,,RUZ,ROSTUERTOL,Russia,N
4288,Rutas Aereas,\N,,RUC,RUTACA,Venezuela,N
4289,Rutland Aviation,\N,,RND,RUTLAND,United Kingdom,N
4290,Rwanda Airlines,\N,,RUA,,Rwanda,N
4291,Rwanda Airways,\N,,RWA,,Rwanda,N
4292,Rwandair Express,\N,WB,RWD,RWANDAIR,Rwanda,Y
4293,Ryan Air Service,\N,,RCT,ARCTIC TRANPORT,United States,N
4294,Ryan Air Services,\N,,RYA,RYAN AIR,United States,Y
4295,Ryan International Airlines,\N,RD,RYN,RYAN INTERNATIONAL,United States,Y
4296,Ryanair,\N,FR,RYR,RYANAIR,Ireland,Y
4297,Ryazan State Air Enterprise,\N,,RYZ,RYAZAN AIR,Russia,N
4298,Rynes Aviation,\N,,RAA,RYNES AVIATION,United States,N
4299,Régional,\N,YS,RAE,REGIONAL EUROPE,France,Y
4302,South Asian Airlines,\N,,BDS,SOUTH ASIAN,Bangladesh,N
4303,State Air Company Berkut,\N,,BEC,,Kazakhstan,N
4304,SATA International,\N,S4,RZO,AIR AZORES,Portugal,Y
4305,South African Airways,SAA South African Airways,SA,SAA,SPRINGBOK,South Africa,Y
4306,Sky Way Air,\N,,SAB,SKY WORKER,Kyrgyzstan,N
4307,SASCO Airlines,\N,,SAC,SASCO,Sudan,N
4308,Republic of Singapore Air Force,\N,,SAF,SINGA,Singapore,N
4309,SOS Flygambulans,\N,,SAG,MEDICAL AIR,Sweden,N
4310,Sayakhat Airlines,\N,,SAH,SAYAKHAT,Kazakhstan,N
4311,Shaheen Air International,\N,NL,SAI,SHAHEEN AIR,Pakistan,Y
4312,Golden Eagle Air Services (FAA),\N,,SAJ,SAJEN,Canada,N
4313,Red Arrows Display Squadron,\N,,SAK,RED ARROWS,United Kingdom,N
4314,SAM Colombia,\N,MM,SAM,SAM,Colombia,N
4315,Servicios Aereos Nacionales (SAN),\N,,SAN,AEREOS,Ecuador,N
4316,Sahel Aviation Service,\N,,SAO,SAVSER,Mali,N
4317,Secretaria de Marina,\N,,ANX,SECRETARIA DEMARINA,Mexico,N
4318,Springbank Aviation,\N,,SAQ,SPRINGBANK,Canada,N
4319,Scandinavian Airlines System,SAS Scandinavian Airlines,SK,SAS,SCANDINAVIAN,Sweden,Y
4320,Samal Air,\N,,SAV,,Kazakhstan,N
4321,Sham Wing Airlines,\N,,SAW,SHAMWING,Syrian Arab Republic,N
4322,Sabah Air,\N,,SAX,SABAH AIR,Malaysia,N
4323,ScotAirways,\N,,SAY,SUCKLING,United Kingdom,Y
4324,Swiss Air-Ambulance,\N,,SAZ,SWISS AMBULANCE,Switzerland,N
4325,STA-MALI,\N,,SBA,STA-MALI,Mali,N
4326,Steinman Aviation,\N,,SBB,SABER EXPRESS,United States,N
4327,Seven Bar Flying Service,\N,,SBF,S-BAR,United States,N
4328,Sabre Incorporated,\N,,SBG,,United States,N
4329,S7 Airlines,Sibir Airlines,S7,SBI,SIBERIAN AIRLINES,Russia,Y
4330,Sobel Airlines of Ghana,\N,,SBL,SOBGHANA,Ghana,N
4331,Sky Bahamas,\N,,SBM,SKY BAHAMAS,Bahamas,N
4332,Stabo Air,\N,,SBO,STABAIR,Zambia,N
4333,Smithkline Beacham Clinical Labs,\N,,SBQ,SKIBBLE,United States,N
4334,Saber Aviation,\N,,SBR,FREIGHTER,United States,N
4335,Seaborne Airlines,\N,BB,SBS,SEABORNE,United States,Y
4336,Saint Barth Commuter,\N,,SBU,BLACK FIN,France,N
4337,Star Air,\N,,URJ,STARAV,Pakistan,N
4338,Scibe Airlift,\N,,SBZ,SCIBE AIRLIFT,Democratic Republic of the Congo,N
4339,Spanish Air Force,\N,,AME,AIRMIL,Spain,N
4340,South Central Air,\N,,SCA,SOUTH CENTRAL,United States,N
4341,Seacoast Airlines,\N,,SCC,SEA-COASTER,United States,N
4342,Scenic Airlines,\N,,SCE,SCENIC,United States,Y
4343,Socofer,\N,,SCF,SOCOFER,Angola,N
4344,Servicios Aereos San Cristobal,\N,,SCI,SAN CRISTOBAL,Mexico,N
4345,Sky Cam,\N,,SCK,SKYCAM,France,N
4346,Switfair Cargo,\N,,SCL,SWIFTAIR,Canada,N
4347,South American Airlines,\N,,SCN,SOUTH AMERICAN,Peru,N
4348,Servicios Aereos de Chihuahua Aerochisa,\N,,AHI,AEROCHISA,Mexico,N
4349,SriLankan Airlines,\N,UL,ALK,SRILANKAN,Sri Lanka,Y
4350,Sunbird Airlines,\N,,CDL,CAROLINA,United States,N
4351,Scorpio,\N,,SCP,SCORPIO,Egypt,N
4352,Si-Chang Flying Service,\N,,SCR,SICHART,Thailand,N
4353,South African Non Scheduled Airways,\N,,SCS,SOUTHERN CHARTERS,South Africa,N
4354,SAAB-Aircraft,\N,,SCT,SAAB-CRAFT,Sweden,N
4355,Servicios Aereos Del Centro,\N,,SCV,SACSA,Mexico,N
4356,Sun Country Airlines,\N,SY,SCX,SUN COUNTRY,United States,Y
4357,St. Andrews Airways,\N,,SDA,SAINT ANDREWS,Canada,N
4358,Sukhoi Design Bureau Company,\N,,SDB,SU-CRAFT,Russia,N
4359,Sunrise Airlines,\N,,SDC,SUNDANCE,United States,N
4360,Skymaster Air Taxi,\N,,SDD,SKY DANCE,United States,N
4361,Sundorph Aeronautical Corporation,\N,,SDF,SUNDORPH,United States,N
4362,Servicio De Helicopteros,\N,,SDH,ARCOS,Spain,N
4363,SADELCA - Sociedad Aerea Del Caqueta,\N,,SDK,SADELCA,Colombia,N
4364,Skydrift,\N,,SDL,SKYDRIFT,United Kingdom,N
4365,Spirit of Africa Airlines,\N,,SDN,BLUE NILE,Sudan,N
4366,Sud Airlines,\N,,SDU,SUD LINES,France,N
4367,Servicios Aereos Del Vaupes,\N,,SDV,SELVA,Colombia,N
4368,Servicio Tecnico Aero De Mexico,\N,,SDX,SERVICIO TECNICO,Mexico,N
4369,Sudan Pezetel for Aviation,\N,,SDZ,SUDANA,Sudan,N
4370,Southeast Air,\N,,SEA,SOUTHEAST AIR,United States,Y
4371,Servicios Aereos Luce,\N,,SEB,SERVILUCE,Mexico,N
4372,Sedona Air Center,\N,,SED,SEDONA AIR,United States,N
4373,Shaheen Air Cargo,\N,,SEE,SHAHEEN CARGO,Pakistan,N
4374,Sky Express,\N,G3,SEH,AIR CRETE,Greece,Y
4375,Spicejet,\N,SG,SEJ,SPICEJET,India,Y
4376,Skyjet,\N,,SEK,SKALA,Kazakhstan,N
4377,Sentel Corporation,\N,,SEL,SENTEL,United States,N
4378,Selcon Airlines,\N,,SEO,SELCON AIR,Nigeria,N
4379,Sky Eyes,\N,I6,SEQ,SKY EYES,Thailand,N
4380,Servicio Aereo Saltillo,\N,,SES,SERVISAL,Mexico,N
4381,SAETA,\N,EH,SET,SAETA,Ecuador,N
4382,Serair Transworld Press,\N,,SEV,CARGOPRESS,Spain,N
4383,SEFA,\N,,SFA,SEFA,France,N
4384,Shuswap Flight Centre,\N,,SFC,SHUSWAP,Canada,N
4385,Sefofane Air Charters,\N,,SFE,SEFOFANE,South Africa,N
4386,Safewings Aviation Company,\N,,SFF,SWIFTWING,United States,N
4387,Sun Freight Logistics,\N,,SFG,AERO GULF,Thailand,N
4388,Star Flyer,\N,7G,SFJ,STARFLYER,Japan,Y
4389,Southflight Aviation,\N,,SFL,SOUTHFLIGHT,New Zealand,N
4390,Safiran Airlines,\N,,SFN,SAFIRAN,Iran,N
4391,Safe Air,\N,,SFP,SAFE AIR,Pakistan,N
4392,Safair,\N,FA,SFR,CARGO,South Africa,N
4393,Southern Frontier Air Transport,\N,,SFS,SOUTHERN FRONTIER,Canada,N
4394,Skyfreight,\N,,SFT,SKYFREIGHT,United States,N
4395,Solent Flight,\N,,SFU,SAINTS,United Kingdom,N
4396,S.K. Logistics,\N,,SFX,SWAMP FOX,United States,N
4397,Sky King,\N,,SGB,SONGBIRD,United States,N
4398,Southern Right Air Charter,\N,,SGC,SOUTHERNRIGHT,South Africa,N
4399,Sky Gate International Aviation,\N,,SGD,AIR BISHKEK,Kyrgyzstan,N
4400,STAC Swiss Government Flights,\N,,SGF,STAC,Switzerland,N
4401,Servisair,\N,,SGH,SERVISAIR,United Kingdom,N
4402,Servicios Aereos Agricolas,\N,,SGI,SERAGRI,Chile,N
4403,Skyward Aviation,\N,,SGK,SKYWARD,Canada,N
4404,Sky Aircraft Service,\N,,SGM,SIGMA,Netherlands,N
4405,Siam GA,\N,,SGN,SIAM,Thailand,N
4406,Sagolair Transportes Ejecutivos,\N,,SGP,SAGOLAIR,Spain,N
4407,Saskatchewan Government Executive Air Service,\N,,SGS,SASKATCHEWAN,Canada,N
4408,Skygate,\N,,SGT,SKYGATE,Netherlands,N
4409,Samgau,\N,,SGU,RAUSHAN,Kazakhstan,N
4410,Saga Airlines,\N,,SGX,,Turkey,N
4411,Skagway Air Service,\N,N5,SGY,SKAGWAY AIR,United States,Y
4412,Shabair,\N,,SHB,SHABAIR,Democratic Republic of the Congo,N
4413,Sky Harbor Air Service,\N,,SHC,SKY HARBOR CHEYENNE,United States,N
4414,Sahara Airlines,\N,,SHD,,Algeria,Y
4415,Shell Aircraft,\N,,SHE,SHELL,United Kingdom,N
4416,Shoprite Group,\N,,SHG,SHOP AIR,United Kingdom,N
4417,Seoul Air International,\N,,SHI,SEOUL AIR,Republic of Korea,N
4418,Sharjah Ruler's Flight,\N,,SHJ,SHARJAH,United Arab Emirates,N
4419,Shorouk Air,\N,,SHK,,Egypt,N
4420,Samson Aviation,\N,,SHL,SAMSON,United Kingdom,N
4421,Sheltam Aviation,\N,,SHM,SHELTAM,South Africa,N
4422,Shaheen Airport Services,\N,,SHN,SUGAR ALFA,Pakistan,N
4423,Sheremetyevo-Cargo,\N,,SHO,,Russia,N
4424,Service Aerien Francais,\N,,SHP,SAF,France,N
4425,Shanghai Airlines Cargo,\N,,SHQ,SHANGHAI CARGO,China,N
4426,Shooter Air Courier,\N,,SHR,SHOOTER,Canada,N
4427,Shura Air Transport Services,\N,,SHS,SHURA AIR,Ethiopia,N
4428,Sakhalinskie Aviatrassy (SAT),\N,,SHU,SATAIR,Russia,N
4429,SATA Air Acores,\N,SP,SAT,SATA,Portugal,Y
4430,Scorpio Aviation,\N,8S,,,,N
4431,Shavano Air,\N,,SHV,SHAVANO,United States,N
4432,Shawnee Airline,\N,,SHW,SHAWNEE,United States,N
4433,Slim Aviation Services,\N,,SHX,SLIM AIR,Nigeria,N
4434,Sky Airlines,\N,,SHY,ANTALYA BIRD,Turkey,N
4435,Singapore Airlines,\N,SQ,SIA,SINGAPORE,Singapore,Y
4436,Sibaviatrans,\N,5M,SIB,SIBAVIA,Russia,Y
4437,Sierra Express,\N,,SIE,SEREX,United States,N
4438,Skynet Airlines,\N,SI,SIH,BLUEJET,Ireland,Y
4439,Seco International,\N,,SIJ,,Japan,N
4440,Servicios Aeronauticos Integrales,\N,,SIL,SERVICIOS INTEGRALES,Mexico,N
4441,Star Air,\N,,SIM,,Sierra Leone,N
4442,Sirio,\N,,SIO,SIRIO,Italy,N
4443,Salair,\N,,SIR,SALAIR,United States,N
4444,Saber Airlines,\N,,SIS,,Egypt,N
4445,SITA,\N,XS,SIT,,Belgium,N
4446,Slovenian Armed Forces,\N,,SIV,SLOVENIAN,Slovenia,N
4447,Sirio Executive,\N,,SIW,SIRIO EXECUTIVE,Italy,N
4448,Servicios Aereos Especiales De Jalisco,\N,,SJA,SERVICIOJAL,Mexico,N
4449,Servicios Ejecutivos Continental,\N,,SJC,SERVIEJECUTIVO,Mexico,N
4450,Sunair 2001,\N,,SJE,SUNBIZ,South Africa,N
4451,Spirit Aviation,\N,,SJJ,SPIRIT JET,United States,N
4452,Servicios Especiales Del Pacifico Jalisco,\N,,SJL,SERVICIOS JALISCO,Mexico,N
4453,Swiss Jet,\N,,SJT,SWISS JET,Switzerland,N
4454,Sriwijaya Air,\N,SJ,SJY,SRIWIJAYA,Indonesia,Y
4455,Sama Airlines,\N,ZS,SMY,NAJIM,Saudi Arabia,Y
4456,Southern Jersey Airways,\N,,ALC, Inc.,ACOM,N
4457,SPASA,\N,,SPS,SALDUERO,Spain,N
4458,Speed Aviation,\N,,SPT,SPEED AVIATION,Bangladesh,N
4459,Southeast Airmotive,\N,,SPU,SPUTTER,United States,N
4460,Servicios Privados De Aviacion,\N,,SPV,SERVICIOS PRIVADOS,Mexico,N
4461,Speedwings,\N,,SPW,SPEEDWING,Switzerland,N
4462,Service People Gesselschaft fur Charter und Service,\N,,SPX,,Germany,N
4463,Slovak National Aeroclub,\N,,SQA,SLOVAK AEROCLUB,Slovakia,N
4465,Slovak Air Force,\N,,SQF,SLOVAK AIRFORCE,Slovakia,N
4466,Servicos De Alquiler,\N,,SQL,ALQUILER,Mexico,N
4467,Sair Aviation,\N,,SRA,SAIR,Canada,N
4468,Searca,\N,,SRC,SEARCA,Colombia,N
4469,Siem Reap Airways,\N,FT,SRH,SIEMREAP AIR,Cambodia,Y
4470,Sky Work Airlines,SkyWork,SX,SRK,SKYFOX,Switzerland,N
4471,Swedline Express,\N,SM,SRL,Starline,Sweden,N
4472,Servicios Aeronauticos Aero Personal,\N,,SRL,SERVICIOS PERSONAL,Mexico,N
4473,Sirair,\N,,SRN,SIRAIR,Russia,N
4474,Servicios Aereos Ejecutivos Saereo,\N,,SRO,SAEREO,Ecuador,N
4475,South East Asian Airlines,\N,DG,SRQ,SEAIR,Philippines,Y
4476,Star Air,\N,,SRR,WHITESTAR,Denmark,N
4477,Selkirk Remote Sensing,\N,,SRS,PHOTO CHARLIE,Canada,N
4478,Star Up,\N,,SRU,STAR-UP,Peru,N
4479,Sarit Airlines,\N,,SRW,SARIA,Sudan,N
4480,Sierra Expressway Airlines,\N,,SRX,SIERRA EX,United States,N
4481,Strato Air Services,\N,,SRZ,STRATO,South Africa,N
4482,Sasair,\N,,SSB,SASIR,Canada,N
4483,Southern Seaplane,\N,,SSC,SOUTHERN SKIES,United States,N
4484,Star Service International,\N,,SSD,STAR SERVICE,France,N
4485,Servicios Aereos Sunset,\N,,SSE,SUNSET,Mexico,N
4486,Severstal,\N,,SSF,SEVERSTAL,Russia,N
4487,Slovak Government Flying Service,\N,,SSG,SLOVAK GOVERNMENT,Slovakia,N
4488,SwedJet Airways,\N,VD,BBB,BLACKBIRD,Sweden,N
4489,Skystar International,\N,,SSK,SKYSTAR,United States,N
4490,Special Scope Limited,\N,,SSO,DOPE,United Kingdom,N
4491,Starspeed,\N,,SSP,STARSPEED,United Kingdom,N
4492,Sunstate Airlines,\N,,SSQ,SUNSTATE,Australia,N
4493,SAESA,\N,,SSS,SAESA,Spain,N
4494,Sunwest Airlines,\N,,SST,SUNFLIGHT,Canada,N
4495,SASCA,\N,,SSU,SASCA,Venezuela,N
4496,Skyservice Airlines,\N,5G,SSV,SKYTOUR,Canada,Y
4497,Streamline Aviation,\N,,SSW,STREAMLINE,United Kingdom,N
4498,Sky Aviation,\N,,SSY,SIERRA SKY,Sierra Leone,N
4499,Specsavers Aviation,\N,,SSZ,SPECSAVERS,United Kingdom,N
4500,Star Aviation,\N,,STA,STAR,United Kingdom,N
4501,Status-Alpha Airline,\N,,STB,STATUS-ALPHA,Ukraine,N
4502,Stadium City Limited,\N,,STC,STADIUM,United Kingdom,N
4503,Servicios De Aerotransportacion De Aguascalientes,\N,,STD,AERO AGUASCALINETES,Mexico,N
4504,Semitool Europe,\N,,STE,SEMITRANS,United Kingdom,N
4505,SFT-Sudanese Flight,\N,,STF,,Sudan,N
4506,Sedalia,\N,,STG, Marshall, Boonville Stage Line,N
4507,South-Airlines,\N,,STH,,Armenia,N
4508,Sontair,\N,,STI,SONTAIR,Canada,N
4509,Sella Aviation,\N,,STJ,STELLAVIA,Netherlands,N
4510,Stapleford Flight Centre,\N,,STL,STAPLEFORD,United Kingdom,N
4511,Streamline Ops,\N,,STO,SLOPS,Russia,N
4512,Star Air,\N,,STQ,STERA,Indonesia,N
4513,Servicios de Transportes A,\N,FS,STU,FUEGUINO,Argentina,Y
4514,Star African Air,\N,,STU,STARSOM,Somali Republic,N
4515,Southern Aviation,\N,,STV,SOUTHERN AVIATION,Ghana,N
4516,South West Air Corporation,\N,,STW,SIERRA WHISKEY,Philippines,N
4517,Stars Away Aviation,\N,,STX,STARSAWAY,South Africa,N
4518,Styrian Airways,\N,,STY,STYRIAN,Austria,N
4519,Silesia Air,\N,,SUA,AIR SILESIA,Czech Republic,N
4520,Suburban Air Freight,\N,,SUB,SUB AIR,United States,N
4521,Sudan Airways,\N,SD,SUD,SUDANAIR,Sudan,Y
4522,Sun Air (Fiji),\N,PI,SUF,SUNFLOWER,Fiji,N
4523,Sunu Air,\N,,SUG,SUNU AIR,Senegal,N
4524,Sun Light,\N,,SUH,LIGHT AIR,Kyrgyzstan,N
4525,Swiss Air Force,\N,,SUI,SWISS AIR FORCE,Switzerland,N
4526,Superior Aviation Services,\N,,SUK,SKYCARGO,Kenya,N
4527,State Unitary Air Enterprise,\N,,SUM,SUMES,Russia,N
4528,Sun Air,\N,,SUR,,Egypt,N
4529,Sun Air of Scandinavia,\N,EZ,SUS,SUNSCAN,Denmark,N
4530,Sistemas Aeronauuticos 2000,\N,,SUT,SISTEMAS AERONAUTICOS,Mexico,N
4531,Star West Aviation,\N,,SUU,SUNSTAR,United States,N
4532,Sundance Air,\N,,SUV,DANCEAIR,Venezuela,N
4533,Saudi Arabian Airlines,\N,SV,SVA,SAUDIA,Saudi Arabia,Y
4534,St. Vincent Grenadines Air (1990),\N,,SVD,GRENADINES,Saint Vincent and the Grenadines,N
4535,Swedish Armed Forces,\N,,SVF,SWEDEFORCE,Sweden,N
4536,Sahel Airlines,\N,,AWJ,SAHEL AIRLINES,Niger,N
4537,Sterling Helicopters,\N,,SVH,SILVER,United Kingdom,N
4538,Servicios De Transporte Aereo,\N,,SVI,SETRA,Mexico,N
4539,Sabre Pacific,\N,,APD,,Australia,N
4540,Silver Air,\N,,SVJ,,Djibouti,N
4541,Sevastopol-Avia,\N,,SVL,SEVAVIA,Ukraine,N
4542,Savanair (Angola),\N,,SVN,SAVANAIR,Angola,N
4543,Servicios Aeronauticos De Oriente,\N,,SVO,SERVIORIENTE,Mexico,N
4544,Servicios Aereos Saar,\N,,SVS,AEREOS SAAR,Mexico,N
4545,Seven Four Eight Air Services,\N,,SVT,SIERRA SERVICES,Luxembourg,N
4546,Security Aviation,\N,,SVX,SECURITY AIR,United States,N
4547,Southwest Airlines,\N,WN,SWA,SOUTHWEST,United States,Y
4548,Swissboogie Parapro,\N,,SWB,SWISSBOOGIE,Switzerland,N
4549,South West Air,\N,,SWC,SAINT CLAIR,Canada,N
4550,Southern Winds Airlines,\N,A4,SWD,SOUTHERN WINDS,Argentina,Y
4551,Swedeways,\N,,SWE,SWEDELINE,Sweden,N
4552,Spurling Aviation,\N,,ASL,AIR SEATTLE,United States,N
4553,Sunwing Airlines,\N,WG,SWG,SUNWING,Canada,N
4554,Sunworld Airlines,\N,,SWI,SUNWORLD,United States,N
4555,Stateswest Airlines,\N,,SWJ,STATES,United States,N
4556,Surninam International Victory Airline,\N,,SWO,SIVA,Suriname,N
4557,Star Work Sky,\N,,SWP,STAR WORK,Italy,N
4558,Swift Air (Interstate Equipment Leasing),\N,,SWQ,SWIFTFLIGHT,United States,N
4559,Swiss International Air Lines,Swiss Airlines,LX,SWR,SWISS,Switzerland,Y
4560,Swissair,\N,SR,SWR,Swissair,Switzerland,Y
4561,Sunwest Aviation (Lindquist Investment),\N,,SWS,SUNNY WEST,United States,N
4562,Swiftair,\N,,SWT,SWIFT,Spain,N
4563,Swiss European Air Lines,Swiss European,,SWU,EUROSWISS,Switzerland,Y
4564,Swe Fly,\N,WV,SWV,FLYING SWEDE,Sweden,Y
4565,Shovkoviy Shlyah,\N,S8,SWW,WAY AERO,Ukraine,N
4566,Swazi Express Airways,\N,,SWX,SWAZI EXPRESS,Swaziland,N
4567,Sky Jet,\N,,SWY,SWISSLINK,Switzerland,N
4568,Servair,\N,,SWZ, Private Charter,SWISSBIRD,N
4569,Southern Cross Aviation,\N,,SXA,FERRY,United States,N
4570,Sky Exec Aviation Services,\N,,SXC,SKY EXEC,Nigeria,N
4571,Southeast Express Airlines,\N,,SXE,DOGWOOD EXPRESS,United States,N
4572,Servicios Aereos Sepecializados Mexicanos,\N,,SXM,SERVIMEX,Mexico,N
4573,SunExpress,\N,XQ,SXS,SUNEXPRESS,Turkey,Y
4574,Servicios De Taxi Aereo,\N,,SXT,SERTA,Mexico,N
4575,Satellite Aero,\N,,SXX,SATELLITE EXPRESS,United States,N
4576,Safari Express Cargo,\N,,SXY,SAFARI EXPRESS,Kenya,N
4577,Skyways,\N,,SYA,LINEAS CARDINAL,Argentina,N
4578,Systec 2000,\N,,SYC,SYSTEC,United States,N
4579,Sheba Aviation,\N,,SYE,,Yemen,N
4580,Sky One Express Airlines,\N,,SYF,SKY FIRST,United States,N
4581,Synergy Aviation,\N,,SYG,SYNERGY,United Kingdom,N
4582,Sonalysts,\N,,SYI,,United States,N
4583,Slate Falls Airways,\N,,SYJ,,Canada,N
4584,Satsair,\N,,SYK,AEROCAB,United States,N
4585,Syncrude Canada,\N,,SYN,SYNCRUDE,Canada,N
4586,Syrian Arab Airlines,\N,RB,SYR,SYRIANAIR,Syrian Arab Republic,Y
4587,Shawbury Flying Training Unit,\N,,SYS,SHAWBURY,United Kingdom,N
4588,Special Aviation Systems,\N,,SYV,SPECIAL SYSTEM,United States,N
4589,Skywalk Airlines,\N,AL,SYX,SKYWAY-EX,United States,Y
4590,Silk Way Airlines,\N,ZP,AZQ,SILK LINE,Azerbaijan,N
4591,South African Historic Flight,\N,,SYY,SKY COACH,South Africa,N
4592,Servicios Aeronauticos Z,\N,,SZT,AERO ZEE,Mexico,N
4593,Specavia Air Company,\N,,BHV,AVIASPEC,Russia,N
4594,Starair,\N,,BLY,BLARNEY,Ireland,N
4595,Sundance Air,\N,,BNC,BARNACLE AIR,United States,N
4596,Samara Airlines,\N,E5,BRZ,BERYOZA,Russia,N
4597,Spark+ Joint-Stock Company,\N,,BVV,SPARK,Russia,N
4598,Swedish Civil Aviation Administration,\N,,CBN,CALIBRATION,Sweden,N
4599,Shandong Airlines,\N,SC,CDG,SHANDONG,China,Y
4600,Spectrem Air,\N,,CDS,SPECDAS,South Africa,N
4601,Servicios Aereos Centrales,\N,,CEE,CENTRA AEREOS,Mexico,N
4602,Swedish Airlines,\N,,CFL,SWEDISH,Sweden,N
4603,Seagle Air,\N,,CGL,SEAGLE,Slovakia,N
4604,Sirius-Aero,\N,,CIG,SIRIUS AERO,Russia,N
4605,Sunwest Home Aviation,\N,,CNK,CHINOOK,Canada,N
4606,SAS Braathens,\N,,CNO,SCANOR,Norway,Y
4607,Spring Airlines,\N,9S,CQH,AIR SPRING,China,Y
4608,Sichuan Airlines,\N,3U,CSC,SI CHUAN,China,Y
4609,Shanghai Airlines,\N,FM,CSH,SHANGHAI AIR,China,Y
4610,Shuangyang General Aviation,\N,,CSY,SHUANGYANG,China,N
4611,Shenzhen Airlines,\N,ZH,CSZ,SHENZHEN AIR,China,Y
4612,Shanxi Airlines,\N,8C,CXI,SHANXI,China,N
4613,Sioux Falls Aviation,\N,,DKT,DAKOTA,United States,N
4614,Servicios Aereos Elite,\N,,DKY,DAKOY,Spain,N
4615,Servicios Aereos Denim,\N,,DNI,AERO DENIM,Mexico,N
4616,Swiss Eagle,\N,,EAB,SWISS EAGLE,Switzerland,N
4617,Skypower Express Airways,\N,,EAN,NIGERIA EXPRESS,Nigeria,N
4618,Scenic Air,\N,,ENR,,Namibia,N
4619,Sun D'Or,\N,7L,ERO,ECHO ROMEO,Israel,Y
4620,SkyEurope,\N,NE,ESK,RELAX,Slovakia,Y
4621,Sunshine Express Airlines,\N,CQ,EXL,,Australia,N
4622,South African Express,\N,,EXY,EXPRESSWAYS,South Africa,N
4623,Stuttgarter Flugdienst,\N,,FFD,FIRST FLIGHT,Germany,N
4624,Shalom Air Services,\N,,FFH,PEACE AIR,Nigeria,N
4625,Silverjet,\N,,FJE,ENVOY,United Kingdom,N
4626,Sky Bus,\N,,FLH,MILE HIGH,United States,N
4627,South Coast Aviation,\N,,GAD,SOUTHCOAST,United Kingdom,N
4628,Servicios Aereos Gadel,\N,,GDE,GADEL,Mexico,N
4629,S.P. Aviation,\N,,GDG,GOLDEN GATE,United States,N
4630,Seba Airlines,\N,,GIK,SEBA,Guinea,N
4631,Servicios Aereos Gana,\N,,GNA,SERVIGANA,Mexico,N
4632,Sky Wings Airlines,\N,,GSW,,Greece,N
4633,Star XL German Airlines,\N,,GXL,STARDUST,Germany,N
4634,Skyhaul,\N,,HAU,SKYHAUL,Japan,N
4635,Starship,\N,,HIP,STARSA,Mexico,N
4636,Servicios Ejecutivos Gosa,\N,,HJE,GOSA,Mexico,N
4637,Superior Aviation,\N,SO,HKA,SPEND AIR,United States,N
4638,Samaritan Air Service,\N,,HLO,HALO,Canada,N
4639,Skyraidybos Mokymo Centras,\N,,HRI,HELIRIM,Lithuania,N
4640,Sky Europe Airlines,\N,,HSK,MATRA,Slovakia,Y
4641,Svenska Direktflyg,\N,,HSV,HIGHSWEDE,Sweden,N
4642,Sky Helicopteros,\N,,HSY,HELISKY,Spain,N
4643,Skytaxi,\N,,IGA,IGUANA,Poland,N
4644,Silvair,\N,,IJS,,United States,N
4645,Servicios Aereos Ilsa,\N,,ILS,SERVICIOS ILSA,Mexico,N
4646,Sincom-Avia,\N,,INK,SINCOM AVIA,Ukraine,N
4647,Safat Airlines,\N,,IRV,SAFAT AIR,Iran,N
4648,Saha Airlines Services,\N,,IRZ,SAHA,Iran,N
4649,Sunline Express,\N,,JAM,SUNTRACK,Kenya,N
4650,Secure Air Charter,\N,,JCM,SECUREAIR,United States,N
4651,Sark International Airways,\N,,JIM,SARK,United Kingdom,N
4652,Spanair,\N,JK,JKK,SPANAIR,Spain,Y
4653,Salem,\N,,KKS,KOKSHE,Kazakhstan,N
4654,Servicios Aereos Copters,\N,,KOP,COPTERS,Chile,N
4655,Servicios Aereos Expecializados En Transportes Petroleros,\N,,KSP,SAEP,Colombia,N
4656,Skybridge Airops,\N,,KYB,SKY AIROPS,Italy,N
4657,Sky Aeronautical Services,\N,,KYR,SKY AERONAUTICAL,Mexico,N
4658,Servicios Aereos Ejecutivos De La Laguna,\N,,LGU,LAGUNA,Mexico,N
4659,Servico Leo Lopex,\N,,LLA,LEO LOPOZ,Mexico,N
4660,Servicios Aereos Estrella,\N,,LLS,SERVIESTRELLA,Mexico,N
4661,South African Air Force,\N,,LMG,SOUTH AFRICAN,South Africa,N
4662,Sky Limo Corporation,\N,,LMO,SKY LIMO,United States,N
4663,SANSA,\N,,LRS,,Costa Rica,N
4664,Spectrum Aviation Incorporated,\N,,LSP,AIR TONY,United Kingdom,N
4665,SOS Helikoptern Gotland,\N,,MCG,MEDICOPTER,Sweden,N
4666,Sundt Air,\N,,MDT,MIDNIGHT,Norway,N
4667,Servicios Aereos Milenio,\N,,MLO,MILENIO,Mexico,N
4668,SAAD (A320) Limited,\N,,MMS,MUSAAD AIR,Cayman Islands,N
4669,Servicios Aereos Moritani,\N,,MRI,MORITANI,Mexico,N
4670,San Juan Airlines,\N,2G,MRR,MARINER,United States,N
4671,Servico Aereo Regional,\N,,MSG,SAR-REGIONAL,Mozambique,N
4672,Servicio De Viglancia Aerea Del Ministerio De Seguridad Publica,\N,,MSP,SEGURIDAD,Costa Rica,N
4673,Servicios Aereos MTT,\N,,MTG,,Mexico,N
4674,Sabre Pacific,\N,1Z,,,Australia,N
4675,Sabre,\N,1S,,,United States,N
4676,Sierra Nevada Airlines,\N,1I,,,United States,N
4677,Siren-Travel,\N,1H,,,Russia,N
4678,Sirena,\N,1Q,,,Russia,N
4679,Sky Trek International Airlines,\N,1I,,,,N
4680,Southern Cross Distribution,\N,1K,,,Australia,N
4681,Sutra,\N,1K,,,United States,N
4682,SNCF,\N,2C,,,France,N
4683,Star Equatorial Airlines,\N,2S,,,Guinea,N
4684,Seulawah Nad Air,\N,,NAD,SEULAWAH,Indonesia,N
4685,Servicios Aereos del Nazas SA de CV,\N,,NAZ,NAZAS,Mexico,N
4686,Simpson Air Ltd,\N,,NCS,COMMUTER-CANADA,Canada,N
4687,Spirit Airlines,\N,NK,NKS,SPIRIT WINGS,United States,Y
4688,Servicios Aereos Latinoamericanos,\N,,NON,SERVICIOS LATINO,Mexico,N
4689,Servicios Aereos Monarrez,\N,,NRZ,MONARREZ,Mexico,N
4690,Societe De Transport Aerien De Mauritanie,\N,,NSC,TRANS-SOCIETE,Mauritania,N
4691,SATENA,\N,9R,NSE,SATENA,Colombia,Y
4692,Servicios Aereos Del Norte,\N,,NTB,SERVINORTE,Mexico,N
4693,Servicios Intergrales De Aviacion,\N,,NTG,INTERGRALES,Mexico,N
4694,Slok Air Gambia,\N,S0,OKS,SLOK GAMBIA,Gambia,N
4695,Soko Aviation,\N,,OKT,SOKO AIR,Spain,N
4696,Solar Cargo,\N,,OLC,SOLARCARGO,Venezuela,N
4697,Soloflight,\N,,OLO,SOLO,United Kingdom,N
4698,Sonnig SA,\N,,ONG,SONNIG,Switzerland,N
4699,Sosoliso Airlines,\N,SO,OSL,SOSOLISO,Nigeria,N
4700,Servicios Aereos Noticiosos,\N,,OSS,NOTICIOSOS,Mexico,N
4701,South Airlines,\N,,OTL,SOUTHLINE,Ukraine,N
4702,Skywest Airlines,\N,,OZW,OZWEST,Australia,Y
4703,Sokol,\N,,PIV,AEROSOKOL,Russia,N
4704,South Carolina Aeronautics Commission,\N,,PLT,PALMETTO,United States,N
4705,Servicios Aereos Premier,\N,,PMR,SERVICIOS PREMIER,Mexico,N
4706,Survey Udara (Penas),\N,,PNS,PENAS,Indonesia,N
4707,Servicios Aereos Poblanos,\N,,POB,POBLANOS,Mexico,N
4708,Servicios Aereos Profesionales,\N,,PSV,PROSERVICIOS,Dominican Republic,N
4709,Southeastern Airways,\N,,PTM,POSTMAN,United States,N
4710,Spurwing Airlines,\N,,PUR,SPURWING,South Africa,N
4711,Sky Trek International Airlines,\N,1I,PZR,PHAZER,United States,N
4712,Shandong Airlines Rainbow Jet,\N,,RBW,CAI HONG,China,N
4713,SA Airlink Regional,\N,,REJ,REGIONAL LINK,South Africa,N
4714,Servicio Aereo Regional Regair,\N,,RER,REGAIR,Ecuador,N
4715,Scoala Superioara De Aviatie Civila,\N,,RFT,ROMANIAN ACADEMY,Romania,N
4716,Servicios Aereos Regiomontanos,\N,,RGC,REGIOMONTANO,Mexico,N
4717,S-Air,\N,,RLS,S-AIRLINES,Russia,N
4718,Servicios De Rampa Y Mostrador,\N,,RMP,SERAMSA,Mexico,N
4719,SNAS Aviation,\N,,RSE,RED SEA,Saudi Arabia,N
4720,Skybus Airlines,\N,SX,SKB,SKYBUS,United States,N
4721,Skymaster Air Lines,\N,,SKC,SKYMASTER AIR,United States,N
4722,Sky Harbor Air Service,\N,,SKD,SKY DAWG,United States,N
4723,Sky Tours,\N,,SKE,SKYISLE,United States,N
4724,Sakaviaservice,\N,,AZG,SAKSERVICE,Georgia,N
4725,Skycraft,\N,,SKF,SKYCRAFT,United States,N
4726,Skycraft Air Transport,\N,,SKG,SKYCRAFT-CANADA,Canada,N
4727,Skyking,\N,,SKI,SKYKING,United Kingdom,N
4728,SkyKing Turks and Caicos Airways,\N,RU,SKI,SKYKING,Turks and Caicos Islands,N
4729,Skylink Aviation,\N,,SKK,SKYLINK,Canada,N
4730,Skycharter (Malton),\N,,SKL,SKYCHARTER,Canada,N
4731,Skyline Aviation Services,\N,,SKN,SKYLINER,United States,N
4732,Scottish Airways Flyers,\N,,SKO,SKYWORK,United Kingdom,N
4733,Skyscapes Air Charters,\N,,SKR,SKYSCAPES,South Africa,N
4734,Sky Service,\N,,SKS,SKY SERVICE,Belgium,N
4735,Santa Barbara Airlines,\N,S3,BBR,SANTA BARBARA,Venezuela,Y
4736,Skystar Airways,\N,,SKT,SKY YOU,Thailand,N
4737,Sky Airline,\N,H2,SKU,AEROSKY,Chile,Y
4738,SkyWest,\N,OO,SKW,SKYWEST,United States,Y
4739,Skyways Express,\N,JZ,SKX,SKY EXPRESS,Sweden,Y
4740,Skymark Airlines,\N,BC,SKY,SKYMARK,Japan,Y
4741,Skyway Enterprises,\N,,SKZ,SKYWAY-INC,United States,N
4742,Sierra National Airlines,\N,LJ,SLA,SELAIR,Sierra Leone,N
4743,Slok Air,\N,,SLB,SLOK AIR,Nigeria,N
4744,Silver Air,\N,,SLD,SILVERLINE,Czech Republic,N
4745,Streamline,\N,,SLE,SLIPSTREAM,South Africa,N
4746,Starfly,\N,,SLF,ELISTARFLY,Italy,N
4747,Saskatchewan Government,\N,,SLG,LIFEGUARD,Canada,N
4748,Silverhawk Aviation,\N,,SLH,SILVERHAWK,United States,N
4749,Servicios Aereos de Los Angeles,\N,,AGE,AEROANGEL,Mexico,N
4750,SilkAir,\N,MI,SLK,SILKAIR,Singapore,Y
4751,Slovak Airlines,\N,6Q,SLL,SLOV LINE,Slovakia,N
4752,Surinam Airways,\N,PY,SLM,SURINAM,Suriname,Y
4753,Sloane Aviation,\N,,SLN,SLOANE,United Kingdom,N
4754,Salpa Aviation,\N,,SLP,SALPA,Sudan,N
4755,Servicios Aereos Slainte,\N,,SLS,SERVICIOS SLAINTE,Mexico,N
4756,Stella Aviation,\N,,SLV,AVISTELLA,Mauritania,N
4757,Salama Airlines Nigeria,\N,,SLW,SALMA AIR,Nigeria,N
4758,Sete Linhas Aereas,\N,,SLX,SETE,Brazil,N
4759,Sky Line for Air Services,\N,,SLY,SKYCO,Sudan,N
4760,Super Luza,\N,,SLZ,LUZA,Angola,N
4761,SMA Airlines,\N,,SMA,SESAME,Nigeria,N
4762,Sabang Merauke Raya Air Charter,\N,,SMC,SAMER,Indonesia,N
4763,Servicios Aereos La Marquesa,\N,,SMD,SERVICIOS MARQUESA,Mexico,N
4764,Servant Air,\N,8D,,,United States,N
4765,Semos,\N,,SME,SEMICH,Kazakhstan,N
4766,Smalandsflyg,\N,,SMF,GORDON,Sweden,N
4767,Smithair,\N,,SMH,SMITHAIR,United States,N
4768,Semeyavia,\N,,SMK,ERTIS,Kazakhstan,N
4769,Smith Air (1976),\N,,SML,SMITH AIR,Canada,N
4770,Summit Airlines,\N,,SMM,SUMMIT-AIR,United States,N
4771,Samar Air,\N,,SMQ,SAMAR AIR,Tajikistan,N
4772,Somon Air,\N,,SMR,SOMON AIR,Tajikistan,N
4773,Skyline,\N,,SMT,SKYLIMIT,Nigeria,N
4774,Servicios Aerecs Del Sol,\N,,AOS, S.A. de C.V.,AEROSOL,N
4775,Senator Aviation Charter,\N,,SNA,SENATOR,Germany,N
4776,Sterling Airlines,\N,NB,SNB,STERLING,Denmark,Y
4777,Servicios Aereos De Nicaragua (SANSA),\N,,SNE,SANSA,Nicaragua,N
4778,Shans Air,\N,,SNF,SHANS AIR,Russia,N
4779,Senair Services,\N,,SNH,SENSERVICE,Senegal,N
4780,Savanah Airlines,\N,,SNI,SAVANAHLINE,Nigeria,N
4781,Skynet Asia Airways,\N,6J,SNJ,NEWSKY,Japan,Y
4782,Southeast Airlines (Sun Jet International),\N,,SNK,SUN KING,United States,N
4783,Soonair Lines,\N,,SNL,SOONAIR,United States,N
4784,Servizi Aerei,\N,,SNM,SERVIZI AEREI,Italy,N
4785,Sun Pacific International,\N,,SNP,SUN PACIFIC,United States,N
4786,Sun Quest Executive Air Charter,\N,,SNQ,EXECU-QUEST,United States,N
4787,Societe Centrafricaine De Transport Aerien,\N,,SNS,,Central African Republic,N
4788,Suncoast Aviation,\N,,SNT,SUNCOAST,United States,N
4789,Snunit Aviation,\N,,SNU,,Israel,N
4790,Sudanese States Aviation,\N,,SNV,SUDANESE,Sudan,N
4791,Sun West Airlines,\N,,SNW,SUN WEST,United States,N
4792,Sun Air Aviation Services,\N,,SNX,SUNEX,Canada,N
4793,Stabo Freight,\N,,SOB,STABO,Zambia,N
4794,Southern Cargo Air Lines,\N,,SOC,,Russia,N
4795,Southern Ohio Aviation Sales,\N,,SOH,SOUTHERN OHIO,United States,N
4796,Southern Aviation,\N,,SOI,SOAVAIR,Zambia,N
4797,Solomon Airlines,\N,IE,SOL,SOLOMON,Solomon Islands,Y
4798,Somali Airlines,\N,,SOM,SOMALAIR,Somali Republic,N
4799,Sunshine Air Tours,\N,,SON,SUNSHINE TOURS,United States,N
4800,Southern Air,\N,,SOO,SOUTHERN AIR,United States,N
4801,Solinair,\N,,SOP,SOLINAIR,Slovenia,N
4802,Sonair Servico Aereo,\N,,SOR,SONAIR,Angola,N
4803,Southeast Correct Craft,\N,,SOT,SOUTH COURIER,United States,N
4804,Southern Airways,\N,,SOU,SOUTHERN EXPRESS,United States,Y
4805,Saratov Aviation Division,\N,6W,SOV,SARATOV AIR,Russia,Y
4806,Sowind Air,\N,,SOW,SOWIND,Canada,N
4807,Solid Air,\N,,SOX,SOLIDAIR,Netherland,N
4808,Sat Airlines,\N,HZ,SOZ,SATCO,Kazakhstan,Y
4809,Sierra Pacific Airlines,\N,,SPA,SIERRA PACIFIC,United States,N
4810,Springbok Classic Air,\N,,SPB,SPRING CLASSIC,South Africa,N
4811,Skyworld Airlines,\N,,SPC,PORT,United States,N
4812,Sprague Electric Company,\N,,SPE,SPRAGUE,United States,N
4813,Space World Airline,\N,,SPF,SPACE WORLD,Nigeria,N
4814,Springdale Air Service,\N,,SPG,SPRING AIR,United States,N
4815,Sapphire Executive Air,\N,,SPH,SAPPHIRE-CHARTER,South Africa,N
4816,South Pacific Island Airways,\N,,SPI,SOUTH PACIFIC,United States,Y
4817,Servicios Corporativos Aereos De La Laguna,\N,,SPL,CORPORATIVOS LAGUNA,Mexico,N
4818,Skorpion Air,\N,,SPN,AIR SKORPIO,Bulgaria,N
4819,Sapphire Aviation,\N,,SPP,SAPPHIRE,United States,N
4820,Servicios Aereos Palenque,\N,,SPQ,SERVICOS PALENQUE,Mexico,N
4821,Servicios Aereos Tribasa,\N,,TBS,TRIBASA,Mexico,N
4822,Shuttle America,\N,S5,TCF,MERCURY,United States,Y
4823,SAAB Nyge Aero,\N,,TGT,TARGET,Sweden,N
4824,Spark Air,\N,,THB,THAI SABAI,Thailand,N
4825,S C Ion Tiriac,\N,,TIH,TIRIAC AIR,Romania,N
4826,Starlite Aviation,\N,,TRL,STARSTREAM,South Africa,N
4827,Servicios Aereos Corporativos,\N,,TRN,AEROTRON,Mauritania,N
4828,Societe Tout Tranport Mauritanien,\N,,TTM,TOUT-AIR,Mauritania,N
4829,Servicios Aereos Tamazula,\N,,TZU,TAMAZULA,Mexico,N
4830,Shar Ink,\N,,UGP,SHARINK,Russia,N
4831,Second Sverdlovsk Air Enterprise,\N,,UKU,PYSHMA,Russia,N
4832,Servicios Aereos Universitarios,\N,,UNT,UNIVERSITARIO,Mexico,N
4833,Skif-Air,\N,,USK,SKIF-AIR,Ukraine,N
4834,Smarkand Aero Servise,\N,,USN,SAMAS,Uzbekistan,N
4835,Samarkand Airways,\N,,UZS,SOGDIANA,Uzbekistan,N
4836,Servicios Aereos Avandaro,\N,,VDO,AVANDARO,Mexico,N
4837,Stichting Vliegschool 16Hoven,\N,,VGS,SMART,Netherlands,N
4838,Silverback Cargo Freighters,\N,,VRB,SILVERBACK,Rwanda,N
4839,Sirvair,\N,,VRS,VAIRSA,Mexico,N
4840,Scat Air,,DV,VSV,VLASTA,Kazakhstan,Y
4841,Sunset Aviation,\N,,VXN,VIXEN,United States,N
4842,Sport Air Travel,\N,,WCC,WEST COAST,United States,N
4843,Swift Copters,\N,,WFC,SWIFTCOPTERS,Switzerland,N
4844,Skyrover CC,\N,,WLK,SKYWATCH,South Africa,N
4845,Safarilink Aviation,\N,,XLK,SAFARILINK,Kenya,N
4846,SENEAM,\N,,XMX,SENEAM,Mexico,N
4847,Southport Air Service,\N,,XPG,,United States,N
4848,Spectrum Air Service,\N,,XSA,,United States,N
4849,Stephenville Aviation Services,\N,,XSN,,Canada,N
4850,Servicios Aereos Textra,\N,,XTA,TEXTRA,Mexico,N
4851,Sector Airlines,\N,,XTR,EXTER,Canada,N
4852,Skyplan Services,\N,,XXS,,Canada,N
4853,Stewart Aviation Services,\N,,YBE,YELLOW BIRD,United States,N
4854,Sirin,\N,R1,,,,N
4855,Star Air,\N,S6,,,Denmark,N
4856,Servicios Aereos Integrales / Flyant,\N,,FYA,FLYANT,Spain,N
4857,State Flight Academy of Ukraine,\N,,UFA,FLIGHT ACADEMY,Ukraine,N
4860,Trans Jet Airways,\N,,SWL,TRANSJET,Sweden,N
4861,Turbot Air Cargo,\N,,TAC,TURBOT,Senegal,N
4862,Tranporte Aereo Dominicano,\N,,TAD,TRANS DOMINICAN,Dominican Republic,N
4863,TAME,\N,EQ,TAE,TAME,Ecuador,Y
4864,Trans International Express Aviation,\N,,BAP,BIG APPLE,United States,N
4865,TAG Aviation USA,\N,,TAG,TAG U-S,United States,N
4866,Talair,\N,,TAL,TALAIR,Papua New Guinea,N
4867,TAM Brazilian Airlines,\N,JJ,TAM,TAM,Brazil,Y
4868,Transaustralian Air Express,\N,,AUC,AUSCARGO,Australia,N
4869,TAP Portugal,TAP Air Portugal,TP,TAP,AIR PORTUGAL,Portugal,Y
4870,Tunisair,\N,TU,TAR,TUNAIR,Tunisia,Y
4871,Transportes Aereos Tauro,\N,,TAU,TRANSTAURO,Mexico,N
4872,Travelair,\N,,TAX,TRAVELAIR,Germany,N
4873,TNT Airways,\N,3V,TAY,QUALITY,Belgium,N
4874,Transbrasil,\N,,TBA,TRANSBRASIL,Brazil,N
4875,Turbine Air Cargo UK,\N,,TBC,,United Kingdom,N
4876,Thunderbird Tours,\N,,TBD,ORCA,Canada,N
4877,Tropical Airways,\N,M7,TBG,,Haiti,N
4878,Trinity Air Bahamas,\N,,TBH,,Bahamas,N
4879,TAB Express International,\N,,TBI,TAB INTERNATIONAL,United States,N
4880,Taban Air Lines,\N,,TBM,TABAN AIR,Iran,N
4881,Teebah Airlines,\N,,TBN,TEEBAH,Sierra Leone,N
4882,Tubelair,\N,,TBR,TUBELAIR,Tunisia,N
4883,Tobago Express,\N,,TBX,TABEX,Trinidad and Tobago,N
4884,Tropican Air Services,\N,,TCA,TROPICANA,Egypt,N
4885,Transporte del Caribe,\N,,TCB,TRANSCARIBE,Colombia,N
4886,Trans Continental Airlines,\N,,TCC,TRANSCAL,Sudan,N
4887,Tchad Airlines,\N,,TCD,TCHADLINES,Chad,N
4888,Trans-Colorado Airlines,\N,,TCE,TRANS-COLORADO,United States,N
4889,Thai Air Cargo,\N,T2,TCG,THAI CARGO,Thailand,Y
4890,Transcontinental Air,\N,,TCH,TRANS GULF,Bahrain,N
4891,Teledyne Continental Motors,\N,,TCM,TELEDYNE,United States,N
4892,Trans Continental Airlines,\N,,TCN,TRANSCON,United States,N
4893,Transcorp Airways,\N,,TCP,TRANSCORP,United Kingdom,N
4894,Transcontinental Sur,\N,,TCT,TRANS-CONT,Uruguay,N
4895,Transglobal Airways Corporation,\N,,TCU,TRANSGLOBAL,Philippines,N
4896,Thomas Cook Airlines,\N,FQ,TCW,THOMAS COOK,Belgium,Y
4897,Thomas Cook Airlines,\N,MT,TCX,KESTREL,United Kingdom,Y
4898,Twin Cities Air Service,\N,,TCY,TWIN CITY,United States,N
4899,Tadair,\N,,TDC,TADAIR,Spain,N
4900,Tellavia / Flight One,\N,,TDE,TELLURIDE,United States,N
4901,Transportes Aereos de Ixtlan,\N,,TDI,TRANSIXTLAN,Mexico,N
4902,Tandem Aero,\N,TQ,TDM,TANDEM,Moldova,N
4903,TRADO,\N,,TDO,TRADO,Dominican Republic,N
4904,Trade Air,\N,,TDR,TRADEAIR,Croatia,N
4905,Taxi Aero Nacional Del Evora,\N,,TDV,TAXI EVORA,Mexico,N
4906,Tradewinds Airlines,\N,,TDX,TRADEWINDS EXPRESS,United States,N
4907,Tenir Airlines,\N,,TEB,TENIR AIR,Kyrgyzstan,N
4908,Teamline Air,\N,L9,TLW,Teamline,Austria,N
4909,Tecnicas Fotograficas,\N,,TEF,TECFOTO,Spain,N
4910,Tempelhof Airways,\N,,TEH,TEMPELHOF,United States,N
4911,Telford Aviation,\N,,TEL,TELFORD,United States,N
4912,Tech-Mont Helicopter Company,\N,,TEM,TECHMONT,Slovakia,N
4913,Tennessee Airways,\N,,TEN,TENNESSEE,United States,N
4914,Territorial Airlines,\N,,TER,TERRI-AIRE,United States,N
4915,TAES Transportes Aereos de El Salvador,\N,,TES,TAES,El Salvador,N
4916,Transeuropean Airlines,\N,UE,TEP,TRANSEURLINE,Russia,N
4917,Tepavia-Trans Airlines,\N,,TET,TEPAVIA,Moldova,N
4918,Trans-Florida Airlines,\N,,TFA,TRANS FLORIDA,United States,N
4919,Tair Airways,\N,,TFB,ROYAL TEE-AIR,Philippines,N
4920,Talon Air,\N,,TFF,TALON FLIGHT,United States,N
4921,Trabajos Aereos Murcianos,\N,,AIM,PIJO,Spain,N
4922,Thai Flying Helicopter Service,\N,,TFH,THAI HELICOPTER,Thailand,N
4923,Transport Facilitators,\N,,TFI,,United States,N
4924,Transafrik International,\N,,TFK,,Sao Tome and Principe,N
4925,Transportes Aereos del Pacifico,\N,,TFO,TRANSPORTES PACIFICO,Mexico,N
4926,Thai Flying Service,\N,,TFT,THAI FLYING,Thailand,N
4927,Transportes Aereos San Rafael,\N,,SRF,SAN RAFEAL,Chile,N
4928,Tayside Aviation,\N,,TFY,TAYSIDE,United Kingdom,N
4929,TG Aviation,\N,,TGC,THANET,United Kingdom,N
4930,Trabajos Aereos,\N,,TGE,TASA,Spain,N
4931,Transportes Aereos Regionales,\N,,TGI,TRANSPORTE REGIONAL,Mexico,N
4932,TAG Aviation Espana,\N,,TGM,TAG ESPANA,Spain,N
4933,Trigana Air Service,\N,,TGN,TRIGANA,Indonesia,Y
4934,Titan Airways,\N,ZT,AWC,ZAP,United Kingdom,N
4935,Transport Canada,\N,,TGO,TRANSPORT,Canada,N
4936,Tiger Airways,\N,TR,TGW,GO CAT,Singapore,Y
4937,Tiger Airways Australia,\N,TT,TGW,GO CAT,Australia,Y
4938,Transair Gabon],\N,,TGX,TRANSGABON,Gabon,N
4939,Trans Guyana Airways,\N,,TGY,TRANS GUYANA,Guyana,N
4940,Thai Airways International,\N,TG,THA,THAI,Thailand,Y
4941,Tar Heel Aviation,\N,,THC,TARHEEL,United States,N
4942,Toumai Air Tchad,\N,,THE,TOUMAI AIR,Chad,N
4943,Touraine Helicoptere,\N,,THF,TOURAINE HELICO,France,N
4944,Thai Global Airline,\N,,THG,THAI GLOBAL,Thailand,N
4945,Thai Jet Intergroup,\N,,THJ,THAI JET,Thailand,N
4946,Turk Hava Kurumu Hava Taksi Isletmesi,\N,,THK,HUR KUS,Turkey,Y
4947,Thai AirAsia,Thai Air Asia,FD,AIQ,THAI ASIA,Thailand,Y
4948,TACA De Honduras,\N,,THO,LEMPIRA,Honduras,N
4949,Tehran Airline,\N,,THR,TEHRAN AIR,Iran,N
4950,Thunder Airlines,\N,,THU,AIR THUNDER,Canada,N
4951,Turkish Airlines,\N,TK,THY,TURKAIR,Turkey,Y
4952,Trans Helicoptere Service,\N,,THZ,LYON HELIJET,France,N
4953,Trans International Airlines,\N,,TIA,TRANS INTERNATIONAL,United States,N
4954,Travel International Air Charters,\N,,TIC,TRAVEL INTERNATIONAL,Zambia,N
4955,Time Air,\N,,TIE,TIME AIR,Czech Republic,N
4956,Tic Air,\N,,TIK,TICAIR,Australia,N
4957,Tajikistan International Airlines,\N,,TIL,TIL,Tajikistan,Y
4958,TEAM Transportes Aereos,\N,,TIM,TEAM BRASIL,Brazil,N
4959,Taino Tours,\N,,TIN,TAINO,Dominican Republic,N
4960,Tesis,\N,,TIS,TESIS,Russia,N
4961,Transcarga Intl Airways,\N,,TIW,TIACA,Venezuela,N
4962,Tajikair,\N,,TJK,TAJIKAIR,Tajikistan,N
4963,Tien-Shan,\N,,TJN,NERON,Kazakhstan,N
4964,Tyrolean Jet Service,\N,,TJS,TYROLJET,Austria,N
4965,Twin Jet,\N,T7,TJT,TWINJET,France,Y
4966,Tikal Jets Airlines,\N,,TKC,TIKAL,Guatemala,N
4967,Take Air Line,\N,,TKE,ISLAND BIRD,France,N
4968,Tropical International Airways,\N,,TKX,TROPEXPRESS,Saint Kitts and Nevis,N
4969,Thai Sky Airlines,\N,9I,TKY,THAI SKY,Thailand,N
4970,Translift Airways,\N,,TLA,TRANSLIFT,Ireland,Y
4971,Transport Africa,\N,,TLF,TRANS-LEONE,Sierra Leone,N
4972,Trans Atlantic Airlines,\N,,TLL,ATLANTIC LEONE,Sierra Leone,N
4973,Eagle Canyon Airlines,\N,,TLO,TALON AIR,United States,N
4974,Tulip Air,\N,TD,TLP,TULIPAIR,Netherland,N
4975,Tunisavia,\N,,TAJ,TUNISAVIA,Tunisia,N
4976,TLC Air,\N,,TLS,TEALSY,United States,N
4977,Turtle Airways,\N,,TLT,TURTLE,Fiji,N
4978,Travelair,\N,,TLV,PAJAROS,Uruguay,N
4979,Telesis Transair,\N,,TLX,TELESIS,United States,N
4980,Top-Fly,\N,,TLY,TOPFLY,Spain,N
4981,Trans Mediterranean Airlines,\N,TL,TMA,TANGO LIMA,Lebanon,Y
4982,Transmed Airlines,\N,,TMD,,Egypt,N
4983,Tri-MG Intra Asia Airlines,\N,GY,TMG,TRILINES,Indonesia,N
4984,Taxis Turisticos Marakame,\N,,TMH,TAXIMARAKAME,Mexico,N
4985,Tamir Airways,\N,,TMI,TAMIRWAYS,Israel,N
4986,Tomahawk Airways,\N,,TMK,TOMAHAWK,United States,N
4987,TAM,\N,,TML,TAM AIRLINE,Madagascar,N
4988,TMC Airlines,\N,,TMM,WILLOW RUN,United States,N
4989,TRAM,\N,,TMQ,TRAM AIR,Mauritania,N
4990,Timberline Air,\N,,TMR,TIMBER,Canada,N
4991,Temsco Helicopters,\N,,TMS,TEMSCO,United States,N
4992,Trans Midwest Airlines,\N,,TMT,TRANS MIDWEST,United States,N
4993,Tramon Air,\N,,TMX,TRAMON,South Africa,N
4994,Transportes Aereos del Mundo Maya,\N,,TMY,MUNDO MAYA,Mexico,N
4995,Tranporte Amazonair,\N,,TMZ,TRANS AMAZON,Venezuela,N
4996,Trans Air-Benin,\N,,TNB,TRANS-BENIN,Benin,N
4997,Taxis Aereos del Noroeste,\N,,TNE,TAXINOROESTE,Mexico,N
4998,Transafricaine,\N,,TNF,TRANSFAS,Burkina Faso,N
4999,Tennessee Air National Guard 164th Airlift Group,\N,,TNG,,United States,N
5000,Transair International Linhas Aereas,\N,,TNI,TRANSINTER,United States,N
5001,Tengeriyn Ulaach Shine,\N,,TNL,SKY HORSE,Mongolia,N
5002,Tiara Air,\N,3P,TNM,TIARA,Aruba,Y
5003,Transped Aviation,\N,,TNP,TRANSPED,Austria,N
5004,Tanana Air Services,\N,,TNR,TAN AIR,United States,N
5005,Trans North Turbo Air,\N,,TNT,TRANS NORTH,Canada,N
5006,Transnorthern,\N,,TNV,TRANSNORTHERN,United States,N
5007,Trans Nation Airways,\N,,TNW,TRANS-NATION,Ethiopia,N
5008,Trener Ltd,\N,,TNX,TRAINER,Hungary,N
5009,Twin Town Leasing Company,\N,,TNY,TWINCAL,United States,N
5010,Tobruk Air,\N,7T,TOB,TOBRUK AIR,Libya,N
5011,TOJ Airlines,\N,,TOJ,TOJ AIRLINE,Tajikistan,N
5012,Tol-Air Services,\N,TI,TOL,TOL AIR,United States,N
5013,Thomsonfly,\N,BY,TOM,TOMSON,United Kingdom,Y
5014,Top Air,\N,,TOP,AIR TOP,Turkey,N
5015,Toronto Airways,\N,,TOR,TORONTAIR,Canada,N
5016,Tropic Air,\N,PM,TOS,TROPISER,Belize,Y
5017,Totavia,\N,,TOT,,Canada,N
5018,Tower Air,\N,FF,TOW,TEE AIR,United States,N
5019,Toyota Canada,\N,,TOY,TOYOTA,Canada,N
5020,TAMPA,\N,QT,TPA,TAMPA,Colombia,Y
5021,Top Speed,\N,,TPD,TOP SPEED,Austria,N
5022,Taxis Aereos del Pacifico,\N,,TPF,TAXIPACIFICO,Mexico,N
5023,Transportes Aereos Pegaso,\N,,TPG,TRANSPEGASO,Mexico,N
5024,TAR Interpilot,\N,,TPL,INTERPILOT,Mauritania,N
5025,Transpais Aereo,\N,,TPM,TRANSPAIS,Mexico,N
5026,Transportacion Aerea del Norte,\N,,TPN,AEREA DELNORTE,Mexico,N
5027,Transpac Express,\N,,TPP,TRANS EXPRESS,Australia,N
5028,Taxis Aereos De Parral,\N,,TPR,TAXIS PARRAL,Mexico,N
5029,TAPSA Transportes Aereos Petroleros,\N,,TPS,TAPSA,Argentina,N
5030,Transportes Aereo del Sureste,\N,,TPT,TASSA,Mexico,N
5031,Trans American Airlines (Trans Am),\N,,TPU,TRANS PERU,Peru,N
5032,Thai Pacific Airlines Business,\N,,TPV,THAI PACIFIC,Thailand,N
5033,Transportes Aereos De Xalapa,\N,,TPX,TRANSXALAPA,Mexico,N
5034,Trans-Provincial Airlines,\N,,TPY,TRANS PROVINCIAL,Canada,N
5035,Transportes La Paz,\N,,TPZ,TRANSPAZ,Mexico,N
5036,Taquan Air Services,\N,,TQN,TAQUAN,United States,N
5037,Transportacion Aerea De Queretaro,\N,,TQR,TRANSQUERETARO,Mexico,N
5038,TransAsia Airways,\N,GE,TNA,TransAsia,Taiwan,Y
5039,Transavia Holland,\N,HV,TRA,TRANSAVIA,Netherlands,Y
5040,Trans Air Charter,\N,,TRC,TRACKER,United States,N
5041,TACV,\N,VR,TCV,CABOVERDE,Portugal,Y
5042,Trans Island Air,\N,,TRD,TRANS ISLAND,Barbados,N
5043,Taxi Air Fret,\N,,TRF,TAXI JET,France,N
5044,TRAGSA (Medios Aereos),\N,,TRG,,Spain,N
5045,Trans Euro Air,\N,,TRJ,HIGH TIDE,United Kingdom,N
5046,Transport Aerien de Mauritanie,\N,,TRM,SOTRANS,Mauritania,N
5047,Tropic Airlines-Air Molokai,\N,,TRO,MOLOKAI,United States,N
5048,Turdus Airways,\N,,TRQ,HUNTER,Netherland,N
5049,Tramson Limited,\N,,TRR,TRAMSON,Sudan,N
5050,Trans Arabian Air Transport,\N,,TRT,TRANS ARABIAN,Sudan,N
5051,Triangle Airline (Uganda),\N,,TRU,TRI AIR,Uganda,N
5052,Transwestern Airlines of Utah,\N,,TRW,TRANS-WEST,United States,N
5053,Tristar Airlines,\N,,TRY,TRISTAR AIR,United States,N
5054,TransMeridian Airlines,\N,T9,TRZ,TRANS-MERIDIAN,United States,N
5055,Transair France,\N,,TSA,AIRTRAF,France,N
5056,TAF-Linhas Aereas,\N,,TSD,TAFI,Brazil,N
5057,Transmile Air Services,\N,TH,TSE,TRANSMILE,Malaysia,N
5058,Trans-Air-Congo,\N,,TSG,TRANS-CONGO,Republic of the Congo,N
5059,Transport'air,\N,,TSI,TRANSPORTAIR,France,N
5060,Trast Aero,\N,S5,TSJ,TRAST AERO,Kyrgyzstan,N
5061,Trast Aero,\N,,TSK,TOMSK-AVIA,Kyrgyzstan,N
5062,Thai Aviation Services,\N,,TSL,THAI AVIATION,Thailand,N
5063,Taftan Airlines,\N,,SBT,TAFTAN,Iran,N
5064,Transwest Air,\N,9T,ABS,ATHABASKA,Canada,Y
5065,Trans Sayegh Airport Services,\N,,TSM,,Lebanon,N
5066,Trans-Air Services,\N,,TSN,AIR TRANS,Nigeria,N
5067,Transaero Airlines,\N,UN,TSO,TRANSOVIET,Russia,Y
5068,Transportes Aereos Inter,\N,,TSP,TRANSPO-INTER,Guatemala,N
5069,Trans Service Airlift,\N,,TSR,TRANS SERVICE,Democratic Republic of Congo,N
5070,Tri-State Aero,\N,,TSS,TRI-STATE,United States,N
5071,TRAST,\N,,TST,TRAST,Kazakhstan,N
5072,Tropair Airservices,\N,,TSV,TROPIC,United Kingdom,N
5073,Transwings,\N,,TSW,SWISSTRANS,Switzerland,N
5074,Thai Star Airlines,\N,T9,TSX,THAI STAR,Thailand,N
5075,Tristar Air,\N,,TSY,TRIPLE STAR,Egypt,N
5076,TTA - Sociedade de Transporte e Trabalho Aereo,\N,,TTA,KANIMANBO,Mozambique,N
5077,Transteco,\N,,TTC,TRANSTECO,Angola,N
5078,Tarhan Tower Airlines,\N,,TTH,,Turkey,N
5079,Total Linhas Aereas,\N,,TTL,TOTAL,Brazil,N
5080,Triple O Aviation,\N,,TTP,MIGHTY WING,Nigeria,N
5081,Transportaciones Y Servicios Aereos,\N,,TTR,TRANSPORTACIONES,Mexico,N
5082,Transporte Aereo Tecnico Ejecutivo,\N,,TTS,TECNICO,Mexico,N
5083,Turkmenistan Airlines,Turkmenhovayollary,T5,TUA,TURKMENISTAN,Turkmenistan,Y
5084,Turismo Aereo de Chile,\N,,TUC,TURICHILE,Chile,N
5085,Tuninter,\N,UG,TUI,,Tunisia,Y
5086,Tulpar Air,\N,,TUL,URSAL,Russia,N
5087,Tyumenspecavia,\N,,TUM,TUMTEL,Russia,N
5088,Taxi Aereo Turistico,\N,,TUO,TURISTICO,Mexico,N
5089,Tulpar Air Service,\N,,TUX,TULPA,Kazakhstan,N
5090,Tuna Aero,\N,,TUZ,TUNA,Sweden,N
5091,Trans America Airlines,\N,,TVA,TRANS-AMERICA,United States,N
5092,Trabajos Aereos Vascongados,\N,,TVH,TRAVASA,Spain,N
5093,Tiramavia,\N,,TVI,TIRAMAVIA,Moldova,N
5094,Travel Service,\N,,TVL,TRAVEL SERVICE,Hungary,N
5095,Transavio,\N,,TVO,TRANS-BALLERIO,Italy,N
5096,Tavrey Airlines,\N,T6,TVR,TAVREY,Ukraine,N
5097,Travel Service,\N,QS,TVS,SKYTRAVEL,Czech Republic,Y
5098,Trans World Airlines,TWA,TW,TWA,TWA,United States,N
5099,Transwede Airways,\N,,TWE,TRANSWEDE,Sweden,N
5100,Twinjet Aircraft Sales,\N,,TWJ,,United Kingdom,N
5101,Tradewinds Aviation,\N,,TWL,TRADEWINDS CANADA,Canada,N
5102,Transairways,\N,,TWM,,Mozambique,N
5103,Twente Airlines,\N,,TWO,COLIBRI,Netherlands,N
5104,Trans Air Welwitchia,\N,,TWW,WELWITCHIA,Angola,N
5105,Texair Charter,\N,,TXA,OKAY AIR,United States,N
5106,Transaviaexport,\N,AL,TXC,TRANSEXPORT,Belarus,N
5107,Transilvania Express,\N,,TXE,TRANSAIR EXPRESS,Romania,N
5108,Taxi Aereo Cozatl,\N,,TXL,TAXI COZATL,Mexico,N
5109,Taxi Aereo de Mexico,\N,,TXM,TAXIMEX,Mexico,N
5110,Texas National Airlines,\N,,TXN,TEXAS NATIONAL,United States,N
5111,Taxis Aereos de Sinaloa,\N,,TXO,TAXIS SINALOA,Mexico,N
5112,Taxirey,\N,,TXR,TAXIREY,Mexico,N
5113,Texas Airlines,\N,,TXS,TEXAIR,United States,N
5114,Texas Air Charters,\N,,TXT,TEXAS CHARTER,United States,N
5115,Tex Star Air Freight,\N,,TXZ,TEX STAR,United States,N
5116,Tayflite,\N,,TYF,TAYFLITE,United Kingdom,N
5117,Trygg-Flyg,\N,,TYG,TRYGG,Sweden,N
5118,Tyrol Air Ambulance,\N,,TYW,TYROL AMBULANCE,Austria,N
5119,Transporte Aereo Ernesto Saenz,\N,,TZE,TRANSPORTE SAENZ,Mexico,N
5120,Tajikstan,\N,,TZK,TAJIKSTAN,Tajikistan,N
5121,Thyssen Krupp AG,\N,,BLI,BLUELINE,Germany,N
5122,TUIfly Nordic,\N,6B,BLX,BLUESCAN,Sweden,Y
5123,Transportes Aereos Bolivianos,\N,,BOL,BOL,Bolivia,N
5124,Tiphook PLC,\N,,BOX,BOX,United Kingdom,N
5125,Top Flight Air Service,\N,,CHE,CHECK AIR,United States,N
5126,Trans America,\N,,CLR,CLINTON AIRWAYS,United States,N
5127,Triple Alpha,\N,,CLU,CAROLUS,Germany,N
5128,Tashkent Aircraft Production Corporation,\N,,CTP,CORTAS,Uzbekistan,N
5129,Texas Airways,\N,,CWT,TEXAS AIRWAYS,United States,N
5130,Transportes Aereos Don Carlos,\N,,DCL,DON CARLOS,Chile,N
5131,Telnic Limited,\N,,DOT,DOT TEL,United Kingdom,N
5132,Triton Airlines,\N,,DRC,TRITON AIR,Canada,N
5133,TAAG Angola Airlines,\N,DT,DTA,DTA,Angola,Y
5134,Tassili Airlines,\N,SF,DTH,TASSILI AIR,Algeria,N
5135,Transporte Ejecutivo Aereo,\N,,EAR,EJECUTIVO-AEREO,Mexico,N
5136,Transportes Aereos Nacionales De Selva Tans,\N,,ELV,AEREOS SELVA,Peru,N
5137,TAG Farnborough Airport,\N,,FBO,,United Kingdom,N
5138,Transaviaservice,\N,,FNV,TRANSAVIASERVICE,Georgia,N
5139,TAG Aviation,\N,,FPG,TAG AVIATION,Switzerland,N
5140,The 955 Preservation Group,\N,,GFN,GRIFFON,United Kingdom,N
5141,Trans-Air-Link,\N,,GJB,SKY TRUCK,United States,N
5142,Tradewind Aviation,\N,,GPD,GOODSPEED,United States,N
5143,Trail Lake Flying Service,\N,,HBA,HARBOR AIR,United States,N
5144,TAF Helicopters,\N,,HET,HELITAF,Spain,N
5145,Tango Bravo,\N,,HTO,HELI TANGO,France,N
5146,Turkish Air Force,\N,,HVK,TURKISH AIRFORCE,Turkey,Y
5147,Thryluthjonustan,\N,,IHS,,Iceland,N
5148,TA-Air Airline,\N,,IRF,TA-AIR,Iran,N
5149,Tara Air Line,\N,,IRR,TARAIR,Iran,N
5150,Trading Air Cargo,\N,,JCH,TRADING CARGO,Mauritania,N
5151,Trans-Kiev,\N,,KCA,TRANS-KIEV,Ukraine,N
5152,Tal Air Charters,\N,,JEL,JETEL,Canada,N
5153,Transcontinental Airlines,\N,,KRA,REGATA,Kazakhstan,N
5154,Transaviabaltika,\N,,KTB,TRANSBALTIKA,Lithuania,N
5155,Transair-Gyraintiee,\N,,KTS,KOTAIR,Russia,N
5156,TAM Mercosur,\N,PZ,LAP,PARAGUAYA,Paraguay,Y
5157,The Lancair Company,\N,,LCC,LANCAIR,United States,N
5158,The Army Aviation Heritage Foundation,\N,,LEG,LEGACY,United States,N
5159,Top Sky International,\N,,LKW,TOPINTER,Indonesia,N
5160,Trans States Airlines,\N,AX,LOF,WATERSKI,United States,Y
5161,Trans Atlantis,\N,,LTA,LANTRA,Canada,N
5162,Transportacion Aerea Del Mar De Cortes,\N,,MCT,TRANS CORTES,Mexico,N
5163,Transporte Aero MGM,\N,,MGM,AERO EMM-GEE-EMM,Mexico,N
5164,Tigerfly,\N,,MOH,MOTH,United Kingdom,N
5165,Transportes Aereos Amparo,\N,,MPO,AMPARO,Mexico,N
5166,Trans Air,\N,,MUI,MAUI,United States,N
5167,Transportes Aereos Mexiquenses,\N,,MXQ,MEXIQUENSES,Mexico,N
5168,Travelsky Technology,\N,1E,,,China,N
5169,Thalys,\N,2H,,,Belgium,N
5171,TNT International Aviation,\N,,NTR,NITRO,United Kingdom,N
5172,Open Skies Consultative Commission,\N,1L,OSY,OPEN SKIES,United States,N
5173,Trans-Pacific Orient Airways,\N,,PCW,PACIFIC ORIENT,Philippines,N
5174,Tauranga Aer Club,\N,,PGS,,New Zealand,N
5175,TSSKB-Progress,\N,,PSS,PROGRESS,Russia,N
5176,Trans World Express,\N,,RBD,RED BIRD,United States,N
5177,Trans Reco,\N,,REC,TRANS-RECO,Mauritania,N
5178,Tas Aviation,\N,,RMS,TASS AIR,United States,N
5179,Tarom,\N,RO,ROT,TAROM,Romania,Y
5180,Transportes Aereos I.R. Crusoe,\N,,ROU,ROBINSON CRUSOE,Chile,N
5181,Transportes Aereos Sierra,\N,,RRT,SIERRA ALTA,Mexico,N
5182,Tbilisi Aviation University,\N,,RRY,AIRFERRY,Georgia,N
5183,Trans Am Compania,\N,,RTM,AERO TRANSAM,Ecuador,N
5184,Trans Sahara Air,\N,,SBJ,TRANS SAHARA,Nigeria,N
5185,Transportes Aereos Sierra Madre,\N,,SEI,TRANSPORTE SIERRA,Mexico,N
5186,Trans Asian Airlines,\N,,SRT,TRASER,Kazakhstan,N
5187,Turan Air,\N,3T,URN,TURAN,Azerbaijan,Y
5188,TRIP Linhas A,\N,8R,TIB,TRIP,Brazil,Y
5189,Tusheti,\N,,USB,TUSHETI,Georgia,N
5190,TAPC Aviatrans Aircompany,\N,,UTM,AVIATAPS,Uzbekistan,N
5191,Trans-Ulgii,\N,,UTN,TRANS-ULGII,Mongolia,N
5192,Transarabian Transportation Services,\N,,UTT,ARABIAN TRANSPORT,Uganda,N
5193,Transaven,\N,,VEN,TRANSAVEN AIRLINE,Venezuela,N
5194,Tag Aviation UK,\N,,VIP,SOVEREIGN,United Kingdom,N
5195,Tbilaviamsheni,\N,,VNZ,TBILAVIA,Georgia,N
5196,Taxi de Veracruz,\N,,VRC,VERACRUZ,Mexico,N
5197,Travel Express Aviation Services,\N,,XAR,TRAVEL EXPRESS,Indonesia,N
5198,Taxi Aero Del Norte,\N,,XNR,TAXI NORTE,Mexico,N
5199,Transports et Travaux A,\N,OF,,,,N
5202,USAF Air Mobility Operations Control Center,\N,,DOD,,United States,N
5203,U.S. Department of the Interior,\N,,DOI,INTERIOR,United States,N
5204,United Kingdom Civil Aviation Authority,\N,,EXM,EXAM,United Kingdom,N
5205,Union des Transports Africains de Guinee,\N,,GIH,TRANSPORT AFRICAIN,Guinea,N
5206,US Army Parachute Team,\N,,GKA,GOLDEN KNIGHTS,United States,N
5207,USA3000 Airlines,\N,U5,GWY,GETAWAY,United States,Y
5208,United Arabian Airlines,\N,,UAB,UNITED ARABIAN,Sudan,N
5209,United Airlines,\N,UA,UAL,UNITED,United States,Y
5210,United Air Charters,\N,,UAC,UNITAIR,Zimbabwe,Y
5211,United Carriers Systems,\N,,UCS,UNITED CARRIERS,United States,N
5212,United Eagle Airlines,\N,,UEA,UNITED EAGLE,China,N
5213,United Feeder Service,\N,U2,,,United States,N
5214,United Kingdom Civil Aviation Authority,\N,,CFU,MINAIR,United Kingdom,N
5215,United Kingdom Royal VIP Flights,\N,,KRF,KITTYHAWK,United Kingdom,N
5216,United Kingdom Royal VIP Flight,\N,,KRH,SPARROWHAWK,United Kingdom,N
5217,United Kingdom Civil Aviation Authority,\N,,SDS,STANDARDS,United Kingdom,N
5218,United Kingdom Royal VIP Flights,\N,,TQF,RAINBOW,United Kingdom,N
5219,United States Coast Guard Auxiliary,\N,,CGX,COASTGUARD AUXAIR,United States,N
5220,United States Department Of Agriculture,\N,,AGR,AGRICULTURE,United States,N
5221,University Air Squadron,\N,,UAD,,United Kingdom,N
5222,University Air Squadron,\N,,UAJ,,United Kingdom,N
5223,University Air Squadron,\N,,UAA,,United Kingdom,N
5224,University Air Squadron,\N,,UAH,,United Kingdom,N
5225,Universal Avia,\N,,HBU,KHARKIV UNIVERSAL,Ukraine,N
5226,UK HEMS,\N,,HLE,HELIMED,United Kingdom,N
5227,USA Jet Airlines,\N,U7,JUS,JET USA,United States,N
5228,Unijet,\N,,LEA,LEADAIR,France,N
5229,US Marshals Service,\N,,MSH,MARSHALAIR,United States,N
5230,University of North Dakota,\N,,NDU,SIOUX,United States,N
5231,Universal Airlines,\N,,PNA,PACIFIC NORTHERN,United States,N
5232,Uganda Royal Airways,\N,,RAU,UGANDA ROYAL,Uganda,N
5233,United Aviation Services,\N,,SAU,UNISERVE,Spain,N
5234,Ural Airlines,\N,U6,SVR,SVERDLOVSK AIR,Russia,Y
5235,Ukraine Transavia,\N,,TRB,KIROVTRANS,Ukraine,N
5236,United Arab Emirates Air Force,\N,,UAF,UNIFORCE,United Arab Emirates,N
5237,Union Africaine des Transports,\N,,UAI,UNAIR,Ivory Coast,N
5238,Uganda Air Cargo,\N,,UCC,UGANDA CARGO,Uganda,N
5239,US Airports Air Charter,\N,,UCH,US CHARTER,United States,N
5240,Ucoaviacion,\N,,UCO,UCOAVIACION,Spain,N
5241,Ues-Avia Aircompany,\N,,UES,AVIASYSTEM,Ukraine,N
5242,UFS,\N,,UFS,FEEDER EXPRESS,United States,N
5243,Uganda Airlines,\N,,UGA,UGANDA,Uganda,N
5244,Urgemer Canarias,\N,,UGC,URGEMER,Spain,N
5245,Ukrainian Helicopters,\N,,UHL,UKRAINE COPTERS,Ukraine,N
5246,Ulyanovsk Higher Civil Aviation School,\N,,UHS,PILOT AIR,Russia,N
5247,Universal Jet Rental de Mexico,\N,,UJR,UNIVERSAL JET,Mexico,N
5248,Universal Jet Aviation,\N,,UJT,UNI-JET,United States,N
5249,UK International Airlines,\N,,UKI,KHALIQ,United Kingdom,N
5250,Ukraine Air Alliance,\N,,UKL,UKRAINE ALLIANCE,Ukraine,N
5251,UM Airlines,\N,UF,UKM,UKRAINE MEDITERRANEE,Ukraine,Y
5252,Ukraine Air Enterprise,\N,,UKN,ENTERPRISE UKRAINE,Ukraine,N
5253,United Kingdom Home Office,\N,,UKP,POLICE,United Kingdom,N
5254,Ukrainian Cargo Airways,\N,6Z,UKS,CARGOTRANS,Ukraine,N
5255,Ultrair,\N,,ULT,ULTRAIR,United States,N
5256,Unitemp-M,\N,,UMS,TOPAZ,Russia,N
5257,Uni-Fly,\N,,UNC,UNICOPTER,Denmark,N
5258,Union Flights,\N,,UNF,UNION FLIGHTS,United States,N
5259,Universal Jet,\N,,UNJ,PROJET,Spain,N
5260,Unsped Paket Servisi,\N,,UNS,UNSPED,Turkey,N
5261,Unifly Servizi Aerei,\N,,UNU,UNIEURO,Italy,N
5262,Ukrainian Pilot School,\N,,UPL,PILOT SCHOOL,Ukraine,N
5263,United Parcel Service,\N,5X,UPS,UPS,United States,N
5264,Uraiavia,\N,,URV,URAI,Russia,N
5265,US Airways,\N,US,USA,U S AIR,United States,Y
5266,US Check Airlines,\N,,USC,STAR CHECK,United States,N
5267,USAfrica Airways,\N,,USF,AFRICA EXPRESS,United States,N
5268,US Helicopter,\N,,USH,US-HELI,United States,Y
5269,US Jet,\N,,USJ,USJET,United States,N
5270,US Express,\N,,USX,AIR EXPRESS,United States,N
5271,UTair Aviation,\N,UT,UTA,UTAIR,Russia,Y
5272,UTAGE,\N,,UTG,UTAGE,Equatorial Guinea,N
5273,Utair South Africa,\N,,UTR,AIRUT,South Africa,N
5274,Ukrainian State Air Traffic Service Enterprise,\N,,UTS,AIRRUH,Ukraine,N
5275,Urartu-Air,\N,,UTU,,Armenia,N
5276,Universal Airways,\N,,UVA,UNIVERSAL,United States,N
5277,Universal Airlines,\N,,UVG,GUYANA JET,Guyana,N
5278,Uvavemex,\N,,UVM,UVAVEMEX,Mexico,N
5279,United States Air Force,\N,,AIO,AIR CHIEF,United States,Y
5280,United Aviation,\N,,UVN,UNITED AVIATION,Kuwait,N
5281,Uzbekistan Airways,\N,HY,UZB,UZBEK,Uzbekistan,Y
5282,Ukraine International Airlines,\N,PS,AUI,UKRAINE INTERNATIONAL,Ukraine,Y
5283,Universal Airlines,\N,,WEC,AIRGO,United States,N
5284,US Helicopter Corporation,\N,UH,,,United States,Y
5288,V Australia Airlines,\N,VA,VAU,KANGA,Australia,N
5289,V Bird Airlines Netherlands,\N,,VBA,VEEBEE,Netherlands,N
5290,V-avia Airline,\N,,WIW,VEE-AVIA,Ukraine,N
5291,V-Berd-Avia,\N,,VBD,VEEBIRD-AVIA,Armenia,N
5292,Vacationair,\N,,VAC,VACATIONAIR,Canada,N
5293,Valair AG (Helicoptere),\N,,RDW,ROADWATCH,Switzerland,N
5294,Valan International Cargo Charter,\N,,VLA,NALAU,South Africa,N
5295,Valan Limited,\N,,VLN,VALAN,Moldova,N
5296,Valfell-Verkflug,\N,,EHR,ROTOR,Iceland,N
5297,Valuair,\N,VF,VLU,VALUAIR,Singapore,Y
5298,Van Air Europe,\N,,VAA,EUROVAN,Czech Republic,N
5299,Vanguardia en Aviacion en Colima,\N,,VGC,VANGUARDIA COLIMA,Mexico,N
5300,Vanguard Airlines,\N,,VGD,VANGUARD AIR,United States,N
5301,Vasco Air,\N,,VFC,VASCO AIR,Vietnam,Y
5302,Vega,\N,,VAG,SEGA,Kazakhstan,N
5303,Vega Air Company,\N,,WGA,WEGA FRANKO,Ukraine,N
5304,Veles,\N,,WEL, Ukrainian Aviation Company,VELES,N
5305,Verataxis,\N,,VTX,VERATAXIS,Mexico,N
5306,Veritair Ltd,\N,,BTP,NET RAIL,United Kingdom,N
5307,Voyageur Airways,\N,VC,VAL,VOYAGEUR,Canada,N
5308,Vernicos Aviation,\N,,GRV,NIGHT RIDER,Greece,N
5309,Vietnam Airlines,\N,VN,HVN,VIET NAM AIRLINES,Vietnam,Y
5310,Vozdushnaya Academy,\N,,KWA,VOZAIR,Kazakhstan,N
5311,VIM Airlines,\N,NN,MOV,MOV AIR,Russia,Y
5312,VIM-Aviaservice,\N,,MVY,,Russia,N
5313,VIA Rail Canada,\N,2R,,,Canada,N
5314,Victoria Aviation,\N,,ENV,ENDEAVOUR,United Kingdom,N
5315,Viscount Air Service,\N,,VCT,VISCOUNT AIR,United States,N
5317,Vision Airlines,\N,,SSI,SUPER JET,Nigeria,N
5318,VIP Air Charter,\N,,FXF,FOX FLIGHT,United States,N
5319,VIP Avia,\N,,PAV,NICOL,Kazakhstan,N
5320,VIP Avia,\N,,PRX,PAREX,Latvia,N
5321,Visionair,\N,,VAT,VISIONAIR,Ireland,N
5322,Viasa,\N,VA,,,Venezuela,N
5323,VICA - Viacao Charter Aereos,\N,,VCA,VICA,Brazil,N
5324,Volare Air Charter Company,\N,,VCM,CARMEN,United States,N
5325,Volaris,\N,Y4,VOI,VOLARIS,Mexico,Y
5326,Volga-Dnepr Airlines,\N,VI,VDA,VOLGA-DNEPR,Russia,Y
5327,Vega Airlines,\N,,VEA,VEGA AIRLINES,Bulgaria,N
5328,Venescar Internacional,\N,,VEC,VECAR,Venezuela,N
5329,Victor Echo,\N,,VEE,VICTOR ECHO,Spain,N
5330,Virgin Express Ireland,\N,,VEI,GREEN ISLE,Ireland,N
5331,Virgin America,\N,VX,VRD,REDWOOD,United States,Y
5332,Vieques Air Link,\N,,VES,VIEQUES,United States,N
5333,Virgin Express,\N,TV,VEX,VIRGIN EXPRESS,Belgium,Y
5334,VZ Flights,\N,,VFT,ZETA FLIGHTS,Mexico,N
5335,Virgin Nigeria Airways,\N,VK,VGN,VIRGIN NIGERIA,Nigeria,Y
5336,Vologda State Air Enterprise,\N,,VGV,VOLOGDA AIR,Russia,N
5337,VH-Air Industrie,\N,,VHA,AIR V-H,Angola,N
5338,VHM Schul-und-Charterflug,\N,,VHM,EARLY BIRD,Germany,N
5339,Vibroair Flugservice,\N,,VIB,VITUS,Germany,N
5340,VIP Servicios Aereos Ejecutivos,\N,,VIC,VIP-EJECUTIVO,Mexico,N
5341,VIP Empresarial,\N,,VIE,VIP EMPRESARIAL,Mexico,N
5342,VIF Luftahrt,\N,,VIF,VIENNA FLIGHT,Austria,N
5343,Vega Aviation,\N,,VIG,VEGA AVIATION,Sudan,N
5344,Vichi,\N,,VIH,VICHI,Moldova,N
5345,Viking Airlines,\N,,VIK,SWEDJET,Sweden,N
5346,Vinair - Helicoptereos,\N,,VIN,VINAIR,Portugal,N
5347,Virgin Atlantic Airways,\N,VS,VIR,VIRGIN,United Kingdom,Y
5348,Viajes Ejecutivos Mexicanos,\N,,VJM,VIAJES MEXICANOS,Mexico,N
5349,Vistajet,\N,,VJT,VISTA,Canada,N
5350,Viva Macau,\N,ZG,VVM,JACKPOT,Macao,Y
5351,Volare Airlines,\N,VE,VLE,VOLA,Italy,Y
5352,Vueling Airlines,\N,VY,VLG,VUELING,Spain,Y
5353,Vladivostok Air,\N,XF,VLK,VLADAIR,Russia,Y
5354,Varig Log,\N,LC,VLO,VELOG,Brazil,Y
5355,Vertical-T Air Company,\N,,VLT,VERTICAL,Russia,N
5356,Vero Monmouth Airlines,\N,,VMA,VERO MONMOUTH,United States,N
5357,Vipport Joint Stock Company,\N,,VNK,,Russia,N
5358,Viaggio Air,\N,VM,VOA,VIAGGIO,Bulgaria,N
5359,Voyager Airlines,\N,,VOG,VOYAGER AIR,Bangladesh,N
5360,Virgin Australia,\N,VA,VOZ,VIRGIN,Australia,Y
5361,Virgin Blue,\N,DJ,VBI,BLUEY,Australia,N
5362,VIP Air,\N,,VPA,VIP TAXI,Slovakia,N
5363,Veteran Air,\N,,VPB,VETERAN,Ukraine,N
5364,VIP-Avia,\N,,VPV,VIP AVIA,Georgia,N
5365,Vertair,\N,,VRA,VERITAIR,United Kingdom,N
5366,Volare Airlines,\N,,VRE,UKRAINE VOLARE,Ukraine,N
5367,Voar Lda,\N,,VRL,VOAR LINHAS,Angola,N
5368,VRG Linhas Aereas,Varig,RG,VRN,VARIG,Brazil,Y
5369,Vickers Limited,\N,,VSB,VICKERS,United Kingdom,N
5370,Visig Operaciones Aereas,\N,,VSG,VISIG,Spain,N
5371,Vision Airways Corporation,\N,,VSN,VISION,Canada,N
5372,Voronezh Aircraft Manufacturing Society,\N,,VSO,VASO,Russia,N
5373,VASP,\N,VP,VSP,VASP,Brazil,Y
5374,Virign Islands Seaplane Shuttle,\N,,VSS,WATERBIRD,United States,N
5375,Vuelos Especializados Tollocan,\N,,VTC,VUELOS TOLLOCAN,Mexico,N
5376,Vuelos Corporativos de Tehuacan,\N,,VTH,VUELOS TEHUACAN,Mexico,N
5377,Vostok Airlines,\N,,VTK,VOSTOK,Russia,N
5378,Victor Tagle Larrain,\N,,VTL,VITALA,Chile,N
5379,Vointeh,\N,,VTV,VOINTEH,Bulgaria,N
5380,Vuelos Internos Privados VIP,\N,,VUR,VIPEC,Ecuador,N
5381,Vuela Bus,\N,,VUS,VUELA BUS,Mexico,N
5382,Vzlyet,\N,,VZL,VZLYET,Russia,N
5383,VLM Airlines,\N,VG,VLM,RUBENS,Belgium,Y
5384,Viking Express,\N,,WCY,TITAN AIR,United States,N
5385,Victoria International Airways,\N,,WEV,VICTORIA UGANDA,Uganda,N
5386,Volga Aviaexpress,\N,,WLG,GOUMRAK,Russia,N
5387,WDL Aviation,\N,,WDL,WDL,Germany,N
5388,WRA Inc,\N,,WRR,WRAP AIR,United States,N
5389,WSI Corporation,\N,,XWS,,United States,N
5390,Walmart Aviation,\N,,CGG,CHARGE,United States,N
5391,Walsten Air Services,\N,,WAS,WALSTEN,Canada,N
5392,Walter I Linkoping,\N,,GOT,GOTHIC,Sweden,N
5393,Wapiti Aviation,\N,,WPT,WAPITI,Canada,N
5394,Warbelow's Air Ventures,\N,,WAV,WARBELOW,United States,N
5395,Warwickshire Aerocentre Ltd.,\N,,ATX,AIRTAX,United Kingdom,N
5396,Wasaya Airways,\N,,WSG,WASAYA,Canada,N
5397,Wayraper,\N,7W,,WAYRAPER,Peru,Y
5398,Weasua Air Transport Company,\N,,WTC,WATCO,Liberia,N
5399,WebJet Linhas A,\N,WJ,WEB,WEB-BRASIL,Brazil,Y
5400,Welch Aviation,\N,,TDB,THUNDER BAY,United States,N
5401,Welcome Air,\N,2W,WLC,WELCOMEAIR,Austria,Y
5402,Wermlandsflyg AB,\N,,BLW,BLUESTAR,Sweden,N
5403,West Africa Airlines,\N,,WCB,KILO YANKEE,Ghana,N
5404,West African Air Transport,\N,,WTF,WESTAF AIRTRANS,Senegal,N
5405,West African Airlines,\N,,WSF,,Benin,N
5406,West African Cargo Airlines,\N,,WAC,WESTAF CARGO,Mauritania,N
5407,West Air Luxembourg,\N,,WLX,WEST LUX,Luxembourg,N
5408,West Air Sweden,\N,PT,SWN,AIR SWEDEN,Sweden,N
5409,West Caribbean Airways,\N,,WCW,,Colombia,N
5410,West Caribbean Costa Rica,\N,,WCR,WEST CARIBBEAN,Costa Rica,N
5411,West Coast Air,\N,8O,,,Canada,Y
5412,West Coast Airlines,\N,,WCG,WHISKY INDIA,Ghana,N
5413,West Coast Airways,\N,,WCA,WEST-LEONE,Sierra Leone,N
5414,West Freugh DTEO,\N,,TEE,TEEBIRD,United Kingdom,N
5415,West Wind Aviation,\N,,WEW,WESTWIND,Canada,N
5416,WestJet,\N,WS,WJA,WESTJET,Canada,Y
5417,Westair Aviation,\N,,WAA,WESTAIR WINGS,Namibia,N
5418,Westair Cargo Airlines,\N,,WSC,WESTCAR,Cote d'Ivoire,N
5419,Westair Industries,\N,,PCM,PAC VALLEY,United States,N
5420,Westcoast Energy,\N,,BLK,BLUE FLAME,Canada,N
5421,Western Air,\N,,WST,WESTERN BAHAMAS,Bahamas,N
5422,Western Air Couriers,\N,,NPC,NORPAC,United States,N
5423,Western Air Express,\N,,WAE,WESTERN EXPRESS,United States,N
5424,Western Airlines,\N,WA,WAL,WESTERN,United States,Y
5425,Western Arctic Air,\N,,WAL,WESTERN ARCTIC,Canada,N
5426,Western Aviators,\N,,WTV,WESTAVIA,United States,N
5427,Western Express Air Lines,\N,,WES,WEST EX,Canada,N
5428,Western Pacific Airlines,\N,,KMR,KOMSTAR,United States,N
5429,Western Pacific Airservice,\N,,WPA,WESTPAC,Solomon Islands,N
5430,Westflight Aviation,\N,,WSL,WEST LINE,United Kingdom,N
5431,Westgates Airlines,\N,,WSA,WESTATES,United States,N
5432,Westland Helicopters,\N,,WHE,WESTLAND,United Kingdom,N
5433,Westpoint Air,\N,,WTP,WESTPOINT,Canada,N
5434,Westward Airways,\N,CN,WWD,WESTWARD,United States,N
5435,White,\N,,WHT,YOUNG SKY,Portugal,N
5436,White Eagle Aviation,\N,,WEA,WHITE EAGLE,Poland,N
5437,White River Air Services,\N,,WRA,,Canada,N
5438,Whyalla Airlines,\N,,WWL,,Australia,N
5439,Widerøe,\N,WF,WIF,WIDEROE,Norway,Y
5440,Wiggins Airways,\N,,WIG,WIGGINS AIRWAYS,United States,N
5441,Wiking Helikopter Service,\N,,WHS,WEEKING,Germany,N
5442,Wilbur's Flight Operations,\N,,WFO,WILBURS,United States,N
5443,Williams Air,\N,,WLS,WILLIAMS AIR,United States,N
5444,Williams Grand Prix Engineering,\N,,WGP,GRAND PRIX,United Kingdom,N
5445,Wimbi Dira Airways,\N,,WDA,WIMBI DIRA,Democratic Republic of Congo,N
5446,Winair,\N,,WNA,WINAIR,United States,N
5447,Wind Jet,\N,IV,JET,GHIBLI,Italy,Y
5448,Wind Spirit Air,\N,,WSI,WIND SPIRIT,United States,N
5449,Windrose Air,\N,,QGA,QUADRIGA,Germany,N
5450,Windward Islands Airways International,\N,,WIA,WINDWARD,Netherland,N
5451,Wings Air,\N,IW,WON,WINGS ABADI,Indonesia,Y
5452,Wings Air Transport,\N,,WAT,,Sudan,N
5453,Wings Airways,\N,,WAW,WING SHUTTLE,United States,N
5454,Wings Aviation,\N,,WOL,WINGJET,Guyana,N
5455,Wings Express,\N,,WEX,WINGS EXPRESS,United States,N
5456,Wings of Alaska,\N,K5,WAK,WINGS ALASKA,United States,N
5457,Wings of Lebanon Aviation,\N,,WLB,WING LEBANON,Lebanon,N
5458,Winlink,\N,,WIN,WINLINK,Saint Lucia,N
5459,Wisconsin Air National Guard,\N,,WAG,,United States,N
5460,Wisman Aviation,\N,,WSM,WISMAN,United States,N
5461,Wizz Air,\N,W6,WZZ,WIZZ AIR,Hungary,Y
5462,Wizz Air Hungary,\N,8Z,WVL,WIZZBUL,Bulgaria,Y
5463,Wondair on Demand Aviation,\N,,WNR,WONDAIR,Spain,N
5464,Woodgate Executive Air Charter,\N,,CWY,CAUSEWAY,United Kingdom,N
5465,World Airways,\N,WO,WOA,WORLD,United States,Y
5466,World Weatherwatch,\N,,XWW,,Canada,N
5467,World Wing Aviation,\N,,WWM,MANAS WING,Kyrgyzstan,N
5468,Worldspan,\N,1P,,,United States,N
5469,Worldwide Air Charter Systems,\N,,CSW,CHARTER SYSTEMS,Canada,N
5470,Worldwide Aviation Services,\N,,WWS,,Pakistan,N
5471,Worldwide Jet Charter,\N,,WWI,WORLDWIDE,United States,N
5472,Wright Air Lines,\N,,WRT,WRIGHT-AIR,United States,N
5473,Wright Air Service,\N,8V,WRF,WRIGHT FLYER,United States,N
5474,Wuhan Airlines,\N,,CWU,WUHAN AIR,China,N
5475,Wycombe Air Centre,\N,,WYC,WYCOMBE,United Kingdom,N
5476,Wyoming Airlines,\N,,WYG,WYOMING,United States,N
5479,XL Airways France,\N,SE,SEU,STARWAY,France,Y
5480,XP Internation,\N,,XPS,XP PARCEL,Netherlands,N
5481,Xabre Aerolineas,\N,,XAB,AERO XABRE,Mexico,N
5482,Xclusive Jet Charter Limited,\N,,XJC,EXCLUSIVE JET,United Kingdom,N
5483,Xerox Corporation,\N,,XER,XEROX,United States,N
5484,Xiamen Airlines,\N,MF,CXA,XIAMEN AIR,China,Y
5485,Xinjiang Airlines,\N,,CXJ,XINJIANG,China,N
5486,Xjet Limited,\N,,XJT,XRAY,United Kingdom,N
5487,Xtra Airways,\N,XP,CXP,RUBY MOUNTAIN,United States,N
5490,Yak-Service,\N,,AKY,YAK-SERVICE,Russia,N
5491,Yakolev,\N,,YAK,YAK AVIA,Russia,N
5492,Yamal Airlines,\N,YL,LLM,YAMAL,Russia,Y
5493,Yana Airlines,\N,,CYG,VICAIR,Cambodia,N
5494,Yangtze River Express,\N,Y8,YZR,YANGTZE RIVER,China,N
5495,Yellow Wings Air Services,\N,,ELW,YELLOW WINGS,Kenya,N
5496,Yemenia,\N,IY,IYE,YEMENI,Yemen,Y
5497,Yerevan-Avia,\N,,ERV,YEREVAN-AVIA,Armenia,N
5498,Young Flying Service,\N,,YFS,YOUNG AIR,United States,N
5499,Yuhi Air Lines,\N,,AYU,,Japan,N
5500,Yuzhnaya Aircompany,\N,,UGN,PLUTON,Kazakhstan,N
5501,Yuzhmashavia,\N,,UMK,YUZMASH,Ukraine,Y
5504,Z-Avia,\N,,RZV,ZEDAVIA,Armenia,N
5505,Zagros Airlines,\N,,IZG,ZAGROS,Iran,N
5506,Zaire Aero Service,\N,,ZAI,ZASAIR,Democratic Republic of Congo,N
5507,Zairean Airlines,\N,,ZAR,ZAIREAN,Democratic Republic of Congo,N
5508,Zambian Airways,\N,Q3,MBN,ZAMBIANA,Zambia,N
5509,Zambezi Airlines,\N,,ZMA,ZAMBEZI WINGS,Zambia,N
5510,Zanair,\N,,TAN,ZANAIR,Tanzania,Y
5511,Zantop International Airlines,\N,,ZAN,ZANTOP,United States,N
5512,ZAS Airlines of Egypt,\N,,ZAS,ZAS AIRLINES,Egypt,N
5513,Zenith Air,\N,,AZR,ZENAIR,South Africa,N
5514,Zenmour Airlines,\N,,EMR,ZENMOUR,Mauritania,N
5515,Zephyr Express,\N,,RZR,RECOVERY,United States,N
5516,Zhejiang Airlines,\N,,CJG,ZHEJIANG,China,N
5517,Zhetysu,\N,,JTU,ZHETYSU,Kazakhstan,N
5518,Zhez Air,\N,,KZH,,Kazakhstan,N
5519,Zhongfei General Aviation,\N,,CFZ,ZHONGFEI,China,N
5520,Zhongyuan Aviation,\N,,CYN,ZHONGYUAN,China,N
5521,Zip,\N,3J,WZP,ZIPPER,Canada,N
5522,Zimex Aviation,\N,C4,IMX,ZIMEX,Switzerland,N
5523,Zoom Airlines,\N,Z4,OOM,ZOOM,Canada,Y
5524,Zoom Airways,\N,,ZAW,ZED AIR,Bangladesh,N
5525,Zorex,\N,,ORZ,ZOREX,Spain,N
5526,Zracno Pristaniste Mali Losinj,\N,,MLU,MALI LOSINJ,Croatia,N
5533,Tyrolean Airways,\N,\N,TYR,TYROLEAN,\N,Y
5556,buzz,\N,UK,BUZ,\N,\N,N
5559,Maldivian Air Taxi,\N,8Q,\N,\N,Maldives,Y
5640,Yellow Air Taxi,,Y0,\N,,United States,N
5651,Royal Air Cambodge,,VJ,RAC,,Cambodia,Y
5813,Air Mandalay,,6T,\N,Six Tango,Burma,Y
5833,TEA Switzerland,,,4SW,,Switzerland,N
5905,TAN-SAHSA,,SH,\N,Sierra Hotel,Honduras,N
-1,Unknown,\N,-,N/A,\N,\N,Y
5982,Air Busan,\N,BX,ABL,Air Busan,Republic of Korea,Y
6002,TUI Airlines Belgium,now Jetairlfy,TB,TUB,BEAUTY,Belgium,N
5584,Sky Express,SkyExpress,XW,SXR,SKYSTORM,Russia,Y
6182,Arctic Air,,,AKR,Arctic Norway,Norway,N
6183,Braathens,Braathens SAFE,BU,BRA,Braathens,Norway,N
6196,Globus,,GH,GLP,,Russia,Y
6222,Air Kazakhstan,,9Y,KZK,Kazakh,Kazakhstan,Y
6557,Japan Air System,,JD,JAS,Air System,Japan,Y
6578,Annsett New Zealand (NZA),,ZQ,\N,,New Zealand,N
6855,EasyJet (DS),,DS,\N,,Switzerland,Y
6856,Rheintalflug,,,RTL,Rheintal,Austria,N
6860,Dan-Air London,,,DAN,,United Kingdom,N
6862,Fred. Olsen,,,FOF,,Norway,N
8359,Star Peru (2I),,2I,\N,,Peru,Y
8434,Robin Hood Aviation,,,RHA,Sherwood,Austria,N
8461,Carnival Air Lines,,KW,\N,Carnival Air,United States,Y
8463,United Airways,,4H,UBD,UNITED BANGLADESH,Bangladesh,Y
8523,Inter European Airways,,,IEA,,United Kingdom,N
8568,Trans Maldivian Airways,,M8,TMW,,Maldives,N
8576,Fly540,,5H,FFV,SWIFT TANGO,Kenya,Y
8745,Transavia France,,TO,TVF,FRENCH SUN,France,Y
8809,Island Air (WP),,WP,MKU,,United States,Y
9018,1-2-go,fly 1-2-go,OG,\N,,Thailand,N
9082,Uni Air,,B7,UIA,Glory,Taiwan,Y
9135,Gomelavia,,YD,\N,,Belarus,Y
9225,NordWind Airlines,NordWind Airlines,,NWS,NORDLAND,Russia,N
9239,Red Wings,Avialinii 400,WZ,RWZ,AIR RED,Russia,Y
9286,Hellenic Star Airways,,,HST,Hellenic Star,Greece,N
9287,BAE Systems Flight Training,,,BAZ,,Australia,N
9335,TUIfly (X3),,11,\N,,Germany,Y
9343,Felix Airways,,FU,FXX,,Yemen,Y
9344,Kostromskie avialinii,,K1,KOQ,,Russia,Y
9373,Greenfly,,XX,GFY,,Spain,Y
9531,Tajik Air,,7J,\N,,Tajikistan,Y
9541,Air Mozambique,,TM,\N,,Mozambique,Y
9577,ELK Airways,,--,ELK,,Estonia,Y
9620,Gabon Airlines,\N,GY,GBK,GABON AIRLINES,Gabon,Y
9626,MCA Airlines,,,MCA,CALSON,Sweden,Y
9656,Maldivo Airlines,,ML,MAV,Maldivo,Maldives,Y
9666,Virgin Pacific,,VH,VNP,,Fiji,Y
9764,Zest Air,,Z2,\N,,Philippines,Y
9784,Yangon Airways,,HK,\N,Hotel Kilo,Burma,Y
9793,Transport Aérien Transrégional,TAT,IJ,\N,,France,N
9798,LÍO Flugmennt,,,LIO,,Iceland,N
9808,Minerva Airlines,,N4,\N,,Italy,N
9809,Eastar Jet,,ZE,ESR,Eastar,South Korea,Y
9810,Jin Air,,LJ,JNA,Jin Air,South Korea,Y
9811,Wataniya Airways,,,KW1,,Kuwait,Y
9814,Aéris (Priv),,SH,\N,,France,N
9818,Air Arabia Maroc,,3O,\N,Air Arabia,Morocco,Y
9825,Baltic Air lines,,B1,BA1,Baltic,Latvia,Y
9828,Ciel Canadien,,YC,YCC,Ciel,Canada,Y
9829,Canadian National Airways,,CN,YCP,CaNational,Canada,Y
9833,Epic Holiday,Epic Holidays,FA,4AA,Epic,United States,Y
9838,Indochina Airlines,,,AXC,Airspup,Vietnam,Y
9840,JetWind,,,JWW,Jetwind,United States,N
9851,Air Comet Chile,,3I,\N,,Chile,Y
9853,Nazca,,-.,\N,,Peru,N
9859,German Air Force - FLB,,,FLB,FLB,Germany,Y
10069,City-Air Germany,,,CIP,,Germany,N
10094,Voronezhskie Airlanes,,DN,\N,,Russia,N
10114,Line Blue,,L8,LBL,Bluebird,Germany,Y
10117,FlyLAL Charters,,,LLC,,Lithuania,Y
10118,Blue Sky America,,BU,BKY,,United States,N
10119,Texas Spirit,,XS,TXP,,United States,N
10121,Illinois Airways,,IL,ILW,,United States,N
10122,Salzburg arrows,SZA,SZ,\N,SZA,Austria,Y
10123,Texas Wings,,TQ,TXW,TXW,United States,Y
10124,California Western,,KT,CWS,,United States,N
10128,Dennis Sky,Dennis Sky Holding,DH,DSY,DSY,Israel,Y
10224,Zz,,ZZ,\N,,Belgium,Y
10226,Atifly,,A1,A1F,atifly,United States,Y
10318,Air UK,,UK,\N,,United Kingdom,N
10334,Suckling Airways,,CB,\N,,United Kingdom,N
10367,Reno Sky,,RY,\N,,United States,N
10371,Aerolineas heredas santa maria,,,SZB,,Dominican Republic,Y
10372,Ciao Air,,99,\N,,Italy,Y
10642,Jc royal.britannica,,,JRB,,United Kingdom,Y
10646,Birmingham European,,VB,\N,,United Kingdom,N
10650,Pal airlines,,5P,\N,,Chile,Y
10673,CanXpress,,C1,CA1,CAX,Canada,Y
10674,Danube Wings (V5),,V5,\N,,Slovakia,Y
10675,Sharp Airlines,,SH,SHA,SHARP,Australia,Y
10683,CanXplorer,,C2,CAP,,Canada,Y
10715,Click (Mexicana),,QA,\N,,Mexico,Y
10735,World Experience Airline,WEA,W1,WE1,WEA,Canada,Y
10737,ALAK,,J4,\N,,Russia,Y
10738,AJT Air International,,E9,\N,,Russia,N
10739,Air Choice One,,3E,\N,,United States,Y
10740,Tianjin Airlines,,,GCR,,China,Y
10741,China United,,KN,\N,,China,Y
10748,Locair,,ZQ,LOC,LOCAIR,United States,Y
10758,Safi Airlines,,4Q,\N,,Afghanistan,Y
10765,SeaPort Airlines,,K5,SQH,SASQUATCH,United States,Y
10776,Salmon Air,,S6,\N,,United States,Y
10796,Fly Illi,,IL,ILY,,United States,N
10798,Bobb Air Freight,,01,\N,,Germany,Y
10800,Star1 Airlines,,V9,HCW,,Lithuania,Y
10845,Pelita,,6D,\N,,Indonesia,Y
10902,Alpi Eagles (E8),,E8,\N,,Italy,N
10912,Alaska Seaplane Service,,J5,\N,,United States,Y
10929,TAN,,T8,\N,,Argentina,N
10945,Enerjet,,,ENJ,ENERJET AIR,Canada,Y
10955,MexicanaLink,,I6,MXI,LINK,Mexico,Y
10960,Island Spirit,,IP,ISX,,Iceland,Y
10969,TACA Peru,,T0,\N,TACA PERU,Peru,Y
11700,Orbest,,,OBS,ORBEST,Portugal,Y
11719,Southern Air Charter,,,SOA,,Bahamas,Y
11724,SVG Air,,,SVG,Grenadines,Saint Vincent and the Grenadines,Y
11726,Air Century,,,CEY,,Dominican Republic,Y
11731,Pan Am World Airways Dominicana,PAWA Dominicana,7Q,\N,PAWA,Dominican Republic,Y
11732,Primera Air,,PF,\N,PRIMERA,Iceland,Y
11741,Air Antilles Express,,3S,\N,GREEN BIRD,Guadeloupe,Y
11751,Sol Lineas Aereas,,,OLS,FLIGHT SOL,Argentina,Y
11755,Regional Paraguaya,,P7,REP,REGIOPAR,Paraguay,Y
11761,VIP Ecuador,,V6,\N,,Ecuador,Y
11762,Transportes Aereos Cielos Andinos,,,NDN,ANDINOS,Peru,Y
11763,Peruvian Airlines,,P9,\N,,Peru,Y
11765,EasyFly,,,EFY,EASYFLY,Colombia,Y
11767,Polar Airlines,,ЯП,\N,,Russia,Y
11794,Catovair,,OC,\N,CATOVAIR,Mauritius,Y
11795,Andalus Lineas Aereas,,,ANU,Andalus,Spain,Y
11798,Air 26,,,DCD,DUCARD,Angola,Y
11800,Mauritania Airways,,,MTW,MAURITANIA AIRWAYS,Mauritania,Y
11802,CEIBA Intercontinental,,,CEL,CEIBA LINE,Equatorial Guinea,Y
11804,Halcyonair,,7Z,\N,CREOLE,Cape Verde,Y
11805,Business Aviation,,4P,\N,AFRICAN BUSINESS,Congo (Kinshasa),Y
11806,Compagnie Africaine d\\'Aviation,,E9,\N,AFRICOMPANY,Congo (Kinshasa),Y
11808,Zambia Skyways,,K8,\N,ZAMBIA SKIES,Zambia,Y
11811,AlMasria Universal Airlines,,UJ,LMU,ALMASRIA,Egypt,Y
11813,EgyptAir Express,,,MSE,EGYPTAIR EXPRESS,Egypt,Y
11814,SmartLynx Airlines,,6Y,\N,,Latvia,Y
11815,Air Italy Egypt,,,EUD,,Egypt,Y
11816,KoralBlue Airlines,,K7,KBR,KORAL BLUE,Egypt,Y
11820,Wind Rose Aviation,,,WRC,WIND ROSE,Ukraine,Y
11823,Elysian Airlines,,E4,GIE,,Cameroon,Y
11833,Sevenair,,,SEN,SEVENAIR,Tunisia,Y
11834,Hellenic Imperial Airways,,HT,IMP,IMPERIAL,Greece,Y
11836,Amsterdam Airlines,,WD,AAN,AMSTEL,Netherlands,Y
11838,Arik Niger,,Q9,NAK,,Niger,Y
11839,Dana Air,,DA,\N,DANACO,Nigeria,Y
11840,STP Airways,,8F,STP,SAOTOME AIRWAYS,Sao Tome and Principe,Y
11843,Med Airways,,7Y,\N,FLYING CARPET,Lebanon,Y
11850,Skyjet Airlines,,UQ,SJU,SKYJET,Uganda,Y
11855,Air Volga,,G6,\N,GOUMRAK,Russia,Y
11856,Transavia Denmark,,,TDK,,Denmark,Y
11857,Royal Falcon,,RL,RFJ,,Jordan,Y
11873,Euroline,,4L,MJX,GEO-LINE,Georgia,Y
11922,Worldways,,WG,WGC,,Canada,N
11943,Turkuaz Airlines,,,TRK,TURKU,Turkey,Y
11947,Athens Airways,,ZF,\N,ATHENSAIR,Greece,Y
11948,Viking Hellas,,VQ,VKH,DELPHI,Greece,Y
11949,Norlandair,,,FNA,NORLAND,Iceland,Y
11950,Flugfelag Vestmannaeyja,,,FVM,ELEGANT,Iceland,Y
11963,Starline.kz,,DZ,\N,ALUNK,Kazakhstan,Y
11995,Euro Harmony,Euro Harmony ,EH,EHM,,Portugal,N
12960,Lugansk Airlines,,L7,\N,ENTERPRISE LUHANSK,Ukraine,Y
12961,Gryphon Airlines,,6P,\N,,United States,Y
12962,Gadair European Airlines,,GP,GDR,GADAIR,Spain,Y
12965,Spirit of Manila Airlines,,SM,MNP,MANILA SKY,Philippines,Y
12975,Chongqing Airlines,,OQ,CQN,CHONG QING,China,Y
12976,Grand China Air,,,GDC,GRAND CHINA,China,Y
12978,West Air China,,PN,CHB,WEST CHINA,China,Y
12990,Falcon Air (IH),,IH,\N,,Sweden,N
12997,QatXpress,qatXpress,C3,QAX,,Qatar,Y
13076,OneChina,OneChina,1C,1CH,,China,Y
13088,NordStar Airlines,,Y7,\N,,Russia,Y
13089,Joy Air,,JR,JOY,JOY AIR,China,Y
13105,Air India Regional,,CD,\N,ALLIED,India,Y
13106,MDLR Airlines,,9H,\N,MDLR,India,Y
13107,Jagson Airlines,,,JGN,JAGSON,India,Y
13108,Maldivian,,Q2,\N,ISLAND AVIATION,Maldives,Y
13130,Xpressair,,XN,\N,,Indonesia,Y
13178,Strategic Airlines,,VC,\N,,Australia,Y
13181,Fars Air Qeshm,,,QFZ,FARS AIR,Iran,Y
13187,Eastok Avia,,,EAA,,Kyrgyzstan,Y
13188,Jupiter Airlines,,,JPU,JUPITERAIR,United Arab Emirates,Y
13189,Vision Air International,,,VIS,,Pakistan,Y
13190,Al-Naser Airlines,,NA,\N,,Iraq,Y
13200,Fuji Dream Airlines,,JH,\N,FUJI DREAM,Japan,Y
13202,Korea Express Air,,,KEA,,South Korea,Y
13209,Eznis Airways,,,EZA,EZNIS,Mongolia,Y
13211,Pacific Flier,,,PFL,KOROR,Palau,Y
13217,Syrian Pearl Airlines,,,PSB,,Syria,Y
13218,SGA Airlines,,5E,\N,SIAM,Thailand,Y
13242,Air2there,,F8,\N,,New Zealand,Y
13254,Avianova (Russia),,AO,\N,Nova,Russia,Y
13303,Parmiss Airlines (IPV),,PA,IPV,IPV,Iran,Y
13304,EuropeSky,,ES,EUV,EuropeSky,Germany,Y
13306,BRAZIL AIR,BRAZIL AIR,GB,BZE,BRAZIL AIR,Brazil,Y
13335,Homer Air,Homer Sky,MR,OME,,Germany,Y
13388,Court Line,,??,???,,United Kingdom,N
13389,South West Africa Territory Force,SWATF,??,***,,Namibia,N
13390,Lombards Air,,++,---,,South Africa,N
13391,U.S. Air,,-+,--+,,United States,Y
13392,Flitestar,,GM,===,,South Africa,N
13394,Jayrow,,\\',\\'\\,,Australia,Y
13395,Llloyd Helicopters,,::,:::,,Australia,N
13397,Wilderness Air,,;;,\N,,Australia,Y
13398,Whitaker Air,,^^,\N,,Australia,Y
13633,PanAm World Airways,,WQ,PQW,,United States,Y
13690,Virginwings,,YY,VWA,,Germany,Y
13704,KSY,Kreta Sky,KY,KSY,KSY,Greece,Y
13732,Buquebus Líneas Aéreas,,BQ,BQB,,Uruguay,Y
13734,SOCHI AIR,SOCHI,CQ,KOL,SLOW FROG,Russia,Y
13757,Wizz Air Ukraine,,WU,WAU,WIZZAIR UKRAINE,Ukraine,Y
13781,88,,47,VVN,,Cyprus,Y
13815,LCM AIRLINES,,LQ,LMM,,Russia,Y
13838,Aero Brazil,,BZ,BZL,,Brazil,N
13899,Cambodia Angkor Air (K6),,K6,\N,,Cambodia,Y
13905,Skyline nepc,,D5,\N,,India,N
13923,THREE,,H3,T33,,China,N
13936,Royal European Airlines,,69,\N,,United Kingdom,Y
13947,Tom\\'s & co airliners,Tom\\'s air,&T,T&O,T&,France,Y
13983,Azul,Azul Linhas Aéreas Brasileiras,AD,AZU,,Brazil,Y
14061,LSM Airlines,slowbird,PQ,LOO,slowbird,Russia,Y
14069,Zapolyarie Airlines,Zapolyarye Airlines,,PZY,,Russia,Y
14073,Finlandian,,,FN1,,Finland,Y
14094,LionXpress,lionXpress,C4,LIX,LIX,Cameroon,Y
14109,Nik Airways,,X1,\N,,Saudi Arabia,N
14118,Genesis,,GK,\N,,Pakistan,Y
14388,Congo Express,,XZ,\N,EXPRESSWAYS,Congo (Kinshasa),Y
14485,Fly Dubai,,FZ,FDB,,United Arab Emirates,Y
14620,Domenican Airlines,Domenican,D1,MDO,Domenican,Dominican Republic,Y
14679,ConneX European Airline,,,2CO,ConneX,Austria,Y
14725,Air Atlantic,,9A,\N,Atlantic,Canada,N
14728,Air Ops,,CR,\N,,Sweden,N
14849,Aereonautica militare,,JY,AXZ,,Italy,Y
14858,Kal Star Aviation,,,KLS,,Indonesia,Y
14881,LSM AIRLINES ,Russian. Yours Air Lines ,YZ,YZZ,Moscow frog ,Russia,Y
14921,Aero Lloyd (YP),,YP,AEF,,Germany,N
15814,UTair-Express,,UR,\N,,Russia,Y
15837,Huaxia,HUAXIA,G5,\N,,China,Y
15867,Zabaykalskii Airlines,Baikal Airlines,ZP,ZZZ,Lakeair,Russia,Y
15887,CBM America,,,XBM,AIRMAX,United States,Y
15893,Marysya Airlines,MARYSYA AIRLINES,M4,1QA,MARSHAK AIR,Russia,Y
15897,N1,,N1,\N,,Peru,Y
15930,Airlink (SAA),,4Z,\N,,South Africa,Y
15939,Westfalia Express VA,,,WFX,,Germany,Y
15953,JobAir,,3B,\N,,Czech Republic,Y
15970,Zuliana de Aviacion,Zuliana,OD,ULA,,Venezuela,N
15975,Black Stallion Airways,,BZ,BSA,Stallion,United States,Y
15984,German International Air Lines,Germanair,GM,GER,,Germany,Y
15985,TrasBrasil,,TB,TBZ,,Brazil,Y
15989,TransBrasil Airlines,,TH,THS,,Brazil,Y
15999,China SSS,Chunqiu Airlines,9C,\N,,China,Y
16025,Nihon.jet,,NJ,\N,,Japan,N
16084,AIR INDOCHINE,,,IIA,,Vietnam,Y
16093,Transportes Aéreos Nacionales de Selva,,TJ,\N,Aereos Selva,Peru,N
16100,Happy Air,,,HPY,,Thailand,Y
16101,Solar Air,,,SRB,Solar Air,Thailand,Y
16103,Air Mekong,,P8,MKG,Air Mekong,Vietnam,Y
16110,Harbour Air (Priv),,H3,\N,,Canada,Y
16116,Air Hamburg (AHO),,HH,AHO,Air Hamburg,Germany,Y
16120,ZABAIKAL AIRLINES,ZABAIKAL ,Z6,ZTT,BAIKAL ,Russia,Y
16127,TransHolding,Trans,TI,THI,,Brazil,Y
16130,SUR Lineas Aereas,,,SZZ,,Argentina,Y
16131,Aerolineas Africanas,,,AA1,,Guinea,Y
16133,Yeti Airways,,YT,\N,,Nepal,Y
16134,Georgian Airlines,,,GEG,,Georgia,N
16135,Yellowstone Club Private Shuttle,,Y1,\N,YCS,United States,Y
16136,Caucasus Airlines,,NS,\N,,Georgia,Y
16139,Serbian Airlines,,S1,SA1,,Serbia,Y
16149,Windward Islands Airways,,WM,\N,Winair,Netherlands Antilles,Y
16150,TransHolding System,,YO,TYS,,Brazil,Y
16151,CCML Airlines,,CB,CCC,,Colombia,Y
16197,Air Charter International,,SF,\N,,France,N
16198,Small Planet Airlines,,,ELC,,Lithuania,Y
16234,Fly Brasil,Fly Brasil,F1,FBL,FBL,Brazil,Y
16260,AUOS,AUOS,,AUK,AUOS,United Kingdom,N
16261,CB Airways UK ( Interliging Flights ),,1F,CIF,,United Kingdom,Y
16262,Fly Colombia ( Interliging Flights ),,3F,3FF,,Colombia,Y
16264,Trans Pas Air,,T6,TP6,,United States,Y
16273,KMV,,МИ,\N,Air Minvody,Russia,Y
16323,Himalayan Airlines,Himalaya,HC,HYM,Himalayan,Nepal,Y
16327,Indya Airline Group,Indya1,G1,IG1,Indya1,India,Y
16329,Sunwing,,WG,\N,sunwing,Canada,Y
16358,Turkish Wings Domestic,,,TWD,TWD,Turkey,Y
16359,Japan Regio,,ZX,ZXY,,Japan,Y
16362,OCEAN AIR CARGO,,,IXO,,India,Y
16363,Norte Lineas Aereas,NORTE,N0,\N,,Argentina,Y
16364,Austral Brasil,Austral Brasil lineas aereas,W7,\N,,Brazil,Y
16373,PEGASUS AIRLINES-,,H9,\N,,Turkey,Y
16399,AirLiberté,,IJ,\N,,France,N
16409,Nihon.jet connect,,,NX1,,Kyrgyzstan,Y
16415,Camair-co,,QC,\N,,Cameroon,Y
16437,Aerocontinente (Priv),,N6,\N,,Peru,N
16459,Sky Regional,Air Canada Express,RS,\N,Sky Regional,Canada,Y
16475,TUR Avrupa Hava YollarÄ±,TUR European Airways,YI,\N,TurAvrupa,Turkey,N
16487,Cruzeiro do Sul Servicos Aereos,,,CRZ,,Brazil,N
16507,LSM International ,Moskva-air,II,UWW,moose,Russia,Y
16508,Baikotovitchestrian Airlines ,,BU,BUU,,American Samoa,Y
16511,Luchsh Airlines ,Air luch,L4,LJJ,russian sky,Russia,Y
16556,ENTERair,,,QQQ,,Poland,Y
16570,Zimbabwean Airlines,,Z7,\N,,Zimbabwe,N
16585,Air Cargo Germany,,6U,\N,Loadmaster,Germany,Y
16615,Mongolian International Air Lines ,Mongol Air ,7M,ZTF,Mongol_AIr ,Mongolia,Y
16616,Alaniya Airlines,Алания,2D,\N,,Russia,N
16624,Tway Airlines,,TW,TWB,TWAY AIR,South Korea,Y
16625,Papillon Grand Canyon Helicopters,,HI,\N,,United States,Y
16628,Jusur airways,,JX,JSR,,Egypt,Y
16645,NEXT Brasil,NEXT,XB,NXB,XB,Brazil,Y
16660,AeroWorld ,Sovet Air ,W4,WER,sovet,Russia,Y
16671,Cook Island Air,,KH,\N,,Cook Islands,N
16675,US Africa Airways,,E8,\N,,United States,N
16695,GNB Linhas Aereas,,GN,\N,,Brazil,Y
16702,Usa Sky Cargo,USky,E1,ES2,USKY,United States,Y
16707,Hankook Airline,,HN,HNX,HNX,South Korea,Y
16715,Red Jet America,,RR,\N,,United States,N
16717,REDjet,,Z7,\N,,Barbados,Y
16719,Hellenic Airways,,1H,HEY,Hellenic,Greece,N
16720,Red Jet Andes,,PT,\N,,Peru,Y
16721,Red Jet Canada,,QY,\N,,Canada,Y
16723,Sprintair,,,SRN,,Poland,Y
16724,Red Jet Mexico,,4X,\N,,Mexico,Y
16725,Marusya Airways,Marusya Air,Y8,MRS,snowball,Russia,Y
16726,Era Alaska,,7H,ERR,ERAH,United States,Y
16728,AirRussia,RussianConector,R8,RRJ,russiancloud,Russia,Y
16735,Hankook Air US,,H1,HA1,,United States,Y
16738,NEPC Airlines,,D5,\N,,India,N
16753,Canadian World,,10,CNN,Canadian,Canada,N
16760,Pim Air,,,PHJ,Pim Air,Netherlands,N
16794,Carpatair Flight Training,,,SMW,Smartwings,Romania,Y
16796,I-Fly,,H5,RSY,RUSSIAN SKY,Russia,Y
16798,T.A.T,,IJ,\N,,France,N
16804,Compania de Aviacion Faucett,,,CFP,,Peru,N
16808,Kar-Air,,,KAR,,Finland,N
16811,Alinord,,DN,\N,,Italy,N
16825,Pacific Express,,VB,\N,,United States,N
16826,Whitejets,,,WTJ,WHITEJET,Brazil,Y
16837,VickJet,,KT,VKJ,Vickjet,France,Y
16844,BVI Airways,,XV,\N,,British Virgin Islands,Y
16858,Hamburg Airways,,,HAY,,Germany,Y
16860,Salsa d\\'Haiti,,SO,SLC,SALSA,Haiti,Y
16867,Zambezi Airlines (ZMA),,ZJ,\N,,Zambia,Y
16868,Kan Air,,,KND,Kan Air,Thailand,Y
16881,Air Cudlua,Air Cudlua,,CUD,Cudlua,United Kingdom,Y
16882,Polet Airlines (Priv),,YQ,\N,,Russia,Y
16895,Air Explore,,,AXE,,Slovakia,Y
16900,TROPICAL LINHAS AEREAS,TROPICAL,T1,TP3,,Brazil,N
16901,12 North,,12,N12,12N,India,Y
16919,Holidays Czech Airlines,,,HCC,,Czech Republic,Y
16921,Comtel Air,,,COE,,Austria,Y
16926,Mint Airways,,,MIC,,Spain,Y
16932,Orbit Airlines,Orbit,,OBT,Orbit,United States,Y
16939,Air Bucharest,,,BUR,,Romania,Y
16940,AlbaStar,,,LAV,,Spain,Y
16942,Mauritania Airlines International,,L6,MAI,,Mauritania,Y
16956,MAT Airways,,6F,MKD,,Macedonia,Y
16960,Asian Wings Airways,,AW,AWM,Asian Star,Burma,Y
16963,Air Arabia Egypt,,E5,RBG,,Egypt,Y
16967,Eagles Airlines,,,EGS,EAGLES,Italy,Y
16973,YES Airways,,,YEP,,Poland,Y
16975,Alitalia Cityliner,,CT,\N,,Italy,Y
16983,Direct Aero Services,,,DSV,,Romania,Y
16985,Medallion Air,,,MDP,MEDALS,Romania,Y
17022,Orchid Airlines,,OI,ORC,,Australia,Y
17023,Asia Wings,,Y5,AWA,,Kazakhstan,Y
17026,Georgian International Airlines,,,GNN,GEO-LINE,Georgia,Y
17027,Air Batumi,,,BTM,,Georgia,Y
17082,Skywest Australia,,XR,\N,,Australia,Y
17083,Nile Air,,NP,NIA,NILEBIRD,Egypt,Y
17086,Feeder Airlines,,,FDD,,Sudan,Y
17094,Senegal Airlines,,DN,SGG,,Senegal,Y
17095,Fly 6ix,,6I,\N,,Sierra Leone,Y
17099,Starbow Airlines,,S9,\N,,Ghana,Y
17115,Copenhagen Express,,0X,CX0,Copex,Denmark,Y
17408,BusinessAir,,8B,BCC,,Thailand,Y
17519,SENIC AIRLINES,,YR,\N,,United States,Y
17557,Compass Airlines (Australia),Southern Cross Airlines,,CYM,Compair,Australia,N
17563,XOJET,,,XOJ,,United States,Y
17568,Dexter (DXT),Декстер авиатакси,,DXT,Dexter,Russia,N
17571,Sky Wing Pacific,,C7,CR7,,South Korea,Y
17572,Bateleur Air,,,BEU,,South Africa,Y
17574,Air Indus,Indus Airlines Pak,PP,AI0,AIPL,Pakistan,Y
17575,Samurai Airlines,Samurai Airlines (DUMMY),07,770,Sam,Pakistan,N
17614,AirOne Continental,AirOne Continental,00,000,Eastern,Slovakia,N
17618,AirOne Polska,AirOne Polska,U1,001,AOC,Poland,N
17628,Orbit International Airlines,,,OAI,OA,United States,Y
17629,Orbit Regional Airlines,,,OAR,OA,United States,Y
17630,Orbit Atlantic Airways,,,OAN,,United States,Y
17658,Volotea,,,VOO,Volotea,Spain,Y
17666,Go Fly (United Kingdom),,,GOE,Go Flight,United Kingdom,N
17675,Peach Aviation,,MM,\N,Air Peach,Japan,Y
17694,Helitt Líneas Aéreas,,,HTH,,Spain,Y
17695,Russia State Transport,Federal State Budget Inst,,RSD,STATE AERO,Russia,Y
17726,Malaysia Wings,,,MWI,MWI,Malaysia,Y
17750,Aviabus,,U1,ABI,,Russia,Y
17780,Michael Airlines,Javi,DF,MJG,MJG,Puerto Rico,Y
17786,Korongo Airlines,,ZC,KGO,KORONGO,Congo (Kinshasa),Y
17794,Indonesia Sky,,I5,IDS,,Indonesia,Y
17822,Pelangi ,,9P,\N,,Malaysia,N
17841,Aws express,,B0,666,aws,United States,Y
17859,Southjet,,76,SJS,,United States,Y
17860,Southjet connect,,77,ZCS,,United States,Y
17861,Air Cape,,KP,\N,,South Africa,N
17862,Southjet cargo,,78,XAN,,United States,Y
17881,Iberia Express,,I2,IBS,,Spain,Y
17885,Interjet (ABC Aerolineas),,4O,\N,INTERJET,Mexico,Y
17889,AirOnix,,OG,\N,,Ukraine,Y
17890,Nordic Global Airlines,,NJ,NGB,Nordic Global,Finland,Y
17891,Scoot,,TZ,SCO,,Singapore,Y
17893,Starling Airlines Spain,,SX,STS,STARLING,Spain,N
17909,Hi Fly (5K),,5K,\N,,Portugal,Y
17911,China Northwest Airlines (WH),,WH,\N,,China,Y
17935,Zenith International Airline,Zenith,ZN,ZNA,ZENITH,Thailand,Y
17936,Orbit Airlines Azerbaijan,Orbit Azerbaijan,O1,OAB,Orbitaz,Azerbaijan,Y
17955,Air Engiadina-BRN,Air Engiadina-BRN,,RQX,Engiadina,Switzerland,N
17963,VG Airlines (IV),,,FVG,Nico,Belgium,N
17989,Air Alps Aviation (A6),,A6,\N,ALPAV,Austria,Y
18011,Austrian Airtransport,,,AAT,,Austria,N
18076,Flying kangaroo Airline,Skippy,,FKA,Skippy,Australia,Y
18083,RusJet,,,RSJ,,Russia,Y
18118,VietJet Air,VietJet,,VJC,VIETJETAIR,Vietnam,Y
18140,Spantax S.A.,,,BXS,,Spain,N
18169,Patriot Airways,,P4,\N,,United States,Y
18178,Vision Airlines (V2),,V2,RBY,RUBY,United States,Y
18228,Chicago Express (C8),ATA Connection,C8,\N,,United States,N
18232,BQB Lineas Aereas,Buquebus,5Q,\N,,Uruguay,Y
18237,AirAsia Japan,,,WAJ,WING ASIA,Japan,Y
18239,Yellowtail,,YE,YEL,,United States,Y
18241,Royal Airways,Royal Inc.,KG,RAW,RAW,United States,Y
18252,FlyHigh Airlines Ireland (FH),,FH,FHI,FLYHIRELAND,Ireland,Y
18257,Executive AirShare,,,XSR,,United States,Y
18475,Hebei Airlines,,,HBH,Hebei Air,China,Y
18476,Air KBZ,,,KBZ,Air KBZ,Burma,Y
18477,Aero VIP (2D),,2D,\N,,Portugal,Y
18497,Yangon Airways Ltd.,,YH,\N,,Burma,Y
18529,T.J. Air,,TJ,TJA,T.J. Air,United States,Y
18543,SkyWork Airlines ,,SX,\N,SKYFOX,Switzerland,Y
18545,ValueJet,,J7,VJA,CRITTER,United States,N
18553,Maastricht Airlines,,W2,\N,,Netherlands,Y
18592,CheapFlyingInternational,CheapFlying,WL,FQR,cheapflying,France,N
18616,Aviaexpresscruise,,E6,\N,,Russia,N
18617,Euro Jet,,24,\N,,Germany,Y
18621,Ukraine Atlantic,,,UAT,,Ukraine,Y
18637,AirOne Atlantic,AirOneAtlantic,00,AO1,,Slovakia,N
18643,HQ- Business Express,,HQ,\N,BizEx,United States,N
18668,Nesma Airlines,,,NMA,Nesma Airlines,Egypt,Y
18672,East Horizon,,,EHN,EAST HORIZON,Afghanistan,Y
18673,Royal Southern Airlines.,Royal Southern,R1,RS1,RSA,Colombia,N
18676,Air Majoro,,,MJP,Air Majoro,Peru,Y
18681,Fly Zoom,,,UKZ,,United Kingdom,N
18692,Rotana Jet,,,RJD,ROTANA,United Arab Emirates,Y
18700,SOCHI AIR CHATER,Sochi Air ,Q3,QER,russian doll,Russia,Y
18702,Denim Air ,FlyNonstop,J7,\N,DNM,Norway,Y
18722,WestAir,,OE,\N,,United States,N
18723,WestAir Airlines,,OE,\N,,United States,N
18724,WestAir Airlines ,,OE,\N,,United States,N
18727,North Pacific Airlines,,NO,\N,,United States,N
18732,Malindo Air,,OD,MXD,Malindo,Malaysia,Y
18749,Tramm Airlines,Tramm Airlines,9F,TLM,9F,Netherlands Antilles,N
18762,Lina Congo,,GC,\N,,Congo (Brazzaville),N
18781,Hermes Airlines,,,HRM,HERMES,Greece,Y
18825,Flightlink Tanzania,Flightlink,Z9,\N,,Tanzania,Y
18828,IzAvia,,I8,\N,,Russia,Y
18860,Катэкавиа,,,КТК,,Russia,Y
18863,Псковавиа,Псков Авиа,,PKV,,Russia,Y
18896,3 Valleys Airlines,,3V,VA3,3 Valleys,France,N
18930,Maryland Air,Maryland,M1,M1F,Maryland Flight,United States,Y
18944,Insel Air (7I/INC) (Priv),,7I,\N,,Netherlands Antilles,Y
18946,VivaColombia,,5Z,VVC,,Colombia,Y
18952,Flybe Finland Oy,,,FCM,FINNCOMM,Finland,Y
18959,Bingo Airways,Bingo,,BGY,,Poland,Y
19007,Bluebird Airways (BZ),,,BBG,,Greece,Y
19016,Apache Air,Apache,ZM,IWA,APACHE,United States,Y
19025,Taunus Air Gmbh,,,TAQ,Taunusair,Germany,N
19026,MHS Aviation GmbH,,M2,\N,,Germany,Y
19030,Jettor Airlines,Jettor,NR,JTO,JETHAPPY,Hong Kong,Y
19204,Eastern Atlantic Charters,,,EDI,,United States,N
19208,GoDutch,,GD,GOD,,Netherlands,N
19215,Flyme (VP),,,VQI,,Maldives,Y
19225,Thai Lion Air,,SL,\N,,Thailand,Y
19231,Deutsche Luftverkehrsgesellschaft,,DW,DLT,,Germany,N
19232,Nürnberger Flugdienst,,,NFD,,Germany,N
19244,Golden Myanmar Airlines,,,GMR,Golden Myanmar,Burma,Y
19262,ViznAir,,,VZA,Brian,United States,N
19276,Canaryfly,,,CNF,,Spain,Y
19280,Sunrise Airways,,,KSZ,,Haiti,Y
19287,National Air Cargo,,N8,NCR,,United States,Y
19290,Eastern Atlantic Virtual Airlines,,13,EAV,EAVA,United States,Y
19305,Citilink Indonesia,,QG,\N,SUPERGREEN,Indonesia,Y
19317,Gulisano airways,,GU,GU1,,Italy,N
19337,Transair,,,TTZ,,Canada,Y
19350,Comfort Express Virtual Charters Albany,,,EVC,Comfort Express,United States,Y
19351,Comfort Express Virtual Charters,,,CEO,,United States,Y
19358,Caribbean Wings,,XP,ZYZ,caribbean Wings,Turks and Caicos Islands,N
19359,FLYJET,,,FYJ,Fast Jet,Poland,Y
19361,Snowbird Airlines,,S8,SBD,,Finland,Y
19363,Russkie Krylya,,,KRY,,Russia,Y
19367,Kharkiv Airlines,,KH,KHK,,Ukraine,Y
19423,Key Air,,,KWY,KeyAir,United States,N
19433,XAIR USA,,XA,XAU,XAIR,United States,Y
19451,Air Costa,,LB,\N,,India,Y
19459,Simrik Airlines,,,RMK,,Nepal,Y
19465,Global Freightways,,F5,GF5,Freight,United States,N
19473,XPTO,XPTO  ,XP,XPT,XPTO,Portugal,Y
19474,Royal Flight,,,DME,,Russia,Y
19525,BBN-Airways,BlackBurn,,EGH,BBN,United Kingdom,Y
19531,Tomsk-Avia,,,TKS,,Russia,Y
19538,Vintage Props and Jets,,,VPP,VINTAGE,United States,N
19541,Malawian Airlines,,3W,\N,,Malawi,Y
19548,Yeti Airlines ,,,NYT,,Nepal,Y
19567,Avilu,Avilu' SA,..,...,,Switzerland,Y
19599,Skyline Ulasim Ticaret A.S.,Skyline Ulasim Ticaret A.S.,,KCU,Kocoglu,Turkey,Y
19619,Envoy Air,,,ENY,Envoy,United States,Y
19651,CARICOM AIRWAYS (BARBADOS) INC.,CARICOM AIRWAYS,,CCB,,Barbados,Y
19674,Rainbow Air (RAI),Rainbow Air (RAI),RN,RAB,Rainbow,United States,Y
19675,Rainbow Air Canada,Rainbow Air CAN,RY,RAY,Rainbow CAN,Canada,Y
19676,Rainbow Air Polynesia,Rainbow Air POL,RX,RPO,Rainbow Air,United States,Y
19677,Rainbow Air Euro,Rainbow Air EU,RU,RUE,Rainbow Air,United Kingdom,Y
19678,Rainbow Air US,Rainbow Air US,RM,RNY,Rainbow Air,United States,Y
19745,Transilvania,,,TNS,,Romania,Y
19751,Dobrolet,Добролёт,QD,DOB,DOBROLET,Russia,Y
19774,Spike Airlines,Aero Spike,S0,SAL,Spike Air,United States,Y
19776,Grand Cru Airlines,,,GCA,,Lithuania,Y
19785,Go2Sky,,,RLX,RELAX,Slovakia,Y
19803,All Argentina,All Argentina,L1,AL1,,Argentina,Y
19804,All America,All America,A2,AL2,,United States,Y
19805,All Asia,All Asia,L9,AL3,,China,Y
19806,All Africa,All Africa,9A,99F,,South Africa,Y
19807,Regionalia México,Regionalia México,N4,J88,,Mexico,Y
19808,All Europe,All Europe,N9,N99,,United Kingdom,Y
19809,All Spain,All Spain,N7,N77,,Spain,Y
19810,Regional Air Iceland,Regional Air Iceland,9N,N78,,Iceland,Y
19811,British Air Ferries,,??,??!,,United Kingdom,N
19812,Voestar,Voestar Brasil,8K,K88,,Brazil,Y
19813,All Colombia,All Colombia,7O,7KK,,Colombia,Y
19814,Regionalia Uruguay,Regionalia Uruguay,2X,2K2,,Uruguay,Y
19815,Regionalia Venezuela,Regionalia Venezuela,9X,9XX,,Venezuela,Y
19827,Regionalia Chile,Regionalia Chile,9J,CR1,,Chile,Y
19828,Vuela Cuba,Vuela Cuba,6C,6CC,,Cuba,Y
19830,All Australia,All Australia,88,8K8,,Australia,Y
19831,Fly Europa,,ER,RWW,,Spain,Y
19834,FlyPortugal,,PO,FPT,FlyPortugal,Portugal,Y
19845,FTI Fluggesellschaft,,,FTI,,Germany,N
`