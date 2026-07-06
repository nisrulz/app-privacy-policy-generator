window.contentMixin = {
  methods: {
    getHtml(id, target) {
      let content = getContent(id);
      let title = getTitle(id);
      let rawHTML = getRawHTML(content, title);
      this.contentRenderType = 2;
      loadInTextView(target, rawHTML);
    },

    getMarkdown(id, target) {
      let content = getContent(id);
      let title = getTitle(id);
      let rawHTML = getRawHTML(content, title);
      let markdown = convertHtmlToMd(rawHTML);
      this.contentRenderType = 2;
      loadInTextView(target, markdown);
    },

    deployFcSimple() {
      fc_deploy_simple();
    },
  },
};
