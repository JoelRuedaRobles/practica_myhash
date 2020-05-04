package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var claveHex string
	var claveByte []byte
	var clave string

	fmt.Print("Ingresa la contraseña: ")
	fmt.Scanln(&clave)

	/*
	*	Si la cadena hexadecimal es menor
	*	a 8 retorna false
	 */
	if len(clave) < 8 {
		fmt.Println("\033[31mLa contraseña debe tener al menos 8 caracteres")

		return
	}

	claveByte = []byte(clave)
	claveHex = hex.EncodeToString(claveByte)

	//*/ Impresion de datos

	fmt.Println("\nClave: \033[35m", clave, "\033[0m")
	fmt.Println("Clave en bytes: \033[35m", claveByte, "\033[0m")
	fmt.Println("Clave en hexadecimal: \033[35m", claveHex, "\033[0m")

	//*/

	claveHex = filtro(claveHex)

	//*/ Impresion de datos

	fmt.Println("Clave en hexadecimal con filtro: \033[35m", claveHex, "\033[0m")

	matriz, orden := makeMirrorMatrix(claveHex)
	nuevaClave, err := strconv.ParseInt(makeHash(matriz, orden), 16, 64)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Nueva Clave en Bytes: \033[32m%s\033[0m\n", makeHash(matriz, orden))
	fmt.Printf("Nueva Clave en Hexadecimal: \033[32m%d\033[0m\n", nuevaClave)

}

func filtro(hexKey string) string {

	letras := 0               // Cuenta la cantidad de letras que hay en hexKey
	keyByte := []byte(hexKey) // Arreglo que contiene los valores decimales de cada caracter de hexkey
	keyInt := 0               // Valor de la sumatoria total de cada caracter de la cadena hexKey
	factor := 3               // Numero caotico
	dinamicHexKey := hexKey   // Copia de la variable hexKey que cambiara su valor en caso de ser necesario
	minimasLetras := 4        // Minomo de letras permitidos en la cadena hexKey

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
		if letras < minimasLetras || letras == len(hexKey) {

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
		dinamicHexKey += "0"
	}
	//*/

	return dinamicHexKey
}

func makeMirrorMatrix(hexKey string) (mirrorMatrix [][]byte, orden int) {

	//keyByte := []byte(hexKey)
	//var matriz [][]byte // Matriz de slice de bytes
	dim := 1 // Orden de la matriz

	/*
	*	Ciclo que determinara de cuanto sera el orden de la matriz
	 */
	for {
		if (dim * dim) < len(hexKey) {
			dim++
		} else {
			break
		}
	}

	/*
	*	Se aumenta el orden para que halla espacio para
	*	almacenar todos los caracteres de hexKey en la matriz central
	 */
	/*
		if math.Pow(float64(dim), float64(dim)) < float64(len(hexKey)) {
			dim++
		}
	*/
	/*
	*	Dira cuantos caracteres hara falta para completar la matriz
	 */
	faltantes := (dim * dim) - len(hexKey)

	/*
	*	Aqui se llena los espacios vacios
	*	de la matrzi central
	 */

	for {
		if faltantes > 0 {
			//hexKey += hexKey[0 : faltantes-2]
			hexKey += string(hexKey[faltantes-1])
			faltantes--
		} else {
			break
		}
	}

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

	//*/ Impresion de datos

	fmt.Println("Clave en hexadecimal completa: \033[35m", hexKey, "\033[0m")
	fmt.Println("Orden de la matriz central: \033[35m", dim, "\033[0m")

	//*/

	//*/ Impresion de matriz espejo
	fmt.Println("\033[0m", "\n\t\t.: Matriz Espejo en Bytes:.")
	fmt.Println("\033[34m")

	for i := 0; i < dim+2; i++ {
		fmt.Print("\t")
		for j := 0; j < dim+2; j++ {

			//fmt.Printf("%d\t", matriz[i][j])

			if i != 0 && i != dim+1 && j != 0 && j != dim+1 {
				fmt.Printf("\033[1;32m%d\t\033[0m", matriz[i][j])
			} else {
				fmt.Printf("\033[34m%d\t\033[0m", matriz[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println("\033[0m")
	//*/

	//*/ Impresion de matriz espejo
	fmt.Println("\033[0m", "\n\t\t.: Matriz Espejo en Hexadecimal :.")
	fmt.Println("\033[34m")

	for i := 0; i < dim+2; i++ {
		fmt.Print("\t")
		for j := 0; j < dim+2; j++ {

			//fmt.Printf("%s ", string(matriz[i][j]))

			if i != 0 && i != dim+1 && j != 0 && j != dim+1 {
				fmt.Printf("\033[1;32m%s\t\033[0m", string(matriz[i][j]))
			} else {
				fmt.Printf("\033[34m%s\t\033[0m", string(matriz[i][j]))
			}
		}
		fmt.Println()
	}

	fmt.Println("\033[0m")
	//*/

	return matriz, dim
}

func makeHash(mirrorMatrix [][]byte, orderMatrix int) string {

	var nuevaKey []byte
	flag := false // Variable que cambia la orientacion

	for i := 1; i <= orderMatrix; i++ {
		for j := 1; j <= orderMatrix; j++ {

			switch mirrorMatrix[i][j] {

			case 97: // Caracter `a`
				if flag == false {
					nuevaKey = append(nuevaKey, mirrorMatrix[i-1][j])
				} else {
					nuevaKey = append(nuevaKey, mirrorMatrix[i-1][j-1])
				}
			case 98: // Caracter `b`
				if flag == false {
					nuevaKey = append(nuevaKey, mirrorMatrix[i][j+1])
				} else {
					nuevaKey = append(nuevaKey, mirrorMatrix[i-1][j+1])
				}
			case 99: // Caracter `c`
				nuevaKey = append(nuevaKey, 99)
			case 100: // Caracter `d`
				if flag == false {
					nuevaKey = append(nuevaKey, mirrorMatrix[i+1][j])
				} else {
					nuevaKey = append(nuevaKey, mirrorMatrix[i+1][j+1])
				}
			case 101: // Caracter `e`
				if flag == false {
					nuevaKey = append(nuevaKey, mirrorMatrix[i][j-1])
				} else {
					nuevaKey = append(nuevaKey, mirrorMatrix[i+1][j-1])
				}
			case 102: // Caracter `f`
				nuevaKey = append(nuevaKey, 102)
				flag = !flag
			}
		}
	}

	return string(nuevaKey)
}
