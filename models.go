/*
 * Copyright (c) 2021 Enix, SAS
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 *
 * Authors:
 * Paul Laffitte <paul.laffitte@enix.fr>
 * Arthur Chaloin <arthur.chaloin@enix.fr>
 * Alexandre Buisine <alexandre.buisine@enix.fr>
 */

package dothill

import "strconv"

// model : interface to allow generic conversion from raw response to user-object
type model interface {
	fillFromObject(obj *Object) error
}

// Volume : volume-view representation
type Volume struct {
	LUN int
}

func (m *Volume) fillFromObject(obj *Object) error {
	lun, err := strconv.Atoi(obj.PropertiesMap["lun"].Data)
	if err != nil {
		return err
	}

	m.LUN = lun
	return nil
}
