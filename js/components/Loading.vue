<template>
  <div :class="loadingClass">
    <SkewLoader :size="loadingSize"></SkewLoader>
    <em>{{ text }}</em>
  </div>
</template>

<script lang="ts">
import "reflect-metadata";
import { Prop, Component } from "vue-property-decorator";
import Vue from "vue";

import SkewLoader from "vue-spinner/src/SkewLoader.vue";

@Component({
  components: { SkewLoader }
})
export default class LoadingComponent extends Vue {
  @Prop() text!: String;
  @Prop() inline!: boolean;

  get loadingClass(): String {
    return this.inline ? "loading-inline" : "loading";
  }

  get loadingSize(): String {
    return this.inline ? "0.5em" : "1.5em";
  }
}
</script>

<style lang="scss">
/* Not using a scoped style to override the spinners color with css */

@import "../../sass/_variables.scss";

.v-skew {
  border-bottom-color: $main-accent !important;
}

.loading,
.loading-inline {
  cursor: wait;
  color: $main-accent;
  user-select: none;
}

.loading {
  text-align: center;

  em {
    display: inline-block;
    margin-top: 1em;
  }
}

.loading-inline {
  display: inline-block;
  margin: 0.25em;
  .v-spinner {
    display: inline-block;
  }
}
</style>