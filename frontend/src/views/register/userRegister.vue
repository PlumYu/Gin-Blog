<template>
    <div class="register">
      <b-row class="mt-5">
        <b-col md="8" offset-md="2" lg="6" offset-lg="3">
          <b-card title="register">
            <b-form>
              <b-form-group label = "姓名">
                <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="输入您的名称（选填）"
                :state="validateState('name')"
                >
              </b-form-input>
              </b-form-group>
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
                <b-button variant="outline-primary" @click="register()" block>register</b-button>
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

import { mapMutations } from 'vuex';

export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      name: {

      },
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
    ...mapMutations('userModule', ['SET_TOKEN', 'SET_USERINFO']),
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      this.userRegister(this.user).then(() => {
        // 跳转请求成功
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
      console.log('register');
    },
  },
};
</script>

<style lang="scss" scoped>

</style>
