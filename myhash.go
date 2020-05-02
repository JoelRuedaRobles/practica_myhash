package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var claveHex string
	var claveByte []byte
	var clave string

	fmt.Println("Ingresa la contraseña")
	fmt.Scanln(&clave)

	/*
	*	Si la cadena hexadecimal es menor
	*	a 8 retorna false
	 */
	if len(clave) < 8 {
		return
	}

	claveByte = []byte(clave)
	claveHex = hex.EncodeToString(claveByte)

	claveHex = filtro(claveHex)
	decodeClaveHex, err := hex.DecodeString(claveHex)

	if err != nil {
		panic(err)
	}

	makeMirrorMatrix(claveHex)

	fmt.Println(claveHex)
	fmt.Println(decodeClaveHex)

}

func filtro(hexKey string) string {

	letras := 0               // Cuenta la cantidad de letras que hay en hexKey
	keyByte := []byte(hexKey) // Arreglo que contiene los valores decimales de cada caracter de hexkey
	keyInt := 0               // Valor de la sumatoria total de cada caracter de la cadena hexKey
	factor := 3               // Numero caotico
	dinamicHexKey := hexKey   // Copia de la variable hexKey que cambiara su valor en caso de ser necesario

	/*
	*	Busca las letras que hay en la
	*	cadena hexadecimal
	 */
	for i := 0; i < 6; i++ {
		letras += strings.Count(hexKey, string(97+i))
	}

	/*
	*	Ciclo que cambiarael valor hexadecimal de la cadena
	*	hasta que tenga las letras suficientes (por lo menos 5)
	 */
	for {
		if letras < 6 || letras == len(hexKey) {

			letras = 0 // Cuenta la cantidad de letras que hay en dinamicHexKey

			/*
			*	Sumatoria total de los valores decimales de cada caracter
			*	en dinamicHexKey
			 */
			for i := 0; i < len(keyByte); i++ {
				keyInt += int(keyByte[i])
			}

			keyInt *= factor
			dinamicHexKey = strconv.FormatInt(int64(keyInt), 16)
			keyByte = []byte(dinamicHexKey)

			/*
			*	Busca las letras que hay en la
			*	cadena hexadecimal
			 */
			for i := 0; i < 6; i++ {
				letras += strings.Count(dinamicHexKey, string(97+i))
			}

		} else {
			break
		}
	}
	//*/
	if len(dinamicHexKey)%2 != 0 {
		dinamicHexKey += "a"
	}
	//*/

	return dinamicHexKey
}

func makeMirrorMatrix(hexKey string) {

	//keyByte := []byte(hexKey)
	//var matriz [][]byte // Matriz de slice de bytes
	dim := 2 // Orden de la matriz

	/*
	*	Ciclo que determinara de cuanto sera el orden de la matriz
	 */
	for {
		if math.Pow(float64(dim), float64(dim)) < float64(len(hexKey)) {
			dim++
		} else {
			break
		}
	}

	/*
	*	Se aumenta el orden para que halla espacio para
	*	almacenar todos los caracteres de hexKey en la matriz central
	 */
	dim++

	/*
	*	Dira cuantos caracteres hara falta para completar la matriz
	 */
	faltantes := ((dim) * (dim)) - len(hexKey)

	/*
	*	Se aumentan 2 unidades a `dim` para la matriz espejo
	*	y una mas para que la matriz central pueda almacenar
	* 	la hexKey
	 */
	var matriz [][]byte

	/*
	* Se inicializa la matriz con 0's para poder trabajar con ella
	* como si fuera una matriz
	 */
	for i := 0; i < dim+2; i++ {
		var slice []byte

		for j := 0; j < dim+2; j++ {
			slice = append(slice, 0)
		}

		matriz = append(matriz, slice)
	}

	/*
	*	Aqui se llena los espacios vacios
	*	de la matrzi central
	 */

	if faltantes > 0 {
		hexKey += hexKey[0:faltantes]
		fmt.Println("Faltantes añadidos", hexKey[0:faltantes])
	}

	/*
	*	Se llena la matriz central con los
	* 	valore de hexKey
	 */
	for i := 1; i < dim+1; i++ {
		for j := 1; j < dim+1; j++ {
			matriz[i][j] = hexKey[(dim)*(i-1)+(j-1)]
		}
	}

	/*
	*	Creacion de la matriz espejo
	 */
	//*/
	for i := 1; i < dim+1; i++ {
		for j := 1; j < dim+1; j++ {

			if i == 1 {
				matriz[dim+1][j] = matriz[i][j]
			}
			if i == dim {
				matriz[0][j] = matriz[i][j]
			}
			if j == 1 {
				matriz[i][dim+1] = matriz[i][j]
			}
			if j == dim {
				matriz[i][0] = matriz[i][j]
			}
		}
	}

	/*
	*	Las esquinas de la matriz se llenan
	 */
	matriz[0][0] = matriz[dim][dim]
	matriz[dim+1][0] = matriz[1][dim]
	matriz[0][dim+1] = matriz[dim][1]
	matriz[dim+1][dim+1] = matriz[1][1]

	//*/ Impresion de matriz espejo
	fmt.Println("\033[33m")

	for i := 0; i < dim+2; i++ {
		for j := 0; j < dim+2; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}

	fmt.Println("\033[0m")
	//*/

	fmt.Println("Faltantes:", faltantes)
	fmt.Println(matriz)

}
