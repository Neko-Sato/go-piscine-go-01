/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ultimatedivmod.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hshimizu <hshimizu@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 21:18:01 by hshimizu          #+#    #+#             */
/*   Updated: 2024/06/25 00:01:15 by hshimizu         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func UltimateDivMod(a *int, b *int) {
	*a, *b = *a / *b, *a%*b
}