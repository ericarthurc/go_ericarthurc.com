// @ts-nocheck: the hast tree is complex and the types are not available
import { bundledLanguages, createHighlighter } from "shiki";
import { unified } from "unified";
import rehypeStringify from "rehype-stringify";
import remarkParse from "remark-parse";
import remarkRehype from "remark-rehype";
import { visit } from "unist-util-visit";
import { toMdast } from "hast-util-to-mdast";

const netlifyTheme = JSON.parse(
  Deno.readTextFileSync(import.meta.dirname + "/netlify-dark.json")
);

const shiki = await createHighlighter({
  themes: [netlifyTheme],
  langs: [...Object.keys(bundledLanguages)],
});

const FALLBACK_LANGUAGE = "text";

export function rehypeAero() {
  return async (tree: any) => {
    visit(tree, "element", (node, index, parent) => {
      if (node.tagName === "pre" && parent) {
        const codeNode = node.children.find(
          (child: any) => child.tagName === "code"
        );

        // Need to parse the meta data and inject html class attributes on the lines/text
        // const meta = codeNode?.data.meta;

        if (codeNode) {
          // convert the node to Mdast format to pull the syntax language for shiki and the html span
          const codeMdast = toMdast(node);
          // grab the code block text and trim any white space off the end
          const codeBlock = codeNode.children[0].value.trimEnd();

          // style the code block with shiki
          const styledCodeBlock = shiki.codeToHast(codeBlock, {
            lang: codeMdast.lang || FALLBACK_LANGUAGE,
            theme: "Netlify",

            transformers: [
              {
                pre(hast: { children: any }) {
                  return {
                    type: "element",
                    tagName: "div",
                    properties: {
                      className: "code-block",
                    },
                    children: [
                      {
                        type: "element",
                        tagName: "p",
                        properties: {
                          className: "code-block-header",
                        },
                        children: [
                          {
                            type: "element",
                            tagName: "span",
                            properties: {
                              className: "language-name",
                            },
                            children: [{ type: "text", value: codeMdast.lang }],
                          },
                        ],
                      },
                      {
                        type: "element",
                        tagName: "pre",
                        properties: {
                          className: "aero",
                        },
                        children: hast.children,
                      },
                    ],
                  };
                },
                line(hast: { children: any }, line: any) {
                  return {
                    type: "element",
                    tagName: "span",
                    properties: {
                      className: "line",
                    },
                    children: hast.children,
                  };
                },
                span(hast: { properties: any; children: any }, line: any) {
                  return {
                    type: "element",
                    tagName: "span",
                    properties: hast.properties,
                    children: hast.children,
                  };
                },
              },
            ],
          });

          parent.children.splice(index, 1, styledCodeBlock);
        }
      }
    });
  };
}

const data = Deno.args[0];
async function main(data: string) {
  const encoder = new TextEncoder();

  const parsed = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypeAero)
    .use(rehypeStringify)
    .process(data);

  encoder.encode(parsed.toString());
  Deno.stdout.writeSync(encoder.encode(parsed.toString()));
}

await main(data);
