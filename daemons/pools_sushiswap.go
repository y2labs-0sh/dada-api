package daemons

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/y2labs-0sh/dada-api/types"
)

const (
	DaemonNameSushiswapPools = "sushiswapPools"
	SushiswapGraphURI        = "https://api.thegraph.com/subgraphs/name/zippoxer/sushiswap-subgraph-fork"
)

var (
	sushiswapPoolsOnce   = sync.Once{}
	sushiswapPoolsDaemon *sushiswapPools
)

type sushiswapPools struct {
	fileStorage

	graphQL string
	logger  Logger

	list     PoolInfos
	listLock sync.RWMutex
}

type SushiswapPoolInfo struct {
	ID     string `json:"id"`
	Token0 struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Symbol         string `json:"symbol"`
		TotalLiquidity string `json:"totalLiquidity"`
	} `json:"token0"`
	Token1 struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Symbol         string `json:"symbol"`
		TotalLiquidity string `json:"totalLiquidity"`
	} `json:"token1"`
	Reserve0           string `json:"reserve0"`
	Reserve1           string `json:"reserve1"`
	ReserveUSD         string `json:"reserveUSD"`
	ReserveETH         string `json:"reserveETH"`
	TotalSupply        string `json:"totalSupply"`
	VolumeUSD          string `json:"volumeUSD"`
	VolumeToken0       string `json:"volumeToken0"`
	VolumeToken1       string `json:"volumeToken1"`
	Token0Price        string `json:"token0Price"`
	Token1Price        string `json:"token1Price"`
	TrackedReserveETH  string `json:"tracktedReserveETH"`
	UntractedVolumeUSD string `json:"untractedVolumeUSD"`
}

func NewSushiswapPoolsDaemon(l Logger) Daemon {
	sushiswapPoolsOnce.Do(func() {
		newSushiswapPoolsDaemon(l)
	})
	return sushiswapPoolsDaemon
}

func newSushiswapPoolsDaemon(l Logger) {
	const query = `{"operationName":"pairs","variables":{"allPairs":["0x397ff1542f962076d0bfe58ea045ffa2d347aca0","0xc3d03e4f041fd4cd388c549ee2a29a9e5075882f","0x06da0fd433c1a5d7a4faa01111c044910a184553","0x795065dcc9f64b5614c407a6efdc400da6221fb0","0x088ee5007c98a9677165d78dd2109ae4a3d04d0c","0xc40d16476380e4037e6b1a2594caf6a6cc8da967","0xf1f85b2c54a2bd284b1cf4141d64fd171bd85539","0xa1d7b2d891e3a1f9ef4bbc5be20630c2feb1c470","0x31503dcb60119a812fee820bb7042752019f2355","0xa75f7c2f025f470355515482bde9efa8153536a8","0xd75ea151a61d06868e31f8988d28dfe5e9df57b4","0xceff51756c56ceffca006cd410b03ffc46dd3a58","0xdafd66636e2561b0284edde37e42d192f2844d40","0xcb2286d9471cc185281c4f763d34a962ed212962","0x001b6450083e531a5a7bf310bd2c1af4247e23d4","0x611cde65dea90918c0078ac0400a72b0d25b9bb1","0xfcff3b04c499a57778ae2cf05584ab24278a7fcb","0x0289b9cd5859476ce325aca04309d36addcebdaa","0x117d4288b3635021a3d612fe05a3cbf5c717fef2","0x382c4a5147fd4090f7be3a9ff398f95638f5d39e","0x58dc5a51fe44589beb22e8ce67720b5bc5378009","0xee6d78755e06c31ae7a5ea2b29b35c073dfc00a9","0x0f82e57804d0b1f6fab2370a43dcfad3c7cb239c","0xd3f85d18206829f917929bbbf738c1e0ce9af7fc","0xd597924b16cc1904d808285bc9044fd51ceeead7","0x9fc5b87b74b9bd239879491056752eb90188106d","0x5e63360e891bd60c69445970256c260b0a6a54c6","0x69b39b89f9274a16e8a19b78e5eb47a4d91dac9e","0x6f58a1aa0248a9f794d13dc78e74fc75140956d7","0xbcd6a2ddafbaa7f424698ed69e717c0c0f1e99bf","0x36e2fcccc59e5747ff63a03ea2e5c0c2c14911e7","0x97f34c8e5992eb985c5f740e7ee8c7e48a1de76a","0x4f871f310ad0e8a170db0021c0ce066859d37469","0xf169cea51eb51774cf107c88309717dda20be167","0x15e86e6f65ef7ea1dbb72a5e51a07926fb1c82e3","0xba13afecda9beb75de5c56bbaf696b880a5a50dd","0xc5fa164247d2f8d68804139457146efbde8370f6","0xc926990039045611eb1de520c1e249fd0d20a8ea","0xc4b26b26d720467d96e18f08664a888d4116cea6","0x95b54c8da12bb23f7a5f6e26c38d04acc6f81820","0x6463bd6026a2e7bfab5851b62969a92f7cca0eb6","0x9cbc2a6ab3f10edf7d71c9cf3b6bdb7ee5629550","0x0bff31d8179da718a7ee3669853cf9978c90a24a","0x34b13f8cd184f55d0bd4dd1fe6c07d46f245c7ed","0xe4455fdec181561e9ffe909dde46aaeaedc55283","0x53aabccae8c1713a6a150d9981d2ee867d0720e8","0x00c70e8b3c9d0e0adb85993382abaae2a11c5d96","0xfcfee7769c012578877dc787a5d9a339cc5920a1","0x807fdbcb54dc5d01b99a5aa7fce883a759d27cbe","0xd11684e2ea44128c26877376cb75b9c36e8381dd","0xdfb87fd2f9f5eec75930a2d8cfe67f791801fe30","0x5a2943b25ce0678dc0b351928d2db331a55d94ea","0xe1996640c0e7ae4ddca6310a25879b8255c7e92f","0x17b3c19bd640a59e832ab73eccf716cb47419846","0x2d2bc284c24b5debe489da59557862e0af884f23","0x742c15d71ea7444964bc39b0ed729b3729adc361","0x1ccc8ee624fbe8c10917b96d1f0672b480f0d671","0x561aef36a2237f86696e15b131f7ff3e141feab2","0x680a025da7b1be2c204d7745e809919bce074026","0xdd77c93199064a53e1db19ee0930bcdf7c9999f4","0x269db91fc3c7fcc275c2e6f22e5552504512811c","0x0be88ac4b5c81700acf3a606a52a31c261a24a35","0xd989f9aca3dff9db16d3b0cebc252596bb8aa63e","0xc9116009b44afc73f319c2011d7c6c698873bf13","0x0cfe7968e7c34a51217a7c9b9dc1690f416e027e","0xd86a120a06255df8d4e2248ab04d4267e23adfaa","0x4fccb922148284270ce8b2fca8c00519563a4dda","0x3e5e560db64c87e71260878e58d74406e6d693af","0xdd2f8d990d7352bb38bd0340c394af4691d9957c","0xf05fd86c9b2f5e883165fc7de5e41f65ac5494c7","0xad1894d8b708a11f627d6a0355c5e6b2e0beb606","0xd626015edfbbeba3def1d764a162a0b4d7bf22f0","0x8e14b803a730cc3bcc8af4ffcd9c4b10ebff7215","0x06f6f3f79bb7236473c144bc43b687428ac0262e","0x9a67737d2ab958d80a763bbbd4b1c826249745f4","0x33f172212d30b559f8d8b8a54911ce10a2d42296","0x686a01120827cae5230bb81d5c2a74667c8b7552","0x7bb6c127d9ca17475fcce1f8c0d61b77d327a981","0x5a1993285c17507da692653c96916bb04ca9f555","0x81ea504d5b278f651398faef51cc65a7b5791663","0x68c6d02d44e16f1c20088731ab032f849100d70f","0xc0ed9ef67d40c9226a6d3ae065e5a5113abf371b","0x167ab22d8745aacf25eca9d4b6bc49f2e75764ef","0x588a41f0f764a335798e0a3e7212a46b5029cc2c","0x47ff5a2ad7a36cfcf7867539f5851a4a573bf4e1","0xcfd111c4c0daeae6351bcdc80a993e78923845ab","0x168d3fcbb4213fc91d5c40acc302bfccd89112aa","0xcd397987bfbf91e5c64e50226a361e80e598fc44","0x94607d90f79ca3416a0ccec1aa7442ea26df1af0","0x378b4c5f2a8a0796a8d4c798ef737cf00ae8e667","0x36b66b447b7a2ffaee307785f663ec8d5a598ba1","0xc4de5cc1232f6493cc7bf7bcb12f905eb9742bd7","0xe22fb95aecfae1324b09d7a34229a006adf5431d","0x53bb862238b852f387c3f8b8008e4622548e82d1","0x44d34985826578e5ba24ec78c93be968549bb918","0x35a0d9579b1e886702375364fe9c540f97e4517b","0x7c570e184361b4f730c7fa19c9f4ee83562ae461","0x04940b1776b0bf5786dac58986112c4f722bfc63","0xf65b8e3f9da7a111c63a283618dd494324cdfec3","0x4cdfe65f19d045b7cacfd074f246e48119064a79","0xee42b08a7e2343e20f0dbc64d38f09bec78c0234","0xf7b515cf4dbcfa0b215bab0ccde0f4befc15e3c8","0x04631328f056046ab6a5ae62df2d4fe6bead5334","0xf139473f526f7a139298ef2892f1915031c4f4da","0x75e66d062a2a8be5c99d71db11d72ab3b82ba8b3","0x715134a16acb73c86e81df5542e1cf759eeb6fc7","0x998bf04788c1c631c0e02bd1eed3d945308bf0a3","0x4ae7a9e5cd73bb3d886febcd3e33a60e9f354be3","0x43cdbf2177cf8164b9ade4e89c216f860f7a4bce","0x8933f521c9b4652cf4c21a330314f85a9e4a1dfa","0x517f396b0582a46c00a20de9904717b55b4dbe46","0x145f07dad471ca72917fb1df304c804a14d5a65f","0xe128c09e570d2da542f9270444b3efc5b3315728","0xf727fc12f07f510df7f5169f5dc4ca91b9a05753","0x9ff54fdfee314c3c404409d6c1b48220a736c5b7","0x046d8c308fd8b1d287fecec297dea805a6d925f0","0xea2b2abec815a2782b25ea2b3d998e1a5eafc259","0xb9e4b63716252ec73dc43c545bfd58d18f040251","0xbb7ab09971e56aaf248dcc6c3df865af69d97372","0xba87dc891945dbb3caeeaf822de208d7ea89b298","0x7212fb4efc7c1d5eb56dea974becb3079c805c6b","0x8283dd24f8d8e7b083e2855a4c3ea1a7feba08c2","0x055cedfe14bce33f985c41d9a1934b7654611aac","0x48f2339da7d4db105ade880a2809ab8a38899f21","0x1cdb9a5b489efb99c46d317022921906dea73cde","0xdca8bedda2a15b740d3bc58bb14027c48f2afbf6","0x7c46ff058f4ec60fdfa99d97e9d8563437f95e01","0x3a99e53b8f5bb7acc583701ae37f1c7305ff79dd","0x033ecd066376afec5e6383bc9f1f15be4c62dc89","0xaf988aff99d3d0cb870812c325c588d8d8cb7de8","0x59befe124e09ce71aee9e8d1b072a08c102aa8d3","0xd7c2a4aa31e1bf08dc7ff44c9980fa8573e10c1b","0x364248b2f1f57c5402d244b2d469a35b4c0e9dab","0x8454326e091a32bb004779128c2466e4874e0144","0xc8703a9c40edb9915c815f29762827f6242278bf","0xc965bc2cf23d54ccf41a7d3af10572580e9b8078","0x7d7428c778424db4c3a0930d18ca89a46b66f042","0x6b5535346d6e8620285b8ee08d2f3651e6c75d68","0x9d7a071735dfc0557bc2e9adacbe951201f55cb6","0xe69d51fe01608edffca53d0881897fff9ae09bf6","0x8b599ddb45e82908fb663a89f09bef0c5dd15f81","0x648806dae0b6b1bf070324dff045c4af5bee8c61","0xbdacbf5cc2400303c312948d2c80ae6fa38d83e2","0x954e5e6b11f3adbd1fffa90e4b718103938a3468","0x9d61d9c3a7daf6056e61236b0f7d0cfe798a32b0","0x87d3d823f7ee1a86f0939371cc3a2ef46889b1e5","0x4bb3382c2b429cf5ff72bc26304bd4f60df5fdee","0x02eb135eba0d5b36c2cf0bb9190c36927c5eed6d","0x844aa51d309bdc4a44146192f8d4097cad51d96e","0xe26433358fdb9addcc59ba3f2df0e8490b015ff4","0x13cb03bfbe8bfb8c9923fae884a1f1b16dd2adef","0x718ff3f9957ab5586ecff2df2aa11121f71f3d67","0x1923ede7104c67d4af0a970bf8e8d566eaea2680","0x6515671641c9028e1a27117f1214dfd08132509c","0xf5d5c8125192c7968540e0718a883874a9038c2d","0x5862d8d46f6a42332dec8d8d8c27796c0f38264b","0xdf2c408f0ad496b222bac96ecadc68703b9e0b2a","0x7a8c6b6e44cc12df550096f88a5ba746c38ed9b0","0xa4d456b9dbe75b033b13b138caadbe4bd15432ae","0x51dcb8d09fac22c974cb2ef169310cc501063f46","0x9307539a322ec5afb08632bc97d8250f2d94544d","0xb2c98ff85c650d642402a16e02d87d8a433f1212","0x3e5c11703f6ee0fbe73462be65d2f44b94c925ee","0x73bf5d088c5ef4673f194dd32f38703a32362d7f","0xa49e7dc858f84f909a7e9589eaaa2826795c0d8a","0x71b50dd6ca7433b5d528f472b93f115ce00af9c8","0x9609f1398ff6b982c97fc4437cfdeba78ccde07b","0x2e3e2d13ff250308e3a9eb25739b7eed1316249b","0x28b5e0090ff9c1192e1134135deb17e0b1c1cdbf","0xfa0e7f0eb396154f71037b8c820f755c4475e9bf","0xfd0a96bba496ef72e925c3aceca8877e4600d275","0xe5ccd204616e0c13321087c74215a6d1e505f382","0x209e2f044bd334c53451900c576ab18e8cea9b16","0x131c209140db8d882ff9b67523874f9e22e9b1de","0xaaf5110db6e744ff70fb339de037b990a20bdace","0x32fdb63ac07682d012255d359299246f7d5fdfa4","0x9bea554d1ffc33f7ec498693d6f512a63769d59c","0x6d37eb2364f8bb28002bd404cf301fa8106dcb14","0x97a707e8f4cda8b0cfb13f20ecc576096f6ebed2","0x6d0083fb8c8d4c62441fee5b5be2cc21182b5fe3","0x121626bce68a417069279ca626b7305a6c1e4eeb","0xb75132cc840d71f401c29fa3f1a6c83516405cdd","0x73d5285f1221564bce82521f901cbf23fbcb876b","0x0087af23607d845d9391137360d497c1c36d3634","0xc66ed7c8034e11ebd32499389c35b3e6e9ee3e78","0x161388deb2c1147d5bed3003277d59d479d5a228","0x1417bac01f58f68ebf547a3d5b9d0f6e8a495963","0xf1934eab3f6b9cac3967cf6d1732c5e8f4c43230","0xa179d8f12f99cea60700f32f602a57c25f1f6e89","0x313bb8421607d7c100bbf3299fe505d0bd1f821c","0x6ba975118545f1bcbda6376d08417494ef9a6e16","0x24aac088f9ec4f6ecd26473204fa91669bd64072","0xd14f2f65915c8c4bb75c78209a53f8d0362997a7","0xbf8722f17e8017c96389c20b41ecb62e1f34e4bd","0xdae2bf063d208611f2e44378197a3315ea569a16","0xef4f1d5007b4ff88c1a56261fec00264af6001fb","0xf2ae5aaef14d08e0ea618abb1384127941bcbda6","0x6367eed0ce97f8ab3c237674ce50d0b746901f33","0x39043f4a27e4a030ad4bf33132d9d1535cd91ef6","0x5a875034bb718d583e490b1e6831a436e73db2e4"]},"query":"fragment PairFields on Pair {\n  id\n  txCount\n  token0 {\n    id\n    symbol\n    name\n    totalLiquidity\n    derivedETH\n    __typename\n  }\n  token1 {\n    id\n    symbol\n    name\n    totalLiquidity\n    derivedETH\n    __typename\n  }\n  reserve0\n  reserve1\n  reserveUSD\n  totalSupply\n  trackedReserveETH\n  reserveETH\n  volumeUSD\n  untrackedVolumeUSD\n  token0Price\n  token1Price\n  createdAtTimestamp\n  __typename\n}\n\nquery pairs($allPairs: [Bytes]!) {\n  pairs(where: {id_in: $allPairs}, orderBy: trackedReserveETH, orderDirection: desc) {\n    ...PairFields\n    __typename\n  }\n}\n"}`
	sushiswapPoolsDaemon = &sushiswapPools{
		fileStorage: fileStorage{
			FilePath: "./resources/sushiswap-pools.json",
			LifeSpan: DefaultLifeSpan,
		},
		graphQL:  query,
		logger:   l,
		listLock: sync.RWMutex{},
	}
	daemons[DaemonNameSushiswapPools] = sushiswapPoolsDaemon
}

func (s *sushiswapPools) GetData() IMap {
	s.listLock.RLock()
	defer s.listLock.RUnlock()
	return s.list
}

func (s *sushiswapPools) run() {
	if !s.isStorageValid() {
		s.fetch()
	} else {
		if len(s.list) == 0 || s.list == nil {
			bs, err := s.fileStorage.read()
			if err != nil {
				s.logger.Error("Sushiswap Pools Daemon: ", err)
			} else {
				var l []types.PoolInfo
				if err := json.Unmarshal(bs, &l); err != nil {
					s.logger.Error("Sushiswap Pools Daemon: ", err)
				}
				s.listLock.Lock()
				s.list = l
				s.listLock.Unlock()
			}
		}
	}
}

func (s *sushiswapPools) Run(ctx context.Context) {
	s.run()
	go func(ctx context.Context) {
		for {
			time.Sleep(DefaultLifeSpanHalf)
			select {
			case <-ctx.Done():
				return
			default:
				s.run()
			}
		}
	}(ctx)
}
