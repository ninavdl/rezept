<template>
  <li class="column is-half">
    <div class="box">
      <b-field grouped>
        <b-field>
          <b-input
            v-model="thisIngredient.Amount"
            controls-position="compact"
            required
            pattern="\d+((,|.)\d+)?"
            validation-message="Please enter a number"
            placeholder="Amount"
            v-on:change="$emit('update:ingredient', thisIngredient)"
          />
        </b-field>
        <b-field expanded>
          <b-input
            v-model="thisIngredient.Unit"
            required
            maxlength="10"
            placeholder="Unit"
            expanded
            v-on:change="$emit('update:ingredient', thisIngredient)"
          />
        </b-field>
      </b-field>
      <b-field grouped>
        <b-field expanded>
          <b-input
            v-model="thisIngredient.Name"
            required
            expanded
            placeholder="Name"
            v-on:change="$emit('update:ingredient', thisIngredient)"
          />
        </b-field>
        <p class="control">
          <b-button v-on:click="$emit('remove')" type="is-danger" icon-left="delete">Delete</b-button>
        </p>
      </b-field>
      <b-field>
        <b-field>
          <b-input
            v-model="thisIngredient.Note"
            required
            placeholder="Note"
            v-on:change="$emit('update:ingredient', thisIngredient)"
          />
        </b-field>
      </b-field>
    </div>
  </li>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  props: ["ingredient"],
  data: function() {
    return {
      thisIngredient: this.ingredient
    };
  }
});
</script>

<style lang="scss" scoped>
@import "../../sass/_variables.scss";

.ingredient {
  display: grid;
  align-items: flex-end;
  grid-template-columns: 22% 22% 50% 6%;
  grid-template-rows: auto auto;

  width: 100%;
  padding: 0.75em 0;
  &:not(:first-child) {
    border-top: 1px solid $main-accent;
  }

  label {
    &:not(:last-of-type) {
      grid-column: auto;
      grid-row: 1;
    }
    &:last-of-type {
      grid-column-start: 1;
      grid-column-end: 5;
      grid-row: 2;
      margin-top: 0.25em;
    }
    width: 100%;
    padding-right: 0.75em;
  }

  .controls {
    grid-row: 1;
    grid-column: 4;

    padding: 0 0.75em 0 0;
    text-align: right;
    align-self: flex-end;
  }

  @media (max-width: $min-pagewidth) {
    grid-template-columns: auto auto auto;
    grid-template-rows: auto auto auto;

    label {
      &:nth-child(1) {
        grid-column: 1;
        grid-row: 1;
      }
      &:nth-child(2) {
        grid-column: 2;
        grid-row: 1;
      }
      &:nth-child(3) {
        grid-column: 1 / 3;
        grid-row: 2;
      }
      &:nth-child(4) {
        grid-column: 1 / 3;
        grid-row: 3;
      }
    }

    .controls {
      margin-top: 0.25em;
      grid-row: 4;
      grid-column: 1/4;
      input {
        width: 100%;
      }
    }
  }

  input {
    width: 100%;
  }
}
</style>
