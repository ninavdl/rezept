<script lang="ts">
import "reflect-metadata";
import { Component, Prop } from "vue-property-decorator";
import Vue from "vue";

import Step from "../models/Step";

@Component({})
export default class RecipeStepComponent extends Vue {
  @Prop()
  text!: string;
  @Prop()
  factor!: number;

  tokens: any[];

  created() {
    this.tokens = RecipeStepComponent.tokenize(this.text);
  }

  render(createElement) {
    return createElement(
      "p",
      this.tokens.map(token => {
        if (token.type == "text") return token.value;
        if (token.type == "number")
          return createElement("span", [
            this.$options.filters
              .round(this.factor * parseFloat(token.value), 2)
              .toString()
          ]);
      })
    );
  }

  static tokenize(text) {
    return text.split(/(\[\d+(?:\.\d+)?\])/).map(value => {
      const match = value.match(/\[(\d+(?:\.\d+)?)\]/);
      if (match != null) {
        return {
          type: "number",
          value: match[1]
        };
      } else {
        return {
          type: "text",
          value: value
        };
      }
    });
  }
}
</script>
