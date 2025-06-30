<template>
  <div class="p-6 max-w-2xl mx-auto">
    <h1 class="text-xl font-bold mb-4">Nova Loja</h1>

    <form @submit.prevent="criarLoja" class="space-y-4">
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block mb-1 font-semibold">Número da Loja</label>
          <input v-model="form.numero_loja" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Nome</label>
          <input v-model="form.nome" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Razão Social</label>
          <input v-model="form.razao_social" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Endereço</label>
          <input v-model="form.endereco" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Número</label>
          <input v-model="form.numero" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Cidade</label>
          <input v-model="form.cidade" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Estado</label>
          <input v-model="form.estado" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">CEP</label>
          <input v-model="form.cep" class="w-full border p-2 rounded" required />
        </div>

        <div class="col-span-2">
          <label class="block mb-1 font-semibold">Estabelecimento</label>
          <select v-model="form.estabelecimento_id" class="w-full border p-2 rounded" required>
            <option value="" disabled>Selecione um estabelecimento</option>
            <option v-for="estabelecimento in estabelecimentos" :key="estabelecimento.id" :value="estabelecimento.id">
              {{ estabelecimento.nome }}
            </option>
          </select>
        </div>
      </div>

      <button
        type="submit"
        class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700"
      >
        Salvar
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useNuxtApp } from '#app'

const router = useRouter()
const { $api } = useNuxtApp()

interface Loja {
  numero_loja: string
  nome: string
  razao_social: string
  endereco: string
  cidade: string
  estado: string
  cep: string
  numero: string
  estabelecimento_id: number | null
}

const form = ref<Loja>({
  numero_loja: '',
  nome: '',
  razao_social: '',
  endereco: '',
  cidade: '',
  estado: '',
  cep: '',
  numero: '',
  estabelecimento_id: null,
})

const estabelecimentos = ref<{ id: number; nome: string }[]>([])

const carregarEstabelecimentos = async () => {
  try {
    const response = await $api<{ data: any[] }>('/estabelecimentos')
    estabelecimentos.value = response.data
  } catch (error) {
    console.error('Erro ao carregar estabelecimentos', error)
  }
}

const criarLoja = async () => {
  try {
    await $api('/lojas', {
      method: 'POST',
      body: form.value,
    })
    alert('Loja criada com sucesso!')
    router.push('/')
  } catch (error) {
    alert('Erro ao criar loja')
    console.error(error)
  }
}

onMounted(carregarEstabelecimentos)
</script>
