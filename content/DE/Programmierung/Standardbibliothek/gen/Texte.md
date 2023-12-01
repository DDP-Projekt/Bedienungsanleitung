+++
title = "Texte"
weight = 1
+++
# Duden/Texte Funktionen
<details>
<summary><h2>LeererText</h2></summary>
<ul>
<pre>
Gibt "" zurück.
Nutzen: Der Text t ist ein leerer Text.
</pre>
</li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;ein leerer Text&#34;</code></li>
	<li><code>&#34;einen leeren Text&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib "" zurück.

</code>
</pre>
</details>

<details>
<summary><h2>TrimAnfang</h2></summary>
<ul>
<pre>
Entfernt alle gegebenen Buchstaben vom Anfang eines gegebenen Textes.
z.B.: 
Der Text t ist "aaaaaaahallo"
Entferne alle 'a' vor t.
t: "hallo"
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Entferne alle &lt;zeichen&gt; vor &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von text gleich 0 ist, verlasse die Funktion.
Die Zahl index ist 1.
Solange (text an der Stelle index) gleich zeichen ist, mache:
	Erhöhe index um 1.
	Wenn index größer als die Länge von text ist, dann:
		Speichere "" in text.
		Verlasse die Funktion.
Speichere text von index bis (die Länge von text) in text.

</code>
</pre>
</details>

<details>
<summary><h2>TrimAnfangWert</h2></summary>
<ul>
<pre>
Siehe TrimAnfang
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; mit allen &lt;zeichen&gt; davor entfernt&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Entferne alle zeichen vor text.
Gib text zurück.

</code>
</pre>
</details>

<details>
<summary><h2>TrimEnde</h2></summary>
<ul>
<pre>
Entfernt alle gegebenen Buchstaben vom Ende eines gegebenen Textes.
z.B.: 
Der Text t ist "hallo!!!!!!!!!!"
Entferne alle '!' nach t.
t: "hallo"
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Entferne alle &lt;zeichen&gt; nach &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von text gleich 0 ist, verlasse die Funktion.

Die Zahl index ist die Länge von text.
Wenn index gleich 0 ist, verlasse die Funktion.
Solange (text an der Stelle index) gleich zeichen ist, mache:
	Verringere index um 1.
	Wenn index kleiner als 1 ist, dann:
		Speichere "" in text.
		Verlasse die Funktion.

Speichere text von 1 bis index in text.

</code>
</pre>
</details>

<details>
<summary><h2>TrimEndeWert</h2></summary>
<ul>
<pre>
Siehe TrimEnde
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; mit allen &lt;zeichen&gt; danach entfernt&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Entferne alle zeichen nach text.
Gib text zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Trim</h2></summary>
<ul>
<pre>
Entfernt alle gegebenen Buchstaben vom Anfang und Ende eines gegebenen Textes.
z.B.: 
Der Text t ist "!!!!!hallo!!!!!!!!!!"
Entferne alle '!' vor und nach t.
t: "hallo"
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Entferne alle &lt;zeichen&gt; vor und nach &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von text gleich 0 ist, verlasse die Funktion.

Die Zahl startIndex ist 1.
Die Zahl stopIndex ist die Länge von text.
Solange (text an der Stelle startIndex) gleich zeichen ist und startIndex kleiner als die Länge von text ist, erhöhe startIndex um 1.
Solange (text an der Stelle stopIndex) gleich zeichen ist und stopIndex ungleich 1 ist, verringere stopIndex um 1.
Wenn startIndex gleich die Länge von text ist und stopIndex gleich 1 ist, dann:
  	Speichere "" in text.
	Verlasse die Funktion.

Speichere text von startIndex bis stopIndex in text.

</code>
</pre>
</details>

<details>
<summary><h2>TrimWert</h2></summary>
<ul>
<pre>
Siehe Trim
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; mit allen &lt;zeichen&gt; davor und danach entfernt&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Entferne alle zeichen vor und nach text.
Gib text zurück.

</code>
</pre>
</details>

<details>
<summary><h2>EnthältBuchstabe</h2></summary>
<ul>
<pre>
Gibt zurück ob der gegebenen Text den gegebenen Buchstaben enthält.
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; &lt;zeichen&gt; enthält&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Für jeden Buchstaben b in text, wenn b gleich zeichen ist, gib wahr zurück.
Gib falsch zurück.

</code>
</pre>
</details>

<details>
<summary><h2>EnthältText</h2></summary>
<ul>
<pre>
Gibt zurück ob der gegebene Text (text) den Subtext (suchText) enthält.
</pre>
	<li>Parameter: <code>text</code>, <code>suchText</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; &lt;suchText&gt; enthält&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn text gleich suchText ist, gib wahr zurück.

Die Zahl startIndex ist 0.
Die Zahl endIndex ist die Länge von suchText.
Wenn endIndex gleich 0 ist, gib falsch zurück.

Solange endIndex kleiner als, oder die Länge von text ist, mache:
	Der Text subtext ist text von startIndex bis endIndex.

	Wenn subtext gleich suchText ist, gib wahr zurück.
	
	Speichere startIndex plus die Länge von suchText in endIndex.
	Erhöhe startIndex um 1.
Gib falsch zurück.

</code>
</pre>
</details>

<details>
<summary><h2>BeginntMitBuchstabe</h2></summary>
<ul>
<pre>
Gibt zurück ob der gegebene Text mit dem gegebenen Buchstaben anfängt.
</pre>
	<li>Parameter: <code>text</code>, <code>buchstabe</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;buchstabe&gt; am Anfang von &lt;text&gt; steht&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von text gleich 0 ist, gib falsch zurück.
Gib (text an der Stelle 1) gleich buchstabe ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>BeginntMitText</h2></summary>
<ul>
<pre>
Gibt zurück ob der gegebene Text mit dem gegebenen Text (suchText) anfängt.
</pre>
	<li>Parameter: <code>text</code>, <code>suchText</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;suchText&gt; am Anfang von &lt;text&gt; steht&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von text gleich 0 ist oder die Länge von suchText gleich 0 ist, gib falsch zurück.
Gib (text von 1 bis (die Länge von suchText)) gleich suchText ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>EndetMitBuchstabe</h2></summary>
<ul>
<pre>
Gibt zurück ob der gegebene Text mit dem gegebenen Buchstaben endet.
</pre>
	<li>Parameter: <code>text</code>, <code>buchstabe</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;buchstabe&gt; am Ende von &lt;text&gt; steht&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von text gleich 0 ist, gib falsch zurück.
Gib (text an der Stelle (die Länge von text)) gleich buchstabe ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>EndetMitText</h2></summary>
<ul>
<pre>
Gibt zurück ob der gegebene Text mit dem gegebenen Text (suchText) endet.
</pre>
	<li>Parameter: <code>text</code>, <code>suchText</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;suchText&gt; am Ende von &lt;text&gt; steht&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von text gleich 0 ist oder die Länge von suchText gleich 0 ist, gib falsch zurück.
Gib (text von die Länge von text minus die Länge von suchText plus 1 bis (die Länge von text)) gleich suchText ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Text_leeren</h2></summary>
<ul>
<pre>
Speichert einen leeren Text in text.
</pre>
	<li>Parameter: <code>text</code></li>
	<li>Parameter Typ: <code>Text Referenz</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Leere &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere "" in text.

</code>
</pre>
</details>

<details>
<summary><h2>Text_an_Text_fügen</h2></summary>
<ul>
<pre>
Fügt zwei Texte aneinander.
f("ha", "lo") -> "halo"
</pre>
	<li>Parameter: <code>text</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;text&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere text verkettet mit elm in text.

</code>
</pre>
</details>

<details>
<summary><h2>Buchstabe_an_Text_fügen</h2></summary>
<ul>
<pre>
Fügt einen Buchstaben an einen Text.
f("may", 'o') -> "mayo"
</pre>
	<li>Parameter: <code>text</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Füge &lt;elm&gt; an &lt;text&gt; an&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere text verkettet mit elm in text.

</code>
</pre>
</details>

<details>
<summary><h2>Text_in_Text_einfügen</h2></summary>
<ul>
<pre>
Fügt einen Text
</pre>
	<li>Parameter: <code>text</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Zahl</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere text von 1 bis (index minus 1) verkettet mit elm verkettet mit text von index bis (die Länge von text) in text.

</code>
</pre>
</details>

<details>
<summary><h2>Buchstabe_in_Text_einfügen</h2></summary>
<ul>
	<li>Parameter: <code>text</code>, <code>index</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Zahl</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Setze &lt;elm&gt; an die Stelle &lt;index&gt; von &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere text von 1 bis (index minus 1) verkettet mit elm verkettet mit text von index bis (die Länge von text) in text.

</code>
</pre>
</details>

<details>
<summary><h2>Text_vor_Text_stellen</h2></summary>
<ul>
<pre>
Fügt einen Text am Anfang eines Textes ein.
f("hallo", " welt!") -> "hallo welt!"
</pre>
	<li>Parameter: <code>text</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Text</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere elm verkettet mit text in text.

</code>
</pre>
</details>

<details>
<summary><h2>Buchstabe_vor_Text_stellen</h2></summary>
<ul>
<pre>
Fügt einen Buchstaben am Anfang eines Textes ein.
f("allo", 'h') -> "hallo"
</pre>
	<li>Parameter: <code>text</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Stelle &lt;elm&gt; vor &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere elm verkettet mit text in text.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Text</h2></summary>
<ul>
<pre>
Entfernt den Buchstaben an der Stelle index vom Text
</pre>
	<li>Parameter: <code>text</code>, <code>index</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche das Element an der Stelle &lt;index&gt; aus &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von text gleich 0 ist, verlasse die Funktion.

Wenn index gleich 1 ist und die Länge von text größer als 1 ist, dann:
	Speichere "" in text.
Wenn aber index gleich 1 ist, dann:
	Speichere text von 2 bis (die Länge von text) in text.
Wenn aber index gleich die Länge von text ist, dann:
	Speichere text von 1 bis (die Länge von text minus 1) in text.
Sonst:
	Speichere text von 1 bis (index minus 1) verkettet mit text von (index plus 1) bis (die Länge von text) in text.

</code>
</pre>
</details>

<details>
<summary><h2>Lösche_Text_Bereich</h2></summary>
<ul>
<pre>
Entfernt einen Bereich vom Text
</pre>
	<li>Parameter: <code>text</code>, <code>start</code>, <code>end</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Zahl</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Lösche alle Elemente von &lt;start&gt; bis &lt;end&gt; aus &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn start gleich 1 ist, dann:
	Speichere text von end plus 1 bis (die Länge von text) in text.
Sonst:
	Speichere text von 1 bis (start minus 1) verkettet mit text von (end plus 1) bis (die Länge von text) in text.

</code>
</pre>
</details>

<details>
<summary><h2>Fülle_Text</h2></summary>
<ul>
<pre>
Füllt den Text mit dem gegebenen Buchstaben
</pre>
	<li>Parameter: <code>text</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text Referenz</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Fülle &lt;text&gt; mit &lt;elm&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text neuerText ist "".
Wiederhole:
	Speichere neuerText verkettet mit elm in neuerText.
Die Länge von text Mal.

Speichere neuerText in text.

</code>
</pre>
</details>

<details>
<summary><h2>IndexVonText</h2></summary>
<ul>
<pre>
Gibt den index des gegebenen Buchstaben im Text zurück oder -1 falls es nicht gefunden wurde.
</pre>
	<li>Parameter: <code>text</code>, <code>elm</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Zahl</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;der Index von &lt;elm&gt; in &lt;text&gt;&#34;</code></li>
	<li><code>&#34;den Index von &lt;elm&gt; in &lt;text&gt;&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Wenn die Länge von text gleich 0 ist, gib -1 zurück.
Für jede Zahl i von 1 bis (die Länge von text), Wenn text an der Stelle i gleich elm ist, gib i zurück.
Gib -1 zurück.

</code>
</pre>
</details>

<details>
<summary><h2>IstTextLeer</h2></summary>
<ul>
<pre>
Gibt ob der gegebene Text leer ist zurück
</pre>
	<li>Parameter: <code>text</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; leer ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Gib [wahr wenn] die Länge von text gleich 0 ist zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Ist_Zahl</h2></summary>
<ul>
<pre>
Gibt zurück ob ein Text in eine Zahl umgewandelt werden kann
</pre>
	<li>Parameter: <code>t</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Boolean</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;t&gt; eine Zahl ist&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl l ist die Länge von t.
Wenn l kleiner als 1 ist, gib falsch zurück.

Der Buchstabe Vorzeichen ist t an der Stelle 1.
Wenn Vorzeichen eine Zahl ist, gib wahr zurück.
Wenn Vorzeichen ungleich '+' ist und Vorzeichen ungleich '-' ist, gib falsch zurück.
Wenn l kleiner als 2 ist oder nicht (t an der Stelle 2) eine Zahl ist, gib falsch zurück.
Gib wahr zurück.

</code>
</pre>
</details>

<details>
<summary><h2>GroßschreibenWert</h2></summary>
<ul>
<pre>
Wandelt jeden Buchstaben des gegebenen Textes in die groß geschriebene Variante
</pre>
	<li>Parameter: <code>text</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; groß geschrieben&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text neuerText ist "".
Für jeden Buchstaben b in text, mache:
	Füge (b als großer Buchstabe) an neuerText an.
Gib neuerText zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Großschreiben</h2></summary>
<ul>
<pre>
Wandelt jeden Buchstaben des gegebenen Textes in die groß geschriebene Variante
</pre>
	<li>Parameter: <code>text</code></li>
	<li>Parameter Typ: <code>Text Referenz</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe &lt;text&gt; groß&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere text groß geschrieben in text.

</code>
</pre>
</details>

<details>
<summary><h2>KleinschreibenWert</h2></summary>
<ul>
<pre>
Wandelt jeden Buchstaben des gegebenen Textes in die klein geschriebene Variante
</pre>
	<li>Parameter: <code>text</code></li>
	<li>Parameter Typ: <code>Text</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; klein geschrieben&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text neuerText ist "".
Für jeden Buchstaben b in text, mache:
	Füge (b als kleiner Buchstabe) an neuerText an.
Gib neuerText zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Kleinschreiben</h2></summary>
<ul>
<pre>
Wandelt jeden Buchstaben des gegebenen Textes in die klein geschriebene Variante
</pre>
	<li>Parameter: <code>text</code></li>
	<li>Parameter Typ: <code>Text Referenz</code></li>
	<li>Rückgabe Typ: <code>nichts</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;Schreibe &lt;text&gt; klein&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Speichere text klein geschrieben in text.

</code>
</pre>
</details>

<details>
<summary><h2>PolsterLinks</h2></summary>
<ul>
<pre>
z.B.:
f("hallo", ' ', 8) -> "   hallo"
f("hey", ' ', 8) -> "     hey"
f("programm", ' ', 8) -> "programm"
f("", 'o', 8) -> "oooooooo"
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code>, <code>endlänge</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; mit &lt;endlänge&gt; &lt;zeichen&gt; links gepolstert&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl länge ist die Länge von text.
Die Zahl gesuchteLänge ist endlänge minus länge.
Wenn gesuchteLänge kleiner als, oder 0 ist, dann:
	Gib text zurück.

Wiederhole:
	Stelle zeichen vor text.
gesuchteLänge Mal.

Gib text zurück.

</code>
</pre>
</details>

<details>
<summary><h2>PolsterRechts</h2></summary>
<ul>
<pre>
z.B.:
f("hallo", ' ', 8) -> "hallo   "
f("hey", ' ', 8) -> "hey     "
f("programm", ' ', 8) -> "programm"
f("", 'o', 8) -> "oooooooo"
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code>, <code>endlänge</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code>, <code>Zahl</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; mit &lt;endlänge&gt; &lt;zeichen&gt; rechts gepolstert&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Zahl länge ist die Länge von text.
Die Zahl gesuchteLänge ist endlänge minus länge.
Wenn gesuchteLänge kleiner als, oder 0 ist, dann:
	Gib text zurück.

Wiederhole:
	Füge zeichen an text an.
gesuchteLänge Mal.

Gib text zurück.

</code>
</pre>
</details>

<details>
<summary><h2>Spalte</h2></summary>
<ul>
<pre>
Spaltet den gegebenen Text anhand des angegebenen Buchstaben in Teiltexte.
</pre>
	<li>Parameter: <code>text</code>, <code>zeichen</code></li>
	<li>Parameter Typen: <code>Text</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Text Liste</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;text&gt; an &lt;zeichen&gt; gespalten&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Die Text Liste endliste ist eine leere Text Liste.
Die Zahl endIndex ist der Index von zeichen in text.
Solange endIndex ungleich -1 ist und endIndex ungleich die Länge von text ist, mache:
	Wenn endIndex ungleich 1 ist, dann:
		Speichere endliste verkettet mit text von 1 bis (endIndex minus 1) in endliste.
		Speichere text von endIndex plus 1 bis (die Länge von text) in text.
	Sonst:
		Speichere endliste verkettet mit "" in endliste.
		Speichere text von 2 bis (die Länge von text) in text.
	Speichere der Index von zeichen in text in endIndex.

Speichere endliste verkettet mit text in endliste.
Gib endliste zurück.

</code>
</pre>
</details>

<details>
<summary><h2>VerbindenT</h2></summary>
<ul>
<pre>
Verkettet alle Elemente der Liste mit dem Trennzeichen und gibt den Text zurück.
z.B.:
f(["hi", "", "yo"], '-') -> "hi--yo"
</pre>
	<li>Parameter: <code>liste</code>, <code>trennzeichen</code></li>
	<li>Parameter Typen: <code>Text Liste</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; mit dem Trennzeichen &lt;trennzeichen&gt; zum Text verbunden&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text ret ist ein leerer Text.
Für jede Zahl i von 1 bis die Länge von liste, mache:
	Wenn i kleiner als die Länge von liste ist, Speichere ret verkettet mit liste an der Stelle i verkettet mit trennzeichen in ret.
	Sonst Speichere ret verkettet mit liste an der Stelle i in ret.
Gib ret zurück.

</code>
</pre>
</details>

<details>
<summary><h2>VerbindenZ</h2></summary>
<ul>
<pre>
Verkettet alle Elemente der Liste mit dem Trennzeichen und gibt den Text zurück.
z.B.:
f([1, 234, 56789, 0], '-') -> "1-234-56789-0"
</pre>
	<li>Parameter: <code>liste</code>, <code>trennzeichen</code></li>
	<li>Parameter Typen: <code>Zahlen Liste</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; mit dem Trennzeichen &lt;trennzeichen&gt; zum Text verbunden&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text ret ist ein leerer Text.
Für jede Zahl i von 1 bis die Länge von liste, mache:
	Wenn i kleiner als die Länge von liste ist, Speichere ret verkettet mit (liste an der Stelle i) als Text verkettet mit trennzeichen in ret.
	Sonst Speichere ret verkettet mit (liste an der Stelle i) als Text in ret.
Gib ret zurück.

</code>
</pre>
</details>

<details>
<summary><h2>VerbindenK</h2></summary>
<ul>
<pre>
Verkettet alle Elemente der Liste mit dem Trennzeichen und gibt den Text zurück.
z.B.:
f([1,4, 0 durch 0, 23,0], '-') -> "1,4-nan-23"
</pre>
	<li>Parameter: <code>liste</code>, <code>trennzeichen</code></li>
	<li>Parameter Typen: <code>Kommazahlen Liste</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; mit dem Trennzeichen &lt;trennzeichen&gt; zum Text verbunden&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text ret ist ein leerer Text.
Für jede Zahl i von 1 bis die Länge von liste, mache:
	Wenn i kleiner als die Länge von liste ist, Speichere ret verkettet mit (liste an der Stelle i) als Text verkettet mit trennzeichen in ret.
	Sonst Speichere ret verkettet mit (liste an der Stelle i) als Text in ret.
Gib ret zurück.

</code>
</pre>
</details>

<details>
<summary><h2>VerbindenB</h2></summary>
<ul>
<pre>
Verkettet alle Elemente der Liste mit dem Trennzeichen und gibt den Text zurück.
z.B.:
f(['a', 'b', 'c'], '-') -> "a-b-c"
</pre>
	<li>Parameter: <code>liste</code>, <code>trennzeichen</code></li>
	<li>Parameter Typen: <code>Buchstaben Liste</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; mit dem Trennzeichen &lt;trennzeichen&gt; zum Text verbunden&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text ret ist ein leerer Text.
Für jede Zahl i von 1 bis die Länge von liste, mache:
	Wenn i kleiner als die Länge von liste ist, Speichere ret verkettet mit (liste an der Stelle i) verkettet mit trennzeichen in ret.
	Sonst Speichere ret verkettet mit (liste an der Stelle i) in ret.
Gib ret zurück.

</code>
</pre>
</details>

<details>
<summary><h2>VerbindenBool</h2></summary>
<ul>
<pre>
Verkettet alle Elemente der Liste mit dem Trennzeichen und gibt den Text zurück.
z.B.:
f([wahr, falsch, falsch], '-') -> "wahr-falsch-falsch"
</pre>
	<li>Parameter: <code>liste</code>, <code>trennzeichen</code></li>
	<li>Parameter Typen: <code>Boolean Liste</code>, <code>Buchstabe</code></li>
	<li>Rückgabe Typ: <code>Text</code></li>
</ul>

<h3>Aliase</h3>
<ol>
	<li><code>&#34;&lt;liste&gt; mit dem Trennzeichen &lt;trennzeichen&gt; zum Text verbunden&#34;</code></li>
</ol>

<h3>Implementation</h3>
<pre class="language-ddp" tabindex="0">
<code class="language-ddp">
Der Text ret ist ein leerer Text.
Für jede Zahl i von 1 bis die Länge von liste, mache:
	Wenn i kleiner als die Länge von liste ist, Speichere ret verkettet mit (liste an der Stelle i) als Text verkettet mit trennzeichen in ret.
	Sonst Speichere ret verkettet mit (liste an der Stelle i) als Text in ret.
Gib ret zurück.

</code>
</pre>
</details>


