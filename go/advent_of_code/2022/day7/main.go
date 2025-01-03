package day7

import "fmt"

func main() {
	root := buildTree(input())
	totalSize := bfsLessThan100000(root)
	fmt.Println(totalSize)
}

func input() string {
	return `$ cd /
$ ls
dir gqcclj
dir lmtpm
dir nhqwt
dir qcq
dir vwqwlqrt
$ cd gqcclj
$ ls
62425 dqp.gjm
174181 hrtw.qsd
273712 pflp.mdw
169404 zlthnlhf.mtn
180878 zprprf
$ cd ..
$ cd lmtpm
$ ls
dir clffsvcw
163587 cvcl.jqh
dir dcqnblb
dir dtpwln
dir fvt
dir hrcrw
dir jdqzmqn
236754 nrdmlj
205959 pflp.mdw
dir qcq
dir rsn
129926 vdgcqdn.sqd
dir zprprf
$ cd clffsvcw
$ ls
6997 dcqnblb.wbh
145711 dqp
159225 pflp.mdw
$ cd ..
$ cd dcqnblb
$ ls
dir dcqnblb
dir gfn
dir lpswsp
dir lvt
dir zprprf
$ cd dcqnblb
$ ls
2020 grpdmd.ggz
dir zpswzfvg
$ cd zpswzfvg
$ ls
206998 zprprf.gnw
$ cd ..
$ cd ..
$ cd gfn
$ ls
277530 rhbvtblc.mvw
$ cd ..
$ cd lpswsp
$ ls
173180 dcqnblb
$ cd ..
$ cd lvt
$ ls
dir hjllwsvl
dir ptbt
$ cd hjllwsvl
$ ls
dir wqnc
$ cd wqnc
$ ls
64695 grpdmd.ggz
$ cd ..
$ cd ..
$ cd ptbt
$ ls
150880 vvbt.gtp
$ cd ..
$ cd ..
$ cd zprprf
$ ls
dir ldzslndn
dir qftt
$ cd ldzslndn
$ ls
dir bwqqsbhg
129454 vbn
$ cd bwqqsbhg
$ ls
108701 zprprf.gss
$ cd ..
$ cd ..
$ cd qftt
$ ls
64268 cvcl.jqh
$ cd ..
$ cd ..
$ cd ..
$ cd dtpwln
$ ls
196215 cvcl.jqh
dir dpwg
dir ldzslndn
dir znnsqqh
$ cd dpwg
$ ls
192388 gmh
47754 grgzh.qdl
99449 hqsh
dir pbmf
50061 pflp.mdw
192902 qcq.pgg
dir rmpvj
dir scgc
$ cd pbmf
$ ls
210083 wpfnwbl.mgf
$ cd ..
$ cd rmpvj
$ ls
125738 nmlnbvrd
226214 zprprf.jnp
114257 zprprf.srs
$ cd ..
$ cd scgc
$ ls
182115 rrc.rcc
$ cd ..
$ cd ..
$ cd ldzslndn
$ ls
201992 qcrm.cpd
$ cd ..
$ cd znnsqqh
$ ls
85635 cvcl.jqh
$ cd ..
$ cd ..
$ cd fvt
$ ls
dir dcqnblb
dir gnc
75864 vfn
$ cd dcqnblb
$ ls
dir dcqnblb
dir lbnflwsh
$ cd dcqnblb
$ ls
269901 cvcl.jqh
$ cd ..
$ cd lbnflwsh
$ ls
33336 grpdmd.ggz
42861 phg.wmc
$ cd ..
$ cd ..
$ cd gnc
$ ls
dir jhjbjsp
dir jjppr
$ cd jhjbjsp
$ ls
96177 ldzslndn
$ cd ..
$ cd jjppr
$ ls
181016 dqp
$ cd ..
$ cd ..
$ cd ..
$ cd hrcrw
$ ls
261376 dtjfpppr.dww
54658 vsrgvw.pfn
$ cd ..
$ cd jdqzmqn
$ ls
52342 dcpndc.vlg
171946 gggpchh.tbb
dir ldzslndn
11156 nbfrfvv.gzw
$ cd ldzslndn
$ ls
107873 cvcl.jqh
216034 gfdjrbz
68844 pqllfrrh.jcf
$ cd ..
$ cd ..
$ cd qcq
$ ls
152886 ldzslndn.ltn
105125 vwplh.vbf
$ cd ..
$ cd rsn
$ ls
15385 hqcmjdgv.jjv
105735 qcq.bzg
58805 snczcsp
26668 vbn
$ cd ..
$ cd zprprf
$ ls
dir chbmq
dir dcqnblb
dir dqp
dir nfspb
89506 zprprf.hnt
$ cd chbmq
$ ls
dir cnjvw
dir dqp
151434 frsvrdnt
dir msztjvcb
240689 qcq.jlh
dir sjzrcg
97312 vnr.zfr
dir zprprf
$ cd cnjvw
$ ls
dir bpbs
252403 cqhtshc
dir djmjhn
10935 fhqmswr
6582 pdwml.ldd
dir qcq
219282 rfmd
$ cd bpbs
$ ls
147582 bnhwsnsj.gdm
61362 cvcl.jqh
152857 vdgcqdn.sqd
$ cd ..
$ cd djmjhn
$ ls
dir bjdbcjbb
dir dcqnblb
dir dqp
dir lgdwtt
$ cd bjdbcjbb
$ ls
110710 cvcl.jqh
252792 hmshctr.lgz
dir mjhtmbj
189745 shsswcgr
dir tfnhp
194940 vbn
dir zprprf
$ cd mjhtmbj
$ ls
dir dqp
dir hbthpcmb
$ cd dqp
$ ls
200832 sbcrz.qgw
$ cd ..
$ cd hbthpcmb
$ ls
55191 ffcntg
$ cd ..
$ cd ..
$ cd tfnhp
$ ls
276825 dqp
161538 gqmr.wgb
$ cd ..
$ cd zprprf
$ ls
287638 dcqnblb.ssp
41274 hgmrvj.mwf
249118 sbb.gsf
105141 wwrg.gqz
$ cd ..
$ cd ..
$ cd dcqnblb
$ ls
1957 btmmc
32386 dtzbzg.dhm
dir mmrbj
98283 ntmhfgtl.pmf
dir zprprf
$ cd mmrbj
$ ls
273194 wnsq
251527 zprprf
$ cd ..
$ cd zprprf
$ ls
27678 ldzslndn.rrl
62866 ljf.fdj
148502 qcq.dlg
dir rvgqvm
179231 tllnmhn.pjp
64033 vbn
dir zcdrj
$ cd rvgqvm
$ ls
dir ntbv
262324 prhgj.szz
dir qbvdh
$ cd ntbv
$ ls
116608 cgv.fvj
175200 swpswq.twt
$ cd ..
$ cd qbvdh
$ ls
160353 sdhfrb.wjn
$ cd ..
$ cd ..
$ cd zcdrj
$ ls
283262 ctl
$ cd ..
$ cd ..
$ cd ..
$ cd dqp
$ ls
dir jfzm
111438 rdrgb.mjf
64194 wgtmqrq
dir zprprf
$ cd jfzm
$ ls
158774 pflp.mdw
$ cd ..
$ cd zprprf
$ ls
215264 sgsstcp
$ cd ..
$ cd ..
$ cd lgdwtt
$ ls
dir qcq
$ cd qcq
$ ls
165461 ldzslndn.vvb
$ cd ..
$ cd ..
$ cd ..
$ cd qcq
$ ls
dir dpd
165044 grpdmd.ggz
82343 ldzslndn
dir mwg
176689 psjcwp.wct
44404 qcq.zwd
$ cd dpd
$ ls
84087 dqp
227386 zprprf.gfs
$ cd ..
$ cd mwg
$ ls
214086 pflp.mdw
dir sjjsdn
225859 wcdt
158892 zprprf.frs
$ cd sjjsdn
$ ls
260121 gplgp.dfn
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd dqp
$ ls
dir hcrwclpg
dir zphd
$ cd hcrwclpg
$ ls
dir cmqntjj
16393 ldzslndn.qbm
91152 qqdtc.zdq
$ cd cmqntjj
$ ls
272266 ldzslndn.pll
$ cd ..
$ cd ..
$ cd zphd
$ ls
165711 chftwcsw.fqw
256871 cvcl.jqh
251168 zprprf.gfv
$ cd ..
$ cd ..
$ cd msztjvcb
$ ls
206231 brzn.lmn
dir dcqnblb
21571 dqp
dir fmn
45779 mlfctz.cjr
288827 pflp.mdw
220578 qcq.fqf
$ cd dcqnblb
$ ls
198121 ghbwgs
93681 nmqhl.vpq
$ cd ..
$ cd fmn
$ ls
29407 mdfws.qvs
$ cd ..
$ cd ..
$ cd sjzrcg
$ ls
155120 ddclvsjr.rpq
136029 ldzslndn.dcm
dir vhzh
$ cd vhzh
$ ls
212446 vbn
$ cd ..
$ cd ..
$ cd zprprf
$ ls
240335 crt.gqh
185363 gnmm.qgh
dir ldzslndn
dir nwl
dir qll
277043 vbn
217796 vtvgpdl.vtm
$ cd ldzslndn
$ ls
273570 cvcl.jqh
68510 fgdmz.hrc
dir npq
dir swjrzzrm
$ cd npq
$ ls
97923 dzcjsqwt
$ cd ..
$ cd swjrzzrm
$ ls
180599 tmpgn.bjf
$ cd ..
$ cd ..
$ cd nwl
$ ls
171833 dlwrfhh.qgn
$ cd ..
$ cd qll
$ ls
219926 dcqnblb.bvn
$ cd ..
$ cd ..
$ cd ..
$ cd dcqnblb
$ ls
dir lvpb
276198 tbgcm.qct
$ cd lvpb
$ ls
142590 bvhjlld
268259 gnjfg.sgb
dir qcq
206220 qcq.zsg
258137 rrsw.dnb
dir tmr
215549 vbn
$ cd qcq
$ ls
dir mmpgd
dir tdsz
dir tmfvsjwc
$ cd mmpgd
$ ls
70793 jwbnpwnn
$ cd ..
$ cd tdsz
$ ls
246310 tdvrhhg.bzq
$ cd ..
$ cd tmfvsjwc
$ ls
103899 grpdmd.ggz
287850 ldzslndn
125930 llhr
$ cd ..
$ cd ..
$ cd tmr
$ ls
83344 fbtfcg.hqp
$ cd ..
$ cd ..
$ cd ..
$ cd dqp
$ ls
dir lbgmcbv
dir nbg
$ cd lbgmcbv
$ ls
81776 wzdzzdp
$ cd ..
$ cd nbg
$ ls
dir mfsgjp
155574 pflp.mdw
$ cd mfsgjp
$ ls
199400 vdgcqdn.sqd
$ cd ..
$ cd ..
$ cd ..
$ cd nfspb
$ ls
262412 csrdtbs
73867 vbn
136389 zqps.hjt
$ cd ..
$ cd ..
$ cd ..
$ cd nhqwt
$ ls
123766 cvcl.jqh
dir dhrtvctp
222086 grpdmd.ggz
dir gzg
26005 lhpmz.tgz
dir mcnjwwfr
117122 msn.gst
$ cd dhrtvctp
$ ls
224079 vdgcqdn.sqd
$ cd ..
$ cd gzg
$ ls
124395 dqp
dir wqdbtqm
$ cd wqdbtqm
$ ls
237354 pflp.mdw
212019 vdgcqdn.sqd
$ cd ..
$ cd ..
$ cd mcnjwwfr
$ ls
92504 cshdztf
dir dctl
dir dqp
dir flcrmhlj
161879 grpdmd.ggz
dir gtt
dir hlbnhchz
220093 mdtdsgvm.zgg
dir twntr
287192 vbn
$ cd dctl
$ ls
dir bbhch
155396 hrrj.jzm
164971 pblqmwj.vdb
dir wnlgfpvf
$ cd bbhch
$ ls
dir dpqtp
dir jvdrcw
$ cd dpqtp
$ ls
174135 gwb.qrb
$ cd ..
$ cd jvdrcw
$ ls
215993 dcqnblb.cqp
200800 stjttf.ngc
$ cd ..
$ cd ..
$ cd wnlgfpvf
$ ls
135978 cvcl.jqh
dir dqp
54018 lbrfmt
$ cd dqp
$ ls
270516 dcqnblb.jqw
dir dqp
144626 grpdmd.ggz
157731 hvcv.rhp
133773 lnnt
76250 vdgcqdn.sqd
$ cd dqp
$ ls
41504 zprprf.cmc
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd dqp
$ ls
dir dqp
dir ldzslndn
236737 mqzcvm.fjh
239746 nhcdz.ncj
dir rpchqq
248824 vdgcqdn.sqd
250937 zrchht.mwg
$ cd dqp
$ ls
203381 qcq.djm
$ cd ..
$ cd ldzslndn
$ ls
dir dqp
dir fptnzlv
dir gmbnpm
dir vhvblt
$ cd dqp
$ ls
19579 qcq.lhg
$ cd ..
$ cd fptnzlv
$ ls
209930 dcqnblb
$ cd ..
$ cd gmbnpm
$ ls
dir ldzslndn
dir qcq
$ cd ldzslndn
$ ls
11075 pflp.mdw
$ cd ..
$ cd qcq
$ ls
dir tdp
$ cd tdp
$ ls
40741 vdgcqdn.sqd
$ cd ..
$ cd ..
$ cd ..
$ cd vhvblt
$ ls
dir lzr
$ cd lzr
$ ls
62245 gbnj.llg
$ cd ..
$ cd ..
$ cd ..
$ cd rpchqq
$ ls
dir bcs
dir dcqnblb
dir fvjzn
dir lrphzrv
$ cd bcs
$ ls
179794 bbn.dzb
242069 cmjdmzjf.zgf
1703 cvcl.jqh
dir gnmhwj
dir ldzslndn
152520 qltpsz.jsj
dir sqqjfps
$ cd gnmhwj
$ ls
dir gvs
201600 hptn.ftf
dir hzrnb
dir qcq
dir sqhl
$ cd gvs
$ ls
152358 zprprf.mlh
$ cd ..
$ cd hzrnb
$ ls
94290 gplsfd
$ cd ..
$ cd qcq
$ ls
91909 vmqd.bmg
$ cd ..
$ cd sqhl
$ ls
238673 vdgcqdn.sqd
262885 zmdvr.nfg
$ cd ..
$ cd ..
$ cd ldzslndn
$ ls
240461 mdz
84303 qtj
$ cd ..
$ cd sqqjfps
$ ls
88753 fwn.tff
$ cd ..
$ cd ..
$ cd dcqnblb
$ ls
dir dqp
189996 dqp.pvp
$ cd dqp
$ ls
dir qvfjz
196506 vbn
$ cd qvfjz
$ ls
209316 pflp.mdw
107459 rwpbh.vpt
$ cd ..
$ cd ..
$ cd ..
$ cd fvjzn
$ ls
241464 cvcl.jqh
dir dqp
dir ldzslndn
dir msp
125 pflp.mdw
131895 vbn
$ cd dqp
$ ls
34019 pflp.mdw
202957 vbn
$ cd ..
$ cd ldzslndn
$ ls
147492 cvcl.jqh
248719 spc.rfv
$ cd ..
$ cd msp
$ ls
184407 cvcl.jqh
$ cd ..
$ cd ..
$ cd lrphzrv
$ ls
dir bbwqmbg
81858 cvcl.jqh
dir dqp
248670 gqqsww.tsn
199141 grpdmd.ggz
dir ldzslndn
34514 ldzslndn.ctw
dir tln
214615 zprprf.fwm
$ cd bbwqmbg
$ ls
129750 flf
dir pvlw
dir qcq
126 sqcqphz.tbm
$ cd pvlw
$ ls
198005 jfvj.hdv
$ cd ..
$ cd qcq
$ ls
dir wgdzws
$ cd wgdzws
$ ls
253522 ldzslndn.qwt
$ cd ..
$ cd ..
$ cd ..
$ cd dqp
$ ls
281993 cvcl.jqh
dir hwqjlwcb
50532 msccz.qgm
102187 trv.tnq
111 wplnmj.bfl
$ cd hwqjlwcb
$ ls
267580 dhjqb.dsb
153195 ldzslndn.jqv
41526 mvwcwc.zsc
$ cd ..
$ cd ..
$ cd ldzslndn
$ ls
58666 cvcl.jqh
79950 dqp.tmc
242217 hns.lrb
dir njswzh
240692 vdgcqdn.sqd
dir zvmjvcdm
52909 zzh
$ cd njswzh
$ ls
149732 cvcl.jqh
dir rnmfd
$ cd rnmfd
$ ls
75368 dqp.hmv
14350 vbn
$ cd ..
$ cd ..
$ cd zvmjvcdm
$ ls
dir jgczt
$ cd jgczt
$ ls
dir qcq
95941 qzvvwshv.jwc
$ cd qcq
$ ls
273942 pflp.mdw
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd tln
$ ls
dir bmcng
1518 lrg
dir vnjfrhp
$ cd bmcng
$ ls
38917 fqcrt
$ cd ..
$ cd vnjfrhp
$ ls
dir dcqnblb
dir dqp
247186 grpdmd.ggz
dir ldzslndn
169216 pflp.mdw
206487 vdgcqdn.sqd
16976 vlsrzjmb.mmc
257938 wjl
$ cd dcqnblb
$ ls
dir dqp
$ cd dqp
$ ls
184133 qcq
$ cd ..
$ cd ..
$ cd dqp
$ ls
dir dcqnblb
31612 dqp.pnt
212283 ldzslndn
61600 vdbfc.ddj
197189 wpv.wff
$ cd dcqnblb
$ ls
62412 tfzllmrj
dir zprprf
$ cd zprprf
$ ls
dir bqnpsl
dir dszrvpzc
$ cd bqnpsl
$ ls
261548 spbsbbsw.cmn
$ cd ..
$ cd dszrvpzc
$ ls
188232 sggpqslr.smn
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ldzslndn
$ ls
dir bgnhd
dir pgvcdzwz
dir qgzhm
$ cd bgnhd
$ ls
56989 cvcl.jqh
$ cd ..
$ cd pgvcdzwz
$ ls
110034 qhgnndv
$ cd ..
$ cd qgzhm
$ ls
247232 grpdmd.ggz
269292 ldzslndn
153843 tpz
dir vnschqwr
162392 wnq.btb
$ cd vnschqwr
$ ls
43005 fvtvzfqm.jvc
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd flcrmhlj
$ ls
245668 dcqnblb.sdj
dir lffj
229909 pflp.mdw
280176 vbn
$ cd lffj
$ ls
116451 jmzz.jdd
dir pjlwb
162815 pmhlqq.snr
226183 zffth
$ cd pjlwb
$ ls
67518 qcq.hjq
$ cd ..
$ cd ..
$ cd ..
$ cd gtt
$ ls
52105 grpdmd.ggz
126869 zprprf.fgj
$ cd ..
$ cd hlbnhchz
$ ls
3064 dqp.lrw
278756 grpdmd.ggz
177208 ldzslndn.wlv
141685 vbn
$ cd ..
$ cd twntr
$ ls
63747 cvcl.jqh
$ cd ..
$ cd ..
$ cd ..
$ cd qcq
$ ls
226858 cwblp.zgp
dir jjqsmfhr
dir rjbqtrq
dir vwmpnbts
141715 wdbhdch
286381 zprprf
$ cd jjqsmfhr
$ ls
dir btmm
dir fqndtlgq
$ cd btmm
$ ls
4031 dqp.lrr
dir fzdd
$ cd fzdd
$ ls
dir vnwpn
$ cd vnwpn
$ ls
dir bzlgsl
dir ztvzrrbv
$ cd bzlgsl
$ ls
9294 ldzslndn.sqr
$ cd ..
$ cd ztvzrrbv
$ ls
256017 cvcl.jqh
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd fqndtlgq
$ ls
271528 ccbmgp.bwd
$ cd ..
$ cd ..
$ cd rjbqtrq
$ ls
122150 ldzslndn
46467 tpdvp.pjf
$ cd ..
$ cd vwmpnbts
$ ls
47518 fcrwfzvm
263343 gmc.lrt
212764 qcq
$ cd ..
$ cd ..
$ cd vwqwlqrt
$ ls
dir psrs
$ cd psrs
$ ls
281998 zprprf.hml`
}