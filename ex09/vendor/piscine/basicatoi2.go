/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basicatoi2.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hshimizu <hshimizu@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 21:18:01 by hshimizu          #+#    #+#             */
/*   Updated: 2024/06/25 00:29:12 by hshimizu         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func BasicAtoi2(s string) int {
	var t int
	for _, c := range s {
		if c < '0' || '9' < c {
			return 0
		}
		t = t*10 + int(c) - '0'
	}
	return t
}
