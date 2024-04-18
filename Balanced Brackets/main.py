def esta_balanceado(sequencia):
    pilha = []
    mapeamento = {')': '(', '}': '{', ']': '['}
    for caractere in sequencia:
        if caractere in mapeamento.values():
            pilha.append(caractere)
        elif caractere in mapeamento.keys():
            if not pilha or mapeamento[caractere] != pilha.pop():
                return False
    return not pilha

sequencia = "{[{([()()({}[{}])(((())))()({}[])({})])({}[()]){}(({}))]}([{}])({[]}){[()]}()({}{})"
if esta_balanceado(sequencia):
    print("A sequência está balanceada.")
else:
    print("A sequência não está balanceada.")
