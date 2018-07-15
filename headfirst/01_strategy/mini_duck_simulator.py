# Duck Interface
class Duck:
    def __init__(self):
        self.flyer = None
        self.quacker = None
        raise "Duck is an interface, can't initialize instance"
    
    def swim(self):
        print("All ducks float, even decoys!")

    def display(self):
        raise NotImplementedError()

    def fly(self):
        self.flyer.fly()

    def quack(self):
        self.quacker.quack()

# Fly Behaviors
class FlyWithWings:
    def fly(self):
        print("I'm flying")

class FlyNoWay:
    def fly(self):
        print("I can't fly")

class FlyRocketPowered:
    def fly(self):
        print("I'm flying with a racket!")

# Quack Behaviors
class Quack:
    def quack(self):
        print("Quack")

class MuteQuack:
    def quack(self):
        print("<< Silence >>")

class Squeak:
    def quack(self):
        print("Squeak")

# Concrete Duck Type
class MallardDuck(Duck):
    def __init__(self, quacker=None, flyer=None):
        self.quacker = quacker
        self.flyer = flyer

    def display(self):
        print("I'm a real Mallard duck  😀")

class ModelDuck(Duck):
    def __init__(self, quacker=None, flyer=None):
        self.quacker = quacker
        self.flyer = flyer

    def display(self):
        print("I'm a model duck  😈")

def mini_duck_simulator():
    duck = MallardDuck(Quack(), FlyWithWings())
    duck.display()
    duck.swim()
    duck.quack()
    duck.fly()

    duck = ModelDuck(Quack(), FlyWithWings())
    duck.display()
    duck.swim()
    duck.quack()
    duck.fly()
    duck.quacker = Squeak()
    duck.flyer = FlyRocketPowered()
    duck.quack()
    duck.fly()

if __name__ == '__main__':
    mini_duck_simulator()
    Duck()