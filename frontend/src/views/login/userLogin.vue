<template>
    <div class="register">
      <b-row class="mt-5">
        <b-col md="8" offset-md="2" lg="6" offset-lg="3">
          <b-card title="Login">
            <b-form>
              <b-form-group label = "手机号">
                  <b-form-input
                  v-model="$v.user.telephone.$model"
                  type="number"
                  placeholder="输入您的电话号码"
                  :state="validateState('telephone')"
                  ></b-form-input>
                  <b-form-invalid-feedback :state="validateState('telephone')">
                    手机号不符合要求
                  </b-form-invalid-feedback>
              </b-form-group>
              <b-form-group label = "密码">
                  <b-form-input
                  v-model="$v.user.password.$model"
                  type="password"
                  placeholder="输入您的密码"
                  :state="validateState('password')"
                  ></b-form-input>
                  <b-form-invalid-feedback :state="validateState('password')">
                    密码必须大于等于 6 位
                  </b-form-invalid-feedback>
              </b-form-group>
              <b-form-group>
                <b-button variant="outline-primary" @click="login()" block>Login</b-button>
              </b-form-group>
            </b-form>
          </b-card>
        </b-col>
      </b-row>
    </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators';

import customValidator from '@/helper/validator';

import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    ...mapActions('userModule', { userlogin: 'login' }),
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      this.userlogin(this.user).then(() => {
        // 跳转页面
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
    },
  },
};
</script>

<style lang="scss" scoped>

</style>
