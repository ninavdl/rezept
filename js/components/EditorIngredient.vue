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
          <b-button v-on:click="$emit('remove')" type="is-danger" icon-left="delete">
            Delete
          </b-button>
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

<script lang="ts">
import 'reflect-metadata';
import { Component, Prop } from 'vue-property-decorator';
import Vue from 'vue';

import Ingredient from '../models/Ingredient';

@Component({})
export default class EditorIngredientComponent extends Vue {
  @Prop()
  ingredient!: Ingredient;

  thisIngredient: Ingredient = this.ingredient;
}

</script>
