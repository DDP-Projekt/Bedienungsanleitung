# Duden/Zeichen Funktionen
<details>
<summary><h2>IstLeer</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn der Buchstabe b ein Leerzeichen (' '), eine neue Zeile ('\n'), ein Tabulator ('\t') oder ein Wagenrücklauf ('\r') ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; ein leeres Zeichen ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib b gleich ' ' ist oder b gleich '\n' ist oder b gleich '\t' ist oder b gleich '\r' ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstGroß</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn der Buchstabe b groß ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; ein großer Buchstabe ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl z ist b als Zahl.
Der Boolean erg ist 
	(z größer als, oder 65 ist und z kleiner als, oder 90 ist) oder
	(z größer als, oder 192 ist und z kleiner als, oder 214 ist) oder
	(z größer als, oder 216 ist und z kleiner als, oder 222 ist).
	[... es gibt noch viel mehr ...]
Gib erg zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstKlein</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn der Buchstabe b klein ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; ein kleiner Buchstabe ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl z ist b als Zahl.
Der Boolean erg ist 
	(z größer als, oder 96 ist und z kleiner als, oder 122 ist) oder
	(z größer als, oder 223 ist und z kleiner als, oder 246 ist) oder
	(z größer als, oder 248 ist und z kleiner als, oder 255 ist).
	[... es gibt noch viel mehr ...]
Gib erg zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstLeerzeichen</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn der Buchstabe b ein Leerzeichen (' ') ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; ein Leerzeichen ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib b gleich ' ' ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstZahl</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn der Buchstabe b eine Zahl (Code 48-57) ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; eine Zahl ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib b als Zahl größer als, oder 48 ist und b als Zahl kleiner als, oder 57 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstKontroll</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn der Buchstabe b ein Kontrollzeichen (Code 0-31) ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; ein Kontrollzeichen ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib b als Zahl größer als, oder 0 ist und b als Zahl kleiner als, oder 31 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstDeutscherBuchstabe</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn der Buchstabe b ein deutscher Buchstabe (a-Z, äöü, ÄÖÜ und ß) ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; ein deutscher Buchstabe ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl z ist b als Zahl.
Gib 
	(z größer als, oder 65 ist und z kleiner als, oder 90 ist) [a-z] oder
	(z größer als, oder 97 ist und z kleiner als, oder 122 ist) [A-Z] oder
	(z gleich 196 ist oder z gleich 228 ist) [Ää] oder 
	(z gleich 214 ist oder z gleich 246 ist) [Öö] oder
	(z gleich 220 ist oder z gleich 252 ist) [Üü] oder
	(z gleich 223 ist) [ß]
	zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstDeutscherBuchstabeOderZahl</h2></summary>
<ul>
<pre>
Gibt wahr zurück wenn der Buchstabe b ein deutscher Buchstabe oder eine Zahl ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; ein deutscher Buchstabe oder eine Zahl ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib b ein deutscher Buchstabe ist oder b eine Zahl ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Großgeschrieben</h2></summary>
<ul>
<pre>
Gibt den gegebenen Buchstaben als großgeschriebe Variante zurück. 
Gibt den selben Buchstaben zurück wenn es schon großgeschrieben ist oder kein deutscher Buchstabe (siehe IstDeutscherBuchstabe) ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Buchstabe</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; als großer Buchstabe&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn nicht b ein deutscher Buchstabe ist, dann:
	Gib b zurück.

Gib (b als Zahl logisch und 223) als Buchstabe zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Kleingeschrieben</h2></summary>
<ul>
<pre>
Gibt den gegebenen Buchstaben als kleingeschriebene Variante zurück. 
Gibt den selben Buchstaben zurück wenn es schon kleingeschrieben ist oder kein deutscher Buchstabe (siehe IstDeutscherBuchstabe) ist.
</pre>
	<li>Parameter: <code>b</code></li>
	<li>Parameter Typ: <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Buchstabe</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;b&gt; als kleiner Buchstabe&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn nicht b ein deutscher Buchstabe ist oder b ein kleiner Buchstabe ist, dann:
	Gib b zurück.

Gib (b als Zahl plus 32) als Buchstabe zurück.

</code>
</pre>
</details>

<details>
<summary><h2>ASCII_Zeichen</h2></summary>
<ul>
<pre>
Gibt den Zeichen mit der gegebenen ASCII Nummer zurück.
</pre>
	<li>Parameter: <code>id</code></li>
	<li>Parameter Typ: <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Buchstabe</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der ASCII Zeichen mit der Nummer &lt;id&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib id als Buchstabe zurück.

</code>
</pre>
</details>


