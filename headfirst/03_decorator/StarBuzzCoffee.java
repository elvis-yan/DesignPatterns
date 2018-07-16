abstract class Beverage {
    String description = "Unknown Beverage";

    String getDescription() {
        return description;
    }

    abstract double cost();
}

abstract class CondimentDecorator extends Beverage {
    abstract String getDescription();
}

//_________________________________ Some Kinds of Beverage ____________________________________
class Espresso extends Beverage {
    Espresso() {
        description = "Espresso";
    }

    double cost() {
        return 1.99;
    }
}

class HouseBlend extends Beverage {
    HouseBlend() {
        description = "House Blend Coffee";
    }

    double cost() {
        return .89;
    }
}

//_____________________________________ Some Condiment ____________________________________________

class Mocha extends CondimentDecorator {
    Beverage beverage;

    Mocha(Beverage beverage){
        this.beverage = beverage;
    }

    String getDescription() {
        return beverage.getDescription() + ", Mocha";
    }

    double cost() {
        return .20 + beverage.cost();
    }
}

class Soy extends CondimentDecorator {
    Beverage beverage;

    Soy(Beverage beverage) {
        this.beverage = beverage;
    }

    String getDescription() {
        return beverage.getDescription() + ", Soy";
    }

    double cost() {
        return .15 + beverage.cost();
    }
}

class Whip extends CondimentDecorator {
    Beverage beverage;

    Whip(Beverage beverage) {
        this.beverage = beverage;
    }

    String getDescription() {
        return beverage.getDescription() + ", Whip";
    }

    double cost() {
        return .10 + beverage.cost();
    }
}

//______________________________________ StarBuzz Coffee ____________________________________________

public class StarBuzzCoffee {
    public static void main(String[] args) {
        Beverage beverage = new Espresso();
        System.out.println(beverage.getDescription() + " $ " + beverage.cost());
        
        Beverage beverage2 = new HouseBlend();
        beverage2 = new Mocha(beverage2);
        beverage2 = new Mocha(beverage2);
        beverage2 = new Whip(beverage2);
        System.out.println(beverage2.getDescription() + " $ " + beverage2.cost());
        
        Beverage beverage3 = new Espresso();
        beverage3 = new Soy(beverage3);
        beverage3 = new Mocha(beverage3);
        beverage3 = new Whip(beverage3);
        System.out.println(beverage3.getDescription() + " $ " + beverage3.cost());
    }
}