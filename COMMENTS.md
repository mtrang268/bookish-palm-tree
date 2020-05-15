## Requirements
1. Go (https://golang.org/doc/install)
2. Setup GOPATH to include current directory.
```
> pwd
/Users/gocode/slcsp
> echo $GOPATH
/Users/gocode
```

## Building
Run the following command in the root directory.
```
go build
```

## Running
```
./slcsp --help
Usage of ./slcsp:
  -planFile string
    	File containing all the health plans (default "testdata/plans.csv")
  -slcspFile string
    	File containing zipCodes to compute SLCSP information (default "testdata/slcsp.csv")
  -targetValue string
    	Sets the level whose second lowest cost plan will be returned. (default "Silver")
  -zipCodeFile string
    	File containing a mapping of ZIP code to county/counties & rate area(s) (default "testdata/zips.csv")
```

When running from the root path, it will use the test data found in `testdata/`. No options need to be provided.

### Sample
Running against `/testdata`
```
./slcsp
zipcode,rate
64148,245.20
67118,212.35
40813,
18229,231.48
51012,252.76
79168,243.68
54923,
67651,249.44
49448,221.63
27702,283.08
47387,326.98
50014,287.30
33608,268.49
06239,
54919,243.77
46706,
14846,
48872,
43343,
77052,243.72
07734,
95327,
12961,
26716,291.76
48435,
53181,306.56
52654,242.39
58703,297.93
91945,
52146,254.56
56097,
21777,
42330,
38849,285.69
77586,243.72
39745,265.73
03299,240.45
63359,
60094,209.95
15935,184.97
39845,325.64
48418,
28411,307.51
37333,219.29
75939,234.50
07184,
86313,292.90
61232,222.38
20047,
47452,
31551,290.60
```