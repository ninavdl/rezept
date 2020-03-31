<template>
  <div class="box">
    <p>{{message}}</p>
    <b-upload drag-drop v-on:input="setFile" v-if="thisImage == null">
      <div class="content has-text-centered">
        <p>
          <b-icon icon="upload" size="is-large"></b-icon>
        </p>
        <p>Drop your file here or click to add new image</p>
      </div>
    </b-upload>
    <div class="container">
      <img v-bind:src="thisImage.ThumbnailURL" v-if="thisImage" />
    </div>
    <b-button
      v-if="thisImage"
      v-on:click="removeImage"
      type="is-danger"
      icon-left="delete"
    >Delete image</b-button>
    <progress v-bind:max="fileSize" v-bind:value="uploaded" v-if="uploading"></progress>
  </div>
</template>

<script>
import Vue from "vue";
import { Upload, Icon } from "buefy";
import Image from "../models/Image";

Vue.use(Upload);
Vue.use(Icon);

export default Vue.extend({
  methods: {
    async setFile(file) {
      if (!file.type.match("image/(jpeg|png|webp)")) {
        this.message = "File must be of type jpeg, png or webp";
        return;
      }

      if (file.size > 5 * 1024 * 1024) {
        this.message = "Max. file size: 5MB";
        return;
      }

      this.uploading = true;
      try {
        const uploadedImage = await this.$controller.uploadImage(
          file,
          (uploaded, size) => {
            this.uploaded = uploaded;
            this.fileSize = size;
          }
        );
        this.thisImage = uploadedImage;
        this.$emit("setImage", this.thisImage);
      } catch (e) {
        this.message = e;
      } finally {
        this.uploading = false;
      }
    },
    removeImage() {
      this.thisImage = null;
      this.$emit("setImage", null);
    }
  },
  props: ["image"],
  data: function() {
    return {
      thisImage: this.image,
      message: "",
      uploading: false,
      fileSize: 0,
      uploaded: 0
    };
  }
});
</script>

<style lang="scss" scoped>
img {
  max-width: 10em;
  max-height: 10em;
}
</style>