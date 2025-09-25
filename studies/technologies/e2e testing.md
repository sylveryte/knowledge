# e2e testing

#testing

No doubts on which modern framework to choose. Answer is 'playwright'

## Playwright(pl) vs Selenium(sl)

- Speed : pl is much faster than sl, due to implmentation and better than sl's BIDI implmentation
- Parallel excution: pl native support, sl can do using sl grid but code needs to be changed
- Auto waiting: pl out of the box support, sl doesn't have; minimal flakiness
- Long term concerns: PL is backed by Microsoft, sl is old and stable and community driven
- Runners: pl no setup req, sl reqs
- Reports: pl has richer report units, screenshots, screen recording, html, auto retries
- Config: pl one file config
- CodeGen: pl has easy code gen tool from cli only, sl needs sl ide
- Tooling: pl no need but has vscode plugin, sl needs sl ide
- Tests: pl can do accessibility and api testing as well has aria locators, sl doesn't have aria locators
