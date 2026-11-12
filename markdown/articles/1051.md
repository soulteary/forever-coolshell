# 如何在Java中避免equals方法的隐藏陷阱

**译者注** :你可能会觉得Java很简单，Object的equals实现也会非常简单，但是事实并不是你想象的这样，耐心的读完本文，你会发现你对Java了解的是如此的少。如果这篇文章是一份Java程序员的入职笔试，那么不知道有多少人会掉落到这样的陷阱中。原文转自<http://www.artima.com/lejava/articles/equality.html> 三位作者都是不同领域的大拿，有兴趣的读者可以从上面这个连接直接去阅读原文。

**摘要**  

本文描述重载equals方法的技术，这种技术即使是具现类的子类增加了字段也能保证equal语义的正确性。  

在《Effective Java》的第8项中，Josh Bloch描述了当继承类作为面向对象语言中的等价关系的基础问题，要保证派生类的equal正确性语义所会面对的困难。Bloch这样写到：

> **除非你忘记了面向对象抽象的好处，否则在当你继承一个新类或在类中增加了一个值组件时你无法同时保证equal的语义依然正确**

在《Programming in Scala》中的第28章演示了一种方法，这种方法允许即使继承了新类，增加了新的值组件，equal的语义仍然能得到保证。虽然在这本书中这项技术是在使用Scala类环境中，但是这项技术同样可以应用于Java定义的类中。在本文中的描述来自于Programming in Scala中的文字描述，但是代码被我从scala翻译成了Java

目录


- [常见的等价方法陷阱](#%E5%B8%B8%E8%A7%81%E7%9A%84%E7%AD%89%E4%BB%B7%E6%96%B9%E6%B3%95%E9%99%B7%E9%98%B1 "常见的等价方法陷阱")
- [陷阱1：定义错误equals方法签名(signature)](#%E9%99%B7%E9%98%B11%EF%BC%9A%E5%AE%9A%E4%B9%89%E9%94%99%E8%AF%AFequals%E6%96%B9%E6%B3%95%E7%AD%BE%E5%90%8Dsignature "陷阱1：定义错误equals方法签名(signature)")
- [陷阱2：重载了equals的但没有同时重载hashCode的方法](#%E9%99%B7%E9%98%B12%EF%BC%9A%E9%87%8D%E8%BD%BD%E4%BA%86equals%E7%9A%84%E4%BD%86%E6%B2%A1%E6%9C%89%E5%90%8C%E6%97%B6%E9%87%8D%E8%BD%BDhashCode%E7%9A%84%E6%96%B9%E6%B3%95 "陷阱2：重载了equals的但没有同时重载hashCode的方法")
- [陷阱3：建立在会变化字段上的equals定义](#%E9%99%B7%E9%98%B13%EF%BC%9A%E5%BB%BA%E7%AB%8B%E5%9C%A8%E4%BC%9A%E5%8F%98%E5%8C%96%E5%AD%97%E6%AE%B5%E4%B8%8A%E7%9A%84equals%E5%AE%9A%E4%B9%89 "陷阱3：建立在会变化字段上的equals定义")
- [陷阱4：不满足等价关系的equals错误定义](#%E9%99%B7%E9%98%B14%EF%BC%9A%E4%B8%8D%E6%BB%A1%E8%B6%B3%E7%AD%89%E4%BB%B7%E5%85%B3%E7%B3%BB%E7%9A%84equals%E9%94%99%E8%AF%AF%E5%AE%9A%E4%B9%89 "陷阱4：不满足等价关系的equals错误定义")
- [canEqual 方法](#canEqual_%E6%96%B9%E6%B3%95 "canEqual 方法")
#### 常见的等价方法陷阱

java.lang.Object 类定义了equals这个方法，它的子类可以通过重载来覆盖它。不幸的是，在面向对象中写出正确的equals方法是非常困难的。事实上，在研究了大量的Java代码后，2007 paper的作者得出了如下的一个结论：

> 几乎所有的equals方法的实现都是错误的！

这个问题是因为等价是和很多其他的事物相关联。例如其中之一，一个的类型C的错误等价方法可能意味着你无法将这个类型C的对象可信赖的放入到容器中。比如说，你有两个元素elem1和elem2他们都是类型C的对象，并且他们是相等，即elem1.equals(elm2)返回ture。但是，只要这个equals方法是错误的实现，那么你就有可能会看见如下的一些行为：

```
Set hashSet<C> = new java.util.HashSet<C>();
hashSet.add(elem1);
hashSet.contains(elem2);    // returns false!
```

当equals重载时，这里有4个会引发equals行为不一致的常见陷阱：

1. 定义了错误的equals方法签名(signature) Defining equals with the wrong signature.
2. 重载了equals的但没有同时重载hashCode的方法。 Changing equals without also changing hashCode.
3. 建立在会变化字域上的equals定义。 Defining equals in terms of mutable fields.
4. 不满足等价关系的equals错误定义 Failing to define equals as an equivalence relation.

在剩下的章节中我们将依次讨论这4中陷阱。

#### 陷阱1：定义错误equals方法签名(signature)

考虑为下面这个简单类Point增加一个等价性方法：

```
public class Point {

    private final int x;
    private final int y;

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public int getX() {
        return x;
    }

    public int getY() {
        return y;
    }

    // ...
}
```

看上去非常明显，但是按照这种方式来定义equals就是错误的。

```
// An utterly wrong definition of equals
public boolean equals(Point other) {
  return (this.getX() == other.getX() && this.getY() == other.getY());
}
```

这个方法有什么问题呢？初看起来，它工作的非常完美：

```
Point p1 = new Point(1, 2);
Point p2 = new Point(1, 2);

Point q = new Point(2, 3);

System.out.println(p1.equals(p2)); // prints true

System.out.println(p1.equals(q)); // prints false
```

然而，当我们一旦把这个Point类的实例放入到一个容器中问题就出现了：

```
import java.util.HashSet;

HashSet<Point> coll = new HashSet<Point>();
coll.add(p1);

System.out.println(coll.contains(p2)); // prints false
```

为什么coll中没有包含p2呢？甚至是p1也被加到集合里面，p1和p2是是等价的对象吗？在下面的程序中，我们可以找到其中的一些原因，定义p2a是一个指向p2的对象，但是p2a的类型是Object而非Point类型：

```
Object p2a = p2;
```

现在我们重复第一个比较，但是不再使用p2而是p2a,我们将会得到如下的结果：

```
System.out.println(p1.equals(p2a)); // prints false
```

到底是那里出了了问题？事实上，之前所给出的equals版本并没有覆盖Object类的equals方法，因为他的类型不同。下面是Object的equals方法的定义

```
public boolean equals(Object other)
```

因为Point类中的equals方法使用的是以Point类而非Object类做为参数，因此它并没有覆盖Object中的equals方法。而是一种变化了的重载。在Java中重载被解析为静态的参数类型而非运行期的类型，因此当静态参数类型是Point,Point的equals方法就被调用。然而当静态参数类型是Object时，Object类的equals就被调用。因为这个方法并没有被覆盖，因此它仍然是实现成比较对象标示。这就是为什么虽然p1和p2a具有同样的x,y值，”p1.equals(p2a)”仍然返回了false。这也是会什么HasSet的contains方法返回false的原因，因为这个方法操作的是泛型，他调用的是一般化的Object上equals方法而非Point类上变化了的重载方法equals

一个更好但不完美的equals方法定义如下：

```
// A better definition, but still not perfect
@Override public boolean equals(Object other) {
    boolean result = false;
    if (other instanceof Point) {
        Point that = (Point) other;
        result = (this.getX() == that.getX() && this.getY() == that.getY());
    }
    return result;
}
```

现在equals有了正确的类型，它使用了一个Object类型的参数和一个返回布尔型的结果。这个方法的实现使用instanceof操作和做了一个造型。它首先检查这个对象是否是一个Point类，如果是，他就比较两个点的坐标并返回结果，否则返回false。

#### 陷阱2：重载了equals的但没有同时重载hashCode的方法

如果你使用上一个定义的Point类进行p1和p2a的反复比较，你都会得到你预期的true的结果。但是如果你将这个类对象放入到HashSet.contains()方法中测试，你就有可能仍然得到false的结果：

```
Point p1 = new Point(1, 2);
Point p2 = new Point(1, 2);

HashSet<Point> coll = new HashSet<Point>();
coll.add(p1);

System.out.println(coll.contains(p2)); // 打印 false (有可能)
```

事实上，这个个结果不是100%的false，你也可能有返回ture的经历。如果你得到的结果是true的话，那么你试试其他的坐标值，最终你一定会得到一个在集合中不包含的结果。导致这个结果的原因是Point重载了equals却没有重载hashCode。  

注意上面例子的的容器是一个HashSet，这就意味着容器中的元素根据他们的哈希码被被放入到”哈希桶 hash buckets”中。contains方法首先根据哈希码在哈希桶中查找，然后让桶中的所有元素和所给的参数进行比较。现在，虽然最后一个Point类的版本重定义了equals方法，但是它并没有同时重定义hashCode。因此，hashCode仍然是Object类的那个版本，即：所分配对象的一个地址的变换。所以p1和p2的哈希码理所当然的不同了，甚至是即时这两个点的坐标完全相同。不同的哈希码导致他们具有极高的可能性被放入到集合中不同的哈希桶中。contains方法将会去找p2的哈希码对应哈希桶中的匹配元素。但是大多数情况下，p1一定是在另外一个桶中，因此，p2永远找不到p1进行匹配。当然p2和p2也可能偶尔会被放入到一个桶中，在这种情况下，contains的结果就为true了。

最新一个Point类实现的问题是，它的实现违背了作为Object类的定义的hashCode的语义。

> **如果两个对象根据equals(Object)方法是相等的，那么在这两个对象上调用hashCode方法应该产生同样的值**

事实上，在Java中，hashCode和equals需要一起被重定义是众所周知的。此外，hashCode只可以依赖于equals依赖的域来产生值。对于Point这个类来说，下面的的hashCode定义是一个非常合适的定义。

```
public class Point {

    private final int x;
    private final int y;

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public int getX() {
        return x;
    }

    public int getY() {
        return y;
    }

    @Override public boolean equals(Object other) {
        boolean result = false;
        if (other instanceof Point) {
            Point that = (Point) other;
            result = (this.getX() == that.getX() && this.getY() == that.getY());
        }
        return result;
    }
    @Override public int hashCode() {
        return (41 * (41 + getX()) + getY());
    }

}
```

这只是hashCode一个可能的实现。x域加上常量41后的结果再乘与41并将结果在加上y域的值。这样做就可以以低成本的运行时间和低成本代码大小得到一个哈希码的合理的分布(**译者注：**性价比相对较高的做法)。  

增加hashCode方法重载修正了定义类似Point类等价性的问题。然而，关于类的等价性仍然有其他的问题点待发现。

#### 陷阱3：建立在会变化字段上的equals定义

让我们在Point类做一个非常微小的变化

```
public class Point {
    private int x;
    private int y;

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public int getX() {
        return x;
    }

    public int getY() {
        return y;
    }
    public void setX(int x) { // Problematic
        this.x = x;
    }

    public void setY(int y) {
        this.y = y;
    }

    @Override public boolean equals(Object other) {
        boolean result = false;
        if (other instanceof Point) {
            Point that = (Point) other;
            result = (this.getX() == that.getX() && this.getY() == that.getY());
        }
        return result;
    }

    @Override public int hashCode() {
        return (41 * (41 + getX()) + getY());
    }
}
```

唯一的不同是x和y域不再是final，并且两个set方法被增加到类中来，并允许客户改变x和y的值。equals和hashCode这个方法的定义现在是基于在这两个会发生变化的域上，因此当他们的域的值改变时，结果也就跟着改变。因此一旦你将这个point对象放入到集合中你将会看到非常神奇的效果。

```
Point p = new Point(1, 2);

HashSet<Point> coll = new HashSet<Point>();
coll.add(p);

System.out.println(coll.contains(p)); // 打印 true
```

现在如果你改变p中的一个域，这个集合中还会包含point吗，我们将拭目以待。

```
p.setX(p.getX() + 1);

System.out.println(coll.contains(p)); // (有可能)打印 false
```

看起来非常的奇怪。p去那里去了？如果你通过集合的迭代器来检查p是否包含，你将会得到更奇怪的结果。

```
Iterator<Point> it = coll.iterator();
boolean containedP = false;
while (it.hasNext()) {
    Point nextP = it.next();
    if (nextP.equals(p)) {
        containedP = true;
        break;
    }
}

System.out.println(containedP); // 打印 true
```

结果是，集合中不包含p，但是p在集合的元素中！到底发生了什么！当然，所有的这一切都是在x域的修改后才发生的，p最终的的hashCode是在集合coll错误的哈希桶中。即，原始哈希桶不再有其新值对应的哈希码。换句话说，p已经在集合coll的是视野范围之外，虽然他仍然属于coll的元素。

从这个例子所得到的教训是，当equals和hashCode依赖于会变化的状态时，那么就会给用户带来问题。如果这样的对象被放入到集合中，用户必须小心，不要修改这些这些对象所依赖的状态，这是一个小陷阱。如果你需要根据对象当前的状态进行比较的话，你应该不要再重定义equals，应该起其他的方法名字而不是equals。对于我们的Point类的最后的定义，我们最好省略掉hashCode的重载，并将比较的方法名命名为equalsContents，或其他不同于equals的名字。那么Point将会继承原来默认的equals和hashCode的实现，因此当我们修改了x域后p依然会呆在其原来在容器中应该在位置。

#### 陷阱4：不满足等价关系的equals错误定义

Object中的equals的规范阐述了equals方法必须实现在非null对象上的等价关系：

- 自反原则：对于任何非null值X,表达式x.equals(x)总返回true。
- 等价性：对于任何非空值x和y，那么当且仅当y.equals(x)返回真时，x.equals(y)返回真。
- 传递性：对于任何非空值x,y,和z，如果x.equals(y)返回真，且y.equals(z)也返回真，那么x.equals(z)也应该返回真。
- 一致性：对于非空x,y，多次调用x.equals(y)应该一致的返回真或假。提供给equals方法比较使用的信息不应该包含改过的信息。
- 对于任何非空值x,x.equals(null)应该总返回false.

Point类的equals定义已经被开发成了足够满足equals规范的定义。然而，当考虑到继承的时候，事情就开始变得非常复杂起来。比如说有一个Point的子类ColoredPoint，它比Point多增加了一个类型是Color的color域。假设Color被定义为一个枚举类型：

```
public enum Color {
    RED, ORANGE, YELLOW, GREEN, BLUE, INDIGO, VIOLET;
}
```

ColoredPoint重载了equals方法，并考虑到新加入color域，代码如下：

```
public class ColoredPoint extends Point { // Problem: equals not symmetric

    private final Color color;

    public ColoredPoint(int x, int y, Color color) {
        super(x, y);
        this.color = color;
    }

    @Override public boolean equals(Object other) {
        boolean result = false;
        if (other instanceof ColoredPoint) {
            ColoredPoint that = (ColoredPoint) other;
            result = (this.color.equals(that.color) && super.equals(that));
        }
        return result;
    }
}
```

这是很多程序员都有可能写成的代码。注意在本例中，类ColoredPointed不需要重载hashCode，因为新的ColoredPoint类上的equals定义，严格的重载了Point上equals的定义。hashCode的规范仍然是有效，如果两个着色点(colored point)相等，其坐标必定相等，因此它的hashCode也保证了具有同样的值。

对于ColoredPoint类自身对象的比较是没有问题的，但是如果使用ColoredPoint和Point混合进行比较就要出现问题。

```
Point p = new Point(1, 2);

ColoredPoint cp = new ColoredPoint(1, 2, Color.RED);

System.out.println(p.equals(cp)); // 打印真 true

System.out.println(cp.equals(p)); // 打印假 false
```

“p等价于cp”的比较这个调用的是定义在Point类上的equals方法。这个方法只考虑两个点的坐标。因此比较返回真。在另外一方面，“cp等价于p”的比较这个调用的是定义在ColoredPoint类上的equals方法，返回的结果却是false，这是因为p不是ColoredPoint，所以equals这个定义违背了对称性。

违背对称性对于集合来说将导致不可以预期的后果，例如：

```
Set<Point> hashSet1 = new java.util.HashSet<Point>();
hashSet1.add(p);
System.out.println(hashSet1.contains(cp));    // 打印 false

Set<Point> hashSet2 = new java.util.HashSet<Point>();
hashSet2.add(cp);
System.out.println(hashSet2.contains(p));    // 打印 true
```

因此虽然p和cp是等价的，但是contains测试中一个返回成功，另外一个却返回失败。  

你如何修改equals的定义，才能使得这个方法满足对称性？本质上说有两种方法，你可以使得这种关系变得更一般化或更严格。更一般化的意思是这一对对象，a和b，被用于进行对比，无论是a比b还是b比a 都返回true，下面是代码：

```
public class ColoredPoint extends Point { // Problem: equals not transitive

    private final Color color;

    public ColoredPoint(int x, int y, Color color) {
        super(x, y);
        this.color = color;
    }

    @Override public boolean equals(Object other) {
        boolean result = false;
        if (other instanceof ColoredPoint) {
            ColoredPoint that = (ColoredPoint) other;
            result = (this.color.equals(that.color) && super.equals(that));
        }
        else if (other instanceof Point) {
            Point that = (Point) other;
            result = that.equals(this);
        }
        return result;
    }
}
```

在ColoredPoint中的equals的新定义比老定义中检查了更多的情况:如果对象是一个Point对象而不是ColoredPoint，方法就转变为Point类的equals方法调用。这个所希望达到的效果就是equals的对称性，不管”cp.equals(p)”还是”p.equals(cp)”的结果都是true。然而这种方法，equals的规范还是被破坏了，现在的问题是这个新等价性不满足传递性。考虑下面的一段代码实例，定义了一个点和这个点上上两种不同颜色点：

```
ColoredPoint redP = new ColoredPoint(1, 2, Color.RED);
ColoredPoint blueP = new ColoredPoint(1, 2, Color.BLUE);
```

redP等价于p，p等价于blueP

```
System.out.println(redP.equals(p)); // prints true

System.out.println(p.equals(blueP)); // prints true
```

然而，对比redP和blueP的结果是false:

```
System.out.println(redP.equals(blueP)); // 打印 false
```

因此，equals的传递性就被违背了。  

使equals的关系更一般化似乎会将我们带入到死胡同。我们应该采用更严格化的方法。一种更严格化的equals方法是认为不同类的对象是不同的。这个可以通过修改Point类和ColoredPoint类的equals方法来达到。你能增加额外的比较来检查是否运行态的这个Point类和那个Point类是同一个类，就像如下所示的代码一样：

```
// A technically valid, but unsatisfying, equals method
public class Point {

    private final int x;
    private final int y;

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public int getX() {
        return x;
    }

    public int getY() {
        return y;
    }

    @Override public boolean equals(Object other) {
        boolean result = false;
        if (other instanceof Point) {
            Point that = (Point) other;
            result = (this.getX() == that.getX() && this.getY() == that.getY()
                    && this.getClass().equals(that.getClass()));
        }
        return result;
    }

    @Override public int hashCode() {
        return (41 * (41 + getX()) + getY());
    }
}
```

你现在可以将ColoredPoint类的equals实现用回刚才那个不满足对称性要的equals实现了。

```
public class ColoredPoint extends Point { // 不再违反对称性需求

    private final Color color;

    public ColoredPoint(int x, int y, Color color) {
        super(x, y);
        this.color = color;
    }

    @Override public boolean equals(Object other) {
        boolean result = false;
        if (other instanceof ColoredPoint) {
            ColoredPoint that = (ColoredPoint) other;
            result = (this.color.equals(that.color) && super.equals(that));
        }
        return result;
    }
}
```

这里，Point类的实例只有当和另外一个对象是同样类，并且有同样的坐标时候，他们才被认为是相等的，即意味着 .getClass()返回的是同样的值。这个新定义的等价关系满足了对称性和传递性因为对于比较对象是不同的类时结果总是false。所以着色点(colored point)永远不会等于点(point)。通常这看起来非常合理，但是这里也存在着另外一种争论——这样的比较过于严格了。

考虑我们如下这种稍微的迂回的方式来定义我们的坐标点(1,2)

```
Point pAnon = new Point(1, 1) {
    @Override public int getY() {
        return 2;
    }
};
```

pAnon等于p吗？答案是假，因为p和pAnon的java.lang.Class对象不同。p是Point，而pAnon是Point的一个匿名派生类。但是，非常清晰的是pAnon的确是在坐标1，2上的另外一个点。所以将他们认为是不同的点是没有理由的。

#### canEqual 方法

到此，我们看其来似乎是遇到阻碍了，存在着一种正常的方式不仅可以在不同类继承层次上定义等价性，并且保证其等价的规范性吗？事实上，的确存在这样的一种方法，但是这就要求除了重定义equals和hashCode外还要另外的定义一个方法。基本思路就是在重载equals(和hashCode)的同时，它应该也要要明确的声明这个类的对象永远不等价于其他的实现了不同等价方法的超类的对象。为了达到这个目标，我们对每一个重载了equals的类新增一个方法canEqual方法。这个方法的方法签名是：

```
public boolean canEqual(Object other)
```

如果other 对象是canEquals(重)定义那个类的实例时，那么这个方法应该返回真，否则返回false。这个方法由equals方法调用，并保证了两个对象是可以相互比较的。下面Point类的新的也是最终的实现：

```
public class Point {

    private final int x;
    private final int y;

    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public int getX() {
        return x;
    }

    public int getY() {
        return y;
    }

    @Override public boolean equals(Object other) {
        boolean result = false;
        if (other instanceof Point) {
            Point that = (Point) other;
            result =(that.canEqual(this) && this.getX() == that.getX() && this.getY() == that.getY());
        }
        return result;
    }

    @Override public int hashCode() {
        return (41 * (41 + getX()) + getY());
    }
    public boolean canEqual(Object other) {
        return (other instanceof Point);
    }

}
```

这个版本的Point类的equals方法中包含了一个额外的需求，通过canEquals方法来决定另外一个对象是否是是满足可以比较的对象。在Point中的canEqual宣称了所有的Point类实例都能被比较。

下面是ColoredPoint相应的实现

```
public class ColoredPoint extends Point { // 不再违背对称性

    private final Color color;

    public ColoredPoint(int x, int y, Color color) {
        super(x, y);
        this.color = color;
    }

    @Override public boolean equals(Object other) {
        boolean result = false;
        if (other instanceof ColoredPoint) {
            ColoredPoint that = (ColoredPoint) other;
            result = (that.canEqual(this) && this.color.equals(that.color) && super.equals(that));
        }
        return result;
    }

    @Override public int hashCode() {
        return (41 * super.hashCode() + color.hashCode());
    }
    @Override public boolean canEqual(Object other) {
        return (other instanceof ColoredPoint);
    }
}
```

在上显示的新版本的Point类和ColoredPoint类定义保证了等价的规范。等价是对称和可传递的。比较一个Point和ColoredPoint类总是返回false。因为点p和着色点cp,“p.equals(cp)返回的是假。并且，因为cp.canEqual(p)总返回false。相反的比较，cp.equals(p)同样也返回false，由于p不是一个ColoredPoint，所以在ColoredPoint的equals方法体内的第一个instanceof检查就失败了。

另外一个方面，不同的Point子类的实例却是可以比较的，同样没有重定义等价性方法的类也是可以比较的。对于这个新类的定义，p和pAnon的比较将总返回true。下面是一些例子：

```
Point p = new Point(1, 2);

ColoredPoint cp = new ColoredPoint(1, 2, Color.INDIGO);

Point pAnon = new Point(1, 1) {
    @Override public int getY() {
        return 2;
    }
};

Set<Point> coll = new java.util.HashSet<Point>();
coll.add(p);

System.out.println(coll.contains(p)); // 打印 true

System.out.println(coll.contains(cp)); // 打印 false

System.out.println(coll.contains(pAnon)); // 打印 true
```

这些例子显示了如果父类在equals的实现定义并调用了canEquals，那么开发人员实现的子类就能决定这个子类是否可以和它父类的实例进行比较。例如ColoredPoint，因为它以”一个着色点永远不可以等于普通不带颜色的点重载了” canEqual，所以他们就不能比较。但是因为pAnon引用的匿名子类没有重载canEqual,因此它的实例就可以和Point的实例进行对比。

canEqual方法的一个潜在的争论是它是否违背了Liskov替换准则(LSP)。例如，通过比较运行态的类来实现的比较技术(**译者注：** canEqual的前一版本，使用.getClass()的那个版本)，将导致不能定义出一个子类，这个子类的实例可以和其父类进行比较，因此就违背了LSP。这是因为，LSP原则是这样的，在任何你能使用父类的地方你都可以使用子类去替换它。在之前例子中，虽然cp的x,y坐标匹配那些在集合中的点，然而”coll.contains(cp)”仍然返回false，这看起来似乎违背得了LSP准则，因为你不能这里能使用Point的地方使用一个ColoredPointed。但是我们认为这种解释是错误的，因为LSP原则并没有要求子类和父类的行为一致，而仅要求其行为能一种方式满足父类的规范。

通过比较运行态的类来编写equals方法(**译者注：** canEqual的前一版本，使用.getClass()的那个版本)的问题并不是违背LSP准则的问题，但是它也没有为你指明一种创建派生类的实例能和父类实例进行对比的的方法。例如，我们使用这种运行态比较的技术在之前的”coll.contains(pAnon)”将会返回false，并且这并不是我们希望的。相反我们希望“coll.contains(cp)”返回false，因为通过在ColoredPoint中重载的equals，我基本上可以说，一个在坐标1，2上着色点和一个坐标1，2上的普通点并不是一回事。然而，在最后的例子中，我们能传递Point两种不同的子类实例到集合中contains方法，并且我们能得到两个不同的答案，并且这两个答案都正确。

**–全文完–**

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
