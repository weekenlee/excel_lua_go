function add(a , b)
    return a+b
end

function sum(...)
    local arg={...}
    local s
    for k, v in ipairs(arg) do
        s = s + v
    end
    return s
end