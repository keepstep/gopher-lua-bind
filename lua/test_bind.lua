require "function"
-- local http = require("http")
-- local json = require("json")
local aaa = require("aaa")
local bbb = require("bbb")
local ccc = require("ccc")
local tesla = require("tesla")
local teslae = require("modele")

local function get()
    a = aaa.new()
    b = bbb.new()
    c = ccc.new()
    n,e = a:DoAdd(1,20)
    printInfo("AAA.DoAdd %d,%s",n,e)
    nn,aa,cc,ee = a:Do(1,b,c)
    printInfo("AAA.DoAdd %s,%s,%s,%s",nn,aa,cc,ee)
    a:Age(2)
    ag = a:Age()
    a:Age(ag*100)
    printInfo("AAA.Age %d,%d",ag,a:Age())
    dumptag("AAA",getmetatable(a))

    r1,r2,r3,r4 = a:DoMap(100,63.333,{name=100,age=10})
    printInfo("a.DoMap %d,%f,%s,%s",r1,r2,r3,r4)
    dumptag("a.DoMap r3",r3)

    r1,r2 = a:DoSlice({10000,1000,100,10,1})
    printInfo("a.DoSlice %s,%s",r1,r2)
    dumptag("a.DoSlice r1",r1)
    r1 = a:DoMapAny(100,"ABC",{qq=1,ww="B",ee=true})
    printInfo("a.DoMapAny %s",r1)
    dumptag("a.DoMapAny r1",r1)
    

    r1 = tesla.TeslaCompare("ABC")
    printInfo("tesla.TeslaCompare %d",r1)

    r1,r2,r3,r4 = tesla.TeslaTest(100,63.333,{name=100,age=10})
    printInfo("tesla.TeslaTest %d,%f,%s,%s",r1,r2,r3,r4)
    dumptag("tesla.TeslaTest r3",r3)

    local t =  tesla.new()
    dumptag("tesla.Cars 1",t:Cars())
    t:Cars({"e","x","y"})
    dumptag("tesla.Cars 2",t:Cars())

    dumptag("tesla.Models 1",t:Models())
    t:Models({e=10,x=100,y=100})
    dumptag("tesla.Models 2",t:Models())
    
    r1 = tesla.TeslaGetAAA("55")
    printInfo("tesla.TeslaGetAAA %s %s %d",r1,r1:Name(),r1:Age())

    local cb = function (name)
        printInfo("tesla.RunCb %s",name)
        return "callback_lua"
    end
    t:RunCb(cb)
    t:Run("run01",100)
    t:Run("run02",200)
    t:Run("run03",300)

    t:Models({e=10,x=100,y=100})
    dumptag("tesla.Models 2",t:Models())

    rc = tesla.TeslaGetCmp("BeyondCompare")
    printInfo("tesla Cmp BeyondCompare  %s",rc:Equal("BeyondCompare"))
    printInfo("tesla Cmp Diff  %s",rc:Equal("Diff"))

    rcb = t:RunCallback("callback",function (name,age,bs,ms)
        printInfo("RunCallback  %s %d %s %s",name,age,bs,ms)
        dumptag("RunCallback []bool",bs)
        dumptag("RunCallback map[string]int",ms)
        return "abcdefg12580",true
    end)
    printInfo("tesla RunCallback  %s",rcb)

    local te = teslae.new()
    te:Speed(1000)
    t:ModelE(te)
    t:RunE()
    te:Run(200)

    dumptag("tesla",getmetatable(t))


end
get()
return "AAA",1000