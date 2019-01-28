require("INC_Class")

cAnimal = setclass("Animal")

function cAnimal.methods:init(action, cutename)
    self.superaction = acetion
    self.supercutename = cutename
end

cTiger = setclass("Tiger", cAnimal)

function cTiger.methods:init(cutename)
    self:init_super("HUNT(Tiger)", "Zoo Animal(Tiger)")
    self.action = "FOAR FOR ME"
    self.cutename = cutename
end

Triger1 = cAnimal:new("HUNT", "Zoo Animal")
Triger2 = cTiger:new("Mr Grumpy")
Triger3 = cTiger:new("Mr Hungry")

print("Classname for tiger1", Triger1:classname())
print("Classname for tiger2", Triger2:classname())
print("Classname for tiger3", Triger3:classname())
