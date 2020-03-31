<template>
  <div :class="loadingClass">
    <SkewLoader :size="loadingSize"></SkewLoader>
    <em>{{ text }}</em>
  </div>
</template>

<script>
import SkewLoader from "vue-spinner/src/SkewLoader.vue";
import Vue from "vue";

export default Vue.extend({
  props: {
    text: String,
    inline: Boolean
  },
  components: {
    SkewLoader
  },
  computed: {
    loadingClass() {
      if (this.inline) return "loading-inline";
      else return "loading";
    },
    loadingSize() {
        if (this.inline) return "0.5em";
        else return "1.5em";
    }
  }
});
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