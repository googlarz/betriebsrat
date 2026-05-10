---
name: betriebsrat
description: "German works council advisor: rights-check, deadlines, BetrVG decisions, and situation-specific playbooks — offline, in your terminal. Trigger phrases: `check Betriebsrat rights`, `does BR have co-determination`, `BetrVG paragraph`, `Betriebsrat deadline`, `works council rights Germany`, `Kündigung Betriebsrat`, `Betriebsänderung`, `Software Einführung Betriebsrat`, `use betriebsrat`, `run betriebsrat`."
author: "Dawid Piaskowski"
license: "Apache-2.0"
argument-hint: "<command> [args] | install cli|mcp"
allowed-tools: "Read Bash"
metadata:
  openclaw:
    requires:
      bins:
        - betriebsrat-pp-cli
---

# Betriebsrat — Printing Press CLI

Full knowledge base of betriebsrat.de — offline, in your terminal — with decision support that no website offers: instant answers to "do we have a say in this?", legal deadline calculation, situation checklists, and plain-language BetrVG paragraph explanations. Built for works council members who need answers in minutes, not hours of reading.

## Prerequisites: Install the CLI

This skill drives the `betriebsrat-pp-cli` binary. **You must verify the CLI is installed before invoking any command from this skill.** If it is missing, install it first:

1. Install via the Printing Press installer:
   ```bash
   npx -y @mvanhorn/printing-press install betriebsrat --cli-only
   ```
2. Verify: `betriebsrat-pp-cli --version`
3. Ensure `$GOPATH/bin` (or `$HOME/go/bin`) is on `$PATH`.

Run `betriebsrat-pp-cli doctor` to verify setup.

---

## Auto-Session Protocol (Always Follow This)

**When this skill is activated with any situation described, do the following immediately — without waiting to be asked:**

### A — Auto-classify the situation

Run all three classification commands in parallel before saying anything.
Detect the user's language from their message and add `--lang en` if they're writing in English:

```bash
betriebsrat-pp-cli rights-check "<situation>" --agent [--lang en]
betriebsrat-pp-cli decide "<situation>" --agent [--lang en]
betriebsrat-pp-cli consequences "<situation_type>" --agent [--lang en]  # if situation type is clear
```

Then present findings: applicable §§, co-determination type, key deadlines, and what happens if the BR misses the window.

### A2 — Proactive command chaining

After running auto-classification, automatically chain follow-up commands when the situation triggers them — without waiting to be asked:

| If auto-classification shows... | Also run automatically |
|---------------------------------|------------------------|
| `ki-check` finds co-determination triggered | `auskunft --topic ki` — draft the information request letter |
| `widerspruch-check` finds strong grounds | `letter kündigung --type widerspruch` + `protokoll` — draft both documents |
| `massenentlassung` threshold triggered | `sozialplan-calc` (ask for salary/years) + `deadline` for § 17 KSchG |
| `nachteilsausgleich` claim > 0 | `sozialplan-calc` (same inputs) — compare claim vs. Sozialplan |
| `check-anhoerung` finds incomplete Anhörung | `consequences kündigung` — explain implications of the clock not running |
| `decide` returns MitbestimmungErzwingbar | `bv-template <topic>` — offer to generate the BV skeleton |

Present chained results together with the primary result, clearly labelled.

### B — Load company profile (if set)

```bash
betriebsrat-pp-cli context show --agent
```

If a profile exists, use it to filter advice: skip §§ that don't apply (e.g., § 111 only applies at ≥20 AN; § 106 only at ≥100 AN). If the person is a BR member, flag § 103 automatically.

### C — Ask company profile questions (if profile is missing or incomplete)

After auto-classification, ask these questions **once per session** to calibrate all subsequent advice. Do not ask them again after the user answers:

1. **Wie viele Arbeitnehmer hat der Betrieb?** (Bestimmt welche BetrVG-§§ gelten: <20 → kein § 111, <100 → kein Wirtschaftsausschuss, ≥200 → Vollfreistellung)
2. **Gibt es einen Tarifvertrag? Wenn ja, welchen?** (Tarifvorbehalt schränkt Betriebsvereinbarungen ein)
3. **Ist die betroffene Person ein Betriebsratsmitglied?** (Ja → § 103 statt § 102; Zustimmung statt Anhörung erforderlich)
4. **Gibt es bereits eine Betriebsvereinbarung zu diesem Thema?** (Bestehende BV kann helfen oder eine neue blockieren)
5. **Wie groß ist der Betriebsrat?** (Quorum-Regeln für BR-Beschlüsse: mindestens 50% anwesend)

Save the answers to the profile for this session:
```bash
betriebsrat-pp-cli context set --employees <n> --tariff [true/false] --br-size <n> --bvs "<topic>" 
```

**Do not ask for salary, names, or personal data unless the user explicitly volunteers it.**

---

## How to Use This Skill: Core Workflow

**Every advisory session follows this three-step pattern:**

### Step 1 — Classify (instant, always run)

Run the classification commands. They use the embedded BetrVG knowledge base and work offline:

```bash
betriebsrat-pp-cli rights-check "<situation>" --agent
betriebsrat-pp-cli codetermination-type "<topic>" --agent
betriebsrat-pp-cli deadline "<type>" --from YYYY-MM-DD --agent
```

These answer: Does BR have a say? What kind? What's the deadline?

The output includes `topic_url` fields pointing to authoritative betriebsrat.de topic pages — use these URLs as citation sources in your advisory response.

### Step 2 — Deepen (for complete answers)

Chain additional commands to get the full picture:

```bash
# Get the applicable law explained in plain German
betriebsrat-pp-cli law <paragraph_number> --agent

# Get step-by-step action checklist
betriebsrat-pp-cli checklist "<situation>" --agent

# Get structured decision support with recommended action
betriebsrat-pp-cli decide "<situation>" --agent

# Get meeting preparation (agenda, quorum, questions for employer)
betriebsrat-pp-cli prepare-meeting "<topic>" --agent
```

### Step 3 — Compose (advisory response)

Build your advisory response from the command outputs. Use this structure:

```
Rechtslage: [co-determination type + applicable §§]
Ihr Recht: [what BR can do — block, demand BV, consult, or inform?]
Frist: [deadline if applicable]
Empfohlene Schritte: [ordered action list from checklist]
Weitere Informationen: [topic_url values from command output]
```

---

## When to Use Each Command

| Situation | Primary Command | Follow-up |
|-----------|----------------|-----------|
| "Does BR have a say?" | `rights-check` | `decide` for full decision |
| "What kind of right do we have?" | `codetermination-type` | `law` for paragraph detail |
| "When must we respond?" | `deadline` | `checklist` for full process |
| "What do we do step by step?" | `checklist` | `law` for legal basis |
| "What does § X mean?" | `law <n>` | `rights-check` for situation match |
| "How do I prepare this meeting?" | `prepare-meeting` | `checklist` for pre-meeting steps |
| "Help me decide what to do" | `decide` | all follow-ups |
| "Draft a formal BR response" | `letter` | `consequences` for leverage |
| "What happens if we miss deadline?" | `consequences` | `deadline` for exact date |
| "What if employer acts without consent?" | `consequences` | `decide` for action plan |
| "How much Sozialplan is this employee entitled to?" | `sozialplan-calc` | `law 112` for legal basis |
| "Store/update company profile" | `context set` | `context show` to verify |
| "Is this Anhörungsschreiben valid? Does the clock run?" | `check-anhoerung` | `deadline` for exact due date |
| "Draft a Betriebsvereinbarung for homeoffice/software" | `bv-template` | `law 87` for legal basis |
| "Export deadline to calendar" | `deadline ... --ical` | pipe to `.ics` file |
| "Do these layoffs trigger § 17 KSchG?" | `massenentlassung` | `law 17` for legal detail |
| "What are the strongest Widerspruch grounds?" | `widerspruch-check` | `letter kündigung --type widerspruch` |
| "Generate BR meeting minutes template" | `protokoll` | — |
| "Calculate Sozialplan for all affected employees" | `sozialplan-calc --csv` | `sozialplan-calc` per individual |
| "Request documents/data from employer" | `auskunft` | `consequences` for enforcement |
| "Does this AI system trigger co-determination?" | `ki-check` | `bv-template software` for draft BV |
| "Employer bypassed Interessenausgleich — what can employees claim?" | `nachteilsausgleich` | `sozialplan-calc` for comparison |
| "Send a BR training request to the employer" | `schulungsantrag` | `law 37` for legal detail |
| "Can we conclude a BV on this topic, or does the Tarifvertrag block it?" | `tarifvertrag-check` | `bv-template` to draft if allowed |
| "Draft a BV for camera surveillance" | `bv-template videoüberwachung` | `ki-check` if digital system involved |
| "Draft a BV for performance appraisals" | `bv-template leistungsbeurteilung` | `law 94` for legal basis |

---

## Scenario Playbooks

### Kündigung (Dismissal) — § 102 BetrVG

The employer is dismissing an employee. BR must be heard before every dismissal.

```bash
# 1. Calculate your deadline (runs first — deadlines are the #1 risk)
betriebsrat-pp-cli deadline "ordentliche Kündigung" --from $(date +%Y-%m-%d) --agent
# For extraordinary dismissal:
betriebsrat-pp-cli deadline "außerordentliche Kündigung" --from $(date +%Y-%m-%d) --agent

# 2. Get the full action checklist
betriebsrat-pp-cli checklist "Kündigung" --agent

# 3. Get the legal basis explained
betriebsrat-pp-cli law 102 --agent

# 4. Check co-determination type
betriebsrat-pp-cli codetermination-type "Kündigung Anhörung" --agent

# 5. Draft the formal response (Stellungnahme or Widerspruch)
betriebsrat-pp-cli letter kündigung --type widerspruch --employee "Name" --ground "fehlerhafte Sozialauswahl" --agent
betriebsrat-pp-cli letter kündigung --type zustimmung --employee "Name" --agent

# 6. Understand consequences of missed deadline
betriebsrat-pp-cli consequences kündigung --agent
```

**Key facts for Kündigung:**
- Ordinary dismissal: BR has **1 week** to respond (§ 102 Abs. 2)
- Extraordinary dismissal: BR has **3 days** (§ 102 Abs. 2 S. 3)
- Silence = consent — BR MUST respond within the window or forfeits rights
- BR can: consent, object (Widerspruch), or express concern
- Widerspruch grounds (§ 102 Abs. 3): wrong social selection, BV violation, Weiterbeschäftigung elsewhere possible, missing retraining, seniority ignored
- Widerspruch triggers right to **Weiterbeschäftigung** during appeal (§ 102 Abs. 5)

---

### Betriebsänderung (Operational Change) — §§ 111–113 BetrVG

Employer is restructuring: closing sites, significant layoffs, outsourcing, mergers.

```bash
# 1. Verify co-determination rights and scope
betriebsrat-pp-cli rights-check "Betriebsänderung Schließung Standort" --agent
betriebsrat-pp-cli codetermination-type "Betriebsänderung" --agent

# 2. Get full step-by-step checklist
betriebsrat-pp-cli checklist "Betriebsänderung" --agent

# 3. Understand the legal instruments
betriebsrat-pp-cli law 111 --agent  # What counts as Betriebsänderung
betriebsrat-pp-cli law 112 --agent  # Interessenausgleich + Sozialplan

# 4. Prepare the first meeting
betriebsrat-pp-cli prepare-meeting "Betriebsänderung § 111" --agent

# 5. Get structured decision support
betriebsrat-pp-cli decide "Arbeitgeber plant Schließung eines Standorts" --agent

# 6. Send formal letters
betriebsrat-pp-cli letter betriebsänderung --type unterrichtung --measure "Schließung Filiale Hamburg" --affected 45 --employer "Firma GmbH" --agent
betriebsrat-pp-cli letter betriebsänderung --type interessenausgleich --measure "Verlagerung Produktion nach Polen" --affected 120 --agent
```

**Key facts for Betriebsänderung:**
- BR must be informed and consulted **before** the decision is implemented (not just before execution)
- Two instruments: **Interessenausgleich** (try to avoid/limit) + **Sozialplan** (compensate those affected)
- Sozialplan is **erzwingbar** — BR can force it via Einigungsstelle
- Interessenausgleich is NOT erzwingbar — employer can act without agreement but must pay **Nachteilsausgleich** (§ 113)
- Threshold: generally ≥20% of workforce or absolute numbers from § 111 BetrVG (varies by company size)
- Trigger early: BR information rights start immediately upon employer decision-making — not just when announced publicly

---

### Software-Einführung / KI-Systeme — § 87 Abs. 1 Nr. 6 BetrVG

Employer wants to introduce new software, monitoring tools, AI systems, or performance tracking.

```bash
# 1. Check co-determination right (usually erzwingbar under § 87 Nr. 6)
betriebsrat-pp-cli rights-check "Einführung Software Leistungsüberwachung KI" --agent

# 2. Classify the right type precisely
betriebsrat-pp-cli codetermination-type "Überwachungssoftware" --agent

# 3. Get the legal basis
betriebsrat-pp-cli law 87 --agent  # Social co-determination, § 87 Abs. 1 Nr. 6

# 4. Get decision framework
betriebsrat-pp-cli decide "Arbeitgeber will KI-System einführen das Mitarbeiter bewertet" --agent

# 5. Prepare for negotiation
betriebsrat-pp-cli prepare-meeting "Einführung KI-System § 87" --agent
```

**Key facts for Software-Einführung:**
- § 87 Abs. 1 Nr. 6: **erzwingbare Mitbestimmung** for technical equipment *capable* of monitoring employee behavior or performance
- The monitoring capability triggers the right — even if the employer says "we won't use it for monitoring"
- Applies to: surveillance cameras, keyloggers, time-tracking, AI tools with employee data, Teams/Slack analytics, GitHub telemetry, code-review AI
- BR can **block the introduction** without an agreed Betriebsvereinbarung (BV)
- BV must cover: purpose, data collected, access rights, retention/deletion schedule, prohibition on disciplinary use

---

### Einstellung / Versetzung (Hiring / Transfer) — §§ 99–101 BetrVG

Employer wants to hire someone or transfer an existing employee to a different role/location.

```bash
# 1. Check co-determination right
betriebsrat-pp-cli rights-check "Einstellung Neueinstellung Versetzung" --agent

# 2. Get the legal framework
betriebsrat-pp-cli law 99 --agent   # Consent requirement for hiring/transfer
betriebsrat-pp-cli law 100 --agent  # Provisional measures without consent

# 3. Understand refusal grounds
betriebsrat-pp-cli decide "Arbeitgeber will neuen Mitarbeiter einstellen ohne BR zu fragen" --agent

# 4. Get checklist for the process
betriebsrat-pp-cli checklist "Einstellung" --agent
```

**Key facts for Einstellung/Versetzung:**
- **Zustimmungsvorbehalt** — employer needs BR consent (§ 99 Abs. 1)
- BR has **1 week** to respond (silence = consent)
- Grounds to refuse consent (§ 99 Abs. 2): BV violation, legal violation, existing employee disadvantage, no internal job posting (§ 93), wrong social selection
- Employer can proceed **without consent** in urgent cases (§ 100) but must apply to labor court within **3 days**
- If labor court rejects: employer must reverse the measure

---

### Massenentlassung (Mass Dismissal) — § 17 KSchG + §§ 111–113 BetrVG

Employer plans large-scale layoffs. § 17 KSchG adds a notification procedure on top of the regular § 102 hearing.

```bash
# 1. Check if § 17 KSchG threshold is met
betriebsrat-pp-cli massenentlassung --employees 200 --planned 25 --agent

# 2. If triggered: check BR rights for the Betriebsänderung (runs in parallel)
betriebsrat-pp-cli rights-check "Massenentlassung Stellenabbau" --agent
betriebsrat-pp-cli law 17 --agent   # § 17 KSchG notification procedure
betriebsrat-pp-cli law 112 --agent  # Sozialplan (erzwingbar)

# 3. Get Betriebsänderung checklist and decision support
betriebsrat-pp-cli checklist "Betriebsänderung Massenentlassung" --agent
betriebsrat-pp-cli decide "Arbeitgeber kündigt 25 von 200 Mitarbeitern" --agent

# 4. Calculate Sozialplan for all affected employees (batch mode)
betriebsrat-pp-cli sozialplan-calc --csv affected_employees.csv --factor 0.75 --max-cap 80000 --agent

# 5. Advise on Widerspruch grounds for each individual dismissal
betriebsrat-pp-cli widerspruch-check --type betriebsbedingt --seniority-ignored --other-position --agent

# 6. Generate BR resolution minutes for the Widerspruch vote
betriebsrat-pp-cli protokoll --topic "Massenentlassung: § 102-Anhörung und Widerspruch" --br-size 7 --agent

# 7. Draft the Betriebsrat Stellungnahme letter
betriebsrat-pp-cli letter betriebsänderung --type unterrichtung --measure "Abbau von 25 Stellen" --affected 25 --agent
```

**Key facts for Massenentlassung:**
- § 17 KSchG is **in addition to** § 102 BetrVG — both must be satisfied
- Failing to file the Massenentlassungsanzeige makes ALL terminations void (BAG 2016)
- The BR Stellungnahme must be attached to the Anzeige to the Agentur für Arbeit
- 1-month Sperrfrist after Anzeige before dismissals can take effect (extendable to 2 months)
- Sozialplan is **erzwingbar** — if negotiations fail, the Einigungsstelle decides

---

### Homeoffice / Mobile Work — §§ 87 Abs. 1 Nr. 2, 14 BetrVG

Employer wants to introduce, change, or end homeoffice/remote work arrangements.

```bash
# 1. Check rights
betriebsrat-pp-cli rights-check "Homeoffice mobile Arbeit Telearbeit Einführung Abschaffung" --agent

# 2. Get the legal basis
betriebsrat-pp-cli law 87 --agent   # Working hours (Nr. 2) + mobile work (Nr. 14)

# 3. Get structured decision support
betriebsrat-pp-cli decide "Arbeitgeber will Homeoffice abschaffen" --agent

# 4. Prepare the meeting
betriebsrat-pp-cli prepare-meeting "Homeoffice-Regelung Betriebsvereinbarung" --agent
```

**Key facts for Homeoffice:**
- § 87 Abs. 1 Nr. 2 (working hours) and Nr. 14 (mobile work, added 2021): **erzwingbare Mitbestimmung**
- BV should cover: who qualifies, equipment (employer provides?), ergonomics, reachability hours, data protection, accident coverage, cost reimbursement
- Employer **cannot unilaterally end** homeoffice governed by a BV without renegotiating
- Individual agreements do not replace a BV — BV governs the framework for everyone

---

## Unique Capabilities

### Decision support
- **`rights-check`** — Answers 'Does the Betriebsrat have a say in this?' — maps situation to BetrVG paragraphs and co-determination type

  ```bash
  betriebsrat-pp-cli rights-check "employer wants to introduce home office policy" --agent
  ```

- **`decide`** — Step-by-step decision support: classify situation, find applicable §§, determine BR rights, recommend action

  ```bash
  betriebsrat-pp-cli decide "Arbeitgeber kündigt 15 Mitarbeiter" --agent
  ```

- **`checklist`** — Generates step-by-step action checklist for BR in a given situation

  ```bash
  betriebsrat-pp-cli checklist "Betriebsänderung" --agent
  ```

- **`codetermination-type`** — Classifies BR rights as: Mitbestimmung (erzwingbar) / Mitwirkung / Unterrichtung / keine

  ```bash
  betriebsrat-pp-cli codetermination-type "Versetzung" --agent
  ```

### Legal deadlines
- **`deadline`** — Calculates legal deadlines for BR response

  ```bash
  betriebsrat-pp-cli deadline "ordentliche Kündigung" --from 2026-05-10 --agent
  ```

### Meeting tools
- **`prepare-meeting`** — Generates agenda, quorum rules, questions for employer for a BR meeting on a specific topic

  ```bash
  betriebsrat-pp-cli prepare-meeting "Einführung KI-System" --agent
  ```

### Legal reference
- **`law`** — Plain-language explanation of any BetrVG paragraph

  ```bash
  betriebsrat-pp-cli law 87 --agent
  ```

### Document drafting
- **`letter`** — Draft a formal BR letter: Stellungnahme, Widerspruch, Zustimmung, Verweigerung, Unterrichtungsverlangen, Interessenausgleich

  _The most practical command for day-to-day BR work. Generates a ready-to-edit German letter with correct legal references and structure._

  ```bash
  betriebsrat-pp-cli letter kündigung --type widerspruch --employee "Max Mustermann" --ground "fehlerhafte Sozialauswahl" --agent
  betriebsrat-pp-cli letter einstellung --type verweigerung --employee "Anna Schmidt" --ground "Verstoß gegen § 93 BetrVG" --agent
  betriebsrat-pp-cli letter versetzung --type zustimmung --employee "Peter Müller" --agent
  betriebsrat-pp-cli letter betriebsänderung --type unterrichtung --measure "Schließung Standort X" --affected 60 --agent
  betriebsrat-pp-cli letter betriebsänderung --type interessenausgleich --measure "Verlagerung Produktion" --affected 120 --agent
  ```

  Types for `kündigung`: `zustimmung` | `bedenken` | `widerspruch`
  Types for `einstellung`/`versetzung`: `zustimmung` | `verweigerung`
  Types for `betriebsänderung`: `unterrichtung` | `interessenausgleich`
  Flags for `betriebsänderung`: `--measure "<Maßnahme>"` `--affected <Anzahl>`

### Consequences
- **`consequences`** — What happens if BR misses a deadline or employer acts without consent?

  _Critical for understanding leverage and urgency. Know the exact legal consequences before deciding how to respond._

  ```bash
  betriebsrat-pp-cli consequences kündigung --agent
  betriebsrat-pp-cli consequences einstellung --agent
  betriebsrat-pp-cli consequences betriebsänderung --agent
  betriebsrat-pp-cli consequences software --agent
  ```

  Situations: `kündigung` | `einstellung` | `versetzung` | `betriebsänderung` | `software` | `br-deadline`

### Sozialplan calculation
- **`sozialplan-calc`** — Calculates individual or batch Sozialplan entitlement using the Munich formula

  _Use when a Betriebsänderung is happening and you need to estimate what each affected employee is entitled to. Use `--csv` for batch mode across all affected employees._

  ```bash
  # Single employee
  betriebsrat-pp-cli sozialplan-calc --salary 4500 --years 8 --age 42 --factor 0.75 --agent
  betriebsrat-pp-cli sozialplan-calc --salary 6000 --years 15 --age 58 --disabled --children 2 --factor 1.0 --agent
  # Batch mode — CSV: name,salary,years,age,disabled,children[,factor[,max_cap]]
  betriebsrat-pp-cli sozialplan-calc --csv employees.csv --factor 0.75 --max-cap 80000 --agent
  ```

  Formula: `Betriebszugehörigkeit × Monatsgehalt × Faktor`  
  Adjustments: +25% disabled, +10%/child (max 3), +5% if age ≥55  
  Factors: 0.5 (floor) · 0.75 (standard) · 1.0 (typical) · 1.5 (strong BR position)

### Massenentlassung threshold check
- **`massenentlassung`** — Checks whether § 17 KSchG applies and generates the complete compliance procedure

  _Always run this when large-scale dismissals are planned. Missing the Massenentlassungsanzeige makes all terminations void._

  ```bash
  betriebsrat-pp-cli massenentlassung --employees 150 --planned 25 --agent
  betriebsrat-pp-cli massenentlassung --employees 500 --planned 35 --agent
  ```

  Thresholds: 21–59 AN → ≥6 | 60–499 AN → ≥10% or ≥26 | ≥500 AN → ≥30  
  Output: triggered/not, 7-step procedure with deadlines, consequences if skipped

### Widerspruch grounds advisor
- **`widerspruch-check`** — Advises which § 102 Abs. 3 Widerspruch grounds are available and strongest

  _A Widerspruch (§ 102 Abs. 3) — unlike Bedenken (§ 102 Abs. 2) — gives the employee the right to continued employment during appeal (§ 102 Abs. 5). Use this command to pick the right grounds._

  ```bash
  betriebsrat-pp-cli widerspruch-check --type betriebsbedingt --wrong-social-selection --other-position --employee "Max Mustermann" --agent
  betriebsrat-pp-cli widerspruch-check --type verhaltensbedingt --no-warning --agent
  betriebsrat-pp-cli widerspruch-check --type betriebsbedingt --seniority-ignored --retraining --agent
  ```

  Grounds (§ 102 Abs. 3 Nr. 1–5): BV violation · wrong social selection · other position exists · retraining possible · changed terms possible  
  Output: applicable grounds ranked by strength, draft Widerspruch text ready to use in a letter

### Information requests
- **`auskunft`** — Drafts a formal § 80 BetrVG information request letter to the employer

  _The BR's most-used leverage tool. Use it to demand social data for Sozialauswahl, org charts, salary structures, AI system documentation, or any other information needed for the BR's statutory tasks. The letter includes the enforcement threat (labour court application)._

  ```bash
  betriebsrat-pp-cli auskunft --topic sozialdaten --reason "Prüfung Sozialauswahl § 102" --employer "Firma GmbH"
  betriebsrat-pp-cli auskunft --topic ki --reason "Einführung KI-Bewertungssystem" --deadline-days 10 --agent
  betriebsrat-pp-cli auskunft --topic custom --custom "Überstundenaufstellungen letzter 12 Monate" --lang en
  ```

  Topics: `sozialdaten` · `stellenplan` · `gehaelter` · `planung` · `auswahlrichtlinien` · `ki` · `wirtschaft` · `custom`  
  Letter is always in German (legal document); metadata/notes switch with `--lang en`

### AI/IT co-determination check
- **`ki-check`** — Analyses whether an AI or IT system triggers § 87 Abs. 1 Nr. 6 co-determination

  _The most important tool for the current wave of AI deployments. § 87 Nr. 6 is triggered by the capability to monitor employees, not actual use. Use this to determine whether to block deployment and what the BV must cover._

  ```bash
  betriebsrat-pp-cli ki-check --system "Workday People Analytics" --monitors-performance --influences-hr --auto-decision --lang en
  betriebsrat-pp-cli ki-check --system "GitHub Copilot" --data "keystrokes,accepted suggestions" --agent
  betriebsrat-pp-cli ki-check --system "Slack Workforce Analytics" --monitors-comms --monitors-performance --agent
  ```

  Flags: `--monitors-performance` · `--monitors-location` · `--monitors-comms` · `--influences-hr` · `--biometric` · `--auto-decision`  
  Output: triggered/not, risk level, required BV clauses, what employer cannot do without BV, 4 key BAG rulings

### Nachteilsausgleich calculator
- **`nachteilsausgleich`** — Calculates the § 113 BetrVG disadvantage compensation claim

  _When the employer implements a Betriebsänderung without attempting an Interessenausgleich (or deviates from one already agreed), every affected employee has a personal claim. This is separate from — and additive to — the Sozialplan. Use it to quantify leverage during negotiations._

  ```bash
  betriebsrat-pp-cli nachteilsausgleich --salary 5000 --years 12 --measure "Standortschließung" --no-ia-attempted --factor 0.75 --lang en
  betriebsrat-pp-cli nachteilsausgleich --salary 6000 --years 15 --measure "Verlagerung" --ia-deviated --agent
  ```

  Key rule: any existing Sozialplan payment is offset against the Nachteilsausgleich claim (§ 113 Abs. 3 Hs. 2)  
  Statutory cap: 12 × monthly salary (§ 10 KSchG analogy)

### BR meeting minutes
- **`protokoll`** — Generates a formal BR Sitzungsprotokoll template with quorum calculation

  _BR resolutions are invalid without proper minutes signed by the chair and secretary (§ 34 BetrVG). This covers all required fields._

  ```bash
  betriebsrat-pp-cli protokoll --topic "Kündigung Max Mustermann § 102" --br-size 7 --date 2026-05-15 --agent
  betriebsrat-pp-cli protokoll --topic "Homeoffice-BV Abstimmung" --br-size 11 --employer "Musterfirma GmbH" --agent
  ```

  Output: complete template with attendance sheet, quorum check, TOP structure (with voting rows), and signature block

### Company profile
- **`context`** — Stores and displays company profile for context-aware advice

  ```bash
  betriebsrat-pp-cli context set --employees 150 --sector IT --tariff --tariff-name "TV-L" --br-size 7 --bvs "Homeoffice,Arbeitszeit"
  betriebsrat-pp-cli context show --agent
  betriebsrat-pp-cli context reset
  ```

  Thresholds applied automatically:
  - `employees ≥ 20` → § 111 Betriebsänderung rights active
  - `employees ≥ 100` → § 106 Wirtschaftsausschuss mandatory
  - `employees ≥ 200` → § 38 full-time BR member release required

---

## Command Reference

**articles** — Individual articles and guides from betriebsrat.de
- `betriebsrat-pp-cli articles` — Search for articles within a topic area

**cases** — Recent Betriebsrat case law (Rechtsprechung)
- `betriebsrat-pp-cli cases` — Fetch recent court decisions relevant to works councils

**glossary** — Legal terms and definitions (Lexikon) for works council members
- `betriebsrat-pp-cli glossary list` — Browse legal terms glossary
- `betriebsrat-pp-cli glossary search` — Search for a specific legal term

**topics** — Betriebsrat topic areas with articles, guides, and practical tips
- `betriebsrat-pp-cli topics get` — Fetch full topic overview page with articles and guides
- `betriebsrat-pp-cli topics list` — List all topic areas (35+ Betriebsrat topics A-Z)

**context** — Store and display company profile for calibrated, threshold-aware advice
- `betriebsrat-pp-cli context set --employees <n> [--sector <s>] [--tariff] [--br-size <n>] [--bvs "<topics>"]`
- `betriebsrat-pp-cli context show` — Display profile and applicable BetrVG thresholds
- `betriebsrat-pp-cli context reset` — Delete stored profile

**sozialplan-calc** — Calculate Sozialplan entitlement (Munich formula), single or batch
- `betriebsrat-pp-cli sozialplan-calc --salary <eur> --years <n> [--age <n>] [--factor <f>] [--disabled] [--children <n>] [--max-cap <eur>]`
- `betriebsrat-pp-cli sozialplan-calc --csv <file> [--factor <f>] [--max-cap <eur>]` — CSV: `name,salary,years,age,disabled,children[,factor[,max_cap]]`

**massenentlassung** — Check § 17 KSchG threshold and generate compliance procedure
- `betriebsrat-pp-cli massenentlassung --employees <n> --planned <n>` — both flags required
- Output: triggered/not, 7-step procedure (BR consultation → Stellungnahme → Interessenausgleich/Sozialplan → Anzeige → Sperrfrist → § 102 per person → Kündigung), legal consequences

**widerspruch-check** — Advise on § 102 Abs. 3 Widerspruch grounds and draft ground text
- `betriebsrat-pp-cli widerspruch-check [--type betriebsbedingt|verhaltensbedingt|personenbedingt] [--wrong-social-selection] [--seniority-ignored] [--other-position] [--retraining] [--reduced-hours] [--bv-violation] [--no-warning] [--employee "<name>"]`
- Output: applicable grounds ranked by strength, draft Widerspruch text, deadline reminder

**auskunft** — Draft a formal § 80 BetrVG information request letter
- `betriebsrat-pp-cli auskunft --topic <topic> [--custom "<text>"] [--reason "<text>"] [--employer "<name>"] [--deadline-days <n>] [--date YYYY-MM-DD] [--lang en|de]`
- Topics: `sozialdaten` · `stellenplan` · `gehaelter` · `planung` · `auswahlrichtlinien` · `ki` · `wirtschaft` · `custom`

**ki-check** — Check § 87 Nr. 6 co-determination for an AI/IT system
- `betriebsrat-pp-cli ki-check --system "<description>" [--purpose "<text>"] [--data "<categories>"] [--monitors-performance] [--monitors-location] [--monitors-comms] [--influences-hr] [--biometric] [--auto-decision] [--lang en|de]`
- Output: triggered/not, risk level (low/medium/high/very high), required BV clauses, employer prohibitions, 4 BAG rulings

**nachteilsausgleich** — Calculate § 113 BetrVG disadvantage compensation claim
- `betriebsrat-pp-cli nachteilsausgleich --salary <eur> --years <n> [--age <n>] [--measure "<text>"] --no-ia-attempted | --ia-deviated [--factor <f>] [--lang en|de]`
- Cap: 12 × monthly salary; Sozialplan offset applies; evidence checklist included

**protokoll** — Generate formal BR Sitzungsprotokoll template
- `betriebsrat-pp-cli protokoll [--topic "<text>"] [--date YYYY-MM-DD] [--br-size <n>] [--location "<text>"] [--employer "<text>"]`
- Output: complete template with quorum calculation, attendance sheet, TOP structure, voting rows, signature block

**check-anhoerung** — Check a § 102 Anhörungsschreiben for completeness
- `betriebsrat-pp-cli check-anhoerung "<letter text>" [--type ordentlich|außerordentlich]`
- Reports: which required fields are present/missing, whether 7-day clock is running, severity per gap

**bv-template** — Generate a skeleton Betriebsvereinbarung
- `betriebsrat-pp-cli bv-template <topic> [--employer "<name>"] [--date YYYY-MM-DD]`
- Topics: `homeoffice` | `software` | `arbeitszeit` | `datenschutz` | `videoüberwachung` | `leistungsbeurteilung`

**schulungsantrag** — Draft a § 37 Abs. 6 BetrVG training request letter
- `betriebsrat-pp-cli schulungsantrag --topic <topic> [--training-name "<name>"] [--provider "<name>"] [--employer "<name>"]`
- Topics: `betrvg` | `arbeitsrecht` | `betriebsrat-praxis` | `kuendigung` | `sozialplan` | `datenschutz` | `gesundheit` | `custom`
- Output: complete letter with legal justification, including cost and release-from-work claims
- `--lang en` supported; letter body stays in German (formal legal document)

**tarifvertrag-check** — Check § 77 Abs. 3 Tarifvorbehalt before drafting a BV
- `betriebsrat-pp-cli tarifvertrag-check --topic <topic> [--tv-type "<type>"] [--tv-covers] [--opening-clause]`
- Topics: `lohn` | `arbeitszeit` | `urlaub` | `zulagen` | `homeoffice` | `software` | `gesundheit` | `custom`
- Output: blocked/not-blocked verdict, what the BV can and cannot cover, legal basis
- Always run this before drafting a BV in a TV-regulated area

**deadline** (updated) — now supports `--ical` flag
- `betriebsrat-pp-cli deadline "ordentliche Kündigung" --from 2026-05-10 --ical > frist.ics`
- Outputs a standard iCalendar file with a 1-day-before reminder; importable into Apple Calendar, Outlook, Google Calendar

**sync** — Populate or refresh the local SQLite knowledge base
- `betriebsrat-pp-cli sync` — Sync all topic areas (run once; safe to re-run)

**search** — Full-text search across the synced knowledge base
- `betriebsrat-pp-cli search "<query>" --data-source local` — Find passages in synced data

### Finding the right command

```bash
betriebsrat-pp-cli which "<capability in your own words>"
```

---

## Auth Setup

No authentication required. Run `betriebsrat-pp-cli doctor` to verify setup.

---

## Agent Mode

Add `--agent` to any command. Expands to: `--json --compact --no-input --no-color --yes`.

- **Pipeable** — JSON on stdout, errors on stderr
- **Filterable** — `--select` keeps a subset of fields
- **Offline-capable** — all embedded knowledge commands work without network access (`rights-check`, `decide`, `deadline`, `checklist`, `law`, `codetermination-type`, `consequences`, `letter`, `sozialplan-calc`, `context`, `check-anhoerung`, `bv-template`, `massenentlassung`, `widerspruch-check`, `protokoll`, `auskunft`, `ki-check`, `nachteilsausgleich`, `schulungsantrag`, `tarifvertrag-check`)
- **Bilingual** — add `--lang en` to any command for English output. Legal document bodies always stay in German.

### Response envelope

```json
{
  "meta": {"source": "live" | "local", "synced_at": "...", "reason": "..."},
  "results": <data>
}
```

---

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 2 | Usage error (wrong arguments) |
| 3 | Resource not found |
| 5 | API error (upstream issue) |
| 7 | Rate limited (wait and retry) |
| 10 | Config error |

---

## Argument Parsing

Parse `$ARGUMENTS`:

1. **Empty, `help`, or `--help`** → show `betriebsrat-pp-cli --help` output
2. **Starts with `install`** → ends with `mcp` → MCP installation; otherwise → see Prerequisites above
3. **Anything else** → Direct Use (execute as CLI command with `--agent`)

---

## MCP Server Installation

```bash
claude mcp add betriebsrat-pp-mcp -- betriebsrat-pp-mcp
```

Verify: `claude mcp list`

---

## Direct Use

1. Check if installed: `which betriebsrat-pp-cli`
   If not found, offer to install (see Prerequisites).
2. Match the user query to the best scenario playbook or command.
3. Execute with the `--agent` flag — chain multiple commands for a complete picture.
4. Compose the advisory response using the Classify → Deepen → Compose structure.
