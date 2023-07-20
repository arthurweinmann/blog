# More Complex Than DNA

The human DNA, comprised of approximately 3.5 billion base pairs, equates to about 750 Megabytes of data. This constitutes all the information necessary to construct a human being. Moreover, it encapsulates the instructions on how to interpret this data.

When I first compiled the Monaco Editor, the sum of all files amounted to around 1.4 Gigabytes - twice the size of the source code required for human construction! This includes files produced for testing, files created for various bundling systems like webpack, vite, parcel, among others, and files to support different programming languages in the editor. For sure, it is not only for the editor in itself. Nevertheless, it still represents a significant level of complexity. One would say more complexity than needed for a browser-based code editor - which, by the way, is quite impressive!

<div class="center_img">
    <img src="/assets/images/dna_complexity.png" />
</div>

The following script allows us to quantify the lines of code per file type in a repository:

```bash
find . -name '*.ts' | xargs wc -l
```

As of July 19, 2023, and at merge commit 0f95ee0d84b8425100bb4953c231e837402f400b, the Monaco Editor's codebase, excluding dependencies, comprises:

- 123302 lines of Typescrit (notably, TypeScript requires a dedicated compiler, which inherently introduces additional layers of complexity).
- 175849 lines of JavaScript
- 36205 lines of JSON, lines of JSON configuration, used predominantly for dependency management and versioning.

This sums up to an impressive 918,477 total lines of code, of which 514,194 are located within the .git repository. Though versioning mechanisms, such as Git, indeed introduce additional elements into the system, they can arguably reduce overall complexity by facilitating code comprehension and management for developers.

For context, complexity is typically defined as:

- The state or characteristic of having intricate or complicated components.
- The quality of being extremely detailed or complex, as exemplified by a highly interconnected system.

In the realm of computer science, the definition expands to characterize something with many parts where those parts interact with each other in multiple ways, culminating in a higher order of emergence greater than the sum of its parts. The intuitive criterion of complexity can be formulated as follows: a system would be more complex if more parts could be distinguished, and if more connections between them existed.

The Monaco Editor also maintains a substantial file: ./src/language/typescript/lib/typescriptServices.js, dedicated to TypeScript, which is incorporated via the npm run import-typescript command. TypeScript operates as a standalone system, necessitating individualized maintenance and support, thus contributing to the overall complexity.

This leads to the intriguing question: Is it feasible to reproduce TypeScript's capabilities purely in vanilla JavaScript, thereby eliminating the need for a separate framework and the associated maintenance overhead? Such a solution could potentially alleviate some of the aforementioned complexity.

<div class="wrap-collabsible">
  <input id="collapsible" class="toggle" type="checkbox"> 
  <label for="collapsible" class="lbl-toggle">The split of the code lines</label>
  <div class="collapsible-content">
  <pre><code class="language-bash">
919138 total
   514194 ./.git/objects/pack/pack-dfab111b0b254443247c3abac391016aba75078f.pack
   169809 ./src/language/typescript/lib/typescriptServices.js
    31764 ./test/manual/samples/run-editor-sample-f12-css.txt
    12530 ./package-lock.json
     7388 ./samples/package-lock.json
     7161 ./src/language/typescript/lib/typescriptServices.d.ts
     6210 ./samples/browser-esm-webpack-typescript-react/package-lock.json
     5361 ./src/basic-languages/freemarker2/freemarker2-auto-dollar.test.ts
     5358 ./src/basic-languages/freemarker2/freemarker2-auto-bracket.test.ts
     5346 ./website/static/monarch-static.html
     5215 ./.git/objects/pack/pack-dfab111b0b254443247c3abac391016aba75078f.idx
     3953 ./samples/browser-esm-parcel/package-lock.json
     3915 ./src/basic-languages/wgsl/wgsl.test.ts
     3463 ./website/src/website/data/home-samples/sample.c.txt
     3449 ./website/yarn.lock
     3122 ./webpack-plugin/package-lock.json
     3099 ./src/basic-languages/freemarker2/freemarker2-angle-dollar.test.ts
     3099 ./src/basic-languages/freemarker2/freemarker2-angle-bracket.test.ts
     3086 ./src/basic-languages/freemarker2/freemarker2-bracket-dollar.test.ts
     3086 ./src/basic-languages/freemarker2/freemarker2-bracket-bracket.test.ts
     2422 ./src/basic-languages/scss/scss.test.ts
     2409 ./src/basic-languages/php/php.test.ts
     2363 ./src/basic-languages/coffee/coffee.test.ts
     2314 ./samples/browser-esm-vite-react/package-lock.json
     2073 ./src/basic-languages/protobuf/protobuf.test.ts
     1560 ./src/basic-languages/less/less.test.ts
     1544 ./CHANGELOG.md
     1492 ./src/basic-languages/go/go.test.ts
     1390 ./src/basic-languages/abap/abap.ts
     1388 ./src/basic-languages/solidity/solidity.ts
     1269 ./src/language/typescript/languageFeatures.ts
     1267 ./src/basic-languages/systemverilog/systemverilog.test.ts
     1230 ./src/basic-languages/freemarker2/freemarker2.ts
     1228 ./src/basic-languages/hcl/hcl.test.ts
     1208 ./docs/debugging-core.gif
     1063 ./src/basic-languages/bicep/bicep.test.ts
     1050 ./src/basic-languages/postiats/postiats.test.ts
     1047 ./src/basic-languages/twig/twig.test.ts
     1004 ./src/language/common/lspLanguageFeatures.ts
      979 ./src/basic-languages/cpp/cpp.test.ts
      976 ./src/basic-languages/csharp/csharp.test.ts
      973 ./src/basic-languages/typescript/typescript.test.ts
      943 ./src/basic-languages/clojure/clojure.test.ts
      943 ./docs/debugging-languages.gif
      924 ./src/basic-languages/javascript/javascript.test.ts
      917 ./src/basic-languages/powerquery/powerquery.ts
      885 ./src/basic-languages/mysql/mysql.ts
      884 ./src/basic-languages/powershell/powershell.test.ts
      878 ./src/basic-languages/java/java.test.ts
      860 ./src/basic-languages/sql/sql.ts
      858 ./src/basic-languages/pgsql/pgsql.ts
      827 ./website/src/website/data/home-samples/sample.css.txt
      827 ./src/basic-languages/pgsql/keywords.postgresql.txt
      815 ./src/basic-languages/redshift/redshift.ts
      810 ./src/language/typescript/monaco.contribution.ts
      792 ./src/basic-languages/clojure/clojure.ts
      753 ./src/basic-languages/kotlin/kotlin.test.ts
      750 ./src/basic-languages/mysql/keywords.mysql.txt
      748 ./src/basic-languages/postiats/postiats.ts
      745 ./website/src/website/pages/playground/PlaygroundModel.ts
      743 ./src/basic-languages/html/html.test.ts
      727 ./src/basic-languages/scala/scala.test.ts
      706 ./src/basic-languages/xml/xml.test.ts
      700 ./src/basic-languages/apex/apex.test.ts
      692 ./src/basic-languages/sql/sql.test.ts
      670 ./src/basic-languages/yaml/yaml.test.ts
      664 ./src/basic-languages/perl/perl.ts
      656 ./src/basic-languages/elixir/elixir.ts
      636 ./src/basic-languages/php/php.ts
      612 ./src/basic-languages/sql/keywords.js
      606 ./src/basic-languages/systemverilog/systemverilog.ts
      600 ./src/basic-languages/r/r.test.ts
      598 ./src/basic-languages/vb/vb.test.ts
      598 ./src/basic-languages/pla/pla.test.ts
      595 ./src/basic-languages/redshift/redshift.test.ts
      595 ./src/basic-languages/pgsql/pgsql.test.ts
      591 ./src/basic-languages/css/css.test.ts
      590 ./src/basic-languages/ruby/ruby.ts
      571 ./src/basic-languages/razor/razor.ts
      561 ./src/basic-languages/mysql/mysql.test.ts
      558 ./src/basic-languages/dart/dart.test.ts
      550 ./src/basic-languages/julia/julia.ts
      510 ./website/src/website/pages/playground/PlaygroundPageContent.tsx
      509 ./src/language/typescript/tsWorker.ts
      493 ./test/manual/samples/run-editor-sample-bom-cs.txt
      491 ./src/basic-languages/fsharp/fsharp.test.ts
      485 ./src/basic-languages/wgsl/wgsl.ts
      481 ./src/basic-languages/ecl/ecl.ts
      480 ./website/src/website/data/playground-samples/extending-language-services/symbols-provider-example/sample.js
      475 ./src/basic-languages/perl/perl.test.ts
      473 ./src/basic-languages/pug/pug.test.ts
      463 ./src/basic-languages/protobuf/protobuf.ts
      454 ./src/basic-languages/restructuredtext/restructuredtext.test.ts
      451 ./test/manual/diff.js
      447 ./src/basic-languages/elixir/elixir.test.ts
      444 ./src/basic-languages/msdax/msdax.test.ts
      441 ./src/basic-languages/st/st.ts
      434 ./src/basic-languages/pug/pug.ts
      432 ./src/basic-languages/handlebars/handlebars.ts
      431 ./src/basic-languages/cpp/cpp.ts
      427 ./test/manual/index.js
      426 ./build/build-monaco-editor.ts
      417 ./src/basic-languages/twig/twig.ts
      416 ./src/basic-languages/powerquery/powerquery.test.ts
      414 ./website/src/website/pages/playground/getNpmVersionsSync.ts
      413 ./src/basic-languages/scala/scala.ts
      412 ./website/src/website/pages/playground/SettingsDialog.tsx
      397 ./src/basic-languages/bat/bat.test.ts
      394 ./src/basic-languages/vb/vb.ts
      392 ./ThirdPartyNotices.txt
      382 ./src/basic-languages/msdax/msdax.ts
      368 ./webpack-plugin/src/index.ts
      368 ./src/basic-languages/objective-c/objective-c.test.ts
      364 ./src/basic-languages/typescript/typescript.ts
      362 ./src/language/css/monaco.contribution.ts
      361 ./src/basic-languages/rust/rust.ts
      360 ./src/basic-languages/abap/abap.test.ts
      355 ./src/basic-languages/html/html.ts
      355 ./src/basic-languages/apex/apex.ts
      354 ./src/basic-languages/handlebars/handlebars.test.ts
      352 ./src/basic-languages/csharp/csharp.ts
      347 ./docs/out-folders.dio.svg
      345 ./src/basic-languages/shell/shell.test.ts
      335 ./src/basic-languages/swift/swift.ts
      329 ./src/language/html/monaco.contribution.ts
      327 ./src/basic-languages/cypher/cypher.test.ts
      324 ./.git/packed-refs
      320 ./src/basic-languages/sb/sb.test.ts
      309 ./src/basic-languages/redis/redis.ts
      307 ./src/basic-languages/qsharp/qsharp.ts
      307 ./.git/index
      306 ./src/basic-languages/redis/redis.test.ts
      305 ./src/basic-languages/dart/dart.ts
      302 ./src/basic-languages/sophia/sophia.test.ts
      297 ./website/static/monarch/monarch.css
      289 ./build/utils.ts
      287 ./test/manual/samples.js
      286 ./test/manual/mouse-scrollable-body.html
      284 ./src/basic-languages/scss/scss.ts
      280 ./build/releaseMetadata.ts
      275 ./docs/code-structure.dio.svg
      274 ./src/basic-languages/cypher/cypher.ts
      273 ./src/language/json/tokenization.ts
      273 ./src/basic-languages/kotlin/kotlin.ts
      272 ./src/basic-languages/solidity/solidity.test.ts
      271 ./website/src/website/components/monaco/MonacoEditor.tsx
      271 ./src/basic-languages/python/python.ts
      270 ./src/basic-languages/pascal/pascal.ts
      263 ./website/src/website/data/home-samples/sample.graphql.txt
      263 ./src/basic-languages/r/r.ts
      263 ./src/basic-languages/powershell/powershell.ts
      259 ./src/basic-languages/tcl/tcl.ts
      258 ./src/basic-languages/coffee/coffee.ts
      256 ./src/basic-languages/markdown/markdown.ts
      254 ./src/basic-languages/swift/swift.test.ts
      252 ./src/basic-languages/java/java.ts
      251 ./src/basic-languages/liquid/liquid.ts
      249 ./website/src/website/pages/home/Home.tsx
      246 ./website/src/website/data/home-samples/sample.cpp.txt
      246 ./src/basic-languages/shell/shell.ts
      239 ./src/language/css/cssWorker.ts
      239 ./src/basic-languages/mips/mips.test.ts
      238 ./src/basic-languages/go/go.ts
      238 ./src/basic-languages/dockerfile/dockerfile.test.ts
      237 ./src/basic-languages/fsharp/fsharp.ts
      236 ./src/basic-languages/yaml/yaml.ts
      229 ./docs/integrate-esm.md
      227 ./src/language/json/monaco.contribution.ts
      226 ./src/basic-languages/m3/m3.ts
      223 ./src/basic-languages/mips/mips.ts
      221 ./test/manual/samples/run-editor-sample-js.txt
      220 ./src/basic-languages/sophia/sophia.ts
      219 ./src/basic-languages/sparql/sparql.ts
      219 ./src/basic-languages/objective-c/objective-c.ts
      218 ./test/manual/dev-setup.js
      216 ./src/language/json/jsonWorker.ts
      215 ./website/src/website/data/diff-sample/original.txt
      215 ./src/basic-languages/css/css.ts
      213 ./website/src/website/data/home-samples/sample.javascript.txt
      212 ./website/src/website/data/diff-sample/modified.txt
      212 ./build/build-languages.ts
      210 ./test/manual/typescript/custom-worker.html
      209 ./test/smoke/smoke.test.js
      209 ./src/basic-languages/liquid/liquid.test.ts
      208 ./src/basic-languages/rust/rust.test.ts
      204 ./test/manual/mouse-scrollable-element.html
      198 ./src/basic-languages/restructuredtext/restructuredtext.ts
      198 ./src/basic-languages/less/less.ts
      196 ./website/src/website/data/playground-samples/all.js
      193 ./src/basic-languages/python/python.test.ts
      193 ./src/basic-languages/cameligo/cameligo.ts
      192 ./build/importTypescript.ts
      189 ./website/src/website/pages/playground/SettingsModel.ts
      189 ./src/basic-languages/hcl/hcl.ts
      187 ./src/basic-languages/sparql/sparql.test.ts
      183 ./src/basic-languages/pascaligo/pascaligo.ts
      179 ./test/manual/typescript/index.html
      179 ./src/basic-languages/lexon/lexon.ts
      178 ./website/static/monarch/monarch.js
      178 ./src/basic-languages/lua/lua.ts
      176 ./src/language/typescript/tsMode.ts
      175 ./src/language/json/jsonMode.ts
      175 ./src/basic-languages/razor/razor.test.ts
      174 ./src/basic-languages/graphql/graphql.ts
      173 ./.git/hooks/fsmonitor-watchman.sample
      172 ./src/basic-languages/azcli/azcli.test.ts
      171 ./src/language/html/htmlMode.ts
      171 ./src/basic-languages/mdx/mdx.test.ts
      171 ./build/simpleserver.ts
      169 ./.git/hooks/pre-rebase.sample
      168 ./website/src/website/data/playground-samples/extending-language-services/semantic-tokens-provider-example/sample.js
      168 ./samples/browser-esm-webpack-small/index.js
      164 ./src/language/html/htmlWorker.ts
      163 ./src/basic-languages/mdx/mdx.ts
      162 ./src/basic-languages/qsharp/qsharp.test.ts
      160 ./src/basic-languages/flow9/flow9.ts
      158 ./src/basic-languages/ruby/ruby.test.ts
      157 ./src/basic-languages/pla/pla.ts
      151 ./website/webpack.config.ts
      151 ./src/basic-languages/pascal/pascal.test.ts
      150 ./website/src/website/data/playground-samples/customizing-the-appearence/exposed-colors/sample.js
      148 ./src/basic-languages/flow9/flow9.test.ts
      145 ./src/language/css/cssMode.ts
      145 ./src/basic-languages/lexon/lexon.test.ts
      145 ./src/basic-languages/dockerfile/dockerfile.ts
      144 ./src/basic-languages/cameligo/cameligo.test.ts
      143 ./src/basic-languages/pascaligo/pascaligo.test.ts
      140 ./samples/browser-esm-webpack-small/generate-imports.js
      139 ./src/basic-languages/st/st.test.ts
      138 ./website/src/website/data/home-samples/sample.sol.txt
      138 ./src/basic-languages/graphql/graphql.test.ts
      137 ./website/src/website/data/playground-samples/extending-language-services/custom-languages/sample.js
      134 ./src/basic-languages/sb/sb.ts
      129 ./src/basic-languages/bicep/bicep.ts
      128 ./src/basic-languages/scheme/scheme.ts
      128 ./.git/hooks/update.sample
      125 ./CONTRIBUTING.md
      124 ./website/src/website/data/home-samples/sample.typescript.txt
      122 ./README.md
      121 ./src/basic-languages/bat/bat.ts
      119 ./src/basic-languages/m3/m3.test.ts
      116 ./webpack-plugin/README.md
      114 ./website/src/website/data/playground-samples/interacting-with-the-editor/listening-to-mouse-events/sample.js
      111 ./website/src/website/data/home-samples/sample.go.txt
      111 ./src/basic-languages/freemarker2/freemarker2.contribution.ts
      107 ./website/src/monaco-loader.ts
      107 ./src/basic-languages/tcl/tcl.test.ts
      106 ./src/basic-languages/xml/xml.ts
      105 ./samples/browser-amd-monarch/index.html
      104 ./website/src/website/data/home-samples/sample.scheme.txt
      104 ./website/src/runner/index.ts
      103 ./website/src/website/pages/DocsPage.tsx
      103 ./website/src/website/data/home-samples/sample.markdown.txt
      103 ./.github/workflows/publish/computeState.js
      101 ./src/language/typescript/workerManager.ts
      100 ./website/src/website/data/home-samples/sample.html.txt
       97 ./website/src/website/pages/playground/Preview.tsx
       96 ./samples/browser-esm-esbuild/build.js
       95 ./src/basic-languages/markdown/markdown.test.ts
       94 ./scripts/lib/index.ts
       93 ./website/src/website/style.scss
       93 ./test/smoke/runner.js
       92 ./website/index/samples/sample.mdx.txt
       91 ./website/src/website/data/home-samples/sample.mdx.txt
       91 ./src/language/json/workerManager.ts
       90 ./src/language/html/workerManager.ts
       90 ./src/language/css/workerManager.ts
       89 ./website/src/website/utils/ObservableHistory.ts
       89 ./src/basic-languages/scheme/scheme.test.ts
       88 ./website/src/website/pages/playground/playgroundExamples.tsx
       87 ./src/language/typescript/lib/lib.ts
       87 ./src/language/typescript/lib/lib.index.ts
       86 ./src/basic-languages/lua/lua.test.ts
       85 ./src/basic-languages/monaco.contribution.ts
       82 ./src/basic-languages/ini/ini.ts
       81 ./website/src/website/data/home-samples/sample.php.txt
       81 ./package.json
       80 ./website/src/website/data/home-samples/sample.elixir.txt
       80 ./src/basic-languages/_.contribution.ts
       80 ./build/fs.ts
       80 ./.azure-pipelines/publish-stable.yml
       78 ./scripts/ci/monaco-editor-core-prepare.ts
       78 ./.git/hooks/push-to-checkout.sample
       77 ./src/basic-languages/javascript/javascript.ts
       77 ./src/basic-languages/azcli/azcli.ts
       75 ./website/src/website/components/Select.tsx
       75 ./website/src/website/components/Nav.tsx
       75 ./.github/workflows/ci.yml
       75 ./.azure-pipelines/publish-nightly.yml
       74 ./website/src/website/data/playground-samples/extending-language-services/folding-provider-example/sample.js
       74 ./samples/browser-amd-diff-editor/modified.txt
       72 ./scripts/ci/monaco-editor-prepare.ts
       71 ./website/src/website/data/playground-samples/extending-language-services/color-provider-example/sample.js
       70 ./samples/browser-amd-iframe/index.html
       69 ./website/src/website/data/playground-samples/extending-language-services/completion-provider-example/sample.js
       68 ./.github/workflows/website.yml
       67 ./website/src/website/data/home-samples/sample.json.txt
       65 ./website/src/website/data/home-samples/sample.freemarker2.tag-auto.interpolation-dollar.txt
       65 ./website/src/website/data/home-samples/sample.freemarker2.tag-auto.interpolation-bracket.txt
       64 ./website/src/website/data/home-samples/sample.postiats.txt
       64 ./test/unit/all.js
       63 ./webpack-plugin/src/plugins/AddWorkerEntryPointPlugin.ts
       63 ./src/basic-languages/test/testRunner.ts
       62 ./website/src/website/data/playground-samples/extending-language-services/hover-provider-example/sample.js
       62 ./samples/browser-amd-diff-editor/index.html
       61 ./website/src/website/data/home-samples/sample.freemarker2.txt
       61 ./website/src/website/data/home-samples/sample.freemarker2.tag-bracket.interpolation-dollar.txt
       61 ./website/src/website/data/home-samples/sample.freemarker2.tag-bracket.interpolation-bracket.txt
       61 ./website/src/website/data/home-samples/sample.freemarker2.tag-angle.interpolation-dollar.txt
       61 ./website/src/website/data/home-samples/sample.freemarker2.tag-angle.interpolation-bracket.txt
       61 ./samples/browser-amd-diff-editor/original.txt
       60 ./test/smoke/package-esbuild.ts
       59 ./website/src/website/data/home-samples/sample.vb.txt
       59 ./website/package.json
       59 ./test/manual/typescript/custom-worker.js
       59 ./samples/browser-esm-webpack-typescript-react/webpack.config.js
       58 ./website/src/website/utils/utils.ts
       58 ./website/scripts/check-playground-samples-js.ts
       58 ./src/basic-languages/csp/csp.ts
       57 ./.vscode/launch.json
       57 ./samples/nwjs-amd-v2/index.html
       57 ./.github/ISSUE_TEMPLATE/1_bug_report.yaml
       56 ./test/smoke/package-webpack.ts
       56 ./test/manual/shadow-dom.html
       53 ./website/src/website/data/home-samples/sample.scala.txt
       53 ./website/src/website/data/home-samples/sample.java.txt
       53 ./website/src/website/data/home-samples/sample.clojure.txt
       53 ./.git/hooks/pre-push.sample
       53 ./build/check-samples.ts
       52 ./website/src/website/data/playground-samples/interacting-with-the-editor/adding-an-action-to-an-editor-instance/sample.js
       52 ./website/src/website/data/home-samples/sample.sql.txt
       52 ./website/src/website/data/home-samples/sample.objective-c.txt
       51 ./website/.gitignore
       50 ./website/src/website/data/playground-samples/extending-language-services/model-markers-example/sample.js
       50 ./samples/browser-amd-shadow-dom/index.html
       49 ./website/src/website/utils/ObservablePromise.ts
       49 ./website/src/website/utils/hotComponent.tsx
       49 ./website/src/website/data/playground-samples/interacting-with-the-editor/adding-a-keybinding-to-an-existing-command/sample.js
       49 ./website/src/website/data/home-samples/sample.swift.txt
       49 ./test/manual/samples/run-editor-sample-dynamic.txt
       49 ./.git/hooks/pre-commit.sample
       48 ./website/src/website/data/home-samples/sample.wgsl.txt
       48 ./website/src/website/data/home-samples/sample.hcl.txt
       48 ./samples/nwjs-amd/index.html
       47 ./website/src/website/data/playground-samples/extending-language-services/inlay-hints-provider-example/sample.js
       47 ./MAINTAINING.md
       46 ./website/src/website/data/playground-samples/extending-language-services/configure-javascript-defaults/sample.js
       46 ./website/src/website/data/home-samples/sample.razor.txt
       46 ./website/src/website/data/home-samples/sample.less.txt
       46 ./samples/electron-amd-nodeIntegration/electron-index.html
       46 ./build/fillers/vscode-nls.ts
       45 ./website/src/website/data/home-samples/sample.qsharp.txt
       45 ./website/src/website/data/home-samples/sample.dart.txt
       45 ./webpack-plugin/package.json
       44 ./test/manual/index.html
       44 ./src/basic-languages/xml/xml.contribution.ts
       44 ./samples/browser-esm-webpack-typescript/webpack.config.js
       44 ./build/npm/installAll.ts
       43 ./website/src/website/data/playground-samples/interacting-with-the-editor/adding-a-command-to-an-editor-instance/sample.js
       43 ./test/manual/samples/run-editor-failing-js.txt
       42 ./website/src/website/data/playground-samples/extending-language-services/configure-json-defaults/sample.js
       42 ./samples/browser-amd-trusted-types/index.html
       42 ./.git/hooks/prepare-commit-msg.sample
       41 ./website/src/website/data/playground-samples/interacting-with-the-editor/revealing-a-position/sample.js
       41 ./website/src/website/data/home-samples/sample.shell.txt
       41 ./website/src/website/data/home-samples/sample.r.txt
       41 ./webpack-plugin/src/loaders/include.ts
       41 ./test/smoke/parcel/index.js
       41 ./SECURITY.md
       40 ./website/src/website/data/playground-samples/customizing-the-appearence/tokens-and-colors/sample.js
       39 ./website/src/website/utils/ref.ts
       39 ./website/src/website/utils/lzmaCompressor.ts
       39 ./website/src/website/data/playground-samples/extending-language-services/codelens-provider-example/sample.js
       39 ./src/basic-languages/systemverilog/systemverilog.contribution.ts
       38 ./website/src/website/data/home-samples/sample.csharp.txt
       38 ./website/src/website/components/monaco/MonacoLoader.tsx
       38 ./src/basic-languages/julia/julia.test.ts
       38 ./src/basic-languages/cpp/cpp.contribution.ts
       38 ./samples/browser-esm-webpack-typescript-react/src/components/Editor.tsx
       38 ./.github/ISSUE_TEMPLATE/2_feature_request.yaml
       37 ./website/src/website/data/playground-samples/interacting-with-the-editor/customizing-the-line-numbers/sample.js
       37 ./website/src/website/data/home-samples/sample.aes.txt
       37 ./website/src/shared.ts
       37 ./samples/electron-amd-nodeIntegration/main.js
       37 ./samples/electron-amd/electron-index.html
       36 ./website/src/website/data/home-samples/sample.scss.txt
       36 ./samples/README.md
       35 ./website/src/website/data/playground-samples/creating-the-editor/web-component/sample.js
       35 ./website/src/website/data/playground-samples/creating-the-editor/syntax-highlighting-for-html-elements/sample.html
       35 ./website/src/website/data/home-samples/sample.verilog.txt
       35 ./website/src/runner/style.scss
       35 ./samples/browser-esm-webpack-small/webpack.config.js
       34 ./scripts/ci/monaco-editor-core.sh
       34 ./samples/browser-amd-iframe/inner.html
       33 ./website/src/website/data/home-samples/sample.st.txt
       33 ./test/manual/samples/run-editor-intellisense-js.txt
       32 ./website/src/website/data/playground-samples/customizing-the-appearence/scrollbars/sample.js
       32 ./website/src/website/components/bootstrap.tsx
       32 ./src/basic-languages/mysql/keywords.js
       32 ./samples/package.json
       32 ./samples/electron-esm-webpack/main.js
       32 ./samples/electron-amd/main.js
       31 ./website/src/website/data/home-samples/sample.handlebars.txt
       31 ./website/src/website/data/home-samples/sample.dockerfile.txt
       31 ./test/smoke/vite/index.js
       31 ./test/manual/mouse-fixed.html
       30 ./website/src/website/data/playground-samples/interacting-with-the-editor/rendering-glyphs-in-the-margin/sample.js
       30 ./website/src/website/data/playground-samples/interacting-with-the-editor/line-and-inline-decorations/sample.js
       30 ./.vscode/tasks.json
       30 ./src/basic-languages/pgsql/keywords.js
       30 ./samples/browser-esm-webpack/webpack.config.js
       30 ./samples/browser-amd-localized/index.html
       29 ./website/src/website/data/home-samples/sample.rust.txt
       29 ./website/src/website/data/home-samples/sample.powershell.txt
       29 ./website/src/website/data/home-samples/sample.bicep.txt
       29 ./website/src/website/data/home-samples/sample.abap.txt
       29 ./test/manual/transform.html
       29 ./test/manual/samples/run-editor-jquery-min-js.txt
       29 ./test/manual/iframe-inner.html
       29 ./samples/electron-esm-webpack/webpack.config.js
       29 ./samples/browser-script-editor/index.html
       29 ./samples/browser-esm-parcel/src/index.js
       29 ./samples/browser-amd-shared-model/index.html
       29 ./.github/workflows/info-needed-closer.yml
       28 ./website/src/website/utils/types.d.ts
       28 ./website/src/website/utils/Debouncer.ts
       28 ./website/src/website/pages/routes.ts
       28 ./website/src/website/data/home-samples/sample.systemverilog.txt
       28 ./website/src/website/data/home-samples/sample.pascal.txt
       28 ./website/src/website/data/home-samples/sample.coffeescript.txt
       28 ./test/unit/setup.js
       28 ./scripts/ci/monaco-editor.sh
       28 ./samples/browser-esm-webpack-monaco-plugin/webpack.config.js
       28 ./docs/integrate-amd.md
       27 ./website/src/website/data/home-samples/sample.kotlin.txt
       27 ./src/basic-languages/javascript/javascript.contribution.ts
       27 ./samples/electron-esm-webpack/dist/electron-index.html
       27 ./samples/browser-esm-vite-react/src/userWorker.ts
       27 ./.github/workflows/locker.yml
       26 ./website/src/website/data/playground-samples/creating-the-diffeditor/navigating-a-diff/sample.js
       26 ./website/src/website/data/home-samples/sample.restructuredtext.txt
       26 ./test/smoke/esbuild/index.js
       26 ./test/manual/cross-origin.html
       26 ./samples/browser-esm-webpack-typescript/src/index.ts
       26 ./samples/browser-esm-webpack-small/index.html
       26 ./samples/browser-amd-requirejs/index.html
       25 ./src/basic-languages/yaml/yaml.contribution.ts
       25 ./src/basic-languages/typescript/typescript.contribution.ts
       25 ./src/basic-languages/twig/twig.contribution.ts
       25 ./src/basic-languages/swift/swift.contribution.ts
       25 ./src/basic-languages/scss/scss.contribution.ts
       25 ./src/basic-languages/scala/scala.contribution.ts
       25 ./src/basic-languages/ruby/ruby.contribution.ts
       25 ./src/basic-languages/razor/razor.contribution.ts
       25 ./src/basic-languages/python/python.contribution.ts
       25 ./src/basic-languages/php/php.contribution.ts
       25 ./src/basic-languages/pascal/pascal.contribution.ts
       25 ./src/basic-languages/mips/mips.contribution.ts
       25 ./src/basic-languages/liquid/liquid.contribution.ts
       25 ./src/basic-languages/less/less.contribution.ts
       25 ./src/basic-languages/kotlin/kotlin.contribution.ts
       25 ./src/basic-languages/java/java.contribution.ts
       25 ./src/basic-languages/ini/ini.contribution.ts
       25 ./src/basic-languages/html/html.contribution.ts
       25 ./src/basic-languages/handlebars/handlebars.contribution.ts
       25 ./src/basic-languages/graphql/graphql.contribution.ts
       25 ./src/basic-languages/dockerfile/dockerfile.contribution.ts
       25 ./src/basic-languages/dart/dart.contribution.ts
       25 ./src/basic-languages/css/css.contribution.ts
       25 ./src/basic-languages/coffee/coffee.contribution.ts
       25 ./src/basic-languages/apex/apex.contribution.ts
       25 ./samples/browser-esm-vite-react/src/components/Editor.tsx
       24 ./src/basic-languages/wgsl/wgsl.contribution.ts
       24 ./src/basic-languages/vb/vb.contribution.ts
       24 ./src/basic-languages/tcl/tcl.contribution.ts
       24 ./src/basic-languages/st/st.contribution.ts
       24 ./src/basic-languages/sql/sql.contribution.ts
       24 ./src/basic-languages/sparql/sparql.contribution.ts
       24 ./src/basic-languages/sophia/sophia.contribution.ts
       24 ./src/basic-languages/solidity/solidity.contribution.ts
       24 ./src/basic-languages/shell/shell.contribution.ts
       24 ./src/basic-languages/scheme/scheme.contribution.ts
       24 ./src/basic-languages/sb/sb.contribution.ts
       24 ./src/basic-languages/rust/rust.contribution.ts
       24 ./src/basic-languages/r/r.contribution.ts
       24 ./src/basic-languages/restructuredtext/restructuredtext.contribution.ts
       24 ./src/basic-languages/redshift/redshift.contribution.ts
       24 ./src/basic-languages/redis/redis.contribution.ts
       24 ./src/basic-languages/qsharp/qsharp.contribution.ts
       24 ./src/basic-languages/pug/pug.contribution.ts
       24 ./src/basic-languages/protobuf/protobuf.contribution.ts
       24 ./src/basic-languages/powershell/powershell.contribution.ts
       24 ./src/basic-languages/powerquery/powerquery.contribution.ts
       24 ./src/basic-languages/postiats/postiats.contribution.ts
       24 ./src/basic-languages/pgsql/pgsql.contribution.ts
       24 ./src/basic-languages/perl/perl.contribution.ts
       24 ./src/basic-languages/pascaligo/pascaligo.contribution.ts
       24 ./src/basic-languages/objective-c/objective-c.contribution.ts
       24 ./src/basic-languages/mysql/mysql.contribution.ts
       24 ./src/basic-languages/msdax/msdax.contribution.ts
       24 ./src/basic-languages/mdx/mdx.contribution.ts
       24 ./src/basic-languages/markdown/markdown.contribution.ts
       24 ./src/basic-languages/m3/m3.contribution.ts
       24 ./src/basic-languages/lua/lua.contribution.ts
       24 ./src/basic-languages/lexon/lexon.contribution.ts
       24 ./src/basic-languages/julia/julia.contribution.ts
       24 ./src/basic-languages/hcl/hcl.contribution.ts
       24 ./src/basic-languages/go/go.contribution.ts
       24 ./src/basic-languages/fsharp/fsharp.contribution.ts
       24 ./src/basic-languages/flow9/flow9.contribution.ts
       24 ./src/basic-languages/elixir/elixir.contribution.ts
       24 ./src/basic-languages/ecl/ecl.contribution.ts
       24 ./src/basic-languages/cypher/cypher.contribution.ts
       24 ./src/basic-languages/csp/csp.contribution.ts
       24 ./src/basic-languages/csharp/csharp.contribution.ts
       24 ./src/basic-languages/clojure/clojure.contribution.ts
       24 ./src/basic-languages/cameligo/cameligo.contribution.ts
       24 ./src/basic-languages/bicep/bicep.contribution.ts
       24 ./src/basic-languages/bat/bat.contribution.ts
       24 ./src/basic-languages/azcli/azcli.contribution.ts
       24 ./src/basic-languages/abap/abap.contribution.ts
       24 ./samples/electron-esm-webpack/index.js
       24 ./samples/browser-esm-webpack/index.js
       24 ./samples/browser-esm-esbuild/index.js
       24 ./samples/browser-amd-editor/index.html
       24 ./.git/hooks/pre-receive.sample
       24 ./.git/hooks/commit-msg.sample
       23 ./website/src/website/data/playground-samples/creating-the-diffeditor/inline-diff-example/sample.js
       23 ./website/src/website/data/home-samples/sample.julia.txt
       23 ./website/src/website/components/Radio.tsx
       23 ./test/manual/samples/run-editor-sample-html.txt
       23 ./test/manual/iframe.html
       23 ./src/basic-languages/pla/pla.contribution.ts
       23 ./build/postinstall.ts
       21 ./website/src/website/pages/App.tsx
       21 ./website/src/website/data/home-samples/sample.lex.txt
       21 ./website/src/website/components/TextBox.tsx
       21 ./webpack-plugin/LICENSE
       21 ./samples/browser-esm-webpack-typescript-react/package.json
       21 ./LICENSE.txt
       20 ./website/src/website/data/playground-samples/creating-the-diffeditor/multi-line-example/sample.js
       20 ./website/src/website/data/playground-samples/creating-the-diffeditor/hello-diff-world/sample.js
       20 ./website/src/website/data/home-samples/sample.ruby.txt
       20 ./website/src/website/data/home-samples/sample.pascaligo.txt
       20 ./test/smoke/amd/index.html
       20 ./test/manual/diff.html
       20 ./samples/browser-esm-webpack-monaco-plugin/index.html
       20 ./samples/browser-esm-webpack/index.html
       20 ./samples/browser-esm-vite-react/package.json
       20 ./build/npm/removeAll.ts
       19 ./website/tsconfig.json
       19 ./website/src/website/pages/playground/PlaygroundPage.tsx
       19 ./website/src/website/data/playground-samples/creating-the-editor/editor-basic-options/sample.js
       19 ./test/smoke/package-vite.ts
       19 ./src/language/typescript/ts.worker.ts
       19 ./samples/browser-esm-parcel/index.html
       18 ./website/src/website/data/home-samples/sample.tcl.txt
       18 ./samples/electron-amd-nodeIntegration/index.html
       18 ./samples/electron-amd/index.html
       18 ./samples/browser-esm-webpack-typescript-react/tsconfig.json
       18 ./samples/browser-esm-esbuild/index.html
       17 ./website/src/website/data/home-samples/sample.pug.txt
       17 ./website/src/website/data/home-samples/sample.perl.txt
       17 ./website/src/website/data/home-samples/sample.lexon.txt
       17 ./website/src/website/data/home-samples/sample.cameligo.txt
       17 ./.prettierignore
       16 ./website/src/website/pages/MonarchPage.tsx
       16 ./website/src/website/monaco-loader-chunk.ts
       16 ./website/src/website/data/home-samples/sample.liquid.txt
       16 ./.vscode/settings.json
       16 ./samples/browser-esm-webpack-typescript/tsconfig.json
       16 ./samples/browser-esm-webpack-monaco-plugin/index.js
       15 ./website/src/website/data/home-samples/sample.ini.txt
       15 ./.git/hooks/applypatch-msg.sample
       14 ./website/src/website/data/home-samples/sample.yaml.txt
       14 ./website/src/website/data/home-samples/sample.flow9.txt
       14 ./test/manual/index.css
       14 ./src/tsconfig.json
       14 ./src/language/json/json.worker.ts
       14 ./src/language/html/html.worker.ts
       14 ./src/language/css/css.worker.ts
       14 ./.git/hooks/pre-applypatch.sample
       13 ./website/static/monarch/monarch-34px.png
       13 ./website/src/website/data/playground-samples/interacting-with-the-editor/line-and-inline-decorations/sample.css
       13 ./website/src/website/data/playground-samples/customizing-the-appearence/scrollbars/sample.css
       13 ./website/src/website/data/playground-samples/creating-the-editor/hard-wrapping/sample.js
       13 ./website/src/website/data/home-samples/sample.xml.txt
       13 ./website/src/website/data/home-samples/sample.python.txt
       13 ./website/src/website/data/home-samples/sample.mips.txt
       13 ./website/src/types.d.ts
       13 ./.git/hooks/pre-merge-commit.sample
       13 ./editor.code-workspace
       12 ./website/src/website/data/home-samples/sample.twig.txt
       12 ./website/src/website/data/home-samples/sample.bat.txt
       12 ./samples/browser-esm-webpack-typescript-react/src/index.tsx
       12 ./samples/browser-esm-vite-react/index.html
       12 ./samples/browser-esm-esbuild/dist/index.html
       11 ./website/typedoc/typedoc.json
       11 ./website/src/website/data/playground-samples/interacting-with-the-editor/listening-to-key-events/sample.js
       11 ./website/src/website/data/home-samples/sample.proto.txt
       11 ./website/src/website/data/home-samples/sample.powerquery.txt
       11 ./website/src/website/data/home-samples/sample.lua.txt
       11 ./website/src/website/components/Page.tsx
       11 ./webpack-plugin/tsconfig.json
       11 ./test/smoke/parcel/index.html
       11 ./test/smoke/common.js
       11 ./test/manual/samples/run-editor-sample-cr-ps1.txt
       11 ./src/fillers/monaco-editor-core-amd.ts
       11 ./samples/browser-esm-webpack-small/dist/index.html
       11 ./samples/browser-esm-webpack/dist/index.html
       11 ./samples/browser-esm-vite-react/tsconfig.json
       11 ./samples/browser-esm-vite-react/src/main.tsx
       11 ./samples/browser-esm-parcel/src/index.html
       11 ./.git/config
       10 ./website/src/website/index.tsx
       10 ./website/src/website/data/playground-samples/creating-the-editor/hello-world/sample.js
       10 ./webpack-plugin/src/types.ts
       10 ./test/smoke/webpack/index.html
       10 ./test/smoke/vite/index.html
       10 ./test/smoke/esbuild/index.html
       10 ./samples/browser-esm-webpack-typescript-react/src/index.html
       10 ./samples/browser-esm-parcel/package.json
        9 ./website/src/website/data/home-samples/sample.sb.txt
        9 ./website/src/website/data/home-samples/sample.redshift.txt
        9 ./test/manual/samples/run-editor-korean.txt
        9 ./samples/browser-esm-webpack-monaco-plugin/dist/index.html
        9 ./gulpfile.js
        8 ./website/typedoc/tsconfig.json
        8 ./website/src/website/monacoEditorVersion.ts
        8 ./website/src/website/data/playground-samples/creating-the-editor/web-component/sample.html
        8 ./website/src/website/data/home-samples/sample.plaintext.txt
        8 ./website/src/website/data/home-samples/sample.fsharp.txt
        8 ./src/language/typescript/lib/editor.worker.d.ts
        8 ./src/basic-languages/ecl/ecl.test.ts
        8 ./src/basic-languages/csp/csp.test.ts
        8 ./samples/electron-esm-webpack/package.json
        8 ./.prettierrc
        8 ./.gitignore
        8 ./.git/hooks/post-update.sample
        8 ./.devcontainer/devcontainer.json
        8 ./build/tsconfig.json
        7 ./website/src/website/data/playground-samples/interacting-with-the-editor/rendering-glyphs-in-the-margin/sample.css
        7 ./website/src/website/data/home-samples/sample.sparql.txt
        7 ./website/src/website/data/home-samples/sample.pgsql.txt
        7 ./website/src/website/data/home-samples/sample.mysql.txt
        7 ./website/src/website/data/home-samples/sample.ecl.txt
        7 ./website/src/website/data/home-samples/sample.apex.txt
        7 ./samples/electron-amd/package.json
        7 ./samples/electron-amd-nodeIntegration/package.json
        7 ./samples/browser-esm-webpack-typescript/package.json
        7 ./samples/browser-esm-webpack-small/package.json
        7 ./samples/browser-esm-vite-react/vite.config.ts
        7 ./.github/publish-failure-issue-template.md
        6 ./website/tslint.json
        6 ./website/src/website/data/playground-samples/interacting-with-the-editor/listening-to-mouse-events/sample.css
        6 ./website/src/website/bootstrap.scss
        6 ./website/.prettierrc.json
        6 ./test/smoke/parcel/package.json
        6 ./src/fillers/monaco-editor-core.ts
        6 ./scripts/ci/env.ts
        6 ./samples/browser-esm-webpack/package.json
        6 ./samples/browser-esm-webpack-monaco-plugin/package.json
        6 ./samples/browser-esm-esbuild/package.json
        6 ./.git/info/exclude
        5 ./website/src/website/data/playground-samples/interacting-with-the-editor/listening-to-mouse-events/sample.html
        5 ./website/src/website/data/home-samples/sample.m3.txt
        5 ./webpack-plugin/src/loader-utils.d.ts
        5 ./src/language/typescript/lib/typescriptServicesMetadata.ts
        5 ./samples/browser-esm-webpack-typescript/src/index.css
        5 ./samples/browser-esm-webpack-typescript-react/src/index.css
        5 ./samples/browser-esm-webpack-typescript-react/.gitignore
        5 ./samples/browser-esm-webpack-typescript/.gitignore
        5 ./.github/ISSUE_TEMPLATE/config.yml
        4 ./website/src/website/data/playground-samples/creating-the-editor/web-component/sample.json
        4 ./website/src/website/data/playground-samples/creating-the-editor/syntax-highlighting-for-html-elements/sample.js
        4 ./website/src/website/data/playground-samples/creating-the-editor/hello-world/sample.json
        4 ./website/src/website/data/playground-samples/creating-the-editor/chapter.json
        4 ./website/src/website/data/home-samples/sample.pla.txt
        4 ./website/src/website/data/home-samples/sample.azcli.txt
        4 ./test/smoke/webpack/index.js
        4 ./src/fillers/editor.api.d.ts
        4 ./samples/browser-esm-vite-react/src/components/Editor.module.css
        4 ./.mocharc.json
        4 ./.husky/pre-commit
        4 ./.github/ISSUE_TEMPLATE/3_other.md
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/revealing-a-position/sample.json
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/rendering-glyphs-in-the-margin/sample.json
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/listening-to-mouse-events/sample.json
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/listening-to-key-events/sample.json
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/line-and-inline-decorations/sample.json
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/customizing-the-line-numbers/sample.json
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/chapter.json
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/adding-an-action-to-an-editor-instance/sample.json
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/adding-a-keybinding-to-an-existing-command/sample.json
        3 ./website/src/website/data/playground-samples/interacting-with-the-editor/adding-a-command-to-an-editor-instance/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/symbols-provider-example/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/semantic-tokens-provider-example/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/model-markers-example/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/inlay-hints-provider-example/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/hover-provider-example/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/folding-provider-example/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/custom-languages/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/configure-json-defaults/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/configure-javascript-defaults/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/completion-provider-example/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/color-provider-example/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/codelens-provider-example/sample.json
        3 ./website/src/website/data/playground-samples/extending-language-services/chapter.json
        3 ./website/src/website/data/playground-samples/customizing-the-appearence/tokens-and-colors/sample.json
        3 ./website/src/website/data/playground-samples/customizing-the-appearence/scrollbars/sample.json
        3 ./website/src/website/data/playground-samples/customizing-the-appearence/exposed-colors/sample.json
        3 ./website/src/website/data/playground-samples/customizing-the-appearence/chapter.json
        3 ./website/src/website/data/playground-samples/creating-the-editor/syntax-highlighting-for-html-elements/sample.json
        3 ./website/src/website/data/playground-samples/creating-the-editor/hard-wrapping/sample.json
        3 ./website/src/website/data/playground-samples/creating-the-editor/editor-basic-options/sample.json
        3 ./website/src/website/data/playground-samples/creating-the-diffeditor/navigating-a-diff/sample.json
        3 ./website/src/website/data/playground-samples/creating-the-diffeditor/multi-line-example/sample.json
        3 ./website/src/website/data/playground-samples/creating-the-diffeditor/inline-diff-example/sample.json
        3 ./website/src/website/data/playground-samples/creating-the-diffeditor/hello-diff-world/sample.json
        3 ./website/src/website/data/playground-samples/creating-the-diffeditor/chapter.json
        3 ./website/src/website/data/home-samples/sample.redis.txt
        3 ./website/src/website/data/home-samples/sample.cypher.txt
        3 ./webpack-plugin/.npmignore
        3 ./samples/electron-esm-webpack/.gitignore
        3 ./samples/browser-esm-webpack-small/.gitignore
        3 ./samples/browser-esm-webpack-monaco-plugin/.gitignore
        3 ./samples/browser-esm-esbuild/.gitignore
        2 ./samples/browser-esm-webpack/.gitignore
        2 ./samples/browser-esm-vite-react/.gitignore
        2 ./samples/browser-esm-parcel/.gitignore
        1 ./website/src/website/data/playground-samples/interacting-with-the-editor/revealing-a-position/sample.html
        1 ./website/src/website/data/playground-samples/interacting-with-the-editor/rendering-glyphs-in-the-margin/sample.html
        1 ./website/src/website/data/playground-samples/interacting-with-the-editor/listening-to-key-events/sample.html
        1 ./website/src/website/data/playground-samples/interacting-with-the-editor/line-and-inline-decorations/sample.html
        1 ./website/src/website/data/playground-samples/interacting-with-the-editor/customizing-the-line-numbers/sample.html
        1 ./website/src/website/data/playground-samples/interacting-with-the-editor/adding-an-action-to-an-editor-instance/sample.html
        1 ./website/src/website/data/playground-samples/interacting-with-the-editor/adding-a-keybinding-to-an-existing-command/sample.html
        1 ./website/src/website/data/playground-samples/interacting-with-the-editor/adding-a-command-to-an-editor-instance/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/symbols-provider-example/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/semantic-tokens-provider-example/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/model-markers-example/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/inlay-hints-provider-example/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/hover-provider-example/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/folding-provider-example/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/custom-languages/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/configure-json-defaults/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/configure-javascript-defaults/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/completion-provider-example/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/color-provider-example/sample.html
        1 ./website/src/website/data/playground-samples/extending-language-services/codelens-provider-example/sample.html
        1 ./website/src/website/data/playground-samples/customizing-the-appearence/tokens-and-colors/sample.html
        1 ./website/src/website/data/playground-samples/customizing-the-appearence/scrollbars/sample.html
        1 ./website/src/website/data/playground-samples/customizing-the-appearence/exposed-colors/sample.html
        1 ./website/src/website/data/playground-samples/creating-the-editor/hello-world/sample.html
        1 ./website/src/website/data/playground-samples/creating-the-editor/hard-wrapping/sample.html
        1 ./website/src/website/data/playground-samples/creating-the-editor/editor-basic-options/sample.html
        1 ./website/src/website/data/playground-samples/creating-the-diffeditor/navigating-a-diff/sample.html
        1 ./website/src/website/data/playground-samples/creating-the-diffeditor/multi-line-example/sample.html
        1 ./website/src/website/data/playground-samples/creating-the-diffeditor/inline-diff-example/sample.html
        1 ./website/src/website/data/playground-samples/creating-the-diffeditor/hello-diff-world/sample.html
        1 ./samples/browser-esm-vite-react/src/vite-env.d.ts
        1 ./.git/refs/remotes/origin/HEAD
        1 ./.git/refs/heads/main
        1 ./.git/logs/refs/remotes/origin/HEAD
        1 ./.git/logs/refs/heads/main
        1 ./.git/logs/HEAD
        1 ./.git/HEAD
        1 ./.git/description
        1 ./.gitattributes
```
</code></pre>
</div>
</div>

# Boxing this complexity away of the eye and mind

The cognitive load associated with managing a high degree of complexity can adversely affect a programmer's ability to code efficiently and effectively.

In an attempt to mitigate this, I have focused on encapsulating Monaco Editor's complexity within a singular JavaScript file. My goal was to circumvent the need for the nodejs + npm + typescript + webpack build system and the learning curve it implies.

I have made my fork available at https://github.com/arthurweinmann/boxed-monaco-editor. This is a work-in-progress solution, tailored specifically to my needs.

My initial approach was to encapsulate everything into a single file, but this proved difficult. Instead, I settled on a functional encapsulation, which comprises a collection of files that the browser can asynchronously load as required. For instance, the language-specific web workers. This encapsulation is implemented in the 'boxed' folder.

One noteworthy technique involves the method of directing the Monaco Editor to the required JavaScript files when these are not located at the root of your website. I achieved this by assigning an ID to the import script:

```html
<script id="boxedMonacoScript" src="/js/lib/app.boxedmonaco.js"></script>
```

This ID is then used in the webpack entry file to deduce the path of the folder containing the boxed Monaco files:

```javascript
import * as monaco from 'monaco-editor';
// or import * as monaco from 'monaco-editor/esm/vs/editor/editor.api';
// if shipping only a subset of the features & languages is desired

self.MonacoEnvironment = {
	getWorkerUrl: function (moduleId, label) {
		let script = document.getElementById("boxedMonacoScript");
		if (script !== undefined) {
			let src = script.getAttribute('src');
			let spl = src.split("/");
			spl.pop();
			if (spl !== undefined) {
				src = spl.join("/");

				if (src.endsWith('/')) {
					src = src.slice(0, -1);
				}

				if (label === 'json') {
					return src + '/json.worker.boxedmonaco.js';
				}
				if (label === 'css' || label === 'scss' || label === 'less') {
					return src + '/css.worker.boxedmonaco.js';
				}
				if (label === 'html' || label === 'handlebars' || label === 'razor') {
					return src + '/html.worker.boxedmonaco.js';
				}
				if (label === 'typescript' || label === 'javascript') {
					return src + '/ts.worker.boxedmonaco.js';
				}
				return src + '/editor.worker.boxedmonaco.js';
			}
		}

		if (label === 'json') {
			return './json.worker.boxedmonaco.js';
		}
		if (label === 'css' || label === 'scss' || label === 'less') {
			return './css.worker.boxedmonaco.js';
		}
		if (label === 'html' || label === 'handlebars' || label === 'razor') {
			return './html.worker.boxedmonaco.js';
		}
		if (label === 'typescript' || label === 'javascript') {
			return './ts.worker.boxedmonaco.js';
		}
		return './editor.worker.boxedmonaco.js';
	}
};

window.boxedMonaco = monaco;
```

# Conclusion

The prototype of the boxed Monaco Editor is presently in a rudimentary state, yet it adequately fulfills my individual needs at this time.

I believe that our digital systems often possess a level of complexity that surpasses necessity. Software engineering should not merely concern itself with system implementation, but also with the crucial task of minimizing complexity, paring it down to the bare essentials required to accomplish the intended functionality and behavior. 

I feel like, occasionally, what we perceive as simplifications or practical choices may give us the impression of reducing complexity, when in fact we moved it elsewhere, allowing it to grow and grow outside of our immediate awareness.

For a deeper overview of how DNA relates to computing, consider reading this great work: https://berthub.eu/articles/posts/amazing-dna/.