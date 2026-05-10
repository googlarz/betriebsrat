# betriebsrat

**Know your rights at work — for employees and works council members.**

> **[🇩🇪 Deutsche Version weiter unten](#deutsch)**

German labour law is full of rights most people don't know they have. This tool makes them instantly accessible — in plain language, in German or English, without needing to know any law yourself.

---

## What can I ask?

Just describe your situation. Here are examples of what the tool considers when answering:

**As an employee:**

> "Can I get severance? I've worked here 7 years and earn €3,500 a month."

*Considers: years of service and salary against the Munich formula, whether a Sozialplan exists or was bypassed, company size threshold.*

> "I received a dismissal notice. Was the works council properly consulted?"

*Considers: whether a hearing notice was issued at all, completeness of social data and stated reason, applicable deadline (§ 102 BetrVG), consequence of missing or defective consultation.*

> "I'm pregnant and my employer wants to end my contract."

*Considers: Mutterschutzgesetz protection period, 2-week retroactive notification rule, whether the employer knew about the pregnancy, SGB IX if a disability is also involved.*

> "My employer introduced monitoring software. Did the works council have to agree?"

*Considers: whether the software can track performance or behaviour — directly or via logs — triggering § 87 Abs. 1 Nr. 6 BetrVG, whether a Betriebsvereinbarung was negotiated, remedy if it wasn't.*

> "I was transferred to a different location without my agreement. Is that allowed?"

*Considers: whether the BR was asked for consent (§ 99 BetrVG), how significantly the working conditions change, valid grounds for refusal, and what you can do if consent was skipped.*

---

**As a works council member:**

> "We received a dismissal hearing notice. What must we do and by when?"

*Considers: type of dismissal (ordinary vs. immediate) → 1-week or 3-day deadline, completeness of social data and reason, grounds for objection under § 102 Abs. 3, effect of objection on continued employment.*

> "The employer wants to introduce an AI writing assistant. Do we have co-determination rights?"

*Considers: whether the tool can monitor output or behaviour even indirectly (§ 87 Nr. 6), DSGVO Art. 35 data protection impact assessment, what a Betriebsvereinbarung must cover, whether rollout can be blocked.*

> "We're facing a mass layoff of 20 people. What's our role?"

*Considers: thresholds under § 17 KSchG, 30-day consultation requirement before the employer files with the employment agency, Sozialplan calculation using Munich formula, Einigungsstelle route if employer refuses.*

> "Can we force a Sozialplan if the employer refuses to negotiate?"

*Considers: erzwingbare Mitbestimmung under § 112 Abs. 4 BetrVG, how to call the Einigungsstelle, what it can and cannot decide, timeline.*

---

## How to get it

Tell Claude to install it:

> Install this: https://github.com/googlarz/betriebsrat

Claude will handle the rest. Then just describe your situation and ask away.

---

## Why trust it?

The knowledge base is built on content from [betriebsrat.de](https://www.betriebsrat.de) — the recognised German authority on works council practice — combined with the complete text of the *Betriebsverfassungsgesetz*. Every answer cites the specific paragraph that applies, and links directly back to the relevant topic on betriebsrat.de so you can verify it yourself.

It reasons from that structured knowledge base — not free-form AI guesswork. If something falls outside what the law covers, it says so.

The disclaimer at the bottom is there for a reason: this tool gives you the knowledge to understand your situation and know when you need a lawyer — not to replace one.

---

## What it knows

- Every paragraph of the *Betriebsverfassungsgesetz* (BetrVG) — the German works constitution law
- Co-determination rights: when the works council can block, must consent, or is only informed
- Legal deadlines — including the critical **3-week window** to challenge a dismissal in court
- Sozialplan and severance calculations (Munich formula)
- AGG (discrimination), Mutterschutz (maternity), Elternzeit (parental leave), SGB IX (disability)
- Procedure violations — what happens if the employer didn't follow the rules

---

## License

Apache 2.0 — see [LICENSE](LICENSE).

*This tool provides legal orientation, not legal advice. For your specific situation, consult a labour law specialist or your trade union.*

---

<details>
<summary>For developers — CLI reference</summary>

### Install (CLI)

```bash
go install github.com/googlarz/betriebsrat/cmd/betriebsrat@latest
betriebsrat doctor
```

### Key commands

```bash
betriebsrat ask "Ich wurde entlassen. Was nun?"
betriebsrat rights-check "KI-Überwachungssystem einführen" --agent
betriebsrat decide "Massenentlassung 20 Personen" --json
betriebsrat sozialplan-calc --salary 3500 --years 7
betriebsrat nachteilsausgleich --salary 3500 --years 7 --no-ia-attempted
betriebsrat check-anhoerung "<text des Anhörungsschreibens>"
betriebsrat deadline kuendigung --from 2026-05-10
betriebsrat law 102
betriebsrat bv-template homeoffice
betriebsrat ki-check "Teams-Analysefunktion zur Produktivitätsmessung"
betriebsrat massenentlassung --employees 120 --affected 25
```

### MCP server

```bash
# Claude Code
claude mcp add betriebsrat betriebsrat-pp-mcp

# Claude Desktop
{
  "mcpServers": {
    "betriebsrat": { "command": "betriebsrat-pp-mcp" }
  }
}
```

### Output flags

`--json` · `--agent` · `--lang en` · `--lang de`

</details>

---

---

<a name="deutsch"></a>

# betriebsrat — Deutsche Dokumentation

**Arbeitsrechte kennen — für Arbeitnehmer und Betriebsratsmitglieder.**

Das deutsche Arbeitsrecht steckt voller Rechte, von denen die meisten Menschen nichts wissen. Dieses Tool macht sie sofort zugänglich — in einfacher Sprache, auf Deutsch oder Englisch, ohne dass Sie selbst Jura studiert haben müssen.

---

## Was kann ich fragen?

Beschreiben Sie einfach Ihre Situation. Hier sind Beispiele dafür, was das Tool bei der Antwort berücksichtigt:

**Als Arbeitnehmer:**

> „Habe ich Anspruch auf eine Abfindung? Ich arbeite hier seit 7 Jahren und verdiene 3.500 € monatlich."

*Berücksichtigt: Betriebszugehörigkeit und Gehalt gegen die Münchner Formel, ob ein Sozialplan besteht oder übergangen wurde, Betriebsgröße.*

> „Ich habe eine Kündigung erhalten. Wurde der Betriebsrat ordnungsgemäß angehört?"

*Berücksichtigt: ob überhaupt ein Anhörungsschreiben ergangen ist, Vollständigkeit von Sozialdaten und Kündigungsgrund, einzuhaltende Frist (§ 102 BetrVG), Folgen einer fehlenden oder fehlerhaften Anhörung.*

> „Ich bin schwanger und mein Arbeitgeber will meinen Vertrag beenden."

*Berücksichtigt: Schutzfrist nach dem Mutterschutzgesetz, 2-Wochen-Regel zur rückwirkenden Mitteilung, ob der Arbeitgeber Kenntnis hatte, zusätzlicher Schutz nach SGB IX bei gleichzeitiger Behinderung.*

> „Mein Arbeitgeber hat eine Überwachungssoftware eingeführt. Musste der Betriebsrat zustimmen?"

*Berücksichtigt: ob die Software Leistung oder Verhalten erfassen kann — auch mittelbar über Logs — (§ 87 Abs. 1 Nr. 6 BetrVG), ob eine Betriebsvereinbarung ausgehandelt wurde, mögliche Rechtsfolgen bei fehlender Einigung.*

> „Ich wurde ohne mein Einverständnis versetzt. Ist das erlaubt?"

*Berücksichtigt: ob der BR um Zustimmung gebeten wurde (§ 99 BetrVG), wie stark sich die Arbeitsbedingungen ändern, Verweigerungsgründe und Handlungsmöglichkeiten bei übergangener Beteiligung.*

---

**Als Betriebsratsmitglied:**

> „Wir haben ein Anhörungsschreiben für eine Kündigung erhalten. Was müssen wir tun und bis wann?"

*Berücksichtigt: Kündigungsart (ordentlich vs. fristlos) → Frist 1 Woche oder 3 Tage, Vollständigkeit von Sozialdaten und Grund, Widerspruchsgründe nach § 102 Abs. 3, Wirkung eines Widerspruchs auf Weiterbeschäftigung.*

> „Der Arbeitgeber will einen KI-Schreibassistenten einführen. Haben wir ein Mitbestimmungsrecht?"

*Berücksichtigt: ob das Tool Leistung oder Verhalten auch mittelbar erfassen kann (§ 87 Nr. 6), Datenschutz-Folgenabschätzung nach DSGVO Art. 35, notwendige Inhalte einer Betriebsvereinbarung, ob die Einführung blockiert werden kann.*

> „Uns droht eine Massenentlassung von 20 Personen. Was ist unsere Rolle?"

*Berücksichtigt: Schwellenwerte nach § 17 KSchG, 30-tägige Konsultationspflicht vor der Anzeige bei der Agentur für Arbeit, Sozialplanberechnung nach der Münchner Formel, Einigungsstellenverfahren bei Verweigerung des Arbeitgebers.*

> „Können wir einen Sozialplan erzwingen, wenn der Arbeitgeber nicht verhandeln will?"

*Berücksichtigt: erzwingbare Mitbestimmung nach § 112 Abs. 4 BetrVG, Einleitung des Einigungsstellenverfahrens, Entscheidungsbefugnis der Einigungsstelle, Zeitrahmen.*

---

## Wie bekomme ich das Tool?

Sagen Sie Claude, es soll installieren:

> Install this: https://github.com/googlarz/betriebsrat

Claude übernimmt den Rest. Dann beschreiben Sie einfach Ihre Situation.

---

## Warum kann ich dem Tool vertrauen?

Die Wissensdatenbank basiert auf Inhalten von [betriebsrat.de](https://www.betriebsrat.de) — der anerkannten deutschen Referenz für Betriebsratsrecht — ergänzt um den vollständigen Text des Betriebsverfassungsgesetzes. Jede Antwort nennt den konkreten Paragrafen und verweist direkt auf das entsprechende Thema auf betriebsrat.de, damit Sie die Grundlage selbst nachprüfen können.

Das Tool arbeitet auf Basis dieser strukturierten Wissensdatenbank — kein freies KI-Raten. Wenn etwas außerhalb des gesetzlichen Rahmens liegt, wird das klar gesagt.

Der Hinweis am Ende hat seinen Grund: Dieses Tool gibt Ihnen das Wissen, um Ihre Situation zu verstehen und zu erkennen, wann Sie einen Anwalt brauchen — es ersetzt ihn nicht.

---

## Was das Tool weiß

- Alle Paragrafen des Betriebsverfassungsgesetzes (BetrVG)
- Mitbestimmungsrechte: wann der Betriebsrat blockieren kann, zustimmen muss oder nur informiert wird
- Gesetzliche Fristen — einschließlich der kritischen **3-Wochen-Frist** für die Kündigungsschutzklage
- Sozialplan- und Abfindungsberechnungen (Münchner Formel)
- AGG (Diskriminierungsschutz), Mutterschutz, Elternzeit (BEEG), SGB IX (Schwerbehinderung)
- Verfahrensfehler — was passiert, wenn der Arbeitgeber die Regeln nicht eingehalten hat

---

## Lizenz

Apache 2.0 — siehe [LICENSE](LICENSE).

*Dieses Tool bietet rechtliche Orientierung, keine Rechtsberatung. Für Ihren konkreten Fall wenden Sie sich an einen Fachanwalt für Arbeitsrecht oder Ihre Gewerkschaft.*

---

<details>
<summary>Für Entwickler — CLI-Referenz</summary>

### Installation (CLI)

```bash
go install github.com/googlarz/betriebsrat/cmd/betriebsrat@latest
betriebsrat doctor
```

### Wichtige Befehle

```bash
betriebsrat ask "Ich wurde entlassen. Was nun?"
betriebsrat rights-check "KI-Überwachungssystem einführen" --agent
betriebsrat decide "Massenentlassung 20 Personen" --json
betriebsrat sozialplan-calc --salary 3500 --years 7
betriebsrat nachteilsausgleich --salary 3500 --years 7 --no-ia-attempted
betriebsrat check-anhoerung "<text des Anhörungsschreibens>"
betriebsrat deadline kuendigung --from 2026-05-10
betriebsrat law 102
betriebsrat bv-template homeoffice
betriebsrat ki-check "Teams-Analysefunktion zur Produktivitätsmessung"
betriebsrat massenentlassung --employees 120 --affected 25
```

### Ausgabeformate

`--json` · `--agent` · `--lang en` · `--lang de`

</details>
