package rfc6979_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"math/big"
	"testing"

	"github.com/nspcc-dev/rfc6979"
)

type ecdsaFixture struct {
	name    string
	key     *ecdsaKey
	alg     func() hash.Hash
	message string
	r, s    string
}

type ecdsaKey struct {
	key      *ecdsa.PrivateKey
	subgroup int
}

var p224 = &ecdsaKey{
	key: &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P224(),
			X:     ecdsaLoadInt("00CF08DA5AD719E42707FA431292DEA11244D64FC51610D94B130D6C"),
			Y:     ecdsaLoadInt("EEAB6F3DEBE455E3DBF85416F7030CBD94F34F2D6F232C69F3C1385A"),
		},
		D: ecdsaLoadInt("F220266E1105BFE3083E03EC7A3A654651F45E37167E88600BF257C1"),
	},
	subgroup: 224,
}

var p256 = &ecdsaKey{
	key: &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     ecdsaLoadInt("60FED4BA255A9D31C961EB74C6356D68C049B8923B61FA6CE669622E60F29FB6"),
			Y:     ecdsaLoadInt("7903FE1008B8BC99A41AE9E95628BC64F2F1B20C2D7E9F5177A3C294D4462299"),
		},
		D: ecdsaLoadInt("C9AFA9D845BA75166B5C215767B1D6934E50C3DB36E89B127B8A622B120F6721"),
	},
	subgroup: 256,
}

var p384 = &ecdsaKey{
	key: &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P384(),
			X:     ecdsaLoadInt("EC3A4E415B4E19A4568618029F427FA5DA9A8BC4AE92E02E06AAE5286B300C64DEF8F0EA9055866064A254515480BC13"),
			Y:     ecdsaLoadInt("8015D9B72D7D57244EA8EF9AC0C621896708A59367F9DFB9F54CA84B3F1C9DB1288B231C3AE0D4FE7344FD2533264720"),
		},
		D: ecdsaLoadInt("6B9D3DAD2E1B8C1C05B19875B6659F4DE23C3B667BF297BA9AA47740787137D896D5724E4C70A825F872C9EA60D2EDF5"),
	},
	subgroup: 384,
}

var p521 = &ecdsaKey{
	key: &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P521(),
			X:     ecdsaLoadInt("1894550D0785932E00EAA23B694F213F8C3121F86DC97A04E5A7167DB4E5BCD371123D46E45DB6B5D5370A7F20FB633155D38FFA16D2BD761DCAC474B9A2F5023A4"),
			Y:     ecdsaLoadInt("0493101C962CD4D2FDDF782285E64584139C2F91B47F87FF82354D6630F746A28A0DB25741B5B34A828008B22ACC23F924FAAFBD4D33F81EA66956DFEAA2BFDFCF5"),
		},
		D: ecdsaLoadInt("0FAD06DAA62BA3B25D2FB40133DA757205DE67F5BB0018FEE8C86E1B68C7E75CAA896EB32F1F47C70855836A6D16FCC1466F6D8FBEC67DB89EC0C08B0E996B83538"),
	},
	subgroup: 521,
}

var fixtures = []ecdsaFixture{
	// ECDSA, 224 Bits (Prime Field)
	// https://tools.ietf.org/html/rfc6979#appendix-A.2.4
	{
		name:    "P224/SHA-1 #1",
		key:     p224,
		alg:     sha1.New,
		message: "sample",
		r:       "22226F9D40A96E19C4A301CE5B74B115303C0F3A4FD30FC257FB57AC",
		s:       "66D1CDD83E3AF75605DD6E2FEFF196D30AA7ED7A2EDF7AF475403D69",
	},
	{
		name:    "P224/SHA-224 #1",
		key:     p224,
		alg:     sha256.New224,
		message: "sample",
		r:       "1CDFE6662DDE1E4A1EC4CDEDF6A1F5A2FB7FBD9145C12113E6ABFD3E",
		s:       "A6694FD7718A21053F225D3F46197CA699D45006C06F871808F43EBC",
	},
	{
		name:    "P224/SHA-256 #1",
		key:     p224,
		alg:     sha256.New,
		message: "sample",
		r:       "61AA3DA010E8E8406C656BC477A7A7189895E7E840CDFE8FF42307BA",
		s:       "BC814050DAB5D23770879494F9E0A680DC1AF7161991BDE692B10101",
	},
	{
		name:    "P224/SHA-384 #1",
		key:     p224,
		alg:     sha512.New384,
		message: "sample",
		r:       "0B115E5E36F0F9EC81F1325A5952878D745E19D7BB3EABFABA77E953",
		s:       "830F34CCDFE826CCFDC81EB4129772E20E122348A2BBD889A1B1AF1D",
	},
	{
		name:    "P224/SHA-512 #1",
		key:     p224,
		alg:     sha512.New,
		message: "sample",
		r:       "074BD1D979D5F32BF958DDC61E4FB4872ADCAFEB2256497CDAC30397",
		s:       "A4CECA196C3D5A1FF31027B33185DC8EE43F288B21AB342E5D8EB084",
	},
	{
		name:    "P224/SHA-1 #2",
		key:     p224,
		alg:     sha1.New,
		message: "test",
		r:       "DEAA646EC2AF2EA8AD53ED66B2E2DDAA49A12EFD8356561451F3E21C",
		s:       "95987796F6CF2062AB8135271DE56AE55366C045F6D9593F53787BD2",
	},
	{
		name:    "P224/SHA-224 #2",
		key:     p224,
		alg:     sha256.New224,
		message: "test",
		r:       "C441CE8E261DED634E4CF84910E4C5D1D22C5CF3B732BB204DBEF019",
		s:       "902F42847A63BDC5F6046ADA114953120F99442D76510150F372A3F4",
	},
	{
		name:    "P224/SHA-256 #2",
		key:     p224,
		alg:     sha256.New,
		message: "test",
		r:       "AD04DDE87B84747A243A631EA47A1BA6D1FAA059149AD2440DE6FBA6",
		s:       "178D49B1AE90E3D8B629BE3DB5683915F4E8C99FDF6E666CF37ADCFD",
	},
	{
		name:    "P224/SHA-384 #2",
		key:     p224,
		alg:     sha512.New384,
		message: "test",
		r:       "389B92682E399B26518A95506B52C03BC9379A9DADF3391A21FB0EA4",
		s:       "414A718ED3249FF6DBC5B50C27F71F01F070944DA22AB1F78F559AAB",
	},
	{
		name:    "P224/SHA-512 #2",
		key:     p224,
		alg:     sha512.New,
		message: "test",
		r:       "049F050477C5ADD858CAC56208394B5A55BAEBBE887FDF765047C17C",
		s:       "077EB13E7005929CEFA3CD0403C7CDCC077ADF4E44F3C41B2F60ECFF",
	},
	// ECDSA, 256 Bits (Prime Field)
	// https://tools.ietf.org/html/rfc6979#appendix-A.2.5
	{
		name:    "P256/SHA-1 #1",
		key:     p256,
		alg:     sha1.New,
		message: "sample",
		r:       "61340C88C3AAEBEB4F6D667F672CA9759A6CCAA9FA8811313039EE4A35471D32",
		s:       "6D7F147DAC089441BB2E2FE8F7A3FA264B9C475098FDCF6E00D7C996E1B8B7EB",
	},
	{
		name:    "P256/SHA-224 #1",
		key:     p256,
		alg:     sha256.New224,
		message: "sample",
		r:       "53B2FFF5D1752B2C689DF257C04C40A587FABABB3F6FC2702F1343AF7CA9AA3F",
		s:       "B9AFB64FDC03DC1A131C7D2386D11E349F070AA432A4ACC918BEA988BF75C74C",
	},
	{
		name:    "P256/SHA-256 #1",
		key:     p256,
		alg:     sha256.New,
		message: "sample",
		r:       "EFD48B2AACB6A8FD1140DD9CD45E81D69D2C877B56AAF991C34D0EA84EAF3716",
		s:       "F7CB1C942D657C41D436C7A1B6E29F65F3E900DBB9AFF4064DC4AB2F843ACDA8",
	},
	{
		name:    "P256/SHA-384 #1",
		key:     p256,
		alg:     sha512.New384,
		message: "sample",
		r:       "0EAFEA039B20E9B42309FB1D89E213057CBF973DC0CFC8F129EDDDC800EF7719",
		s:       "4861F0491E6998B9455193E34E7B0D284DDD7149A74B95B9261F13ABDE940954",
	},
	{
		name:    "P256/SHA-512 #1",
		key:     p256,
		alg:     sha512.New,
		message: "sample",
		r:       "8496A60B5E9B47C825488827E0495B0E3FA109EC4568FD3F8D1097678EB97F00",
		s:       "2362AB1ADBE2B8ADF9CB9EDAB740EA6049C028114F2460F96554F61FAE3302FE",
	},
	{
		name:    "P256/SHA-1 #2",
		key:     p256,
		alg:     sha1.New,
		message: "test",
		r:       "0CBCC86FD6ABD1D99E703E1EC50069EE5C0B4BA4B9AC60E409E8EC5910D81A89",
		s:       "01B9D7B73DFAA60D5651EC4591A0136F87653E0FD780C3B1BC872FFDEAE479B1",
	},
	{
		name:    "P256/SHA-224 #2",
		key:     p256,
		alg:     sha256.New224,
		message: "test",
		r:       "C37EDB6F0AE79D47C3C27E962FA269BB4F441770357E114EE511F662EC34A692",
		s:       "C820053A05791E521FCAAD6042D40AEA1D6B1A540138558F47D0719800E18F2D",
	},
	{
		name:    "P256/SHA-256 #2",
		key:     p256,
		alg:     sha256.New,
		message: "test",
		r:       "F1ABB023518351CD71D881567B1EA663ED3EFCF6C5132B354F28D3B0B7D38367",
		s:       "019F4113742A2B14BD25926B49C649155F267E60D3814B4C0CC84250E46F0083",
	},
	{
		name:    "P256/SHA-384 #2",
		key:     p256,
		alg:     sha512.New384,
		message: "test",
		r:       "83910E8B48BB0C74244EBDF7F07A1C5413D61472BD941EF3920E623FBCCEBEB6",
		s:       "8DDBEC54CF8CD5874883841D712142A56A8D0F218F5003CB0296B6B509619F2C",
	},
	{
		name:    "P256/SHA-512 #2",
		key:     p256,
		alg:     sha512.New,
		message: "test",
		r:       "461D93F31B6540894788FD206C07CFA0CC35F46FA3C91816FFF1040AD1581A04",
		s:       "39AF9F15DE0DB8D97E72719C74820D304CE5226E32DEDAE67519E840D1194E55",
	},
	// ECDSA, 384 Bits (Prime Field)
	// https://tools.ietf.org/html/rfc6979#appendix-A.2.6
	{
		name:    "P384/SHA-1 #1",
		key:     p384,
		alg:     sha1.New,
		message: "sample",
		r:       "EC748D839243D6FBEF4FC5C4859A7DFFD7F3ABDDF72014540C16D73309834FA37B9BA002899F6FDA3A4A9386790D4EB2",
		s:       "A3BCFA947BEEF4732BF247AC17F71676CB31A847B9FF0CBC9C9ED4C1A5B3FACF26F49CA031D4857570CCB5CA4424A443",
	},
	{
		name:    "P384/SHA-224 #1",
		key:     p384,
		alg:     sha256.New224,
		message: "sample",
		r:       "42356E76B55A6D9B4631C865445DBE54E056D3B3431766D0509244793C3F9366450F76EE3DE43F5A125333A6BE060122",
		s:       "9DA0C81787064021E78DF658F2FBB0B042BF304665DB721F077A4298B095E4834C082C03D83028EFBF93A3C23940CA8D",
	},
	{
		name:    "P384/SHA-256 #1",
		key:     p384,
		alg:     sha256.New,
		message: "sample",
		r:       "21B13D1E013C7FA1392D03C5F99AF8B30C570C6F98D4EA8E354B63A21D3DAA33BDE1E888E63355D92FA2B3C36D8FB2CD",
		s:       "F3AA443FB107745BF4BD77CB3891674632068A10CA67E3D45DB2266FA7D1FEEBEFDC63ECCD1AC42EC0CB8668A4FA0AB0",
	},
	{
		name:    "P384/SHA-384 #1",
		key:     p384,
		alg:     sha512.New384,
		message: "sample",
		r:       "94EDBB92A5ECB8AAD4736E56C691916B3F88140666CE9FA73D64C4EA95AD133C81A648152E44ACF96E36DD1E80FABE46",
		s:       "99EF4AEB15F178CEA1FE40DB2603138F130E740A19624526203B6351D0A3A94FA329C145786E679E7B82C71A38628AC8",
	},
	{
		name:    "P384/SHA-512 #1",
		key:     p384,
		alg:     sha512.New,
		message: "sample",
		r:       "ED0959D5880AB2D869AE7F6C2915C6D60F96507F9CB3E047C0046861DA4A799CFE30F35CC900056D7C99CD7882433709",
		s:       "512C8CCEEE3890A84058CE1E22DBC2198F42323CE8ACA9135329F03C068E5112DC7CC3EF3446DEFCEB01A45C2667FDD5",
	},
	{
		name:    "P384/SHA-1 #2",
		key:     p384,
		alg:     sha1.New,
		message: "test",
		r:       "4BC35D3A50EF4E30576F58CD96CE6BF638025EE624004A1F7789A8B8E43D0678ACD9D29876DAF46638645F7F404B11C7",
		s:       "D5A6326C494ED3FF614703878961C0FDE7B2C278F9A65FD8C4B7186201A2991695BA1C84541327E966FA7B50F7382282",
	},
	{
		name:    "P384/SHA-224 #2",
		key:     p384,
		alg:     sha256.New224,
		message: "test",
		r:       "E8C9D0B6EA72A0E7837FEA1D14A1A9557F29FAA45D3E7EE888FC5BF954B5E62464A9A817C47FF78B8C11066B24080E72",
		s:       "07041D4A7A0379AC7232FF72E6F77B6DDB8F09B16CCE0EC3286B2BD43FA8C6141C53EA5ABEF0D8231077A04540A96B66",
	},
	{
		name:    "P384/SHA-256 #2",
		key:     p384,
		alg:     sha256.New,
		message: "test",
		r:       "6D6DEFAC9AB64DABAFE36C6BF510352A4CC27001263638E5B16D9BB51D451559F918EEDAF2293BE5B475CC8F0188636B",
		s:       "2D46F3BECBCC523D5F1A1256BF0C9B024D879BA9E838144C8BA6BAEB4B53B47D51AB373F9845C0514EEFB14024787265",
	},
	{
		name:    "P384/SHA-384 #2",
		key:     p384,
		alg:     sha512.New384,
		message: "test",
		r:       "8203B63D3C853E8D77227FB377BCF7B7B772E97892A80F36AB775D509D7A5FEB0542A7F0812998DA8F1DD3CA3CF023DB",
		s:       "DDD0760448D42D8A43AF45AF836FCE4DE8BE06B485E9B61B827C2F13173923E06A739F040649A667BF3B828246BAA5A5",
	},
	{
		name:    "P384/SHA-512 #2",
		key:     p384,
		alg:     sha512.New,
		message: "test",
		r:       "A0D5D090C9980FAF3C2CE57B7AE951D31977DD11C775D314AF55F76C676447D06FB6495CD21B4B6E340FC236584FB277",
		s:       "976984E59B4C77B0E8E4460DCA3D9F20E07B9BB1F63BEEFAF576F6B2E8B224634A2092CD3792E0159AD9CEE37659C736",
	},
	// ECDSA, 521 Bits (Prime Field)
	// https://tools.ietf.org/html/rfc6979#appendix-A.2.7
	{
		name:    "P521/SHA-1 #1",
		key:     p521,
		alg:     sha1.New,
		message: "sample",
		r:       "0343B6EC45728975EA5CBA6659BBB6062A5FF89EEA58BE3C80B619F322C87910FE092F7D45BB0F8EEE01ED3F20BABEC079D202AE677B243AB40B5431D497C55D75D",
		s:       "0E7B0E675A9B24413D448B8CC119D2BF7B2D2DF032741C096634D6D65D0DBE3D5694625FB9E8104D3B842C1B0E2D0B98BEA19341E8676AEF66AE4EBA3D5475D5D16",
	},
	{
		name:    "P521/SHA-224 #1",
		key:     p521,
		alg:     sha256.New224,
		message: "sample",
		r:       "1776331CFCDF927D666E032E00CF776187BC9FDD8E69D0DABB4109FFE1B5E2A30715F4CC923A4A5E94D2503E9ACFED92857B7F31D7152E0F8C00C15FF3D87E2ED2E",
		s:       "050CB5265417FE2320BBB5A122B8E1A32BD699089851128E360E620A30C7E17BA41A666AF126CE100E5799B153B60528D5300D08489CA9178FB610A2006C254B41F",
	},
	{
		name:    "P521/SHA-256 #1",
		key:     p521,
		alg:     sha256.New,
		message: "sample",
		r:       "1511BB4D675114FE266FC4372B87682BAECC01D3CC62CF2303C92B3526012659D16876E25C7C1E57648F23B73564D67F61C6F14D527D54972810421E7D87589E1A7",
		s:       "04A171143A83163D6DF460AAF61522695F207A58B95C0644D87E52AA1A347916E4F7A72930B1BC06DBE22CE3F58264AFD23704CBB63B29B931F7DE6C9D949A7ECFC",
	},
	{
		name:    "P521/SHA-384 #1",
		key:     p521,
		alg:     sha512.New384,
		message: "sample",
		r:       "1EA842A0E17D2DE4F92C15315C63DDF72685C18195C2BB95E572B9C5136CA4B4B576AD712A52BE9730627D16054BA40CC0B8D3FF035B12AE75168397F5D50C67451",
		s:       "1F21A3CEE066E1961025FB048BD5FE2B7924D0CD797BABE0A83B66F1E35EEAF5FDE143FA85DC394A7DEE766523393784484BDF3E00114A1C857CDE1AA203DB65D61",
	},
	{
		name:    "P521/SHA-512 #1",
		key:     p521,
		alg:     sha512.New,
		message: "sample",
		r:       "0C328FAFCBD79DD77850370C46325D987CB525569FB63C5D3BC53950E6D4C5F174E25A1EE9017B5D450606ADD152B534931D7D4E8455CC91F9B15BF05EC36E377FA",
		s:       "0617CCE7CF5064806C467F678D3B4080D6F1CC50AF26CA209417308281B68AF282623EAA63E5B5C0723D8B8C37FF0777B1A20F8CCB1DCCC43997F1EE0E44DA4A67A",
	},
	{
		name:    "P521/SHA-1 #2",
		key:     p521,
		alg:     sha1.New,
		message: "test",
		r:       "13BAD9F29ABE20DE37EBEB823C252CA0F63361284015A3BF430A46AAA80B87B0693F0694BD88AFE4E661FC33B094CD3B7963BED5A727ED8BD6A3A202ABE009D0367",
		s:       "1E9BB81FF7944CA409AD138DBBEE228E1AFCC0C890FC78EC8604639CB0DBDC90F717A99EAD9D272855D00162EE9527567DD6A92CBD629805C0445282BBC916797FF",
	},
	{
		name:    "P521/SHA-224 #2",
		key:     p521,
		alg:     sha256.New224,
		message: "test",
		r:       "1C7ED902E123E6815546065A2C4AF977B22AA8EADDB68B2C1110E7EA44D42086BFE4A34B67DDC0E17E96536E358219B23A706C6A6E16BA77B65E1C595D43CAE17FB",
		s:       "177336676304FCB343CE028B38E7B4FBA76C1C1B277DA18CAD2A8478B2A9A9F5BEC0F3BA04F35DB3E4263569EC6AADE8C92746E4C82F8299AE1B8F1739F8FD519A4",
	},
	{
		name:    "P521/SHA-256 #2",
		key:     p521,
		alg:     sha256.New,
		message: "test",
		r:       "00E871C4A14F993C6C7369501900C4BC1E9C7B0B4BA44E04868B30B41D8071042EB28C4C250411D0CE08CD197E4188EA4876F279F90B3D8D74A3C76E6F1E4656AA8",
		s:       "0CD52DBAA33B063C3A6CD8058A1FB0A46A4754B034FCC644766CA14DA8CA5CA9FDE00E88C1AD60CCBA759025299079D7A427EC3CC5B619BFBC828E7769BCD694E86",
	},
	{
		name:    "P521/SHA-384 #2",
		key:     p521,
		alg:     sha512.New384,
		message: "test",
		r:       "14BEE21A18B6D8B3C93FAB08D43E739707953244FDBE924FA926D76669E7AC8C89DF62ED8975C2D8397A65A49DCC09F6B0AC62272741924D479354D74FF6075578C",
		s:       "133330865C067A0EAF72362A65E2D7BC4E461E8C8995C3B6226A21BD1AA78F0ED94FE536A0DCA35534F0CD1510C41525D163FE9D74D134881E35141ED5E8E95B979",
	},
	{
		name:    "P521/SHA-512 #2",
		key:     p521,
		alg:     sha512.New,
		message: "test",
		r:       "13E99020ABF5CEE7525D16B69B229652AB6BDF2AFFCAEF38773B4B7D08725F10CDB93482FDCC54EDCEE91ECA4166B2A7C6265EF0CE2BD7051B7CEF945BABD47EE6D",
		s:       "1FBD0013C674AA79CB39849527916CE301C66EA7CE8B80682786AD60F98F7E78A19CA69EFF5C57400E3B3A0AD66CE0978214D13BAF4E9AC60752F7B155E2DE4DCE3",
	},
}

func TestECDSA(t *testing.T) {
	for _, f := range fixtures {
		testEcsaFixture(&f, t)
	}
}

func ecdsaLoadInt(s string) (n *big.Int) {
	n, _ = new(big.Int).SetString(s, 16)
	return
}

func testEcsaFixture(f *ecdsaFixture, t *testing.T) {
	t.Logf("Testing %s", f.name)

	h := f.alg()
	h.Write([]byte(f.message))
	digest := h.Sum(nil)

	g := f.key.subgroup / 8
	if len(digest) > g {
		digest = digest[0:g]
	}

	r, s := rfc6979.SignECDSA(f.key.key, digest, f.alg)
	expectedR := ecdsaLoadInt(f.r)
	expectedS := ecdsaLoadInt(f.s)

	if r.Cmp(expectedR) != 0 {
		t.Errorf("%s: Expected R of %X, got %X", f.name, expectedR, r)
	}

	if s.Cmp(expectedS) != 0 {
		t.Errorf("%s: Expected S of %X, got %X", f.name, expectedS, s)
	}
}
