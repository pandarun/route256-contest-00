package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var ______testOk1 = `1
otLLLrRuEe256LLLN`

var ______testOkResult1 = `route
256
-
`

func TestTerminal(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk1))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult1 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult1)
	}
}

var ______testOk2 = `1
LRLUUDE`

var ______testOkResult2 = `
-
`

func TestTerminal2(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk2))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult2 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult2)
	}
}

var ______testOk3 = `1
itisatest`

var ______testOkResult3 = `itisatest
-
`

func TestTerminal3(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk3))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult3 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult3)
	}
}

var ______testOk4 = `1
abNcdLLLeUfNxNx`

var ______testOkResult4 = `af
x
xb
ecd
-
`

func TestTerminal4(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk4))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult4 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult4)
	}
}

var ______testOk5 = `4
otLLLrRuEe256LLLN
LRLUUDE
itisatest
abNcdLLLeUfNxNx`

var ______testOkResult5 = `route
256
-

-
itisatest
-
af
x
xb
ecd
-
`

func TestTerminal5(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk5))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult5 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult5)
	}
}

var ______testOk6 = `2
UNR
mUo`

var ______testOkResult6 = `

-
mo
-
`

func TestTerminal6(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk6))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult6 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult6)
	}
}

var ______testOk7 = `1
UNR`

var ______testOkResult7 = `

-
`

func TestTerminal7(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk7))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult7 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult7)
	}
}

var ______testOk12 = `100
5ELN6qt3uNLBqNqfLqmB
UERL9NLioeLRBDweE8UD05L3Dzdx1
RreEfB7D2tNUzE11BN7s9kpUg9ivhBDlEBUxEBkNvUvRLg
L242vlLysD2Nwym9bq4BxEULc0NUjlDLL4mkEUyE8bwRL9
xRdLDLdUbaukREUUB0sRtNDDtsaL0lsg1nD
NgvU6DL8lszu
g4xyRDLUwyssR
ogEyELDBNz
RyLLEEqcuiRDD1RNRBb7B
Dv9NDc5ck0UNgEBeB5U8gmEoUBpk1L0EBL0vRBuUNbENDLB
ENUEi
UBzEgDbUNcDUBBLxLeftEdN5ijLaUkBjNEBDBdwzUhaLLtEBB
dNELpiiEE6xBUnEDcEN3bcBrttRBdENgRLmLR8RL9E4NNw7E
R566DesNwUrEylUeL1vREwDEw9Uo93dL91Uosv
EBLbdroUeBDRU87LqEUB1LgUn
RkB
ogDBhLUBUzE
BB9yzUhbByaBquljNNE83UeigRExR2UEL
BBURE06RhLmNLNDiBLLEc1EdRuDUE
BzLi261LRw4LUD7DLELwLBdRUyrDN97ybEL
RNaUL06iDRbUqDtahNBkhDzl9LyE1gNvLU2tLLUE
qxhU
BD0UEg1L9jnURiDxqtB6RU1Uv56N3vRRmbsRNUU
bEtRLBg1Doyh
kBo
bR8k5R7Li1uh4lDRqliEDEBDvUNs
E
99UEcrfoRxcL1ose0
DkBEsNr27l9tUUREBrNlmrNcrgN0k
EBEjUuDDERBmsbakUp18DEeBxLNEE
DLEUBEmb5DLoLLzaBrafR
3UkkBLBi3dddqu2
UUnsULERRDRLpNUc7dRtButlDBc17hE7jUDB9DEqD
NrBR1vDBBhLBpRsnhn2DNBq
q
BBN
pqtsULywDEN6N96N04Ewp6l5DENRDmgLUxNBUL8j3ytURxRl2
aBDtUuD3ui5aDLEoELBEBlR
0o9EB4cef
N4EwRcBzLUNE4B4RRRfRDD8DEl
U1nUBUbBpUxmvBByDU6U28
RqUEBUkDEcNDzBjDUDNvNc4U6wDut5RD
UN7NLBLE6LEmLL
joBBBfNN
LEBNzpLrREmjd9DDlDNULEEhz0EENNENrk1fnoDBwkR
BfRLj0iBpvot4B4sBDvRd
DuBEEBRDRdNl
021
weERh4NRNBE8fyDNUqENRqbBEx
fnwN9cEUsD12NNmRdLtd
BEEUsLlR
D7sU4NRyEN
LNLezBDuNUnBRtv
LDeD5tixlBUycB5Ll6LsoEDLftpN7BEc0
DDR
RunmN4EsNoDmR
BRUkEEhmEoBBE
yzBaysfdRUDDs0NU
dNzRxL2DfDbR04D8DDLm1DBhtUyLlz4UoN
BENDU7NRc9m8Dx5j6NBREwBu0DNE6BqLNR
UmNRzRh5k3DDUdwBl17EBtyDmBUilE
h93NBxBiEREfE0zBuLDiELNzBBmlRm3kLDvNLUEn
8NBEsL0fBupNr
LEULU8DBqUEDUnuLNE20lraNuLz11Dc
cLRLDLDg0ND5ERUbUBRRN27kULl
L8mveEwNRytL2NtfL8Lq6
UBk79LLNkNNBRjvxLU2UeE1uUnciN8Ba
9t8hc9c9UDhmngBEyDUNwNjw3ayUpNRtEvvR
3rhND97N2rBweULuqU9ivmENRL2xtDE
mBzREqR7ERUuR5fk8
EsDwNkoRLefrf0BsvsbBcBRxLaLENcjf81
LcB2fUu9p0DoLcBDEbRB
ELup5kDgReyEqEN7U1NdUvBNllb4Nj1DtRRpBv
LDBkNDBNvcL29tU2NB8NN1lz27UeERpUhDDuyULtxaUR
EN5536Ngm2BuBEjLaaREDedDzEUjUlNDR6N4mEpRajND
jU6UqEBRqoByNNLm52sEkkBmsNRd2UB1jUBl6D2DiNfEE
srg4lLoRLELBnExpDEB9D6E2
BLDh2E
e7LBriUvjUUBNjRL2fNNaRB6E4kBBDmihwnRUEDDnE8t2
UUr
E0mDDDRR9BEzkpRs3aBUqRv
cBcDRqBENRrxLDrR63gLmoEdNB4N3EDx8zBbjL2tURpmU2RaD
NLDjUpcEmfsRBBohDmBdvLLLBBUBph8UjUDELUL09U3kvNU
ENDw198LEvBLDUB2DBtUlErsNNyg07jDh2U0Nuhw1rERE
DuDpkRR5qL4EED7LDms37BBEs
f8RsDBzLNRldiBhD
R9N6LBDUiibUj24LoDBDU6BE7
decrEBEtNEpED98RflNRev
NyR9wUDEpls1UoN1dokRNDND5bt4UkBBlDEk
ug3DBDr9UqpvEibNULdER6DE64xNB
lNEoDRpeUcBD2LURBhNkxwLRtN1gEBRc4UD
LieqU
EUBRwRR
gBDRURBdBDBvELaposfmERRo7
bxd583vRfB0wNnDRNRlNU
tvi9NBR1beUwN8U9atlNBl4Dli
sUUvsEEDNELRlERlULtkxp6LdB8aB0jLBLNDk
D7hNwsLi3LL2Nh
cR
rrjE1gNDRLdtDoxt9R
`

var ______testOkResult12 = `
6qt3u
q
qqmf5
-
9
w03zdx15eioe8
-
kgv
vxg9ivh
l7s9kpz72t11
ref
-
jl24y2vysc08b9w
4mk2
xwym9bq4l
-
0sdt
ts0lsg1nabaukxd
-
6
8lszugv
-
g4xwyssy
-

zogy
-
yqcui1
b7
-
u
b0vpk01v8gm9o

5eg
c5ck0
-
i

-
j
eftthakxzgbd
dwz5iaj
c
-
nd
pici6x
drtt3bc
m89g4

w7
-
5r6o9391osvd6esyl1vew
ww9
-
gn1b8q7droe
-
k
-
zhog
-
qulj
eigx2
ya9yzhb83
-
06m

ihc1du
-
diyr
97yb261w74wz
-
06qi
abtah
2tkhzly91g
v
-
qxh
-
601v56
3vg9mbsj
n1ixqt
-
g1oyhbt
-
ok
-
v
sb8k5i1uh4l7qli
-

-
99crfox1ose0c
-
r
lmr
crg
0kks
r27l9t
-

xmsbakp18jue
-
rafmzabo5
-
i3dddqu23kk
-
utlc7dntp
9c17hs7jq
-

phsnhn2
qr1v
-
q
-


-
pqtyws
6
96xl2
8j3yt0x
4wp6l5
mg
-
ltu3ui5aao
-
4cef0o9
-

44f
z4w8cl
-
y628pxmvb1n
-
kqc
j
v6w
c4zut5
-

7
6m
-
f

jo
-

zrpmjd9lhz0


rk1fno
wk
-
v4dspvot4j0if
-
ud
l
-
021
-
weh4

q8fy
qbx

-
fnsw
9c12

mtdd
-
ls
-
7s4
y

-

ntvu
ez
-
lso65yce5tixftp
7lc0
-

-
unm
4s
om
-
khmo
-
aysfdys0
z
-
dlz4o
y
htz2fbx04m18
-
7
c9m8
x5j6
u0

qw6
-
iltyl17mdw
zhm5k3
-
h93
iuixf0
mlzm3vn
kz
-
8
up
r0fs
-
q8n
u20lra
z11cu
-
gl0
27kb
5c
-
8mvew
y2
tq68ft
-
knci
a8
ke1u
2
7jvx9
-
9t8hc9c9hmngy
wp
tvv
jw3ay
-
3rh9ivm
2xt
9uq7
we2r
-
zmq7u5fk8
-
sw
caxsvsbkefrf0o
cjf81
-
2fu9p0cocb
-

llb4
j1uv1
vdpt5kpgeyq
7
-
k
2
8h
etxap
1luyz27
v29tc
-
l

56
4m536jpaj

ugm2aajedz
-
y
l6
1j2ms
md2i
f52sjqo6qkk
-
96nsrg4olxp2
-
h2
-

j2f

nmihwn6arivje74k8t2
-
r
-
q0vm9zkps3a
-
ccq
rrx26a3mogd
4pm
b2tj3x8z
-
ph093kv
8johpcmfs
dvjm
-
2lrs

yg07j0
uhw1r
tw198h2v
-
upk54qms377s
-

hzldif8s
-
6iibj2o497
6
-
decrt
p98fl
ev
-
o
1dok

lk
5bt4y9wpls1k
-
dr9qpvug3ib6
64x

-
h
kxwt
1c4glc
2ope
-
ieq
-
w
-
vdaposfmgo7
-
0w
nb
xl
d583vf
-
t9atl
l4viw
89li
1be
-

0j8astkxpd6vs
kll
-
7h
w2
hi3s
-
c
-
rrj1g
dtoxt9
-`

func TestTerminal12(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk12))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult12 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult12)
	}
}

var ______testOk13 = `1
UERL9NLioeLRBDweE8`

var ______testOkResult13 = `9
weioe8
-
`

func TestTerminal13(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk13))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult13 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult13)
	}
}

var ______testOk14 = `1
UERL9NLioeLRBDweE8UD05L3Dzdx1`

var ______testOkResult14 = `9
w03zdx15eioe8
-
`

func TestTerminal14(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(______testOk14))
	out := new(bytes.Buffer)
	err := terminal(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != ______testOkResult14 {
		t.Errorf("test for OK failed - results not match\n %v %v", x, ______testOkResult14)
	}
}
