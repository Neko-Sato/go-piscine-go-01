/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sortintegertable.go                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hshimizu <hshimizu@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 21:18:01 by hshimizu          #+#    #+#             */
/*   Updated: 2024/06/25 00:46:50 by hshimizu         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func SortIntegerTable(table []int) {
	len := 0
	for range table {
		len++
	}
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
