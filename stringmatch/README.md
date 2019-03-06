## AlgorithmGo.stringmatch

-----

Algorithm.stringmatch provide three different algorithm to match the pattern(or template) string in master string. It's **Naive** algorithm(or BF), **KMP** algorithm and the **Boyer-Moore** algorithm.

Obviously these algorithm's performance is incremental: BM faster than KMP faster than Naive.

In order to compare the performance of these three algorithm, I write this go package. But I didn't give you the function which compare these algorithm's performance. I hope you do it yourself.

-----

## Package Information:

### Package: stringmatch

#### Function: Naive(master, temp string) (find bool, pos int);

#### Function: KMP(master, temp string) (find bool, pos int);

#### Function: Boyer-Moore(master, temp string) (find bool, pos int);

-----

## Some Details:

More internal function like **getNext()** please read my code by yourself.

-----

## In the End:

### Contributor: ShiinaOrez
