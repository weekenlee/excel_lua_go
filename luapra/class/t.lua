t = {11,22,33,you="one",me="two"}
s={}
mt={__newindex = s, __index = _G}
setmetatable(t, mt)

print(t.you)
x="wow"
print(t.x)
t[5]=009

print("_________ s")
for k,v in pairs(s) do
    print(k,v)
end

print("_________ t")
for k,v in pairs(t) do
    print(k,v)
end

print("_________ _G")
for k,v in pairs(_G) do
    print(k,v)
end
