<template>
  <li>
    <span>{{ Amount | round(2) }}</span>
    <span>{{ Unit }}</span>
    <span>{{ Name }}</span>
    <small>{{ Note }}</small>
  </li>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  props: ["Name", "Amount", "Unit", "Note"]
});
</script>