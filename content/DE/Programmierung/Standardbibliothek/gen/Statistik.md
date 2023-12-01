+++
title = "Statistik"
weight = 1
+++
# Duden/Statistik Funktionen
<details>
<summary><h2>Höchste_ListeZ</h2></summary>
<ul>
<pre>
Gibt den höchsten Wert der Zahlen Liste zurück.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Zahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der höchste Wert aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl maxNum ist der minimale Wert einer Zahl.
Für jede Zahl z in liste, speichere die größere Zahl von z und maxNum in maxNum.
Gib maxNum zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Höchste_ListeK</h2></summary>
<ul>
<pre>
Gibt den höchsten Wert der Kommazahlen Liste zurück.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der höchste Wert aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahl maxNum ist der minimale Wert einer Kommazahl.
Für jede Kommazahl z in liste, speichere die größere Zahl von z und maxNum in maxNum.
Gib maxNum zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Kleinste_ListeZ</h2></summary>
<ul>
<pre>
Gibt den kleinsten Wert der Zahlen Liste zurück.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Zahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der kleinste Wert aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl maxNum ist der maximale Wert einer Zahl.
Für jede Zahl z in liste, speichere die kleinere Zahl von z und maxNum in maxNum.
Gib maxNum zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Kleinste_ListeK</h2></summary>
<ul>
<pre>
Gibt den kleinsten Wert der Kommazahlen Liste zurück.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der kleinste Wert aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahl maxNum ist der maximale Wert einer Kommazahl.
Für jede Kommazahl z in liste, speichere die kleinere Zahl von z und maxNum in maxNum.
Gib maxNum zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Mindestens_Liste</h2></summary>
<ul>
<pre>
Gibt die Summe der relativen Häufigkeiten aller Zahlen größer als, oder x.
</pre>
	<li>Parameter: <code>x</code>, <code>liste</code></li>
	<li>Parameter Typen: <code>Kommazahl</code>, <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;wie viel Prozent der Zahlen aus &lt;liste&gt; mindestens &lt;x&gt; sind&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahl a ist 0,0.
Für jede Kommazahl z in liste, wenn z kleiner als, oder x ist, erhöhe a um 1.
Gib a durch die Länge von liste zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Höchstens_Liste</h2></summary>
<ul>
<pre>
Gibt die Summe der relativen Häufigkeiten aller Zahlen kleiner als, oder x.
</pre>
	<li>Parameter: <code>x</code>, <code>liste</code></li>
	<li>Parameter Typen: <code>Kommazahl</code>, <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;wie viel Prozent der Zahlen aus &lt;liste&gt; höchstens &lt;x&gt; sind&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahl a ist 0,0.
Für jede Kommazahl z in liste, wenn z größer als, oder x ist, erhöhe a um 1.
Gib a durch die Länge von liste zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Zwischen_Liste</h2></summary>
<ul>
<pre>
Gibt die Summe der relativen Häufigkeiten aller Zahlen zwischen x und y.
</pre>
	<li>Parameter: <code>x</code>, <code>y</code>, <code>liste</code></li>
	<li>Parameter Typen: <code>Kommazahl</code>, <code>Kommazahl</code>, <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;wie viel Prozent der Zahlen aus &lt;liste&gt; zwischen &lt;x&gt; und &lt;y&gt; sind&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahl a ist 0,0.
Für jede Kommazahl z in liste, wenn z größer als, oder x ist und z kleiner als, oder y ist, erhöhe a um 1.
Gib a durch die Länge von liste zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Absolute_Häufigkeit</h2></summary>
<ul>
<pre>
Gibt die absolute Häufigkeit von x aus der gegebenen Liste zurück.
</pre>
	<li>Parameter: <code>liste</code>, <code>x</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die absolute Häufigkeit von &lt;x&gt; in &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl anzahl ist 0.
Für jede Kommazahl z in liste, wenn z gleich x ist, erhöhe anzahl um 1.
Gib anzahl zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Relative_Häufigkeit</h2></summary>
<ul>
<pre>
Gibt die absolute Häufigkeit von x aus der gegebenen Liste zurück.
</pre>
	<li>Parameter: <code>liste</code>, <code>x</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die relative Häufigkeit von &lt;x&gt; in &lt;liste&gt;&#34;</code></li>
	<li><code>&#34;wie viel Prozent der Zahlen aus &lt;liste&gt; gleich &lt;x&gt; sind&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib die absolute Häufigkeit von x in liste durch die Länge von liste zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Summe</h2></summary>
<ul>
<pre>
Gibt die Summe aller Zahlen der gegebenen Liste zurück.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Summe aller zahlen aus &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahl summe ist 0,0.
Für jede Kommazahl z in liste, erhöhe summe um z.
Gib summe zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Mittelwert</h2></summary>
<ul>
<pre>
Gibt den Mittelwert (arithmetisches Mittel) der gegebenen Liste zurück.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Mittelwert von &lt;liste&gt;&#34;</code></li>
	<li><code>&#34;das arithmetische Mittel von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib die Summe aller zahlen aus liste durch die Länge von liste zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Median</h2></summary>
<ul>
<pre>
Es muss eine sortierte Liste übergeben werden!
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Median von &lt;liste&gt;&#34;</code></li>
	<li><code>&#34;der Zentralwert von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahl m ist (die Länge von liste) durch 2.
Wenn (die Länge von liste) eine gerade Zahl ist, dann:
	Gib (liste an der Stelle (m nach unten gerundet als Zahl) plus liste an der Stelle (m nach oben gerundet als Zahl)) durch 2 zurück.
Gib liste an der Stelle (m als Zahl plus 1) zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Modalwert</h2></summary>
<ul>
<pre>
Gibt eine Liste der am häufigsten auftretenden Kommazahlen aus der gegebenen Liste zurück.
Gibt eine leere Liste zurück falls die gegebene Liste leer ist.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahlen Liste</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Modalwert von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahlen Liste modalwerte ist eine leere Kommazahlen Liste.
Die Zahl maxAbs ist 0.
Für jede Kommazahl z in liste, mache:
	Speichere die größere Zahl von (die absolute Häufigkeit von z in liste) und maxAbs in maxAbs.

Für jede Kommazahl z in liste, mache:
	Wenn nicht modalwerte z enthält und die absolute Häufigkeit von z in liste gleich maxAbs ist, füge z an modalwerte an.
Gib modalwerte zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Quantil</h2></summary>
<ul>
<pre>
Gibt das p-Quantil der gegebenen Liste zurück.
Es gibt an über welchen Wert p% der Daten befinden. 

Es muss eine sortierte Liste übergeben werden!
</pre>
	<li>Parameter: <code>liste</code>, <code>p</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Kommazahl</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;das &lt;p&gt;-Quantil von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahl np ist die Länge von liste mal p.
Wenn np eine ganze Zahl ist, dann:
	Gib (liste an der Stelle np als Zahl plus liste an der Stelle (np als Zahl plus 1)) durch 2 zurück.
Gib liste an der Stelle (np nach oben gerundet) als Zahl zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Varianz</h2></summary>
<ul>
<pre>
Berechnet die Varianz der gegebenen Liste.
Es gibt an wie sehr die Listenwerte um ihren Mittelwert streuen.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Varianz von &lt;liste&gt;&#34;</code></li>
	<li><code>&#34;der Varianz von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Kommazahl x ist 0,0.
Die Kommazahl m ist der Mittelwert von liste.
Für jede Kommazahl z in liste, erhöhe x um (z minus m) mal (z minus m).
Gib x durch (die Länge von liste minus 1) zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Standardabweichung</h2></summary>
<ul>
<pre>
Berechnet die Standardabweichung (Quadratwurzel der Varianz) der Liste.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Standardabweichung von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib die 2. Wurzel von der Varianz von liste zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Spannweite</h2></summary>
<ul>
<pre>
Gibt die Differenz des höchsten und niedristen Wertes zurück.
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die Spannweite von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib der höchste Wert aus liste minus der kleinste Wert aus liste zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Interquartilabstand</h2></summary>
<ul>
<pre>
Gibt die Differenz des 0,75-Quantils und 0,25-Quantils zurück.
Es muss eine sortierte Liste übergeben werden!
</pre>
	<li>Parameter: <code>liste</code></li>
	<li>Parameter Typ: <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Interquartilabstand von &lt;liste&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib das 0,75-Quantil von liste minus das 0,25-Quantil von liste zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Kovarianz</h2></summary>
<ul>
<pre>
Berechnet die Kovarianz zweier gleich langen Listen.
Beide Listen MÜSSEN gleich lang sein, sonst wird ein Laufzeitfehler geworfen.
</pre>
	<li>Parameter: <code>liste1</code>, <code>liste2</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;die empirische Kovarianz von &lt;liste1&gt; und &lt;liste2&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von liste1 ungleich die Länge von liste2 ist, löse einen Laufzeitfehler mit der Nachricht "Beide Listen müssen die gleiche Länge haben" und dem Code 1 aus.

Die Kommazahl x ist 0,0.
Die Kommazahl m1 ist der Mittelwert von liste1.
Die Kommazahl m2 ist der Mittelwert von liste2.
Für jede Zahl i von 1 bis die Länge von liste1, mache:
	Erhöhe x um (liste1 an der Stelle i minus m1) mal (liste2 an der Stelle i minus m2).
Gib x durch (die Länge von liste1 minus 1) zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Korrelationskoeffizient</h2></summary>
<ul>
<pre>
Berechnet den Korrelationskoeffizienten zweier gleichlangen Listen.
Beide Listen MÜSSEN gleich lang sein, sonst wird ein Laufzeitfehler geworfen.
</pre>
	<li>Parameter: <code>liste1</code>, <code>liste2</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der empirische Korrelationskoeffizient von &lt;liste1&gt; und &lt;liste2&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib (die empirische Kovarianz von liste1 und liste2) durch (die Standardabweichung von liste1 mal die Standardabweichung von liste2) zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Bestimmtheitsmaß</h2></summary>
<ul>
<pre>
Berechnet den Bestimmtheitsmaß zweier gleichlangen Listen (Quadrat des Korrelationskoeffizienten).
Beide Listen MÜSSEN gleich lang sein, sonst wird ein Laufzeitfehler geworfen.
</pre>
	<li>Parameter: <code>liste1</code>, <code>liste2</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Kommazahlen Liste</code></li>
	<li>Rückgabe Typ: <code>Kommazahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Bestimmtheitsmaß von &lt;liste1&gt; und &lt;liste2&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib der empirische Korrelationskoeffizient von liste1 und liste2 hoch 2 zurück.

</code>
</pre>
</details>


