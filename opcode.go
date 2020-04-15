package main

const (
	OP_0                   = 0x00
	OP_FALSE               = 0x00
	OP_DATA_1              = 0x01
	OP_DATA_2              = 0x02
	OP_DATA_3              = 0x03
	OP_DATA_4              = 0x04
	OP_DATA_5              = 0x05
	OP_DATA_6              = 0x06
	OP_DATA_7              = 0x07
	OP_DATA_8              = 0x08
	OP_DATA_9              = 0x09
	OP_DATA_10             = 0x0a
	OP_DATA_11             = 0x0b
	OP_DATA_12             = 0x0c
	OP_DATA_13             = 0x0d
	OP_DATA_14             = 0x0e
	OP_DATA_15             = 0x0f
	OP_DATA_16             = 0x10
	OP_DATA_17             = 0x11
	OP_DATA_18             = 0x12
	OP_DATA_19             = 0x13
	OP_DATA_20             = 0x14
	OP_DATA_21             = 0x15
	OP_DATA_22             = 0x16
	OP_DATA_23             = 0x17
	OP_DATA_24             = 0x18
	OP_DATA_25             = 0x19
	OP_DATA_26             = 0x1a
	OP_DATA_27             = 0x1b
	OP_DATA_28             = 0x1c
	OP_DATA_29             = 0x1d
	OP_DATA_30             = 0x1e
	OP_DATA_31             = 0x1f
	OP_DATA_32             = 0x20
	OP_DATA_33             = 0x21
	OP_DATA_34             = 0x22
	OP_DATA_35             = 0x23
	OP_DATA_36             = 0x24
	OP_DATA_37             = 0x25
	OP_DATA_38             = 0x26
	OP_DATA_39             = 0x27
	OP_DATA_40             = 0x28
	OP_DATA_41             = 0x29
	OP_DATA_42             = 0x2a
	OP_DATA_43             = 0x2b
	OP_DATA_44             = 0x2c
	OP_DATA_45             = 0x2d
	OP_DATA_46             = 0x2e
	OP_DATA_47             = 0x2f
	OP_DATA_48             = 0x30
	OP_DATA_49             = 0x31
	OP_DATA_50             = 0x32
	OP_DATA_51             = 0x33
	OP_DATA_52             = 0x34
	OP_DATA_53             = 0x35
	OP_DATA_54             = 0x36
	OP_DATA_55             = 0x37
	OP_DATA_56             = 0x38
	OP_DATA_57             = 0x39
	OP_DATA_58             = 0x3a
	OP_DATA_59             = 0x3b
	OP_DATA_60             = 0x3c
	OP_DATA_61             = 0x3d
	OP_DATA_62             = 0x3e
	OP_DATA_63             = 0x3f
	OP_DATA_64             = 0x40
	OP_DATA_65             = 0x41
	OP_DATA_66             = 0x42
	OP_DATA_67             = 0x43
	OP_DATA_68             = 0x44
	OP_DATA_69             = 0x45
	OP_DATA_70             = 0x46
	OP_DATA_71             = 0x47
	OP_DATA_72             = 0x48
	OP_DATA_73             = 0x49
	OP_DATA_74             = 0x4a
	OP_DATA_75             = 0x4b
	OP_PUSHDATA1           = 0x4c
	OP_PUSHDATA2           = 0x4d
	OP_PUSHDATA4           = 0x4e
	OP_1NEGATE             = 0x4f
	OP_RESERVED            = 0x50
	OP_1                   = 0x51
	OP_TRUE                = 0x51
	OP_2                   = 0x52
	OP_3                   = 0x53
	OP_4                   = 0x54
	OP_5                   = 0x55
	OP_6                   = 0x56
	OP_7                   = 0x57
	OP_8                   = 0x58
	OP_9                   = 0x59
	OP_10                  = 0x5a
	OP_11                  = 0x5b
	OP_12                  = 0x5c
	OP_13                  = 0x5d
	OP_14                  = 0x5e
	OP_15                  = 0x5f
	OP_16                  = 0x60
	OP_NOP                 = 0x61
	OP_VER                 = 0x62
	OP_IF                  = 0x63
	OP_NOTIF               = 0x64
	OP_VERIF               = 0x65
	OP_VERNOTIF            = 0x66
	OP_ELSE                = 0x67
	OP_ENDIF               = 0x68
	OP_VERIFY              = 0x69
	OP_RETURN              = 0x6a
	OP_TOALTSTACK          = 0x6b
	OP_FROMALTSTACK        = 0x6c
	OP_2DROP               = 0x6d
	OP_2DUP                = 0x6e
	OP_3DUP                = 0x6f
	OP_2OVER               = 0x70
	OP_2ROT                = 0x71
	OP_2SWAP               = 0x72
	OP_IFDUP               = 0x73
	OP_DEPTH               = 0x74
	OP_DROP                = 0x75
	OP_DUP                 = 0x76
	OP_NIP                 = 0x77
	OP_OVER                = 0x78
	OP_PICK                = 0x79
	OP_ROLL                = 0x7a
	OP_ROT                 = 0x7b
	OP_SWAP                = 0x7c
	OP_TUCK                = 0x7d
	OP_CAT                 = 0x7e
	OP_SUBSTR              = 0x7f
	OP_LEFT                = 0x80
	OP_RIGHT               = 0x81
	OP_SIZE                = 0x82
	OP_INVERT              = 0x83
	OP_AND                 = 0x84
	OP_OR                  = 0x85
	OP_XOR                 = 0x86
	OP_EQUAL               = 0x87
	OP_EQUALVERIFY         = 0x88
	OP_RESERVED1           = 0x89
	OP_RESERVED2           = 0x8a
	OP_1ADD                = 0x8b
	OP_1SUB                = 0x8c
	OP_2MUL                = 0x8d
	OP_2DIV                = 0x8e
	OP_NEGATE              = 0x8f
	OP_ABS                 = 0x90
	OP_NOT                 = 0x91
	OP_0NOTEQUAL           = 0x92
	OP_ADD                 = 0x93
	OP_SUB                 = 0x94
	OP_MUL                 = 0x95
	OP_DIV                 = 0x96
	OP_MOD                 = 0x97
	OP_LSHIFT              = 0x98
	OP_RSHIFT              = 0x99
	OP_BOOLAND             = 0x9a
	OP_BOOLOR              = 0x9b
	OP_NUMEQUAL            = 0x9c
	OP_NUMEQUALVERIFY      = 0x9d
	OP_NUMNOTEQUAL         = 0x9e
	OP_LESSTHAN            = 0x9f
	OP_GREATERTHAN         = 0xa0
	OP_LESSTHANOREQUAL     = 0xa1
	OP_GREATERTHANOREQUAL  = 0xa2
	OP_MIN                 = 0xa3
	OP_MAX                 = 0xa4
	OP_WITHIN              = 0xa5
	OP_RIPEMD160           = 0xa6
	OP_SHA1                = 0xa7
	OP_SHA256              = 0xa8
	OP_HASH160             = 0xa9
	OP_HASH256             = 0xaa
	OP_CODESEPARATOR       = 0xab
	OP_CHECKSIG            = 0xac
	OP_CHECKSIGVERIFY      = 0xad
	OP_CHECKMULTISIG       = 0xae
	OP_CHECKMULTISIGVERIFY = 0xaf
	OP_NOP1                = 0xb0
	OP_NOP2                = 0xb1
	OP_CHECKLOCKTIMEVERIFY = 0xb1
	OP_NOP3                = 0xb2
	OP_CHECKSEQUENCEVERIFY = 0xb2
	OP_NOP4                = 0xb3
	OP_NOP5                = 0xb4
	OP_NOP6                = 0xb5
	OP_NOP7                = 0xb6
	OP_NOP8                = 0xb7
	OP_NOP9                = 0xb8
	OP_NOP10               = 0xb9
	OP_UNKNOWN186          = 0xba
	OP_UNKNOWN187          = 0xbb
	OP_UNKNOWN188          = 0xbc
	OP_UNKNOWN189          = 0xbd
	OP_UNKNOWN190          = 0xbe
	OP_UNKNOWN191          = 0xbf
	OP_UNKNOWN192          = 0xc0
	OP_UNKNOWN193          = 0xc1
	OP_UNKNOWN194          = 0xc2
	OP_UNKNOWN195          = 0xc3
	OP_UNKNOWN196          = 0xc4
	OP_UNKNOWN197          = 0xc5
	OP_UNKNOWN198          = 0xc6
	OP_UNKNOWN199          = 0xc7
	OP_UNKNOWN200          = 0xc8
	OP_UNKNOWN201          = 0xc9
	OP_UNKNOWN202          = 0xca
	OP_UNKNOWN203          = 0xcb
	OP_UNKNOWN204          = 0xcc
	OP_UNKNOWN205          = 0xcd
	OP_UNKNOWN206          = 0xce
	OP_UNKNOWN207          = 0xcf
	OP_UNKNOWN208          = 0xd0
	OP_UNKNOWN209          = 0xd1
	OP_UNKNOWN210          = 0xd2
	OP_UNKNOWN211          = 0xd3
	OP_UNKNOWN212          = 0xd4
	OP_UNKNOWN213          = 0xd5
	OP_UNKNOWN214          = 0xd6
	OP_UNKNOWN215          = 0xd7
	OP_UNKNOWN216          = 0xd8
	OP_UNKNOWN217          = 0xd9
	OP_UNKNOWN218          = 0xda
	OP_UNKNOWN219          = 0xdb
	OP_UNKNOWN220          = 0xdc
	OP_UNKNOWN221          = 0xdd
	OP_UNKNOWN222          = 0xde
	OP_UNKNOWN223          = 0xdf
	OP_UNKNOWN224          = 0xe0
	OP_UNKNOWN225          = 0xe1
	OP_UNKNOWN226          = 0xe2
	OP_UNKNOWN227          = 0xe3
	OP_UNKNOWN228          = 0xe4
	OP_UNKNOWN229          = 0xe5
	OP_UNKNOWN230          = 0xe6
	OP_UNKNOWN231          = 0xe7
	OP_UNKNOWN232          = 0xe8
	OP_UNKNOWN233          = 0xe9
	OP_UNKNOWN234          = 0xea
	OP_UNKNOWN235          = 0xeb
	OP_UNKNOWN236          = 0xec
	OP_UNKNOWN237          = 0xed
	OP_UNKNOWN238          = 0xee
	OP_UNKNOWN239          = 0xef
	OP_UNKNOWN240          = 0xf0
	OP_UNKNOWN241          = 0xf1
	OP_UNKNOWN242          = 0xf2
	OP_UNKNOWN243          = 0xf3
	OP_UNKNOWN244          = 0xf4
	OP_UNKNOWN245          = 0xf5
	OP_UNKNOWN246          = 0xf6
	OP_UNKNOWN247          = 0xf7
	OP_UNKNOWN248          = 0xf8
	OP_UNKNOWN249          = 0xf9
	OP_SMALLINTEGER        = 0xfa
	OP_PUBKEYS             = 0xfb
	OP_UNKNOWN252          = 0xfc
	OP_PUBKEYHASH          = 0xfd
	OP_PUBKEY              = 0xfe
	OP_INVALIDOPCODE       = 0xff
)

var opcodeArray = [256]opcode{
	OP_FALSE:     {OP_FALSE, "OP_0", 1},
	OP_DATA_1:    {OP_DATA_1, "OP_DATA_1", 2},
	OP_DATA_2:    {OP_DATA_2, "OP_DATA_2", 3},
	OP_DATA_3:    {OP_DATA_3, "OP_DATA_3", 4},
	OP_DATA_4:    {OP_DATA_4, "OP_DATA_4", 5},
	OP_DATA_5:    {OP_DATA_5, "OP_DATA_5", 6},
	OP_DATA_6:    {OP_DATA_6, "OP_DATA_6", 7},
	OP_DATA_7:    {OP_DATA_7, "OP_DATA_7", 8,},
	OP_DATA_8:    {OP_DATA_8, "OP_DATA_8", 9},
	OP_DATA_9:    {OP_DATA_9, "OP_DATA_9", 10},
	OP_DATA_10:   {OP_DATA_10, "OP_DATA_10", 11},
	OP_DATA_11:   {OP_DATA_11, "OP_DATA_11", 12},
	OP_DATA_12:   {OP_DATA_12, "OP_DATA_12", 13},
	OP_DATA_13:   {OP_DATA_13, "OP_DATA_13", 14},
	OP_DATA_14:   {OP_DATA_14, "OP_DATA_14", 15},
	OP_DATA_15:   {OP_DATA_15, "OP_DATA_15", 16},
	OP_DATA_16:   {OP_DATA_16, "OP_DATA_16", 17},
	OP_DATA_17:   {OP_DATA_17, "OP_DATA_17", 18},
	OP_DATA_18:   {OP_DATA_18, "OP_DATA_18", 19},
	OP_DATA_19:   {OP_DATA_19, "OP_DATA_19", 20},
	OP_DATA_20:   {OP_DATA_20, "OP_DATA_20", 21},
	OP_DATA_21:   {OP_DATA_21, "OP_DATA_21", 22},
	OP_DATA_22:   {OP_DATA_22, "OP_DATA_22", 23},
	OP_DATA_23:   {OP_DATA_23, "OP_DATA_23", 24},
	OP_DATA_24:   {OP_DATA_24, "OP_DATA_24", 25},
	OP_DATA_25:   {OP_DATA_25, "OP_DATA_25", 26},
	OP_DATA_26:   {OP_DATA_26, "OP_DATA_26", 27},
	OP_DATA_27:   {OP_DATA_27, "OP_DATA_27", 28},
	OP_DATA_28:   {OP_DATA_28, "OP_DATA_28", 29},
	OP_DATA_29:   {OP_DATA_29, "OP_DATA_29", 30},
	OP_DATA_30:   {OP_DATA_30, "OP_DATA_30", 31},
	OP_DATA_31:   {OP_DATA_31, "OP_DATA_31", 32},
	OP_DATA_32:   {OP_DATA_32, "OP_DATA_32", 33},
	OP_DATA_33:   {OP_DATA_33, "OP_DATA_33", 34},
	OP_DATA_34:   {OP_DATA_34, "OP_DATA_34", 35},
	OP_DATA_35:   {OP_DATA_35, "OP_DATA_35", 36},
	OP_DATA_36:   {OP_DATA_36, "OP_DATA_36", 37},
	OP_DATA_37:   {OP_DATA_37, "OP_DATA_37", 38},
	OP_DATA_38:   {OP_DATA_38, "OP_DATA_38", 39},
	OP_DATA_39:   {OP_DATA_39, "OP_DATA_39", 40},
	OP_DATA_40:   {OP_DATA_40, "OP_DATA_40", 41},
	OP_DATA_41:   {OP_DATA_41, "OP_DATA_41", 42},
	OP_DATA_42:   {OP_DATA_42, "OP_DATA_42", 43},
	OP_DATA_43:   {OP_DATA_43, "OP_DATA_43", 44},
	OP_DATA_44:   {OP_DATA_44, "OP_DATA_44", 45},
	OP_DATA_45:   {OP_DATA_45, "OP_DATA_45", 46},
	OP_DATA_46:   {OP_DATA_46, "OP_DATA_46", 47},
	OP_DATA_47:   {OP_DATA_47, "OP_DATA_47", 48},
	OP_DATA_48:   {OP_DATA_48, "OP_DATA_48", 49},
	OP_DATA_49:   {OP_DATA_49, "OP_DATA_49", 50},
	OP_DATA_50:   {OP_DATA_50, "OP_DATA_50", 51},
	OP_DATA_51:   {OP_DATA_51, "OP_DATA_51", 52},
	OP_DATA_52:   {OP_DATA_52, "OP_DATA_52", 53},
	OP_DATA_53:   {OP_DATA_53, "OP_DATA_53", 54},
	OP_DATA_54:   {OP_DATA_54, "OP_DATA_54", 55},
	OP_DATA_55:   {OP_DATA_55, "OP_DATA_55", 56},
	OP_DATA_56:   {OP_DATA_56, "OP_DATA_56", 57},
	OP_DATA_57:   {OP_DATA_57, "OP_DATA_57", 58},
	OP_DATA_58:   {OP_DATA_58, "OP_DATA_58", 59},
	OP_DATA_59:   {OP_DATA_59, "OP_DATA_59", 60},
	OP_DATA_60:   {OP_DATA_60, "OP_DATA_60", 61},
	OP_DATA_61:   {OP_DATA_61, "OP_DATA_61", 62},
	OP_DATA_62:   {OP_DATA_62, "OP_DATA_62", 63},
	OP_DATA_63:   {OP_DATA_63, "OP_DATA_63", 64},
	OP_DATA_64:   {OP_DATA_64, "OP_DATA_64", 65},
	OP_DATA_65:   {OP_DATA_65, "OP_DATA_65", 66},
	OP_DATA_66:   {OP_DATA_66, "OP_DATA_66", 67},
	OP_DATA_67:   {OP_DATA_67, "OP_DATA_67", 68},
	OP_DATA_68:   {OP_DATA_68, "OP_DATA_68", 69},
	OP_DATA_69:   {OP_DATA_69, "OP_DATA_69", 70},
	OP_DATA_70:   {OP_DATA_70, "OP_DATA_70", 71},
	OP_DATA_71:   {OP_DATA_71, "OP_DATA_71", 72},
	OP_DATA_72:   {OP_DATA_72, "OP_DATA_72", 73},
	OP_DATA_73:   {OP_DATA_73, "OP_DATA_73", 74},
	OP_DATA_74:   {OP_DATA_74, "OP_DATA_74", 75},
	OP_DATA_75:   {OP_DATA_75, "OP_DATA_75", 76},
	OP_PUSHDATA1: {OP_PUSHDATA1, "OP_PUSHDATA1", -1},
	OP_PUSHDATA2: {OP_PUSHDATA2, "OP_PUSHDATA2", -2},
	OP_PUSHDATA4: {OP_PUSHDATA4, "OP_PUSHDATA4", -4},
	OP_1NEGATE:   {OP_1NEGATE, "OP_1NEGATE", 1},
	OP_RESERVED:  {OP_RESERVED, "OP_RESERVED", 1},
	OP_TRUE:      {OP_TRUE, "OP_1", 1},
	OP_2:         {OP_2, "OP_2", 1},
	OP_3:         {OP_3, "OP_3", 1},
	OP_4:         {OP_4, "OP_4", 1},
	OP_5:         {OP_5, "OP_5", 1},
	OP_6:         {OP_6, "OP_6", 1},
	OP_7:         {OP_7, "OP_7", 1},
	OP_8:         {OP_8, "OP_8", 1},
	OP_9:         {OP_9, "OP_9", 1},
	OP_10:        {OP_10, "OP_10", 1},
	OP_11:        {OP_11, "OP_11", 1},
	OP_12:        {OP_12, "OP_12", 1},
	OP_13:        {OP_13, "OP_13", 1},
	OP_14:        {OP_14, "OP_14", 1},
	OP_15:        {OP_15, "OP_15", 1},
	OP_16:        {OP_16, "OP_16", 1},
	OP_NOP:                 {OP_NOP, "OP_NOP", 1 },
	OP_VER:                 {OP_VER, "OP_VER", 1 },
	OP_IF:                  {OP_IF, "OP_IF", 1, },
	OP_NOTIF:               {OP_NOTIF, "OP_NOTIF", 1 },
	OP_VERIF:               {OP_VERIF, "OP_VERIF", 1 },
	OP_VERNOTIF:            {OP_VERNOTIF, "OP_VERNOTIF", 1 },
	OP_ELSE:                {OP_ELSE, "OP_ELSE", 1,},
	OP_ENDIF:               {OP_ENDIF, "OP_ENDIF", 1},
	OP_VERIFY:              {OP_VERIFY, "OP_VERIFY", 1},
	OP_RETURN:              {OP_RETURN, "OP_RETURN", 1},
	OP_CHECKLOCKTIMEVERIFY: {OP_CHECKLOCKTIMEVERIFY, "OP_CHECKLOCKTIMEVERIFY", 1},
	OP_CHECKSEQUENCEVERIFY: {OP_CHECKSEQUENCEVERIFY, "OP_CHECKSEQUENCEVERIFY", 1},
	OP_TOALTSTACK:   {OP_TOALTSTACK, "OP_TOALTSTACK", 1},
	OP_FROMALTSTACK: {OP_FROMALTSTACK, "OP_FROMALTSTACK", 1},
	OP_2DROP:        {OP_2DROP, "OP_2DROP", 1},
	OP_2DUP:         {OP_2DUP, "OP_2DUP", 1},
	OP_3DUP:         {OP_3DUP, "OP_3DUP", 1},
	OP_2OVER:        {OP_2OVER, "OP_2OVER", 1},
	OP_2ROT:         {OP_2ROT, "OP_2ROT", 1},
	OP_2SWAP:        {OP_2SWAP, "OP_2SWAP", 1},
	OP_IFDUP:        {OP_IFDUP, "OP_IFDUP", 1},
	OP_DEPTH:        {OP_DEPTH, "OP_DEPTH", 1},
	OP_DROP:         {OP_DROP, "OP_DROP", 1},
	OP_DUP:          {OP_DUP, "OP_DUP", 1},
	OP_NIP:          {OP_NIP, "OP_NIP", 1},
	OP_OVER:         {OP_OVER, "OP_OVER", 1},
	OP_PICK:         {OP_PICK, "OP_PICK", 1},
	OP_ROLL:         {OP_ROLL, "OP_ROLL", 1},
	OP_ROT:          {OP_ROT, "OP_ROT", 1},
	OP_SWAP:         {OP_SWAP, "OP_SWAP", 1},
	OP_TUCK:         {OP_TUCK, "OP_TUCK", 1},
	OP_CAT:    {OP_CAT, "OP_CAT", 1},
	OP_SUBSTR: {OP_SUBSTR, "OP_SUBSTR", 1},
	OP_LEFT:   {OP_LEFT, "OP_LEFT", 1},
	OP_RIGHT:  {OP_RIGHT, "OP_RIGHT", 1},
	OP_SIZE:   {OP_SIZE, "OP_SIZE", 1},
	OP_INVERT:      {OP_INVERT, "OP_INVERT", 1},
	OP_AND:         {OP_AND, "OP_AND", 1},
	OP_OR:          {OP_OR, "OP_OR", 1},
	OP_XOR:         {OP_XOR, "OP_XOR", 1},
	OP_EQUAL:       {OP_EQUAL, "OP_EQUAL", 1},
	OP_EQUALVERIFY: {OP_EQUALVERIFY, "OP_EQUALVERIFY", 1},
	OP_RESERVED1:   {OP_RESERVED1, "OP_RESERVED1", 1},
	OP_RESERVED2:   {OP_RESERVED2, "OP_RESERVED2", 1},
	OP_1SUB:               {OP_1SUB, "OP_1SUB", 1},
	OP_2MUL:               {OP_2MUL, "OP_2MUL", 1},
	OP_2DIV:               {OP_2DIV, "OP_2DIV", 1},
	OP_NEGATE:             {OP_NEGATE, "OP_NEGATE", 1},
	OP_ABS:                {OP_ABS, "OP_ABS", 1},
	OP_NOT:                {OP_NOT, "OP_NOT", 1},
	OP_0NOTEQUAL:          {OP_0NOTEQUAL, "OP_0NOTEQUAL", 1},
	OP_ADD:                {OP_ADD, "OP_ADD", 1},
	OP_SUB:                {OP_SUB, "OP_SUB", 1},
	OP_MUL:                {OP_MUL, "OP_MUL", 1},
	OP_DIV:                {OP_DIV, "OP_DIV", 1},
	OP_MOD:                {OP_MOD, "OP_MOD", 1},
	OP_LSHIFT:             {OP_LSHIFT, "OP_LSHIFT", 1},
	OP_RSHIFT:             {OP_RSHIFT, "OP_RSHIFT", 1},
	OP_BOOLAND:            {OP_BOOLAND, "OP_BOOLAND", 1},
	OP_BOOLOR:             {OP_BOOLOR, "OP_BOOLOR", 1},
	OP_NUMEQUAL:           {OP_NUMEQUAL, "OP_NUMEQUAL", 1},
	OP_NUMEQUALVERIFY:     {OP_NUMEQUALVERIFY, "OP_NUMEQUALVERIFY", 1},
	OP_NUMNOTEQUAL:        {OP_NUMNOTEQUAL, "OP_NUMNOTEQUAL", 1},
	OP_LESSTHAN:           {OP_LESSTHAN, "OP_LESSTHAN", 1},
	OP_GREATERTHAN:        {OP_GREATERTHAN, "OP_GREATERTHAN", 1},
	OP_LESSTHANOREQUAL:    {OP_LESSTHANOREQUAL, "OP_LESSTHANOREQUAL", 1},
	OP_GREATERTHANOREQUAL: {OP_GREATERTHANOREQUAL, "OP_GREATERTHANOREQUAL", 1},
	OP_MIN:                {OP_MIN, "OP_MIN", 1},
	OP_MAX:                {OP_MAX, "OP_MAX", 1},
	OP_WITHIN:             {OP_WITHIN, "OP_WITHIN", 1},
	OP_RIPEMD160:           {OP_RIPEMD160, "OP_RIPEMD160", 1},
	OP_SHA1:                {OP_SHA1, "OP_SHA1", 1},
	OP_SHA256:              {OP_SHA256, "OP_SHA256", 1},
	OP_HASH160:             {OP_HASH160, "OP_HASH160", 1},
	OP_HASH256:             {OP_HASH256, "OP_HASH256", 1},
	OP_CODESEPARATOR:       {OP_CODESEPARATOR, "OP_CODESEPARATOR", 1},
	OP_CHECKSIG:            {OP_CHECKSIG, "OP_CHECKSIG", 1},
	OP_CHECKSIGVERIFY:      {OP_CHECKSIGVERIFY, "OP_CHECKSIGVERIFY", 1},
	OP_CHECKMULTISIG:       {OP_CHECKMULTISIG, "OP_CHECKMULTISIG", 1},
	OP_CHECKMULTISIGVERIFY: {OP_CHECKMULTISIGVERIFY, "OP_CHECKMULTISIGVERIFY", 1},
	OP_NOP1:  {OP_NOP1, "OP_NOP1", 1},
	OP_NOP4:  {OP_NOP4, "OP_NOP4", 1},
	OP_NOP5:  {OP_NOP5, "OP_NOP5", 1},
	OP_NOP6:  {OP_NOP6, "OP_NOP6", 1},
	OP_NOP7:  {OP_NOP7, "OP_NOP7", 1},
	OP_NOP8:  {OP_NOP8, "OP_NOP8", 1},
	OP_NOP9:  {OP_NOP9, "OP_NOP9", 1},
	OP_NOP10: {OP_NOP10, "OP_NOP10", 1},
	OP_UNKNOWN186: {OP_UNKNOWN186, "OP_UNKNOWN186", 1},
	OP_UNKNOWN187: {OP_UNKNOWN187, "OP_UNKNOWN187", 1},
	OP_UNKNOWN188: {OP_UNKNOWN188, "OP_UNKNOWN188", 1},
	OP_UNKNOWN189: {OP_UNKNOWN189, "OP_UNKNOWN189", 1},
	OP_UNKNOWN190: {OP_UNKNOWN190, "OP_UNKNOWN190", 1},
	OP_UNKNOWN191: {OP_UNKNOWN191, "OP_UNKNOWN191", 1},
	OP_UNKNOWN192: {OP_UNKNOWN192, "OP_UNKNOWN192", 1},
	OP_UNKNOWN193: {OP_UNKNOWN193, "OP_UNKNOWN193", 1},
	OP_UNKNOWN194: {OP_UNKNOWN194, "OP_UNKNOWN194", 1},
	OP_UNKNOWN195: {OP_UNKNOWN195, "OP_UNKNOWN195", 1},
	OP_UNKNOWN196: {OP_UNKNOWN196, "OP_UNKNOWN196", 1},
	OP_UNKNOWN197: {OP_UNKNOWN197, "OP_UNKNOWN197", 1},
	OP_UNKNOWN198: {OP_UNKNOWN198, "OP_UNKNOWN198", 1},
	OP_UNKNOWN199: {OP_UNKNOWN199, "OP_UNKNOWN199", 1},
	OP_UNKNOWN200: {OP_UNKNOWN200, "OP_UNKNOWN200", 1},
	OP_UNKNOWN201: {OP_UNKNOWN201, "OP_UNKNOWN201", 1},
	OP_UNKNOWN202: {OP_UNKNOWN202, "OP_UNKNOWN202", 1},
	OP_UNKNOWN203: {OP_UNKNOWN203, "OP_UNKNOWN203", 1},
	OP_UNKNOWN204: {OP_UNKNOWN204, "OP_UNKNOWN204", 1},
	OP_UNKNOWN205: {OP_UNKNOWN205, "OP_UNKNOWN205", 1},
	OP_UNKNOWN206: {OP_UNKNOWN206, "OP_UNKNOWN206", 1},
	OP_UNKNOWN207: {OP_UNKNOWN207, "OP_UNKNOWN207", 1},
	OP_UNKNOWN208: {OP_UNKNOWN208, "OP_UNKNOWN208", 1},
	OP_UNKNOWN209: {OP_UNKNOWN209, "OP_UNKNOWN209", 1},
	OP_UNKNOWN210: {OP_UNKNOWN210, "OP_UNKNOWN210", 1},
	OP_UNKNOWN211: {OP_UNKNOWN211, "OP_UNKNOWN211", 1},
	OP_UNKNOWN212: {OP_UNKNOWN212, "OP_UNKNOWN212", 1},
	OP_UNKNOWN213: {OP_UNKNOWN213, "OP_UNKNOWN213", 1},
	OP_UNKNOWN214: {OP_UNKNOWN214, "OP_UNKNOWN214", 1},
	OP_UNKNOWN215: {OP_UNKNOWN215, "OP_UNKNOWN215", 1},
	OP_UNKNOWN216: {OP_UNKNOWN216, "OP_UNKNOWN216", 1},
	OP_UNKNOWN217: {OP_UNKNOWN217, "OP_UNKNOWN217", 1},
	OP_UNKNOWN218: {OP_UNKNOWN218, "OP_UNKNOWN218", 1},
	OP_UNKNOWN219: {OP_UNKNOWN219, "OP_UNKNOWN219", 1},
	OP_UNKNOWN220: {OP_UNKNOWN220, "OP_UNKNOWN220", 1},
	OP_UNKNOWN221: {OP_UNKNOWN221, "OP_UNKNOWN221", 1},
	OP_UNKNOWN222: {OP_UNKNOWN222, "OP_UNKNOWN222", 1},
	OP_UNKNOWN223: {OP_UNKNOWN223, "OP_UNKNOWN223", 1},
	OP_UNKNOWN224: {OP_UNKNOWN224, "OP_UNKNOWN224", 1},
	OP_UNKNOWN225: {OP_UNKNOWN225, "OP_UNKNOWN225", 1},
	OP_UNKNOWN226: {OP_UNKNOWN226, "OP_UNKNOWN226", 1},
	OP_UNKNOWN227: {OP_UNKNOWN227, "OP_UNKNOWN227", 1},
	OP_UNKNOWN228: {OP_UNKNOWN228, "OP_UNKNOWN228", 1},
	OP_UNKNOWN229: {OP_UNKNOWN229, "OP_UNKNOWN229", 1},
	OP_UNKNOWN230: {OP_UNKNOWN230, "OP_UNKNOWN230", 1},
	OP_UNKNOWN231: {OP_UNKNOWN231, "OP_UNKNOWN231", 1},
	OP_UNKNOWN232: {OP_UNKNOWN232, "OP_UNKNOWN232", 1},
	OP_UNKNOWN233: {OP_UNKNOWN233, "OP_UNKNOWN233", 1},
	OP_UNKNOWN234: {OP_UNKNOWN234, "OP_UNKNOWN234", 1},
	OP_UNKNOWN235: {OP_UNKNOWN235, "OP_UNKNOWN235", 1},
	OP_UNKNOWN236: {OP_UNKNOWN236, "OP_UNKNOWN236", 1},
	OP_UNKNOWN237: {OP_UNKNOWN237, "OP_UNKNOWN237", 1},
	OP_UNKNOWN238: {OP_UNKNOWN238, "OP_UNKNOWN238", 1},
	OP_UNKNOWN239: {OP_UNKNOWN239, "OP_UNKNOWN239", 1},
	OP_UNKNOWN240: {OP_UNKNOWN240, "OP_UNKNOWN240", 1},
	OP_UNKNOWN241: {OP_UNKNOWN241, "OP_UNKNOWN241", 1},
	OP_UNKNOWN242: {OP_UNKNOWN242, "OP_UNKNOWN242", 1},
	OP_UNKNOWN243: {OP_UNKNOWN243, "OP_UNKNOWN243", 1},
	OP_UNKNOWN244: {OP_UNKNOWN244, "OP_UNKNOWN244", 1},
	OP_UNKNOWN245: {OP_UNKNOWN245, "OP_UNKNOWN245", 1},
	OP_UNKNOWN246: {OP_UNKNOWN246, "OP_UNKNOWN246", 1},
	OP_UNKNOWN247: {OP_UNKNOWN247, "OP_UNKNOWN247", 1},
	OP_UNKNOWN248: {OP_UNKNOWN248, "OP_UNKNOWN248", 1},
	OP_UNKNOWN249: {OP_UNKNOWN249, "OP_UNKNOWN249", 1},
	OP_SMALLINTEGER: {OP_SMALLINTEGER, "OP_SMALLINTEGER", 1},
	OP_PUBKEYS:      {OP_PUBKEYS, "OP_PUBKEYS", 1},
	OP_UNKNOWN252:   {OP_UNKNOWN252, "OP_UNKNOWN252", 1},
	OP_PUBKEYHASH:   {OP_PUBKEYHASH, "OP_PUBKEYHASH", 1},
	OP_PUBKEY:       {OP_PUBKEY, "OP_PUBKEY", 1},
	OP_INVALIDOPCODE: {OP_INVALIDOPCODE, "OP_INVALIDOPCODE", 1},
}

type opcode struct {
	Value Hex 	`json:"value"`
	Name string `json:"name"`
	length int
}

type parsedOpcode struct {
	Opcode *opcode	`json:"opcode"`
	Data Hash		`json:"data"`
}

