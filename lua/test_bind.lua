require "function"
-- local http = require("http")
-- local json = require("json")
local driver = require("driver")
local han = require("han")
local modely = require("modely")

local function get()
    local d = driver.new()
    local h = han.new()
    local y = modely.new()

    dumptag("d",getmetatable(d))
    dumptag("h",getmetatable(h))
    dumptag("y",getmetatable(y))

    local r = driver.GetByCb("getbycb", function(str, int, bool)
        return str,true
    end)
    printInfo("driver.GetByCb %s",r)

    local ss, mm = driver.GetBrand("get_brand", 99, 3.1415626)
    dumptag("driver.GetBrand ss",ss)
    dumptag("driver.GetBrand mm",mm)

    local ss, mm = driver.GetCars("get_cars", 555, 23.45)
    dumptag("driver.GetCars ss",ss)
    dumptag("driver.GetCars mm",mm)

    for i,v in ipairs(ss) do
        printInfo("driver.GetCars ss %d,%v,%v",i,v:Name(),v:Price())
    end
    for i,v in pairs(mm) do
        printInfo("driver.GetCars mm before %d,%v,%v",i,v:Name(),v:Mpg())
        v:Name(v:Name() .. "__dev")
        v:Mpg(v:Mpg() * 100)
    end
    for i,v in pairs(mm) do
        printInfo("driver.GetCars mm after  %d,%v,%v",i,v:Name(),v:Mpg())
    end

    n,e = d:DoAdd(1,20)
    printInfo("d.DoAdd %d,%s",n,e)

    local r1,r2,r3,r4 = d:DoMap(100,63.333,{name=100,age=10})
    printInfo("d.DoMap %d,%f,%s,%s",r1,r2,r3,r4)
    dumptag("d.DoMap r3",r3)

    local r1,r2 = d:DoSlice({10000,1000,100,10,1})
    printInfo("d.DoSlice %s,%s",r1,r2)
    dumptag("d.DoSlice r1",r1)

    local r1,r2 = d:DoMapAny(100,"ABC",{qq=1,ww="B",ee=true})
    printInfo("d.DoMapAny %s",r1)
    dumptag("d.DoMapAny r2",r2)

    y:SpeedUp(50)
    local r = d:Drive(y,function (i)
        printInfo("d.Drive check %d %d",i,y:CurrSpeed())
        sleep(1)
        if i<10 then
            return true
        else
            return false
        end
    end)
    y:SpeedUp(-50)
    printInfo("d.Drive %s %d",r,y:CurrSpeed())



end
get()
return "AAA",1000