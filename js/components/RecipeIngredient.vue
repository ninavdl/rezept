<template>
  <li>
    <span>{{ amount | round(2) }}</span>
    <span>{{ ingredient.Unit }}</span>
    <span>{{ ingredient.Name }}</span>
    <small>{{ ingredient.Note }}</small>
  </li>
</template>

<script lang="ts">
import 'reflect-metadata';
import { Prop, Component } from 'vue-property-decorator';
import Vue from 'vue';

import Ingredient from '../models/Ingredient';

@Component({})
export default class RecipeIngredientComponent extends Vue {
  @Prop()
  ingredient!: Ingredient;

  @Prop()
  amount!: number;
}
</script>
